// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: glory_road.proto

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

// 玩家球星之路信息
type GloryroadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonId       int32         `protobuf:"varint,1,opt,name=seasonId,proto3" json:"seasonId,omitempty"`             // 赛季id
	HighestCup     int32         `protobuf:"varint,2,opt,name=highestCup,proto3" json:"highestCup,omitempty"`         // 最高杯数
	CurrentCup     int32         `protobuf:"varint,3,opt,name=currentCup,proto3" json:"currentCup,omitempty"`         // 当前杯数
	Received       []int32       `protobuf:"varint,4,rep,packed,name=received,proto3" json:"received,omitempty"`      // 已领取的奖励，对应杯数
	STime          int64         `protobuf:"varint,5,opt,name=sTime,proto3" json:"sTime,omitempty"`                   // 赛季剩余秒数
	CupResetShowed bool          `protobuf:"varint,6,opt,name=cupResetShowed,proto3" json:"cupResetShowed,omitempty"` // 赛季杯数重置是否已显示过
	SeasonReward   []*ItemUpdate `protobuf:"bytes,7,rep,name=seasonReward,proto3" json:"seasonReward,omitempty"`      // 赛季奖励是否已领取
}

func (x *GloryroadInfo) Reset() {
	*x = GloryroadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glory_road_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GloryroadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GloryroadInfo) ProtoMessage() {}

func (x *GloryroadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_glory_road_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GloryroadInfo.ProtoReflect.Descriptor instead.
func (*GloryroadInfo) Descriptor() ([]byte, []int) {
	return file_glory_road_proto_rawDescGZIP(), []int{0}
}

func (x *GloryroadInfo) GetSeasonId() int32 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *GloryroadInfo) GetHighestCup() int32 {
	if x != nil {
		return x.HighestCup
	}
	return 0
}

func (x *GloryroadInfo) GetCurrentCup() int32 {
	if x != nil {
		return x.CurrentCup
	}
	return 0
}

func (x *GloryroadInfo) GetReceived() []int32 {
	if x != nil {
		return x.Received
	}
	return nil
}

func (x *GloryroadInfo) GetSTime() int64 {
	if x != nil {
		return x.STime
	}
	return 0
}

func (x *GloryroadInfo) GetCupResetShowed() bool {
	if x != nil {
		return x.CupResetShowed
	}
	return false
}

func (x *GloryroadInfo) GetSeasonReward() []*ItemUpdate {
	if x != nil {
		return x.SeasonReward
	}
	return nil
}

// 上赛季杯数与被扣除多少
type CupAndLess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cup  int32 `protobuf:"varint,1,opt,name=cup,proto3" json:"cup,omitempty"`   // 上赛季杯数
	Less int32 `protobuf:"varint,2,opt,name=less,proto3" json:"less,omitempty"` // 扣除
}

func (x *CupAndLess) Reset() {
	*x = CupAndLess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glory_road_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CupAndLess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CupAndLess) ProtoMessage() {}

func (x *CupAndLess) ProtoReflect() protoreflect.Message {
	mi := &file_glory_road_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CupAndLess.ProtoReflect.Descriptor instead.
func (*CupAndLess) Descriptor() ([]byte, []int) {
	return file_glory_road_proto_rawDescGZIP(), []int{1}
}

func (x *CupAndLess) GetCup() int32 {
	if x != nil {
		return x.Cup
	}
	return 0
}

func (x *CupAndLess) GetLess() int32 {
	if x != nil {
		return x.Less
	}
	return 0
}

// 上赛季球员杯数信息
type GloryroadLastSeasonActorCup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorCupMap map[int32]*CupAndLess `protobuf:"bytes,1,rep,name=actorCupMap,proto3" json:"actorCupMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 上赛季球员杯数<actorId,Cup>
}

func (x *GloryroadLastSeasonActorCup) Reset() {
	*x = GloryroadLastSeasonActorCup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glory_road_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GloryroadLastSeasonActorCup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GloryroadLastSeasonActorCup) ProtoMessage() {}

func (x *GloryroadLastSeasonActorCup) ProtoReflect() protoreflect.Message {
	mi := &file_glory_road_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GloryroadLastSeasonActorCup.ProtoReflect.Descriptor instead.
func (*GloryroadLastSeasonActorCup) Descriptor() ([]byte, []int) {
	return file_glory_road_proto_rawDescGZIP(), []int{2}
}

func (x *GloryroadLastSeasonActorCup) GetActorCupMap() map[int32]*CupAndLess {
	if x != nil {
		return x.ActorCupMap
	}
	return nil
}

// 领取球星之路奖励
type GloryroadRewardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cup    int32  `protobuf:"varint,1,opt,name=cup,proto3" json:"cup,omitempty"`       //杯数
	Belong uint32 `protobuf:"varint,2,opt,name=belong,proto3" json:"belong,omitempty"` // 球员id，如果奖励是球员熟练度
}

func (x *GloryroadRewardReq) Reset() {
	*x = GloryroadRewardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glory_road_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GloryroadRewardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GloryroadRewardReq) ProtoMessage() {}

func (x *GloryroadRewardReq) ProtoReflect() protoreflect.Message {
	mi := &file_glory_road_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GloryroadRewardReq.ProtoReflect.Descriptor instead.
func (*GloryroadRewardReq) Descriptor() ([]byte, []int) {
	return file_glory_road_proto_rawDescGZIP(), []int{3}
}

func (x *GloryroadRewardReq) GetCup() int32 {
	if x != nil {
		return x.Cup
	}
	return 0
}

func (x *GloryroadRewardReq) GetBelong() uint32 {
	if x != nil {
		return x.Belong
	}
	return 0
}

// 领取球星之路奖励返回
type GloryroadRewardResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info         *GloryroadInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`                 // 球星之路信息
	RewardResult *RewardResultRes `protobuf:"bytes,2,opt,name=rewardResult,proto3" json:"rewardResult,omitempty"` // 领奖结果
}

