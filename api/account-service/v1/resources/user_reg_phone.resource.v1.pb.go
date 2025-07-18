// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v6.30.2
// source: api/account-service/v1/resources/user_reg_phone.resource.v1.proto

package resourcev1

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

// UserRegPhone 用户_注册_手机
type UserRegPhone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// deleted_time 删除时间
	DeletedTime uint64 `protobuf:"varint,4,opt,name=deleted_time,json=deletedTime,proto3" json:"deleted_time,omitempty"`
	// user_id uid
	UserId uint64 `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// user_phone 手机
	UserPhone string `protobuf:"bytes,6,opt,name=user_phone,json=userPhone,proto3" json:"user_phone,omitempty"`
}

func (x *UserRegPhone) Reset() {
	*x = UserRegPhone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegPhone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegPhone) ProtoMessage() {}

func (x *UserRegPhone) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegPhone.ProtoReflect.Descriptor instead.
func (*UserRegPhone) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegPhone) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRegPhone) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *UserRegPhone) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *UserRegPhone) GetDeletedTime() uint64 {
	if x != nil {
		return x.DeletedTime
	}
	return 0
}

func (x *UserRegPhone) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRegPhone) GetUserPhone() string {
	if x != nil {
		return x.UserPhone
	}
	return ""
}

var File_api_account_service_v1_resources_user_reg_phone_resource_v1_proto protoreflect.FileDescriptor

var file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x41, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31,
	0x22, 0xbf, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x42, 0x8b, 0x01, 0x0a, 0x1b, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x76, 0x31, 0x42, 0x18, 0x53, 0x61, 0x61, 0x73, 0x41, 0x70, 0x69, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDescOnce sync.Once
	file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDescData = file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDesc
)

func file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDescGZIP() []byte {
	file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDescData)
	})
	return file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDescData
}

var file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_goTypes = []any{
	(*UserRegPhone)(nil), // 0: saas.api.account.resourcev1.UserRegPhone
}
var file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_init() }
func file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_init() {
	if File_api_account_service_v1_resources_user_reg_phone_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserRegPhone); i {
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
			RawDescriptor: file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_msgTypes,
	}.Build()
	File_api_account_service_v1_resources_user_reg_phone_resource_v1_proto = out.File
	file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_rawDesc = nil
	file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_goTypes = nil
	file_api_account_service_v1_resources_user_reg_phone_resource_v1_proto_depIdxs = nil
}
