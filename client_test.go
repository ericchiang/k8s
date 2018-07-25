package k8s_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

const skipMsg = `
warning: this package's test run using the default context of your "kubectl" command,
and will create resources on your cluster (mostly configmaps).

If you wish to continue set the following environment variable:

	export K8S_CLIENT_TEST=1

To suppress this message, set:

	export K8S_CLIENT_TEST=0
`

func newTestClient(t *testing.T) *k8s.Client {
	if os.Getenv("K8S_CLIENT_TEST") == "0" {
		t.Skip("")
	}
	if os.Getenv("K8S_CLIENT_TEST") != "1" {
		t.Skip(skipMsg)
	}

	cmd := exec.Command("kubectl", "config", "view", "-o", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("'kubectl config view -o json': %v %s", err, out)
	}

	config := new(k8s.Config)
	if err := json.Unmarshal(out, config); err != nil {
		t.Fatalf("parse kubeconfig: %v '%s'", err, out)
	}
	client, err := k8s.NewClient(config)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	return client
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func withNamespace(t *testing.T, test func(client *k8s.Client, namespace string)) {
	client := newTestClient(t)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	name := "k8s-test-" + string(b)
	namespace := corev1.Namespace{
		Metadata: &metav1.ObjectMeta{
			Name: &name,
		},
	}
	if err := client.Create(context.TODO(), &namespace); err != nil {
		t.Fatalf("create namespace: %v", err)
	}
	defer func() {
		if err := client.Delete(context.TODO(), &namespace); err != nil {
			t.Fatalf("delete namespace: %v", err)
		}
	}()

	test(client, name)
}

func TestNewTestClient(t *testing.T) {
	newTestClient(t)
}

func TestListNodes(t *testing.T) {
	client := newTestClient(t)
	var nodes corev1.NodeList
	if err := client.List(context.TODO(), "", &nodes); err != nil {
		t.Fatal(err)
	}
	if len(nodes.Items) == 0 {
		t.Skip("no nodes in cluster")
	}

	for _, node := range nodes.Items {
		if node.Metadata.Annotations == nil {
			node.Metadata.Annotations = map[string]string{}
		}
		node.Metadata.Annotations["foo"] = "bar"
		if err := client.Update(context.TODO(), node); err != nil {
			t.Fatal(err)
		}
		delete(node.Metadata.Annotations, "foo")
		if err := client.Update(context.TODO(), node); err != nil {
			t.Fatal(err)
		}
	}
}

func TestWithNamespace(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {})
}

func TestUpdateNamespaceStatus(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		var ns corev1.Namespace
		if err := client.Get(context.TODO(), "", namespace, &ns); err != nil {
			t.Errorf("get namespace: %v", err)
			return
		}

		if err := client.Update(context.TODO(), &ns, k8s.Subresource("status")); err != nil {
			t.Errorf("update namespace status subresource: %v", err)
		}

		if err := client.Update(context.TODO(), &ns, k8s.Subresource("idontexist")); err == nil {
			t.Errorf("updated invalid subresource")
		}
	})
}

func TestCreateConfigMap(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		cm := &corev1.ConfigMap{
			Metadata: &metav1.ObjectMeta{
				Name:      k8s.String("my-configmap"),
				Namespace: &namespace,
			},
		}
		if err := client.Create(context.TODO(), cm); err != nil {
			t.Errorf("create configmap: %v", err)
			return
		}
		got := new(corev1.ConfigMap)
		if err := client.Get(context.TODO(), namespace, *cm.Metadata.Name, got); err != nil {
			t.Errorf("get configmap: %v", err)
			return
		}
		if !reflect.DeepEqual(cm, got) {
			t.Errorf("expected configmap %#v, got=%#v", cm, got)
		}

		if err := client.Delete(context.TODO(), cm); err != nil {
			t.Errorf("delete configmap: %v", err)
			return
		}
	})
}

func TestListConfigMap(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		for i := 0; i < 5; i++ {
			cm := &corev1.ConfigMap{
				Metadata: &metav1.ObjectMeta{
					Name:      k8s.String(fmt.Sprintf("my-configmap-%d", i)),
					Namespace: &namespace,
				},
			}
			if err := client.Create(context.TODO(), cm); err != nil {
				t.Errorf("create configmap: %v", err)
				return
			}
		}

		var configMapList corev1.ConfigMapList
		if err := client.List(context.TODO(), namespace, &configMapList); err != nil {
			t.Errorf("list configmaps: %v", err)
			return
		}

		if n := len(configMapList.Items); n != 5 {
			t.Errorf("expected 5 configmaps, got %d", n)
		}
	})
}

