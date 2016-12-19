/*
Package k8s implements a Kubernetes client.
*/
package k8s

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/ericchiang/k8s/api/unversioned"
)

// String returns a pointer to a string. Useful for creating API objects
// that take pointers instead of literals.
//
//		cm := &v1.ConfigMap{
//			Metadata: &v1.ObjectMeta{
//				Name:      k8s.String("myconfigmap"),
//				Namespace: k8s.String("default"),
//			},
//			Data: map[string]string{
//				"foo": "bar",
//			},
//		}
//
func String(s string) *string { return &s }
func Int(i int) *int          { return &i }
func Bool(b bool) *bool       { return &b }

// Client is a Kuberntes client.
type Client struct {
	// The URL of the API server.
	Endpoint string

	// Default namespaces for objects that don't supply a namespace in
	// their object metadata.
	Namespace string

	Client *http.Client
}

// Option represents optional call parameters, such as label selectors.
type Option interface {
	queryParam() (key, val string)
}

type resourceVersionOption string

func (r resourceVersionOption) queryParam() (string, string) {
	return "resourceVersion", string(r)
}

// ResourceVersion causes watch operations to only show changes since
// a particular version of a resource.
func ResourceVersion(resourceVersion string) Option {
	return resourceVersionOption(resourceVersion)
}

type timeoutSeconds string

func (t timeoutSeconds) queryParam() (string, string) {
	return "timeoutSeconds", string(t)
}

// Timeout declares the timeout for list and watch operations. Timeout
// is only accurate to the second.
func Timeout(d time.Duration) Option {
	return timeoutSeconds(strconv.FormatInt(int64(d/time.Second), 10))
}

// NewClient initializes a client from a client config.
func NewClient(config *Config) (*Client, error) {
	if len(config.Contexts) == 0 {
		if config.CurrentContext != "" {
			return nil, fmt.Errorf("no contexts with name %q", config.CurrentContext)
		}

		if n := len(config.Clusters); n == 0 {
			return nil, errors.New("no clusters provided")
		} else if n > 1 {
			return nil, errors.New("multiple clusters but no current context")
		}
		if n := len(config.AuthInfos); n == 0 {
			return nil, errors.New("no users provided")
		} else if n > 1 {
			return nil, errors.New("multiple users but no current context")
		}

		return newClient(config.Clusters[0].Cluster, config.AuthInfos[0].AuthInfo, "")
	}

	var ctx Context
	if config.CurrentContext == "" {
		if n := len(config.Contexts); n == 0 {
			return nil, errors.New("no contexts provided")
		} else if n > 1 {
			return nil, errors.New("multiple contexts but no current context")
		}
		ctx = config.Contexts[0].Context
	} else {
		for _, c := range config.Contexts {
			if c.Name == config.CurrentContext {
				ctx = c.Context
				goto configFound
			}
		}
		return nil, fmt.Errorf("no config named %q", config.CurrentContext)
	configFound:
	}

	if ctx.Cluster == "" {
		return nil, fmt.Errorf("context doesn't have a cluster")
	}
	if ctx.AuthInfo == "" {
		return nil, fmt.Errorf("context doesn't have a user")
	}
	var (
		user    AuthInfo
		cluster Cluster
	)

	for _, u := range config.AuthInfos {
		if u.Name == ctx.AuthInfo {
			user = u.AuthInfo
			goto userFound
		}
	}
	return nil, fmt.Errorf("no user named %q", ctx.AuthInfo)
userFound:

	for _, c := range config.Clusters {
		if c.Name == ctx.Cluster {
			cluster = c.Cluster
			goto clusterFound
		}
	}
	return nil, fmt.Errorf("no cluster named %q", ctx.Cluster)
clusterFound:

	return newClient(cluster, user, ctx.Namespace)
}

// NewInClusterClient returns a client that uses the service account bearer token mounted
// into Kubernetes pods.
func NewInClusterClient() (*Client, error) {
	host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")
	if len(host) == 0 || len(port) == 0 {
		return nil, errors.New("unable to load in-cluster configuration, KUBERNETES_SERVICE_HOST and KUBERNETES_SERVICE_PORT must be defined")
	}
	namespace, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		return nil, err
	}

	cluster := Cluster{
		Server:               "https://" + host + ":" + port,
		CertificateAuthority: "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt",
	}
	user := AuthInfo{TokenFile: "/var/run/secrets/kubernetes.io/serviceaccount/token"}
	return newClient(cluster, user, string(namespace))
}

func load(filepath string, data []byte) (out []byte, err error) {
	if filepath != "" {
		data, err = ioutil.ReadFile(filepath)
	}
	return data, err
}

