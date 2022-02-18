// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: actor.proto

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

type ActorStatusEnum int32

const (
	ActorStatusEnum_ASE_None      ActorStatusEnum = 0 // 不展示
	ActorStatusEnum_ASE_Not_STart ActorStatusEnum = 1 // 未开始
	ActorStatusEnum_ASE_Review    ActorStatusEnum = 2 // 预览
	ActorStatusEnum_ASE_Output    ActorStatusEnum = 3 // 产出
	ActorStatusEnum_ASE_Off       ActorStatusEnum = 4 // 关闭（指上新活动不再开启，但仍能购买）
	ActorStatusEnum_ASE_End       ActorStatusEnum = 5 // 停止产出
)

// Enum value maps for ActorStatusEnum.
var (
	ActorStatusEnum_name = map[int32]string{
		0: "ASE_None",
		1: "ASE_Not_STart",
		2: "ASE_Review",
		3: "ASE_Output",
		4: "ASE_Off",
		5: "ASE_End",
	}
	ActorStatusEnum_value = map[string]int32{
		"ASE_None":      0,
		"ASE_Not_STart": 1,
		"ASE_Review":    2,
		"ASE_Output":    3,
		"ASE_Off":       4,
		"ASE_End":       5,
	}
)

func (x ActorStatusEnum) Enum() *ActorStatusEnum {
	p := new(ActorStatusEnum)
	*p = x
	return p
}

func (x ActorStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActorStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_actor_proto_enumTypes[0].Descriptor()
}

func (ActorStatusEnum) Type() protoreflect.EnumType {
	return &file_actor_proto_enumTypes[0]
}

