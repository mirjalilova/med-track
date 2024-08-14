// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: lifestyle_data.proto

package health_sync

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

type LifeStyle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DataType      string  `protobuf:"bytes,3,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	DataValue     string  `protobuf:"bytes,4,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
	HeartRate     int32   `protobuf:"varint,5,opt,name=heart_rate,json=heartRate,proto3" json:"heart_rate,omitempty"`
	BloodPressure string  `protobuf:"bytes,6,opt,name=blood_pressure,json=bloodPressure,proto3" json:"blood_pressure,omitempty"`
	Weight        float32 `protobuf:"fixed32,8,opt,name=weight,proto3" json:"weight,omitempty"`
	Height        float32 `protobuf:"fixed32,9,opt,name=height,proto3" json:"height,omitempty"`
	Temperature   float32 `protobuf:"fixed32,10,opt,name=temperature,proto3" json:"temperature,omitempty"`
	StressLevel   int32   `protobuf:"varint,11,opt,name=stress_level,json=stressLevel,proto3" json:"stress_level,omitempty"`
	SleepDuration float32 `protobuf:"fixed32,12,opt,name=sleep_duration,json=sleepDuration,proto3" json:"sleep_duration,omitempty"`
	RecordedDate  string  `protobuf:"bytes,13,opt,name=recorded_date,json=recordedDate,proto3" json:"recorded_date,omitempty"`
	CreatedAt     string  `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     int32   `protobuf:"varint,16,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *LifeStyle) Reset() {
	*x = LifeStyle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifestyle_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifeStyle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifeStyle) ProtoMessage() {}

func (x *LifeStyle) ProtoReflect() protoreflect.Message {
	mi := &file_lifestyle_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifeStyle.ProtoReflect.Descriptor instead.
func (*LifeStyle) Descriptor() ([]byte, []int) {
	return file_lifestyle_data_proto_rawDescGZIP(), []int{0}
}

func (x *LifeStyle) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LifeStyle) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LifeStyle) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *LifeStyle) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *LifeStyle) GetHeartRate() int32 {
	if x != nil {
		return x.HeartRate
	}
	return 0
}

func (x *LifeStyle) GetBloodPressure() string {
	if x != nil {
		return x.BloodPressure
	}
	return ""
}

func (x *LifeStyle) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *LifeStyle) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *LifeStyle) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *LifeStyle) GetStressLevel() int32 {
	if x != nil {
		return x.StressLevel
	}
	return 0
}

func (x *LifeStyle) GetSleepDuration() float32 {
	if x != nil {
		return x.SleepDuration
	}
	return 0
}

func (x *LifeStyle) GetRecordedDate() string {
	if x != nil {
		return x.RecordedDate
	}
	return ""
}

func (x *LifeStyle) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *LifeStyle) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *LifeStyle) GetDeletedAt() int32 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type LifestyleCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DataType      string  `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	DataValue     string  `protobuf:"bytes,3,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
	HeartRate     int32   `protobuf:"varint,4,opt,name=heart_rate,json=heartRate,proto3" json:"heart_rate,omitempty"`
	BloodPressure string  `protobuf:"bytes,5,opt,name=blood_pressure,json=bloodPressure,proto3" json:"blood_pressure,omitempty"`
	Weight        float32 `protobuf:"fixed32,6,opt,name=weight,proto3" json:"weight,omitempty"`
	Height        float32 `protobuf:"fixed32,7,opt,name=height,proto3" json:"height,omitempty"`
	Temperature   float32 `protobuf:"fixed32,8,opt,name=temperature,proto3" json:"temperature,omitempty"`
	StressLevel   int32   `protobuf:"varint,9,opt,name=stress_level,json=stressLevel,proto3" json:"stress_level,omitempty"`
	SleepDuration float32 `protobuf:"fixed32,10,opt,name=sleep_duration,json=sleepDuration,proto3" json:"sleep_duration,omitempty"`
	RecordedDate  string  `protobuf:"bytes,11,opt,name=recorded_date,json=recordedDate,proto3" json:"recorded_date,omitempty"`
}

