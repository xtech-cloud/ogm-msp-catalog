// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/catalog/volume.proto

package catalog

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 创建的请求
type VolumeCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`   // 键
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // 名称
}

func (x *VolumeCreateRequest) Reset() {
	*x = VolumeCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeCreateRequest) ProtoMessage() {}

func (x *VolumeCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeCreateRequest.ProtoReflect.Descriptor instead.
func (*VolumeCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{0}
}

func (x *VolumeCreateRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *VolumeCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 更新的请求
type VolumeUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`   // 键
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` // 名称
}

func (x *VolumeUpdateRequest) Reset() {
	*x = VolumeUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeUpdateRequest) ProtoMessage() {}

func (x *VolumeUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeUpdateRequest.ProtoReflect.Descriptor instead.
func (*VolumeUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{1}
}

func (x *VolumeUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *VolumeUpdateRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *VolumeUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 删除的请求
type VolumeDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
}

func (x *VolumeDeleteRequest) Reset() {
	*x = VolumeDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeDeleteRequest) ProtoMessage() {}

func (x *VolumeDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeDeleteRequest.ProtoReflect.Descriptor instead.
func (*VolumeDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{2}
}

func (x *VolumeDeleteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 列举的请求
type VolumeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
}

func (x *VolumeListRequest) Reset() {
	*x = VolumeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeListRequest) ProtoMessage() {}

func (x *VolumeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeListRequest.ProtoReflect.Descriptor instead.
func (*VolumeListRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{3}
}

func (x *VolumeListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *VolumeListRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 搜索的请求
type VolumeSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
	Key    string `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`       // 键
	Name   string `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`     // 名称
}

func (x *VolumeSearchRequest) Reset() {
	*x = VolumeSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeSearchRequest) ProtoMessage() {}

func (x *VolumeSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeSearchRequest.ProtoReflect.Descriptor instead.
func (*VolumeSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{4}
}

func (x *VolumeSearchRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *VolumeSearchRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *VolumeSearchRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *VolumeSearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 列举的回复
type VolumeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  uint64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  // 总数
	Entity []*VolumeEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"` // 实体列表
}

func (x *VolumeListResponse) Reset() {
	*x = VolumeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeListResponse) ProtoMessage() {}

func (x *VolumeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeListResponse.ProtoReflect.Descriptor instead.
func (*VolumeListResponse) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{5}
}

func (x *VolumeListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *VolumeListResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *VolumeListResponse) GetEntity() []*VolumeEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

// 获取的请求
type VolumeGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
}

func (x *VolumeGetRequest) Reset() {
	*x = VolumeGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeGetRequest) ProtoMessage() {}

