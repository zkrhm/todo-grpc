// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spec/todo.proto

package todo

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

type Todo struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Completed            bool     `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_c6a4641cdd85bf86, []int{0}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (dst *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(dst, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Todo) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Todo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type TodoToSave struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoToSave) Reset()         { *m = TodoToSave{} }
func (m *TodoToSave) String() string { return proto.CompactTextString(m) }
func (*TodoToSave) ProtoMessage()    {}
func (*TodoToSave) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_c6a4641cdd85bf86, []int{1}
}
func (m *TodoToSave) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoToSave.Unmarshal(m, b)
}
func (m *TodoToSave) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoToSave.Marshal(b, m, deterministic)
}
func (dst *TodoToSave) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoToSave.Merge(dst, src)
}
func (m *TodoToSave) XXX_Size() int {
	return xxx_messageInfo_TodoToSave.Size(m)
}
func (m *TodoToSave) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoToSave.DiscardUnknown(m)
}

var xxx_messageInfo_TodoToSave proto.InternalMessageInfo

func (m *TodoToSave) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TodoToSave) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type Ack struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_c6a4641cdd85bf86, []int{2}
}
func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (dst *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(dst, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Ack) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SearchReq struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Completed            bool     `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReq) Reset()         { *m = SearchReq{} }
func (m *SearchReq) String() string { return proto.CompactTextString(m) }
func (*SearchReq) ProtoMessage()    {}
func (*SearchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_c6a4641cdd85bf86, []int{3}
}
func (m *SearchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReq.Unmarshal(m, b)
}
func (m *SearchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReq.Marshal(b, m, deterministic)
}
func (dst *SearchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReq.Merge(dst, src)
}
func (m *SearchReq) XXX_Size() int {
	return xxx_messageInfo_SearchReq.Size(m)
}
func (m *SearchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReq.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReq proto.InternalMessageInfo

func (m *SearchReq) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *SearchReq) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type SearchRes struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRes) Reset()         { *m = SearchRes{} }
func (m *SearchRes) String() string { return proto.CompactTextString(m) }
func (*SearchRes) ProtoMessage()    {}
func (*SearchRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_c6a4641cdd85bf86, []int{4}
}
func (m *SearchRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRes.Unmarshal(m, b)
}
func (m *SearchRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRes.Marshal(b, m, deterministic)
}
func (dst *SearchRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRes.Merge(dst, src)
}
func (m *SearchRes) XXX_Size() int {
	return xxx_messageInfo_SearchRes.Size(m)
}
func (m *SearchRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRes.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRes proto.InternalMessageInfo

func (m *SearchRes) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

func init() {
	proto.RegisterType((*Todo)(nil), "Todo")
	proto.RegisterType((*TodoToSave)(nil), "TodoToSave")
	proto.RegisterType((*Ack)(nil), "Ack")
	proto.RegisterType((*SearchReq)(nil), "SearchReq")
	proto.RegisterType((*SearchRes)(nil), "SearchRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error)
	Save(ctx context.Context, in *TodoToSave, opts ...grpc.CallOption) (*Ack, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
} 

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error) {
	out := new(SearchRes)
	err := c.cc.Invoke(ctx, "/TodoService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Save(ctx context.Context, in *TodoToSave, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/TodoService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	Search(context.Context, *SearchReq) (*SearchRes, error)
	Save(context.Context, *TodoToSave) (*Ack, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Search(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoToSave)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Save(ctx, req.(*TodoToSave))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _TodoService_Search_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _TodoService_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spec/todo.proto",
}

func init() { proto.RegisterFile("spec/todo.proto", fileDescriptor_todo_c6a4641cdd85bf86) }

var fileDescriptor_todo_c6a4641cdd85bf86 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x26, 0x29, 0xe4, 0x32, 0x20, 0x9d, 0x18, 0x4c, 0x61, 0x88, 0x32, 0x65, 0x32,
	0x52, 0xbb, 0x22, 0xa1, 0xae, 0x8c, 0x4e, 0x5e, 0x20, 0x38, 0x27, 0x40, 0xa5, 0xb9, 0x10, 0x5b,
	0xe5, 0xf5, 0x91, 0x1d, 0x5a, 0x77, 0xca, 0x76, 0xff, 0x49, 0xdf, 0x27, 0xdf, 0x6f, 0xb8, 0xb3,
	0x23, 0x99, 0x67, 0xc7, 0x3d, 0xab, 0x71, 0x62, 0xc7, 0x55, 0x0b, 0x69, 0xcb, 0x3d, 0xa3, 0x84,
	0x1b, 0xc3, 0x83, 0xa3, 0xc1, 0x49, 0x51, 0x8a, 0x3a, 0xd7, 0xe7, 0x88, 0x4f, 0x90, 0x1b, 0x3e,
	0x8e, 0xdf, 0xe4, 0xa8, 0x97, 0xab, 0x52, 0xd4, 0xb7, 0x3a, 0x2e, 0xf0, 0x1e, 0x32, 0xfe, 0x1d,
	0x68, 0x92, 0x49, 0xa0, 0xe6, 0x50, 0xbd, 0x00, 0x78, 0x6b, 0xcb, 0x4d, 0x77, 0xa2, 0x05, 0xf7,
	0x85, 0x5e, 0x5d, 0xd3, 0x3b, 0x48, 0xf6, 0xe6, 0x80, 0x08, 0xa9, 0xe1, 0x9e, 0x02, 0x93, 0xe9,
	0x30, 0x7b, 0xd5, 0x91, 0xac, 0xed, 0x3e, 0xe8, 0x1f, 0x39, 0xc7, 0xea, 0x15, 0xf2, 0x86, 0xba,
	0xc9, 0x7c, 0x6a, 0xfa, 0x89, 0x5e, 0x71, 0xe5, 0x5d, 0xbe, 0xa4, 0xaa, 0xa3, 0xc0, 0xe2, 0x23,
	0x64, 0xbe, 0x24, 0x2b, 0x45, 0x99, 0xd4, 0xc5, 0x36, 0x53, 0xfe, 0x1c, 0x3d, 0xef, 0xb6, 0x6f,
	0x50, 0xf8, 0xd8, 0xd0, 0x74, 0xfa, 0x32, 0x84, 0x25, 0xac, 0x67, 0x10, 0x41, 0x5d, 0x9e, 0xb0,
	0x89, 0xb3, 0xc5, 0x07, 0x48, 0x43, 0x11, 0x85, 0x8a, 0xad, 0x6c, 0x52, 0xb5, 0x37, 0x87, 0xf7,
	0x75, 0xf8, 0x86, 0xdd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0xed, 0xdf, 0x0c, 0x99, 0x01,
	0x00, 0x00,
}