func (x ActorStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActorStatusEnum.Descriptor instead.
func (ActorStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{0}
}

type Actor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId     uint32          `protobuf:"varint,1,opt,name=actorId,proto3" json:"actorId,omitempty"`                                                                                                   // 球员id（道具id）
	Level       int32           `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`                                                                                                       // 当前等级
	Cup         int32           `protobuf:"varint,3,opt,name=cup,proto3" json:"cup,omitempty"`                                                                                                           // 该球员总杯数
	Proficiency int32           `protobuf:"varint,4,opt,name=proficiency,proto3" json:"proficiency,omitempty"`                                                                                           // 剩余熟练度
	MaxCup      int32           `protobuf:"varint,5,opt,name=maxCup,proto3" json:"maxCup,omitempty"`                                                                                                     // 该球员历史最大杯数(用于等级计算)
	CreateAt    int64           `protobuf:"varint,6,opt,name=createAt,proto3" json:"createAt,omitempty"`                                                                                                 // 球员创建时间
	EquipSkill  []int32         `protobuf:"varint,7,rep,packed,name=equipSkill,proto3" json:"equipSkill,omitempty"`                                                                                      // 装备的技能
	HaveSkill   []int32         `protobuf:"varint,8,rep,packed,name=haveSkill,proto3" json:"haveSkill,omitempty"`                                                                                        // 已解锁的技能
	Skin        int32           `protobuf:"varint,9,opt,name=skin,proto3" json:"skin,omitempty"`                                                                                                         // 皮肤
	Collocation map[int32]int32 `protobuf:"bytes,10,rep,name=collocation,proto3" json:"collocation,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 部位，服饰id
}

func (x *Actor) Reset() {
	*x = Actor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actor) ProtoMessage() {}

func (x *Actor) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actor.ProtoReflect.Descriptor instead.
func (*Actor) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{0}
}

func (x *Actor) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *Actor) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Actor) GetCup() int32 {
	if x != nil {
		return x.Cup
	}
	return 0
}

func (x *Actor) GetProficiency() int32 {
	if x != nil {
		return x.Proficiency
	}
	return 0
}

func (x *Actor) GetMaxCup() int32 {
	if x != nil {
		return x.MaxCup
	}
	return 0
}

func (x *Actor) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Actor) GetEquipSkill() []int32 {
	if x != nil {
		return x.EquipSkill
	}
	return nil
}

func (x *Actor) GetHaveSkill() []int32 {
	if x != nil {
		return x.HaveSkill
	}
	return nil
}

func (x *Actor) GetSkin() int32 {
	if x != nil {
		return x.Skin
	}
	return 0
}

func (x *Actor) GetCollocation() map[int32]int32 {
	if x != nil {
		return x.Collocation
	}
	return nil
}

type ActorOutPutTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeShow   int64 `protobuf:"varint,1,opt,name=timeShow,proto3" json:"timeShow,omitempty"`     // 倒计时 开始显示时间
	TimeOutPut int64 `protobuf:"varint,2,opt,name=timeOutPut,proto3" json:"timeOutPut,omitempty"` // 倒计时 开始产出时间
	TimeUnable int64 `protobuf:"varint,3,opt,name=timeUnable,proto3" json:"timeUnable,omitempty"` // 倒计时 停止产出时间
}

func (x *ActorOutPutTime) Reset() {
	*x = ActorOutPutTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorOutPutTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorOutPutTime) ProtoMessage() {}

func (x *ActorOutPutTime) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorOutPutTime.ProtoReflect.Descriptor instead.
func (*ActorOutPutTime) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{1}
}

func (x *ActorOutPutTime) GetTimeShow() int64 {
	if x != nil {
		return x.TimeShow
	}
	return 0
}

func (x *ActorOutPutTime) GetTimeOutPut() int64 {
	if x != nil {
		return x.TimeOutPut
	}
	return 0
}

func (x *ActorOutPutTime) GetTimeUnable() int64 {
	if x != nil {
		return x.TimeUnable
	}
	return 0
}

type ActorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actor    map[uint32]*Actor `protobuf:"bytes,1,rep,name=actor,proto3" json:"actor,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`        // 球员列表 基础存储数据  actorId: Actor
	Group    []*ItemPair       `protobuf:"bytes,2,rep,name=group,proto3" json:"group,omitempty"`                                                                                                 // 3v3上阵列表  球员id，位置 ; 球员id，位置 ; 球员id，位置
	Show     uint32            `protobuf:"varint,3,opt,name=show,proto3" json:"show,omitempty"`                                                                                                  // 显示的球员（默认的球员）
	BoxCount map[int32]int32   `protobuf:"bytes,4,rep,name=boxCount,proto3" json:"boxCount,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 宝箱开启计数   球员稀有度：已开启宝箱个数
}

func (x *ActorInfo) Reset() {
	*x = ActorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorInfo) ProtoMessage() {}

func (x *ActorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorInfo.ProtoReflect.Descriptor instead.
func (*ActorInfo) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{2}
}

func (x *ActorInfo) GetActor() map[uint32]*Actor {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *ActorInfo) GetGroup() []*ItemPair {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *ActorInfo) GetShow() uint32 {
	if x != nil {
		return x.Show
	}
	return 0
}

func (x *ActorInfo) GetBoxCount() map[int32]int32 {
	if x != nil {
		return x.BoxCount
	}
	return nil
}

type ActorOutPutStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status map[uint32]*ActorOutPutTime `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 球员产出状态
}

func (x *ActorOutPutStatus) Reset() {
	*x = ActorOutPutStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorOutPutStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorOutPutStatus) ProtoMessage() {}

func (x *ActorOutPutStatus) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorOutPutStatus.ProtoReflect.Descriptor instead.
func (*ActorOutPutStatus) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{3}
}

func (x *ActorOutPutStatus) GetStatus() map[uint32]*ActorOutPutTime {
	if x != nil {
		return x.Status
	}
	return nil
}

type ActorGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group []*ItemPair `protobuf:"bytes,1,rep,name=group,proto3" json:"group,omitempty"` //`required` 球员组
}

