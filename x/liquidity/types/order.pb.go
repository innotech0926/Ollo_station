// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/liquidity/v1/order.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// OrderType enumerates order types.
type OrderType int32

const (
	// ORDER_TYPE_UNSPECIFIED specifies unknown order type.
	OrderTypeUnspecified OrderType = 0
	// ORDER_TYPE_LIMIT specifies limit order type.
	OrderTypeLimit OrderType = 1
	// ORDER_TYPE_MARKET specifies market order type.
	OrderTypeMarket OrderType = 2
	// ORDER_TYPE_MM specifies MM(market making) order type.
	OrderTypeMM OrderType = 3
)

var OrderType_name = map[int32]string{
	0: "ORDER_TYPE_UNSPECIFIED",
	1: "ORDER_TYPE_LIMIT",
	2: "ORDER_TYPE_MARKET",
	3: "ORDER_TYPE_MM",
}

var OrderType_value = map[string]int32{
	"ORDER_TYPE_UNSPECIFIED": 0,
	"ORDER_TYPE_LIMIT":       1,
	"ORDER_TYPE_MARKET":      2,
	"ORDER_TYPE_MM":          3,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}

func (OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{0}
}

// OrderDirection enumerates order directions.
type OrderDirection int32

const (
	// ORDER_DIRECTION_UNSPECIFIED specifies unknown order direction
	OrderDirectionUnspecified OrderDirection = 0
	// ORDER_DIRECTION_BUY specifies buy(swap quote coin to base coin) order direction
	OrderDirectionBuy OrderDirection = 1
	// ORDER_DIRECTION_SELL specifies sell(swap base coin to quote coin) order direction
	OrderDirectionSell OrderDirection = 2
)

var OrderDirection_name = map[int32]string{
	0: "ORDER_DIRECTION_UNSPECIFIED",
	1: "ORDER_DIRECTION_BUY",
	2: "ORDER_DIRECTION_SELL",
}

var OrderDirection_value = map[string]int32{
	"ORDER_DIRECTION_UNSPECIFIED": 0,
	"ORDER_DIRECTION_BUY":         1,
	"ORDER_DIRECTION_SELL":        2,
}

func (x OrderDirection) String() string {
	return proto.EnumName(OrderDirection_name, int32(x))
}

func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{1}
}

// RequestStatus enumerates request statuses.
type RequestStatus int32

const (
	// REQUEST_STATUS_UNSPECIFIED specifies unknown request status
	RequestStatusUnspecified RequestStatus = 0
	// REQUEST_STATUS_NOT_EXECUTED indicates the request is not executed yet
	RequestStatusNotExecuted RequestStatus = 1
	// REQUEST_STATUS_SUCCEEDED indicates the request has been succeeded
	RequestStatusSucceeded RequestStatus = 2
	// REQUEST_STATUS_FAILED indicates the request is failed
	RequestStatusFailed RequestStatus = 3
)

var RequestStatus_name = map[int32]string{
	0: "REQUEST_STATUS_UNSPECIFIED",
	1: "REQUEST_STATUS_NOT_EXECUTED",
	2: "REQUEST_STATUS_SUCCEEDED",
	3: "REQUEST_STATUS_FAILED",
}

var RequestStatus_value = map[string]int32{
	"REQUEST_STATUS_UNSPECIFIED":  0,
	"REQUEST_STATUS_NOT_EXECUTED": 1,
	"REQUEST_STATUS_SUCCEEDED":    2,
	"REQUEST_STATUS_FAILED":       3,
}

func (x RequestStatus) String() string {
	return proto.EnumName(RequestStatus_name, int32(x))
}

func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{2}
}

// OrderStatus enumerates order statuses.
type OrderStatus int32

const (
	// ORDER_STATUS_UNSPECIFIED specifies unknown order status
	OrderStatusUnspecified OrderStatus = 0
	// ORDER_STATUS_NOT_EXECUTED indicates the order has not been executed yet
	OrderStatusNotExecuted OrderStatus = 1
	// ORDER_STATUS_NOT_MATCHED indicates the order has been executed but has no match
	OrderStatusNotMatched OrderStatus = 2
	// ORDER_STATUS_PARTIALLY_MATCHED indicates the order has been partially matched
	OrderStatusPartiallyMatched OrderStatus = 3
	// ORDER_STATUS_COMPLETED indicates the order has been fully matched and completed
	OrderStatusCompleted OrderStatus = 4
	// ORDER_STATUS_CANCELED indicates the order has been canceled
	OrderStatusCanceled OrderStatus = 5
	// ORDER_STATUS_EXPIRED indicates the order has been expired
	OrderStatusExpired OrderStatus = 6
)

