// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fabex.proto

package fabex

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

type RequestRange struct {
	Startblock           int64    `protobuf:"varint,1,opt,name=startblock,proto3" json:"startblock,omitempty"`
	Endblock             int64    `protobuf:"varint,2,opt,name=endblock,proto3" json:"endblock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestRange) Reset()         { *m = RequestRange{} }
func (m *RequestRange) String() string { return proto.CompactTextString(m) }
func (*RequestRange) ProtoMessage()    {}
func (*RequestRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d7206373264ff4, []int{0}
}

func (m *RequestRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestRange.Unmarshal(m, b)
}
func (m *RequestRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestRange.Marshal(b, m, deterministic)
}
func (m *RequestRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestRange.Merge(m, src)
}
func (m *RequestRange) XXX_Size() int {
	return xxx_messageInfo_RequestRange.Size(m)
}
func (m *RequestRange) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestRange.DiscardUnknown(m)
}

var xxx_messageInfo_RequestRange proto.InternalMessageInfo

func (m *RequestRange) GetStartblock() int64 {
	if m != nil {
		return m.Startblock
	}
	return 0
}

func (m *RequestRange) GetEndblock() int64 {
	if m != nil {
		return m.Endblock
	}
	return 0
}

type Reply struct {
	Txid                 string   `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Blocknum             uint64   `protobuf:"varint,3,opt,name=blocknum,proto3" json:"blocknum,omitempty"`
	Payload              string   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d7206373264ff4, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *Reply) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Reply) GetBlocknum() uint64 {
	if m != nil {
		return m.Blocknum
	}
	return 0
}

