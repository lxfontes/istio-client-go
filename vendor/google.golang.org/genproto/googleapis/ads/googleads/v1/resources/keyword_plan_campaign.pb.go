// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/keyword_plan_campaign.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A Keyword Plan campaign.
// Max number of keyword plan campaigns per plan allowed: 1.
type KeywordPlanCampaign struct {
	// The resource name of the Keyword Plan campaign.
	// KeywordPlanCampaign resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanCampaigns/{kp_campaign_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The keyword plan this campaign belongs to.
	KeywordPlan *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan,json=keywordPlan,proto3" json:"keyword_plan,omitempty"`
	// The ID of the Keyword Plan campaign.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Keyword Plan campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The languages targeted for the Keyword Plan campaign.
	// Max allowed: 1.
	LanguageConstants []*wrappers.StringValue `protobuf:"bytes,5,rep,name=language_constants,json=languageConstants,proto3" json:"language_constants,omitempty"`
	// Targeting network.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,6,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v1.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// A default max cpc bid in micros, and in the account currency, for all ad
	// groups under the campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,7,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// The geo targets.
	// Max number allowed: 20.
	GeoTargets           []*KeywordPlanGeoTarget `protobuf:"bytes,8,rep,name=geo_targets,json=geoTargets,proto3" json:"geo_targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordPlanCampaign) Reset()         { *m = KeywordPlanCampaign{} }
func (m *KeywordPlanCampaign) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCampaign) ProtoMessage()    {}
func (*KeywordPlanCampaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f4e57860a6ad48d, []int{0}
}

func (m *KeywordPlanCampaign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCampaign.Unmarshal(m, b)
}
func (m *KeywordPlanCampaign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCampaign.Marshal(b, m, deterministic)
}
func (m *KeywordPlanCampaign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCampaign.Merge(m, src)
}
func (m *KeywordPlanCampaign) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCampaign.Size(m)
}
func (m *KeywordPlanCampaign) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCampaign.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCampaign proto.InternalMessageInfo

func (m *KeywordPlanCampaign) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanCampaign) GetKeywordPlan() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlan
	}
	return nil
}

func (m *KeywordPlanCampaign) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanCampaign) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlanCampaign) GetLanguageConstants() []*wrappers.StringValue {
	if m != nil {
		return m.LanguageConstants
	}
	return nil
}

func (m *KeywordPlanCampaign) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

func (m *KeywordPlanCampaign) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *KeywordPlanCampaign) GetGeoTargets() []*KeywordPlanGeoTarget {
	if m != nil {
		return m.GeoTargets
	}
	return nil
}

// A geo target.
// Next ID: 3
type KeywordPlanGeoTarget struct {
	// Required. The resource name of the geo target.
	GeoTargetConstant    *wrappers.StringValue `protobuf:"bytes,1,opt,name=geo_target_constant,json=geoTargetConstant,proto3" json:"geo_target_constant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KeywordPlanGeoTarget) Reset()         { *m = KeywordPlanGeoTarget{} }
func (m *KeywordPlanGeoTarget) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanGeoTarget) ProtoMessage()    {}
func (*KeywordPlanGeoTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f4e57860a6ad48d, []int{1}
}

func (m *KeywordPlanGeoTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanGeoTarget.Unmarshal(m, b)
}
func (m *KeywordPlanGeoTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanGeoTarget.Marshal(b, m, deterministic)
}
func (m *KeywordPlanGeoTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanGeoTarget.Merge(m, src)
}
func (m *KeywordPlanGeoTarget) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanGeoTarget.Size(m)
}
func (m *KeywordPlanGeoTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanGeoTarget.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanGeoTarget proto.InternalMessageInfo

func (m *KeywordPlanGeoTarget) GetGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstant
	}
	return nil
}

func init() {
	proto.RegisterType((*KeywordPlanCampaign)(nil), "google.ads.googleads.v1.resources.KeywordPlanCampaign")
	proto.RegisterType((*KeywordPlanGeoTarget)(nil), "google.ads.googleads.v1.resources.KeywordPlanGeoTarget")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/keyword_plan_campaign.proto", fileDescriptor_1f4e57860a6ad48d)
}

