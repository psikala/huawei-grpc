// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sensors/huawei-pic/huawei-pic.proto

package huawei_pic

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

type Pic_PortStatistics_Statistic_PortType int32

const (
	Pic_PortStatistics_Statistic_PortType_ATM       Pic_PortStatistics_Statistic_PortType = 0
	Pic_PortStatistics_Statistic_PortType_C37       Pic_PortStatistics_Statistic_PortType = 1
	Pic_PortStatistics_Statistic_PortType_CODIR     Pic_PortStatistics_Statistic_PortType = 2
	Pic_PortStatistics_Statistic_PortType_CPOS      Pic_PortStatistics_Statistic_PortType = 3
	Pic_PortStatistics_Statistic_PortType_E1        Pic_PortStatistics_Statistic_PortType = 4
	Pic_PortStatistics_Statistic_PortType_E3        Pic_PortStatistics_Statistic_PortType = 5
	Pic_PortStatistics_Statistic_PortType_ECMSERIAL Pic_PortStatistics_Statistic_PortType = 6
	Pic_PortStatistics_Statistic_PortType_ETH       Pic_PortStatistics_Statistic_PortType = 7
	Pic_PortStatistics_Statistic_PortType_FLEXEPHY  Pic_PortStatistics_Statistic_PortType = 8
	Pic_PortStatistics_Statistic_PortType_FLEXE     Pic_PortStatistics_Statistic_PortType = 9
	Pic_PortStatistics_Statistic_PortType_OTN       Pic_PortStatistics_Statistic_PortType = 10
	Pic_PortStatistics_Statistic_PortType_POS       Pic_PortStatistics_Statistic_PortType = 11
	Pic_PortStatistics_Statistic_PortType_VOICE     Pic_PortStatistics_Statistic_PortType = 12
)

var Pic_PortStatistics_Statistic_PortType_name = map[int32]string{
	0:  "PortType_ATM",
	1:  "PortType_C37",
	2:  "PortType_CODIR",
	3:  "PortType_CPOS",
	4:  "PortType_E1",
	5:  "PortType_E3",
	6:  "PortType_ECMSERIAL",
	7:  "PortType_ETH",
	8:  "PortType_FLEXEPHY",
	9:  "PortType_FLEXE",
	10: "PortType_OTN",
	11: "PortType_POS",
	12: "PortType_VOICE",
}

var Pic_PortStatistics_Statistic_PortType_value = map[string]int32{
	"PortType_ATM":       0,
	"PortType_C37":       1,
	"PortType_CODIR":     2,
	"PortType_CPOS":      3,
	"PortType_E1":        4,
	"PortType_E3":        5,
	"PortType_ECMSERIAL": 6,
	"PortType_ETH":       7,
	"PortType_FLEXEPHY":  8,
	"PortType_FLEXE":     9,
	"PortType_OTN":       10,
	"PortType_POS":       11,
	"PortType_VOICE":     12,
}

func (x Pic_PortStatistics_Statistic_PortType) String() string {
	return proto.EnumName(Pic_PortStatistics_Statistic_PortType_name, int32(x))
}

func (Pic_PortStatistics_Statistic_PortType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_322574f333bd178c, []int{0, 0, 0, 0}
}

