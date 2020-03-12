// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/share/share.proto

package fullbottle_srv_share

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	bottle "github.com/vegchic/fullbottle/bottle/proto/bottle"
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

type ShareStatusRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ViewerId             int64    `protobuf:"varint,2,opt,name=viewer_id,json=viewerId,proto3" json:"viewer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareStatusRequest) Reset()         { *m = ShareStatusRequest{} }
func (m *ShareStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ShareStatusRequest) ProtoMessage()    {}
func (*ShareStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{0}
}

func (m *ShareStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareStatusRequest.Unmarshal(m, b)
}
func (m *ShareStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareStatusRequest.Marshal(b, m, deterministic)
}
func (m *ShareStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareStatusRequest.Merge(m, src)
}
func (m *ShareStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ShareStatusRequest.Size(m)
}
func (m *ShareStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShareStatusRequest proto.InternalMessageInfo

func (m *ShareStatusRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ShareStatusRequest) GetViewerId() int64 {
	if m != nil {
		return m.ViewerId
	}
	return 0
}

type ShareStatusResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	OwnerId              int64    `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	IsPublic             bool     `protobuf:"varint,3,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareStatusResponse) Reset()         { *m = ShareStatusResponse{} }
func (m *ShareStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ShareStatusResponse) ProtoMessage()    {}
func (*ShareStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{1}
}

func (m *ShareStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareStatusResponse.Unmarshal(m, b)
}
func (m *ShareStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareStatusResponse.Marshal(b, m, deterministic)
}
func (m *ShareStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareStatusResponse.Merge(m, src)
}
func (m *ShareStatusResponse) XXX_Size() int {
	return xxx_messageInfo_ShareStatusResponse.Size(m)
}
func (m *ShareStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShareStatusResponse proto.InternalMessageInfo

func (m *ShareStatusResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ShareStatusResponse) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *ShareStatusResponse) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

type AccessShareRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	ViewerId             int64    `protobuf:"varint,3,opt,name=viewer_id,json=viewerId,proto3" json:"viewer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessShareRequest) Reset()         { *m = AccessShareRequest{} }
func (m *AccessShareRequest) String() string { return proto.CompactTextString(m) }
func (*AccessShareRequest) ProtoMessage()    {}
func (*AccessShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{2}
}