func (x *VolumeGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeGetRequest.ProtoReflect.Descriptor instead.
func (*VolumeGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{6}
}

func (x *VolumeGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 获取的回复
type VolumeGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Entity *VolumeEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"` // 实体
}

func (x *VolumeGetResponse) Reset() {
	*x = VolumeGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeGetResponse) ProtoMessage() {}

func (x *VolumeGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeGetResponse.ProtoReflect.Descriptor instead.
func (*VolumeGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{7}
}

func (x *VolumeGetResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *VolumeGetResponse) GetEntity() *VolumeEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

// 导出的请求
type VolumeExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
}

func (x *VolumeExportRequest) Reset() {
	*x = VolumeExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeExportRequest) ProtoMessage() {}

func (x *VolumeExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeExportRequest.ProtoReflect.Descriptor instead.
func (*VolumeExportRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{8}
}

func (x *VolumeExportRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 导出的回复
type VolumeExportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Dump   string  `protobuf:"bytes,2,opt,name=dump,proto3" json:"dump,omitempty"`     // 导出文本
}

func (x *VolumeExportResponse) Reset() {
	*x = VolumeExportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_volume_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeExportResponse) ProtoMessage() {}

func (x *VolumeExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_volume_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeExportResponse.ProtoReflect.Descriptor instead.
func (*VolumeExportResponse) Descriptor() ([]byte, []int) {
	return file_proto_catalog_volume_proto_rawDescGZIP(), []int{9}
}

func (x *VolumeExportResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *VolumeExportResponse) GetDump() string {
	if x != nil {
		return x.Dump
	}
	return ""
}

var File_proto_catalog_volume_proto protoreflect.FileDescriptor

var file_proto_catalog_volume_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3b, 0x0a, 0x13, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x13,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a,
	0x13, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x13, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7e, 0x0a, 0x12, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x10, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x67,
	0x0a, 0x11, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x29, 0x0a, 0x13, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x22, 0x51, 0x0a, 0x14, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x75, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x75, 0x6d, 0x70, 0x32, 0xc2, 0x03, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55,
	0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55, 0x75, 0x69, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3b, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_catalog_volume_proto_rawDescOnce sync.Once
	file_proto_catalog_volume_proto_rawDescData = file_proto_catalog_volume_proto_rawDesc
)

func file_proto_catalog_volume_proto_rawDescGZIP() []byte {
	file_proto_catalog_volume_proto_rawDescOnce.Do(func() {
		file_proto_catalog_volume_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_catalog_volume_proto_rawDescData)
	})
	return file_proto_catalog_volume_proto_rawDescData
}

var file_proto_catalog_volume_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_catalog_volume_proto_goTypes = []interface{}{
	(*VolumeCreateRequest)(nil),  // 0: group.VolumeCreateRequest
	(*VolumeUpdateRequest)(nil),  // 1: group.VolumeUpdateRequest
	(*VolumeDeleteRequest)(nil),  // 2: group.VolumeDeleteRequest
	(*VolumeListRequest)(nil),    // 3: group.VolumeListRequest
	(*VolumeSearchRequest)(nil),  // 4: group.VolumeSearchRequest
	(*VolumeListResponse)(nil),   // 5: group.VolumeListResponse
	(*VolumeGetRequest)(nil),     // 6: group.VolumeGetRequest
	(*VolumeGetResponse)(nil),    // 7: group.VolumeGetResponse
	(*VolumeExportRequest)(nil),  // 8: group.VolumeExportRequest
	(*VolumeExportResponse)(nil), // 9: group.VolumeExportResponse
	(*Status)(nil),               // 10: group.Status
	(*VolumeEntity)(nil),         // 11: group.VolumeEntity
	(*UuidResponse)(nil),         // 12: group.UuidResponse
}
var file_proto_catalog_volume_proto_depIdxs = []int32{
	10, // 0: group.VolumeListResponse.status:type_name -> group.Status
	11, // 1: group.VolumeListResponse.entity:type_name -> group.VolumeEntity
	10, // 2: group.VolumeGetResponse.status:type_name -> group.Status
	11, // 3: group.VolumeGetResponse.entity:type_name -> group.VolumeEntity
	10, // 4: group.VolumeExportResponse.status:type_name -> group.Status
	0,  // 5: group.Volume.Create:input_type -> group.VolumeCreateRequest
	1,  // 6: group.Volume.Update:input_type -> group.VolumeUpdateRequest
	2,  // 7: group.Volume.Delete:input_type -> group.VolumeDeleteRequest
	6,  // 8: group.Volume.Get:input_type -> group.VolumeGetRequest
	3,  // 9: group.Volume.List:input_type -> group.VolumeListRequest
	4,  // 10: group.Volume.Search:input_type -> group.VolumeSearchRequest
	8,  // 11: group.Volume.Export:input_type -> group.VolumeExportRequest
	12, // 12: group.Volume.Create:output_type -> group.UuidResponse
	12, // 13: group.Volume.Update:output_type -> group.UuidResponse
	12, // 14: group.Volume.Delete:output_type -> group.UuidResponse
	7,  // 15: group.Volume.Get:output_type -> group.VolumeGetResponse
	5,  // 16: group.Volume.List:output_type -> group.VolumeListResponse
	5,  // 17: group.Volume.Search:output_type -> group.VolumeListResponse
	9,  // 18: group.Volume.Export:output_type -> group.VolumeExportResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_catalog_volume_proto_init() }
func file_proto_catalog_volume_proto_init() {
	if File_proto_catalog_volume_proto != nil {
		return
	}
	file_proto_catalog_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_catalog_volume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_volume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_volume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_volume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_volume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeSearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_volume_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_volume_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeGetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_volume_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeGetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_volume_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeExportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_volume_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeExportResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_catalog_volume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_catalog_volume_proto_goTypes,
		DependencyIndexes: file_proto_catalog_volume_proto_depIdxs,
		MessageInfos:      file_proto_catalog_volume_proto_msgTypes,
	}.Build()
	File_proto_catalog_volume_proto = out.File
	file_proto_catalog_volume_proto_rawDesc = nil
	file_proto_catalog_volume_proto_goTypes = nil
	file_proto_catalog_volume_proto_depIdxs = nil
}