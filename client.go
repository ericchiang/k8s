package k8s

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/ericchiang/k8s/api/v1"
	"github.com/gogo/protobuf/proto"
)

const contentTypePB = "application/vnd.kubernetes.protobuf"

type Client struct {
	Endpoint  string
	Namespace string

	Client *http.Client
}

type HTTPError interface {
	error
	StatusCode() int
}

var _ HTTPError = (*httpError)(nil)

type httpError struct {
	statusCode int
	body       []byte
}

func (e *httpError) Error() string   { return string(e.body) }
func (e *httpError) StatusCode() int { return e.statusCode }

type object interface {
	GetMetadata() *v1.ObjectMeta
}

func (c *Client) client() *http.Client {
	if c.Client == nil {
		return http.DefaultClient
	}
	return c.Client
}

func (c *Client) namespaceFor(ns string, namespaced bool) string {
	if !namespaced {
		return ""
	}
	if ns != "" {
		return ns
	}
	if c.Namespace != "" {
		return c.Namespace
	}
	return "default"
}

func (c *Client) urlFor(apiGroup, apiVersion, namespace, resource, name string) string {
	basePath := "apis/"
	if apiGroup == "" {
		basePath = "api/"
	}

	var p string
	if namespace != "" {
		p = path.Join(basePath, apiGroup, apiVersion, "namespaces", namespace, resource, name)
	} else {
		p = path.Join(basePath, apiGroup, apiVersion, resource, name)
	}
	if strings.HasSuffix(c.Endpoint, "/") {
		return c.Endpoint + p
	}
	return c.Endpoint + "/" + p
}

func (c *Client) create(ctx context.Context, url string, req proto.Marshaler, resp proto.Unmarshaler) error {
	body, err := req.Marshal()
	if err != nil {
		return err
	}
	r, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", contentTypePB)
	r.Header.Set("Accept", contentTypePB)

	re, err := c.client().Do(r)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	respBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return fmt.Errorf("read body: %v", err)
	}

	if re.StatusCode != http.StatusOK {
		return &httpError{re.StatusCode, respBody}
	}
	return resp.Unmarshal(respBody)
}

func (c *Client) delete(ctx context.Context, url string, name string) error {
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	re, err := c.client().Do(r)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	respBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return fmt.Errorf("read body: %v", err)
	}

	if re.StatusCode != http.StatusOK {
		return &httpError{re.StatusCode, respBody}
	}
	return nil
}

func (c *Client) get(ctx context.Context, url string, resp proto.Unmarshaler) error {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	r.Header.Set("Accept", contentTypePB)
	re, err := c.client().Do(r)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	respBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return fmt.Errorf("read body: %v", err)
	}

	if re.StatusCode != http.StatusOK {
		return &httpError{re.StatusCode, respBody}
	}
	return resp.Unmarshal(respBody)
}
