// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: task/task.proto

package task

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DefinitionClient is the client API for Definition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DefinitionClient interface {
	Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Task, error)
	List(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Tasks, error)
	Create(ctx context.Context, in *Task, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	Update(ctx context.Context, in *Task, opts ...grpc.CallOption) (*UpdateTaskReponse, error)
	Delete(ctx context.Context, in *Task, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
}

type definitionClient struct {
	cc grpc.ClientConnInterface
}

func NewDefinitionClient(cc grpc.ClientConnInterface) DefinitionClient {
	return &definitionClient{cc}
}

func (c *definitionClient) Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/task.Definition/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionClient) List(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Tasks, error) {
	out := new(Tasks)
	err := c.cc.Invoke(ctx, "/task.Definition/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionClient) Create(ctx context.Context, in *Task, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/task.Definition/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionClient) Update(ctx context.Context, in *Task, opts ...grpc.CallOption) (*UpdateTaskReponse, error) {
	out := new(UpdateTaskReponse)
	err := c.cc.Invoke(ctx, "/task.Definition/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionClient) Delete(ctx context.Context, in *Task, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, "/task.Definition/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DefinitionServer is the server API for Definition service.
// All implementations must embed UnimplementedDefinitionServer
// for forward compatibility
type DefinitionServer interface {
	Get(context.Context, *Filter) (*Task, error)
	List(context.Context, *Filter) (*Tasks, error)
	Create(context.Context, *Task) (*CreateTaskResponse, error)
	Update(context.Context, *Task) (*UpdateTaskReponse, error)
	Delete(context.Context, *Task) (*DeleteTaskResponse, error)
	mustEmbedUnimplementedDefinitionServer()
}

// UnimplementedDefinitionServer must be embedded to have forward compatible implementations.
type UnimplementedDefinitionServer struct {
}

func (UnimplementedDefinitionServer) Get(context.Context, *Filter) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDefinitionServer) List(context.Context, *Filter) (*Tasks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDefinitionServer) Create(context.Context, *Task) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDefinitionServer) Update(context.Context, *Task) (*UpdateTaskReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDefinitionServer) Delete(context.Context, *Task) (*DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDefinitionServer) mustEmbedUnimplementedDefinitionServer() {}

// UnsafeDefinitionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DefinitionServer will
// result in compilation errors.
type UnsafeDefinitionServer interface {
	mustEmbedUnimplementedDefinitionServer()
}

func RegisterDefinitionServer(s grpc.ServiceRegistrar, srv DefinitionServer) {
	s.RegisterService(&Definition_ServiceDesc, srv)
}

func _Definition_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefinitionServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Definition/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefinitionServer).Get(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Definition_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefinitionServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Definition/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefinitionServer).List(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Definition_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefinitionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Definition/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefinitionServer).Create(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Definition_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefinitionServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Definition/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefinitionServer).Update(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Definition_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefinitionServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Definition/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefinitionServer).Delete(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

// Definition_ServiceDesc is the grpc.ServiceDesc for Definition service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Definition_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.Definition",
	HandlerType: (*DefinitionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Definition_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Definition_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Definition_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Definition_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Definition_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task/task.proto",
}

// ExecClient is the client API for Exec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecClient interface {
	Start(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Execution, error)
	Stop(ctx context.Context, in *Execution, opts ...grpc.CallOption) (*StopExecutionResponse, error)
	Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Execution, error)
	List(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Executions, error)
}

type execClient struct {
	cc grpc.ClientConnInterface
}

func NewExecClient(cc grpc.ClientConnInterface) ExecClient {
	return &execClient{cc}
}

