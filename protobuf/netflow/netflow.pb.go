// Source: https://github.com/OpenNMS/opennms/blob/master/features/telemetry/protocols/netflow/transport/src/main/proto/netflow.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: netflow.proto

package netflow

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Direction int32

const (
	Direction_INGRESS Direction = 0
	Direction_EGRESS  Direction = 1
	Direction_UNKNOWN Direction = 255
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0:   "INGRESS",
		1:   "EGRESS",
		255: "UNKNOWN",
	}
	Direction_value = map[string]int32{
		"INGRESS": 0,
		"EGRESS":  1,
		"UNKNOWN": 255,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_netflow_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_netflow_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_netflow_proto_rawDescGZIP(), []int{0}
}

type SamplingAlgorithm int32

const (
	SamplingAlgorithm_UNASSIGNED                                               SamplingAlgorithm = 0
	SamplingAlgorithm_SYSTEMATIC_COUNT_BASED_SAMPLING                          SamplingAlgorithm = 1
	SamplingAlgorithm_SYSTEMATIC_TIME_BASED_SAMPLING                           SamplingAlgorithm = 2
	SamplingAlgorithm_RANDOM_N_OUT_OF_N_SAMPLING                               SamplingAlgorithm = 3
	SamplingAlgorithm_UNIFORM_PROBABILISTIC_SAMPLING                           SamplingAlgorithm = 4
	SamplingAlgorithm_PROPERTY_MATCH_FILTERING                                 SamplingAlgorithm = 5
	SamplingAlgorithm_HASH_BASED_FILTERING                                     SamplingAlgorithm = 6
	SamplingAlgorithm_FLOW_STATE_DEPENDENT_INTERMEDIATE_FLOW_SELECTION_PROCESS SamplingAlgorithm = 7
)

// Enum value maps for SamplingAlgorithm.
var (
	SamplingAlgorithm_name = map[int32]string{
		0: "UNASSIGNED",
		1: "SYSTEMATIC_COUNT_BASED_SAMPLING",
		2: "SYSTEMATIC_TIME_BASED_SAMPLING",
		3: "RANDOM_N_OUT_OF_N_SAMPLING",
		4: "UNIFORM_PROBABILISTIC_SAMPLING",
		5: "PROPERTY_MATCH_FILTERING",
		6: "HASH_BASED_FILTERING",
		7: "FLOW_STATE_DEPENDENT_INTERMEDIATE_FLOW_SELECTION_PROCESS",
	}
	SamplingAlgorithm_value = map[string]int32{
		"UNASSIGNED":                                               0,
		"SYSTEMATIC_COUNT_BASED_SAMPLING":                          1,
		"SYSTEMATIC_TIME_BASED_SAMPLING":                           2,
		"RANDOM_N_OUT_OF_N_SAMPLING":                               3,
		"UNIFORM_PROBABILISTIC_SAMPLING":                           4,
		"PROPERTY_MATCH_FILTERING":                                 5,
		"HASH_BASED_FILTERING":                                     6,
		"FLOW_STATE_DEPENDENT_INTERMEDIATE_FLOW_SELECTION_PROCESS": 7,
	}
)

func (x SamplingAlgorithm) Enum() *SamplingAlgorithm {
	p := new(SamplingAlgorithm)
	*p = x
	return p
}

func (x SamplingAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SamplingAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_netflow_proto_enumTypes[1].Descriptor()
}

func (SamplingAlgorithm) Type() protoreflect.EnumType {
	return &file_netflow_proto_enumTypes[1]
}

