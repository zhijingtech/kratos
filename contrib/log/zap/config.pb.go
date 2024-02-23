// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.3
// source: third_party/contrib/log/zap/config.proto

package zap

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

	Level             string            `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	Encoding          string            `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding,omitempty"`
	Development       bool              `protobuf:"varint,3,opt,name=development,proto3" json:"development,omitempty"`
	DisableCaller     bool              `protobuf:"varint,4,opt,name=disableCaller,proto3" json:"disableCaller,omitempty"`
	DisableStacktrace bool              `protobuf:"varint,5,opt,name=disableStacktrace,proto3" json:"disableStacktrace,omitempty"`
	OutputPaths       []string          `protobuf:"bytes,6,rep,name=outputPaths,proto3" json:"outputPaths,omitempty"`
	ErrorOutputPaths  []string          `protobuf:"bytes,7,rep,name=errorOutputPaths,proto3" json:"errorOutputPaths,omitempty"`
	InitialFields     map[string]string `protobuf:"bytes,8,rep,name=initialFields,proto3" json:"initialFields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Sampling          *SamplingConfig   `protobuf:"bytes,9,opt,name=sampling,proto3" json:"sampling,omitempty"`
	EncoderConfig     *EncoderConfig    `protobuf:"bytes,10,opt,name=encoderConfig,proto3" json:"encoderConfig,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_contrib_log_zap_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_contrib_log_zap_config_proto_msgTypes[0]
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
	return file_third_party_contrib_log_zap_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Config) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *Config) GetDevelopment() bool {
	if x != nil {
		return x.Development
	}
	return false
}

func (x *Config) GetDisableCaller() bool {
	if x != nil {
		return x.DisableCaller
	}
	return false
}

func (x *Config) GetDisableStacktrace() bool {
	if x != nil {
		return x.DisableStacktrace
	}
	return false
}

func (x *Config) GetOutputPaths() []string {
	if x != nil {
		return x.OutputPaths
	}
	return nil
}

func (x *Config) GetErrorOutputPaths() []string {
	if x != nil {
		return x.ErrorOutputPaths
	}
	return nil
}

func (x *Config) GetInitialFields() map[string]string {
	if x != nil {
		return x.InitialFields
	}
	return nil
}

func (x *Config) GetSampling() *SamplingConfig {
	if x != nil {
		return x.Sampling
	}
	return nil
}

func (x *Config) GetEncoderConfig() *EncoderConfig {
	if x != nil {
		return x.EncoderConfig
	}
	return nil
}

type SamplingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Initial    int32 `protobuf:"varint,1,opt,name=initial,proto3" json:"initial,omitempty"`
	Thereafter int32 `protobuf:"varint,2,opt,name=thereafter,proto3" json:"thereafter,omitempty"`
}

func (x *SamplingConfig) Reset() {
	*x = SamplingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_contrib_log_zap_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamplingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamplingConfig) ProtoMessage() {}

func (x *SamplingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_contrib_log_zap_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamplingConfig.ProtoReflect.Descriptor instead.
func (*SamplingConfig) Descriptor() ([]byte, []int) {
	return file_third_party_contrib_log_zap_config_proto_rawDescGZIP(), []int{1}
}

func (x *SamplingConfig) GetInitial() int32 {
	if x != nil {
		return x.Initial
	}
	return 0
}

func (x *SamplingConfig) GetThereafter() int32 {
	if x != nil {
		return x.Thereafter
	}
	return 0
}

type EncoderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageKey       string `protobuf:"bytes,1,opt,name=messageKey,proto3" json:"messageKey,omitempty"`
	LevelKey         string `protobuf:"bytes,2,opt,name=levelKey,proto3" json:"levelKey,omitempty"`
	TimeKey          string `protobuf:"bytes,3,opt,name=timeKey,proto3" json:"timeKey,omitempty"`
	NameKey          string `protobuf:"bytes,4,opt,name=nameKey,proto3" json:"nameKey,omitempty"`
	CallerKey        string `protobuf:"bytes,5,opt,name=callerKey,proto3" json:"callerKey,omitempty"`
	FunctionKey      string `protobuf:"bytes,6,opt,name=functionKey,proto3" json:"functionKey,omitempty"`
	StacktraceKey    string `protobuf:"bytes,7,opt,name=stacktraceKey,proto3" json:"stacktraceKey,omitempty"`
	LineEnding       string `protobuf:"bytes,8,opt,name=lineEnding,proto3" json:"lineEnding,omitempty"`
	ConsoleSeparator string `protobuf:"bytes,9,opt,name=consoleSeparator,proto3" json:"consoleSeparator,omitempty"`
	SkipLineEnding   bool   `protobuf:"varint,10,opt,name=skipLineEnding,proto3" json:"skipLineEnding,omitempty"`
}

