// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/huawei-dialout/huawei-grpc-dialout.proto

package huawei_dialout

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

type ServiceArgs struct {
	ReqId int64 `protobuf:"varint,1,opt,name=ReqId,proto3" json:"ReqId,omitempty"`
	// Types that are valid to be assigned to MessageData:
	//	*ServiceArgs_Data
	//	*ServiceArgs_DataJson
	MessageData          isServiceArgs_MessageData `protobuf_oneof:"MessageData"`
	Errors               string                    `protobuf:"bytes,3,opt,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ServiceArgs) Reset()         { *m = ServiceArgs{} }
func (m *ServiceArgs) String() string { return proto.CompactTextString(m) }
func (*ServiceArgs) ProtoMessage()    {}
func (*ServiceArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f724a8af9c73fa3, []int{0}
}

func (m *ServiceArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceArgs.Unmarshal(m, b)
}
func (m *ServiceArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceArgs.Marshal(b, m, deterministic)
}
func (m *ServiceArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceArgs.Merge(m, src)
}
func (m *ServiceArgs) XXX_Size() int {
	return xxx_messageInfo_ServiceArgs.Size(m)
}
func (m *ServiceArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceArgs proto.InternalMessageInfo

func (m *ServiceArgs) GetReqId() int64 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

type isServiceArgs_MessageData interface {
	isServiceArgs_MessageData()
}

type ServiceArgs_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type ServiceArgs_DataJson struct {
	DataJson string `protobuf:"bytes,4,opt,name=data_json,json=dataJson,proto3,oneof"`
}

func (*ServiceArgs_Data) isServiceArgs_MessageData() {}

func (*ServiceArgs_DataJson) isServiceArgs_MessageData() {}

func (m *ServiceArgs) GetMessageData() isServiceArgs_MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *ServiceArgs) GetData() []byte {
	if x, ok := m.GetMessageData().(*ServiceArgs_Data); ok {
		return x.Data
	}
	return nil
}

func (m *ServiceArgs) GetDataJson() string {
	if x, ok := m.GetMessageData().(*ServiceArgs_DataJson); ok {
		return x.DataJson
	}
	return ""
}

func (m *ServiceArgs) GetErrors() string {
	if m != nil {
		return m.Errors
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServiceArgs) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServiceArgs_Data)(nil),
		(*ServiceArgs_DataJson)(nil),
	}
}

func init() {
	proto.RegisterType((*ServiceArgs)(nil), "huawei_dialout.serviceArgs")
}

func init() {
	proto.RegisterFile("proto/huawei-dialout/huawei-grpc-dialout.proto", fileDescriptor_5f724a8af9c73fa3)
}

var fileDescriptor_5f724a8af9c73fa3 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x28, 0x4d, 0x2c, 0x4f, 0xcd, 0xd4, 0x4d, 0xc9, 0x4c, 0xcc, 0xc9, 0x2f, 0x2d,
	0x81, 0x71, 0xd3, 0x8b, 0x0a, 0x92, 0x61, 0x62, 0x10, 0x85, 0x42, 0x7c, 0x10, 0xa9, 0x78, 0xa8,
	0xa8, 0x52, 0x3d, 0x17, 0x77, 0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x63, 0x51, 0x7a, 0xb1,
	0x90, 0x08, 0x17, 0x6b, 0x50, 0x6a, 0xa1, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10,
	0x84, 0x23, 0x24, 0xc2, 0xc5, 0x92, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe3,
	0xc1, 0x10, 0x04, 0xe6, 0x09, 0xc9, 0x72, 0x71, 0x82, 0xe8, 0xf8, 0xac, 0xe2, 0xfc, 0x3c, 0x09,
	0x16, 0x05, 0x46, 0x0d, 0x4e, 0x0f, 0x86, 0x20, 0x0e, 0x90, 0x90, 0x57, 0x71, 0x7e, 0x9e, 0x90,
	0x18, 0x17, 0x5b, 0x6a, 0x51, 0x51, 0x7e, 0x51, 0xb1, 0x04, 0x33, 0x48, 0x2e, 0x08, 0xca, 0x73,
	0xe2, 0xe5, 0xe2, 0xf6, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x75, 0x49, 0x2c, 0x49, 0x34, 0x4a,
	0xe0, 0xe2, 0x4f, 0x0f, 0x0a, 0x70, 0x06, 0xb1, 0xa1, 0x0e, 0x11, 0xf2, 0xe5, 0xe2, 0x06, 0x99,
	0x12, 0x50, 0x9a, 0x94, 0x93, 0x59, 0x9c, 0x21, 0x24, 0xad, 0x87, 0xea, 0x66, 0x3d, 0x24, 0x07,
	0x4b, 0xe1, 0x93, 0x54, 0x62, 0xd0, 0x60, 0x34, 0x60, 0x4c, 0x62, 0x03, 0xfb, 0xdc, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xdb, 0x7a, 0x94, 0xc7, 0x2b, 0x01, 0x00, 0x00,
}
