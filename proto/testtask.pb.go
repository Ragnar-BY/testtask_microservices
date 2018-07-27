// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testtask.proto

package proto

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

// The request message containing player name.
type CreatePlayerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlayerRequest) Reset()         { *m = CreatePlayerRequest{} }
func (m *CreatePlayerRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerRequest) ProtoMessage()    {}
func (*CreatePlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_testtask_0da8bd1a6661b039, []int{0}
}
func (m *CreatePlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayerRequest.Unmarshal(m, b)
}
func (m *CreatePlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayerRequest.Marshal(b, m, deterministic)
}
func (dst *CreatePlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerRequest.Merge(dst, src)
}
func (m *CreatePlayerRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePlayerRequest.Size(m)
}
func (m *CreatePlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerRequest proto.InternalMessageInfo

func (m *CreatePlayerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response containing player id
type CreatePlayerReply struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlayerReply) Reset()         { *m = CreatePlayerReply{} }
func (m *CreatePlayerReply) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerReply) ProtoMessage()    {}
func (*CreatePlayerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_testtask_0da8bd1a6661b039, []int{1}
}
func (m *CreatePlayerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayerReply.Unmarshal(m, b)
}
func (m *CreatePlayerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayerReply.Marshal(b, m, deterministic)
}
func (dst *CreatePlayerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerReply.Merge(dst, src)
}
func (m *CreatePlayerReply) XXX_Size() int {
	return xxx_messageInfo_CreatePlayerReply.Size(m)
}
func (m *CreatePlayerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerReply proto.InternalMessageInfo

func (m *CreatePlayerReply) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The request message containing player ID.
type PlayerIDRequest struct {
	PlayerID             int32    `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerIDRequest) Reset()         { *m = PlayerIDRequest{} }
func (m *PlayerIDRequest) String() string { return proto.CompactTextString(m) }
func (*PlayerIDRequest) ProtoMessage()    {}
func (*PlayerIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_testtask_0da8bd1a6661b039, []int{2}
}
func (m *PlayerIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerIDRequest.Unmarshal(m, b)
}
func (m *PlayerIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerIDRequest.Marshal(b, m, deterministic)
}
func (dst *PlayerIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerIDRequest.Merge(dst, src)
}
func (m *PlayerIDRequest) XXX_Size() int {
	return xxx_messageInfo_PlayerIDRequest.Size(m)
}
func (m *PlayerIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerIDRequest proto.InternalMessageInfo

func (m *PlayerIDRequest) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

// The response message containing player balance.
type PlayerBalanceReply struct {
	Balance              float32  `protobuf:"fixed32,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerBalanceReply) Reset()         { *m = PlayerBalanceReply{} }
func (m *PlayerBalanceReply) String() string { return proto.CompactTextString(m) }
func (*PlayerBalanceReply) ProtoMessage()    {}
func (*PlayerBalanceReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_testtask_0da8bd1a6661b039, []int{3}
}
func (m *PlayerBalanceReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerBalanceReply.Unmarshal(m, b)
}
func (m *PlayerBalanceReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerBalanceReply.Marshal(b, m, deterministic)
}
func (dst *PlayerBalanceReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerBalanceReply.Merge(dst, src)
}
func (m *PlayerBalanceReply) XXX_Size() int {
	return xxx_messageInfo_PlayerBalanceReply.Size(m)
}
func (m *PlayerBalanceReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerBalanceReply.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerBalanceReply proto.InternalMessageInfo

func (m *PlayerBalanceReply) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

// The request message containing player ID and points.
type PlayerIDPointRequest struct {
	PlayerID             int32    `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Points               float32  `protobuf:"fixed32,2,opt,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerIDPointRequest) Reset()         { *m = PlayerIDPointRequest{} }
func (m *PlayerIDPointRequest) String() string { return proto.CompactTextString(m) }
func (*PlayerIDPointRequest) ProtoMessage()    {}
func (*PlayerIDPointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_testtask_0da8bd1a6661b039, []int{4}
}
func (m *PlayerIDPointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerIDPointRequest.Unmarshal(m, b)
}
func (m *PlayerIDPointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerIDPointRequest.Marshal(b, m, deterministic)
}
func (dst *PlayerIDPointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerIDPointRequest.Merge(dst, src)
}
func (m *PlayerIDPointRequest) XXX_Size() int {
	return xxx_messageInfo_PlayerIDPointRequest.Size(m)
}
func (m *PlayerIDPointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerIDPointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerIDPointRequest proto.InternalMessageInfo

func (m *PlayerIDPointRequest) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *PlayerIDPointRequest) GetPoints() float32 {
	if m != nil {
		return m.Points
	}
	return 0
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_testtask_0da8bd1a6661b039, []int{5}
}
func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (dst *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(dst, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreatePlayerRequest)(nil), "testtask.CreatePlayerRequest")
	proto.RegisterType((*CreatePlayerReply)(nil), "testtask.CreatePlayerReply")
	proto.RegisterType((*PlayerIDRequest)(nil), "testtask.PlayerIDRequest")
	proto.RegisterType((*PlayerBalanceReply)(nil), "testtask.PlayerBalanceReply")
	proto.RegisterType((*PlayerIDPointRequest)(nil), "testtask.PlayerIDPointRequest")
	proto.RegisterType((*Nothing)(nil), "testtask.Nothing")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlayerServiceClient is the client API for PlayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlayerServiceClient interface {
	// CreateNewPlayer creates new player and returns its id.
	CreateNewPlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerReply, error)
	// GetPlayerPoints returns player points.
	GetPlayerPoints(ctx context.Context, in *PlayerIDRequest, opts ...grpc.CallOption) (*PlayerBalanceReply, error)
	// TakePointsFromPlayer take points from player and returns new balance.
	TakePointsFromPlayer(ctx context.Context, in *PlayerIDPointRequest, opts ...grpc.CallOption) (*PlayerBalanceReply, error)
	// FundPointsFromPlayer fund points to player and returns new balance.
	FundPointsToPlayer(ctx context.Context, in *PlayerIDPointRequest, opts ...grpc.CallOption) (*PlayerBalanceReply, error)
	RemovePlayer(ctx context.Context, in *PlayerIDRequest, opts ...grpc.CallOption) (*Nothing, error)
}

type playerServiceClient struct {
	cc *grpc.ClientConn
}

func NewPlayerServiceClient(cc *grpc.ClientConn) PlayerServiceClient {
	return &playerServiceClient{cc}
}

func (c *playerServiceClient) CreateNewPlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerReply, error) {
	out := new(CreatePlayerReply)
	err := c.cc.Invoke(ctx, "/testtask.PlayerService/CreateNewPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) GetPlayerPoints(ctx context.Context, in *PlayerIDRequest, opts ...grpc.CallOption) (*PlayerBalanceReply, error) {
	out := new(PlayerBalanceReply)
	err := c.cc.Invoke(ctx, "/testtask.PlayerService/GetPlayerPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) TakePointsFromPlayer(ctx context.Context, in *PlayerIDPointRequest, opts ...grpc.CallOption) (*PlayerBalanceReply, error) {
	out := new(PlayerBalanceReply)
	err := c.cc.Invoke(ctx, "/testtask.PlayerService/TakePointsFromPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) FundPointsToPlayer(ctx context.Context, in *PlayerIDPointRequest, opts ...grpc.CallOption) (*PlayerBalanceReply, error) {
	out := new(PlayerBalanceReply)
	err := c.cc.Invoke(ctx, "/testtask.PlayerService/FundPointsToPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) RemovePlayer(ctx context.Context, in *PlayerIDRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/testtask.PlayerService/RemovePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServiceServer is the server API for PlayerService service.
type PlayerServiceServer interface {
	// CreateNewPlayer creates new player and returns its id.
	CreateNewPlayer(context.Context, *CreatePlayerRequest) (*CreatePlayerReply, error)
	// GetPlayerPoints returns player points.
	GetPlayerPoints(context.Context, *PlayerIDRequest) (*PlayerBalanceReply, error)
	// TakePointsFromPlayer take points from player and returns new balance.
	TakePointsFromPlayer(context.Context, *PlayerIDPointRequest) (*PlayerBalanceReply, error)
	// FundPointsFromPlayer fund points to player and returns new balance.
	FundPointsToPlayer(context.Context, *PlayerIDPointRequest) (*PlayerBalanceReply, error)
	RemovePlayer(context.Context, *PlayerIDRequest) (*Nothing, error)
}

func RegisterPlayerServiceServer(s *grpc.Server, srv PlayerServiceServer) {
	s.RegisterService(&_PlayerService_serviceDesc, srv)
}

func _PlayerService_CreateNewPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).CreateNewPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testtask.PlayerService/CreateNewPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).CreateNewPlayer(ctx, req.(*CreatePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_GetPlayerPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).GetPlayerPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testtask.PlayerService/GetPlayerPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).GetPlayerPoints(ctx, req.(*PlayerIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_TakePointsFromPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerIDPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).TakePointsFromPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testtask.PlayerService/TakePointsFromPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).TakePointsFromPlayer(ctx, req.(*PlayerIDPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_FundPointsToPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerIDPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).FundPointsToPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testtask.PlayerService/FundPointsToPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).FundPointsToPlayer(ctx, req.(*PlayerIDPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_RemovePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).RemovePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testtask.PlayerService/RemovePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).RemovePlayer(ctx, req.(*PlayerIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlayerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testtask.PlayerService",
	HandlerType: (*PlayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewPlayer",
			Handler:    _PlayerService_CreateNewPlayer_Handler,
		},
		{
			MethodName: "GetPlayerPoints",
			Handler:    _PlayerService_GetPlayerPoints_Handler,
		},
		{
			MethodName: "TakePointsFromPlayer",
			Handler:    _PlayerService_TakePointsFromPlayer_Handler,
		},
		{
			MethodName: "FundPointsToPlayer",
			Handler:    _PlayerService_FundPointsToPlayer_Handler,
		},
		{
			MethodName: "RemovePlayer",
			Handler:    _PlayerService_RemovePlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testtask.proto",
}

func init() { proto.RegisterFile("testtask.proto", fileDescriptor_testtask_0da8bd1a6661b039) }

var fileDescriptor_testtask_0da8bd1a6661b039 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4b, 0x4f, 0xbb, 0x40,
	0x14, 0xc5, 0x5b, 0xfe, 0x7f, 0xfb, 0xb8, 0xd1, 0x36, 0xbd, 0x36, 0x06, 0xf1, 0x11, 0x33, 0x6e,
	0x74, 0x21, 0x0b, 0xdd, 0xbb, 0x50, 0x53, 0xa3, 0x31, 0xb5, 0x41, 0xe2, 0x7e, 0x5a, 0x6e, 0x94,
	0x14, 0x18, 0x84, 0x69, 0x0d, 0xdf, 0xc5, 0x0f, 0x6b, 0x9c, 0x19, 0xda, 0x5a, 0x9f, 0x0b, 0x77,
	0x9c, 0x39, 0x67, 0x7e, 0xdc, 0x7b, 0x00, 0x5a, 0x92, 0x72, 0x29, 0x79, 0x3e, 0x76, 0xd3, 0x4c,
	0x48, 0x81, 0x8d, 0x52, 0xb3, 0x43, 0x58, 0x3f, 0xcf, 0x88, 0x4b, 0x1a, 0x44, 0xbc, 0xa0, 0xcc,
	0xa3, 0xa7, 0x09, 0xe5, 0x12, 0x11, 0xfe, 0x27, 0x3c, 0x26, 0xbb, 0xba, 0x57, 0x3d, 0x68, 0x7a,
	0xea, 0x99, 0xed, 0x43, 0xe7, 0x7d, 0x34, 0x8d, 0x0a, 0x6c, 0x81, 0x15, 0x06, 0x2a, 0xb6, 0xe2,
	0x59, 0x61, 0xc0, 0x8e, 0xa0, 0xad, 0xed, 0xab, 0x8b, 0x92, 0xe5, 0x40, 0x23, 0x35, 0x47, 0x26,
	0x38, 0xd3, 0xcc, 0x05, 0xd4, 0xf1, 0x33, 0x1e, 0xf1, 0x64, 0x44, 0x1a, 0x6a, 0x43, 0x7d, 0xa8,
	0xb5, 0xba, 0x60, 0x79, 0xa5, 0x64, 0xd7, 0xd0, 0x2d, 0xf1, 0x03, 0x11, 0x26, 0xf2, 0x17, 0xef,
	0xc0, 0x0d, 0xa8, 0xa5, 0x6f, 0xd9, 0xdc, 0xb6, 0x14, 0xcc, 0x28, 0xd6, 0x84, 0x7a, 0x5f, 0xc8,
	0xc7, 0x30, 0x79, 0x38, 0x7e, 0xf9, 0x07, 0x6b, 0x9a, 0x7b, 0x47, 0xd9, 0x34, 0x1c, 0x11, 0xde,
	0x42, 0x5b, 0x2f, 0xdb, 0xa7, 0x67, 0xed, 0xe0, 0x8e, 0x3b, 0x6b, 0xf1, 0x93, 0xca, 0x9c, 0xad,
	0xaf, 0xec, 0x34, 0x2a, 0x58, 0x05, 0x6f, 0xa0, 0x7d, 0x49, 0x52, 0x9f, 0xa9, 0xd1, 0x73, 0xdc,
	0x9c, 0xdf, 0x58, 0xea, 0xcc, 0xd9, 0x5e, 0xb6, 0x16, 0xfb, 0x61, 0x15, 0xbc, 0x87, 0xae, 0xcf,
	0xc7, 0xa4, 0x41, 0xbd, 0x4c, 0xc4, 0x66, 0xc6, 0xdd, 0x8f, 0xc8, 0xc5, 0x9e, 0x7e, 0xe4, 0xfa,
	0x80, 0xbd, 0x49, 0x12, 0x68, 0xae, 0x2f, 0xfe, 0x88, 0x7a, 0x0a, 0xab, 0x1e, 0xc5, 0x62, 0x6a,
	0x2a, 0xf9, 0x6e, 0xf1, 0xce, 0xdc, 0x32, 0x1f, 0x87, 0x55, 0x86, 0x35, 0xf5, 0xd7, 0x9e, 0xbc,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xc4, 0x2c, 0xfd, 0xc7, 0x02, 0x00, 0x00,
}