// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: mockoidc/v1/mockoidc.proto

package mockoidcv1

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

type PushUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject           string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Email             string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PreferredUsername string `protobuf:"bytes,3,opt,name=preferred_username,json=preferredUsername,proto3" json:"preferred_username,omitempty"`
}

func (x *PushUserRequest) Reset() {
	*x = PushUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockoidc_v1_mockoidc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushUserRequest) ProtoMessage() {}

func (x *PushUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockoidc_v1_mockoidc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushUserRequest.ProtoReflect.Descriptor instead.
func (*PushUserRequest) Descriptor() ([]byte, []int) {
	return file_mockoidc_v1_mockoidc_proto_rawDescGZIP(), []int{0}
}

func (x *PushUserRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *PushUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PushUserRequest) GetPreferredUsername() string {
	if x != nil {
		return x.PreferredUsername
	}
	return ""
}

type PushUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushUserResponse) Reset() {
	*x = PushUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockoidc_v1_mockoidc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushUserResponse) ProtoMessage() {}

func (x *PushUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mockoidc_v1_mockoidc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushUserResponse.ProtoReflect.Descriptor instead.
func (*PushUserResponse) Descriptor() ([]byte, []int) {
	return file_mockoidc_v1_mockoidc_proto_rawDescGZIP(), []int{1}
}

var File_mockoidc_v1_mockoidc_proto protoreflect.FileDescriptor

var file_mockoidc_v1_mockoidc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x63, 0x6b, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x63, 0x6b, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x6f,
	0x63, 0x6b, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x76, 0x31, 0x22, 0x70, 0x0a, 0x0f, 0x50, 0x75, 0x73,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2d, 0x0a, 0x12,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x50,
	0x75, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x5c, 0x0a, 0x0f, 0x4d, 0x6f, 0x63, 0x6b, 0x4f, 0x49, 0x44, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d,
	0x6f, 0x63, 0x6b, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x6f, 0x69, 0x64,
	0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x6f, 0x69,
	0x64, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x63, 0x6b, 0x6f, 0x69, 0x64, 0x63, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockoidc_v1_mockoidc_proto_rawDescOnce sync.Once
	file_mockoidc_v1_mockoidc_proto_rawDescData = file_mockoidc_v1_mockoidc_proto_rawDesc
)

func file_mockoidc_v1_mockoidc_proto_rawDescGZIP() []byte {
	file_mockoidc_v1_mockoidc_proto_rawDescOnce.Do(func() {
		file_mockoidc_v1_mockoidc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockoidc_v1_mockoidc_proto_rawDescData)
	})
	return file_mockoidc_v1_mockoidc_proto_rawDescData
}

var file_mockoidc_v1_mockoidc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mockoidc_v1_mockoidc_proto_goTypes = []interface{}{
	(*PushUserRequest)(nil),  // 0: mockoidc.v1.PushUserRequest
	(*PushUserResponse)(nil), // 1: mockoidc.v1.PushUserResponse
}
var file_mockoidc_v1_mockoidc_proto_depIdxs = []int32{
	0, // 0: mockoidc.v1.MockOIDCService.PushUser:input_type -> mockoidc.v1.PushUserRequest
	1, // 1: mockoidc.v1.MockOIDCService.PushUser:output_type -> mockoidc.v1.PushUserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mockoidc_v1_mockoidc_proto_init() }
func file_mockoidc_v1_mockoidc_proto_init() {
	if File_mockoidc_v1_mockoidc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockoidc_v1_mockoidc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushUserRequest); i {
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
		file_mockoidc_v1_mockoidc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushUserResponse); i {
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
			RawDescriptor: file_mockoidc_v1_mockoidc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mockoidc_v1_mockoidc_proto_goTypes,
		DependencyIndexes: file_mockoidc_v1_mockoidc_proto_depIdxs,
		MessageInfos:      file_mockoidc_v1_mockoidc_proto_msgTypes,
	}.Build()
	File_mockoidc_v1_mockoidc_proto = out.File
	file_mockoidc_v1_mockoidc_proto_rawDesc = nil
	file_mockoidc_v1_mockoidc_proto_goTypes = nil
	file_mockoidc_v1_mockoidc_proto_depIdxs = nil
}
