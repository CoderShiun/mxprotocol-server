// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topup.proto

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

type GetTopUpHistoryRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopUpHistoryRequest) Reset()         { *m = GetTopUpHistoryRequest{} }
func (m *GetTopUpHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopUpHistoryRequest) ProtoMessage()    {}
func (*GetTopUpHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{0}
}

func (m *GetTopUpHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpHistoryRequest.Unmarshal(m, b)
}
func (m *GetTopUpHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetTopUpHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpHistoryRequest.Merge(m, src)
}
func (m *GetTopUpHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopUpHistoryRequest.Size(m)
}
func (m *GetTopUpHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpHistoryRequest proto.InternalMessageInfo

func (m *GetTopUpHistoryRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetTopUpHistoryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetTopUpHistoryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type TopUpHistory struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	MoneyType            string   `protobuf:"bytes,5,opt,name=money_type,json=moneyType,proto3" json:"money_type,omitempty"`
	TxHash               string   `protobuf:"bytes,6,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopUpHistory) Reset()         { *m = TopUpHistory{} }
func (m *TopUpHistory) String() string { return proto.CompactTextString(m) }
func (*TopUpHistory) ProtoMessage()    {}
func (*TopUpHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{1}
}

func (m *TopUpHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopUpHistory.Unmarshal(m, b)
}
func (m *TopUpHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopUpHistory.Marshal(b, m, deterministic)
}
func (m *TopUpHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopUpHistory.Merge(m, src)
}
func (m *TopUpHistory) XXX_Size() int {
	return xxx_messageInfo_TopUpHistory.Size(m)
}
func (m *TopUpHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_TopUpHistory.DiscardUnknown(m)
}

var xxx_messageInfo_TopUpHistory proto.InternalMessageInfo

func (m *TopUpHistory) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TopUpHistory) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TopUpHistory) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TopUpHistory) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *TopUpHistory) GetMoneyType() string {
	if m != nil {
		return m.MoneyType
	}
	return ""
}

func (m *TopUpHistory) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

type GetTopUpHistoryResponse struct {
	Count                int64            `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	TopupHistory         []*TopUpHistory  `protobuf:"bytes,2,rep,name=topup_history,json=topupHistory,proto3" json:"topup_history,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetTopUpHistoryResponse) Reset()         { *m = GetTopUpHistoryResponse{} }
func (m *GetTopUpHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetTopUpHistoryResponse) ProtoMessage()    {}
func (*GetTopUpHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{2}
}

func (m *GetTopUpHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpHistoryResponse.Unmarshal(m, b)
}
func (m *GetTopUpHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetTopUpHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpHistoryResponse.Merge(m, src)
}
func (m *GetTopUpHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetTopUpHistoryResponse.Size(m)
}
func (m *GetTopUpHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpHistoryResponse proto.InternalMessageInfo

func (m *GetTopUpHistoryResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetTopUpHistoryResponse) GetTopupHistory() []*TopUpHistory {
	if m != nil {
		return m.TopupHistory
	}
	return nil
}

func (m *GetTopUpHistoryResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetTopUpDestinationRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	MoneyAbbr            Money    `protobuf:"varint,2,opt,name=money_abbr,json=moneyAbbr,proto3,enum=api.Money" json:"money_abbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopUpDestinationRequest) Reset()         { *m = GetTopUpDestinationRequest{} }
func (m *GetTopUpDestinationRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopUpDestinationRequest) ProtoMessage()    {}
func (*GetTopUpDestinationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{3}
}

func (m *GetTopUpDestinationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpDestinationRequest.Unmarshal(m, b)
}
func (m *GetTopUpDestinationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpDestinationRequest.Marshal(b, m, deterministic)
}
func (m *GetTopUpDestinationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpDestinationRequest.Merge(m, src)
}
func (m *GetTopUpDestinationRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopUpDestinationRequest.Size(m)
}
func (m *GetTopUpDestinationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpDestinationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpDestinationRequest proto.InternalMessageInfo

func (m *GetTopUpDestinationRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetTopUpDestinationRequest) GetMoneyAbbr() Money {
	if m != nil {
		return m.MoneyAbbr
	}
	return Money_Ether
}

type GetTopUpDestinationResponse struct {
	ActiveAccount        string           `protobuf:"bytes,1,opt,name=active_account,json=activeAccount,proto3" json:"active_account,omitempty"`
	UserProfile          *ProfileResponse `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetTopUpDestinationResponse) Reset()         { *m = GetTopUpDestinationResponse{} }
func (m *GetTopUpDestinationResponse) String() string { return proto.CompactTextString(m) }
func (*GetTopUpDestinationResponse) ProtoMessage()    {}
func (*GetTopUpDestinationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{4}
}

func (m *GetTopUpDestinationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpDestinationResponse.Unmarshal(m, b)
}
func (m *GetTopUpDestinationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpDestinationResponse.Marshal(b, m, deterministic)
}
func (m *GetTopUpDestinationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpDestinationResponse.Merge(m, src)
}
func (m *GetTopUpDestinationResponse) XXX_Size() int {
	return xxx_messageInfo_GetTopUpDestinationResponse.Size(m)
}
func (m *GetTopUpDestinationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpDestinationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpDestinationResponse proto.InternalMessageInfo

func (m *GetTopUpDestinationResponse) GetActiveAccount() string {
	if m != nil {
		return m.ActiveAccount
	}
	return ""
}

func (m *GetTopUpDestinationResponse) GetUserProfile() *ProfileResponse {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTopUpHistoryRequest)(nil), "api.GetTopUpHistoryRequest")
	proto.RegisterType((*TopUpHistory)(nil), "api.TopUpHistory")
	proto.RegisterType((*GetTopUpHistoryResponse)(nil), "api.GetTopUpHistoryResponse")
	proto.RegisterType((*GetTopUpDestinationRequest)(nil), "api.GetTopUpDestinationRequest")
	proto.RegisterType((*GetTopUpDestinationResponse)(nil), "api.GetTopUpDestinationResponse")
}

func init() { proto.RegisterFile("topup.proto", fileDescriptor_8eec749941d0cb6c) }

var fileDescriptor_8eec749941d0cb6c = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8f, 0x12, 0x41,
	0x10, 0xcd, 0x0c, 0x0b, 0x86, 0xe6, 0xc3, 0x6c, 0xbb, 0xee, 0x4e, 0x58, 0x54, 0x32, 0xd1, 0x64,
	0x3d, 0x2c, 0x24, 0x6c, 0xd4, 0x83, 0x27, 0x12, 0x13, 0xd7, 0x83, 0x89, 0x19, 0xd7, 0xa3, 0x4e,
	0x1a, 0x68, 0xa0, 0x13, 0x98, 0x6a, 0xbb, 0x6b, 0x36, 0x90, 0x0d, 0x17, 0xff, 0x82, 0xf1, 0xe2,
	0x49, 0x7f, 0x93, 0x7f, 0xc1, 0x1f, 0x62, 0xa8, 0xee, 0x89, 0xae, 0x60, 0xdc, 0x1b, 0xf5, 0xea,
	0xf1, 0xaa, 0xea, 0xbd, 0x1e, 0x56, 0x43, 0xd0, 0xb9, 0xee, 0x6a, 0x03, 0x08, 0xbc, 0x24, 0xb4,
	0x6a, 0xb5, 0xa7, 0x00, 0xd3, 0xb9, 0xec, 0x09, 0xad, 0x7a, 0x22, 0xcb, 0x00, 0x05, 0x2a, 0xc8,
	0xac, 0xa3, 0xb4, 0x1a, 0xda, 0xc0, 0x44, 0xcd, 0xa5, 0x2f, 0xf7, 0xe5, 0x12, 0x53, 0x31, 0x1a,
	0x41, 0x9e, 0xa1, 0x83, 0xe2, 0xf7, 0xec, 0xf0, 0xa5, 0xc4, 0x0b, 0xd0, 0xef, 0xf4, 0xb9, 0xb2,
	0x08, 0x66, 0x95, 0xc8, 0x8f, 0xb9, 0xb4, 0xc8, 0xef, 0xb2, 0x0a, 0x98, 0x69, 0xaa, 0xc6, 0x51,
	0xd0, 0x09, 0x4e, 0x4a, 0x49, 0x19, 0xcc, 0xf4, 0xd5, 0x98, 0x1f, 0xb2, 0x0a, 0x4c, 0x26, 0x56,
	0x62, 0x14, 0x12, 0xec, 0x2b, 0x7e, 0xc0, 0xca, 0x73, 0xb5, 0x50, 0x18, 0x95, 0x1c, 0x9b, 0x8a,
	0xf8, 0x7b, 0xc0, 0xea, 0x7f, 0x8a, 0x73, 0xce, 0xf6, 0x26, 0x06, 0x16, 0xa4, 0x59, 0x4d, 0xe8,
	0x37, 0x6f, 0xb2, 0x10, 0x81, 0xe4, 0xaa, 0x49, 0x88, 0xb0, 0x19, 0x21, 0x16, 0x9b, 0x1d, 0x49,
	0x2b, 0x48, 0x7c, 0xc5, 0xef, 0x31, 0x36, 0x32, 0x52, 0xa0, 0x1c, 0xa7, 0x02, 0xa3, 0x3d, 0xe2,
	0x57, 0x3d, 0x32, 0xa0, 0xf6, 0x02, 0x32, 0xb9, 0x4a, 0x71, 0xa5, 0x65, 0x54, 0x76, 0x6d, 0x42,
	0x2e, 0x56, 0x5a, 0xf2, 0x23, 0x76, 0x0b, 0x97, 0xe9, 0x4c, 0xd8, 0x59, 0x54, 0xa1, 0x5e, 0x05,
	0x97, 0xe7, 0xc2, 0xce, 0xe2, 0x6f, 0x01, 0x3b, 0xda, 0xf2, 0xc0, 0x6a, 0xc8, 0xac, 0xdc, 0x5c,
	0x45, 0x6e, 0x15, 0x1e, 0x50, 0xc1, 0x9f, 0xb2, 0x06, 0x05, 0x91, 0xce, 0x1c, 0x3d, 0x0a, 0x3b,
	0xa5, 0x93, 0x5a, 0x7f, 0xbf, 0x2b, 0xb4, 0xea, 0x5e, 0xd3, 0xa9, 0x13, 0xaf, 0x38, 0xfe, 0x19,
	0xab, 0xe7, 0x56, 0x9a, 0xd4, 0xa7, 0x42, 0xe7, 0xd5, 0xfa, 0x07, 0xf4, 0xb7, 0x37, 0x0e, 0x2b,
	0x26, 0x27, 0xb5, 0x0d, 0xd3, 0x83, 0xf1, 0x07, 0xd6, 0x2a, 0x36, 0x7c, 0x21, 0x2d, 0xaa, 0x8c,
	0x52, 0xfe, 0x4f, 0x52, 0x8f, 0x0b, 0x3f, 0xc4, 0x70, 0x68, 0xc8, 0xde, 0x66, 0x9f, 0xd1, 0xac,
	0xd7, 0x1b, 0xd8, 0x7b, 0x33, 0x18, 0x0e, 0x4d, 0xbc, 0x66, 0xc7, 0x3b, 0xf5, 0xbd, 0x0b, 0x8f,
	0x58, 0x53, 0x8c, 0x50, 0x5d, 0xca, 0xe2, 0xf1, 0xf8, 0xf8, 0x1a, 0x0e, 0x1d, 0x38, 0x70, 0xeb,
	0xbc, 0xf0, 0x86, 0xe7, 0xf5, 0xbf, 0x86, 0xfe, 0x95, 0xbc, 0x95, 0xe6, 0x52, 0x8d, 0x24, 0x37,
	0xec, 0xf6, 0x5f, 0x89, 0xf0, 0x63, 0x92, 0xd9, 0xfd, 0x56, 0x5b, 0xed, 0xdd, 0x4d, 0x37, 0x2b,
	0x7e, 0xf8, 0xe9, 0xc7, 0xcf, 0xcf, 0xe1, 0x7d, 0xde, 0xa6, 0xaf, 0x04, 0x41, 0x9f, 0xe6, 0xba,
	0x77, 0xe5, 0x2c, 0x5b, 0xf7, 0x7c, 0x86, 0xfc, 0x4b, 0xc0, 0xee, 0xec, 0x30, 0x81, 0x3f, 0xb8,
	0xa6, 0xbd, 0x6d, 0x7f, 0xab, 0xf3, 0x6f, 0x82, 0x5f, 0xe0, 0x39, 0x2d, 0xf0, 0x84, 0x9f, 0xed,
	0x5c, 0xc0, 0xe6, 0x5a, 0x9a, 0x0c, 0xc6, 0xf2, 0xd4, 0xbb, 0xdb, 0xbb, 0xfa, 0x9d, 0xdb, 0x7a,
	0x58, 0xa1, 0x0f, 0xf5, 0xec, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0xea, 0xb5, 0x3d, 0xfc,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopUpServiceClient is the client API for TopUpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopUpServiceClient interface {
	GetTopUpHistory(ctx context.Context, in *GetTopUpHistoryRequest, opts ...grpc.CallOption) (*GetTopUpHistoryResponse, error)
	GetTopUpDestination(ctx context.Context, in *GetTopUpDestinationRequest, opts ...grpc.CallOption) (*GetTopUpDestinationResponse, error)
}

type topUpServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopUpServiceClient(cc *grpc.ClientConn) TopUpServiceClient {
	return &topUpServiceClient{cc}
}

func (c *topUpServiceClient) GetTopUpHistory(ctx context.Context, in *GetTopUpHistoryRequest, opts ...grpc.CallOption) (*GetTopUpHistoryResponse, error) {
	out := new(GetTopUpHistoryResponse)
	err := c.cc.Invoke(ctx, "/api.TopUpService/GetTopUpHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topUpServiceClient) GetTopUpDestination(ctx context.Context, in *GetTopUpDestinationRequest, opts ...grpc.CallOption) (*GetTopUpDestinationResponse, error) {
	out := new(GetTopUpDestinationResponse)
	err := c.cc.Invoke(ctx, "/api.TopUpService/GetTopUpDestination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopUpServiceServer is the server API for TopUpService service.
type TopUpServiceServer interface {
	GetTopUpHistory(context.Context, *GetTopUpHistoryRequest) (*GetTopUpHistoryResponse, error)
	GetTopUpDestination(context.Context, *GetTopUpDestinationRequest) (*GetTopUpDestinationResponse, error)
}

// UnimplementedTopUpServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopUpServiceServer struct {
}

func (*UnimplementedTopUpServiceServer) GetTopUpHistory(ctx context.Context, req *GetTopUpHistoryRequest) (*GetTopUpHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopUpHistory not implemented")
}
func (*UnimplementedTopUpServiceServer) GetTopUpDestination(ctx context.Context, req *GetTopUpDestinationRequest) (*GetTopUpDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopUpDestination not implemented")
}

func RegisterTopUpServiceServer(s *grpc.Server, srv TopUpServiceServer) {
	s.RegisterService(&_TopUpService_serviceDesc, srv)
}

func _TopUpService_GetTopUpHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopUpHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopUpServiceServer).GetTopUpHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TopUpService/GetTopUpHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopUpServiceServer).GetTopUpHistory(ctx, req.(*GetTopUpHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopUpService_GetTopUpDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopUpDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopUpServiceServer).GetTopUpDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TopUpService/GetTopUpDestination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopUpServiceServer).GetTopUpDestination(ctx, req.(*GetTopUpDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopUpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TopUpService",
	HandlerType: (*TopUpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopUpHistory",
			Handler:    _TopUpService_GetTopUpHistory_Handler,
		},
		{
			MethodName: "GetTopUpDestination",
			Handler:    _TopUpService_GetTopUpDestination_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topup.proto",
}
