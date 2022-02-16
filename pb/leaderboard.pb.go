// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: leaderboard.proto

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

type AREA_CODE int32

const (
	AREA_CODE_AREA_NONE    AREA_CODE = 0
	AREA_CODE_Anhui        AREA_CODE = 1
	AREA_CODE_Zhejiang     AREA_CODE = 2
	AREA_CODE_Jiangxi      AREA_CODE = 3
	AREA_CODE_Jiangsu      AREA_CODE = 4
	AREA_CODE_Jilin        AREA_CODE = 5
	AREA_CODE_Qinghai      AREA_CODE = 6
	AREA_CODE_Fujian       AREA_CODE = 7
	AREA_CODE_Heilongjiang AREA_CODE = 8
	AREA_CODE_Henan        AREA_CODE = 9
	AREA_CODE_Hebei        AREA_CODE = 10
	AREA_CODE_Hunan        AREA_CODE = 11
	AREA_CODE_Hubei        AREA_CODE = 12
	AREA_CODE_Xinjiang     AREA_CODE = 13
	AREA_CODE_Xizang       AREA_CODE = 14
	AREA_CODE_Gansu        AREA_CODE = 15
	AREA_CODE_Guangxi      AREA_CODE = 16
	AREA_CODE_Guizhou      AREA_CODE = 18
	AREA_CODE_Liaoning     AREA_CODE = 19
	AREA_CODE_Nei_Mongol   AREA_CODE = 20
	AREA_CODE_Ningxia      AREA_CODE = 21
	AREA_CODE_Beijing      AREA_CODE = 22
	AREA_CODE_Shanghai     AREA_CODE = 23
	AREA_CODE_Shanxi       AREA_CODE = 24
	AREA_CODE_Shandong     AREA_CODE = 25
	AREA_CODE_Shaanxi      AREA_CODE = 26
	AREA_CODE_Tianjin      AREA_CODE = 28
	AREA_CODE_Yunnan       AREA_CODE = 29
	AREA_CODE_Guangdong    AREA_CODE = 30
	AREA_CODE_Hainan       AREA_CODE = 31
	AREA_CODE_Sichuan      AREA_CODE = 32
	AREA_CODE_Chongqing    AREA_CODE = 33
	AREA_CODE_Taiwan       AREA_CODE = 34
	AREA_CODE_Hong_Kong    AREA_CODE = 35
	AREA_CODE_Macao        AREA_CODE = 36
	AREA_CODE_OTHER        AREA_CODE = 37
)

// Enum value maps for AREA_CODE.
var (
	AREA_CODE_name = map[int32]string{
		0:  "AREA_NONE",
		1:  "Anhui",
		2:  "Zhejiang",
		3:  "Jiangxi",
		4:  "Jiangsu",
		5:  "Jilin",
		6:  "Qinghai",
		7:  "Fujian",
		8:  "Heilongjiang",
		9:  "Henan",
		10: "Hebei",
		11: "Hunan",
		12: "Hubei",
		13: "Xinjiang",
		14: "Xizang",
		15: "Gansu",
		16: "Guangxi",
		18: "Guizhou",
		19: "Liaoning",
		20: "Nei_Mongol",
		21: "Ningxia",
		22: "Beijing",
		23: "Shanghai",
		24: "Shanxi",
		25: "Shandong",
		26: "Shaanxi",
		28: "Tianjin",
		29: "Yunnan",
		30: "Guangdong",
		31: "Hainan",
		32: "Sichuan",
		33: "Chongqing",
		34: "Taiwan",
		35: "Hong_Kong",
		36: "Macao",
		37: "OTHER",
	}
	AREA_CODE_value = map[string]int32{
		"AREA_NONE":    0,
		"Anhui":        1,
		"Zhejiang":     2,
		"Jiangxi":      3,
		"Jiangsu":      4,
		"Jilin":        5,
		"Qinghai":      6,
		"Fujian":       7,
		"Heilongjiang": 8,
		"Henan":        9,
		"Hebei":        10,
		"Hunan":        11,
		"Hubei":        12,
		"Xinjiang":     13,
		"Xizang":       14,
		"Gansu":        15,
		"Guangxi":      16,
		"Guizhou":      18,
		"Liaoning":     19,
		"Nei_Mongol":   20,
		"Ningxia":      21,
		"Beijing":      22,
		"Shanghai":     23,
		"Shanxi":       24,
		"Shandong":     25,
		"Shaanxi":      26,
		"Tianjin":      28,
		"Yunnan":       29,
		"Guangdong":    30,
		"Hainan":       31,
		"Sichuan":      32,
		"Chongqing":    33,
		"Taiwan":       34,
		"Hong_Kong":    35,
		"Macao":        36,
		"OTHER":        37,
	}
)