func (x *LifestyleCreate) Reset() {
	*x = LifestyleCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifestyle_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifestyleCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifestyleCreate) ProtoMessage() {}

func (x *LifestyleCreate) ProtoReflect() protoreflect.Message {
	mi := &file_lifestyle_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifestyleCreate.ProtoReflect.Descriptor instead.
func (*LifestyleCreate) Descriptor() ([]byte, []int) {
	return file_lifestyle_data_proto_rawDescGZIP(), []int{1}
}

func (x *LifestyleCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LifestyleCreate) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *LifestyleCreate) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *LifestyleCreate) GetHeartRate() int32 {
	if x != nil {
		return x.HeartRate
	}
	return 0
}

func (x *LifestyleCreate) GetBloodPressure() string {
	if x != nil {
		return x.BloodPressure
	}
	return ""
}

func (x *LifestyleCreate) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *LifestyleCreate) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *LifestyleCreate) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *LifestyleCreate) GetStressLevel() int32 {
	if x != nil {
		return x.StressLevel
	}
	return 0
}

func (x *LifestyleCreate) GetSleepDuration() float32 {
	if x != nil {
		return x.SleepDuration
	}
	return 0
}

func (x *LifestyleCreate) GetRecordedDate() string {
	if x != nil {
		return x.RecordedDate
	}
	return ""
}

type LifestyleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DataType      string  `protobuf:"bytes,3,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	DataValue     string  `protobuf:"bytes,4,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
	HeartRate     int32   `protobuf:"varint,5,opt,name=heart_rate,json=heartRate,proto3" json:"heart_rate,omitempty"`
	BloodPressure string  `protobuf:"bytes,6,opt,name=blood_pressure,json=bloodPressure,proto3" json:"blood_pressure,omitempty"`
	Weight        float32 `protobuf:"fixed32,7,opt,name=weight,proto3" json:"weight,omitempty"`
	Height        float32 `protobuf:"fixed32,8,opt,name=height,proto3" json:"height,omitempty"`
	Temperature   float32 `protobuf:"fixed32,9,opt,name=temperature,proto3" json:"temperature,omitempty"`
	StressLevel   int32   `protobuf:"varint,10,opt,name=stress_level,json=stressLevel,proto3" json:"stress_level,omitempty"`
	SleepDuration float32 `protobuf:"fixed32,11,opt,name=sleep_duration,json=sleepDuration,proto3" json:"sleep_duration,omitempty"`
	RecordedDate  string  `protobuf:"bytes,12,opt,name=recorded_date,json=recordedDate,proto3" json:"recorded_date,omitempty"`
}

func (x *LifestyleRes) Reset() {
	*x = LifestyleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifestyle_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifestyleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifestyleRes) ProtoMessage() {}

func (x *LifestyleRes) ProtoReflect() protoreflect.Message {
	mi := &file_lifestyle_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifestyleRes.ProtoReflect.Descriptor instead.
func (*LifestyleRes) Descriptor() ([]byte, []int) {
	return file_lifestyle_data_proto_rawDescGZIP(), []int{2}
}

func (x *LifestyleRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LifestyleRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LifestyleRes) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *LifestyleRes) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *LifestyleRes) GetHeartRate() int32 {
	if x != nil {
		return x.HeartRate
	}
	return 0
}

func (x *LifestyleRes) GetBloodPressure() string {
	if x != nil {
		return x.BloodPressure
	}
	return ""
}

func (x *LifestyleRes) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *LifestyleRes) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *LifestyleRes) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *LifestyleRes) GetStressLevel() int32 {
	if x != nil {
		return x.StressLevel
	}
	return 0
}

func (x *LifestyleRes) GetSleepDuration() float32 {
	if x != nil {
		return x.SleepDuration
	}
	return 0
}

func (x *LifestyleRes) GetRecordedDate() string {
	if x != nil {
		return x.RecordedDate
	}
	return ""
}

type LifestyleUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataType      string  `protobuf:"bytes,1,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	DataValue     string  `protobuf:"bytes,2,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
	HeartRate     int32   `protobuf:"varint,3,opt,name=heart_rate,json=heartRate,proto3" json:"heart_rate,omitempty"`
	BloodPressure string  `protobuf:"bytes,4,opt,name=blood_pressure,json=bloodPressure,proto3" json:"blood_pressure,omitempty"`
	Weight        float32 `protobuf:"fixed32,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Height        float32 `protobuf:"fixed32,6,opt,name=height,proto3" json:"height,omitempty"`
	Temperature   float32 `protobuf:"fixed32,7,opt,name=temperature,proto3" json:"temperature,omitempty"`
	StressLevel   int32   `protobuf:"varint,8,opt,name=stress_level,json=stressLevel,proto3" json:"stress_level,omitempty"`
	SleepDuration float32 `protobuf:"fixed32,9,opt,name=sleep_duration,json=sleepDuration,proto3" json:"sleep_duration,omitempty"`
	Id            string  `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LifestyleUpdate) Reset() {
	*x = LifestyleUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifestyle_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifestyleUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifestyleUpdate) ProtoMessage() {}

func (x *LifestyleUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_lifestyle_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifestyleUpdate.ProtoReflect.Descriptor instead.
func (*LifestyleUpdate) Descriptor() ([]byte, []int) {
	return file_lifestyle_data_proto_rawDescGZIP(), []int{3}
}

func (x *LifestyleUpdate) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *LifestyleUpdate) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *LifestyleUpdate) GetHeartRate() int32 {
	if x != nil {
		return x.HeartRate
	}
	return 0
}

func (x *LifestyleUpdate) GetBloodPressure() string {
	if x != nil {
		return x.BloodPressure
	}
	return ""
}

func (x *LifestyleUpdate) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *LifestyleUpdate) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *LifestyleUpdate) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *LifestyleUpdate) GetStressLevel() int32 {
	if x != nil {
		return x.StressLevel
	}
	return 0
}

func (x *LifestyleUpdate) GetSleepDuration() float32 {
	if x != nil {
		return x.SleepDuration
	}
	return 0
}

