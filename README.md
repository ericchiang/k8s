# A simple Go client for Kubernetes

[![GoDoc](https://godoc.org/github.com/ericchiang/k8s?status.svg)](https://godoc.org/github.com/ericchiang/k8s)

This package holds a slimmed down Kubernetes client. It imports a [single external dependency][go-proto], compiles much faster than either of the offical clients, and won't bloat `vendor` directories.

The client uses Kubernetes' new support for [protobuf][protobuf] serialization. Are types are generated from canonical `.proto` files in the Kubernetes repo, and this package understands the custom wire format used to talk to the API server. However, the package API looks similar to the official client:

```go
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

## Project status

__DO NOT USE THIS CODE.__

This package is still in development and may change in unexpected ways. It's an experiment to demonstrate generating a client is possible, though it may evolve into a more mature library.

Until this package becomes more mature use Kubernetes' Go client instead: https://github.com/kubernetes/client-go

## Requirements

* Go 1.7+ (this package uses "context" features added in 1.7)
* Kubernetes 1.3+ (protobuf support was added in 1.3)

## Versioned clients?

This client grabs `.proto` files from multiple versions of the Kubernetes codebase. This means that it supports every API group version present in Kubernetes since 1.3.

This client currently doesn't support the discovery API for determining what version of the API your client is talking to. Progress for that feature can be found [here](https://github.com/ericchiang/k8s/issues/3).

## Configuration

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

### Creating resources

Use the generated API types directly to create resources.

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
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "os"

    "github.com/ericchiang/k8s"
    "github.com/ghodss/yaml"
)

func loadClient() (*k8s.Client, error) {
    data, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".kube/config"))
    if err != nil {
        return nil, fmt.Errorf("load kubeconfig: %v", err)
    }

    // Create a new config and parse it.
    config := new(k8s.Config)
    if err := yaml.Unmarshal(data, config); err != nil {
        return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
    }

    // Create client from config.
    return k8s.NewClient(config)
}

func main() {
    client, err := loadClient()
    if err != nil {
        fmt.Fprintf(os.Stderr, "load client: %v\n", err)
        os.Exit(2)
    }

    nodes, err := client.CoreV1().ListNodes(context.Background())
    if err != nil {
        fmt.Fprintf(os.Stderr, "list nodes: %v\n", err)
        os.Exit(2)
    }

    for _, node := range nodes {
        fmt.Println(node.Name)
    }
}
```

## Errors

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

[go-proto]: https://godoc.org/github.com/golang/protobuf/proto
[protobuf]: https://developers.google.com/protocol-buffers/
[unversioned-status]: https://godoc.org/github.com/ericchiang/k8s/api/unversioned#Status
[k8s-error]: https://godoc.org/github.com/ericchiang/k8s#APIError
[config]: https://godoc.org/github.com/ericchiang/k8s#Config
[string]: https://godoc.org/github.com/ericchiang/k8s#String
