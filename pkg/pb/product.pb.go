// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: product.proto

package pb

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

type SearchProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *SearchProductRequest) Reset() {
	*x = SearchProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductRequest) ProtoMessage() {}

func (x *SearchProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductRequest.ProtoReflect.Descriptor instead.
func (*SearchProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *SearchProductRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type SearchProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product []*Product `protobuf:"bytes,1,rep,name=product,proto3" json:"product,omitempty"`
}

func (x *SearchProductResponse) Reset() {
	*x = SearchProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductResponse) ProtoMessage() {}

func (x *SearchProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductResponse.ProtoReflect.Descriptor instead.
func (*SearchProductResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *SearchProductResponse) GetProduct() []*Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type Rom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit SizeUnit `protobuf:"varint,1,opt,name=unit,proto3,enum=pb.SizeUnit" json:"unit,omitempty"`
	Size int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Rom) Reset() {
	*x = Rom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rom) ProtoMessage() {}

func (x *Rom) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rom.ProtoReflect.Descriptor instead.
func (*Rom) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *Rom) GetUnit() SizeUnit {
	if x != nil {
		return x.Unit
	}
	return SizeUnit_UNKNOWN
}

func (x *Rom) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Ram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit SizeUnit `protobuf:"varint,1,opt,name=unit,proto3,enum=pb.SizeUnit" json:"unit,omitempty"`
	Size int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Ram) Reset() {
	*x = Ram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ram) ProtoMessage() {}

func (x *Ram) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ram.ProtoReflect.Descriptor instead.
func (*Ram) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *Ram) GetUnit() SizeUnit {
	if x != nil {
		return x.Unit
	}
	return SizeUnit_UNKNOWN
}

func (x *Ram) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type AdvanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rom *Rom    `protobuf:"bytes,1,opt,name=rom,proto3" json:"rom,omitempty"`
	Ram *Ram    `protobuf:"bytes,2,opt,name=ram,proto3" json:"ram,omitempty"`
	Cpu float32 `protobuf:"fixed32,3,opt,name=cpu,proto3" json:"cpu,omitempty"`
}

func (x *AdvanceInfo) Reset() {
	*x = AdvanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvanceInfo) ProtoMessage() {}

func (x *AdvanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvanceInfo.ProtoReflect.Descriptor instead.
func (*AdvanceInfo) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *AdvanceInfo) GetRom() *Rom {
	if x != nil {
		return x.Rom
	}
	return nil
}

func (x *AdvanceInfo) GetRam() *Ram {
	if x != nil {
		return x.Ram
	}
	return nil
}

func (x *AdvanceInfo) GetCpu() float32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string       `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Color       string       `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	Producer    string       `protobuf:"bytes,3,opt,name=producer,proto3" json:"producer,omitempty"`
	Series      string       `protobuf:"bytes,4,opt,name=series,proto3" json:"series,omitempty"`
	AdvanceInfo *AdvanceInfo `protobuf:"bytes,5,opt,name=advance_info,json=advanceInfo,proto3" json:"advance_info,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *Product) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Product) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Product) GetProducer() string {
	if x != nil {
		return x.Producer
	}
	return ""
}

func (x *Product) GetSeries() string {
	if x != nil {
		return x.Series
	}
	return ""
}

func (x *Product) GetAdvanceInfo() *AdvanceInfo {
	if x != nil {
		return x.AdvanceInfo
	}
	return nil
}

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{6}
}

func (x *CreateProductRequest) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type ClientStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *ClientStreamRequest) Reset() {
	*x = ClientStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamRequest) ProtoMessage() {}

func (x *ClientStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamRequest.ProtoReflect.Descriptor instead.
func (*ClientStreamRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{7}
}

func (x *ClientStreamRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ClientStreamRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4d, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3e, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0x3b, 0x0a, 0x03, 0x52, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x55, 0x6e,
	0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3b, 0x0a, 0x03,
	0x52, 0x61, 0x6d, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x55, 0x0a, 0x0b, 0x41, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x03, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6d, 0x52, 0x03,
	0x72, 0x6f, 0x6d, 0x12, 0x19, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61, 0x6d, 0x52, 0x03, 0x72, 0x61, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x70, 0x75,
	0x22, 0x97, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0c, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3f, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x3b, 0x0a, 0x13, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x6f, 0x63, 0x62, 0x61, 0x6e, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_product_proto_goTypes = []interface{}{
	(*SearchProductRequest)(nil),  // 0: pb.SearchProductRequest
	(*SearchProductResponse)(nil), // 1: pb.SearchProductResponse
	(*Rom)(nil),                   // 2: pb.Rom
	(*Ram)(nil),                   // 3: pb.Ram
	(*AdvanceInfo)(nil),           // 4: pb.AdvanceInfo
	(*Product)(nil),               // 5: pb.Product
	(*CreateProductRequest)(nil),  // 6: pb.CreateProductRequest
	(*ClientStreamRequest)(nil),   // 7: pb.ClientStreamRequest
	(*PaginationRequest)(nil),     // 8: pb.PaginationRequest
	(SizeUnit)(0),                 // 9: pb.SizeUnit
}
var file_product_proto_depIdxs = []int32{
	8, // 0: pb.SearchProductRequest.pagination:type_name -> pb.PaginationRequest
	5, // 1: pb.SearchProductResponse.product:type_name -> pb.Product
	9, // 2: pb.Rom.unit:type_name -> pb.SizeUnit
	9, // 3: pb.Ram.unit:type_name -> pb.SizeUnit
	2, // 4: pb.AdvanceInfo.rom:type_name -> pb.Rom
	3, // 5: pb.AdvanceInfo.ram:type_name -> pb.Ram
	4, // 6: pb.Product.advance_info:type_name -> pb.AdvanceInfo
	5, // 7: pb.CreateProductRequest.products:type_name -> pb.Product
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	file_pagination_proto_init()
	file_types_proto_init()
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProductRequest); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProductResponse); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rom); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ram); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvanceInfo); i {
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
		file_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamRequest); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
