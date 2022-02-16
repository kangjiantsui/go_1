// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: item.proto

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

// 奖励后非道具变化后值
type UpdateAfter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorMaxCup            []*ItemPair     `protobuf:"bytes,1,rep,name=actorMaxCup,proto3" json:"actorMaxCup,omitempty"`                                                                                      // 球员最大奖杯数更新  id:num
	PlayerMaxCup           int32           `protobuf:"varint,2,opt,name=playerMaxCup,proto3" json:"playerMaxCup,omitempty"`                                                                                   // 玩家最大杯数
	PlayerLevel            int32           `protobuf:"varint,3,opt,name=playerLevel,proto3" json:"playerLevel,omitempty"`                                                                                     // 玩家等级
	MedalLimit             int32           `protobuf:"varint,4,opt,name=medalLimit,proto3" json:"medalLimit,omitempty"`                                                                                       // 奖章限定值
	RecoverFrom            int64           `protobuf:"varint,5,opt,name=recoverFrom,proto3" json:"recoverFrom,omitempty"`                                                                                     // 奖章恢复倒计时
	BattlepassLevel        int32           `protobuf:"varint,6,opt,name=battlepassLevel,proto3" json:"battlepassLevel,omitempty"`                                                                             //赛券等级
	Medal                  int32           `protobuf:"varint,7,opt,name=medal,proto3" json:"medal,omitempty"`                                                                                                 //奖章数
	NowPoint               int32           `protobuf:"varint,8,opt,name=nowPoint,proto3" json:"nowPoint,omitempty"`                                                                                           // 当前赛点
	NowMode                int32           `protobuf:"varint,9,opt,name=nowMode,proto3" json:"nowMode,omitempty"`                                                                                             // 当前赛点玩法
	Actor                  *Actor          `protobuf:"bytes,10,opt,name=actor,proto3" json:"actor,omitempty"`                                                                                                 // 当前球员信息
	BoxCount               map[int32]int32 `protobuf:"bytes,11,rep,name=boxCount,proto3" json:"boxCount,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 宝箱开启计数   球员稀有度：已开启宝箱个数
	Star                   int32           `protobuf:"varint,12,opt,name=star,proto3" json:"star,omitempty"`                                                                                                  // 变化后的星数
	StarLevel              int32           `protobuf:"varint,13,opt,name=starLevel,proto3" json:"starLevel,omitempty"`                                                                                        // 变化后星级
	KrSection              int32           `protobuf:"varint,14,opt,name=krSection,proto3" json:"krSection,omitempty"`                                                                                        // 路人王结算后的站点
	KrLose                 int32           `protobuf:"varint,15,opt,name=krLose,proto3" json:"krLose,omitempty"`                                                                                              // 路人王结算后的失败次数
	AchievementLevel       int32           `protobuf:"varint,16,opt,name=achievementLevel,proto3" json:"achievementLevel,omitempty"`                                                                          // 成就等级
	AchievementLevelReward []*ItemUpdate   `protobuf:"bytes,17,rep,name=achievementLevelReward,proto3" json:"achievementLevelReward,omitempty"`                                                               // 成就等级奖励
}

func (x *UpdateAfter) Reset() {
	*x = UpdateAfter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAfter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAfter) ProtoMessage() {}

func (x *UpdateAfter) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAfter.ProtoReflect.Descriptor instead.
func (*UpdateAfter) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAfter) GetActorMaxCup() []*ItemPair {
	if x != nil {
		return x.ActorMaxCup
	}
	return nil
}

func (x *UpdateAfter) GetPlayerMaxCup() int32 {
	if x != nil {
		return x.PlayerMaxCup
	}
	return 0
}

func (x *UpdateAfter) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *UpdateAfter) GetMedalLimit() int32 {
	if x != nil {
		return x.MedalLimit
	}
	return 0
}

func (x *UpdateAfter) GetRecoverFrom() int64 {
	if x != nil {
		return x.RecoverFrom
	}
	return 0
}

func (x *UpdateAfter) GetBattlepassLevel() int32 {
	if x != nil {
		return x.BattlepassLevel
	}
	return 0
}

func (x *UpdateAfter) GetMedal() int32 {
	if x != nil {
		return x.Medal
	}
	return 0
}

func (x *UpdateAfter) GetNowPoint() int32 {
	if x != nil {
		return x.NowPoint
	}
	return 0
}

func (x *UpdateAfter) GetNowMode() int32 {
	if x != nil {
		return x.NowMode
	}
	return 0
}

func (x *UpdateAfter) GetActor() *Actor {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *UpdateAfter) GetBoxCount() map[int32]int32 {
	if x != nil {
		return x.BoxCount
	}
	return nil
}

func (x *UpdateAfter) GetStar() int32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *UpdateAfter) GetStarLevel() int32 {
	if x != nil {
		return x.StarLevel
	}
	return 0
}

func (x *UpdateAfter) GetKrSection() int32 {
	if x != nil {
		return x.KrSection
	}
	return 0
}

func (x *UpdateAfter) GetKrLose() int32 {
	if x != nil {
		return x.KrLose
	}
	return 0
}

func (x *UpdateAfter) GetAchievementLevel() int32 {
	if x != nil {
		return x.AchievementLevel
	}
	return 0
}

func (x *UpdateAfter) GetAchievementLevelReward() []*ItemUpdate {
	if x != nil {
		return x.AchievementLevelReward
	}
	return nil
}

// 道具变更 （奖励返回）
type ItemUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        int32         `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`                 // `required` 道具类型
	Id          uint32        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`                     // 道具id
	Num         int32         `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`                   // 道具数量
	Total       int32         `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`               // 修改后的值
	Belong      uint32        `protobuf:"varint,5,opt,name=belong,proto3" json:"belong,omitempty"`             // 对于球员战力能量，该值为球员id
	IsHide      int32         `protobuf:"varint,6,opt,name=isHide,proto3" json:"isHide,omitempty"`             // 是否是隐藏奖励 1 为额外奖励
	Box         []*ItemUpdate `protobuf:"bytes,7,rep,name=box,proto3" json:"box,omitempty"`                    // 宝箱道具  如果这个道具奖励
	Source      SOURCE        `protobuf:"varint,8,opt,name=source,proto3,enum=SOURCE" json:"source,omitempty"` // 奖励来源
	IsConverted int32         `protobuf:"varint,9,opt,name=isConverted,proto3" json:"isConverted,omitempty"`   // 已经转化过，不实际增加，也不显示
}

