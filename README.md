# A simple Go client for Kubernetes

[![GoDoc](https://godoc.org/github.com/ericchiang/k8s?status.svg)](https://godoc.org/github.com/ericchiang/k8s)

A slimmed down Go client generated using Kubernetes' new [protocol buffer][protobuf] support. This package behaves similarly to [official Kubernetes' Go client][client-go], but only imports two external dependencies.

```go
import (
    "context"
    "fmt"
    "log"

    "github.com/ericchiang/k8s"
)

func main() {
    client, err := k8s.NewInClusterClient()
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
* [github.com/golang/protobuf/proto][go-proto] (protobuf serialization)
* [golang.org/x/net/http2][go-http2] (HTTP/2 support)

## Versioned supported

This client supports every API group version present since 1.3.

## Usage

### Namespaces

Clients are initialized with a default namespace. For in-cluster clients, this is the namespace the pod was deployed in.

```go
pods, err := client.ListPods(ctx, "") // Pods in the current namespace.
```

This can be overridden by explicitly passing a namespace.

```go
pods, err := client.ListPods(ctx, "custom-namespace") // Pods from the "custom-namespace"
```

### Label selectors

Label selectors can be provided to any list operation.

```go
l := new(k8s.LabelSelector)
l.Eq("tier", "production")
l.In("app", "database", "frontend")

pods, err := client.CoreV1().ListPods(ctx, "", l.Selector())
```

### Working with resources

Use the generated API types directly to create and modify resources.

```go
import (
    "context"

    "github.com/ericchiang/k8s"
    "github.com/ericchiang/k8s/api/v1"
)

func createConfigMap(client *k8s.Client, name string, values map[string]string) error {
    cm := &v1.ConfigMap{
        Metadata: &v1.ObjectMeta{
            Name:      &name,
        },
        Data: values,
    }
    // Will return the created configmap as well.
    _, err := client.CoreV1().Create(context.TODO(), cm, client.Namespace)
    return err
}
```

API structs use pointers to `int`, `bool`, and `string` types to differentiate between the zero value and an unsupplied one. This package provides [convenience methods][string] for creating pointers to literals of basic types.

### Creating out-of-cluster clients

Out-of-cluster clients can be constructed by either creating an `http.Client` manually or parsing a [`Config`][config] object. The following is an example of creating a client from a kubeconfig:

```go
import (
    "io/ioutil"

    "github.com/ericchiang/k8s"
    "github.com/ghodss/yaml"
)

func loadClient(kubeconfigPath string) (*k8s.Client, error) {
    data, err := ioutil.ReadFile(kubeconfigPath)
    if err != nil {
        return nil, fmt.Errorf("read kubeconfig: %v", err)
    }

    var config k8s.Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
    }
    return k8s.NewClient(&config)
}
```

### Errors

Errors returned by the Kubernetes API are formatted as [`unversioned.Status`][unversioned-status] objects and surfaced by clients as [`*k8s.APIError`][k8s-error]s. Programs that need to inspect error codes or failure details can use a type cast to access this information.

```go
configMap := &v1.ConfigMap{
    Metadata: &v1.ObjectMeta{
        Name:      "test",
        Namespace: "default",
    },
    Data: map[string]string{"foo": "bar"},
}

_, err := client.CoreV1().CreateConfigMap(ctx, configMap)
if err != nil {
    if apiErr, ok := err.(*k8s.APIError); ok {
        // Resource already exists. Carry on.
        if apiErr.Status.Code == http.StatusConflict {
            return nil
        }
    }
    return fmt.Errorf("create configmap: %v", err)
}
return nil
```

[client-go]: https://github.com/kubernetes/client-go
[go-proto]: https://godoc.org/github.com/golang/protobuf/proto
[go-http2]: https://godoc.org/golang.org/x/net/http2
[protobuf]: https://developers.google.com/protocol-buffers/
[unversioned-status]: https://godoc.org/github.com/ericchiang/k8s/api/unversioned#Status
[k8s-error]: https://godoc.org/github.com/ericchiang/k8s#APIError
[config]: https://godoc.org/github.com/ericchiang/k8s#Config
[string]: https://godoc.org/github.com/ericchiang/k8s#String
