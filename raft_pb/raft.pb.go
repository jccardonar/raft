// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raft.proto

package raft_pb

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

// RequestVoreRequest are initiated by node in Candidate state, on election.
type RequestVoteRequest struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CommittedSequence    int64    `protobuf:"varint,2,opt,name=committed_sequence,json=committedSequence,proto3" json:"committed_sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVoteRequest) Reset()         { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()    {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{0}
}

func (m *RequestVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteRequest.Unmarshal(m, b)
}
func (m *RequestVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteRequest.Marshal(b, m, deterministic)
}
func (m *RequestVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteRequest.Merge(m, src)
}
func (m *RequestVoteRequest) XXX_Size() int {
	return xxx_messageInfo_RequestVoteRequest.Size(m)
}
func (m *RequestVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteRequest proto.InternalMessageInfo

func (m *RequestVoteRequest) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetCommittedSequence() int64 {
	if m != nil {
		return m.CommittedSequence
	}
	return 0
}

type RequestVoteReply struct {
	Ack bool `protobuf:"varint,1,opt,name=Ack,proto3" json:"Ack,omitempty"`
	// This reply is in response to this request.
	Request              *AppendEntryRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RequestVoteReply) Reset()         { *m = RequestVoteReply{} }
func (m *RequestVoteReply) String() string { return proto.CompactTextString(m) }
func (*RequestVoteReply) ProtoMessage()    {}
func (*RequestVoteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{1}
}

func (m *RequestVoteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteReply.Unmarshal(m, b)
}
func (m *RequestVoteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteReply.Marshal(b, m, deterministic)
}
func (m *RequestVoteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteReply.Merge(m, src)
}
func (m *RequestVoteReply) XXX_Size() int {
	return xxx_messageInfo_RequestVoteReply.Size(m)
}
func (m *RequestVoteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteReply.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteReply proto.InternalMessageInfo

func (m *RequestVoteReply) GetAck() bool {
	if m != nil {
		return m.Ack
	}
	return false
}

func (m *RequestVoteReply) GetRequest() *AppendEntryRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// AppendEntryRequest are initiated by Leader nodes to request replication of log.
// AppendEntryRequest with no log entries are sent by Leader as keepalives,
// to stave off elections, and to convey highest committed sequence.
type AppendEntryRequest struct {
	// Term of leadership.
	Term int64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// Largest sequence committed.
	CommittedSequence int64 `protobuf:"varint,2,opt,name=committed_sequence,json=committedSequence,proto3" json:"committed_sequence,omitempty"`
	// Sequence for this log entry.
	Sequence int64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Log entry for this sequence.
	LogEntry             []byte   `protobuf:"bytes,4,opt,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEntryRequest) Reset()         { *m = AppendEntryRequest{} }
func (m *AppendEntryRequest) String() string { return proto.CompactTextString(m) }
func (*AppendEntryRequest) ProtoMessage()    {}
func (*AppendEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{2}
}

func (m *AppendEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntryRequest.Unmarshal(m, b)
}
func (m *AppendEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntryRequest.Marshal(b, m, deterministic)
}
func (m *AppendEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntryRequest.Merge(m, src)
}
func (m *AppendEntryRequest) XXX_Size() int {
	return xxx_messageInfo_AppendEntryRequest.Size(m)
}
func (m *AppendEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntryRequest proto.InternalMessageInfo

func (m *AppendEntryRequest) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntryRequest) GetCommittedSequence() int64 {
	if m != nil {
		return m.CommittedSequence
	}
	return 0
}