func (x SamplingAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SamplingAlgorithm.Descriptor instead.
func (SamplingAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_netflow_proto_rawDescGZIP(), []int{1}
}

type NetflowVersion int32

const (
	NetflowVersion_V5    NetflowVersion = 0
	NetflowVersion_V9    NetflowVersion = 1
	NetflowVersion_IPFIX NetflowVersion = 2
)

// Enum value maps for NetflowVersion.
var (
	NetflowVersion_name = map[int32]string{
		0: "V5",
		1: "V9",
		2: "IPFIX",
	}
	NetflowVersion_value = map[string]int32{
		"V5":    0,
		"V9":    1,
		"IPFIX": 2,
	}
)

func (x NetflowVersion) Enum() *NetflowVersion {
	p := new(NetflowVersion)
	*p = x
	return p
}

func (x NetflowVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetflowVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_netflow_proto_enumTypes[2].Descriptor()
}

func (NetflowVersion) Type() protoreflect.EnumType {
	return &file_netflow_proto_enumTypes[2]
}

func (x NetflowVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetflowVersion.Descriptor instead.
func (NetflowVersion) EnumDescriptor() ([]byte, []int) {
	return file_netflow_proto_rawDescGZIP(), []int{2}
}

type FlowMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp     uint64                  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                              // Flow timestamp in milliseconds.
	NumBytes      *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=num_bytes,json=numBytes,proto3" json:"num_bytes,omitempty"`                 // Number of bytes transferred in the flow
	Direction     Direction               `protobuf:"varint,3,opt,name=direction,proto3,enum=netflow.Direction" json:"direction,omitempty"`       // Direction of the flow (egress vs ingress)
	DstAddress    string                  `protobuf:"bytes,4,opt,name=dst_address,json=dstAddress,proto3" json:"dst_address,omitempty"`           //  Destination address.
	DstHostname   string                  `protobuf:"bytes,5,opt,name=dst_hostname,json=dstHostname,proto3" json:"dst_hostname,omitempty"`        // Destination address hostname.
	DstAs         *wrapperspb.UInt64Value `protobuf:"bytes,6,opt,name=dst_as,json=dstAs,proto3" json:"dst_as,omitempty"`                          // Destination autonomous system (AS).
	DstMaskLen    *wrapperspb.UInt32Value `protobuf:"bytes,7,opt,name=dst_mask_len,json=dstMaskLen,proto3" json:"dst_mask_len,omitempty"`         // The number of contiguous bits in the source address subnet mask.
	DstPort       *wrapperspb.UInt32Value `protobuf:"bytes,8,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`                    // Destination port.
	EngineId      *wrapperspb.UInt32Value `protobuf:"bytes,9,opt,name=engine_id,json=engineId,proto3" json:"engine_id,omitempty"`                 // Slot number of the flow-switching engine.
	EngineType    *wrapperspb.UInt32Value `protobuf:"bytes,10,opt,name=engine_type,json=engineType,proto3" json:"engine_type,omitempty"`          // Type of flow-switching engine.
	DeltaSwitched *wrapperspb.UInt64Value `protobuf:"bytes,11,opt,name=delta_switched,json=deltaSwitched,proto3" json:"delta_switched,omitempty"` // Unix timestamp in ms at which the previous exported packet-
	// -associated with this flow was switched.
	FirstSwitched *wrapperspb.UInt64Value `protobuf:"bytes,12,opt,name=first_switched,json=firstSwitched,proto3" json:"first_switched,omitempty"` // Unix timestamp in ms at which the first packet-
	// -associated with this flow was switched.
	LastSwitched      *wrapperspb.UInt64Value `protobuf:"bytes,13,opt,name=last_switched,json=lastSwitched,proto3" json:"last_switched,omitempty"`
	NumFlowRecords    *wrapperspb.UInt32Value `protobuf:"bytes,14,opt,name=num_flow_records,json=numFlowRecords,proto3" json:"num_flow_records,omitempty"`                                        // Number of flow records in the associated packet.
	NumPackets        *wrapperspb.UInt64Value `protobuf:"bytes,15,opt,name=num_packets,json=numPackets,proto3" json:"num_packets,omitempty"`                                                      // Number of packets in the flow.
	FlowSeqNum        *wrapperspb.UInt64Value `protobuf:"bytes,16,opt,name=flow_seq_num,json=flowSeqNum,proto3" json:"flow_seq_num,omitempty"`                                                    // Flow packet sequence number.
	InputSnmpIfindex  *wrapperspb.UInt32Value `protobuf:"bytes,17,opt,name=input_snmp_ifindex,json=inputSnmpIfindex,proto3" json:"input_snmp_ifindex,omitempty"`                                  // Input SNMP ifIndex.
	OutputSnmpIfindex *wrapperspb.UInt32Value `protobuf:"bytes,18,opt,name=output_snmp_ifindex,json=outputSnmpIfindex,proto3" json:"output_snmp_ifindex,omitempty"`                               // Output SNMP ifIndex.
	IpProtocolVersion *wrapperspb.UInt32Value `protobuf:"bytes,19,opt,name=ip_protocol_version,json=ipProtocolVersion,proto3" json:"ip_protocol_version,omitempty"`                               // IPv4 vs IPv6.
	NextHopAddress    string                  `protobuf:"bytes,20,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`                                        // Next hop IpAddress.
	NextHopHostname   string                  `protobuf:"bytes,21,opt,name=next_hop_hostname,json=nextHopHostname,proto3" json:"next_hop_hostname,omitempty"`                                     // Next hop hostname.
	Protocol          *wrapperspb.UInt32Value `protobuf:"bytes,22,opt,name=protocol,proto3" json:"protocol,omitempty"`                                                                            // IP protocol number i.e 6 for TCP, 17 for UDP
	SamplingAlgorithm SamplingAlgorithm       `protobuf:"varint,23,opt,name=sampling_algorithm,json=samplingAlgorithm,proto3,enum=netflow.SamplingAlgorithm" json:"sampling_algorithm,omitempty"` // Sampling algorithm ID.
	SamplingInterval  *wrapperspb.DoubleValue `protobuf:"bytes,24,opt,name=sampling_interval,json=samplingInterval,proto3" json:"sampling_interval,omitempty"`                                    // Sampling interval.
	SrcAddress        string                  `protobuf:"bytes,26,opt,name=src_address,json=srcAddress,proto3" json:"src_address,omitempty"`                                                      // Source address.
	SrcHostname       string                  `protobuf:"bytes,27,opt,name=src_hostname,json=srcHostname,proto3" json:"src_hostname,omitempty"`                                                   // Source hostname.
	SrcAs             *wrapperspb.UInt64Value `protobuf:"bytes,28,opt,name=src_as,json=srcAs,proto3" json:"src_as,omitempty"`                                                                     // Source AS number.
	SrcMaskLen        *wrapperspb.UInt32Value `protobuf:"bytes,29,opt,name=src_mask_len,json=srcMaskLen,proto3" json:"src_mask_len,omitempty"`                                                    // The number of contiguous bits in the destination address subnet mask.
	SrcPort           *wrapperspb.UInt32Value `protobuf:"bytes,30,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`                                                               // Source port.
	TcpFlags          *wrapperspb.UInt32Value `protobuf:"bytes,31,opt,name=tcp_flags,json=tcpFlags,proto3" json:"tcp_flags,omitempty"`                                                            // TCP Flags.
	Tos               *wrapperspb.UInt32Value `protobuf:"bytes,32,opt,name=tos,proto3" json:"tos,omitempty"`                                                                                      // TOS
	NetflowVersion    NetflowVersion          `protobuf:"varint,33,opt,name=netflow_version,json=netflowVersion,proto3,enum=netflow.NetflowVersion" json:"netflow_version,omitempty"`             // Netflow version
	Vlan              *wrapperspb.UInt32Value `protobuf:"bytes,34,opt,name=vlan,proto3" json:"vlan,omitempty"`                                                                                    // VLAN ID.
	NodeIdentifier    string                  `protobuf:"bytes,35,opt,name=node_identifier,json=nodeIdentifier,proto3" json:"node_identifier,omitempty"`                                          // node lookup identifier.
}

func (x *FlowMessage) Reset() {
	*x = FlowMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowMessage) ProtoMessage() {}

func (x *FlowMessage) ProtoReflect() protoreflect.Message {
	mi := &file_netflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowMessage.ProtoReflect.Descriptor instead.
func (*FlowMessage) Descriptor() ([]byte, []int) {
	return file_netflow_proto_rawDescGZIP(), []int{0}
}

func (x *FlowMessage) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *FlowMessage) GetNumBytes() *wrapperspb.UInt64Value {
	if x != nil {
		return x.NumBytes
	}
	return nil
}

func (x *FlowMessage) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_INGRESS
}