func (x *LifestyleUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllLifestyleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DataType string  `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Filter   *Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetAllLifestyleReq) Reset() {
	*x = GetAllLifestyleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifestyle_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllLifestyleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllLifestyleReq) ProtoMessage() {}

func (x *GetAllLifestyleReq) ProtoReflect() protoreflect.Message {
	mi := &file_lifestyle_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllLifestyleReq.ProtoReflect.Descriptor instead.
func (*GetAllLifestyleReq) Descriptor() ([]byte, []int) {
	return file_lifestyle_data_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllLifestyleReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetAllLifestyleReq) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *GetAllLifestyleReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetAllLifestyleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lifestyles []*LifestyleRes `protobuf:"bytes,1,rep,name=lifestyles,proto3" json:"lifestyles,omitempty"`
	Count      int32           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllLifestyleRes) Reset() {
	*x = GetAllLifestyleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifestyle_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllLifestyleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllLifestyleRes) ProtoMessage() {}

func (x *GetAllLifestyleRes) ProtoReflect() protoreflect.Message {
	mi := &file_lifestyle_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllLifestyleRes.ProtoReflect.Descriptor instead.
func (*GetAllLifestyleRes) Descriptor() ([]byte, []int) {
	return file_lifestyle_data_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllLifestyleRes) GetLifestyles() []*LifestyleRes {
	if x != nil {
		return x.Lifestyles
	}
	return nil
}

func (x *GetAllLifestyleRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_lifestyle_data_proto protoreflect.FileDescriptor

var file_lifestyle_data_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x66, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73,
	0x79, 0x6e, 0x63, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd4, 0x03, 0x0a, 0x09, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x5f, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6c, 0x6f,
	0x6f, 0x64, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xed, 0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x50,
	0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xfa, 0x02, 0x0a, 0x0c, 0x4c, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x50, 0x72, 0x65, 0x73,
	0x73, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6c, 0x65,
	0x65, 0x70, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0d, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xbf, 0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x79, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x5f, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6c,
	0x6f, 0x6f, 0x64, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x77, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4c, 0x69, 0x66, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x65, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x79, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x79, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x73, 0x74, 0x79,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x52, 0x0a, 0x6c, 0x69, 0x66, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xc7, 0x02, 0x0a, 0x10, 0x4c, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x79, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x1a, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x1a, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x14, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x14, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x6e, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4c, 0x69, 0x66, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x66, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lifestyle_data_proto_rawDescOnce sync.Once
	file_lifestyle_data_proto_rawDescData = file_lifestyle_data_proto_rawDesc
)

func file_lifestyle_data_proto_rawDescGZIP() []byte {
	file_lifestyle_data_proto_rawDescOnce.Do(func() {
		file_lifestyle_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_lifestyle_data_proto_rawDescData)
	})
	return file_lifestyle_data_proto_rawDescData
}

var file_lifestyle_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_lifestyle_data_proto_goTypes = []any{
	(*LifeStyle)(nil),          // 0: health_sync.LifeStyle
	(*LifestyleCreate)(nil),    // 1: health_sync.LifestyleCreate
	(*LifestyleRes)(nil),       // 2: health_sync.LifestyleRes
	(*LifestyleUpdate)(nil),    // 3: health_sync.LifestyleUpdate
	(*GetAllLifestyleReq)(nil), // 4: health_sync.GetAllLifestyleReq
	(*GetAllLifestyleRes)(nil), // 5: health_sync.GetAllLifestyleRes
	(*Filter)(nil),             // 6: health_sync.Filter
	(*GetById)(nil),            // 7: health_sync.GetById
	(*Void)(nil),               // 8: health_sync.Void
}
var file_lifestyle_data_proto_depIdxs = []int32{
	6, // 0: health_sync.GetAllLifestyleReq.filter:type_name -> health_sync.Filter
	2, // 1: health_sync.GetAllLifestyleRes.lifestyles:type_name -> health_sync.LifestyleRes
	1, // 2: health_sync.LifestyleService.Create:input_type -> health_sync.LifestyleCreate
	3, // 3: health_sync.LifestyleService.Update:input_type -> health_sync.LifestyleUpdate
	7, // 4: health_sync.LifestyleService.Delete:input_type -> health_sync.GetById
	7, // 5: health_sync.LifestyleService.Get:input_type -> health_sync.GetById
	4, // 6: health_sync.LifestyleService.List:input_type -> health_sync.GetAllLifestyleReq
	8, // 7: health_sync.LifestyleService.Create:output_type -> health_sync.Void
	8, // 8: health_sync.LifestyleService.Update:output_type -> health_sync.Void
	8, // 9: health_sync.LifestyleService.Delete:output_type -> health_sync.Void
	2, // 10: health_sync.LifestyleService.Get:output_type -> health_sync.LifestyleRes
	5, // 11: health_sync.LifestyleService.List:output_type -> health_sync.GetAllLifestyleRes
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lifestyle_data_proto_init() }
func file_lifestyle_data_proto_init() {
	if File_lifestyle_data_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lifestyle_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LifeStyle); i {
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
		file_lifestyle_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LifestyleCreate); i {
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
		file_lifestyle_data_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LifestyleRes); i {
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
		file_lifestyle_data_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LifestyleUpdate); i {
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
		file_lifestyle_data_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllLifestyleReq); i {
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
		file_lifestyle_data_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllLifestyleRes); i {
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
			RawDescriptor: file_lifestyle_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lifestyle_data_proto_goTypes,
		DependencyIndexes: file_lifestyle_data_proto_depIdxs,
		MessageInfos:      file_lifestyle_data_proto_msgTypes,
	}.Build()
	File_lifestyle_data_proto = out.File
	file_lifestyle_data_proto_rawDesc = nil
	file_lifestyle_data_proto_goTypes = nil
	file_lifestyle_data_proto_depIdxs = nil
}
