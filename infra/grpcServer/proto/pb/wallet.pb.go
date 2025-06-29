// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: wallet.proto

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

type Wallet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wallet        uint64                 `protobuf:"varint,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Coin          uint64                 `protobuf:"varint,3,opt,name=coin,proto3" json:"coin,omitempty"`
	Amount        string                 `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	mi := &file_wallet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *Wallet) GetWallet() uint64 {
	if x != nil {
		return x.Wallet
	}
	return 0
}

func (x *Wallet) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *Wallet) GetCoin() uint64 {
	if x != nil {
		return x.Coin
	}
	return 0
}

func (x *Wallet) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type WalletWithCoin struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wallet        uint64                 `protobuf:"varint,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Coin          uint64                 `protobuf:"varint,3,opt,name=coin,proto3" json:"coin,omitempty"`
	Amount        string                 `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Symbol        string                 `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Active        bool                   `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WalletWithCoin) Reset() {
	*x = WalletWithCoin{}
	mi := &file_wallet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WalletWithCoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletWithCoin) ProtoMessage() {}

func (x *WalletWithCoin) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletWithCoin.ProtoReflect.Descriptor instead.
func (*WalletWithCoin) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *WalletWithCoin) GetWallet() uint64 {
	if x != nil {
		return x.Wallet
	}
	return 0
}

func (x *WalletWithCoin) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *WalletWithCoin) GetCoin() uint64 {
	if x != nil {
		return x.Coin
	}
	return 0
}

func (x *WalletWithCoin) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *WalletWithCoin) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *WalletWithCoin) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type GetWalletWithCoinResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wallet        *Wallet                `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Coin          uint64                 `protobuf:"varint,2,opt,name=coin,proto3" json:"coin,omitempty"`
	Symbol        string                 `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Active        bool                   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWalletWithCoinResponse) Reset() {
	*x = GetWalletWithCoinResponse{}
	mi := &file_wallet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWalletWithCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletWithCoinResponse) ProtoMessage() {}

func (x *GetWalletWithCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletWithCoinResponse.ProtoReflect.Descriptor instead.
func (*GetWalletWithCoinResponse) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *GetWalletWithCoinResponse) GetWallet() *Wallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

func (x *GetWalletWithCoinResponse) GetCoin() uint64 {
	if x != nil {
		return x.Coin
	}
	return 0
}

func (x *GetWalletWithCoinResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *GetWalletWithCoinResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type GetWalletWithCoinRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exchange      uint64                 `protobuf:"varint,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Coin          uint64                 `protobuf:"varint,2,opt,name=coin,proto3" json:"coin,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWalletWithCoinRequest) Reset() {
	*x = GetWalletWithCoinRequest{}
	mi := &file_wallet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWalletWithCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletWithCoinRequest) ProtoMessage() {}

func (x *GetWalletWithCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletWithCoinRequest.ProtoReflect.Descriptor instead.
func (*GetWalletWithCoinRequest) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *GetWalletWithCoinRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *GetWalletWithCoinRequest) GetCoin() uint64 {
	if x != nil {
		return x.Coin
	}
	return 0
}

type ListWalletWithCoinResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wallet        []*WalletWithCoin      `protobuf:"bytes,1,rep,name=wallet,proto3" json:"wallet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListWalletWithCoinResponse) Reset() {
	*x = ListWalletWithCoinResponse{}
	mi := &file_wallet_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListWalletWithCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWalletWithCoinResponse) ProtoMessage() {}

func (x *ListWalletWithCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWalletWithCoinResponse.ProtoReflect.Descriptor instead.
func (*ListWalletWithCoinResponse) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *ListWalletWithCoinResponse) GetWallet() []*WalletWithCoin {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type ListWalletWithCoinRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exchange      uint64                 `protobuf:"varint,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListWalletWithCoinRequest) Reset() {
	*x = ListWalletWithCoinRequest{}
	mi := &file_wallet_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListWalletWithCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWalletWithCoinRequest) ProtoMessage() {}

