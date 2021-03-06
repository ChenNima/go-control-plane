// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/listener/tls_inspector/v3alpha/tls_inspector.proto

package envoy_extensions_filters_listener_tls_inspector_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TlsInspector struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TlsInspector) Reset()         { *m = TlsInspector{} }
func (m *TlsInspector) String() string { return proto.CompactTextString(m) }
func (*TlsInspector) ProtoMessage()    {}
func (*TlsInspector) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f368c19e8e3815, []int{0}
}

func (m *TlsInspector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsInspector.Unmarshal(m, b)
}
func (m *TlsInspector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsInspector.Marshal(b, m, deterministic)
}
func (m *TlsInspector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsInspector.Merge(m, src)
}
func (m *TlsInspector) XXX_Size() int {
	return xxx_messageInfo_TlsInspector.Size(m)
}
func (m *TlsInspector) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsInspector.DiscardUnknown(m)
}

var xxx_messageInfo_TlsInspector proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TlsInspector)(nil), "envoy.extensions.filters.listener.tls_inspector.v3alpha.TlsInspector")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/listener/tls_inspector/v3alpha/tls_inspector.proto", fileDescriptor_c2f368c19e8e3815)
}

var fileDescriptor_c2f368c19e8e3815 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0x2f, 0xc9,
	0x29, 0x8e, 0xcf, 0xcc, 0x2b, 0x2e, 0x48, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0x2f, 0x33, 0x4e, 0xcc,
	0x29, 0xc8, 0x48, 0x44, 0x15, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x07, 0x1b, 0xa6,
	0x87, 0x30, 0x4c, 0x0f, 0x6a, 0x98, 0x1e, 0xcc, 0x30, 0x3d, 0x54, 0x6d, 0x50, 0xc3, 0xa4, 0x14,
	0x4b, 0x53, 0x0a, 0x12, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0xc0, 0xae, 0x28, 0x4b,
	0x2d, 0x02, 0x99, 0x90, 0x99, 0x97, 0x0e, 0x31, 0x5b, 0x29, 0x90, 0x8b, 0x27, 0x24, 0xa7, 0xd8,
	0x13, 0xa6, 0xd5, 0xca, 0x71, 0xd6, 0xd1, 0x0e, 0x39, 0x1b, 0x2e, 0x2b, 0x88, 0x95, 0xc9, 0xf9,
	0x79, 0x69, 0x99, 0xe9, 0x50, 0xeb, 0x70, 0xda, 0x66, 0xa4, 0x87, 0x6c, 0x84, 0x53, 0x0c, 0x97,
	0x6b, 0x66, 0xbe, 0x1e, 0xd8, 0x80, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0x32, 0x9d, 0xef, 0x24,
	0x88, 0x6c, 0x6c, 0x00, 0xc8, 0xb9, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0x77, 0x1b, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xb8, 0x2b, 0xaf, 0xf9, 0x62, 0x01, 0x00, 0x00,
}
