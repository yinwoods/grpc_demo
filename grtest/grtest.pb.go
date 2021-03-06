// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grtest.proto

package grtest

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Ping struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	Aux                  string   `protobuf:"bytes,2,opt,name=aux,proto3" json:"aux,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_6951cfd7d0369690, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

func (m *Ping) GetAux() string {
	if m != nil {
		return m.Aux
	}
	return ""
}

type Pong struct {
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	Aux                  string   `protobuf:"bytes,2,opt,name=aux,proto3" json:"aux,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_6951cfd7d0369690, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func (m *Pong) GetAux() string {
	if m != nil {
		return m.Aux
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "Ping")
	proto.RegisterType((*Pong)(nil), "Pong")
}

func init() { proto.RegisterFile("grtest.proto", fileDescriptor_6951cfd7d0369690) }

var fileDescriptor_6951cfd7d0369690 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2f, 0x2a, 0x49,
	0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe1, 0x62, 0x09, 0xc8, 0xcc, 0x4b,
	0x17, 0x12, 0xe2, 0x62, 0x29, 0xc8, 0xcc, 0x4b, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0xb3, 0x85, 0x04, 0xb8, 0x98, 0x13, 0x4b, 0x2b, 0x24, 0x98, 0xc0, 0x42, 0x20, 0x26, 0x58, 0x75,
	0x3e, 0x54, 0x75, 0x3e, 0x92, 0xea, 0x7c, 0x6c, 0xaa, 0x8d, 0x5c, 0xb8, 0xb8, 0x43, 0x52, 0x8b,
	0x4b, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xa4, 0xb9, 0xb8, 0x41, 0x56, 0x05, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xb1, 0xea, 0x81, 0x78, 0x52, 0xac, 0x7a, 0x60, 0x13, 0x25,
	0xb8, 0xd8, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0x51, 0xc5, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0x0e,
	0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x63, 0x37, 0x7f, 0x23, 0xb8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	PingRequest(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
	Stream(ctx context.Context, in *Ping, opts ...grpc.CallOption) (TestService_StreamClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) PingRequest(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/TestService/PingRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Stream(ctx context.Context, in *Ping, opts ...grpc.CallOption) (TestService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/TestService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_StreamClient interface {
	Recv() (*Pong, error)
	grpc.ClientStream
}

type testServiceStreamClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	PingRequest(context.Context, *Ping) (*Pong, error)
	Stream(*Ping, TestService_StreamServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_PingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestService/PingRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingRequest(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Ping)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).Stream(m, &testServiceStreamServer{stream})
}

type TestService_StreamServer interface {
	Send(*Pong) error
	grpc.ServerStream
}

type testServiceStreamServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingRequest",
			Handler:    _TestService_PingRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _TestService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grtest.proto",
}
