// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: attachment.proto

package attachment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code int32

const (
	Code_OK                        Code = 0
	Code_ERROR_UNKNOWN             Code = 1
	Code_ERROR_ATTACHMENT_NOTFOUND Code = 2
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0: "OK",
		1: "ERROR_UNKNOWN",
		2: "ERROR_ATTACHMENT_NOTFOUND",
	}
	Code_value = map[string]int32{
		"OK":                        0,
		"ERROR_UNKNOWN":             1,
		"ERROR_ATTACHMENT_NOTFOUND": 2,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_attachment_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_attachment_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{0}
}

// The response message containing the greetings
type AttachmentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // Attachment id
	CardId        uint64 `protobuf:"varint,2,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`                      // 所属卡片id
	Link          string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`                                         // Attachment链接
	DownloadCount uint64 `protobuf:"varint,4,opt,name=download_count,json=downloadCount,proto3" json:"download_count,omitempty"` //下载量
	CreatedAt     string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`              // Attachment创建时间
	UpdatedAt     string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`              // Attachment更新时间
}

func (x *AttachmentInfo) Reset() {
	*x = AttachmentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentInfo) ProtoMessage() {}

func (x *AttachmentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentInfo.ProtoReflect.Descriptor instead.
func (*AttachmentInfo) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{0}
}

func (x *AttachmentInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttachmentInfo) GetCardId() uint64 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *AttachmentInfo) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *AttachmentInfo) GetDownloadCount() uint64 {
	if x != nil {
		return x.DownloadCount
	}
	return 0
}

func (x *AttachmentInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AttachmentInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AddAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId uint64 `protobuf:"varint,1,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"` // 所属卡片id
	Link   string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`                    // Attachment链接
}

func (x *AddAttachmentRequest) Reset() {
	*x = AddAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAttachmentRequest) ProtoMessage() {}

func (x *AddAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAttachmentRequest.ProtoReflect.Descriptor instead.
func (*AddAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{1}
}

func (x *AddAttachmentRequest) GetCardId() uint64 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *AddAttachmentRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type UpdateAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"` // Attachment链接
}

func (x *UpdateAttachmentRequest) Reset() {
	*x = UpdateAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttachmentRequest) ProtoMessage() {}

func (x *UpdateAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAttachmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAttachmentRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetAttachmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attachment *AttachmentInfo `protobuf:"bytes,1,opt,name=Attachment,proto3" json:"Attachment,omitempty"`
	Code       Code            `protobuf:"varint,2,opt,name=code,proto3,enum=attachment.Code" json:"code,omitempty"`
}

func (x *GetAttachmentResponse) Reset() {
	*x = GetAttachmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachmentResponse) ProtoMessage() {}

func (x *GetAttachmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachmentResponse.ProtoReflect.Descriptor instead.
func (*GetAttachmentResponse) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{3}
}

func (x *GetAttachmentResponse) GetAttachment() *AttachmentInfo {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *GetAttachmentResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

type GetAttachmentByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAttachmentByIdRequest) Reset() {
	*x = GetAttachmentByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachmentByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachmentByIdRequest) ProtoMessage() {}

