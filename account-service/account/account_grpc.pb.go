// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: account/account.proto

package account

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Account_GetPublicProfileSection_FullMethodName  = "/account.Account/GetPublicProfileSection"
	Account_GetPrivateProfileSection_FullMethodName = "/account.Account/GetPrivateProfileSection"
	Account_UpdateProfileSection_FullMethodName     = "/account.Account/UpdateProfileSection"
	Account_GetPrivateBasicUserInfo_FullMethodName  = "/account.Account/GetPrivateBasicUserInfo"
	Account_GetPublicBasicUserInfo_FullMethodName   = "/account.Account/GetPublicBasicUserInfo"
	Account_UpdateBasicUserInfo_FullMethodName      = "/account.Account/UpdateBasicUserInfo"
	Account_GetCustomerReviews_FullMethodName       = "/account.Account/GetCustomerReviews"
	Account_AddCustomerReview_FullMethodName        = "/account.Account/AddCustomerReview"
)

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	GetPublicProfileSection(ctx context.Context, in *GetPublicProfileSectionRequest, opts ...grpc.CallOption) (*ProfileSectionResponse, error)
	GetPrivateProfileSection(ctx context.Context, in *GetPrivateProfileSectionRequest, opts ...grpc.CallOption) (*ProfileSectionResponse, error)
	UpdateProfileSection(ctx context.Context, in *UpdateProfileSectionRequest, opts ...grpc.CallOption) (*ProfileSectionResponse, error)
	GetPrivateBasicUserInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PrivateBasicUserInfo, error)
	GetPublicBasicUserInfo(ctx context.Context, in *GetPublicBasicUserInfoRequest, opts ...grpc.CallOption) (*PublicBasicUserInfo, error)
	UpdateBasicUserInfo(ctx context.Context, in *UpdatePrivateBasicUserInfoRequest, opts ...grpc.CallOption) (*PrivateBasicUserInfo, error)
	GetCustomerReviews(ctx context.Context, in *GetCustomerReviewsRequest, opts ...grpc.CallOption) (*GetCustomerReviewsResponse, error)
	AddCustomerReview(ctx context.Context, in *AddCustomerReviewRequest, opts ...grpc.CallOption) (*CustomerReview, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) GetPublicProfileSection(ctx context.Context, in *GetPublicProfileSectionRequest, opts ...grpc.CallOption) (*ProfileSectionResponse, error) {
	out := new(ProfileSectionResponse)
	err := c.cc.Invoke(ctx, Account_GetPublicProfileSection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetPrivateProfileSection(ctx context.Context, in *GetPrivateProfileSectionRequest, opts ...grpc.CallOption) (*ProfileSectionResponse, error) {
	out := new(ProfileSectionResponse)
	err := c.cc.Invoke(ctx, Account_GetPrivateProfileSection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateProfileSection(ctx context.Context, in *UpdateProfileSectionRequest, opts ...grpc.CallOption) (*ProfileSectionResponse, error) {
	out := new(ProfileSectionResponse)
	err := c.cc.Invoke(ctx, Account_UpdateProfileSection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetPrivateBasicUserInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PrivateBasicUserInfo, error) {
	out := new(PrivateBasicUserInfo)
	err := c.cc.Invoke(ctx, Account_GetPrivateBasicUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetPublicBasicUserInfo(ctx context.Context, in *GetPublicBasicUserInfoRequest, opts ...grpc.CallOption) (*PublicBasicUserInfo, error) {
	out := new(PublicBasicUserInfo)
	err := c.cc.Invoke(ctx, Account_GetPublicBasicUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateBasicUserInfo(ctx context.Context, in *UpdatePrivateBasicUserInfoRequest, opts ...grpc.CallOption) (*PrivateBasicUserInfo, error) {
	out := new(PrivateBasicUserInfo)
	err := c.cc.Invoke(ctx, Account_UpdateBasicUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetCustomerReviews(ctx context.Context, in *GetCustomerReviewsRequest, opts ...grpc.CallOption) (*GetCustomerReviewsResponse, error) {
	out := new(GetCustomerReviewsResponse)
	err := c.cc.Invoke(ctx, Account_GetCustomerReviews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) AddCustomerReview(ctx context.Context, in *AddCustomerReviewRequest, opts ...grpc.CallOption) (*CustomerReview, error) {
	out := new(CustomerReview)
	err := c.cc.Invoke(ctx, Account_AddCustomerReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	GetPublicProfileSection(context.Context, *GetPublicProfileSectionRequest) (*ProfileSectionResponse, error)
	GetPrivateProfileSection(context.Context, *GetPrivateProfileSectionRequest) (*ProfileSectionResponse, error)
	UpdateProfileSection(context.Context, *UpdateProfileSectionRequest) (*ProfileSectionResponse, error)
	GetPrivateBasicUserInfo(context.Context, *empty.Empty) (*PrivateBasicUserInfo, error)
	GetPublicBasicUserInfo(context.Context, *GetPublicBasicUserInfoRequest) (*PublicBasicUserInfo, error)
	UpdateBasicUserInfo(context.Context, *UpdatePrivateBasicUserInfoRequest) (*PrivateBasicUserInfo, error)
	GetCustomerReviews(context.Context, *GetCustomerReviewsRequest) (*GetCustomerReviewsResponse, error)
	AddCustomerReview(context.Context, *AddCustomerReviewRequest) (*CustomerReview, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) GetPublicProfileSection(context.Context, *GetPublicProfileSectionRequest) (*ProfileSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicProfileSection not implemented")
}
func (UnimplementedAccountServer) GetPrivateProfileSection(context.Context, *GetPrivateProfileSectionRequest) (*ProfileSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateProfileSection not implemented")
}
func (UnimplementedAccountServer) UpdateProfileSection(context.Context, *UpdateProfileSectionRequest) (*ProfileSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfileSection not implemented")
}
func (UnimplementedAccountServer) GetPrivateBasicUserInfo(context.Context, *empty.Empty) (*PrivateBasicUserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateBasicUserInfo not implemented")
}
func (UnimplementedAccountServer) GetPublicBasicUserInfo(context.Context, *GetPublicBasicUserInfoRequest) (*PublicBasicUserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicBasicUserInfo not implemented")
}
func (UnimplementedAccountServer) UpdateBasicUserInfo(context.Context, *UpdatePrivateBasicUserInfoRequest) (*PrivateBasicUserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBasicUserInfo not implemented")
}
func (UnimplementedAccountServer) GetCustomerReviews(context.Context, *GetCustomerReviewsRequest) (*GetCustomerReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerReviews not implemented")
}
func (UnimplementedAccountServer) AddCustomerReview(context.Context, *AddCustomerReviewRequest) (*CustomerReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomerReview not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_GetPublicProfileSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicProfileSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetPublicProfileSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetPublicProfileSection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetPublicProfileSection(ctx, req.(*GetPublicProfileSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetPrivateProfileSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateProfileSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetPrivateProfileSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetPrivateProfileSection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetPrivateProfileSection(ctx, req.(*GetPrivateProfileSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UpdateProfileSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdateProfileSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_UpdateProfileSection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdateProfileSection(ctx, req.(*UpdateProfileSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetPrivateBasicUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetPrivateBasicUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetPrivateBasicUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetPrivateBasicUserInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetPublicBasicUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicBasicUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetPublicBasicUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetPublicBasicUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetPublicBasicUserInfo(ctx, req.(*GetPublicBasicUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UpdateBasicUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrivateBasicUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdateBasicUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_UpdateBasicUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdateBasicUserInfo(ctx, req.(*UpdatePrivateBasicUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetCustomerReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetCustomerReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetCustomerReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetCustomerReviews(ctx, req.(*GetCustomerReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_AddCustomerReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCustomerReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).AddCustomerReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_AddCustomerReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).AddCustomerReview(ctx, req.(*AddCustomerReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicProfileSection",
			Handler:    _Account_GetPublicProfileSection_Handler,
		},
		{
			MethodName: "GetPrivateProfileSection",
			Handler:    _Account_GetPrivateProfileSection_Handler,
		},
		{
			MethodName: "UpdateProfileSection",
			Handler:    _Account_UpdateProfileSection_Handler,
		},
		{
			MethodName: "GetPrivateBasicUserInfo",
			Handler:    _Account_GetPrivateBasicUserInfo_Handler,
		},
		{
			MethodName: "GetPublicBasicUserInfo",
			Handler:    _Account_GetPublicBasicUserInfo_Handler,
		},
		{
			MethodName: "UpdateBasicUserInfo",
			Handler:    _Account_UpdateBasicUserInfo_Handler,
		},
		{
			MethodName: "GetCustomerReviews",
			Handler:    _Account_GetCustomerReviews_Handler,
		},
		{
			MethodName: "AddCustomerReview",
			Handler:    _Account_AddCustomerReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/account.proto",
}