func (x *FlowMessage) GetDstAddress() string {
	if x != nil {
		return x.DstAddress
	}
	return ""
}

func (x *FlowMessage) GetDstHostname() string {
	if x != nil {
		return x.DstHostname
	}
	return ""
}

func (x *FlowMessage) GetDstAs() *wrapperspb.UInt64Value {
	if x != nil {
		return x.DstAs
	}
	return nil
}

func (x *FlowMessage) GetDstMaskLen() *wrapperspb.UInt32Value {
	if x != nil {
		return x.DstMaskLen
	}
	return nil
}

func (x *FlowMessage) GetDstPort() *wrapperspb.UInt32Value {
	if x != nil {
		return x.DstPort
	}
	return nil
}

func (x *FlowMessage) GetEngineId() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EngineId
	}
	return nil
}

func (x *FlowMessage) GetEngineType() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EngineType
	}
	return nil
}

func (x *FlowMessage) GetDeltaSwitched() *wrapperspb.UInt64Value {
	if x != nil {
		return x.DeltaSwitched
	}
	return nil
}

func (x *FlowMessage) GetFirstSwitched() *wrapperspb.UInt64Value {
	if x != nil {
		return x.FirstSwitched
	}
	return nil
}

func (x *FlowMessage) GetLastSwitched() *wrapperspb.UInt64Value {
	if x != nil {
		return x.LastSwitched
	}
	return nil
}

