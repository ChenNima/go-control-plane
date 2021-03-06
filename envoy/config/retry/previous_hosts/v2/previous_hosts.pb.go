// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/retry/previous_hosts/v2/previous_hosts.proto

package envoy_config_retry_previous_hosts_v2

import (
	fmt "fmt"
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

type PreviousHostsPredicate struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreviousHostsPredicate) Reset()         { *m = PreviousHostsPredicate{} }
func (m *PreviousHostsPredicate) String() string { return proto.CompactTextString(m) }
func (*PreviousHostsPredicate) ProtoMessage()    {}
func (*PreviousHostsPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4d889d3482108fe, []int{0}
}

func (m *PreviousHostsPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreviousHostsPredicate.Unmarshal(m, b)
}
func (m *PreviousHostsPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreviousHostsPredicate.Marshal(b, m, deterministic)
}
func (m *PreviousHostsPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousHostsPredicate.Merge(m, src)
}
func (m *PreviousHostsPredicate) XXX_Size() int {
	return xxx_messageInfo_PreviousHostsPredicate.Size(m)
}
func (m *PreviousHostsPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousHostsPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousHostsPredicate proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PreviousHostsPredicate)(nil), "envoy.config.retry.previous_hosts.v2.PreviousHostsPredicate")
}

func init() {
	proto.RegisterFile("envoy/config/retry/previous_hosts/v2/previous_hosts.proto", fileDescriptor_e4d889d3482108fe)
}

var fileDescriptor_e4d889d3482108fe = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0x29, 0xaa, 0xd4, 0x2f,
	0x28, 0x4a, 0x2d, 0xcb, 0xcc, 0x2f, 0x2d, 0x8e, 0xcf, 0xc8, 0x2f, 0x2e, 0x29, 0xd6, 0x2f, 0x33,
	0x42, 0x13, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x01, 0x6b, 0xd5, 0x83, 0x68, 0xd5,
	0x03, 0x6b, 0xd5, 0x43, 0x53, 0x58, 0x66, 0xa4, 0x24, 0xc1, 0x25, 0x16, 0x00, 0x15, 0xf4, 0x00,
	0x89, 0x05, 0x14, 0xa5, 0xa6, 0x64, 0x26, 0x27, 0x96, 0xa4, 0x3a, 0x79, 0x71, 0x19, 0x65, 0xe6,
	0xeb, 0x81, 0x0d, 0x29, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0xc6, 0x3c, 0x27, 0x21, 0x34, 0xd3,
	0xf2, 0x4b, 0xf2, 0x03, 0x18, 0x93, 0xd8, 0xc0, 0x4e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x90, 0x70, 0xbc, 0xba, 0xcf, 0x00, 0x00, 0x00,
}
