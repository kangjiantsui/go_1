// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: shop_new.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 商品状态
type PRODUCT_STATUS int32

const (
	PRODUCT_STATUS_PRODUCT_CAN_BUY     PRODUCT_STATUS = 0 // 未购买，可购买
	PRODUCT_STATUS_PRODUCT_ALREADY_BUY PRODUCT_STATUS = 1 // 已购买
	PRODUCT_STATUS_PRODUCT_ALREADY_GET PRODUCT_STATUS = 2 // 已通过其他方式获取，不能买
)

// Enum value maps for PRODUCT_STATUS.
var (
	PRODUCT_STATUS_name = map[int32]string{
		0: "PRODUCT_CAN_BUY",
		1: "PRODUCT_ALREADY_BUY",
		2: "PRODUCT_ALREADY_GET",
	}
	PRODUCT_STATUS_value = map[string]int32{
		"PRODUCT_CAN_BUY":     0,
		"PRODUCT_ALREADY_BUY": 1,
		"PRODUCT_ALREADY_GET": 2,
	}
)

func (x PRODUCT_STATUS) Enum() *PRODUCT_STATUS {
	p := new(PRODUCT_STATUS)
	*p = x
	return p
}

func (x PRODUCT_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PRODUCT_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_shop_new_proto_enumTypes[0].Descriptor()
}

func (PRODUCT_STATUS) Type() protoreflect.EnumType {
	return &file_shop_new_proto_enumTypes[0]
}

func (x PRODUCT_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PRODUCT_STATUS.Descriptor instead.
func (PRODUCT_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{0}
}

// 商店刷新请求
type ShopNewRefreshReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubType int32 `protobuf:"varint,1,opt,name=subType,proto3" json:"subType,omitempty"` // table.0_types.@Types.ShopSubType
}

func (x *ShopNewRefreshReq) Reset() {
	*x = ShopNewRefreshReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopNewRefreshReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopNewRefreshReq) ProtoMessage() {}

func (x *ShopNewRefreshReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopNewRefreshReq.ProtoReflect.Descriptor instead.
func (*ShopNewRefreshReq) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{0}
}

func (x *ShopNewRefreshReq) GetSubType() int32 {
	if x != nil {
		return x.SubType
	}
	return 0
}

// 商店刷新返回
type ShopNewRefreshResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info             *ShopNewInfo      `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`                         // 商店信息
	ItemUpdateResult *ItemUpdateResult `protobuf:"bytes,2,opt,name=itemUpdateResult,proto3" json:"itemUpdateResult,omitempty"` // 刷新消耗
}

func (x *ShopNewRefreshResp) Reset() {
	*x = ShopNewRefreshResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopNewRefreshResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopNewRefreshResp) ProtoMessage() {}

func (x *ShopNewRefreshResp) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopNewRefreshResp.ProtoReflect.Descriptor instead.
func (*ShopNewRefreshResp) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{1}
}

func (x *ShopNewRefreshResp) GetInfo() *ShopNewInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ShopNewRefreshResp) GetItemUpdateResult() *ItemUpdateResult {
	if x != nil {
		return x.ItemUpdateResult
	}
	return nil
}

// 资源商店,当日各类型宝箱计数
type BoxCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BigChestBuyCount   int32 `protobuf:"varint,1,opt,name=bigChestBuyCount,proto3" json:"bigChestBuyCount,omitempty"`     // 大宝箱的计数
	SuperChestBuyCount int32 `protobuf:"varint,2,opt,name=superChestBuyCount,proto3" json:"superChestBuyCount,omitempty"` // 超级宝箱计数
}

func (x *BoxCount) Reset() {
	*x = BoxCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoxCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoxCount) ProtoMessage() {}

func (x *BoxCount) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoxCount.ProtoReflect.Descriptor instead.
func (*BoxCount) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{2}
}

func (x *BoxCount) GetBigChestBuyCount() int32 {
	if x != nil {
		return x.BigChestBuyCount
	}
	return 0
}

func (x *BoxCount) GetSuperChestBuyCount() int32 {
	if x != nil {
		return x.SuperChestBuyCount
	}
	return 0
}

