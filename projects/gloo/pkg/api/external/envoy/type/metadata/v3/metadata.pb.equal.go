// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/metadata/v3/metadata.proto

package v3

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
)

// Equal function
func (m *MetadataKey) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MetadataKey)
	if !ok {
		that2, ok := that.(MetadataKey)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if len(m.GetPath()) != len(target.GetPath()) {
		return false
	}
	for idx, v := range m.GetPath() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPath()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPath()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *MetadataKind) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MetadataKind)
	if !ok {
		that2, ok := that.(MetadataKind)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Kind.(type) {

	case *MetadataKind_Request_:

		if h, ok := interface{}(m.GetRequest()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequest()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRequest(), target.GetRequest()) {
				return false
			}
		}

	case *MetadataKind_Route_:

		if h, ok := interface{}(m.GetRoute()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRoute()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRoute(), target.GetRoute()) {
				return false
			}
		}

	case *MetadataKind_Cluster_:

		if h, ok := interface{}(m.GetCluster()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCluster()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCluster(), target.GetCluster()) {
				return false
			}
		}

	case *MetadataKind_Host_:

		if h, ok := interface{}(m.GetHost()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHost()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHost(), target.GetHost()) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *MetadataKey_PathSegment) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MetadataKey_PathSegment)
	if !ok {
		that2, ok := that.(MetadataKey_PathSegment)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Segment.(type) {

	case *MetadataKey_PathSegment_Key:

		if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *MetadataKind_Request) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MetadataKind_Request)
	if !ok {
		that2, ok := that.(MetadataKind_Request)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *MetadataKind_Route) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MetadataKind_Route)
	if !ok {
		that2, ok := that.(MetadataKind_Route)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *MetadataKind_Cluster) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MetadataKind_Cluster)
	if !ok {
		that2, ok := that.(MetadataKind_Cluster)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *MetadataKind_Host) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MetadataKind_Host)
	if !ok {
		that2, ok := that.(MetadataKind_Host)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}
