// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: base.proto

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

// 请求包装
type ReqWrap1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`                    //版本号, 默认为1, 必填
	GameId   string `protobuf:"bytes,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`        ///游戏ID，必填
	Cmd      int32  `protobuf:"varint,3,opt,name=cmd,proto3" json:"cmd,omitempty"`                           ///请求命令字, 必填
	Seq      uint32 `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`                           ///请求序列号
	AuthKey  string `protobuf:"bytes,5,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`     ///鉴权key, 传token值, sdk必填
	AuthType uint32 `protobuf:"varint,6,opt,name=auth_type,json=authType,proto3" json:"auth_type,omitempty"` ///鉴权类型
	Uid      uint64 `protobuf:"varint,7,opt,name=uid,proto3" json:"uid,omitempty"`                           ///用户唯一标示，保留
	PlayerId string `protobuf:"bytes,8,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`  ///后端生成的玩家ID, 接入层负责校验合法性, 必填
	Body     []byte `protobuf:"bytes,9,opt,name=body,proto3" json:"body,omitempty"`                          ///
}

func (x *ReqWrap1) Reset() {
	*x = ReqWrap1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqWrap1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqWrap1) ProtoMessage() {}

func (x *ReqWrap1) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqWrap1.ProtoReflect.Descriptor instead.
func (*ReqWrap1) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *ReqWrap1) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ReqWrap1) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *ReqWrap1) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *ReqWrap1) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *ReqWrap1) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

func (x *ReqWrap1) GetAuthType() uint32 {
	if x != nil {
		return x.AuthType
	}
	return 0
}

func (x *ReqWrap1) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReqWrap1) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *ReqWrap1) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// 推送包装
type BstWrap1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`             ///版本号, 默认为1,, 必填
	GameId  string `protobuf:"bytes,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"` ///游戏ID，必填
	Cmd     int32  `protobuf:"varint,3,opt,name=cmd,proto3" json:"cmd,omitempty"`                    ///请求命令字, 必填
	Body    []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`                   ///
}

func (x *BstWrap1) Reset() {
	*x = BstWrap1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BstWrap1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BstWrap1) ProtoMessage() {}

func (x *BstWrap1) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BstWrap1.ProtoReflect.Descriptor instead.
func (*BstWrap1) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *BstWrap1) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *BstWrap1) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *BstWrap1) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *BstWrap1) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

///统一回包，err_code和err_msg为框架错误信息，业务错误码在body中
type RspWrap1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  int32  `protobuf:"varint,1,opt,name=cmd,proto3" json:"cmd,omitempty"`   // 业务码
	Seq  uint32 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`   ///请求序列号
	Code int32  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"` ///错误码(0~10000为系统错误)
	Msg  string `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`    ///错误信息
	Body []byte `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`  ///
}

func (x *RspWrap1) Reset() {
	*x = RspWrap1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspWrap1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspWrap1) ProtoMessage() {}

func (x *RspWrap1) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspWrap1.ProtoReflect.Descriptor instead.
func (*RspWrap1) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *RspWrap1) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *RspWrap1) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *RspWrap1) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RspWrap1) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RspWrap1) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// 踢下线
type RspKick struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"` //  kick 原因
}

func (x *RspKick) Reset() {
	*x = RspKick{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspKick) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspKick) ProtoMessage() {}

func (x *RspKick) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspKick.ProtoReflect.Descriptor instead.
func (*RspKick) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *RspKick) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x71, 0x57, 0x72, 0x61, 0x70, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x63, 0x0a, 0x08, 0x42,
	0x73, 0x74, 0x57, 0x72, 0x61, 0x70, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x68, 0x0a, 0x08, 0x52, 0x73, 0x70, 0x57, 0x72, 0x61, 0x70, 0x31, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x21, 0x0a, 0x07, 0x52, 0x73,
	0x70, 0x4b, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x0b, 0x5a,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_base_proto_goTypes = []interface{}{
	(*ReqWrap1)(nil), // 0: ReqWrap1
	(*BstWrap1)(nil), // 1: BstWrap1
	(*RspWrap1)(nil), // 2: RspWrap1
	(*RspKick)(nil),  // 3: RspKick
}
var file_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqWrap1); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BstWrap1); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspWrap1); i {
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
		file_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspKick); i {
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
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
