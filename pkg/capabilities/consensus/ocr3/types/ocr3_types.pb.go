// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: ocr3_types.proto

package types

import (
	pb "github.com/smartcontractkit/chainlink-common/pkg/values/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// per-workflow aggregation outcome
type AggregationOutcome struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	EncodableOutcome *pb.Map                `protobuf:"bytes,1,opt,name=encodableOutcome,proto3" json:"encodableOutcome,omitempty"` // passed to encoders
	Metadata         []byte                 `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`                 // internal to the aggregator
	ShouldReport     bool                   `protobuf:"varint,3,opt,name=shouldReport,proto3" json:"shouldReport,omitempty"`
	LastSeenAt       uint64                 `protobuf:"varint,4,opt,name=lastSeenAt,proto3" json:"lastSeenAt,omitempty"`      // this is the equivalent of a SeqNr
	Timestamp        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`         // current time of the node
	EncoderName      string                 `protobuf:"bytes,6,opt,name=encoderName,proto3" json:"encoderName,omitempty"`     // optional dynamic encoder override
	EncoderConfig    *pb.Map                `protobuf:"bytes,7,opt,name=encoderConfig,proto3" json:"encoderConfig,omitempty"` // optional dynamic encoder config
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *AggregationOutcome) Reset() {
	*x = AggregationOutcome{}
	mi := &file_ocr3_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregationOutcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregationOutcome) ProtoMessage() {}

func (x *AggregationOutcome) ProtoReflect() protoreflect.Message {
	mi := &file_ocr3_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregationOutcome.ProtoReflect.Descriptor instead.
func (*AggregationOutcome) Descriptor() ([]byte, []int) {
	return file_ocr3_types_proto_rawDescGZIP(), []int{0}
}

func (x *AggregationOutcome) GetEncodableOutcome() *pb.Map {
	if x != nil {
		return x.EncodableOutcome
	}
	return nil
}

func (x *AggregationOutcome) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AggregationOutcome) GetShouldReport() bool {
	if x != nil {
		return x.ShouldReport
	}
	return false
}

func (x *AggregationOutcome) GetLastSeenAt() uint64 {
	if x != nil {
		return x.LastSeenAt
	}
	return 0
}

func (x *AggregationOutcome) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AggregationOutcome) GetEncoderName() string {
	if x != nil {
		return x.EncoderName
	}
	return ""
}

func (x *AggregationOutcome) GetEncoderConfig() *pb.Map {
	if x != nil {
		return x.EncoderConfig
	}
	return nil
}

type Query struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the requests to get consensus on.
	Ids           []*Id `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Query) Reset() {
	*x = Query{}
	mi := &file_ocr3_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_ocr3_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_ocr3_types_proto_rawDescGZIP(), []int{1}
}

func (x *Query) GetIds() []*Id {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Id struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	WorkflowExecutionId      string                 `protobuf:"bytes,1,opt,name=workflowExecutionId,proto3" json:"workflowExecutionId,omitempty"`
	WorkflowId               string                 `protobuf:"bytes,2,opt,name=workflowId,proto3" json:"workflowId,omitempty"`
	WorkflowOwner            string                 `protobuf:"bytes,3,opt,name=workflowOwner,proto3" json:"workflowOwner,omitempty"`
	WorkflowName             string                 `protobuf:"bytes,4,opt,name=workflowName,proto3" json:"workflowName,omitempty"`
	ReportId                 string                 `protobuf:"bytes,6,opt,name=reportId,proto3" json:"reportId,omitempty"`
	WorkflowDonId            uint32                 `protobuf:"varint,7,opt,name=workflowDonId,proto3" json:"workflowDonId,omitempty"`
	WorkflowDonConfigVersion uint32                 `protobuf:"varint,8,opt,name=workflowDonConfigVersion,proto3" json:"workflowDonConfigVersion,omitempty"`
	KeyId                    string                 `protobuf:"bytes,9,opt,name=keyId,proto3" json:"keyId,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *Id) Reset() {
	*x = Id{}
	mi := &file_ocr3_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_ocr3_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_ocr3_types_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetWorkflowExecutionId() string {
	if x != nil {
		return x.WorkflowExecutionId
	}
	return ""
}

func (x *Id) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *Id) GetWorkflowOwner() string {
	if x != nil {
		return x.WorkflowOwner
	}
	return ""
}

func (x *Id) GetWorkflowName() string {
	if x != nil {
		return x.WorkflowName
	}
	return ""
}

func (x *Id) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

func (x *Id) GetWorkflowDonId() uint32 {
	if x != nil {
		return x.WorkflowDonId
	}
	return 0
}

func (x *Id) GetWorkflowDonConfigVersion() uint32 {
	if x != nil {
		return x.WorkflowDonConfigVersion
	}
	return 0
}

func (x *Id) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type Observation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    *Id                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// list of observations defined in inputs.observations
	Observations            *pb.List `protobuf:"bytes,4,opt,name=observations,proto3" json:"observations,omitempty"`
	OverriddenEncoderName   string   `protobuf:"bytes,5,opt,name=overriddenEncoderName,proto3" json:"overriddenEncoderName,omitempty"`
	OverriddenEncoderConfig *pb.Map  `protobuf:"bytes,6,opt,name=overriddenEncoderConfig,proto3" json:"overriddenEncoderConfig,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Observation) Reset() {
	*x = Observation{}
	mi := &file_ocr3_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Observation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observation) ProtoMessage() {}

