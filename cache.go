package k8s

import (
	"encoding/json"
	"io"

	"github.com/golang/protobuf/proto"

	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

type rawMeta struct {
	Meta *metav1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
}

func (r *rawMeta) ProtoMessage()  {}
func (r *rawMeta) String() string { return "rawMeta" }
func (r *rawMeta) Reset()         { *r = rawMeta{} }

// rawResource is the an instance of a resource that has defered parsing of the
// body until later.
type rawResource struct {
	name       string
	namespace  string
	apiVersion string
	apiGroup   string

	contentType string // contentType of the payload, protobuf or JSON
	payload     []byte // protobuf or JSON encoded resource
	meta        *metav1.ObjectMeta
}

func (r *rawResource) GetMetadata() *metav1.ObjectMeta { return r.meta }
func (r *rawResource) ProtoMessage()                   {}
func (r *rawResource) String() string                  { return "rawResource" }
func (r *rawResource) Reset()                          { *r = rawResource{} }

func (r *rawResource) UnmarshalJSON(b []byte) error {
	var raw rawMeta
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	r.meta = raw.Meta
	r.payload = b
	r.contentType = contentTypeJSON
	return nil
}

func (r *rawResource) Unmarshal(b []byte) error {
	var raw rawMeta
	if err := proto.Unmarshal(b, &raw); err != nil {
		return err
	}
	r.meta = raw.Meta
	r.payload = b
	r.contentType = contentTypePB
	return nil
}

type rawListMeta struct {
	Meta  *metav1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Items []*rawResource   `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (r *rawListMeta) ProtoMessage()  {}
func (r *rawListMeta) String() string { return "rawMeta" }
func (r *rawListMeta) Reset()         { *r = rawListMeta{} }

type rawResourceList struct {
	namespace  string
	apiVersion string
	apiGroup   string

	contentType string
	items       []*rawResource
	meta        *metav1.ListMeta
}

func (r *rawResourceList) GetMetadata() *metav1.ListMeta { return r.meta }
func (r *rawResourceList) ProtoMessage()                 {}
func (r *rawResourceList) String() string                { return "rawResourceList" }
func (r *rawResourceList) Reset()                        { *r = rawResourceList{} }

func (r *rawResourceList) UnmarshalJSON(b []byte) error {
	var raw rawListMeta
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	r.meta = raw.Meta
	r.items = raw.Items
	r.contentType = contentTypeJSON
	return nil
}

func (r *rawResourceList) Unmarshal(b []byte) error {
	r.contentType = contentTypePB

	// proto.Unmarshal only respects proto.Unmarshaler if the top level type implements
	// the interface. Fields aren't checked.
	//
	// For example, the following code won't call rawResource's Unmarshal method, and
	// attempts to treat rawResource as a normal struct:
	//
	//		var message struct {
	//			Items []*rawResource `protobuf:"bytes,2,rep,name=items"`
	//		}
	//		proto.Unmarshal(b, &message)
	//
	// Instead, manually unpack each field then match on the field number.

	buff := proto.NewBuffer(b)

	for {
		// Read the next item.
		key, err := buff.DecodeVarint()
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				return nil
			}
			return err
		}

		data, err := buff.DecodeRawBytes(true)
		if err != nil {
			return err
		}

		// Lower three bits of key are the message type. The rest is the
		// field number.
		//
		// https://developers.google.com/protocol-buffers/docs/encoding#structure
		fieldNum := key >> 3

		switch fieldNum {
		case 1:
			var meta metav1.ListMeta
			if err := proto.Unmarshal(data, &meta); err != nil {
				return err
			}
			r.meta = &meta
		case 2:
			var raw rawResource
			if err := proto.Unmarshal(data, &raw); err != nil {
				return err
			}
			r.items = append(r.items, &raw)
		}
	}
	return nil
}