func (x *FlowMessage) GetNumFlowRecords() *wrapperspb.UInt32Value {
	if x != nil {
		return x.NumFlowRecords
	}
	return nil
}

func (x *FlowMessage) GetNumPackets() *wrapperspb.UInt64Value {
	if x != nil {
		return x.NumPackets
	}
	return nil
}

func (x *FlowMessage) GetFlowSeqNum() *wrapperspb.UInt64Value {
	if x != nil {
		return x.FlowSeqNum
	}
	return nil
}

func (x *FlowMessage) GetInputSnmpIfindex() *wrapperspb.UInt32Value {
	if x != nil {
		return x.InputSnmpIfindex
	}
	return nil
}

func (x *FlowMessage) GetOutputSnmpIfindex() *wrapperspb.UInt32Value {
	if x != nil {
		return x.OutputSnmpIfindex
	}
	return nil
}

func (x *FlowMessage) GetIpProtocolVersion() *wrapperspb.UInt32Value {
	if x != nil {
		return x.IpProtocolVersion
	}
	return nil
}

func (x *FlowMessage) GetNextHopAddress() string {
	if x != nil {
		return x.NextHopAddress
	}
	return ""
}

func (x *FlowMessage) GetNextHopHostname() string {
	if x != nil {
		return x.NextHopHostname
	}
	return ""
}

func (x *FlowMessage) GetProtocol() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Protocol
	}
	return nil
}

func (x *FlowMessage) GetSamplingAlgorithm() SamplingAlgorithm {
	if x != nil {
		return x.SamplingAlgorithm
	}
	return SamplingAlgorithm_UNASSIGNED
}

func (x *FlowMessage) GetSamplingInterval() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SamplingInterval
	}
	return nil
}

func (x *FlowMessage) GetSrcAddress() string {
	if x != nil {
		return x.SrcAddress
	}
	return ""
}

func (x *FlowMessage) GetSrcHostname() string {
	if x != nil {
		return x.SrcHostname
	}
	return ""
}

func (x *FlowMessage) GetSrcAs() *wrapperspb.UInt64Value {
	if x != nil {
		return x.SrcAs
	}
	return nil
}

func (x *FlowMessage) GetSrcMaskLen() *wrapperspb.UInt32Value {
	if x != nil {
		return x.SrcMaskLen
	}
	return nil
}

func (x *FlowMessage) GetSrcPort() *wrapperspb.UInt32Value {
	if x != nil {
		return x.SrcPort
	}
	return nil
}

func (x *FlowMessage) GetTcpFlags() *wrapperspb.UInt32Value {
	if x != nil {
		return x.TcpFlags
	}
	return nil
}

func (x *FlowMessage) GetTos() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Tos
	}
	return nil
}

