package k8s

import (
	"context"
	"fmt"

	apiv1 "github.com/ericchiang/k8s/api/v1"
	appsv1alpha1 "github.com/ericchiang/k8s/apis/apps/v1alpha1"
	authenticationv1beta1 "github.com/ericchiang/k8s/apis/authentication/v1beta1"
	authorizationv1beta1 "github.com/ericchiang/k8s/apis/authorization/v1beta1"
	autoscalingv1 "github.com/ericchiang/k8s/apis/autoscaling/v1"
	batchv1 "github.com/ericchiang/k8s/apis/batch/v1"
	batchv2alpha1 "github.com/ericchiang/k8s/apis/batch/v2alpha1"
	certificatesv1alpha1 "github.com/ericchiang/k8s/apis/certificates/v1alpha1"
	extensionsv1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
	imagepolicyv1alpha1 "github.com/ericchiang/k8s/apis/imagepolicy/v1alpha1"
	policyv1alpha1 "github.com/ericchiang/k8s/apis/policy/v1alpha1"
	rbacv1alpha1 "github.com/ericchiang/k8s/apis/rbac/v1alpha1"
	storagev1beta1 "github.com/ericchiang/k8s/apis/storage/v1beta1"
)

// CoreV1 returns a client for interacting with the /v1 API group.
func (c *Client) CoreV1() *CoreV1 {
	return &CoreV1{c}
}

// CoreV1 is a client for interacting with the /v1 API group.
type CoreV1 struct {
	client *Client
}