func (m *AppendEntryRequest) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *AppendEntryRequest) GetLogEntry() []byte {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

type AppendEntryReply struct {
	// This reply is in response to this request.
	Request              *AppendEntryRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AppendEntryReply) Reset()         { *m = AppendEntryReply{} }
func (m *AppendEntryReply) String() string { return proto.CompactTextString(m) }
func (*AppendEntryReply) ProtoMessage()    {}
func (*AppendEntryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{3}
}

func (m *AppendEntryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntryReply.Unmarshal(m, b)
}
func (m *AppendEntryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntryReply.Marshal(b, m, deterministic)
}
func (m *AppendEntryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntryReply.Merge(m, src)
}
func (m *AppendEntryReply) XXX_Size() int {
	return xxx_messageInfo_AppendEntryReply.Size(m)
}
func (m *AppendEntryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntryReply.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntryReply proto.InternalMessageInfo

func (m *AppendEntryReply) GetRequest() *AppendEntryRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// TimeoutNowRequest
type TimeoutNowRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeoutNowRequest) Reset()         { *m = TimeoutNowRequest{} }
func (m *TimeoutNowRequest) String() string { return proto.CompactTextString(m) }
func (*TimeoutNowRequest) ProtoMessage()    {}
func (*TimeoutNowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{4}
}

func (m *TimeoutNowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeoutNowRequest.Unmarshal(m, b)
}
func (m *TimeoutNowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeoutNowRequest.Marshal(b, m, deterministic)
}
func (m *TimeoutNowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeoutNowRequest.Merge(m, src)
}
func (m *TimeoutNowRequest) XXX_Size() int {
	return xxx_messageInfo_TimeoutNowRequest.Size(m)
}
func (m *TimeoutNowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeoutNowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimeoutNowRequest proto.InternalMessageInfo

type TimeoutNowReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeoutNowReply) Reset()         { *m = TimeoutNowReply{} }
func (m *TimeoutNowReply) String() string { return proto.CompactTextString(m) }
func (*TimeoutNowReply) ProtoMessage()    {}
func (*TimeoutNowReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{5}
}

func (m *TimeoutNowReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeoutNowReply.Unmarshal(m, b)
}
func (m *TimeoutNowReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeoutNowReply.Marshal(b, m, deterministic)
}
func (m *TimeoutNowReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeoutNowReply.Merge(m, src)
}
func (m *TimeoutNowReply) XXX_Size() int {
	return xxx_messageInfo_TimeoutNowReply.Size(m)
}
func (m *TimeoutNowReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeoutNowReply.DiscardUnknown(m)
}

var xxx_messageInfo_TimeoutNowReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RequestVoteRequest)(nil), "raft_pb.RequestVoteRequest")
	proto.RegisterType((*RequestVoteReply)(nil), "raft_pb.RequestVoteReply")
	proto.RegisterType((*AppendEntryRequest)(nil), "raft_pb.AppendEntryRequest")
	proto.RegisterType((*AppendEntryReply)(nil), "raft_pb.AppendEntryReply")
	proto.RegisterType((*TimeoutNowRequest)(nil), "raft_pb.TimeoutNowRequest")
	proto.RegisterType((*TimeoutNowReply)(nil), "raft_pb.TimeoutNowReply")
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor_b042552c306ae59b) }

var fileDescriptor_b042552c306ae59b = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x31, 0x4f, 0xf3, 0x30,
	0x14, 0x94, 0xbf, 0x56, 0x5f, 0xcb, 0x0b, 0x82, 0xe4, 0xb1, 0x84, 0x74, 0xa9, 0x32, 0x75, 0x21,
	0x43, 0x11, 0x3f, 0xa0, 0x43, 0x90, 0x58, 0x18, 0x5c, 0x04, 0x03, 0x43, 0xd4, 0xa6, 0x6e, 0x15,
	0x91, 0xc4, 0xc6, 0x75, 0x41, 0xf9, 0x17, 0xfc, 0x47, 0xfe, 0x08, 0xb2, 0x71, 0x23, 0x57, 0xa9,
	0x90, 0x90, 0xd8, 0xce, 0xbe, 0xd3, 0xbd, 0x3b, 0xbd, 0x07, 0x20, 0x17, 0x6b, 0x95, 0x08, 0xc9,
	0x15, 0xc7, 0x81, 0xc6, 0x99, 0x58, 0xc6, 0x4f, 0x80, 0x94, 0xbd, 0xee, 0xd8, 0x56, 0x3d, 0x72,
	0xc5, 0x2c, 0x44, 0x84, 0xbe, 0x62, 0xb2, 0x0a, 0xc9, 0x98, 0x4c, 0x7a, 0xd4, 0x60, 0xbc, 0x02,
	0xcc, 0x79, 0x55, 0x15, 0x4a, 0xb1, 0x55, 0xb6, 0xd5, 0xc2, 0x3a, 0x67, 0xe1, 0x3f, 0xa3, 0x08,
	0x5a, 0x66, 0x6e, 0x89, 0xf8, 0x19, 0xfc, 0x03, 0x63, 0x51, 0x36, 0xe8, 0x43, 0x6f, 0x96, 0xbf,
	0x18, 0xd7, 0x21, 0xd5, 0x10, 0x6f, 0x60, 0x20, 0xbf, 0x55, 0xc6, 0xc9, 0x9b, 0x8e, 0x12, 0x9b,
	0x2c, 0x99, 0x09, 0xc1, 0xea, 0x55, 0x5a, 0x2b, 0xd9, 0x58, 0x23, 0xba, 0xd7, 0xc6, 0x1f, 0x04,
	0xb0, 0xcb, 0xff, 0x41, 0x6c, 0x8c, 0x60, 0xd8, 0x8a, 0x7a, 0x46, 0xd4, 0xbe, 0x71, 0x04, 0x27,
	0x25, 0xdf, 0x64, 0x4c, 0x8f, 0x0c, 0xfb, 0x63, 0x32, 0x39, 0xa5, 0xc3, 0x92, 0x6f, 0x4c, 0x84,
	0xf8, 0x0e, 0xfc, 0x83, 0x44, 0xba, 0xaf, 0xd3, 0x8e, 0xfc, 0xa2, 0xdd, 0x05, 0x04, 0x0f, 0x45,
	0xc5, 0xf8, 0x4e, 0xdd, 0xf3, 0x77, 0xcb, 0xc6, 0x01, 0x9c, 0xbb, 0x9f, 0xa2, 0x6c, 0xa6, 0x9f,
	0x04, 0x3c, 0xba, 0x58, 0xab, 0x39, 0x93, 0x6f, 0x45, 0xce, 0x30, 0x05, 0xcf, 0xb1, 0xc5, 0x9f,
	0x86, 0x45, 0x97, 0xc7, 0x49, 0x9d, 0x3a, 0x05, 0xcf, 0xd9, 0x9c, 0x63, 0xd3, 0x3d, 0x14, 0xc7,
	0xa6, 0xb3, 0xec, 0x5b, 0x38, 0xb3, 0x81, 0xf7, 0xeb, 0x89, 0x5a, 0x71, 0xa7, 0x5e, 0x14, 0x1e,
	0xe5, 0x44, 0xd9, 0x2c, 0xff, 0x9b, 0x8b, 0xbd, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xce,
	0x26, 0x29, 0xbf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RaftServiceClient is the client API for RaftService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RaftServiceClient interface {
	AppendEntry(ctx context.Context, in *AppendEntryRequest, opts ...grpc.CallOption) (*AppendEntryReply, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteReply, error)
	TimeoutRequest(ctx context.Context, in *TimeoutNowRequest, opts ...grpc.CallOption) (*TimeoutNowReply, error)
}