func (x *ListWalletWithCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWalletWithCoinRequest.ProtoReflect.Descriptor instead.
func (*ListWalletWithCoinRequest) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *ListWalletWithCoinRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

type CreateWalletRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wallet        uint64                 `protobuf:"varint,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Coin          uint64                 `protobuf:"varint,3,opt,name=coin,proto3" json:"coin,omitempty"`
	Amount        string                 `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWalletRequest) Reset() {
	*x = CreateWalletRequest{}
	mi := &file_wallet_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRequest) ProtoMessage() {}

func (x *CreateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRequest.ProtoReflect.Descriptor instead.
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *CreateWalletRequest) GetWallet() uint64 {
	if x != nil {
		return x.Wallet
	}
	return 0
}

func (x *CreateWalletRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *CreateWalletRequest) GetCoin() uint64 {
	if x != nil {
		return x.Coin
	}
	return 0
}

func (x *CreateWalletRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type CreateWalletResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wallet        *Wallet                `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWalletResponse) Reset() {
	*x = CreateWalletResponse{}
	mi := &file_wallet_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletResponse) ProtoMessage() {}

func (x *CreateWalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletResponse.ProtoReflect.Descriptor instead.
func (*CreateWalletResponse) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *CreateWalletResponse) GetWallet() *Wallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type UpdateWalletRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wallet        []*Wallet              `protobuf:"bytes,1,rep,name=wallet,proto3" json:"wallet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWalletRequest) Reset() {
	*x = UpdateWalletRequest{}
	mi := &file_wallet_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWalletRequest) ProtoMessage() {}

