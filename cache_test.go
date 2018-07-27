package k8s

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"

	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

func TestUnmarshalRawResourceJSON(t *testing.T) {
	data := []byte(`{
		"metadata": {
			"name": "my-configmap",
			"namespace": "my-namespace"
        },
        "data": {
			"hello": "world"
        }
	}`)

	want := &metav1.ObjectMeta{
		Name:      String("my-configmap"),
		Namespace: String("my-namespace"),
	}

	var raw rawResource
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("failed to unmarshal struct: %v", err)
	}

	got := raw.GetMetadata()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want= %#v", want)
		t.Errorf("got=  %#v", got)
	}
}

func TestUnmarshalRawResouceListJSON(t *testing.T) {
	data := []byte(`{
		"metadata": {
			"resourceVersion": "12345"
		},
		"items": [
			{
				"metadata": {
					"name": "my-configmap",
					"namespace": "my-namespace"
				},
				"data": {
					"hello": "world"
				}
			},
			{
				"metadata": {
					"name": "my-configmap-2",
					"namespace": "my-namespace"
				},
				"data": {
					"hello": "world"
				}
			}
		]
	}`)

	want := &metav1.ListMeta{
		ResourceVersion: String("12345"),
	}

	var raw rawResourceList
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("failed to unmarshal struct: %v", err)
	}

	got := raw.GetMetadata()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want= %#v", want)
		t.Errorf("got=  %#v", got)
	}
}

func TestUnmarshalRawResourceProto(t *testing.T) {
	// cm := &v1.ConfigMap{
	//         Metadata: &metav1.ObjectMeta{
	//                 Name:      k8s.String("my-configmap"),
	//                 Namespace: k8s.String("my-namespace"),
	//         },
	//         Data: map[string]string{"hello": "world"},
	// }
	// data, _ := proto.Marshal(cm)
	// fmt.Println(base64.StdEncoding.EncodeToString(data))
	data, _ := base64.StdEncoding.DecodeString(
		"ChwKDG15LWNvbmZpZ21hcBoMbXktbmFtZXNwYWNlEg4KBWhlbGxvEgV3b3JsZA==",
	)

	want := &metav1.ObjectMeta{
		Name:      String("my-configmap"),
		Namespace: String("my-namespace"),
	}

	var raw rawResource
	if err := proto.Unmarshal(data, &raw); err != nil {
		t.Fatalf("failed to unmarshal struct: %v", err)
	}

	got := raw.GetMetadata()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want= %#v", want)
		t.Errorf("got=  %#v", got)
	}
}

func TestUnmarshalRawResourceListProto(t *testing.T) {
	// list := &v1.ConfigMapList{
	//         Metadata: &metav1.ListMeta{
	//                 ResourceVersion: k8s.String("12345"),
	//         },
	//         Items: []*v1.ConfigMap{
	//                 {
	//                         Metadata: &metav1.ObjectMeta{
	//                                 Name:      k8s.String("my-configmap"),
	//                                 Namespace: k8s.String("my-namespace"),
	//                         },
	//                         Data: map[string]string{"hello": "world"},
	//                 },
	//                 {
	//                         Metadata: &metav1.ObjectMeta{
	//                                 Name:      k8s.String("my-configmap-2"),
	//                                 Namespace: k8s.String("my-namespace"),
	//                         },
	//                         Data: map[string]string{"hello": "world"},
	//                 },
	//         },
	// }
	// data, _ := proto.Marshal(list)
	// fmt.Println(base64.StdEncoding.EncodeToString(data))
	data, _ := base64.StdEncoding.DecodeString(
		"CgcSBTEyMzQ1Ei4KHAoMbXktY29uZmlnbWFwGgxteS1uYW1lc3BhY2USDgoFaGVsbG8SBXdvcmxkEjAKHgoObXktY29uZmlnbWFwLTIaDG15LW5hbWVzcGFjZRIOCgVoZWxsbxIFd29ybGQ=",
	)

	want := &metav1.ListMeta{
		ResourceVersion: String("12345"),
	}

	var raw rawResourceList
	if err := proto.Unmarshal(data, &raw); err != nil {
		t.Fatalf("failed to unmarshal struct: %v", err)
	}

	got := raw.GetMetadata()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want= %#v", want)
		t.Errorf("got=  %#v", got)
	}
}
