// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type GetUserRequest struct {
	// Username.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Plain text password.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// The client connection's id.
	Clientid             string   `protobuf:"bytes,3,opt,name=clientid,proto3" json:"clientid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *GetUserRequest) GetClientid() string {
	if m != nil {
		return m.Clientid
	}
	return ""
}

type GetSuperuserRequest struct {
	// Username.
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSuperuserRequest) Reset()         { *m = GetSuperuserRequest{} }
func (m *GetSuperuserRequest) String() string { return proto.CompactTextString(m) }
func (*GetSuperuserRequest) ProtoMessage()    {}
func (*GetSuperuserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *GetSuperuserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSuperuserRequest.Unmarshal(m, b)
}
func (m *GetSuperuserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSuperuserRequest.Marshal(b, m, deterministic)
}
func (m *GetSuperuserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSuperuserRequest.Merge(m, src)
}
func (m *GetSuperuserRequest) XXX_Size() int {
	return xxx_messageInfo_GetSuperuserRequest.Size(m)
}
func (m *GetSuperuserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSuperuserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSuperuserRequest proto.InternalMessageInfo

func (m *GetSuperuserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type CheckAclRequest struct {
	// Username.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Topic to be checked for.
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// The client connection's id.
	Clientid string `protobuf:"bytes,3,opt,name=clientid,proto3" json:"clientid,omitempty"`
	// Topic access.
	Acc                  int32    `protobuf:"varint,4,opt,name=acc,proto3" json:"acc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAclRequest) Reset()         { *m = CheckAclRequest{} }
