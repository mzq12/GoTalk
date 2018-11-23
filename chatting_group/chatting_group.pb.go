// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatting_group.proto

package pb

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

type GroupInfo struct {
	Gid                  int64    `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Appid                int64    `protobuf:"varint,2,opt,name=appid,proto3" json:"appid,omitempty"`
	Uid                  int64    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	GName                string   `protobuf:"bytes,4,opt,name=g_name,json=gName,proto3" json:"g_name,omitempty"`
	Super                bool     `protobuf:"varint,5,opt,name=super,proto3" json:"super,omitempty"`
	Announcement         string   `protobuf:"bytes,6,opt,name=announcement,proto3" json:"announcement,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupInfo) Reset()         { *m = GroupInfo{} }
func (m *GroupInfo) String() string { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()    {}
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatting_group_99f87f33d2c8f5c0, []int{0}
}
func (m *GroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupInfo.Unmarshal(m, b)
}
func (m *GroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupInfo.Marshal(b, m, deterministic)
}
func (dst *GroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupInfo.Merge(dst, src)
}
func (m *GroupInfo) XXX_Size() int {
	return xxx_messageInfo_GroupInfo.Size(m)
}
func (m *GroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GroupInfo proto.InternalMessageInfo

func (m *GroupInfo) GetGid() int64 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *GroupInfo) GetAppid() int64 {
	if m != nil {
		return m.Appid
	}
	return 0
}

func (m *GroupInfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GroupInfo) GetGName() string {
	if m != nil {
		return m.GName
	}
	return ""
}

func (m *GroupInfo) GetSuper() bool {
	if m != nil {
		return m.Super
	}
	return false
}

func (m *GroupInfo) GetAnnouncement() string {
	if m != nil {
		return m.Announcement
	}
	return ""
}

type GroupMember struct {
	Gid                  int64    `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	InTime               int64    `protobuf:"varint,3,opt,name=in_time,json=inTime,proto3" json:"in_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupMember) Reset()         { *m = GroupMember{} }
func (m *GroupMember) String() string { return proto.CompactTextString(m) }
func (*GroupMember) ProtoMessage()    {}
func (*GroupMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatting_group_99f87f33d2c8f5c0, []int{1}
}
func (m *GroupMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupMember.Unmarshal(m, b)
}
func (m *GroupMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupMember.Marshal(b, m, deterministic)
}
func (dst *GroupMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupMember.Merge(dst, src)
}
func (m *GroupMember) XXX_Size() int {
	return xxx_messageInfo_GroupMember.Size(m)
}
func (m *GroupMember) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupMember.DiscardUnknown(m)
}

var xxx_messageInfo_GroupMember proto.InternalMessageInfo

func (m *GroupMember) GetGid() int64 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *GroupMember) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GroupMember) GetInTime() int64 {
	if m != nil {
		return m.InTime
	}
	return 0
}

type FindAllGroupRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllGroupRequest) Reset()         { *m = FindAllGroupRequest{} }
func (m *FindAllGroupRequest) String() string { return proto.CompactTextString(m) }
func (*FindAllGroupRequest) ProtoMessage()    {}
func (*FindAllGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatting_group_99f87f33d2c8f5c0, []int{2}
}
func (m *FindAllGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllGroupRequest.Unmarshal(m, b)
}
func (m *FindAllGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllGroupRequest.Marshal(b, m, deterministic)
}
func (dst *FindAllGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllGroupRequest.Merge(dst, src)
}
func (m *FindAllGroupRequest) XXX_Size() int {
	return xxx_messageInfo_FindAllGroupRequest.Size(m)
}
func (m *FindAllGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllGroupRequest proto.InternalMessageInfo

type FindAllGroupResponse struct {
	Groups               []*GroupInfo `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	ErrorCode            int32        `protobuf:"varint,2,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FindAllGroupResponse) Reset()         { *m = FindAllGroupResponse{} }
func (m *FindAllGroupResponse) String() string { return proto.CompactTextString(m) }
func (*FindAllGroupResponse) ProtoMessage()    {}
func (*FindAllGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatting_group_99f87f33d2c8f5c0, []int{3}
}
func (m *FindAllGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllGroupResponse.Unmarshal(m, b)
}
func (m *FindAllGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllGroupResponse.Marshal(b, m, deterministic)
}
func (dst *FindAllGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllGroupResponse.Merge(dst, src)
}
func (m *FindAllGroupResponse) XXX_Size() int {
	return xxx_messageInfo_FindAllGroupResponse.Size(m)
}
func (m *FindAllGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllGroupResponse proto.InternalMessageInfo

func (m *FindAllGroupResponse) GetGroups() []*GroupInfo {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *FindAllGroupResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type GetGroupMembersRequest struct {
	Gid                  int64    `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupMembersRequest) Reset()         { *m = GetGroupMembersRequest{} }
func (m *GetGroupMembersRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupMembersRequest) ProtoMessage()    {}
func (*GetGroupMembersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatting_group_99f87f33d2c8f5c0, []int{4}
}
func (m *GetGroupMembersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupMembersRequest.Unmarshal(m, b)
}
func (m *GetGroupMembersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupMembersRequest.Marshal(b, m, deterministic)
}
func (dst *GetGroupMembersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupMembersRequest.Merge(dst, src)
}
func (m *GetGroupMembersRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupMembersRequest.Size(m)
}
func (m *GetGroupMembersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupMembersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupMembersRequest proto.InternalMessageInfo

func (m *GetGroupMembersRequest) GetGid() int64 {
	if m != nil {
		return m.Gid
	}
	return 0
}

type GetGroupMembersResponse struct {
	Members              []*GroupMember `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	ErrorCode            int32          `protobuf:"varint,2,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetGroupMembersResponse) Reset()         { *m = GetGroupMembersResponse{} }
func (m *GetGroupMembersResponse) String() string { return proto.CompactTextString(m) }
func (*GetGroupMembersResponse) ProtoMessage()    {}
func (*GetGroupMembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatting_group_99f87f33d2c8f5c0, []int{5}
}
func (m *GetGroupMembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupMembersResponse.Unmarshal(m, b)
}
func (m *GetGroupMembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupMembersResponse.Marshal(b, m, deterministic)
}
func (dst *GetGroupMembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupMembersResponse.Merge(dst, src)
}
func (m *GetGroupMembersResponse) XXX_Size() int {
	return xxx_messageInfo_GetGroupMembersResponse.Size(m)
}
func (m *GetGroupMembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupMembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupMembersResponse proto.InternalMessageInfo

func (m *GetGroupMembersResponse) GetMembers() []*GroupMember {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *GetGroupMembersResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*GroupInfo)(nil), "pb.GroupInfo")
	proto.RegisterType((*GroupMember)(nil), "pb.GroupMember")
	proto.RegisterType((*FindAllGroupRequest)(nil), "pb.FindAllGroupRequest")
	proto.RegisterType((*FindAllGroupResponse)(nil), "pb.FindAllGroupResponse")
	proto.RegisterType((*GetGroupMembersRequest)(nil), "pb.GetGroupMembersRequest")
	proto.RegisterType((*GetGroupMembersResponse)(nil), "pb.GetGroupMembersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChattingGroupClient is the client API for ChattingGroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChattingGroupClient interface {
	// 查询所有群信息
	FindAllGroup(ctx context.Context, in *FindAllGroupRequest, opts ...grpc.CallOption) (*FindAllGroupResponse, error)
	// 查询群成员信息
	GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error)
}

type chattingGroupClient struct {
	cc *grpc.ClientConn
}

func NewChattingGroupClient(cc *grpc.ClientConn) ChattingGroupClient {
	return &chattingGroupClient{cc}
}

func (c *chattingGroupClient) FindAllGroup(ctx context.Context, in *FindAllGroupRequest, opts ...grpc.CallOption) (*FindAllGroupResponse, error) {
	out := new(FindAllGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ChattingGroup/FindAllGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chattingGroupClient) GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error) {
	out := new(GetGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/pb.ChattingGroup/GetGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChattingGroupServer is the server API for ChattingGroup service.
type ChattingGroupServer interface {
	// 查询所有群信息
	FindAllGroup(context.Context, *FindAllGroupRequest) (*FindAllGroupResponse, error)
	// 查询群成员信息
	GetGroupMembers(context.Context, *GetGroupMembersRequest) (*GetGroupMembersResponse, error)
}

func RegisterChattingGroupServer(s *grpc.Server, srv ChattingGroupServer) {
	s.RegisterService(&_ChattingGroup_serviceDesc, srv)
}

func _ChattingGroup_FindAllGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChattingGroupServer).FindAllGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChattingGroup/FindAllGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChattingGroupServer).FindAllGroup(ctx, req.(*FindAllGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChattingGroup_GetGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChattingGroupServer).GetGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChattingGroup/GetGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChattingGroupServer).GetGroupMembers(ctx, req.(*GetGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChattingGroup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ChattingGroup",
	HandlerType: (*ChattingGroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllGroup",
			Handler:    _ChattingGroup_FindAllGroup_Handler,
		},
		{
			MethodName: "GetGroupMembers",
			Handler:    _ChattingGroup_GetGroupMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatting_group.proto",
}

func init() {
	proto.RegisterFile("chatting_group.proto", fileDescriptor_chatting_group_99f87f33d2c8f5c0)
}

var fileDescriptor_chatting_group_99f87f33d2c8f5c0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xff, 0x34, 0x7f, 0x52, 0x3a, 0x6d, 0x55, 0x64, 0x5a, 0x6a, 0x15, 0x16, 0x91, 0x25,
	0xa4, 0xc0, 0xa2, 0x8b, 0xf2, 0x04, 0xa8, 0x12, 0x15, 0x08, 0x58, 0x44, 0xec, 0x58, 0x54, 0x49,
	0x33, 0x04, 0x4b, 0x8d, 0x6d, 0x72, 0x79, 0x15, 0x1e, 0x80, 0x27, 0x45, 0x71, 0xdc, 0x1b, 0x8d,
	0xd8, 0xc5, 0x67, 0xce, 0x4c, 0xbe, 0x33, 0x1a, 0x18, 0xae, 0x3e, 0xc2, 0xa2, 0xe0, 0x22, 0x59,
	0x26, 0x99, 0x2c, 0xd5, 0x54, 0x65, 0xb2, 0x90, 0xa4, 0xa5, 0x22, 0xf6, 0x65, 0x41, 0x67, 0x51,
	0x69, 0x0f, 0xe2, 0x5d, 0x92, 0x53, 0xb0, 0x13, 0x1e, 0x53, 0xcb, 0xb3, 0x7c, 0x3b, 0xa8, 0x3e,
	0xc9, 0x10, 0x9c, 0x50, 0x29, 0x1e, 0xd3, 0x96, 0xd6, 0xea, 0x47, 0xe5, 0x2b, 0x79, 0x4c, 0xed,
	0xda, 0x57, 0xf2, 0x98, 0x8c, 0xc0, 0x4d, 0x96, 0x22, 0x4c, 0x91, 0xfe, 0xf7, 0x2c, 0xbf, 0x13,
	0x38, 0xc9, 0x4b, 0x98, 0x62, 0xd5, 0x9e, 0x97, 0x0a, 0x33, 0xea, 0x78, 0x96, 0x7f, 0x12, 0xd4,
	0x0f, 0xc2, 0xa0, 0x17, 0x0a, 0x21, 0x4b, 0xb1, 0xc2, 0x14, 0x45, 0x41, 0x5d, 0xdd, 0x72, 0xa0,
	0xb1, 0x47, 0xe8, 0x6a, 0xae, 0x67, 0x4c, 0x23, 0xcc, 0x1a, 0xc8, 0x0c, 0x43, 0x6b, 0xc7, 0x30,
	0x86, 0x36, 0x17, 0xcb, 0x82, 0xa7, 0x68, 0xc8, 0x5c, 0x2e, 0x5e, 0x79, 0x8a, 0x6c, 0x04, 0x67,
	0xf7, 0x5c, 0xc4, 0x77, 0xeb, 0xb5, 0x1e, 0x19, 0xe0, 0x67, 0x89, 0x79, 0xc1, 0xde, 0x60, 0x78,
	0x28, 0xe7, 0x4a, 0x8a, 0x1c, 0xc9, 0x15, 0xb8, 0x7a, 0x4d, 0x39, 0xb5, 0x3c, 0xdb, 0xef, 0xce,
	0xfa, 0x53, 0x15, 0x4d, 0xb7, 0x4b, 0x0a, 0x4c, 0x91, 0x5c, 0x42, 0x07, 0xb3, 0x4c, 0x66, 0x73,
	0x19, 0xa3, 0xc6, 0x70, 0x82, 0x9d, 0xc0, 0x6e, 0xe0, 0x7c, 0x81, 0xc5, 0x5e, 0x84, 0xdc, 0xfc,
	0xf6, 0x38, 0x0a, 0x8b, 0x60, 0x7c, 0xe4, 0x35, 0x2c, 0xd7, 0xd0, 0x4e, 0x6b, 0xc9, 0xc0, 0x0c,
	0xb6, 0x30, 0xb5, 0x35, 0xd8, 0xd4, 0xff, 0xe6, 0x99, 0x7d, 0x5b, 0xd0, 0x9f, 0x9b, 0x2b, 0xd0,
	0xed, 0x64, 0x0e, 0xbd, 0xfd, 0xf8, 0x64, 0x5c, 0x4d, 0x6e, 0xd8, 0xd3, 0x84, 0x1e, 0x17, 0x6a,
	0x3a, 0xf6, 0x8f, 0x3c, 0xc1, 0xe0, 0x17, 0x3a, 0x99, 0x68, 0xc2, 0xc6, 0xec, 0x93, 0x8b, 0xc6,
	0xda, 0x66, 0x5a, 0xe4, 0xea, 0xc3, 0xbc, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xe5, 0x67,
	0x84, 0xb0, 0x02, 0x00, 0x00,
}
