package k8s

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/ericchiang/k8s/api/v1"
)

const skipMsg = `
warning: this package's test run using the default context of your "kubeclt" command,
and will create resources on your cluster (mostly configmaps).

If you wish to continue set the following environment variable:

	export K8S_CLIENT_TEST=1

To suppress this message, set:

	export K8S_CLIENT_TEST=0
`

func newTestClient(t *testing.T) *Client {
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

	config := new(Config)
	if err := json.Unmarshal(out, config); err != nil {
		t.Fatalf("parse kubeconfig: %v '%s'", err, out)
	}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	return client
}

func newName() string {
	b := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func TestNewTestClient(t *testing.T) {
	newTestClient(t)
}

func TestHTTP2(t *testing.T) {
	client := newTestClient(t)
	req, err := http.NewRequest("GET", client.urlForPath("/api"), nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if !strings.HasPrefix(resp.Proto, "HTTP/2") {
		t.Errorf("expected proto=HTTP/2.X, got=", resp.Proto)
	}
}

func TestListNodes(t *testing.T) {
	client := newTestClient(t)
	if _, err := client.CoreV1().ListNodes(context.Background()); err != nil {
		t.Fatal("failed to list nodes: %v", err)
	}
}

func TestConfigMaps(t *testing.T) {
	client := newTestClient(t).CoreV1()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	name := newName()
	labelVal := newName()

	cm := &v1.ConfigMap{
		Metadata: &v1.ObjectMeta{
			Name:      String(name),
			Namespace: String("default"),
			Labels: map[string]string{
				"testLabel": labelVal,
			},
		},
		Data: map[string]string{
			"foo": "bar",
		},
	}
	got, err := client.CreateConfigMap(ctx, cm)
	if err != nil {
		t.Fatalf("create config map: %v", err)
	}
	got.Data["zam"] = "spam"
	_, err = client.UpdateConfigMap(ctx, got)
	if err != nil {
		t.Fatalf("update config map: %v", err)
	}

	tests := []struct {
		labelVal string
		expNum   int
	}{
		{labelVal, 1},
		{newName(), 0},
	}
	for _, test := range tests {
		l := new(LabelSelector)
		l.Eq("testLabel", test.labelVal)

		configMaps, err := client.ListConfigMaps(ctx, "default", l.Selector())
		if err != nil {
			t.Errorf("failed to list configmaps: %v", err)
			continue
		}
		got := len(configMaps.Items)
		if got != test.expNum {
			t.Errorf("expected selector to return %d items got %d", test.expNum, got)
		}
	}

	if err := client.DeleteConfigMap(ctx, *cm.Metadata.Name, *cm.Metadata.Namespace); err != nil {
		t.Fatalf("delete config map: %v", err)
	}

}

func TestWatch(t *testing.T) {
	client := newTestClient(t).CoreV1()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w, err := client.WatchConfigMaps(ctx, "default")
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()

	name := newName()
	labelVal := newName()

	cm := &v1.ConfigMap{
		Metadata: &v1.ObjectMeta{
			Name:      String(name),
			Namespace: String("default"),
			Labels: map[string]string{
				"testLabel": labelVal,
			},
		},
		Data: map[string]string{
			"foo": "bar",
		},
	}
	got, err := client.CreateConfigMap(ctx, cm)
	if err != nil {
		t.Fatalf("create config map: %v", err)
	}

	if event, gotFromWatch, err := w.Next(); err != nil {
		t.Errorf("failed to get next watch: %v", err)
	} else {
		if *event.Type != EventAdded {
			t.Errorf("expected event type %q got %q", EventAdded, *event.Type)
		}
		if !reflect.DeepEqual(got, gotFromWatch) {
			t.Errorf("object from add event did not match expected value")
		}
	}

	got.Data["zam"] = "spam"
	got, err = client.UpdateConfigMap(ctx, got)
	if err != nil {
		t.Fatalf("update config map: %v", err)
	}

	if event, gotFromWatch, err := w.Next(); err != nil {
		t.Errorf("failed to get next watch: %v", err)
	} else {
		if *event.Type != EventModified {
			t.Errorf("expected event type %q got %q", EventModified, *event.Type)
		}
		if !reflect.DeepEqual(got, gotFromWatch) {
			t.Errorf("object from modified event did not match expected value")
		}
	}

	tests := []struct {
		labelVal string
		expNum   int
	}{
		{labelVal, 1},
		{newName(), 0},
	}
	for _, test := range tests {
		l := new(LabelSelector)
		l.Eq("testLabel", test.labelVal)

		configMaps, err := client.ListConfigMaps(ctx, "default", l.Selector())
		if err != nil {
			t.Errorf("failed to list configmaps: %v", err)
			continue
		}
		got := len(configMaps.Items)
		if got != test.expNum {
			t.Errorf("expected selector to return %d items got %d", test.expNum, got)
		}
	}

	if err := client.DeleteConfigMap(ctx, *cm.Metadata.Name, *cm.Metadata.Namespace); err != nil {
		t.Fatalf("delete config map: %v", err)
	}
	if event, gotFromWatch, err := w.Next(); err != nil {
		t.Errorf("failed to get next watch: %v", err)
	} else {
		if *event.Type != EventDeleted {
			t.Errorf("expected event type %q got %q", EventDeleted, *event.Type)
		}

		// Resource version will be different after a delete
		got.Metadata.ResourceVersion = String("")
		gotFromWatch.Metadata.ResourceVersion = String("")

		if !reflect.DeepEqual(got, gotFromWatch) {
			t.Errorf("object from deleted event did not match expected value")
		}
	}
}
