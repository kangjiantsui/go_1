// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: king_round.proto

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

// 路人王信息
type KingRoundInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 全局数据
	WeekPeriod int32           `protobuf:"varint,1,opt,name=weekPeriod,proto3" json:"weekPeriod,omitempty"`                                                                                  // 周期数   unique week + 1/2/3/4
	EndTime    int64           `protobuf:"varint,2,opt,name=endTime,proto3" json:"endTime,omitempty"`                                                                                        // 当前周期结束倒计时  秒
	PlayId     map[int32]int32 `protobuf:"bytes,3,rep,name=playId,proto3" json:"playId,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 当前周期内随机出来的玩法， 轮次：玩法
	// 历史数据
	RWeekPeriod int32 `protobuf:"varint,4,opt,name=rWeekPeriod,proto3" json:"rWeekPeriod,omitempty"` // 上次领奖卡片周期数
	SWeekPeriod int32 `protobuf:"varint,5,opt,name=sWeekPeriod,proto3" json:"sWeekPeriod,omitempty"` // 当前数据所处于的周期数
	// 周期数据  每次获取数据时根据周期数刷新
	TotalWin  int32 `protobuf:"varint,6,opt,name=totalWin,proto3" json:"totalWin,omitempty"`   // 当前周期胜利总次数
	TotalLose int32 `protobuf:"varint,7,opt,name=totalLose,proto3" json:"totalLose,omitempty"` // 当前周期失败总次数
	CurRound  int32 `protobuf:"varint,8,opt,name=curRound,proto3" json:"curRound,omitempty"`   // 当前周期内的报名轮次   0为未报名  1-3为已报名的轮次
	// 轮次数据  报名才会重置，报名的条件是奖励全部领取
	Section           int32          `protobuf:"varint,9,opt,name=section,proto3" json:"section,omitempty"`                                                                                         // 当前轮次已通关的站点
	ActId             int32          `protobuf:"varint,10,opt,name=actId,proto3" json:"actId,omitempty"`                                                                                            // 当前轮次活动ID
	Lose              int32          `protobuf:"varint,11,opt,name=lose,proto3" json:"lose,omitempty"`                                                                                              // 当前轮次失败次数
	Reward            map[int32]bool `protobuf:"bytes,12,rep,name=reward,proto3" json:"reward,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 当前轮次奖励领取记录  section:true
	MaxWinStreak      int32          `protobuf:"varint,13,opt,name=maxWinStreak,proto3" json:"maxWinStreak,omitempty"`                                                                              // 最大连胜数
	TenWinStreakCount int32          `protobuf:"varint,14,opt,name=tenWinStreakCount,proto3" json:"tenWinStreakCount,omitempty"`                                                                    // 十连胜次数
}

func (x *KingRoundInfo) Reset() {
	*x = KingRoundInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_king_round_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KingRoundInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KingRoundInfo) ProtoMessage() {}

func (x *KingRoundInfo) ProtoReflect() protoreflect.Message {
	mi := &file_king_round_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KingRoundInfo.ProtoReflect.Descriptor instead.
func (*KingRoundInfo) Descriptor() ([]byte, []int) {
	return file_king_round_proto_rawDescGZIP(), []int{0}
}

func (x *KingRoundInfo) GetWeekPeriod() int32 {
	if x != nil {
		return x.WeekPeriod
	}
	return 0
}

func (x *KingRoundInfo) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *KingRoundInfo) GetPlayId() map[int32]int32 {
	if x != nil {
		return x.PlayId
	}
	return nil
}

func (x *KingRoundInfo) GetRWeekPeriod() int32 {
	if x != nil {
		return x.RWeekPeriod
	}
	return 0
}

func (x *KingRoundInfo) GetSWeekPeriod() int32 {
	if x != nil {
		return x.SWeekPeriod
	}
	return 0
}

func (x *KingRoundInfo) GetTotalWin() int32 {
	if x != nil {
		return x.TotalWin
	}
	return 0
}

func (x *KingRoundInfo) GetTotalLose() int32 {
	if x != nil {
		return x.TotalLose
	}
	return 0
}

func (x *KingRoundInfo) GetCurRound() int32 {
	if x != nil {
		return x.CurRound
	}
	return 0
}

func (x *KingRoundInfo) GetSection() int32 {
	if x != nil {
		return x.Section
	}
	return 0
}

func (x *KingRoundInfo) GetActId() int32 {
	if x != nil {
		return x.ActId
	}
	return 0
}

func (x *KingRoundInfo) GetLose() int32 {
	if x != nil {
		return x.Lose
	}
	return 0
}

func (x *KingRoundInfo) GetReward() map[int32]bool {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *KingRoundInfo) GetMaxWinStreak() int32 {
	if x != nil {
		return x.MaxWinStreak
	}
	return 0
}

func (x *KingRoundInfo) GetTenWinStreakCount() int32 {
	if x != nil {
		return x.TenWinStreakCount
	}
	return 0
}

// 操作后路人王信息
type OptKRInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *KingRoundInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"` // 操作后路人王信息
	Ir   *ItemUpdateResult `protobuf:"bytes,2,opt,name=ir,proto3" json:"ir,omitempty"`     // 资源变化
}

