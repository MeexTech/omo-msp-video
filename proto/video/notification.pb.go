// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/video/notification.proto

package video

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

// 通知
type Notification struct {
	// 任务ID
	TaskID string `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	// 状态
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c471ff4906b8dcb, []int{0}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *Notification) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Notification)(nil), "video.Notification")
}

func init() {
	proto.RegisterFile("proto/video/notification.proto", fileDescriptor_9c471ff4906b8dcb)
}

var fileDescriptor_9c471ff4906b8dcb = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xcb, 0x4c, 0x49, 0xcd, 0xd7, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e,
	0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x4b, 0x08, 0xb1, 0x82, 0x65, 0x94, 0xec, 0xb8, 0x78, 0xfc,
	0x90, 0x24, 0x85, 0xc4, 0xb8, 0xd8, 0x4a, 0x12, 0x8b, 0xb3, 0x3d, 0x5d, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xa0, 0x3c, 0x90, 0x78, 0x71, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x04, 0x13, 0x44,
	0x1c, 0xc2, 0x4b, 0x62, 0x03, 0x9b, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x43, 0x06, 0x67,
	0xc2, 0x6f, 0x00, 0x00, 0x00,
}