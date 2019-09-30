// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kademlia.proto

package kademlia

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

type PingRequest struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Destination          string   `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79477522ae6a503a, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PingRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type PingResponse struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Response             string   `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79477522ae6a503a, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type StoreRequest struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79477522ae6a503a, []int{2}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *StoreRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type StoreResponse struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79477522ae6a503a, []int{3}
}

func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (m *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(m, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *StoreResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type FindNodeRequest struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNodeRequest) Reset()         { *m = FindNodeRequest{} }
func (m *FindNodeRequest) String() string { return proto.CompactTextString(m) }
func (*FindNodeRequest) ProtoMessage()    {}
func (*FindNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79477522ae6a503a, []int{4}
}

func (m *FindNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNodeRequest.Unmarshal(m, b)
}
func (m *FindNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNodeRequest.Marshal(b, m, deterministic)
}
func (m *FindNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNodeRequest.Merge(m, src)
}
func (m *FindNodeRequest) XXX_Size() int {
	return xxx_messageInfo_FindNodeRequest.Size(m)
}
func (m *FindNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindNodeRequest proto.InternalMessageInfo

func (m *FindNodeRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *FindNodeRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type FindNodeResponse struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Ids                  []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNodeResponse) Reset()         { *m = FindNodeResponse{} }
func (m *FindNodeResponse) String() string { return proto.CompactTextString(m) }
func (*FindNodeResponse) ProtoMessage()    {}
func (*FindNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79477522ae6a503a, []int{5}
}

func (m *FindNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNodeResponse.Unmarshal(m, b)
}
func (m *FindNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNodeResponse.Marshal(b, m, deterministic)
}
func (m *FindNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNodeResponse.Merge(m, src)
}
func (m *FindNodeResponse) XXX_Size() int {
	return xxx_messageInfo_FindNodeResponse.Size(m)
}
func (m *FindNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindNodeResponse proto.InternalMessageInfo

func (m *FindNodeResponse) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *FindNodeResponse) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type FindValueResponse struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindValueResponse) Reset()         { *m = FindValueResponse{} }
func (m *FindValueResponse) String() string { return proto.CompactTextString(m) }
func (*FindValueResponse) ProtoMessage()    {}
func (*FindValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79477522ae6a503a, []int{6}
}

func (m *FindValueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindValueResponse.Unmarshal(m, b)
}
func (m *FindValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindValueResponse.Marshal(b, m, deterministic)
}
func (m *FindValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindValueResponse.Merge(m, src)
}
func (m *FindValueResponse) XXX_Size() int {
	return xxx_messageInfo_FindValueResponse.Size(m)
}
func (m *FindValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindValueResponse proto.InternalMessageInfo

func (m *FindValueResponse) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *FindValueResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "PingRequest")
	proto.RegisterType((*PingResponse)(nil), "PingResponse")
	proto.RegisterType((*StoreRequest)(nil), "StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "StoreResponse")
	proto.RegisterType((*FindNodeRequest)(nil), "FindNodeRequest")
	proto.RegisterType((*FindNodeResponse)(nil), "FindNodeResponse")
	proto.RegisterType((*FindValueResponse)(nil), "FindValueResponse")
}

func init() { proto.RegisterFile("kademlia.proto", fileDescriptor_79477522ae6a503a) }

var fileDescriptor_79477522ae6a503a = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x24, 0xfd, 0x08, 0xc9, 0x34, 0xa9, 0xe9, 0x22, 0x25, 0xe4, 0x14, 0x56, 0x90, 0x9e, 0x56,
	0xd4, 0xa3, 0xbd, 0xd4, 0x83, 0x22, 0x82, 0x48, 0x04, 0xef, 0x91, 0x5d, 0xec, 0x62, 0xdd, 0xad,
	0xd9, 0xd4, 0x3f, 0xe7, 0x9f, 0x93, 0x24, 0x6b, 0x4d, 0x6d, 0x25, 0xb7, 0x37, 0x49, 0x66, 0xde,
	0x9b, 0x99, 0x60, 0xfc, 0x96, 0x73, 0xf1, 0xbe, 0x92, 0x39, 0x5b, 0x17, 0xba, 0xd4, 0xf4, 0x16,
	0xa3, 0x47, 0xa9, 0x5e, 0x33, 0xf1, 0xb1, 0x11, 0xa6, 0x24, 0x53, 0xb8, 0x46, 0x28, 0x2e, 0x8a,
	0xd8, 0x49, 0x9d, 0x99, 0x9f, 0x59, 0x44, 0x52, 0x8c, 0xb8, 0x30, 0xa5, 0x54, 0x79, 0x29, 0xb5,
	0x8a, 0x7b, 0xf5, 0xcb, 0xf6, 0x23, 0x7a, 0x8d, 0xa0, 0x11, 0x32, 0x6b, 0xad, 0x8c, 0xf8, 0x57,
	0x29, 0x81, 0x57, 0xd8, 0x6f, 0xac, 0xcc, 0x16, 0xd3, 0x39, 0x82, 0xa7, 0x52, 0x17, 0xa2, 0xeb,
	0x9a, 0x63, 0x0c, 0x3f, 0xf3, 0xd5, 0xa6, 0x11, 0x08, 0xb2, 0x06, 0xd0, 0x2b, 0x84, 0x96, 0xdd,
	0x71, 0x02, 0xc1, 0x60, 0x99, 0x9b, 0xa5, 0x5d, 0x5f, 0xcf, 0x74, 0x81, 0xa3, 0x1b, 0xa9, 0xf8,
	0x83, 0xe6, 0x9d, 0xdb, 0xa7, 0x70, 0x95, 0xe6, 0xe2, 0x8e, 0x5b, 0x01, 0x8b, 0xe8, 0x1c, 0xd1,
	0xaf, 0x44, 0xc7, 0x09, 0x11, 0xfa, 0x92, 0x9b, 0xb8, 0x97, 0xf6, 0x67, 0x7e, 0x56, 0x8d, 0x74,
	0x81, 0x49, 0xc5, 0x7e, 0xae, 0xac, 0x74, 0xd2, 0x0f, 0x06, 0x70, 0xf1, 0xe5, 0xc0, 0xbb, 0xb7,
	0xf5, 0x92, 0x13, 0x0c, 0xaa, 0x3e, 0x48, 0xc0, 0x5a, 0xfd, 0x26, 0x21, 0xdb, 0x29, 0xe9, 0x14,
	0xc3, 0x3a, 0x32, 0x12, 0xb2, 0x76, 0xf0, 0xc9, 0x98, 0xed, 0x26, 0x79, 0x06, 0xef, 0xc7, 0x1a,
	0x89, 0xd8, 0x9f, 0xa0, 0x92, 0x09, 0xdb, 0xf3, 0x7d, 0x0e, 0x7f, 0xeb, 0xe6, 0x00, 0x83, 0xb0,
	0x3d, 0xaf, 0x2f, 0x6e, 0xfd, 0x43, 0x5e, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x35, 0x46, 0x89,
	0x84, 0xa2, 0x02, 0x00, 0x00,
}
