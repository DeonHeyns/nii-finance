// Code generated by protoc-gen-go.
// source: github.com/nii236/nii-forex/tickRecorder/proto/tickRecorder.proto
// DO NOT EDIT!

/*
Package tickRecorder is a generated protocol buffer package.

It is generated from these files:
	github.com/nii236/nii-forex/tickRecorder/proto/tickRecorder.proto

It has these top-level messages:
	Tick
*/
package tickRecorder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Tick struct {
	Time int64  `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Bid  int32  `protobuf:"varint,2,opt,name=bid" json:"bid,omitempty"`
	Ask  int32  `protobuf:"varint,3,opt,name=ask" json:"ask,omitempty"`
	Last int32  `protobuf:"varint,4,opt,name=last" json:"last,omitempty"`
	Pair string `protobuf:"bytes,5,opt,name=pair" json:"pair,omitempty"`
}

func (m *Tick) Reset()                    { *m = Tick{} }
func (m *Tick) String() string            { return proto.CompactTextString(m) }
func (*Tick) ProtoMessage()               {}
func (*Tick) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Tick)(nil), "Tick")
}

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcb, 0xcc, 0x34, 0x32, 0x36, 0x03, 0x51, 0xba,
	0x69, 0xf9, 0x45, 0xa9, 0x15, 0xfa, 0x25, 0x99, 0xc9, 0xd9, 0x41, 0xa9, 0xc9, 0xf9, 0x45, 0x29,
	0xa9, 0x45, 0xfa, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x28, 0x42, 0x7a, 0x60, 0x21, 0x25, 0x2f, 0x2e,
	0x96, 0x10, 0xa0, 0xa8, 0x10, 0x0f, 0x17, 0x4b, 0x49, 0x66, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0xb3, 0x10, 0x37, 0x17, 0x73, 0x52, 0x66, 0x8a, 0x04, 0x13, 0x90, 0xc3, 0x0a, 0xe2, 0x24,
	0x16, 0x67, 0x4b, 0x30, 0x83, 0x39, 0x40, 0x75, 0x39, 0x89, 0xc5, 0x25, 0x12, 0x2c, 0x30, 0x5e,
	0x41, 0x62, 0x66, 0x91, 0x04, 0x2b, 0x90, 0xc7, 0x99, 0xc4, 0x06, 0x36, 0xd2, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x13, 0xf8, 0x69, 0xb5, 0x97, 0x00, 0x00, 0x00,
}