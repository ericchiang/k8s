package k8s

import (
	"testing"
	"time"

	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

// Redefine types since all API groups import "github.com/ericchiang/k8s"
// We can't use them here because it'll create a circular import cycle.

type Pod struct {
	Metadata *metav1.ObjectMeta
}

type PodList struct {
	Metadata *metav1.ListMeta
}

func (p *Pod) GetMetadata() *metav1.ObjectMeta   { return p.Metadata }
func (p *PodList) GetMetadata() *metav1.ListMeta { return p.Metadata }

type Deployment struct {
	Metadata *metav1.ObjectMeta
}

type DeploymentList struct {
	Metadata *metav1.ListMeta
}

func (p *Deployment) GetMetadata() *metav1.ObjectMeta   { return p.Metadata }
func (p *DeploymentList) GetMetadata() *metav1.ListMeta { return p.Metadata }

type ClusterRole struct {
	Metadata *metav1.ObjectMeta
}

type ClusterRoleList struct {
	Metadata *metav1.ListMeta
}

func (p *ClusterRole) GetMetadata() *metav1.ObjectMeta   { return p.Metadata }
func (p *ClusterRoleList) GetMetadata() *metav1.ListMeta { return p.Metadata }

func init() {
	Register("", "v1", "pods", true, &Pod{})
	RegisterList("", "v1", "pods", true, &PodList{})

	Register("apps", "v1beta2", "deployments", true, &Deployment{})
	RegisterList("apps", "v1beta2", "deployments", true, &DeploymentList{})

	Register("rbac.authorization.k8s.io", "v1", "clusterroles", false, &ClusterRole{})
	RegisterList("rbac.authorization.k8s.io", "v1", "clusterroles", false, &ClusterRoleList{})
}

func TestResourceURL(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		resource Resource
		withName bool
		options  []Option
		want     string
		wantErr  bool
	}{
		{
			name:     "pod",
			endpoint: "https://example.com",
			resource: &Pod{
				Metadata: &metav1.ObjectMeta{
					Namespace: String("my-namespace"),
					Name:      String("my-pod"),
				},
			},
			want: "https://example.com/api/v1/namespaces/my-namespace/pods",
		},
		{
			name:     "deployment",
			endpoint: "https://example.com",
			resource: &Deployment{
				Metadata: &metav1.ObjectMeta{
					Namespace: String("my-namespace"),
					Name:      String("my-deployment"),
				},
			},
			want: "https://example.com/apis/apps/v1beta2/namespaces/my-namespace/deployments",
		},
		{
			name:     "deployment-with-name",
			endpoint: "https://example.com",
			resource: &Deployment{
				Metadata: &metav1.ObjectMeta{
					Namespace: String("my-namespace"),
					Name:      String("my-deployment"),
				},
			},
			withName: true,
			want:     "https://example.com/apis/apps/v1beta2/namespaces/my-namespace/deployments/my-deployment",
		},
		{
			name:     "deployment-with-subresource",
			endpoint: "https://example.com",
			resource: &Deployment{
				Metadata: &metav1.ObjectMeta{
					Namespace: String("my-namespace"),
					Name:      String("my-deployment"),
				},
			},
			withName: true,
			options: []Option{
				Subresource("status"),
			},
			want: "https://example.com/apis/apps/v1beta2/namespaces/my-namespace/deployments/my-deployment/status",
		},
		{
			name:     "pod-with-timeout",
			endpoint: "https://example.com",
			resource: &Pod{
				Metadata: &metav1.ObjectMeta{
					Namespace: String("my-namespace"),
					Name:      String("my-pod"),
				},
			},
			options: []Option{
				Timeout(time.Minute),
			},
			want: "https://example.com/api/v1/namespaces/my-namespace/pods?timeoutSeconds=60",
		},
		{
			name:     "pod-with-resource-version",
			endpoint: "https://example.com",
			resource: &Pod{
				Metadata: &metav1.ObjectMeta{
					Namespace: String("my-namespace"),
					Name:      String("my-pod"),
				},
			},
			options: []Option{
				ResourceVersion("foo"),
			},
			want: "https://example.com/api/v1/namespaces/my-namespace/pods?resourceVersion=foo",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := resourceURL(test.endpoint, test.resource, test.withName, test.options...)
			if err != nil {
				if test.wantErr {
					return
				}
				t.Fatalf("constructing resource URL: %v", err)
			}
			if test.wantErr {
				t.Fatal("expected error")
			}
			if test.want != got {
				t.Errorf("wanted=%q, got=%q", test.want, got)
			}
		})
	}
}

