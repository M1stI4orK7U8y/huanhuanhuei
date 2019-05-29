// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/huanhuancore/api/grpc/huanhuan.proto

package huanhuan

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/reply"
	model "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/model"
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

type HuanHuanRequest struct {
	From                 model.TokenType `protobuf:"varint,1,opt,name=from,proto3,enum=model.token.TokenType" json:"from,omitempty"`
	FromTxid             string          `protobuf:"bytes,2,opt,name=fromTxid,proto3" json:"fromTxid,omitempty"`
	To                   model.TokenType `protobuf:"varint,3,opt,name=to,proto3,enum=model.token.TokenType" json:"to,omitempty"`
	Receiver             string          `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HuanHuanRequest) Reset()         { *m = HuanHuanRequest{} }
func (m *HuanHuanRequest) String() string { return proto.CompactTextString(m) }
func (*HuanHuanRequest) ProtoMessage()    {}
func (*HuanHuanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_89bd64f88c09226b, []int{0}
}

func (m *HuanHuanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HuanHuanRequest.Unmarshal(m, b)
}
func (m *HuanHuanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HuanHuanRequest.Marshal(b, m, deterministic)
}
func (m *HuanHuanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HuanHuanRequest.Merge(m, src)
}
func (m *HuanHuanRequest) XXX_Size() int {
	return xxx_messageInfo_HuanHuanRequest.Size(m)
}
func (m *HuanHuanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HuanHuanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HuanHuanRequest proto.InternalMessageInfo

func (m *HuanHuanRequest) GetFrom() model.TokenType {
	if m != nil {
		return m.From
	}
	return model.TokenType_BTC
}

func (m *HuanHuanRequest) GetFromTxid() string {
	if m != nil {
		return m.FromTxid
	}
	return ""
}

func (m *HuanHuanRequest) GetTo() model.TokenType {
	if m != nil {
		return m.To
	}
	return model.TokenType_BTC
}

func (m *HuanHuanRequest) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func init() {
	proto.RegisterType((*HuanHuanRequest)(nil), "api.grpc.huanhuan.HuanHuanRequest")
}

func init() {
	proto.RegisterFile("src/huanhuancore/api/grpc/huanhuan.proto", fileDescriptor_89bd64f88c09226b)
}

var fileDescriptor_89bd64f88c09226b = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0xb5, 0x75, 0x91, 0x3a, 0x07, 0xc5, 0x1c, 0xa4, 0xf4, 0xb4, 0x14, 0x91, 0xe2, 0x21, 0x81,
	0xf5, 0x0f, 0x44, 0x64, 0xcf, 0xa1, 0x27, 0x6f, 0xd9, 0x76, 0xd4, 0xe2, 0xee, 0x4e, 0x4c, 0x53,
	0x71, 0x3f, 0xc6, 0x7f, 0x95, 0x69, 0x1a, 0x05, 0x0b, 0x7b, 0x98, 0x17, 0x26, 0x33, 0xef, 0xcd,
	0xcc, 0x83, 0xaa, 0x77, 0x8d, 0x7a, 0x1b, 0xcc, 0x9e, 0xa3, 0x21, 0x87, 0xca, 0xd8, 0x4e, 0xbd,
	0x3a, 0xfb, 0xf7, 0x2b, 0xad, 0x23, 0x4f, 0xe2, 0xca, 0xd8, 0x4e, 0x72, 0x41, 0xc6, 0x42, 0x71,
	0xc3, 0xe4, 0xd6, 0x78, 0xb3, 0x31, 0x3d, 0xaa, 0x1d, 0xb5, 0xb8, 0x55, 0x0e, 0xed, 0xf6, 0x10,
	0x30, 0x10, 0x8b, 0x72, 0x36, 0x22, 0x74, 0x7a, 0x7a, 0xc7, 0x49, 0xbc, 0xfc, 0x4e, 0xe0, 0x72,
	0x3d, 0x98, 0x3d, 0x87, 0xc6, 0x8f, 0x01, 0x7b, 0x2f, 0xee, 0x60, 0xf1, 0xe2, 0x68, 0x97, 0x27,
	0xcb, 0xa4, 0xba, 0x58, 0x5d, 0xcb, 0x91, 0x25, 0x03, 0xab, 0x66, 0xac, 0x0f, 0x16, 0xf5, 0xd8,
	0x23, 0x0a, 0xc8, 0xf8, 0xad, 0xbf, 0xba, 0x36, 0x4f, 0x97, 0x49, 0x75, 0xae, 0x7f, 0x73, 0x71,
	0x0b, 0xa9, 0xa7, 0xfc, 0xf4, 0xa8, 0x4a, 0xea, 0x89, 0x35, 0x1c, 0x36, 0xd8, 0x7d, 0xa2, 0xcb,
	0x17, 0x41, 0x23, 0xe6, 0x2b, 0x0d, 0xd9, 0x7a, 0xba, 0x40, 0x3c, 0x01, 0x3c, 0x52, 0x5c, 0x56,
	0x94, 0x72, 0xe6, 0x8b, 0xfc, 0x77, 0x49, 0x21, 0xa6, 0xa9, 0xc1, 0x15, 0xcd, 0x58, 0x9e, 0x3c,
	0xc0, 0x73, 0x16, 0x19, 0x9b, 0xb3, 0xd1, 0x86, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd,
	0xb8, 0x40, 0xe4, 0x8f, 0x01, 0x00, 0x00,
}
