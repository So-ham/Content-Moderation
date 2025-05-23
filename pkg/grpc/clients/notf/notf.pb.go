// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: notf.proto

package notf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetCardsMap messages
type SendFlaggedNotificationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   uint32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Severity string `protobuf:"bytes,3,opt,name=severity,proto3" json:"severity,omitempty"`
}

func (x *SendFlaggedNotificationReq) Reset() {
	*x = SendFlaggedNotificationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendFlaggedNotificationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendFlaggedNotificationReq) ProtoMessage() {}

func (x *SendFlaggedNotificationReq) ProtoReflect() protoreflect.Message {
	mi := &file_notf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendFlaggedNotificationReq.ProtoReflect.Descriptor instead.
func (*SendFlaggedNotificationReq) Descriptor() ([]byte, []int) {
	return file_notf_proto_rawDescGZIP(), []int{0}
}

func (x *SendFlaggedNotificationReq) GetUserID() uint32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *SendFlaggedNotificationReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendFlaggedNotificationReq) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

var File_notf_proto protoreflect.FileDescriptor

var file_notf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f,
	0x74, 0x66, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6a, 0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x67, 0x65, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x32, 0x62, 0x0a, 0x0b, 0x4e,
	0x6f, 0x74, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x17, 0x53, 0x65,
	0x6e, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x66, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x46, 0x6c, 0x61, 0x67, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x2a, 0x5a, 0x28, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x4d, 0x6f, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_notf_proto_rawDescOnce sync.Once
	file_notf_proto_rawDescData = file_notf_proto_rawDesc
)

func file_notf_proto_rawDescGZIP() []byte {
	file_notf_proto_rawDescOnce.Do(func() {
		file_notf_proto_rawDescData = protoimpl.X.CompressGZIP(file_notf_proto_rawDescData)
	})
	return file_notf_proto_rawDescData
}

var file_notf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_notf_proto_goTypes = []interface{}{
	(*SendFlaggedNotificationReq)(nil), // 0: notf.SendFlaggedNotificationReq
	(*emptypb.Empty)(nil),              // 1: google.protobuf.Empty
}
var file_notf_proto_depIdxs = []int32{
	0, // 0: notf.NotfService.SendFlaggedNotification:input_type -> notf.SendFlaggedNotificationReq
	1, // 1: notf.NotfService.SendFlaggedNotification:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notf_proto_init() }
func file_notf_proto_init() {
	if File_notf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendFlaggedNotificationReq); i {
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
			RawDescriptor: file_notf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notf_proto_goTypes,
		DependencyIndexes: file_notf_proto_depIdxs,
		MessageInfos:      file_notf_proto_msgTypes,
	}.Build()
	File_notf_proto = out.File
	file_notf_proto_rawDesc = nil
	file_notf_proto_goTypes = nil
	file_notf_proto_depIdxs = nil
}
