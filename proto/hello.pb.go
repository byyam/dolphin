// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GreetingRequest struct {
	Question             string   `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingRequest) Reset()         { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()    {}
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7116fbf19896b02, []int{0}
}

func (m *GreetingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingRequest.Unmarshal(m, b)
}
func (m *GreetingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingRequest.Marshal(b, m, deterministic)
}
func (m *GreetingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingRequest.Merge(m, src)
}
func (m *GreetingRequest) XXX_Size() int {
	return xxx_messageInfo_GreetingRequest.Size(m)
}
func (m *GreetingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingRequest proto.InternalMessageInfo

func (m *GreetingRequest) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

type GreetingResponse struct {
	Answer               string   `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingResponse) Reset()         { *m = GreetingResponse{} }
func (m *GreetingResponse) String() string { return proto.CompactTextString(m) }
func (*GreetingResponse) ProtoMessage()    {}
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7116fbf19896b02, []int{1}
}

func (m *GreetingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingResponse.Unmarshal(m, b)
}
func (m *GreetingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingResponse.Marshal(b, m, deterministic)
}
func (m *GreetingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingResponse.Merge(m, src)
}
func (m *GreetingResponse) XXX_Size() int {
	return xxx_messageInfo_GreetingResponse.Size(m)
}
func (m *GreetingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingResponse proto.InternalMessageInfo

func (m *GreetingResponse) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetingRequest)(nil), "hello.GreetingRequest")
	proto.RegisterType((*GreetingResponse)(nil), "hello.GreetingResponse")
}

func init() { proto.RegisterFile("proto/hello.proto", fileDescriptor_f7116fbf19896b02) }

var fileDescriptor_f7116fbf19896b02 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x1c, 0x25, 0x5d,
	0x2e, 0x7e, 0xf7, 0xa2, 0xd4, 0xd4, 0x92, 0xcc, 0xbc, 0xf4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x29, 0x2e, 0x0e, 0x30, 0x23, 0x33, 0x3f, 0x4f, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x08, 0xce, 0x57, 0xd2, 0xe2, 0x12, 0x40, 0x28, 0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12,
	0xe3, 0x62, 0x4b, 0xcc, 0x2b, 0x2e, 0x4f, 0x2d, 0x82, 0xaa, 0x86, 0xf2, 0x8c, 0xdc, 0xb8, 0x58,
	0x3d, 0x40, 0x76, 0x08, 0xd9, 0x72, 0x71, 0xc0, 0x34, 0x09, 0x89, 0xe9, 0x41, 0x1c, 0x81, 0x66,
	0xa9, 0x94, 0x38, 0x86, 0x38, 0xc4, 0x74, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x83, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x9c, 0x8f, 0xcb, 0xe8, 0xc5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error) {
	out := new(GreetingResponse)
	err := c.cc.Invoke(ctx, "/hello.Hello/Greeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Greeting(context.Context, *GreetingRequest) (*GreetingResponse, error)
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) Greeting(ctx context.Context, req *GreetingRequest) (*GreetingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Greeting(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _Hello_Greeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello.proto",
}
