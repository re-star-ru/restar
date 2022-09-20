// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: api/proto/v1/diagnostic.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Item_ItemType int32

const (
	Item_ITEM_TYPE_UNKNOWN Item_ItemType = 0
	Item_ITEM_TYPE_PRODUCT Item_ItemType = 1
	Item_ITEM_TYPE_SERVICE Item_ItemType = 2
)

// Enum value maps for Item_ItemType.
var (
	Item_ItemType_name = map[int32]string{
		0: "ITEM_TYPE_UNKNOWN",
		1: "ITEM_TYPE_PRODUCT",
		2: "ITEM_TYPE_SERVICE",
	}
	Item_ItemType_value = map[string]int32{
		"ITEM_TYPE_UNKNOWN": 0,
		"ITEM_TYPE_PRODUCT": 1,
		"ITEM_TYPE_SERVICE": 2,
	}
)

func (x Item_ItemType) Enum() *Item_ItemType {
	p := new(Item_ItemType)
	*p = x
	return p
}

func (x Item_ItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Item_ItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_v1_diagnostic_proto_enumTypes[0].Descriptor()
}

func (Item_ItemType) Type() protoreflect.EnumType {
	return &file_api_proto_v1_diagnostic_proto_enumTypes[0]
}

func (x Item_ItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Item_ItemType.Descriptor instead.
func (Item_ItemType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_v1_diagnostic_proto_rawDescGZIP(), []int{1, 0}
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_diagnostic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_diagnostic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_diagnostic_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemType Item_ItemType `protobuf:"varint,2,opt,name=itemType,proto3,enum=Item_ItemType" json:"itemType,omitempty"`
	Name     string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_diagnostic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_diagnostic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_diagnostic_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetItemType() Item_ItemType {
	if x != nil {
		return x.ItemType
	}
	return Item_ITEM_TYPE_UNKNOWN
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Diagnostic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Version       uint32                 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DefinedNumber string                 `protobuf:"bytes,5,opt,name=definedNumber,proto3" json:"definedNumber,omitempty"`
	SKU           string                 `protobuf:"bytes,6,opt,name=SKU,proto3" json:"SKU,omitempty"`
	Items         []*Item                `protobuf:"bytes,7,rep,name=items,proto3" json:"items,omitempty"`
	Images        []*Image               `protobuf:"bytes,8,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *Diagnostic) Reset() {
	*x = Diagnostic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_diagnostic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Diagnostic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic) ProtoMessage() {}

func (x *Diagnostic) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_diagnostic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic.ProtoReflect.Descriptor instead.
func (*Diagnostic) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_diagnostic_proto_rawDescGZIP(), []int{2}
}

func (x *Diagnostic) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Diagnostic) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Diagnostic) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Diagnostic) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Diagnostic) GetDefinedNumber() string {
	if x != nil {
		return x.DefinedNumber
	}
	return ""
}

func (x *Diagnostic) GetSKU() string {
	if x != nil {
		return x.SKU
	}
	return ""
}

func (x *Diagnostic) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Diagnostic) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type DiagnosticList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Diagnostic `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *DiagnosticList) Reset() {
	*x = DiagnosticList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_diagnostic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiagnosticList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiagnosticList) ProtoMessage() {}

func (x *DiagnosticList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_diagnostic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiagnosticList.ProtoReflect.Descriptor instead.
func (*DiagnosticList) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_diagnostic_proto_rawDescGZIP(), []int{3}
}

func (x *DiagnosticList) GetList() []*Diagnostic {
	if x != nil {
		return x.List
	}
	return nil
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_diagnostic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_diagnostic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_diagnostic_proto_rawDescGZIP(), []int{4}
}

func (x *ID) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_proto_v1_diagnostic_proto protoreflect.FileDescriptor

var file_api_proto_v1_diagnostic_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xa7,
	0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x22, 0x9f, 0x02, 0x0a, 0x0a, 0x44, 0x69, 0x61,
	0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x53,
	0x4b, 0x55, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x4b, 0x55, 0x12, 0x1b, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x0e, 0x44, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x44, 0x69, 0x61,
	0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x14, 0x0a,
	0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xb9, 0x01, 0x0a, 0x11, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x44, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x04, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0b, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x1a, 0x0b, 0x2e,
	0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x44,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42,
	0x16, 0x5a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_diagnostic_proto_rawDescOnce sync.Once
	file_api_proto_v1_diagnostic_proto_rawDescData = file_api_proto_v1_diagnostic_proto_rawDesc
)

func file_api_proto_v1_diagnostic_proto_rawDescGZIP() []byte {
	file_api_proto_v1_diagnostic_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_diagnostic_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_diagnostic_proto_rawDescData)
	})
	return file_api_proto_v1_diagnostic_proto_rawDescData
}

var file_api_proto_v1_diagnostic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_v1_diagnostic_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_v1_diagnostic_proto_goTypes = []interface{}{
	(Item_ItemType)(0),            // 0: Item.ItemType
	(*Image)(nil),                 // 1: Image
	(*Item)(nil),                  // 2: Item
	(*Diagnostic)(nil),            // 3: Diagnostic
	(*DiagnosticList)(nil),        // 4: DiagnosticList
	(*ID)(nil),                    // 5: ID
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_api_proto_v1_diagnostic_proto_depIdxs = []int32{
	0,  // 0: Item.itemType:type_name -> Item.ItemType
	6,  // 1: Diagnostic.createdAt:type_name -> google.protobuf.Timestamp
	6,  // 2: Diagnostic.updatedAt:type_name -> google.protobuf.Timestamp
	2,  // 3: Diagnostic.items:type_name -> Item
	1,  // 4: Diagnostic.images:type_name -> Image
	3,  // 5: DiagnosticList.list:type_name -> Diagnostic
	7,  // 6: DiagnosticService.Create:input_type -> google.protobuf.Empty
	5,  // 7: DiagnosticService.Read:input_type -> ID
	3,  // 8: DiagnosticService.Update:input_type -> Diagnostic
	7,  // 9: DiagnosticService.List:input_type -> google.protobuf.Empty
	3,  // 10: DiagnosticService.Create:output_type -> Diagnostic
	3,  // 11: DiagnosticService.Read:output_type -> Diagnostic
	3,  // 12: DiagnosticService.Update:output_type -> Diagnostic
	4,  // 13: DiagnosticService.List:output_type -> DiagnosticList
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_v1_diagnostic_proto_init() }
func file_api_proto_v1_diagnostic_proto_init() {
	if File_api_proto_v1_diagnostic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_diagnostic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_api_proto_v1_diagnostic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_api_proto_v1_diagnostic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Diagnostic); i {
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
		file_api_proto_v1_diagnostic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiagnosticList); i {
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
		file_api_proto_v1_diagnostic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
			RawDescriptor: file_api_proto_v1_diagnostic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_diagnostic_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_diagnostic_proto_depIdxs,
		EnumInfos:         file_api_proto_v1_diagnostic_proto_enumTypes,
		MessageInfos:      file_api_proto_v1_diagnostic_proto_msgTypes,
	}.Build()
	File_api_proto_v1_diagnostic_proto = out.File
	file_api_proto_v1_diagnostic_proto_rawDesc = nil
	file_api_proto_v1_diagnostic_proto_goTypes = nil
	file_api_proto_v1_diagnostic_proto_depIdxs = nil
}
