// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/listener/http_inspector/v2/http_inspector.proto

package envoy_config_filter_listener_http_inspector_v2

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

type HttpInspector struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpInspector) Reset()         { *m = HttpInspector{} }
func (m *HttpInspector) String() string { return proto.CompactTextString(m) }
func (*HttpInspector) ProtoMessage()    {}
func (*HttpInspector) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d6cdfe6d97a924e, []int{0}
}

func (m *HttpInspector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpInspector.Unmarshal(m, b)
}
func (m *HttpInspector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpInspector.Marshal(b, m, deterministic)
}
func (m *HttpInspector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpInspector.Merge(m, src)
}
func (m *HttpInspector) XXX_Size() int {
	return xxx_messageInfo_HttpInspector.Size(m)
}
func (m *HttpInspector) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpInspector.DiscardUnknown(m)
}

var xxx_messageInfo_HttpInspector proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HttpInspector)(nil), "envoy.config.filter.listener.http_inspector.v2.HttpInspector")
}

func init() {
	proto.RegisterFile("envoy/config/filter/listener/http_inspector/v2/http_inspector.proto", fileDescriptor_0d6cdfe6d97a924e)
}

var fileDescriptor_0d6cdfe6d97a924e = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0xcf, 0x28, 0x29, 0x29, 0x88, 0xcf, 0xcc,
	0x2b, 0x2e, 0x48, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0x2f, 0x33, 0x42, 0x13, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xd2, 0x03, 0x1b, 0xa2, 0x07, 0x31, 0x44, 0x0f, 0x62, 0x88, 0x1e, 0xcc, 0x10,
	0x3d, 0x34, 0x2d, 0x65, 0x46, 0x52, 0x72, 0xa5, 0x29, 0x05, 0x89, 0xfa, 0x89, 0x79, 0x79, 0xf9,
	0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0xfa, 0xb9, 0x99, 0xe9, 0x45, 0x89, 0x25, 0xa9, 0x10,
	0xf3, 0x94, 0xf8, 0xb9, 0x78, 0x3d, 0x4a, 0x4a, 0x0a, 0x3c, 0x61, 0x7a, 0x9c, 0xa6, 0x30, 0x7e,
	0x9a, 0xf1, 0xaf, 0x9f, 0xd5, 0x4a, 0xc8, 0x02, 0x62, 0x53, 0x6a, 0x45, 0x49, 0x6a, 0x5e, 0x31,
	0x48, 0x27, 0xd4, 0xb6, 0x62, 0xdc, 0xd6, 0x19, 0x27, 0xe6, 0x14, 0x64, 0x24, 0x72, 0xd9, 0x64,
	0xe6, 0x43, 0x9c, 0x59, 0x50, 0x94, 0x5f, 0x51, 0x49, 0xa2, 0x8b, 0x9d, 0x84, 0x50, 0xdc, 0x13,
	0x00, 0x72, 0x65, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0xb9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb7, 0xf0, 0xe9, 0xc7, 0x45, 0x01, 0x00, 0x00,
}
