// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/envoy/http/service_control/requirement.proto

package google_api_envoy_http_service_control

import (
	fmt "fmt"
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

type APIKeyLocation struct {
	// Types that are valid to be assigned to Key:
	//	*APIKeyLocation_Query
	//	*APIKeyLocation_Header
	//	*APIKeyLocation_Cookie
	Key                  isAPIKeyLocation_Key `protobuf_oneof:"key"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *APIKeyLocation) Reset()         { *m = APIKeyLocation{} }
func (m *APIKeyLocation) String() string { return proto.CompactTextString(m) }
func (*APIKeyLocation) ProtoMessage()    {}
func (*APIKeyLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dda622705118f3d, []int{0}
}

func (m *APIKeyLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIKeyLocation.Unmarshal(m, b)
}
func (m *APIKeyLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIKeyLocation.Marshal(b, m, deterministic)
}
func (m *APIKeyLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIKeyLocation.Merge(m, src)
}
func (m *APIKeyLocation) XXX_Size() int {
	return xxx_messageInfo_APIKeyLocation.Size(m)
}
func (m *APIKeyLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_APIKeyLocation.DiscardUnknown(m)
}

var xxx_messageInfo_APIKeyLocation proto.InternalMessageInfo

type isAPIKeyLocation_Key interface {
	isAPIKeyLocation_Key()
}

type APIKeyLocation_Query struct {
	Query string `protobuf:"bytes,1,opt,name=query,proto3,oneof"`
}

type APIKeyLocation_Header struct {
	Header string `protobuf:"bytes,2,opt,name=header,proto3,oneof"`
}

type APIKeyLocation_Cookie struct {
	Cookie string `protobuf:"bytes,3,opt,name=cookie,proto3,oneof"`
}

func (*APIKeyLocation_Query) isAPIKeyLocation_Key() {}

func (*APIKeyLocation_Header) isAPIKeyLocation_Key() {}

func (*APIKeyLocation_Cookie) isAPIKeyLocation_Key() {}

func (m *APIKeyLocation) GetKey() isAPIKeyLocation_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *APIKeyLocation) GetQuery() string {
	if x, ok := m.GetKey().(*APIKeyLocation_Query); ok {
		return x.Query
	}
	return ""
}

func (m *APIKeyLocation) GetHeader() string {
	if x, ok := m.GetKey().(*APIKeyLocation_Header); ok {
		return x.Header
	}
	return ""
}

func (m *APIKeyLocation) GetCookie() string {
	if x, ok := m.GetKey().(*APIKeyLocation_Cookie); ok {
		return x.Cookie
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*APIKeyLocation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*APIKeyLocation_Query)(nil),
		(*APIKeyLocation_Header)(nil),
		(*APIKeyLocation_Cookie)(nil),
	}
}

type APIKeyRequirement struct {
	Locations            []*APIKeyLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	AllowWithoutApiKey   bool              `protobuf:"varint,2,opt,name=allow_without_api_key,json=allowWithoutApiKey,proto3" json:"allow_without_api_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *APIKeyRequirement) Reset()         { *m = APIKeyRequirement{} }
func (m *APIKeyRequirement) String() string { return proto.CompactTextString(m) }
func (*APIKeyRequirement) ProtoMessage()    {}
func (*APIKeyRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dda622705118f3d, []int{1}
}

func (m *APIKeyRequirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIKeyRequirement.Unmarshal(m, b)
}
func (m *APIKeyRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIKeyRequirement.Marshal(b, m, deterministic)
}
func (m *APIKeyRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIKeyRequirement.Merge(m, src)
}
func (m *APIKeyRequirement) XXX_Size() int {
	return xxx_messageInfo_APIKeyRequirement.Size(m)
}
func (m *APIKeyRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_APIKeyRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_APIKeyRequirement proto.InternalMessageInfo

func (m *APIKeyRequirement) GetLocations() []*APIKeyLocation {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *APIKeyRequirement) GetAllowWithoutApiKey() bool {
	if m != nil {
		return m.AllowWithoutApiKey
	}
	return false
}

type MetricCost struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cost                 int64    `protobuf:"varint,2,opt,name=cost,proto3" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricCost) Reset()         { *m = MetricCost{} }
func (m *MetricCost) String() string { return proto.CompactTextString(m) }
func (*MetricCost) ProtoMessage()    {}
func (*MetricCost) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dda622705118f3d, []int{2}
}

func (m *MetricCost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricCost.Unmarshal(m, b)
}
func (m *MetricCost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricCost.Marshal(b, m, deterministic)
}
func (m *MetricCost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricCost.Merge(m, src)
}
func (m *MetricCost) XXX_Size() int {
	return xxx_messageInfo_MetricCost.Size(m)
}
func (m *MetricCost) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricCost.DiscardUnknown(m)
}

var xxx_messageInfo_MetricCost proto.InternalMessageInfo

func (m *MetricCost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricCost) GetCost() int64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

type Requirement struct {
	ServiceName          string             `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	OperationName        string             `protobuf:"bytes,2,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	ApiKey               *APIKeyRequirement `protobuf:"bytes,3,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	CustomLabels         []string           `protobuf:"bytes,4,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty"`
	ApiName              string             `protobuf:"bytes,5,opt,name=api_name,json=apiName,proto3" json:"api_name,omitempty"`
	ApiVersion           string             `protobuf:"bytes,6,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	SkipServiceControl   bool               `protobuf:"varint,7,opt,name=skip_service_control,json=skipServiceControl,proto3" json:"skip_service_control,omitempty"`
	MetricCosts          []*MetricCost      `protobuf:"bytes,8,rep,name=metric_costs,json=metricCosts,proto3" json:"metric_costs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Requirement) Reset()         { *m = Requirement{} }
func (m *Requirement) String() string { return proto.CompactTextString(m) }
func (*Requirement) ProtoMessage()    {}
func (*Requirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dda622705118f3d, []int{3}
}

func (m *Requirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Requirement.Unmarshal(m, b)
}
func (m *Requirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Requirement.Marshal(b, m, deterministic)
}
func (m *Requirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Requirement.Merge(m, src)
}
func (m *Requirement) XXX_Size() int {
	return xxx_messageInfo_Requirement.Size(m)
}
func (m *Requirement) XXX_DiscardUnknown() {
	xxx_messageInfo_Requirement.DiscardUnknown(m)
}

var xxx_messageInfo_Requirement proto.InternalMessageInfo

func (m *Requirement) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Requirement) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *Requirement) GetApiKey() *APIKeyRequirement {
	if m != nil {
		return m.ApiKey
	}
	return nil
}

func (m *Requirement) GetCustomLabels() []string {
	if m != nil {
		return m.CustomLabels
	}
	return nil
}

func (m *Requirement) GetApiName() string {
	if m != nil {
		return m.ApiName
	}
	return ""
}

func (m *Requirement) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *Requirement) GetSkipServiceControl() bool {
	if m != nil {
		return m.SkipServiceControl
	}
	return false
}