func (x *FlowMessage) GetNetflowVersion() NetflowVersion {
	if x != nil {
		return x.NetflowVersion
	}
	return NetflowVersion_V5
}

func (x *FlowMessage) GetVlan() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Vlan
	}
	return nil
}

func (x *FlowMessage) GetNodeIdentifier() string {
	if x != nil {
		return x.NodeIdentifier
	}
	return ""
}

var File_netflow_proto protoreflect.FileDescriptor

var file_netflow_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x0f, 0x0a, 0x0b, 0x46, 0x6c, 0x6f,
	0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x39, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x30, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x73, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x73, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x73, 0x74, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x64, 0x73, 0x74, 0x5f, 0x61,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x64, 0x73, 0x74, 0x41, 0x73, 0x12, 0x3e, 0x0a, 0x0c,
	0x64, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0a, 0x64, 0x73, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x4c, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x08,
	0x64, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x64, 0x73,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x0b, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x43, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x64, 0x12, 0x43, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x10,
	0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x71, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x71,
	0x4e, 0x75, 0x6d, 0x12, 0x4a, 0x0a, 0x12, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x6e, 0x6d,
	0x70, 0x5f, 0x69, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x53, 0x6e, 0x6d, 0x70, 0x49, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x4c, 0x0a, 0x13, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x69,
	0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x53, 0x6e, 0x6d, 0x70, 0x49, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x4c, 0x0a,
	0x13, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f,
	0x70, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x49, 0x0a, 0x12, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x52, 0x11, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x49, 0x0a, 0x11, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x10, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x72, 0x63, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x72, 0x63, 0x48, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x73, 0x18,
	0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x73, 0x72, 0x63, 0x41, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x73, 0x72,
	0x63, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a,
	0x73, 0x72, 0x63, 0x4d, 0x61, 0x73, 0x6b, 0x4c, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x72,
	0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x74, 0x63, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x74, 0x63, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x2e,
	0x0a, 0x03, 0x74, 0x6f, 0x73, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x74, 0x6f, 0x73, 0x12, 0x40,
	0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x4e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0e, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x30, 0x0a, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x76, 0x6c,
	0x61, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2a, 0x32, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x47, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0xff, 0x01, 0x2a,
	0xa6, 0x02, 0x0a, 0x11, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x41, 0x53, 0x53, 0x49, 0x47,
	0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x41,
	0x54, 0x49, 0x43, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x44, 0x5f,
	0x53, 0x41, 0x4d, 0x50, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x59,
	0x53, 0x54, 0x45, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x42, 0x41,
	0x53, 0x45, 0x44, 0x5f, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x1e,
	0x0a, 0x1a, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x5f, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f,
	0x46, 0x5f, 0x4e, 0x5f, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x22,
	0x0a, 0x1e, 0x55, 0x4e, 0x49, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x43, 0x5f, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x49, 0x4e, 0x47,
	0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x4d,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x05,
	0x12, 0x18, 0x0a, 0x14, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x44, 0x5f, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x3c, 0x0a, 0x38, 0x46, 0x4c,
	0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x45, 0x4e, 0x44, 0x45,
	0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x5f,
	0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x07, 0x2a, 0x2b, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x66,
	0x6c, 0x6f, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x35,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x39, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x50,
	0x46, 0x49, 0x58, 0x10, 0x02, 0x42, 0x53, 0x0a, 0x38, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x6e, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x74, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x42, 0x0a, 0x46, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a,
	0x09, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_netflow_proto_rawDescOnce sync.Once
	file_netflow_proto_rawDescData = file_netflow_proto_rawDesc
)

func file_netflow_proto_rawDescGZIP() []byte {
	file_netflow_proto_rawDescOnce.Do(func() {
		file_netflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_netflow_proto_rawDescData)
	})
	return file_netflow_proto_rawDescData
}

