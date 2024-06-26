// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: prospect.proto

package proto

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

type ProspectFlowInitJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CampaignId string `protobuf:"bytes,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
}

func (x *ProspectFlowInitJob) Reset() {
	*x = ProspectFlowInitJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prospect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProspectFlowInitJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProspectFlowInitJob) ProtoMessage() {}

func (x *ProspectFlowInitJob) ProtoReflect() protoreflect.Message {
	mi := &file_prospect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProspectFlowInitJob.ProtoReflect.Descriptor instead.
func (*ProspectFlowInitJob) Descriptor() ([]byte, []int) {
	return file_prospect_proto_rawDescGZIP(), []int{0}
}

func (x *ProspectFlowInitJob) GetCampaignId() string {
	if x != nil {
		return x.CampaignId
	}
	return ""
}

type HotCampaignEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *HotCampaignEntry) Reset() {
	*x = HotCampaignEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prospect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotCampaignEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotCampaignEntry) ProtoMessage() {}

func (x *HotCampaignEntry) ProtoReflect() protoreflect.Message {
	mi := &file_prospect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotCampaignEntry.ProtoReflect.Descriptor instead.
func (*HotCampaignEntry) Descriptor() ([]byte, []int) {
	return file_prospect_proto_rawDescGZIP(), []int{1}
}

func (x *HotCampaignEntry) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

var File_prospect_proto protoreflect.FileDescriptor

var file_prospect_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x22, 0x36, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x69, 0x74, 0x4a, 0x6f,
	0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x49, 0x64, 0x22, 0x2d, 0x0a, 0x10, 0x48, 0x6f, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prospect_proto_rawDescOnce sync.Once
	file_prospect_proto_rawDescData = file_prospect_proto_rawDesc
)

func file_prospect_proto_rawDescGZIP() []byte {
	file_prospect_proto_rawDescOnce.Do(func() {
		file_prospect_proto_rawDescData = protoimpl.X.CompressGZIP(file_prospect_proto_rawDescData)
	})
	return file_prospect_proto_rawDescData
}

var file_prospect_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_prospect_proto_goTypes = []interface{}{
	(*ProspectFlowInitJob)(nil), // 0: goosechase.data.prospect.ProspectFlowInitJob
	(*HotCampaignEntry)(nil),    // 1: goosechase.data.prospect.HotCampaignEntry
}
var file_prospect_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prospect_proto_init() }
func file_prospect_proto_init() {
	if File_prospect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prospect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProspectFlowInitJob); i {
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
		file_prospect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotCampaignEntry); i {
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
			RawDescriptor: file_prospect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prospect_proto_goTypes,
		DependencyIndexes: file_prospect_proto_depIdxs,
		MessageInfos:      file_prospect_proto_msgTypes,
	}.Build()
	File_prospect_proto = out.File
	file_prospect_proto_rawDesc = nil
	file_prospect_proto_goTypes = nil
	file_prospect_proto_depIdxs = nil
}
