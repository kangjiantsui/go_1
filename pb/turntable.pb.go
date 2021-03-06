// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: turntable.proto

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

// 玩家转盘信息
type TurntableInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level               int32                  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`                                                                                                   // 转盘等级
	TimeLeft            int64                  `protobuf:"varint,2,opt,name=timeLeft,proto3" json:"timeLeft,omitempty"`                                                                                             // 转盘等级重置倒计时秒
	AdNum               int32                  `protobuf:"varint,3,opt,name=adNum,proto3" json:"adNum,omitempty"`                                                                                                   // 广告看了多少次(玩转盘会消耗这个值)
	AdCount             int32                  `protobuf:"varint,4,opt,name=adCount,proto3" json:"adCount,omitempty"`                                                                                               // 当日广告观看总次数
	LatticeMap          map[int32]*LatticeInfo `protobuf:"bytes,5,rep,name=latticeMap,proto3" json:"latticeMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // <格子,格子信息>
	LevelUpProgressRate int32                  `protobuf:"varint,6,opt,name=levelUpProgressRate,proto3" json:"levelUpProgressRate,omitempty"`                                                                       // 转盘升级进度(当前五连抽几次了)
	ItemProbability     []*ItemProbability     `protobuf:"bytes,7,rep,name=itemProbability,proto3" json:"itemProbability,omitempty"`                                                                                // 各种物品出现概率
	AlreadyUsedFree     bool                   `protobuf:"varint,8,opt,name=alreadyUsedFree,proto3" json:"alreadyUsedFree,omitempty"`                                                                               // 每日免费次数是否已被使用
}

func (x *TurntableInfo) Reset() {
	*x = TurntableInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turntable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TurntableInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TurntableInfo) ProtoMessage() {}

func (x *TurntableInfo) ProtoReflect() protoreflect.Message {
	mi := &file_turntable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TurntableInfo.ProtoReflect.Descriptor instead.
func (*TurntableInfo) Descriptor() ([]byte, []int) {
	return file_turntable_proto_rawDescGZIP(), []int{0}
}

func (x *TurntableInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *TurntableInfo) GetTimeLeft() int64 {
	if x != nil {
		return x.TimeLeft
	}
	return 0
}

func (x *TurntableInfo) GetAdNum() int32 {
	if x != nil {
		return x.AdNum
	}
	return 0
}

func (x *TurntableInfo) GetAdCount() int32 {
	if x != nil {
		return x.AdCount
	}
	return 0
}

func (x *TurntableInfo) GetLatticeMap() map[int32]*LatticeInfo {
	if x != nil {
		return x.LatticeMap
	}
	return nil
}

func (x *TurntableInfo) GetLevelUpProgressRate() int32 {
	if x != nil {
		return x.LevelUpProgressRate
	}
	return 0
}

func (x *TurntableInfo) GetItemProbability() []*ItemProbability {
	if x != nil {
		return x.ItemProbability
	}
	return nil
}

func (x *TurntableInfo) GetAlreadyUsedFree() bool {
	if x != nil {
		return x.AlreadyUsedFree
	}
	return false
}

type ItemProbability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        ITEM  `protobuf:"varint,1,opt,name=type,proto3,enum=ITEM" json:"type,omitempty"`     // item类型
	Id          int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`                   // id
	Probability int32 `protobuf:"varint,3,opt,name=probability,proto3" json:"probability,omitempty"` // 出现概率
}

func (x *ItemProbability) Reset() {
	*x = ItemProbability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turntable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemProbability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemProbability) ProtoMessage() {}

func (x *ItemProbability) ProtoReflect() protoreflect.Message {
	mi := &file_turntable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemProbability.ProtoReflect.Descriptor instead.
func (*ItemProbability) Descriptor() ([]byte, []int) {
	return file_turntable_proto_rawDescGZIP(), []int{1}
}

func (x *ItemProbability) GetType() ITEM {
	if x != nil {
		return x.Type
	}
	return ITEM_ZERO
}

func (x *ItemProbability) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemProbability) GetProbability() int32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

type LatticeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index   int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`     // 对应产出表index
	Replace bool  `protobuf:"varint,2,opt,name=replace,proto3" json:"replace,omitempty"` // 是否被替换,已拥有的皮肤表情等
	Belong  int32 `protobuf:"varint,3,opt,name=belong,proto3" json:"belong,omitempty"`   // 未指定球员的熟练度随机出来的球员id
}

