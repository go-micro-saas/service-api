// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v6.30.2
// source: api/account-service/v1/services/account_event.service.v1.proto

package servicev1

import (
	resources "github.com/go-micro-saas/service-api/api/account-service/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_account_service_v1_services_account_event_service_v1_proto protoreflect.FileDescriptor

var file_api_account_service_v1_services_account_event_service_v1_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb8, 0x03, 0x0a, 0x11, 0x53,
	0x72, 0x76, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31,
	0x12, 0x77, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x25,
	0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x9a, 0x01, 0x0a, 0x1b, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x43, 0x6f, 0x64, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x2e, 0x73, 0x61, 0x61, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x3c, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x6f, 0x70, 0x53,
	0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x36, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x37, 0x2e, 0x73, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x87, 0x01, 0x0a, 0x1a, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x76, 0x31, 0x42, 0x17, 0x53, 0x61, 0x61, 0x73, 0x41, 0x70, 0x69, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x50, 0x01, 0x5a,
	0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_account_service_v1_services_account_event_service_v1_proto_goTypes = []any{
	(*resources.PingReq)(nil),                         // 0: saas.api.account.resourcev1.PingReq
	(*resources.SubscribeSendEmailCodeEventReq)(nil),  // 1: saas.api.account.resourcev1.SubscribeSendEmailCodeEventReq
	(*resources.StopSendEmailCodeEventReq)(nil),       // 2: saas.api.account.resourcev1.StopSendEmailCodeEventReq
	(*resources.PingResp)(nil),                        // 3: saas.api.account.resourcev1.PingResp
	(*resources.SubscribeSendEmailCodeEventResp)(nil), // 4: saas.api.account.resourcev1.SubscribeSendEmailCodeEventResp
	(*resources.StopSendEmailCodeEventResp)(nil),      // 5: saas.api.account.resourcev1.StopSendEmailCodeEventResp
}
var file_api_account_service_v1_services_account_event_service_v1_proto_depIdxs = []int32{
	0, // 0: saas.api.account.servicev1.SrvAccountEventV1.Ping:input_type -> saas.api.account.resourcev1.PingReq
	1, // 1: saas.api.account.servicev1.SrvAccountEventV1.SubscribeSendEmailCodeEvent:input_type -> saas.api.account.resourcev1.SubscribeSendEmailCodeEventReq
	2, // 2: saas.api.account.servicev1.SrvAccountEventV1.StopSendEmailCodedEvent:input_type -> saas.api.account.resourcev1.StopSendEmailCodeEventReq
	3, // 3: saas.api.account.servicev1.SrvAccountEventV1.Ping:output_type -> saas.api.account.resourcev1.PingResp
	4, // 4: saas.api.account.servicev1.SrvAccountEventV1.SubscribeSendEmailCodeEvent:output_type -> saas.api.account.resourcev1.SubscribeSendEmailCodeEventResp
	5, // 5: saas.api.account.servicev1.SrvAccountEventV1.StopSendEmailCodedEvent:output_type -> saas.api.account.resourcev1.StopSendEmailCodeEventResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_account_service_v1_services_account_event_service_v1_proto_init() }
func file_api_account_service_v1_services_account_event_service_v1_proto_init() {
	if File_api_account_service_v1_services_account_event_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_account_service_v1_services_account_event_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_account_service_v1_services_account_event_service_v1_proto_goTypes,
		DependencyIndexes: file_api_account_service_v1_services_account_event_service_v1_proto_depIdxs,
	}.Build()
	File_api_account_service_v1_services_account_event_service_v1_proto = out.File
	file_api_account_service_v1_services_account_event_service_v1_proto_rawDesc = nil
	file_api_account_service_v1_services_account_event_service_v1_proto_goTypes = nil
	file_api_account_service_v1_services_account_event_service_v1_proto_depIdxs = nil
}
