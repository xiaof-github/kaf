// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/topic/topic.proto

package topic

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

type CreateTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTopicRequest) Reset()         { *m = CreateTopicRequest{} }
func (m *CreateTopicRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTopicRequest) ProtoMessage()    {}
func (*CreateTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c55aa52b213962a7, []int{0}
}

func (m *CreateTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicRequest.Unmarshal(m, b)
}
func (m *CreateTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicRequest.Marshal(b, m, deterministic)
}
func (m *CreateTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicRequest.Merge(m, src)
}
func (m *CreateTopicRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTopicRequest.Size(m)
}
func (m *CreateTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicRequest proto.InternalMessageInfo

type CreateTopicResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTopicResponse) Reset()         { *m = CreateTopicResponse{} }
func (m *CreateTopicResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTopicResponse) ProtoMessage()    {}
func (*CreateTopicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c55aa52b213962a7, []int{1}
}

func (m *CreateTopicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicResponse.Unmarshal(m, b)
}
func (m *CreateTopicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicResponse.Marshal(b, m, deterministic)
}
func (m *CreateTopicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicResponse.Merge(m, src)
}
func (m *CreateTopicResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTopicResponse.Size(m)
}
func (m *CreateTopicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicResponse proto.InternalMessageInfo

type ListTopicsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopicsRequest) Reset()         { *m = ListTopicsRequest{} }
func (m *ListTopicsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTopicsRequest) ProtoMessage()    {}
func (*ListTopicsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c55aa52b213962a7, []int{2}
}

func (m *ListTopicsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsRequest.Unmarshal(m, b)
}
func (m *ListTopicsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsRequest.Marshal(b, m, deterministic)
}
func (m *ListTopicsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsRequest.Merge(m, src)
}
func (m *ListTopicsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTopicsRequest.Size(m)
}
func (m *ListTopicsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsRequest proto.InternalMessageInfo

type ListTopicsResponse struct {
	Topics               []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopicsResponse) Reset()         { *m = ListTopicsResponse{} }
func (m *ListTopicsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTopicsResponse) ProtoMessage()    {}
func (*ListTopicsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c55aa52b213962a7, []int{3}
}

func (m *ListTopicsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsResponse.Unmarshal(m, b)
}
func (m *ListTopicsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsResponse.Marshal(b, m, deterministic)
}
func (m *ListTopicsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsResponse.Merge(m, src)
}
func (m *ListTopicsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTopicsResponse.Size(m)
}
func (m *ListTopicsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsResponse proto.InternalMessageInfo

func (m *ListTopicsResponse) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type GetTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicRequest) Reset()         { *m = GetTopicRequest{} }
func (m *GetTopicRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicRequest) ProtoMessage()    {}
func (*GetTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c55aa52b213962a7, []int{4}
}

func (m *GetTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicRequest.Unmarshal(m, b)
}
func (m *GetTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicRequest.Marshal(b, m, deterministic)
}
func (m *GetTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicRequest.Merge(m, src)
}
func (m *GetTopicRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicRequest.Size(m)
}
func (m *GetTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicRequest proto.InternalMessageInfo

type DeleteTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTopicRequest) Reset()         { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()    {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c55aa52b213962a7, []int{5}
}

func (m *DeleteTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTopicRequest.Unmarshal(m, b)
}
func (m *DeleteTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTopicRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTopicRequest.Merge(m, src)
}
func (m *DeleteTopicRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTopicRequest.Size(m)
}
func (m *DeleteTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTopicRequest proto.InternalMessageInfo

type UpdateTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTopicRequest) Reset()         { *m = UpdateTopicRequest{} }
func (m *UpdateTopicRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTopicRequest) ProtoMessage()    {}
func (*UpdateTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c55aa52b213962a7, []int{6}
}

func (m *UpdateTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTopicRequest.Unmarshal(m, b)
}
func (m *UpdateTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTopicRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTopicRequest.Merge(m, src)
}
func (m *UpdateTopicRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTopicRequest.Size(m)
}
func (m *UpdateTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTopicRequest proto.InternalMessageInfo

type Topic struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_c55aa52b213962a7, []int{7}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateTopicRequest)(nil), "kaf.topic.CreateTopicRequest")
	proto.RegisterType((*CreateTopicResponse)(nil), "kaf.topic.CreateTopicResponse")
	proto.RegisterType((*ListTopicsRequest)(nil), "kaf.topic.ListTopicsRequest")
	proto.RegisterType((*ListTopicsResponse)(nil), "kaf.topic.ListTopicsResponse")
	proto.RegisterType((*GetTopicRequest)(nil), "kaf.topic.GetTopicRequest")
	proto.RegisterType((*DeleteTopicRequest)(nil), "kaf.topic.DeleteTopicRequest")
	proto.RegisterType((*UpdateTopicRequest)(nil), "kaf.topic.UpdateTopicRequest")
	proto.RegisterType((*Topic)(nil), "kaf.topic.Topic")
}

func init() { proto.RegisterFile("pkg/topic/topic.proto", fileDescriptor_c55aa52b213962a7) }

