# A Simple Go Client for Kubernetes

[![GoDoc](https://godoc.org/github.com/ericchiang/k8s?status.svg)](https://godoc.org/github.com/ericchiang/k8s)

This package holds a slimmed down Kubernetes client. It imports a [single external dependency][gogo-proto], compiles much faster than either of the offical clients, and won't bloat `vendor` directories.

The client uses Kubernetes' new support for [protobuf][protobuf] serialization. Are types are generated from canonical `.proto` files in the Kubernetes repo, and this package understands the custom wire format used to talk to the API server. However, the package API looks similar to the official client:

```
import (
    "context"
    "fmt"
    "log"

    "github.com/ericchiang/k8s"
)

func main() {
    client, err := k8s.InClusterClient()
    if err != nil {
        log.Fatal(err)
    }

    nodes, err := client.CoreV1().ListNodes(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    for _, node := range nodes.Items {
        fmt.Printf("name=%q schedulable=%t\n", node.Metadata.Name, !node.Spec.Unschedulable)
    }
}
```

## Requirements

* Go 1.7+ (this package uses "context" features added in 1.7)
* Kubernetes 1.3+ (protobuf support was added in 1.3)

## Configurability

Clients are initialized with a default namespace. For in-cluster clients, this is the namespace the pod was deployed in.

```
pods, err := client.ListPods(ctx) // Pods in the current namespace.
```

Clients that wish to query a different namespace can do so using a context key.

```
ctxWithNamespace = k8s.NamespaceContext(ctx, "custom-namespace")
pods, err := client.ListPods(ctxWithNamespace) // Pods from the "custom-namespace"
```

Out-of-cluster clients can be constructed by creating a `Client` manually. The following is an example of creating a client which uses TLS client auth:

```
// Load client cert.
clientCert, err := tls.LoadX509KeyPair("client.crt", "client.key")
if err != nil {
    // handle error
}

// Load API server's CA.
caData, err := ioutil.ReadFile("ca.crt")
if err != nil {
    // handle error
}
rootCAs := x509.NewCertPool()
if !rootCAs.AppendCertsFromPEM(caData) {
    // handle error
}

// Create a client with a custom TLS config.
client := &k8s.Client{
    Endpoint:  "https://node1.example.com:443",
    Namespace: "kube-system",
    Client: &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                RootCAs:      rootCAs,
                Certificates: []tls.Certificate{clientCert},
            },
        },
    },
}
```

## Errors

Errors returned by the Kubernetes API are formatted as [`unversioned.Status`][unversioned-status] objects and surfaced by clients as [`*k8s.Error`][k8s-error]s. Programs that need to inspect error codes or failure details can use a type cast to access this information.

```
configMap := &v1.ConfigMap{
    Metadata: &v1.ObjectMeta{
        Name:      "test",
        Namespace: "default",
    },
    Data: map[string]string{"foo": "bar"},
}

_, err := client.CoreV1().CreateConfigMap(ctx, configMap)
if err != nil {
    if k8sErr, ok := err.(*k8s.Error); ok {
        // Resource already exists. Carry on.
        if k8sErr.Status.Code == http.StatusConflict {
            return nil
        }
    }
    return fmt.Errorf("create configmap: %v", err)
}
return nil
```

[gogo-proto]: https://godoc.org/github.com/gogo/protobuf/proto
[protobuf]: https://developers.google.com/protocol-buffers/
[unversioned-status]: https://godoc.org/github.com/ericchiang/k8s/api/unversioned#Status
[k8s-error]: https://godoc.org/github.com/ericchiang/k8s#Error
