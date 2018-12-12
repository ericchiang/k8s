package k8s_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

// configMapJSON is used to test the JSON serialization watch.
type configMapJSON struct {
	Metadata *metav1.ObjectMeta `json:"metadata"`
	Data     map[string]string  `json:"data"`
}

func (c *configMapJSON) GetMetadata() *metav1.ObjectMeta {
	return c.Metadata
}

func init() {
	k8s.Register("", "v1", "configmaps", true, &configMapJSON{})
}

func wantEvent(t *testing.T, w *k8s.Watcher, eventType string, got, want k8s.Resource) {
	t.Helper()
	eT, err := w.Next(got)
	if err != nil {
		t.Errorf("decode watch event: %v", err)
		return
	}
	if eT != eventType {
		t.Errorf("expected event type %q got %q", eventType, eT)
	}
	want.GetMetadata().ResourceVersion = k8s.String("")
	got.GetMetadata().ResourceVersion = k8s.String("")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("configmaps didn't match")
		t.Errorf("want: %#v", want)
		t.Errorf(" got: %#v", got)
	}
}

func testWatch(t *testing.T, client *k8s.Client, namespace string, r k8s.Resource, newCM func() k8s.Resource, update func(cm k8s.Resource)) {
	cm := newCM()

	if r.GetMetadata() != nil {
		// Individual watch must created beforehand
		if err := client.Create(context.TODO(), cm); err != nil {
			t.Errorf("create configmap: %v", err)
			return
		}
	}
	w, err := client.Watch(context.TODO(), namespace, r)
	if err != nil {
		t.Fatalf("watch configmaps: %v", err)
	}
	defer w.Close()

	if r.GetMetadata() == nil {
		if err := client.Create(context.TODO(), cm); err != nil {
			t.Errorf("create configmap: %v", err)
			return
		}
		wantEvent(t, w, k8s.EventAdded, newCM(), cm)
	}

	update(cm)

	if err := client.Update(context.TODO(), cm); err != nil {
		t.Errorf("update configmap: %v", err)
		return
	}
	wantEvent(t, w, k8s.EventModified, newCM(), cm)

	if err := client.Delete(context.TODO(), cm); err != nil {
		t.Errorf("Delete configmap: %v", err)
		return
	}
	wantEvent(t, w, k8s.EventDeleted, newCM(), cm)
}

func TestWatchConfigMapJSON(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		newCM := func() k8s.Resource {
			return &configMapJSON{
				Metadata: &metav1.ObjectMeta{
					Name:      k8s.String("my-configmap"),
					Namespace: &namespace,
				},
			}
		}

		updateCM := func(cm k8s.Resource) {
			(cm.(*configMapJSON)).Data = map[string]string{"hello": "world"}
		}
		testWatch(t, client, namespace, &configMapJSON{}, newCM, updateCM)
	})
}

func TestWatchConfigMapProto(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		newCM := func() k8s.Resource {
			return &corev1.ConfigMap{
				Metadata: &metav1.ObjectMeta{
					Name:      k8s.String("my-configmap"),
					Namespace: &namespace,
				},
			}
		}

		updateCM := func(cm k8s.Resource) {
			(cm.(*corev1.ConfigMap)).Data = map[string]string{"hello": "world"}
		}
		testWatch(t, client, namespace, &corev1.ConfigMap{}, newCM, updateCM)
	})
}

func TestWatchIndividualConfigMap(t *testing.T) {
	withNamespace(t, func(client *k8s.Client, namespace string) {
		newCM := func() k8s.Resource {
			return &corev1.ConfigMap{
				Metadata: &metav1.ObjectMeta{
					Name:      k8s.String("my-configmap"),
					Namespace: &namespace,
				},
			}
		}

		updateCM := func(cm k8s.Resource) {
			(cm.(*corev1.ConfigMap)).Data = map[string]string{"hello": "world"}
		}
		testWatch(t, client, namespace, newCM(), newCM, updateCM)
	})
}
