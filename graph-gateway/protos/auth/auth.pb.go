// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: auth/auth.proto

package auth

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Validate2FARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TwoFaCode string `protobuf:"bytes,1,opt,name=two_fa_code,json=twoFaCode,proto3" json:"two_fa_code,omitempty"`
}

func (x *Validate2FARequest) Reset() {
	*x = Validate2FARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validate2FARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validate2FARequest) ProtoMessage() {}

func (x *Validate2FARequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validate2FARequest.ProtoReflect.Descriptor instead.
func (*Validate2FARequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *Validate2FARequest) GetTwoFaCode() string {
	if x != nil {
		return x.TwoFaCode
	}
	return ""
}

type RegisterAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RegisterAccountRequest) Reset() {
	*x = RegisterAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAccountRequest) ProtoMessage() {}

func (x *RegisterAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAccountRequest.ProtoReflect.Descriptor instead.
func (*RegisterAccountRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ValidateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewUserToken string `protobuf:"bytes,1,opt,name=new_user_token,json=newUserToken,proto3" json:"new_user_token,omitempty"`
}

func (x *ValidateEmailRequest) Reset() {
	*x = ValidateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateEmailRequest) ProtoMessage() {}

func (x *ValidateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateEmailRequest.ProtoReflect.Descriptor instead.
func (*ValidateEmailRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateEmailRequest) GetNewUserToken() string {
	if x != nil {
		return x.NewUserToken
	}
	return ""
}

type SetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password          string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirmed string `protobuf:"bytes,2,opt,name=password_confirmed,json=passwordConfirmed,proto3" json:"password_confirmed,omitempty"`
}

func (x *SetPasswordRequest) Reset() {
	*x = SetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordRequest) ProtoMessage() {}

func (x *SetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordRequest.ProtoReflect.Descriptor instead.
func (*SetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *SetPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SetPasswordRequest) GetPasswordConfirmed() string {
	if x != nil {
		return x.PasswordConfirmed
	}
	return ""
}

type Set2FARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode string `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *Set2FARequest) Reset() {
	*x = Set2FARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set2FARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set2FARequest) ProtoMessage() {}

func (x *Set2FARequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set2FARequest.ProtoReflect.Descriptor instead.
func (*Set2FARequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *Set2FARequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type ForgotPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ForgotPasswordRequest) Reset() {
	*x = ForgotPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordRequest) ProtoMessage() {}

func (x *ForgotPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordRequest.ProtoReflect.Descriptor instead.
func (*ForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *ForgotPasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// If reset_token is missing or empty, there must be a valid JWT
type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResetToken        *string `protobuf:"bytes,1,opt,name=reset_token,json=resetToken,proto3,oneof" json:"reset_token,omitempty"`
	Password          string  `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirmed string  `protobuf:"bytes,3,opt,name=password_confirmed,json=passwordConfirmed,proto3" json:"password_confirmed,omitempty"`
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{7}
}

func (x *ResetPasswordRequest) GetResetToken() string {
	if x != nil && x.ResetToken != nil {
		return *x.ResetToken
	}
	return ""
}

func (x *ResetPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ResetPasswordRequest) GetPasswordConfirmed() string {
	if x != nil {
		return x.PasswordConfirmed
	}
	return ""
}

var File_auth_auth_proto protoreflect.FileDescriptor

var file_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x34, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x32, 0x46, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0b,
	0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x77, 0x6f, 0x46, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2e, 0x0a, 0x16,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3c, 0x0a, 0x14,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5f, 0x0a, 0x12, 0x53, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x12,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x22, 0x23, 0x0a, 0x0d, 0x53,
	0x65, 0x74, 0x32, 0x46, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x22, 0x2d, 0x0a, 0x15, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x97, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc8, 0x04, 0x0a, 0x04, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x33, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x32, 0x46, 0x41, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x32, 0x46, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x43, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x32, 0x46,
	0x41, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x74, 0x32, 0x46, 0x41, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38,
	0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67,
	0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x43, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_auth_proto_rawDescOnce sync.Once
	file_auth_auth_proto_rawDescData = file_auth_auth_proto_rawDesc
)

func file_auth_auth_proto_rawDescGZIP() []byte {
	file_auth_auth_proto_rawDescOnce.Do(func() {
		file_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_auth_proto_rawDescData)
	})
	return file_auth_auth_proto_rawDescData
}

var file_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_auth_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),           // 0: auth.LoginRequest
	(*Validate2FARequest)(nil),     // 1: auth.Validate2FARequest
	(*RegisterAccountRequest)(nil), // 2: auth.RegisterAccountRequest
	(*ValidateEmailRequest)(nil),   // 3: auth.ValidateEmailRequest
	(*SetPasswordRequest)(nil),     // 4: auth.SetPasswordRequest
	(*Set2FARequest)(nil),          // 5: auth.Set2FARequest
	(*ForgotPasswordRequest)(nil),  // 6: auth.ForgotPasswordRequest
	(*ResetPasswordRequest)(nil),   // 7: auth.ResetPasswordRequest
	(*empty.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_auth_auth_proto_depIdxs = []int32{
	0, // 0: auth.Auth.Login:input_type -> auth.LoginRequest
	1, // 1: auth.Auth.Validate2FA:input_type -> auth.Validate2FARequest
	2, // 2: auth.Auth.RegisterAccount:input_type -> auth.RegisterAccountRequest
	3, // 3: auth.Auth.ValidateEmail:input_type -> auth.ValidateEmailRequest
	4, // 4: auth.Auth.SetPassword:input_type -> auth.SetPasswordRequest
	5, // 5: auth.Auth.Set2FA:input_type -> auth.Set2FARequest
	8, // 6: auth.Auth.Logout:input_type -> google.protobuf.Empty
	6, // 7: auth.Auth.ForgotPassword:input_type -> auth.ForgotPasswordRequest
	7, // 8: auth.Auth.ResetPassword:input_type -> auth.ResetPasswordRequest
	8, // 9: auth.Auth.Login:output_type -> google.protobuf.Empty
	8, // 10: auth.Auth.Validate2FA:output_type -> google.protobuf.Empty
	8, // 11: auth.Auth.RegisterAccount:output_type -> google.protobuf.Empty
	8, // 12: auth.Auth.ValidateEmail:output_type -> google.protobuf.Empty
	8, // 13: auth.Auth.SetPassword:output_type -> google.protobuf.Empty
	8, // 14: auth.Auth.Set2FA:output_type -> google.protobuf.Empty
	8, // 15: auth.Auth.Logout:output_type -> google.protobuf.Empty
	8, // 16: auth.Auth.ForgotPassword:output_type -> google.protobuf.Empty
	8, // 17: auth.Auth.ResetPassword:output_type -> google.protobuf.Empty
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validate2FARequest); i {
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
		file_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAccountRequest); i {
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
		file_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateEmailRequest); i {
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
		file_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordRequest); i {
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
		file_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set2FARequest); i {
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
		file_auth_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPasswordRequest); i {
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
		file_auth_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordRequest); i {
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
	file_auth_auth_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
		MessageInfos:      file_auth_auth_proto_msgTypes,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_rawDesc = nil
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}
