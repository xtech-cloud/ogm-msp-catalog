// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/catalog/path.proto

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

// 添加的请求
type PathAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume   string            `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`                                                                                                              // 卷的uuid
	Uri      string            `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`                                                                                                                    // 资源标识
	Name     string            `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`                                                                                                                 // 名称
	NameI18N map[string]string `protobuf:"bytes,11,rep,name=name_i18n,json=nameI18n,proto3" json:"name_i18n,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 名称的多国语言
	Label    []string          `protobuf:"bytes,20,rep,name=label,proto3" json:"label,omitempty"`                                                                                                               // 预设标签列表
	Tag      []string          `protobuf:"bytes,21,rep,name=tag,proto3" json:"tag,omitempty"`                                                                                                                   // 自定义标签列表
	Resource []string          `protobuf:"bytes,30,rep,name=resource,proto3" json:"resource,omitempty"`                                                                                                         // 内容列表
}

func (x *PathAddRequest) Reset() {
	*x = PathAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_path_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathAddRequest) ProtoMessage() {}

func (x *PathAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_path_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathAddRequest.ProtoReflect.Descriptor instead.
func (*PathAddRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_path_proto_rawDescGZIP(), []int{0}
}

func (x *PathAddRequest) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *PathAddRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *PathAddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PathAddRequest) GetNameI18N() map[string]string {
	if x != nil {
		return x.NameI18N
	}
	return nil
}

func (x *PathAddRequest) GetLabel() []string {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *PathAddRequest) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *PathAddRequest) GetResource() []string {
	if x != nil {
		return x.Resource
	}
	return nil
}

// 更新的请求
type PathUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`                                                                                                                  // uuid
	Uri      string            `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`                                                                                                                    // 资源标识
	Name     string            `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`                                                                                                                 // 名称
	NameI18N map[string]string `protobuf:"bytes,11,rep,name=name_i18n,json=nameI18n,proto3" json:"name_i18n,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 名称的多国语言
	Label    []string          `protobuf:"bytes,20,rep,name=label,proto3" json:"label,omitempty"`                                                                                                               // 预设标签列表
	Tag      []string          `protobuf:"bytes,21,rep,name=tag,proto3" json:"tag,omitempty"`                                                                                                                   // 自定义标签列表
	Resource []string          `protobuf:"bytes,30,rep,name=resource,proto3" json:"resource,omitempty"`                                                                                                         // 内容列表
}

func (x *PathUpdateRequest) Reset() {
	*x = PathUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_path_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathUpdateRequest) ProtoMessage() {}

func (x *PathUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_path_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathUpdateRequest.ProtoReflect.Descriptor instead.
func (*PathUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_path_proto_rawDescGZIP(), []int{1}
}

func (x *PathUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PathUpdateRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *PathUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PathUpdateRequest) GetNameI18N() map[string]string {
	if x != nil {
		return x.NameI18N
	}
	return nil
}

func (x *PathUpdateRequest) GetLabel() []string {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *PathUpdateRequest) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *PathUpdateRequest) GetResource() []string {
	if x != nil {
		return x.Resource
	}
	return nil
}

// 删除的请求
type PathDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
}

func (x *PathDeleteRequest) Reset() {
	*x = PathDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_path_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathDeleteRequest) ProtoMessage() {}

func (x *PathDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_path_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathDeleteRequest.ProtoReflect.Descriptor instead.
func (*PathDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_path_proto_rawDescGZIP(), []int{2}
}

func (x *PathDeleteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 列举的请求
type PathListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
}

func (x *PathListRequest) Reset() {
	*x = PathListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_path_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathListRequest) ProtoMessage() {}

func (x *PathListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_path_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathListRequest.ProtoReflect.Descriptor instead.
func (*PathListRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_path_proto_rawDescGZIP(), []int{3}
}

func (x *PathListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PathListRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 搜索的请求
type PathSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
	Volume string `protobuf:"bytes,3,opt,name=volume,proto3" json:"volume,omitempty"`  // 卷的uuid
	Uri    string `protobuf:"bytes,10,opt,name=uri,proto3" json:"uri,omitempty"`       // 资源标识
	Name   string `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`     // 名称
}

func (x *PathSearchRequest) Reset() {
	*x = PathSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_path_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathSearchRequest) ProtoMessage() {}

func (x *PathSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_path_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathSearchRequest.ProtoReflect.Descriptor instead.
func (*PathSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_path_proto_rawDescGZIP(), []int{4}
}

func (x *PathSearchRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PathSearchRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PathSearchRequest) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *PathSearchRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *PathSearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 列举的回复
type PathListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  uint64        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  // 总数
	Entity []*PathEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"` // 实体列表
}

func (x *PathListResponse) Reset() {
	*x = PathListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_path_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathListResponse) ProtoMessage() {}

func (x *PathListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_path_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathListResponse.ProtoReflect.Descriptor instead.
func (*PathListResponse) Descriptor() ([]byte, []int) {
	return file_proto_catalog_path_proto_rawDescGZIP(), []int{5}
}

func (x *PathListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *PathListResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PathListResponse) GetEntity() []*PathEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

// 获取的请求
type PathGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
}

func (x *PathGetRequest) Reset() {
	*x = PathGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_path_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathGetRequest) ProtoMessage() {}

