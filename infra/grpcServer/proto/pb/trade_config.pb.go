// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: trade_config.proto

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

type TradeConfig struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	TradeConfig             uint64                 `protobuf:"varint,1,opt,name=trade_config,json=tradeConfig,proto3" json:"trade_config,omitempty"`
	User                    uint64                 `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Modality                uint64                 `protobuf:"varint,3,opt,name=modality,proto3" json:"modality,omitempty"`
	Strategy                uint64                 `protobuf:"varint,4,opt,name=strategy,proto3" json:"strategy,omitempty"`
	StrategyEnabled         bool                   `protobuf:"varint,5,opt,name=strategy_enabled,json=strategyEnabled,proto3" json:"strategy_enabled,omitempty"`
	StrategyVariant         uint64                 `protobuf:"varint,6,opt,name=strategy_variant,json=strategyVariant,proto3" json:"strategy_variant,omitempty"`
	Parity                  uint64                 `protobuf:"varint,7,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange                uint64                 `protobuf:"varint,8,opt,name=exchange,proto3" json:"exchange,omitempty"`
	OperationQuantity       uint64                 `protobuf:"varint,9,opt,name=operation_quantity,json=operationQuantity,proto3" json:"operation_quantity,omitempty"`
	OperationAmount         string                 `protobuf:"bytes,10,opt,name=operation_amount,json=operationAmount,proto3" json:"operation_amount,omitempty"`
	Enabled                 bool                   `protobuf:"varint,11,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DefaultProfitPercentage string                 `protobuf:"bytes,12,opt,name=default_profit_percentage,json=defaultProfitPercentage,proto3" json:"default_profit_percentage,omitempty"`
	WalletValueLimit        string                 `protobuf:"bytes,13,opt,name=wallet_value_limit,json=walletValueLimit,proto3" json:"wallet_value_limit,omitempty"`
	UserName                string                 `protobuf:"bytes,14,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	ModalityName            string                 `protobuf:"bytes,15,opt,name=modality_name,json=modalityName,proto3" json:"modality_name,omitempty"`
	StrategyName            string                 `protobuf:"bytes,16,opt,name=strategy_name,json=strategyName,proto3" json:"strategy_name,omitempty"`
	StrategyVariantName     string                 `protobuf:"bytes,17,opt,name=strategy_variant_name,json=strategyVariantName,proto3" json:"strategy_variant_name,omitempty"`
	StrategyVariantEnabled  bool                   `protobuf:"varint,18,opt,name=strategy_variant_enabled,json=strategyVariantEnabled,proto3" json:"strategy_variant_enabled,omitempty"`
	SymbolName              string                 `protobuf:"bytes,19,opt,name=symbol_name,json=symbolName,proto3" json:"symbol_name,omitempty"`
	ExchangeName            string                 `protobuf:"bytes,20,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty"`
	ParitySymbol            string                 `protobuf:"bytes,21,opt,name=parity_symbol,json=paritySymbol,proto3" json:"parity_symbol,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *TradeConfig) Reset() {
	*x = TradeConfig{}
	mi := &file_trade_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TradeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeConfig) ProtoMessage() {}

func (x *TradeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeConfig.ProtoReflect.Descriptor instead.
func (*TradeConfig) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{0}
}

func (x *TradeConfig) GetTradeConfig() uint64 {
	if x != nil {
		return x.TradeConfig
	}
	return 0
}

func (x *TradeConfig) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *TradeConfig) GetModality() uint64 {
	if x != nil {
		return x.Modality
	}
	return 0
}

func (x *TradeConfig) GetStrategy() uint64 {
	if x != nil {
		return x.Strategy
	}
	return 0
}

func (x *TradeConfig) GetStrategyEnabled() bool {
	if x != nil {
		return x.StrategyEnabled
	}
	return false
}

func (x *TradeConfig) GetStrategyVariant() uint64 {
	if x != nil {
		return x.StrategyVariant
	}
	return 0
}

func (x *TradeConfig) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *TradeConfig) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *TradeConfig) GetOperationQuantity() uint64 {
	if x != nil {
		return x.OperationQuantity
	}
	return 0
}

func (x *TradeConfig) GetOperationAmount() string {
	if x != nil {
		return x.OperationAmount
	}
	return ""
}

func (x *TradeConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *TradeConfig) GetDefaultProfitPercentage() string {
	if x != nil {
		return x.DefaultProfitPercentage
	}
	return ""
}

func (x *TradeConfig) GetWalletValueLimit() string {
	if x != nil {
		return x.WalletValueLimit
	}
	return ""
}

func (x *TradeConfig) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *TradeConfig) GetModalityName() string {
	if x != nil {
		return x.ModalityName
	}
	return ""
}

func (x *TradeConfig) GetStrategyName() string {
	if x != nil {
		return x.StrategyName
	}
	return ""
}

func (x *TradeConfig) GetStrategyVariantName() string {
	if x != nil {
		return x.StrategyVariantName
	}
	return ""
}

func (x *TradeConfig) GetStrategyVariantEnabled() bool {
	if x != nil {
		return x.StrategyVariantEnabled
	}
	return false
}

func (x *TradeConfig) GetSymbolName() string {
	if x != nil {
		return x.SymbolName
	}
	return ""
}

func (x *TradeConfig) GetExchangeName() string {
	if x != nil {
		return x.ExchangeName
	}
	return ""
}

func (x *TradeConfig) GetParitySymbol() string {
	if x != nil {
		return x.ParitySymbol
	}
	return ""
}

type TradeConfigResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TradeConfig   []*TradeConfig         `protobuf:"bytes,1,rep,name=trade_config,json=tradeConfig,proto3" json:"trade_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TradeConfigResponse) Reset() {
	*x = TradeConfigResponse{}
	mi := &file_trade_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TradeConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeConfigResponse) ProtoMessage() {}

