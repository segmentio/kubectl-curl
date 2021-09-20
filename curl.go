package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/segmentio/kubectl-curl/curl"
	"github.com/spf13/pflag"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

var (
	curlOptions = curl.NewOptionSet()

	help   bool
	debug  bool
	usage  string
	flags  *pflag.FlagSet
	config *genericclioptions.ConfigFlags
)

func init() {
	runtime.ErrorHandlers = nil // disables default kubernetes error logging
	rand.Seed(time.Now().UnixNano())

	log.SetOutput(os.Stderr)
	log.SetPrefix("* ")

	flags = pflag.NewFlagSet("kubectl curl", pflag.ExitOnError)
	flags.BoolVarP(&help, "help", "h", false, "Prints the kubectl plugin help.")
	flags.BoolVarP(&debug, "debug", "", false, "Enable debug mode to print more details about the kubectl command execution.")

	for _, opt := range curlOptions {
		name := strings.TrimPrefix(opt.Name, "--")
		short := strings.TrimPrefix(opt.Short, "-")

		// Rewrite option names that conflict with the kubectl default options:
		switch name {
		case "user": // * Change curl's "--user" option to "--userinfo"
			name = "userinfo"
		}

		switch short {
		case "n", "s":
			// Remove short names that conflict with the kubectl default options:
			// * "-n" conflicts between kubectl's "--namespace" and curl's "--netrc"
			// * "-s" conflicts between kubectl's "--server" and curl's "--silent"
			short = ""
		}

		flags.VarP(opt.Value, name, short, opt.Help)
	}

	config = genericclioptions.NewConfigFlags(false)
	config.AddFlags(flags)
	usage = flags.FlagUsages()
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "* ERROR: %s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	flags.Parse(os.Args[1:])

	if help {
		printUsage()
		return nil
	}

	var stdout io.Writer
	var stderr io.Writer
	if debug {
		stdout = os.Stdout
		stderr = os.Stderr
	} else {
		log.SetOutput(io.Discard)
	}

	var args = flags.Args()
	var query string
	var containerName string
	switch len(args) {
	case 2:
		query, containerName = args[0], args[1]
	case 1:
		query = args[0]
	default:
		return fmt.Errorf("too many arguments passed in the command line invocation of kubectl curl [URL] [container]")
	}

	requestURL, err := url.Parse(query)
	if err != nil {
		return fmt.Errorf("malformed URL: %w", err)
	}
	switch requestURL.Scheme {
	case "http", "https":
	case "":
		return fmt.Errorf("missing scheme in query URL: %s", query)
	default:
		return fmt.Errorf("unsupposed scheme in query URL: %s", query)
	}

	podName, podPort, err := net.SplitHostPort(requestURL.Host)
	if err != nil {
		podName = requestURL.Host
		podPort = ""
	}

	kubeConfig := config.ToRawKubeConfigLoader()
	namespace, _, err := kubeConfig.Namespace()
	if err != nil {
		return err
	}
	restConfig, err := config.ToRESTConfig()
	if err != nil {
		return err
	}
	client, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	log.Printf("kubectl get -n %s pod/%s", namespace, podName)
	pod, err := client.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if pod.Status.Phase != corev1.PodRunning {
		return fmt.Errorf("unable to forward port because pod is not running. Current status=%v", pod.Status.Phase)
	}

	const minPort = 10200
	const maxPort = 16383
	localPort := rand.Int31n(maxPort-minPort) + minPort
	remotePort := int32(0)
	portName := requestURL.Scheme

	if podPort != "" {
		p, err := strconv.ParseInt(podPort, 10, 32)
		if err != nil {
			portName = podPort
		} else {
			remotePort = int32(p)
		}
	}

	if remotePort == 0 {
		selectedContainerName, selectedContainerPort, err := selectContainerPort(pod, containerName, portName)
		if err != nil {
			return err
		}
		containerName = selectedContainerName
		remotePort = selectedContainerPort.ContainerPort
	}

	log.Printf("forwarding local port %d to port %d of %s", localPort, remotePort, containerName)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	f, err := openPortForwarder(ctx, portForwarderConfig{
		config:     restConfig,
		namespace:  namespace,
		podName:    podName,
		localPort:  localPort,
		remotePort: remotePort,
		stdout:     stdout,
		stderr:     stderr,
	})
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()
	defer log.Printf("waiting for port forwarder to stop")

	defer cancel()
	defer log.Printf("shutting down port forwarder")

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer f.Close()

		if err := f.ForwardPorts(); err != nil {
			log.Print(err)
		}
	}()

	log.Printf("waiting for port fowarding to be established")
	select {
	case <-f.Ready:
	case <-ctx.Done():
		return nil
	}

	requestURL.Host = net.JoinHostPort("localhost", strconv.Itoa(int(localPort)))
	options := append(curlOptions,
		// The -s option is taken by -s,--server from the default kubectl
		// configuration. Force --silent because we don't really need to
		// print the dynamic progress view for the scenarios in which this
		// plugin is useful for.
		curl.Silent(true),
	)
	cmd := curl.Command(ctx, requestURL.String(), options...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("curl %s", strings.Join(cmd.Args[1:], "\n\t"))
	return cmd.Run()
}

func selectContainerPort(pod *corev1.Pod, containerName, portName string) (selectedContainerName string, selectedContainerPort corev1.ContainerPort, err error) {
	for _, container := range pod.Spec.Containers {
		if containerName != "" && container.Name != containerName {
			continue
		}
		for _, port := range container.Ports {
			if port.Name != portName || port.Protocol != corev1.ProtocolTCP {
				continue
			}
			if selectedContainerPort.Name != "" {
				err = fmt.Errorf("pod %[1]s has multiple containers with a %[2]s port, use kubectl %[1]s [container] to specify which one to profile", pod.Name, portName)
				return
			}
			selectedContainerName = container.Name
			selectedContainerPort = port
		}
	}
	if selectedContainerPort.Name == "" {
		err = fmt.Errorf("pod %s had no containers exposing a %s port", pod.Name, portName)
	}
	return
}

type portForwarderConfig struct {
	config     *rest.Config
	namespace  string
	podName    string
	localPort  int32
	remotePort int32
	stdout     io.Writer
	stderr     io.Writer
}

func openPortForwarder(ctx context.Context, fwd portForwarderConfig) (*portforward.PortForwarder, error) {
	transport, upgrader, err := spdy.RoundTripperFor(fwd.config)
	if err != nil {
		return nil, err
	}

	host := strings.TrimLeft(fwd.config.Host, "htps:/")
	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", fwd.namespace, fwd.podName)

	client := &http.Client{
		Transport: transport,
	}

	dialer := spdy.NewDialer(upgrader, client, http.MethodPost, &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	})

	ports := []string{
		fmt.Sprintf("%d:%d", fwd.localPort, fwd.remotePort),
	}

	if fwd.stdout == nil {
		fwd.stdout = io.Discard
	}

	if fwd.stderr == nil {
		fwd.stderr = io.Discard
	}

	return portforward.New(dialer, ports, ctx.Done(), make(chan struct{}), fwd.stdout, fwd.stderr)
}

func printUsage() {
	fmt.Printf(`Run curl against kubernetes pods

Usage:
  kubectl curl [options] URL [container]

Options:
%s
`, usage)
}
