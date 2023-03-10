// hero/hero.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: hero/hero.proto

package hero

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

type HeroById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *HeroById) Reset() {
	*x = HeroById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_hero_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeroById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeroById) ProtoMessage() {}

func (x *HeroById) ProtoReflect() protoreflect.Message {
	mi := &file_hero_hero_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeroById.ProtoReflect.Descriptor instead.
func (*HeroById) Descriptor() ([]byte, []int) {
	return file_hero_hero_proto_rawDescGZIP(), []int{0}
}

func (x *HeroById) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Hero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Hero) Reset() {
	*x = Hero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_hero_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero) ProtoMessage() {}

func (x *Hero) ProtoReflect() protoreflect.Message {
	mi := &file_hero_hero_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero.ProtoReflect.Descriptor instead.
func (*Hero) Descriptor() ([]byte, []int) {
	return file_hero_hero_proto_rawDescGZIP(), []int{1}
}

func (x *Hero) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Hero) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ServerStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ServerStreamRequest) Reset() {
	*x = ServerStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_hero_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamRequest) ProtoMessage() {}

func (x *ServerStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hero_hero_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamRequest.ProtoReflect.Descriptor instead.
func (*ServerStreamRequest) Descriptor() ([]byte, []int) {
	return file_hero_hero_proto_rawDescGZIP(), []int{2}
}

func (x *ServerStreamRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ServerStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ServerStreamResponse) Reset() {
	*x = ServerStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_hero_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamResponse) ProtoMessage() {}

func (x *ServerStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hero_hero_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamResponse.ProtoReflect.Descriptor instead.
func (*ServerStreamResponse) Descriptor() ([]byte, []int) {
	return file_hero_hero_proto_rawDescGZIP(), []int{3}
}

func (x *ServerStreamResponse) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ClientStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ClientStreamRequest) Reset() {
	*x = ClientStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_hero_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamRequest) ProtoMessage() {}

func (x *ClientStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hero_hero_proto_msgTypes[4]
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
	return file_hero_hero_proto_rawDescGZIP(), []int{4}
}

func (x *ClientStreamRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ClientStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ClientStreamResponse) Reset() {
	*x = ClientStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_hero_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamResponse) ProtoMessage() {}

func (x *ClientStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hero_hero_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamResponse.ProtoReflect.Descriptor instead.
func (*ClientStreamResponse) Descriptor() ([]byte, []int) {
	return file_hero_hero_proto_rawDescGZIP(), []int{5}
}

func (x *ClientStreamResponse) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type TwoWayStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *TwoWayStreamRequest) Reset() {
	*x = TwoWayStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_hero_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwoWayStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwoWayStreamRequest) ProtoMessage() {}

func (x *TwoWayStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hero_hero_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwoWayStreamRequest.ProtoReflect.Descriptor instead.
func (*TwoWayStreamRequest) Descriptor() ([]byte, []int) {
	return file_hero_hero_proto_rawDescGZIP(), []int{6}
}

func (x *TwoWayStreamRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type TwoWayStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *TwoWayStreamResponse) Reset() {
	*x = TwoWayStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_hero_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwoWayStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwoWayStreamResponse) ProtoMessage() {}

func (x *TwoWayStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hero_hero_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwoWayStreamResponse.ProtoReflect.Descriptor instead.
func (*TwoWayStreamResponse) Descriptor() ([]byte, []int) {
	return file_hero_hero_proto_rawDescGZIP(), []int{7}
}

func (x *TwoWayStreamResponse) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_hero_hero_proto protoreflect.FileDescriptor

var file_hero_hero_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x65, 0x72, 0x6f, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x48, 0x65, 0x72, 0x6f, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x04, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x27, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x28, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x22, 0x27, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x28, 0x0a, 0x14, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x27, 0x0a, 0x13, 0x54, 0x77, 0x6f, 0x57, 0x61, 0x79, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x28,
	0x0a, 0x14, 0x54, 0x77, 0x6f, 0x57, 0x61, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x32, 0x9b, 0x02, 0x0a, 0x0d, 0x48, 0x65, 0x72,
	0x6f, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x46, 0x69,
	0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x0e, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x48, 0x65, 0x72,
	0x6f, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x48, 0x65, 0x72,
	0x6f, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x49,
	0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19,
	0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x65, 0x72, 0x6f,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4b, 0x0a, 0x0c, 0x54, 0x77, 0x6f,
	0x57, 0x61, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x72, 0x6f,
	0x2e, 0x54, 0x77, 0x6f, 0x57, 0x61, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x54, 0x77, 0x6f, 0x57,
	0x61, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x68, 0x65, 0x72, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hero_hero_proto_rawDescOnce sync.Once
	file_hero_hero_proto_rawDescData = file_hero_hero_proto_rawDesc
)

func file_hero_hero_proto_rawDescGZIP() []byte {
	file_hero_hero_proto_rawDescOnce.Do(func() {
		file_hero_hero_proto_rawDescData = protoimpl.X.CompressGZIP(file_hero_hero_proto_rawDescData)
	})
	return file_hero_hero_proto_rawDescData
}

var file_hero_hero_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_hero_hero_proto_goTypes = []interface{}{
	(*HeroById)(nil),             // 0: hero.HeroById
	(*Hero)(nil),                 // 1: hero.Hero
	(*ServerStreamRequest)(nil),  // 2: hero.ServerStreamRequest
	(*ServerStreamResponse)(nil), // 3: hero.ServerStreamResponse
	(*ClientStreamRequest)(nil),  // 4: hero.ClientStreamRequest
	(*ClientStreamResponse)(nil), // 5: hero.ClientStreamResponse
	(*TwoWayStreamRequest)(nil),  // 6: hero.TwoWayStreamRequest
	(*TwoWayStreamResponse)(nil), // 7: hero.TwoWayStreamResponse
}
var file_hero_hero_proto_depIdxs = []int32{
	0, // 0: hero.HeroesService.FindOne:input_type -> hero.HeroById
	2, // 1: hero.HeroesService.ServerStream:input_type -> hero.ServerStreamRequest
	4, // 2: hero.HeroesService.ClientStream:input_type -> hero.ClientStreamRequest
	6, // 3: hero.HeroesService.TwoWayStream:input_type -> hero.TwoWayStreamRequest
	1, // 4: hero.HeroesService.FindOne:output_type -> hero.Hero
	3, // 5: hero.HeroesService.ServerStream:output_type -> hero.ServerStreamResponse
	5, // 6: hero.HeroesService.ClientStream:output_type -> hero.ClientStreamResponse
	7, // 7: hero.HeroesService.TwoWayStream:output_type -> hero.TwoWayStreamResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hero_hero_proto_init() }
func file_hero_hero_proto_init() {
	if File_hero_hero_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hero_hero_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeroById); i {
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
		file_hero_hero_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hero); i {
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
		file_hero_hero_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamRequest); i {
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
		file_hero_hero_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamResponse); i {
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
		file_hero_hero_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_hero_hero_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamResponse); i {
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
		file_hero_hero_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwoWayStreamRequest); i {
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
		file_hero_hero_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwoWayStreamResponse); i {
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
			RawDescriptor: file_hero_hero_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hero_hero_proto_goTypes,
		DependencyIndexes: file_hero_hero_proto_depIdxs,
		MessageInfos:      file_hero_hero_proto_msgTypes,
	}.Build()
	File_hero_hero_proto = out.File
	file_hero_hero_proto_rawDesc = nil
	file_hero_hero_proto_goTypes = nil
	file_hero_hero_proto_depIdxs = nil
}
