// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/notes_api.proto

package api

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

type CreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SubId                int64    `protobuf:"varint,2,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetSubId() int64 {
	if m != nil {
		return m.SubId
	}
	return 0
}

type CreateResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SubId                int64    `protobuf:"varint,4,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateResponse) GetSubId() int64 {
	if m != nil {
		return m.SubId
	}
	return 0
}

type ReadRequest struct {
	// id записи
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{2}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadResponse struct {
	Topics               []*ReadResponse_Topics `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{3}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetTopics() []*ReadResponse_Topics {
	if m != nil {
		return m.Topics
	}
	return nil
}

type ReadResponse_Topics struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SubId                int64    `protobuf:"varint,4,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse_Topics) Reset()         { *m = ReadResponse_Topics{} }
func (m *ReadResponse_Topics) String() string { return proto.CompactTextString(m) }
func (*ReadResponse_Topics) ProtoMessage()    {}
func (*ReadResponse_Topics) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{3, 0}
}

func (m *ReadResponse_Topics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse_Topics.Unmarshal(m, b)
}
func (m *ReadResponse_Topics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse_Topics.Marshal(b, m, deterministic)
}
func (m *ReadResponse_Topics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse_Topics.Merge(m, src)
}
func (m *ReadResponse_Topics) XXX_Size() int {
	return xxx_messageInfo_ReadResponse_Topics.Size(m)
}
func (m *ReadResponse_Topics) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse_Topics.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse_Topics proto.InternalMessageInfo

func (m *ReadResponse_Topics) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadResponse_Topics) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ReadResponse_Topics) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadResponse_Topics) GetSubId() int64 {
	if m != nil {
		return m.SubId
	}
	return 0
}

type UpdateRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SubId                int64    `protobuf:"varint,3,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetSubId() int64 {
	if m != nil {
		return m.SubId
	}
	return 0
}

type UpdateResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SubId                int64    `protobuf:"varint,4,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{5}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateResponse) GetSubId() int64 {
	if m != nil {
		return m.SubId
	}
	return 0
}

type DeleteRequest struct {
	Md                   []*DeleteRequest_MultiDelete `protobuf:"bytes,1,rep,name=md,proto3" json:"md,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetMd() []*DeleteRequest_MultiDelete {
	if m != nil {
		return m.Md
	}
	return nil
}

type DeleteRequest_MultiDelete struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest_MultiDelete) Reset()         { *m = DeleteRequest_MultiDelete{} }
func (m *DeleteRequest_MultiDelete) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest_MultiDelete) ProtoMessage()    {}
func (*DeleteRequest_MultiDelete) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{6, 0}
}

func (m *DeleteRequest_MultiDelete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest_MultiDelete.Unmarshal(m, b)
}
func (m *DeleteRequest_MultiDelete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest_MultiDelete.Marshal(b, m, deterministic)
}
func (m *DeleteRequest_MultiDelete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest_MultiDelete.Merge(m, src)
}
func (m *DeleteRequest_MultiDelete) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest_MultiDelete.Size(m)
}
func (m *DeleteRequest_MultiDelete) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest_MultiDelete.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest_MultiDelete proto.InternalMessageInfo

func (m *DeleteRequest_MultiDelete) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteRequest_MultiDelete) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteResponse struct {
	Message              []string `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff16eb62a4823ab, []int{7}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetMessage() []string {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*ReadResponse_Topics)(nil), "ReadResponse.Topics")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteRequest_MultiDelete)(nil), "DeleteRequest.MultiDelete")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
}

func init() {
	proto.RegisterFile("api/notes_api.proto", fileDescriptor_4ff16eb62a4823ab)
}