func (m *Reply) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type RequestFilter struct {
	Txid                 string   `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Blocknum             uint64   `protobuf:"varint,3,opt,name=blocknum,proto3" json:"blocknum,omitempty"`
	Payload              string   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFilter) Reset()         { *m = RequestFilter{} }
func (m *RequestFilter) String() string { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()    {}
func (*RequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d7206373264ff4, []int{2}
}

func (m *RequestFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFilter.Unmarshal(m, b)
}
func (m *RequestFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFilter.Marshal(b, m, deterministic)
}
func (m *RequestFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFilter.Merge(m, src)
}
func (m *RequestFilter) XXX_Size() int {
	return xxx_messageInfo_RequestFilter.Size(m)
}
func (m *RequestFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFilter proto.InternalMessageInfo

func (m *RequestFilter) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *RequestFilter) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *RequestFilter) GetBlocknum() uint64 {
	if m != nil {
		return m.Blocknum
	}
	return 0
}

func (m *RequestFilter) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestRange)(nil), "fabex.RequestRange")
	proto.RegisterType((*Reply)(nil), "fabex.Reply")
	proto.RegisterType((*RequestFilter)(nil), "fabex.RequestFilter")
}

func init() { proto.RegisterFile("fabex.proto", fileDescriptor_d7d7206373264ff4) }

var fileDescriptor_d7d7206373264ff4 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0x4d, 0xac, 0x19, 0xdb, 0xcb, 0xa8, 0x10, 0x7a, 0x90, 0x92, 0x53, 0x4f, 0xc1,
	0x3f, 0x78, 0xf4, 0x12, 0xb0, 0x52, 0x4f, 0xb2, 0xf8, 0x05, 0x36, 0x66, 0x6a, 0x8b, 0xdb, 0xec,
	0x9a, 0x6c, 0x21, 0xfb, 0x6d, 0xfd, 0x28, 0xb2, 0x13, 0x1b, 0x2a, 0x5e, 0x7a, 0xe9, 0x6d, 0xde,
	0x1b, 0x7e, 0xbb, 0x8f, 0xc7, 0xc0, 0xf9, 0x52, 0x16, 0xd4, 0x66, 0xa6, 0xd6, 0x56, 0x63, 0xc4,
	0x22, 0x7d, 0x81, 0x91, 0xa0, 0xaf, 0x2d, 0x35, 0x56, 0xc8, 0xea, 0x83, 0xf0, 0x1a, 0xa0, 0xb1,
	0xb2, 0xb6, 0x85, 0xd2, 0xef, 0x9f, 0x49, 0x30, 0x0d, 0x66, 0x03, 0xb1, 0xe7, 0xe0, 0x04, 0xce,
	0xa8, 0x2a, 0xbb, 0xed, 0x09, 0x6f, 0x7b, 0x9d, 0x12, 0x44, 0x82, 0x8c, 0x72, 0x88, 0x10, 0xda,
	0x76, 0x5d, 0x32, 0x1e, 0x0b, 0x9e, 0xbd, 0xb7, 0x92, 0xcd, 0x8a, 0xa1, 0x58, 0xf0, 0xec, 0x1f,
	0x63, 0xb2, 0xda, 0x6e, 0x92, 0xc1, 0x34, 0x98, 0x85, 0xa2, 0xd7, 0x98, 0xc0, 0xd0, 0x48, 0xa7,
	0xb4, 0x2c, 0x93, 0x90, 0x91, 0x9d, 0x4c, 0x37, 0x30, 0xfe, 0x8d, 0x3c, 0x5f, 0x2b, 0x4b, 0xf5,
	0x71, 0xbf, 0xbb, 0xfb, 0x0e, 0x20, 0x9a, 0xfb, 0xae, 0x30, 0x83, 0xe1, 0x53, 0x6b, 0x94, 0xae,
	0x09, 0x2f, 0xb2, 0xae, 0xcb, 0xfd, 0xee, 0x26, 0xa3, 0xde, 0x34, 0xca, 0xdd, 0x04, 0x78, 0x0b,
	0xf1, 0x33, 0xd9, 0xdc, 0xbd, 0xb5, 0x8b, 0x12, 0x2f, 0xff, 0x12, 0x5d, 0xf4, 0x7f, 0xc8, 0x03,
	0x8c, 0x19, 0xc9, 0x77, 0xb9, 0x0e, 0xc3, 0x1e, 0xe1, 0xca, 0x63, 0x1e, 0x5a, 0x54, 0x4b, 0x9d,
	0xbb, 0xd7, 0x2e, 0xfc, 0x61, 0x78, 0x71, 0xca, 0x27, 0x71, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xfb, 0x38, 0x87, 0x2b, 0x21, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FabexClient is the client API for Fabex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FabexClient interface {
	Explore(ctx context.Context, in *RequestRange, opts ...grpc.CallOption) (Fabex_ExploreClient, error)
	GetByTxId(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Fabex_GetByTxIdClient, error)
	GetByBlocknum(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Fabex_GetByBlocknumClient, error)
	GetBlockInfoByPayload(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Fabex_GetBlockInfoByPayloadClient, error)
}

type fabexClient struct {
	cc *grpc.ClientConn
}

func NewFabexClient(cc *grpc.ClientConn) FabexClient {
	return &fabexClient{cc}
}

func (c *fabexClient) Explore(ctx context.Context, in *RequestRange, opts ...grpc.CallOption) (Fabex_ExploreClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[0], "/fabex.Fabex/Explore", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexExploreClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_ExploreClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type fabexExploreClient struct {
	grpc.ClientStream
}

func (x *fabexExploreClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fabexClient) GetByTxId(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Fabex_GetByTxIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[1], "/fabex.Fabex/GetByTxId", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexGetByTxIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_GetByTxIdClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type fabexGetByTxIdClient struct {
	grpc.ClientStream
}

func (x *fabexGetByTxIdClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fabexClient) GetByBlocknum(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Fabex_GetByBlocknumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[2], "/fabex.Fabex/GetByBlocknum", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexGetByBlocknumClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_GetByBlocknumClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type fabexGetByBlocknumClient struct {
	grpc.ClientStream
}

func (x *fabexGetByBlocknumClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fabexClient) GetBlockInfoByPayload(ctx context.Context, in *RequestFilter, opts ...grpc.CallOption) (Fabex_GetBlockInfoByPayloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[3], "/fabex.Fabex/GetBlockInfoByPayload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexGetBlockInfoByPayloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_GetBlockInfoByPayloadClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type fabexGetBlockInfoByPayloadClient struct {
	grpc.ClientStream
}

func (x *fabexGetBlockInfoByPayloadClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FabexServer is the server API for Fabex service.
type FabexServer interface {
	Explore(*RequestRange, Fabex_ExploreServer) error
	GetByTxId(*RequestFilter, Fabex_GetByTxIdServer) error
	GetByBlocknum(*RequestFilter, Fabex_GetByBlocknumServer) error
	GetBlockInfoByPayload(*RequestFilter, Fabex_GetBlockInfoByPayloadServer) error
}

// UnimplementedFabexServer can be embedded to have forward compatible implementations.
type UnimplementedFabexServer struct {
}

func (*UnimplementedFabexServer) Explore(req *RequestRange, srv Fabex_ExploreServer) error {
	return status.Errorf(codes.Unimplemented, "method Explore not implemented")
}
func (*UnimplementedFabexServer) GetByTxId(req *RequestFilter, srv Fabex_GetByTxIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetByTxId not implemented")
}
func (*UnimplementedFabexServer) GetByBlocknum(req *RequestFilter, srv Fabex_GetByBlocknumServer) error {
	return status.Errorf(codes.Unimplemented, "method GetByBlocknum not implemented")
}
func (*UnimplementedFabexServer) GetBlockInfoByPayload(req *RequestFilter, srv Fabex_GetBlockInfoByPayloadServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlockInfoByPayload not implemented")
}

func RegisterFabexServer(s *grpc.Server, srv FabexServer) {
	s.RegisterService(&_Fabex_serviceDesc, srv)
}

func _Fabex_Explore_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestRange)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).Explore(m, &fabexExploreServer{stream})
}

type Fabex_ExploreServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type fabexExploreServer struct {
	grpc.ServerStream
}

func (x *fabexExploreServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func _Fabex_GetByTxId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).GetByTxId(m, &fabexGetByTxIdServer{stream})
}

type Fabex_GetByTxIdServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type fabexGetByTxIdServer struct {
	grpc.ServerStream
}

func (x *fabexGetByTxIdServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func _Fabex_GetByBlocknum_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).GetByBlocknum(m, &fabexGetByBlocknumServer{stream})
}

type Fabex_GetByBlocknumServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type fabexGetByBlocknumServer struct {
	grpc.ServerStream
}

func (x *fabexGetByBlocknumServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func _Fabex_GetBlockInfoByPayload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).GetBlockInfoByPayload(m, &fabexGetBlockInfoByPayloadServer{stream})
}

type Fabex_GetBlockInfoByPayloadServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type fabexGetBlockInfoByPayloadServer struct {
	grpc.ServerStream
}

func (x *fabexGetBlockInfoByPayloadServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

var _Fabex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fabex.Fabex",
	HandlerType: (*FabexServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Explore",
			Handler:       _Fabex_Explore_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetByTxId",
			Handler:       _Fabex_GetByTxId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetByBlocknum",
			Handler:       _Fabex_GetByBlocknum_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBlockInfoByPayload",
			Handler:       _Fabex_GetBlockInfoByPayload_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fabex.proto",
}