type Pic struct {
	PortStatistics       *Pic_PortStatistics `protobuf:"bytes,1,opt,name=port_statistics,json=port-statistics,proto3" json:"port_statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Pic) Reset()         { *m = Pic{} }
func (m *Pic) String() string { return proto.CompactTextString(m) }
func (*Pic) ProtoMessage()    {}
func (*Pic) Descriptor() ([]byte, []int) {
	return fileDescriptor_322574f333bd178c, []int{0}
}

func (m *Pic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pic.Unmarshal(m, b)
}
func (m *Pic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pic.Marshal(b, m, deterministic)
}
func (m *Pic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pic.Merge(m, src)
}
func (m *Pic) XXX_Size() int {
	return xxx_messageInfo_Pic.Size(m)
}
func (m *Pic) XXX_DiscardUnknown() {
	xxx_messageInfo_Pic.DiscardUnknown(m)
}

var xxx_messageInfo_Pic proto.InternalMessageInfo

func (m *Pic) GetPortStatistics() *Pic_PortStatistics {
	if m != nil {
		return m.PortStatistics
	}
	return nil
}

type Pic_PortStatistics struct {
	Statistic            []*Pic_PortStatistics_Statistic `protobuf:"bytes,1,rep,name=statistic,proto3" json:"statistic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Pic_PortStatistics) Reset()         { *m = Pic_PortStatistics{} }
func (m *Pic_PortStatistics) String() string { return proto.CompactTextString(m) }
func (*Pic_PortStatistics) ProtoMessage()    {}
func (*Pic_PortStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_322574f333bd178c, []int{0, 0}
}

func (m *Pic_PortStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pic_PortStatistics.Unmarshal(m, b)
}
func (m *Pic_PortStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pic_PortStatistics.Marshal(b, m, deterministic)
}
func (m *Pic_PortStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pic_PortStatistics.Merge(m, src)
}
func (m *Pic_PortStatistics) XXX_Size() int {
	return xxx_messageInfo_Pic_PortStatistics.Size(m)
}
func (m *Pic_PortStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Pic_PortStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Pic_PortStatistics proto.InternalMessageInfo

func (m *Pic_PortStatistics) GetStatistic() []*Pic_PortStatistics_Statistic {
	if m != nil {
		return m.Statistic
	}
	return nil
}

type Pic_PortStatistics_Statistic struct {
	ChassisId            string                                `protobuf:"bytes,1,opt,name=chassis_id,json=chassis-id,proto3" json:"chassis_id,omitempty"`
	BoardId              uint32                                `protobuf:"varint,2,opt,name=board_id,json=board-id,proto3" json:"board_id,omitempty"`
	PortType             Pic_PortStatistics_Statistic_PortType `protobuf:"varint,3,opt,name=port_type,json=port-type,proto3,enum=huawei_pic.Pic_PortStatistics_Statistic_PortType" json:"port_type,omitempty"`
	TotalPortNum         uint32                                `protobuf:"varint,4,opt,name=total_port_num,json=total-port-num,proto3" json:"total_port_num,omitempty"`
	PhyUpPortNum         uint32                                `protobuf:"varint,5,opt,name=phy_up_port_num,json=phy-up-port-num,proto3" json:"phy_up_port_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *Pic_PortStatistics_Statistic) Reset()         { *m = Pic_PortStatistics_Statistic{} }
func (m *Pic_PortStatistics_Statistic) String() string { return proto.CompactTextString(m) }
func (*Pic_PortStatistics_Statistic) ProtoMessage()    {}
func (*Pic_PortStatistics_Statistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_322574f333bd178c, []int{0, 0, 0}
}

func (m *Pic_PortStatistics_Statistic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pic_PortStatistics_Statistic.Unmarshal(m, b)
}
func (m *Pic_PortStatistics_Statistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pic_PortStatistics_Statistic.Marshal(b, m, deterministic)
}
func (m *Pic_PortStatistics_Statistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pic_PortStatistics_Statistic.Merge(m, src)
}
func (m *Pic_PortStatistics_Statistic) XXX_Size() int {
	return xxx_messageInfo_Pic_PortStatistics_Statistic.Size(m)
}
func (m *Pic_PortStatistics_Statistic) XXX_DiscardUnknown() {
	xxx_messageInfo_Pic_PortStatistics_Statistic.DiscardUnknown(m)
}

var xxx_messageInfo_Pic_PortStatistics_Statistic proto.InternalMessageInfo

func (m *Pic_PortStatistics_Statistic) GetChassisId() string {
	if m != nil {
		return m.ChassisId
	}
	return ""
}

func (m *Pic_PortStatistics_Statistic) GetBoardId() uint32 {
	if m != nil {
		return m.BoardId
	}
	return 0
}

func (m *Pic_PortStatistics_Statistic) GetPortType() Pic_PortStatistics_Statistic_PortType {
	if m != nil {
		return m.PortType
	}
	return Pic_PortStatistics_Statistic_PortType_ATM
}

func (m *Pic_PortStatistics_Statistic) GetTotalPortNum() uint32 {
	if m != nil {
		return m.TotalPortNum
	}
	return 0
}

func (m *Pic_PortStatistics_Statistic) GetPhyUpPortNum() uint32 {
	if m != nil {
		return m.PhyUpPortNum
	}
	return 0
}

func init() {
	proto.RegisterEnum("huawei_pic.Pic_PortStatistics_Statistic_PortType", Pic_PortStatistics_Statistic_PortType_name, Pic_PortStatistics_Statistic_PortType_value)
	proto.RegisterType((*Pic)(nil), "huawei_pic.Pic")
	proto.RegisterType((*Pic_PortStatistics)(nil), "huawei_pic.Pic.PortStatistics")
	proto.RegisterType((*Pic_PortStatistics_Statistic)(nil), "huawei_pic.Pic.PortStatistics.Statistic")
}

func init() {
	proto.RegisterFile("proto/sensors/huawei-pic/huawei-pic.proto", fileDescriptor_322574f333bd178c)
}

var fileDescriptor_322574f333bd178c = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xea, 0xda, 0x40,
	0x10, 0xc7, 0x1b, 0xa3, 0xd6, 0x4c, 0xfc, 0xb3, 0x2e, 0xb4, 0x04, 0x0f, 0x12, 0x3c, 0x94, 0xf4,
	0x90, 0x88, 0x7a, 0xe8, 0x59, 0x6c, 0x44, 0x41, 0x9b, 0x10, 0x43, 0x69, 0x4f, 0x21, 0xc6, 0x80,
	0x0b, 0xd5, 0x2c, 0xd9, 0x0d, 0xc5, 0x7b, 0x9f, 0xa3, 0x4f, 0xd4, 0x97, 0xe9, 0x1b, 0x94, 0x6c,
	0x30, 0xe9, 0x5e, 0x7e, 0xfc, 0x0e, 0xc2, 0xf8, 0x99, 0xcf, 0x77, 0x66, 0x12, 0x02, 0x1f, 0x69,
	0x9e, 0xf1, 0x6c, 0xce, 0xd2, 0x3b, 0xcb, 0x72, 0x36, 0xbf, 0x16, 0xf1, 0xcf, 0x94, 0xd8, 0x94,
	0x24, 0xff, 0x95, 0x8e, 0x70, 0x30, 0x54, 0x24, 0xa2, 0x24, 0x99, 0xfd, 0xee, 0x80, 0xea, 0x93,
	0x04, 0xef, 0x60, 0x44, 0xb3, 0x9c, 0x47, 0x8c, 0xc7, 0x9c, 0x30, 0x4e, 0x12, 0x66, 0x28, 0xa6,
	0x62, 0xe9, 0xcb, 0xa9, 0xd3, 0xd8, 0x8e, 0x5f, 0xfe, 0xb2, 0x9c, 0x9f, 0x6a, 0x2b, 0x10, 0x31,
	0xbb, 0x89, 0x4d, 0xfe, 0xb4, 0x61, 0x28, 0x3b, 0x78, 0x0b, 0x5a, 0x2d, 0x18, 0x8a, 0xa9, 0x5a,
	0xfa, 0xd2, 0x7a, 0x79, 0xac, 0x53, 0x97, 0x41, 0x13, 0x9d, 0xfc, 0x55, 0x41, 0xab, 0x1b, 0x78,
	0x0a, 0x90, 0x5c, 0x63, 0xc6, 0x08, 0x8b, 0xc8, 0x45, 0x5c, 0xab, 0x05, 0x4f, 0x62, 0x93, 0x0b,
	0x9e, 0x40, 0xef, 0x9c, 0xc5, 0xf9, 0xa5, 0xec, 0xb6, 0x4c, 0xc5, 0x1a, 0x04, 0xd5, 0xff, 0xb2,
	0xe7, 0x81, 0x26, 0x1e, 0x97, 0x3f, 0x68, 0x6a, 0xa8, 0xa6, 0x62, 0x0d, 0x97, 0x8b, 0xd7, 0x5e,
	0x24, 0x1a, 0xe1, 0x83, 0xa6, 0x81, 0x98, 0x61, 0x97, 0x33, 0xf0, 0x07, 0x18, 0xf2, 0x8c, 0xc7,
	0x3f, 0x22, 0x31, 0xf6, 0x5e, 0xdc, 0x8c, 0xb6, 0x58, 0x59, 0x51, 0x5b, 0x88, 0xf7, 0xe2, 0x86,
	0x2d, 0x18, 0xd1, 0xeb, 0x23, 0x2a, 0x68, 0x23, 0x76, 0x84, 0x58, 0x62, 0xbb, 0xa0, 0xb5, 0x39,
	0xfb, 0xd5, 0x82, 0xde, 0x73, 0x13, 0x46, 0xd0, 0x7f, 0xd6, 0xd1, 0x3a, 0x3c, 0xa2, 0x37, 0x12,
	0xd9, 0xac, 0x3e, 0x21, 0x05, 0xe3, 0xea, 0xbd, 0x57, 0xc4, 0xfb, 0xbc, 0x0f, 0x50, 0x0b, 0x8f,
	0x61, 0xd0, 0x30, 0xdf, 0x3b, 0x21, 0x15, 0x8f, 0x40, 0xaf, 0x91, 0xbb, 0x40, 0x6d, 0x19, 0xac,
	0x50, 0x07, 0xbf, 0x07, 0xdc, 0x80, 0xcd, 0xf1, 0xe4, 0x06, 0xfb, 0xf5, 0x01, 0x75, 0xa5, 0x95,
	0x6e, 0xb8, 0x43, 0x6f, 0xf1, 0x3b, 0x18, 0xd7, 0x64, 0x7b, 0x70, 0xbf, 0xb9, 0xfe, 0xee, 0x3b,
	0xea, 0x49, 0x97, 0x08, 0x8c, 0x34, 0x29, 0xec, 0x85, 0x5f, 0x10, 0x48, 0xa4, 0x3c, 0x4d, 0x97,
	0x72, 0x5f, 0xbd, 0xfd, 0xc6, 0x45, 0xfd, 0x73, 0x57, 0x7c, 0xb3, 0xab, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x5c, 0x03, 0x9a, 0xfa, 0xe0, 0x02, 0x00, 0x00,
}
