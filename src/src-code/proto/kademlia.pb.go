// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/kademlia.proto

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
	return fileDescriptor_5ad3b431acfc5603, []int{0}
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
	return fileDescriptor_5ad3b431acfc5603, []int{1}
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
	return fileDescriptor_5ad3b431acfc5603, []int{2}
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
	return fileDescriptor_5ad3b431acfc5603, []int{3}
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
	NodeID               *Contact `protobuf:"bytes,2,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNodeRequest) Reset()         { *m = FindNodeRequest{} }
func (m *FindNodeRequest) String() string { return proto.CompactTextString(m) }
func (*FindNodeRequest) ProtoMessage()    {}
func (*FindNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad3b431acfc5603, []int{4}
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

func (m *FindNodeRequest) GetNodeID() *Contact {
	if m != nil {
		return m.NodeID
	}
	return nil
}

type FindValueRequest struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindValueRequest) Reset()         { *m = FindValueRequest{} }
func (m *FindValueRequest) String() string { return proto.CompactTextString(m) }
func (*FindValueRequest) ProtoMessage()    {}
func (*FindValueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad3b431acfc5603, []int{5}
}

func (m *FindValueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindValueRequest.Unmarshal(m, b)
}
func (m *FindValueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindValueRequest.Marshal(b, m, deterministic)
}
func (m *FindValueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindValueRequest.Merge(m, src)
}
func (m *FindValueRequest) XXX_Size() int {
	return xxx_messageInfo_FindValueRequest.Size(m)
}
func (m *FindValueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindValueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindValueRequest proto.InternalMessageInfo

func (m *FindValueRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *FindValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type FindNodeResponse struct {
	Sender               string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Ids                  []*Contact `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindNodeResponse) Reset()         { *m = FindNodeResponse{} }
