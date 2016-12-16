package k8s

import (
	"context"
	"errors"

	"github.com/ericchiang/k8s/api/v1"
)

type ThirdPartyResources struct {
	c *Client

	apiGroup   string
	apiVersion string
}

// ThirdPartyResources returns a client for interacting with a ThirdPartyResource
// API group.
func (c *Client) ThirdPartyResources(apiGroup, apiVersion string) *ThirdPartyResources {
	return &ThirdPartyResources{c, apiGroup, apiVersion}
}

func checkResource(apiGroup, apiVersion, resource, namespace, name string) error {
	if apiGroup == "" {
		return errors.New("no api group provided")
	}
	if apiVersion == "" {
		return errors.New("no api version provided")
	}
	if resource == "" {
		return errors.New("no resource version provided")
	}
	if namespace == "" {
		return errors.New("no namespace provided")
	}
	if name == "" {
		return errors.New("no resource name provided")
	}
	return nil
}

type object interface {
	GetMetadata() *v1.ObjectMeta
}

func (t *ThirdPartyResources) Create(ctx context.Context, resource, namespace string, req, resp interface{}) error {
	ns := t.c.namespaceFor(namespace)
	if err := checkResource(t.apiGroup, t.apiVersion, resource, ns, "not required"); err != nil {
		return err
	}
	url := t.c.urlFor(t.apiGroup, t.apiVersion, ns, resource, "")
	return t.c.create(ctx, jsonCodec, url, req, resp)
}

func (t *ThirdPartyResources) Get(ctx context.Context, resource, namespace, name string, resp interface{}) error {
	ns := t.c.namespaceFor(namespace)
	if err := checkResource(t.apiGroup, t.apiVersion, resource, ns, name); err != nil {
		return err
	}
	url := t.c.urlFor(t.apiGroup, t.apiVersion, ns, resource, name)
	return t.c.get(ctx, jsonCodec, url, resp)
}

func (t *ThirdPartyResources) Delete(ctx context.Context, resource, namespace, name string) error {
	ns := t.c.namespaceFor(namespace)
	if err := checkResource(t.apiGroup, t.apiVersion, resource, ns, name); err != nil {
		return err
	}
	url := t.c.urlFor(t.apiGroup, t.apiVersion, ns, resource, name)
	return t.c.delete(ctx, jsonCodec, url)
}

func (t *ThirdPartyResources) List(ctx context.Context, resource, namespace string, resp interface{}) error {
	ns := t.c.namespaceFor(namespace)
	if err := checkResource(t.apiGroup, t.apiVersion, resource, ns, "name not required"); err != nil {
		return err
	}
	url := t.c.urlFor(t.apiGroup, t.apiVersion, ns, resource, "")
	return t.c.get(ctx, jsonCodec, url, resp)
}