func (x *EncoderConfig) Reset() {
	*x = EncoderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_contrib_log_zap_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncoderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncoderConfig) ProtoMessage() {}

func (x *EncoderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_contrib_log_zap_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncoderConfig.ProtoReflect.Descriptor instead.
func (*EncoderConfig) Descriptor() ([]byte, []int) {
	return file_third_party_contrib_log_zap_config_proto_rawDescGZIP(), []int{2}
}

func (x *EncoderConfig) GetMessageKey() string {
	if x != nil {
		return x.MessageKey
	}
	return ""
}

func (x *EncoderConfig) GetLevelKey() string {
	if x != nil {
		return x.LevelKey
	}
	return ""
}

func (x *EncoderConfig) GetTimeKey() string {
	if x != nil {
		return x.TimeKey
	}
	return ""
}

func (x *EncoderConfig) GetNameKey() string {
	if x != nil {
		return x.NameKey
	}
	return ""
}

func (x *EncoderConfig) GetCallerKey() string {
	if x != nil {
		return x.CallerKey
	}
	return ""
}

func (x *EncoderConfig) GetFunctionKey() string {
	if x != nil {
		return x.FunctionKey
	}
	return ""
}

func (x *EncoderConfig) GetStacktraceKey() string {
	if x != nil {
		return x.StacktraceKey
	}
	return ""
}

func (x *EncoderConfig) GetLineEnding() string {
	if x != nil {
		return x.LineEnding
	}
	return ""
}

func (x *EncoderConfig) GetConsoleSeparator() string {
	if x != nil {
		return x.ConsoleSeparator
	}
	return ""
}

func (x *EncoderConfig) GetSkipLineEnding() bool {
	if x != nil {
		return x.SkipLineEnding
	}
	return false
}

var File_third_party_contrib_log_zap_config_proto protoreflect.FileDescriptor

var file_third_party_contrib_log_zap_config_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x7a, 0x61, 0x70, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x7a, 0x61, 0x70, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf1, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x61, 0x70, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x7a, 0x61, 0x70,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x08, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x0d, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x7a, 0x61, 0x70, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x40, 0x0a, 0x12, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a, 0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x72, 0x65, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x72, 0x65, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x22, 0xd9, 0x02, 0x0a, 0x0d, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x61, 0x6d,
	0x65, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x61, 0x6d, 0x65,
	0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x6e,
	0x65, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x69, 0x6e, 0x65, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x70, 0x61,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x70, 0x4c, 0x69, 0x6e,
	0x65, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73,
	0x6b, 0x69, 0x70, 0x4c, 0x69, 0x6e, 0x65, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x69, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x7a, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x3b,
	0x7a, 0x61, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_contrib_log_zap_config_proto_rawDescOnce sync.Once
	file_third_party_contrib_log_zap_config_proto_rawDescData = file_third_party_contrib_log_zap_config_proto_rawDesc
)

func file_third_party_contrib_log_zap_config_proto_rawDescGZIP() []byte {
	file_third_party_contrib_log_zap_config_proto_rawDescOnce.Do(func() {
		file_third_party_contrib_log_zap_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_contrib_log_zap_config_proto_rawDescData)
	})
	return file_third_party_contrib_log_zap_config_proto_rawDescData
}

var file_third_party_contrib_log_zap_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_third_party_contrib_log_zap_config_proto_goTypes = []interface{}{
	(*Config)(nil),         // 0: zap.Config
	(*SamplingConfig)(nil), // 1: zap.SamplingConfig
	(*EncoderConfig)(nil),  // 2: zap.EncoderConfig
	nil,                    // 3: zap.Config.InitialFieldsEntry
}
var file_third_party_contrib_log_zap_config_proto_depIdxs = []int32{
	3, // 0: zap.Config.initialFields:type_name -> zap.Config.InitialFieldsEntry
	1, // 1: zap.Config.sampling:type_name -> zap.SamplingConfig
	2, // 2: zap.Config.encoderConfig:type_name -> zap.EncoderConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_third_party_contrib_log_zap_config_proto_init() }
func file_third_party_contrib_log_zap_config_proto_init() {
	if File_third_party_contrib_log_zap_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_contrib_log_zap_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_third_party_contrib_log_zap_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamplingConfig); i {
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
		file_third_party_contrib_log_zap_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncoderConfig); i {
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
			RawDescriptor: file_third_party_contrib_log_zap_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_contrib_log_zap_config_proto_goTypes,
		DependencyIndexes: file_third_party_contrib_log_zap_config_proto_depIdxs,
		MessageInfos:      file_third_party_contrib_log_zap_config_proto_msgTypes,
	}.Build()
	File_third_party_contrib_log_zap_config_proto = out.File
	file_third_party_contrib_log_zap_config_proto_rawDesc = nil
	file_third_party_contrib_log_zap_config_proto_goTypes = nil
	file_third_party_contrib_log_zap_config_proto_depIdxs = nil
}