func TestListConfigMapDefaultNamespace(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		client.Namespace = namespace
		for i := 0; i < 5; i++ {
			cm := &corev1.ConfigMap{
				Metadata: &metav1.ObjectMeta{
					Name: k8s.String(fmt.Sprintf("my-configmap-%d", i)),
				},
			}
			if err := client.Create(context.TODO(), cm); err != nil {
				t.Errorf("create configmap: %v", err)
				return
			}
		}

		var configMapList corev1.ConfigMapList
		if err := client.List(context.TODO(), "", &configMapList); err != nil {
			t.Errorf("list configmaps: %v", err)
			return
		}

		if n := len(configMapList.Items); n != 5 {
			t.Errorf("expected 5 configmaps, got %d", n)
		}
	})
}

func TestCreateConfigMapDefaultNamespace(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		client.Namespace = namespace
		cm := &corev1.ConfigMap{
			Metadata: &metav1.ObjectMeta{
				Name: k8s.String("my-configmap"),
			},
		}
		if err := client.Create(context.TODO(), cm); err != nil {
			t.Errorf("create configmap: %v", err)
			return
		}
		got := new(corev1.ConfigMap)
		if err := client.Get(context.TODO(), "", *cm.Metadata.Name, got); err != nil {
			t.Errorf("get configmap: %v", err)
			return
		}
		if !reflect.DeepEqual(cm, got) {
			t.Errorf("expected configmap %#v, got=%#v", cm, got)
		}

		got.Metadata.Namespace = nil
		if err := client.Delete(context.TODO(), cm); err != nil {
			t.Errorf("delete configmap: %v", err)
			return
		}
	})
}

func TestDefaultNamespace(t *testing.T) {
	c := &k8s.Config{
		Clusters: []k8s.NamedCluster{
			{
				Name: "local",
				Cluster: k8s.Cluster{
					Server: "http://localhost:8080",
				},
			},
		},
		AuthInfos: []k8s.NamedAuthInfo{
			{
				Name: "local",
			},
		},
	}
	cli, err := k8s.NewClient(c)
	if err != nil {
		t.Fatal(err)
	}
	if cli.Namespace != "default" {
		t.Errorf("expected namespace=%q got=%q", "default", cli.Namespace)
	}
}

func TestDelete(t *testing.T) {
    t.Skip("this test requires a controller manager to run, see issue #95")

	withNamespace(t, func(client *k8s.Client, namespace string) {
		cm := &corev1.ConfigMap{
			Metadata: &metav1.ObjectMeta{
				Name:      k8s.String("my-configmap"),
				Namespace: &namespace,
			},
		}
		if err := client.Create(context.TODO(), cm); err != nil {
			t.Errorf("create configmap: %v", err)
			return
		}

		cm2 := &corev1.ConfigMap{
			Metadata: &metav1.ObjectMeta{
				Name:      k8s.String("my-configmap-2"),
				Namespace: &namespace,
				OwnerReferences: []*metav1.OwnerReference{
					{
						ApiVersion:         k8s.String("v1"),
						Kind:               k8s.String("ConfigMap"),
						Name:               cm.Metadata.Name,
						Uid:                cm.Metadata.Uid,
						BlockOwnerDeletion: k8s.Bool(true),
					},
				},
			},
		}
		if err := client.Create(context.TODO(), cm2); err != nil {
			t.Errorf("create configmap: %v", err)
			return
		}

		if err := client.Delete(context.TODO(), cm, k8s.DeletePropagationForeground()); err != nil {
			t.Fatalf("delete configmap: %v", err)
		}

		var err error
		for i := 0; i < 50; i++ {
			err = func() error {
				err := client.Get(context.TODO(), namespace, *cm2.Metadata.Name, cm2)
				if err == nil {
					return fmt.Errorf("expected 404 error")
				}
				apiErr, ok := err.(*k8s.APIError)
				if !ok {
					return fmt.Errorf("error was not of type APIError: %T %v", err, err)
				}
				if apiErr.Code != 404 {
					return fmt.Errorf("expected 404 error code, got %d", apiErr.Code)
				}
				return nil
			}()
			if err == nil {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		if err != nil {
			t.Errorf("getting parent configmap: %v", err)
		}
	})
}

func TestDeleteOrphan(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		cm := &corev1.ConfigMap{
			Metadata: &metav1.ObjectMeta{
				Name:      k8s.String("my-configmap"),
				Namespace: &namespace,
			},
		}
		if err := client.Create(context.TODO(), cm); err != nil {
			t.Errorf("create configmap: %v", err)
			return
		}

		cm2 := &corev1.ConfigMap{
			Metadata: &metav1.ObjectMeta{
				Name:      k8s.String("my-configmap-2"),
				Namespace: &namespace,
				OwnerReferences: []*metav1.OwnerReference{
					{
						ApiVersion: k8s.String("v1"),
						Kind:       k8s.String("ConfigMap"),
						Name:       cm.Metadata.Name,
						Uid:        cm.Metadata.Uid,
					},
				},
			},
		}
		if err := client.Create(context.TODO(), cm2); err != nil {
			t.Errorf("create configmap: %v", err)
			return
		}

		if err := client.Delete(context.TODO(), cm, k8s.DeletePropagationOrphan()); err != nil {
			t.Fatalf("delete configmap: %v", err)
		}

		// We're attempting to affirm a negative, that this won't eventually be deleted.
		// Since there's no explicit even for us to wait for
		time.Sleep(time.Second)

		err := client.Get(context.TODO(), namespace, *cm2.Metadata.Name, cm2)
		if err != nil {
			t.Errorf("failed to get configmap: %v", err)
		}
	})
}