var OrderStatus_name = map[int32]string{
	0: "ORDER_STATUS_UNSPECIFIED",
	1: "ORDER_STATUS_NOT_EXECUTED",
	2: "ORDER_STATUS_NOT_MATCHED",
	3: "ORDER_STATUS_PARTIALLY_MATCHED",
	4: "ORDER_STATUS_COMPLETED",
	5: "ORDER_STATUS_CANCELED",
	6: "ORDER_STATUS_EXPIRED",
}

var OrderStatus_value = map[string]int32{
	"ORDER_STATUS_UNSPECIFIED":       0,
	"ORDER_STATUS_NOT_EXECUTED":      1,
	"ORDER_STATUS_NOT_MATCHED":       2,
	"ORDER_STATUS_PARTIALLY_MATCHED": 3,
	"ORDER_STATUS_COMPLETED":         4,
	"ORDER_STATUS_CANCELED":          5,
	"ORDER_STATUS_EXPIRED":           6,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{3}
}

// Order defines an order.
type Order struct {
	// type specifies the typo of the order
	Type OrderType `protobuf:"varint,1,opt,name=type,proto3,enum=ollo.liquidity.v1.OrderType" json:"type,omitempty"`
	// id specifies the id of the order
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// pair_id specifies the pair id
	PairId uint64 `protobuf:"varint,3,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty"`
	// msg_height specifies the block height when the order is stored for the batch execution
	MsgHeight int64 `protobuf:"varint,4,opt,name=msg_height,json=msgHeight,proto3" json:"msg_height,omitempty"`
	// orderer specifies the bech32-encoded address that makes an order
	Orderer string `protobuf:"bytes,5,opt,name=orderer,proto3" json:"orderer,omitempty" yaml:"orderer"`
	// direction specifies the order direction; either buy or sell
	Direction OrderDirection `protobuf:"varint,6,opt,name=direction,proto3,enum=ollo.liquidity.v1.OrderDirection" json:"direction,omitempty"`
	OfferCoin types.Coin     `protobuf:"bytes,7,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin"`
	// remaining_offer_coin specifies the remaining offer coin
	RemainingOfferCoin types.Coin `protobuf:"bytes,8,opt,name=remaining_offer_coin,json=remainingOfferCoin,proto3" json:"remaining_offer_coin"`
	// received_coin specifies the received coin after the swap
	ReceivedCoin types.Coin `protobuf:"bytes,9,opt,name=received_coin,json=receivedCoin,proto3" json:"received_coin"`
	// price specifies the price that an orderer is willing to swap
	Price      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	Amount     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	OpenAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,12,opt,name=open_amount,json=openAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"open_amount"`
	// batch_id specifies the pair's batch id when the request is stored
	BatchId  uint64      `protobuf:"varint,13,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	ExpireAt time.Time   `protobuf:"bytes,14,opt,name=expire_at,json=expireAt,proto3,stdtime" json:"expire_at"`
	Status   OrderStatus `protobuf:"varint,15,opt,name=status,proto3,enum=ollo.liquidity.v1.OrderStatus" json:"status,omitempty"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ollo.liquidity.v1.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("ollo.liquidity.v1.OrderDirection", OrderDirection_name, OrderDirection_value)
	proto.RegisterEnum("ollo.liquidity.v1.RequestStatus", RequestStatus_name, RequestStatus_value)
	proto.RegisterEnum("ollo.liquidity.v1.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterType((*Order)(nil), "ollo.liquidity.v1.Order")
}

func init() { proto.RegisterFile("ollo/liquidity/v1/order.proto", fileDescriptor_71b117091565b732) }

var fileDescriptor_71b117091565b732 = []byte{
	// 1084 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0x4d, 0x6f, 0xe2, 0x56,
	0x17, 0xc7, 0x31, 0x10, 0x12, 0x6e, 0x26, 0x0c, 0xb9, 0x79, 0x19, 0xe3, 0x4c, 0xc0, 0x4f, 0xf4,
	0xa8, 0x42, 0x91, 0x62, 0x26, 0x69, 0xd5, 0x37, 0xb5, 0xa9, 0x78, 0xb9, 0xd1, 0x58, 0x85, 0xc0,
	0x18, 0x23, 0x4d, 0xba, 0x41, 0xc6, 0xbe, 0x21, 0x57, 0x63, 0x30, 0x63, 0x5f, 0xa2, 0xb0, 0xeb,
	0xb2, 0xf2, 0x6a, 0xbe, 0x80, 0x37, 0xed, 0x57, 0xe8, 0xa2, 0xdb, 0xee, 0xb2, 0x8c, 0xba, 0xaa,
	0xba, 0x48, 0xdb, 0xe4, 0x1b, 0xf4, 0x13, 0x54, 0x7e, 0x81, 0xd8, 0xee, 0x8c, 0x34, 0xed, 0x2a,
	0xb9, 0x3e, 0xff, 0xdf, 0xdf, 0xe7, 0x1c, 0x9f, 0x73, 0x05, 0xd8, 0x35, 0x74, 0xdd, 0xa8, 0xe8,
	0xe4, 0xf5, 0x94, 0x68, 0x84, 0xce, 0x2a, 0x97, 0x87, 0x15, 0xc3, 0xd4, 0xb0, 0x29, 0x4c, 0x4c,
	0x83, 0x1a, 0x70, 0xdd, 0x0d, 0x0b, 0x8b, 0xb0, 0x70, 0x79, 0xc8, 0x6d, 0x0e, 0x8d, 0xa1, 0xe1,
	0x45, 0x2b, 0xee, 0x7f, 0xbe, 0x90, 0x2b, 0xa8, 0x86, 0x35, 0x32, 0xac, 0xbe, 0x1f, 0xf0, 0x0f,
	0x41, 0xa8, 0xe8, 0x9f, 0x2a, 0x03, 0xc5, 0xc2, 0x95, 0xcb, 0xc3, 0x01, 0xa6, 0xca, 0x61, 0x45,
	0x35, 0xc8, 0x38, 0x88, 0x97, 0x86, 0x86, 0x31, 0xd4, 0x71, 0xc5, 0x3b, 0x0d, 0xa6, 0xe7, 0x15,
	0x4a, 0x46, 0xd8, 0xa2, 0xca, 0x68, 0x32, 0x37, 0x88, 0x0b, 0xb4, 0xa9, 0xa9, 0x50, 0x62, 0x04,
	0x06, 0x7b, 0x37, 0x19, 0xb0, 0xd4, 0x76, 0x93, 0x86, 0xcf, 0x40, 0x9a, 0xce, 0x26, 0x98, 0x65,
	0x78, 0xa6, 0x9c, 0x3b, 0x7a, 0x2a, 0xfc, 0x23, 0x7b, 0xc1, 0xd3, 0xc9, 0xb3, 0x09, 0x96, 0x3c,
	0x25, 0xcc, 0x81, 0x24, 0xd1, 0xd8, 0x24, 0xcf, 0x94, 0xd3, 0x52, 0x92, 0x68, 0xf0, 0x09, 0x58,
	0x9e, 0x28, 0xc4, 0xec, 0x13, 0x8d, 0x4d, 0x79, 0x0f, 0x33, 0xee, 0x51, 0xd4, 0xe0, 0x2e, 0x00,
	0x23, 0x6b, 0xd8, 0xbf, 0xc0, 0x64, 0x78, 0x41, 0xd9, 0x34, 0xcf, 0x94, 0x53, 0x52, 0x76, 0x64,
	0x0d, 0x9f, 0x7b, 0x0f, 0xe0, 0x31, 0x58, 0xf6, 0xfa, 0x86, 0x4d, 0x76, 0x89, 0x67, 0xca, 0xd9,
	0xda, 0xff, 0xff, 0xba, 0x2d, 0xe5, 0x66, 0xca, 0x48, 0xff, 0x7c, 0x2f, 0x08, 0xec, 0xfd, 0xf2,
	0xe3, 0x41, 0x2e, 0xe8, 0x4c, 0x55, 0xd3, 0x4c, 0x6c, 0x59, 0xd2, 0x1c, 0x82, 0x5f, 0x81, 0xac,
	0x46, 0x4c, 0xac, 0xba, 0x65, 0xb1, 0x19, 0x2f, 0xfd, 0xff, 0xbd, 0x2b, 0xfd, 0xc6, 0x5c, 0x28,
	0x3d, 0x30, 0xf0, 0x18, 0x00, 0xe3, 0xfc, 0x1c, 0x9b, 0x7d, 0xb7, 0xb3, 0xec, 0x32, 0xcf, 0x94,
	0x57, 0x8f, 0x0a, 0x42, 0xf0, 0x3a, 0xb7, 0xf5, 0x42, 0xd0, 0x7a, 0xa1, 0x6e, 0x90, 0x71, 0x2d,
	0x7d, 0x7d, 0x5b, 0x4a, 0x48, 0x59, 0x0f, 0x71, 0x1f, 0xc0, 0x17, 0x60, 0xd3, 0xc4, 0x23, 0x85,
	0x8c, 0xc9, 0x78, 0xd8, 0x0f, 0x39, 0xad, 0xbc, 0x9f, 0x13, 0x5c, 0xc0, 0xed, 0x85, 0x65, 0x03,
	0xac, 0x99, 0x58, 0xc5, 0xe4, 0x12, 0x6b, 0xbe, 0x57, 0xf6, 0xfd, 0xbc, 0x1e, 0xcd, 0xa9, 0xc0,
	0x65, 0x69, 0x62, 0x12, 0x15, 0xb3, 0xc0, 0xeb, 0xab, 0xe0, 0x4a, 0x7e, 0xbb, 0x2d, 0x7d, 0x30,
	0x24, 0xf4, 0x62, 0x3a, 0x10, 0x54, 0x63, 0x14, 0x8c, 0x5b, 0xf0, 0xe7, 0xc0, 0xd2, 0x5e, 0x55,
	0xdc, 0x6f, 0x6b, 0x09, 0x0d, 0xac, 0x4a, 0x3e, 0x0c, 0x4f, 0x40, 0x46, 0x19, 0x19, 0xd3, 0x31,
	0x65, 0x57, 0xff, 0xb5, 0x8d, 0x38, 0xa6, 0x52, 0x40, 0xc3, 0x36, 0x58, 0x35, 0x26, 0x78, 0xdc,
	0x0f, 0xcc, 0x1e, 0xfd, 0x27, 0x33, 0xe0, 0x5a, 0x54, 0x7d, 0xc3, 0x02, 0x58, 0x19, 0x28, 0x54,
	0xbd, 0x70, 0x27, 0x6e, 0xcd, 0x9b, 0xb8, 0x65, 0xef, 0x2c, 0x6a, 0xb0, 0x0a, 0xb2, 0xf8, 0x6a,
	0x42, 0x4c, 0xdc, 0x57, 0x28, 0x9b, 0xf3, 0x7a, 0xc7, 0x09, 0xfe, 0x2e, 0x08, 0xf3, 0x5d, 0x10,
	0xe4, 0xf9, 0xb2, 0xd4, 0x56, 0xdc, 0x2c, 0xde, 0xfc, 0x5e, 0x62, 0xa4, 0x15, 0x1f, 0xab, 0x52,
	0xf8, 0x31, 0xc8, 0x58, 0x54, 0xa1, 0x53, 0x8b, 0x7d, 0xec, 0xcd, 0x54, 0xf1, 0x5d, 0x33, 0xd5,
	0xf5, 0x54, 0x52, 0xa0, 0xde, 0xff, 0x99, 0x01, 0xd9, 0xc5, 0xaa, 0xc0, 0x8f, 0xc0, 0x76, 0x5b,
	0x6a, 0x20, 0xa9, 0x2f, 0x9f, 0x75, 0x50, 0xbf, 0x77, 0xda, 0xed, 0xa0, 0xba, 0x78, 0x22, 0xa2,
	0x46, 0x3e, 0xc1, 0xb1, 0xb6, 0xc3, 0x6f, 0x2e, 0xa4, 0xbd, 0xb1, 0x35, 0xc1, 0x2a, 0x39, 0x27,
	0x58, 0x83, 0x65, 0x90, 0x0f, 0x51, 0x4d, 0xb1, 0x25, 0xca, 0x79, 0x86, 0x83, 0xb6, 0xc3, 0xe7,
	0x16, 0xfa, 0x26, 0x19, 0x11, 0x0a, 0xf7, 0xc1, 0x7a, 0x48, 0xd9, 0xaa, 0x4a, 0x5f, 0x23, 0x39,
	0x9f, 0xe4, 0x36, 0x6c, 0x87, 0x7f, 0xbc, 0x90, 0xb6, 0x14, 0xf3, 0x15, 0xa6, 0x70, 0x0f, 0xac,
	0x85, 0xb5, 0xad, 0x7c, 0x8a, 0x7b, 0x6c, 0x3b, 0xfc, 0xea, 0x83, 0xae, 0xc5, 0xa5, 0xbf, 0xfb,
	0xa1, 0x98, 0xd8, 0xff, 0x89, 0x01, 0xb9, 0xe8, 0xbe, 0xc0, 0x63, 0xb0, 0xe3, 0xc3, 0x0d, 0x51,
	0x42, 0x75, 0x59, 0x6c, 0x9f, 0xc6, 0xaa, 0xd9, 0xb5, 0x1d, 0xbe, 0x10, 0x85, 0xc2, 0x25, 0x09,
	0x60, 0x23, 0xce, 0xd7, 0x7a, 0x67, 0x79, 0x86, 0xdb, 0xb2, 0x1d, 0x7e, 0x3d, 0xca, 0xd5, 0xa6,
	0x33, 0xf8, 0x0c, 0x6c, 0xc6, 0xf5, 0x5d, 0xd4, 0x6c, 0xe6, 0x93, 0xdc, 0xb6, 0xed, 0xf0, 0x30,
	0x0a, 0x74, 0xb1, 0xae, 0x07, 0xa9, 0x7f, 0x9b, 0x04, 0x6b, 0x12, 0x7e, 0x3d, 0xc5, 0x16, 0xf5,
	0x3f, 0x0c, 0xfc, 0x02, 0x70, 0x12, 0x7a, 0xd1, 0x43, 0x5d, 0xb9, 0xdf, 0x95, 0xab, 0x72, 0xaf,
	0x1b, 0x4b, 0xfc, 0xa9, 0xed, 0xf0, 0x6c, 0x04, 0x09, 0xe7, 0xfd, 0x25, 0xd8, 0x89, 0xd1, 0xa7,
	0x6d, 0xb9, 0x8f, 0x5e, 0xa2, 0x7a, 0x4f, 0x46, 0x8d, 0x3c, 0xf3, 0x16, 0xfc, 0xd4, 0xa0, 0xe8,
	0x0a, 0xab, 0x53, 0x8a, 0x35, 0xf8, 0x29, 0x60, 0x63, 0x78, 0xb7, 0x57, 0xaf, 0x23, 0xd4, 0x40,
	0x8d, 0x7c, 0x92, 0xe3, 0x6c, 0x87, 0xdf, 0x8e, 0xb0, 0xdd, 0xa9, 0xaa, 0x62, 0xac, 0x61, 0x0d,
	0x1e, 0x81, 0xad, 0x18, 0x79, 0x52, 0x15, 0x9b, 0xa8, 0x91, 0x4f, 0x71, 0x4f, 0x6c, 0x87, 0xdf,
	0x88, 0x60, 0x27, 0x0a, 0xd1, 0xb1, 0x16, 0xb4, 0xe0, 0xfb, 0x14, 0x58, 0x0d, 0x4d, 0xa6, 0x9b,
	0x83, 0xdf, 0xca, 0xb7, 0x96, 0xef, 0xe5, 0x10, 0x92, 0x87, 0x8b, 0xff, 0x0c, 0x14, 0x22, 0x64,
	0xac, 0xf4, 0x38, 0x1a, 0x2e, 0xfc, 0x93, 0xd8, 0x4b, 0x5d, 0xb4, 0x55, 0x95, 0xeb, 0xcf, 0xbd,
	0xc2, 0x0b, 0xb6, 0xc3, 0x6f, 0x45, 0xc9, 0x96, 0xbb, 0xba, 0x58, 0x83, 0x75, 0x50, 0x8c, 0x80,
	0x9d, 0xaa, 0x24, 0x8b, 0xd5, 0x66, 0xf3, 0x6c, 0x81, 0xa7, 0xb8, 0x92, 0xed, 0xf0, 0x3b, 0x21,
	0xbc, 0xa3, 0x98, 0x94, 0x28, 0xba, 0x3e, 0x9b, 0x9b, 0x2c, 0xd6, 0x2e, 0x30, 0xa9, 0xb7, 0x5b,
	0x9d, 0x26, 0x72, 0xb3, 0x4e, 0x87, 0xd6, 0xce, 0x87, 0xeb, 0xc6, 0x68, 0xa2, 0x63, 0xea, 0xb7,
	0x3c, 0x4a, 0x55, 0x4f, 0xeb, 0xc8, 0x6d, 0xf9, 0x92, 0xdf, 0xf2, 0x30, 0xa4, 0x8c, 0x55, 0xac,
	0x63, 0xed, 0x61, 0x4e, 0x03, 0x06, 0xbd, 0xec, 0x88, 0x12, 0x6a, 0xe4, 0x33, 0xa1, 0x39, 0xf5,
	0x11, 0xe4, 0xdd, 0x2c, 0xc1, 0x47, 0xaa, 0x35, 0xaf, 0xff, 0x2c, 0x26, 0xae, 0xef, 0x8a, 0xcc,
	0xcd, 0x5d, 0x91, 0xf9, 0xe3, 0xae, 0xc8, 0xbc, 0xb9, 0x2f, 0x26, 0x6e, 0xee, 0x8b, 0x89, 0x5f,
	0xef, 0x8b, 0x89, 0x6f, 0x84, 0xd0, 0x75, 0xe8, 0x5e, 0x3b, 0x07, 0xee, 0xe5, 0x42, 0x8c, 0xb1,
	0x77, 0xa8, 0x5c, 0x85, 0x7e, 0x75, 0x78, 0x57, 0xe3, 0x20, 0xe3, 0x5d, 0x6a, 0x1f, 0xfe, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x6e, 0x04, 0x6b, 0xd3, 0x94, 0x08, 0x00, 0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x78
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpireAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintOrder(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x72
	if m.BatchId != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.BatchId))
		i--
		dAtA[i] = 0x68
	}
	{
		size := m.OpenAmount.Size()
		i -= size
		if _, err := m.OpenAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size, err := m.ReceivedCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size, err := m.RemainingOfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size, err := m.OfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.Direction != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Orderer) > 0 {
		i -= len(m.Orderer)
		copy(dAtA[i:], m.Orderer)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Orderer)))
		i--
		dAtA[i] = 0x2a
	}
	if m.MsgHeight != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.MsgHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.PairId != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovOrder(uint64(m.Type))
	}
	if m.Id != 0 {
		n += 1 + sovOrder(uint64(m.Id))
	}
	if m.PairId != 0 {
		n += 1 + sovOrder(uint64(m.PairId))
	}
	if m.MsgHeight != 0 {
		n += 1 + sovOrder(uint64(m.MsgHeight))
	}
	l = len(m.Orderer)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.Direction != 0 {
		n += 1 + sovOrder(uint64(m.Direction))
	}
	l = m.OfferCoin.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.RemainingOfferCoin.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.ReceivedCoin.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.OpenAmount.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.BatchId != 0 {
		n += 1 + sovOrder(uint64(m.BatchId))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireAt)
	n += 1 + l + sovOrder(uint64(l))
	if m.Status != 0 {
		n += 1 + sovOrder(uint64(m.Status))
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= OrderType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgHeight", wireType)
			}
			m.MsgHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orderer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orderer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= OrderDirection(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingOfferCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainingOfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReceivedCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OpenAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchId", wireType)
			}
			m.BatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpireAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OrderStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)