var fileDescriptor_1f4e57860a6ad48d = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x8a, 0x13, 0x31,
	0x14, 0x86, 0xe9, 0xb4, 0xae, 0x9a, 0xd6, 0x05, 0xb3, 0x7b, 0x31, 0xac, 0x8b, 0x74, 0x57, 0x16,
	0x0a, 0x42, 0xc6, 0xae, 0xa2, 0x32, 0x22, 0x32, 0x5d, 0xa4, 0xea, 0xea, 0x52, 0xaa, 0x14, 0x91,
	0xc2, 0x90, 0x4e, 0x62, 0x08, 0xed, 0x24, 0x43, 0x92, 0xd9, 0xa2, 0xf7, 0xbe, 0x88, 0x97, 0x3e,
	0x85, 0xd7, 0x3e, 0x8a, 0x4f, 0x21, 0xcd, 0x4c, 0x66, 0x8b, 0xdb, 0x75, 0xbc, 0x3b, 0x93, 0xf3,
	0x7f, 0xe7, 0xfc, 0x73, 0x4e, 0x02, 0x9e, 0x33, 0x29, 0xd9, 0x82, 0x06, 0x98, 0xe8, 0xa0, 0x08,
	0x57, 0xd1, 0x79, 0x3f, 0x50, 0x54, 0xcb, 0x5c, 0x25, 0x54, 0x07, 0x73, 0xfa, 0x65, 0x29, 0x15,
	0x89, 0xb3, 0x05, 0x16, 0x71, 0x82, 0xd3, 0x0c, 0x73, 0x26, 0x50, 0xa6, 0xa4, 0x91, 0xf0, 0xa0,
	0x60, 0x10, 0x26, 0x1a, 0x55, 0x38, 0x3a, 0xef, 0xa3, 0x0a, 0xdf, 0x7b, 0x7a, 0x55, 0x07, 0x2a,
	0xf2, 0xf4, 0xaf, 0xea, 0x82, 0x9a, 0xa5, 0x54, 0xf3, 0xa2, 0xf8, 0xde, 0xbe, 0x23, 0x33, 0x1e,
	0x60, 0x21, 0xa4, 0xc1, 0x86, 0x4b, 0xa1, 0xcb, 0xec, 0xdd, 0x32, 0x6b, 0xbf, 0x66, 0xf9, 0xe7,
	0x60, 0xa9, 0x70, 0x96, 0x51, 0x55, 0xe6, 0x0f, 0x7f, 0xb6, 0xc0, 0xce, 0x69, 0x51, 0x7c, 0xb4,
	0xc0, 0xe2, 0xa4, 0x34, 0x0e, 0xef, 0x81, 0x5b, 0xce, 0x5c, 0x2c, 0x70, 0x4a, 0xfd, 0x46, 0xb7,
	0xd1, 0xbb, 0x39, 0xee, 0xb8, 0xc3, 0x33, 0x9c, 0x52, 0xf8, 0x02, 0x74, 0xd6, 0x8d, 0xf9, 0x5e,
	0xb7, 0xd1, 0x6b, 0x1f, 0xef, 0x97, 0xff, 0x88, 0x5c, 0x4f, 0xf4, 0xde, 0x28, 0x2e, 0xd8, 0x04,
	0x2f, 0x72, 0x3a, 0x6e, 0xcf, 0x2f, 0xba, 0xc1, 0xfb, 0xc0, 0xe3, 0xc4, 0x6f, 0x5a, 0xec, 0xce,
	0x25, 0xec, 0xb5, 0x30, 0x8f, 0x1f, 0x15, 0x94, 0xc7, 0x09, 0x7c, 0x00, 0x5a, 0xd6, 0x49, 0xeb,
	0x3f, 0xba, 0x58, 0x25, 0x3c, 0x05, 0x70, 0x81, 0x05, 0xcb, 0x31, 0xa3, 0x71, 0x22, 0x85, 0x36,
	0x58, 0x18, 0xed, 0x5f, 0xeb, 0x36, 0x6b, 0xf9, 0xdb, 0x8e, 0x3b, 0x71, 0x18, 0xfc, 0x0a, 0x76,
	0x37, 0x6d, 0xc1, 0xdf, 0xea, 0x36, 0x7a, 0xdb, 0xc7, 0xaf, 0xd0, 0x55, 0x3b, 0xb6, 0x0b, 0x44,
	0x6b, 0x33, 0x3e, 0x2b, 0xc0, 0x97, 0x22, 0x4f, 0x37, 0x1c, 0x8f, 0xe1, 0xfc, 0xd2, 0x19, 0x8c,
	0xc0, 0x76, 0x92, 0x25, 0xf1, 0x8c, 0x93, 0x38, 0xe5, 0x89, 0x92, 0xda, 0xbf, 0x5e, 0x3f, 0xb3,
	0x4e, 0x92, 0x25, 0x03, 0x4e, 0xde, 0x59, 0x00, 0x7e, 0x04, 0x6d, 0x46, 0x65, 0x6c, 0xb0, 0x62,
	0xd4, 0x68, 0xff, 0x86, 0x1d, 0xc2, 0x13, 0x54, 0x7b, 0x33, 0xd7, 0x2d, 0x0e, 0xa9, 0xfc, 0x60,
	0xf9, 0x31, 0x60, 0x2e, 0xd4, 0x87, 0x04, 0xec, 0x6e, 0xd2, 0xc0, 0xb7, 0x60, 0xe7, 0xa2, 0x63,
	0x35, 0x7f, 0x7b, 0x91, 0x6a, 0xc7, 0x5f, 0x95, 0x77, 0xf3, 0x1f, 0x7c, 0xf3, 0xc0, 0x51, 0x22,
	0xd3, 0x7a, 0xc3, 0x03, 0x7f, 0xc3, 0x7d, 0x1e, 0xad, 0xba, 0x8c, 0x1a, 0x9f, 0xde, 0x94, 0x38,
	0x93, 0xab, 0x05, 0x23, 0xa9, 0x58, 0xc0, 0xa8, 0xb0, 0x1e, 0xdc, 0xb3, 0xcb, 0xb8, 0xfe, 0xc7,
	0x3b, 0x7f, 0x56, 0x45, 0xdf, 0xbd, 0xe6, 0x30, 0x8a, 0x7e, 0x78, 0x07, 0xc3, 0xa2, 0x64, 0x44,
	0x34, 0x2a, 0xc2, 0x55, 0x34, 0xe9, 0xa3, 0xb1, 0x53, 0xfe, 0x72, 0x9a, 0x69, 0x44, 0xf4, 0xb4,
	0xd2, 0x4c, 0x27, 0xfd, 0x69, 0xa5, 0xf9, 0xed, 0x1d, 0x15, 0x89, 0x30, 0x8c, 0x88, 0x0e, 0xc3,
	0x4a, 0x15, 0x86, 0x93, 0x7e, 0x18, 0x56, 0xba, 0xd9, 0x96, 0x35, 0xfb, 0xf0, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe9, 0xdf, 0x0e, 0x07, 0x93, 0x04, 0x00, 0x00,
}