func (x *GloryroadRewardResp) Reset() {
	*x = GloryroadRewardResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glory_road_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GloryroadRewardResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GloryroadRewardResp) ProtoMessage() {}

func (x *GloryroadRewardResp) ProtoReflect() protoreflect.Message {
	mi := &file_glory_road_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GloryroadRewardResp.ProtoReflect.Descriptor instead.
func (*GloryroadRewardResp) Descriptor() ([]byte, []int) {
	return file_glory_road_proto_rawDescGZIP(), []int{4}
}

func (x *GloryroadRewardResp) GetInfo() *GloryroadInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *GloryroadRewardResp) GetRewardResult() *RewardResultRes {
	if x != nil {
		return x.RewardResult
	}
	return nil
}

var File_glory_road_proto protoreflect.FileDescriptor

var file_glory_road_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x6c, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6,
	0x01, 0x0a, 0x0d, 0x47, 0x6c, 0x6f, 0x72, 0x79, 0x72, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x43, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x43, 0x75, 0x70, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x63, 0x75, 0x70, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x77, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x75, 0x70, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0x32, 0x0a, 0x0a, 0x43, 0x75, 0x70, 0x41, 0x6e,
	0x64, 0x4c, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x63, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x65, 0x73, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x1b,
	0x47, 0x6c, 0x6f, 0x72, 0x79, 0x72, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x70, 0x12, 0x4f, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x70, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x47, 0x6c, 0x6f, 0x72, 0x79, 0x72, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x73, 0x74,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x70, 0x2e, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x70, 0x4d, 0x61, 0x70, 0x1a, 0x4b, 0x0a, 0x10,
	0x41, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x75, 0x70, 0x41, 0x6e, 0x64, 0x4c, 0x65, 0x73, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x12, 0x47, 0x6c, 0x6f,
	0x72, 0x79, 0x72, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x75,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x6f, 0x0a, 0x13, 0x47, 0x6c, 0x6f,
	0x72, 0x79, 0x72, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x47, 0x6c, 0x6f, 0x72, 0x79, 0x72, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0c, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x52, 0x0c, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_glory_road_proto_rawDescOnce sync.Once
	file_glory_road_proto_rawDescData = file_glory_road_proto_rawDesc
)

func file_glory_road_proto_rawDescGZIP() []byte {
	file_glory_road_proto_rawDescOnce.Do(func() {
		file_glory_road_proto_rawDescData = protoimpl.X.CompressGZIP(file_glory_road_proto_rawDescData)
	})
	return file_glory_road_proto_rawDescData
}

var file_glory_road_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_glory_road_proto_goTypes = []interface{}{
	(*GloryroadInfo)(nil),               // 0: GloryroadInfo
	(*CupAndLess)(nil),                  // 1: CupAndLess
	(*GloryroadLastSeasonActorCup)(nil), // 2: GloryroadLastSeasonActorCup
	(*GloryroadRewardReq)(nil),          // 3: GloryroadRewardReq
	(*GloryroadRewardResp)(nil),         // 4: GloryroadRewardResp
	nil,                                 // 5: GloryroadLastSeasonActorCup.ActorCupMapEntry
	(*ItemUpdate)(nil),                  // 6: ItemUpdate
	(*RewardResultRes)(nil),             // 7: RewardResultRes
}
var file_glory_road_proto_depIdxs = []int32{
	6, // 0: GloryroadInfo.seasonReward:type_name -> ItemUpdate
	5, // 1: GloryroadLastSeasonActorCup.actorCupMap:type_name -> GloryroadLastSeasonActorCup.ActorCupMapEntry
	0, // 2: GloryroadRewardResp.info:type_name -> GloryroadInfo
	7, // 3: GloryroadRewardResp.rewardResult:type_name -> RewardResultRes
	1, // 4: GloryroadLastSeasonActorCup.ActorCupMapEntry.value:type_name -> CupAndLess
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_glory_road_proto_init() }
func file_glory_road_proto_init() {
	if File_glory_road_proto != nil {
		return
	}
	file_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_glory_road_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GloryroadInfo); i {
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
		file_glory_road_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CupAndLess); i {
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
		file_glory_road_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GloryroadLastSeasonActorCup); i {
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
		file_glory_road_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GloryroadRewardReq); i {
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
		file_glory_road_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GloryroadRewardResp); i {
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
			RawDescriptor: file_glory_road_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_glory_road_proto_goTypes,
		DependencyIndexes: file_glory_road_proto_depIdxs,
		MessageInfos:      file_glory_road_proto_msgTypes,
	}.Build()
	File_glory_road_proto = out.File
	file_glory_road_proto_rawDesc = nil
	file_glory_road_proto_goTypes = nil
	file_glory_road_proto_depIdxs = nil
}
