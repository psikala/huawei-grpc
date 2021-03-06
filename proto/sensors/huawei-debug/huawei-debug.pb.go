// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sensors/huawei-debug/huawei-debug.proto

package huawei_debug

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

type Debug struct {
	CpuInfos             *Debug_CpuInfos        `protobuf:"bytes,1,opt,name=cpu_infos,json=cpu-infos,proto3" json:"cpu_infos,omitempty"`
	ServiceCpuInfos      *Debug_ServiceCpuInfos `protobuf:"bytes,2,opt,name=service_cpu_infos,json=service-cpu-infos,proto3" json:"service_cpu_infos,omitempty"`
	MemoryInfos          *Debug_MemoryInfos     `protobuf:"bytes,3,opt,name=memory_infos,json=memory-infos,proto3" json:"memory_infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Debug) Reset()         { *m = Debug{} }
func (m *Debug) String() string { return proto.CompactTextString(m) }
func (*Debug) ProtoMessage()    {}
func (*Debug) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4e74524612840b, []int{0}
}

func (m *Debug) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Debug.Unmarshal(m, b)
}
func (m *Debug) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Debug.Marshal(b, m, deterministic)
}
func (m *Debug) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Debug.Merge(m, src)
}
func (m *Debug) XXX_Size() int {
	return xxx_messageInfo_Debug.Size(m)
}
func (m *Debug) XXX_DiscardUnknown() {
	xxx_messageInfo_Debug.DiscardUnknown(m)
}

var xxx_messageInfo_Debug proto.InternalMessageInfo

func (m *Debug) GetCpuInfos() *Debug_CpuInfos {
	if m != nil {
		return m.CpuInfos
	}
	return nil
}

func (m *Debug) GetServiceCpuInfos() *Debug_ServiceCpuInfos {
	if m != nil {
		return m.ServiceCpuInfos
	}
	return nil
}

func (m *Debug) GetMemoryInfos() *Debug_MemoryInfos {
	if m != nil {
		return m.MemoryInfos
	}
	return nil
}

type Debug_CpuInfos struct {
	CpuInfo              []*Debug_CpuInfos_CpuInfo `protobuf:"bytes,1,rep,name=cpu_info,json=cpu-info,proto3" json:"cpu_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Debug_CpuInfos) Reset()         { *m = Debug_CpuInfos{} }
func (m *Debug_CpuInfos) String() string { return proto.CompactTextString(m) }
func (*Debug_CpuInfos) ProtoMessage()    {}
func (*Debug_CpuInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4e74524612840b, []int{0, 0}
}

func (m *Debug_CpuInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Debug_CpuInfos.Unmarshal(m, b)
}
func (m *Debug_CpuInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Debug_CpuInfos.Marshal(b, m, deterministic)
}
func (m *Debug_CpuInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Debug_CpuInfos.Merge(m, src)
}
func (m *Debug_CpuInfos) XXX_Size() int {
	return xxx_messageInfo_Debug_CpuInfos.Size(m)
}
func (m *Debug_CpuInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_Debug_CpuInfos.DiscardUnknown(m)
}

var xxx_messageInfo_Debug_CpuInfos proto.InternalMessageInfo

func (m *Debug_CpuInfos) GetCpuInfo() []*Debug_CpuInfos_CpuInfo {
	if m != nil {
		return m.CpuInfo
	}
	return nil
}

type Debug_CpuInfos_CpuInfo struct {
	Position                string   `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	OverloadThreshold       uint32   `protobuf:"varint,2,opt,name=overload_threshold,json=overload-threshold,proto3" json:"overload_threshold,omitempty"`
	UnoverloadThreshold     uint32   `protobuf:"varint,3,opt,name=unoverload_threshold,json=unoverload-threshold,proto3" json:"unoverload_threshold,omitempty"`
	Interval                uint32   `protobuf:"varint,4,opt,name=interval,proto3" json:"interval,omitempty"`
	Index                   uint32   `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	SystemCpuUsage          uint32   `protobuf:"varint,6,opt,name=system_cpu_usage,json=system-cpu-usage,proto3" json:"system_cpu_usage,omitempty"`
	MonitorNumber           uint32   `protobuf:"varint,7,opt,name=monitor_number,json=monitor-number,proto3" json:"monitor_number,omitempty"`
	MonitorCycle            uint32   `protobuf:"varint,8,opt,name=monitor_cycle,json=monitor-cycle,proto3" json:"monitor_cycle,omitempty"`
	OverloadStateChangeTime string   `protobuf:"bytes,9,opt,name=overload_state_change_time,json=overload-state-change-time,proto3" json:"overload_state_change_time,omitempty"`
	CurrentOverloadState    string   `protobuf:"bytes,10,opt,name=current_overload_state,json=current-overload-state,proto3" json:"current_overload_state,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Debug_CpuInfos_CpuInfo) Reset()         { *m = Debug_CpuInfos_CpuInfo{} }
func (m *Debug_CpuInfos_CpuInfo) String() string { return proto.CompactTextString(m) }
func (*Debug_CpuInfos_CpuInfo) ProtoMessage()    {}
func (*Debug_CpuInfos_CpuInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4e74524612840b, []int{0, 0, 0}
}

func (m *Debug_CpuInfos_CpuInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Debug_CpuInfos_CpuInfo.Unmarshal(m, b)
}
func (m *Debug_CpuInfos_CpuInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Debug_CpuInfos_CpuInfo.Marshal(b, m, deterministic)
}
func (m *Debug_CpuInfos_CpuInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Debug_CpuInfos_CpuInfo.Merge(m, src)
}
func (m *Debug_CpuInfos_CpuInfo) XXX_Size() int {
	return xxx_messageInfo_Debug_CpuInfos_CpuInfo.Size(m)
}
func (m *Debug_CpuInfos_CpuInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Debug_CpuInfos_CpuInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Debug_CpuInfos_CpuInfo proto.InternalMessageInfo

func (m *Debug_CpuInfos_CpuInfo) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Debug_CpuInfos_CpuInfo) GetOverloadThreshold() uint32 {
	if m != nil {
		return m.OverloadThreshold
	}
	return 0
}

func (m *Debug_CpuInfos_CpuInfo) GetUnoverloadThreshold() uint32 {
	if m != nil {
		return m.UnoverloadThreshold
	}
	return 0
}

func (m *Debug_CpuInfos_CpuInfo) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *Debug_CpuInfos_CpuInfo) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Debug_CpuInfos_CpuInfo) GetSystemCpuUsage() uint32 {
	if m != nil {
		return m.SystemCpuUsage
	}
	return 0
}

func (m *Debug_CpuInfos_CpuInfo) GetMonitorNumber() uint32 {
	if m != nil {
		return m.MonitorNumber
	}
	return 0
}

func (m *Debug_CpuInfos_CpuInfo) GetMonitorCycle() uint32 {
	if m != nil {
		return m.MonitorCycle
	}
	return 0
}

func (m *Debug_CpuInfos_CpuInfo) GetOverloadStateChangeTime() string {
	if m != nil {
		return m.OverloadStateChangeTime
	}
	return ""
}

func (m *Debug_CpuInfos_CpuInfo) GetCurrentOverloadState() string {
	if m != nil {
		return m.CurrentOverloadState
	}
	return ""
}

type Debug_ServiceCpuInfos struct {
	ServiceCpuInfo       []*Debug_ServiceCpuInfos_ServiceCpuInfo `protobuf:"bytes,1,rep,name=service_cpu_info,json=service-cpu-info,proto3" json:"service_cpu_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *Debug_ServiceCpuInfos) Reset()         { *m = Debug_ServiceCpuInfos{} }
func (m *Debug_ServiceCpuInfos) String() string { return proto.CompactTextString(m) }
func (*Debug_ServiceCpuInfos) ProtoMessage()    {}
func (*Debug_ServiceCpuInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4e74524612840b, []int{0, 1}
}

func (m *Debug_ServiceCpuInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Debug_ServiceCpuInfos.Unmarshal(m, b)
}
func (m *Debug_ServiceCpuInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Debug_ServiceCpuInfos.Marshal(b, m, deterministic)
}
func (m *Debug_ServiceCpuInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Debug_ServiceCpuInfos.Merge(m, src)
}
func (m *Debug_ServiceCpuInfos) XXX_Size() int {
	return xxx_messageInfo_Debug_ServiceCpuInfos.Size(m)
}
func (m *Debug_ServiceCpuInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_Debug_ServiceCpuInfos.DiscardUnknown(m)
}

var xxx_messageInfo_Debug_ServiceCpuInfos proto.InternalMessageInfo

func (m *Debug_ServiceCpuInfos) GetServiceCpuInfo() []*Debug_ServiceCpuInfos_ServiceCpuInfo {
	if m != nil {
		return m.ServiceCpuInfo
	}
	return nil
}

type Debug_ServiceCpuInfos_ServiceCpuInfo struct {
	Position             string   `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=service-name,proto3" json:"service_name,omitempty"`
	ServiceCpuUsage      uint32   `protobuf:"varint,3,opt,name=service_cpu_usage,json=service-cpu-usage,proto3" json:"service_cpu_usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) Reset()         { *m = Debug_ServiceCpuInfos_ServiceCpuInfo{} }
func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) String() string { return proto.CompactTextString(m) }
func (*Debug_ServiceCpuInfos_ServiceCpuInfo) ProtoMessage()    {}
func (*Debug_ServiceCpuInfos_ServiceCpuInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4e74524612840b, []int{0, 1, 0}
}

func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Debug_ServiceCpuInfos_ServiceCpuInfo.Unmarshal(m, b)
}
func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Debug_ServiceCpuInfos_ServiceCpuInfo.Marshal(b, m, deterministic)
}
func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Debug_ServiceCpuInfos_ServiceCpuInfo.Merge(m, src)
}
func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) XXX_Size() int {
	return xxx_messageInfo_Debug_ServiceCpuInfos_ServiceCpuInfo.Size(m)
}
func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Debug_ServiceCpuInfos_ServiceCpuInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Debug_ServiceCpuInfos_ServiceCpuInfo proto.InternalMessageInfo

func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Debug_ServiceCpuInfos_ServiceCpuInfo) GetServiceCpuUsage() uint32 {
	if m != nil {
		return m.ServiceCpuUsage
	}
	return 0
}

type Debug_MemoryInfos struct {
	MemoryInfo           []*Debug_MemoryInfos_MemoryInfo `protobuf:"bytes,1,rep,name=memory_info,json=memory-info,proto3" json:"memory_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Debug_MemoryInfos) Reset()         { *m = Debug_MemoryInfos{} }
func (m *Debug_MemoryInfos) String() string { return proto.CompactTextString(m) }
func (*Debug_MemoryInfos) ProtoMessage()    {}
func (*Debug_MemoryInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4e74524612840b, []int{0, 2}
}

func (m *Debug_MemoryInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Debug_MemoryInfos.Unmarshal(m, b)
}
func (m *Debug_MemoryInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Debug_MemoryInfos.Marshal(b, m, deterministic)
}
func (m *Debug_MemoryInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Debug_MemoryInfos.Merge(m, src)
}
func (m *Debug_MemoryInfos) XXX_Size() int {
	return xxx_messageInfo_Debug_MemoryInfos.Size(m)
}
func (m *Debug_MemoryInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_Debug_MemoryInfos.DiscardUnknown(m)
}

var xxx_messageInfo_Debug_MemoryInfos proto.InternalMessageInfo

func (m *Debug_MemoryInfos) GetMemoryInfo() []*Debug_MemoryInfos_MemoryInfo {
	if m != nil {
		return m.MemoryInfo
	}
	return nil
}

type Debug_MemoryInfos_MemoryInfo struct {
	Position                string   `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	OverloadThreshold       uint32   `protobuf:"varint,2,opt,name=overload_threshold,json=overload-threshold,proto3" json:"overload_threshold,omitempty"`
	UnoverloadThreshold     uint32   `protobuf:"varint,3,opt,name=unoverload_threshold,json=unoverload-threshold,proto3" json:"unoverload_threshold,omitempty"`
	Index                   uint32   `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	OsMemoryTotal           uint32   `protobuf:"varint,5,opt,name=os_memory_total,json=os-memory-total,proto3" json:"os_memory_total,omitempty"`
	OsMemoryUse             uint32   `protobuf:"varint,6,opt,name=os_memory_use,json=os-memory-use,proto3" json:"os_memory_use,omitempty"`
	OsMemoryFree            uint32   `protobuf:"varint,7,opt,name=os_memory_free,json=os-memory-free,proto3" json:"os_memory_free,omitempty"`
	OsMemoryUsage           uint32   `protobuf:"varint,8,opt,name=os_memory_usage,json=os-memory-usage,proto3" json:"os_memory_usage,omitempty"`
	DoMemoryTotal           uint32   `protobuf:"varint,9,opt,name=do_memory_total,json=do-memory-total,proto3" json:"do_memory_total,omitempty"`
	DoMemoryUse             uint32   `protobuf:"varint,10,opt,name=do_memory_use,json=do-memory-use,proto3" json:"do_memory_use,omitempty"`
	DoMemoryFree            uint32   `protobuf:"varint,11,opt,name=do_memory_free,json=do-memory-free,proto3" json:"do_memory_free,omitempty"`
	DoMemoryUsage           uint32   `protobuf:"varint,12,opt,name=do_memory_usage,json=do-memory-usage,proto3" json:"do_memory_usage,omitempty"`
	SimpleMemoryTotal       uint32   `protobuf:"varint,13,opt,name=simple_memory_total,json=simple-memory-total,proto3" json:"simple_memory_total,omitempty"`
	SimpleMemoryUse         uint32   `protobuf:"varint,14,opt,name=simple_memory_use,json=simple-memory-use,proto3" json:"simple_memory_use,omitempty"`
	SimpleMemoryFree        uint32   `protobuf:"varint,15,opt,name=simple_memory_free,json=simple-memory-free,proto3" json:"simple_memory_free,omitempty"`
	SimpleMemoryUsage       uint32   `protobuf:"varint,16,opt,name=simple_memory_usage,json=simple-memory-usage,proto3" json:"simple_memory_usage,omitempty"`
	OverloadStateChangeTime string   `protobuf:"bytes,17,opt,name=overload_state_change_time,json=overload-state-change-time,proto3" json:"overload_state_change_time,omitempty"`
	CurrentOverloadState    string   `protobuf:"bytes,18,opt,name=current_overload_state,json=current-overload-state,proto3" json:"current_overload_state,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Debug_MemoryInfos_MemoryInfo) Reset()         { *m = Debug_MemoryInfos_MemoryInfo{} }
func (m *Debug_MemoryInfos_MemoryInfo) String() string { return proto.CompactTextString(m) }
func (*Debug_MemoryInfos_MemoryInfo) ProtoMessage()    {}
func (*Debug_MemoryInfos_MemoryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f4e74524612840b, []int{0, 2, 0}
}

func (m *Debug_MemoryInfos_MemoryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Debug_MemoryInfos_MemoryInfo.Unmarshal(m, b)
}
func (m *Debug_MemoryInfos_MemoryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Debug_MemoryInfos_MemoryInfo.Marshal(b, m, deterministic)
}
func (m *Debug_MemoryInfos_MemoryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Debug_MemoryInfos_MemoryInfo.Merge(m, src)
}
func (m *Debug_MemoryInfos_MemoryInfo) XXX_Size() int {
	return xxx_messageInfo_Debug_MemoryInfos_MemoryInfo.Size(m)
}
func (m *Debug_MemoryInfos_MemoryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Debug_MemoryInfos_MemoryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Debug_MemoryInfos_MemoryInfo proto.InternalMessageInfo

func (m *Debug_MemoryInfos_MemoryInfo) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Debug_MemoryInfos_MemoryInfo) GetOverloadThreshold() uint32 {
	if m != nil {
		return m.OverloadThreshold
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetUnoverloadThreshold() uint32 {
	if m != nil {
		return m.UnoverloadThreshold
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetOsMemoryTotal() uint32 {
	if m != nil {
		return m.OsMemoryTotal
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetOsMemoryUse() uint32 {
	if m != nil {
		return m.OsMemoryUse
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetOsMemoryFree() uint32 {
	if m != nil {
		return m.OsMemoryFree
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetOsMemoryUsage() uint32 {
	if m != nil {
		return m.OsMemoryUsage
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetDoMemoryTotal() uint32 {
	if m != nil {
		return m.DoMemoryTotal
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetDoMemoryUse() uint32 {
	if m != nil {
		return m.DoMemoryUse
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetDoMemoryFree() uint32 {
	if m != nil {
		return m.DoMemoryFree
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetDoMemoryUsage() uint32 {
	if m != nil {
		return m.DoMemoryUsage
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetSimpleMemoryTotal() uint32 {
	if m != nil {
		return m.SimpleMemoryTotal
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetSimpleMemoryUse() uint32 {
	if m != nil {
		return m.SimpleMemoryUse
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetSimpleMemoryFree() uint32 {
	if m != nil {
		return m.SimpleMemoryFree
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetSimpleMemoryUsage() uint32 {
	if m != nil {
		return m.SimpleMemoryUsage
	}
	return 0
}

func (m *Debug_MemoryInfos_MemoryInfo) GetOverloadStateChangeTime() string {
	if m != nil {
		return m.OverloadStateChangeTime
	}
	return ""
}

func (m *Debug_MemoryInfos_MemoryInfo) GetCurrentOverloadState() string {
	if m != nil {
		return m.CurrentOverloadState
	}
	return ""
}

func init() {
	proto.RegisterType((*Debug)(nil), "huawei_debug.Debug")
	proto.RegisterType((*Debug_CpuInfos)(nil), "huawei_debug.Debug.CpuInfos")
	proto.RegisterType((*Debug_CpuInfos_CpuInfo)(nil), "huawei_debug.Debug.CpuInfos.CpuInfo")
	proto.RegisterType((*Debug_ServiceCpuInfos)(nil), "huawei_debug.Debug.ServiceCpuInfos")
	proto.RegisterType((*Debug_ServiceCpuInfos_ServiceCpuInfo)(nil), "huawei_debug.Debug.ServiceCpuInfos.ServiceCpuInfo")
	proto.RegisterType((*Debug_MemoryInfos)(nil), "huawei_debug.Debug.MemoryInfos")
	proto.RegisterType((*Debug_MemoryInfos_MemoryInfo)(nil), "huawei_debug.Debug.MemoryInfos.MemoryInfo")
}

func init() {
	proto.RegisterFile("proto/sensors/huawei-debug/huawei-debug.proto", fileDescriptor_6f4e74524612840b)
}

var fileDescriptor_6f4e74524612840b = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x95, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0xd5, 0x5f, 0xb7, 0x2e, 0x3d, 0xfd, 0x3b, 0xff, 0x26, 0x14, 0x45, 0x48, 0x4c, 0x63,
	0x42, 0xd5, 0x84, 0x3b, 0x54, 0x24, 0x2e, 0xb8, 0x40, 0x48, 0xe3, 0x06, 0x09, 0x2e, 0x08, 0xf7,
	0x44, 0x59, 0xeb, 0xad, 0x91, 0x9a, 0xb8, 0xb2, 0x9d, 0xc1, 0x6e, 0x78, 0x16, 0x5e, 0x80, 0x67,
	0xe1, 0x59, 0xb8, 0x47, 0x08, 0xf9, 0xd8, 0x4e, 0xe2, 0xae, 0x2a, 0x48, 0x5c, 0x70, 0x17, 0x9f,
	0xf3, 0x39, 0xc7, 0x5f, 0x3b, 0xdf, 0x93, 0x00, 0x5d, 0x0b, 0xae, 0xf8, 0xb9, 0x64, 0x85, 0xe4,
	0x42, 0x9e, 0x2f, 0xcb, 0xf4, 0x23, 0xcb, 0xe8, 0x82, 0x5d, 0x96, 0xd7, 0xde, 0x62, 0x8a, 0x1c,
	0xe9, 0x9b, 0x58, 0x82, 0xb1, 0x93, 0x6f, 0x43, 0xd8, 0x7f, 0xa5, 0x9f, 0xc8, 0x73, 0xe8, 0xce,
	0xd7, 0x65, 0x92, 0x15, 0x57, 0x5c, 0x86, 0xad, 0xe3, 0xd6, 0xa4, 0x37, 0xbb, 0x3f, 0x6d, 0xb2,
	0x53, 0xe4, 0xa6, 0x17, 0xeb, 0xf2, 0xb5, 0x66, 0x62, 0x8d, 0x53, 0xc4, 0xc9, 0x3b, 0x38, 0x94,
	0x4c, 0xdc, 0x64, 0x73, 0x96, 0xd4, 0x3d, 0xfe, 0xc3, 0x1e, 0x0f, 0xb7, 0xf5, 0x78, 0x6f, 0xe0,
	0xaa, 0x95, 0xab, 0xa6, 0x75, 0xcb, 0x0b, 0xe8, 0xe7, 0x2c, 0xe7, 0xe2, 0xd6, 0x76, 0x6b, 0x63,
	0xb7, 0x07, 0xdb, 0xba, 0xbd, 0x45, 0xce, 0x74, 0xb2, 0x45, 0xa6, 0x49, 0xf4, 0xa3, 0x0d, 0x81,
	0xdb, 0x84, 0xbc, 0x84, 0xc0, 0x89, 0x0b, 0x5b, 0xc7, 0xed, 0x49, 0x6f, 0x76, 0xba, 0xeb, 0x7c,
	0xee, 0x21, 0x0e, 0x9c, 0xa8, 0xe8, 0x6b, 0x1b, 0x0e, 0x6c, 0x94, 0x44, 0x10, 0xac, 0xb9, 0xcc,
	0x54, 0xc6, 0x0b, 0xbc, 0xad, 0x6e, 0x5c, 0xad, 0xc9, 0x14, 0x08, 0xbf, 0x61, 0x62, 0xc5, 0xd3,
	0x45, 0xa2, 0x96, 0x82, 0xc9, 0x25, 0x5f, 0x2d, 0xf0, 0x3e, 0x06, 0x71, 0x95, 0xa1, 0x55, 0x86,
	0xcc, 0xe0, 0xa8, 0x2c, 0xb6, 0x54, 0xb4, 0xb1, 0xa2, 0x91, 0x6b, 0xd4, 0x44, 0x10, 0x64, 0x85,
	0x62, 0xe2, 0x26, 0x5d, 0x85, 0x7b, 0xc8, 0x55, 0x6b, 0x72, 0x04, 0xfb, 0x59, 0xb1, 0x60, 0x9f,
	0xc2, 0x7d, 0x4c, 0x98, 0x05, 0x39, 0x83, 0xb1, 0xbc, 0x95, 0x8a, 0xe5, 0xf8, 0x8e, 0x4a, 0x99,
	0x5e, 0xb3, 0xb0, 0x83, 0x80, 0x8d, 0xe3, 0xed, 0x63, 0x9c, 0x3c, 0x82, 0x61, 0xce, 0x8b, 0x4c,
	0x71, 0x91, 0x14, 0x65, 0x7e, 0xc9, 0x44, 0x78, 0x80, 0xa4, 0x8b, 0x52, 0x13, 0x25, 0xa7, 0x30,
	0x70, 0xdc, 0xfc, 0x76, 0xbe, 0x62, 0x61, 0x80, 0x98, 0x0b, 0x52, 0x0c, 0x92, 0x17, 0x10, 0x55,
	0xa7, 0x93, 0x2a, 0x55, 0x2c, 0x99, 0x2f, 0xd3, 0xe2, 0x9a, 0x25, 0x2a, 0xcb, 0x59, 0xd8, 0xc5,
	0xdb, 0xab, 0x08, 0x8a, 0x04, 0x35, 0x04, 0xd5, 0x04, 0x79, 0x06, 0xf7, 0xe6, 0xa5, 0x10, 0xac,
	0x50, 0x89, 0xdf, 0x27, 0x04, 0xac, 0x75, 0x59, 0xea, 0xf7, 0x88, 0xbe, 0xb7, 0x60, 0xb4, 0x61,
	0x35, 0xf2, 0x01, 0xc6, 0x9b, 0x56, 0xb5, 0x6e, 0x98, 0xfd, 0x81, 0x53, 0x37, 0xd6, 0xf1, 0x78,
	0xd3, 0xb8, 0xd1, 0x67, 0x18, 0xfa, 0xcc, 0x4e, 0xa7, 0x9c, 0x40, 0xdf, 0xa9, 0x29, 0xd2, 0x9c,
	0xa1, 0x47, 0xba, 0xb1, 0x8b, 0x51, 0x1d, 0x23, 0x8f, 0xfd, 0xe1, 0x32, 0x2f, 0xce, 0x58, 0xc3,
	0x9b, 0x1b, 0x4c, 0x44, 0x3f, 0x3b, 0xd0, 0x6b, 0x0c, 0x04, 0x79, 0x03, 0xbd, 0xc6, 0x1c, 0xd9,
	0xa3, 0x9e, 0xfd, 0x66, 0x8c, 0x1a, 0xcf, 0x71, 0xaf, 0x31, 0x51, 0xd1, 0x97, 0x0e, 0x40, 0x9d,
	0xfb, 0xe7, 0x43, 0x50, 0x19, 0x7d, 0xaf, 0x69, 0xf4, 0x09, 0x8c, 0xb8, 0x4c, 0xec, 0xa9, 0x15,
	0x57, 0xe9, 0xca, 0x0e, 0xc2, 0x88, 0x4b, 0x6a, 0x4f, 0x83, 0x61, 0x6d, 0xdf, 0x9a, 0x2c, 0xa5,
	0x9b, 0x87, 0x41, 0xcd, 0x95, 0x12, 0x87, 0xa1, 0xa6, 0xae, 0x04, 0x63, 0x6e, 0x18, 0x6a, 0x4c,
	0x47, 0xfd, 0x7d, 0xcd, 0x6b, 0x0a, 0x36, 0xf7, 0x35, 0xe3, 0x35, 0x81, 0xd1, 0x82, 0xfb, 0x0a,
	0xbb, 0x86, 0x5c, 0xf0, 0x3b, 0x0a, 0x6b, 0x52, 0x2b, 0x04, 0xa3, 0xb0, 0xe6, 0xac, 0xc2, 0x9a,
	0x42, 0x85, 0x3d, 0xa3, 0xb0, 0xc6, 0x9c, 0xc2, 0x66, 0x37, 0xad, 0xb0, 0xbf, 0xb9, 0xaf, 0x51,
	0xf8, 0x04, 0xfe, 0x97, 0x59, 0xbe, 0x5e, 0x31, 0x5f, 0xe5, 0x00, 0x69, 0x9b, 0xf2, 0x95, 0x6a,
	0x9b, 0x7a, 0x15, 0x5a, 0xed, 0xd0, 0xda, 0xd4, 0xe3, 0xb5, 0xe2, 0x29, 0x10, 0x9f, 0x46, 0xd5,
	0x23, 0xe3, 0x0e, 0x1f, 0x47, 0xe5, 0x77, 0xf4, 0x18, 0xf5, 0xe3, 0x6d, 0x7a, 0xcc, 0x09, 0x76,
	0x7f, 0x74, 0x0e, 0xff, 0xe2, 0xa3, 0x43, 0x76, 0x7d, 0x74, 0x2e, 0x3b, 0xf8, 0x9b, 0x7d, 0xfa,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xd8, 0xdb, 0xd4, 0x97, 0x07, 0x00, 0x00,
}
