// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/user/user.proto

package user

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

// User list
type UserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserListReq) Reset() {
	*x = UserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReq) ProtoMessage() {}

func (x *UserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListReq.ProtoReflect.Descriptor instead.
func (*UserListReq) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserListReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type UserListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *UserListRes) Reset() {
	*x = UserListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListRes) ProtoMessage() {}

func (x *UserListRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListRes.ProtoReflect.Descriptor instead.
func (*UserListRes) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserListRes) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

// User store
type UserStoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Role     string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserStoreReq) Reset() {
	*x = UserStoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStoreReq) ProtoMessage() {}

func (x *UserStoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStoreReq.ProtoReflect.Descriptor instead.
func (*UserStoreReq) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserStoreReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserStoreReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserStoreReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserStoreReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type UserStoreRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserStoreRes) Reset() {
	*x = UserStoreRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStoreRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStoreRes) ProtoMessage() {}

func (x *UserStoreRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStoreRes.ProtoReflect.Descriptor instead.
func (*UserStoreRes) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{3}
}

// User detail
type UserDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserDetailReq) Reset() {
	*x = UserDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailReq) ProtoMessage() {}

func (x *UserDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailReq.ProtoReflect.Descriptor instead.
func (*UserDetailReq) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserDetailReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserDetailRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserDetailRes) Reset() {
	*x = UserDetailRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailRes) ProtoMessage() {}

func (x *UserDetailRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailRes.ProtoReflect.Descriptor instead.
func (*UserDetailRes) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserDetailRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// User update
type UserUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Role     string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserUpdateReq) Reset() {
	*x = UserUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateReq) ProtoMessage() {}

func (x *UserUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateReq.ProtoReflect.Descriptor instead.
func (*UserUpdateReq) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserUpdateReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserUpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserUpdateReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserUpdateReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserUpdateReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type UserUpdateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserUpdateRes) Reset() {
	*x = UserUpdateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateRes) ProtoMessage() {}

func (x *UserUpdateRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateRes.ProtoReflect.Descriptor instead.
func (*UserUpdateRes) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{7}
}

// User delete
type UserDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserDeleteReq) Reset() {
	*x = UserDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteReq) ProtoMessage() {}

func (x *UserDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteReq.ProtoReflect.Descriptor instead.
func (*UserDeleteReq) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserDeleteReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserDeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserDeleteRes) Reset() {
	*x = UserDeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteRes) ProtoMessage() {}

func (x *UserDeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteRes.ProtoReflect.Descriptor instead.
func (*UserDeleteRes) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{9}
}

// ping
type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[10]
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
	return file_proto_user_user_proto_rawDescGZIP(), []int{10}
}

type PingRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingRes) Reset() {
	*x = PingRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRes) ProtoMessage() {}

func (x *PingRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRes.ProtoReflect.Descriptor instead.
func (*PingRes) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{11}
}

func (x *PingRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// user
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Role      string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{12}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// user login
type UserLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLoginReq) Reset() {
	*x = UserLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginReq) ProtoMessage() {}

func (x *UserLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginReq.ProtoReflect.Descriptor instead.
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{13}
}

func (x *UserLoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserLoginRes) Reset() {
	*x = UserLoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRes) ProtoMessage() {}

func (x *UserLoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRes.ProtoReflect.Descriptor instead.
func (*UserLoginRes) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{14}
}

func (x *UserLoginRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_proto_user_user_proto protoreflect.FileDescriptor

var file_proto_user_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x35, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x22, 0x2f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x68, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0x0e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x1f, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2f, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x79, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x0f, 0x0a, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x1f, 0x0a,
	0x0d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0f,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x09, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x22, 0x23, 0x0a, 0x07, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xae, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x40, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x2e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x32, 0xed, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_user_proto_rawDescOnce sync.Once
	file_proto_user_user_proto_rawDescData = file_proto_user_user_proto_rawDesc
)

