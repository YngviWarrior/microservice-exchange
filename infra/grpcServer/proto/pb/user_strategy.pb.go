// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: user_strategy.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserStrategy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserStrategy  uint64                 `protobuf:"varint,1,opt,name=user_strategy,json=userStrategy,proto3" json:"user_strategy,omitempty"`
	User          uint64                 `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	TradeConfig   string                 `protobuf:"bytes,3,opt,name=trade_config,json=tradeConfig,proto3" json:"trade_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserStrategy) Reset() {
	*x = UserStrategy{}
	mi := &file_user_strategy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStrategy) ProtoMessage() {}

func (x *UserStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_user_strategy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStrategy.ProtoReflect.Descriptor instead.
func (*UserStrategy) Descriptor() ([]byte, []int) {
	return file_user_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *UserStrategy) GetUserStrategy() uint64 {
	if x != nil {
		return x.UserStrategy
	}
	return 0
}

func (x *UserStrategy) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *UserStrategy) GetTradeConfig() string {
	if x != nil {
		return x.TradeConfig
	}
	return ""
}

type UserStrategyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserStrategy  []*UserStrategy        `protobuf:"bytes,1,rep,name=user_strategy,json=userStrategy,proto3" json:"user_strategy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserStrategyResponse) Reset() {
	*x = UserStrategyResponse{}
	mi := &file_user_strategy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserStrategyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStrategyResponse) ProtoMessage() {}

func (x *UserStrategyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_strategy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStrategyResponse.ProtoReflect.Descriptor instead.
func (*UserStrategyResponse) Descriptor() ([]byte, []int) {
	return file_user_strategy_proto_rawDescGZIP(), []int{1}
}

func (x *UserStrategyResponse) GetUserStrategy() []*UserStrategy {
	if x != nil {
		return x.UserStrategy
	}
	return nil
}

type CreateUserStrategyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          uint64                 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	TradeConfig   string                 `protobuf:"bytes,2,opt,name=trade_config,json=tradeConfig,proto3" json:"trade_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserStrategyRequest) Reset() {
	*x = CreateUserStrategyRequest{}
	mi := &file_user_strategy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserStrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserStrategyRequest) ProtoMessage() {}

func (x *CreateUserStrategyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_strategy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserStrategyRequest.ProtoReflect.Descriptor instead.
func (*CreateUserStrategyRequest) Descriptor() ([]byte, []int) {
	return file_user_strategy_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserStrategyRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CreateUserStrategyRequest) GetTradeConfig() string {
	if x != nil {
		return x.TradeConfig
	}
	return ""
}

type ListUserStrategyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserStrategyRequest) Reset() {
	*x = ListUserStrategyRequest{}
	mi := &file_user_strategy_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserStrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserStrategyRequest) ProtoMessage() {}

func (x *ListUserStrategyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_strategy_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserStrategyRequest.ProtoReflect.Descriptor instead.
func (*ListUserStrategyRequest) Descriptor() ([]byte, []int) {
	return file_user_strategy_proto_rawDescGZIP(), []int{3}
}

var File_user_strategy_proto protoreflect.FileDescriptor

const file_user_strategy_proto_rawDesc = "" +
	"\n" +
	"\x13user_strategy.proto\x12\x02pb\"j\n" +
	"\fUserStrategy\x12#\n" +
	"\ruser_strategy\x18\x01 \x01(\x04R\fuserStrategy\x12\x12\n" +
	"\x04user\x18\x02 \x01(\x04R\x04user\x12!\n" +
	"\ftrade_config\x18\x03 \x01(\tR\vtradeConfig\"M\n" +
	"\x14UserStrategyResponse\x125\n" +
	"\ruser_strategy\x18\x01 \x03(\v2\x10.pb.UserStrategyR\fuserStrategy\"R\n" +
	"\x19CreateUserStrategyRequest\x12\x12\n" +
	"\x04user\x18\x01 \x01(\x04R\x04user\x12!\n" +
	"\ftrade_config\x18\x02 \x01(\tR\vtradeConfig\"\x19\n" +
	"\x17ListUserStrategyRequestB\x06Z\x04./pbb\x06proto3"

var (
	file_user_strategy_proto_rawDescOnce sync.Once
	file_user_strategy_proto_rawDescData []byte
)

func file_user_strategy_proto_rawDescGZIP() []byte {
	file_user_strategy_proto_rawDescOnce.Do(func() {
		file_user_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_strategy_proto_rawDesc), len(file_user_strategy_proto_rawDesc)))
	})
	return file_user_strategy_proto_rawDescData
}

var file_user_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_strategy_proto_goTypes = []any{
	(*UserStrategy)(nil),              // 0: pb.UserStrategy
	(*UserStrategyResponse)(nil),      // 1: pb.UserStrategyResponse
	(*CreateUserStrategyRequest)(nil), // 2: pb.CreateUserStrategyRequest
	(*ListUserStrategyRequest)(nil),   // 3: pb.ListUserStrategyRequest
}
var file_user_strategy_proto_depIdxs = []int32{
	0, // 0: pb.UserStrategyResponse.user_strategy:type_name -> pb.UserStrategy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_strategy_proto_init() }
func file_user_strategy_proto_init() {
	if File_user_strategy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_strategy_proto_rawDesc), len(file_user_strategy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_strategy_proto_goTypes,
		DependencyIndexes: file_user_strategy_proto_depIdxs,
		MessageInfos:      file_user_strategy_proto_msgTypes,
	}.Build()
	File_user_strategy_proto = out.File
	file_user_strategy_proto_goTypes = nil
	file_user_strategy_proto_depIdxs = nil
}
