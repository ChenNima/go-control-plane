// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/fault/v3alpha/fault.proto

package envoy_extensions_filters_http_fault_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha2 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3alpha"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/common/fault/v3alpha"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type FaultAbort struct {
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType            isFaultAbort_ErrorType     `protobuf_oneof:"error_type"`
	Percentage           *v3alpha.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_6343d697c6c7821d, []int{0}
}

func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort.Unmarshal(m, b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort.Size(m)
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetPercentage() *v3alpha.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

type HTTPFault struct {
	Delay                           *v3alpha1.FaultDelay      `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	Abort                           *FaultAbort               `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	UpstreamCluster                 string                    `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	Headers                         []*v3alpha2.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	DownstreamNodes                 []string                  `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	MaxActiveFaults                 *wrappers.UInt32Value     `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	ResponseRateLimit               *v3alpha1.FaultRateLimit  `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	DelayPercentRuntime             string                    `protobuf:"bytes,8,opt,name=delay_percent_runtime,json=delayPercentRuntime,proto3" json:"delay_percent_runtime,omitempty"`
	AbortPercentRuntime             string                    `protobuf:"bytes,9,opt,name=abort_percent_runtime,json=abortPercentRuntime,proto3" json:"abort_percent_runtime,omitempty"`
	DelayDurationRuntime            string                    `protobuf:"bytes,10,opt,name=delay_duration_runtime,json=delayDurationRuntime,proto3" json:"delay_duration_runtime,omitempty"`
	AbortHttpStatusRuntime          string                    `protobuf:"bytes,11,opt,name=abort_http_status_runtime,json=abortHttpStatusRuntime,proto3" json:"abort_http_status_runtime,omitempty"`
	MaxActiveFaultsRuntime          string                    `protobuf:"bytes,12,opt,name=max_active_faults_runtime,json=maxActiveFaultsRuntime,proto3" json:"max_active_faults_runtime,omitempty"`
	ResponseRateLimitPercentRuntime string                    `protobuf:"bytes,13,opt,name=response_rate_limit_percent_runtime,json=responseRateLimitPercentRuntime,proto3" json:"response_rate_limit_percent_runtime,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                  `json:"-"`
	XXX_unrecognized                []byte                    `json:"-"`
	XXX_sizecache                   int32                     `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_6343d697c6c7821d, []int{1}
}

func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFault.Unmarshal(m, b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return xxx_messageInfo_HTTPFault.Size(m)
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v3alpha1.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*v3alpha2.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v3alpha1.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func (m *HTTPFault) GetDelayPercentRuntime() string {
	if m != nil {
		return m.DelayPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortPercentRuntime() string {
	if m != nil {
		return m.AbortPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetDelayDurationRuntime() string {
	if m != nil {
		return m.DelayDurationRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortHttpStatusRuntime() string {
	if m != nil {
		return m.AbortHttpStatusRuntime
	}
	return ""
}

func (m *HTTPFault) GetMaxActiveFaultsRuntime() string {
	if m != nil {
		return m.MaxActiveFaultsRuntime
	}
	return ""
}

func (m *HTTPFault) GetResponseRateLimitPercentRuntime() string {
	if m != nil {
		return m.ResponseRateLimitPercentRuntime
	}
	return ""
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.extensions.filters.http.fault.v3alpha.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.extensions.filters.http.fault.v3alpha.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/fault/v3alpha/fault.proto", fileDescriptor_6343d697c6c7821d)
}

var fileDescriptor_6343d697c6c7821d = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xaf, 0x9b, 0x7e, 0x4e, 0x6e, 0xd5, 0xd6, 0xbd, 0xb7, 0xd7, 0xb7, 0x42, 0x25, 0x14,
	0x21, 0xa5, 0x14, 0xc6, 0x22, 0x41, 0xaa, 0x52, 0x89, 0x45, 0xd3, 0x52, 0xa5, 0xa8, 0x85, 0xc8,
	0x14, 0xb6, 0xd6, 0x34, 0x9e, 0x24, 0x96, 0xec, 0x19, 0x6b, 0x66, 0x9c, 0x26, 0x6f, 0xc0, 0x0b,
	0xb0, 0xe1, 0x7d, 0x90, 0xd8, 0xb2, 0xe3, 0x51, 0x10, 0x2b, 0x34, 0x67, 0x6c, 0x27, 0x69, 0xa8,
	0x50, 0xd9, 0xc5, 0x73, 0xce, 0xef, 0x9c, 0xff, 0xf9, 0x0a, 0x3a, 0xa0, 0x6c, 0xc0, 0x47, 0x2e,
	0x1d, 0x2a, 0xca, 0x64, 0xc8, 0x99, 0x74, 0xbb, 0x61, 0xa4, 0xa8, 0x90, 0x6e, 0x5f, 0xa9, 0xc4,
	0xed, 0x92, 0x34, 0x52, 0xee, 0xa0, 0x4e, 0xa2, 0xa4, 0x4f, 0xcc, 0x17, 0x4e, 0x04, 0x57, 0xdc,
	0xde, 0x07, 0x10, 0x8f, 0x41, 0x9c, 0x81, 0x58, 0x83, 0xd8, 0xb8, 0x66, 0xe0, 0xf6, 0x33, 0x93,
	0xa5, 0xc3, 0x59, 0x37, 0xec, 0xb9, 0x82, 0xa7, 0x8a, 0x16, 0x41, 0xe1, 0xcb, 0xef, 0xf0, 0x38,
	0xe1, 0x8c, 0x32, 0x25, 0x4d, 0xfc, 0xed, 0xc6, 0xad, 0xc2, 0x3a, 0x3c, 0x8e, 0x39, 0xbb, 0x5d,
	0xda, 0x76, 0xc5, 0xa0, 0x6a, 0x94, 0x8c, 0xb3, 0x24, 0x54, 0x74, 0x28, 0xcb, 0x3d, 0x76, 0x7a,
	0x9c, 0xf7, 0x22, 0xea, 0xc2, 0xd7, 0x55, 0xda, 0x75, 0xaf, 0x05, 0x49, 0x12, 0x2d, 0xde, 0xd8,
	0x1f, 0xa4, 0x41, 0x42, 0x5c, 0xc2, 0x18, 0x57, 0x44, 0x41, 0xf2, 0x01, 0x15, 0x5a, 0x45, 0xc8,
	0x7a, 0x99, 0xcb, 0x7f, 0x03, 0x12, 0x85, 0x01, 0xd1, 0x85, 0x64, 0x3f, 0x8c, 0x61, 0xf7, 0xab,
	0x85, 0xd0, 0xa9, 0x56, 0x73, 0x74, 0xc5, 0x85, 0xb2, 0x31, 0x2a, 0xeb, 0x86, 0xf8, 0x52, 0x11,
	0x95, 0x4a, 0x67, 0xae, 0x62, 0x55, 0x57, 0x9b, 0xe5, 0x1f, 0xcd, 0xe5, 0xc7, 0x8b, 0xeb, 0xdf,
	0xe6, 0xab, 0x5f, 0xac, 0xd6, 0x5f, 0x1e, 0xd2, 0x1e, 0x6f, 0xc1, 0xc1, 0x7e, 0x89, 0x50, 0xa6,
	0x95, 0xf4, 0xa8, 0x53, 0xaa, 0x58, 0xd5, 0x72, 0xed, 0x11, 0x36, 0xcd, 0xd6, 0x15, 0xe5, 0x3d,
	0xc5, 0xa7, 0x82, 0x74, 0xb4, 0x36, 0x12, 0xb5, 0x8d, 0xbf, 0x37, 0x01, 0x1e, 0xd6, 0x3f, 0x7d,
	0xfe, 0xb0, 0x83, 0xd1, 0x13, 0x03, 0x9a, 0xc6, 0x67, 0x13, 0x9a, 0x1a, 0x50, 0x0d, 0x8f, 0xb5,
	0x36, 0x37, 0x10, 0xa2, 0x42, 0x70, 0xe1, 0xeb, 0x44, 0x76, 0xe9, 0x7b, 0xd3, 0x7a, 0x35, 0xbf,
	0x6c, 0xad, 0xcf, 0xed, 0x7e, 0x5c, 0x42, 0x2b, 0xad, 0xcb, 0xcb, 0x36, 0xf8, 0xda, 0x6f, 0xd0,
	0x42, 0x40, 0x23, 0x32, 0x72, 0x2c, 0x50, 0xd7, 0xc0, 0xb7, 0xae, 0x82, 0x19, 0xd5, 0xf4, 0x32,
	0x98, 0x84, 0x27, 0x3a, 0x80, 0x67, 0xe2, 0xd8, 0x17, 0x68, 0x81, 0x68, 0x01, 0xd0, 0x9d, 0x72,
	0xed, 0x00, 0xdf, 0x61, 0xb7, 0x26, 0xf4, 0x7b, 0x26, 0x8a, 0xbd, 0x87, 0xd6, 0xd3, 0x44, 0x2a,
	0x41, 0x49, 0xec, 0x77, 0xa2, 0x54, 0x2a, 0x2a, 0xa0, 0x91, 0x2b, 0xde, 0x5a, 0xfe, 0x7e, 0x6c,
	0x9e, 0xed, 0x63, 0xb4, 0xd4, 0xa7, 0x24, 0xa0, 0x42, 0x3a, 0xf3, 0x95, 0x52, 0xb5, 0x5c, 0xdb,
	0xc3, 0x53, 0x1d, 0x83, 0xe5, 0x2c, 0x52, 0xb5, 0xc0, 0xf5, 0x82, 0xa8, 0x4e, 0x9f, 0x0a, 0x2f,
	0x27, 0x75, 0xbe, 0x80, 0x5f, 0xb3, 0x2c, 0x23, 0xe3, 0x01, 0x95, 0xce, 0x42, 0xa5, 0xa4, 0xf3,
	0x8d, 0xdf, 0x5f, 0xeb, 0x67, 0xbb, 0x85, 0x36, 0x62, 0x32, 0xf4, 0xf5, 0xe4, 0x06, 0xd4, 0x87,
	0x42, 0xa4, 0xb3, 0x08, 0x55, 0xdf, 0xc3, 0x66, 0x29, 0x71, 0xbe, 0x94, 0xf8, 0xdd, 0x19, 0x53,
	0xf5, 0xda, 0x7b, 0x12, 0xa5, 0xd4, 0x5b, 0x8b, 0xc9, 0xf0, 0x08, 0x28, 0xa8, 0x57, 0xda, 0x31,
	0xda, 0x14, 0x54, 0x26, 0x9c, 0x49, 0xea, 0x0b, 0xa2, 0xa8, 0x1f, 0x85, 0x71, 0xa8, 0x9c, 0x25,
	0x88, 0xf5, 0xe2, 0x4f, 0x46, 0xe2, 0x11, 0x45, 0xcf, 0x75, 0x10, 0x6f, 0x23, 0x8f, 0x5c, 0x3c,
	0xd9, 0x35, 0xf4, 0x2f, 0xcc, 0xca, 0xcf, 0x76, 0xcc, 0x17, 0x29, 0x53, 0x61, 0x4c, 0x9d, 0x65,
	0x68, 0xec, 0x26, 0x18, 0xf3, 0x45, 0x34, 0x26, 0xcd, 0xc0, 0x40, 0x66, 0x98, 0x15, 0xc3, 0x80,
	0xf1, 0x06, 0xf3, 0x1c, 0x6d, 0x99, 0x3c, 0x41, 0x2a, 0xe0, 0xf4, 0x0a, 0x08, 0x01, 0xf4, 0x0f,
	0x58, 0x4f, 0x32, 0x63, 0x4e, 0x35, 0xd0, 0xff, 0x26, 0xd3, 0xc4, 0xa9, 0x15, 0x60, 0x19, 0xc0,
	0x2d, 0x70, 0x68, 0x15, 0x87, 0x36, 0x81, 0xce, 0x4c, 0xa4, 0x40, 0xff, 0x36, 0xe8, 0x8d, 0xde,
	0xe7, 0xe8, 0x39, 0x7a, 0xf8, 0x8b, 0x11, 0xcc, 0x54, 0xbb, 0x0a, 0x41, 0xee, 0xcf, 0xf4, 0x74,
	0xba, 0xf2, 0xc3, 0x9a, 0xbe, 0xd8, 0xa7, 0x68, 0xff, 0xf7, 0x17, 0x5b, 0x5c, 0x62, 0xf3, 0x0c,
	0x35, 0x42, 0x6e, 0x66, 0x9d, 0x08, 0x3e, 0x1c, 0xdd, 0xe5, 0x70, 0x9a, 0xe6, 0x5f, 0xaa, 0xad,
	0xb7, 0xad, 0x6d, 0x5d, 0x2d, 0xc2, 0xda, 0xd5, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xc0,
	0xaf, 0x2c, 0x11, 0x06, 0x00, 0x00,
}