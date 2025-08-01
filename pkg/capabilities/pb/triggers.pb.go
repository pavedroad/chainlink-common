// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: triggers.proto

package pb

import (
	pb "github.com/smartcontractkit/chainlink-common/pkg/values/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type OCRTriggerEvent struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	ConfigDigest  []byte                           `protobuf:"bytes,1,opt,name=configDigest,proto3" json:"configDigest,omitempty"`
	SeqNr         uint64                           `protobuf:"varint,2,opt,name=seqNr,proto3" json:"seqNr,omitempty"`
	Report        []byte                           `protobuf:"bytes,3,opt,name=report,proto3" json:"report,omitempty"` // marshalled OCRTriggerReport
	Sigs          []*OCRAttributedOnchainSignature `protobuf:"bytes,4,rep,name=sigs,proto3" json:"sigs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OCRTriggerEvent) Reset() {
	*x = OCRTriggerEvent{}
	mi := &file_triggers_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OCRTriggerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCRTriggerEvent) ProtoMessage() {}

func (x *OCRTriggerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_triggers_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCRTriggerEvent.ProtoReflect.Descriptor instead.
func (*OCRTriggerEvent) Descriptor() ([]byte, []int) {
	return file_triggers_proto_rawDescGZIP(), []int{0}
}

func (x *OCRTriggerEvent) GetConfigDigest() []byte {
	if x != nil {
		return x.ConfigDigest
	}
	return nil
}

func (x *OCRTriggerEvent) GetSeqNr() uint64 {
	if x != nil {
		return x.SeqNr
	}
	return 0
}

func (x *OCRTriggerEvent) GetReport() []byte {
	if x != nil {
		return x.Report
	}
	return nil
}

func (x *OCRTriggerEvent) GetSigs() []*OCRAttributedOnchainSignature {
	if x != nil {
		return x.Sigs
	}
	return nil
}

type OCRAttributedOnchainSignature struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Signature []byte                 `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// signer is actually a uint8 but uint32 is the smallest supported by protobuf
	Signer        uint32 `protobuf:"varint,2,opt,name=signer,proto3" json:"signer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OCRAttributedOnchainSignature) Reset() {
	*x = OCRAttributedOnchainSignature{}
	mi := &file_triggers_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OCRAttributedOnchainSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCRAttributedOnchainSignature) ProtoMessage() {}

func (x *OCRAttributedOnchainSignature) ProtoReflect() protoreflect.Message {
	mi := &file_triggers_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCRAttributedOnchainSignature.ProtoReflect.Descriptor instead.
func (*OCRAttributedOnchainSignature) Descriptor() ([]byte, []int) {
	return file_triggers_proto_rawDescGZIP(), []int{1}
}

func (x *OCRAttributedOnchainSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *OCRAttributedOnchainSignature) GetSigner() uint32 {
	if x != nil {
		return x.Signer
	}
	return 0
}

type OCRTriggerReport struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventID       string                 `protobuf:"bytes,1,opt,name=eventID,proto3" json:"eventID,omitempty"`      // unique, scoped to the trigger capability
	Timestamp     uint64                 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // used to enforce freshness
	Outputs       *pb.Map                `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`      // contains trigger-specific data
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OCRTriggerReport) Reset() {
	*x = OCRTriggerReport{}
	mi := &file_triggers_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OCRTriggerReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCRTriggerReport) ProtoMessage() {}

func (x *OCRTriggerReport) ProtoReflect() protoreflect.Message {
	mi := &file_triggers_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCRTriggerReport.ProtoReflect.Descriptor instead.
func (*OCRTriggerReport) Descriptor() ([]byte, []int) {
	return file_triggers_proto_rawDescGZIP(), []int{2}
}

func (x *OCRTriggerReport) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

func (x *OCRTriggerReport) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *OCRTriggerReport) GetOutputs() *pb.Map {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_triggers_proto protoreflect.FileDescriptor

const file_triggers_proto_rawDesc = "" +
	"\n" +
	"\x0etriggers.proto\x12\x02v1\x1a\x16values/v1/values.proto\"\x9a\x01\n" +
	"\x0fOCRTriggerEvent\x12\"\n" +
	"\fconfigDigest\x18\x01 \x01(\fR\fconfigDigest\x12\x14\n" +
	"\x05seqNr\x18\x02 \x01(\x04R\x05seqNr\x12\x16\n" +
	"\x06report\x18\x03 \x01(\fR\x06report\x125\n" +
	"\x04sigs\x18\x04 \x03(\v2!.v1.OCRAttributedOnchainSignatureR\x04sigs\"U\n" +
	"\x1dOCRAttributedOnchainSignature\x12\x1c\n" +
	"\tsignature\x18\x01 \x01(\fR\tsignature\x12\x16\n" +
	"\x06signer\x18\x02 \x01(\rR\x06signer\"t\n" +
	"\x10OCRTriggerReport\x12\x18\n" +
	"\aeventID\x18\x01 \x01(\tR\aeventID\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\x04R\ttimestamp\x12(\n" +
	"\aoutputs\x18\x03 \x01(\v2\x0e.values.v1.MapR\aoutputsBBZ@github.com/smartcontractkit/chainlink-common/pkg/capabilities/pbb\x06proto3"

var (
	file_triggers_proto_rawDescOnce sync.Once
	file_triggers_proto_rawDescData []byte
)

func file_triggers_proto_rawDescGZIP() []byte {
	file_triggers_proto_rawDescOnce.Do(func() {
		file_triggers_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_triggers_proto_rawDesc), len(file_triggers_proto_rawDesc)))
	})
	return file_triggers_proto_rawDescData
}

var file_triggers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_triggers_proto_goTypes = []any{
	(*OCRTriggerEvent)(nil),               // 0: v1.OCRTriggerEvent
	(*OCRAttributedOnchainSignature)(nil), // 1: v1.OCRAttributedOnchainSignature
	(*OCRTriggerReport)(nil),              // 2: v1.OCRTriggerReport
	(*pb.Map)(nil),                        // 3: values.v1.Map
}
var file_triggers_proto_depIdxs = []int32{
	1, // 0: v1.OCRTriggerEvent.sigs:type_name -> v1.OCRAttributedOnchainSignature
	3, // 1: v1.OCRTriggerReport.outputs:type_name -> values.v1.Map
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_triggers_proto_init() }
func file_triggers_proto_init() {
	if File_triggers_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_triggers_proto_rawDesc), len(file_triggers_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_triggers_proto_goTypes,
		DependencyIndexes: file_triggers_proto_depIdxs,
		MessageInfos:      file_triggers_proto_msgTypes,
	}.Build()
	File_triggers_proto = out.File
	file_triggers_proto_goTypes = nil
	file_triggers_proto_depIdxs = nil
}
