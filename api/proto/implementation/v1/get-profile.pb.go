// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: v1/get-profile.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId       string `protobuf:"bytes,1,opt,name=profileId,proto3" json:"profileId,omitempty"`
	ProfileNickname string `protobuf:"bytes,2,opt,name=profileNickname,proto3" json:"profileNickname,omitempty"`
	RequesterId     string `protobuf:"bytes,3,opt,name=requesterId,proto3" json:"requesterId,omitempty"`
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_v1_get_profile_proto_rawDescGZIP(), []int{0}
}

func (x *GetProfileRequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *GetProfileRequest) GetProfileNickname() string {
	if x != nil {
		return x.ProfileNickname
	}
	return ""
}

func (x *GetProfileRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                 *string                `protobuf:"bytes,2,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Nickname              *string                `protobuf:"bytes,3,opt,name=nickname,proto3,oneof" json:"nickname,omitempty"`
	Role                  *string                `protobuf:"bytes,4,opt,name=role,proto3,oneof" json:"role,omitempty"`
	RegistrationTimestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=registrationTimestamp,proto3" json:"registrationTimestamp,omitempty"`
	IsBlocked             bool                   `protobuf:"varint,6,opt,name=isBlocked,proto3" json:"isBlocked,omitempty"`
	OAuth                 *OAuthConnections      `protobuf:"bytes,7,opt,name=oAuth,proto3" json:"oAuth,omitempty"`
	FirstName             *string                `protobuf:"bytes,8,opt,name=firstName,proto3,oneof" json:"firstName,omitempty"`
	LastName              *string                `protobuf:"bytes,9,opt,name=lastName,proto3,oneof" json:"lastName,omitempty"`
	Description           *string                `protobuf:"bytes,10,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Avatar                *string                `protobuf:"bytes,11,opt,name=avatar,proto3,oneof" json:"avatar,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_v1_get_profile_proto_rawDescGZIP(), []int{1}
}

func (x *GetProfileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProfileResponse) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *GetProfileResponse) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *GetProfileResponse) GetRole() string {
	if x != nil && x.Role != nil {
		return *x.Role
	}
	return ""
}

func (x *GetProfileResponse) GetRegistrationTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.RegistrationTimestamp
	}
	return nil
}

func (x *GetProfileResponse) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *GetProfileResponse) GetOAuth() *OAuthConnections {
	if x != nil {
		return x.OAuth
	}
	return nil
}

func (x *GetProfileResponse) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *GetProfileResponse) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *GetProfileResponse) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *GetProfileResponse) GetAvatar() string {
	if x != nil && x.Avatar != nil {
		return *x.Avatar
	}
	return ""
}

type OAuthConnections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoogleId *string `protobuf:"bytes,1,opt,name=googleId,proto3,oneof" json:"googleId,omitempty"`
	VkId     *int64  `protobuf:"varint,2,opt,name=vkId,proto3,oneof" json:"vkId,omitempty"`
}

func (x *OAuthConnections) Reset() {
	*x = OAuthConnections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthConnections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthConnections) ProtoMessage() {}

func (x *OAuthConnections) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthConnections.ProtoReflect.Descriptor instead.
func (*OAuthConnections) Descriptor() ([]byte, []int) {
	return file_v1_get_profile_proto_rawDescGZIP(), []int{2}
}

func (x *OAuthConnections) GetGoogleId() string {
	if x != nil && x.GoogleId != nil {
		return *x.GoogleId
	}
	return ""
}

func (x *OAuthConnections) GetVkId() int64 {
	if x != nil && x.VkId != nil {
		return *x.VkId
	}
	return 0
}

var File_v1_get_profile_proto protoreflect.FileDescriptor

var file_v1_get_profile_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0xf3, 0x03, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x50, 0x0a, 0x15, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x15, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x6f, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x22, 0x62, 0x0a, 0x10, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x76, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x04, 0x76, 0x6b, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x76, 0x6b, 0x49, 0x64, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63,
	0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_get_profile_proto_rawDescOnce sync.Once
	file_v1_get_profile_proto_rawDescData = file_v1_get_profile_proto_rawDesc
)

func file_v1_get_profile_proto_rawDescGZIP() []byte {
	file_v1_get_profile_proto_rawDescOnce.Do(func() {
		file_v1_get_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_get_profile_proto_rawDescData)
	})
	return file_v1_get_profile_proto_rawDescData
}

var file_v1_get_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_get_profile_proto_goTypes = []interface{}{
	(*GetProfileRequest)(nil),     // 0: v1.GetProfileRequest
	(*GetProfileResponse)(nil),    // 1: v1.GetProfileResponse
	(*OAuthConnections)(nil),      // 2: v1.OAuthConnections
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_v1_get_profile_proto_depIdxs = []int32{
	3, // 0: v1.GetProfileResponse.registrationTimestamp:type_name -> google.protobuf.Timestamp
	2, // 1: v1.GetProfileResponse.oAuth:type_name -> v1.OAuthConnections
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_get_profile_proto_init() }
func file_v1_get_profile_proto_init() {
	if File_v1_get_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_get_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileRequest); i {
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
		file_v1_get_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
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
		file_v1_get_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthConnections); i {
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
	file_v1_get_profile_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_v1_get_profile_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_get_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_get_profile_proto_goTypes,
		DependencyIndexes: file_v1_get_profile_proto_depIdxs,
		MessageInfos:      file_v1_get_profile_proto_msgTypes,
	}.Build()
	File_v1_get_profile_proto = out.File
	file_v1_get_profile_proto_rawDesc = nil
	file_v1_get_profile_proto_goTypes = nil
	file_v1_get_profile_proto_depIdxs = nil
}