func (x *LatticeInfo) Reset() {
	*x = LatticeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turntable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatticeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatticeInfo) ProtoMessage() {}

func (x *LatticeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_turntable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatticeInfo.ProtoReflect.Descriptor instead.
func (*LatticeInfo) Descriptor() ([]byte, []int) {
	return file_turntable_proto_rawDescGZIP(), []int{2}
}

func (x *LatticeInfo) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LatticeInfo) GetReplace() bool {
	if x != nil {
		return x.Replace
	}
	return false
}

func (x *LatticeInfo) GetBelong() int32 {
	if x != nil {
		return x.Belong
	}
	return 0
}

type AdBeginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId string `protobuf:"bytes,1,opt,name=adId,proto3" json:"adId,omitempty"`
}

func (x *AdBeginResp) Reset() {
	*x = AdBeginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turntable_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdBeginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdBeginResp) ProtoMessage() {}

func (x *AdBeginResp) ProtoReflect() protoreflect.Message {
	mi := &file_turntable_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdBeginResp.ProtoReflect.Descriptor instead.
func (*AdBeginResp) Descriptor() ([]byte, []int) {
	return file_turntable_proto_rawDescGZIP(), []int{3}
}

func (x *AdBeginResp) GetAdId() string {
	if x != nil {
		return x.AdId
	}
	return ""
}

// 看广告完成
type AdFinishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId string `protobuf:"bytes,1,opt,name=adId,proto3" json:"adId,omitempty"`
}

func (x *AdFinishReq) Reset() {
	*x = AdFinishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turntable_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdFinishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdFinishReq) ProtoMessage() {}

func (x *AdFinishReq) ProtoReflect() protoreflect.Message {
	mi := &file_turntable_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdFinishReq.ProtoReflect.Descriptor instead.
func (*AdFinishReq) Descriptor() ([]byte, []int) {
	return file_turntable_proto_rawDescGZIP(), []int{4}
}

func (x *AdFinishReq) GetAdId() string {
	if x != nil {
		return x.AdId
	}
	return ""
}

// 玩转盘
type TurntableReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FiveConsecutive bool `protobuf:"varint,1,opt,name=fiveConsecutive,proto3" json:"fiveConsecutive,omitempty"` // 是否五连抽
	Free            bool `protobuf:"varint,2,opt,name=free,proto3" json:"free,omitempty"`                       // 是否使用免费次数
}

func (x *TurntableReq) Reset() {
	*x = TurntableReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turntable_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TurntableReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TurntableReq) ProtoMessage() {}

func (x *TurntableReq) ProtoReflect() protoreflect.Message {
	mi := &file_turntable_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TurntableReq.ProtoReflect.Descriptor instead.
func (*TurntableReq) Descriptor() ([]byte, []int) {
	return file_turntable_proto_rawDescGZIP(), []int{5}
}

func (x *TurntableReq) GetFiveConsecutive() bool {
	if x != nil {
		return x.FiveConsecutive
	}
	return false
}

func (x *TurntableReq) GetFree() bool {
	if x != nil {
		return x.Free
	}
	return false
}

// 玩转盘结果
type TurntableResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info             *TurntableInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	ItemUpdateResult *ItemUpdateResult `protobuf:"bytes,2,opt,name=itemUpdateResult,proto3" json:"itemUpdateResult,omitempty"`
	HitLattice       []int32           `protobuf:"varint,3,rep,packed,name=hitLattice,proto3" json:"hitLattice,omitempty"` // 转到的格子
}

func (x *TurntableResp) Reset() {
	*x = TurntableResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turntable_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TurntableResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TurntableResp) ProtoMessage() {}

func (x *TurntableResp) ProtoReflect() protoreflect.Message {
	mi := &file_turntable_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TurntableResp.ProtoReflect.Descriptor instead.
func (*TurntableResp) Descriptor() ([]byte, []int) {
	return file_turntable_proto_rawDescGZIP(), []int{6}
}

func (x *TurntableResp) GetInfo() *TurntableInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *TurntableResp) GetItemUpdateResult() *ItemUpdateResult {
	if x != nil {
		return x.ItemUpdateResult
	}
	return nil
}

func (x *TurntableResp) GetHitLattice() []int32 {
	if x != nil {
		return x.HitLattice
	}
	return nil
}