// 商店信息
type ShopNewInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductListMap  map[int32]*ProductList `protobuf:"bytes,1,rep,name=productListMap,proto3" json:"productListMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // <table.0_types.@Types.ShopType,products>
	BoxCount        *BoxCount              `protobuf:"bytes,2,opt,name=boxCount,proto3" json:"boxCount,omitempty"`                                                                                                      // 资源商店,当日各类型宝箱计数
	RefreshCost     map[int32]*ItemUpdate  `protobuf:"bytes,3,rep,name=refreshCost,proto3" json:"refreshCost,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`       // 商店刷新消耗<table.0_types.@Types.ShopSubType,item>
	NextRefreshTime int64                  `protobuf:"varint,4,opt,name=nextRefreshTime,proto3" json:"nextRefreshTime,omitempty"`                                                                                       // 商店刷新倒计时,红点用
	ShopShowTime    map[int32]int64        `protobuf:"bytes,5,rep,name=shopShowTime,proto3" json:"shopShowTime,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`    // 商品出现倒计时,<table.Shop.Index,timeLeft>,预告用
}

func (x *ShopNewInfo) Reset() {
	*x = ShopNewInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopNewInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopNewInfo) ProtoMessage() {}

func (x *ShopNewInfo) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopNewInfo.ProtoReflect.Descriptor instead.
func (*ShopNewInfo) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{3}
}

func (x *ShopNewInfo) GetProductListMap() map[int32]*ProductList {
	if x != nil {
		return x.ProductListMap
	}
	return nil
}

func (x *ShopNewInfo) GetBoxCount() *BoxCount {
	if x != nil {
		return x.BoxCount
	}
	return nil
}

func (x *ShopNewInfo) GetRefreshCost() map[int32]*ItemUpdate {
	if x != nil {
		return x.RefreshCost
	}
	return nil
}

func (x *ShopNewInfo) GetNextRefreshTime() int64 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

func (x *ShopNewInfo) GetShopShowTime() map[int32]int64 {
	if x != nil {
		return x.ShopShowTime
	}
	return nil
}

// 商店单个subType的列表,已按规则排序
type ProductList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*ProductInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *ProductList) Reset() {
	*x = ProductList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductList) ProtoMessage() {}

func (x *ProductList) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductList.ProtoReflect.Descriptor instead.
func (*ProductList) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{4}
}

func (x *ProductList) GetInfos() []*ProductInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

// 商品单个信息
type ProductInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     int32          `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`                     // table.Shop.index
	State     PRODUCT_STATUS `protobuf:"varint,2,opt,name=state,proto3,enum=PRODUCT_STATUS" json:"state,omitempty"` // 商品状态
	TimeLeft  int64          `protobuf:"varint,4,opt,name=timeLeft,proto3" json:"timeLeft,omitempty"`               // 商品剩余时间
	Items     []*ItemUpdate  `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`                      // 商品内容
	Price     *ItemUpdate    `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`                      // 商品价格
	BirthTime int64          `protobuf:"varint,7,opt,name=birthTime,proto3" json:"birthTime,omitempty"`             // 刷出来的时间,排序用
}

func (x *ProductInfo) Reset() {
	*x = ProductInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductInfo) ProtoMessage() {}

func (x *ProductInfo) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductInfo.ProtoReflect.Descriptor instead.
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{5}
}

func (x *ProductInfo) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ProductInfo) GetState() PRODUCT_STATUS {
	if x != nil {
		return x.State
	}
	return PRODUCT_STATUS_PRODUCT_CAN_BUY
}

func (x *ProductInfo) GetTimeLeft() int64 {
	if x != nil {
		return x.TimeLeft
	}
	return 0
}

func (x *ProductInfo) GetItems() []*ItemUpdate {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ProductInfo) GetPrice() *ItemUpdate {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *ProductInfo) GetBirthTime() int64 {
	if x != nil {
		return x.BirthTime
	}
	return 0
}

// 商品购买请求
type ShopNewBuyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index  int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`   // table.Shop.Index
	Belong int32 `protobuf:"varint,2,opt,name=belong,proto3" json:"belong,omitempty"` // 熟练度买给哪个球员,table.Actor.Id
}

func (x *ShopNewBuyReq) Reset() {
	*x = ShopNewBuyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopNewBuyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopNewBuyReq) ProtoMessage() {}

func (x *ShopNewBuyReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopNewBuyReq.ProtoReflect.Descriptor instead.
func (*ShopNewBuyReq) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{6}
}

func (x *ShopNewBuyReq) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ShopNewBuyReq) GetBelong() int32 {
	if x != nil {
		return x.Belong
	}
	return 0
}

// 商品购买返回
type ShopNewBuyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info             *ShopNewInfo      `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	ItemUpdateResult *ItemUpdateResult `protobuf:"bytes,2,opt,name=itemUpdateResult,proto3" json:"itemUpdateResult,omitempty"`
}

func (x *ShopNewBuyResp) Reset() {
	*x = ShopNewBuyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopNewBuyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopNewBuyResp) ProtoMessage() {}