func (x *Observation) ProtoReflect() protoreflect.Message {
	mi := &file_ocr3_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observation.ProtoReflect.Descriptor instead.
func (*Observation) Descriptor() ([]byte, []int) {
	return file_ocr3_types_proto_rawDescGZIP(), []int{3}
}

func (x *Observation) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Observation) GetObservations() *pb.List {
	if x != nil {
		return x.Observations
	}
	return nil
}

func (x *Observation) GetOverriddenEncoderName() string {
	if x != nil {
		return x.OverriddenEncoderName
	}
	return ""
}

func (x *Observation) GetOverriddenEncoderConfig() *pb.Map {
	if x != nil {
		return x.OverriddenEncoderConfig
	}
	return nil
}

type Observations struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// batched observations for multiple workflow execution IDs
	Observations []*Observation `protobuf:"bytes,1,rep,name=observations,proto3" json:"observations,omitempty"`
	// the workflow IDs that are registered in the node
	RegisteredWorkflowIds []string `protobuf:"bytes,2,rep,name=registeredWorkflowIds,proto3" json:"registeredWorkflowIds,omitempty"`
	// the node's current current time
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Observations) Reset() {
	*x = Observations{}
	mi := &file_ocr3_types_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Observations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observations) ProtoMessage() {}

func (x *Observations) ProtoReflect() protoreflect.Message {
	mi := &file_ocr3_types_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observations.ProtoReflect.Descriptor instead.
func (*Observations) Descriptor() ([]byte, []int) {
	return file_ocr3_types_proto_rawDescGZIP(), []int{4}
}

func (x *Observations) GetObservations() []*Observation {
	if x != nil {
		return x.Observations
	}
	return nil
}

func (x *Observations) GetRegisteredWorkflowIds() []string {
	if x != nil {
		return x.RegisteredWorkflowIds
	}
	return nil
}

func (x *Observations) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Report struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *Id                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Outcome       *AggregationOutcome    `protobuf:"bytes,2,opt,name=outcome,proto3" json:"outcome,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Report) Reset() {
	*x = Report{}
	mi := &file_ocr3_types_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_ocr3_types_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_ocr3_types_proto_rawDescGZIP(), []int{5}
}

func (x *Report) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Report) GetOutcome() *AggregationOutcome {
	if x != nil {
		return x.Outcome
	}
	return nil
}

type ReportInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *Id                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShouldReport  bool                   `protobuf:"varint,2,opt,name=should_report,json=shouldReport,proto3" json:"should_report,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReportInfo) Reset() {
	*x = ReportInfo{}
	mi := &file_ocr3_types_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportInfo) ProtoMessage() {}

func (x *ReportInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ocr3_types_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportInfo.ProtoReflect.Descriptor instead.
func (*ReportInfo) Descriptor() ([]byte, []int) {
	return file_ocr3_types_proto_rawDescGZIP(), []int{6}
}

func (x *ReportInfo) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ReportInfo) GetShouldReport() bool {
	if x != nil {
		return x.ShouldReport
	}
	return false
}