func (c *CoreV1) CreateBinding(ctx context.Context, obj *apiv1.Binding) (*apiv1.Binding, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "bindings", "")
	resp := new(apiv1.Binding)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteBinding(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "bindings", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetBinding(ctx context.Context, name string) (*apiv1.Binding, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "bindings", name)
	resp := new(apiv1.Binding)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateComponentStatus(ctx context.Context, obj *apiv1.ComponentStatus) (*apiv1.ComponentStatus, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", md.Namespace, "componentstatuses", "")
	resp := new(apiv1.ComponentStatus)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteComponentStatus(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "componentstatuses", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetComponentStatus(ctx context.Context, name string) (*apiv1.ComponentStatus, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "componentstatuses", name)
	resp := new(apiv1.ComponentStatus)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListComponentStatuses(ctx context.Context) (*apiv1.ComponentStatusList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "componentstatuses", "")
	resp := new(apiv1.ComponentStatusList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateConfigMap(ctx context.Context, obj *apiv1.ConfigMap) (*apiv1.ConfigMap, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "configmaps", "")
	resp := new(apiv1.ConfigMap)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteConfigMap(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "configmaps", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetConfigMap(ctx context.Context, name string) (*apiv1.ConfigMap, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "configmaps", name)
	resp := new(apiv1.ConfigMap)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListConfigMaps(ctx context.Context) (*apiv1.ConfigMapList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "configmaps", "")
	resp := new(apiv1.ConfigMapList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateEndpoints(ctx context.Context, obj *apiv1.Endpoints) (*apiv1.Endpoints, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "endpointses", "")
	resp := new(apiv1.Endpoints)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteEndpoints(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "endpointses", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetEndpoints(ctx context.Context, name string) (*apiv1.Endpoints, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "endpointses", name)
	resp := new(apiv1.Endpoints)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListEndpointses(ctx context.Context) (*apiv1.EndpointsList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "endpointses", "")
	resp := new(apiv1.EndpointsList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateEvent(ctx context.Context, obj *apiv1.Event) (*apiv1.Event, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "events", "")
	resp := new(apiv1.Event)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteEvent(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "events", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetEvent(ctx context.Context, name string) (*apiv1.Event, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "events", name)
	resp := new(apiv1.Event)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListEvents(ctx context.Context) (*apiv1.EventList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "events", "")
	resp := new(apiv1.EventList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateLimitRange(ctx context.Context, obj *apiv1.LimitRange) (*apiv1.LimitRange, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "limitranges", "")
	resp := new(apiv1.LimitRange)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteLimitRange(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "limitranges", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetLimitRange(ctx context.Context, name string) (*apiv1.LimitRange, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "limitranges", name)
	resp := new(apiv1.LimitRange)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListLimitRanges(ctx context.Context) (*apiv1.LimitRangeList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "limitranges", "")
	resp := new(apiv1.LimitRangeList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateNamespace(ctx context.Context, obj *apiv1.Namespace) (*apiv1.Namespace, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", md.Namespace, "namespaces", "")
	resp := new(apiv1.Namespace)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteNamespace(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "namespaces", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetNamespace(ctx context.Context, name string) (*apiv1.Namespace, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "namespaces", name)
	resp := new(apiv1.Namespace)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListNamespaces(ctx context.Context) (*apiv1.NamespaceList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "namespaces", "")
	resp := new(apiv1.NamespaceList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateNode(ctx context.Context, obj *apiv1.Node) (*apiv1.Node, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", md.Namespace, "nodes", "")
	resp := new(apiv1.Node)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteNode(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "nodes", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetNode(ctx context.Context, name string) (*apiv1.Node, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "nodes", name)
	resp := new(apiv1.Node)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListNodes(ctx context.Context) (*apiv1.NodeList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "nodes", "")
	resp := new(apiv1.NodeList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreatePersistentVolume(ctx context.Context, obj *apiv1.PersistentVolume) (*apiv1.PersistentVolume, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", md.Namespace, "persistentvolumes", "")
	resp := new(apiv1.PersistentVolume)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeletePersistentVolume(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "persistentvolumes", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetPersistentVolume(ctx context.Context, name string) (*apiv1.PersistentVolume, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "persistentvolumes", name)
	resp := new(apiv1.PersistentVolume)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListPersistentVolumes(ctx context.Context) (*apiv1.PersistentVolumeList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("", "v1", ns, "persistentvolumes", "")
	resp := new(apiv1.PersistentVolumeList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreatePersistentVolumeClaim(ctx context.Context, obj *apiv1.PersistentVolumeClaim) (*apiv1.PersistentVolumeClaim, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "persistentvolumeclaims", "")
	resp := new(apiv1.PersistentVolumeClaim)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeletePersistentVolumeClaim(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "persistentvolumeclaims", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetPersistentVolumeClaim(ctx context.Context, name string) (*apiv1.PersistentVolumeClaim, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "persistentvolumeclaims", name)
	resp := new(apiv1.PersistentVolumeClaim)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListPersistentVolumeClaims(ctx context.Context) (*apiv1.PersistentVolumeClaimList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "persistentvolumeclaims", "")
	resp := new(apiv1.PersistentVolumeClaimList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreatePod(ctx context.Context, obj *apiv1.Pod) (*apiv1.Pod, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "pods", "")
	resp := new(apiv1.Pod)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeletePod(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "pods", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetPod(ctx context.Context, name string) (*apiv1.Pod, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "pods", name)
	resp := new(apiv1.Pod)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListPods(ctx context.Context) (*apiv1.PodList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "pods", "")
	resp := new(apiv1.PodList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreatePodStatusResult(ctx context.Context, obj *apiv1.PodStatusResult) (*apiv1.PodStatusResult, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "podstatusresults", "")
	resp := new(apiv1.PodStatusResult)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeletePodStatusResult(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "podstatusresults", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetPodStatusResult(ctx context.Context, name string) (*apiv1.PodStatusResult, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "podstatusresults", name)
	resp := new(apiv1.PodStatusResult)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreatePodTemplate(ctx context.Context, obj *apiv1.PodTemplate) (*apiv1.PodTemplate, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "podtemplates", "")
	resp := new(apiv1.PodTemplate)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeletePodTemplate(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "podtemplates", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetPodTemplate(ctx context.Context, name string) (*apiv1.PodTemplate, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "podtemplates", name)
	resp := new(apiv1.PodTemplate)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListPodTemplates(ctx context.Context) (*apiv1.PodTemplateList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "podtemplates", "")
	resp := new(apiv1.PodTemplateList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreatePodTemplateSpec(ctx context.Context, obj *apiv1.PodTemplateSpec) (*apiv1.PodTemplateSpec, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "podtemplatespecs", "")
	resp := new(apiv1.PodTemplateSpec)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeletePodTemplateSpec(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "podtemplatespecs", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetPodTemplateSpec(ctx context.Context, name string) (*apiv1.PodTemplateSpec, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "podtemplatespecs", name)
	resp := new(apiv1.PodTemplateSpec)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateRangeAllocation(ctx context.Context, obj *apiv1.RangeAllocation) (*apiv1.RangeAllocation, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "rangeallocations", "")
	resp := new(apiv1.RangeAllocation)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteRangeAllocation(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "rangeallocations", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetRangeAllocation(ctx context.Context, name string) (*apiv1.RangeAllocation, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "rangeallocations", name)
	resp := new(apiv1.RangeAllocation)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateReplicationController(ctx context.Context, obj *apiv1.ReplicationController) (*apiv1.ReplicationController, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "replicationcontrollers", "")
	resp := new(apiv1.ReplicationController)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteReplicationController(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "replicationcontrollers", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetReplicationController(ctx context.Context, name string) (*apiv1.ReplicationController, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "replicationcontrollers", name)
	resp := new(apiv1.ReplicationController)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListReplicationControllers(ctx context.Context) (*apiv1.ReplicationControllerList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "replicationcontrollers", "")
	resp := new(apiv1.ReplicationControllerList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateResourceQuota(ctx context.Context, obj *apiv1.ResourceQuota) (*apiv1.ResourceQuota, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "resourcequotas", "")
	resp := new(apiv1.ResourceQuota)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteResourceQuota(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "resourcequotas", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetResourceQuota(ctx context.Context, name string) (*apiv1.ResourceQuota, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "resourcequotas", name)
	resp := new(apiv1.ResourceQuota)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListResourceQuotas(ctx context.Context) (*apiv1.ResourceQuotaList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "resourcequotas", "")
	resp := new(apiv1.ResourceQuotaList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateSecret(ctx context.Context, obj *apiv1.Secret) (*apiv1.Secret, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "secrets", "")
	resp := new(apiv1.Secret)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteSecret(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "secrets", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetSecret(ctx context.Context, name string) (*apiv1.Secret, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "secrets", name)
	resp := new(apiv1.Secret)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListSecrets(ctx context.Context) (*apiv1.SecretList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "secrets", "")
	resp := new(apiv1.SecretList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateService(ctx context.Context, obj *apiv1.Service) (*apiv1.Service, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "services", "")
	resp := new(apiv1.Service)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteService(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "services", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetService(ctx context.Context, name string) (*apiv1.Service, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "services", name)
	resp := new(apiv1.Service)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListServices(ctx context.Context) (*apiv1.ServiceList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "services", "")
	resp := new(apiv1.ServiceList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) CreateServiceAccount(ctx context.Context, obj *apiv1.ServiceAccount) (*apiv1.ServiceAccount, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", md.Namespace, "serviceaccounts", "")
	resp := new(apiv1.ServiceAccount)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) DeleteServiceAccount(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "serviceaccounts", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CoreV1) GetServiceAccount(ctx context.Context, name string) (*apiv1.ServiceAccount, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "serviceaccounts", name)
	resp := new(apiv1.ServiceAccount)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CoreV1) ListServiceAccounts(ctx context.Context) (*apiv1.ServiceAccountList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("", "v1", ns, "serviceaccounts", "")
	resp := new(apiv1.ServiceAccountList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// AppsV1 returns a client for interacting with the apps/v1alpha1 API group.
func (c *Client) AppsV1() *AppsV1 {
	return &AppsV1{c}
}

// AppsV1 is a client for interacting with the apps/v1alpha1 API group.
type AppsV1 struct {
	client *Client
}

func (c *AppsV1) CreatePetSet(ctx context.Context, obj *appsv1alpha1.PetSet) (*appsv1alpha1.PetSet, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("apps", "v1alpha1", md.Namespace, "petsets", "")
	resp := new(appsv1alpha1.PetSet)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AppsV1) DeletePetSet(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("apps", "v1alpha1", ns, "petsets", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *AppsV1) GetPetSet(ctx context.Context, name string) (*appsv1alpha1.PetSet, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("apps", "v1alpha1", ns, "petsets", name)
	resp := new(appsv1alpha1.PetSet)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AppsV1) ListPetSets(ctx context.Context) (*appsv1alpha1.PetSetList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("apps", "v1alpha1", ns, "petsets", "")
	resp := new(appsv1alpha1.PetSetList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// AuthenticationV1 returns a client for interacting with the authentication.k8s.io/v1beta1 API group.
func (c *Client) AuthenticationV1() *AuthenticationV1 {
	return &AuthenticationV1{c}
}

// AuthenticationV1 is a client for interacting with the authentication.k8s.io/v1beta1 API group.
type AuthenticationV1 struct {
	client *Client
}

func (c *AuthenticationV1) CreateTokenReview(ctx context.Context, obj *authenticationv1beta1.TokenReview) (*authenticationv1beta1.TokenReview, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("authentication.k8s.io", "v1beta1", md.Namespace, "tokenreviews", "")
	resp := new(authenticationv1beta1.TokenReview)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AuthenticationV1) DeleteTokenReview(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("authentication.k8s.io", "v1beta1", ns, "tokenreviews", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *AuthenticationV1) GetTokenReview(ctx context.Context, name string) (*authenticationv1beta1.TokenReview, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("authentication.k8s.io", "v1beta1", ns, "tokenreviews", name)
	resp := new(authenticationv1beta1.TokenReview)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// AuthorizationV1 returns a client for interacting with the authorization.k8s.io/v1beta1 API group.
func (c *Client) AuthorizationV1() *AuthorizationV1 {
	return &AuthorizationV1{c}
}

// AuthorizationV1 is a client for interacting with the authorization.k8s.io/v1beta1 API group.
type AuthorizationV1 struct {
	client *Client
}

func (c *AuthorizationV1) CreateLocalSubjectAccessReview(ctx context.Context, obj *authorizationv1beta1.LocalSubjectAccessReview) (*authorizationv1beta1.LocalSubjectAccessReview, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("authorization.k8s.io", "v1beta1", md.Namespace, "localsubjectaccessreviews", "")
	resp := new(authorizationv1beta1.LocalSubjectAccessReview)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AuthorizationV1) DeleteLocalSubjectAccessReview(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("authorization.k8s.io", "v1beta1", ns, "localsubjectaccessreviews", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *AuthorizationV1) GetLocalSubjectAccessReview(ctx context.Context, name string) (*authorizationv1beta1.LocalSubjectAccessReview, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("authorization.k8s.io", "v1beta1", ns, "localsubjectaccessreviews", name)
	resp := new(authorizationv1beta1.LocalSubjectAccessReview)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AuthorizationV1) CreateSelfSubjectAccessReview(ctx context.Context, obj *authorizationv1beta1.SelfSubjectAccessReview) (*authorizationv1beta1.SelfSubjectAccessReview, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("authorization.k8s.io", "v1beta1", md.Namespace, "selfsubjectaccessreviews", "")
	resp := new(authorizationv1beta1.SelfSubjectAccessReview)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AuthorizationV1) DeleteSelfSubjectAccessReview(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("authorization.k8s.io", "v1beta1", ns, "selfsubjectaccessreviews", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *AuthorizationV1) GetSelfSubjectAccessReview(ctx context.Context, name string) (*authorizationv1beta1.SelfSubjectAccessReview, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("authorization.k8s.io", "v1beta1", ns, "selfsubjectaccessreviews", name)
	resp := new(authorizationv1beta1.SelfSubjectAccessReview)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AuthorizationV1) CreateSubjectAccessReview(ctx context.Context, obj *authorizationv1beta1.SubjectAccessReview) (*authorizationv1beta1.SubjectAccessReview, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("authorization.k8s.io", "v1beta1", md.Namespace, "subjectaccessreviews", "")
	resp := new(authorizationv1beta1.SubjectAccessReview)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AuthorizationV1) DeleteSubjectAccessReview(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("authorization.k8s.io", "v1beta1", ns, "subjectaccessreviews", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *AuthorizationV1) GetSubjectAccessReview(ctx context.Context, name string) (*authorizationv1beta1.SubjectAccessReview, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("authorization.k8s.io", "v1beta1", ns, "subjectaccessreviews", name)
	resp := new(authorizationv1beta1.SubjectAccessReview)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// AutoscalingV1 returns a client for interacting with the autoscaling/v1 API group.
func (c *Client) AutoscalingV1() *AutoscalingV1 {
	return &AutoscalingV1{c}
}

// AutoscalingV1 is a client for interacting with the autoscaling/v1 API group.
type AutoscalingV1 struct {
	client *Client
}

func (c *AutoscalingV1) CreateHorizontalPodAutoscaler(ctx context.Context, obj *autoscalingv1.HorizontalPodAutoscaler) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("autoscaling", "v1", md.Namespace, "horizontalpodautoscalers", "")
	resp := new(autoscalingv1.HorizontalPodAutoscaler)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AutoscalingV1) DeleteHorizontalPodAutoscaler(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("autoscaling", "v1", ns, "horizontalpodautoscalers", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *AutoscalingV1) GetHorizontalPodAutoscaler(ctx context.Context, name string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("autoscaling", "v1", ns, "horizontalpodautoscalers", name)
	resp := new(autoscalingv1.HorizontalPodAutoscaler)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AutoscalingV1) ListHorizontalPodAutoscalers(ctx context.Context) (*autoscalingv1.HorizontalPodAutoscalerList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("autoscaling", "v1", ns, "horizontalpodautoscalers", "")
	resp := new(autoscalingv1.HorizontalPodAutoscalerList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AutoscalingV1) CreateScale(ctx context.Context, obj *autoscalingv1.Scale) (*autoscalingv1.Scale, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("autoscaling", "v1", md.Namespace, "scales", "")
	resp := new(autoscalingv1.Scale)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AutoscalingV1) DeleteScale(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("autoscaling", "v1", ns, "scales", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *AutoscalingV1) GetScale(ctx context.Context, name string) (*autoscalingv1.Scale, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("autoscaling", "v1", ns, "scales", name)
	resp := new(autoscalingv1.Scale)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// BatchV1 returns a client for interacting with the batch/v1 API group.
func (c *Client) BatchV1() *BatchV1 {
	return &BatchV1{c}
}

// BatchV1 is a client for interacting with the batch/v1 API group.
type BatchV1 struct {
	client *Client
}

func (c *BatchV1) CreateJob(ctx context.Context, obj *batchv1.Job) (*batchv1.Job, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v1", md.Namespace, "jobs", "")
	resp := new(batchv1.Job)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *BatchV1) DeleteJob(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v1", ns, "jobs", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *BatchV1) GetJob(ctx context.Context, name string) (*batchv1.Job, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v1", ns, "jobs", name)
	resp := new(batchv1.Job)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *BatchV1) ListJobs(ctx context.Context) (*batchv1.JobList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v1", ns, "jobs", "")
	resp := new(batchv1.JobList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// BatchV2 returns a client for interacting with the batch/v2alpha1 API group.
func (c *Client) BatchV2() *BatchV2 {
	return &BatchV2{c}
}

// BatchV2 is a client for interacting with the batch/v2alpha1 API group.
type BatchV2 struct {
	client *Client
}

func (c *BatchV2) CreateJob(ctx context.Context, obj *batchv2alpha1.Job) (*batchv2alpha1.Job, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", md.Namespace, "jobs", "")
	resp := new(batchv2alpha1.Job)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *BatchV2) DeleteJob(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", ns, "jobs", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *BatchV2) GetJob(ctx context.Context, name string) (*batchv2alpha1.Job, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", ns, "jobs", name)
	resp := new(batchv2alpha1.Job)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *BatchV2) ListJobs(ctx context.Context) (*batchv2alpha1.JobList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", ns, "jobs", "")
	resp := new(batchv2alpha1.JobList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *BatchV2) CreateJobTemplate(ctx context.Context, obj *batchv2alpha1.JobTemplate) (*batchv2alpha1.JobTemplate, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", md.Namespace, "jobtemplates", "")
	resp := new(batchv2alpha1.JobTemplate)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *BatchV2) DeleteJobTemplate(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", ns, "jobtemplates", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *BatchV2) GetJobTemplate(ctx context.Context, name string) (*batchv2alpha1.JobTemplate, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", ns, "jobtemplates", name)
	resp := new(batchv2alpha1.JobTemplate)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *BatchV2) CreateScheduledJob(ctx context.Context, obj *batchv2alpha1.ScheduledJob) (*batchv2alpha1.ScheduledJob, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", md.Namespace, "scheduledjobs", "")
	resp := new(batchv2alpha1.ScheduledJob)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *BatchV2) DeleteScheduledJob(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", ns, "scheduledjobs", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *BatchV2) GetScheduledJob(ctx context.Context, name string) (*batchv2alpha1.ScheduledJob, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", ns, "scheduledjobs", name)
	resp := new(batchv2alpha1.ScheduledJob)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *BatchV2) ListScheduledJobs(ctx context.Context) (*batchv2alpha1.ScheduledJobList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("batch", "v2alpha1", ns, "scheduledjobs", "")
	resp := new(batchv2alpha1.ScheduledJobList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// CertificatesV1 returns a client for interacting with the certificates.k8s.io/v1alpha1 API group.
func (c *Client) CertificatesV1() *CertificatesV1 {
	return &CertificatesV1{c}
}

// CertificatesV1 is a client for interacting with the certificates.k8s.io/v1alpha1 API group.
type CertificatesV1 struct {
	client *Client
}

func (c *CertificatesV1) CreateCertificateSigningRequest(ctx context.Context, obj *certificatesv1alpha1.CertificateSigningRequest) (*certificatesv1alpha1.CertificateSigningRequest, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("certificates.k8s.io", "v1alpha1", md.Namespace, "certificatesigningrequests", "")
	resp := new(certificatesv1alpha1.CertificateSigningRequest)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CertificatesV1) DeleteCertificateSigningRequest(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("certificates.k8s.io", "v1alpha1", ns, "certificatesigningrequests", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *CertificatesV1) GetCertificateSigningRequest(ctx context.Context, name string) (*certificatesv1alpha1.CertificateSigningRequest, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("certificates.k8s.io", "v1alpha1", ns, "certificatesigningrequests", name)
	resp := new(certificatesv1alpha1.CertificateSigningRequest)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CertificatesV1) ListCertificateSigningRequests(ctx context.Context) (*certificatesv1alpha1.CertificateSigningRequestList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("certificates.k8s.io", "v1alpha1", ns, "certificatesigningrequests", "")
	resp := new(certificatesv1alpha1.CertificateSigningRequestList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// ExtensionsV1 returns a client for interacting with the extensions/v1beta1 API group.
func (c *Client) ExtensionsV1() *ExtensionsV1 {
	return &ExtensionsV1{c}
}

// ExtensionsV1 is a client for interacting with the extensions/v1beta1 API group.
type ExtensionsV1 struct {
	client *Client
}

func (c *ExtensionsV1) CreateDaemonSet(ctx context.Context, obj *extensionsv1beta1.DaemonSet) (*extensionsv1beta1.DaemonSet, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "daemonsets", "")
	resp := new(extensionsv1beta1.DaemonSet)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteDaemonSet(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "daemonsets", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetDaemonSet(ctx context.Context, name string) (*extensionsv1beta1.DaemonSet, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "daemonsets", name)
	resp := new(extensionsv1beta1.DaemonSet)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListDaemonSets(ctx context.Context) (*extensionsv1beta1.DaemonSetList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "daemonsets", "")
	resp := new(extensionsv1beta1.DaemonSetList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreateDeployment(ctx context.Context, obj *extensionsv1beta1.Deployment) (*extensionsv1beta1.Deployment, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "deployments", "")
	resp := new(extensionsv1beta1.Deployment)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteDeployment(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "deployments", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetDeployment(ctx context.Context, name string) (*extensionsv1beta1.Deployment, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "deployments", name)
	resp := new(extensionsv1beta1.Deployment)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListDeployments(ctx context.Context) (*extensionsv1beta1.DeploymentList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "deployments", "")
	resp := new(extensionsv1beta1.DeploymentList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreateHorizontalPodAutoscaler(ctx context.Context, obj *extensionsv1beta1.HorizontalPodAutoscaler) (*extensionsv1beta1.HorizontalPodAutoscaler, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "horizontalpodautoscalers", "")
	resp := new(extensionsv1beta1.HorizontalPodAutoscaler)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteHorizontalPodAutoscaler(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "horizontalpodautoscalers", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetHorizontalPodAutoscaler(ctx context.Context, name string) (*extensionsv1beta1.HorizontalPodAutoscaler, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "horizontalpodautoscalers", name)
	resp := new(extensionsv1beta1.HorizontalPodAutoscaler)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListHorizontalPodAutoscalers(ctx context.Context) (*extensionsv1beta1.HorizontalPodAutoscalerList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "horizontalpodautoscalers", "")
	resp := new(extensionsv1beta1.HorizontalPodAutoscalerList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreateIngress(ctx context.Context, obj *extensionsv1beta1.Ingress) (*extensionsv1beta1.Ingress, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "ingresses", "")
	resp := new(extensionsv1beta1.Ingress)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteIngress(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "ingresses", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetIngress(ctx context.Context, name string) (*extensionsv1beta1.Ingress, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "ingresses", name)
	resp := new(extensionsv1beta1.Ingress)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListIngresses(ctx context.Context) (*extensionsv1beta1.IngressList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "ingresses", "")
	resp := new(extensionsv1beta1.IngressList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreateJob(ctx context.Context, obj *extensionsv1beta1.Job) (*extensionsv1beta1.Job, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "jobs", "")
	resp := new(extensionsv1beta1.Job)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteJob(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "jobs", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetJob(ctx context.Context, name string) (*extensionsv1beta1.Job, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "jobs", name)
	resp := new(extensionsv1beta1.Job)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListJobs(ctx context.Context) (*extensionsv1beta1.JobList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "jobs", "")
	resp := new(extensionsv1beta1.JobList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreateNetworkPolicy(ctx context.Context, obj *extensionsv1beta1.NetworkPolicy) (*extensionsv1beta1.NetworkPolicy, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "networkpolicies", "")
	resp := new(extensionsv1beta1.NetworkPolicy)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteNetworkPolicy(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "networkpolicies", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetNetworkPolicy(ctx context.Context, name string) (*extensionsv1beta1.NetworkPolicy, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "networkpolicies", name)
	resp := new(extensionsv1beta1.NetworkPolicy)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListNetworkPolicies(ctx context.Context) (*extensionsv1beta1.NetworkPolicyList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "networkpolicies", "")
	resp := new(extensionsv1beta1.NetworkPolicyList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreatePodSecurityPolicy(ctx context.Context, obj *extensionsv1beta1.PodSecurityPolicy) (*extensionsv1beta1.PodSecurityPolicy, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "podsecuritypolicies", "")
	resp := new(extensionsv1beta1.PodSecurityPolicy)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeletePodSecurityPolicy(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("extensions", "v1beta1", ns, "podsecuritypolicies", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetPodSecurityPolicy(ctx context.Context, name string) (*extensionsv1beta1.PodSecurityPolicy, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("extensions", "v1beta1", ns, "podsecuritypolicies", name)
	resp := new(extensionsv1beta1.PodSecurityPolicy)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListPodSecurityPolicies(ctx context.Context) (*extensionsv1beta1.PodSecurityPolicyList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("extensions", "v1beta1", ns, "podsecuritypolicies", "")
	resp := new(extensionsv1beta1.PodSecurityPolicyList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreateReplicaSet(ctx context.Context, obj *extensionsv1beta1.ReplicaSet) (*extensionsv1beta1.ReplicaSet, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "replicasets", "")
	resp := new(extensionsv1beta1.ReplicaSet)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteReplicaSet(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "replicasets", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetReplicaSet(ctx context.Context, name string) (*extensionsv1beta1.ReplicaSet, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "replicasets", name)
	resp := new(extensionsv1beta1.ReplicaSet)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListReplicaSets(ctx context.Context) (*extensionsv1beta1.ReplicaSetList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "replicasets", "")
	resp := new(extensionsv1beta1.ReplicaSetList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreateScale(ctx context.Context, obj *extensionsv1beta1.Scale) (*extensionsv1beta1.Scale, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "scales", "")
	resp := new(extensionsv1beta1.Scale)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteScale(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "scales", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetScale(ctx context.Context, name string) (*extensionsv1beta1.Scale, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "scales", name)
	resp := new(extensionsv1beta1.Scale)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreateThirdPartyResource(ctx context.Context, obj *extensionsv1beta1.ThirdPartyResource) (*extensionsv1beta1.ThirdPartyResource, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "thirdpartyresources", "")
	resp := new(extensionsv1beta1.ThirdPartyResource)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteThirdPartyResource(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("extensions", "v1beta1", ns, "thirdpartyresources", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetThirdPartyResource(ctx context.Context, name string) (*extensionsv1beta1.ThirdPartyResource, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("extensions", "v1beta1", ns, "thirdpartyresources", name)
	resp := new(extensionsv1beta1.ThirdPartyResource)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListThirdPartyResources(ctx context.Context) (*extensionsv1beta1.ThirdPartyResourceList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("extensions", "v1beta1", ns, "thirdpartyresources", "")
	resp := new(extensionsv1beta1.ThirdPartyResourceList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) CreateThirdPartyResourceData(ctx context.Context, obj *extensionsv1beta1.ThirdPartyResourceData) (*extensionsv1beta1.ThirdPartyResourceData, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", md.Namespace, "thirdpartyresourcedatas", "")
	resp := new(extensionsv1beta1.ThirdPartyResourceData)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) DeleteThirdPartyResourceData(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "thirdpartyresourcedatas", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ExtensionsV1) GetThirdPartyResourceData(ctx context.Context, name string) (*extensionsv1beta1.ThirdPartyResourceData, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "thirdpartyresourcedatas", name)
	resp := new(extensionsv1beta1.ThirdPartyResourceData)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ExtensionsV1) ListThirdPartyResourceDatas(ctx context.Context) (*extensionsv1beta1.ThirdPartyResourceDataList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("extensions", "v1beta1", ns, "thirdpartyresourcedatas", "")
	resp := new(extensionsv1beta1.ThirdPartyResourceDataList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// ImagepolicyV1 returns a client for interacting with the imagepolicy/v1alpha1 API group.
func (c *Client) ImagepolicyV1() *ImagepolicyV1 {
	return &ImagepolicyV1{c}
}

// ImagepolicyV1 is a client for interacting with the imagepolicy/v1alpha1 API group.
type ImagepolicyV1 struct {
	client *Client
}

func (c *ImagepolicyV1) CreateImageReview(ctx context.Context, obj *imagepolicyv1alpha1.ImageReview) (*imagepolicyv1alpha1.ImageReview, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("imagepolicy", "v1alpha1", md.Namespace, "imagereviews", "")
	resp := new(imagepolicyv1alpha1.ImageReview)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImagepolicyV1) DeleteImageReview(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("imagepolicy", "v1alpha1", ns, "imagereviews", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *ImagepolicyV1) GetImageReview(ctx context.Context, name string) (*imagepolicyv1alpha1.ImageReview, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("imagepolicy", "v1alpha1", ns, "imagereviews", name)
	resp := new(imagepolicyv1alpha1.ImageReview)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// PolicyV1 returns a client for interacting with the policy/v1alpha1 API group.
func (c *Client) PolicyV1() *PolicyV1 {
	return &PolicyV1{c}
}

// PolicyV1 is a client for interacting with the policy/v1alpha1 API group.
type PolicyV1 struct {
	client *Client
}

func (c *PolicyV1) CreateEviction(ctx context.Context, obj *policyv1alpha1.Eviction) (*policyv1alpha1.Eviction, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("policy", "v1alpha1", md.Namespace, "evictions", "")
	resp := new(policyv1alpha1.Eviction)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PolicyV1) DeleteEviction(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("policy", "v1alpha1", ns, "evictions", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *PolicyV1) GetEviction(ctx context.Context, name string) (*policyv1alpha1.Eviction, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("policy", "v1alpha1", ns, "evictions", name)
	resp := new(policyv1alpha1.Eviction)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PolicyV1) CreatePodDisruptionBudget(ctx context.Context, obj *policyv1alpha1.PodDisruptionBudget) (*policyv1alpha1.PodDisruptionBudget, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("policy", "v1alpha1", md.Namespace, "poddisruptionbudgets", "")
	resp := new(policyv1alpha1.PodDisruptionBudget)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PolicyV1) DeletePodDisruptionBudget(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("policy", "v1alpha1", ns, "poddisruptionbudgets", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *PolicyV1) GetPodDisruptionBudget(ctx context.Context, name string) (*policyv1alpha1.PodDisruptionBudget, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("policy", "v1alpha1", ns, "poddisruptionbudgets", name)
	resp := new(policyv1alpha1.PodDisruptionBudget)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PolicyV1) ListPodDisruptionBudgets(ctx context.Context) (*policyv1alpha1.PodDisruptionBudgetList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("policy", "v1alpha1", ns, "poddisruptionbudgets", "")
	resp := new(policyv1alpha1.PodDisruptionBudgetList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// RBACV1 returns a client for interacting with the rbac.authorization.k8s.io/v1alpha1 API group.
func (c *Client) RBACV1() *RBACV1 {
	return &RBACV1{c}
}

// RBACV1 is a client for interacting with the rbac.authorization.k8s.io/v1alpha1 API group.
type RBACV1 struct {
	client *Client
}

func (c *RBACV1) CreateClusterRole(ctx context.Context, obj *rbacv1alpha1.ClusterRole) (*rbacv1alpha1.ClusterRole, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", md.Namespace, "clusterroles", "")
	resp := new(rbacv1alpha1.ClusterRole)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) DeleteClusterRole(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "clusterroles", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *RBACV1) GetClusterRole(ctx context.Context, name string) (*rbacv1alpha1.ClusterRole, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "clusterroles", name)
	resp := new(rbacv1alpha1.ClusterRole)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) ListClusterRoles(ctx context.Context) (*rbacv1alpha1.ClusterRoleList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "clusterroles", "")
	resp := new(rbacv1alpha1.ClusterRoleList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) CreateClusterRoleBinding(ctx context.Context, obj *rbacv1alpha1.ClusterRoleBinding) (*rbacv1alpha1.ClusterRoleBinding, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", md.Namespace, "clusterrolebindings", "")
	resp := new(rbacv1alpha1.ClusterRoleBinding)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) DeleteClusterRoleBinding(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "clusterrolebindings", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *RBACV1) GetClusterRoleBinding(ctx context.Context, name string) (*rbacv1alpha1.ClusterRoleBinding, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "clusterrolebindings", name)
	resp := new(rbacv1alpha1.ClusterRoleBinding)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) ListClusterRoleBindings(ctx context.Context) (*rbacv1alpha1.ClusterRoleBindingList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "clusterrolebindings", "")
	resp := new(rbacv1alpha1.ClusterRoleBindingList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) CreateRole(ctx context.Context, obj *rbacv1alpha1.Role) (*rbacv1alpha1.Role, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", md.Namespace, "roles", "")
	resp := new(rbacv1alpha1.Role)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) DeleteRole(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "roles", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *RBACV1) GetRole(ctx context.Context, name string) (*rbacv1alpha1.Role, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "roles", name)
	resp := new(rbacv1alpha1.Role)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) ListRoles(ctx context.Context) (*rbacv1alpha1.RoleList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "roles", "")
	resp := new(rbacv1alpha1.RoleList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) CreateRoleBinding(ctx context.Context, obj *rbacv1alpha1.RoleBinding) (*rbacv1alpha1.RoleBinding, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", md.Namespace, "rolebindings", "")
	resp := new(rbacv1alpha1.RoleBinding)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) DeleteRoleBinding(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "rolebindings", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *RBACV1) GetRoleBinding(ctx context.Context, name string) (*rbacv1alpha1.RoleBinding, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "rolebindings", name)
	resp := new(rbacv1alpha1.RoleBinding)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RBACV1) ListRoleBindings(ctx context.Context) (*rbacv1alpha1.RoleBindingList, error) {
	ns := c.client.namespaceFor(ctx, true)
	url := c.client.urlFor("rbac.authorization.k8s.io", "v1alpha1", ns, "rolebindings", "")
	resp := new(rbacv1alpha1.RoleBindingList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}


// StorageV1 returns a client for interacting with the storage.k8s.io/v1beta1 API group.
func (c *Client) StorageV1() *StorageV1 {
	return &StorageV1{c}
}

// StorageV1 is a client for interacting with the storage.k8s.io/v1beta1 API group.
type StorageV1 struct {
	client *Client
}

func (c *StorageV1) CreateStorageClass(ctx context.Context, obj *storagev1beta1.StorageClass) (*storagev1beta1.StorageClass, error) {
	md := obj.GetMetadata()
	if md.Name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	md.Namespace = c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("storage.k8s.io", "v1beta1", md.Namespace, "storageclasses", "")
	resp := new(storagev1beta1.StorageClass)
	err := c.client.create(ctx, pbCodec, url, obj, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *StorageV1) DeleteStorageClass(ctx context.Context, name string) error {
	if name == "" {
		return fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("storage.k8s.io", "v1beta1", ns, "storageclasses", name)
	return c.client.delete(ctx, pbCodec, url)
}

func (c *StorageV1) GetStorageClass(ctx context.Context, name string) (*storagev1beta1.StorageClass, error) {
	if name == "" {
		return nil, fmt.Errorf("create: no name for given object")
	}
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("storage.k8s.io", "v1beta1", ns, "storageclasses", name)
	resp := new(storagev1beta1.StorageClass)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *StorageV1) ListStorageClasses(ctx context.Context) (*storagev1beta1.StorageClassList, error) {
	ns := c.client.namespaceFor(ctx, false)
	url := c.client.urlFor("storage.k8s.io", "v1beta1", ns, "storageclasses", "")
	resp := new(storagev1beta1.StorageClassList)
	if err := c.client.get(ctx, pbCodec, url, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

