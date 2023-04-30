// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: v1/properties.proto

package proto

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

type PropertiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Direction can only be "in" or "out"
	Direction string `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *PropertiesRequest) Reset() {
	*x = PropertiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_properties_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertiesRequest) ProtoMessage() {}

func (x *PropertiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertiesRequest.ProtoReflect.Descriptor instead.
func (*PropertiesRequest) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{0}
}

func (x *PropertiesRequest) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *PropertiesRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type PropertiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node       string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Properties []string `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *PropertiesResponse) Reset() {
	*x = PropertiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_properties_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertiesResponse) ProtoMessage() {}

func (x *PropertiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertiesResponse.ProtoReflect.Descriptor instead.
func (*PropertiesResponse) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{1}
}

func (x *PropertiesResponse) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *PropertiesResponse) GetProperties() []string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type BulkPropertiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// Direction can only be "in" or "out"
	Direction string `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *BulkPropertiesRequest) Reset() {
	*x = BulkPropertiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_properties_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkPropertiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkPropertiesRequest) ProtoMessage() {}

func (x *BulkPropertiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkPropertiesRequest.ProtoReflect.Descriptor instead.
func (*BulkPropertiesRequest) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{2}
}

func (x *BulkPropertiesRequest) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *BulkPropertiesRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type BulkPropertiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*PropertiesResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BulkPropertiesResponse) Reset() {
	*x = BulkPropertiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_properties_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkPropertiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkPropertiesResponse) ProtoMessage() {}

func (x *BulkPropertiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkPropertiesResponse.ProtoReflect.Descriptor instead.
func (*BulkPropertiesResponse) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{3}
}

func (x *BulkPropertiesResponse) GetData() []*PropertiesResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_v1_properties_proto protoreflect.FileDescriptor

var file_v1_properties_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x45, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x12,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x15, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x16, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_properties_proto_rawDescOnce sync.Once
	file_v1_properties_proto_rawDescData = file_v1_properties_proto_rawDesc
)

func file_v1_properties_proto_rawDescGZIP() []byte {
	file_v1_properties_proto_rawDescOnce.Do(func() {
		file_v1_properties_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_properties_proto_rawDescData)
	})
	return file_v1_properties_proto_rawDescData
}

var file_v1_properties_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_properties_proto_goTypes = []interface{}{
	(*PropertiesRequest)(nil),      // 0: datacommons.v1.PropertiesRequest
	(*PropertiesResponse)(nil),     // 1: datacommons.v1.PropertiesResponse
	(*BulkPropertiesRequest)(nil),  // 2: datacommons.v1.BulkPropertiesRequest
	(*BulkPropertiesResponse)(nil), // 3: datacommons.v1.BulkPropertiesResponse
}
var file_v1_properties_proto_depIdxs = []int32{
	1, // 0: datacommons.v1.BulkPropertiesResponse.data:type_name -> datacommons.v1.PropertiesResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_properties_proto_init() }
func file_v1_properties_proto_init() {
	if File_v1_properties_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_properties_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertiesRequest); i {
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
		file_v1_properties_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertiesResponse); i {
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
		file_v1_properties_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkPropertiesRequest); i {
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
		file_v1_properties_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkPropertiesResponse); i {
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
			RawDescriptor: file_v1_properties_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_properties_proto_goTypes,
		DependencyIndexes: file_v1_properties_proto_depIdxs,
		MessageInfos:      file_v1_properties_proto_msgTypes,
	}.Build()
	File_v1_properties_proto = out.File
	file_v1_properties_proto_rawDesc = nil
	file_v1_properties_proto_goTypes = nil
	file_v1_properties_proto_depIdxs = nil
}