func (x *GetAttachmentByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachmentByIdRequest.ProtoReflect.Descriptor instead.
func (*GetAttachmentByIdRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{4}
}

func (x *GetAttachmentByIdRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAttachmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attachment []*AttachmentInfo `protobuf:"bytes,1,rep,name=Attachment,proto3" json:"Attachment,omitempty"`
	Code       Code              `protobuf:"varint,2,opt,name=code,proto3,enum=attachment.Code" json:"code,omitempty"`
}

func (x *GetAttachmentsResponse) Reset() {
	*x = GetAttachmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachmentsResponse) ProtoMessage() {}

func (x *GetAttachmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachmentsResponse.ProtoReflect.Descriptor instead.
func (*GetAttachmentsResponse) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{5}
}

func (x *GetAttachmentsResponse) GetAttachment() []*AttachmentInfo {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *GetAttachmentsResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

type GetAttachmentsByCardIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId uint64 `protobuf:"varint,1,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
}

func (x *GetAttachmentsByCardIdRequest) Reset() {
	*x = GetAttachmentsByCardIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachmentsByCardIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachmentsByCardIdRequest) ProtoMessage() {}

func (x *GetAttachmentsByCardIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachmentsByCardIdRequest.ProtoReflect.Descriptor instead.
func (*GetAttachmentsByCardIdRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{6}
}

func (x *GetAttachmentsByCardIdRequest) GetCardId() uint64 {
	if x != nil {
		return x.CardId
	}
	return 0
}

type AddAttachmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code   `protobuf:"varint,1,opt,name=code,proto3,enum=attachment.Code" json:"code,omitempty"`
	Id   uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddAttachmentResponse) Reset() {
	*x = AddAttachmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAttachmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAttachmentResponse) ProtoMessage() {}

func (x *AddAttachmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAttachmentResponse.ProtoReflect.Descriptor instead.
func (*AddAttachmentResponse) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{7}
}

func (x *AddAttachmentResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

func (x *AddAttachmentResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateAttachmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=attachment.Code" json:"code,omitempty"`
}

func (x *UpdateAttachmentResponse) Reset() {
	*x = UpdateAttachmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAttachmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttachmentResponse) ProtoMessage() {}

func (x *UpdateAttachmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAttachmentResponse.ProtoReflect.Descriptor instead.
func (*UpdateAttachmentResponse) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateAttachmentResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

type DeleteAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAttachmentRequest) Reset() {
	*x = DeleteAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAttachmentRequest) ProtoMessage() {}

func (x *DeleteAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAttachmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAttachmentRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAttachmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=attachment.Code" json:"code,omitempty"`
}

func (x *DeleteAttachmentResponse) Reset() {
	*x = DeleteAttachmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAttachmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAttachmentResponse) ProtoMessage() {}

func (x *DeleteAttachmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAttachmentResponse.ProtoReflect.Descriptor instead.
func (*DeleteAttachmentResponse) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAttachmentResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

var File_attachment_proto protoreflect.FileDescriptor

var file_attachment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xb2,
	0x01, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x25,
	0x0a, 0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x43, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x2d, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x79, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7a,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x38, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x40, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x2a, 0x40, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x41,
	0x54, 0x54, 0x41, 0x43, 0x48, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x02, 0x32, 0x86, 0x08, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x24,
	0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x69, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x29, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x29, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x56, 0x0a,
	0x0d, 0x41, 0x64, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20,
	0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x2e, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x5f, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x23, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x5f, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x69, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x23, 0x2e, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x3b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attachment_proto_rawDescOnce sync.Once
	file_attachment_proto_rawDescData = file_attachment_proto_rawDesc
)

func file_attachment_proto_rawDescGZIP() []byte {
	file_attachment_proto_rawDescOnce.Do(func() {
		file_attachment_proto_rawDescData = protoimpl.X.CompressGZIP(file_attachment_proto_rawDescData)
	})
	return file_attachment_proto_rawDescData
}

var file_attachment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_attachment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_attachment_proto_goTypes = []interface{}{
	(Code)(0),                             // 0: attachment.Code
	(*AttachmentInfo)(nil),                // 1: attachment.AttachmentInfo
	(*AddAttachmentRequest)(nil),          // 2: attachment.AddAttachmentRequest
	(*UpdateAttachmentRequest)(nil),       // 3: attachment.UpdateAttachmentRequest
	(*GetAttachmentResponse)(nil),         // 4: attachment.GetAttachmentResponse
	(*GetAttachmentByIdRequest)(nil),      // 5: attachment.GetAttachmentByIdRequest
	(*GetAttachmentsResponse)(nil),        // 6: attachment.GetAttachmentsResponse
	(*GetAttachmentsByCardIdRequest)(nil), // 7: attachment.GetAttachmentsByCardIdRequest
	(*AddAttachmentResponse)(nil),         // 8: attachment.AddAttachmentResponse
	(*UpdateAttachmentResponse)(nil),      // 9: attachment.UpdateAttachmentResponse
	(*DeleteAttachmentRequest)(nil),       // 10: attachment.DeleteAttachmentRequest
	(*DeleteAttachmentResponse)(nil),      // 11: attachment.DeleteAttachmentResponse
}
var file_attachment_proto_depIdxs = []int32{
	1,  // 0: attachment.GetAttachmentResponse.Attachment:type_name -> attachment.AttachmentInfo
	0,  // 1: attachment.GetAttachmentResponse.code:type_name -> attachment.Code
	1,  // 2: attachment.GetAttachmentsResponse.Attachment:type_name -> attachment.AttachmentInfo
	0,  // 3: attachment.GetAttachmentsResponse.code:type_name -> attachment.Code
	0,  // 4: attachment.AddAttachmentResponse.code:type_name -> attachment.Code
	0,  // 5: attachment.UpdateAttachmentResponse.code:type_name -> attachment.Code
	0,  // 6: attachment.DeleteAttachmentResponse.code:type_name -> attachment.Code
	5,  // 7: attachment.Attachment.GetAttachmentById:input_type -> attachment.GetAttachmentByIdRequest
	5,  // 8: attachment.Attachment.GetAttachmentByIdStream:input_type -> attachment.GetAttachmentByIdRequest
	7,  // 9: attachment.Attachment.GetAttachmentsByCardId:input_type -> attachment.GetAttachmentsByCardIdRequest
	7,  // 10: attachment.Attachment.GetAttachmentsByCardIdStream:input_type -> attachment.GetAttachmentsByCardIdRequest
	2,  // 11: attachment.Attachment.AddAttachment:input_type -> attachment.AddAttachmentRequest
	2,  // 12: attachment.Attachment.AddAttachmentStream:input_type -> attachment.AddAttachmentRequest
	3,  // 13: attachment.Attachment.UpdateAttachment:input_type -> attachment.UpdateAttachmentRequest
	3,  // 14: attachment.Attachment.UpdateAttachmentStream:input_type -> attachment.UpdateAttachmentRequest
	10, // 15: attachment.Attachment.DeleteAttachment:input_type -> attachment.DeleteAttachmentRequest
	10, // 16: attachment.Attachment.DeleteAttachmentStream:input_type -> attachment.DeleteAttachmentRequest
	4,  // 17: attachment.Attachment.GetAttachmentById:output_type -> attachment.GetAttachmentResponse
	4,  // 18: attachment.Attachment.GetAttachmentByIdStream:output_type -> attachment.GetAttachmentResponse
	6,  // 19: attachment.Attachment.GetAttachmentsByCardId:output_type -> attachment.GetAttachmentsResponse
	6,  // 20: attachment.Attachment.GetAttachmentsByCardIdStream:output_type -> attachment.GetAttachmentsResponse
	8,  // 21: attachment.Attachment.AddAttachment:output_type -> attachment.AddAttachmentResponse
	8,  // 22: attachment.Attachment.AddAttachmentStream:output_type -> attachment.AddAttachmentResponse
	9,  // 23: attachment.Attachment.UpdateAttachment:output_type -> attachment.UpdateAttachmentResponse
	9,  // 24: attachment.Attachment.UpdateAttachmentStream:output_type -> attachment.UpdateAttachmentResponse
	11, // 25: attachment.Attachment.DeleteAttachment:output_type -> attachment.DeleteAttachmentResponse
	11, // 26: attachment.Attachment.DeleteAttachmentStream:output_type -> attachment.DeleteAttachmentResponse
	17, // [17:27] is the sub-list for method output_type
	7,  // [7:17] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_attachment_proto_init() }
func file_attachment_proto_init() {
	if File_attachment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attachment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAttachmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAttachmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachmentByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachmentsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachmentsByCardIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAttachmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAttachmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAttachmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attachment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAttachmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_attachment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attachment_proto_goTypes,
		DependencyIndexes: file_attachment_proto_depIdxs,
		EnumInfos:         file_attachment_proto_enumTypes,
		MessageInfos:      file_attachment_proto_msgTypes,
	}.Build()
	File_attachment_proto = out.File
	file_attachment_proto_rawDesc = nil
	file_attachment_proto_goTypes = nil
	file_attachment_proto_depIdxs = nil
}
