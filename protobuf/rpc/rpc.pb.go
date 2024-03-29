// Source: https://github.com/OpenNMS/opennms/blob/master/core/ipc/rpc/kafka/src/main/proto/kafka-rpc.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: rpc.proto

package rpc

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

type RpcMessageProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpcId              string            `protobuf:"bytes,1,opt,name=rpc_id,json=rpcId,proto3" json:"rpc_id,omitempty"`
	RpcContent         []byte            `protobuf:"bytes,2,opt,name=rpc_content,json=rpcContent,proto3" json:"rpc_content,omitempty"`
	SystemId           string            `protobuf:"bytes,3,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	ExpirationTime     uint64            `protobuf:"varint,4,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	CurrentChunkNumber int32             `protobuf:"varint,5,opt,name=current_chunk_number,json=currentChunkNumber,proto3" json:"current_chunk_number,omitempty"`
	TotalChunks        int32             `protobuf:"varint,6,opt,name=total_chunks,json=totalChunks,proto3" json:"total_chunks,omitempty"`
	TracingInfo        map[string]string `protobuf:"bytes,7,rep,name=tracing_info,json=tracingInfo,proto3" json:"tracing_info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ModuleId           string            `protobuf:"bytes,8,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
}

func (x *RpcMessageProto) Reset() {
	*x = RpcMessageProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcMessageProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcMessageProto) ProtoMessage() {}

func (x *RpcMessageProto) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcMessageProto.ProtoReflect.Descriptor instead.
func (*RpcMessageProto) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RpcMessageProto) GetRpcId() string {
	if x != nil {
		return x.RpcId
	}
	return ""
}

func (x *RpcMessageProto) GetRpcContent() []byte {
	if x != nil {
		return x.RpcContent
	}
	return nil
}

func (x *RpcMessageProto) GetSystemId() string {
	if x != nil {
		return x.SystemId
	}
	return ""
}

func (x *RpcMessageProto) GetExpirationTime() uint64 {
	if x != nil {
		return x.ExpirationTime
	}
	return 0
}

func (x *RpcMessageProto) GetCurrentChunkNumber() int32 {
	if x != nil {
		return x.CurrentChunkNumber
	}
	return 0
}

func (x *RpcMessageProto) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

func (x *RpcMessageProto) GetTracingInfo() map[string]string {
	if x != nil {
		return x.TracingInfo
	}
	return nil
}

func (x *RpcMessageProto) GetModuleId() string {
	if x != nil {
		return x.ModuleId
	}
	return ""
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63,
	0x22, 0x8b, 0x03, 0x0a, 0x0f, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x70, 0x63, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x48, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x3e,
	0x0a, 0x10, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x39,
	0x0a, 0x24, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6e, 0x6d, 0x73, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x08, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_proto_goTypes = []interface{}{
	(*RpcMessageProto)(nil), // 0: rpc.RpcMessageProto
	nil,                     // 1: rpc.RpcMessageProto.TracingInfoEntry
}
var file_rpc_proto_depIdxs = []int32{
	1, // 0: rpc.RpcMessageProto.tracing_info:type_name -> rpc.RpcMessageProto.TracingInfoEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcMessageProto); i {
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
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
