// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/adaptive_concurrency/v3alpha/adaptive_concurrency.proto

package envoy_extensions_filters_http_adaptive_concurrency_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3alpha"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GradientControllerConfig struct {
	SampleAggregatePercentile *v3alpha.Percent                                            `protobuf:"bytes,1,opt,name=sample_aggregate_percentile,json=sampleAggregatePercentile,proto3" json:"sample_aggregate_percentile,omitempty"`
	ConcurrencyLimitParams    *GradientControllerConfig_ConcurrencyLimitCalculationParams `protobuf:"bytes,2,opt,name=concurrency_limit_params,json=concurrencyLimitParams,proto3" json:"concurrency_limit_params,omitempty"`
	MinRttCalcParams          *GradientControllerConfig_MinimumRTTCalculationParams       `protobuf:"bytes,3,opt,name=min_rtt_calc_params,json=minRttCalcParams,proto3" json:"min_rtt_calc_params,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                                                    `json:"-"`
	XXX_unrecognized          []byte                                                      `json:"-"`
	XXX_sizecache             int32                                                       `json:"-"`
}

func (m *GradientControllerConfig) Reset()         { *m = GradientControllerConfig{} }
func (m *GradientControllerConfig) String() string { return proto.CompactTextString(m) }
func (*GradientControllerConfig) ProtoMessage()    {}
func (*GradientControllerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_778e4cdafe4c14e9, []int{0}
}

func (m *GradientControllerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig.Unmarshal(m, b)
}
func (m *GradientControllerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig.Merge(m, src)
}
func (m *GradientControllerConfig) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig.Size(m)
}
func (m *GradientControllerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig proto.InternalMessageInfo

func (m *GradientControllerConfig) GetSampleAggregatePercentile() *v3alpha.Percent {
	if m != nil {
		return m.SampleAggregatePercentile
	}
	return nil
}

func (m *GradientControllerConfig) GetConcurrencyLimitParams() *GradientControllerConfig_ConcurrencyLimitCalculationParams {
	if m != nil {
		return m.ConcurrencyLimitParams
	}
	return nil
}

func (m *GradientControllerConfig) GetMinRttCalcParams() *GradientControllerConfig_MinimumRTTCalculationParams {
	if m != nil {
		return m.MinRttCalcParams
	}
	return nil
}

type GradientControllerConfig_ConcurrencyLimitCalculationParams struct {
	MaxConcurrencyLimit       *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_concurrency_limit,json=maxConcurrencyLimit,proto3" json:"max_concurrency_limit,omitempty"`
	ConcurrencyUpdateInterval *duration.Duration    `protobuf:"bytes,3,opt,name=concurrency_update_interval,json=concurrencyUpdateInterval,proto3" json:"concurrency_update_interval,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}              `json:"-"`
	XXX_unrecognized          []byte                `json:"-"`
	XXX_sizecache             int32                 `json:"-"`
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) Reset() {
	*m = GradientControllerConfig_ConcurrencyLimitCalculationParams{}
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) String() string {
	return proto.CompactTextString(m)
}
func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) ProtoMessage() {}
func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_778e4cdafe4c14e9, []int{0, 0}
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Unmarshal(m, b)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Merge(m, src)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Size(m)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams proto.InternalMessageInfo

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetMaxConcurrencyLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrencyLimit
	}
	return nil
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetConcurrencyUpdateInterval() *duration.Duration {
	if m != nil {
		return m.ConcurrencyUpdateInterval
	}
	return nil
}

type GradientControllerConfig_MinimumRTTCalculationParams struct {
	Interval             *duration.Duration    `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	RequestCount         *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	Jitter               *v3alpha.Percent      `protobuf:"bytes,3,opt,name=jitter,proto3" json:"jitter,omitempty"`
	MinConcurrency       *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=min_concurrency,json=minConcurrency,proto3" json:"min_concurrency,omitempty"`
	Buffer               *v3alpha.Percent      `protobuf:"bytes,5,opt,name=buffer,proto3" json:"buffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) Reset() {
	*m = GradientControllerConfig_MinimumRTTCalculationParams{}
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) String() string {
	return proto.CompactTextString(m)
}
func (*GradientControllerConfig_MinimumRTTCalculationParams) ProtoMessage() {}
func (*GradientControllerConfig_MinimumRTTCalculationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_778e4cdafe4c14e9, []int{0, 1}
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Unmarshal(m, b)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Merge(m, src)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Size(m)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams proto.InternalMessageInfo

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetRequestCount() *wrappers.UInt32Value {
	if m != nil {
		return m.RequestCount
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetJitter() *v3alpha.Percent {
	if m != nil {
		return m.Jitter
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetMinConcurrency() *wrappers.UInt32Value {
	if m != nil {
		return m.MinConcurrency
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetBuffer() *v3alpha.Percent {
	if m != nil {
		return m.Buffer
	}
	return nil
}

type AdaptiveConcurrency struct {
	// Types that are valid to be assigned to ConcurrencyControllerConfig:
	//	*AdaptiveConcurrency_GradientControllerConfig
	ConcurrencyControllerConfig isAdaptiveConcurrency_ConcurrencyControllerConfig `protobuf_oneof:"concurrency_controller_config"`
	Enabled                     *v3alpha1.RuntimeFeatureFlag                      `protobuf:"bytes,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                                          `json:"-"`
	XXX_unrecognized            []byte                                            `json:"-"`
	XXX_sizecache               int32                                             `json:"-"`
}

