// Holds the ConversationReference

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: message_contents/conversation_reference.proto

package message_contents

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

// A light pointer for a conversation that contains no decryption keys
type ConversationReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic       string                `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	PeerAddress string                `protobuf:"bytes,2,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	CreatedNs   uint64                `protobuf:"varint,3,opt,name=created_ns,json=createdNs,proto3" json:"created_ns,omitempty"`
	Context     *InvitationV1_Context `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *ConversationReference) Reset() {
	*x = ConversationReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_contents_conversation_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationReference) ProtoMessage() {}

func (x *ConversationReference) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_conversation_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationReference.ProtoReflect.Descriptor instead.
func (*ConversationReference) Descriptor() ([]byte, []int) {
	return file_message_contents_conversation_reference_proto_rawDescGZIP(), []int{0}
}

func (x *ConversationReference) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ConversationReference) GetPeerAddress() string {
	if x != nil {
		return x.PeerAddress
	}
	return ""
}

func (x *ConversationReference) GetCreatedNs() uint64 {
	if x != nil {
		return x.CreatedNs
	}
	return 0
}

func (x *ConversationReference) GetContext() *InvitationV1_Context {
	if x != nil {
		return x.Context
	}
	return nil
}

var File_message_contents_conversation_reference_proto protoreflect.FileDescriptor

var file_message_contents_conversation_reference_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x21, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x15, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x73, 0x12, 0x45, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x78,
	0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x42, 0x4f, 0x0a, 0x1f, 0x6f, 0x72, 0x67, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f,
	0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_contents_conversation_reference_proto_rawDescOnce sync.Once
	file_message_contents_conversation_reference_proto_rawDescData = file_message_contents_conversation_reference_proto_rawDesc
)

func file_message_contents_conversation_reference_proto_rawDescGZIP() []byte {
	file_message_contents_conversation_reference_proto_rawDescOnce.Do(func() {
		file_message_contents_conversation_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_contents_conversation_reference_proto_rawDescData)
	})
	return file_message_contents_conversation_reference_proto_rawDescData
}

var file_message_contents_conversation_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_contents_conversation_reference_proto_goTypes = []interface{}{
	(*ConversationReference)(nil), // 0: xmtp.message_contents.ConversationReference
	(*InvitationV1_Context)(nil),  // 1: xmtp.message_contents.InvitationV1.Context
}
var file_message_contents_conversation_reference_proto_depIdxs = []int32{
	1, // 0: xmtp.message_contents.ConversationReference.context:type_name -> xmtp.message_contents.InvitationV1.Context
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_contents_conversation_reference_proto_init() }
func file_message_contents_conversation_reference_proto_init() {
	if File_message_contents_conversation_reference_proto != nil {
		return
	}
	file_message_contents_invitation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_contents_conversation_reference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationReference); i {
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
			RawDescriptor: file_message_contents_conversation_reference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_contents_conversation_reference_proto_goTypes,
		DependencyIndexes: file_message_contents_conversation_reference_proto_depIdxs,
		MessageInfos:      file_message_contents_conversation_reference_proto_msgTypes,
	}.Build()
	File_message_contents_conversation_reference_proto = out.File
	file_message_contents_conversation_reference_proto_rawDesc = nil
	file_message_contents_conversation_reference_proto_goTypes = nil
	file_message_contents_conversation_reference_proto_depIdxs = nil
}
