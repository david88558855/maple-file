// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/file/repo.proto

package file

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

type Driver int32

const (
	Driver_DRIVER_UNSPECIFIED Driver = 0
	Driver_DRIVER_LOCAL       Driver = 1
	Driver_DRIVER_WEBDAV      Driver = 2
)

// Enum value maps for Driver.
var (
	Driver_name = map[int32]string{
		0: "DRIVER_UNSPECIFIED",
		1: "DRIVER_LOCAL",
		2: "DRIVER_WEBDAV",
	}
	Driver_value = map[string]int32{
		"DRIVER_UNSPECIFIED": 0,
		"DRIVER_LOCAL":       1,
		"DRIVER_WEBDAV":      2,
	}
)

func (x Driver) Enum() *Driver {
	p := new(Driver)
	*p = x
	return p
}

func (x Driver) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Driver) Descriptor() protoreflect.EnumDescriptor {
	return file_api_file_repo_proto_enumTypes[0].Descriptor()
}

func (Driver) Type() protoreflect.EnumType {
	return &file_api_file_repo_proto_enumTypes[0]
}

func (x Driver) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Driver.Descriptor instead.
func (Driver) EnumDescriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{0}
}

type Repo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"primary_key;auto_increment"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key;auto_increment"`
	// @gotags: gorm:"serializer:protobuf_timestamp;type:datetime"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"serializer:protobuf_timestamp;type:datetime"`
	// @gotags: gorm:"serializer:protobuf_timestamp;type:datetime"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"serializer:protobuf_timestamp;type:datetime"`
	// @gotags: gorm:"not null;unique"
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" gorm:"not null;unique"`
	Desc   string `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	Path   string `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Driver string `protobuf:"bytes,7,opt,name=driver,proto3" json:"driver,omitempty"`
	Option string `protobuf:"bytes,8,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *Repo) Reset() {
	*x = Repo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repo) ProtoMessage() {}

func (x *Repo) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repo.ProtoReflect.Descriptor instead.
func (*Repo) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{0}
}

func (x *Repo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Repo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Repo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Repo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Repo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Repo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Repo) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Repo) GetOption() string {
	if x != nil {
		return x.Option
	}
	return ""
}

type ListReposRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum   int32  `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListReposRequest) Reset() {
	*x = ListReposRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReposRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReposRequest) ProtoMessage() {}

func (x *ListReposRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReposRequest.ProtoReflect.Descriptor instead.
func (*ListReposRequest) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{1}
}

func (x *ListReposRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListReposRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListReposResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Repo `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListReposResponse) Reset() {
	*x = ListReposResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReposResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReposResponse) ProtoMessage() {}

func (x *ListReposResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReposResponse.ProtoReflect.Descriptor instead.
func (*ListReposResponse) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{2}
}

func (x *ListReposResponse) GetResults() []*Repo {
	if x != nil {
		return x.Results
	}
	return nil
}

type CreateRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *Repo `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *CreateRepoRequest) Reset() {
	*x = CreateRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepoRequest) ProtoMessage() {}

func (x *CreateRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepoRequest.ProtoReflect.Descriptor instead.
func (*CreateRepoRequest) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRepoRequest) GetPayload() *Repo {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CreateRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Repo `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateRepoResponse) Reset() {
	*x = CreateRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepoResponse) ProtoMessage() {}

func (x *CreateRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepoResponse.ProtoReflect.Descriptor instead.
func (*CreateRepoResponse) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRepoResponse) GetResult() *Repo {
	if x != nil {
		return x.Result
	}
	return nil
}

type TestRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *Repo `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *TestRepoRequest) Reset() {
	*x = TestRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRepoRequest) ProtoMessage() {}

func (x *TestRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRepoRequest.ProtoReflect.Descriptor instead.
func (*TestRepoRequest) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{5}
}

func (x *TestRepoRequest) GetPayload() *Repo {
	if x != nil {
		return x.Payload
	}
	return nil
}

type TestRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TestRepoResponse) Reset() {
	*x = TestRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRepoResponse) ProtoMessage() {}

func (x *TestRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRepoResponse.ProtoReflect.Descriptor instead.
func (*TestRepoResponse) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{6}
}

func (x *TestRepoResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UpdateRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *Repo `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *UpdateRepoRequest) Reset() {
	*x = UpdateRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRepoRequest) ProtoMessage() {}

func (x *UpdateRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRepoRequest.ProtoReflect.Descriptor instead.
func (*UpdateRepoRequest) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRepoRequest) GetPayload() *Repo {
	if x != nil {
		return x.Payload
	}
	return nil
}

type UpdateRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Repo `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateRepoResponse) Reset() {
	*x = UpdateRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRepoResponse) ProtoMessage() {}

func (x *UpdateRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRepoResponse.ProtoReflect.Descriptor instead.
func (*UpdateRepoResponse) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRepoResponse) GetResult() *Repo {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRepoRequest) Reset() {
	*x = DeleteRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRepoRequest) ProtoMessage() {}