func (x *ActorGroupReq) Reset() {
	*x = ActorGroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorGroupReq) ProtoMessage() {}

func (x *ActorGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorGroupReq.ProtoReflect.Descriptor instead.
func (*ActorGroupReq) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{4}
}

func (x *ActorGroupReq) GetGroup() []*ItemPair {
	if x != nil {
		return x.Group
	}
	return nil
}

type ActorShowReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Show uint32 `protobuf:"varint,1,opt,name=show,proto3" json:"show,omitempty"` //`required` 显示的球员ID
}

func (x *ActorShowReq) Reset() {
	*x = ActorShowReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorShowReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorShowReq) ProtoMessage() {}

func (x *ActorShowReq) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorShowReq.ProtoReflect.Descriptor instead.
func (*ActorShowReq) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{5}
}

func (x *ActorShowReq) GetShow() uint32 {
	if x != nil {
		return x.Show
	}
	return 0
}

// 请求球员升级
type ActorUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId uint32 `protobuf:"varint,1,opt,name=actorId,proto3" json:"actorId,omitempty"` // `required` 球员id（道具id）
}

func (x *ActorUpdateReq) Reset() {
	*x = ActorUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorUpdateReq) ProtoMessage() {}

func (x *ActorUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorUpdateReq.ProtoReflect.Descriptor instead.
func (*ActorUpdateReq) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{6}
}

func (x *ActorUpdateReq) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

// 装备技能
type ActorSkillReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId int32 `protobuf:"varint,1,opt,name=skillId,proto3" json:"skillId,omitempty"`
}

func (x *ActorSkillReq) Reset() {
	*x = ActorSkillReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorSkillReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorSkillReq) ProtoMessage() {}

func (x *ActorSkillReq) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorSkillReq.ProtoReflect.Descriptor instead.
func (*ActorSkillReq) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{7}
}

func (x *ActorSkillReq) GetSkillId() int32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

// 获取球员全部技能,0表示全部球员
type GetAllSkillReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId uint32 `protobuf:"varint,1,opt,name=actorId,proto3" json:"actorId,omitempty"`
	Ic      uint64 `protobuf:"varint,2,opt,name=ic,proto3" json:"ic,omitempty"`
}

func (x *GetAllSkillReq) Reset() {
	*x = GetAllSkillReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSkillReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSkillReq) ProtoMessage() {}

func (x *GetAllSkillReq) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSkillReq.ProtoReflect.Descriptor instead.
func (*GetAllSkillReq) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllSkillReq) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *GetAllSkillReq) GetIc() uint64 {
	if x != nil {
		return x.Ic
	}
	return 0
}

// 请求一个稀有度所有能获得的球员
type ActorsAvailableReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rarity int32 `protobuf:"varint,1,opt,name=rarity,proto3" json:"rarity,omitempty"` // 稀有度   如果为0则返回所有可获得的球员
}

func (x *ActorsAvailableReq) Reset() {
	*x = ActorsAvailableReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorsAvailableReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorsAvailableReq) ProtoMessage() {}

func (x *ActorsAvailableReq) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorsAvailableReq.ProtoReflect.Descriptor instead.
func (*ActorsAvailableReq) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{9}
}

func (x *ActorsAvailableReq) GetRarity() int32 {
	if x != nil {
		return x.Rarity
	}
	return 0
}

// 返回稀有度/所有 可获得的球员
type ActorsAvailableResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId []int32 `protobuf:"varint,1,rep,packed,name=actorId,proto3" json:"actorId,omitempty"` // 可获得的球员id列表
}

func (x *ActorsAvailableResp) Reset() {
	*x = ActorsAvailableResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorsAvailableResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorsAvailableResp) ProtoMessage() {}

func (x *ActorsAvailableResp) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorsAvailableResp.ProtoReflect.Descriptor instead.
func (*ActorsAvailableResp) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{10}
}

