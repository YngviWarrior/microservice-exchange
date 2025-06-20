// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: candle.proto

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

type Candle struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=Parity,proto3" json:"Parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=Exchange,proto3" json:"Exchange,omitempty"`
	Mts           uint64                 `protobuf:"varint,3,opt,name=Mts,proto3" json:"Mts,omitempty"`
	Open          string                 `protobuf:"bytes,4,opt,name=Open,proto3" json:"Open,omitempty"`
	Close         string                 `protobuf:"bytes,5,opt,name=Close,proto3" json:"Close,omitempty"`
	High          string                 `protobuf:"bytes,6,opt,name=High,proto3" json:"High,omitempty"`
	Low           string                 `protobuf:"bytes,7,opt,name=Low,proto3" json:"Low,omitempty"`
	Volume        string                 `protobuf:"bytes,8,opt,name=Volume,proto3" json:"Volume,omitempty"`
	Roc           string                 `protobuf:"bytes,9,opt,name=Roc,proto3" json:"Roc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Candle) Reset() {
	*x = Candle{}
	mi := &file_candle_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Candle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candle) ProtoMessage() {}

func (x *Candle) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candle.ProtoReflect.Descriptor instead.
func (*Candle) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{0}
}

func (x *Candle) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *Candle) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *Candle) GetMts() uint64 {
	if x != nil {
		return x.Mts
	}
	return 0
}

func (x *Candle) GetOpen() string {
	if x != nil {
		return x.Open
	}
	return ""
}

func (x *Candle) GetClose() string {
	if x != nil {
		return x.Close
	}
	return ""
}

func (x *Candle) GetHigh() string {
	if x != nil {
		return x.High
	}
	return ""
}

func (x *Candle) GetLow() string {
	if x != nil {
		return x.Low
	}
	return ""
}

func (x *Candle) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *Candle) GetRoc() string {
	if x != nil {
		return x.Roc
	}
	return ""
}

type GetCandleFirstMtsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCandleFirstMtsRequest) Reset() {
	*x = GetCandleFirstMtsRequest{}
	mi := &file_candle_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCandleFirstMtsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandleFirstMtsRequest) ProtoMessage() {}

func (x *GetCandleFirstMtsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandleFirstMtsRequest.ProtoReflect.Descriptor instead.
func (*GetCandleFirstMtsRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{1}
}

func (x *GetCandleFirstMtsRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *GetCandleFirstMtsRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

type GetCandleFirstMtsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCandleFirstMtsResponse) Reset() {
	*x = GetCandleFirstMtsResponse{}
	mi := &file_candle_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCandleFirstMtsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandleFirstMtsResponse) ProtoMessage() {}

func (x *GetCandleFirstMtsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandleFirstMtsResponse.ProtoReflect.Descriptor instead.
func (*GetCandleFirstMtsResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{2}
}

func (x *GetCandleFirstMtsResponse) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type GetLastTwoCandlesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLastTwoCandlesRequest) Reset() {
	*x = GetLastTwoCandlesRequest{}
	mi := &file_candle_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLastTwoCandlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastTwoCandlesRequest) ProtoMessage() {}

func (x *GetLastTwoCandlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastTwoCandlesRequest.ProtoReflect.Descriptor instead.
func (*GetLastTwoCandlesRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{3}
}

func (x *GetLastTwoCandlesRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *GetLastTwoCandlesRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

type GetLastTwoCandlesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLastTwoCandlesResponse) Reset() {
	*x = GetLastTwoCandlesResponse{}
	mi := &file_candle_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLastTwoCandlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastTwoCandlesResponse) ProtoMessage() {}

func (x *GetLastTwoCandlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastTwoCandlesResponse.ProtoReflect.Descriptor instead.
func (*GetLastTwoCandlesResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{4}
}

func (x *GetLastTwoCandlesResponse) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type CreateCandlesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCandlesRequest) Reset() {
	*x = CreateCandlesRequest{}
	mi := &file_candle_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCandlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCandlesRequest) ProtoMessage() {}

func (x *CreateCandlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCandlesRequest.ProtoReflect.Descriptor instead.
func (*CreateCandlesRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCandlesRequest) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type CreateCandlesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCandlesResponse) Reset() {
	*x = CreateCandlesResponse{}
	mi := &file_candle_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCandlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCandlesResponse) ProtoMessage() {}

func (x *CreateCandlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCandlesResponse.ProtoReflect.Descriptor instead.
func (*CreateCandlesResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCandlesResponse) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type ListCandleLimitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Limit         uint64                 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCandleLimitRequest) Reset() {
	*x = ListCandleLimitRequest{}
	mi := &file_candle_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCandleLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCandleLimitRequest) ProtoMessage() {}

func (x *ListCandleLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCandleLimitRequest.ProtoReflect.Descriptor instead.
func (*ListCandleLimitRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{7}
}

func (x *ListCandleLimitRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *ListCandleLimitRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *ListCandleLimitRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListCandleLimitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCandleLimitResponse) Reset() {
	*x = ListCandleLimitResponse{}
	mi := &file_candle_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCandleLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCandleLimitResponse) ProtoMessage() {}

func (x *ListCandleLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCandleLimitResponse.ProtoReflect.Descriptor instead.
func (*ListCandleLimitResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{8}
}

func (x *ListCandleLimitResponse) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type ListAvgPricesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	To            uint64                 `protobuf:"varint,1,opt,name=to,proto3" json:"to,omitempty"`
	From          uint64                 `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAvgPricesRequest) Reset() {
	*x = ListAvgPricesRequest{}
	mi := &file_candle_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAvgPricesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAvgPricesRequest) ProtoMessage() {}