func (x *ItemUpdate) Reset() {
	*x = ItemUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemUpdate) ProtoMessage() {}

func (x *ItemUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemUpdate.ProtoReflect.Descriptor instead.
func (*ItemUpdate) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{1}
}

func (x *ItemUpdate) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ItemUpdate) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemUpdate) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *ItemUpdate) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ItemUpdate) GetBelong() uint32 {
	if x != nil {
		return x.Belong
	}
	return 0
}

func (x *ItemUpdate) GetIsHide() int32 {
	if x != nil {
		return x.IsHide
	}
	return 0
}

func (x *ItemUpdate) GetBox() []*ItemUpdate {
	if x != nil {
		return x.Box
	}
	return nil
}

func (x *ItemUpdate) GetSource() SOURCE {
	if x != nil {
		return x.Source
	}
	return SOURCE_S_NONE
}

func (x *ItemUpdate) GetIsConverted() int32 {
	if x != nil {
		return x.IsConverted
	}
	return 0
}

// 道具变更列表 （奖励与扣除）
type ItemUpdateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList    []*ItemUpdate `protobuf:"bytes,1,rep,name=itemList,proto3" json:"itemList,omitempty"`       // 奖励列表 道具变化
	UpdateAfter *UpdateAfter  `protobuf:"bytes,2,opt,name=updateAfter,proto3" json:"updateAfter,omitempty"` // 非道具变化后值
}

func (x *ItemUpdateResult) Reset() {
	*x = ItemUpdateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemUpdateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemUpdateResult) ProtoMessage() {}

func (x *ItemUpdateResult) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemUpdateResult.ProtoReflect.Descriptor instead.
func (*ItemUpdateResult) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{2}
}

func (x *ItemUpdateResult) GetItemList() []*ItemUpdate {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *ItemUpdateResult) GetUpdateAfter() *UpdateAfter {
	if x != nil {
		return x.UpdateAfter
	}
	return nil
}

type RewardResultRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reward          *ItemUpdateResult  `protobuf:"bytes,1,opt,name=reward,proto3" json:"reward,omitempty"`                    // 结算奖励
	Mission         *MissionUpdateResp `protobuf:"bytes,2,opt,name=mission,proto3" json:"mission,omitempty"`                  // 任务更新
	CoinTimeGoldNum int32              `protobuf:"varint,3,opt,name=CoinTimeGoldNum,proto3" json:"CoinTimeGoldNum,omitempty"` // 千金一刻奖励金币数量
}

func (x *RewardResultRes) Reset() {
	*x = RewardResultRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardResultRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardResultRes) ProtoMessage() {}

func (x *RewardResultRes) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardResultRes.ProtoReflect.Descriptor instead.
func (*RewardResultRes) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{3}
}

func (x *RewardResultRes) GetReward() *ItemUpdateResult {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *RewardResultRes) GetMission() *MissionUpdateResp {
	if x != nil {
		return x.Mission
	}
	return nil
}

func (x *RewardResultRes) GetCoinTimeGoldNum() int32 {
	if x != nil {
		return x.CoinTimeGoldNum
	}
	return 0
}

type ItemRecordOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64         `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ItemList  []*ItemUpdate `protobuf:"bytes,2,rep,name=itemList,proto3" json:"itemList,omitempty"`
}

func (x *ItemRecordOne) Reset() {
	*x = ItemRecordOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemRecordOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemRecordOne) ProtoMessage() {}

func (x *ItemRecordOne) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemRecordOne.ProtoReflect.Descriptor instead.
func (*ItemRecordOne) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{4}
}

func (x *ItemRecordOne) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ItemRecordOne) GetItemList() []*ItemUpdate {
	if x != nil {
		return x.ItemList
	}
	return nil
}

type ItemRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record []*ItemRecordOne `protobuf:"bytes,1,rep,name=record,proto3" json:"record,omitempty"`
}

func (x *ItemRecord) Reset() {
	*x = ItemRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemRecord) ProtoMessage() {}

func (x *ItemRecord) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemRecord.ProtoReflect.Descriptor instead.
func (*ItemRecord) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{5}
}

func (x *ItemRecord) GetRecord() []*ItemRecordOne {
	if x != nil {
		return x.Record
	}
	return nil
}

var File_item_proto protoreflect.FileDescriptor

var file_item_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa4, 0x05, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x66, 0x74,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x78, 0x43, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x78, 0x43, 0x75, 0x70, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x78, 0x43, 0x75, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x78,
	0x43, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x65, 0x64, 0x61, 0x6c,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x70, 0x61, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x70, 0x61, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x77, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x6f, 0x77, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x62,
	0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x78, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x6f, 0x78, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x72, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x72, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x72, 0x4c, 0x6f, 0x73, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6b, 0x72, 0x4c, 0x6f, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x43, 0x0a, 0x16, 0x61, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x16, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x1a, 0x3b, 0x0a, 0x0d,
	0x42, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xea, 0x01, 0x0a, 0x0a, 0x49, 0x74,
	0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x73, 0x48, 0x69, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73,
	0x48, 0x69, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x62, 0x6f, 0x78, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x03,
	0x62, 0x6f, 0x78, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x22, 0x6b, 0x0a, 0x10, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x08, 0x69, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x66, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x22, 0x94, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x07, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x47, 0x6f, 0x6c, 0x64,
	0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x47, 0x6f, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x22, 0x56, 0x0a, 0x0d, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x34, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x6e, 0x65,
	0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x0b, 0x5a, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_item_proto_rawDescOnce sync.Once
	file_item_proto_rawDescData = file_item_proto_rawDesc
)

func file_item_proto_rawDescGZIP() []byte {
	file_item_proto_rawDescOnce.Do(func() {
		file_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_proto_rawDescData)
	})
	return file_item_proto_rawDescData
}

var file_item_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_item_proto_goTypes = []interface{}{
	(*UpdateAfter)(nil),       // 0: UpdateAfter
	(*ItemUpdate)(nil),        // 1: ItemUpdate
	(*ItemUpdateResult)(nil),  // 2: ItemUpdateResult
	(*RewardResultRes)(nil),   // 3: RewardResultRes
	(*ItemRecordOne)(nil),     // 4: ItemRecordOne
	(*ItemRecord)(nil),        // 5: ItemRecord
	nil,                       // 6: UpdateAfter.BoxCountEntry
	(*ItemPair)(nil),          // 7: ItemPair
	(*Actor)(nil),             // 8: Actor
	(SOURCE)(0),               // 9: SOURCE
	(*MissionUpdateResp)(nil), // 10: MissionUpdateResp
}
var file_item_proto_depIdxs = []int32{
	7,  // 0: UpdateAfter.actorMaxCup:type_name -> ItemPair
	8,  // 1: UpdateAfter.actor:type_name -> Actor
	6,  // 2: UpdateAfter.boxCount:type_name -> UpdateAfter.BoxCountEntry
	1,  // 3: UpdateAfter.achievementLevelReward:type_name -> ItemUpdate
	1,  // 4: ItemUpdate.box:type_name -> ItemUpdate
	9,  // 5: ItemUpdate.source:type_name -> SOURCE
	1,  // 6: ItemUpdateResult.itemList:type_name -> ItemUpdate
	0,  // 7: ItemUpdateResult.updateAfter:type_name -> UpdateAfter
	2,  // 8: RewardResultRes.reward:type_name -> ItemUpdateResult
	10, // 9: RewardResultRes.mission:type_name -> MissionUpdateResp
	1,  // 10: ItemRecordOne.itemList:type_name -> ItemUpdate
	4,  // 11: ItemRecord.record:type_name -> ItemRecordOne
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_item_proto_init() }
func file_item_proto_init() {
	if File_item_proto != nil {
		return
	}
	file_game_proto_init()
	file_define_proto_init()
	file_mission_proto_init()
	file_actor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAfter); i {
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
		file_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemUpdate); i {
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
		file_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemUpdateResult); i {
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
		file_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardResultRes); i {
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
		file_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemRecordOne); i {
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
		file_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemRecord); i {
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
			RawDescriptor: file_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_item_proto_goTypes,
		DependencyIndexes: file_item_proto_depIdxs,
		MessageInfos:      file_item_proto_msgTypes,
	}.Build()
	File_item_proto = out.File
	file_item_proto_rawDesc = nil
	file_item_proto_goTypes = nil
	file_item_proto_depIdxs = nil
}