func (x *PathGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_path_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathGetRequest.ProtoReflect.Descriptor instead.
func (*PathGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_path_proto_rawDescGZIP(), []int{6}
}

func (x *PathGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 获取的回复
type PathGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Entity *PathEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"` // 实体
}

func (x *PathGetResponse) Reset() {
	*x = PathGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_path_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathGetResponse) ProtoMessage() {}

func (x *PathGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_path_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathGetResponse.ProtoReflect.Descriptor instead.
func (*PathGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_catalog_path_proto_rawDescGZIP(), []int{7}
}

func (x *PathGetResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *PathGetResponse) GetEntity() *PathEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

var File_proto_catalog_path_proto protoreflect.FileDescriptor

var file_proto_catalog_path_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f,
	0x70, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02,
	0x0a, 0x0e, 0x50, 0x61, 0x74, 0x68, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x31, 0x38, 0x6e, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x31, 0x38,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x31, 0x38, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x15, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x31, 0x38, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x93, 0x02, 0x0a, 0x11, 0x50, 0x61, 0x74, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x43, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x31, 0x38, 0x6e, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74,
	0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x49, 0x31, 0x38, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6e, 0x61,
	0x6d, 0x65, 0x49, 0x31, 0x38, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x61, 0x67, 0x18, 0x15, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x4e, 0x61,
	0x6d, 0x65, 0x49, 0x31, 0x38, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x27, 0x0a, 0x11, 0x50, 0x61, 0x74, 0x68, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x3f, 0x0a, 0x0f, 0x50, 0x61, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x7f, 0x0a, 0x11, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x7a, 0x0a, 0x10, 0x50, 0x61, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x24,
	0x0a, 0x0e, 0x50, 0x61, 0x74, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x0f, 0x50, 0x61, 0x74, 0x68, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xe3, 0x02, 0x0a, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x33, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x18, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55,
	0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74,
	0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61,
	0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x74,
	0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x17, 0x5a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x3b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_catalog_path_proto_rawDescOnce sync.Once
	file_proto_catalog_path_proto_rawDescData = file_proto_catalog_path_proto_rawDesc
)

func file_proto_catalog_path_proto_rawDescGZIP() []byte {
	file_proto_catalog_path_proto_rawDescOnce.Do(func() {
		file_proto_catalog_path_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_catalog_path_proto_rawDescData)
	})
	return file_proto_catalog_path_proto_rawDescData
}

var file_proto_catalog_path_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_catalog_path_proto_goTypes = []interface{}{
	(*PathAddRequest)(nil),    // 0: group.PathAddRequest
	(*PathUpdateRequest)(nil), // 1: group.PathUpdateRequest
	(*PathDeleteRequest)(nil), // 2: group.PathDeleteRequest
	(*PathListRequest)(nil),   // 3: group.PathListRequest
	(*PathSearchRequest)(nil), // 4: group.PathSearchRequest
	(*PathListResponse)(nil),  // 5: group.PathListResponse
	(*PathGetRequest)(nil),    // 6: group.PathGetRequest
	(*PathGetResponse)(nil),   // 7: group.PathGetResponse
	nil,                       // 8: group.PathAddRequest.NameI18nEntry
	nil,                       // 9: group.PathUpdateRequest.NameI18nEntry
	(*Status)(nil),            // 10: group.Status
	(*PathEntity)(nil),        // 11: group.PathEntity
	(*UuidResponse)(nil),      // 12: group.UuidResponse
}
var file_proto_catalog_path_proto_depIdxs = []int32{
	8,  // 0: group.PathAddRequest.name_i18n:type_name -> group.PathAddRequest.NameI18nEntry
	9,  // 1: group.PathUpdateRequest.name_i18n:type_name -> group.PathUpdateRequest.NameI18nEntry
	10, // 2: group.PathListResponse.status:type_name -> group.Status
	11, // 3: group.PathListResponse.entity:type_name -> group.PathEntity
	10, // 4: group.PathGetResponse.status:type_name -> group.Status
	11, // 5: group.PathGetResponse.entity:type_name -> group.PathEntity
	0,  // 6: group.Path.Add:input_type -> group.PathAddRequest
	1,  // 7: group.Path.Update:input_type -> group.PathUpdateRequest
	2,  // 8: group.Path.Remove:input_type -> group.PathDeleteRequest
	6,  // 9: group.Path.Get:input_type -> group.PathGetRequest
	3,  // 10: group.Path.List:input_type -> group.PathListRequest
	4,  // 11: group.Path.Search:input_type -> group.PathSearchRequest
	12, // 12: group.Path.Add:output_type -> group.UuidResponse
	12, // 13: group.Path.Update:output_type -> group.UuidResponse
	12, // 14: group.Path.Remove:output_type -> group.UuidResponse
	7,  // 15: group.Path.Get:output_type -> group.PathGetResponse
	5,  // 16: group.Path.List:output_type -> group.PathListResponse
	5,  // 17: group.Path.Search:output_type -> group.PathListResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_catalog_path_proto_init() }
func file_proto_catalog_path_proto_init() {
	if File_proto_catalog_path_proto != nil {
		return
	}
	file_proto_catalog_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_catalog_path_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathAddRequest); i {
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
		file_proto_catalog_path_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathUpdateRequest); i {
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
		file_proto_catalog_path_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathDeleteRequest); i {
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
		file_proto_catalog_path_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathListRequest); i {
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
		file_proto_catalog_path_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathSearchRequest); i {
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
		file_proto_catalog_path_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathListResponse); i {
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
		file_proto_catalog_path_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathGetRequest); i {
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
		file_proto_catalog_path_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathGetResponse); i {
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
			RawDescriptor: file_proto_catalog_path_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_catalog_path_proto_goTypes,
		DependencyIndexes: file_proto_catalog_path_proto_depIdxs,
		MessageInfos:      file_proto_catalog_path_proto_msgTypes,
	}.Build()
	File_proto_catalog_path_proto = out.File
	file_proto_catalog_path_proto_rawDesc = nil
	file_proto_catalog_path_proto_goTypes = nil
	file_proto_catalog_path_proto_depIdxs = nil
}
