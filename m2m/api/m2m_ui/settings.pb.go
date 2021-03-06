// Code generated by protoc-gen-go. DO NOT EDIT.
// source: settings.proto

package m2m_ui

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type GetSettingsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSettingsRequest) Reset()         { *m = GetSettingsRequest{} }
func (m *GetSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*GetSettingsRequest) ProtoMessage()    {}
func (*GetSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7cab62fa432213, []int{0}
}

func (m *GetSettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSettingsRequest.Unmarshal(m, b)
}
func (m *GetSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSettingsRequest.Marshal(b, m, deterministic)
}
func (m *GetSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSettingsRequest.Merge(m, src)
}
func (m *GetSettingsRequest) XXX_Size() int {
	return xxx_messageInfo_GetSettingsRequest.Size(m)
}
func (m *GetSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSettingsRequest proto.InternalMessageInfo

type GetSettingsResponse struct {
	LowBalanceWarning          int64    `protobuf:"varint,1,opt,name=lowBalanceWarning,proto3" json:"lowBalanceWarning,omitempty"`
	DownlinkFee                int64    `protobuf:"varint,2,opt,name=downlinkFee,proto3" json:"downlinkFee,omitempty"`
	TransactionPercentageShare int64    `protobuf:"varint,3,opt,name=transactionPercentageShare,proto3" json:"transactionPercentageShare,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *GetSettingsResponse) Reset()         { *m = GetSettingsResponse{} }
func (m *GetSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*GetSettingsResponse) ProtoMessage()    {}
func (*GetSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7cab62fa432213, []int{1}
}

func (m *GetSettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSettingsResponse.Unmarshal(m, b)
}
func (m *GetSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSettingsResponse.Marshal(b, m, deterministic)
}
func (m *GetSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSettingsResponse.Merge(m, src)
}
func (m *GetSettingsResponse) XXX_Size() int {
	return xxx_messageInfo_GetSettingsResponse.Size(m)
}
func (m *GetSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSettingsResponse proto.InternalMessageInfo

func (m *GetSettingsResponse) GetLowBalanceWarning() int64 {
	if m != nil {
		return m.LowBalanceWarning
	}
	return 0
}

func (m *GetSettingsResponse) GetDownlinkFee() int64 {
	if m != nil {
		return m.DownlinkFee
	}
	return 0
}

func (m *GetSettingsResponse) GetTransactionPercentageShare() int64 {
	if m != nil {
		return m.TransactionPercentageShare
	}
	return 0
}

type ModifySettingsRequest struct {
	LowBalanceWarning          *wrappers.Int64Value `protobuf:"bytes,1,opt,name=lowBalanceWarning,proto3" json:"lowBalanceWarning,omitempty"`
	DownlinkFee                *wrappers.Int64Value `protobuf:"bytes,2,opt,name=downlinkFee,proto3" json:"downlinkFee,omitempty"`
	TransactionPercentageShare *wrappers.Int64Value `protobuf:"bytes,3,opt,name=transactionPercentageShare,proto3" json:"transactionPercentageShare,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}             `json:"-"`
	XXX_unrecognized           []byte               `json:"-"`
	XXX_sizecache              int32                `json:"-"`
}

func (m *ModifySettingsRequest) Reset()         { *m = ModifySettingsRequest{} }
func (m *ModifySettingsRequest) String() string { return proto.CompactTextString(m) }
func (*ModifySettingsRequest) ProtoMessage()    {}
func (*ModifySettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7cab62fa432213, []int{2}
}

func (m *ModifySettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifySettingsRequest.Unmarshal(m, b)
}
func (m *ModifySettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifySettingsRequest.Marshal(b, m, deterministic)
}
func (m *ModifySettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifySettingsRequest.Merge(m, src)
}
func (m *ModifySettingsRequest) XXX_Size() int {
	return xxx_messageInfo_ModifySettingsRequest.Size(m)
}
func (m *ModifySettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifySettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifySettingsRequest proto.InternalMessageInfo

func (m *ModifySettingsRequest) GetLowBalanceWarning() *wrappers.Int64Value {
	if m != nil {
		return m.LowBalanceWarning
	}
	return nil
}

func (m *ModifySettingsRequest) GetDownlinkFee() *wrappers.Int64Value {
	if m != nil {
		return m.DownlinkFee
	}
	return nil
}

func (m *ModifySettingsRequest) GetTransactionPercentageShare() *wrappers.Int64Value {
	if m != nil {
		return m.TransactionPercentageShare
	}
	return nil
}

type ModifySettingsResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifySettingsResponse) Reset()         { *m = ModifySettingsResponse{} }
func (m *ModifySettingsResponse) String() string { return proto.CompactTextString(m) }
func (*ModifySettingsResponse) ProtoMessage()    {}
func (*ModifySettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7cab62fa432213, []int{3}
}

func (m *ModifySettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifySettingsResponse.Unmarshal(m, b)
}
func (m *ModifySettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifySettingsResponse.Marshal(b, m, deterministic)
}
func (m *ModifySettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifySettingsResponse.Merge(m, src)
}
func (m *ModifySettingsResponse) XXX_Size() int {
	return xxx_messageInfo_ModifySettingsResponse.Size(m)
}
func (m *ModifySettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifySettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifySettingsResponse proto.InternalMessageInfo

func (m *ModifySettingsResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*GetSettingsRequest)(nil), "m2m_ui.GetSettingsRequest")
	proto.RegisterType((*GetSettingsResponse)(nil), "m2m_ui.GetSettingsResponse")
	proto.RegisterType((*ModifySettingsRequest)(nil), "m2m_ui.ModifySettingsRequest")
	proto.RegisterType((*ModifySettingsResponse)(nil), "m2m_ui.ModifySettingsResponse")
}

