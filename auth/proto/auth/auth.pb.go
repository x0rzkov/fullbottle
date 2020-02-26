// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package fullbottle_srv_auth

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

type GenerateJwtTokenRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Expire               int64    `protobuf:"varint,2,opt,name=expire,proto3" json:"expire,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateJwtTokenRequest) Reset()         { *m = GenerateJwtTokenRequest{} }
func (m *GenerateJwtTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateJwtTokenRequest) ProtoMessage()    {}
func (*GenerateJwtTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *GenerateJwtTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateJwtTokenRequest.Unmarshal(m, b)
}
func (m *GenerateJwtTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateJwtTokenRequest.Marshal(b, m, deterministic)
}
func (m *GenerateJwtTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateJwtTokenRequest.Merge(m, src)
}
func (m *GenerateJwtTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateJwtTokenRequest.Size(m)
}
func (m *GenerateJwtTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateJwtTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateJwtTokenRequest proto.InternalMessageInfo

func (m *GenerateJwtTokenRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GenerateJwtTokenRequest) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

type GenerateJwtTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateJwtTokenResponse) Reset()         { *m = GenerateJwtTokenResponse{} }
func (m *GenerateJwtTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateJwtTokenResponse) ProtoMessage()    {}
func (*GenerateJwtTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *GenerateJwtTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateJwtTokenResponse.Unmarshal(m, b)
}
func (m *GenerateJwtTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateJwtTokenResponse.Marshal(b, m, deterministic)
}
func (m *GenerateJwtTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateJwtTokenResponse.Merge(m, src)
}
func (m *GenerateJwtTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateJwtTokenResponse.Size(m)
}
func (m *GenerateJwtTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateJwtTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateJwtTokenResponse proto.InternalMessageInfo

func (m *GenerateJwtTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ParseJwtTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseJwtTokenRequest) Reset()         { *m = ParseJwtTokenRequest{} }
func (m *ParseJwtTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ParseJwtTokenRequest) ProtoMessage()    {}
func (*ParseJwtTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *ParseJwtTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseJwtTokenRequest.Unmarshal(m, b)
}
func (m *ParseJwtTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseJwtTokenRequest.Marshal(b, m, deterministic)
}
func (m *ParseJwtTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseJwtTokenRequest.Merge(m, src)
}
func (m *ParseJwtTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ParseJwtTokenRequest.Size(m)
}
func (m *ParseJwtTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseJwtTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParseJwtTokenRequest proto.InternalMessageInfo

func (m *ParseJwtTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ParseJwtTokenResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseJwtTokenResponse) Reset()         { *m = ParseJwtTokenResponse{} }
func (m *ParseJwtTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ParseJwtTokenResponse) ProtoMessage()    {}
func (*ParseJwtTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{3}
}

func (m *ParseJwtTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseJwtTokenResponse.Unmarshal(m, b)
}
func (m *ParseJwtTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseJwtTokenResponse.Marshal(b, m, deterministic)
}
func (m *ParseJwtTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseJwtTokenResponse.Merge(m, src)
}
func (m *ParseJwtTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ParseJwtTokenResponse.Size(m)
}
func (m *ParseJwtTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseJwtTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParseJwtTokenResponse proto.InternalMessageInfo

func (m *ParseJwtTokenResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CheckFolderAccessRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FolderId             int64    `protobuf:"varint,2,opt,name=folderId,proto3" json:"folderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckFolderAccessRequest) Reset()         { *m = CheckFolderAccessRequest{} }
func (m *CheckFolderAccessRequest) String() string { return proto.CompactTextString(m) }
func (*CheckFolderAccessRequest) ProtoMessage()    {}
func (*CheckFolderAccessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{4}
}