func (m *CheckAclRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAclRequest) ProtoMessage()    {}
func (*CheckAclRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *CheckAclRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAclRequest.Unmarshal(m, b)
}
func (m *CheckAclRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAclRequest.Marshal(b, m, deterministic)
}
func (m *CheckAclRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAclRequest.Merge(m, src)
}
func (m *CheckAclRequest) XXX_Size() int {
	return xxx_messageInfo_CheckAclRequest.Size(m)
}
func (m *CheckAclRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAclRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAclRequest proto.InternalMessageInfo

func (m *CheckAclRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CheckAclRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *CheckAclRequest) GetClientid() string {
	if m != nil {
		return m.Clientid
	}
	return ""
}

func (m *CheckAclRequest) GetAcc() int32 {
	if m != nil {
		return m.Acc
	}
	return 0
}

type AuthResponse struct {
	// If the user is authorized/authenticated.
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type NameResponse struct {
	// The name of the gRPC backend.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameResponse) Reset()         { *m = NameResponse{} }
func (m *NameResponse) String() string { return proto.CompactTextString(m) }
func (*NameResponse) ProtoMessage()    {}
func (*NameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *NameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameResponse.Unmarshal(m, b)
}
func (m *NameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameResponse.Marshal(b, m, deterministic)
}
func (m *NameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameResponse.Merge(m, src)
}
func (m *NameResponse) XXX_Size() int {
	return xxx_messageInfo_NameResponse.Size(m)
}
func (m *NameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NameResponse proto.InternalMessageInfo

func (m *NameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "grpc.GetUserRequest")
	proto.RegisterType((*GetSuperuserRequest)(nil), "grpc.GetSuperuserRequest")
	proto.RegisterType((*CheckAclRequest)(nil), "grpc.CheckAclRequest")
	proto.RegisterType((*AuthResponse)(nil), "grpc.AuthResponse")
	proto.RegisterType((*NameResponse)(nil), "grpc.NameResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x14, 0x84, 0x52, 0x14, 0x9f, 0x04, 0xcd, 0x8a, 0xa6, 0x62, 0x62, 0xc8, 0x9e, 0x38, 0x95, 0xa8,
	0x31, 0x7a, 0x33, 0xc4, 0x18, 0x3c, 0x79, 0x28, 0xf1, 0x07, 0x94, 0xe5, 0x09, 0x0d, 0x85, 0x5d,
	0xf6, 0x43, 0xe3, 0xcf, 0xf2, 0x1f, 0x9a, 0xed, 0xd2, 0xa6, 0x92, 0xd4, 0x70, 0xdb, 0x99, 0xd9,
	0xd9, 0x37, 0x9d, 0x3e, 0x80, 0xd8, 0xe8, 0x45, 0x28, 0x24, 0xd7, 0x9c, 0xf8, 0x73, 0x29, 0x58,
	0xef, 0x6a, 0xce, 0xf9, 0x3c, 0xc5, 0x61, 0xc6, 0x4d, 0xcd, 0xc7, 0x10, 0x57, 0x42, 0x7f, 0xbb,
	0x2b, 0x74, 0x06, 0x9d, 0x31, 0xea, 0x77, 0x85, 0x32, 0xc2, 0x8d, 0x41, 0xa5, 0x49, 0x0f, 0x5a,
	0x46, 0xa1, 0x5c, 0xc7, 0x2b, 0x0c, 0xea, 0xfd, 0xfa, 0xe0, 0x28, 0x2a, 0xb0, 0xd5, 0x44, 0xac,
	0xd4, 0x17, 0x97, 0xb3, 0xc0, 0x73, 0x5a, 0x8e, 0xad, 0xc6, 0xd2, 0x04, 0xd7, 0x3a, 0x99, 0x05,
	0x0d, 0xa7, 0xe5, 0x98, 0xde, 0xc0, 0xd9, 0x18, 0xf5, 0xc4, 0x08, 0x94, 0x66, 0xbf, 0x51, 0x74,
	0x03, 0x27, 0xcf, 0x0b, 0x64, 0xcb, 0x11, 0x4b, 0xf7, 0x49, 0xd6, 0x85, 0xa6, 0xe6, 0x22, 0x61,
	0xdb, 0x58, 0x0e, 0xfc, 0x97, 0x89, 0x9c, 0x42, 0x23, 0x66, 0x2c, 0xf0, 0xfb, 0xf5, 0x41, 0x33,
	0xb2, 0x47, 0x7a, 0x0d, 0xed, 0x91, 0xd1, 0x8b, 0x08, 0x95, 0xe0, 0x6b, 0x85, 0xa4, 0x03, 0x1e,
	0x5f, 0x66, 0x93, 0x5a, 0x91, 0xc7, 0x97, 0x94, 0x42, 0xfb, 0x2d, 0x5e, 0x61, 0xa1, 0x13, 0xf0,
	0x4b, 0x59, 0xb2, 0xf3, 0xed, 0x8f, 0x07, 0xc7, 0xf6, 0x91, 0x09, 0xca, 0xcf, 0x84, 0x21, 0xb9,
	0x87, 0xc3, 0x6d, 0xbf, 0xa4, 0x1b, 0xda, 0xdf, 0x11, 0xfe, 0xad, 0xbb, 0x47, 0x1c, 0x5b, 0x1e,
	0x4c, 0x6b, 0xe4, 0x09, 0xda, 0xe5, 0xc2, 0xc8, 0x65, 0xe1, 0xdd, 0x2d, 0xb1, 0xe2, 0x81, 0x07,
	0x68, 0xe5, 0xf5, 0x91, 0x73, 0x77, 0x63, 0xa7, 0xce, 0x4a, 0xa3, 0x0d, 0x6c, 0xbf, 0x93, 0x5c,
	0x84, 0x6e, 0x73, 0xc2, 0x7c, 0x73, 0xc2, 0x17, 0xbb, 0x39, 0xb9, 0xb1, 0xdc, 0x05, 0xad, 0x91,
	0x47, 0xf0, 0x5f, 0xe3, 0x54, 0x57, 0xba, 0x2a, 0x78, 0x5a, 0x9b, 0x1e, 0x64, 0xcc, 0xdd, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0x08, 0xf7, 0xf9, 0xbb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	// GetUser tries to authenticate a user.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// GetSuperuser checks if a user is a superuser.
	GetSuperuser(ctx context.Context, in *GetSuperuserRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// CheckAcl checks user's authorization for the given topic.
	CheckAcl(ctx context.Context, in *CheckAclRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// GetName retrieves the name of the backend.
	GetName(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NameResponse, error)
	// Halt signals the backend to halt.
	Halt(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetSuperuser(ctx context.Context, in *GetSuperuserRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/GetSuperuser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckAcl(ctx context.Context, in *CheckAclRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/CheckAcl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetName(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Halt(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/Halt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	// GetUser tries to authenticate a user.
	GetUser(context.Context, *GetUserRequest) (*AuthResponse, error)
	// GetSuperuser checks if a user is a superuser.
	GetSuperuser(context.Context, *GetSuperuserRequest) (*AuthResponse, error)
	// CheckAcl checks user's authorization for the given topic.
	CheckAcl(context.Context, *CheckAclRequest) (*AuthResponse, error)
	// GetName retrieves the name of the backend.
	GetName(context.Context, *empty.Empty) (*NameResponse, error)
	// Halt signals the backend to halt.
	Halt(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedAuthServiceServer) GetSuperuser(ctx context.Context, req *GetSuperuserRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuperuser not implemented")
}
func (*UnimplementedAuthServiceServer) CheckAcl(ctx context.Context, req *CheckAclRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAcl not implemented")
}
func (*UnimplementedAuthServiceServer) GetName(ctx context.Context, req *empty.Empty) (*NameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}
func (*UnimplementedAuthServiceServer) Halt(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Halt not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetSuperuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSuperuserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetSuperuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/GetSuperuser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetSuperuser(ctx, req.(*GetSuperuserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/CheckAcl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckAcl(ctx, req.(*CheckAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetName(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Halt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Halt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/Halt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Halt(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _AuthService_GetUser_Handler,
		},
		{
			MethodName: "GetSuperuser",
			Handler:    _AuthService_GetSuperuser_Handler,
		},
		{
			MethodName: "CheckAcl",
			Handler:    _AuthService_CheckAcl_Handler,
		},
		{
			MethodName: "GetName",
			Handler:    _AuthService_GetName_Handler,
		},
		{
			MethodName: "Halt",
			Handler:    _AuthService_Halt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