func Test404(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		var configMap corev1.ConfigMap
		err := client.Get(context.TODO(), namespace, "i-dont-exist", &configMap)
		if err == nil {
			t.Errorf("expected 404 error")
			return
		}
		apiErr, ok := err.(*k8s.APIError)
		if !ok {
			t.Errorf("error was not of type APIError: %T %v", err, err)
			return
		}
		if apiErr.Code != 404 {
			t.Errorf("expected 404 error code, got %d", apiErr.Code)
		}
	})
}

func Test404JSON(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		// Use a JSON object to ensure JSON payload errors deserialize correctly.
		var configMap configMapJSON
		err := client.Get(context.TODO(), namespace, "i-dont-exist", &configMap)
		if err == nil {
			t.Errorf("expected 404 error")
			return
		}
		apiErr, ok := err.(*k8s.APIError)
		if !ok {
			t.Errorf("error was not of type APIError: %T %v", err, err)
			return
		}
		if apiErr.Code != 404 {
			t.Errorf("expected 404 error code, got %d", apiErr.Code)
		}
	})
}

func TestLabelSelector(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		for i := 0; i < 5; i++ {
			cm := &corev1.ConfigMap{
				Metadata: &metav1.ObjectMeta{
					Name:      k8s.String(fmt.Sprintf("my-configmap-%d", i)),
					Namespace: &namespace,
					Labels: map[string]string{
						"configmap": "true",
						"n":         strconv.Itoa(i),
						"m":         strconv.Itoa(i % 2),
					},
				},
			}
			if err := client.Create(context.TODO(), cm); err != nil {
				t.Errorf("create configmap: %v", err)
				return
			}
		}

		tests := []struct {
			setup func(l *k8s.LabelSelector)
			want  int
		}{
			{
				func(l *k8s.LabelSelector) {
					l.Eq("configmap", "true")
				},
				5,
			},
			{
				func(l *k8s.LabelSelector) {
					l.Eq("configmap", "true")
					l.NotEq("n", "4")
				},
				4,
			},
			{
				func(l *k8s.LabelSelector) {
					l.Eq("configmap", "false")
				},
				0,
			},
			{
				func(l *k8s.LabelSelector) {
					l.Eq("configmap", "true")
					l.Eq("n", "4")
				},
				1,
			},
			{
				func(l *k8s.LabelSelector) {
					l.Eq("configmap", "true")
					l.Eq("m", "0")
				},
				3,
			},
		}

		for _, test := range tests {
			var configmaps corev1.ConfigMapList
			l := new(k8s.LabelSelector)
			test.setup(l)
			ctx := context.TODO()
			if err := client.List(ctx, namespace, &configmaps, l.Selector()); err != nil {
				t.Fatalf("list configmaps: %v", err)
			}
			if len(configmaps.Items) != test.want {
				t.Errorf("label selector %s expected %d items, got %d",
					l, test.want, len(configmaps.Items))
			}
		}
	})
}