type raftServiceClient struct {
	cc *grpc.ClientConn
}

func NewRaftServiceClient(cc *grpc.ClientConn) RaftServiceClient {
	return &raftServiceClient{cc}
}

func (c *raftServiceClient) AppendEntry(ctx context.Context, in *AppendEntryRequest, opts ...grpc.CallOption) (*AppendEntryReply, error) {
	out := new(AppendEntryReply)
	err := c.cc.Invoke(ctx, "/raft_pb.RaftService/AppendEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteReply, error) {
	out := new(RequestVoteReply)
	err := c.cc.Invoke(ctx, "/raft_pb.RaftService/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) TimeoutRequest(ctx context.Context, in *TimeoutNowRequest, opts ...grpc.CallOption) (*TimeoutNowReply, error) {
	out := new(TimeoutNowReply)
	err := c.cc.Invoke(ctx, "/raft_pb.RaftService/TimeoutRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftServiceServer is the server API for RaftService service.
type RaftServiceServer interface {
	AppendEntry(context.Context, *AppendEntryRequest) (*AppendEntryReply, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteReply, error)
	TimeoutRequest(context.Context, *TimeoutNowRequest) (*TimeoutNowReply, error)
}

// UnimplementedRaftServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRaftServiceServer struct {
}

func (*UnimplementedRaftServiceServer) AppendEntry(ctx context.Context, req *AppendEntryRequest) (*AppendEntryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntry not implemented")
}
func (*UnimplementedRaftServiceServer) RequestVote(ctx context.Context, req *RequestVoteRequest) (*RequestVoteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (*UnimplementedRaftServiceServer) TimeoutRequest(ctx context.Context, req *TimeoutNowRequest) (*TimeoutNowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeoutRequest not implemented")
}

func RegisterRaftServiceServer(s *grpc.Server, srv RaftServiceServer) {
	s.RegisterService(&_RaftService_serviceDesc, srv)
}

func _RaftService_AppendEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).AppendEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft_pb.RaftService/AppendEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).AppendEntry(ctx, req.(*AppendEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft_pb.RaftService/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_TimeoutRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeoutNowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).TimeoutRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft_pb.RaftService/TimeoutRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).TimeoutRequest(ctx, req.(*TimeoutNowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "raft_pb.RaftService",
	HandlerType: (*RaftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEntry",
			Handler:    _RaftService_AppendEntry_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _RaftService_RequestVote_Handler,
		},
		{
			MethodName: "TimeoutRequest",
			Handler:    _RaftService_TimeoutRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft.proto",
}