func (m *AdaptiveConcurrency) Reset()         { *m = AdaptiveConcurrency{} }
func (m *AdaptiveConcurrency) String() string { return proto.CompactTextString(m) }
func (*AdaptiveConcurrency) ProtoMessage()    {}
func (*AdaptiveConcurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_778e4cdafe4c14e9, []int{1}
}

func (m *AdaptiveConcurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdaptiveConcurrency.Unmarshal(m, b)
}
func (m *AdaptiveConcurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdaptiveConcurrency.Marshal(b, m, deterministic)
}
func (m *AdaptiveConcurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdaptiveConcurrency.Merge(m, src)
}
func (m *AdaptiveConcurrency) XXX_Size() int {
	return xxx_messageInfo_AdaptiveConcurrency.Size(m)
}
func (m *AdaptiveConcurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_AdaptiveConcurrency.DiscardUnknown(m)
}

var xxx_messageInfo_AdaptiveConcurrency proto.InternalMessageInfo

type isAdaptiveConcurrency_ConcurrencyControllerConfig interface {
	isAdaptiveConcurrency_ConcurrencyControllerConfig()
}

type AdaptiveConcurrency_GradientControllerConfig struct {
	GradientControllerConfig *GradientControllerConfig `protobuf:"bytes,1,opt,name=gradient_controller_config,json=gradientControllerConfig,proto3,oneof"`
}

func (*AdaptiveConcurrency_GradientControllerConfig) isAdaptiveConcurrency_ConcurrencyControllerConfig() {
}

func (m *AdaptiveConcurrency) GetConcurrencyControllerConfig() isAdaptiveConcurrency_ConcurrencyControllerConfig {
	if m != nil {
		return m.ConcurrencyControllerConfig
	}
	return nil
}

func (m *AdaptiveConcurrency) GetGradientControllerConfig() *GradientControllerConfig {
	if x, ok := m.GetConcurrencyControllerConfig().(*AdaptiveConcurrency_GradientControllerConfig); ok {
		return x.GradientControllerConfig
	}
	return nil
}

func (m *AdaptiveConcurrency) GetEnabled() *v3alpha1.RuntimeFeatureFlag {
	if m != nil {
		return m.Enabled
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdaptiveConcurrency) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdaptiveConcurrency_GradientControllerConfig)(nil),
	}
}

func init() {
	proto.RegisterType((*GradientControllerConfig)(nil), "envoy.extensions.filters.http.adaptive_concurrency.v3alpha.GradientControllerConfig")
	proto.RegisterType((*GradientControllerConfig_ConcurrencyLimitCalculationParams)(nil), "envoy.extensions.filters.http.adaptive_concurrency.v3alpha.GradientControllerConfig.ConcurrencyLimitCalculationParams")
	proto.RegisterType((*GradientControllerConfig_MinimumRTTCalculationParams)(nil), "envoy.extensions.filters.http.adaptive_concurrency.v3alpha.GradientControllerConfig.MinimumRTTCalculationParams")
	proto.RegisterType((*AdaptiveConcurrency)(nil), "envoy.extensions.filters.http.adaptive_concurrency.v3alpha.AdaptiveConcurrency")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/adaptive_concurrency/v3alpha/adaptive_concurrency.proto", fileDescriptor_778e4cdafe4c14e9)
}