func (x *ActorsAvailableResp) GetActorId() []int32 {
	if x != nil {
		return x.ActorId
	}
	return nil
}

type ActorStatusAndTimeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`   // 产出状态
	EndTime int64 `protobuf:"varint,2,opt,name=endTime,proto3" json:"endTime,omitempty"` // 结束倒计时
	BuyType int32 `protobuf:"varint,3,opt,name=buyType,proto3" json:"buyType,omitempty"` // 产出途径
}

func (x *ActorStatusAndTimeItem) Reset() {
	*x = ActorStatusAndTimeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorStatusAndTimeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorStatusAndTimeItem) ProtoMessage() {}

func (x *ActorStatusAndTimeItem) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorStatusAndTimeItem.ProtoReflect.Descriptor instead.
func (*ActorStatusAndTimeItem) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{11}
}

func (x *ActorStatusAndTimeItem) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ActorStatusAndTimeItem) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ActorStatusAndTimeItem) GetBuyType() int32 {
	if x != nil {
		return x.BuyType
	}
	return 0
}

// 返回球员产出途径与途径结束时间
type ActorStatusAndTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actors map[int32]*ActorStatusAndTimeItem `protobuf:"bytes,1,rep,name=actors,proto3" json:"actors,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // id:ActorStatusAndTimeItem
}

func (x *ActorStatusAndTime) Reset() {
	*x = ActorStatusAndTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorStatusAndTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorStatusAndTime) ProtoMessage() {}

func (x *ActorStatusAndTime) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorStatusAndTime.ProtoReflect.Descriptor instead.
func (*ActorStatusAndTime) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{12}
}

func (x *ActorStatusAndTime) GetActors() map[int32]*ActorStatusAndTimeItem {
	if x != nil {
		return x.Actors
	}
	return nil
}

var File_actor_proto protoreflect.FileDescriptor

var file_actor_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a, 0x05, 0x41, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x63, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x43, 0x75,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x43, 0x75, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0a, 0x65, 0x71, 0x75, 0x69, 0x70, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x68,
	0x61, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09,
	0x68, 0x61, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x6e, 0x12, 0x39, 0x0a,
	0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3e, 0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6d, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x6f,
	0x72, 0x4f, 0x75, 0x74, 0x50, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x68, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x4f,
	0x75, 0x74, 0x50, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x4f, 0x75, 0x74, 0x50, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x55,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x55, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xa2, 0x02, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x77, 0x12, 0x34, 0x0a, 0x08, 0x62, 0x6f, 0x78, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x40, 0x0a,
	0x0a, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3b, 0x0a, 0x0d, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x98, 0x01, 0x0a,
	0x11, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x50, 0x75, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x50, 0x75, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x4b, 0x0a, 0x0b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x4f, 0x75, 0x74, 0x50, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f,
	0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x77, 0x22, 0x2a, 0x0a,
	0x0e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x63,
	0x22, 0x2c, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x22, 0x2f,
	0x0a, 0x13, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22,
	0x64, 0x0a, 0x16, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x75, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x75,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x52, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x41, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x6c, 0x0a, 0x0f, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x53, 0x45, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x53,
	0x45, 0x5f, 0x4e, 0x6f, 0x74, 0x5f, 0x53, 0x54, 0x61, 0x72, 0x74, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x41, 0x53, 0x45, 0x5f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x41, 0x53, 0x45, 0x5f, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x41, 0x53, 0x45, 0x5f, 0x4f, 0x66, 0x66, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x53,
	0x45, 0x5f, 0x45, 0x6e, 0x64, 0x10, 0x05, 0x42, 0x0b, 0x5a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_actor_proto_rawDescOnce sync.Once
	file_actor_proto_rawDescData = file_actor_proto_rawDesc
)

