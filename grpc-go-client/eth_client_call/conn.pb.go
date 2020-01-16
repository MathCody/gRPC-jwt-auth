// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conn.proto

package eth_client_call

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

type ClientCallInterface struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientCallInterface) Reset()         { *m = ClientCallInterface{} }
func (m *ClientCallInterface) String() string { return proto.CompactTextString(m) }
func (*ClientCallInterface) ProtoMessage()    {}
func (*ClientCallInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{0}
}

func (m *ClientCallInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCallInterface.Unmarshal(m, b)
}
func (m *ClientCallInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCallInterface.Marshal(b, m, deterministic)
}
func (m *ClientCallInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCallInterface.Merge(m, src)
}
func (m *ClientCallInterface) XXX_Size() int {
	return xxx_messageInfo_ClientCallInterface.Size(m)
}
func (m *ClientCallInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCallInterface.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCallInterface proto.InternalMessageInfo

func (m *ClientCallInterface) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *ClientCallInterface) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type ClientCallRequest struct {
	CallInterface        *ClientCallInterface `protobuf:"bytes,1,opt,name=callInterface,proto3" json:"callInterface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClientCallRequest) Reset()         { *m = ClientCallRequest{} }
func (m *ClientCallRequest) String() string { return proto.CompactTextString(m) }
func (*ClientCallRequest) ProtoMessage()    {}
func (*ClientCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{1}
}

func (m *ClientCallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCallRequest.Unmarshal(m, b)
}
func (m *ClientCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCallRequest.Marshal(b, m, deterministic)
}
func (m *ClientCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCallRequest.Merge(m, src)
}
func (m *ClientCallRequest) XXX_Size() int {
	return xxx_messageInfo_ClientCallRequest.Size(m)
}
func (m *ClientCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCallRequest proto.InternalMessageInfo

func (m *ClientCallRequest) GetCallInterface() *ClientCallInterface {
	if m != nil {
		return m.CallInterface
	}
	return nil
}

type ClientCallResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientCallResponse) Reset()         { *m = ClientCallResponse{} }
func (m *ClientCallResponse) String() string { return proto.CompactTextString(m) }
func (*ClientCallResponse) ProtoMessage()    {}
func (*ClientCallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{2}
}

func (m *ClientCallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCallResponse.Unmarshal(m, b)
}
func (m *ClientCallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCallResponse.Marshal(b, m, deterministic)
}
func (m *ClientCallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCallResponse.Merge(m, src)
}
func (m *ClientCallResponse) XXX_Size() int {
	return xxx_messageInfo_ClientCallResponse.Size(m)
}
func (m *ClientCallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCallResponse proto.InternalMessageInfo

func (m *ClientCallResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientCallInterface)(nil), "eth_client_call.ClientCallInterface")
	proto.RegisterType((*ClientCallRequest)(nil), "eth_client_call.ClientCallRequest")
	proto.RegisterType((*ClientCallResponse)(nil), "eth_client_call.ClientCallResponse")
}

func init() { proto.RegisterFile("conn.proto", fileDescriptor_f401a58c1fc7ceef) }

var fileDescriptor_f401a58c1fc7ceef = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4a, 0xc6, 0x30,
	0x14, 0x85, 0xad, 0xc3, 0x2f, 0xbd, 0x22, 0xd2, 0x08, 0x52, 0x9c, 0x24, 0x3a, 0x38, 0x48, 0x91,
	0xfa, 0x06, 0xd6, 0xa5, 0x8e, 0x71, 0x11, 0x97, 0x12, 0xd3, 0x2b, 0x16, 0x6e, 0x93, 0x98, 0xa4,
	0x42, 0xdf, 0x5e, 0x6c, 0x23, 0xd6, 0xf2, 0xd3, 0xf1, 0xe3, 0x24, 0xe7, 0x7e, 0x1c, 0x00, 0x65,
	0xb4, 0x2e, 0xac, 0x33, 0xc1, 0xb0, 0x53, 0x0c, 0x1f, 0x8d, 0xa2, 0x0e, 0x75, 0x68, 0x94, 0x24,
	0xe2, 0x35, 0x9c, 0x55, 0x13, 0x56, 0x92, 0xa8, 0xd6, 0x01, 0xdd, 0xbb, 0x54, 0xc8, 0x72, 0x38,
	0x52, 0xa6, 0xef, 0xa5, 0x6e, 0xf3, 0xe4, 0x32, 0xb9, 0x49, 0xc5, 0x2f, 0xfe, 0x24, 0x56, 0x8e,
	0x64, 0x64, 0x9b, 0x1f, 0xce, 0x49, 0x44, 0xde, 0x40, 0xf6, 0x57, 0x25, 0xf0, 0x73, 0x40, 0x1f,
	0xd8, 0x13, 0x9c, 0xa8, 0x65, 0xf3, 0x54, 0x77, 0x5c, 0x5e, 0x17, 0x2b, 0x91, 0x62, 0x8f, 0x85,
	0xf8, 0xff, 0x95, 0xdf, 0x02, 0x5b, 0x1e, 0xf0, 0xd6, 0x68, 0x8f, 0xec, 0x1c, 0x76, 0x0e, 0xfd,
	0x40, 0x21, 0x9a, 0x46, 0x2a, 0xfb, 0xa5, 0xce, 0x33, 0xba, 0xaf, 0x4e, 0x21, 0x7b, 0x81, 0x54,
	0x0c, 0xfa, 0x11, 0x2d, 0x99, 0x91, 0xf1, 0x0d, 0x89, 0xe8, 0x7f, 0x71, 0xb5, 0xf9, 0x66, 0x56,
	0xe0, 0x07, 0x77, 0xc9, 0x43, 0xf6, 0xba, 0xde, 0xf6, 0x6d, 0x37, 0x6d, 0x7e, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xa0, 0x34, 0xb6, 0x32, 0x81, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientCallServiceClient is the client API for ClientCallService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientCallServiceClient interface {
	RunDeploy(ctx context.Context, in *ClientCallRequest, opts ...grpc.CallOption) (ClientCallService_RunDeployClient, error)
}

type clientCallServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientCallServiceClient(cc *grpc.ClientConn) ClientCallServiceClient {
	return &clientCallServiceClient{cc}
}

func (c *clientCallServiceClient) RunDeploy(ctx context.Context, in *ClientCallRequest, opts ...grpc.CallOption) (ClientCallService_RunDeployClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientCallService_serviceDesc.Streams[0], "/eth_client_call.ClientCallService/RunDeploy", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientCallServiceRunDeployClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientCallService_RunDeployClient interface {
	Recv() (*ClientCallResponse, error)
	grpc.ClientStream
}

type clientCallServiceRunDeployClient struct {
	grpc.ClientStream
}

func (x *clientCallServiceRunDeployClient) Recv() (*ClientCallResponse, error) {
	m := new(ClientCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientCallServiceServer is the server API for ClientCallService service.
type ClientCallServiceServer interface {
	RunDeploy(*ClientCallRequest, ClientCallService_RunDeployServer) error
}

// UnimplementedClientCallServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientCallServiceServer struct {
}

func (*UnimplementedClientCallServiceServer) RunDeploy(req *ClientCallRequest, srv ClientCallService_RunDeployServer) error {
	return status.Errorf(codes.Unimplemented, "method RunDeploy not implemented")
}

func RegisterClientCallServiceServer(s *grpc.Server, srv ClientCallServiceServer) {
	s.RegisterService(&_ClientCallService_serviceDesc, srv)
}

func _ClientCallService_RunDeploy_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientCallRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientCallServiceServer).RunDeploy(m, &clientCallServiceRunDeployServer{stream})
}

type ClientCallService_RunDeployServer interface {
	Send(*ClientCallResponse) error
	grpc.ServerStream
}

type clientCallServiceRunDeployServer struct {
	grpc.ServerStream
}

func (x *clientCallServiceRunDeployServer) Send(m *ClientCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ClientCallService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eth_client_call.ClientCallService",
	HandlerType: (*ClientCallServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunDeploy",
			Handler:       _ClientCallService_RunDeploy_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "conn.proto",
}