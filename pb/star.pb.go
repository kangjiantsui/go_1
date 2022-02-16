// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: star.proto

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

type StarRewardType int32

const (
	StarRewardType_RewardNone  StarRewardType = 0 // 没有奖励
	StarRewardType_RewardDaily StarRewardType = 1 // 每日奖励
	StarRewardType_RewardUp    StarRewardType = 2 // 升级奖励
)

// Enum value maps for StarRewardType.
var (
	StarRewardType_name = map[int32]string{
		0: "RewardNone",
		1: "RewardDaily",
		2: "RewardUp",
	}
	StarRewardType_value = map[string]int32{
		"RewardNone":  0,
		"RewardDaily": 1,
		"RewardUp":    2,
	}
)

func (x StarRewardType) Enum() *StarRewardType {
	p := new(StarRewardType)
	*p = x
	return p
}

func (x StarRewardType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StarRewardType) Descriptor() protoreflect.EnumDescriptor {
	return file_star_proto_enumTypes[0].Descriptor()
}

func (StarRewardType) Type() protoreflect.EnumType {
	return &file_star_proto_enumTypes[0]
}

func (x StarRewardType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StarRewardType.Descriptor instead.
func (StarRewardType) EnumDescriptor() ([]byte, []int) {
	return file_star_proto_rawDescGZIP(), []int{0}
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WinState WinState `protobuf:"varint,1,opt,name=winState,proto3,enum=WinState" json:"winState,omitempty"`
	Count    int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_star_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_star_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_star_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetWinState() WinState {
	if x != nil {
		return x.WinState
	}
	return WinState_W_WIN
}

func (x *Record) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StarInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NowStar                  int32          `protobuf:"varint,1,opt,name=nowStar,proto3" json:"nowStar,omitempty"`                                    // 当前星辉数
	NowStarLevel             int32          `protobuf:"varint,2,opt,name=nowStarLevel,proto3" json:"nowStarLevel,omitempty"`                          // 当前星辉等级 （约战等级）
	MaxStar                  int32          `protobuf:"varint,3,opt,name=maxStar,proto3" json:"maxStar,omitempty"`                                    // 历史最大星辉数
	MaxStarLevel             int32          `protobuf:"varint,4,opt,name=maxStarLevel,proto3" json:"maxStarLevel,omitempty"`                          // 历史最大星辉等级 （约战等级）
	DailyStarGet             int32          `protobuf:"varint,5,opt,name=dailyStarGet,proto3" json:"dailyStarGet,omitempty"`                          // 今日获得的星辉数     to delete
	TodayRecord              bool           `protobuf:"varint,6,opt,name=todayRecord,proto3" json:"todayRecord,omitempty"`                            // 当天奖励是否已领取    to delete
	NowRewardType            StarRewardType `protobuf:"varint,7,opt,name=nowRewardType,proto3,enum=StarRewardType" json:"nowRewardType,omitempty"`    // 当前可领取的奖励类型    to delete
	RefreshTime              int64          `protobuf:"varint,8,opt,name=refreshTime,proto3" json:"refreshTime,omitempty"`                            // 奖励下次刷新倒计时     to delete
	NextRewardScale          int32          `protobuf:"varint,9,opt,name=nextRewardScale,proto3" json:"nextRewardScale,omitempty"`                    // 下一次可领奖的刻度  根据maxStar计算
	NextRewardBoxId          uint32         `protobuf:"varint,10,opt,name=nextRewardBoxId,proto3" json:"nextRewardBoxId,omitempty"`                   // 下一次可领奖的宝箱id  根据maxStar计算
	BestAvailableRewardBoxId uint32         `protobuf:"varint,11,opt,name=bestAvailableRewardBoxId,proto3" json:"bestAvailableRewardBoxId,omitempty"` // 可领奖的最好的宝箱id
	Record                   []*Record      `protobuf:"bytes,12,rep,name=record,proto3" json:"record,omitempty"`                                      // 战绩
}

func (x *StarInfoResp) Reset() {
	*x = StarInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_star_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StarInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StarInfoResp) ProtoMessage() {}

func (x *StarInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_star_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StarInfoResp.ProtoReflect.Descriptor instead.
func (*StarInfoResp) Descriptor() ([]byte, []int) {
	return file_star_proto_rawDescGZIP(), []int{1}
}

func (x *StarInfoResp) GetNowStar() int32 {
	if x != nil {
		return x.NowStar
	}
	return 0
}

func (x *StarInfoResp) GetNowStarLevel() int32 {
	if x != nil {
		return x.NowStarLevel
	}
	return 0
}

func (x *StarInfoResp) GetMaxStar() int32 {
	if x != nil {
		return x.MaxStar
	}
	return 0
}

func (x *StarInfoResp) GetMaxStarLevel() int32 {
	if x != nil {
		return x.MaxStarLevel
	}
	return 0
}

func (x *StarInfoResp) GetDailyStarGet() int32 {
	if x != nil {
		return x.DailyStarGet
	}
	return 0
}

func (x *StarInfoResp) GetTodayRecord() bool {
	if x != nil {
		return x.TodayRecord
	}
	return false
}

func (x *StarInfoResp) GetNowRewardType() StarRewardType {
	if x != nil {
		return x.NowRewardType
	}
	return StarRewardType_RewardNone
}

func (x *StarInfoResp) GetRefreshTime() int64 {
	if x != nil {
		return x.RefreshTime
	}
	return 0
}

func (x *StarInfoResp) GetNextRewardScale() int32 {
	if x != nil {
		return x.NextRewardScale
	}
	return 0
}

func (x *StarInfoResp) GetNextRewardBoxId() uint32 {
	if x != nil {
		return x.NextRewardBoxId
	}
	return 0
}

func (x *StarInfoResp) GetBestAvailableRewardBoxId() uint32 {
	if x != nil {
		return x.BestAvailableRewardBoxId
	}
	return 0
}

func (x *StarInfoResp) GetRecord() []*Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type EngageRecommendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"` // 请求页面数
}

func (x *EngageRecommendReq) Reset() {
	*x = EngageRecommendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_star_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EngageRecommendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EngageRecommendReq) ProtoMessage() {}

func (x *EngageRecommendReq) ProtoReflect() protoreflect.Message {
	mi := &file_star_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EngageRecommendReq.ProtoReflect.Descriptor instead.
func (*EngageRecommendReq) Descriptor() ([]byte, []int) {
	return file_star_proto_rawDescGZIP(), []int{2}
}

func (x *EngageRecommendReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type EngageRecommendPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ic           uint64 `protobuf:"varint,1,opt,name=ic,proto3" json:"ic,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar       string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	PlayerAvatar int32  `protobuf:"varint,4,opt,name=playerAvatar,proto3" json:"playerAvatar,omitempty"`
	Star         int32  `protobuf:"varint,5,opt,name=star,proto3" json:"star,omitempty"`
	StarLevel    int32  `protobuf:"varint,6,opt,name=starLevel,proto3" json:"starLevel,omitempty"`     // 星辉等级（约战等级）
	FightStatus  int32  `protobuf:"varint,7,opt,name=fightStatus,proto3" json:"fightStatus,omitempty"` // 战斗状态 0空闲 1战斗中 2约战房间已满
	NowCup       int32  `protobuf:"varint,8,opt,name=nowCup,proto3" json:"nowCup,omitempty"`           // 当前杯数
	AvatarFrame  int32  `protobuf:"varint,9,opt,name=avatarFrame,proto3" json:"avatarFrame,omitempty"` // 头像框
	Title        int32  `protobuf:"varint,10,opt,name=title,proto3" json:"title,omitempty"`            // 炫彩名
}

func (x *EngageRecommendPlayer) Reset() {
	*x = EngageRecommendPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_star_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EngageRecommendPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EngageRecommendPlayer) ProtoMessage() {}

func (x *EngageRecommendPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_star_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EngageRecommendPlayer.ProtoReflect.Descriptor instead.
func (*EngageRecommendPlayer) Descriptor() ([]byte, []int) {
	return file_star_proto_rawDescGZIP(), []int{3}
}

func (x *EngageRecommendPlayer) GetIc() uint64 {
	if x != nil {
		return x.Ic
	}
	return 0
}

func (x *EngageRecommendPlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EngageRecommendPlayer) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *EngageRecommendPlayer) GetPlayerAvatar() int32 {
	if x != nil {
		return x.PlayerAvatar
	}
	return 0
}

func (x *EngageRecommendPlayer) GetStar() int32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *EngageRecommendPlayer) GetStarLevel() int32 {
	if x != nil {
		return x.StarLevel
	}
	return 0
}

func (x *EngageRecommendPlayer) GetFightStatus() int32 {
	if x != nil {
		return x.FightStatus
	}
	return 0
}

func (x *EngageRecommendPlayer) GetNowCup() int32 {
	if x != nil {
		return x.NowCup
	}
	return 0
}

func (x *EngageRecommendPlayer) GetAvatarFrame() int32 {
	if x != nil {
		return x.AvatarFrame
	}
	return 0
}

func (x *EngageRecommendPlayer) GetTitle() int32 {
	if x != nil {
		return x.Title
	}
	return 0
}

type EngageRecommendList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	E        []*EngageRecommendPlayer `protobuf:"bytes,1,rep,name=e,proto3" json:"e,omitempty"`                // 在线玩家推荐
	NextPage int32                    `protobuf:"varint,2,opt,name=nextPage,proto3" json:"nextPage,omitempty"` // 下次请求的页面
	F        []*EngageRecommendPlayer `protobuf:"bytes,3,rep,name=f,proto3" json:"f,omitempty"`                // 在线好友
}

func (x *EngageRecommendList) Reset() {
	*x = EngageRecommendList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_star_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EngageRecommendList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EngageRecommendList) ProtoMessage() {}

func (x *EngageRecommendList) ProtoReflect() protoreflect.Message {
	mi := &file_star_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EngageRecommendList.ProtoReflect.Descriptor instead.
func (*EngageRecommendList) Descriptor() ([]byte, []int) {
	return file_star_proto_rawDescGZIP(), []int{4}
}

func (x *EngageRecommendList) GetE() []*EngageRecommendPlayer {
	if x != nil {
		return x.E
	}
	return nil
}

func (x *EngageRecommendList) GetNextPage() int32 {
	if x != nil {
		return x.NextPage
	}
	return 0
}

func (x *EngageRecommendList) GetF() []*EngageRecommendPlayer {
	if x != nil {
		return x.F
	}
	return nil
}

var File_star_proto protoreflect.FileDescriptor

var file_star_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x06, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x57, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xda, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x6e, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x6e, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61,
	0x78, 0x53, 0x74, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22,
	0x0a, 0x0c, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x72, 0x47, 0x65, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x72, 0x47,
	0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x35, 0x0a, 0x0d, 0x6e, 0x6f, 0x77, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6e, 0x6f,
	0x77, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x6f, 0x78, 0x49,
	0x64, 0x12, 0x3a, 0x0a, 0x18, 0x62, 0x65, 0x73, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x18, 0x62, 0x65, 0x73, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x28,
	0x0a, 0x12, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x9b, 0x02, 0x0a, 0x15, 0x45, 0x6e, 0x67,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69, 0x67, 0x68, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x77, 0x43, 0x75, 0x70,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x77, 0x43, 0x75, 0x70, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x7d, 0x0a, 0x13, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x01, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x45, 0x6e, 0x67, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x01, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x24, 0x0a, 0x01, 0x66, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x45, 0x6e, 0x67,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x01, 0x66, 0x2a, 0x3f, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x55, 0x70, 0x10, 0x02, 0x42, 0x0b, 0x5a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_star_proto_rawDescOnce sync.Once
	file_star_proto_rawDescData = file_star_proto_rawDesc
)

func file_star_proto_rawDescGZIP() []byte {
	file_star_proto_rawDescOnce.Do(func() {
		file_star_proto_rawDescData = protoimpl.X.CompressGZIP(file_star_proto_rawDescData)
	})
	return file_star_proto_rawDescData
}

var file_star_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_star_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_star_proto_goTypes = []interface{}{
	(StarRewardType)(0),           // 0: StarRewardType
	(*Record)(nil),                // 1: Record
	(*StarInfoResp)(nil),          // 2: StarInfoResp
	(*EngageRecommendReq)(nil),    // 3: EngageRecommendReq
	(*EngageRecommendPlayer)(nil), // 4: EngageRecommendPlayer
	(*EngageRecommendList)(nil),   // 5: EngageRecommendList
	(WinState)(0),                 // 6: WinState
}
var file_star_proto_depIdxs = []int32{
	6, // 0: Record.winState:type_name -> WinState
	0, // 1: StarInfoResp.nowRewardType:type_name -> StarRewardType
	1, // 2: StarInfoResp.record:type_name -> Record
	4, // 3: EngageRecommendList.e:type_name -> EngageRecommendPlayer
	4, // 4: EngageRecommendList.f:type_name -> EngageRecommendPlayer
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_star_proto_init() }
func file_star_proto_init() {
	if File_star_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_star_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_star_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StarInfoResp); i {
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
		file_star_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EngageRecommendReq); i {
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
		file_star_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EngageRecommendPlayer); i {
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
		file_star_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EngageRecommendList); i {
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
			RawDescriptor: file_star_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_star_proto_goTypes,
		DependencyIndexes: file_star_proto_depIdxs,
		EnumInfos:         file_star_proto_enumTypes,
		MessageInfos:      file_star_proto_msgTypes,
	}.Build()
	File_star_proto = out.File
	file_star_proto_rawDesc = nil
	file_star_proto_goTypes = nil
	file_star_proto_depIdxs = nil
}
