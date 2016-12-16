package k8s

import "github.com/ericchiang/k8s/api/v1"

type MyTRP struct {
	v1.ObjectMeta `json:"metadata,omitempty"`

	MyValue string `json:"myValue"`
}

type myTRP struct {
}
