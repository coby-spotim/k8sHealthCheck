# K8sHealthCheck

A small and simple health check for Kubernetes Applications that can be used on any non-HTTP server app as well.

One of the biggest issues that I've had with using Kubernetes is the fact that the typical way to tell K8S that your pod is alive is by raising a server with an endpoint that returns a 200 OK response.

To me, this is an annoying overhead that I despise, since I write non-HTTP server applications as well, so I wrote a simple library out of the last health check I wrote for one of these so that I can easily use and run a small HTTP endpoint without copying and pasting my code everywhere.