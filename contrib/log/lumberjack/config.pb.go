// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.3
// source: third_party/contrib/log/lumberjack/config.proto

package lumberjack

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Maxsize    int32  `protobuf:"varint,2,opt,name=maxsize,proto3" json:"maxsize,omitempty"`
	Maxage     int32  `protobuf:"varint,3,opt,name=maxage,proto3" json:"maxage,omitempty"`
	Maxbackups int32  `protobuf:"varint,4,opt,name=maxbackups,proto3" json:"maxbackups,omitempty"`
	Localtime  bool   `protobuf:"varint,5,opt,name=localtime,proto3" json:"localtime,omitempty"`
	Compress   bool   `protobuf:"varint,6,opt,name=compress,proto3" json:"compress,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_contrib_log_lumberjack_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_contrib_log_lumberjack_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_third_party_contrib_log_lumberjack_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Config) GetMaxsize() int32 {
	if x != nil {
		return x.Maxsize
	}
	return 0
}

func (x *Config) GetMaxage() int32 {
	if x != nil {
		return x.Maxage
	}
	return 0
}

func (x *Config) GetMaxbackups() int32 {
	if x != nil {
		return x.Maxbackups
	}
	return 0
}

func (x *Config) GetLocaltime() bool {
	if x != nil {
		return x.Localtime
	}
	return false
}

func (x *Config) GetCompress() bool {
	if x != nil {
		return x.Compress
	}
	return false
}

var File_third_party_contrib_log_lumberjack_config_proto protoreflect.FileDescriptor

var file_third_party_contrib_log_lumberjack_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x6a, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x6a, 0x61, 0x63, 0x6b, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb0, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6d, 0x61, 0x78, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61,
	0x78, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x6c, 0x6f, 0x67,
	0x2f, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x6a, 0x61, 0x63, 0x6b, 0x3b, 0x6c, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x6a, 0x61, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_contrib_log_lumberjack_config_proto_rawDescOnce sync.Once
	file_third_party_contrib_log_lumberjack_config_proto_rawDescData = file_third_party_contrib_log_lumberjack_config_proto_rawDesc
)

func file_third_party_contrib_log_lumberjack_config_proto_rawDescGZIP() []byte {
	file_third_party_contrib_log_lumberjack_config_proto_rawDescOnce.Do(func() {
		file_third_party_contrib_log_lumberjack_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_contrib_log_lumberjack_config_proto_rawDescData)
	})
	return file_third_party_contrib_log_lumberjack_config_proto_rawDescData
}

var file_third_party_contrib_log_lumberjack_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_third_party_contrib_log_lumberjack_config_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: lumberjack.Config
}
var file_third_party_contrib_log_lumberjack_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_third_party_contrib_log_lumberjack_config_proto_init() }
func file_third_party_contrib_log_lumberjack_config_proto_init() {
	if File_third_party_contrib_log_lumberjack_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_contrib_log_lumberjack_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_third_party_contrib_log_lumberjack_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_contrib_log_lumberjack_config_proto_goTypes,
		DependencyIndexes: file_third_party_contrib_log_lumberjack_config_proto_depIdxs,
		MessageInfos:      file_third_party_contrib_log_lumberjack_config_proto_msgTypes,
	}.Build()
	File_third_party_contrib_log_lumberjack_config_proto = out.File
	file_third_party_contrib_log_lumberjack_config_proto_rawDesc = nil
	file_third_party_contrib_log_lumberjack_config_proto_goTypes = nil
	file_third_party_contrib_log_lumberjack_config_proto_depIdxs = nil
}
