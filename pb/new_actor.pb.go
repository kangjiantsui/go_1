// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: new_actor.proto

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

type NewActorActivityItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPass     bool `protobuf:"varint,1,opt,name=isPass,proto3" json:"isPass,omitempty"`         // 是否通过
	IsRewarded bool `protobuf:"varint,2,opt,name=isRewarded,proto3" json:"isRewarded,omitempty"` // 是否领取过奖励
}

func (x *NewActorActivityItem) Reset() {
	*x = NewActorActivityItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_actor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewActorActivityItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewActorActivityItem) ProtoMessage() {}

func (x *NewActorActivityItem) ProtoReflect() protoreflect.Message {
	mi := &file_new_actor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewActorActivityItem.ProtoReflect.Descriptor instead.
func (*NewActorActivityItem) Descriptor() ([]byte, []int) {
	return file_new_actor_proto_rawDescGZIP(), []int{0}
}

func (x *NewActorActivityItem) GetIsPass() bool {
	if x != nil {
		return x.IsPass
	}
	return false
}

func (x *NewActorActivityItem) GetIsRewarded() bool {
	if x != nil {
		return x.IsRewarded
	}
	return false
}

type NewActorActivityRound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreCD   int64                           `protobuf:"varint,1,opt,name=preCD,proto3" json:"preCD,omitempty"`                                                                                       // 活动预热开始时间倒计时
	StartCD int64                           `protobuf:"varint,2,opt,name=startCD,proto3" json:"startCD,omitempty"`                                                                                   // 活动开始倒计时
	CloseCD int64                           `protobuf:"varint,3,opt,name=closeCD,proto3" json:"closeCD,omitempty"`                                                                                   // 活动无法进入战斗时间倒计时
	EndCD   int64                           `protobuf:"varint,4,opt,name=endCD,proto3" json:"endCD,omitempty"`                                                                                       // 活动结束倒计时
	Item    map[int32]*NewActorActivityItem `protobuf:"bytes,5,rep,name=item,proto3" json:"item,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 场次记录
}

func (x *NewActorActivityRound) Reset() {
	*x = NewActorActivityRound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_actor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewActorActivityRound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewActorActivityRound) ProtoMessage() {}

func (x *NewActorActivityRound) ProtoReflect() protoreflect.Message {
	mi := &file_new_actor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewActorActivityRound.ProtoReflect.Descriptor instead.
func (*NewActorActivityRound) Descriptor() ([]byte, []int) {
	return file_new_actor_proto_rawDescGZIP(), []int{1}
}

func (x *NewActorActivityRound) GetPreCD() int64 {
	if x != nil {
		return x.PreCD
	}
	return 0
}

func (x *NewActorActivityRound) GetStartCD() int64 {
	if x != nil {
		return x.StartCD
	}
	return 0
}

func (x *NewActorActivityRound) GetCloseCD() int64 {
	if x != nil {
		return x.CloseCD
	}
	return 0
}

func (x *NewActorActivityRound) GetEndCD() int64 {
	if x != nil {
		return x.EndCD
	}
	return 0
}

func (x *NewActorActivityRound) GetItem() map[int32]*NewActorActivityItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type NewActorActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round map[int32]*NewActorActivityRound `protobuf:"bytes,1,rep,name=round,proto3" json:"round,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 期数记录
}

func (x *NewActorActivity) Reset() {
	*x = NewActorActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_actor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewActorActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewActorActivity) ProtoMessage() {}

func (x *NewActorActivity) ProtoReflect() protoreflect.Message {
	mi := &file_new_actor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewActorActivity.ProtoReflect.Descriptor instead.
func (*NewActorActivity) Descriptor() ([]byte, []int) {
	return file_new_actor_proto_rawDescGZIP(), []int{2}
}

func (x *NewActorActivity) GetRound() map[int32]*NewActorActivityRound {
	if x != nil {
		return x.Round
	}
	return nil
}

type NewActorActivityUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rid     int32                 `protobuf:"varint,1,opt,name=rid,proto3" json:"rid,omitempty"`        // 期数id
	Iid     int32                 `protobuf:"varint,2,opt,name=iid,proto3" json:"iid,omitempty"`        // 场次id
	AUpdate *NewActorActivityItem `protobuf:"bytes,3,opt,name=aUpdate,proto3" json:"aUpdate,omitempty"` // 记录更新后值
	IUpdate *ItemUpdateResult     `protobuf:"bytes,4,opt,name=iUpdate,proto3" json:"iUpdate,omitempty"` // 奖励更新
}