var fileDescriptor_c55aa52b213962a7 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x05, 0xa8, 0x2d, 0x5c, 0x90, 0xa0, 0xd7, 0x14, 0x21, 0x43, 0x25, 0x94, 0xa9, 0x93,
	0x23, 0x95, 0x85, 0x29, 0x03, 0x14, 0x21, 0x24, 0x26, 0xfe, 0x2c, 0x6c, 0x6d, 0xb9, 0x46, 0x51,
	0x0b, 0x36, 0xb1, 0x8b, 0xc4, 0xa3, 0xf2, 0x36, 0xa8, 0x76, 0x03, 0x76, 0x9c, 0x2e, 0x51, 0x7c,
	0xdf, 0x9d, 0xef, 0x77, 0x9f, 0x0f, 0xfa, 0x72, 0x91, 0xa7, 0x5a, 0xc8, 0x62, 0x66, 0xbf, 0x5c,
	0x96, 0x42, 0x0b, 0x3c, 0x58, 0x4c, 0xe6, 0xdc, 0x04, 0xd8, 0x59, 0x2e, 0x44, 0xbe, 0xa4, 0xd4,
	0x08, 0xd3, 0xd5, 0x3c, 0xa5, 0x77, 0xa9, 0xbf, 0x6d, 0x5e, 0x12, 0x03, 0xde, 0x94, 0x34, 0xd1,
	0xf4, 0xbc, 0xce, 0x7d, 0xa4, 0xcf, 0x15, 0x29, 0x9d, 0xf4, 0xa1, 0xe7, 0x45, 0x95, 0x14, 0x1f,
	0x8a, 0x92, 0x1e, 0x74, 0x1f, 0x0a, 0xa5, 0x4d, 0x50, 0x55, 0xb9, 0x19, 0xa0, 0x1b, 0xb4, 0xa9,
	0x38, 0x84, 0xb6, 0xe9, 0xae, 0x4e, 0x77, 0x2e, 0xf6, 0x86, 0xd1, 0xe8, 0x98, 0xff, 0x01, 0x71,
	0x7b, 0xe9, 0x46, 0x4f, 0xba, 0x70, 0x74, 0x47, 0xda, 0x6b, 0x1f, 0x03, 0x8e, 0x69, 0x49, 0x35,
	0xa8, 0x18, 0xf0, 0x45, 0xbe, 0xd5, 0x51, 0x3b, 0xd0, 0x32, 0xe7, 0xd1, 0xcf, 0x2e, 0x1c, 0x9a,
	0xbf, 0x27, 0x2a, 0xbf, 0x8a, 0x19, 0x61, 0x06, 0x91, 0x33, 0x04, 0x0e, 0x1c, 0x82, 0x70, 0x64,
	0x16, 0x00, 0xe2, 0x15, 0xec, 0x57, 0x60, 0xc8, 0x1c, 0xb5, 0x46, 0xdb, 0x50, 0x99, 0x41, 0xe4,
	0x90, 0x7a, 0x9d, 0xc3, 0x09, 0x1a, 0xea, 0xef, 0x01, 0xfe, 0x2d, 0xc5, 0x73, 0x47, 0x0f, 0xec,
	0x67, 0x83, 0x2d, 0xea, 0xe6, 0x1d, 0xc6, 0x10, 0x39, 0x56, 0x7a, 0x28, 0xa1, 0xc5, 0xec, 0x84,
	0xdb, 0x5d, 0xe1, 0xd5, 0xae, 0xf0, 0xdb, 0xf5, 0xae, 0x5c, 0x77, 0x5e, 0x5b, 0xa6, 0x66, 0xda,
	0x36, 0xc2, 0xe5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xce, 0x98, 0x4d, 0x76, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicServiceClient interface {
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type topicServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopicServiceClient(cc *grpc.ClientConn) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/kaf.topic.TopicService/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/kaf.topic.TopicService/GetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/kaf.topic.TopicService/UpdateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, "/kaf.topic.TopicService/ListTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kaf.topic.TopicService/DeleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
type TopicServiceServer interface {
	CreateTopic(context.Context, *CreateTopicRequest) (*Topic, error)
	GetTopic(context.Context, *GetTopicRequest) (*Topic, error)
	UpdateTopic(context.Context, *UpdateTopicRequest) (*Topic, error)
	ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)
	DeleteTopic(context.Context, *DeleteTopicRequest) (*empty.Empty, error)
}

// UnimplementedTopicServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (*UnimplementedTopicServiceServer) CreateTopic(ctx context.Context, req *CreateTopicRequest) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (*UnimplementedTopicServiceServer) GetTopic(ctx context.Context, req *GetTopicRequest) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (*UnimplementedTopicServiceServer) UpdateTopic(ctx context.Context, req *UpdateTopicRequest) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}
func (*UnimplementedTopicServiceServer) ListTopics(ctx context.Context, req *ListTopicsRequest) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (*UnimplementedTopicServiceServer) DeleteTopic(ctx context.Context, req *DeleteTopicRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}

func RegisterTopicServiceServer(s *grpc.Server, srv TopicServiceServer) {
	s.RegisterService(&_TopicService_serviceDesc, srv)
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.topic.TopicService/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.topic.TopicService/GetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.topic.TopicService/UpdateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).UpdateTopic(ctx, req.(*UpdateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.topic.TopicService/ListTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).ListTopics(ctx, req.(*ListTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.topic.TopicService/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kaf.topic.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _TopicService_GetTopic_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _TopicService_UpdateTopic_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _TopicService_ListTopics_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _TopicService_DeleteTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/topic/topic.proto",
}