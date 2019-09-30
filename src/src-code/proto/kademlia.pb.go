
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
	NodeID               string   `protobuf:"bytes,2,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
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

func (m *FindNodeRequest) GetNodeID() string {
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
	return fileDescriptor_79477522ae6a503a, []int{5}
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
	return fileDescriptor_79477522ae6a503a, []int{6}
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
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindValueResponse) Reset()         { *m = FindValueResponse{} }
func (m *FindValueResponse) String() string { return proto.CompactTextString(m) }
func (*FindValueResponse) ProtoMessage()    {}
func (*FindValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79477522ae6a503a, []int{7}
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
	return fileDescriptor_79477522ae6a503a, []int{8}
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

func init() { proto.RegisterFile("kademlia.proto", fileDescriptor_79477522ae6a503a) }

var fileDescriptor_79477522ae6a503a = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xd9, 0xba, 0x3f, 0xdd, 0x59, 0xb7, 0x77, 0x0b, 0x2f, 0xa3, 0xf4, 0x6a, 0x44, 0x90,
	0x5d, 0x45, 0x9c, 0x97, 0xee, 0x66, 0x2a, 0x13, 0x11, 0x44, 0x2a, 0xe8, 0x75, 0x5c, 0x82, 0x0b,
	0x9b, 0xc9, 0x6c, 0x32, 0xc1, 0xcf, 0xe6, 0x97, 0x93, 0x36, 0x69, 0xe9, 0xdc, 0xa4, 0x77, 0x79,
	0xd2, 0x9e, 0x5f, 0x4e, 0x9e, 0xf3, 0x04, 0xfa, 0x6b, 0xca, 0xf8, 0xfb, 0x46, 0x50, 0xb2, 0x4d,
	0x94, 0x51, 0xf8, 0x16, 0xba, 0x8f, 0x42, 0xbe, 0xc5, 0xfc, 0x63, 0xc7, 0xb5, 0x41, 0x23, 0x68,
	0x69, 0x2e, 0x19, 0x4f, 0xc2, 0xda, 0xb8, 0x36, 0xe9, 0xc4, 0x4e, 0xa1, 0x31, 0x74, 0x19, 0xd7,
	0x46, 0x48, 0x6a, 0x84, 0x92, 0x61, 0x3d, 0xfb, 0x58, 0xde, 0xc2, 0x57, 0x10, 0x58, 0x90, 0xde,
	0x2a, 0xa9, 0xf9, 0x9f, 0xa4, 0x08, 0xfc, 0xc4, 0xfd, 0xe3, 0x30, 0x85, 0xc6, 0x33, 0x08, 0x9e,
	0x8c, 0x4a, 0x78, 0x55, 0x37, 0xff, 0xa1, 0xf9, 0x49, 0x37, 0x3b, 0x0b, 0x08, 0x62, 0x2b, 0xf0,
	0x25, 0xf4, 0x5c, 0x75, 0x45, 0x0b, 0x08, 0x1a, 0x2b, 0xaa, 0x57, 0xee, 0xf8, 0x6c, 0x8d, 0xe7,
	0xf0, 0x6f, 0x21, 0x24, 0x7b, 0x50, 0xac, 0xf2, 0xf4, 0x11, 0xb4, 0xa4, 0x62, 0xfc, 0xee, 0xc6,
	0x01, 0x9c, 0xc2, 0x33, 0x18, 0xa4, 0x88, 0xe7, 0xb4, 0x99, 0x2a, 0xc6, 0x00, 0xbc, 0x35, 0xff,
	0x72, 0x80, 0x74, 0x89, 0x17, 0xb6, 0xda, 0x36, 0x50, 0xe9, 0xa1, 0x27, 0x98, 0x0e, 0xeb, 0x63,
	0x6f, 0xd2, 0x9d, 0xfa, 0xe4, 0x5a, 0x49, 0x43, 0x97, 0x26, 0x4e, 0x37, 0xf1, 0x1c, 0x86, 0xa5,
	0x2e, 0x2a, 0x40, 0xc7, 0x8d, 0x7c, 0x81, 0xb6, 0x43, 0x16, 0x77, 0x65, 0x79, 0xa1, 0x55, 0x28,
	0x84, 0x36, 0x65, 0x2c, 0xe1, 0x5a, 0xbb, 0x3b, 0xe4, 0x32, 0x9d, 0x2f, 0x13, 0xda, 0x50, 0xb9,
	0xe4, 0xa1, 0x67, 0xe7, 0x9b, 0xeb, 0xe9, 0x77, 0x0d, 0xfc, 0x7b, 0x97, 0x3f, 0x74, 0x02, 0x8d,
	0x34, 0x30, 0x28, 0x20, 0xa5, 0x00, 0x46, 0x3d, 0xb2, 0x97, 0xa2, 0x53, 0x68, 0x66, 0x33, 0x45,
	0x3d, 0x52, 0x4e, 0x46, 0xd4, 0x27, 0xfb, 0xa3, 0x3e, 0x03, 0x3f, 0x77, 0x0f, 0x0d, 0xc8, 0xaf,
	0x49, 0x46, 0x43, 0x72, 0x60, 0xed, 0x39, 0x74, 0x0a, 0x9b, 0x8e, 0x54, 0x20, 0x72, 0x60, 0xe2,
	0x6b, 0x2b, 0x7b, 0x31, 0x17, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xa6, 0xed, 0x28, 0x43,
	0x03, 0x00, 0x00,
}

