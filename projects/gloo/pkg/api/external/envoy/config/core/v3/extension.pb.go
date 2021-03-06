// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/extension.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Message type for extension configuration.
// [#next-major-version: revisit all existing typed_config that doesn't use this wrapper.].
type TypedExtensionConfig struct {
	// The name of an extension. This is not used to select the extension, instead
	// it serves the role of an opaque identifier.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The typed config for the extension. The type URL will be used to identify
	// the extension. In the case that the type URL is *udpa.type.v1.TypedStruct*,
	// the inner type URL of *TypedStruct* will be utilized. See the
	// :ref:`extension configuration overview
	// <config_overview_extension_configuration>` for further details.
	TypedConfig          *types.Any `protobuf:"bytes,2,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TypedExtensionConfig) Reset()         { *m = TypedExtensionConfig{} }
func (m *TypedExtensionConfig) String() string { return proto.CompactTextString(m) }
func (*TypedExtensionConfig) ProtoMessage()    {}
func (*TypedExtensionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf325a4a2e8752e, []int{0}
}
func (m *TypedExtensionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedExtensionConfig.Unmarshal(m, b)
}
func (m *TypedExtensionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedExtensionConfig.Marshal(b, m, deterministic)
}
func (m *TypedExtensionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedExtensionConfig.Merge(m, src)
}
func (m *TypedExtensionConfig) XXX_Size() int {
	return xxx_messageInfo_TypedExtensionConfig.Size(m)
}
func (m *TypedExtensionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedExtensionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TypedExtensionConfig proto.InternalMessageInfo

func (m *TypedExtensionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TypedExtensionConfig) GetTypedConfig() *types.Any {
	if m != nil {
		return m.TypedConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*TypedExtensionConfig)(nil), "envoy.config.core.v3.TypedExtensionConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/extension.proto", fileDescriptor_5bf325a4a2e8752e)
}

var fileDescriptor_5bf325a4a2e8752e = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4a, 0x3b, 0x31,
	0x14, 0xc6, 0xc9, 0xd0, 0x7f, 0xff, 0x75, 0x2a, 0x52, 0xca, 0x80, 0xb5, 0xa2, 0x94, 0xae, 0xba,
	0x31, 0x0f, 0xec, 0x09, 0x3a, 0xe2, 0xca, 0x4d, 0x29, 0x82, 0xe0, 0x46, 0xd2, 0x69, 0x1a, 0xa3,
	0xd3, 0xbc, 0x90, 0x64, 0x86, 0xce, 0xce, 0x1b, 0x78, 0x07, 0x57, 0x9e, 0xc1, 0x13, 0xb8, 0xf5,
	0x0a, 0x1e, 0xa3, 0x2b, 0x99, 0xa4, 0x75, 0x25, 0xe8, 0xee, 0x9b, 0xf7, 0x7d, 0xdf, 0xe4, 0xc7,
	0x7b, 0xf1, 0x8d, 0x90, 0xee, 0xbe, 0x98, 0xd3, 0x0c, 0x57, 0x60, 0x31, 0xc7, 0x33, 0x89, 0x20,
	0x72, 0x44, 0xd0, 0x06, 0x1f, 0x78, 0xe6, 0x6c, 0xf8, 0x62, 0x5a, 0x02, 0x5f, 0x3b, 0x6e, 0x14,
	0xcb, 0x81, 0xab, 0x12, 0x2b, 0xc8, 0x50, 0x2d, 0xa5, 0x80, 0x0c, 0x0d, 0x87, 0x72, 0xec, 0x5d,
	0x65, 0x25, 0x2a, 0xaa, 0x0d, 0x3a, 0xec, 0x26, 0x3e, 0x45, 0x43, 0x8a, 0xd6, 0x29, 0x5a, 0x8e,
	0xfb, 0x47, 0x02, 0x51, 0xe4, 0x1c, 0x7c, 0x66, 0x5e, 0x2c, 0x81, 0xa9, 0x2a, 0x14, 0xfa, 0x27,
	0xc5, 0x42, 0x33, 0x60, 0x4a, 0xa1, 0x63, 0x4e, 0xa2, 0xb2, 0x60, 0x1d, 0x73, 0x85, 0xdd, 0xda,
	0x87, 0x25, 0xcb, 0xe5, 0x82, 0x39, 0x0e, 0x3b, 0xb1, 0x35, 0x12, 0x81, 0x02, 0xbd, 0x84, 0x5a,
	0x85, 0xe9, 0xb0, 0x8c, 0x93, 0xeb, 0x4a, 0xf3, 0xc5, 0xe5, 0x0e, 0xeb, 0xc2, 0x83, 0x74, 0x8f,
	0xe3, 0x86, 0x62, 0x2b, 0xde, 0x23, 0x03, 0x32, 0xda, 0x4b, 0xff, 0x6f, 0xd2, 0x86, 0x89, 0x3a,
	0x64, 0xe6, 0x87, 0xdd, 0x49, 0xbc, 0xef, 0xea, 0xd2, 0x5d, 0xa0, 0xee, 0x45, 0x03, 0x32, 0x6a,
	0x9f, 0x27, 0x34, 0x40, 0xd3, 0x1d, 0x34, 0x9d, 0xa8, 0x2a, 0x6d, 0x6d, 0xd2, 0x7f, 0x2f, 0x24,
	0x6a, 0x91, 0x59, 0xdb, 0x77, 0xc2, 0xff, 0xd3, 0x67, 0xf2, 0xfa, 0x79, 0x4a, 0xde, 0x9e, 0xde,
	0x3f, 0x9a, 0x51, 0x27, 0x8a, 0x87, 0x12, 0xa9, 0x5f, 0x84, 0x36, 0xb8, 0xae, 0xe8, 0x4f, 0x3b,
	0x49, 0x0f, 0xbe, 0x19, 0xa7, 0xf5, 0x03, 0x53, 0x72, 0x7b, 0xf5, 0xb7, 0xa3, 0xe8, 0x47, 0xf1,
	0xfb, 0x61, 0xe6, 0x4d, 0x8f, 0x3d, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xc6, 0x58, 0xd8,
	0xea, 0x01, 0x00, 0x00,
}

func (this *TypedExtensionConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TypedExtensionConfig)
	if !ok {
		that2, ok := that.(TypedExtensionConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.TypedConfig.Equal(that1.TypedConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
