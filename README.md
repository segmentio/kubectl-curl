# kubectl-curl

Kubectl plugin to run curl commands against kubernetes pods

## Motivation

Sending http requests to kubernetes pods is unnecessarily complicated, this
plugin makes it easy.

The plugin creates a port forwarding from the local network to the kubernetes
pod that was selected to receive the request, then instantiate a curl command
by rewriting the URL to connect to the forwarded local port, and passing all
curl options that were given on the command line.

## Installation

If `$GOPATH/bin` is in the `PATH`, the plugin can be installed with:
```
$ go install github.com/segmentio/kubectl-curl@latest
```

If it was installed properly, it will be visibile when listing kubectl plugins:
```
$ kubectl plugin list
The following compatible plugins are available:

/.../kubectl-curl
```

## Usage

```
kubectl curl [options] URL [container]
```

* In the URL, the host part must be the name of the pod to send the request to.
* If no port number is specified, the request will be sent to a `http` port.
* If there are multiple containers with a `http` port, the name of the container
  to send to the request to must be specified after the URL.

## Examples

This section records common use cases for this kubectl plugin.

### Collecting profiles of Go programs

```
$ kubectl curl -o profile "http://{pod}/debug/pprof/profile?debug=1&seconds=10"
$ go tool pprof -http :6060 ./profile
```

* Full documentation: [net/http/pprof](https://pkg.go.dev/net/http/pprof)

### Retrieving prometheus metrics

```
$ kubectl curl http://{pod}/metrics
...
```
