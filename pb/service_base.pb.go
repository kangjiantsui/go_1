// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: service_base.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_service_base_proto protoreflect.FileDescriptor

var file_service_base_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe3, 0x04, 0x0a, 0x04, 0x42, 0x61, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x0e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x09, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x62, 0x01, 0x2a, 0xca, 0x3e, 0x02,
	0x08, 0x02, 0x12, 0x49, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x09,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x3a, 0x01, 0x2a, 0x62, 0x01, 0x2a, 0xca, 0x3e, 0x03, 0x08, 0xe8, 0x07, 0x12, 0x4c, 0x0a,
	0x0a, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x0f,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a,
	0x01, 0x2a, 0x62, 0x01, 0x2a, 0xca, 0x3e, 0x03, 0x08, 0xe3, 0x07, 0x12, 0x4f, 0x0a, 0x0b, 0x57,
	0x65, 0x63, 0x68, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x57, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x3a, 0x01, 0x2a, 0x62, 0x01, 0x2a, 0xca, 0x3e, 0x03, 0x08, 0xe7, 0x07, 0x12, 0x43, 0x0a, 0x07,
	0x51, 0x51, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0b, 0x2e, 0x51, 0x51, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x71, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x62, 0x01, 0x2a, 0xca, 0x3e, 0x03, 0x08, 0xe6,
	0x07, 0x12, 0x4f, 0x0a, 0x0b, 0x4f, 0x68, 0x61, 0x79, 0x6f, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x0f, 0x2e, 0x4f, 0x68, 0x61, 0x79, 0x6f, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x09, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x24, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x68, 0x61, 0x79, 0x6f, 0x6f,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x62, 0x01, 0x2a, 0xca, 0x3e, 0x03, 0x08,
	0xe5, 0x07, 0x12, 0x3b, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0a, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x3a, 0x01, 0x2a, 0x62, 0x01, 0x2a, 0xca, 0x3e, 0x03, 0x08, 0xa0, 0x06, 0x12,
	0x45, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x62, 0x01, 0x2a, 0xca,
	0x3e, 0x04, 0x08, 0x89, 0xa4, 0x01, 0x1a, 0x06, 0xd0, 0x3e, 0x00, 0xe0, 0x3e, 0x02, 0x42, 0x0b,
	0x5a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_service_base_proto_goTypes = []interface{}{
	(*CheckLoginReq)(nil),  // 0: CheckLoginReq
	(*BaseLoginReq)(nil),   // 1: BaseLoginReq
	(*AppleLoginReq)(nil),  // 2: AppleLoginReq
	(*WechatLoginReq)(nil), // 3: WechatLoginReq
	(*QQLoginReq)(nil),     // 4: QQLoginReq
	(*OhayooLoginReq)(nil), // 5: OhayooLoginReq
	(*Empty)(nil),          // 6: Empty
	(*LoginRsp)(nil),       // 7: LoginRsp
	(*NoticeResp)(nil),     // 8: NoticeResp
}
var file_service_base_proto_depIdxs = []int32{
	0, // 0: Base.CheckLogin:input_type -> CheckLoginReq
	1, // 1: Base.BaseLogin:input_type -> BaseLoginReq
	2, // 2: Base.AppleLogin:input_type -> AppleLoginReq
	3, // 3: Base.WechatLogin:input_type -> WechatLoginReq
	4, // 4: Base.QQLogin:input_type -> QQLoginReq
	5, // 5: Base.OhayooLogin:input_type -> OhayooLoginReq
	6, // 6: Base.HeartBeat:input_type -> Empty
	6, // 7: Base.GetNotice:input_type -> Empty
	7, // 8: Base.CheckLogin:output_type -> LoginRsp
	7, // 9: Base.BaseLogin:output_type -> LoginRsp
	7, // 10: Base.AppleLogin:output_type -> LoginRsp
	7, // 11: Base.WechatLogin:output_type -> LoginRsp
	7, // 12: Base.QQLogin:output_type -> LoginRsp
	7, // 13: Base.OhayooLogin:output_type -> LoginRsp
	6, // 14: Base.HeartBeat:output_type -> Empty
	8, // 15: Base.GetNotice:output_type -> NoticeResp
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_base_proto_init() }
func file_service_base_proto_init() {
	if File_service_base_proto != nil {
		return
	}
	file_extend_proto_init()
	file_game_proto_init()
	file_player_proto_init()
	file_notice_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_base_proto_goTypes,
		DependencyIndexes: file_service_base_proto_depIdxs,
	}.Build()
	File_service_base_proto = out.File
	file_service_base_proto_rawDesc = nil
	file_service_base_proto_goTypes = nil
	file_service_base_proto_depIdxs = nil
}
