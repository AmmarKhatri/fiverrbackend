// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: schedule/schedule.proto

package schedule

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
	Scheduler_ListProviderDays_FullMethodName  = "/schedule.Scheduler/ListProviderDays"
	Scheduler_ListCustomerDays_FullMethodName  = "/schedule.Scheduler/ListCustomerDays"
	Scheduler_UpdateCalendar_FullMethodName    = "/schedule.Scheduler/UpdateCalendar"
	Scheduler_AddAppointment_FullMethodName    = "/schedule.Scheduler/AddAppointment"
	Scheduler_CancelAppointment_FullMethodName = "/schedule.Scheduler/CancelAppointment"
	Scheduler_GetDayDetails_FullMethodName     = "/schedule.Scheduler/GetDayDetails"
)

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerClient interface {
	ListProviderDays(ctx context.Context, in *ListDaysRequest, opts ...grpc.CallOption) (*ListProviderDaysResponse, error)
	ListCustomerDays(ctx context.Context, in *ListDaysRequest, opts ...grpc.CallOption) (*ListCustomerDaysResponse, error)
	UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*UpdateCalendarResponse, error)
	AddAppointment(ctx context.Context, in *AddAppointmentRequest, opts ...grpc.CallOption) (*Appointment, error)
	CancelAppointment(ctx context.Context, in *CancelAppointmentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetDayDetails(ctx context.Context, in *DayInput, opts ...grpc.CallOption) (*DayDetails, error)
}

type schedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) ListProviderDays(ctx context.Context, in *ListDaysRequest, opts ...grpc.CallOption) (*ListProviderDaysResponse, error) {
	out := new(ListProviderDaysResponse)
	err := c.cc.Invoke(ctx, Scheduler_ListProviderDays_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) ListCustomerDays(ctx context.Context, in *ListDaysRequest, opts ...grpc.CallOption) (*ListCustomerDaysResponse, error) {
	out := new(ListCustomerDaysResponse)
	err := c.cc.Invoke(ctx, Scheduler_ListCustomerDays_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*UpdateCalendarResponse, error) {
	out := new(UpdateCalendarResponse)
	err := c.cc.Invoke(ctx, Scheduler_UpdateCalendar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AddAppointment(ctx context.Context, in *AddAppointmentRequest, opts ...grpc.CallOption) (*Appointment, error) {
	out := new(Appointment)
	err := c.cc.Invoke(ctx, Scheduler_AddAppointment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) CancelAppointment(ctx context.Context, in *CancelAppointmentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Scheduler_CancelAppointment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetDayDetails(ctx context.Context, in *DayInput, opts ...grpc.CallOption) (*DayDetails, error) {
	out := new(DayDetails)
	err := c.cc.Invoke(ctx, Scheduler_GetDayDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
// All implementations must embed UnimplementedSchedulerServer
// for forward compatibility
type SchedulerServer interface {
	ListProviderDays(context.Context, *ListDaysRequest) (*ListProviderDaysResponse, error)
	ListCustomerDays(context.Context, *ListDaysRequest) (*ListCustomerDaysResponse, error)
	UpdateCalendar(context.Context, *UpdateCalendarRequest) (*UpdateCalendarResponse, error)
	AddAppointment(context.Context, *AddAppointmentRequest) (*Appointment, error)
	CancelAppointment(context.Context, *CancelAppointmentRequest) (*empty.Empty, error)
	GetDayDetails(context.Context, *DayInput) (*DayDetails, error)
	mustEmbedUnimplementedSchedulerServer()
}

// UnimplementedSchedulerServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (UnimplementedSchedulerServer) ListProviderDays(context.Context, *ListDaysRequest) (*ListProviderDaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProviderDays not implemented")
}
func (UnimplementedSchedulerServer) ListCustomerDays(context.Context, *ListDaysRequest) (*ListCustomerDaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCustomerDays not implemented")
}
func (UnimplementedSchedulerServer) UpdateCalendar(context.Context, *UpdateCalendarRequest) (*UpdateCalendarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCalendar not implemented")
}
func (UnimplementedSchedulerServer) AddAppointment(context.Context, *AddAppointmentRequest) (*Appointment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAppointment not implemented")
}
func (UnimplementedSchedulerServer) CancelAppointment(context.Context, *CancelAppointmentRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAppointment not implemented")
}
func (UnimplementedSchedulerServer) GetDayDetails(context.Context, *DayInput) (*DayDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDayDetails not implemented")
}
func (UnimplementedSchedulerServer) mustEmbedUnimplementedSchedulerServer() {}

// UnsafeSchedulerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerServer will
// result in compilation errors.
type UnsafeSchedulerServer interface {
	mustEmbedUnimplementedSchedulerServer()
}

func RegisterSchedulerServer(s grpc.ServiceRegistrar, srv SchedulerServer) {
	s.RegisterService(&Scheduler_ServiceDesc, srv)
}

func _Scheduler_ListProviderDays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).ListProviderDays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_ListProviderDays_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).ListProviderDays(ctx, req.(*ListDaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_ListCustomerDays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).ListCustomerDays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_ListCustomerDays_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).ListCustomerDays(ctx, req.(*ListDaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_UpdateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UpdateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_UpdateCalendar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UpdateCalendar(ctx, req.(*UpdateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AddAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AddAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_AddAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AddAppointment(ctx, req.(*AddAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_CancelAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).CancelAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_CancelAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).CancelAppointment(ctx, req.(*CancelAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetDayDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DayInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetDayDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_GetDayDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetDayDetails(ctx, req.(*DayInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Scheduler_ServiceDesc is the grpc.ServiceDesc for Scheduler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scheduler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProviderDays",
			Handler:    _Scheduler_ListProviderDays_Handler,
		},
		{
			MethodName: "ListCustomerDays",
			Handler:    _Scheduler_ListCustomerDays_Handler,
		},
		{
			MethodName: "UpdateCalendar",
			Handler:    _Scheduler_UpdateCalendar_Handler,
		},
		{
			MethodName: "AddAppointment",
			Handler:    _Scheduler_AddAppointment_Handler,
		},
		{
			MethodName: "CancelAppointment",
			Handler:    _Scheduler_CancelAppointment_Handler,
		},
		{
			MethodName: "GetDayDetails",
			Handler:    _Scheduler_GetDayDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedule/schedule.proto",
}
