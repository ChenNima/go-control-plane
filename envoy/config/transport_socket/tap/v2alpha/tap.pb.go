// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/transport_socket/tap/v2alpha/tap.proto

package envoy_config_transport_socket_tap_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2alpha "github.com/envoyproxy/go-control-plane/envoy/config/common/tap/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Tap struct {
	CommonConfig         *v2alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	TransportSocket      *core.TransportSocket          `protobuf:"bytes,2,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_07cb8c0b42756e40, []int{0}
}

func (m *Tap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tap.Unmarshal(m, b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return xxx_messageInfo_Tap.Size(m)
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v2alpha.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *Tap) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.config.transport_socket.tap.v2alpha.Tap")
}

func init() {
	proto.RegisterFile("envoy/config/transport_socket/tap/v2alpha/tap.proto", fileDescriptor_07cb8c0b42756e40)
}

var fileDescriptor_07cb8c0b42756e40 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcf, 0x4a, 0xf3, 0x40,
	0x10, 0x67, 0xfb, 0xd1, 0x8f, 0x12, 0x15, 0x4b, 0x2e, 0x96, 0x22, 0x22, 0x3d, 0x29, 0xc8, 0xae,
	0xa4, 0xa0, 0xf7, 0x16, 0xef, 0x45, 0x03, 0x1e, 0xcb, 0x34, 0x5d, 0xeb, 0x62, 0xb3, 0xb3, 0xec,
	0xae, 0xa1, 0xc5, 0x37, 0xf0, 0xe2, 0xd5, 0x17, 0xf2, 0x85, 0x3c, 0x7a, 0x10, 0xc9, 0x4e, 0x02,
	0x4d, 0xbc, 0x78, 0x1b, 0x66, 0x7e, 0x7f, 0x66, 0x7e, 0x13, 0x8d, 0xa5, 0x2e, 0x70, 0x2b, 0x32,
	0xd4, 0x0f, 0x6a, 0x25, 0xbc, 0x05, 0xed, 0x0c, 0x5a, 0x3f, 0x77, 0x98, 0x3d, 0x49, 0x2f, 0x3c,
	0x18, 0x51, 0x24, 0xb0, 0x36, 0x8f, 0x50, 0xd6, 0xdc, 0x58, 0xf4, 0x18, 0x9f, 0x07, 0x12, 0x27,
	0x12, 0x6f, 0x93, 0x78, 0x09, 0xac, 0x48, 0xc3, 0x63, 0xd2, 0x07, 0xa3, 0x44, 0x91, 0x88, 0x0c,
	0xad, 0x14, 0x0b, 0x70, 0x92, 0x84, 0x86, 0x17, 0x0d, 0xf7, 0x0c, 0xf3, 0x1c, 0x75, 0xc3, 0x93,
	0x5a, 0x15, 0xfa, 0xe4, 0x79, 0x69, 0x40, 0x80, 0xd6, 0xe8, 0xc1, 0x2b, 0xd4, 0x4e, 0xe4, 0x6a,
	0x65, 0xc1, 0xd7, 0x6a, 0x47, 0x05, 0xac, 0xd5, 0x12, 0xbc, 0x14, 0x75, 0x41, 0x83, 0xd1, 0x07,
	0x8b, 0xfe, 0xa5, 0x60, 0x62, 0x19, 0x1d, 0x90, 0xe0, 0x9c, 0x1c, 0x07, 0xec, 0x94, 0x9d, 0xed,
	0x25, 0x57, 0xbc, 0x71, 0x4f, 0xe5, 0xb9, 0x73, 0x05, 0x9f, 0x86, 0xd6, 0xcd, 0xc6, 0x4b, 0xed,
	0x14, 0xea, 0x69, 0x00, 0x4e, 0x7a, 0x5f, 0x93, 0xee, 0x2b, 0xeb, 0xf4, 0xd9, 0xed, 0x3e, 0x71,
	0xa8, 0x1f, 0xdf, 0x47, 0xfd, 0x76, 0x26, 0x83, 0x4e, 0x70, 0x1a, 0x55, 0x4e, 0x60, 0x14, 0x2f,
	0x12, 0x5e, 0xc6, 0xc1, 0xd3, 0x1a, 0x7a, 0x17, 0x90, 0x3b, 0xaa, 0x87, 0xbe, 0x35, 0x7a, 0xf9,
	0x7c, 0xff, 0x7e, 0xeb, 0x5e, 0xc6, 0x95, 0x8a, 0xac, 0xf7, 0x71, 0xbf, 0x7e, 0xe0, 0x68, 0xfd,
	0x71, 0x58, 0x3f, 0xba, 0x56, 0x48, 0x14, 0x63, 0x71, 0xb3, 0xe5, 0x7f, 0xfe, 0xde, 0xa4, 0x97,
	0x82, 0x99, 0x95, 0x11, 0xce, 0xd8, 0xe2, 0x7f, 0xc8, 0x72, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x22, 0x31, 0x3b, 0xc1, 0x32, 0x02, 0x00, 0x00,
}