func newClient(cluster Cluster, user AuthInfo, namespace string) (*Client, error) {
	if cluster.Server == "" {
		// NOTE: kubectl defaults to localhost:8080, but it's probably better to just
		// be strict.
		return nil, fmt.Errorf("no cluster endpoint provided")
	}

	ca, err := load(cluster.CertificateAuthority, cluster.CertificateAuthorityData)
	if err != nil {
		return nil, fmt.Errorf("loading certificate authority: %v", err)
	}

	clientCert, err := load(user.ClientCertificate, user.ClientCertificateData)
	if err != nil {
		return nil, fmt.Errorf("load client cert: %v", err)
	}
	clientKey, err := load(user.ClientKey, user.ClientKeyData)
	if err != nil {
		return nil, fmt.Errorf("load client cert: %v", err)
	}

	// See https://github.com/gtank/cryptopasta
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	if len(ca) != 0 {
		tlsConfig.RootCAs = x509.NewCertPool()
		if !tlsConfig.RootCAs.AppendCertsFromPEM(ca) {
			return nil, errors.New("certificate authority doesn't contain any certificates")
		}
	}
	if len(clientCert) != 0 {
		cert, err := tls.X509KeyPair(clientCert, clientKey)
		if err != nil {
			return nil, fmt.Errorf("invalid client cert and key pair: %v", err)
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}

	token := user.Token
	if user.TokenFile != "" {
		data, err := ioutil.ReadFile(user.TokenFile)
		if err != nil {
			return nil, fmt.Errorf("load token file: %v", err)
		}
		token = string(data)
	}

	var transport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSClientConfig:       tlsConfig,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if token != "" {
		transport = &bearerTokenTransport{transport, token}
	}
	if user.Username != "" && user.Password != "" {
		transport = &basicAuthTransport{transport, user.Username, user.Password}
	}

	return &Client{
		Endpoint:  cluster.Server,
		Namespace: namespace,
		Client:    &http.Client{Transport: transport},
	}, nil
}

// copyReq creates a shallow copy of an http.Request while.
func copyReq(req *http.Request) *http.Request {
	r := new(http.Request)
	*r = *req
	r.Header = make(http.Header, len(req.Header)+1) // assume that a header will be added.
	for k, s := range req.Header {
		r.Header[k] = append([]string(nil), s...)
	}
	return r
}

type bearerTokenTransport struct {
	base  http.RoundTripper
	token string
}

func (t *bearerTokenTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	r := copyReq(req)
	r.Header.Set("Authorization", "Bearer "+t.token)
	return t.base.RoundTrip(r)
}

type basicAuthTransport struct {
	base     http.RoundTripper
	username string
	password string
}

func (t *basicAuthTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	r := copyReq(req)
	r.SetBasicAuth(t.username, t.password)
	return t.base.RoundTrip(r)
}

// APIError is an error from a unexpected status code.
type APIError struct {
	// The status object returned by the Kubernetes API,
	Status *unversioned.Status
}

func (e *APIError) Error() string { return fmt.Sprintf("kubernetes api: %s", e.Status.Message) }

func checkStatusCode(c *codec, statusCode int, body []byte) error {
	if statusCode/100 == 2 {
		return nil
	}

	status := new(unversioned.Status)
	if err := c.unmarshal(body, status); err != nil {
		return fmt.Errorf("decode error status: %v", err)
	}
	return &APIError{status}
}

func (c *Client) client() *http.Client {
	if c.Client == nil {
		return http.DefaultClient
	}
	return c.Client
}

func (c *Client) namespaceFor(namespace string) string {
	if namespace != "" {
		return namespace
	}
	return c.Namespace
}

// The following methods hold the logic for interacting with the Kubernetes API. Generated
// clients are thin wrappers on top of these methods.
//
// This client implements specs in the "API Conventions" developer document, which can be
// found here:
//
//   https://github.com/kubernetes/kubernetes/blob/master/docs/devel/api-conventions.md

func (c *Client) urlFor(apiGroup, apiVersion, namespace, resource, name string, options ...Option) string {
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
	endpoint := ""
	if strings.HasSuffix(c.Endpoint, "/") {
		endpoint = c.Endpoint + p
	} else {
		endpoint = c.Endpoint + "/" + p
	}
	if len(options) == 0 {
		return endpoint
	}

	v := url.Values{}
	for _, option := range options {
		key, val := option.queryParam()
		v.Set(key, val)
	}
	return endpoint + "?" + v.Encode()
}

func (c *Client) create(ctx context.Context, codec *codec, verb, url string, req, resp interface{}) error {
	body, err := codec.marshal(req)
	if err != nil {
		return err
	}

	r, err := http.NewRequest(verb, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", codec.contentType)
	r.Header.Set("Accept", codec.contentType)

	re, err := c.client().Do(r)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	respBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return fmt.Errorf("read body: %v", err)
	}

	if err := checkStatusCode(codec, re.StatusCode, respBody); err != nil {
		return err
	}
	return codec.unmarshal(respBody, resp)
}

func (c *Client) delete(ctx context.Context, codec *codec, url string) error {
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	r.Header.Set("Accept", codec.contentType)
	re, err := c.client().Do(r)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	respBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return fmt.Errorf("read body: %v", err)
	}

	if err := checkStatusCode(codec, re.StatusCode, respBody); err != nil {
		return err
	}
	return nil
}

// get can be used to either get or list a given resource.
func (c *Client) get(ctx context.Context, codec *codec, url string, resp interface{}) error {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	r.Header.Set("Accept", codec.contentType)
	re, err := c.client().Do(r)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	respBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return fmt.Errorf("read body: %v", err)
	}

	if err := checkStatusCode(codec, re.StatusCode, respBody); err != nil {
		return err
	}
	return codec.unmarshal(respBody, resp)
}
