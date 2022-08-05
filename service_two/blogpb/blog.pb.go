// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog.proto

package blogpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Blog struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Blog) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRequest) Reset()         { *m = CreateBlogRequest{} }
func (m *CreateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRequest) ProtoMessage()    {}
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{1}
}

func (m *CreateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRequest.Unmarshal(m, b)
}
func (m *CreateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRequest.Marshal(b, m, deterministic)
}
func (m *CreateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRequest.Merge(m, src)
}
func (m *CreateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRequest.Size(m)
}
func (m *CreateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRequest proto.InternalMessageInfo

func (m *CreateBlogRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateBlogRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogResponse) Reset()         { *m = CreateBlogResponse{} }
func (m *CreateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBlogResponse) ProtoMessage()    {}
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{2}
}

func (m *CreateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogResponse.Unmarshal(m, b)
}
func (m *CreateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogResponse.Marshal(b, m, deterministic)
}
func (m *CreateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogResponse.Merge(m, src)
}
func (m *CreateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBlogResponse.Size(m)
}
func (m *CreateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogResponse proto.InternalMessageInfo

func (m *CreateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type DeleteBlogRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogRequest) Reset()         { *m = DeleteBlogRequest{} }
func (m *DeleteBlogRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogRequest) ProtoMessage()    {}
func (*DeleteBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{3}
}

func (m *DeleteBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogRequest.Unmarshal(m, b)
}
func (m *DeleteBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogRequest.Merge(m, src)
}
func (m *DeleteBlogRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogRequest.Size(m)
}
func (m *DeleteBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogRequest proto.InternalMessageInfo

func (m *DeleteBlogRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogResponse) Reset()         { *m = DeleteBlogResponse{} }
func (m *DeleteBlogResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogResponse) ProtoMessage()    {}
func (*DeleteBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{4}
}

func (m *DeleteBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogResponse.Unmarshal(m, b)
}
func (m *DeleteBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogResponse.Marshal(b, m, deterministic)
}
func (m *DeleteBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogResponse.Merge(m, src)
}
func (m *DeleteBlogResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogResponse.Size(m)
}
func (m *DeleteBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogResponse proto.InternalMessageInfo

func (m *DeleteBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type GetBlogRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlogRequest) Reset()         { *m = GetBlogRequest{} }
func (m *GetBlogRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlogRequest) ProtoMessage()    {}
func (*GetBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{5}
}

func (m *GetBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlogRequest.Unmarshal(m, b)
}
func (m *GetBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlogRequest.Marshal(b, m, deterministic)
}
func (m *GetBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlogRequest.Merge(m, src)
}
func (m *GetBlogRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlogRequest.Size(m)
}
func (m *GetBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlogRequest proto.InternalMessageInfo

func (m *GetBlogRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlogResponse) Reset()         { *m = GetBlogResponse{} }
func (m *GetBlogResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlogResponse) ProtoMessage()    {}
func (*GetBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{6}
}

func (m *GetBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlogResponse.Unmarshal(m, b)
}
func (m *GetBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlogResponse.Marshal(b, m, deterministic)
}
func (m *GetBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlogResponse.Merge(m, src)
}
func (m *GetBlogResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlogResponse.Size(m)
}
func (m *GetBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlogResponse proto.InternalMessageInfo

func (m *GetBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type GetBlogsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlogsRequest) Reset()         { *m = GetBlogsRequest{} }
func (m *GetBlogsRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlogsRequest) ProtoMessage()    {}
func (*GetBlogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{7}
}

func (m *GetBlogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlogsRequest.Unmarshal(m, b)
}
func (m *GetBlogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlogsRequest.Marshal(b, m, deterministic)
}
func (m *GetBlogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlogsRequest.Merge(m, src)
}
func (m *GetBlogsRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlogsRequest.Size(m)
}
func (m *GetBlogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlogsRequest proto.InternalMessageInfo

type GetBlogsResponse struct {
	Blog                 []*Blog  `protobuf:"bytes,1,rep,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlogsResponse) Reset()         { *m = GetBlogsResponse{} }
func (m *GetBlogsResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlogsResponse) ProtoMessage()    {}
func (*GetBlogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{8}
}

func (m *GetBlogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlogsResponse.Unmarshal(m, b)
}
func (m *GetBlogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlogsResponse.Marshal(b, m, deterministic)
}
func (m *GetBlogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlogsResponse.Merge(m, src)
}
func (m *GetBlogsResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlogsResponse.Size(m)
}
func (m *GetBlogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlogsResponse proto.InternalMessageInfo

func (m *GetBlogsResponse) GetBlog() []*Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blogpb.Blog")
	proto.RegisterType((*CreateBlogRequest)(nil), "blogpb.CreateBlogRequest")
	proto.RegisterType((*CreateBlogResponse)(nil), "blogpb.CreateBlogResponse")
	proto.RegisterType((*DeleteBlogRequest)(nil), "blogpb.DeleteBlogRequest")
	proto.RegisterType((*DeleteBlogResponse)(nil), "blogpb.DeleteBlogResponse")
	proto.RegisterType((*GetBlogRequest)(nil), "blogpb.GetBlogRequest")
	proto.RegisterType((*GetBlogResponse)(nil), "blogpb.GetBlogResponse")
	proto.RegisterType((*GetBlogsRequest)(nil), "blogpb.GetBlogsRequest")
	proto.RegisterType((*GetBlogsResponse)(nil), "blogpb.GetBlogsResponse")
}

func init() { proto.RegisterFile("blog.proto", fileDescriptor_6745b25902462fb1) }

var fileDescriptor_6745b25902462fb1 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x4f, 0x83, 0x50,
	0x14, 0xc5, 0x03, 0xad, 0xa5, 0x3d, 0x35, 0x55, 0xee, 0xa0, 0xc8, 0x44, 0x70, 0xe9, 0x84, 0x49,
	0x6b, 0x1c, 0x4c, 0x5c, 0x8a, 0x89, 0x31, 0x71, 0xc2, 0xcd, 0xad, 0xb4, 0x37, 0x0d, 0x09, 0xf2,
	0x10, 0x5e, 0xfd, 0x20, 0x7e, 0x62, 0xc3, 0xff, 0x52, 0xd0, 0x74, 0xe3, 0xbe, 0x73, 0xce, 0x8f,
	0x77, 0x0f, 0x00, 0x7e, 0x28, 0x76, 0x4e, 0x9c, 0x08, 0x29, 0x68, 0x94, 0x3d, 0xc7, 0xbe, 0xfd,
	0x86, 0xe1, 0x2a, 0x14, 0x3b, 0x9a, 0x41, 0x0d, 0xb6, 0x86, 0x62, 0x29, 0xf3, 0x33, 0x4f, 0x0d,
	0xb6, 0x64, 0x62, 0xbc, 0x4f, 0x39, 0x89, 0xd6, 0x9f, 0x6c, 0xa8, 0x96, 0x32, 0x9f, 0x78, 0xf5,
	0x4c, 0x06, 0xb4, 0x8d, 0x88, 0x24, 0x47, 0xd2, 0x18, 0xe4, 0x52, 0x35, 0xda, 0xaf, 0xd0, 0xdd,
	0x84, 0xd7, 0x92, 0x33, 0xa6, 0xc7, 0x5f, 0x7b, 0x4e, 0x65, 0x0b, 0xa5, 0xfc, 0x8d, 0x52, 0xdb,
	0xa8, 0x07, 0xd0, 0x21, 0x2a, 0x8d, 0x45, 0x94, 0x32, 0x59, 0x18, 0x66, 0x17, 0xcf, 0x39, 0xd3,
	0xc5, 0xb9, 0x53, 0x6c, 0xe1, 0xe4, 0x9e, 0x5c, 0xb1, 0x6f, 0xa1, 0x3f, 0x73, 0xc8, 0xed, 0x2b,
	0x1c, 0x6d, 0x97, 0xc1, 0x0f, 0x4d, 0x27, 0xc3, 0x2d, 0xcc, 0x5e, 0x58, 0xfe, 0x47, 0x5e, 0xe2,
	0xa2, 0x76, 0x9c, 0x8c, 0xd5, 0xeb, 0x50, 0x5a, 0x72, 0xed, 0x7b, 0x5c, 0x36, 0x47, 0x1d, 0xd0,
	0xa0, 0x1f, 0xb4, 0xf8, 0x51, 0x31, 0xcd, 0xc6, 0x77, 0x4e, 0xbe, 0x83, 0x0d, 0x93, 0x0b, 0x34,
	0x25, 0xd2, 0x4d, 0x95, 0xe8, 0x7c, 0x23, 0xd3, 0xec, 0x93, 0xca, 0xd7, 0xba, 0x40, 0x53, 0x56,
	0x03, 0xe9, 0xb4, 0xdc, 0x40, 0x7a, 0xba, 0x7d, 0x84, 0x56, 0xee, 0x43, 0x57, 0x95, 0xad, 0x5d,
	0xa5, 0x79, 0xdd, 0x39, 0x2f, 0xb3, 0x4f, 0x18, 0x57, 0x5d, 0xd0, 0xb1, 0xa9, 0x2a, 0xcc, 0x34,
	0xba, 0x42, 0x11, 0x5f, 0x4d, 0x3e, 0xb4, 0xbb, 0x42, 0xf3, 0x47, 0xf9, 0xcf, 0xbf, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x74, 0xa0, 0x16, 0x98, 0x0a, 0x03, 0x00, 0x00,
}