func TestResourceWatchURL(t *testing.T) {
	tests := []struct {
		name      string
		endpoint  string
		namespace string
		resource  Resource
		options   []Option
		want      string
		wantErr   bool
	}{
		{
			name:      "watch_pods",
			namespace: "my-namespace",
			endpoint:  "https://k8s.example.com/foo/",
			resource:  &Pod{},
			want:      "https://k8s.example.com/foo/api/v1/namespaces/my-namespace/pods?watch=true",
		},
		{
			name:     "watch_all_pods",
			endpoint: "https://k8s.example.com/foo/",
			resource: &Pod{},
			want:     "https://k8s.example.com/foo/api/v1/pods?watch=true",
		},
		{
			name:      "watch_deployments",
			namespace: "my-namespace",
			endpoint:  "https://k8s.example.com/foo/",
			resource:  &Deployment{},
			want:      "https://k8s.example.com/foo/apis/apps/v1beta2/namespaces/my-namespace/deployments?watch=true",
		},
		{
			name:      "watch_with_options",
			namespace: "my-namespace",
			endpoint:  "https://k8s.example.com/foo/",
			resource:  &Deployment{},
			options: []Option{
				Timeout(time.Minute),
			},
			want: "https://k8s.example.com/foo/apis/apps/v1beta2/namespaces/my-namespace/deployments?timeoutSeconds=60&watch=true",
		},
		{
			name:     "watch_non_namespaced",
			endpoint: "https://k8s.example.com/foo/",
			resource: &ClusterRole{},
			want:     "https://k8s.example.com/foo/apis/rbac.authorization.k8s.io/v1/clusterroles?watch=true",
		},
		{
			name:      "watch_non_namespaced_with_namespace",
			namespace: "my-namespace",
			endpoint:  "https://k8s.example.com/foo/",
			resource:  &ClusterRole{},
			wantErr:   true, // can't provide a namespace for a non-namespaced resource
		},
		{
			name:     "watch_deployment",
			endpoint: "https://k8s.example.com/foo/",
			resource: &Deployment{
				Metadata: &metav1.ObjectMeta{
					Namespace: String("my-namespace"),
					Name:      String("my-deployment"),
				},
			},
			want: "https://k8s.example.com/foo/apis/apps/v1beta2/namespaces/my-namespace/deployments/my-deployment?watch=true",
		},
		{
			name:      "watch_deployment_ns_in_call",
			endpoint:  "https://k8s.example.com/foo/",
			namespace: "my-namespace",
			resource: &Deployment{
				Metadata: &metav1.ObjectMeta{
					Name: String("my-deployment"),
				},
			},
			want: "https://k8s.example.com/foo/apis/apps/v1beta2/namespaces/my-namespace/deployments/my-deployment?watch=true",
		},
		{
			name:      "watch_deployment_mismatched_ns",
			endpoint:  "https://k8s.example.com/foo/",
			namespace: "my-other-namespace",
			resource: &Deployment{
				Metadata: &metav1.ObjectMeta{
					Namespace: String("my-namespace"),
					Name:      String("my-deployment"),
				},
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := resourceWatchURL(
				test.endpoint,
				test.namespace,
				test.resource,
				test.options...,
			)
			if err != nil {
				if !test.wantErr {
					t.Fatalf("resourceWatchURL: %v", err)
				}
				return
			}
			if got != test.want {
				t.Errorf("want: %q", test.want)
				t.Errorf("got : %q", got)
			}
		})
	}
}
