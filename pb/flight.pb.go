// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: flight.proto

package pb

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

type CreateFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string                 `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int32                  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
}

func (x *CreateFlightRequest) Reset() {
	*x = CreateFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlightRequest) ProtoMessage() {}

func (x *CreateFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlightRequest.ProtoReflect.Descriptor instead.
func (*CreateFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFlightRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateFlightRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFlightRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreateFlightRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *CreateFlightRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *CreateFlightRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateFlightRequest) GetAvailableSlot() int32 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

type Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Date          string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Status        string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int32  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
}

func (x *Flight) Reset() {
	*x = Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flight) ProtoMessage() {}

func (x *Flight) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flight.ProtoReflect.Descriptor instead.
func (*Flight) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{1}
}

func (x *Flight) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Flight) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flight) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Flight) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Flight) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Flight) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Flight) GetAvailableSlot() int32 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

type CreateFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Date          string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Status        string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int32  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
}

func (x *CreateFlightResponse) Reset() {
	*x = CreateFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlightResponse) ProtoMessage() {}

func (x *CreateFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlightResponse.ProtoReflect.Descriptor instead.
func (*CreateFlightResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFlightResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateFlightResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFlightResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreateFlightResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *CreateFlightResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *CreateFlightResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateFlightResponse) GetAvailableSlot() int32 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

type UpdateFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string                 `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int32                  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
}

func (x *UpdateFlightRequest) Reset() {
	*x = UpdateFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlightRequest) ProtoMessage() {}

func (x *UpdateFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlightRequest.ProtoReflect.Descriptor instead.
func (*UpdateFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFlightRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateFlightRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateFlightRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *UpdateFlightRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *UpdateFlightRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *UpdateFlightRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateFlightRequest) GetAvailableSlot() int32 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

type UpdateFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Date          string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Status        string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int32  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
}

func (x *UpdateFlightResponse) Reset() {
	*x = UpdateFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlightResponse) ProtoMessage() {}

func (x *UpdateFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlightResponse.ProtoReflect.Descriptor instead.
func (*UpdateFlightResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateFlightResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateFlightResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateFlightResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *UpdateFlightResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *UpdateFlightResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *UpdateFlightResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateFlightResponse) GetAvailableSlot() int32 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

type SearchFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SearchFlightRequest) Reset() {
	*x = SearchFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlightRequest) ProtoMessage() {}

func (x *SearchFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlightRequest.ProtoReflect.Descriptor instead.
func (*SearchFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{5}
}

func (x *SearchFlightRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Date          string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Status        string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int32  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
}

func (x *SearchFlightResponse) Reset() {
	*x = SearchFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlightResponse) ProtoMessage() {}

func (x *SearchFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlightResponse.ProtoReflect.Descriptor instead.
func (*SearchFlightResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{6}
}

func (x *SearchFlightResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SearchFlightResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchFlightResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SearchFlightResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SearchFlightResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *SearchFlightResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SearchFlightResponse) GetAvailableSlot() int32 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

var File_flight_proto protoreflect.FileDescriptor

var file_flight_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x73, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x06,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x6f,
	0x74, 0x22, 0xb1, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0xcc, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x6c, 0x6f, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73,
	0x6c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xb1, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x6c, 0x6f, 0x74, 0x32, 0x8e, 0x02, 0x0a, 0x0d, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x73, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x73, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x73, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x73, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x73, 0x6d, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x73, 0x6d, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flight_proto_rawDescOnce sync.Once
	file_flight_proto_rawDescData = file_flight_proto_rawDesc
)

func file_flight_proto_rawDescGZIP() []byte {
	file_flight_proto_rawDescOnce.Do(func() {
		file_flight_proto_rawDescData = protoimpl.X.CompressGZIP(file_flight_proto_rawDescData)
	})
	return file_flight_proto_rawDescData
}

var file_flight_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_flight_proto_goTypes = []interface{}{
	(*CreateFlightRequest)(nil),   // 0: booking_asm.CreateFlightRequest
	(*Flight)(nil),                // 1: booking_asm.Flight
	(*CreateFlightResponse)(nil),  // 2: booking_asm.CreateFlightResponse
	(*UpdateFlightRequest)(nil),   // 3: booking_asm.UpdateFlightRequest
	(*UpdateFlightResponse)(nil),  // 4: booking_asm.UpdateFlightResponse
	(*SearchFlightRequest)(nil),   // 5: booking_asm.SearchFlightRequest
	(*SearchFlightResponse)(nil),  // 6: booking_asm.SearchFlightResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_flight_proto_depIdxs = []int32{
	7, // 0: booking_asm.CreateFlightRequest.date:type_name -> google.protobuf.Timestamp
	7, // 1: booking_asm.UpdateFlightRequest.date:type_name -> google.protobuf.Timestamp
	0, // 2: booking_asm.FlightService.CreateFlight:input_type -> booking_asm.CreateFlightRequest
	3, // 3: booking_asm.FlightService.UpdateFlight:input_type -> booking_asm.UpdateFlightRequest
	5, // 4: booking_asm.FlightService.SearchFlight:input_type -> booking_asm.SearchFlightRequest
	2, // 5: booking_asm.FlightService.CreateFlight:output_type -> booking_asm.CreateFlightResponse
	4, // 6: booking_asm.FlightService.UpdateFlight:output_type -> booking_asm.UpdateFlightResponse
	6, // 7: booking_asm.FlightService.SearchFlight:output_type -> booking_asm.SearchFlightResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_flight_proto_init() }
func file_flight_proto_init() {
	if File_flight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFlightRequest); i {
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
		file_flight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flight); i {
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
		file_flight_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFlightResponse); i {
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
		file_flight_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlightRequest); i {
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
		file_flight_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlightResponse); i {
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
		file_flight_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlightRequest); i {
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
		file_flight_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlightResponse); i {
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
			RawDescriptor: file_flight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flight_proto_goTypes,
		DependencyIndexes: file_flight_proto_depIdxs,
		MessageInfos:      file_flight_proto_msgTypes,
	}.Build()
	File_flight_proto = out.File
	file_flight_proto_rawDesc = nil
	file_flight_proto_goTypes = nil
	file_flight_proto_depIdxs = nil
}