func (m *Requirement) GetMetricCosts() []*MetricCost {
	if m != nil {
		return m.MetricCosts
	}
	return nil
}

func init() {
	proto.RegisterType((*APIKeyLocation)(nil), "google.api.envoy.http.service_control.APIKeyLocation")
	proto.RegisterType((*APIKeyRequirement)(nil), "google.api.envoy.http.service_control.APIKeyRequirement")
	proto.RegisterType((*MetricCost)(nil), "google.api.envoy.http.service_control.MetricCost")
	proto.RegisterType((*Requirement)(nil), "google.api.envoy.http.service_control.Requirement")
}

func init() {
	proto.RegisterFile("api/envoy/http/service_control/requirement.proto", fileDescriptor_3dda622705118f3d)
}

var fileDescriptor_3dda622705118f3d = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x26, 0xd9, 0x24, 0x4d, 0x66, 0xd3, 0x4a, 0x58, 0xfc, 0x2c, 0x5c, 0x88, 0x82, 0x90, 0x72,
	0x40, 0xbb, 0x69, 0x01, 0x89, 0x6b, 0xdb, 0x0b, 0xa8, 0x05, 0x81, 0x8b, 0xe0, 0xb8, 0x72, 0xb7,
	0xa3, 0xc6, 0xca, 0xee, 0x8e, 0x6b, 0x3b, 0xa9, 0xf6, 0x59, 0x78, 0x02, 0x5e, 0x81, 0x13, 0xaf,
	0xc3, 0x5b, 0x20, 0xdb, 0x49, 0x53, 0x72, 0x6a, 0x6f, 0xe3, 0xf9, 0xe6, 0x9b, 0xf9, 0xbe, 0x19,
	0xc3, 0x54, 0x28, 0x99, 0x61, 0xbd, 0xa4, 0x26, 0x9b, 0x59, 0xab, 0x32, 0x83, 0x7a, 0x29, 0x0b,
	0xcc, 0x0b, 0xaa, 0xad, 0xa6, 0x32, 0xd3, 0x78, 0xb5, 0x90, 0x1a, 0x2b, 0xac, 0x6d, 0xaa, 0x34,
	0x59, 0x62, 0xaf, 0x2e, 0x89, 0x2e, 0x4b, 0x4c, 0x85, 0x92, 0xa9, 0x27, 0xa6, 0x8e, 0x98, 0x6e,
	0x11, 0x9f, 0x3f, 0x5d, 0x8a, 0x52, 0x5e, 0x08, 0x8b, 0xd9, 0x3a, 0x08, 0xfc, 0x71, 0x01, 0x7b,
	0x87, 0x5f, 0x3e, 0x9e, 0x60, 0x73, 0x4a, 0x85, 0xb0, 0x92, 0x6a, 0xf6, 0x04, 0xba, 0x57, 0x0b,
	0xd4, 0x4d, 0xd2, 0x1a, 0xb5, 0x26, 0x83, 0x0f, 0x0f, 0x78, 0x78, 0xb2, 0x04, 0x7a, 0x33, 0x14,
	0x17, 0xa8, 0x93, 0xf6, 0x0a, 0x58, 0xbd, 0x1d, 0x52, 0x10, 0xcd, 0x25, 0x26, 0xd1, 0x1a, 0x09,
	0xef, 0xa3, 0x2e, 0x44, 0x73, 0x6c, 0xc6, 0x3f, 0x5b, 0xf0, 0x30, 0x4c, 0xe1, 0x1b, 0x03, 0xec,
	0x0c, 0x06, 0xe5, 0x6a, 0xa8, 0x49, 0x5a, 0xa3, 0x68, 0x12, 0x1f, 0xbc, 0x4b, 0xef, 0x64, 0x27,
	0xfd, 0x5f, 0x32, 0xdf, 0xf4, 0x61, 0xfb, 0xf0, 0x58, 0x94, 0x25, 0x5d, 0xe7, 0xd7, 0xd2, 0xce,
	0x68, 0x61, 0x73, 0xa1, 0x64, 0x3e, 0xc7, 0xc6, 0x8b, 0xee, 0x73, 0xe6, 0xc1, 0x1f, 0x01, 0x3b,
	0x54, 0xf2, 0x04, 0x9b, 0xf1, 0x5b, 0x80, 0x4f, 0x68, 0xb5, 0x2c, 0x8e, 0xc9, 0x58, 0xc6, 0xa0,
	0x53, 0x8b, 0x0a, 0x83, 0x7b, 0xee, 0x63, 0x97, 0x2b, 0xc8, 0x58, 0xdf, 0x23, 0xe2, 0x3e, 0x1e,
	0xff, 0x8a, 0x20, 0xbe, 0xed, 0xe6, 0x35, 0x0c, 0xd7, 0x2a, 0x37, 0xfc, 0xa3, 0xc1, 0xef, 0xbf,
	0x7f, 0xa2, 0x8e, 0x6e, 0x8f, 0x5a, 0x3c, 0x5e, 0xc1, 0x9f, 0x5d, 0xc7, 0x29, 0xec, 0x91, 0x42,
	0xed, 0x45, 0x87, 0xfa, 0xf6, 0x76, 0xfd, 0xee, 0x4d, 0x81, 0x67, 0x7c, 0x85, 0x9d, 0xb5, 0x15,
	0xb7, 0xe5, 0xf8, 0xe0, 0xfd, 0xbd, 0x76, 0x75, 0x4b, 0x2a, 0xef, 0x09, 0x6f, 0x9c, 0xbd, 0x84,
	0xdd, 0x62, 0x61, 0x2c, 0x55, 0x79, 0x29, 0xce, 0xb1, 0x34, 0x49, 0x67, 0x14, 0x4d, 0x06, 0x7c,
	0x18, 0x92, 0xa7, 0x3e, 0xc7, 0x9e, 0x41, 0xdf, 0xcd, 0xf5, 0x1a, 0xbb, 0x7e, 0x27, 0x4e, 0x87,
	0x97, 0xf4, 0x02, 0x62, 0x07, 0x2d, 0x51, 0x1b, 0x49, 0x75, 0xd2, 0xf3, 0x28, 0x08, 0x25, 0xbf,
	0x87, 0x0c, 0x9b, 0xc2, 0x23, 0x33, 0x97, 0x2a, 0xdf, 0x92, 0x94, 0xec, 0x84, 0x5b, 0x38, 0xec,
	0x2c, 0x40, 0xc7, 0x01, 0x61, 0xdf, 0x60, 0x58, 0xf9, 0x5b, 0xe4, 0x6e, 0xc9, 0x26, 0xe9, 0xfb,
	0x6f, 0xb1, 0x7f, 0x47, 0xab, 0x9b, 0x33, 0xf2, 0xb8, 0xba, 0x89, 0xcd, 0x79, 0xcf, 0xff, 0xf5,
	0x37, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x06, 0xa9, 0xfe, 0x5f, 0x03, 0x00, 0x00,
}