func (x *OptKRInfo) Reset() {
	*x = OptKRInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_king_round_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptKRInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptKRInfo) ProtoMessage() {}

func (x *OptKRInfo) ProtoReflect() protoreflect.Message {
	mi := &file_king_round_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptKRInfo.ProtoReflect.Descriptor instead.
func (*OptKRInfo) Descriptor() ([]byte, []int) {
	return file_king_round_proto_rawDescGZIP(), []int{1}
}

func (x *OptKRInfo) GetInfo() *KingRoundInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *OptKRInfo) GetIr() *ItemUpdateResult {
	if x != nil {
		return x.Ir
	}
	return nil
}

// 领取站点通关奖励
type TakeKingRoundReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TakeAll   bool  `protobuf:"varint,1,opt,name=takeAll,proto3" json:"takeAll,omitempty"`     // 是否全部领取
	SectionId int32 `protobuf:"varint,2,opt,name=sectionId,proto3" json:"sectionId,omitempty"` // 站点id 不是全部领取时有效
	Actor     int32 `protobuf:"varint,3,opt,name=actor,proto3" json:"actor,omitempty"`         // 熟练度指定球员id
}

func (x *TakeKingRoundReward) Reset() {
	*x = TakeKingRoundReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_king_round_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeKingRoundReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeKingRoundReward) ProtoMessage() {}

func (x *TakeKingRoundReward) ProtoReflect() protoreflect.Message {
	mi := &file_king_round_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeKingRoundReward.ProtoReflect.Descriptor instead.
func (*TakeKingRoundReward) Descriptor() ([]byte, []int) {
	return file_king_round_proto_rawDescGZIP(), []int{2}
}

func (x *TakeKingRoundReward) GetTakeAll() bool {
	if x != nil {
		return x.TakeAll
	}
	return false
}

func (x *TakeKingRoundReward) GetSectionId() int32 {
	if x != nil {
		return x.SectionId
	}
	return 0
}

func (x *TakeKingRoundReward) GetActor() int32 {
	if x != nil {
		return x.Actor
	}
	return 0
}

var File_king_round_proto protoreflect.FileDescriptor

var file_king_round_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7,
	0x04, 0x0a, 0x0d, 0x4b, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x65, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x65, 0x65, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x4b, 0x69, 0x6e,
	0x67, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x49,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x57, 0x65, 0x65, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x57, 0x65, 0x65, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x57, 0x65, 0x65, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x57, 0x65, 0x65, 0x6b, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x69, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x69, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x6f, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x6f, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x73, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x4b,
	0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x57, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6b,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x57, 0x69, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6b, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x65, 0x6e, 0x57, 0x69, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x74, 0x65, 0x6e, 0x57, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a,
	0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52, 0x0a, 0x09, 0x4f, 0x70, 0x74, 0x4b,
	0x52, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x4b, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x02, 0x69, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x02, 0x69, 0x72, 0x22, 0x63, 0x0a, 0x13,
	0x54, 0x61, 0x6b, 0x65, 0x4b, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x6b, 0x65, 0x41, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x61, 0x6b, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x42, 0x0b, 0x5a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_king_round_proto_rawDescOnce sync.Once
	file_king_round_proto_rawDescData = file_king_round_proto_rawDesc
)

func file_king_round_proto_rawDescGZIP() []byte {
	file_king_round_proto_rawDescOnce.Do(func() {
		file_king_round_proto_rawDescData = protoimpl.X.CompressGZIP(file_king_round_proto_rawDescData)
	})
	return file_king_round_proto_rawDescData
}

var file_king_round_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_king_round_proto_goTypes = []interface{}{
	(*KingRoundInfo)(nil),       // 0: KingRoundInfo
	(*OptKRInfo)(nil),           // 1: OptKRInfo
	(*TakeKingRoundReward)(nil), // 2: TakeKingRoundReward
	nil,                         // 3: KingRoundInfo.PlayIdEntry
	nil,                         // 4: KingRoundInfo.RewardEntry
	(*ItemUpdateResult)(nil),    // 5: ItemUpdateResult
}
var file_king_round_proto_depIdxs = []int32{
	3, // 0: KingRoundInfo.playId:type_name -> KingRoundInfo.PlayIdEntry
	4, // 1: KingRoundInfo.reward:type_name -> KingRoundInfo.RewardEntry
	0, // 2: OptKRInfo.info:type_name -> KingRoundInfo
	5, // 3: OptKRInfo.ir:type_name -> ItemUpdateResult
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_king_round_proto_init() }
func file_king_round_proto_init() {
	if File_king_round_proto != nil {
		return
	}
	file_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_king_round_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KingRoundInfo); i {
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
		file_king_round_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptKRInfo); i {
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
		file_king_round_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeKingRoundReward); i {
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
			RawDescriptor: file_king_round_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_king_round_proto_goTypes,
		DependencyIndexes: file_king_round_proto_depIdxs,
		MessageInfos:      file_king_round_proto_msgTypes,
	}.Build()
	File_king_round_proto = out.File
	file_king_round_proto_rawDesc = nil
	file_king_round_proto_goTypes = nil
	file_king_round_proto_depIdxs = nil
}