var fileDescriptor_4ff16eb62a4823ab = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0x8d, 0xed, 0xc4, 0xd9, 0x4c, 0x62, 0x07, 0xb4, 0xbb, 0xac, 0x31, 0x2c, 0x04, 0x41, 0x21,
	0xa4, 0x45, 0xa1, 0xe9, 0xad, 0xbd, 0xb5, 0xbd, 0xa4, 0xd0, 0x1e, 0x44, 0x7b, 0x29, 0x85, 0xa0,
	0x54, 0xa2, 0x08, 0x62, 0x5b, 0x8d, 0xec, 0x3f, 0xd3, 0xff, 0xd1, 0xff, 0x57, 0xfc, 0xa1, 0xc6,
	0x0a, 0x39, 0xf4, 0x92, 0x9b, 0x35, 0xa3, 0x79, 0x6f, 0xde, 0xf3, 0x13, 0xfc, 0x66, 0x4a, 0xce,
	0xd3, 0x2c, 0x17, 0x7a, 0xc5, 0x94, 0x24, 0x6a, 0x9b, 0xe5, 0x19, 0xbe, 0x84, 0xe0, 0x66, 0x2b,
	0x58, 0x2e, 0xa8, 0x78, 0x2f, 0x84, 0xce, 0x11, 0x82, 0x6e, 0xca, 0x12, 0x11, 0x39, 0x13, 0x67,
	0x3a, 0xa0, 0xd5, 0x37, 0xfa, 0x0b, 0xbe, 0x2e, 0xd6, 0x2b, 0xc9, 0x23, 0x77, 0xe2, 0x4c, 0x3d,
	0xda, 0xd3, 0xc5, 0x7a, 0xc9, 0x31, 0x87, 0xd0, 0xcc, 0x6a, 0x95, 0xa5, 0x5a, 0xa0, 0x10, 0x5c,
	0xc9, 0xab, 0x51, 0x8f, 0xba, 0x92, 0xa3, 0x7f, 0xd0, 0x2f, 0xb4, 0xd8, 0xee, 0x26, 0xfd, 0xf2,
	0xb8, 0xe4, 0xdf, 0x2c, 0xde, 0x41, 0x96, 0x6e, 0x9b, 0xe5, 0x3f, 0x0c, 0xa9, 0x60, 0xdc, 0xec,
	0xb7, 0x47, 0x81, 0x3f, 0x1c, 0x18, 0xd5, 0xfd, 0x66, 0x87, 0x33, 0xf0, 0xf3, 0x4c, 0xc9, 0x57,
	0x1d, 0x39, 0x13, 0x6f, 0x3a, 0x5c, 0xfc, 0x21, 0xed, 0x36, 0x79, 0xac, 0x7a, 0xb4, 0xb9, 0x13,
	0xbf, 0x80, 0x5f, 0x57, 0x8e, 0xb2, 0xfb, 0x1d, 0x04, 0x4f, 0x8a, 0xb7, 0xdc, 0xdd, 0x27, 0x31,
	0x58, 0xee, 0x41, 0x2c, 0x6f, 0xcf, 0x6d, 0x83, 0x75, 0x44, 0xb7, 0x53, 0x08, 0x6e, 0xc5, 0x46,
	0xec, 0x36, 0x9e, 0x81, 0x9b, 0xf0, 0xc6, 0xca, 0x98, 0x58, 0x3d, 0x72, 0x5f, 0x6c, 0x72, 0xd9,
	0x94, 0xdc, 0x84, 0xc7, 0xe7, 0x30, 0x6c, 0x95, 0x7e, 0x22, 0x16, 0xcf, 0x20, 0x34, 0x98, 0x8d,
	0xaa, 0x08, 0xfa, 0x89, 0xd0, 0x9a, 0xbd, 0x89, 0x8a, 0x75, 0x40, 0xcd, 0x71, 0xf1, 0xe9, 0x40,
	0xef, 0xa1, 0xcc, 0x2f, 0x3a, 0x05, 0xbf, 0x4e, 0x1e, 0x0a, 0x89, 0x15, 0xdf, 0x78, 0x4c, 0xec,
	0x48, 0xe2, 0x0e, 0x3a, 0x81, 0x6e, 0x99, 0x00, 0x34, 0x22, 0xad, 0x1c, 0xc5, 0x81, 0x15, 0x0b,
	0xdc, 0x29, 0x31, 0x6b, 0x7f, 0x51, 0x48, 0xac, 0x9f, 0x16, 0x8f, 0x89, 0x6d, 0x7c, 0x7d, 0xd9,
	0x88, 0xb4, 0x3d, 0x89, 0xc7, 0xc4, 0xd6, 0x83, 0x3b, 0xd7, 0xf0, 0xfc, 0x6b, 0xce, 0x94, 0xbc,
	0x62, 0x4a, 0xae, 0xfd, 0xea, 0xd9, 0x5d, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x60, 0x60,
	0x1a, 0x8d, 0x03, 0x00, 0x00,
}