func (x *TradeConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeConfigResponse.ProtoReflect.Descriptor instead.
func (*TradeConfigResponse) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{1}
}

func (x *TradeConfigResponse) GetTradeConfig() []*TradeConfig {
	if x != nil {
		return x.TradeConfig
	}
	return nil
}

type CreateTradeConfigRequest struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	Modality                uint64                 `protobuf:"varint,1,opt,name=modality,proto3" json:"modality,omitempty"`
	Strategy                uint64                 `protobuf:"varint,2,opt,name=strategy,proto3" json:"strategy,omitempty"`
	StrategyVariant         uint64                 `protobuf:"varint,3,opt,name=strategy_variant,json=strategyVariant,proto3" json:"strategy_variant,omitempty"`
	Parity                  uint64                 `protobuf:"varint,4,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange                uint64                 `protobuf:"varint,5,opt,name=exchange,proto3" json:"exchange,omitempty"`
	OperationQuantity       uint64                 `protobuf:"varint,6,opt,name=operation_quantity,json=operationQuantity,proto3" json:"operation_quantity,omitempty"`
	OperationAmount         string                 `protobuf:"bytes,7,opt,name=operation_amount,json=operationAmount,proto3" json:"operation_amount,omitempty"`
	DefaultProfitPercentage string                 `protobuf:"bytes,8,opt,name=default_profit_percentage,json=defaultProfitPercentage,proto3" json:"default_profit_percentage,omitempty"`
	WalletValueLimit        string                 `protobuf:"bytes,9,opt,name=wallet_value_limit,json=walletValueLimit,proto3" json:"wallet_value_limit,omitempty"`
	Enabled                 bool                   `protobuf:"varint,10,opt,name=enabled,proto3" json:"enabled,omitempty"`
	User                    uint64                 `protobuf:"varint,11,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *CreateTradeConfigRequest) Reset() {
	*x = CreateTradeConfigRequest{}
	mi := &file_trade_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTradeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeConfigRequest) ProtoMessage() {}

