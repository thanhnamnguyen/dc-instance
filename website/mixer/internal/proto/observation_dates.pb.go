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
// source: v1/observation_dates.proto

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

// Holds the count of entities for a particular facet.
type EntityCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Count of the entity.
	Count float64 `protobuf:"fixed64,1,opt,name=count,proto3" json:"count,omitempty"`
	// The facet hash string.
	Facet string `protobuf:"bytes,2,opt,name=facet,proto3" json:"facet,omitempty"`
}

func (x *EntityCount) Reset() {
	*x = EntityCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observation_dates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityCount) ProtoMessage() {}

func (x *EntityCount) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observation_dates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityCount.ProtoReflect.Descriptor instead.
func (*EntityCount) Descriptor() ([]byte, []int) {
	return file_v1_observation_dates_proto_rawDescGZIP(), []int{0}
}

func (x *EntityCount) GetCount() float64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *EntityCount) GetFacet() string {
	if x != nil {
		return x.Facet
	}
	return ""
}

type ObservationDates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Date of observation
	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	// For each facet, record the number of entities that have observations.
	EntityCount []*EntityCount `protobuf:"bytes,2,rep,name=entity_count,json=entityCount,proto3" json:"entity_count,omitempty"`
}

func (x *ObservationDates) Reset() {
	*x = ObservationDates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observation_dates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationDates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationDates) ProtoMessage() {}

func (x *ObservationDates) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observation_dates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationDates.ProtoReflect.Descriptor instead.
func (*ObservationDates) Descriptor() ([]byte, []int) {
	return file_v1_observation_dates_proto_rawDescGZIP(), []int{1}
}

func (x *ObservationDates) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ObservationDates) GetEntityCount() []*EntityCount {
	if x != nil {
		return x.EntityCount
	}
	return nil
}

type VariableObservationDates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stat var DCID.
	Variable string `protobuf:"bytes,1,opt,name=variable,proto3" json:"variable,omitempty"`
	// A list of observation data information for each stat var.
	ObservationDates []*ObservationDates `protobuf:"bytes,3,rep,name=observation_dates,json=observationDates,proto3" json:"observation_dates,omitempty"`
}

func (x *VariableObservationDates) Reset() {
	*x = VariableObservationDates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observation_dates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariableObservationDates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariableObservationDates) ProtoMessage() {}

func (x *VariableObservationDates) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observation_dates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariableObservationDates.ProtoReflect.Descriptor instead.
func (*VariableObservationDates) Descriptor() ([]byte, []int) {
	return file_v1_observation_dates_proto_rawDescGZIP(), []int{2}
}

func (x *VariableObservationDates) GetVariable() string {
	if x != nil {
		return x.Variable
	}
	return ""
}

func (x *VariableObservationDates) GetObservationDates() []*ObservationDates {
	if x != nil {
		return x.ObservationDates
	}
	return nil
}

type BulkObservationDatesLinkedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The observed entity type
	EntityType string `protobuf:"bytes,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	// The entity that is linked to the observed entity
	LinkedEntity string `protobuf:"bytes,2,opt,name=linked_entity,json=linkedEntity,proto3" json:"linked_entity,omitempty"`
	// The property that links the root and observed entity
	LinkedProperty string `protobuf:"bytes,3,opt,name=linked_property,json=linkedProperty,proto3" json:"linked_property,omitempty"`
	// Variables to query for
	Variables []string `protobuf:"bytes,4,rep,name=variables,proto3" json:"variables,omitempty"`
}

func (x *BulkObservationDatesLinkedRequest) Reset() {
	*x = BulkObservationDatesLinkedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observation_dates_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkObservationDatesLinkedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkObservationDatesLinkedRequest) ProtoMessage() {}

func (x *BulkObservationDatesLinkedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observation_dates_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkObservationDatesLinkedRequest.ProtoReflect.Descriptor instead.
func (*BulkObservationDatesLinkedRequest) Descriptor() ([]byte, []int) {
	return file_v1_observation_dates_proto_rawDescGZIP(), []int{3}
}

func (x *BulkObservationDatesLinkedRequest) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *BulkObservationDatesLinkedRequest) GetLinkedEntity() string {
	if x != nil {
		return x.LinkedEntity
	}
	return ""
}

func (x *BulkObservationDatesLinkedRequest) GetLinkedProperty() string {
	if x != nil {
		return x.LinkedProperty
	}
	return ""
}

func (x *BulkObservationDatesLinkedRequest) GetVariables() []string {
	if x != nil {
		return x.Variables
	}
	return nil
}

type BulkObservationDatesLinkedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatesByVariable []*VariableObservationDates `protobuf:"bytes,1,rep,name=dates_by_variable,json=datesByVariable,proto3" json:"dates_by_variable,omitempty"`
	Facets          map[string]*StatMetadata    `protobuf:"bytes,2,rep,name=facets,proto3" json:"facets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BulkObservationDatesLinkedResponse) Reset() {
	*x = BulkObservationDatesLinkedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observation_dates_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkObservationDatesLinkedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkObservationDatesLinkedResponse) ProtoMessage() {}

func (x *BulkObservationDatesLinkedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observation_dates_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkObservationDatesLinkedResponse.ProtoReflect.Descriptor instead.
func (*BulkObservationDatesLinkedResponse) Descriptor() ([]byte, []int) {
	return file_v1_observation_dates_proto_rawDescGZIP(), []int{4}
}

func (x *BulkObservationDatesLinkedResponse) GetDatesByVariable() []*VariableObservationDates {
	if x != nil {
		return x.DatesByVariable
	}
	return nil
}

func (x *BulkObservationDatesLinkedResponse) GetFacets() map[string]*StatMetadata {
	if x != nil {
		return x.Facets
	}
	return nil
}

var File_v1_observation_dates_proto protoreflect.FileDescriptor

var file_v1_observation_dates_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0b, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x61, 0x63, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x61,
	0x63, 0x65, 0x74, 0x22, 0x66, 0x0a, 0x10, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x18,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x4d, 0x0a, 0x11, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x10, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x21, 0x42, 0x75, 0x6c, 0x6b, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69,
	0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x27, 0x0a, 0x0f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0xa8, 0x02, 0x0a, 0x22, 0x42, 0x75, 0x6c, 0x6b, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x4c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x11, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x73, 0x42, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x56, 0x0a, 0x06, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x1a, 0x54, 0x0a, 0x0b, 0x46,
	0x61, 0x63, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_observation_dates_proto_rawDescOnce sync.Once
	file_v1_observation_dates_proto_rawDescData = file_v1_observation_dates_proto_rawDesc
)

func file_v1_observation_dates_proto_rawDescGZIP() []byte {
	file_v1_observation_dates_proto_rawDescOnce.Do(func() {
		file_v1_observation_dates_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_observation_dates_proto_rawDescData)
	})
	return file_v1_observation_dates_proto_rawDescData
}

var file_v1_observation_dates_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_observation_dates_proto_goTypes = []interface{}{
	(*EntityCount)(nil),                        // 0: datacommons.v1.EntityCount
	(*ObservationDates)(nil),                   // 1: datacommons.v1.ObservationDates
	(*VariableObservationDates)(nil),           // 2: datacommons.v1.VariableObservationDates
	(*BulkObservationDatesLinkedRequest)(nil),  // 3: datacommons.v1.BulkObservationDatesLinkedRequest
	(*BulkObservationDatesLinkedResponse)(nil), // 4: datacommons.v1.BulkObservationDatesLinkedResponse
	nil,                  // 5: datacommons.v1.BulkObservationDatesLinkedResponse.FacetsEntry
	(*StatMetadata)(nil), // 6: datacommons.StatMetadata
}
var file_v1_observation_dates_proto_depIdxs = []int32{
	0, // 0: datacommons.v1.ObservationDates.entity_count:type_name -> datacommons.v1.EntityCount
	1, // 1: datacommons.v1.VariableObservationDates.observation_dates:type_name -> datacommons.v1.ObservationDates
	2, // 2: datacommons.v1.BulkObservationDatesLinkedResponse.dates_by_variable:type_name -> datacommons.v1.VariableObservationDates
	5, // 3: datacommons.v1.BulkObservationDatesLinkedResponse.facets:type_name -> datacommons.v1.BulkObservationDatesLinkedResponse.FacetsEntry
	6, // 4: datacommons.v1.BulkObservationDatesLinkedResponse.FacetsEntry.value:type_name -> datacommons.StatMetadata
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_observation_dates_proto_init() }
func file_v1_observation_dates_proto_init() {
	if File_v1_observation_dates_proto != nil {
		return
	}
	file_stat_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_observation_dates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityCount); i {
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
		file_v1_observation_dates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservationDates); i {
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
		file_v1_observation_dates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariableObservationDates); i {
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
		file_v1_observation_dates_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkObservationDatesLinkedRequest); i {
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
		file_v1_observation_dates_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkObservationDatesLinkedResponse); i {
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
			RawDescriptor: file_v1_observation_dates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_observation_dates_proto_goTypes,
		DependencyIndexes: file_v1_observation_dates_proto_depIdxs,
		MessageInfos:      file_v1_observation_dates_proto_msgTypes,
	}.Build()
	File_v1_observation_dates_proto = out.File
	file_v1_observation_dates_proto_rawDesc = nil
	file_v1_observation_dates_proto_goTypes = nil
	file_v1_observation_dates_proto_depIdxs = nil
}
