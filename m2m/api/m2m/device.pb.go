// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

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

type DeviceMode int32

const (
	DeviceMode_DV_INACTIVE              DeviceMode = 0
	DeviceMode_DV_FREE_GATEWAYS_LIMITED DeviceMode = 1
	DeviceMode_DV_WHOLE_NETWORK         DeviceMode = 2
	DeviceMode_DV_DELETED               DeviceMode = 3
)

var DeviceMode_name = map[int32]string{
	0: "DV_INACTIVE",
	1: "DV_FREE_GATEWAYS_LIMITED",
	2: "DV_WHOLE_NETWORK",
	3: "DV_DELETED",
}

var DeviceMode_value = map[string]int32{
	"DV_INACTIVE":              0,
	"DV_FREE_GATEWAYS_LIMITED": 1,
	"DV_WHOLE_NETWORK":         2,
	"DV_DELETED":               3,
}

func (x DeviceMode) String() string {
	return proto.EnumName(DeviceMode_name, int32(x))
}

func (DeviceMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{0}
}

type GetDeviceListRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceListRequest) Reset()         { *m = GetDeviceListRequest{} }
func (m *GetDeviceListRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceListRequest) ProtoMessage()    {}
func (*GetDeviceListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{0}
}

func (m *GetDeviceListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceListRequest.Unmarshal(m, b)
}
func (m *GetDeviceListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceListRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceListRequest.Merge(m, src)
}
func (m *GetDeviceListRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceListRequest.Size(m)
}
func (m *GetDeviceListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceListRequest proto.InternalMessageInfo

func (m *GetDeviceListRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetDeviceListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetDeviceListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetDeviceListResponse struct {
	DevProfile           []*DeviceProfile `protobuf:"bytes,1,rep,name=dev_profile,json=devProfile,proto3" json:"dev_profile,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetDeviceListResponse) Reset()         { *m = GetDeviceListResponse{} }
func (m *GetDeviceListResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceListResponse) ProtoMessage()    {}
func (*GetDeviceListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{1}
}

func (m *GetDeviceListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceListResponse.Unmarshal(m, b)
}
func (m *GetDeviceListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceListResponse.Marshal(b, m, deterministic)
}
func (m *GetDeviceListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceListResponse.Merge(m, src)
}
func (m *GetDeviceListResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceListResponse.Size(m)
}
func (m *GetDeviceListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceListResponse proto.InternalMessageInfo

func (m *GetDeviceListResponse) GetDevProfile() []*DeviceProfile {
	if m != nil {
		return m.DevProfile
	}
	return nil
}

func (m *GetDeviceListResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetDeviceProfileRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	DevId                int64    `protobuf:"varint,2,opt,name=dev_id,json=devId,proto3" json:"dev_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceProfileRequest) Reset()         { *m = GetDeviceProfileRequest{} }
func (m *GetDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProfileRequest) ProtoMessage()    {}
func (*GetDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{2}
}

func (m *GetDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProfileRequest.Unmarshal(m, b)
}
func (m *GetDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProfileRequest.Merge(m, src)
}
func (m *GetDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProfileRequest.Size(m)
}
func (m *GetDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProfileRequest proto.InternalMessageInfo

func (m *GetDeviceProfileRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetDeviceProfileRequest) GetDevId() int64 {
	if m != nil {
		return m.DevId
	}
	return 0
}

type DeviceProfile struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DevEui               string   `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	FkWallet             int64    `protobuf:"varint,3,opt,name=fk_wallet,json=fkWallet,proto3" json:"fk_wallet,omitempty"`
	Mode                 string   `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastSeenAt           string   `protobuf:"bytes,6,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
	ApplicationId        int64    `protobuf:"varint,7,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProfile) Reset()         { *m = DeviceProfile{} }
func (m *DeviceProfile) String() string { return proto.CompactTextString(m) }
func (*DeviceProfile) ProtoMessage()    {}
func (*DeviceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{3}
}

func (m *DeviceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProfile.Unmarshal(m, b)
}
func (m *DeviceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProfile.Marshal(b, m, deterministic)
}
func (m *DeviceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProfile.Merge(m, src)
}
func (m *DeviceProfile) XXX_Size() int {
	return xxx_messageInfo_DeviceProfile.Size(m)
}
func (m *DeviceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProfile proto.InternalMessageInfo

func (m *DeviceProfile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeviceProfile) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

func (m *DeviceProfile) GetFkWallet() int64 {
	if m != nil {
		return m.FkWallet
	}
	return 0
}

func (m *DeviceProfile) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *DeviceProfile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *DeviceProfile) GetLastSeenAt() string {
	if m != nil {
		return m.LastSeenAt
	}
	return ""
}

func (m *DeviceProfile) GetApplicationId() int64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *DeviceProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetDeviceProfileResponse struct {
	DevProfile           *DeviceProfile   `protobuf:"bytes,1,opt,name=dev_profile,json=devProfile,proto3" json:"dev_profile,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetDeviceProfileResponse) Reset()         { *m = GetDeviceProfileResponse{} }
func (m *GetDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProfileResponse) ProtoMessage()    {}
func (*GetDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{4}
}

func (m *GetDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProfileResponse.Unmarshal(m, b)
}
func (m *GetDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProfileResponse.Merge(m, src)
}
func (m *GetDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProfileResponse.Size(m)
}
func (m *GetDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProfileResponse proto.InternalMessageInfo

func (m *GetDeviceProfileResponse) GetDevProfile() *DeviceProfile {
	if m != nil {
		return m.DevProfile
	}
	return nil
}

func (m *GetDeviceProfileResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetDeviceHistoryRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	DevId                int64    `protobuf:"varint,2,opt,name=dev_id,json=devId,proto3" json:"dev_id,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceHistoryRequest) Reset()         { *m = GetDeviceHistoryRequest{} }
func (m *GetDeviceHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceHistoryRequest) ProtoMessage()    {}
func (*GetDeviceHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{5}
}

func (m *GetDeviceHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceHistoryRequest.Unmarshal(m, b)
}
func (m *GetDeviceHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceHistoryRequest.Merge(m, src)
}
func (m *GetDeviceHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceHistoryRequest.Size(m)
}
func (m *GetDeviceHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceHistoryRequest proto.InternalMessageInfo

func (m *GetDeviceHistoryRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetDeviceHistoryRequest) GetDevId() int64 {
	if m != nil {
		return m.DevId
	}
	return 0
}

func (m *GetDeviceHistoryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetDeviceHistoryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetDeviceHistoryResponse struct {
	DevHistory           string           `protobuf:"bytes,1,opt,name=dev_history,json=devHistory,proto3" json:"dev_history,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetDeviceHistoryResponse) Reset()         { *m = GetDeviceHistoryResponse{} }
func (m *GetDeviceHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceHistoryResponse) ProtoMessage()    {}
func (*GetDeviceHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{6}
}

func (m *GetDeviceHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceHistoryResponse.Unmarshal(m, b)
}
func (m *GetDeviceHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetDeviceHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceHistoryResponse.Merge(m, src)
}
func (m *GetDeviceHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceHistoryResponse.Size(m)
}
func (m *GetDeviceHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceHistoryResponse proto.InternalMessageInfo

func (m *GetDeviceHistoryResponse) GetDevHistory() string {
	if m != nil {
		return m.DevHistory
	}
	return ""
}

func (m *GetDeviceHistoryResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type SetDeviceModeRequest struct {
	OrgId                int64      `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	DevId                int64      `protobuf:"varint,2,opt,name=dev_id,json=devId,proto3" json:"dev_id,omitempty"`
	DevMode              DeviceMode `protobuf:"varint,3,opt,name=dev_mode,json=devMode,proto3,enum=api.DeviceMode" json:"dev_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SetDeviceModeRequest) Reset()         { *m = SetDeviceModeRequest{} }
func (m *SetDeviceModeRequest) String() string { return proto.CompactTextString(m) }
func (*SetDeviceModeRequest) ProtoMessage()    {}
func (*SetDeviceModeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{7}
}

func (m *SetDeviceModeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeviceModeRequest.Unmarshal(m, b)
}
func (m *SetDeviceModeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeviceModeRequest.Marshal(b, m, deterministic)
}
func (m *SetDeviceModeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeviceModeRequest.Merge(m, src)
}
func (m *SetDeviceModeRequest) XXX_Size() int {
	return xxx_messageInfo_SetDeviceModeRequest.Size(m)
}
func (m *SetDeviceModeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeviceModeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeviceModeRequest proto.InternalMessageInfo

func (m *SetDeviceModeRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *SetDeviceModeRequest) GetDevId() int64 {
	if m != nil {
		return m.DevId
	}
	return 0
}

func (m *SetDeviceModeRequest) GetDevMode() DeviceMode {
	if m != nil {
		return m.DevMode
	}
	return DeviceMode_DV_INACTIVE
}

type SetDeviceModeResponse struct {
	Status               bool             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SetDeviceModeResponse) Reset()         { *m = SetDeviceModeResponse{} }
func (m *SetDeviceModeResponse) String() string { return proto.CompactTextString(m) }
func (*SetDeviceModeResponse) ProtoMessage()    {}
func (*SetDeviceModeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{8}
}

func (m *SetDeviceModeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeviceModeResponse.Unmarshal(m, b)
}
func (m *SetDeviceModeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeviceModeResponse.Marshal(b, m, deterministic)
}
func (m *SetDeviceModeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeviceModeResponse.Merge(m, src)
}
func (m *SetDeviceModeResponse) XXX_Size() int {
	return xxx_messageInfo_SetDeviceModeResponse.Size(m)
}
func (m *SetDeviceModeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeviceModeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeviceModeResponse proto.InternalMessageInfo

func (m *SetDeviceModeResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SetDeviceModeResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.DeviceMode", DeviceMode_name, DeviceMode_value)
	proto.RegisterType((*GetDeviceListRequest)(nil), "api.GetDeviceListRequest")
	proto.RegisterType((*GetDeviceListResponse)(nil), "api.GetDeviceListResponse")
	proto.RegisterType((*GetDeviceProfileRequest)(nil), "api.GetDeviceProfileRequest")
	proto.RegisterType((*DeviceProfile)(nil), "api.DeviceProfile")
	proto.RegisterType((*GetDeviceProfileResponse)(nil), "api.GetDeviceProfileResponse")
	proto.RegisterType((*GetDeviceHistoryRequest)(nil), "api.GetDeviceHistoryRequest")
	proto.RegisterType((*GetDeviceHistoryResponse)(nil), "api.GetDeviceHistoryResponse")
	proto.RegisterType((*SetDeviceModeRequest)(nil), "api.SetDeviceModeRequest")
	proto.RegisterType((*SetDeviceModeResponse)(nil), "api.SetDeviceModeResponse")
}

func init() { proto.RegisterFile("device.proto", fileDescriptor_870276a56ac00da5) }

var fileDescriptor_870276a56ac00da5 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x6a, 0xdb, 0x3c,
	0x14, 0xfe, 0x1d, 0x37, 0x69, 0x72, 0xd2, 0xa4, 0x41, 0x24, 0xad, 0xff, 0xac, 0x65, 0xc1, 0xa3,
	0x50, 0x3a, 0xd6, 0x42, 0x7a, 0x31, 0xd8, 0x5d, 0x58, 0xbc, 0xd6, 0x2c, 0x6d, 0x87, 0x13, 0x12,
	0xc6, 0x2e, 0x8c, 0x57, 0x2b, 0xa9, 0xa8, 0x6b, 0x79, 0x96, 0x92, 0x31, 0xb6, 0xdd, 0x0c, 0x06,
	0xbb, 0xdf, 0xa3, 0xed, 0x15, 0xb6, 0x17, 0xd8, 0x13, 0x0c, 0xc9, 0x4a, 0x16, 0xa7, 0x69, 0x29,
	0x85, 0x5d, 0x59, 0x3a, 0x3a, 0xe7, 0x7c, 0xfa, 0xce, 0xf9, 0x74, 0x0c, 0x6b, 0x3e, 0x9e, 0x90,
	0x73, 0xbc, 0x1f, 0xc5, 0x94, 0x53, 0xa4, 0x7b, 0x11, 0xa9, 0x6f, 0x8d, 0x28, 0x1d, 0x05, 0xf8,
	0xc0, 0x8b, 0xc8, 0x81, 0x17, 0x86, 0x94, 0x7b, 0x9c, 0xd0, 0x90, 0x25, 0x2e, 0xf5, 0x52, 0x14,
	0xd3, 0x21, 0x09, 0x54, 0x84, 0xf9, 0x06, 0xaa, 0x47, 0x98, 0xb7, 0x65, 0x92, 0x0e, 0x61, 0xdc,
	0xc1, 0xef, 0xc6, 0x98, 0x71, 0x54, 0x83, 0x1c, 0x8d, 0x47, 0x2e, 0xf1, 0x0d, 0xad, 0xa1, 0xed,
	0xea, 0x4e, 0x96, 0xc6, 0x23, 0xdb, 0x47, 0x1b, 0x90, 0xa3, 0xc3, 0x21, 0xc3, 0xdc, 0xc8, 0x48,
	0xb3, 0xda, 0xa1, 0x2a, 0x64, 0x03, 0x72, 0x45, 0xb8, 0xa1, 0x27, 0xde, 0x72, 0x63, 0x7e, 0xd5,
	0xa0, 0xb6, 0x90, 0x9d, 0x45, 0x34, 0x64, 0x18, 0x1d, 0x42, 0xd1, 0xc7, 0x13, 0x57, 0xdd, 0xc5,
	0xd0, 0x1a, 0xfa, 0x6e, 0xb1, 0x89, 0xf6, 0xbd, 0x88, 0xec, 0x27, 0xde, 0xaf, 0x92, 0x13, 0x07,
	0x7c, 0x3c, 0x51, 0x6b, 0xf4, 0x14, 0xd6, 0xc6, 0x0c, 0xc7, 0xb3, 0x28, 0x71, 0x85, 0x62, 0xb3,
	0x2a, 0xa3, 0xa6, 0xfe, 0x0a, 0xc0, 0x29, 0x0a, 0x4f, 0x65, 0x34, 0x8f, 0x60, 0x73, 0x76, 0x8d,
	0x99, 0xe3, 0xad, 0x3c, 0x6b, 0x90, 0x13, 0xf7, 0x23, 0xbe, 0xe2, 0x99, 0xf5, 0xf1, 0xc4, 0xf6,
	0xcd, 0x5f, 0x1a, 0x94, 0x52, 0x69, 0x50, 0x19, 0x32, 0xb3, 0xd8, 0x0c, 0xf1, 0xd1, 0x26, 0xac,
	0x8a, 0x40, 0x3c, 0x26, 0x32, 0xb2, 0xe0, 0x88, 0x3c, 0xd6, 0x98, 0xa0, 0x07, 0x50, 0x18, 0x5e,
	0xba, 0xef, 0xbd, 0x20, 0xc0, 0xd3, 0x2a, 0xe5, 0x87, 0x97, 0x03, 0xb9, 0x47, 0x08, 0x56, 0xae,
	0xa8, 0x8f, 0x8d, 0x15, 0x19, 0x22, 0xd7, 0x68, 0x1b, 0xe0, 0x3c, 0xc6, 0x1e, 0xc7, 0xbe, 0xeb,
	0x71, 0x23, 0x2b, 0x4f, 0x0a, 0xca, 0xd2, 0xe2, 0xa8, 0x01, 0x6b, 0x81, 0xc7, 0xb8, 0xcb, 0x30,
	0x0e, 0x85, 0x43, 0x4e, 0x3a, 0x80, 0xb0, 0x75, 0x31, 0x0e, 0x5b, 0x1c, 0xed, 0x40, 0xd9, 0x8b,
	0xa2, 0x80, 0x9c, 0xcb, 0xfe, 0x0b, 0x2e, 0xab, 0x12, 0xb6, 0x34, 0x67, 0xb5, 0x7d, 0x81, 0x1d,
	0x7a, 0x57, 0xd8, 0xc8, 0x27, 0xd8, 0x62, 0x6d, 0x7e, 0xd3, 0xc0, 0xb8, 0x5e, 0xb1, 0x9b, 0x7a,
	0xa7, 0xfd, 0xcb, 0xde, 0x8d, 0xe7, 0x7a, 0x77, 0x4c, 0x18, 0xa7, 0xf1, 0x87, 0x7b, 0xf5, 0x6e,
	0x4e, 0xba, 0xfa, 0x72, 0xe9, 0xae, 0xcc, 0x4b, 0x97, 0xcf, 0x15, 0x60, 0x06, 0xab, 0x0a, 0xf0,
	0x30, 0x29, 0xc0, 0x45, 0x62, 0x96, 0xe0, 0x05, 0x49, 0x56, 0x39, 0xde, 0x9f, 0x6c, 0x04, 0xd5,
	0xee, 0x14, 0xf5, 0x84, 0xfa, 0xf7, 0x53, 0x29, 0xda, 0x83, 0xbc, 0x30, 0x4b, 0x45, 0x09, 0xae,
	0xe5, 0xe6, 0xfa, 0x5c, 0x77, 0x64, 0x5e, 0x21, 0x52, 0xb1, 0x30, 0x2f, 0xa0, 0xb6, 0x80, 0xa8,
	0x48, 0x6e, 0x40, 0x8e, 0x71, 0x8f, 0x8f, 0x99, 0x84, 0xcc, 0x3b, 0x6a, 0x77, 0x8d, 0x9b, 0x7e,
	0x47, 0x6e, 0x7b, 0x1e, 0xc0, 0x5f, 0x18, 0xb4, 0x0e, 0xc5, 0x76, 0xdf, 0xb5, 0x4f, 0x5b, 0xcf,
	0x7b, 0x76, 0xdf, 0xaa, 0xfc, 0x87, 0xb6, 0xc0, 0x68, 0xf7, 0xdd, 0x17, 0x8e, 0x65, 0xb9, 0x47,
	0xad, 0x9e, 0x35, 0x68, 0xbd, 0xee, 0xba, 0x1d, 0xfb, 0xc4, 0xee, 0x59, 0xed, 0x8a, 0x86, 0xaa,
	0x50, 0x69, 0xf7, 0xdd, 0xc1, 0xf1, 0x59, 0xc7, 0x72, 0x4f, 0xad, 0xde, 0xe0, 0xcc, 0x79, 0x59,
	0xc9, 0xa0, 0x32, 0x40, 0xbb, 0xef, 0xb6, 0xad, 0x8e, 0x25, 0xbc, 0xf4, 0xe6, 0x6f, 0x7d, 0xfa,
	0x3c, 0xbb, 0x38, 0x16, 0x1f, 0x14, 0x41, 0x29, 0x35, 0x80, 0xd0, 0xff, 0xf2, 0xa2, 0xcb, 0x46,
	0x5e, 0xbd, 0xbe, 0xec, 0x28, 0x61, 0x62, 0xee, 0x7e, 0xf9, 0xf1, 0xf3, 0x7b, 0xc6, 0x44, 0x0d,
	0x39, 0x55, 0x93, 0x99, 0x7b, 0xf0, 0x31, 0xe9, 0xc9, 0x67, 0xb5, 0x7f, 0x12, 0x08, 0x80, 0x4f,
	0x50, 0x59, 0x7c, 0x39, 0x68, 0x2b, 0x9d, 0x39, 0x3d, 0x82, 0xea, 0xdb, 0x37, 0x9c, 0x2a, 0xe8,
	0xc7, 0x12, 0x7a, 0x07, 0x3d, 0xba, 0x0d, 0x5a, 0x75, 0x23, 0x85, 0x3e, 0x55, 0xe3, 0x02, 0x7a,
	0xfa, 0x11, 0x2d, 0xa2, 0x2f, 0x68, 0xfd, 0x6e, 0xe8, 0xea, 0x25, 0x20, 0x06, 0xa5, 0x94, 0x98,
	0x54, 0xb5, 0x97, 0x49, 0x5a, 0x55, 0x7b, 0xa9, 0xf6, 0xa6, 0xa0, 0xe6, 0xad, 0xd5, 0x16, 0xf2,
	0x7e, 0xa6, 0xed, 0xbd, 0xcd, 0xc9, 0x1f, 0xd9, 0xe1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7,
	0x1d, 0xdb, 0x4b, 0x0a, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceServiceClient interface {
	GetDeviceList(ctx context.Context, in *GetDeviceListRequest, opts ...grpc.CallOption) (*GetDeviceListResponse, error)
	GetDeviceProfile(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error)
	GetDeviceHistory(ctx context.Context, in *GetDeviceHistoryRequest, opts ...grpc.CallOption) (*GetDeviceHistoryResponse, error)
	SetDeviceMode(ctx context.Context, in *SetDeviceModeRequest, opts ...grpc.CallOption) (*SetDeviceModeResponse, error)
}

type deviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceServiceClient(cc *grpc.ClientConn) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) GetDeviceList(ctx context.Context, in *GetDeviceListRequest, opts ...grpc.CallOption) (*GetDeviceListResponse, error) {
	out := new(GetDeviceListResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/GetDeviceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDeviceProfile(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error) {
	out := new(GetDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/GetDeviceProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDeviceHistory(ctx context.Context, in *GetDeviceHistoryRequest, opts ...grpc.CallOption) (*GetDeviceHistoryResponse, error) {
	out := new(GetDeviceHistoryResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/GetDeviceHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) SetDeviceMode(ctx context.Context, in *SetDeviceModeRequest, opts ...grpc.CallOption) (*SetDeviceModeResponse, error) {
	out := new(SetDeviceModeResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/SetDeviceMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
type DeviceServiceServer interface {
	GetDeviceList(context.Context, *GetDeviceListRequest) (*GetDeviceListResponse, error)
	GetDeviceProfile(context.Context, *GetDeviceProfileRequest) (*GetDeviceProfileResponse, error)
	GetDeviceHistory(context.Context, *GetDeviceHistoryRequest) (*GetDeviceHistoryResponse, error)
	SetDeviceMode(context.Context, *SetDeviceModeRequest) (*SetDeviceModeResponse, error)
}

// UnimplementedDeviceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (*UnimplementedDeviceServiceServer) GetDeviceList(ctx context.Context, req *GetDeviceListRequest) (*GetDeviceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceList not implemented")
}
func (*UnimplementedDeviceServiceServer) GetDeviceProfile(ctx context.Context, req *GetDeviceProfileRequest) (*GetDeviceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceProfile not implemented")
}
func (*UnimplementedDeviceServiceServer) GetDeviceHistory(ctx context.Context, req *GetDeviceHistoryRequest) (*GetDeviceHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceHistory not implemented")
}
func (*UnimplementedDeviceServiceServer) SetDeviceMode(ctx context.Context, req *SetDeviceModeRequest) (*SetDeviceModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceMode not implemented")
}

func RegisterDeviceServiceServer(s *grpc.Server, srv DeviceServiceServer) {
	s.RegisterService(&_DeviceService_serviceDesc, srv)
}

func _DeviceService_GetDeviceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDeviceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/GetDeviceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDeviceList(ctx, req.(*GetDeviceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDeviceProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDeviceProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/GetDeviceProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDeviceProfile(ctx, req.(*GetDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDeviceHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDeviceHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/GetDeviceHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDeviceHistory(ctx, req.(*GetDeviceHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_SetDeviceMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).SetDeviceMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/SetDeviceMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).SetDeviceMode(ctx, req.(*SetDeviceModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceList",
			Handler:    _DeviceService_GetDeviceList_Handler,
		},
		{
			MethodName: "GetDeviceProfile",
			Handler:    _DeviceService_GetDeviceProfile_Handler,
		},
		{
			MethodName: "GetDeviceHistory",
			Handler:    _DeviceService_GetDeviceHistory_Handler,
		},
		{
			MethodName: "SetDeviceMode",
			Handler:    _DeviceService_SetDeviceMode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device.proto",
}
