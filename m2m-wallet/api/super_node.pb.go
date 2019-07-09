// Code generated by protoc-gen-go. DO NOT EDIT.
// source: super_node.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type GetSuperNodeActiveMoneyAccountRequest struct {
	MoneyAbbr            Money    `protobuf:"varint,1,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSuperNodeActiveMoneyAccountRequest) Reset()         { *m = GetSuperNodeActiveMoneyAccountRequest{} }
func (m *GetSuperNodeActiveMoneyAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetSuperNodeActiveMoneyAccountRequest) ProtoMessage()    {}
func (*GetSuperNodeActiveMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e142dc5bc4ebd3, []int{0}
}

func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.Unmarshal(m, b)
}
func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.Merge(m, src)
}
func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.Size(m)
}
func (m *GetSuperNodeActiveMoneyAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSuperNodeActiveMoneyAccountRequest proto.InternalMessageInfo

func (m *GetSuperNodeActiveMoneyAccountRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

type GetSuperNodeActiveMoneyAccountResponse struct {
	SupernodeActiveAccount string           `protobuf:"bytes,1,opt,name=supernode_active_account,json=supernodeActiveAccount,proto3" json:"supernode_active_account,omitempty"`
	Error                  string           `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	UserProfile            *ProfileResponse `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}         `json:"-"`
	XXX_unrecognized       []byte           `json:"-"`
	XXX_sizecache          int32            `json:"-"`
}

func (m *GetSuperNodeActiveMoneyAccountResponse) Reset() {
	*m = GetSuperNodeActiveMoneyAccountResponse{}
}
func (m *GetSuperNodeActiveMoneyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetSuperNodeActiveMoneyAccountResponse) ProtoMessage()    {}
func (*GetSuperNodeActiveMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e142dc5bc4ebd3, []int{1}
}

func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.Unmarshal(m, b)
}
func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.Marshal(b, m, deterministic)
}
func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.Merge(m, src)
}
func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_Size() int {
	return xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.Size(m)
}
func (m *GetSuperNodeActiveMoneyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSuperNodeActiveMoneyAccountResponse proto.InternalMessageInfo

func (m *GetSuperNodeActiveMoneyAccountResponse) GetSupernodeActiveAccount() string {
	if m != nil {
		return m.SupernodeActiveAccount
	}
	return ""
}

func (m *GetSuperNodeActiveMoneyAccountResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetSuperNodeActiveMoneyAccountResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type AddSuperNodeMoneyAccountRequest struct {
	MoneyAbbr            Money    `protobuf:"varint,1,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	AccountAddr          string   `protobuf:"bytes,2,opt,name=account_addr,json=accountAddr,proto3" json:"account_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSuperNodeMoneyAccountRequest) Reset()         { *m = AddSuperNodeMoneyAccountRequest{} }
func (m *AddSuperNodeMoneyAccountRequest) String() string { return proto.CompactTextString(m) }
func (*AddSuperNodeMoneyAccountRequest) ProtoMessage()    {}
func (*AddSuperNodeMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e142dc5bc4ebd3, []int{2}
}

func (m *AddSuperNodeMoneyAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSuperNodeMoneyAccountRequest.Unmarshal(m, b)
}
func (m *AddSuperNodeMoneyAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSuperNodeMoneyAccountRequest.Marshal(b, m, deterministic)
}
func (m *AddSuperNodeMoneyAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSuperNodeMoneyAccountRequest.Merge(m, src)
}
func (m *AddSuperNodeMoneyAccountRequest) XXX_Size() int {
	return xxx_messageInfo_AddSuperNodeMoneyAccountRequest.Size(m)
}
func (m *AddSuperNodeMoneyAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSuperNodeMoneyAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddSuperNodeMoneyAccountRequest proto.InternalMessageInfo

func (m *AddSuperNodeMoneyAccountRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

func (m *AddSuperNodeMoneyAccountRequest) GetAccountAddr() string {
	if m != nil {
		return m.AccountAddr
	}
	return ""
}

type AddSuperNodeMoneyAccountResponse struct {
	Error                string           `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Status               bool             `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AddSuperNodeMoneyAccountResponse) Reset()         { *m = AddSuperNodeMoneyAccountResponse{} }
func (m *AddSuperNodeMoneyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*AddSuperNodeMoneyAccountResponse) ProtoMessage()    {}
func (*AddSuperNodeMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e142dc5bc4ebd3, []int{3}
}

func (m *AddSuperNodeMoneyAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSuperNodeMoneyAccountResponse.Unmarshal(m, b)
}
func (m *AddSuperNodeMoneyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSuperNodeMoneyAccountResponse.Marshal(b, m, deterministic)
}
func (m *AddSuperNodeMoneyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSuperNodeMoneyAccountResponse.Merge(m, src)
}
func (m *AddSuperNodeMoneyAccountResponse) XXX_Size() int {
	return xxx_messageInfo_AddSuperNodeMoneyAccountResponse.Size(m)
}
func (m *AddSuperNodeMoneyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSuperNodeMoneyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddSuperNodeMoneyAccountResponse proto.InternalMessageInfo

func (m *AddSuperNodeMoneyAccountResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AddSuperNodeMoneyAccountResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *AddSuperNodeMoneyAccountResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSuperNodeActiveMoneyAccountRequest)(nil), "api.GetSuperNodeActiveMoneyAccountRequest")
	proto.RegisterType((*GetSuperNodeActiveMoneyAccountResponse)(nil), "api.GetSuperNodeActiveMoneyAccountResponse")
	proto.RegisterType((*AddSuperNodeMoneyAccountRequest)(nil), "api.AddSuperNodeMoneyAccountRequest")
	proto.RegisterType((*AddSuperNodeMoneyAccountResponse)(nil), "api.AddSuperNodeMoneyAccountResponse")
}

func init() { proto.RegisterFile("super_node.proto", fileDescriptor_02e142dc5bc4ebd3) }

var fileDescriptor_02e142dc5bc4ebd3 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0xa5, 0x67, 0xd9, 0xc5, 0xad, 0x2c, 0xcb, 0xd2, 0x0c, 0x4b, 0x08, 0xa2, 0x63, 0x70, 0x75,
	0x1d, 0x75, 0x02, 0xf1, 0xa0, 0x78, 0xcb, 0x41, 0x3c, 0x29, 0x92, 0xf9, 0x80, 0xd0, 0x49, 0x97,
	0x43, 0xc3, 0xd8, 0x1d, 0xbb, 0x3b, 0x03, 0x22, 0x5e, 0xbc, 0x79, 0xf6, 0xe8, 0x37, 0x78, 0xf0,
	0x5b, 0xfc, 0x01, 0x0f, 0x7e, 0x88, 0xa4, 0xbb, 0x27, 0x3a, 0xc2, 0x38, 0x83, 0x7b, 0x7c, 0xd5,
	0x2f, 0xaf, 0xde, 0xab, 0xaa, 0xc0, 0x99, 0xe9, 0x5a, 0xd4, 0x95, 0x54, 0x1c, 0x67, 0xad, 0x56,
	0x56, 0xd1, 0x03, 0xd6, 0x8a, 0xe4, 0xfa, 0x42, 0xa9, 0xc5, 0x12, 0x33, 0xd6, 0x8a, 0x8c, 0x49,
	0xa9, 0x2c, 0xb3, 0x42, 0x49, 0xe3, 0x29, 0x49, 0xf4, 0x46, 0x49, 0x7c, 0x17, 0xc0, 0xa9, 0x90,
	0x16, 0xb5, 0x64, 0x4b, 0x8f, 0xd3, 0x12, 0x2e, 0x9e, 0xa3, 0x9d, 0xf7, 0xb2, 0x2f, 0x15, 0xc7,
	0xa2, 0xb1, 0x62, 0x85, 0x2f, 0xfa, 0x0f, 0x8a, 0xa6, 0x51, 0x9d, 0xb4, 0x25, 0xbe, 0xed, 0xd0,
	0x58, 0x7a, 0x0f, 0xc0, 0xe9, 0x54, 0xac, 0xae, 0x75, 0x4c, 0x26, 0xe4, 0xf2, 0x34, 0x87, 0x19,
	0x6b, 0xc5, 0xcc, 0xb1, 0xcb, 0x63, 0xf7, 0x5a, 0xd4, 0xb5, 0x4e, 0xbf, 0x11, 0xb8, 0xb3, 0x4b,
	0xd4, 0xb4, 0x4a, 0x1a, 0xa4, 0x4f, 0x20, 0x76, 0x91, 0xfa, 0x44, 0x15, 0x73, 0xbc, 0x8a, 0x79,
	0x8e, 0xeb, 0x71, 0x5c, 0x9e, 0x0f, 0xef, 0x5e, 0x26, 0x28, 0xd0, 0x31, 0x1c, 0x3e, 0xd3, 0x5a,
	0xe9, 0x78, 0xe4, 0x68, 0x1e, 0xd0, 0xc7, 0x70, 0xd2, 0x19, 0xd4, 0x55, 0xab, 0xd5, 0x6b, 0xb1,
	0xc4, 0xf8, 0x60, 0x42, 0x2e, 0xa3, 0x7c, 0xec, 0x7c, 0xbe, 0xf2, 0xb5, 0x75, 0xef, 0x32, 0xea,
	0x99, 0xa1, 0x98, 0x2a, 0xb8, 0x59, 0x70, 0x3e, 0x58, 0xbe, 0xda, 0x04, 0xe8, 0x2d, 0x38, 0x09,
	0x29, 0x2a, 0xc6, 0xf9, 0xda, 0x63, 0x14, 0x6a, 0x05, 0xe7, 0x3a, 0xfd, 0x44, 0x60, 0xb2, 0xbd,
	0x63, 0x18, 0xcf, 0x18, 0x0e, 0xd1, 0x85, 0xf4, 0xb3, 0xf0, 0x80, 0x9e, 0xc3, 0x91, 0xb1, 0xcc,
	0x76, 0xc6, 0xe9, 0x5e, 0x2b, 0x03, 0xfa, 0xef, 0xf0, 0xf9, 0x8f, 0x11, 0x9c, 0x0d, 0x46, 0xe6,
	0xa8, 0x57, 0xa2, 0x41, 0xfa, 0x95, 0xc0, 0x8d, 0x7f, 0x6f, 0x91, 0x4e, 0x9d, 0xf4, 0x5e, 0xf7,
	0x93, 0xdc, 0xdf, 0x8b, 0xeb, 0xdd, 0xa5, 0xf9, 0xc7, 0xef, 0x3f, 0x3f, 0x8f, 0x1e, 0xd0, 0xa9,
	0x3b, 0xe9, 0xe1, 0x02, 0xb2, 0xcd, 0x0b, 0xc9, 0xde, 0xff, 0xde, 0xc7, 0x07, 0xfa, 0x85, 0x40,
	0xbc, 0x6d, 0xa0, 0xf4, 0xb6, 0xeb, 0xbe, 0x63, 0xc3, 0xc9, 0xc5, 0x0e, 0xd6, 0xa6, 0xbb, 0xf4,
	0xee, 0x5f, 0xee, 0xfe, 0xb4, 0x93, 0x31, 0xce, 0x1f, 0x06, 0x9f, 0x4f, 0xc9, 0xb4, 0x3e, 0x72,
	0xbf, 0xdb, 0xa3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0x69, 0xc4, 0xc1, 0xc2, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SuperNodeServiceClient is the client API for SuperNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SuperNodeServiceClient interface {
	GetSuperNodeActiveMoneyAccount(ctx context.Context, in *GetSuperNodeActiveMoneyAccountRequest, opts ...grpc.CallOption) (*GetSuperNodeActiveMoneyAccountResponse, error)
	AddSuperNodeMoneyAccount(ctx context.Context, in *AddSuperNodeMoneyAccountRequest, opts ...grpc.CallOption) (*AddSuperNodeMoneyAccountResponse, error)
}

type superNodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewSuperNodeServiceClient(cc *grpc.ClientConn) SuperNodeServiceClient {
	return &superNodeServiceClient{cc}
}

func (c *superNodeServiceClient) GetSuperNodeActiveMoneyAccount(ctx context.Context, in *GetSuperNodeActiveMoneyAccountRequest, opts ...grpc.CallOption) (*GetSuperNodeActiveMoneyAccountResponse, error) {
	out := new(GetSuperNodeActiveMoneyAccountResponse)
	err := c.cc.Invoke(ctx, "/api.SuperNodeService/GetSuperNodeActiveMoneyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superNodeServiceClient) AddSuperNodeMoneyAccount(ctx context.Context, in *AddSuperNodeMoneyAccountRequest, opts ...grpc.CallOption) (*AddSuperNodeMoneyAccountResponse, error) {
	out := new(AddSuperNodeMoneyAccountResponse)
	err := c.cc.Invoke(ctx, "/api.SuperNodeService/AddSuperNodeMoneyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuperNodeServiceServer is the server API for SuperNodeService service.
type SuperNodeServiceServer interface {
	GetSuperNodeActiveMoneyAccount(context.Context, *GetSuperNodeActiveMoneyAccountRequest) (*GetSuperNodeActiveMoneyAccountResponse, error)
	AddSuperNodeMoneyAccount(context.Context, *AddSuperNodeMoneyAccountRequest) (*AddSuperNodeMoneyAccountResponse, error)
}

func RegisterSuperNodeServiceServer(s *grpc.Server, srv SuperNodeServiceServer) {
	s.RegisterService(&_SuperNodeService_serviceDesc, srv)
}

func _SuperNodeService_GetSuperNodeActiveMoneyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSuperNodeActiveMoneyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperNodeServiceServer).GetSuperNodeActiveMoneyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SuperNodeService/GetSuperNodeActiveMoneyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperNodeServiceServer).GetSuperNodeActiveMoneyAccount(ctx, req.(*GetSuperNodeActiveMoneyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperNodeService_AddSuperNodeMoneyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSuperNodeMoneyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperNodeServiceServer).AddSuperNodeMoneyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SuperNodeService/AddSuperNodeMoneyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperNodeServiceServer).AddSuperNodeMoneyAccount(ctx, req.(*AddSuperNodeMoneyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SuperNodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SuperNodeService",
	HandlerType: (*SuperNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSuperNodeActiveMoneyAccount",
			Handler:    _SuperNodeService_GetSuperNodeActiveMoneyAccount_Handler,
		},
		{
			MethodName: "AddSuperNodeMoneyAccount",
			Handler:    _SuperNodeService_AddSuperNodeMoneyAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "super_node.proto",
}
