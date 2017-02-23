// Code generated by protoc-gen-go.
// source: serverpb.proto
// DO NOT EDIT!

/*
Package serverpb is a generated protocol buffer package.

It is generated from these files:
	serverpb.proto

It has these top-level messages:
	Req
	Msg
	Resp
	StatusCode
*/
package serverpb

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

type Req struct {
	Offset string `protobuf:"bytes,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Req) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

type Msg struct {
	Offset string `protobuf:"bytes,1,opt,name=offset" json:"offset,omitempty"`
	Body   string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Msg) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *Msg) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Resp struct {
	Offset string `protobuf:"bytes,1,opt,name=offset" json:"offset,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Resp) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *Resp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StatusCode struct {
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
}

func (m *StatusCode) Reset()                    { *m = StatusCode{} }
func (m *StatusCode) String() string            { return proto.CompactTextString(m) }
func (*StatusCode) ProtoMessage()               {}
func (*StatusCode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatusCode) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*Req)(nil), "serverpb.Req")
	proto.RegisterType((*Msg)(nil), "serverpb.Msg")
	proto.RegisterType((*Resp)(nil), "serverpb.Resp")
	proto.RegisterType((*StatusCode)(nil), "serverpb.StatusCode")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Omq service

type OmqClient interface {
	PutMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*StatusCode, error)
	Poll(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type omqClient struct {
	cc *grpc.ClientConn
}

func NewOmqClient(cc *grpc.ClientConn) OmqClient {
	return &omqClient{cc}
}

func (c *omqClient) PutMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := grpc.Invoke(ctx, "/serverpb.Omq/PutMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omqClient) Poll(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/serverpb.Omq/poll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Omq service

type OmqServer interface {
	PutMsg(context.Context, *Msg) (*StatusCode, error)
	Poll(context.Context, *Req) (*Resp, error)
}

func RegisterOmqServer(s *grpc.Server, srv OmqServer) {
	s.RegisterService(&_Omq_serviceDesc, srv)
}

func _Omq_PutMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmqServer).PutMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.Omq/PutMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmqServer).PutMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omq_Poll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmqServer).Poll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.Omq/Poll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmqServer).Poll(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Omq_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.Omq",
	HandlerType: (*OmqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutMsg",
			Handler:    _Omq_PutMsg_Handler,
		},
		{
			MethodName: "poll",
			Handler:    _Omq_Poll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverpb.proto",
}

func init() { proto.RegisterFile("serverpb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0x2a, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x64,
	0xb9, 0x98, 0x83, 0x52, 0x0b, 0x85, 0xc4, 0xb8, 0xd8, 0xf2, 0xd3, 0xd2, 0x8a, 0x53, 0x4b, 0x24,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x25, 0x43, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0x24,
	0x69, 0x46, 0x64, 0x69, 0x21, 0x21, 0x2e, 0x96, 0xa4, 0xfc, 0x94, 0x4a, 0xa8, 0x26, 0x30, 0x5b,
	0xc9, 0x80, 0x8b, 0x25, 0x28, 0xb5, 0xb8, 0x00, 0xa7, 0x1e, 0x01, 0x2e, 0xe6, 0xdc, 0xe2, 0x74,
	0xa8, 0x16, 0x10, 0x53, 0x49, 0x81, 0x8b, 0x2b, 0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0xd8, 0x39, 0x3f,
	0x25, 0x15, 0x64, 0x66, 0x72, 0x7e, 0x4a, 0x2a, 0x58, 0x17, 0x6b, 0x10, 0x98, 0x6d, 0x14, 0xcf,
	0xc5, 0xec, 0x9f, 0x5b, 0x28, 0xa4, 0xcf, 0xc5, 0x16, 0x50, 0x5a, 0x02, 0x72, 0x10, 0xaf, 0x1e,
	0xdc, 0x47, 0xbe, 0xc5, 0xe9, 0x52, 0x22, 0x08, 0x2e, 0xc2, 0x24, 0x25, 0x06, 0x21, 0x75, 0x2e,
	0x96, 0x82, 0xfc, 0x9c, 0x1c, 0x64, 0xe5, 0x41, 0xa9, 0x85, 0x52, 0x7c, 0xc8, 0xdc, 0xe2, 0x02,
	0x25, 0x86, 0x24, 0x36, 0x70, 0xb8, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x85, 0x32,
	0x19, 0x29, 0x01, 0x00, 0x00,
}