var File_turntable_proto protoreflect.FileDescriptor

var file_turntable_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x75, 0x72, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x0d,
	0x54, 0x75, 0x72, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x65, 0x66, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x65, 0x66, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x61, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3e, 0x0a, 0x0a, 0x6c, 0x61, 0x74, 0x74, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x4c, 0x61, 0x74, 0x74, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x74, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x12,
	0x30, 0x0a, 0x13, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x55, 0x70, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x69, 0x74,
	0x65, 0x6d, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a,
	0x0f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x55, 0x73, 0x65, 0x64, 0x46, 0x72, 0x65, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x55,
	0x73, 0x65, 0x64, 0x46, 0x72, 0x65, 0x65, 0x1a, 0x4b, 0x0a, 0x0f, 0x4c, 0x61, 0x74, 0x74, 0x69,
	0x63, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x61,
	0x74, 0x74, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x5e, 0x0a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x62,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x49, 0x54, 0x45, 0x4d, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x22, 0x55, 0x0a, 0x0b, 0x4c, 0x61, 0x74, 0x74, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x21, 0x0a, 0x0b, 0x41,
	0x64, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x22, 0x21,
	0x0a, 0x0b, 0x41, 0x64, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x49,
	0x64, 0x22, 0x4c, 0x0a, 0x0c, 0x54, 0x75, 0x72, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x66, 0x69, 0x76, 0x65,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x72, 0x65, 0x65, 0x22,
	0x92, 0x01, 0x0a, 0x0d, 0x54, 0x75, 0x72, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x74, 0x4c, 0x61, 0x74, 0x74, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x74, 0x4c, 0x61, 0x74,
	0x74, 0x69, 0x63, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_turntable_proto_rawDescOnce sync.Once
	file_turntable_proto_rawDescData = file_turntable_proto_rawDesc
)

func file_turntable_proto_rawDescGZIP() []byte {
	file_turntable_proto_rawDescOnce.Do(func() {
		file_turntable_proto_rawDescData = protoimpl.X.CompressGZIP(file_turntable_proto_rawDescData)
	})
	return file_turntable_proto_rawDescData
}

var file_turntable_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_turntable_proto_goTypes = []interface{}{
	(*TurntableInfo)(nil),    // 0: TurntableInfo
	(*ItemProbability)(nil),  // 1: ItemProbability
	(*LatticeInfo)(nil),      // 2: LatticeInfo
	(*AdBeginResp)(nil),      // 3: AdBeginResp
	(*AdFinishReq)(nil),      // 4: AdFinishReq
	(*TurntableReq)(nil),     // 5: TurntableReq
	(*TurntableResp)(nil),    // 6: TurntableResp
	nil,                      // 7: TurntableInfo.LatticeMapEntry
	(ITEM)(0),                // 8: ITEM
	(*ItemUpdateResult)(nil), // 9: ItemUpdateResult
}
var file_turntable_proto_depIdxs = []int32{
	7, // 0: TurntableInfo.latticeMap:type_name -> TurntableInfo.LatticeMapEntry
	1, // 1: TurntableInfo.itemProbability:type_name -> ItemProbability
	8, // 2: ItemProbability.type:type_name -> ITEM
	0, // 3: TurntableResp.info:type_name -> TurntableInfo
	9, // 4: TurntableResp.itemUpdateResult:type_name -> ItemUpdateResult
	2, // 5: TurntableInfo.LatticeMapEntry.value:type_name -> LatticeInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_turntable_proto_init() }
func file_turntable_proto_init() {
	if File_turntable_proto != nil {
		return
	}
	file_item_proto_init()
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_turntable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TurntableInfo); i {
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
		file_turntable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemProbability); i {
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
		file_turntable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatticeInfo); i {
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
		file_turntable_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdBeginResp); i {
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
		file_turntable_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdFinishReq); i {
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
		file_turntable_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TurntableReq); i {
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
		file_turntable_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TurntableResp); i {
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
			RawDescriptor: file_turntable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_turntable_proto_goTypes,
		DependencyIndexes: file_turntable_proto_depIdxs,
		MessageInfos:      file_turntable_proto_msgTypes,
	}.Build()
	File_turntable_proto = out.File
	file_turntable_proto_rawDesc = nil
	file_turntable_proto_goTypes = nil
	file_turntable_proto_depIdxs = nil
}
