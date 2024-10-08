// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.2
// source: v1/get-profiles-min-info.proto

package v1

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

type ProfileMinInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VisibleName *string `protobuf:"bytes,1,opt,name=visibleName,proto3,oneof" json:"visibleName,omitempty"`
	Avatar      *string `protobuf:"bytes,2,opt,name=avatar,proto3,oneof" json:"avatar,omitempty"`
}

func (x *ProfileMinInfo) Reset() {
	*x = ProfileMinInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_profiles_min_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileMinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileMinInfo) ProtoMessage() {}

func (x *ProfileMinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_profiles_min_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileMinInfo.ProtoReflect.Descriptor instead.
func (*ProfileMinInfo) Descriptor() ([]byte, []int) {
	return file_v1_get_profiles_min_info_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileMinInfo) GetVisibleName() string {
	if x != nil && x.VisibleName != nil {
		return *x.VisibleName
	}
	return ""
}

func (x *ProfileMinInfo) GetAvatar() string {
	if x != nil && x.Avatar != nil {
		return *x.Avatar
	}
	return ""
}

type GetProfilesMinInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileIds []string `protobuf:"bytes,1,rep,name=profileIds,proto3" json:"profileIds,omitempty"`
}

func (x *GetProfilesMinInfoRequest) Reset() {
	*x = GetProfilesMinInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_profiles_min_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfilesMinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfilesMinInfoRequest) ProtoMessage() {}

func (x *GetProfilesMinInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_profiles_min_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfilesMinInfoRequest.ProtoReflect.Descriptor instead.
func (*GetProfilesMinInfoRequest) Descriptor() ([]byte, []int) {
	return file_v1_get_profiles_min_info_proto_rawDescGZIP(), []int{1}
}

func (x *GetProfilesMinInfoRequest) GetProfileIds() []string {
	if x != nil {
		return x.ProfileIds
	}
	return nil
}

type GetProfilesMinInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos map[string]*ProfileMinInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetProfilesMinInfoResponse) Reset() {
	*x = GetProfilesMinInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_profiles_min_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfilesMinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfilesMinInfoResponse) ProtoMessage() {}

func (x *GetProfilesMinInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_profiles_min_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfilesMinInfoResponse.ProtoReflect.Descriptor instead.
func (*GetProfilesMinInfoResponse) Descriptor() ([]byte, []int) {
	return file_v1_get_profiles_min_info_proto_rawDescGZIP(), []int{2}
}

func (x *GetProfilesMinInfoResponse) GetInfos() map[string]*ProfileMinInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_v1_get_profiles_min_info_proto protoreflect.FileDescriptor

var file_v1_get_profiles_min_info_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2d, 0x6d, 0x69, 0x6e, 0x2d, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x22, 0x6f, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0b, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x3b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x4d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x4d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x4d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x1a, 0x4c, 0x0a, 0x0a, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x69,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x65, 0x70, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62,
	0x6f, 0x6f, 0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_get_profiles_min_info_proto_rawDescOnce sync.Once
	file_v1_get_profiles_min_info_proto_rawDescData = file_v1_get_profiles_min_info_proto_rawDesc
)

func file_v1_get_profiles_min_info_proto_rawDescGZIP() []byte {
	file_v1_get_profiles_min_info_proto_rawDescOnce.Do(func() {
		file_v1_get_profiles_min_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_get_profiles_min_info_proto_rawDescData)
	})
	return file_v1_get_profiles_min_info_proto_rawDescData
}

var file_v1_get_profiles_min_info_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_get_profiles_min_info_proto_goTypes = []interface{}{
	(*ProfileMinInfo)(nil),             // 0: v1.ProfileMinInfo
	(*GetProfilesMinInfoRequest)(nil),  // 1: v1.GetProfilesMinInfoRequest
	(*GetProfilesMinInfoResponse)(nil), // 2: v1.GetProfilesMinInfoResponse
	nil,                                // 3: v1.GetProfilesMinInfoResponse.InfosEntry
}
var file_v1_get_profiles_min_info_proto_depIdxs = []int32{
	3, // 0: v1.GetProfilesMinInfoResponse.infos:type_name -> v1.GetProfilesMinInfoResponse.InfosEntry
	0, // 1: v1.GetProfilesMinInfoResponse.InfosEntry.value:type_name -> v1.ProfileMinInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_get_profiles_min_info_proto_init() }
func file_v1_get_profiles_min_info_proto_init() {
	if File_v1_get_profiles_min_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_get_profiles_min_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileMinInfo); i {
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
		file_v1_get_profiles_min_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfilesMinInfoRequest); i {
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
		file_v1_get_profiles_min_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfilesMinInfoResponse); i {
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
	file_v1_get_profiles_min_info_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_get_profiles_min_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_get_profiles_min_info_proto_goTypes,
		DependencyIndexes: file_v1_get_profiles_min_info_proto_depIdxs,
		MessageInfos:      file_v1_get_profiles_min_info_proto_msgTypes,
	}.Build()
	File_v1_get_profiles_min_info_proto = out.File
	file_v1_get_profiles_min_info_proto_rawDesc = nil
	file_v1_get_profiles_min_info_proto_goTypes = nil
	file_v1_get_profiles_min_info_proto_depIdxs = nil
}