func (x *ListAvgPricesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAvgPricesRequest.ProtoReflect.Descriptor instead.
func (*ListAvgPricesRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{9}
}

func (x *ListAvgPricesRequest) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *ListAvgPricesRequest) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

type ListAvgPricesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAvgPricesResponse) Reset() {
	*x = ListAvgPricesResponse{}
	mi := &file_candle_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAvgPricesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAvgPricesResponse) ProtoMessage() {}

func (x *ListAvgPricesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAvgPricesResponse.ProtoReflect.Descriptor instead.
func (*ListAvgPricesResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{10}
}

func (x *ListAvgPricesResponse) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type GetFirstPriceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	From          uint64                 `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFirstPriceRequest) Reset() {
	*x = GetFirstPriceRequest{}
	mi := &file_candle_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFirstPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFirstPriceRequest) ProtoMessage() {}

func (x *GetFirstPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFirstPriceRequest.ProtoReflect.Descriptor instead.
func (*GetFirstPriceRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{11}
}

func (x *GetFirstPriceRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *GetFirstPriceRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *GetFirstPriceRequest) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

type GetFirstPriceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Price         string                 `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFirstPriceResponse) Reset() {
	*x = GetFirstPriceResponse{}
	mi := &file_candle_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFirstPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFirstPriceResponse) ProtoMessage() {}