func file_proto_user_user_proto_rawDescGZIP() []byte {
	file_proto_user_user_proto_rawDescOnce.Do(func() {
		file_proto_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_user_proto_rawDescData)
	})
	return file_proto_user_user_proto_rawDescData
}

var file_proto_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_user_user_proto_goTypes = []interface{}{
	(*UserListReq)(nil),   // 0: user.UserListReq
	(*UserListRes)(nil),   // 1: user.UserListRes
	(*UserStoreReq)(nil),  // 2: user.UserStoreReq
	(*UserStoreRes)(nil),  // 3: user.UserStoreRes
	(*UserDetailReq)(nil), // 4: user.UserDetailReq
	(*UserDetailRes)(nil), // 5: user.UserDetailRes
	(*UserUpdateReq)(nil), // 6: user.UserUpdateReq
	(*UserUpdateRes)(nil), // 7: user.UserUpdateRes
	(*UserDeleteReq)(nil), // 8: user.UserDeleteReq
	(*UserDeleteRes)(nil), // 9: user.UserDeleteRes
	(*PingReq)(nil),       // 10: user.PingReq
	(*PingRes)(nil),       // 11: user.PingRes
	(*User)(nil),          // 12: user.User
	(*UserLoginReq)(nil),  // 13: user.UserLoginReq
	(*UserLoginRes)(nil),  // 14: user.UserLoginRes
}
var file_proto_user_user_proto_depIdxs = []int32{
	12, // 0: user.UserListRes.Users:type_name -> user.User
	12, // 1: user.UserDetailRes.User:type_name -> user.User
	12, // 2: user.UserLoginRes.User:type_name -> user.User
	10, // 3: user.UserService.Ping:input_type -> user.PingReq
	0,  // 4: user.UserService.List:input_type -> user.UserListReq
	2,  // 5: user.UserService.Store:input_type -> user.UserStoreReq
	4,  // 6: user.UserService.Detail:input_type -> user.UserDetailReq
	6,  // 7: user.UserService.Update:input_type -> user.UserUpdateReq
	8,  // 8: user.UserService.Delete:input_type -> user.UserDeleteReq
	13, // 9: user.UserService.Login:input_type -> user.UserLoginReq
	11, // 10: user.UserService.Ping:output_type -> user.PingRes
	1,  // 11: user.UserService.List:output_type -> user.UserListRes
	3,  // 12: user.UserService.Store:output_type -> user.UserStoreRes
	5,  // 13: user.UserService.Detail:output_type -> user.UserDetailRes
	7,  // 14: user.UserService.Update:output_type -> user.UserUpdateRes
	9,  // 15: user.UserService.Delete:output_type -> user.UserDeleteRes
	14, // 16: user.UserService.Login:output_type -> user.UserLoginRes
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_user_user_proto_init() }
func file_proto_user_user_proto_init() {
	if File_proto_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListReq); i {
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
		file_proto_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListRes); i {
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
		file_proto_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStoreReq); i {
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
		file_proto_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStoreRes); i {
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
		file_proto_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailReq); i {
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
		file_proto_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailRes); i {
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
		file_proto_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateReq); i {
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
		file_proto_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateRes); i {
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
		file_proto_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteReq); i {
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
		file_proto_user_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteRes); i {
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
		file_proto_user_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_user_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRes); i {
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
		file_proto_user_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_user_user_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginReq); i {
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
		file_proto_user_user_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRes); i {
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
			RawDescriptor: file_proto_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_user_proto_goTypes,
		DependencyIndexes: file_proto_user_user_proto_depIdxs,
		MessageInfos:      file_proto_user_user_proto_msgTypes,
	}.Build()
	File_proto_user_user_proto = out.File
	file_proto_user_user_proto_rawDesc = nil
	file_proto_user_user_proto_goTypes = nil
	file_proto_user_user_proto_depIdxs = nil
}