func (m *AccessShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessShareRequest.Unmarshal(m, b)
}
func (m *AccessShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessShareRequest.Marshal(b, m, deterministic)
}
func (m *AccessShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessShareRequest.Merge(m, src)
}
func (m *AccessShareRequest) XXX_Size() int {
	return xxx_messageInfo_AccessShareRequest.Size(m)
}
func (m *AccessShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessShareRequest proto.InternalMessageInfo

func (m *AccessShareRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccessShareRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *AccessShareRequest) GetViewerId() int64 {
	if m != nil {
		return m.ViewerId
	}
	return 0
}

type AccessShareResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Exp                  int64    `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessShareResponse) Reset()         { *m = AccessShareResponse{} }
func (m *AccessShareResponse) String() string { return proto.CompactTextString(m) }
func (*AccessShareResponse) ProtoMessage()    {}
func (*AccessShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{3}
}

func (m *AccessShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessShareResponse.Unmarshal(m, b)
}
func (m *AccessShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessShareResponse.Marshal(b, m, deterministic)
}
func (m *AccessShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessShareResponse.Merge(m, src)
}
func (m *AccessShareResponse) XXX_Size() int {
	return xxx_messageInfo_AccessShareResponse.Size(m)
}
func (m *AccessShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccessShareResponse proto.InternalMessageInfo

func (m *AccessShareResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AccessShareResponse) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

type GetShareInfoRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ViewerId             int64    `protobuf:"varint,3,opt,name=viewer_id,json=viewerId,proto3" json:"viewer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShareInfoRequest) Reset()         { *m = GetShareInfoRequest{} }
func (m *GetShareInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetShareInfoRequest) ProtoMessage()    {}
func (*GetShareInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{4}
}

func (m *GetShareInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShareInfoRequest.Unmarshal(m, b)
}
func (m *GetShareInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShareInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetShareInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShareInfoRequest.Merge(m, src)
}
func (m *GetShareInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetShareInfoRequest.Size(m)
}
func (m *GetShareInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShareInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShareInfoRequest proto.InternalMessageInfo

func (m *GetShareInfoRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetShareInfoRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetShareInfoRequest) GetViewerId() int64 {
	if m != nil {
		return m.ViewerId
	}
	return 0
}

type GetShareInfoResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SharerId             int64    `protobuf:"varint,2,opt,name=sharer_id,json=sharerId,proto3" json:"sharer_id,omitempty"`
	Exp                  int64    `protobuf:"varint,3,opt,name=exp,proto3" json:"exp,omitempty"`
	DownloadTimes        int64    `protobuf:"varint,4,opt,name=download_times,json=downloadTimes,proto3" json:"download_times,omitempty"`
	ViewTimes            int64    `protobuf:"varint,5,opt,name=view_times,json=viewTimes,proto3" json:"view_times,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShareInfoResponse) Reset()         { *m = GetShareInfoResponse{} }
func (m *GetShareInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetShareInfoResponse) ProtoMessage()    {}
func (*GetShareInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{5}
}

func (m *GetShareInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShareInfoResponse.Unmarshal(m, b)
}
func (m *GetShareInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShareInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetShareInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShareInfoResponse.Merge(m, src)
}
func (m *GetShareInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetShareInfoResponse.Size(m)
}
func (m *GetShareInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShareInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShareInfoResponse proto.InternalMessageInfo

func (m *GetShareInfoResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetShareInfoResponse) GetSharerId() int64 {
	if m != nil {
		return m.SharerId
	}
	return 0
}

func (m *GetShareInfoResponse) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *GetShareInfoResponse) GetDownloadTimes() int64 {
	if m != nil {
		return m.DownloadTimes
	}
	return 0
}

func (m *GetShareInfoResponse) GetViewTimes() int64 {
	if m != nil {
		return m.ViewTimes
	}
	return 0
}

type GetShareFolderRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	ViewerId             int64    `protobuf:"varint,4,opt,name=viewer_id,json=viewerId,proto3" json:"viewer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShareFolderRequest) Reset()         { *m = GetShareFolderRequest{} }
func (m *GetShareFolderRequest) String() string { return proto.CompactTextString(m) }
func (*GetShareFolderRequest) ProtoMessage()    {}
func (*GetShareFolderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{6}
}

func (m *GetShareFolderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShareFolderRequest.Unmarshal(m, b)
}
func (m *GetShareFolderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShareFolderRequest.Marshal(b, m, deterministic)
}
func (m *GetShareFolderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShareFolderRequest.Merge(m, src)
}
func (m *GetShareFolderRequest) XXX_Size() int {
	return xxx_messageInfo_GetShareFolderRequest.Size(m)
}
func (m *GetShareFolderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShareFolderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShareFolderRequest proto.InternalMessageInfo

func (m *GetShareFolderRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetShareFolderRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetShareFolderRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GetShareFolderRequest) GetViewerId() int64 {
	if m != nil {
		return m.ViewerId
	}
	return 0
}

type GetShareFolderResponse struct {
	Folder               *bottle.FolderInfo `protobuf:"bytes,1,opt,name=folder,proto3" json:"folder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetShareFolderResponse) Reset()         { *m = GetShareFolderResponse{} }
func (m *GetShareFolderResponse) String() string { return proto.CompactTextString(m) }
func (*GetShareFolderResponse) ProtoMessage()    {}
func (*GetShareFolderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{7}
}

func (m *GetShareFolderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShareFolderResponse.Unmarshal(m, b)
}
func (m *GetShareFolderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShareFolderResponse.Marshal(b, m, deterministic)
}
func (m *GetShareFolderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShareFolderResponse.Merge(m, src)
}
func (m *GetShareFolderResponse) XXX_Size() int {
	return xxx_messageInfo_GetShareFolderResponse.Size(m)
}
func (m *GetShareFolderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShareFolderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShareFolderResponse proto.InternalMessageInfo

func (m *GetShareFolderResponse) GetFolder() *bottle.FolderInfo {
	if m != nil {
		return m.Folder
	}
	return nil
}

type GetShareDownloadUrlRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	FileId               int64    `protobuf:"varint,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Path                 string   `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	ViewerId             int64    `protobuf:"varint,5,opt,name=viewer_id,json=viewerId,proto3" json:"viewer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShareDownloadUrlRequest) Reset()         { *m = GetShareDownloadUrlRequest{} }
func (m *GetShareDownloadUrlRequest) String() string { return proto.CompactTextString(m) }
func (*GetShareDownloadUrlRequest) ProtoMessage()    {}
func (*GetShareDownloadUrlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{8}
}

func (m *GetShareDownloadUrlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShareDownloadUrlRequest.Unmarshal(m, b)
}
func (m *GetShareDownloadUrlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShareDownloadUrlRequest.Marshal(b, m, deterministic)
}
func (m *GetShareDownloadUrlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShareDownloadUrlRequest.Merge(m, src)
}
func (m *GetShareDownloadUrlRequest) XXX_Size() int {
	return xxx_messageInfo_GetShareDownloadUrlRequest.Size(m)
}
func (m *GetShareDownloadUrlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShareDownloadUrlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShareDownloadUrlRequest proto.InternalMessageInfo

func (m *GetShareDownloadUrlRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetShareDownloadUrlRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetShareDownloadUrlRequest) GetFileId() int64 {
	if m != nil {
		return m.FileId
	}
	return 0
}

func (m *GetShareDownloadUrlRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GetShareDownloadUrlRequest) GetViewerId() int64 {
	if m != nil {
		return m.ViewerId
	}
	return 0
}

type GetShareDownloadUrlResponse struct {
	DownloadToken        string   `protobuf:"bytes,1,opt,name=download_token,json=downloadToken,proto3" json:"download_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShareDownloadUrlResponse) Reset()         { *m = GetShareDownloadUrlResponse{} }
func (m *GetShareDownloadUrlResponse) String() string { return proto.CompactTextString(m) }
func (*GetShareDownloadUrlResponse) ProtoMessage()    {}
func (*GetShareDownloadUrlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{9}
}

func (m *GetShareDownloadUrlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShareDownloadUrlResponse.Unmarshal(m, b)
}
func (m *GetShareDownloadUrlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShareDownloadUrlResponse.Marshal(b, m, deterministic)
}
func (m *GetShareDownloadUrlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShareDownloadUrlResponse.Merge(m, src)
}
func (m *GetShareDownloadUrlResponse) XXX_Size() int {
	return xxx_messageInfo_GetShareDownloadUrlResponse.Size(m)
}
func (m *GetShareDownloadUrlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShareDownloadUrlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShareDownloadUrlResponse proto.InternalMessageInfo

func (m *GetShareDownloadUrlResponse) GetDownloadToken() string {
	if m != nil {
		return m.DownloadToken
	}
	return ""
}

type CreateShareRequest struct {
	SharerId             int64    `protobuf:"varint,1,opt,name=sharer_id,json=sharerId,proto3" json:"sharer_id,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	ParentId             int64    `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	FolderIds            []int64  `protobuf:"varint,4,rep,packed,name=folder_ids,json=folderIds,proto3" json:"folder_ids,omitempty"`
	FileIds              []int64  `protobuf:"varint,5,rep,packed,name=file_ids,json=fileIds,proto3" json:"file_ids,omitempty"`
	Exp                  int64    `protobuf:"varint,6,opt,name=exp,proto3" json:"exp,omitempty"`
	IsPublic             bool     `protobuf:"varint,7,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShareRequest) Reset()         { *m = CreateShareRequest{} }
func (m *CreateShareRequest) String() string { return proto.CompactTextString(m) }
func (*CreateShareRequest) ProtoMessage()    {}
func (*CreateShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{10}
}

func (m *CreateShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShareRequest.Unmarshal(m, b)
}
func (m *CreateShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShareRequest.Marshal(b, m, deterministic)
}
func (m *CreateShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShareRequest.Merge(m, src)
}
func (m *CreateShareRequest) XXX_Size() int {
	return xxx_messageInfo_CreateShareRequest.Size(m)
}
func (m *CreateShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShareRequest proto.InternalMessageInfo

func (m *CreateShareRequest) GetSharerId() int64 {
	if m != nil {
		return m.SharerId
	}
	return 0
}

func (m *CreateShareRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CreateShareRequest) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *CreateShareRequest) GetFolderIds() []int64 {
	if m != nil {
		return m.FolderIds
	}
	return nil
}

func (m *CreateShareRequest) GetFileIds() []int64 {
	if m != nil {
		return m.FileIds
	}
	return nil
}

func (m *CreateShareRequest) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *CreateShareRequest) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

type CreateShareResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShareResponse) Reset()         { *m = CreateShareResponse{} }
func (m *CreateShareResponse) String() string { return proto.CompactTextString(m) }
func (*CreateShareResponse) ProtoMessage()    {}
func (*CreateShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{11}
}

func (m *CreateShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShareResponse.Unmarshal(m, b)
}
func (m *CreateShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShareResponse.Marshal(b, m, deterministic)
}
func (m *CreateShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShareResponse.Merge(m, src)
}
func (m *CreateShareResponse) XXX_Size() int {
	return xxx_messageInfo_CreateShareResponse.Size(m)
}
func (m *CreateShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShareResponse proto.InternalMessageInfo

func (m *CreateShareResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateShareResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UpdateShareRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SharerId             int64    `protobuf:"varint,2,opt,name=sharer_id,json=sharerId,proto3" json:"sharer_id,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Exp                  int64    `protobuf:"varint,7,opt,name=exp,proto3" json:"exp,omitempty"`
	IsPublic             bool     `protobuf:"varint,8,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateShareRequest) Reset()         { *m = UpdateShareRequest{} }
func (m *UpdateShareRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateShareRequest) ProtoMessage()    {}
func (*UpdateShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{12}
}

func (m *UpdateShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateShareRequest.Unmarshal(m, b)
}
func (m *UpdateShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateShareRequest.Marshal(b, m, deterministic)
}
func (m *UpdateShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateShareRequest.Merge(m, src)
}
func (m *UpdateShareRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateShareRequest.Size(m)
}
func (m *UpdateShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateShareRequest proto.InternalMessageInfo

func (m *UpdateShareRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateShareRequest) GetSharerId() int64 {
	if m != nil {
		return m.SharerId
	}
	return 0
}

func (m *UpdateShareRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *UpdateShareRequest) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *UpdateShareRequest) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

type UpdateShareResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateShareResponse) Reset()         { *m = UpdateShareResponse{} }
func (m *UpdateShareResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateShareResponse) ProtoMessage()    {}
func (*UpdateShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{13}
}

func (m *UpdateShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateShareResponse.Unmarshal(m, b)
}
func (m *UpdateShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateShareResponse.Marshal(b, m, deterministic)
}
func (m *UpdateShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateShareResponse.Merge(m, src)
}
func (m *UpdateShareResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateShareResponse.Size(m)
}
func (m *UpdateShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateShareResponse proto.InternalMessageInfo

type CancelShareRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SharerId             int64    `protobuf:"varint,2,opt,name=sharer_id,json=sharerId,proto3" json:"sharer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelShareRequest) Reset()         { *m = CancelShareRequest{} }
func (m *CancelShareRequest) String() string { return proto.CompactTextString(m) }
func (*CancelShareRequest) ProtoMessage()    {}
func (*CancelShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{14}
}

func (m *CancelShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelShareRequest.Unmarshal(m, b)
}
func (m *CancelShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelShareRequest.Marshal(b, m, deterministic)
}
func (m *CancelShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelShareRequest.Merge(m, src)
}
func (m *CancelShareRequest) XXX_Size() int {
	return xxx_messageInfo_CancelShareRequest.Size(m)
}
func (m *CancelShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelShareRequest proto.InternalMessageInfo

func (m *CancelShareRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CancelShareRequest) GetSharerId() int64 {
	if m != nil {
		return m.SharerId
	}
	return 0
}

type CancelShareResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelShareResponse) Reset()         { *m = CancelShareResponse{} }
func (m *CancelShareResponse) String() string { return proto.CompactTextString(m) }
func (*CancelShareResponse) ProtoMessage()    {}
func (*CancelShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{15}
}

func (m *CancelShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelShareResponse.Unmarshal(m, b)
}
func (m *CancelShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelShareResponse.Marshal(b, m, deterministic)
}
func (m *CancelShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelShareResponse.Merge(m, src)
}
func (m *CancelShareResponse) XXX_Size() int {
	return xxx_messageInfo_CancelShareResponse.Size(m)
}
func (m *CancelShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelShareResponse proto.InternalMessageInfo

type RemoveShareEntryRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SharerId             int64    `protobuf:"varint,2,opt,name=sharer_id,json=sharerId,proto3" json:"sharer_id,omitempty"`
	RemoveFiles          []int64  `protobuf:"varint,3,rep,packed,name=remove_files,json=removeFiles,proto3" json:"remove_files,omitempty"`
	RemoveFolders        []int64  `protobuf:"varint,4,rep,packed,name=remove_folders,json=removeFolders,proto3" json:"remove_folders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveShareEntryRequest) Reset()         { *m = RemoveShareEntryRequest{} }
func (m *RemoveShareEntryRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveShareEntryRequest) ProtoMessage()    {}
func (*RemoveShareEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{16}
}

func (m *RemoveShareEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveShareEntryRequest.Unmarshal(m, b)
}
func (m *RemoveShareEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveShareEntryRequest.Marshal(b, m, deterministic)
}
func (m *RemoveShareEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveShareEntryRequest.Merge(m, src)
}
func (m *RemoveShareEntryRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveShareEntryRequest.Size(m)
}
func (m *RemoveShareEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveShareEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveShareEntryRequest proto.InternalMessageInfo

func (m *RemoveShareEntryRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RemoveShareEntryRequest) GetSharerId() int64 {
	if m != nil {
		return m.SharerId
	}
	return 0
}

func (m *RemoveShareEntryRequest) GetRemoveFiles() []int64 {
	if m != nil {
		return m.RemoveFiles
	}
	return nil
}

func (m *RemoveShareEntryRequest) GetRemoveFolders() []int64 {
	if m != nil {
		return m.RemoveFolders
	}
	return nil
}

type RemoveShareEntryResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveShareEntryResponse) Reset()         { *m = RemoveShareEntryResponse{} }
func (m *RemoveShareEntryResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveShareEntryResponse) ProtoMessage()    {}
func (*RemoveShareEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91733f900302165, []int{17}
}

func (m *RemoveShareEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveShareEntryResponse.Unmarshal(m, b)
}
func (m *RemoveShareEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveShareEntryResponse.Marshal(b, m, deterministic)
}
func (m *RemoveShareEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveShareEntryResponse.Merge(m, src)
}
func (m *RemoveShareEntryResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveShareEntryResponse.Size(m)
}
func (m *RemoveShareEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveShareEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveShareEntryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ShareStatusRequest)(nil), "fullbottle.srv.share.ShareStatusRequest")
	proto.RegisterType((*ShareStatusResponse)(nil), "fullbottle.srv.share.ShareStatusResponse")
	proto.RegisterType((*AccessShareRequest)(nil), "fullbottle.srv.share.AccessShareRequest")
	proto.RegisterType((*AccessShareResponse)(nil), "fullbottle.srv.share.AccessShareResponse")
	proto.RegisterType((*GetShareInfoRequest)(nil), "fullbottle.srv.share.GetShareInfoRequest")
	proto.RegisterType((*GetShareInfoResponse)(nil), "fullbottle.srv.share.GetShareInfoResponse")
	proto.RegisterType((*GetShareFolderRequest)(nil), "fullbottle.srv.share.GetShareFolderRequest")
	proto.RegisterType((*GetShareFolderResponse)(nil), "fullbottle.srv.share.GetShareFolderResponse")
	proto.RegisterType((*GetShareDownloadUrlRequest)(nil), "fullbottle.srv.share.GetShareDownloadUrlRequest")
	proto.RegisterType((*GetShareDownloadUrlResponse)(nil), "fullbottle.srv.share.GetShareDownloadUrlResponse")
	proto.RegisterType((*CreateShareRequest)(nil), "fullbottle.srv.share.CreateShareRequest")
	proto.RegisterType((*CreateShareResponse)(nil), "fullbottle.srv.share.CreateShareResponse")
	proto.RegisterType((*UpdateShareRequest)(nil), "fullbottle.srv.share.UpdateShareRequest")
	proto.RegisterType((*UpdateShareResponse)(nil), "fullbottle.srv.share.UpdateShareResponse")
	proto.RegisterType((*CancelShareRequest)(nil), "fullbottle.srv.share.CancelShareRequest")
	proto.RegisterType((*CancelShareResponse)(nil), "fullbottle.srv.share.CancelShareResponse")
	proto.RegisterType((*RemoveShareEntryRequest)(nil), "fullbottle.srv.share.RemoveShareEntryRequest")
	proto.RegisterType((*RemoveShareEntryResponse)(nil), "fullbottle.srv.share.RemoveShareEntryResponse")
}

func init() { proto.RegisterFile("proto/share/share.proto", fileDescriptor_a91733f900302165) }

var fileDescriptor_a91733f900302165 = []byte{
	// 838 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xae, 0xe3, 0x7c, 0x4e, 0xd2, 0xaa, 0xda, 0xf4, 0x23, 0xaf, 0xab, 0x57, 0x6a, 0x2d, 0xbd,
	0x52, 0x5e, 0x3e, 0x1c, 0x28, 0x27, 0x84, 0x84, 0x84, 0x5a, 0x5a, 0x85, 0x13, 0x72, 0xdb, 0x13,
	0x87, 0xc8, 0xb1, 0x37, 0x8d, 0x55, 0xc7, 0x0e, 0x5e, 0x27, 0x05, 0x89, 0x0b, 0x57, 0x7e, 0x00,
	0x17, 0x0e, 0xfc, 0x20, 0x7e, 0x00, 0x7f, 0x07, 0xed, 0x87, 0x13, 0xaf, 0xe3, 0xad, 0x82, 0x72,
	0x69, 0xed, 0xf1, 0xec, 0x3c, 0xf3, 0xcc, 0xcc, 0x3e, 0x13, 0x38, 0x9c, 0xc6, 0x51, 0x12, 0xf5,
	0xc8, 0xd8, 0x89, 0x31, 0xff, 0x6b, 0x31, 0x0b, 0xda, 0x1b, 0xcd, 0x82, 0x60, 0x18, 0x25, 0x49,
	0x80, 0x2d, 0x12, 0xcf, 0x2d, 0xf6, 0xcd, 0x78, 0x7d, 0xeb, 0x27, 0xe3, 0xd9, 0xd0, 0x72, 0xa3,
	0x49, 0x6f, 0x8e, 0x6f, 0xdd, 0xb1, 0xef, 0xf6, 0x96, 0x8e, 0x3d, 0xf1, 0x8f, 0xc7, 0x94, 0x6c,
	0x3c, 0xaa, 0x79, 0x09, 0xe8, 0x8a, 0x06, 0xba, 0x4a, 0x9c, 0x64, 0x46, 0x6c, 0xfc, 0x71, 0x86,
	0x49, 0x82, 0xf6, 0xa0, 0x92, 0x44, 0x77, 0x38, 0xec, 0x68, 0xc7, 0x5a, 0xb7, 0x61, 0xf3, 0x17,
	0x74, 0x04, 0x8d, 0xb9, 0x8f, 0xef, 0x71, 0x3c, 0xf0, 0xbd, 0x4e, 0xe9, 0x58, 0xeb, 0xea, 0x76,
	0x9d, 0x1b, 0xfa, 0x9e, 0x89, 0xa1, 0x2d, 0x05, 0x22, 0xd3, 0x28, 0x24, 0x18, 0x1d, 0x40, 0x95,
	0x30, 0x0b, 0x0b, 0x55, 0xb1, 0xc5, 0x1b, 0xfa, 0x07, 0xea, 0xd1, 0x7d, 0x98, 0x0d, 0x55, 0x63,
	0xef, 0x7d, 0x8f, 0xc2, 0xf8, 0x64, 0x30, 0x9d, 0x0d, 0x03, 0xdf, 0xed, 0xe8, 0xc7, 0x5a, 0xb7,
	0x6e, 0xd7, 0x7d, 0xf2, 0x9e, 0xbd, 0x9b, 0x1f, 0x00, 0xbd, 0x71, 0x5d, 0x4c, 0x08, 0x03, 0x7b,
	0x38, 0x5f, 0x04, 0x65, 0x37, 0xf2, 0x30, 0x8b, 0xdf, 0xb0, 0xd9, 0xb3, 0xcc, 0x41, 0xcf, 0x71,
	0x78, 0x07, 0x6d, 0x29, 0xb8, 0xe0, 0x70, 0x02, 0x2d, 0x87, 0x99, 0x07, 0x59, 0x90, 0x26, 0xb7,
	0x5d, 0x33, 0xa8, 0x5d, 0xd0, 0xf1, 0xa7, 0xa9, 0x60, 0x42, 0x1f, 0xcd, 0x3b, 0x68, 0x5f, 0xe2,
	0x84, 0x05, 0xea, 0x87, 0xa3, 0x28, 0xcd, 0x74, 0x8d, 0x58, 0x0b, 0x32, 0x25, 0x65, 0xf1, 0xf3,
	0x89, 0xff, 0xd0, 0x60, 0x4f, 0x46, 0x13, 0xa9, 0xef, 0x40, 0xc9, 0xf7, 0x18, 0x88, 0x6e, 0x97,
	0x7c, 0x56, 0x5b, 0x36, 0x37, 0xd9, 0x16, 0x72, 0x43, 0xdf, 0x4b, 0x49, 0xe8, 0x0b, 0x12, 0xe8,
	0x3f, 0xd8, 0xf1, 0xa2, 0xfb, 0x30, 0x88, 0x1c, 0x6f, 0x90, 0xf8, 0x13, 0x4c, 0x3a, 0x65, 0xf6,
	0x71, 0x3b, 0xb5, 0x5e, 0x53, 0x23, 0xfa, 0x17, 0x80, 0xa6, 0x22, 0x5c, 0x2a, 0xcc, 0x85, 0x65,
	0xcb, 0x3e, 0x9b, 0x5f, 0x35, 0xd8, 0x4f, 0xb3, 0xbb, 0x88, 0x02, 0x0f, 0xc7, 0x1b, 0x57, 0x03,
	0x41, 0x79, 0xea, 0x24, 0x63, 0x96, 0x6b, 0xc3, 0x66, 0xcf, 0x72, 0x85, 0xca, 0xb9, 0x0a, 0x5d,
	0xc1, 0x41, 0x3e, 0x05, 0x51, 0xa2, 0x97, 0x50, 0x1d, 0x31, 0x0b, 0x43, 0x6f, 0x9e, 0x9e, 0x58,
	0xb9, 0x8b, 0x26, 0x1e, 0xf9, 0x31, 0x56, 0x5d, 0x71, 0xc0, 0xfc, 0xa9, 0x81, 0x91, 0x46, 0x3d,
	0x17, 0x15, 0xb9, 0x89, 0x83, 0x8d, 0xd9, 0x1d, 0x42, 0x6d, 0xe4, 0x07, 0x78, 0xd9, 0xe9, 0x2a,
	0x7d, 0xed, 0x7b, 0x0b, 0xda, 0x65, 0x15, 0xed, 0x4a, 0x8e, 0xf6, 0x39, 0x1c, 0x15, 0x26, 0x28,
	0xb8, 0x4b, 0xfd, 0xcd, 0xe4, 0xb8, 0xec, 0x2f, 0x35, 0x9a, 0xbf, 0x34, 0x40, 0x67, 0x31, 0x76,
	0x12, 0x2c, 0xdd, 0x3a, 0x69, 0x98, 0xb4, 0xdc, 0x30, 0x29, 0x2e, 0xdf, 0xd4, 0x89, 0x71, 0x98,
	0x64, 0x66, 0x98, 0x1b, 0xfa, 0x1e, 0x1d, 0x22, 0x5e, 0xd6, 0x81, 0xef, 0xd1, 0x39, 0xd3, 0xe9,
	0x10, 0x71, 0x4b, 0xdf, 0x63, 0x82, 0x21, 0x6a, 0x42, 0x27, 0x8c, 0x7e, 0xac, 0xf1, 0xa2, 0x90,
	0x74, 0x6e, 0xab, 0xcb, 0xb9, 0x95, 0x24, 0xa4, 0x96, 0x93, 0x90, 0x57, 0xd0, 0x96, 0xc8, 0x28,
	0xae, 0x4a, 0x61, 0x6b, 0xcc, 0x6f, 0x1a, 0xa0, 0x9b, 0xa9, 0x97, 0x2f, 0x85, 0x52, 0x30, 0xd5,
	0xb7, 0x2d, 0x2d, 0x90, 0x9e, 0x29, 0x90, 0x60, 0x52, 0x53, 0x30, 0xa9, 0xe7, 0x98, 0xec, 0x43,
	0x5b, 0xca, 0x85, 0x33, 0xa1, 0x9a, 0x7e, 0xe6, 0x84, 0x2e, 0x0e, 0x36, 0x4c, 0x91, 0xc6, 0x97,
	0x02, 0x89, 0xf8, 0xdf, 0x35, 0x38, 0xb4, 0xf1, 0x24, 0x9a, 0x73, 0xdc, 0xb7, 0x61, 0x12, 0x7f,
	0xde, 0xa0, 0x10, 0x27, 0xd0, 0x8a, 0x59, 0xb4, 0x01, 0x6d, 0x28, 0xe9, 0xe8, 0xac, 0xbb, 0x4d,
	0x6e, 0xbb, 0xa0, 0x26, 0x3a, 0xa7, 0xa9, 0x0b, 0x1b, 0x88, 0x74, 0x3e, 0xb6, 0x85, 0x13, 0x37,
	0x9a, 0x06, 0x74, 0x56, 0xf3, 0xe2, 0x49, 0x9f, 0xfe, 0xae, 0x41, 0x8b, 0x2f, 0x28, 0x1c, 0xcf,
	0x7d, 0x17, 0x23, 0x0f, 0x9a, 0x99, 0x85, 0x85, 0xba, 0x56, 0xd1, 0x7e, 0xb5, 0x56, 0x97, 0xa3,
	0xf1, 0xff, 0x1a, 0x9e, 0xa2, 0x52, 0x5b, 0x14, 0x25, 0xb3, 0x52, 0x54, 0x28, 0xab, 0x2b, 0x4d,
	0x85, 0x52, 0xb0, 0x9f, 0xcc, 0x2d, 0x74, 0x0b, 0xad, 0xac, 0xfc, 0x23, 0xc5, 0xe1, 0x82, 0x85,
	0x64, 0x3c, 0x5a, 0xc7, 0x75, 0x01, 0x34, 0x81, 0x1d, 0x59, 0x46, 0xd1, 0xe3, 0x87, 0xcf, 0x4b,
	0x7a, 0x6f, 0x3c, 0x59, 0xcf, 0x79, 0x01, 0xf7, 0x65, 0xb9, 0x44, 0x33, 0xf2, 0x85, 0x9e, 0x3d,
	0x1c, 0x66, 0x55, 0x8a, 0x8d, 0xe7, 0x7f, 0x71, 0x22, 0xdb, 0xbb, 0x8c, 0x50, 0xa8, 0x7a, 0xb7,
	0x2a, 0x8c, 0xaa, 0xde, 0x15, 0xa8, 0x0e, 0x47, 0xc9, 0x5c, 0x62, 0x15, 0xca, 0xaa, 0xe6, 0xa8,
	0x50, 0x8a, 0x14, 0x81, 0x73, 0x59, 0x5e, 0x65, 0x25, 0x97, 0x15, 0xd9, 0x50, 0x72, 0x29, 0xd0,
	0x85, 0x2d, 0x44, 0x60, 0x37, 0x7f, 0x01, 0xd1, 0xd3, 0xe2, 0x00, 0x0a, 0x01, 0x31, 0xac, 0x75,
	0xdd, 0x53, 0xd0, 0x61, 0x95, 0xfd, 0x92, 0x7d, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x25, 0x94,
	0x4f, 0x04, 0x3a, 0x0b, 0x00, 0x00,
}
