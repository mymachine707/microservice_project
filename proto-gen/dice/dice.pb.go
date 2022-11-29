// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/dice.proto

package dice

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

// The request message containing the user's name.
type RollDiceRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RollDiceRequest) Reset()         { *m = RollDiceRequest{} }
func (m *RollDiceRequest) String() string { return proto.CompactTextString(m) }
func (*RollDiceRequest) ProtoMessage()    {}
func (*RollDiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b299c7e2607893, []int{0}
}

func (m *RollDiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollDiceRequest.Unmarshal(m, b)
}
func (m *RollDiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollDiceRequest.Marshal(b, m, deterministic)
}
func (m *RollDiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollDiceRequest.Merge(m, src)
}
func (m *RollDiceRequest) XXX_Size() int {
	return xxx_messageInfo_RollDiceRequest.Size(m)
}
func (m *RollDiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RollDiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RollDiceRequest proto.InternalMessageInfo

func (m *RollDiceRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

// The response message containing the greetings
type RollDiceResponse struct {
	Dice                 []int32  `protobuf:"varint,1,rep,packed,name=dice,proto3" json:"dice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RollDiceResponse) Reset()         { *m = RollDiceResponse{} }
func (m *RollDiceResponse) String() string { return proto.CompactTextString(m) }
func (*RollDiceResponse) ProtoMessage()    {}
func (*RollDiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b299c7e2607893, []int{1}
}

func (m *RollDiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollDiceResponse.Unmarshal(m, b)
}
func (m *RollDiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollDiceResponse.Marshal(b, m, deterministic)
}
func (m *RollDiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollDiceResponse.Merge(m, src)
}
func (m *RollDiceResponse) XXX_Size() int {
	return xxx_messageInfo_RollDiceResponse.Size(m)
}
func (m *RollDiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RollDiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RollDiceResponse proto.InternalMessageInfo

func (m *RollDiceResponse) GetDice() []int32 {
	if m != nil {
		return m.Dice
	}
	return nil
}

func init() {
	proto.RegisterType((*RollDiceRequest)(nil), "RollDiceRequest")
	proto.RegisterType((*RollDiceResponse)(nil), "RollDiceResponse")
}

func init() { proto.RegisterFile("protos/dice.proto", fileDescriptor_d6b299c7e2607893) }

var fileDescriptor_d6b299c7e2607893 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0xc9, 0x4c, 0x4e, 0xd5, 0x03, 0xb3, 0x95, 0x94, 0xb9, 0xf8, 0x83, 0xf2,
	0x73, 0x72, 0x5c, 0x32, 0x93, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8,
	0x98, 0xf3, 0x4a, 0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x25, 0x35, 0x2e,
	0x01, 0x84, 0xa2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x21, 0x2e, 0x16, 0x90, 0x31, 0x12,
	0x8c, 0x0a, 0xcc, 0x1a, 0xac, 0x41, 0x60, 0xb6, 0x91, 0x2d, 0x17, 0x47, 0x48, 0x69, 0x49, 0x7e,
	0x51, 0x66, 0x62, 0x8e, 0x90, 0x21, 0x17, 0x07, 0x4c, 0x8f, 0x90, 0x80, 0x1e, 0x9a, 0x1d, 0x52,
	0x82, 0x7a, 0xe8, 0x06, 0x2a, 0x31, 0x38, 0x71, 0x44, 0xb1, 0xe9, 0x81, 0xdd, 0x96, 0xc4, 0x06,
	0x76, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x58, 0x1b, 0xc6, 0xb0, 0xb1, 0x00, 0x00, 0x00,
}