type Outcome struct {
	state          protoimpl.MessageState         `protogen:"open.v1"`
	Outcomes       map[string]*AggregationOutcome `protobuf:"bytes,1,rep,name=outcomes,proto3" json:"outcomes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CurrentReports []*Report                      `protobuf:"bytes,2,rep,name=current_reports,json=currentReports,proto3" json:"current_reports,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	mi := &file_ocr3_types_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_ocr3_types_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_ocr3_types_proto_rawDescGZIP(), []int{7}
}

func (x *Outcome) GetOutcomes() map[string]*AggregationOutcome {
	if x != nil {
		return x.Outcomes
	}
	return nil
}

func (x *Outcome) GetCurrentReports() []*Report {
	if x != nil {
		return x.CurrentReports
	}
	return nil
}

var File_ocr3_types_proto protoreflect.FileDescriptor

const file_ocr3_types_proto_rawDesc = "" +
	"\n" +
	"\x10ocr3_types.proto\x12\n" +
	"ocr3_types\x1a\x16values/v1/values.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc2\x02\n" +
	"\x12AggregationOutcome\x12:\n" +
	"\x10encodableOutcome\x18\x01 \x01(\v2\x0e.values.v1.MapR\x10encodableOutcome\x12\x1a\n" +
	"\bmetadata\x18\x02 \x01(\fR\bmetadata\x12\"\n" +
	"\fshouldReport\x18\x03 \x01(\bR\fshouldReport\x12\x1e\n" +
	"\n" +
	"lastSeenAt\x18\x04 \x01(\x04R\n" +
	"lastSeenAt\x128\n" +
	"\ttimestamp\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12 \n" +
	"\vencoderName\x18\x06 \x01(\tR\vencoderName\x124\n" +
	"\rencoderConfig\x18\a \x01(\v2\x0e.values.v1.MapR\rencoderConfig\")\n" +
	"\x05Query\x12 \n" +
	"\x03ids\x18\x01 \x03(\v2\x0e.ocr3_types.IdR\x03ids\"\xba\x02\n" +
	"\x02Id\x120\n" +
	"\x13workflowExecutionId\x18\x01 \x01(\tR\x13workflowExecutionId\x12\x1e\n" +
	"\n" +
	"workflowId\x18\x02 \x01(\tR\n" +
	"workflowId\x12$\n" +
	"\rworkflowOwner\x18\x03 \x01(\tR\rworkflowOwner\x12\"\n" +
	"\fworkflowName\x18\x04 \x01(\tR\fworkflowName\x12\x1a\n" +
	"\breportId\x18\x06 \x01(\tR\breportId\x12$\n" +
	"\rworkflowDonId\x18\a \x01(\rR\rworkflowDonId\x12:\n" +
	"\x18workflowDonConfigVersion\x18\b \x01(\rR\x18workflowDonConfigVersion\x12\x14\n" +
	"\x05keyId\x18\t \x01(\tR\x05keyIdJ\x04\b\x05\x10\x06\"\xe2\x01\n" +
	"\vObservation\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\v2\x0e.ocr3_types.IdR\x02id\x123\n" +
	"\fobservations\x18\x04 \x01(\v2\x0f.values.v1.ListR\fobservations\x124\n" +
	"\x15overriddenEncoderName\x18\x05 \x01(\tR\x15overriddenEncoderName\x12H\n" +
	"\x17overriddenEncoderConfig\x18\x06 \x01(\v2\x0e.values.v1.MapR\x17overriddenEncoderConfig\"\xbb\x01\n" +
	"\fObservations\x12;\n" +
	"\fobservations\x18\x01 \x03(\v2\x17.ocr3_types.ObservationR\fobservations\x124\n" +
	"\x15registeredWorkflowIds\x18\x02 \x03(\tR\x15registeredWorkflowIds\x128\n" +
	"\ttimestamp\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\"b\n" +
	"\x06Report\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\v2\x0e.ocr3_types.IdR\x02id\x128\n" +
	"\aoutcome\x18\x02 \x01(\v2\x1e.ocr3_types.AggregationOutcomeR\aoutcome\"Q\n" +
	"\n" +
	"ReportInfo\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\v2\x0e.ocr3_types.IdR\x02id\x12#\n" +
	"\rshould_report\x18\x02 \x01(\bR\fshouldReport\"\xe2\x01\n" +
	"\aOutcome\x12=\n" +
	"\boutcomes\x18\x01 \x03(\v2!.ocr3_types.Outcome.OutcomesEntryR\boutcomes\x12;\n" +
	"\x0fcurrent_reports\x18\x02 \x03(\v2\x12.ocr3_types.ReportR\x0ecurrentReports\x1a[\n" +
	"\rOutcomesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x124\n" +
	"\x05value\x18\x02 \x01(\v2\x1e.ocr3_types.AggregationOutcomeR\x05value:\x028\x01B#Z!capabilities/consensus/ocr3/typesb\x06proto3"

