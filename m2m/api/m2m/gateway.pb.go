// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GatewayMode int32

const (
	GatewayMode_GW_INACTIVE              GatewayMode = 0
	GatewayMode_GW_FREE_GATEWAYS_LIMITED GatewayMode = 1
	GatewayMode_GW_WHOLE_NETWORK         GatewayMode = 2
	GatewayMode_GW_DELETED               GatewayMode = 3
)

var GatewayMode_name = map[int32]string{
	0: "GW_INACTIVE",
	1: "GW_FREE_GATEWAYS_LIMITED",
	2: "GW_WHOLE_NETWORK",
	3: "GW_DELETED",
}

var GatewayMode_value = map[string]int32{
	"GW_INACTIVE":              0,
	"GW_FREE_GATEWAYS_LIMITED": 1,
	"GW_WHOLE_NETWORK":         2,
	"GW_DELETED":               3,
}

func (x GatewayMode) String() string {
	return proto.EnumName(GatewayMode_name, int32(x))
}

func (GatewayMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

type GetGatewayListRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGatewayListRequest) Reset()         { *m = GetGatewayListRequest{} }
func (m *GetGatewayListRequest) String() string { return proto.CompactTextString(m) }
func (*GetGatewayListRequest) ProtoMessage()    {}
func (*GetGatewayListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

func (m *GetGatewayListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayListRequest.Unmarshal(m, b)
}
func (m *GetGatewayListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayListRequest.Marshal(b, m, deterministic)
}
func (m *GetGatewayListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayListRequest.Merge(m, src)
}
func (m *GetGatewayListRequest) XXX_Size() int {
	return xxx_messageInfo_GetGatewayListRequest.Size(m)
}
func (m *GetGatewayListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayListRequest proto.InternalMessageInfo

func (m *GetGatewayListRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetGatewayListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetGatewayListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetGatewayListResponse struct {
	GwProfile            []*GatewayProfile `protobuf:"bytes,1,rep,name=gw_profile,json=gwProfile,proto3" json:"gw_profile,omitempty"`
	UserProfile          *ProfileResponse  `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetGatewayListResponse) Reset()         { *m = GetGatewayListResponse{} }
func (m *GetGatewayListResponse) String() string { return proto.CompactTextString(m) }
func (*GetGatewayListResponse) ProtoMessage()    {}
func (*GetGatewayListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

func (m *GetGatewayListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayListResponse.Unmarshal(m, b)
}
func (m *GetGatewayListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayListResponse.Marshal(b, m, deterministic)
}
func (m *GetGatewayListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayListResponse.Merge(m, src)
}
func (m *GetGatewayListResponse) XXX_Size() int {
	return xxx_messageInfo_GetGatewayListResponse.Size(m)
}
func (m *GetGatewayListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayListResponse proto.InternalMessageInfo

func (m *GetGatewayListResponse) GetGwProfile() []*GatewayProfile {
	if m != nil {
		return m.GwProfile
	}
	return nil
}

func (m *GetGatewayListResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetGatewayProfileRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	GwId                 int64    `protobuf:"varint,2,opt,name=gw_id,json=gwId,proto3" json:"gw_id,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGatewayProfileRequest) Reset()         { *m = GetGatewayProfileRequest{} }
func (m *GetGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetGatewayProfileRequest) ProtoMessage()    {}
func (*GetGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

func (m *GetGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayProfileRequest.Unmarshal(m, b)
}
func (m *GetGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayProfileRequest.Merge(m, src)
}
func (m *GetGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetGatewayProfileRequest.Size(m)
}
func (m *GetGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayProfileRequest proto.InternalMessageInfo

func (m *GetGatewayProfileRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetGatewayProfileRequest) GetGwId() int64 {
	if m != nil {
		return m.GwId
	}
	return 0
}

func (m *GetGatewayProfileRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetGatewayProfileRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GatewayProfile struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mac                  string   `protobuf:"bytes,2,opt,name=mac,proto3" json:"mac,omitempty"`
	FkGwNs               int64    `protobuf:"varint,3,opt,name=fk_gw_ns,json=fkGwNs,proto3" json:"fk_gw_ns,omitempty"`
	FkWallet             int64    `protobuf:"varint,4,opt,name=fk_wallet,json=fkWallet,proto3" json:"fk_wallet,omitempty"`
	Mode                 string   `protobuf:"bytes,5,opt,name=mode,proto3" json:"mode,omitempty"`
	CreateAt             string   `protobuf:"bytes,6,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	LastSeenAt           string   `protobuf:"bytes,7,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
	OrgId                int64    `protobuf:"varint,8,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Description          string   `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Name                 string   `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayProfile) Reset()         { *m = GatewayProfile{} }
func (m *GatewayProfile) String() string { return proto.CompactTextString(m) }
func (*GatewayProfile) ProtoMessage()    {}
func (*GatewayProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{3}
}

func (m *GatewayProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayProfile.Unmarshal(m, b)
}
func (m *GatewayProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayProfile.Marshal(b, m, deterministic)
}
func (m *GatewayProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayProfile.Merge(m, src)
}
func (m *GatewayProfile) XXX_Size() int {
	return xxx_messageInfo_GatewayProfile.Size(m)
}
func (m *GatewayProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayProfile.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayProfile proto.InternalMessageInfo

func (m *GatewayProfile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GatewayProfile) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *GatewayProfile) GetFkGwNs() int64 {
	if m != nil {
		return m.FkGwNs
	}
	return 0
}

func (m *GatewayProfile) GetFkWallet() int64 {
	if m != nil {
		return m.FkWallet
	}
	return 0
}

func (m *GatewayProfile) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *GatewayProfile) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *GatewayProfile) GetLastSeenAt() string {
	if m != nil {
		return m.LastSeenAt
	}
	return ""
}

func (m *GatewayProfile) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GatewayProfile) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GatewayProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetGatewayProfileResponse struct {
	GwProfile            *GatewayProfile  `protobuf:"bytes,1,opt,name=gw_profile,json=gwProfile,proto3" json:"gw_profile,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetGatewayProfileResponse) Reset()         { *m = GetGatewayProfileResponse{} }
func (m *GetGatewayProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetGatewayProfileResponse) ProtoMessage()    {}
func (*GetGatewayProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{4}
}

func (m *GetGatewayProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayProfileResponse.Unmarshal(m, b)
}
func (m *GetGatewayProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetGatewayProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayProfileResponse.Merge(m, src)
}
func (m *GetGatewayProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetGatewayProfileResponse.Size(m)
}
func (m *GetGatewayProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayProfileResponse proto.InternalMessageInfo

func (m *GetGatewayProfileResponse) GetGwProfile() *GatewayProfile {
	if m != nil {
		return m.GwProfile
	}
	return nil
}

func (m *GetGatewayProfileResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetGatewayHistoryRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	GwId                 int64    `protobuf:"varint,2,opt,name=gw_id,json=gwId,proto3" json:"gw_id,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGatewayHistoryRequest) Reset()         { *m = GetGatewayHistoryRequest{} }
func (m *GetGatewayHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetGatewayHistoryRequest) ProtoMessage()    {}
func (*GetGatewayHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{5}
}

func (m *GetGatewayHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayHistoryRequest.Unmarshal(m, b)
}
func (m *GetGatewayHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetGatewayHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayHistoryRequest.Merge(m, src)
}
func (m *GetGatewayHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetGatewayHistoryRequest.Size(m)
}
func (m *GetGatewayHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayHistoryRequest proto.InternalMessageInfo

func (m *GetGatewayHistoryRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetGatewayHistoryRequest) GetGwId() int64 {
	if m != nil {
		return m.GwId
	}
	return 0
}

func (m *GetGatewayHistoryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetGatewayHistoryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetGatewayHistoryResponse struct {
	GwHistory            string           `protobuf:"bytes,1,opt,name=gw_history,json=gwHistory,proto3" json:"gw_history,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetGatewayHistoryResponse) Reset()         { *m = GetGatewayHistoryResponse{} }
func (m *GetGatewayHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetGatewayHistoryResponse) ProtoMessage()    {}
func (*GetGatewayHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{6}
}

func (m *GetGatewayHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayHistoryResponse.Unmarshal(m, b)
}
func (m *GetGatewayHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetGatewayHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayHistoryResponse.Merge(m, src)
}
func (m *GetGatewayHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetGatewayHistoryResponse.Size(m)
}
func (m *GetGatewayHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayHistoryResponse proto.InternalMessageInfo

func (m *GetGatewayHistoryResponse) GetGwHistory() string {
	if m != nil {
		return m.GwHistory
	}
	return ""
}

func (m *GetGatewayHistoryResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type SetGatewayModeRequest struct {
	OrgId                int64       `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	GwId                 int64       `protobuf:"varint,2,opt,name=gw_id,json=gwId,proto3" json:"gw_id,omitempty"`
	GwMode               GatewayMode `protobuf:"varint,3,opt,name=gw_mode,json=gwMode,proto3,enum=api.GatewayMode" json:"gw_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetGatewayModeRequest) Reset()         { *m = SetGatewayModeRequest{} }
func (m *SetGatewayModeRequest) String() string { return proto.CompactTextString(m) }
func (*SetGatewayModeRequest) ProtoMessage()    {}
func (*SetGatewayModeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{7}
}

func (m *SetGatewayModeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGatewayModeRequest.Unmarshal(m, b)
}
func (m *SetGatewayModeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGatewayModeRequest.Marshal(b, m, deterministic)
}
func (m *SetGatewayModeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGatewayModeRequest.Merge(m, src)
}
func (m *SetGatewayModeRequest) XXX_Size() int {
	return xxx_messageInfo_SetGatewayModeRequest.Size(m)
}
func (m *SetGatewayModeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGatewayModeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetGatewayModeRequest proto.InternalMessageInfo

func (m *SetGatewayModeRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *SetGatewayModeRequest) GetGwId() int64 {
	if m != nil {
		return m.GwId
	}
	return 0
}

func (m *SetGatewayModeRequest) GetGwMode() GatewayMode {
	if m != nil {
		return m.GwMode
	}
	return GatewayMode_GW_INACTIVE
}

type SetGatewayModeResponse struct {
	Status               bool             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SetGatewayModeResponse) Reset()         { *m = SetGatewayModeResponse{} }
func (m *SetGatewayModeResponse) String() string { return proto.CompactTextString(m) }
func (*SetGatewayModeResponse) ProtoMessage()    {}
func (*SetGatewayModeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{8}
}

func (m *SetGatewayModeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGatewayModeResponse.Unmarshal(m, b)
}
func (m *SetGatewayModeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGatewayModeResponse.Marshal(b, m, deterministic)
}
func (m *SetGatewayModeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGatewayModeResponse.Merge(m, src)
}
func (m *SetGatewayModeResponse) XXX_Size() int {
	return xxx_messageInfo_SetGatewayModeResponse.Size(m)
}
func (m *SetGatewayModeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGatewayModeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetGatewayModeResponse proto.InternalMessageInfo

func (m *SetGatewayModeResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SetGatewayModeResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.GatewayMode", GatewayMode_name, GatewayMode_value)
	proto.RegisterType((*GetGatewayListRequest)(nil), "api.GetGatewayListRequest")
	proto.RegisterType((*GetGatewayListResponse)(nil), "api.GetGatewayListResponse")
	proto.RegisterType((*GetGatewayProfileRequest)(nil), "api.GetGatewayProfileRequest")
	proto.RegisterType((*GatewayProfile)(nil), "api.GatewayProfile")
	proto.RegisterType((*GetGatewayProfileResponse)(nil), "api.GetGatewayProfileResponse")
	proto.RegisterType((*GetGatewayHistoryRequest)(nil), "api.GetGatewayHistoryRequest")
	proto.RegisterType((*GetGatewayHistoryResponse)(nil), "api.GetGatewayHistoryResponse")
	proto.RegisterType((*SetGatewayModeRequest)(nil), "api.SetGatewayModeRequest")
	proto.RegisterType((*SetGatewayModeResponse)(nil), "api.SetGatewayModeResponse")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xe1, 0x6a, 0x1a, 0x4b,
	0x14, 0xbe, 0xeb, 0xaa, 0xd1, 0x63, 0xe2, 0xf5, 0x4e, 0x12, 0xd9, 0x6b, 0x92, 0x8b, 0x2c, 0xf7,
	0x47, 0x2a, 0x24, 0x82, 0xa5, 0x14, 0xfa, 0x4f, 0x9a, 0xed, 0x46, 0x6a, 0x92, 0xb2, 0x4a, 0x97,
	0x42, 0x61, 0xd8, 0xe8, 0xb8, 0x5d, 0xd4, 0x5d, 0xbb, 0x33, 0x76, 0x09, 0x21, 0x50, 0x0a, 0x85,
	0xe6, 0x77, 0xdf, 0xa0, 0xaf, 0xd4, 0x57, 0xe8, 0x83, 0x94, 0x99, 0x1d, 0xed, 0x6a, 0x4c, 0x08,
	0x81, 0xf6, 0x97, 0x33, 0x67, 0xce, 0x9c, 0xef, 0xfb, 0xf6, 0x7c, 0x67, 0x84, 0x0d, 0xd7, 0x61,
	0x24, 0x72, 0x2e, 0x0e, 0x27, 0x61, 0xc0, 0x02, 0xa4, 0x3a, 0x13, 0xaf, 0xb2, 0xeb, 0x06, 0x81,
	0x3b, 0x22, 0x75, 0x67, 0xe2, 0xd5, 0x1d, 0xdf, 0x0f, 0x98, 0xc3, 0xbc, 0xc0, 0xa7, 0x71, 0x4a,
	0x65, 0x63, 0x12, 0x06, 0x03, 0x6f, 0x44, 0xe2, 0xad, 0xfe, 0x16, 0xb6, 0x4d, 0xc2, 0xcc, 0xb8,
	0x4a, 0xdb, 0xa3, 0xcc, 0x22, 0xef, 0xa7, 0x84, 0x32, 0xb4, 0x0d, 0xd9, 0x20, 0x74, 0xb1, 0xd7,
	0xd7, 0x94, 0xaa, 0xb2, 0xaf, 0x5a, 0x99, 0x20, 0x74, 0x5b, 0x7d, 0x54, 0x86, 0x6c, 0x30, 0x18,
	0x50, 0xc2, 0xb4, 0x94, 0x08, 0xcb, 0x1d, 0xda, 0x82, 0xcc, 0xc8, 0x1b, 0x7b, 0x4c, 0x53, 0xe3,
	0x6c, 0xb1, 0xd1, 0x3f, 0x2b, 0x50, 0x5e, 0x2e, 0x4f, 0x27, 0x81, 0x4f, 0x09, 0x6a, 0x00, 0xb8,
	0x11, 0x96, 0x64, 0x34, 0xa5, 0xaa, 0xee, 0x17, 0x1a, 0x9b, 0x87, 0xce, 0xc4, 0x3b, 0x94, 0xd9,
	0xaf, 0xe2, 0x23, 0x2b, 0xef, 0x46, 0x72, 0x89, 0x9e, 0xc2, 0xfa, 0x94, 0x92, 0x70, 0x7e, 0x8b,
	0x53, 0x28, 0x34, 0xb6, 0xc4, 0xad, 0x59, 0xba, 0xac, 0x6f, 0x15, 0x78, 0xa6, 0x0c, 0xea, 0x0c,
	0xb4, 0x5f, 0x34, 0xe6, 0x99, 0x77, 0x0a, 0xdd, 0x84, 0x8c, 0x1b, 0xf1, 0x68, 0xac, 0x33, 0xed,
	0x46, 0x0b, 0xea, 0xd5, 0xd5, 0xea, 0xd3, 0x49, 0xf5, 0xd7, 0x29, 0x28, 0x2e, 0x62, 0xa2, 0x22,
	0xa4, 0xe6, 0x40, 0x29, 0xaf, 0x8f, 0x4a, 0xa0, 0x8e, 0x9d, 0x9e, 0xc0, 0xc8, 0x5b, 0x7c, 0x89,
	0x34, 0xc8, 0x0d, 0x86, 0xd8, 0x8d, 0xb0, 0x4f, 0x67, 0x20, 0x83, 0xa1, 0x19, 0x9d, 0x52, 0xb4,
	0x03, 0xf9, 0xc1, 0x10, 0x47, 0xce, 0x68, 0x44, 0x66, 0x40, 0xb9, 0xc1, 0xd0, 0x16, 0x7b, 0x84,
	0x20, 0x3d, 0x0e, 0xfa, 0x44, 0xcb, 0x88, 0x4a, 0x62, 0xcd, 0x2f, 0xf4, 0x42, 0xe2, 0x30, 0x82,
	0x1d, 0xa6, 0x65, 0xc5, 0x41, 0x2e, 0x0e, 0x34, 0x19, 0xaa, 0xc2, 0xfa, 0xc8, 0xa1, 0x0c, 0x53,
	0x42, 0x7c, 0x7e, 0xbe, 0x26, 0xce, 0x81, 0xc7, 0x3a, 0x84, 0xf8, 0xcd, 0xe4, 0x87, 0xc9, 0x25,
	0x3f, 0x4c, 0x15, 0x0a, 0x7d, 0x42, 0x7b, 0xa1, 0x37, 0xe1, 0xb6, 0xd2, 0xf2, 0xe2, 0x5e, 0x32,
	0xc4, 0xb9, 0xf8, 0xce, 0x98, 0x68, 0x10, 0x73, 0xe1, 0x6b, 0xfd, 0x8b, 0x02, 0xff, 0xae, 0x68,
	0xc1, 0x2d, 0x66, 0x50, 0xfe, 0x98, 0x19, 0x8e, 0x3d, 0xca, 0x82, 0xf0, 0xe2, 0xf7, 0x9b, 0x81,
	0x26, 0xf5, 0xcf, 0x51, 0xa5, 0xfe, 0x3d, 0xa1, 0xff, 0x5d, 0x1c, 0x15, 0xd0, 0x79, 0x2e, 0x55,
	0xa6, 0x3d, 0x5c, 0xaa, 0x0f, 0xdb, 0x9d, 0x39, 0xe8, 0x49, 0xd0, 0x7f, 0x90, 0xe9, 0x1f, 0xc1,
	0x9a, 0x1b, 0x61, 0xe1, 0x2e, 0x2e, 0xb4, 0xd8, 0x28, 0x25, 0x3b, 0x23, 0xaa, 0x66, 0xdd, 0x88,
	0xff, 0xea, 0x1e, 0x94, 0x97, 0xf1, 0xa4, 0xc2, 0x32, 0x64, 0x29, 0x73, 0xd8, 0x94, 0x0a, 0xc0,
	0x9c, 0x25, 0x77, 0x37, 0xa4, 0xa9, 0xf7, 0x94, 0x56, 0x3b, 0x87, 0x42, 0x02, 0x07, 0xfd, 0x0d,
	0x05, 0xd3, 0xc6, 0xad, 0xd3, 0xe6, 0xf3, 0x6e, 0xeb, 0xb5, 0x51, 0xfa, 0x0b, 0xed, 0x82, 0x66,
	0xda, 0xf8, 0x85, 0x65, 0x18, 0xd8, 0x6c, 0x76, 0x0d, 0xbb, 0xf9, 0xa6, 0x83, 0xdb, 0xad, 0x93,
	0x56, 0xd7, 0x38, 0x2a, 0x29, 0x68, 0x0b, 0x4a, 0xa6, 0x8d, 0xed, 0xe3, 0xb3, 0xb6, 0x81, 0x4f,
	0x8d, 0xae, 0x7d, 0x66, 0xbd, 0x2c, 0xa5, 0x50, 0x11, 0xc0, 0xb4, 0xf1, 0x91, 0xd1, 0x36, 0x78,
	0x96, 0xda, 0xf8, 0x96, 0x9e, 0x0f, 0x70, 0x87, 0x84, 0x1f, 0xbc, 0x1e, 0x41, 0x53, 0x28, 0x2e,
	0x3e, 0x68, 0xa8, 0x12, 0x7f, 0x8d, 0x55, 0x8f, 0x68, 0x65, 0x67, 0xe5, 0x59, 0x2c, 0x47, 0xaf,
	0x7d, 0xfa, 0xfe, 0xe3, 0x6b, 0xea, 0x7f, 0xa4, 0x8b, 0x97, 0x5a, 0x3e, 0xe4, 0xf5, 0xcb, 0xb8,
	0x2f, 0x57, 0xb3, 0xc0, 0xc1, 0x88, 0x83, 0x5c, 0x2b, 0xf0, 0xcf, 0x8d, 0xf1, 0x41, 0x7b, 0x4b,
	0xe5, 0x17, 0x5f, 0xb6, 0xca, 0x7f, 0xb7, 0x1d, 0x4b, 0x02, 0x4f, 0x04, 0x81, 0x3a, 0x3a, 0xb8,
	0x9b, 0x80, 0x6c, 0x4d, 0xfd, 0x52, 0x58, 0xe3, 0x6a, 0x89, 0xcb, 0xcc, 0xa3, 0xcb, 0x5c, 0x16,
	0x07, 0xeb, 0x06, 0x97, 0xa5, 0x09, 0xb8, 0x2f, 0x17, 0x39, 0x22, 0x73, 0x2e, 0x1f, 0x15, 0x28,
	0x2e, 0x3a, 0x4e, 0xf6, 0x63, 0xa5, 0xed, 0x65, 0x3f, 0x56, 0x5b, 0x74, 0x46, 0x41, 0xaf, 0xdd,
	0x4d, 0x81, 0x0f, 0xc2, 0x0c, 0xff, 0x99, 0x52, 0x3b, 0xcf, 0x8a, 0x3f, 0xd2, 0xc7, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x33, 0x1f, 0xde, 0xf0, 0x8b, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayServiceClient interface {
	GetGatewayList(ctx context.Context, in *GetGatewayListRequest, opts ...grpc.CallOption) (*GetGatewayListResponse, error)
	GetGatewayProfile(ctx context.Context, in *GetGatewayProfileRequest, opts ...grpc.CallOption) (*GetGatewayProfileResponse, error)
	GetGatewayHistory(ctx context.Context, in *GetGatewayHistoryRequest, opts ...grpc.CallOption) (*GetGatewayHistoryResponse, error)
	SetGatewayMode(ctx context.Context, in *SetGatewayModeRequest, opts ...grpc.CallOption) (*SetGatewayModeResponse, error)
}

type gatewayServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayServiceClient(cc *grpc.ClientConn) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) GetGatewayList(ctx context.Context, in *GetGatewayListRequest, opts ...grpc.CallOption) (*GetGatewayListResponse, error) {
	out := new(GetGatewayListResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayService/GetGatewayList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetGatewayProfile(ctx context.Context, in *GetGatewayProfileRequest, opts ...grpc.CallOption) (*GetGatewayProfileResponse, error) {
	out := new(GetGatewayProfileResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayService/GetGatewayProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetGatewayHistory(ctx context.Context, in *GetGatewayHistoryRequest, opts ...grpc.CallOption) (*GetGatewayHistoryResponse, error) {
	out := new(GetGatewayHistoryResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayService/GetGatewayHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) SetGatewayMode(ctx context.Context, in *SetGatewayModeRequest, opts ...grpc.CallOption) (*SetGatewayModeResponse, error) {
	out := new(SetGatewayModeResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayService/SetGatewayMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	GetGatewayList(context.Context, *GetGatewayListRequest) (*GetGatewayListResponse, error)
	GetGatewayProfile(context.Context, *GetGatewayProfileRequest) (*GetGatewayProfileResponse, error)
	GetGatewayHistory(context.Context, *GetGatewayHistoryRequest) (*GetGatewayHistoryResponse, error)
	SetGatewayMode(context.Context, *SetGatewayModeRequest) (*SetGatewayModeResponse, error)
}

// UnimplementedGatewayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (*UnimplementedGatewayServiceServer) GetGatewayList(ctx context.Context, req *GetGatewayListRequest) (*GetGatewayListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayList not implemented")
}
func (*UnimplementedGatewayServiceServer) GetGatewayProfile(ctx context.Context, req *GetGatewayProfileRequest) (*GetGatewayProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayProfile not implemented")
}
func (*UnimplementedGatewayServiceServer) GetGatewayHistory(ctx context.Context, req *GetGatewayHistoryRequest) (*GetGatewayHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayHistory not implemented")
}
func (*UnimplementedGatewayServiceServer) SetGatewayMode(ctx context.Context, req *SetGatewayModeRequest) (*SetGatewayModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGatewayMode not implemented")
}

func RegisterGatewayServiceServer(s *grpc.Server, srv GatewayServiceServer) {
	s.RegisterService(&_GatewayService_serviceDesc, srv)
}

func _GatewayService_GetGatewayList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetGatewayList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayService/GetGatewayList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetGatewayList(ctx, req.(*GetGatewayListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetGatewayProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetGatewayProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayService/GetGatewayProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetGatewayProfile(ctx, req.(*GetGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetGatewayHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetGatewayHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayService/GetGatewayHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetGatewayHistory(ctx, req.(*GetGatewayHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_SetGatewayMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGatewayModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).SetGatewayMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayService/SetGatewayMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).SetGatewayMode(ctx, req.(*SetGatewayModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayList",
			Handler:    _GatewayService_GetGatewayList_Handler,
		},
		{
			MethodName: "GetGatewayProfile",
			Handler:    _GatewayService_GetGatewayProfile_Handler,
		},
		{
			MethodName: "GetGatewayHistory",
			Handler:    _GatewayService_GetGatewayHistory_Handler,
		},
		{
			MethodName: "SetGatewayMode",
			Handler:    _GatewayService_SetGatewayMode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
