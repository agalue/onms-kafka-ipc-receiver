// Source: https://github.com/OpenNMS/opennms/blob/master/core/ipc/sink/common/src/main/proto/sink-message.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: sink.proto

package sink

import (
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

type SinkMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId          string            `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Content            []byte            `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	CurrentChunkNumber int32             `protobuf:"varint,3,opt,name=current_chunk_number,json=currentChunkNumber,proto3" json:"current_chunk_number,omitempty"`
	TotalChunks        int32             `protobuf:"varint,4,opt,name=total_chunks,json=totalChunks,proto3" json:"total_chunks,omitempty"`
	TracingInfo        map[string]string `protobuf:"bytes,5,rep,name=tracing_info,json=tracingInfo,proto3" json:"tracing_info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SinkMessage) Reset() {
	*x = SinkMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SinkMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SinkMessage) ProtoMessage() {}

func (x *SinkMessage) ProtoReflect() protoreflect.Message {
	mi := &file_sink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SinkMessage.ProtoReflect.Descriptor instead.
func (*SinkMessage) Descriptor() ([]byte, []int) {
	return file_sink_proto_rawDescGZIP(), []int{0}
}

func (x *SinkMessage) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *SinkMessage) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *SinkMessage) GetCurrentChunkNumber() int32 {
	if x != nil {
		return x.CurrentChunkNumber
	}
	return 0
}

func (x *SinkMessage) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

func (x *SinkMessage) GetTracingInfo() map[string]string {
	if x != nil {
		return x.TracingInfo
	}
	return nil
}

var File_sink_proto protoreflect.FileDescriptor

var file_sink_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x69,
	0x6e, 0x6b, 0x22, 0xa2, 0x02, 0x0a, 0x0b, 0x53, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x12, 0x45, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x69,
	0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x3e, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3e, 0x0a, 0x1f, 0x6f, 0x72, 0x67, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x6e, 0x6d, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x70, 0x63, 0x2e,
	0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x11, 0x53, 0x69, 0x6e, 0x6b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a,
	0x06, 0x2e, 0x2f, 0x73, 0x69, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sink_proto_rawDescOnce sync.Once
	file_sink_proto_rawDescData = file_sink_proto_rawDesc
)

func file_sink_proto_rawDescGZIP() []byte {
	file_sink_proto_rawDescOnce.Do(func() {
		file_sink_proto_rawDescData = protoimpl.X.CompressGZIP(file_sink_proto_rawDescData)
	})
	return file_sink_proto_rawDescData
}

var file_sink_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sink_proto_goTypes = []interface{}{
	(*SinkMessage)(nil), // 0: sink.SinkMessage
	nil,                 // 1: sink.SinkMessage.TracingInfoEntry
}
var file_sink_proto_depIdxs = []int32{
	1, // 0: sink.SinkMessage.tracing_info:type_name -> sink.SinkMessage.TracingInfoEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sink_proto_init() }
func file_sink_proto_init() {
	if File_sink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SinkMessage); i {
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
			RawDescriptor: file_sink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sink_proto_goTypes,
		DependencyIndexes: file_sink_proto_depIdxs,
		MessageInfos:      file_sink_proto_msgTypes,
	}.Build()
	File_sink_proto = out.File
	file_sink_proto_rawDesc = nil
	file_sink_proto_goTypes = nil
	file_sink_proto_depIdxs = nil
}