func (x *UpdateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWalletRequest.ProtoReflect.Descriptor instead.
func (*UpdateWalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateWalletRequest) GetWallet() []*Wallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type UpdateWalletResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wallet        []*Wallet              `protobuf:"bytes,1,rep,name=wallet,proto3" json:"wallet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWalletResponse) Reset() {
	*x = UpdateWalletResponse{}
	mi := &file_wallet_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWalletResponse) ProtoMessage() {}

func (x *UpdateWalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWalletResponse.ProtoReflect.Descriptor instead.
func (*UpdateWalletResponse) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateWalletResponse) GetWallet() []*Wallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

var File_wallet_proto protoreflect.FileDescriptor

const file_wallet_proto_rawDesc = "" +
	"\n" +
	"\fwallet.proto\x12\x02pb\"h\n" +
	"\x06Wallet\x12\x16\n" +
	"\x06wallet\x18\x01 \x01(\x04R\x06wallet\x12\x1a\n" +
	"\bexchange\x18\x02 \x01(\x04R\bexchange\x12\x12\n" +
	"\x04coin\x18\x03 \x01(\x04R\x04coin\x12\x16\n" +
	"\x06amount\x18\x04 \x01(\tR\x06amount\"\xa0\x01\n" +
	"\x0eWalletWithCoin\x12\x16\n" +
	"\x06wallet\x18\x01 \x01(\x04R\x06wallet\x12\x1a\n" +
	"\bexchange\x18\x02 \x01(\x04R\bexchange\x12\x12\n" +
	"\x04coin\x18\x03 \x01(\x04R\x04coin\x12\x16\n" +
	"\x06amount\x18\x04 \x01(\tR\x06amount\x12\x16\n" +
	"\x06symbol\x18\x06 \x01(\tR\x06symbol\x12\x16\n" +
	"\x06active\x18\a \x01(\bR\x06active\"\x83\x01\n" +
	"\x19GetWalletWithCoinResponse\x12\"\n" +
	"\x06wallet\x18\x01 \x01(\v2\n" +
	".pb.WalletR\x06wallet\x12\x12\n" +
	"\x04coin\x18\x02 \x01(\x04R\x04coin\x12\x16\n" +
	"\x06symbol\x18\x03 \x01(\tR\x06symbol\x12\x16\n" +
	"\x06active\x18\x04 \x01(\bR\x06active\"J\n" +
	"\x18GetWalletWithCoinRequest\x12\x1a\n" +
	"\bexchange\x18\x01 \x01(\x04R\bexchange\x12\x12\n" +
	"\x04coin\x18\x02 \x01(\x04R\x04coin\"H\n" +
	"\x1aListWalletWithCoinResponse\x12*\n" +
	"\x06wallet\x18\x01 \x03(\v2\x12.pb.WalletWithCoinR\x06wallet\"7\n" +
	"\x19ListWalletWithCoinRequest\x12\x1a\n" +
	"\bexchange\x18\x01 \x01(\x04R\bexchange\"u\n" +
	"\x13CreateWalletRequest\x12\x16\n" +
	"\x06wallet\x18\x01 \x01(\x04R\x06wallet\x12\x1a\n" +
	"\bexchange\x18\x02 \x01(\x04R\bexchange\x12\x12\n" +
	"\x04coin\x18\x03 \x01(\x04R\x04coin\x12\x16\n" +
	"\x06amount\x18\x04 \x01(\tR\x06amount\":\n" +
	"\x14CreateWalletResponse\x12\"\n" +
	"\x06wallet\x18\x01 \x01(\v2\n" +
	".pb.WalletR\x06wallet\"9\n" +
	"\x13UpdateWalletRequest\x12\"\n" +
	"\x06wallet\x18\x01 \x03(\v2\n" +
	".pb.WalletR\x06wallet\":\n" +
	"\x14UpdateWalletResponse\x12\"\n" +
	"\x06wallet\x18\x01 \x03(\v2\n" +
	".pb.WalletR\x06walletB\x06Z\x04./pbb\x06proto3"

var (
	file_wallet_proto_rawDescOnce sync.Once
	file_wallet_proto_rawDescData []byte
)

func file_wallet_proto_rawDescGZIP() []byte {
	file_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_wallet_proto_rawDesc), len(file_wallet_proto_rawDesc)))
	})
	return file_wallet_proto_rawDescData
}

var file_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_wallet_proto_goTypes = []any{
	(*Wallet)(nil),                     // 0: pb.Wallet
	(*WalletWithCoin)(nil),             // 1: pb.WalletWithCoin
	(*GetWalletWithCoinResponse)(nil),  // 2: pb.GetWalletWithCoinResponse
	(*GetWalletWithCoinRequest)(nil),   // 3: pb.GetWalletWithCoinRequest
	(*ListWalletWithCoinResponse)(nil), // 4: pb.ListWalletWithCoinResponse
	(*ListWalletWithCoinRequest)(nil),  // 5: pb.ListWalletWithCoinRequest
	(*CreateWalletRequest)(nil),        // 6: pb.CreateWalletRequest
	(*CreateWalletResponse)(nil),       // 7: pb.CreateWalletResponse
	(*UpdateWalletRequest)(nil),        // 8: pb.UpdateWalletRequest
	(*UpdateWalletResponse)(nil),       // 9: pb.UpdateWalletResponse
}
var file_wallet_proto_depIdxs = []int32{
	0, // 0: pb.GetWalletWithCoinResponse.wallet:type_name -> pb.Wallet
	1, // 1: pb.ListWalletWithCoinResponse.wallet:type_name -> pb.WalletWithCoin
	0, // 2: pb.CreateWalletResponse.wallet:type_name -> pb.Wallet
	0, // 3: pb.UpdateWalletRequest.wallet:type_name -> pb.Wallet
	0, // 4: pb.UpdateWalletResponse.wallet:type_name -> pb.Wallet
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_wallet_proto_init() }
func file_wallet_proto_init() {
	if File_wallet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_wallet_proto_rawDesc), len(file_wallet_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_proto_msgTypes,
	}.Build()
	File_wallet_proto = out.File
	file_wallet_proto_goTypes = nil
	file_wallet_proto_depIdxs = nil
}