func (c *execClient) Start(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Execution, error) {
	out := new(Execution)
	err := c.cc.Invoke(ctx, "/task.Exec/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execClient) Stop(ctx context.Context, in *Execution, opts ...grpc.CallOption) (*StopExecutionResponse, error) {
	out := new(StopExecutionResponse)
	err := c.cc.Invoke(ctx, "/task.Exec/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execClient) Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Execution, error) {
	out := new(Execution)
	err := c.cc.Invoke(ctx, "/task.Exec/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execClient) List(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Executions, error) {
	out := new(Executions)
	err := c.cc.Invoke(ctx, "/task.Exec/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecServer is the server API for Exec service.
// All implementations must embed UnimplementedExecServer
// for forward compatibility
type ExecServer interface {
	Start(context.Context, *Task) (*Execution, error)
	Stop(context.Context, *Execution) (*StopExecutionResponse, error)
	Get(context.Context, *Filter) (*Execution, error)
	List(context.Context, *Filter) (*Executions, error)
	mustEmbedUnimplementedExecServer()
}

// UnimplementedExecServer must be embedded to have forward compatible implementations.
type UnimplementedExecServer struct {
}

func (UnimplementedExecServer) Start(context.Context, *Task) (*Execution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedExecServer) Stop(context.Context, *Execution) (*StopExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedExecServer) Get(context.Context, *Filter) (*Execution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedExecServer) List(context.Context, *Filter) (*Executions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedExecServer) mustEmbedUnimplementedExecServer() {}

// UnsafeExecServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecServer will
// result in compilation errors.
type UnsafeExecServer interface {
	mustEmbedUnimplementedExecServer()
}

func RegisterExecServer(s grpc.ServiceRegistrar, srv ExecServer) {
	s.RegisterService(&Exec_ServiceDesc, srv)
}

func _Exec_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Exec/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecServer).Start(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exec_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Execution)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Exec/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecServer).Stop(ctx, req.(*Execution))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exec_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Exec/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecServer).Get(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exec_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Exec/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecServer).List(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// Exec_ServiceDesc is the grpc.ServiceDesc for Exec service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exec_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.Exec",
	HandlerType: (*ExecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Exec_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Exec_Stop_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Exec_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Exec_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task/task.proto",
}

// ScheduleClient is the client API for Schedule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleClient interface {
	Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Calendar2, error)
	List(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Calendar2, error)
	Create(ctx context.Context, in *Calendar2, opts ...grpc.CallOption) (*CreateCalendarResponse, error)
	Delete(ctx context.Context, in *Calendar2, opts ...grpc.CallOption) (*DeleteCalendarResponse, error)
	Update(ctx context.Context, in *Calendar2, opts ...grpc.CallOption) (*UpdateCalendarResponse, error)
}

type scheduleClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleClient(cc grpc.ClientConnInterface) ScheduleClient {
	return &scheduleClient{cc}
}

func (c *scheduleClient) Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Calendar2, error) {
	out := new(Calendar2)
	err := c.cc.Invoke(ctx, "/task.Schedule/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) List(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Calendar2, error) {
	out := new(Calendar2)
	err := c.cc.Invoke(ctx, "/task.Schedule/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) Create(ctx context.Context, in *Calendar2, opts ...grpc.CallOption) (*CreateCalendarResponse, error) {
	out := new(CreateCalendarResponse)
	err := c.cc.Invoke(ctx, "/task.Schedule/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) Delete(ctx context.Context, in *Calendar2, opts ...grpc.CallOption) (*DeleteCalendarResponse, error) {
	out := new(DeleteCalendarResponse)
	err := c.cc.Invoke(ctx, "/task.Schedule/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) Update(ctx context.Context, in *Calendar2, opts ...grpc.CallOption) (*UpdateCalendarResponse, error) {
	out := new(UpdateCalendarResponse)
	err := c.cc.Invoke(ctx, "/task.Schedule/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServer is the server API for Schedule service.
// All implementations must embed UnimplementedScheduleServer
// for forward compatibility
type ScheduleServer interface {
	Get(context.Context, *Filter) (*Calendar2, error)
	List(context.Context, *Filter) (*Calendar2, error)
	Create(context.Context, *Calendar2) (*CreateCalendarResponse, error)
	Delete(context.Context, *Calendar2) (*DeleteCalendarResponse, error)
	Update(context.Context, *Calendar2) (*UpdateCalendarResponse, error)
	mustEmbedUnimplementedScheduleServer()
}

// UnimplementedScheduleServer must be embedded to have forward compatible implementations.
type UnimplementedScheduleServer struct {
}

func (UnimplementedScheduleServer) Get(context.Context, *Filter) (*Calendar2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedScheduleServer) List(context.Context, *Filter) (*Calendar2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedScheduleServer) Create(context.Context, *Calendar2) (*CreateCalendarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedScheduleServer) Delete(context.Context, *Calendar2) (*DeleteCalendarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedScheduleServer) Update(context.Context, *Calendar2) (*UpdateCalendarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedScheduleServer) mustEmbedUnimplementedScheduleServer() {}

// UnsafeScheduleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServer will
// result in compilation errors.
type UnsafeScheduleServer interface {
	mustEmbedUnimplementedScheduleServer()
}

func RegisterScheduleServer(s grpc.ServiceRegistrar, srv ScheduleServer) {
	s.RegisterService(&Schedule_ServiceDesc, srv)
}

func _Schedule_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Schedule/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Get(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Schedule/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).List(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Calendar2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Schedule/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Create(ctx, req.(*Calendar2))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Calendar2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Schedule/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Delete(ctx, req.(*Calendar2))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Calendar2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Schedule/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Update(ctx, req.(*Calendar2))
	}
	return interceptor(ctx, in, info, handler)
}

// Schedule_ServiceDesc is the grpc.ServiceDesc for Schedule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Schedule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.Schedule",
	HandlerType: (*ScheduleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Schedule_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Schedule_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Schedule_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Schedule_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Schedule_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task/task.proto",
}