func (m *CheckFolderAccessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckFolderAccessRequest.Unmarshal(m, b)
}
func (m *CheckFolderAccessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckFolderAccessRequest.Marshal(b, m, deterministic)
}
func (m *CheckFolderAccessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckFolderAccessRequest.Merge(m, src)
}
func (m *CheckFolderAccessRequest) XXX_Size() int {
	return xxx_messageInfo_CheckFolderAccessRequest.Size(m)
}
func (m *CheckFolderAccessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckFolderAccessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckFolderAccessRequest proto.InternalMessageInfo

func (m *CheckFolderAccessRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CheckFolderAccessRequest) GetFolderId() int64 {
	if m != nil {
		return m.FolderId
	}
	return 0
}

type CheckFolderAccessResponse struct {
	Actions              []string `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckFolderAccessResponse) Reset()         { *m = CheckFolderAccessResponse{} }
func (m *CheckFolderAccessResponse) String() string { return proto.CompactTextString(m) }
func (*CheckFolderAccessResponse) ProtoMessage()    {}
func (*CheckFolderAccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{5}
}

func (m *CheckFolderAccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckFolderAccessResponse.Unmarshal(m, b)
}
func (m *CheckFolderAccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckFolderAccessResponse.Marshal(b, m, deterministic)
}
func (m *CheckFolderAccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckFolderAccessResponse.Merge(m, src)
}
func (m *CheckFolderAccessResponse) XXX_Size() int {
	return xxx_messageInfo_CheckFolderAccessResponse.Size(m)
}
func (m *CheckFolderAccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckFolderAccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckFolderAccessResponse proto.InternalMessageInfo

func (m *CheckFolderAccessResponse) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateJwtTokenRequest)(nil), "fullbottle.srv.auth.GenerateJwtTokenRequest")
	proto.RegisterType((*GenerateJwtTokenResponse)(nil), "fullbottle.srv.auth.GenerateJwtTokenResponse")
	proto.RegisterType((*ParseJwtTokenRequest)(nil), "fullbottle.srv.auth.ParseJwtTokenRequest")
	proto.RegisterType((*ParseJwtTokenResponse)(nil), "fullbottle.srv.auth.ParseJwtTokenResponse")
	proto.RegisterType((*CheckFolderAccessRequest)(nil), "fullbottle.srv.auth.CheckFolderAccessRequest")
	proto.RegisterType((*CheckFolderAccessResponse)(nil), "fullbottle.srv.auth.CheckFolderAccessResponse")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xed, 0x4a, 0xc3, 0x40,
	0x10, 0x34, 0x2d, 0x56, 0xbb, 0x22, 0xe8, 0xd9, 0xea, 0x99, 0x5f, 0xe5, 0x7e, 0x55, 0xa9, 0x57,
	0x51, 0x7c, 0x80, 0x22, 0x28, 0xf1, 0x87, 0x48, 0xf4, 0x05, 0xd2, 0x64, 0x4b, 0x42, 0x43, 0x2e,
	0xbd, 0xbb, 0x54, 0x1f, 0xd7, 0x47, 0x91, 0x7c, 0xb4, 0x60, 0x73, 0xa1, 0xf9, 0x13, 0x98, 0xcd,
	0x30, 0x3b, 0x3b, 0x73, 0x30, 0x4c, 0xa5, 0xd0, 0x62, 0xea, 0x65, 0x3a, 0x2c, 0x3e, 0xbc, 0xc0,
	0xe4, 0x62, 0x91, 0xc5, 0xf1, 0x5c, 0x68, 0x1d, 0x23, 0x57, 0x72, 0xcd, 0xf3, 0x5f, 0xcc, 0x81,
	0xab, 0x57, 0x4c, 0x50, 0x7a, 0x1a, 0xdf, 0xbe, 0xf5, 0x97, 0x58, 0x62, 0xe2, 0xe2, 0x2a, 0x43,
	0xa5, 0xc9, 0x25, 0xf4, 0x32, 0x85, 0xd2, 0x09, 0xa8, 0x35, 0xb2, 0xc6, 0x5d, 0xb7, 0x42, 0xf9,
	0x1c, 0x7f, 0xd2, 0x48, 0x22, 0xed, 0x94, 0xf3, 0x12, 0xb1, 0x7b, 0xa0, 0x75, 0x29, 0x95, 0x8a,
	0x44, 0x21, 0x19, 0xc0, 0xa1, 0xce, 0x07, 0x85, 0x54, 0xdf, 0x2d, 0x01, 0x9b, 0xc0, 0xe0, 0xc3,
	0x93, 0xaa, 0xb6, 0xd9, 0xcc, 0x9e, 0xc2, 0x70, 0x87, 0x5d, 0x89, 0x37, 0x18, 0x65, 0xef, 0x40,
	0x9f, 0x43, 0xf4, 0x97, 0x2f, 0x22, 0x0e, 0x50, 0xce, 0x7c, 0x1f, 0x95, 0xda, 0x77, 0x9c, 0x0d,
	0xc7, 0x8b, 0x82, 0xee, 0x04, 0xd5, 0x79, 0x5b, 0xcc, 0x9e, 0xe0, 0xda, 0xa0, 0x57, 0x99, 0xa0,
	0x70, 0xe4, 0xf9, 0x3a, 0x12, 0x89, 0xa2, 0xd6, 0xa8, 0x3b, 0xee, 0xbb, 0x1b, 0xf8, 0xf0, 0xdb,
	0x81, 0x93, 0x59, 0xa6, 0xc3, 0x4f, 0x94, 0xeb, 0xc8, 0x47, 0xb2, 0x82, 0xb3, 0xdd, 0x9c, 0xc8,
	0x84, 0x1b, 0xca, 0xe1, 0x0d, 0xcd, 0xd8, 0x77, 0x2d, 0xd9, 0xa5, 0x35, 0x76, 0x40, 0x42, 0x38,
	0xfd, 0x17, 0x1d, 0xb9, 0x31, 0x2a, 0x98, 0xca, 0xb0, 0x6f, 0xdb, 0x50, 0xb7, 0x9b, 0x34, 0x9c,
	0xd7, 0x32, 0x22, 0x66, 0xbf, 0x4d, 0xdd, 0xd8, 0xbc, 0x2d, 0x7d, 0xb3, 0x75, 0xde, 0x2b, 0x5e,
	0xf8, 0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0x0c, 0x6e, 0xd0, 0xfa, 0x02, 0x00, 0x00,
}
