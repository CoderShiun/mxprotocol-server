// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ext_account.proto

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

type Money int32

const (
	Money_Ether Money = 0
)

var Money_name = map[int32]string{
	0: "Ether",
}

var Money_value = map[string]int32{
	"Ether": 0,
}

func (x Money) String() string {
	return proto.EnumName(Money_name, int32(x))
}

func (Money) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{0}
}

type ModifyMoneyAccountRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	MoneyAbbr            Money    `protobuf:"varint,2,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	CurrentAccount       string   `protobuf:"bytes,3,opt,name=current_account,json=currentAccount,proto3" json:"current_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyMoneyAccountRequest) Reset()         { *m = ModifyMoneyAccountRequest{} }
func (m *ModifyMoneyAccountRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyMoneyAccountRequest) ProtoMessage()    {}
func (*ModifyMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{0}
}

func (m *ModifyMoneyAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyMoneyAccountRequest.Unmarshal(m, b)
}
func (m *ModifyMoneyAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyMoneyAccountRequest.Marshal(b, m, deterministic)
}
func (m *ModifyMoneyAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyMoneyAccountRequest.Merge(m, src)
}
func (m *ModifyMoneyAccountRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyMoneyAccountRequest.Size(m)
}
func (m *ModifyMoneyAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyMoneyAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyMoneyAccountRequest proto.InternalMessageInfo

func (m *ModifyMoneyAccountRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *ModifyMoneyAccountRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

func (m *ModifyMoneyAccountRequest) GetCurrentAccount() string {
	if m != nil {
		return m.CurrentAccount
	}
	return ""
}

type ModifyMoneyAccountResponse struct {
	Status               bool             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ModifyMoneyAccountResponse) Reset()         { *m = ModifyMoneyAccountResponse{} }
func (m *ModifyMoneyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyMoneyAccountResponse) ProtoMessage()    {}
func (*ModifyMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{1}
}

func (m *ModifyMoneyAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyMoneyAccountResponse.Unmarshal(m, b)
}
func (m *ModifyMoneyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyMoneyAccountResponse.Marshal(b, m, deterministic)
}
func (m *ModifyMoneyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyMoneyAccountResponse.Merge(m, src)
}
func (m *ModifyMoneyAccountResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyMoneyAccountResponse.Size(m)
}
func (m *ModifyMoneyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyMoneyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyMoneyAccountResponse proto.InternalMessageInfo

func (m *ModifyMoneyAccountResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ModifyMoneyAccountResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetMoneyAccountChangeHistoryRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	MoneyAbbr            Money    `protobuf:"varint,4,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMoneyAccountChangeHistoryRequest) Reset()         { *m = GetMoneyAccountChangeHistoryRequest{} }
func (m *GetMoneyAccountChangeHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetMoneyAccountChangeHistoryRequest) ProtoMessage()    {}
func (*GetMoneyAccountChangeHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{2}
}

func (m *GetMoneyAccountChangeHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.Unmarshal(m, b)
}
func (m *GetMoneyAccountChangeHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetMoneyAccountChangeHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.Merge(m, src)
}
func (m *GetMoneyAccountChangeHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.Size(m)
}
func (m *GetMoneyAccountChangeHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMoneyAccountChangeHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMoneyAccountChangeHistoryRequest proto.InternalMessageInfo

func (m *GetMoneyAccountChangeHistoryRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetMoneyAccountChangeHistoryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetMoneyAccountChangeHistoryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetMoneyAccountChangeHistoryRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

type MoneyAccountChangeHistory struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoneyAccountChangeHistory) Reset()         { *m = MoneyAccountChangeHistory{} }
func (m *MoneyAccountChangeHistory) String() string { return proto.CompactTextString(m) }
func (*MoneyAccountChangeHistory) ProtoMessage()    {}
func (*MoneyAccountChangeHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{3}
}

func (m *MoneyAccountChangeHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoneyAccountChangeHistory.Unmarshal(m, b)
}
func (m *MoneyAccountChangeHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoneyAccountChangeHistory.Marshal(b, m, deterministic)
}
func (m *MoneyAccountChangeHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoneyAccountChangeHistory.Merge(m, src)
}
func (m *MoneyAccountChangeHistory) XXX_Size() int {
	return xxx_messageInfo_MoneyAccountChangeHistory.Size(m)
}
func (m *MoneyAccountChangeHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_MoneyAccountChangeHistory.DiscardUnknown(m)
}

var xxx_messageInfo_MoneyAccountChangeHistory proto.InternalMessageInfo

func (m *MoneyAccountChangeHistory) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *MoneyAccountChangeHistory) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MoneyAccountChangeHistory) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type GetMoneyAccountChangeHistoryResponse struct {
	Count                int64                        `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	ChangeHistory        []*MoneyAccountChangeHistory `protobuf:"bytes,2,rep,name=change_history,json=changeHistory,proto3" json:"change_history,omitempty"`
	UserProfile          *ProfileResponse             `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetMoneyAccountChangeHistoryResponse) Reset()         { *m = GetMoneyAccountChangeHistoryResponse{} }
func (m *GetMoneyAccountChangeHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetMoneyAccountChangeHistoryResponse) ProtoMessage()    {}
func (*GetMoneyAccountChangeHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{4}
}

func (m *GetMoneyAccountChangeHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.Unmarshal(m, b)
}
func (m *GetMoneyAccountChangeHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetMoneyAccountChangeHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.Merge(m, src)
}
func (m *GetMoneyAccountChangeHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.Size(m)
}
func (m *GetMoneyAccountChangeHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMoneyAccountChangeHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMoneyAccountChangeHistoryResponse proto.InternalMessageInfo

func (m *GetMoneyAccountChangeHistoryResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetMoneyAccountChangeHistoryResponse) GetChangeHistory() []*MoneyAccountChangeHistory {
	if m != nil {
		return m.ChangeHistory
	}
	return nil
}

func (m *GetMoneyAccountChangeHistoryResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetActiveMoneyAccountRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	MoneyAbbr            Money    `protobuf:"varint,2,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetActiveMoneyAccountRequest) Reset()         { *m = GetActiveMoneyAccountRequest{} }
func (m *GetActiveMoneyAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetActiveMoneyAccountRequest) ProtoMessage()    {}
func (*GetActiveMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{5}
}

func (m *GetActiveMoneyAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetActiveMoneyAccountRequest.Unmarshal(m, b)
}
func (m *GetActiveMoneyAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetActiveMoneyAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetActiveMoneyAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetActiveMoneyAccountRequest.Merge(m, src)
}
func (m *GetActiveMoneyAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetActiveMoneyAccountRequest.Size(m)
}
func (m *GetActiveMoneyAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetActiveMoneyAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetActiveMoneyAccountRequest proto.InternalMessageInfo

func (m *GetActiveMoneyAccountRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetActiveMoneyAccountRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

type GetActiveMoneyAccountResponse struct {
	ActiveAccount        string           `protobuf:"bytes,1,opt,name=active_account,json=activeAccount,proto3" json:"active_account,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetActiveMoneyAccountResponse) Reset()         { *m = GetActiveMoneyAccountResponse{} }
func (m *GetActiveMoneyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetActiveMoneyAccountResponse) ProtoMessage()    {}
func (*GetActiveMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ddbf9312a3f483, []int{6}
}

func (m *GetActiveMoneyAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetActiveMoneyAccountResponse.Unmarshal(m, b)
}
func (m *GetActiveMoneyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetActiveMoneyAccountResponse.Marshal(b, m, deterministic)
}
func (m *GetActiveMoneyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetActiveMoneyAccountResponse.Merge(m, src)
}
func (m *GetActiveMoneyAccountResponse) XXX_Size() int {
	return xxx_messageInfo_GetActiveMoneyAccountResponse.Size(m)
}
func (m *GetActiveMoneyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetActiveMoneyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetActiveMoneyAccountResponse proto.InternalMessageInfo

func (m *GetActiveMoneyAccountResponse) GetActiveAccount() string {
	if m != nil {
		return m.ActiveAccount
	}
	return ""
}

func (m *GetActiveMoneyAccountResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.Money", Money_name, Money_value)
	proto.RegisterType((*ModifyMoneyAccountRequest)(nil), "api.ModifyMoneyAccountRequest")
	proto.RegisterType((*ModifyMoneyAccountResponse)(nil), "api.ModifyMoneyAccountResponse")
	proto.RegisterType((*GetMoneyAccountChangeHistoryRequest)(nil), "api.GetMoneyAccountChangeHistoryRequest")
	proto.RegisterType((*MoneyAccountChangeHistory)(nil), "api.MoneyAccountChangeHistory")
	proto.RegisterType((*GetMoneyAccountChangeHistoryResponse)(nil), "api.GetMoneyAccountChangeHistoryResponse")
	proto.RegisterType((*GetActiveMoneyAccountRequest)(nil), "api.GetActiveMoneyAccountRequest")
	proto.RegisterType((*GetActiveMoneyAccountResponse)(nil), "api.GetActiveMoneyAccountResponse")
}

func init() { proto.RegisterFile("ext_account.proto", fileDescriptor_d6ddbf9312a3f483) }

var fileDescriptor_d6ddbf9312a3f483 = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0x75, 0xb2, 0x4d, 0x30, 0xb7, 0x4d, 0xd4, 0x21, 0x2d, 0x71, 0x69, 0x35, 0xae, 0x8a, 0x69,
	0xa1, 0x09, 0x44, 0xa4, 0x50, 0x10, 0x0c, 0x52, 0xaa, 0x0f, 0x05, 0x59, 0x3f, 0x60, 0x9d, 0xec,
	0x4e, 0x92, 0x81, 0x64, 0x67, 0x9d, 0x9d, 0x94, 0x06, 0x29, 0x82, 0x0f, 0xfe, 0x80, 0x88, 0x4f,
	0xfe, 0x84, 0xcf, 0xfe, 0x85, 0xbf, 0xe0, 0x6f, 0x08, 0x92, 0x3b, 0xb3, 0x4d, 0x5a, 0x93, 0x18,
	0x04, 0xdf, 0x72, 0xcf, 0x9c, 0xcc, 0x3d, 0xf7, 0x9c, 0x3b, 0x0b, 0xb7, 0xf8, 0x99, 0x0e, 0x58,
	0x18, 0xca, 0x51, 0xac, 0x1b, 0x89, 0x92, 0x5a, 0x52, 0x87, 0x25, 0xc2, 0xdd, 0xee, 0x49, 0xd9,
	0x1b, 0xf0, 0x26, 0x4b, 0x44, 0x93, 0xc5, 0xb1, 0xd4, 0x4c, 0x0b, 0x19, 0xa7, 0x86, 0xe2, 0x96,
	0x12, 0x25, 0xbb, 0x62, 0xc0, 0x4d, 0xe9, 0x7d, 0x24, 0x70, 0xfb, 0x44, 0x46, 0xa2, 0x3b, 0x3e,
	0x91, 0x31, 0x1f, 0xb7, 0xcd, 0x75, 0x3e, 0x7f, 0x3b, 0xe2, 0xa9, 0xa6, 0x9b, 0x50, 0x90, 0xaa,
	0x17, 0x88, 0xa8, 0x4a, 0x6a, 0xa4, 0xee, 0xf8, 0x79, 0xa9, 0x7a, 0x2f, 0x23, 0xba, 0x0b, 0x30,
	0x9c, 0xb0, 0x03, 0xd6, 0xe9, 0xa8, 0x6a, 0xae, 0x46, 0xea, 0xe5, 0x16, 0x34, 0x58, 0x22, 0x1a,
	0x78, 0x89, 0x5f, 0xc4, 0xd3, 0x76, 0xa7, 0xa3, 0xe8, 0x23, 0xb8, 0x11, 0x8e, 0x94, 0xe2, 0xf1,
	0x85, 0xd4, 0xaa, 0x53, 0x23, 0xf5, 0xa2, 0x5f, 0xb6, 0xb0, 0xed, 0xe8, 0x0d, 0xc1, 0x9d, 0xa7,
	0x23, 0x4d, 0x64, 0x9c, 0x72, 0xba, 0x05, 0x85, 0x54, 0x33, 0x3d, 0x4a, 0x51, 0xc8, 0x75, 0xdf,
	0x56, 0xf4, 0x00, 0x36, 0x46, 0x29, 0x57, 0x81, 0x1d, 0x0a, 0xb5, 0xac, 0xb7, 0x2a, 0xa8, 0xe5,
	0x95, 0xc1, 0xb2, 0x3b, 0xfc, 0xf5, 0x09, 0xd3, 0x82, 0xde, 0x67, 0x02, 0xf7, 0x8f, 0xb9, 0x9e,
	0x6d, 0xf6, 0xbc, 0xcf, 0xe2, 0x1e, 0x7f, 0x21, 0x52, 0x2d, 0xd5, 0xf8, 0x2f, 0x0e, 0x6c, 0x41,
	0x41, 0x76, 0xbb, 0x29, 0xd7, 0xd8, 0xd1, 0xf1, 0x6d, 0x45, 0x2b, 0x90, 0x1f, 0x88, 0xa1, 0x30,
	0x43, 0x3a, 0xbe, 0x29, 0xae, 0xf8, 0xb5, 0xb6, 0xc4, 0x2f, 0xaf, 0x3b, 0x89, 0x63, 0x81, 0x26,
	0x4a, 0x61, 0x8d, 0x45, 0x91, 0x42, 0x29, 0x45, 0x1f, 0x7f, 0xcf, 0x38, 0x93, 0x43, 0x34, 0x73,
	0x66, 0x07, 0x20, 0x54, 0x9c, 0x69, 0x1e, 0x05, 0x2c, 0xf3, 0xbc, 0x68, 0x91, 0xb6, 0xf6, 0xbe,
	0x13, 0x78, 0xb0, 0x7c, 0x7e, 0xeb, 0x7c, 0x05, 0xf2, 0x26, 0x36, 0x3b, 0x3f, 0x16, 0xf4, 0x08,
	0xca, 0x21, 0xd2, 0x83, 0xbe, 0xe1, 0x57, 0x73, 0x35, 0xa7, 0xbe, 0xde, 0xba, 0x33, 0x9d, 0x6a,
	0xee, 0xad, 0xa5, 0xf0, 0xd2, 0x40, 0x57, 0xe3, 0x73, 0x56, 0x8d, 0xef, 0x0d, 0x6c, 0x1f, 0x73,
	0xdd, 0x0e, 0xb5, 0x38, 0xe5, 0xff, 0x65, 0x71, 0xbd, 0xf7, 0xb0, 0xb3, 0xa0, 0x83, 0x35, 0xe6,
	0x21, 0x94, 0x19, 0x9e, 0x5e, 0x2c, 0xb6, 0x89, 0xa5, 0x64, 0x50, 0x4b, 0xff, 0xe7, 0x0d, 0xdd,
	0xa3, 0x90, 0xc7, 0xbe, 0xb4, 0x08, 0xf9, 0x23, 0xdd, 0xe7, 0xea, 0xe6, 0xb5, 0xd6, 0x2f, 0x07,
	0x36, 0x10, 0x7c, 0xcd, 0xd5, 0xa9, 0x08, 0x39, 0xfd, 0x42, 0x80, 0xfe, 0xf9, 0x6c, 0x68, 0x16,
	0xc3, 0x82, 0x77, 0xed, 0xde, 0x5d, 0x78, 0x6e, 0x94, 0x78, 0xcf, 0x3e, 0xfc, 0xf8, 0xf9, 0x29,
	0x77, 0xe8, 0x3e, 0xc1, 0xaf, 0x08, 0x3f, 0xd3, 0xfb, 0x76, 0xc8, 0xe6, 0x3b, 0xe3, 0xeb, 0x79,
	0x73, 0x88, 0xff, 0x9e, 0xe2, 0x53, 0x63, 0xcf, 0x0f, 0xc9, 0x1e, 0xfd, 0x46, 0x30, 0x22, 0x13,
	0xff, 0x6c, 0x8f, 0x2c, 0xfb, 0x3a, 0x6a, 0x58, 0xe1, 0x0d, 0xba, 0xbb, 0x2b, 0x30, 0xad, 0xee,
	0xa7, 0xa8, 0xfb, 0x80, 0x2e, 0xd1, 0x6d, 0x36, 0x70, 0xdf, 0xee, 0xed, 0x25, 0xdd, 0xf4, 0x2b,
	0x81, 0xcd, 0xb9, 0xa9, 0xd3, 0x7b, 0x99, 0x86, 0x85, 0x3b, 0xe7, 0x7a, 0xcb, 0x28, 0xab, 0xeb,
	0x33, 0xeb, 0x33, 0xd7, 0xd7, 0x4e, 0x01, 0x3f, 0xda, 0x8f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x62, 0x85, 0xcd, 0x41, 0xfb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MoneyServiceClient is the client API for MoneyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoneyServiceClient interface {
	ModifyMoneyAccount(ctx context.Context, in *ModifyMoneyAccountRequest, opts ...grpc.CallOption) (*ModifyMoneyAccountResponse, error)
	GetChangeMoneyAccountHistory(ctx context.Context, in *GetMoneyAccountChangeHistoryRequest, opts ...grpc.CallOption) (*GetMoneyAccountChangeHistoryResponse, error)
	GetActiveMoneyAccount(ctx context.Context, in *GetActiveMoneyAccountRequest, opts ...grpc.CallOption) (*GetActiveMoneyAccountResponse, error)
}

type moneyServiceClient struct {
	cc *grpc.ClientConn
}

func NewMoneyServiceClient(cc *grpc.ClientConn) MoneyServiceClient {
	return &moneyServiceClient{cc}
}

func (c *moneyServiceClient) ModifyMoneyAccount(ctx context.Context, in *ModifyMoneyAccountRequest, opts ...grpc.CallOption) (*ModifyMoneyAccountResponse, error) {
	out := new(ModifyMoneyAccountResponse)
	err := c.cc.Invoke(ctx, "/api.MoneyService/ModifyMoneyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moneyServiceClient) GetChangeMoneyAccountHistory(ctx context.Context, in *GetMoneyAccountChangeHistoryRequest, opts ...grpc.CallOption) (*GetMoneyAccountChangeHistoryResponse, error) {
	out := new(GetMoneyAccountChangeHistoryResponse)
	err := c.cc.Invoke(ctx, "/api.MoneyService/GetChangeMoneyAccountHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moneyServiceClient) GetActiveMoneyAccount(ctx context.Context, in *GetActiveMoneyAccountRequest, opts ...grpc.CallOption) (*GetActiveMoneyAccountResponse, error) {
	out := new(GetActiveMoneyAccountResponse)
	err := c.cc.Invoke(ctx, "/api.MoneyService/GetActiveMoneyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoneyServiceServer is the server API for MoneyService service.
type MoneyServiceServer interface {
	ModifyMoneyAccount(context.Context, *ModifyMoneyAccountRequest) (*ModifyMoneyAccountResponse, error)
	GetChangeMoneyAccountHistory(context.Context, *GetMoneyAccountChangeHistoryRequest) (*GetMoneyAccountChangeHistoryResponse, error)
	GetActiveMoneyAccount(context.Context, *GetActiveMoneyAccountRequest) (*GetActiveMoneyAccountResponse, error)
}

// UnimplementedMoneyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMoneyServiceServer struct {
}

func (*UnimplementedMoneyServiceServer) ModifyMoneyAccount(ctx context.Context, req *ModifyMoneyAccountRequest) (*ModifyMoneyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyMoneyAccount not implemented")
}
func (*UnimplementedMoneyServiceServer) GetChangeMoneyAccountHistory(ctx context.Context, req *GetMoneyAccountChangeHistoryRequest) (*GetMoneyAccountChangeHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChangeMoneyAccountHistory not implemented")
}
func (*UnimplementedMoneyServiceServer) GetActiveMoneyAccount(ctx context.Context, req *GetActiveMoneyAccountRequest) (*GetActiveMoneyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveMoneyAccount not implemented")
}

func RegisterMoneyServiceServer(s *grpc.Server, srv MoneyServiceServer) {
	s.RegisterService(&_MoneyService_serviceDesc, srv)
}

func _MoneyService_ModifyMoneyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyMoneyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).ModifyMoneyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MoneyService/ModifyMoneyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).ModifyMoneyAccount(ctx, req.(*ModifyMoneyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoneyService_GetChangeMoneyAccountHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoneyAccountChangeHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).GetChangeMoneyAccountHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MoneyService/GetChangeMoneyAccountHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).GetChangeMoneyAccountHistory(ctx, req.(*GetMoneyAccountChangeHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoneyService_GetActiveMoneyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveMoneyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).GetActiveMoneyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MoneyService/GetActiveMoneyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).GetActiveMoneyAccount(ctx, req.(*GetActiveMoneyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MoneyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MoneyService",
	HandlerType: (*MoneyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModifyMoneyAccount",
			Handler:    _MoneyService_ModifyMoneyAccount_Handler,
		},
		{
			MethodName: "GetChangeMoneyAccountHistory",
			Handler:    _MoneyService_GetChangeMoneyAccountHistory_Handler,
		},
		{
			MethodName: "GetActiveMoneyAccount",
			Handler:    _MoneyService_GetActiveMoneyAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ext_account.proto",
}