func (x *ShopNewBuyResp) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopNewBuyResp.ProtoReflect.Descriptor instead.
func (*ShopNewBuyResp) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{7}
}

func (x *ShopNewBuyResp) GetInfo() *ShopNewInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ShopNewBuyResp) GetItemUpdateResult() *ItemUpdateResult {
	if x != nil {
		return x.ItemUpdateResult
	}
	return nil
}

// 商品显示信息
type ShopNewShowInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index    int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`      // table.Shop.Index
	ShowTime string `protobuf:"bytes,2,opt,name=showTime,proto3" json:"showTime,omitempty"` // 开始展示时间,如 2006-01-02
	SellTime string `protobuf:"bytes,3,opt,name=sellTime,proto3" json:"sellTime,omitempty"` // 开始售卖时间
}

func (x *ShopNewShowInfo) Reset() {
	*x = ShopNewShowInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopNewShowInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopNewShowInfo) ProtoMessage() {}

func (x *ShopNewShowInfo) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopNewShowInfo.ProtoReflect.Descriptor instead.
func (*ShopNewShowInfo) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{8}
}

func (x *ShopNewShowInfo) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ShopNewShowInfo) GetShowTime() string {
	if x != nil {
		return x.ShowTime
	}
	return ""
}

func (x *ShopNewShowInfo) GetSellTime() string {
	if x != nil {
		return x.SellTime
	}
	return ""
}

// 商品显示设置
type ShopNewShowInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShowInfos []*ShopNewShowInfo `protobuf:"bytes,1,rep,name=showInfos,proto3" json:"showInfos,omitempty"`
}

func (x *ShopNewShowInfos) Reset() {
	*x = ShopNewShowInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopNewShowInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopNewShowInfos) ProtoMessage() {}

func (x *ShopNewShowInfos) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopNewShowInfos.ProtoReflect.Descriptor instead.
func (*ShopNewShowInfos) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{9}
}

func (x *ShopNewShowInfos) GetShowInfos() []*ShopNewShowInfo {
	if x != nil {
		return x.ShowInfos
	}
	return nil
}

// 删除商品显示信息
type ShopNewShowDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index []int32 `protobuf:"varint,1,rep,packed,name=index,proto3" json:"index,omitempty"`
}

func (x *ShopNewShowDelReq) Reset() {
	*x = ShopNewShowDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_new_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopNewShowDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopNewShowDelReq) ProtoMessage() {}

func (x *ShopNewShowDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_new_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopNewShowDelReq.ProtoReflect.Descriptor instead.
func (*ShopNewShowDelReq) Descriptor() ([]byte, []int) {
	return file_shop_new_proto_rawDescGZIP(), []int{10}
}

func (x *ShopNewShowDelReq) GetIndex() []int32 {
	if x != nil {
		return x.Index
	}
	return nil
}

var File_shop_new_proto protoreflect.FileDescriptor

var file_shop_new_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x11,
	0x53, 0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x22, 0x75, 0x0a, 0x12, 0x53,
	0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x66, 0x0a, 0x08, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x10, 0x62, 0x69, 0x67, 0x43, 0x68, 0x65, 0x73, 0x74, 0x42, 0x75, 0x79, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x62, 0x69, 0x67, 0x43, 0x68, 0x65,
	0x73, 0x74, 0x42, 0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x43, 0x68, 0x65, 0x73, 0x74, 0x42, 0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x73, 0x75, 0x70, 0x65, 0x72, 0x43, 0x68, 0x65,
	0x73, 0x74, 0x42, 0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8c, 0x04, 0x0a, 0x0b, 0x53,
	0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x61, 0x70, 0x12, 0x25, 0x0a, 0x08, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x08, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x53, 0x68,
	0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x53,
	0x68, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x68,
	0x6f, 0x70, 0x53, 0x68, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x4f, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x10, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x70,
	0x53, 0x68, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x31, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0xca, 0x01, 0x0a,
	0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x4c, 0x65, 0x66, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x0d, 0x53, 0x68, 0x6f,
	0x70, 0x4e, 0x65, 0x77, 0x42, 0x75, 0x79, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x71, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x70,
	0x4e, 0x65, 0x77, 0x42, 0x75, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x4e,
	0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x10,
	0x69, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5f, 0x0a, 0x0f, 0x53,
	0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x10,
	0x53, 0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x2e, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f,
	0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x22, 0x29, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f, 0x77, 0x44,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2a, 0x57, 0x0a, 0x0e, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x42, 0x55, 0x59,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x41, 0x4c,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x47,
	0x45, 0x54, 0x10, 0x02, 0x42, 0x0b, 0x5a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_new_proto_rawDescOnce sync.Once
	file_shop_new_proto_rawDescData = file_shop_new_proto_rawDesc
)

func file_shop_new_proto_rawDescGZIP() []byte {
	file_shop_new_proto_rawDescOnce.Do(func() {
		file_shop_new_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_new_proto_rawDescData)
	})
	return file_shop_new_proto_rawDescData
}

var file_shop_new_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shop_new_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_shop_new_proto_goTypes = []interface{}{
	(PRODUCT_STATUS)(0),        // 0: PRODUCT_STATUS
	(*ShopNewRefreshReq)(nil),  // 1: ShopNewRefreshReq
	(*ShopNewRefreshResp)(nil), // 2: ShopNewRefreshResp
	(*BoxCount)(nil),           // 3: BoxCount
	(*ShopNewInfo)(nil),        // 4: ShopNewInfo
	(*ProductList)(nil),        // 5: ProductList
	(*ProductInfo)(nil),        // 6: ProductInfo
	(*ShopNewBuyReq)(nil),      // 7: ShopNewBuyReq
	(*ShopNewBuyResp)(nil),     // 8: ShopNewBuyResp
	(*ShopNewShowInfo)(nil),    // 9: ShopNewShowInfo
	(*ShopNewShowInfos)(nil),   // 10: ShopNewShowInfos
	(*ShopNewShowDelReq)(nil),  // 11: ShopNewShowDelReq
	nil,                        // 12: ShopNewInfo.ProductListMapEntry
	nil,                        // 13: ShopNewInfo.RefreshCostEntry
	nil,                        // 14: ShopNewInfo.ShopShowTimeEntry
	(*ItemUpdateResult)(nil),   // 15: ItemUpdateResult
	(*ItemUpdate)(nil),         // 16: ItemUpdate
}
var file_shop_new_proto_depIdxs = []int32{
	4,  // 0: ShopNewRefreshResp.info:type_name -> ShopNewInfo
	15, // 1: ShopNewRefreshResp.itemUpdateResult:type_name -> ItemUpdateResult
	12, // 2: ShopNewInfo.productListMap:type_name -> ShopNewInfo.ProductListMapEntry
	3,  // 3: ShopNewInfo.boxCount:type_name -> BoxCount
	13, // 4: ShopNewInfo.refreshCost:type_name -> ShopNewInfo.RefreshCostEntry
	14, // 5: ShopNewInfo.shopShowTime:type_name -> ShopNewInfo.ShopShowTimeEntry
	6,  // 6: ProductList.infos:type_name -> ProductInfo
	0,  // 7: ProductInfo.state:type_name -> PRODUCT_STATUS
	16, // 8: ProductInfo.items:type_name -> ItemUpdate
	16, // 9: ProductInfo.price:type_name -> ItemUpdate
	4,  // 10: ShopNewBuyResp.info:type_name -> ShopNewInfo
	15, // 11: ShopNewBuyResp.itemUpdateResult:type_name -> ItemUpdateResult
	9,  // 12: ShopNewShowInfos.showInfos:type_name -> ShopNewShowInfo
	5,  // 13: ShopNewInfo.ProductListMapEntry.value:type_name -> ProductList
	16, // 14: ShopNewInfo.RefreshCostEntry.value:type_name -> ItemUpdate
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_shop_new_proto_init() }
func file_shop_new_proto_init() {
	if File_shop_new_proto != nil {
		return
	}
	file_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shop_new_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopNewRefreshReq); i {
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
		file_shop_new_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopNewRefreshResp); i {
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
		file_shop_new_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoxCount); i {
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
		file_shop_new_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopNewInfo); i {
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
		file_shop_new_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductList); i {
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
		file_shop_new_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductInfo); i {
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
		file_shop_new_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopNewBuyReq); i {
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
		file_shop_new_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopNewBuyResp); i {
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
		file_shop_new_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopNewShowInfo); i {
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
		file_shop_new_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopNewShowInfos); i {
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
		file_shop_new_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopNewShowDelReq); i {
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
			RawDescriptor: file_shop_new_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shop_new_proto_goTypes,
		DependencyIndexes: file_shop_new_proto_depIdxs,
		EnumInfos:         file_shop_new_proto_enumTypes,
		MessageInfos:      file_shop_new_proto_msgTypes,
	}.Build()
	File_shop_new_proto = out.File
	file_shop_new_proto_rawDesc = nil
	file_shop_new_proto_goTypes = nil
	file_shop_new_proto_depIdxs = nil
}