var fileDescriptor_778e4cdafe4c14e9 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xad, 0xd3, 0xa7, 0x86, 0x57, 0xe5, 0x0a, 0x70, 0x93, 0x52, 0xb5, 0x55, 0x17, 0xa8, 0x12,
	0xb6, 0xd4, 0xec, 0xb2, 0xab, 0x83, 0xfa, 0xe0, 0x51, 0x45, 0x56, 0x83, 0x84, 0x58, 0x58, 0x13,
	0xe7, 0xc6, 0x9d, 0x6a, 0x3c, 0x9e, 0x4e, 0xc6, 0x69, 0xf2, 0x07, 0x88, 0x4f, 0x60, 0xc1, 0x1e,
	0x3e, 0x81, 0x3d, 0x1f, 0xc1, 0x86, 0xaf, 0x40, 0x02, 0x75, 0x85, 0xec, 0x99, 0x24, 0xa6, 0x4d,
	0x03, 0x7d, 0xed, 0x12, 0xdd, 0x7b, 0xce, 0x3d, 0xf7, 0xdc, 0x7b, 0xc7, 0xa8, 0x0e, 0xac, 0x13,
	0xf7, 0x1c, 0xe8, 0x4a, 0x60, 0x6d, 0x12, 0xb3, 0xb6, 0xd3, 0x22, 0x54, 0x82, 0x68, 0x3b, 0x87,
	0x52, 0x72, 0x07, 0x37, 0x31, 0x97, 0xa4, 0x03, 0x7e, 0x10, 0xb3, 0x20, 0x11, 0x02, 0x58, 0xd0,
	0x73, 0x3a, 0x65, 0x4c, 0xf9, 0x21, 0x1e, 0x19, 0xb4, 0xb9, 0x88, 0x65, 0x6c, 0x56, 0x32, 0x5a,
	0x7b, 0x48, 0x6b, 0x6b, 0x5a, 0x3b, 0xa5, 0xb5, 0x47, 0x22, 0x35, 0x6d, 0x71, 0x5d, 0x49, 0x0a,
	0x62, 0xd6, 0x22, 0xa1, 0x13, 0xc4, 0x02, 0x06, 0x15, 0x1b, 0xb8, 0x0d, 0xaa, 0x42, 0x71, 0x45,
	0x65, 0xc9, 0x1e, 0x1f, 0x86, 0x39, 0x88, 0x00, 0x98, 0xd4, 0x19, 0x4b, 0x61, 0x1c, 0x87, 0x14,
	0x1c, 0xcc, 0x89, 0x83, 0x19, 0x8b, 0x25, 0x96, 0x99, 0x12, 0x15, 0x5d, 0xd6, 0xd1, 0xec, 0x5f,
	0x23, 0x69, 0x39, 0xcd, 0x44, 0x64, 0x09, 0x17, 0xc5, 0x4f, 0x04, 0xe6, 0x3c, 0xed, 0x40, 0xc5,
	0x57, 0x93, 0x26, 0xc7, 0x79, 0x5e, 0xa7, 0x03, 0x22, 0x6d, 0x95, 0xb0, 0x50, 0xa7, 0x3c, 0xee,
	0x60, 0x4a, 0x9a, 0x58, 0x82, 0xd3, 0xff, 0xa1, 0x02, 0x6b, 0x3f, 0x11, 0xb2, 0x76, 0x04, 0x6e,
	0x12, 0x60, 0xb2, 0x1a, 0x33, 0x29, 0x62, 0x4a, 0x41, 0x54, 0xb3, 0x86, 0xcd, 0x77, 0xa8, 0xd4,
	0xc6, 0x11, 0xa7, 0xe0, 0xe3, 0x30, 0x14, 0x10, 0x62, 0x09, 0xbe, 0x6e, 0x8c, 0x50, 0xb0, 0x8c,
	0x15, 0xe3, 0xe9, 0x9d, 0xcd, 0x92, 0xad, 0x0c, 0x4e, 0xdb, 0xef, 0x1b, 0x67, 0xd7, 0x54, 0x96,
	0xb7, 0xa8, 0xf0, 0x5b, 0x7d, 0x78, 0x6d, 0x80, 0x36, 0xbf, 0x1a, 0xc8, 0xca, 0x79, 0xee, 0x53,
	0x12, 0x11, 0xe9, 0x73, 0x2c, 0x70, 0xd4, 0xb6, 0x0a, 0x19, 0x75, 0xc7, 0xbe, 0xfa, 0xec, 0xec,
	0x8b, 0xba, 0xb2, 0xab, 0xc3, 0xe4, 0x57, 0x69, 0xcd, 0x2a, 0xa6, 0x41, 0x42, 0x33, 0xe7, 0x6a,
	0x59, 0x75, 0x77, 0xee, 0xd4, 0x9d, 0xfe, 0x60, 0x14, 0xe6, 0x0d, 0xef, 0x51, 0x70, 0x26, 0x59,
	0x65, 0x98, 0x9f, 0x0d, 0xb4, 0x10, 0x11, 0xe6, 0x0b, 0x29, 0xfd, 0x00, 0xd3, 0xa0, 0xaf, 0x7b,
	0x32, 0xd3, 0xcd, 0x6f, 0x45, 0xf7, 0x6b, 0xc2, 0x48, 0x94, 0x44, 0xde, 0xc1, 0xc1, 0x38, 0xc5,
	0xf3, 0x11, 0x61, 0x9e, 0xcc, 0x9a, 0x52, 0xb1, 0xe2, 0x8f, 0x02, 0x5a, 0xfd, 0x67, 0xcf, 0xe6,
	0x5b, 0xf4, 0x30, 0xc2, 0x5d, 0xff, 0xdc, 0x44, 0xf4, 0x28, 0x96, 0x6c, 0xb5, 0x84, 0x76, 0x7f,
	0x09, 0xed, 0xfa, 0x1e, 0x93, 0xe5, 0xcd, 0x37, 0x98, 0x26, 0xe0, 0xce, 0x9e, 0xba, 0x53, 0x1b,
	0x85, 0x95, 0x09, 0x6f, 0x21, 0xc2, 0xdd, 0xb3, 0xb5, 0x4c, 0x40, 0xa5, 0x3c, 0x6d, 0xc2, 0xd3,
	0xfd, 0xf3, 0x09, 0x93, 0x20, 0x3a, 0x98, 0x6a, 0xcf, 0x16, 0xcf, 0x15, 0x78, 0xae, 0xaf, 0xc0,
	0x45, 0xa7, 0xee, 0xec, 0x17, 0x63, 0x6a, 0xce, 0xd8, 0x98, 0xf0, 0x16, 0x73, 0x4c, 0xf5, 0x8c,
	0x68, 0x4f, 0xf3, 0x54, 0x4e, 0x3e, 0x7e, 0x7b, 0xbf, 0x2c, 0x10, 0x57, 0xde, 0xab, 0x9b, 0xd5,
	0xbe, 0x8f, 0xb3, 0x7d, 0xf3, 0xba, 0xeb, 0x52, 0xfc, 0x3e, 0x89, 0x4a, 0x63, 0x86, 0x63, 0x6e,
	0xa1, 0xb9, 0x41, 0xb3, 0xc6, 0x65, 0x9a, 0x1d, 0xc0, 0xcc, 0x17, 0xe8, 0x9e, 0x80, 0xe3, 0x04,
	0xda, 0xd2, 0x0f, 0xe2, 0x84, 0x5d, 0x72, 0x2a, 0x77, 0x35, 0xb6, 0x9a, 0x42, 0xcd, 0x32, 0x9a,
	0x39, 0x22, 0x52, 0x82, 0xd0, 0xce, 0x8f, 0x3d, 0x60, 0x9d, 0x6a, 0xee, 0xa3, 0x07, 0xe9, 0xbe,
	0xe7, 0x9c, 0xb3, 0xa6, 0x2e, 0x23, 0xe1, 0x7e, 0x44, 0x58, 0xce, 0xc9, 0x54, 0x44, 0x23, 0x69,
	0xb5, 0x40, 0x58, 0xd3, 0xff, 0x21, 0x42, 0xa5, 0x56, 0x8e, 0xd3, 0x09, 0x53, 0x74, 0x74, 0xc3,
	0x13, 0x1e, 0x33, 0xbb, 0x4a, 0x3d, 0x2d, 0x59, 0x43, 0xfb, 0x37, 0x5b, 0x72, 0xed, 0x57, 0x01,
	0x2d, 0x6c, 0x69, 0x64, 0xde, 0x96, 0x4f, 0x06, 0x2a, 0x86, 0x1a, 0x94, 0x32, 0x6a, 0x94, 0xaf,
	0x0a, 0xeb, 0xed, 0x39, 0xb8, 0x8d, 0xe7, 0x65, 0xf8, 0x84, 0xec, 0x4e, 0x78, 0x56, 0x78, 0xd1,
	0x27, 0x61, 0x07, 0xcd, 0x02, 0xc3, 0x0d, 0x0a, 0x4d, 0xbd, 0x82, 0xcf, 0xec, 0xbf, 0xac, 0x49,
	0xbf, 0x91, 0x83, 0x5a, 0x5e, 0xc2, 0x24, 0x89, 0x60, 0x1b, 0xb0, 0x4c, 0x04, 0x6c, 0x53, 0x1c,
	0x7a, 0x7d, 0x74, 0xa5, 0x96, 0x1a, 0xfb, 0x12, 0xed, 0x5d, 0xcd, 0xd8, 0x11, 0xde, 0xb9, 0xeb,
	0xe8, 0x49, 0xfe, 0x99, 0x39, 0xe7, 0x9e, 0x39, 0xf9, 0xdb, 0x35, 0xdc, 0x16, 0xda, 0x25, 0xb1,
	0xd2, 0xcc, 0x45, 0xdc, 0xed, 0x5d, 0xc3, 0x4b, 0xd7, 0x1a, 0x21, 0xa3, 0x96, 0x9e, 0x41, 0xcd,
	0x68, 0xcc, 0x64, 0xf7, 0x50, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x60, 0x4f, 0x16, 0x35, 0xd6,
	0x08, 0x00, 0x00,
}