var (
	file_ocr3_types_proto_rawDescOnce sync.Once
	file_ocr3_types_proto_rawDescData []byte
)

func file_ocr3_types_proto_rawDescGZIP() []byte {
	file_ocr3_types_proto_rawDescOnce.Do(func() {
		file_ocr3_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ocr3_types_proto_rawDesc), len(file_ocr3_types_proto_rawDesc)))
	})
	return file_ocr3_types_proto_rawDescData
}

var file_ocr3_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ocr3_types_proto_goTypes = []any{
	(*AggregationOutcome)(nil),    // 0: ocr3_types.AggregationOutcome
	(*Query)(nil),                 // 1: ocr3_types.Query
	(*Id)(nil),                    // 2: ocr3_types.Id
	(*Observation)(nil),           // 3: ocr3_types.Observation
	(*Observations)(nil),          // 4: ocr3_types.Observations
	(*Report)(nil),                // 5: ocr3_types.Report
	(*ReportInfo)(nil),            // 6: ocr3_types.ReportInfo
	(*Outcome)(nil),               // 7: ocr3_types.Outcome
	nil,                           // 8: ocr3_types.Outcome.OutcomesEntry
	(*pb.Map)(nil),                // 9: values.v1.Map
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*pb.List)(nil),               // 11: values.v1.List
}
var file_ocr3_types_proto_depIdxs = []int32{
	9,  // 0: ocr3_types.AggregationOutcome.encodableOutcome:type_name -> values.v1.Map
	10, // 1: ocr3_types.AggregationOutcome.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 2: ocr3_types.AggregationOutcome.encoderConfig:type_name -> values.v1.Map
	2,  // 3: ocr3_types.Query.ids:type_name -> ocr3_types.Id
	2,  // 4: ocr3_types.Observation.id:type_name -> ocr3_types.Id
	11, // 5: ocr3_types.Observation.observations:type_name -> values.v1.List
	9,  // 6: ocr3_types.Observation.overriddenEncoderConfig:type_name -> values.v1.Map
	3,  // 7: ocr3_types.Observations.observations:type_name -> ocr3_types.Observation
	10, // 8: ocr3_types.Observations.timestamp:type_name -> google.protobuf.Timestamp
	2,  // 9: ocr3_types.Report.id:type_name -> ocr3_types.Id
	0,  // 10: ocr3_types.Report.outcome:type_name -> ocr3_types.AggregationOutcome
	2,  // 11: ocr3_types.ReportInfo.id:type_name -> ocr3_types.Id
	8,  // 12: ocr3_types.Outcome.outcomes:type_name -> ocr3_types.Outcome.OutcomesEntry
	5,  // 13: ocr3_types.Outcome.current_reports:type_name -> ocr3_types.Report
	0,  // 14: ocr3_types.Outcome.OutcomesEntry.value:type_name -> ocr3_types.AggregationOutcome
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_ocr3_types_proto_init() }
func file_ocr3_types_proto_init() {
	if File_ocr3_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ocr3_types_proto_rawDesc), len(file_ocr3_types_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ocr3_types_proto_goTypes,
		DependencyIndexes: file_ocr3_types_proto_depIdxs,
		MessageInfos:      file_ocr3_types_proto_msgTypes,
	}.Build()
	File_ocr3_types_proto = out.File
	file_ocr3_types_proto_goTypes = nil
	file_ocr3_types_proto_depIdxs = nil
}
