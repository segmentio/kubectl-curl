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

* In the `URL`, the host part can be:
    * **podName**: pod to send the request to
    * a resource reference, such as **deployment/deploymentName**. The request is sent to a random pod from this resource.
       * NOTE: supported resources: **deployment**, **statefulset**, **daemonset**
       * NOTE: supported abbreviations: **deploy**, **sts**, **ds**
* If no port number is specified, the request will be sent to an `http` port.
* If there are multiple containers with an `http` port, the name of the container
  to send to the request to must be specified after the URL.

## Examples

This section records common use cases for this kubectl plugin.

### Collecting profiles of Go programs

```
$ kubectl curl "http://{podname}/debug/pprof/profile?debug=1&seconds=10" > ./profile
$ go tool pprof -http :6060 ./profile
```

* Full documentation: [net/http/pprof](https://pkg.go.dev/net/http/pprof)

### Retrieving prometheus metrics

All of these variants work:

```
$ kubectl curl http://{podname}/metrics
...

$ kubectl curl http://daemonset/{daemonsetname}/metrics
...

$ kubectl curl http://ds/{daemonsetname}/metrics
...

$ kubectl curl {podname}/metrics
...

$ kubectl curl daemonset/{daemonsetname}/metrics
...

$ kubectl curl ds/{daemonsetname}/metrics
```
