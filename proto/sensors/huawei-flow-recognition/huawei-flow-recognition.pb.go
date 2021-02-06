// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sensors/huawei-flow-recognition/huawei-flow-recognition.proto

package huawei_flow_recognition

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

type FlowRecognition struct {
	Streaminfos          *FlowRecognition_Streaminfos `protobuf:"bytes,1,opt,name=streaminfos,proto3" json:"streaminfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *FlowRecognition) Reset()         { *m = FlowRecognition{} }
func (m *FlowRecognition) String() string { return proto.CompactTextString(m) }
func (*FlowRecognition) ProtoMessage()    {}
func (*FlowRecognition) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fb3299627ae3bd7, []int{0}
}

func (m *FlowRecognition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRecognition.Unmarshal(m, b)
}
func (m *FlowRecognition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRecognition.Marshal(b, m, deterministic)
}
func (m *FlowRecognition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRecognition.Merge(m, src)
}
func (m *FlowRecognition) XXX_Size() int {
	return xxx_messageInfo_FlowRecognition.Size(m)
}
func (m *FlowRecognition) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRecognition.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRecognition proto.InternalMessageInfo

func (m *FlowRecognition) GetStreaminfos() *FlowRecognition_Streaminfos {
	if m != nil {
		return m.Streaminfos
	}
	return nil
}

type FlowRecognition_Streaminfos struct {
	Streaminfo           []*FlowRecognition_Streaminfos_Streaminfo `protobuf:"bytes,1,rep,name=streaminfo,proto3" json:"streaminfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *FlowRecognition_Streaminfos) Reset()         { *m = FlowRecognition_Streaminfos{} }
func (m *FlowRecognition_Streaminfos) String() string { return proto.CompactTextString(m) }
func (*FlowRecognition_Streaminfos) ProtoMessage()    {}
func (*FlowRecognition_Streaminfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fb3299627ae3bd7, []int{0, 0}
}

func (m *FlowRecognition_Streaminfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRecognition_Streaminfos.Unmarshal(m, b)
}
func (m *FlowRecognition_Streaminfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRecognition_Streaminfos.Marshal(b, m, deterministic)
}
func (m *FlowRecognition_Streaminfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRecognition_Streaminfos.Merge(m, src)
}
func (m *FlowRecognition_Streaminfos) XXX_Size() int {
	return xxx_messageInfo_FlowRecognition_Streaminfos.Size(m)
}
func (m *FlowRecognition_Streaminfos) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRecognition_Streaminfos.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRecognition_Streaminfos proto.InternalMessageInfo

func (m *FlowRecognition_Streaminfos) GetStreaminfo() []*FlowRecognition_Streaminfos_Streaminfo {
	if m != nil {
		return m.Streaminfo
	}
	return nil
}

type FlowRecognition_Streaminfos_Streaminfo struct {
	SrcMac               string   `protobuf:"bytes,1,opt,name=src_mac,json=src-mac,proto3" json:"src_mac,omitempty"`
	DstMac               string   `protobuf:"bytes,2,opt,name=dst_mac,json=dst-mac,proto3" json:"dst_mac,omitempty"`
	SrcIpAddr            string   `protobuf:"bytes,3,opt,name=src_ip_addr,json=src-ip-addr,proto3" json:"src_ip_addr,omitempty"`
	DstIpAddr            string   `protobuf:"bytes,4,opt,name=dst_ip_addr,json=dst-ip-addr,proto3" json:"dst_ip_addr,omitempty"`
	SrcPort              uint32   `protobuf:"varint,5,opt,name=src_port,json=src-port,proto3" json:"src_port,omitempty"`
	DstPort              uint32   `protobuf:"varint,6,opt,name=dst_port,json=dst-port,proto3" json:"dst_port,omitempty"`
	Protocol             uint32   `protobuf:"varint,7,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Direction            uint32   `protobuf:"varint,8,opt,name=direction,proto3" json:"direction,omitempty"`
	IfName               string   `protobuf:"bytes,9,opt,name=if_name,json=if-name,proto3" json:"if_name,omitempty"`
	TimeStampSec         uint32   `protobuf:"varint,10,opt,name=time_stamp_sec,json=time-stamp-sec,proto3" json:"time_stamp_sec,omitempty"`
	TimeStampNsec        uint32   `protobuf:"varint,11,opt,name=time_stamp_nsec,json=time-stamp-nsec,proto3" json:"time_stamp_nsec,omitempty"`
	PacketNum            uint64   `protobuf:"varint,12,opt,name=packet_num,json=packet-num,proto3" json:"packet_num,omitempty"`
	BytesNum             uint64   `protobuf:"varint,13,opt,name=bytes_num,json=bytes-num,proto3" json:"bytes_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowRecognition_Streaminfos_Streaminfo) Reset() {
	*m = FlowRecognition_Streaminfos_Streaminfo{}
}
func (m *FlowRecognition_Streaminfos_Streaminfo) String() string { return proto.CompactTextString(m) }
func (*FlowRecognition_Streaminfos_Streaminfo) ProtoMessage()    {}
func (*FlowRecognition_Streaminfos_Streaminfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fb3299627ae3bd7, []int{0, 0, 0}
}

func (m *FlowRecognition_Streaminfos_Streaminfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRecognition_Streaminfos_Streaminfo.Unmarshal(m, b)
}
func (m *FlowRecognition_Streaminfos_Streaminfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRecognition_Streaminfos_Streaminfo.Marshal(b, m, deterministic)
}
func (m *FlowRecognition_Streaminfos_Streaminfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRecognition_Streaminfos_Streaminfo.Merge(m, src)
}
func (m *FlowRecognition_Streaminfos_Streaminfo) XXX_Size() int {
	return xxx_messageInfo_FlowRecognition_Streaminfos_Streaminfo.Size(m)
}
func (m *FlowRecognition_Streaminfos_Streaminfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRecognition_Streaminfos_Streaminfo.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRecognition_Streaminfos_Streaminfo proto.InternalMessageInfo

func (m *FlowRecognition_Streaminfos_Streaminfo) GetSrcMac() string {
	if m != nil {
		return m.SrcMac
	}
	return ""
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetDstMac() string {
	if m != nil {
		return m.DstMac
	}
	return ""
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetSrcIpAddr() string {
	if m != nil {
		return m.SrcIpAddr
	}
	return ""
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetDstIpAddr() string {
	if m != nil {
		return m.DstIpAddr
	}
	return ""
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetIfName() string {
	if m != nil {
		return m.IfName
	}
	return ""
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetTimeStampSec() uint32 {
	if m != nil {
		return m.TimeStampSec
	}
	return 0
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetTimeStampNsec() uint32 {
	if m != nil {
		return m.TimeStampNsec
	}
	return 0
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetPacketNum() uint64 {
	if m != nil {
		return m.PacketNum
	}
	return 0
}

func (m *FlowRecognition_Streaminfos_Streaminfo) GetBytesNum() uint64 {
	if m != nil {
		return m.BytesNum
	}
	return 0
}

func init() {
	proto.RegisterType((*FlowRecognition)(nil), "huawei_flow_recognition.FlowRecognition")
	proto.RegisterType((*FlowRecognition_Streaminfos)(nil), "huawei_flow_recognition.FlowRecognition.Streaminfos")
	proto.RegisterType((*FlowRecognition_Streaminfos_Streaminfo)(nil), "huawei_flow_recognition.FlowRecognition.Streaminfos.Streaminfo")
}

func init() {
	proto.RegisterFile("proto/sensors/huawei-flow-recognition/huawei-flow-recognition.proto", fileDescriptor_8fb3299627ae3bd7)
}

var fileDescriptor_8fb3299627ae3bd7 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0x71, 0xe3, 0xe6, 0xcf, 0xb8, 0x69, 0x40, 0x97, 0x8a, 0x50, 0x8a, 0xe9, 0xa1, 0xf8,
	0x22, 0x07, 0xd2, 0xde, 0x7b, 0x28, 0xf4, 0x01, 0x5c, 0xe8, 0x55, 0x28, 0xb2, 0xdc, 0x8a, 0x8d,
	0x24, 0x23, 0x29, 0x84, 0x7d, 0x8c, 0x7d, 0x83, 0x7d, 0xd1, 0x85, 0x65, 0xe4, 0x35, 0x36, 0x0b,
	0xb9, 0xec, 0x2d, 0xf3, 0xfb, 0x7d, 0xdf, 0xc4, 0x63, 0xc3, 0xaf, 0xde, 0xbb, 0xe8, 0x0e, 0x41,
	0xd9, 0xe0, 0x7c, 0x38, 0xfc, 0xbf, 0x88, 0xab, 0xd2, 0xac, 0x3b, 0xbb, 0x2b, 0xf3, 0x4a, 0xba,
	0x7f, 0x56, 0x47, 0xed, 0xec, 0x2d, 0x5e, 0xa7, 0x36, 0xf9, 0x34, 0x68, 0x8e, 0x9a, 0xcf, 0xf4,
	0xd7, 0xa7, 0x1c, 0x76, 0xbf, 0xcf, 0xee, 0xda, 0x4c, 0x8c, 0xfc, 0x85, 0x22, 0x44, 0xaf, 0x84,
	0xd1, 0xb6, 0x73, 0x81, 0x66, 0x65, 0x56, 0x15, 0xc7, 0x1f, 0xf5, 0x8d, 0x15, 0xf5, 0xab, 0x7a,
	0xfd, 0x67, 0xea, 0x36, 0xf3, 0x45, 0xfb, 0x87, 0x1c, 0x8a, 0x99, 0x24, 0x1c, 0x60, 0xd2, 0x34,
	0x2b, 0x17, 0x55, 0x71, 0xfc, 0xf9, 0x96, 0xbf, 0x99, 0xfd, 0x6e, 0x66, 0x2b, 0xf7, 0x8f, 0x0b,
	0x80, 0x49, 0x11, 0x0a, 0xab, 0xe0, 0x25, 0x37, 0x42, 0xa6, 0x9b, 0x36, 0x0d, 0x8e, 0xcc, 0x08,
	0x89, 0xa6, 0x0d, 0x31, 0x99, 0x77, 0x83, 0x69, 0x43, 0x4c, 0xa6, 0x84, 0x02, 0x3b, 0xba, 0xe7,
	0xa2, 0x6d, 0x3d, 0x5d, 0x24, 0x8b, 0x88, 0xe9, 0x9e, 0x21, 0xc2, 0x04, 0x76, 0xc7, 0x44, 0x3e,
	0x24, 0xb0, 0x3f, 0x26, 0xf6, 0xb0, 0xc6, 0x1d, 0xbd, 0xf3, 0x91, 0xbe, 0x2f, 0xb3, 0x6a, 0xdb,
	0xe0, 0xcc, 0x70, 0x46, 0x87, 0xed, 0xe4, 0x96, 0x83, 0xc3, 0xea, 0xe8, 0xd2, 0xd7, 0x93, 0xee,
	0x4c, 0x57, 0x83, 0x1b, 0x67, 0xf2, 0x19, 0x36, 0xad, 0xf6, 0x4a, 0xe2, 0xab, 0xa0, 0xeb, 0x24,
	0x27, 0x80, 0xf7, 0xe8, 0x8e, 0x5b, 0x61, 0x14, 0xdd, 0x0c, 0xf7, 0xe8, 0x8e, 0xe1, 0x48, 0xbe,
	0xc1, 0xc7, 0xa8, 0x8d, 0xe2, 0x21, 0x0a, 0xd3, 0xf3, 0xa0, 0x24, 0x85, 0x54, 0x4e, 0x94, 0x25,
	0xca, 0x82, 0x92, 0xa4, 0x82, 0xdd, 0x2c, 0x67, 0x31, 0x58, 0xa4, 0xe0, 0x6e, 0x16, 0x44, 0x4c,
	0xbe, 0x00, 0xf4, 0x42, 0xde, 0xa9, 0xc8, 0xed, 0xc5, 0xd0, 0x0f, 0x65, 0x56, 0xe5, 0xcd, 0x0b,
	0x61, 0xf6, 0x62, 0xf0, 0x49, 0x4f, 0xf7, 0x51, 0x85, 0xa4, 0xb7, 0x49, 0x0f, 0x00, 0xed, 0x69,
	0x99, 0x2e, 0xfa, 0xfe, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x41, 0x33, 0x50, 0x8b, 0xe6, 0x02, 0x00,
	0x00,
}
