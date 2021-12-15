// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/component/v1/gripper.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GripperServiceOpenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GripperServiceOpenRequest) Reset() {
	*x = GripperServiceOpenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gripper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GripperServiceOpenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GripperServiceOpenRequest) ProtoMessage() {}

func (x *GripperServiceOpenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gripper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GripperServiceOpenRequest.ProtoReflect.Descriptor instead.
func (*GripperServiceOpenRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gripper_proto_rawDescGZIP(), []int{0}
}

func (x *GripperServiceOpenRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GripperServiceOpenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GripperServiceOpenResponse) Reset() {
	*x = GripperServiceOpenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gripper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GripperServiceOpenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GripperServiceOpenResponse) ProtoMessage() {}

func (x *GripperServiceOpenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gripper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GripperServiceOpenResponse.ProtoReflect.Descriptor instead.
func (*GripperServiceOpenResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gripper_proto_rawDescGZIP(), []int{1}
}

type GripperServiceGrabRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GripperServiceGrabRequest) Reset() {
	*x = GripperServiceGrabRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gripper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GripperServiceGrabRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GripperServiceGrabRequest) ProtoMessage() {}

func (x *GripperServiceGrabRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gripper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GripperServiceGrabRequest.ProtoReflect.Descriptor instead.
func (*GripperServiceGrabRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gripper_proto_rawDescGZIP(), []int{2}
}

func (x *GripperServiceGrabRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GripperServiceGrabResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grabbed bool `protobuf:"varint,1,opt,name=grabbed,proto3" json:"grabbed,omitempty"`
}

func (x *GripperServiceGrabResponse) Reset() {
	*x = GripperServiceGrabResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gripper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GripperServiceGrabResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GripperServiceGrabResponse) ProtoMessage() {}

func (x *GripperServiceGrabResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gripper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GripperServiceGrabResponse.ProtoReflect.Descriptor instead.
func (*GripperServiceGrabResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gripper_proto_rawDescGZIP(), []int{3}
}

func (x *GripperServiceGrabResponse) GetGrabbed() bool {
	if x != nil {
		return x.Grabbed
	}
	return false
}

var File_proto_api_component_v1_gripper_proto protoreflect.FileDescriptor

var file_proto_api_component_v1_gripper_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x19,
	0x47, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a,
	0x1a, 0x47, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x0a, 0x19, 0x47,
	0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x61,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x1a,
	0x47, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72,
	0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72,
	0x61, 0x62, 0x62, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x67, 0x72, 0x61,
	0x62, 0x62, 0x65, 0x64, 0x32, 0xce, 0x02, 0x0a, 0x0e, 0x47, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e,
	0x12, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x1a,
	0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x9c, 0x01, 0x0a, 0x04, 0x47, 0x72, 0x61, 0x62, 0x12,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x61, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x61, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x1a, 0x25,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2f, 0x67, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x2f, 0x67, 0x72, 0x61, 0x62, 0x42, 0x4f, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x5a, 0x27, 0x67,
	0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_component_v1_gripper_proto_rawDescOnce sync.Once
	file_proto_api_component_v1_gripper_proto_rawDescData = file_proto_api_component_v1_gripper_proto_rawDesc
)

func file_proto_api_component_v1_gripper_proto_rawDescGZIP() []byte {
	file_proto_api_component_v1_gripper_proto_rawDescOnce.Do(func() {
		file_proto_api_component_v1_gripper_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_component_v1_gripper_proto_rawDescData)
	})
	return file_proto_api_component_v1_gripper_proto_rawDescData
}

var file_proto_api_component_v1_gripper_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_api_component_v1_gripper_proto_goTypes = []interface{}{
	(*GripperServiceOpenRequest)(nil),  // 0: proto.api.component.v1.GripperServiceOpenRequest
	(*GripperServiceOpenResponse)(nil), // 1: proto.api.component.v1.GripperServiceOpenResponse
	(*GripperServiceGrabRequest)(nil),  // 2: proto.api.component.v1.GripperServiceGrabRequest
	(*GripperServiceGrabResponse)(nil), // 3: proto.api.component.v1.GripperServiceGrabResponse
}
var file_proto_api_component_v1_gripper_proto_depIdxs = []int32{
	0, // 0: proto.api.component.v1.GripperService.Open:input_type -> proto.api.component.v1.GripperServiceOpenRequest
	2, // 1: proto.api.component.v1.GripperService.Grab:input_type -> proto.api.component.v1.GripperServiceGrabRequest
	1, // 2: proto.api.component.v1.GripperService.Open:output_type -> proto.api.component.v1.GripperServiceOpenResponse
	3, // 3: proto.api.component.v1.GripperService.Grab:output_type -> proto.api.component.v1.GripperServiceGrabResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_api_component_v1_gripper_proto_init() }
func file_proto_api_component_v1_gripper_proto_init() {
	if File_proto_api_component_v1_gripper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_component_v1_gripper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GripperServiceOpenRequest); i {
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
		file_proto_api_component_v1_gripper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GripperServiceOpenResponse); i {
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
		file_proto_api_component_v1_gripper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GripperServiceGrabRequest); i {
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
		file_proto_api_component_v1_gripper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GripperServiceGrabResponse); i {
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
			RawDescriptor: file_proto_api_component_v1_gripper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_component_v1_gripper_proto_goTypes,
		DependencyIndexes: file_proto_api_component_v1_gripper_proto_depIdxs,
		MessageInfos:      file_proto_api_component_v1_gripper_proto_msgTypes,
	}.Build()
	File_proto_api_component_v1_gripper_proto = out.File
	file_proto_api_component_v1_gripper_proto_rawDesc = nil
	file_proto_api_component_v1_gripper_proto_goTypes = nil
	file_proto_api_component_v1_gripper_proto_depIdxs = nil
}