func init() { proto.RegisterFile("settings.proto", fileDescriptor_6c7cab62fa432213) }

var fileDescriptor_6c7cab62fa432213 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x65, 0x5a, 0x28, 0x1f, 0x29, 0x5f, 0x8b, 0xd1, 0x96, 0x21, 0xd5, 0x52, 0x66, 0x25, 0x22,
	0x53, 0xa9, 0xe2, 0x42, 0xd0, 0x85, 0x0b, 0xa5, 0x0b, 0x41, 0x5a, 0xd0, 0x85, 0x88, 0xa4, 0xd3,
	0xdb, 0x31, 0x38, 0x4d, 0xc6, 0x24, 0x63, 0x71, 0xeb, 0x2b, 0xf8, 0x0c, 0x3e, 0x91, 0x4b, 0xb7,
	0x3e, 0x87, 0x88, 0x69, 0x06, 0x3a, 0xfd, 0x9b, 0x65, 0xee, 0xb9, 0xb9, 0xe7, 0x9e, 0x73, 0x0f,
	0xaa, 0x28, 0xd0, 0x9a, 0xf1, 0x50, 0xf9, 0xb1, 0x14, 0x5a, 0xe0, 0xd2, 0xb8, 0x33, 0x7e, 0x48,
	0x18, 0xd9, 0x0e, 0x85, 0x08, 0x23, 0x68, 0xd3, 0x98, 0xb5, 0x29, 0xe7, 0x42, 0x53, 0xcd, 0x04,
	0xb7, 0x5d, 0xa4, 0x69, 0x51, 0xf3, 0x1a, 0x24, 0xa3, 0xf6, 0x44, 0xd2, 0x38, 0x06, 0x69, 0x71,
	0x6f, 0x0b, 0xe1, 0x4b, 0xd0, 0x7d, 0x3b, 0xba, 0x07, 0xcf, 0x09, 0x28, 0xed, 0x7d, 0x38, 0x68,
	0x33, 0x53, 0x56, 0xb1, 0xe0, 0x0a, 0xf0, 0x3e, 0xda, 0x88, 0xc4, 0xe4, 0x9c, 0x46, 0x94, 0x07,
	0x70, 0x4b, 0x25, 0x67, 0x3c, 0x74, 0x9d, 0x96, 0xb3, 0x5b, 0xec, 0x2d, 0x02, 0xb8, 0x85, 0xca,
	0x43, 0x31, 0xe1, 0x11, 0xe3, 0x4f, 0x17, 0x00, 0x6e, 0xc1, 0xf4, 0xcd, 0x96, 0xf0, 0x19, 0x22,
	0x5a, 0x52, 0xae, 0x68, 0xf0, 0xb7, 0xf3, 0x35, 0xc8, 0x00, 0xb8, 0xa6, 0x21, 0xf4, 0x1f, 0xa9,
	0x04, 0xb7, 0x68, 0x3e, 0xac, 0xe9, 0xf0, 0x7e, 0x1c, 0x54, 0xbb, 0x12, 0x43, 0x36, 0x7a, 0x9d,
	0x53, 0x80, 0xbb, 0xab, 0x36, 0x2d, 0x77, 0x1a, 0xfe, 0xd4, 0x13, 0x3f, 0xf5, 0xc4, 0xef, 0x72,
	0x7d, 0x7c, 0x74, 0x43, 0xa3, 0x04, 0x96, 0xc9, 0x38, 0x5d, 0x94, 0x91, 0x33, 0x24, 0xa3, 0xf1,
	0x2e, 0x57, 0x63, 0xce, 0xb4, 0x75, 0x06, 0x1c, 0xa0, 0xfa, 0xbc, 0x7e, 0x7b, 0xaa, 0x3a, 0x2a,
	0x29, 0x4d, 0x75, 0xa2, 0x8c, 0xea, 0x7f, 0x3d, 0xfb, 0xea, 0x7c, 0x39, 0xa8, 0x9a, 0x36, 0xf7,
	0x41, 0xbe, 0xb0, 0x00, 0xf0, 0x3d, 0x2a, 0xcf, 0x5c, 0x1b, 0x13, 0x7f, 0x1a, 0x2d, 0x7f, 0x31,
	0x19, 0xa4, 0xb1, 0x14, 0x9b, 0x72, 0x7a, 0xb5, 0xb7, 0xcf, 0xef, 0xf7, 0x42, 0x15, 0xff, 0x37,
	0x61, 0x4c, 0xf3, 0x8a, 0x19, 0xaa, 0x64, 0x97, 0xc4, 0x3b, 0xe9, 0x94, 0xa5, 0xc7, 0x23, 0xcd,
	0x55, 0xb0, 0xe5, 0x71, 0x0d, 0x0f, 0x26, 0x59, 0x9e, 0x13, 0x67, 0x6f, 0x50, 0x32, 0x06, 0x1e,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x67, 0xd0, 0x5f, 0x2d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SettingsServiceClient is the client API for SettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SettingsServiceClient interface {
	GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error)
	ModifySettings(ctx context.Context, in *ModifySettingsRequest, opts ...grpc.CallOption) (*ModifySettingsResponse, error)
}