var file_netflow_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_netflow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_netflow_proto_goTypes = []interface{}{
	(Direction)(0),                 // 0: netflow.Direction
	(SamplingAlgorithm)(0),         // 1: netflow.SamplingAlgorithm
	(NetflowVersion)(0),            // 2: netflow.NetflowVersion
	(*FlowMessage)(nil),            // 3: netflow.FlowMessage
	(*wrapperspb.UInt64Value)(nil), // 4: google.protobuf.UInt64Value
	(*wrapperspb.UInt32Value)(nil), // 5: google.protobuf.UInt32Value
	(*wrapperspb.DoubleValue)(nil), // 6: google.protobuf.DoubleValue
}
var file_netflow_proto_depIdxs = []int32{
	4,  // 0: netflow.FlowMessage.num_bytes:type_name -> google.protobuf.UInt64Value
	0,  // 1: netflow.FlowMessage.direction:type_name -> netflow.Direction
	4,  // 2: netflow.FlowMessage.dst_as:type_name -> google.protobuf.UInt64Value
	5,  // 3: netflow.FlowMessage.dst_mask_len:type_name -> google.protobuf.UInt32Value
	5,  // 4: netflow.FlowMessage.dst_port:type_name -> google.protobuf.UInt32Value
	5,  // 5: netflow.FlowMessage.engine_id:type_name -> google.protobuf.UInt32Value
	5,  // 6: netflow.FlowMessage.engine_type:type_name -> google.protobuf.UInt32Value
	4,  // 7: netflow.FlowMessage.delta_switched:type_name -> google.protobuf.UInt64Value
	4,  // 8: netflow.FlowMessage.first_switched:type_name -> google.protobuf.UInt64Value
	4,  // 9: netflow.FlowMessage.last_switched:type_name -> google.protobuf.UInt64Value
	5,  // 10: netflow.FlowMessage.num_flow_records:type_name -> google.protobuf.UInt32Value
	4,  // 11: netflow.FlowMessage.num_packets:type_name -> google.protobuf.UInt64Value
	4,  // 12: netflow.FlowMessage.flow_seq_num:type_name -> google.protobuf.UInt64Value
	5,  // 13: netflow.FlowMessage.input_snmp_ifindex:type_name -> google.protobuf.UInt32Value
	5,  // 14: netflow.FlowMessage.output_snmp_ifindex:type_name -> google.protobuf.UInt32Value
	5,  // 15: netflow.FlowMessage.ip_protocol_version:type_name -> google.protobuf.UInt32Value
	5,  // 16: netflow.FlowMessage.protocol:type_name -> google.protobuf.UInt32Value
	1,  // 17: netflow.FlowMessage.sampling_algorithm:type_name -> netflow.SamplingAlgorithm
	6,  // 18: netflow.FlowMessage.sampling_interval:type_name -> google.protobuf.DoubleValue
	4,  // 19: netflow.FlowMessage.src_as:type_name -> google.protobuf.UInt64Value
	5,  // 20: netflow.FlowMessage.src_mask_len:type_name -> google.protobuf.UInt32Value
	5,  // 21: netflow.FlowMessage.src_port:type_name -> google.protobuf.UInt32Value
	5,  // 22: netflow.FlowMessage.tcp_flags:type_name -> google.protobuf.UInt32Value
	5,  // 23: netflow.FlowMessage.tos:type_name -> google.protobuf.UInt32Value
	2,  // 24: netflow.FlowMessage.netflow_version:type_name -> netflow.NetflowVersion
	5,  // 25: netflow.FlowMessage.vlan:type_name -> google.protobuf.UInt32Value
	26, // [26:26] is the sub-list for method output_type
	26, // [26:26] is the sub-list for method input_type
	26, // [26:26] is the sub-list for extension type_name
	26, // [26:26] is the sub-list for extension extendee
	0,  // [0:26] is the sub-list for field type_name
}

func init() { file_netflow_proto_init() }
func file_netflow_proto_init() {
	if File_netflow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_netflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowMessage); i {
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
			RawDescriptor: file_netflow_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_netflow_proto_goTypes,
		DependencyIndexes: file_netflow_proto_depIdxs,
		EnumInfos:         file_netflow_proto_enumTypes,
		MessageInfos:      file_netflow_proto_msgTypes,
	}.Build()
	File_netflow_proto = out.File
	file_netflow_proto_rawDesc = nil
	file_netflow_proto_goTypes = nil
	file_netflow_proto_depIdxs = nil
}