func (x *GetFirstPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFirstPriceResponse.ProtoReflect.Descriptor instead.
func (*GetFirstPriceResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{12}
}

func (x *GetFirstPriceResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type GetLastPriceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLastPriceRequest) Reset() {
	*x = GetLastPriceRequest{}
	mi := &file_candle_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLastPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastPriceRequest) ProtoMessage() {}

func (x *GetLastPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastPriceRequest.ProtoReflect.Descriptor instead.
func (*GetLastPriceRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{13}
}

func (x *GetLastPriceRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *GetLastPriceRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

type GetLastPriceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Price         string                 `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLastPriceResponse) Reset() {
	*x = GetLastPriceResponse{}
	mi := &file_candle_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLastPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastPriceResponse) ProtoMessage() {}

func (x *GetLastPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastPriceResponse.ProtoReflect.Descriptor instead.
func (*GetLastPriceResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{14}
}

func (x *GetLastPriceResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

var File_candle_proto protoreflect.FileDescriptor

const file_candle_proto_rawDesc = "" +
	"\n" +
	"\fcandle.proto\x12\x02pb\"\xc8\x01\n" +
	"\x06Candle\x12\x16\n" +
	"\x06Parity\x18\x01 \x01(\x04R\x06Parity\x12\x1a\n" +
	"\bExchange\x18\x02 \x01(\x04R\bExchange\x12\x10\n" +
	"\x03Mts\x18\x03 \x01(\x04R\x03Mts\x12\x12\n" +
	"\x04Open\x18\x04 \x01(\tR\x04Open\x12\x14\n" +
	"\x05Close\x18\x05 \x01(\tR\x05Close\x12\x12\n" +
	"\x04High\x18\x06 \x01(\tR\x04High\x12\x10\n" +
	"\x03Low\x18\a \x01(\tR\x03Low\x12\x16\n" +
	"\x06Volume\x18\b \x01(\tR\x06Volume\x12\x10\n" +
	"\x03Roc\x18\t \x01(\tR\x03Roc\"N\n" +
	"\x18GetCandleFirstMtsRequest\x12\x16\n" +
	"\x06parity\x18\x01 \x01(\x04R\x06parity\x12\x1a\n" +
	"\bexchange\x18\x02 \x01(\x04R\bexchange\"A\n" +
	"\x19GetCandleFirstMtsResponse\x12$\n" +
	"\acandles\x18\x01 \x03(\v2\n" +
	".pb.CandleR\acandles\"N\n" +
	"\x18GetLastTwoCandlesRequest\x12\x16\n" +
	"\x06parity\x18\x01 \x01(\x04R\x06parity\x12\x1a\n" +
	"\bexchange\x18\x02 \x01(\x04R\bexchange\"A\n" +
	"\x19GetLastTwoCandlesResponse\x12$\n" +
	"\acandles\x18\x01 \x03(\v2\n" +
	".pb.CandleR\acandles\"<\n" +
	"\x14CreateCandlesRequest\x12$\n" +
	"\acandles\x18\x01 \x03(\v2\n" +
	".pb.CandleR\acandles\"=\n" +
	"\x15CreateCandlesResponse\x12$\n" +
	"\acandles\x18\x01 \x03(\v2\n" +
	".pb.CandleR\acandles\"b\n" +
	"\x16ListCandleLimitRequest\x12\x16\n" +
	"\x06parity\x18\x01 \x01(\x04R\x06parity\x12\x1a\n" +
	"\bexchange\x18\x02 \x01(\x04R\bexchange\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\x04R\x05limit\"?\n" +
	"\x17ListCandleLimitResponse\x12$\n" +
	"\acandles\x18\x01 \x03(\v2\n" +
	".pb.CandleR\acandles\":\n" +
	"\x14ListAvgPricesRequest\x12\x0e\n" +
	"\x02to\x18\x01 \x01(\x04R\x02to\x12\x12\n" +
	"\x04from\x18\x02 \x01(\x04R\x04from\"=\n" +
	"\x15ListAvgPricesResponse\x12$\n" +
	"\acandles\x18\x01 \x03(\v2\n" +
	".pb.CandleR\acandles\"^\n" +
	"\x14GetFirstPriceRequest\x12\x16\n" +
	"\x06parity\x18\x01 \x01(\x04R\x06parity\x12\x1a\n" +
	"\bexchange\x18\x02 \x01(\x04R\bexchange\x12\x12\n" +
	"\x04from\x18\x03 \x01(\x04R\x04from\"-\n" +
	"\x15GetFirstPriceResponse\x12\x14\n" +
	"\x05price\x18\x01 \x01(\tR\x05price\"I\n" +
	"\x13GetLastPriceRequest\x12\x16\n" +
	"\x06parity\x18\x01 \x01(\x04R\x06parity\x12\x1a\n" +
	"\bexchange\x18\x02 \x01(\x04R\bexchange\",\n" +
	"\x14GetLastPriceResponse\x12\x14\n" +
	"\x05price\x18\x01 \x01(\tR\x05priceB\x06Z\x04./pbb\x06proto3"

var (
	file_candle_proto_rawDescOnce sync.Once
	file_candle_proto_rawDescData []byte
)

func file_candle_proto_rawDescGZIP() []byte {
	file_candle_proto_rawDescOnce.Do(func() {
		file_candle_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_candle_proto_rawDesc), len(file_candle_proto_rawDesc)))
	})
	return file_candle_proto_rawDescData
}

var file_candle_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_candle_proto_goTypes = []any{
	(*Candle)(nil),                    // 0: pb.Candle
	(*GetCandleFirstMtsRequest)(nil),  // 1: pb.GetCandleFirstMtsRequest
	(*GetCandleFirstMtsResponse)(nil), // 2: pb.GetCandleFirstMtsResponse
	(*GetLastTwoCandlesRequest)(nil),  // 3: pb.GetLastTwoCandlesRequest
	(*GetLastTwoCandlesResponse)(nil), // 4: pb.GetLastTwoCandlesResponse
	(*CreateCandlesRequest)(nil),      // 5: pb.CreateCandlesRequest
	(*CreateCandlesResponse)(nil),     // 6: pb.CreateCandlesResponse
	(*ListCandleLimitRequest)(nil),    // 7: pb.ListCandleLimitRequest
	(*ListCandleLimitResponse)(nil),   // 8: pb.ListCandleLimitResponse
	(*ListAvgPricesRequest)(nil),      // 9: pb.ListAvgPricesRequest
	(*ListAvgPricesResponse)(nil),     // 10: pb.ListAvgPricesResponse
	(*GetFirstPriceRequest)(nil),      // 11: pb.GetFirstPriceRequest
	(*GetFirstPriceResponse)(nil),     // 12: pb.GetFirstPriceResponse
	(*GetLastPriceRequest)(nil),       // 13: pb.GetLastPriceRequest
	(*GetLastPriceResponse)(nil),      // 14: pb.GetLastPriceResponse
}
var file_candle_proto_depIdxs = []int32{
	0, // 0: pb.GetCandleFirstMtsResponse.candles:type_name -> pb.Candle
	0, // 1: pb.GetLastTwoCandlesResponse.candles:type_name -> pb.Candle
	0, // 2: pb.CreateCandlesRequest.candles:type_name -> pb.Candle
	0, // 3: pb.CreateCandlesResponse.candles:type_name -> pb.Candle
	0, // 4: pb.ListCandleLimitResponse.candles:type_name -> pb.Candle
	0, // 5: pb.ListAvgPricesResponse.candles:type_name -> pb.Candle
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_candle_proto_init() }
func file_candle_proto_init() {
	if File_candle_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_candle_proto_rawDesc), len(file_candle_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_candle_proto_goTypes,
		DependencyIndexes: file_candle_proto_depIdxs,
		MessageInfos:      file_candle_proto_msgTypes,
	}.Build()
	File_candle_proto = out.File
	file_candle_proto_goTypes = nil
	file_candle_proto_depIdxs = nil
}
