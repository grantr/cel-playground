// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ce.proto

package dev_knative

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

type CloudEvent struct {
	Specversion          string   `protobuf:"bytes,1,opt,name=specversion,proto3" json:"specversion,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Time                 string   `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudEvent) Reset()         { *m = CloudEvent{} }
func (m *CloudEvent) String() string { return proto.CompactTextString(m) }
func (*CloudEvent) ProtoMessage()    {}
func (*CloudEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_85884d17b902dd0c, []int{0}
}

func (m *CloudEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudEvent.Unmarshal(m, b)
}
func (m *CloudEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudEvent.Marshal(b, m, deterministic)
}
func (m *CloudEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudEvent.Merge(m, src)
}
func (m *CloudEvent) XXX_Size() int {
	return xxx_messageInfo_CloudEvent.Size(m)
}
func (m *CloudEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CloudEvent proto.InternalMessageInfo

func (m *CloudEvent) GetSpecversion() string {
	if m != nil {
		return m.Specversion
	}
	return ""
}

func (m *CloudEvent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CloudEvent) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *CloudEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CloudEvent) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func init() {
	proto.RegisterType((*CloudEvent)(nil), "dev.knative.CloudEvent")
}

func init() { proto.RegisterFile("ce.proto", fileDescriptor_85884d17b902dd0c) }

var fileDescriptor_85884d17b902dd0c = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x48, 0x4e, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x49, 0x2d, 0xd3, 0xcb, 0xce, 0x4b, 0x2c, 0xc9, 0x2c,
	0x4b, 0x55, 0xaa, 0xe3, 0xe2, 0x72, 0xce, 0xc9, 0x2f, 0x4d, 0x71, 0x2d, 0x4b, 0xcd, 0x2b, 0x11,
	0x52, 0xe0, 0xe2, 0x2e, 0x2e, 0x48, 0x4d, 0x2e, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x93, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x16, 0x12, 0x12, 0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95,
	0x60, 0x02, 0x4b, 0x81, 0xd9, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0xf9, 0xa5, 0x45, 0xc9, 0xa9, 0x12,
	0xcc, 0x60, 0x51, 0x28, 0x4f, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x05, 0x2c, 0xc6, 0x94,
	0x99, 0x02, 0xd6, 0x9b, 0x99, 0x9b, 0x2a, 0xc1, 0x0a, 0xd5, 0x9b, 0x99, 0x9b, 0x9a, 0xc4, 0x06,
	0x76, 0x93, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xe5, 0xb8, 0xd4, 0x9f, 0x00, 0x00, 0x00,
}