func (x *NewActorActivityUpdate) Reset() {
	*x = NewActorActivityUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_actor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewActorActivityUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewActorActivityUpdate) ProtoMessage() {}

func (x *NewActorActivityUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_new_actor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewActorActivityUpdate.ProtoReflect.Descriptor instead.
func (*NewActorActivityUpdate) Descriptor() ([]byte, []int) {
	return file_new_actor_proto_rawDescGZIP(), []int{3}
}

func (x *NewActorActivityUpdate) GetRid() int32 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *NewActorActivityUpdate) GetIid() int32 {
	if x != nil {
		return x.Iid
	}
	return 0
}

func (x *NewActorActivityUpdate) GetAUpdate() *NewActorActivityItem {
	if x != nil {
		return x.AUpdate
	}
	return nil
}

func (x *NewActorActivityUpdate) GetIUpdate() *ItemUpdateResult {
	if x != nil {
		return x.IUpdate
	}
	return nil
}

var File_new_actor_proto protoreflect.FileDescriptor

var file_new_actor_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a,
	0x14, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x22, 0xfd, 0x01,
	0x0a, 0x15, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x65, 0x43, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x65, 0x43, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x43, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x43,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x64, 0x43, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x65, 0x6e, 0x64, 0x43, 0x44, 0x12, 0x34, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x1a, 0x4e, 0x0a,
	0x09, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4e, 0x65,
	0x77, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x98, 0x01,
	0x0a, 0x10, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x1a, 0x50, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9a, 0x01, 0x0a, 0x16, 0x4e, 0x65, 0x77,
	0x41, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x69, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x61, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x63,
	0x74, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x07, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x69, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x69, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_new_actor_proto_rawDescOnce sync.Once
	file_new_actor_proto_rawDescData = file_new_actor_proto_rawDesc
)

func file_new_actor_proto_rawDescGZIP() []byte {
	file_new_actor_proto_rawDescOnce.Do(func() {
		file_new_actor_proto_rawDescData = protoimpl.X.CompressGZIP(file_new_actor_proto_rawDescData)
	})
	return file_new_actor_proto_rawDescData
}

var file_new_actor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_new_actor_proto_goTypes = []interface{}{
	(*NewActorActivityItem)(nil),   // 0: NewActorActivityItem
	(*NewActorActivityRound)(nil),  // 1: NewActorActivityRound
	(*NewActorActivity)(nil),       // 2: NewActorActivity
	(*NewActorActivityUpdate)(nil), // 3: NewActorActivityUpdate
	nil,                            // 4: NewActorActivityRound.ItemEntry
	nil,                            // 5: NewActorActivity.RoundEntry
	(*ItemUpdateResult)(nil),       // 6: ItemUpdateResult
}
var file_new_actor_proto_depIdxs = []int32{
	4, // 0: NewActorActivityRound.item:type_name -> NewActorActivityRound.ItemEntry
	5, // 1: NewActorActivity.round:type_name -> NewActorActivity.RoundEntry
	0, // 2: NewActorActivityUpdate.aUpdate:type_name -> NewActorActivityItem
	6, // 3: NewActorActivityUpdate.iUpdate:type_name -> ItemUpdateResult
	0, // 4: NewActorActivityRound.ItemEntry.value:type_name -> NewActorActivityItem
	1, // 5: NewActorActivity.RoundEntry.value:type_name -> NewActorActivityRound
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_new_actor_proto_init() }
func file_new_actor_proto_init() {
	if File_new_actor_proto != nil {
		return
	}
	file_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_new_actor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewActorActivityItem); i {
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
		file_new_actor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewActorActivityRound); i {
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
		file_new_actor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewActorActivity); i {
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
		file_new_actor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewActorActivityUpdate); i {
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
			RawDescriptor: file_new_actor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_new_actor_proto_goTypes,
		DependencyIndexes: file_new_actor_proto_depIdxs,
		MessageInfos:      file_new_actor_proto_msgTypes,
	}.Build()
	File_new_actor_proto = out.File
	file_new_actor_proto_rawDesc = nil
	file_new_actor_proto_goTypes = nil
	file_new_actor_proto_depIdxs = nil
}