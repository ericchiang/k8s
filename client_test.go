package k8s

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"os/exec"
	"testing"
)

func newTestClient(t *testing.T) *Client {
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

	configMaps, err := client.ListConfigMaps(ctx, "default")
	if err != nil {
		t.Fatal("failed to list configmaps: %v", err)
	}
	_ = configMaps
}