func (x AREA_CODE) Enum() *AREA_CODE {
	p := new(AREA_CODE)
	*p = x
	return p
}

func (x AREA_CODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AREA_CODE) Descriptor() protoreflect.EnumDescriptor {
	return file_leaderboard_proto_enumTypes[0].Descriptor()
}

func (AREA_CODE) Type() protoreflect.EnumType {
	return &file_leaderboard_proto_enumTypes[0]
}

func (x AREA_CODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AREA_CODE.Descriptor instead.
func (AREA_CODE) EnumDescriptor() ([]byte, []int) {
	return file_leaderboard_proto_rawDescGZIP(), []int{0}
}

type RANK_TYPE int32

const (
	RANK_TYPE_NONE_RANK          RANK_TYPE = 0
	RANK_TYPE_CUP_RANK           RANK_TYPE = 1 // 奖杯榜
	RANK_TYPE_INVITE_BATTLE_RANK RANK_TYPE = 2 // 约战榜
	RANK_TYPE_ACHIEVEMENT_RANK   RANK_TYPE = 3 // 成就榜
)

// Enum value maps for RANK_TYPE.
var (
	RANK_TYPE_name = map[int32]string{
		0: "NONE_RANK",
		1: "CUP_RANK",
		2: "INVITE_BATTLE_RANK",
		3: "ACHIEVEMENT_RANK",
	}
	RANK_TYPE_value = map[string]int32{
		"NONE_RANK":          0,
		"CUP_RANK":           1,
		"INVITE_BATTLE_RANK": 2,
		"ACHIEVEMENT_RANK":   3,
	}
)

func (x RANK_TYPE) Enum() *RANK_TYPE {
	p := new(RANK_TYPE)
	*p = x
	return p
}

func (x RANK_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RANK_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_leaderboard_proto_enumTypes[1].Descriptor()
}

func (RANK_TYPE) Type() protoreflect.EnumType {
	return &file_leaderboard_proto_enumTypes[1]
}

func (x RANK_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RANK_TYPE.Descriptor instead.
func (RANK_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_leaderboard_proto_rawDescGZIP(), []int{1}
}

// 范围类型
type RANGE_TYPE int32

const (
	RANGE_TYPE_NONE_RANGE   RANGE_TYPE = 0
	RANGE_TYPE_AREA_RANGE   RANGE_TYPE = 1 // 地区排行
	RANGE_TYPE_SERVER_RANGE RANGE_TYPE = 2 // 全服排行
	RANGE_TYPE_FRIEND_RANGE RANGE_TYPE = 3 // 好友排行
)

// Enum value maps for RANGE_TYPE.
var (
	RANGE_TYPE_name = map[int32]string{
		0: "NONE_RANGE",
		1: "AREA_RANGE",
		2: "SERVER_RANGE",
		3: "FRIEND_RANGE",
	}
	RANGE_TYPE_value = map[string]int32{
		"NONE_RANGE":   0,
		"AREA_RANGE":   1,
		"SERVER_RANGE": 2,
		"FRIEND_RANGE": 3,
	}
)

func (x RANGE_TYPE) Enum() *RANGE_TYPE {
	p := new(RANGE_TYPE)
	*p = x
	return p
}

func (x RANGE_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RANGE_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_leaderboard_proto_enumTypes[2].Descriptor()
}

func (RANGE_TYPE) Type() protoreflect.EnumType {
	return &file_leaderboard_proto_enumTypes[2]
}

func (x RANGE_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RANGE_TYPE.Descriptor instead.
func (RANGE_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_leaderboard_proto_rawDescGZIP(), []int{2}
}

type TIME_RANGE int32

const (
	TIME_RANGE_NONE_TIME    TIME_RANGE = 0
	TIME_RANGE_WEEK_TIME    TIME_RANGE = 1 // 周榜
	TIME_RANGE_SEASON_TIME  TIME_RANGE = 2 // 赛季榜
	TIME_RANGE_HISTORY_TIME TIME_RANGE = 3 // 历史榜
)

// Enum value maps for TIME_RANGE.
var (
	TIME_RANGE_name = map[int32]string{
		0: "NONE_TIME",
		1: "WEEK_TIME",
		2: "SEASON_TIME",
		3: "HISTORY_TIME",
	}
	TIME_RANGE_value = map[string]int32{
		"NONE_TIME":    0,
		"WEEK_TIME":    1,
		"SEASON_TIME":  2,
		"HISTORY_TIME": 3,
	}
)

func (x TIME_RANGE) Enum() *TIME_RANGE {
	p := new(TIME_RANGE)
	*p = x
	return p
}

func (x TIME_RANGE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TIME_RANGE) Descriptor() protoreflect.EnumDescriptor {
	return file_leaderboard_proto_enumTypes[3].Descriptor()
}

func (TIME_RANGE) Type() protoreflect.EnumType {
	return &file_leaderboard_proto_enumTypes[3]
}

func (x TIME_RANGE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TIME_RANGE.Descriptor instead.
func (TIME_RANGE) EnumDescriptor() ([]byte, []int) {
	return file_leaderboard_proto_rawDescGZIP(), []int{3}
}

type SetPlayerAreaCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviteCode uint64    `protobuf:"varint,1,opt,name=inviteCode,proto3" json:"inviteCode,omitempty"`
	AreaCode   AREA_CODE `protobuf:"varint,2,opt,name=areaCode,proto3,enum=AREA_CODE" json:"areaCode,omitempty"`
	Uid        uint64    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *SetPlayerAreaCodeReq) Reset() {
	*x = SetPlayerAreaCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPlayerAreaCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPlayerAreaCodeReq) ProtoMessage() {}

func (x *SetPlayerAreaCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPlayerAreaCodeReq.ProtoReflect.Descriptor instead.
func (*SetPlayerAreaCodeReq) Descriptor() ([]byte, []int) {
	return file_leaderboard_proto_rawDescGZIP(), []int{0}
}

func (x *SetPlayerAreaCodeReq) GetInviteCode() uint64 {
	if x != nil {
		return x.InviteCode
	}
	return 0
}

func (x *SetPlayerAreaCodeReq) GetAreaCode() AREA_CODE {
	if x != nil {
		return x.AreaCode
	}
	return AREA_CODE_AREA_NONE
}

func (x *SetPlayerAreaCodeReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type AreaCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SavedCode    AREA_CODE `protobuf:"varint,1,opt,name=savedCode,proto3,enum=AREA_CODE" json:"savedCode,omitempty"` // 记录到服务器的地址
	CurCode      AREA_CODE `protobuf:"varint,2,opt,name=curCode,proto3,enum=AREA_CODE" json:"curCode,omitempty"`     // 当前定位地址
	AreaModifyCD int64     `protobuf:"varint,3,opt,name=areaModifyCD,proto3" json:"areaModifyCD,omitempty"`          // 所属区域切换冷却倒计时秒
}

func (x *AreaCodeResp) Reset() {
	*x = AreaCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaCodeResp) ProtoMessage() {}

func (x *AreaCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaCodeResp.ProtoReflect.Descriptor instead.
func (*AreaCodeResp) Descriptor() ([]byte, []int) {
	return file_leaderboard_proto_rawDescGZIP(), []int{1}
}

func (x *AreaCodeResp) GetSavedCode() AREA_CODE {
	if x != nil {
		return x.SavedCode
	}
	return AREA_CODE_AREA_NONE
}

func (x *AreaCodeResp) GetCurCode() AREA_CODE {
	if x != nil {
		return x.CurCode
	}
	return AREA_CODE_AREA_NONE
}

func (x *AreaCodeResp) GetAreaModifyCD() int64 {
	if x != nil {
		return x.AreaModifyCD
	}
	return 0
}

type LeaderboardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankType  RANK_TYPE  `protobuf:"varint,1,opt,name=rankType,proto3,enum=RANK_TYPE" json:"rankType,omitempty"`
	TimeType  TIME_RANGE `protobuf:"varint,2,opt,name=timeType,proto3,enum=TIME_RANGE" json:"timeType,omitempty"`
	RangeType RANGE_TYPE `protobuf:"varint,3,opt,name=rangeType,proto3,enum=RANGE_TYPE" json:"rangeType,omitempty"`
	AreaCode  AREA_CODE  `protobuf:"varint,4,opt,name=areaCode,proto3,enum=AREA_CODE" json:"areaCode,omitempty"` // 地区代号
}

func (x *LeaderboardReq) Reset() {
	*x = LeaderboardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardReq) ProtoMessage() {}

func (x *LeaderboardReq) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardReq.ProtoReflect.Descriptor instead.
func (*LeaderboardReq) Descriptor() ([]byte, []int) {
	return file_leaderboard_proto_rawDescGZIP(), []int{2}
}

func (x *LeaderboardReq) GetRankType() RANK_TYPE {
	if x != nil {
		return x.RankType
	}
	return RANK_TYPE_NONE_RANK
}

func (x *LeaderboardReq) GetTimeType() TIME_RANGE {
	if x != nil {
		return x.TimeType
	}
	return TIME_RANGE_NONE_TIME
}

func (x *LeaderboardReq) GetRangeType() RANGE_TYPE {
	if x != nil {
		return x.RangeType
	}
	return RANGE_TYPE_NONE_RANGE
}

func (x *LeaderboardReq) GetAreaCode() AREA_CODE {
	if x != nil {
		return x.AreaCode
	}
	return AREA_CODE_AREA_NONE
}

type LeaderboardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankType         RANK_TYPE      `protobuf:"varint,1,opt,name=rankType,proto3,enum=RANK_TYPE" json:"rankType,omitempty"`
	RangeType        RANGE_TYPE     `protobuf:"varint,2,opt,name=rangeType,proto3,enum=RANGE_TYPE" json:"rangeType,omitempty"`
	TimeType         TIME_RANGE     `protobuf:"varint,3,opt,name=timeType,proto3,enum=TIME_RANGE" json:"timeType,omitempty"`
	PlayerBoards     []*PlayerBoard `protobuf:"bytes,4,rep,name=playerBoards,proto3" json:"playerBoards,omitempty"`
	OwnRank          int32          `protobuf:"varint,5,opt,name=ownRank,proto3" json:"ownRank,omitempty"` // 自己的排名,-1表示没上榜
	OwnInfo          *PlayerBoard   `protobuf:"bytes,6,opt,name=ownInfo,proto3" json:"ownInfo,omitempty"`
	WeekRankTimeLeft int64          `protobuf:"varint,7,opt,name=weekRankTimeLeft,proto3" json:"weekRankTimeLeft,omitempty"` // 周榜刷新倒计时秒
}

func (x *LeaderboardInfo) Reset() {
	*x = LeaderboardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardInfo) ProtoMessage() {}

func (x *LeaderboardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardInfo.ProtoReflect.Descriptor instead.
func (*LeaderboardInfo) Descriptor() ([]byte, []int) {
	return file_leaderboard_proto_rawDescGZIP(), []int{3}
}

func (x *LeaderboardInfo) GetRankType() RANK_TYPE {
	if x != nil {
		return x.RankType
	}
	return RANK_TYPE_NONE_RANK
}

func (x *LeaderboardInfo) GetRangeType() RANGE_TYPE {
	if x != nil {
		return x.RangeType
	}
	return RANGE_TYPE_NONE_RANGE
}

func (x *LeaderboardInfo) GetTimeType() TIME_RANGE {
	if x != nil {
		return x.TimeType
	}
	return TIME_RANGE_NONE_TIME
}

func (x *LeaderboardInfo) GetPlayerBoards() []*PlayerBoard {
	if x != nil {
		return x.PlayerBoards
	}
	return nil
}

func (x *LeaderboardInfo) GetOwnRank() int32 {
	if x != nil {
		return x.OwnRank
	}
	return 0
}

func (x *LeaderboardInfo) GetOwnInfo() *PlayerBoard {
	if x != nil {
		return x.OwnInfo
	}
	return nil
}

func (x *LeaderboardInfo) GetWeekRankTimeLeft() int64 {
	if x != nil {
		return x.WeekRankTimeLeft
	}
	return 0
}

type PlayerBoard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviteCode      uint64 `protobuf:"varint,1,opt,name=inviteCode,proto3" json:"inviteCode,omitempty"`           // 邀请码
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                        // 昵称
	Avatar          string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`                    // 头像
	PlayerAvatar    int32  `protobuf:"varint,4,opt,name=playerAvatar,proto3" json:"playerAvatar,omitempty"`       // 游戏内头像
	MaxCup          int32  `protobuf:"varint,5,opt,name=maxCup,proto3" json:"maxCup,omitempty"`                   // 杯数
	SeasonIncrCup   int32  `protobuf:"varint,6,opt,name=seasonIncrCup,proto3" json:"seasonIncrCup,omitempty"`     // 赛季增长杯数
	WeekIncrCup     int32  `protobuf:"varint,7,opt,name=weekIncrCup,proto3" json:"weekIncrCup,omitempty"`         // 周增长杯数
	Point           int32  `protobuf:"varint,8,opt,name=point,proto3" json:"point,omitempty"`                     // 成就积分
	SeasonIncrPoint int32  `protobuf:"varint,9,opt,name=seasonIncrPoint,proto3" json:"seasonIncrPoint,omitempty"` // 赛季增长成就积分
	WeekIncrPoint   int32  `protobuf:"varint,10,opt,name=weekIncrPoint,proto3" json:"weekIncrPoint,omitempty"`    // 周增长成就积分
	MaxStar         int32  `protobuf:"varint,11,opt,name=maxStar,proto3" json:"maxStar,omitempty"`                // 星辉
	SeasonIncrStar  int32  `protobuf:"varint,12,opt,name=seasonIncrStar,proto3" json:"seasonIncrStar,omitempty"`  // 赛季增长星辉
	WeekIncrStar    int32  `protobuf:"varint,13,opt,name=weekIncrStar,proto3" json:"weekIncrStar,omitempty"`      // 周增长星辉
	Title           int32  `protobuf:"varint,14,opt,name=title,proto3" json:"title,omitempty"`                    // 称号
	CurCup          int32  `protobuf:"varint,15,opt,name=curCup,proto3" json:"curCup,omitempty"`                  // 当前杯数
	Starlevel       int32  `protobuf:"varint,16,opt,name=starlevel,proto3" json:"starlevel,omitempty"`            // 约战星级
	IsFriend        bool   `protobuf:"varint,17,opt,name=isFriend,proto3" json:"isFriend,omitempty"`              // 是否是好友
	AvatarFrame     int32  `protobuf:"varint,18,opt,name=avatarFrame,proto3" json:"avatarFrame,omitempty"`        // 头像框
	ColorfulName    int32  `protobuf:"varint,19,opt,name=colorfulName,proto3" json:"colorfulName,omitempty"`      // 炫彩名
}

func (x *PlayerBoard) Reset() {
	*x = PlayerBoard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerBoard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerBoard) ProtoMessage() {}

func (x *PlayerBoard) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerBoard.ProtoReflect.Descriptor instead.
func (*PlayerBoard) Descriptor() ([]byte, []int) {
	return file_leaderboard_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerBoard) GetInviteCode() uint64 {
	if x != nil {
		return x.InviteCode
	}
	return 0
}

func (x *PlayerBoard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerBoard) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *PlayerBoard) GetPlayerAvatar() int32 {
	if x != nil {
		return x.PlayerAvatar
	}
	return 0
}

func (x *PlayerBoard) GetMaxCup() int32 {
	if x != nil {
		return x.MaxCup
	}
	return 0
}

func (x *PlayerBoard) GetSeasonIncrCup() int32 {
	if x != nil {
		return x.SeasonIncrCup
	}
	return 0
}

func (x *PlayerBoard) GetWeekIncrCup() int32 {
	if x != nil {
		return x.WeekIncrCup
	}
	return 0
}

func (x *PlayerBoard) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *PlayerBoard) GetSeasonIncrPoint() int32 {
	if x != nil {
		return x.SeasonIncrPoint
	}
	return 0
}

func (x *PlayerBoard) GetWeekIncrPoint() int32 {
	if x != nil {
		return x.WeekIncrPoint
	}
	return 0
}

func (x *PlayerBoard) GetMaxStar() int32 {
	if x != nil {
		return x.MaxStar
	}
	return 0
}

func (x *PlayerBoard) GetSeasonIncrStar() int32 {
	if x != nil {
		return x.SeasonIncrStar
	}
	return 0
}

func (x *PlayerBoard) GetWeekIncrStar() int32 {
	if x != nil {
		return x.WeekIncrStar
	}
	return 0
}

func (x *PlayerBoard) GetTitle() int32 {
	if x != nil {
		return x.Title
	}
	return 0
}

func (x *PlayerBoard) GetCurCup() int32 {
	if x != nil {
		return x.CurCup
	}
	return 0
}

func (x *PlayerBoard) GetStarlevel() int32 {
	if x != nil {
		return x.Starlevel
	}
	return 0
}

func (x *PlayerBoard) GetIsFriend() bool {
	if x != nil {
		return x.IsFriend
	}
	return false
}

func (x *PlayerBoard) GetAvatarFrame() int32 {
	if x != nil {
		return x.AvatarFrame
	}
	return 0
}

func (x *PlayerBoard) GetColorfulName() int32 {
	if x != nil {
		return x.ColorfulName
	}
	return 0
}

var File_leaderboard_proto protoreflect.FileDescriptor

var file_leaderboard_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x61,
	0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e,
	0x41, 0x52, 0x45, 0x41, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x61, 0x76, 0x65, 0x64, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x41, 0x52, 0x45, 0x41,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x52, 0x09, 0x73, 0x61, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x24, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x65, 0x61, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x43, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x72,
	0x65, 0x61, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43, 0x44, 0x22, 0xb4, 0x01, 0x0a, 0x0e, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a,
	0x08, 0x72, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0a, 0x2e, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x08, 0x72, 0x61, 0x6e,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29,
	0x0a, 0x09, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x09,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x72, 0x65,
	0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x41, 0x52,
	0x45, 0x41, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0xad, 0x02, 0x0a, 0x0f, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x52, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a,
	0x09, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x09, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x30, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x26, 0x0a,
	0x07, 0x6f, 0x77, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x07, 0x6f, 0x77,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x10, 0x77, 0x65, 0x65, 0x6b, 0x52, 0x61, 0x6e,
	0x6b, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x65, 0x66, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x77, 0x65, 0x65, 0x6b, 0x52, 0x61, 0x6e, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x65, 0x66,
	0x74, 0x22, 0xd7, 0x04, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x43, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x43, 0x75, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x63, 0x72, 0x43, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x63, 0x72, 0x43, 0x75, 0x70, 0x12,
	0x20, 0x0a, 0x0b, 0x77, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x63, 0x72, 0x43, 0x75, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x77, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x63, 0x72, 0x43, 0x75,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x49, 0x6e, 0x63, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x63, 0x72, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x63, 0x72, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x77, 0x65, 0x65, 0x6b, 0x49, 0x6e,
	0x63, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x74,
	0x61, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61,
	0x72, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x63, 0x72, 0x53,
	0x74, 0x61, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x49, 0x6e, 0x63, 0x72, 0x53, 0x74, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x65, 0x65,
	0x6b, 0x49, 0x6e, 0x63, 0x72, 0x53, 0x74, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x77, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x63, 0x72, 0x53, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x43, 0x75, 0x70, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x75, 0x72, 0x43, 0x75, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x66, 0x75, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x75, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0xdc, 0x03, 0x0a, 0x09,
	0x41, 0x52, 0x45, 0x41, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x52, 0x45,
	0x41, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x6e, 0x68, 0x75,
	0x69, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x5a, 0x68, 0x65, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4a, 0x69, 0x61, 0x6e, 0x67, 0x78, 0x69, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x4a, 0x69, 0x61, 0x6e, 0x67, 0x73, 0x75, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4a,
	0x69, 0x6c, 0x69, 0x6e, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x51, 0x69, 0x6e, 0x67, 0x68, 0x61,
	0x69, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x75, 0x6a, 0x69, 0x61, 0x6e, 0x10, 0x07, 0x12,
	0x10, 0x0a, 0x0c, 0x48, 0x65, 0x69, 0x6c, 0x6f, 0x6e, 0x67, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x10,
	0x08, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x65, 0x6e, 0x61, 0x6e, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05,
	0x48, 0x65, 0x62, 0x65, 0x69, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x75, 0x6e, 0x61, 0x6e,
	0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x75, 0x62, 0x65, 0x69, 0x10, 0x0c, 0x12, 0x0c, 0x0a,
	0x08, 0x58, 0x69, 0x6e, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x10, 0x0d, 0x12, 0x0a, 0x0a, 0x06, 0x58,
	0x69, 0x7a, 0x61, 0x6e, 0x67, 0x10, 0x0e, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x61, 0x6e, 0x73, 0x75,
	0x10, 0x0f, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x75, 0x61, 0x6e, 0x67, 0x78, 0x69, 0x10, 0x10, 0x12,
	0x0b, 0x0a, 0x07, 0x47, 0x75, 0x69, 0x7a, 0x68, 0x6f, 0x75, 0x10, 0x12, 0x12, 0x0c, 0x0a, 0x08,
	0x4c, 0x69, 0x61, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x13, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x65,
	0x69, 0x5f, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x6c, 0x10, 0x14, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x69,
	0x6e, 0x67, 0x78, 0x69, 0x61, 0x10, 0x15, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x65, 0x69, 0x6a, 0x69,
	0x6e, 0x67, 0x10, 0x16, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x68, 0x61, 0x6e, 0x67, 0x68, 0x61, 0x69,
	0x10, 0x17, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x68, 0x61, 0x6e, 0x78, 0x69, 0x10, 0x18, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x68, 0x61, 0x6e, 0x64, 0x6f, 0x6e, 0x67, 0x10, 0x19, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x68, 0x61, 0x61, 0x6e, 0x78, 0x69, 0x10, 0x1a, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x69, 0x61,
	0x6e, 0x6a, 0x69, 0x6e, 0x10, 0x1c, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x75, 0x6e, 0x6e, 0x61, 0x6e,
	0x10, 0x1d, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x75, 0x61, 0x6e, 0x67, 0x64, 0x6f, 0x6e, 0x67, 0x10,
	0x1e, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x61, 0x69, 0x6e, 0x61, 0x6e, 0x10, 0x1f, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x69, 0x63, 0x68, 0x75, 0x61, 0x6e, 0x10, 0x20, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x68,
	0x6f, 0x6e, 0x67, 0x71, 0x69, 0x6e, 0x67, 0x10, 0x21, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x61, 0x69,
	0x77, 0x61, 0x6e, 0x10, 0x22, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x6f, 0x6e, 0x67, 0x5f, 0x4b, 0x6f,
	0x6e, 0x67, 0x10, 0x23, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x61, 0x63, 0x61, 0x6f, 0x10, 0x24, 0x12,
	0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x25, 0x2a, 0x56, 0x0a, 0x09, 0x52, 0x41,
	0x4e, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x4e, 0x45, 0x5f,
	0x52, 0x41, 0x4e, 0x4b, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55, 0x50, 0x5f, 0x52, 0x41,
	0x4e, 0x4b, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x5f, 0x42,
	0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x4b, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10,
	0x41, 0x43, 0x48, 0x49, 0x45, 0x56, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x41, 0x4e, 0x4b,
	0x10, 0x03, 0x2a, 0x50, 0x0a, 0x0a, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x10, 0x03, 0x2a, 0x4d, 0x0a, 0x0a, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10,
	0x02, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x10, 0x03, 0x42, 0x0b, 0x5a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_leaderboard_proto_rawDescOnce sync.Once
	file_leaderboard_proto_rawDescData = file_leaderboard_proto_rawDesc
)

func file_leaderboard_proto_rawDescGZIP() []byte {
	file_leaderboard_proto_rawDescOnce.Do(func() {
		file_leaderboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_leaderboard_proto_rawDescData)
	})
	return file_leaderboard_proto_rawDescData
}

var file_leaderboard_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_leaderboard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_leaderboard_proto_goTypes = []interface{}{
	(AREA_CODE)(0),               // 0: AREA_CODE
	(RANK_TYPE)(0),               // 1: RANK_TYPE
	(RANGE_TYPE)(0),              // 2: RANGE_TYPE
	(TIME_RANGE)(0),              // 3: TIME_RANGE
	(*SetPlayerAreaCodeReq)(nil), // 4: SetPlayerAreaCodeReq
	(*AreaCodeResp)(nil),         // 5: AreaCodeResp
	(*LeaderboardReq)(nil),       // 6: LeaderboardReq
	(*LeaderboardInfo)(nil),      // 7: LeaderboardInfo
	(*PlayerBoard)(nil),          // 8: PlayerBoard
}
var file_leaderboard_proto_depIdxs = []int32{
	0,  // 0: SetPlayerAreaCodeReq.areaCode:type_name -> AREA_CODE
	0,  // 1: AreaCodeResp.savedCode:type_name -> AREA_CODE
	0,  // 2: AreaCodeResp.curCode:type_name -> AREA_CODE
	1,  // 3: LeaderboardReq.rankType:type_name -> RANK_TYPE
	3,  // 4: LeaderboardReq.timeType:type_name -> TIME_RANGE
	2,  // 5: LeaderboardReq.rangeType:type_name -> RANGE_TYPE
	0,  // 6: LeaderboardReq.areaCode:type_name -> AREA_CODE
	1,  // 7: LeaderboardInfo.rankType:type_name -> RANK_TYPE
	2,  // 8: LeaderboardInfo.rangeType:type_name -> RANGE_TYPE
	3,  // 9: LeaderboardInfo.timeType:type_name -> TIME_RANGE
	8,  // 10: LeaderboardInfo.playerBoards:type_name -> PlayerBoard
	8,  // 11: LeaderboardInfo.ownInfo:type_name -> PlayerBoard
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_leaderboard_proto_init() }
func file_leaderboard_proto_init() {
	if File_leaderboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_leaderboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPlayerAreaCodeReq); i {
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
		file_leaderboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaCodeResp); i {
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
		file_leaderboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardReq); i {
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
		file_leaderboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardInfo); i {
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
		file_leaderboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerBoard); i {
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
			RawDescriptor: file_leaderboard_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_leaderboard_proto_goTypes,
		DependencyIndexes: file_leaderboard_proto_depIdxs,
		EnumInfos:         file_leaderboard_proto_enumTypes,
		MessageInfos:      file_leaderboard_proto_msgTypes,
	}.Build()
	File_leaderboard_proto = out.File
	file_leaderboard_proto_rawDesc = nil
	file_leaderboard_proto_goTypes = nil
	file_leaderboard_proto_depIdxs = nil
}