func (x *CreateTradeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradeConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateTradeConfigRequest) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTradeConfigRequest) GetModality() uint64 {
	if x != nil {
		return x.Modality
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetStrategy() uint64 {
	if x != nil {
		return x.Strategy
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetStrategyVariant() uint64 {
	if x != nil {
		return x.StrategyVariant
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetOperationQuantity() uint64 {
	if x != nil {
		return x.OperationQuantity
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetOperationAmount() string {
	if x != nil {
		return x.OperationAmount
	}
	return ""
}

func (x *CreateTradeConfigRequest) GetDefaultProfitPercentage() string {
	if x != nil {
		return x.DefaultProfitPercentage
	}
	return ""
}

func (x *CreateTradeConfigRequest) GetWalletValueLimit() string {
	if x != nil {
		return x.WalletValueLimit
	}
	return ""
}

func (x *CreateTradeConfigRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CreateTradeConfigRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

type ListTradeConfigRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTradeConfigRequest) Reset() {
	*x = ListTradeConfigRequest{}
	mi := &file_trade_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTradeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTradeConfigRequest) ProtoMessage() {}

func (x *ListTradeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTradeConfigRequest.ProtoReflect.Descriptor instead.
func (*ListTradeConfigRequest) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{3}
}

type UpdateTradeConfigRequest struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	TradeConfig             uint64                 `protobuf:"varint,1,opt,name=tradeConfig,proto3" json:"tradeConfig,omitempty"`
	User                    uint64                 `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Modality                uint64                 `protobuf:"varint,3,opt,name=modality,proto3" json:"modality,omitempty"`
	Strategy                uint64                 `protobuf:"varint,4,opt,name=strategy,proto3" json:"strategy,omitempty"`
	StrategyVariant         uint64                 `protobuf:"varint,5,opt,name=strategy_variant,json=strategyVariant,proto3" json:"strategy_variant,omitempty"`
	Parity                  uint64                 `protobuf:"varint,6,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange                uint64                 `protobuf:"varint,7,opt,name=exchange,proto3" json:"exchange,omitempty"`
	OperationQuantity       uint64                 `protobuf:"varint,8,opt,name=operation_quantity,json=operationQuantity,proto3" json:"operation_quantity,omitempty"`
	OperationAmount         string                 `protobuf:"bytes,9,opt,name=operation_amount,json=operationAmount,proto3" json:"operation_amount,omitempty"`
	DefaultProfitPercentage string                 `protobuf:"bytes,10,opt,name=default_profit_percentage,json=defaultProfitPercentage,proto3" json:"default_profit_percentage,omitempty"`
	WalletValueLimit        string                 `protobuf:"bytes,11,opt,name=wallet_value_limit,json=walletValueLimit,proto3" json:"wallet_value_limit,omitempty"`
	Enabled                 bool                   `protobuf:"varint,13,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *UpdateTradeConfigRequest) Reset() {
	*x = UpdateTradeConfigRequest{}
	mi := &file_trade_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTradeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTradeConfigRequest) ProtoMessage() {}

func (x *UpdateTradeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTradeConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateTradeConfigRequest) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTradeConfigRequest) GetTradeConfig() uint64 {
	if x != nil {
		return x.TradeConfig
	}
	return 0
}

func (x *UpdateTradeConfigRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *UpdateTradeConfigRequest) GetModality() uint64 {
	if x != nil {
		return x.Modality
	}
	return 0
}

func (x *UpdateTradeConfigRequest) GetStrategy() uint64 {
	if x != nil {
		return x.Strategy
	}
	return 0
}

func (x *UpdateTradeConfigRequest) GetStrategyVariant() uint64 {
	if x != nil {
		return x.StrategyVariant
	}
	return 0
}

func (x *UpdateTradeConfigRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *UpdateTradeConfigRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *UpdateTradeConfigRequest) GetOperationQuantity() uint64 {
	if x != nil {
		return x.OperationQuantity
	}
	return 0
}

func (x *UpdateTradeConfigRequest) GetOperationAmount() string {
	if x != nil {
		return x.OperationAmount
	}
	return ""
}

func (x *UpdateTradeConfigRequest) GetDefaultProfitPercentage() string {
	if x != nil {
		return x.DefaultProfitPercentage
	}
	return ""
}

func (x *UpdateTradeConfigRequest) GetWalletValueLimit() string {
	if x != nil {
		return x.WalletValueLimit
	}
	return ""
}

func (x *UpdateTradeConfigRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type UpdateTradeConfigResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tradeconfig   *TradeConfig           `protobuf:"bytes,1,opt,name=tradeconfig,proto3" json:"tradeconfig,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTradeConfigResponse) Reset() {
	*x = UpdateTradeConfigResponse{}
	mi := &file_trade_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTradeConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTradeConfigResponse) ProtoMessage() {}

func (x *UpdateTradeConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTradeConfigResponse.ProtoReflect.Descriptor instead.
func (*UpdateTradeConfigResponse) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTradeConfigResponse) GetTradeconfig() *TradeConfig {
	if x != nil {
		return x.Tradeconfig
	}
	return nil
}

type GetTradeConfigRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	User            uint64                 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Strategy        uint64                 `protobuf:"varint,2,opt,name=strategy,proto3" json:"strategy,omitempty"`
	StrategyVariant uint64                 `protobuf:"varint,3,opt,name=strategy_variant,json=strategyVariant,proto3" json:"strategy_variant,omitempty"`
	Parity          uint64                 `protobuf:"varint,4,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange        uint64                 `protobuf:"varint,5,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Modality        uint64                 `protobuf:"varint,6,opt,name=modality,proto3" json:"modality,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetTradeConfigRequest) Reset() {
	*x = GetTradeConfigRequest{}
	mi := &file_trade_config_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTradeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTradeConfigRequest) ProtoMessage() {}

func (x *GetTradeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTradeConfigRequest.ProtoReflect.Descriptor instead.
func (*GetTradeConfigRequest) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{6}
}

func (x *GetTradeConfigRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *GetTradeConfigRequest) GetStrategy() uint64 {
	if x != nil {
		return x.Strategy
	}
	return 0
}

func (x *GetTradeConfigRequest) GetStrategyVariant() uint64 {
	if x != nil {
		return x.StrategyVariant
	}
	return 0
}

func (x *GetTradeConfigRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *GetTradeConfigRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *GetTradeConfigRequest) GetModality() uint64 {
	if x != nil {
		return x.Modality
	}
	return 0
}

type GetTradeConfigResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TradeConfig   *TradeConfig           `protobuf:"bytes,1,opt,name=trade_config,json=tradeConfig,proto3" json:"trade_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTradeConfigResponse) Reset() {
	*x = GetTradeConfigResponse{}
	mi := &file_trade_config_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTradeConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTradeConfigResponse) ProtoMessage() {}

func (x *GetTradeConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTradeConfigResponse.ProtoReflect.Descriptor instead.
func (*GetTradeConfigResponse) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{7}
}

func (x *GetTradeConfigResponse) GetTradeConfig() *TradeConfig {
	if x != nil {
		return x.TradeConfig
	}
	return nil
}

var File_trade_config_proto protoreflect.FileDescriptor

const file_trade_config_proto_rawDesc = "" +
	"\n" +
	"\x12trade_config.proto\x12\x02pb\"\xa4\x06\n" +
	"\vTradeConfig\x12!\n" +
	"\ftrade_config\x18\x01 \x01(\x04R\vtradeConfig\x12\x12\n" +
	"\x04user\x18\x02 \x01(\x04R\x04user\x12\x1a\n" +
	"\bmodality\x18\x03 \x01(\x04R\bmodality\x12\x1a\n" +
	"\bstrategy\x18\x04 \x01(\x04R\bstrategy\x12)\n" +
	"\x10strategy_enabled\x18\x05 \x01(\bR\x0fstrategyEnabled\x12)\n" +
	"\x10strategy_variant\x18\x06 \x01(\x04R\x0fstrategyVariant\x12\x16\n" +
	"\x06parity\x18\a \x01(\x04R\x06parity\x12\x1a\n" +
	"\bexchange\x18\b \x01(\x04R\bexchange\x12-\n" +
	"\x12operation_quantity\x18\t \x01(\x04R\x11operationQuantity\x12)\n" +
	"\x10operation_amount\x18\n" +
	" \x01(\tR\x0foperationAmount\x12\x18\n" +
	"\aenabled\x18\v \x01(\bR\aenabled\x12:\n" +
	"\x19default_profit_percentage\x18\f \x01(\tR\x17defaultProfitPercentage\x12,\n" +
	"\x12wallet_value_limit\x18\r \x01(\tR\x10walletValueLimit\x12\x1b\n" +
	"\tuser_name\x18\x0e \x01(\tR\buserName\x12#\n" +
	"\rmodality_name\x18\x0f \x01(\tR\fmodalityName\x12#\n" +
	"\rstrategy_name\x18\x10 \x01(\tR\fstrategyName\x122\n" +
	"\x15strategy_variant_name\x18\x11 \x01(\tR\x13strategyVariantName\x128\n" +
	"\x18strategy_variant_enabled\x18\x12 \x01(\bR\x16strategyVariantEnabled\x12\x1f\n" +
	"\vsymbol_name\x18\x13 \x01(\tR\n" +
	"symbolName\x12#\n" +
	"\rexchange_name\x18\x14 \x01(\tR\fexchangeName\x12#\n" +
	"\rparity_symbol\x18\x15 \x01(\tR\fparitySymbol\"I\n" +
	"\x13TradeConfigResponse\x122\n" +
	"\ftrade_config\x18\x01 \x03(\v2\x0f.pb.TradeConfigR\vtradeConfig\"\xa3\x03\n" +
	"\x18CreateTradeConfigRequest\x12\x1a\n" +
	"\bmodality\x18\x01 \x01(\x04R\bmodality\x12\x1a\n" +
	"\bstrategy\x18\x02 \x01(\x04R\bstrategy\x12)\n" +
	"\x10strategy_variant\x18\x03 \x01(\x04R\x0fstrategyVariant\x12\x16\n" +
	"\x06parity\x18\x04 \x01(\x04R\x06parity\x12\x1a\n" +
	"\bexchange\x18\x05 \x01(\x04R\bexchange\x12-\n" +
	"\x12operation_quantity\x18\x06 \x01(\x04R\x11operationQuantity\x12)\n" +
	"\x10operation_amount\x18\a \x01(\tR\x0foperationAmount\x12:\n" +
	"\x19default_profit_percentage\x18\b \x01(\tR\x17defaultProfitPercentage\x12,\n" +
	"\x12wallet_value_limit\x18\t \x01(\tR\x10walletValueLimit\x12\x18\n" +
	"\aenabled\x18\n" +
	" \x01(\bR\aenabled\x12\x12\n" +
	"\x04user\x18\v \x01(\x04R\x04user\"\x18\n" +
	"\x16ListTradeConfigRequest\"\xc5\x03\n" +
	"\x18UpdateTradeConfigRequest\x12 \n" +
	"\vtradeConfig\x18\x01 \x01(\x04R\vtradeConfig\x12\x12\n" +
	"\x04user\x18\x02 \x01(\x04R\x04user\x12\x1a\n" +
	"\bmodality\x18\x03 \x01(\x04R\bmodality\x12\x1a\n" +
	"\bstrategy\x18\x04 \x01(\x04R\bstrategy\x12)\n" +
	"\x10strategy_variant\x18\x05 \x01(\x04R\x0fstrategyVariant\x12\x16\n" +
	"\x06parity\x18\x06 \x01(\x04R\x06parity\x12\x1a\n" +
	"\bexchange\x18\a \x01(\x04R\bexchange\x12-\n" +
	"\x12operation_quantity\x18\b \x01(\x04R\x11operationQuantity\x12)\n" +
	"\x10operation_amount\x18\t \x01(\tR\x0foperationAmount\x12:\n" +
	"\x19default_profit_percentage\x18\n" +
	" \x01(\tR\x17defaultProfitPercentage\x12,\n" +
	"\x12wallet_value_limit\x18\v \x01(\tR\x10walletValueLimit\x12\x18\n" +
	"\aenabled\x18\r \x01(\bR\aenabled\"N\n" +
	"\x19UpdateTradeConfigResponse\x121\n" +
	"\vtradeconfig\x18\x01 \x01(\v2\x0f.pb.TradeConfigR\vtradeconfig\"\xc2\x01\n" +
	"\x15GetTradeConfigRequest\x12\x12\n" +
	"\x04user\x18\x01 \x01(\x04R\x04user\x12\x1a\n" +
	"\bstrategy\x18\x02 \x01(\x04R\bstrategy\x12)\n" +
	"\x10strategy_variant\x18\x03 \x01(\x04R\x0fstrategyVariant\x12\x16\n" +
	"\x06parity\x18\x04 \x01(\x04R\x06parity\x12\x1a\n" +
	"\bexchange\x18\x05 \x01(\x04R\bexchange\x12\x1a\n" +
	"\bmodality\x18\x06 \x01(\x04R\bmodality\"L\n" +
	"\x16GetTradeConfigResponse\x122\n" +
	"\ftrade_config\x18\x01 \x01(\v2\x0f.pb.TradeConfigR\vtradeConfigB\x06Z\x04./pbb\x06proto3"

var (
	file_trade_config_proto_rawDescOnce sync.Once
	file_trade_config_proto_rawDescData []byte
)

func file_trade_config_proto_rawDescGZIP() []byte {
	file_trade_config_proto_rawDescOnce.Do(func() {
		file_trade_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_trade_config_proto_rawDesc), len(file_trade_config_proto_rawDesc)))
	})
	return file_trade_config_proto_rawDescData
}

var file_trade_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_trade_config_proto_goTypes = []any{
	(*TradeConfig)(nil),               // 0: pb.TradeConfig
	(*TradeConfigResponse)(nil),       // 1: pb.TradeConfigResponse
	(*CreateTradeConfigRequest)(nil),  // 2: pb.CreateTradeConfigRequest
	(*ListTradeConfigRequest)(nil),    // 3: pb.ListTradeConfigRequest
	(*UpdateTradeConfigRequest)(nil),  // 4: pb.UpdateTradeConfigRequest
	(*UpdateTradeConfigResponse)(nil), // 5: pb.UpdateTradeConfigResponse
	(*GetTradeConfigRequest)(nil),     // 6: pb.GetTradeConfigRequest
	(*GetTradeConfigResponse)(nil),    // 7: pb.GetTradeConfigResponse
}
var file_trade_config_proto_depIdxs = []int32{
	0, // 0: pb.TradeConfigResponse.trade_config:type_name -> pb.TradeConfig
	0, // 1: pb.UpdateTradeConfigResponse.tradeconfig:type_name -> pb.TradeConfig
	0, // 2: pb.GetTradeConfigResponse.trade_config:type_name -> pb.TradeConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_trade_config_proto_init() }
func file_trade_config_proto_init() {
	if File_trade_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_trade_config_proto_rawDesc), len(file_trade_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trade_config_proto_goTypes,
		DependencyIndexes: file_trade_config_proto_depIdxs,
		MessageInfos:      file_trade_config_proto_msgTypes,
	}.Build()
	File_trade_config_proto = out.File
	file_trade_config_proto_goTypes = nil
	file_trade_config_proto_depIdxs = nil
}