func (x *DeleteRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRepoRequest.ProtoReflect.Descriptor instead.
func (*DeleteRepoRequest) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRepoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRepoResponse) Reset() {
	*x = DeleteRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_repo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRepoResponse) ProtoMessage() {}

func (x *DeleteRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_repo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRepoResponse.ProtoReflect.Descriptor instead.
func (*DeleteRepoResponse) Descriptor() ([]byte, []int) {
	return file_api_file_repo_proto_rawDescGZIP(), []int{10}
}

var File_api_file_repo_proto protoreflect.FileDescriptor

var file_api_file_repo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf8, 0x01, 0x0a, 0x04, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x3d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x3c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x45, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x52, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x52,
	0x49, 0x56, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x44, 0x52, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x57, 0x45, 0x42, 0x44, 0x41, 0x56, 0x10, 0x02, 0x42,
	0x99, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x42, 0x09, 0x52, 0x65, 0x70, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x6e, 0x6d, 0x61, 0x70,
	0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x6c, 0x65, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0xa2, 0x02, 0x03, 0x41,
	0x46, 0x58, 0xaa, 0x02, 0x08, 0x41, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0xca, 0x02, 0x08,
	0x41, 0x70, 0x69, 0x5c, 0x46, 0x69, 0x6c, 0x65, 0xe2, 0x02, 0x14, 0x41, 0x70, 0x69, 0x5c, 0x46,
	0x69, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x09, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x46, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_file_repo_proto_rawDescOnce sync.Once
	file_api_file_repo_proto_rawDescData = file_api_file_repo_proto_rawDesc
)

func file_api_file_repo_proto_rawDescGZIP() []byte {
	file_api_file_repo_proto_rawDescOnce.Do(func() {
		file_api_file_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_file_repo_proto_rawDescData)
	})
	return file_api_file_repo_proto_rawDescData
}

var file_api_file_repo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_file_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_file_repo_proto_goTypes = []interface{}{
	(Driver)(0),                   // 0: api.file.Driver
	(*Repo)(nil),                  // 1: api.file.Repo
	(*ListReposRequest)(nil),      // 2: api.file.ListReposRequest
	(*ListReposResponse)(nil),     // 3: api.file.ListReposResponse
	(*CreateRepoRequest)(nil),     // 4: api.file.CreateRepoRequest
	(*CreateRepoResponse)(nil),    // 5: api.file.CreateRepoResponse
	(*TestRepoRequest)(nil),       // 6: api.file.TestRepoRequest
	(*TestRepoResponse)(nil),      // 7: api.file.TestRepoResponse
	(*UpdateRepoRequest)(nil),     // 8: api.file.UpdateRepoRequest
	(*UpdateRepoResponse)(nil),    // 9: api.file.UpdateRepoResponse
	(*DeleteRepoRequest)(nil),     // 10: api.file.DeleteRepoRequest
	(*DeleteRepoResponse)(nil),    // 11: api.file.DeleteRepoResponse
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_api_file_repo_proto_depIdxs = []int32{
	12, // 0: api.file.Repo.created_at:type_name -> google.protobuf.Timestamp
	12, // 1: api.file.Repo.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: api.file.ListReposResponse.results:type_name -> api.file.Repo
	1,  // 3: api.file.CreateRepoRequest.payload:type_name -> api.file.Repo
	1,  // 4: api.file.CreateRepoResponse.result:type_name -> api.file.Repo
	1,  // 5: api.file.TestRepoRequest.payload:type_name -> api.file.Repo
	1,  // 6: api.file.UpdateRepoRequest.payload:type_name -> api.file.Repo
	1,  // 7: api.file.UpdateRepoResponse.result:type_name -> api.file.Repo
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_file_repo_proto_init() }
func file_api_file_repo_proto_init() {
	if File_api_file_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_file_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repo); i {
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
		file_api_file_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReposRequest); i {
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
		file_api_file_repo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReposResponse); i {
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
		file_api_file_repo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepoRequest); i {
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
		file_api_file_repo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepoResponse); i {
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
		file_api_file_repo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRepoRequest); i {
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
		file_api_file_repo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRepoResponse); i {
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
		file_api_file_repo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRepoRequest); i {
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
		file_api_file_repo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRepoResponse); i {
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
		file_api_file_repo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRepoRequest); i {
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
		file_api_file_repo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRepoResponse); i {
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
			RawDescriptor: file_api_file_repo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_file_repo_proto_goTypes,
		DependencyIndexes: file_api_file_repo_proto_depIdxs,
		EnumInfos:         file_api_file_repo_proto_enumTypes,
		MessageInfos:      file_api_file_repo_proto_msgTypes,
	}.Build()
	File_api_file_repo_proto = out.File
	file_api_file_repo_proto_rawDesc = nil
	file_api_file_repo_proto_goTypes = nil
	file_api_file_repo_proto_depIdxs = nil
}
