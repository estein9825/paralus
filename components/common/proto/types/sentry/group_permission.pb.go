// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/types/sentry/group_permission.proto

package sentry

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GroupPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID        string           `protobuf:"bytes,1,opt,name=groupID,proto3" json:"groupID,omitempty"`
	ProjectID      string           `protobuf:"bytes,2,opt,name=projectID,proto3" json:"projectID,omitempty"`
	OrganizationID string           `protobuf:"bytes,3,opt,name=organizationID,proto3" json:"organizationID,omitempty"`
	PartnerID      string           `protobuf:"bytes,4,opt,name=partnerID,proto3" json:"partnerID,omitempty"`
	GroupName      string           `protobuf:"bytes,5,opt,name=groupName,proto3" json:"groupName,omitempty"`
	RoleName       string           `protobuf:"bytes,6,opt,name=roleName,proto3" json:"roleName,omitempty"`
	IsGlobal       bool             `protobuf:"varint,7,opt,name=isGlobal,proto3" json:"isGlobal,omitempty"`
	Scope          string           `protobuf:"bytes,8,opt,name=scope,proto3" json:"scope,omitempty"`
	PermissionName string           `protobuf:"bytes,9,opt,name=permissionName,proto3" json:"permissionName,omitempty"`
	BaseURL        string           `protobuf:"bytes,10,opt,name=baseURL,proto3" json:"baseURL,omitempty"`
	Urls           []*PermissionURL `protobuf:"bytes,11,rep,name=urls,proto3" json:"urls,omitempty"`
	ProjectName    string           `protobuf:"bytes,12,opt,name=projectName,proto3" json:"projectName,omitempty"`
}

func (x *GroupPermission) Reset() {
	*x = GroupPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_sentry_group_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPermission) ProtoMessage() {}

func (x *GroupPermission) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_sentry_group_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPermission.ProtoReflect.Descriptor instead.
func (*GroupPermission) Descriptor() ([]byte, []int) {
	return file_proto_types_sentry_group_permission_proto_rawDescGZIP(), []int{0}
}

func (x *GroupPermission) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *GroupPermission) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *GroupPermission) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

func (x *GroupPermission) GetPartnerID() string {
	if x != nil {
		return x.PartnerID
	}
	return ""
}

func (x *GroupPermission) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GroupPermission) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *GroupPermission) GetIsGlobal() bool {
	if x != nil {
		return x.IsGlobal
	}
	return false
}

func (x *GroupPermission) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *GroupPermission) GetPermissionName() string {
	if x != nil {
		return x.PermissionName
	}
	return ""
}

func (x *GroupPermission) GetBaseURL() string {
	if x != nil {
		return x.BaseURL
	}
	return ""
}

func (x *GroupPermission) GetUrls() []*PermissionURL {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *GroupPermission) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

var File_proto_types_sentry_group_permission_proto protoreflect.FileDescriptor

var file_proto_types_sentry_group_permission_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9a, 0x03, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x26, 0x0a,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55,
	0x52, 0x4c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x52,
	0x4c, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xf8,
	0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x14, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x72,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0xa2,
	0x02, 0x04, 0x52, 0x44, 0x54, 0x53, 0xaa, 0x02, 0x16, 0x52, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x44,
	0x65, 0x76, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0xca,
	0x02, 0x16, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x5c, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0xe2, 0x02, 0x22, 0x52, 0x61, 0x66, 0x61, 0x79,
	0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x53, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19,
	0x52, 0x61, 0x66, 0x61, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x3a, 0x3a, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_types_sentry_group_permission_proto_rawDescOnce sync.Once
	file_proto_types_sentry_group_permission_proto_rawDescData = file_proto_types_sentry_group_permission_proto_rawDesc
)

func file_proto_types_sentry_group_permission_proto_rawDescGZIP() []byte {
	file_proto_types_sentry_group_permission_proto_rawDescOnce.Do(func() {
		file_proto_types_sentry_group_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_sentry_group_permission_proto_rawDescData)
	})
	return file_proto_types_sentry_group_permission_proto_rawDescData
}

var file_proto_types_sentry_group_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_types_sentry_group_permission_proto_goTypes = []interface{}{
	(*GroupPermission)(nil), // 0: rafay.dev.types.sentry.GroupPermission
	(*PermissionURL)(nil),   // 1: rafay.dev.types.sentry.PermissionURL
}
var file_proto_types_sentry_group_permission_proto_depIdxs = []int32{
	1, // 0: rafay.dev.types.sentry.GroupPermission.urls:type_name -> rafay.dev.types.sentry.PermissionURL
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_types_sentry_group_permission_proto_init() }
func file_proto_types_sentry_group_permission_proto_init() {
	if File_proto_types_sentry_group_permission_proto != nil {
		return
	}
	file_proto_types_sentry_account_permission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_types_sentry_group_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPermission); i {
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
			RawDescriptor: file_proto_types_sentry_group_permission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_sentry_group_permission_proto_goTypes,
		DependencyIndexes: file_proto_types_sentry_group_permission_proto_depIdxs,
		MessageInfos:      file_proto_types_sentry_group_permission_proto_msgTypes,
	}.Build()
	File_proto_types_sentry_group_permission_proto = out.File
	file_proto_types_sentry_group_permission_proto_rawDesc = nil
	file_proto_types_sentry_group_permission_proto_goTypes = nil
	file_proto_types_sentry_group_permission_proto_depIdxs = nil
}