func file_actor_proto_rawDescGZIP() []byte {
	file_actor_proto_rawDescOnce.Do(func() {
		file_actor_proto_rawDescData = protoimpl.X.CompressGZIP(file_actor_proto_rawDescData)
	})
	return file_actor_proto_rawDescData
}

var file_actor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_actor_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_actor_proto_goTypes = []interface{}{
	(ActorStatusEnum)(0),           // 0: ActorStatusEnum
	(*Actor)(nil),                  // 1: Actor
	(*ActorOutPutTime)(nil),        // 2: ActorOutPutTime
	(*ActorInfo)(nil),              // 3: ActorInfo
	(*ActorOutPutStatus)(nil),      // 4: ActorOutPutStatus
	(*ActorGroupReq)(nil),          // 5: ActorGroupReq
	(*ActorShowReq)(nil),           // 6: ActorShowReq
	(*ActorUpdateReq)(nil),         // 7: ActorUpdateReq
	(*ActorSkillReq)(nil),          // 8: ActorSkillReq
	(*GetAllSkillReq)(nil),         // 9: GetAllSkillReq
	(*ActorsAvailableReq)(nil),     // 10: ActorsAvailableReq
	(*ActorsAvailableResp)(nil),    // 11: ActorsAvailableResp
	(*ActorStatusAndTimeItem)(nil), // 12: ActorStatusAndTimeItem
	(*ActorStatusAndTime)(nil),     // 13: ActorStatusAndTime
	nil,                            // 14: Actor.CollocationEntry
	nil,                            // 15: ActorInfo.ActorEntry
	nil,                            // 16: ActorInfo.BoxCountEntry
	nil,                            // 17: ActorOutPutStatus.StatusEntry
	nil,                            // 18: ActorStatusAndTime.ActorsEntry
	(*ItemPair)(nil),               // 19: ItemPair
}
var file_actor_proto_depIdxs = []int32{
	14, // 0: Actor.collocation:type_name -> Actor.CollocationEntry
	15, // 1: ActorInfo.actor:type_name -> ActorInfo.ActorEntry
	19, // 2: ActorInfo.group:type_name -> ItemPair
	16, // 3: ActorInfo.boxCount:type_name -> ActorInfo.BoxCountEntry
	17, // 4: ActorOutPutStatus.status:type_name -> ActorOutPutStatus.StatusEntry
	19, // 5: ActorGroupReq.group:type_name -> ItemPair
	18, // 6: ActorStatusAndTime.actors:type_name -> ActorStatusAndTime.ActorsEntry
	1,  // 7: ActorInfo.ActorEntry.value:type_name -> Actor
	2,  // 8: ActorOutPutStatus.StatusEntry.value:type_name -> ActorOutPutTime
	12, // 9: ActorStatusAndTime.ActorsEntry.value:type_name -> ActorStatusAndTimeItem
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_actor_proto_init() }
func file_actor_proto_init() {
	if File_actor_proto != nil {
		return
	}
	file_game_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_actor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actor); i {
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
		file_actor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorOutPutTime); i {
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
		file_actor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorInfo); i {
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
		file_actor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorOutPutStatus); i {
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
		file_actor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorGroupReq); i {
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
		file_actor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorShowReq); i {
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
		file_actor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorUpdateReq); i {
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
		file_actor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorSkillReq); i {
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
		file_actor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSkillReq); i {
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
		file_actor_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorsAvailableReq); i {
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
		file_actor_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorsAvailableResp); i {
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
		file_actor_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorStatusAndTimeItem); i {
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
		file_actor_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorStatusAndTime); i {
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
			RawDescriptor: file_actor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_actor_proto_goTypes,
		DependencyIndexes: file_actor_proto_depIdxs,
		EnumInfos:         file_actor_proto_enumTypes,
		MessageInfos:      file_actor_proto_msgTypes,
	}.Build()
	File_actor_proto = out.File
	file_actor_proto_rawDesc = nil
	file_actor_proto_goTypes = nil
	file_actor_proto_depIdxs = nil
}