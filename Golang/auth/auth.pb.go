// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	UserData
	CallbackData
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserData struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
}

func (m *UserData) Reset()                    { *m = UserData{} }
func (m *UserData) String() string            { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()               {}
func (*UserData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserData) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CallbackData struct {
	IdCallback string    `protobuf:"bytes,1,opt,name=id_callback,json=idCallback" json:"id_callback,omitempty"`
	User       *UserData `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *CallbackData) Reset()                    { *m = CallbackData{} }
func (m *CallbackData) String() string            { return proto.CompactTextString(m) }
func (*CallbackData) ProtoMessage()               {}
func (*CallbackData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CallbackData) GetIdCallback() string {
	if m != nil {
		return m.IdCallback
	}
	return ""
}

func (m *CallbackData) GetUser() *UserData {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*UserData)(nil), "auth.userData")
	proto.RegisterType((*CallbackData)(nil), "auth.callbackData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthService service

type AuthServiceClient interface {
	Authlogin(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*UserData, error)
	WaitCallback(ctx context.Context, opts ...grpc.CallOption) (AuthService_WaitCallbackClient, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authlogin(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := grpc.Invoke(ctx, "/auth.authService/authlogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) WaitCallback(ctx context.Context, opts ...grpc.CallOption) (AuthService_WaitCallbackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AuthService_serviceDesc.Streams[0], c.cc, "/auth.authService/waitCallback", opts...)
	if err != nil {
		return nil, err
	}
	x := &authServiceWaitCallbackClient{stream}
	return x, nil
}

type AuthService_WaitCallbackClient interface {
	Send(*CallbackData) error
	Recv() (*CallbackData, error)
	grpc.ClientStream
}

type authServiceWaitCallbackClient struct {
	grpc.ClientStream
}

func (x *authServiceWaitCallbackClient) Send(m *CallbackData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *authServiceWaitCallbackClient) Recv() (*CallbackData, error) {
	m := new(CallbackData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	Authlogin(context.Context, *UserData) (*UserData, error)
	WaitCallback(AuthService_WaitCallbackServer) error
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Authlogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authlogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.authService/Authlogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authlogin(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_WaitCallback_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuthServiceServer).WaitCallback(&authServiceWaitCallbackServer{stream})
}

type AuthService_WaitCallbackServer interface {
	Send(*CallbackData) error
	Recv() (*CallbackData, error)
	grpc.ServerStream
}

type authServiceWaitCallbackServer struct {
	grpc.ServerStream
}

func (x *authServiceWaitCallbackServer) Send(m *CallbackData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *authServiceWaitCallbackServer) Recv() (*CallbackData, error) {
	m := new(CallbackData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.authService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "authlogin",
			Handler:    _AuthService_Authlogin_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "waitCallback",
			Handler:       _AuthService_WaitCallback_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0x86, 0xd9, 0xb5, 0x18, 0x99, 0x25, 0x1c, 0x26, 0x1e, 0x9a, 0xbd, 0x68, 0x7a, 0xe2, 0x22,
	0x31, 0x78, 0xf5, 0xa6, 0x4f, 0x00, 0x0f, 0x60, 0x86, 0x6d, 0xa3, 0x13, 0x0b, 0x25, 0xdd, 0x22,
	0x89, 0xf1, 0xe1, 0x4d, 0xa7, 0x2c, 0x31, 0xc4, 0xdb, 0xfc, 0xdf, 0x3f, 0x9d, 0xfe, 0x33, 0x00,
	0x74, 0x48, 0x1f, 0x8b, 0x7d, 0x0c, 0x29, 0xa0, 0xca, 0xb5, 0xf9, 0x81, 0x9b, 0x43, 0xef, 0xe2,
	0x2b, 0x25, 0xc2, 0x19, 0xd4, 0x6c, 0x75, 0x75, 0x5f, 0xcd, 0x27, 0xab, 0x9a, 0x2d, 0x22, 0xa8,
	0x1d, 0x6d, 0x9d, 0xae, 0x85, 0x48, 0x8d, 0xb7, 0x30, 0x76, 0x5b, 0x62, 0xaf, 0xaf, 0x04, 0x16,
	0x81, 0x6d, 0x99, 0x22, 0xdd, 0x4a, 0x8c, 0xb3, 0xce, 0xde, 0x9e, 0xfa, 0xfe, 0x18, 0xa2, 0xd5,
	0xe3, 0xe2, 0x0d, 0xda, 0xac, 0x61, 0xda, 0x91, 0xf7, 0x1b, 0xea, 0x3e, 0x25, 0xc1, 0x1d, 0x34,
	0x6c, 0xdf, 0x06, 0x74, 0x8a, 0x02, 0x6c, 0x5f, 0x4e, 0x04, 0x0d, 0xa8, 0x3c, 0x58, 0x22, 0x35,
	0xcb, 0xd9, 0x42, 0xf6, 0x19, 0x16, 0x58, 0x89, 0xb7, 0xfc, 0x86, 0x26, 0xe3, 0xb5, 0x8b, 0x5f,
	0xdc, 0x39, 0x7c, 0x80, 0x49, 0x96, 0x3e, 0xbc, 0xf3, 0x0e, 0x2f, 0x5e, 0xb4, 0x17, 0xda, 0x8c,
	0xf0, 0x19, 0xa6, 0x47, 0xe2, 0x74, 0xfe, 0x11, 0x4b, 0xc7, 0xdf, 0x98, 0xed, 0x3f, 0xcc, 0x8c,
	0xe6, 0xd5, 0x63, 0xb5, 0xb9, 0x96, 0xdb, 0x3e, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xa9,
	0x09, 0x78, 0x69, 0x01, 0x00, 0x00,
}