func (m *FindNodeResponse) String() string { return proto.CompactTextString(m) }
func (*FindNodeResponse) ProtoMessage()    {}
func (*FindNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad3b431acfc5603, []int{6}
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

func (m *FindNodeResponse) GetIds() []*Contact {
	if m != nil {
		return m.Ids
	}
	return nil
}

type FindValueResponse struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindValueResponse) Reset()         { *m = FindValueResponse{} }
func (m *FindValueResponse) String() string { return proto.CompactTextString(m) }
func (*FindValueResponse) ProtoMessage()    {}
func (*FindValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad3b431acfc5603, []int{7}
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

func (m *FindValueResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Contact struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Distance             string   `protobuf:"bytes,3,opt,name=distance,proto3" json:"distance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad3b431acfc5603, []int{8}
}

func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Contact) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Contact) GetDistance() string {
	if m != nil {
		return m.Distance
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "PingRequest")
	proto.RegisterType((*PingResponse)(nil), "PingResponse")
	proto.RegisterType((*StoreRequest)(nil), "StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "StoreResponse")
	proto.RegisterType((*FindNodeRequest)(nil), "FindNodeRequest")
	proto.RegisterType((*FindValueRequest)(nil), "FindValueRequest")
	proto.RegisterType((*FindNodeResponse)(nil), "FindNodeResponse")
	proto.RegisterType((*FindValueResponse)(nil), "FindValueResponse")
	proto.RegisterType((*Contact)(nil), "Contact")
}

func init() { proto.RegisterFile("proto/kademlia.proto", fileDescriptor_5ad3b431acfc5603) }

var fileDescriptor_5ad3b431acfc5603 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x4b, 0xf3, 0x30,
	0x18, 0x65, 0xeb, 0x3e, 0xba, 0x67, 0xdd, 0xfb, 0x6e, 0x61, 0x48, 0xe9, 0xd5, 0x88, 0x20, 0x5e,
	0x65, 0x30, 0x2f, 0xdd, 0x8d, 0x1f, 0x4c, 0x64, 0x20, 0x52, 0x41, 0xaf, 0xe3, 0x12, 0x5c, 0xd9,
	0x4c, 0x66, 0x93, 0x09, 0xfe, 0x37, 0x7f, 0x9c, 0xe4, 0xa3, 0xa3, 0x75, 0x48, 0xef, 0x72, 0x9e,
	0xf4, 0x39, 0x3d, 0x3d, 0xe7, 0x14, 0xc6, 0xbb, 0x5c, 0x6a, 0x39, 0xdd, 0x50, 0xc6, 0xdf, 0xb7,
	0x19, 0x25, 0x16, 0xe2, 0x3b, 0xe8, 0x3f, 0x66, 0xe2, 0x2d, 0xe5, 0x1f, 0x7b, 0xae, 0x34, 0x3a,
	0x81, 0x8e, 0xe2, 0x82, 0xf1, 0x3c, 0x6e, 0x4c, 0x1a, 0xe7, 0xbd, 0xd4, 0x23, 0x34, 0x81, 0x3e,
	0xe3, 0x4a, 0x67, 0x82, 0xea, 0x4c, 0x8a, 0xb8, 0x69, 0x2f, 0xcb, 0x23, 0x7c, 0x0d, 0x91, 0x23,
	0x52, 0x3b, 0x29, 0x14, 0xff, 0x93, 0x29, 0x81, 0x30, 0xf7, 0xcf, 0x78, 0x9a, 0x03, 0xc6, 0x73,
	0x88, 0x9e, 0xb4, 0xcc, 0x79, 0x9d, 0x9a, 0x31, 0xb4, 0x3f, 0xe9, 0x76, 0xef, 0x08, 0xa2, 0xd4,
	0x01, 0x7c, 0x09, 0x03, 0xbf, 0x5d, 0x23, 0x01, 0x41, 0x6b, 0x4d, 0xd5, 0xda, 0xbf, 0xde, 0x9e,
	0xf1, 0x12, 0xfe, 0x2f, 0x32, 0xc1, 0x1e, 0x24, 0xe3, 0xf5, 0x5e, 0x74, 0x84, 0x64, 0xfc, 0xfe,
	0xd6, 0x12, 0xf4, 0x67, 0x21, 0xb9, 0x91, 0x42, 0xd3, 0x95, 0x4e, 0xfd, 0x1c, 0xcf, 0x61, 0x68,
	0xc8, 0x9e, 0x8d, 0xac, 0x3a, 0xb6, 0x21, 0x04, 0x1b, 0xfe, 0xe5, 0xb5, 0x98, 0x23, 0x5e, 0xb8,
	0x6d, 0x27, 0xa5, 0xd6, 0xcd, 0x20, 0x63, 0x2a, 0x6e, 0x4e, 0x82, 0x8a, 0x10, 0x33, 0xc4, 0x57,
	0x30, 0x2a, 0xa9, 0xa8, 0x21, 0xaa, 0x58, 0xda, 0x2b, 0x2c, 0x7d, 0x81, 0xae, 0xa7, 0x34, 0x8b,
	0xf6, 0xeb, 0x58, 0xb1, 0xe8, 0x10, 0x8a, 0xa1, 0x4b, 0x19, 0xcb, 0xb9, 0x52, 0x7e, 0xb5, 0x80,
	0x26, 0x69, 0x96, 0x29, 0x4d, 0xc5, 0x8a, 0xc7, 0x81, 0x4b, 0xba, 0xc0, 0xb3, 0xef, 0x06, 0x84,
	0x4b, 0xdf, 0x44, 0x74, 0x0a, 0x2d, 0x53, 0x1d, 0x14, 0x91, 0x52, 0x15, 0x93, 0x01, 0xa9, 0xf4,
	0xe9, 0x0c, 0xda, 0x36, 0x5d, 0x34, 0x20, 0xe5, 0x8e, 0x24, 0xff, 0x48, 0x35, 0xf4, 0x29, 0x84,
	0x85, 0x7b, 0x68, 0x48, 0x7e, 0x65, 0x9a, 0x8c, 0xc8, 0x91, 0xb5, 0x33, 0xe8, 0x1d, 0x6c, 0x42,
	0xee, 0xbe, 0x1c, 0x5c, 0x82, 0xc8, 0x91, 0x8b, 0xaf, 0x1d, 0xfb, 0xf3, 0x5c, 0xfc, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xff, 0x0a, 0x6f, 0xce, 0x54, 0x03, 0x00, 0x00,
}