type settingsServiceClient struct {
	cc *grpc.ClientConn
}

func NewSettingsServiceClient(cc *grpc.ClientConn) SettingsServiceClient {
	return &settingsServiceClient{cc}
}

func (c *settingsServiceClient) GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error) {
	out := new(GetSettingsResponse)
	err := c.cc.Invoke(ctx, "/m2m_ui.SettingsService/GetSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) ModifySettings(ctx context.Context, in *ModifySettingsRequest, opts ...grpc.CallOption) (*ModifySettingsResponse, error) {
	out := new(ModifySettingsResponse)
	err := c.cc.Invoke(ctx, "/m2m_ui.SettingsService/ModifySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServiceServer is the server API for SettingsService service.
type SettingsServiceServer interface {
	GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error)
	ModifySettings(context.Context, *ModifySettingsRequest) (*ModifySettingsResponse, error)
}

// UnimplementedSettingsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSettingsServiceServer struct {
}

func (*UnimplementedSettingsServiceServer) GetSettings(ctx context.Context, req *GetSettingsRequest) (*GetSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
}
func (*UnimplementedSettingsServiceServer) ModifySettings(ctx context.Context, req *ModifySettingsRequest) (*ModifySettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySettings not implemented")
}

func RegisterSettingsServiceServer(s *grpc.Server, srv SettingsServiceServer) {
	s.RegisterService(&_SettingsService_serviceDesc, srv)
}

func _SettingsService_GetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m2m_ui.SettingsService/GetSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetSettings(ctx, req.(*GetSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_ModifySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).ModifySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m2m_ui.SettingsService/ModifySettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).ModifySettings(ctx, req.(*ModifySettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SettingsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "m2m_ui.SettingsService",
	HandlerType: (*SettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSettings",
			Handler:    _SettingsService_GetSettings_Handler,
		},
		{
			MethodName: "ModifySettings",
			Handler:    _SettingsService_ModifySettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings.proto",
}
