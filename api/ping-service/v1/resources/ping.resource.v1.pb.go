// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.6
// source: api/ping-service/v1/resources/ping.resource.v1.proto

package resourcev1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// PingReq ping请求
//
// ping请求
type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// message 请求消息
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// PingResp ping响应
//
// ping响应
type PingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason   string            `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message  string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data     *PingRespData     `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PingResp) Reset() {
	*x = PingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResp) ProtoMessage() {}

func (x *PingResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResp.ProtoReflect.Descriptor instead.
func (*PingResp) Descriptor() ([]byte, []int) {
	return file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescGZIP(), []int{1}
}

func (x *PingResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PingResp) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *PingResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PingResp) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PingResp) GetData() *PingRespData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PingRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingRespData) Reset() {
	*x = PingRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRespData) ProtoMessage() {}

func (x *PingRespData) ProtoReflect() protoreflect.Message {
	mi := &file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRespData.ProtoReflect.Descriptor instead.
func (*PingRespData) Descriptor() ([]byte, []int) {
	return file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescGZIP(), []int{2}
}

func (x *PingRespData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_ping_service_v1_resources_ping_resource_v1_proto protoreflect.FileDescriptor

var file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x34, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x28, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x97, 0x02, 0x0a, 0x08,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4c,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x98, 0x01, 0x0a, 0x18, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x42, 0x15, 0x53, 0x61,
	0x61, 0x73, 0x41, 0x70, 0x69, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescOnce sync.Once
	file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescData = file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDesc
)

func file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescGZIP() []byte {
	file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescData)
	})
	return file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDescData
}

var file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_ping_service_v1_resources_ping_resource_v1_proto_goTypes = []interface{}{
	(*PingReq)(nil),      // 0: saas.api.ping.resourcev1.PingReq
	(*PingResp)(nil),     // 1: saas.api.ping.resourcev1.PingResp
	(*PingRespData)(nil), // 2: saas.api.ping.resourcev1.PingRespData
	nil,                  // 3: saas.api.ping.resourcev1.PingResp.MetadataEntry
}
var file_api_ping_service_v1_resources_ping_resource_v1_proto_depIdxs = []int32{
	3, // 0: saas.api.ping.resourcev1.PingResp.metadata:type_name -> saas.api.ping.resourcev1.PingResp.MetadataEntry
	2, // 1: saas.api.ping.resourcev1.PingResp.data:type_name -> saas.api.ping.resourcev1.PingRespData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ping_service_v1_resources_ping_resource_v1_proto_init() }
func file_api_ping_service_v1_resources_ping_resource_v1_proto_init() {
	if File_api_ping_service_v1_resources_ping_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
		file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResp); i {
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
		file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRespData); i {
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
			RawDescriptor: file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ping_service_v1_resources_ping_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_ping_service_v1_resources_ping_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_ping_service_v1_resources_ping_resource_v1_proto_msgTypes,
	}.Build()
	File_api_ping_service_v1_resources_ping_resource_v1_proto = out.File
	file_api_ping_service_v1_resources_ping_resource_v1_proto_rawDesc = nil
	file_api_ping_service_v1_resources_ping_resource_v1_proto_goTypes = nil
	file_api_ping_service_v1_resources_ping_resource_v1_proto_depIdxs = nil
}
