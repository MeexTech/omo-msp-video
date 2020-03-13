// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/video/build.proto

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

// 切割策略
type SplitStrategy int32

const (
	SplitStrategy_SIZE SplitStrategy = 0
	SplitStrategy_TIME SplitStrategy = 1
)

var SplitStrategy_name = map[int32]string{
	0: "SIZE",
	1: "TIME",
}

var SplitStrategy_value = map[string]int32{
	"SIZE": 0,
	"TIME": 1,
}

func (x SplitStrategy) String() string {
	return proto.EnumName(SplitStrategy_name, int32(x))
}

func (SplitStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_475d49d3c86cca58, []int{0}
}

// 构建回复
type BuildResponse struct {
	// 状态
	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 任务ID
	TaskID               string   `protobuf:"bytes,2,opt,name=taskID,proto3" json:"taskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildResponse) Reset()         { *m = BuildResponse{} }
func (m *BuildResponse) String() string { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()    {}
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_475d49d3c86cca58, []int{0}
}

func (m *BuildResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildResponse.Unmarshal(m, b)
}
func (m *BuildResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildResponse.Marshal(b, m, deterministic)
}
func (m *BuildResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildResponse.Merge(m, src)
}
func (m *BuildResponse) XXX_Size() int {
	return xxx_messageInfo_BuildResponse.Size(m)
}
func (m *BuildResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuildResponse proto.InternalMessageInfo

func (m *BuildResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BuildResponse) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

// 合并请求
type CombineRequest struct {
	// 视频文件地址
	// 网络地址使用http://或https://前缀
	// 内置存储使用file://前缀
	Video                []string `protobuf:"bytes,1,rep,name=video,proto3" json:"video,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CombineRequest) Reset()         { *m = CombineRequest{} }
func (m *CombineRequest) String() string { return proto.CompactTextString(m) }
func (*CombineRequest) ProtoMessage()    {}
func (*CombineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_475d49d3c86cca58, []int{1}
}

func (m *CombineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CombineRequest.Unmarshal(m, b)
}
func (m *CombineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CombineRequest.Marshal(b, m, deterministic)
}
func (m *CombineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CombineRequest.Merge(m, src)
}
func (m *CombineRequest) XXX_Size() int {
	return xxx_messageInfo_CombineRequest.Size(m)
}
func (m *CombineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CombineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CombineRequest proto.InternalMessageInfo

func (m *CombineRequest) GetVideo() []string {
	if m != nil {
		return m.Video
	}
	return nil
}

// 切割请求
type SplitRequest struct {
	// 视频文件地址
	// 网络地址使用http://或https://前缀
	// 内置存储使用file://前缀
	Video string `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	// 切割策略
	Strategy SplitStrategy `protobuf:"varint,2,opt,name=strategy,proto3,enum=video.SplitStrategy" json:"strategy,omitempty"`
	// 切割因子，strategy=SIZE时，此值的单位是字节，strategy=TIME时，此值的单位时秒
	Factor               int32    `protobuf:"varint,3,opt,name=factor,proto3" json:"factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SplitRequest) Reset()         { *m = SplitRequest{} }
func (m *SplitRequest) String() string { return proto.CompactTextString(m) }
func (*SplitRequest) ProtoMessage()    {}
func (*SplitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_475d49d3c86cca58, []int{2}
}

func (m *SplitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SplitRequest.Unmarshal(m, b)
}
func (m *SplitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SplitRequest.Marshal(b, m, deterministic)
}
func (m *SplitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplitRequest.Merge(m, src)
}
func (m *SplitRequest) XXX_Size() int {
	return xxx_messageInfo_SplitRequest.Size(m)
}
func (m *SplitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SplitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SplitRequest proto.InternalMessageInfo

func (m *SplitRequest) GetVideo() string {
	if m != nil {
		return m.Video
	}
	return ""
}

func (m *SplitRequest) GetStrategy() SplitStrategy {
	if m != nil {
		return m.Strategy
	}
	return SplitStrategy_SIZE
}

func (m *SplitRequest) GetFactor() int32 {
	if m != nil {
		return m.Factor
	}
	return 0
}

//图片
type Image struct {
	// 图片文件地址
	// 网络地址使用http://或https://前缀
	// 内置存储使用file://前缀
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// 图片在视频中的时间戳, 单位为毫秒
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 图片在视频中的持续时间, 单位为毫秒
	Duration             int64    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_475d49d3c86cca58, []int{3}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Image) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Image) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

//音频
type Audio struct {
	// 音频文件地址
	// 网络地址使用http://或https://前缀
	// 内置存储使用file://前缀
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// 音频在视频中的时间戳, 单位为毫秒
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 音频在视频中的持续时间, 单位为毫秒
	Duration             int64    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Audio) Reset()         { *m = Audio{} }
func (m *Audio) String() string { return proto.CompactTextString(m) }
func (*Audio) ProtoMessage()    {}
func (*Audio) Descriptor() ([]byte, []int) {
	return fileDescriptor_475d49d3c86cca58, []int{4}
}

func (m *Audio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Audio.Unmarshal(m, b)
}
func (m *Audio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Audio.Marshal(b, m, deterministic)
}
func (m *Audio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Audio.Merge(m, src)
}
func (m *Audio) XXX_Size() int {
	return xxx_messageInfo_Audio.Size(m)
}
func (m *Audio) XXX_DiscardUnknown() {
	xxx_messageInfo_Audio.DiscardUnknown(m)
}

var xxx_messageInfo_Audio proto.InternalMessageInfo

func (m *Audio) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Audio) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Audio) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

//笔迹
type Pen struct {
	// 图片文件地址
	// 网络地址使用http://或https://前缀
	// 内置存储使用file://前缀
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// 图片在视频中的时间戳, 单位为毫秒
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pen) Reset()         { *m = Pen{} }
func (m *Pen) String() string { return proto.CompactTextString(m) }
func (*Pen) ProtoMessage()    {}
func (*Pen) Descriptor() ([]byte, []int) {
	return fileDescriptor_475d49d3c86cca58, []int{5}
}

func (m *Pen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pen.Unmarshal(m, b)
}
func (m *Pen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pen.Marshal(b, m, deterministic)
}
func (m *Pen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pen.Merge(m, src)
}
func (m *Pen) XXX_Size() int {
	return xxx_messageInfo_Pen.Size(m)
}
func (m *Pen) XXX_DiscardUnknown() {
	xxx_messageInfo_Pen.DiscardUnknown(m)
}

var xxx_messageInfo_Pen proto.InternalMessageInfo

func (m *Pen) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Pen) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// 渲染请求
type RenderRequest struct {
	// 图片集
	Image []*Image `protobuf:"bytes,1,rep,name=image,proto3" json:"image,omitempty"`
	// 音频集
	Audio                []*Audio `protobuf:"bytes,2,rep,name=audio,proto3" json:"audio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenderRequest) Reset()         { *m = RenderRequest{} }
func (m *RenderRequest) String() string { return proto.CompactTextString(m) }
func (*RenderRequest) ProtoMessage()    {}
func (*RenderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_475d49d3c86cca58, []int{6}
}

func (m *RenderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenderRequest.Unmarshal(m, b)
}
func (m *RenderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenderRequest.Marshal(b, m, deterministic)
}
func (m *RenderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenderRequest.Merge(m, src)
}
func (m *RenderRequest) XXX_Size() int {
	return xxx_messageInfo_RenderRequest.Size(m)
}
func (m *RenderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenderRequest proto.InternalMessageInfo

func (m *RenderRequest) GetImage() []*Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *RenderRequest) GetAudio() []*Audio {
	if m != nil {
		return m.Audio
	}
	return nil
}

// 绘制请求
type DrawRequest struct {
	// 笔迹集
	Pen []*Pen `protobuf:"bytes,1,rep,name=pen,proto3" json:"pen,omitempty"`
	// 音频集
	Audio []*Audio `protobuf:"bytes,2,rep,name=audio,proto3" json:"audio,omitempty"`
	// 笔迹数据解析器
	Parser               string   `protobuf:"bytes,3,opt,name=parser,proto3" json:"parser,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DrawRequest) Reset()         { *m = DrawRequest{} }
func (m *DrawRequest) String() string { return proto.CompactTextString(m) }
func (*DrawRequest) ProtoMessage()    {}
func (*DrawRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_475d49d3c86cca58, []int{7}
}

func (m *DrawRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrawRequest.Unmarshal(m, b)
}
func (m *DrawRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrawRequest.Marshal(b, m, deterministic)
}
func (m *DrawRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrawRequest.Merge(m, src)
}
func (m *DrawRequest) XXX_Size() int {
	return xxx_messageInfo_DrawRequest.Size(m)
}
func (m *DrawRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DrawRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DrawRequest proto.InternalMessageInfo

func (m *DrawRequest) GetPen() []*Pen {
	if m != nil {
		return m.Pen
	}
	return nil
}

func (m *DrawRequest) GetAudio() []*Audio {
	if m != nil {
		return m.Audio
	}
	return nil
}

func (m *DrawRequest) GetParser() string {
	if m != nil {
		return m.Parser
	}
	return ""
}

func init() {
	proto.RegisterEnum("video.SplitStrategy", SplitStrategy_name, SplitStrategy_value)
	proto.RegisterType((*BuildResponse)(nil), "video.BuildResponse")
	proto.RegisterType((*CombineRequest)(nil), "video.CombineRequest")
	proto.RegisterType((*SplitRequest)(nil), "video.SplitRequest")
	proto.RegisterType((*Image)(nil), "video.Image")
	proto.RegisterType((*Audio)(nil), "video.Audio")
	proto.RegisterType((*Pen)(nil), "video.Pen")
	proto.RegisterType((*RenderRequest)(nil), "video.RenderRequest")
	proto.RegisterType((*DrawRequest)(nil), "video.DrawRequest")
}

func init() {
	proto.RegisterFile("proto/video/build.proto", fileDescriptor_475d49d3c86cca58)
}

var fileDescriptor_475d49d3c86cca58 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xdd, 0x6c, 0x9a, 0xd2, 0x4e, 0xb7, 0xab, 0xd5, 0xb0, 0x40, 0x54, 0xed, 0xa1, 0x0a, 0x02,
	0x55, 0x1c, 0xba, 0x28, 0x20, 0xe0, 0x0a, 0xec, 0x1e, 0x7a, 0x00, 0xad, 0x5c, 0x10, 0x12, 0x37,
	0x97, 0xcc, 0x16, 0x8b, 0xc6, 0x0e, 0xb6, 0x03, 0xe2, 0x7b, 0xf9, 0x11, 0x64, 0xc7, 0x29, 0x8d,
	0x44, 0x05, 0x48, 0x7b, 0xcb, 0xcc, 0xbc, 0xf7, 0xc6, 0x7e, 0x7e, 0x81, 0x7b, 0x95, 0x56, 0x56,
	0x9d, 0x7f, 0x13, 0x05, 0xa9, 0xf3, 0x55, 0x2d, 0x36, 0xc5, 0xdc, 0x77, 0x30, 0xf1, 0xad, 0x49,
	0xba, 0x3b, 0x37, 0x9f, 0xb9, 0xa6, 0x00, 0xc8, 0xde, 0xc2, 0xf8, 0x95, 0xc3, 0x33, 0x32, 0x95,
	0x92, 0x86, 0xf0, 0x01, 0xf4, 0x8d, 0xe5, 0xb6, 0x36, 0x69, 0x34, 0x8d, 0x66, 0xa3, 0x7c, 0x3c,
	0xf7, 0xac, 0xf9, 0xd2, 0x37, 0x59, 0x18, 0xe2, 0x5d, 0xe8, 0x5b, 0x6e, 0xbe, 0x2c, 0x2e, 0xd2,
	0xc3, 0x69, 0x34, 0x1b, 0xb2, 0x50, 0x65, 0x0f, 0xe1, 0xf8, 0xb5, 0x2a, 0x57, 0x42, 0x12, 0xa3,
	0xaf, 0x35, 0x19, 0x8b, 0xa7, 0xd0, 0x1c, 0x22, 0x8d, 0xa6, 0xf1, 0x6c, 0xc8, 0x9a, 0x22, 0x93,
	0x70, 0xb4, 0xac, 0x36, 0xc2, 0xfe, 0x01, 0x15, 0x6d, 0x51, 0xf8, 0x18, 0x06, 0xc6, 0x6a, 0x6e,
	0x69, 0xfd, 0xc3, 0xef, 0x39, 0xce, 0x4f, 0xdb, 0xe3, 0x38, 0xf2, 0x32, 0xcc, 0xd8, 0x16, 0xe5,
	0xce, 0x75, 0xcd, 0x3f, 0x59, 0xa5, 0xd3, 0x78, 0x1a, 0xcd, 0x12, 0x16, 0xaa, 0xec, 0x3d, 0x24,
	0x8b, 0x92, 0xaf, 0x09, 0x11, 0x7a, 0xd7, 0x62, 0x43, 0x61, 0x8f, 0xff, 0xc6, 0x33, 0x18, 0x5a,
	0x51, 0x92, 0xb1, 0xbc, 0xac, 0xfc, 0x9e, 0x98, 0xfd, 0x6e, 0xe0, 0x04, 0x06, 0x45, 0xad, 0xb9,
	0x15, 0x4a, 0x7a, 0xd1, 0x98, 0x6d, 0x6b, 0x27, 0xfb, 0xb2, 0x2e, 0x84, 0xba, 0x61, 0xd9, 0xe7,
	0x10, 0x5f, 0x91, 0xfc, 0x7f, 0xd1, 0xec, 0x03, 0x8c, 0x19, 0xc9, 0x82, 0x74, 0xeb, 0x6b, 0x06,
	0x89, 0x70, 0xf7, 0xf6, 0xee, 0x8f, 0xf2, 0xa3, 0x60, 0x9f, 0xf7, 0x82, 0x35, 0x23, 0x87, 0xe1,
	0xee, 0x12, 0xe9, 0x61, 0x07, 0xe3, 0x2f, 0xc6, 0x9a, 0x51, 0xb6, 0x86, 0xd1, 0x85, 0xe6, 0xdf,
	0x5b, 0xd9, 0x33, 0x88, 0x2b, 0x92, 0x41, 0x14, 0x02, 0xe1, 0x8a, 0x24, 0x73, 0xed, 0x7f, 0x11,
	0x74, 0x0f, 0x55, 0x71, 0x6d, 0xa8, 0x79, 0xa8, 0x21, 0x0b, 0xd5, 0xa3, 0xfb, 0x30, 0xee, 0xbc,
	0x2d, 0x0e, 0xa0, 0xb7, 0x5c, 0x7c, 0xbc, 0x3c, 0x39, 0x70, 0x5f, 0xef, 0x16, 0x6f, 0x2e, 0x4f,
	0xa2, 0xfc, 0x67, 0x04, 0x89, 0x8f, 0x2d, 0xbe, 0x80, 0x5b, 0x21, 0x6f, 0x78, 0x27, 0xac, 0xe9,
	0xe6, 0x6f, 0xd2, 0x26, 0xa6, 0x13, 0xf3, 0xec, 0x00, 0x9f, 0x42, 0xe2, 0x17, 0xe1, 0xed, 0xdd,
	0x48, 0xfd, 0x8d, 0xf5, 0x0c, 0xfa, 0x8d, 0xc1, 0xd8, 0x22, 0x3a, 0x7e, 0xef, 0xe5, 0xe5, 0xd0,
	0x73, 0xfe, 0x21, 0x86, 0xf9, 0x8e, 0x99, 0xfb, 0x38, 0xab, 0xbe, 0xff, 0x45, 0x9f, 0xfc, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xd8, 0xc8, 0xbe, 0x58, 0xde, 0x03, 0x00, 0x00,
}
