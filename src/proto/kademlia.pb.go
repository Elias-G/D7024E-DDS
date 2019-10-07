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
	RpcID                string   `protobuf:"bytes,1,opt,name=rpcID,proto3" json:"rpcID,omitempty"`
	Sender               *Contact `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Destination          string   `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
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

func (m *PingRequest) GetRpcID() string {
	if m != nil {
		return m.RpcID
	}
	return ""
}

func (m *PingRequest) GetSender() *Contact {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *PingRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type PingResponse struct {
	RpcID                string   `protobuf:"bytes,1,opt,name=rpcID,proto3" json:"rpcID,omitempty"`
	Sender               *Contact `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Response             string   `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
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

func (m *PingResponse) GetRpcID() string {
	if m != nil {
		return m.RpcID
	}
	return ""
}

func (m *PingResponse) GetSender() *Contact {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type StoreRequest struct {
	RpcID                string   `protobuf:"bytes,1,opt,name=rpcID,proto3" json:"rpcID,omitempty"`
	Sender               *Contact `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
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

func (m *StoreRequest) GetRpcID() string {
	if m != nil {
		return m.RpcID
	}
	return ""
}

func (m *StoreRequest) GetSender() *Contact {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *StoreRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type StoreResponse struct {
	RpcID                string   `protobuf:"bytes,1,opt,name=rpcID,proto3" json:"rpcID,omitempty"`
	Sender               *Contact `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
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

func (m *StoreResponse) GetRpcID() string {
	if m != nil {
		return m.RpcID
	}
	return ""
}

func (m *StoreResponse) GetSender() *Contact {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *StoreResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type FindNodeRequest struct {
	RpcID                string   `protobuf:"bytes,1,opt,name=rpcID,proto3" json:"rpcID,omitempty"`
	Sender               *Contact `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	TargetId             string   `protobuf:"bytes,3,opt,name=targetId,proto3" json:"targetId,omitempty"`
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

func (m *FindNodeRequest) GetRpcID() string {
	if m != nil {
		return m.RpcID
	}
	return ""
}

func (m *FindNodeRequest) GetSender() *Contact {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *FindNodeRequest) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

type FindValueRequest struct {
	RpcID                string   `protobuf:"bytes,1,opt,name=rpcID,proto3" json:"rpcID,omitempty"`
	Sender               *Contact `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
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

func (m *FindValueRequest) GetRpcID() string {
	if m != nil {
		return m.RpcID
	}
	return ""
}

func (m *FindValueRequest) GetSender() *Contact {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *FindValueRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type FindNodeResponse struct {
	RpcID                string     `protobuf:"bytes,1,opt,name=rpcID,proto3" json:"rpcID,omitempty"`
	Sender               *Contact   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Contacts             []*Contact `protobuf:"bytes,3,rep,name=contacts,proto3" json:"contacts,omitempty"`
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

func (m *FindNodeResponse) GetRpcID() string {
	if m != nil {
		return m.RpcID
	}
	return ""
}

func (m *FindNodeResponse) GetSender() *Contact {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *FindNodeResponse) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type FindValueResponse struct {
	RpcID                string     `protobuf:"bytes,1,opt,name=rpcID,proto3" json:"rpcID,omitempty"`
	Sender               *Contact   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Value                string     `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Contacts             []*Contact `protobuf:"bytes,4,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
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

func (m *FindValueResponse) GetRpcID() string {
	if m != nil {
		return m.RpcID
	}
	return ""
}

func (m *FindValueResponse) GetSender() *Contact {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *FindValueResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *FindValueResponse) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
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
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x4b, 0xfb, 0x30,
	0x14, 0xa5, 0xbf, 0xfd, 0x6b, 0xef, 0xba, 0xfd, 0xb6, 0x30, 0xa4, 0xf4, 0xa9, 0x54, 0x91, 0x3d,
	0x45, 0x9c, 0x1f, 0x41, 0x11, 0x86, 0x20, 0x52, 0x41, 0x1f, 0x14, 0x21, 0x6b, 0xc2, 0x56, 0x9c,
	0x49, 0x6d, 0x32, 0xbf, 0x82, 0x1f, 0xca, 0x2f, 0x27, 0x6d, 0xd2, 0xd2, 0xba, 0xf9, 0x62, 0xdf,
	0x7a, 0x6e, 0x9b, 0x9c, 0x7b, 0xee, 0xb9, 0xa7, 0x30, 0x7e, 0x25, 0x94, 0xbd, 0x6d, 0x13, 0x82,
	0xd3, 0x4c, 0x28, 0x11, 0xae, 0x61, 0x78, 0x97, 0xf0, 0x75, 0xc4, 0xde, 0x77, 0x4c, 0x2a, 0x34,
	0x83, 0x5e, 0x96, 0xc6, 0xcb, 0x2b, 0xcf, 0x0a, 0xac, 0xb9, 0x13, 0x69, 0x80, 0x02, 0xe8, 0x4b,
	0xc6, 0x29, 0xcb, 0xbc, 0x7f, 0x81, 0x35, 0x1f, 0x2e, 0x6c, 0x7c, 0x29, 0xb8, 0x22, 0xb1, 0x8a,
	0x4c, 0x1d, 0x05, 0x30, 0xa4, 0x4c, 0xaa, 0x84, 0x13, 0x95, 0x08, 0xee, 0x75, 0x8a, 0xd3, 0xf5,
	0x52, 0xb8, 0x02, 0x57, 0x13, 0xc9, 0x54, 0x70, 0xc9, 0xfe, 0xcc, 0xe4, 0x83, 0x9d, 0x99, 0x3b,
	0x0c, 0x4d, 0x85, 0xc3, 0x67, 0x70, 0xef, 0x95, 0xc8, 0x58, 0x5b, 0x35, 0x33, 0xe8, 0x7d, 0x90,
	0xed, 0x4e, 0x13, 0xb8, 0x91, 0x06, 0xe1, 0x13, 0x8c, 0xcc, 0xed, 0x2d, 0x25, 0x20, 0xe8, 0x6e,
	0x88, 0xdc, 0x98, 0xf6, 0x8b, 0xe7, 0x90, 0xc1, 0xff, 0xeb, 0x84, 0xd3, 0x5b, 0x41, 0x5b, 0x77,
	0xef, 0x83, 0xad, 0x48, 0xb6, 0x66, 0x6a, 0x49, 0xcb, 0x09, 0x95, 0x38, 0x7c, 0x81, 0x49, 0x4e,
	0xf3, 0x90, 0x0b, 0x6a, 0xcb, 0x73, 0x48, 0x46, 0xaa, 0xef, 0xd7, 0x32, 0x5a, 0x8e, 0xe9, 0x04,
	0xec, 0x58, 0x97, 0xa4, 0xd7, 0x09, 0x3a, 0x8d, 0x6f, 0xaa, 0x37, 0xe1, 0xa7, 0x05, 0xd3, 0x9a,
	0xa4, 0x96, 0x9c, 0x0d, 0xe7, 0x1d, 0xe3, 0x7c, 0xa3, 0x93, 0xee, 0xaf, 0x9d, 0x3c, 0xc2, 0xc0,
	0x14, 0xd1, 0x11, 0xf4, 0xb9, 0xa0, 0x6c, 0x49, 0x0d, 0xbf, 0x41, 0xc8, 0x83, 0x01, 0xa1, 0x34,
	0x63, 0x52, 0x16, 0x1d, 0x38, 0x51, 0x09, 0x73, 0xd3, 0x68, 0x22, 0x15, 0xe1, 0x71, 0xb5, 0xd6,
	0x25, 0x5e, 0x7c, 0x59, 0x60, 0xdf, 0x98, 0xd8, 0xa2, 0x63, 0xe8, 0xe6, 0x39, 0x42, 0x2e, 0xae,
	0xe5, 0xd6, 0x1f, 0xe1, 0x46, 0xb8, 0x4e, 0xa1, 0x57, 0xac, 0x2a, 0x1a, 0xe1, 0x7a, 0x20, 0xfc,
	0x31, 0x6e, 0x6e, 0xf0, 0x19, 0xd8, 0xa5, 0x5d, 0x68, 0x82, 0x7f, 0x2c, 0xa0, 0x3f, 0xc5, 0x7b,
	0x5e, 0x9e, 0x83, 0x53, 0x0d, 0xfb, 0xc0, 0x09, 0x84, 0xf7, 0xac, 0x58, 0xf5, 0x8b, 0x1f, 0xcd,
	0xc5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x42, 0xd7, 0xba, 0x7a, 0x04, 0x00, 0x00,
}
