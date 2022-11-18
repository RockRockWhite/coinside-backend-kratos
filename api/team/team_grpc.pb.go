// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: team.proto

package api

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

// TeamClient is the client API for Team service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamClient interface {
	// Get a team by id
	GetTeamByID(ctx context.Context, in *GetTeamByIdRequest, opts ...grpc.CallOption) (*GetTeamResponse, error)
	// Get a team by id via stream
	GetTeamByIDStream(ctx context.Context, opts ...grpc.CallOption) (Team_GetTeamByIDStreamClient, error)
	// Get a team by name
	GetTeamsByName(ctx context.Context, in *GetTeamsByNameRequest, opts ...grpc.CallOption) (*GetTeamsResponse, error)
	// Get a user by name via stream
	GetTeamsByNameStream(ctx context.Context, opts ...grpc.CallOption) (Team_GetTeamsByNameStreamClient, error)
	// add a team
	AddTeam(ctx context.Context, in *TeamInfo, opts ...grpc.CallOption) (*AddTeamResponse, error)
	// add a team via stream
	AddTeamStream(ctx context.Context, opts ...grpc.CallOption) (Team_AddTeamStreamClient, error)
	// update a team
	UpdateTeam(ctx context.Context, in *TeamInfo, opts ...grpc.CallOption) (*UpdateTeamResponse, error)
	// update a team via stream
	UpdateTeamSteam(ctx context.Context, in *TeamInfo, opts ...grpc.CallOption) (Team_UpdateTeamSteamClient, error)
	// delete a team
	DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error)
	// delete a team via stream
	DeleteTeamStream(ctx context.Context, opts ...grpc.CallOption) (Team_DeleteTeamStreamClient, error)
	// add  a  Member
	AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*AddMemberResponse, error)
	// add a member  via stream
	AddMemberStream(ctx context.Context, opts ...grpc.CallOption) (Team_AddMemberStreamClient, error)
	// Delete  a  Member
	DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*DeleteMemberResponse, error)
	// Delete a Member  via stream
	DeleteMemberStream(ctx context.Context, opts ...grpc.CallOption) (Team_DeleteMemberStreamClient, error)
	// add an administrator
	AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error)
	// add an administrator  via stream
	AddAdminStream(ctx context.Context, opts ...grpc.CallOption) (Team_AddAdminStreamClient, error)
}

type teamClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamClient(cc grpc.ClientConnInterface) TeamClient {
	return &teamClient{cc}
}

func (c *teamClient) GetTeamByID(ctx context.Context, in *GetTeamByIdRequest, opts ...grpc.CallOption) (*GetTeamResponse, error) {
	out := new(GetTeamResponse)
	err := c.cc.Invoke(ctx, "/team.Team/GetTeamByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) GetTeamByIDStream(ctx context.Context, opts ...grpc.CallOption) (Team_GetTeamByIDStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Team_ServiceDesc.Streams[0], "/team.Team/GetTeamByIDStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamGetTeamByIDStreamClient{stream}
	return x, nil
}

type Team_GetTeamByIDStreamClient interface {
	Send(*GetTeamByIdRequest) error
	Recv() (*GetTeamResponse, error)
	grpc.ClientStream
}

type teamGetTeamByIDStreamClient struct {
	grpc.ClientStream
}

func (x *teamGetTeamByIDStreamClient) Send(m *GetTeamByIdRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teamGetTeamByIDStreamClient) Recv() (*GetTeamResponse, error) {
	m := new(GetTeamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teamClient) GetTeamsByName(ctx context.Context, in *GetTeamsByNameRequest, opts ...grpc.CallOption) (*GetTeamsResponse, error) {
	out := new(GetTeamsResponse)
	err := c.cc.Invoke(ctx, "/team.Team/GetTeamsByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) GetTeamsByNameStream(ctx context.Context, opts ...grpc.CallOption) (Team_GetTeamsByNameStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Team_ServiceDesc.Streams[1], "/team.Team/GetTeamsByNameStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamGetTeamsByNameStreamClient{stream}
	return x, nil
}

type Team_GetTeamsByNameStreamClient interface {
	Send(*GetTeamsByNameRequest) error
	Recv() (*GetTeamsResponse, error)
	grpc.ClientStream
}

type teamGetTeamsByNameStreamClient struct {
	grpc.ClientStream
}

func (x *teamGetTeamsByNameStreamClient) Send(m *GetTeamsByNameRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teamGetTeamsByNameStreamClient) Recv() (*GetTeamsResponse, error) {
	m := new(GetTeamsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teamClient) AddTeam(ctx context.Context, in *TeamInfo, opts ...grpc.CallOption) (*AddTeamResponse, error) {
	out := new(AddTeamResponse)
	err := c.cc.Invoke(ctx, "/team.Team/AddTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) AddTeamStream(ctx context.Context, opts ...grpc.CallOption) (Team_AddTeamStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Team_ServiceDesc.Streams[2], "/team.Team/AddTeamStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamAddTeamStreamClient{stream}
	return x, nil
}

type Team_AddTeamStreamClient interface {
	Send(*TeamInfo) error
	Recv() (*AddTeamResponse, error)
	grpc.ClientStream
}

type teamAddTeamStreamClient struct {
	grpc.ClientStream
}

func (x *teamAddTeamStreamClient) Send(m *TeamInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teamAddTeamStreamClient) Recv() (*AddTeamResponse, error) {
	m := new(AddTeamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teamClient) UpdateTeam(ctx context.Context, in *TeamInfo, opts ...grpc.CallOption) (*UpdateTeamResponse, error) {
	out := new(UpdateTeamResponse)
	err := c.cc.Invoke(ctx, "/team.Team/UpdateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) UpdateTeamSteam(ctx context.Context, in *TeamInfo, opts ...grpc.CallOption) (Team_UpdateTeamSteamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Team_ServiceDesc.Streams[3], "/team.Team/UpdateTeamSteam", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamUpdateTeamSteamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Team_UpdateTeamSteamClient interface {
	Recv() (*UpdateTeamResponse, error)
	grpc.ClientStream
}

type teamUpdateTeamSteamClient struct {
	grpc.ClientStream
}

func (x *teamUpdateTeamSteamClient) Recv() (*UpdateTeamResponse, error) {
	m := new(UpdateTeamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teamClient) DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error) {
	out := new(DeleteTeamResponse)
	err := c.cc.Invoke(ctx, "/team.Team/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) DeleteTeamStream(ctx context.Context, opts ...grpc.CallOption) (Team_DeleteTeamStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Team_ServiceDesc.Streams[4], "/team.Team/DeleteTeamStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamDeleteTeamStreamClient{stream}
	return x, nil
}

type Team_DeleteTeamStreamClient interface {
	Send(*DeleteTeamRequest) error
	Recv() (*DeleteTeamResponse, error)
	grpc.ClientStream
}

type teamDeleteTeamStreamClient struct {
	grpc.ClientStream
}

func (x *teamDeleteTeamStreamClient) Send(m *DeleteTeamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teamDeleteTeamStreamClient) Recv() (*DeleteTeamResponse, error) {
	m := new(DeleteTeamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teamClient) AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*AddMemberResponse, error) {
	out := new(AddMemberResponse)
	err := c.cc.Invoke(ctx, "/team.Team/AddMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) AddMemberStream(ctx context.Context, opts ...grpc.CallOption) (Team_AddMemberStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Team_ServiceDesc.Streams[5], "/team.Team/AddMemberStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamAddMemberStreamClient{stream}
	return x, nil
}

type Team_AddMemberStreamClient interface {
	Send(*AddMemberRequest) error
	Recv() (*AddMemberResponse, error)
	grpc.ClientStream
}

type teamAddMemberStreamClient struct {
	grpc.ClientStream
}

func (x *teamAddMemberStreamClient) Send(m *AddMemberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teamAddMemberStreamClient) Recv() (*AddMemberResponse, error) {
	m := new(AddMemberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teamClient) DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*DeleteMemberResponse, error) {
	out := new(DeleteMemberResponse)
	err := c.cc.Invoke(ctx, "/team.Team/DeleteMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) DeleteMemberStream(ctx context.Context, opts ...grpc.CallOption) (Team_DeleteMemberStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Team_ServiceDesc.Streams[6], "/team.Team/DeleteMemberStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamDeleteMemberStreamClient{stream}
	return x, nil
}

type Team_DeleteMemberStreamClient interface {
	Send(*DeleteMemberRequest) error
	Recv() (*DeleteMemberResponse, error)
	grpc.ClientStream
}

type teamDeleteMemberStreamClient struct {
	grpc.ClientStream
}

func (x *teamDeleteMemberStreamClient) Send(m *DeleteMemberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teamDeleteMemberStreamClient) Recv() (*DeleteMemberResponse, error) {
	m := new(DeleteMemberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teamClient) AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error) {
	out := new(AddAdminResponse)
	err := c.cc.Invoke(ctx, "/team.Team/AddAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) AddAdminStream(ctx context.Context, opts ...grpc.CallOption) (Team_AddAdminStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Team_ServiceDesc.Streams[7], "/team.Team/AddAdminStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamAddAdminStreamClient{stream}
	return x, nil
}

type Team_AddAdminStreamClient interface {
	Send(*AddAdminRequest) error
	Recv() (*AddAdminResponse, error)
	grpc.ClientStream
}

type teamAddAdminStreamClient struct {
	grpc.ClientStream
}

func (x *teamAddAdminStreamClient) Send(m *AddAdminRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teamAddAdminStreamClient) Recv() (*AddAdminResponse, error) {
	m := new(AddAdminResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TeamServer is the server API for Team service.
// All implementations must embed UnimplementedTeamServer
// for forward compatibility
type TeamServer interface {
	// Get a team by id
	GetTeamByID(context.Context, *GetTeamByIdRequest) (*GetTeamResponse, error)
	// Get a team by id via stream
	GetTeamByIDStream(Team_GetTeamByIDStreamServer) error
	// Get a team by name
	GetTeamsByName(context.Context, *GetTeamsByNameRequest) (*GetTeamsResponse, error)
	// Get a user by name via stream
	GetTeamsByNameStream(Team_GetTeamsByNameStreamServer) error
	// add a team
	AddTeam(context.Context, *TeamInfo) (*AddTeamResponse, error)
	// add a team via stream
	AddTeamStream(Team_AddTeamStreamServer) error
	// update a team
	UpdateTeam(context.Context, *TeamInfo) (*UpdateTeamResponse, error)
	// update a team via stream
	UpdateTeamSteam(*TeamInfo, Team_UpdateTeamSteamServer) error
	// delete a team
	DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error)
	// delete a team via stream
	DeleteTeamStream(Team_DeleteTeamStreamServer) error
	// add  a  Member
	AddMember(context.Context, *AddMemberRequest) (*AddMemberResponse, error)
	// add a member  via stream
	AddMemberStream(Team_AddMemberStreamServer) error
	// Delete  a  Member
	DeleteMember(context.Context, *DeleteMemberRequest) (*DeleteMemberResponse, error)
	// Delete a Member  via stream
	DeleteMemberStream(Team_DeleteMemberStreamServer) error
	// add an administrator
	AddAdmin(context.Context, *AddAdminRequest) (*AddAdminResponse, error)
	// add an administrator  via stream
	AddAdminStream(Team_AddAdminStreamServer) error
	mustEmbedUnimplementedTeamServer()
}

// UnimplementedTeamServer must be embedded to have forward compatible implementations.
type UnimplementedTeamServer struct {
}

func (UnimplementedTeamServer) GetTeamByID(context.Context, *GetTeamByIdRequest) (*GetTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamByID not implemented")
}
func (UnimplementedTeamServer) GetTeamByIDStream(Team_GetTeamByIDStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTeamByIDStream not implemented")
}
func (UnimplementedTeamServer) GetTeamsByName(context.Context, *GetTeamsByNameRequest) (*GetTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsByName not implemented")
}
func (UnimplementedTeamServer) GetTeamsByNameStream(Team_GetTeamsByNameStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTeamsByNameStream not implemented")
}
func (UnimplementedTeamServer) AddTeam(context.Context, *TeamInfo) (*AddTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeam not implemented")
}
func (UnimplementedTeamServer) AddTeamStream(Team_AddTeamStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AddTeamStream not implemented")
}
func (UnimplementedTeamServer) UpdateTeam(context.Context, *TeamInfo) (*UpdateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (UnimplementedTeamServer) UpdateTeamSteam(*TeamInfo, Team_UpdateTeamSteamServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateTeamSteam not implemented")
}
func (UnimplementedTeamServer) DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedTeamServer) DeleteTeamStream(Team_DeleteTeamStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteTeamStream not implemented")
}
func (UnimplementedTeamServer) AddMember(context.Context, *AddMemberRequest) (*AddMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMember not implemented")
}
func (UnimplementedTeamServer) AddMemberStream(Team_AddMemberStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AddMemberStream not implemented")
}
func (UnimplementedTeamServer) DeleteMember(context.Context, *DeleteMemberRequest) (*DeleteMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMember not implemented")
}
func (UnimplementedTeamServer) DeleteMemberStream(Team_DeleteMemberStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteMemberStream not implemented")
}
func (UnimplementedTeamServer) AddAdmin(context.Context, *AddAdminRequest) (*AddAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdmin not implemented")
}
func (UnimplementedTeamServer) AddAdminStream(Team_AddAdminStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AddAdminStream not implemented")
}
func (UnimplementedTeamServer) mustEmbedUnimplementedTeamServer() {}

// UnsafeTeamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServer will
// result in compilation errors.
type UnsafeTeamServer interface {
	mustEmbedUnimplementedTeamServer()
}

func RegisterTeamServer(s grpc.ServiceRegistrar, srv TeamServer) {
	s.RegisterService(&Team_ServiceDesc, srv)
}

func _Team_GetTeamByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).GetTeamByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/GetTeamByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).GetTeamByID(ctx, req.(*GetTeamByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_GetTeamByIDStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeamServer).GetTeamByIDStream(&teamGetTeamByIDStreamServer{stream})
}

type Team_GetTeamByIDStreamServer interface {
	Send(*GetTeamResponse) error
	Recv() (*GetTeamByIdRequest, error)
	grpc.ServerStream
}

type teamGetTeamByIDStreamServer struct {
	grpc.ServerStream
}

func (x *teamGetTeamByIDStreamServer) Send(m *GetTeamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teamGetTeamByIDStreamServer) Recv() (*GetTeamByIdRequest, error) {
	m := new(GetTeamByIdRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Team_GetTeamsByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).GetTeamsByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/GetTeamsByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).GetTeamsByName(ctx, req.(*GetTeamsByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_GetTeamsByNameStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeamServer).GetTeamsByNameStream(&teamGetTeamsByNameStreamServer{stream})
}

type Team_GetTeamsByNameStreamServer interface {
	Send(*GetTeamsResponse) error
	Recv() (*GetTeamsByNameRequest, error)
	grpc.ServerStream
}

type teamGetTeamsByNameStreamServer struct {
	grpc.ServerStream
}

func (x *teamGetTeamsByNameStreamServer) Send(m *GetTeamsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teamGetTeamsByNameStreamServer) Recv() (*GetTeamsByNameRequest, error) {
	m := new(GetTeamsByNameRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Team_AddTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).AddTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/AddTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).AddTeam(ctx, req.(*TeamInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_AddTeamStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeamServer).AddTeamStream(&teamAddTeamStreamServer{stream})
}

type Team_AddTeamStreamServer interface {
	Send(*AddTeamResponse) error
	Recv() (*TeamInfo, error)
	grpc.ServerStream
}

type teamAddTeamStreamServer struct {
	grpc.ServerStream
}

func (x *teamAddTeamStreamServer) Send(m *AddTeamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teamAddTeamStreamServer) Recv() (*TeamInfo, error) {
	m := new(TeamInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Team_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).UpdateTeam(ctx, req.(*TeamInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_UpdateTeamSteam_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TeamInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TeamServer).UpdateTeamSteam(m, &teamUpdateTeamSteamServer{stream})
}

type Team_UpdateTeamSteamServer interface {
	Send(*UpdateTeamResponse) error
	grpc.ServerStream
}

type teamUpdateTeamSteamServer struct {
	grpc.ServerStream
}

func (x *teamUpdateTeamSteamServer) Send(m *UpdateTeamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Team_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).DeleteTeam(ctx, req.(*DeleteTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_DeleteTeamStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeamServer).DeleteTeamStream(&teamDeleteTeamStreamServer{stream})
}

type Team_DeleteTeamStreamServer interface {
	Send(*DeleteTeamResponse) error
	Recv() (*DeleteTeamRequest, error)
	grpc.ServerStream
}

type teamDeleteTeamStreamServer struct {
	grpc.ServerStream
}

func (x *teamDeleteTeamStreamServer) Send(m *DeleteTeamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teamDeleteTeamStreamServer) Recv() (*DeleteTeamRequest, error) {
	m := new(DeleteTeamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Team_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/AddMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).AddMember(ctx, req.(*AddMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_AddMemberStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeamServer).AddMemberStream(&teamAddMemberStreamServer{stream})
}

type Team_AddMemberStreamServer interface {
	Send(*AddMemberResponse) error
	Recv() (*AddMemberRequest, error)
	grpc.ServerStream
}

type teamAddMemberStreamServer struct {
	grpc.ServerStream
}

func (x *teamAddMemberStreamServer) Send(m *AddMemberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teamAddMemberStreamServer) Recv() (*AddMemberRequest, error) {
	m := new(AddMemberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Team_DeleteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).DeleteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/DeleteMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).DeleteMember(ctx, req.(*DeleteMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_DeleteMemberStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeamServer).DeleteMemberStream(&teamDeleteMemberStreamServer{stream})
}

type Team_DeleteMemberStreamServer interface {
	Send(*DeleteMemberResponse) error
	Recv() (*DeleteMemberRequest, error)
	grpc.ServerStream
}

type teamDeleteMemberStreamServer struct {
	grpc.ServerStream
}

func (x *teamDeleteMemberStreamServer) Send(m *DeleteMemberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teamDeleteMemberStreamServer) Recv() (*DeleteMemberRequest, error) {
	m := new(DeleteMemberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Team_AddAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).AddAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/AddAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).AddAdmin(ctx, req.(*AddAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_AddAdminStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeamServer).AddAdminStream(&teamAddAdminStreamServer{stream})
}

type Team_AddAdminStreamServer interface {
	Send(*AddAdminResponse) error
	Recv() (*AddAdminRequest, error)
	grpc.ServerStream
}

type teamAddAdminStreamServer struct {
	grpc.ServerStream
}

func (x *teamAddAdminStreamServer) Send(m *AddAdminResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teamAddAdminStreamServer) Recv() (*AddAdminRequest, error) {
	m := new(AddAdminRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Team_ServiceDesc is the grpc.ServiceDesc for Team service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Team_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "team.Team",
	HandlerType: (*TeamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeamByID",
			Handler:    _Team_GetTeamByID_Handler,
		},
		{
			MethodName: "GetTeamsByName",
			Handler:    _Team_GetTeamsByName_Handler,
		},
		{
			MethodName: "AddTeam",
			Handler:    _Team_AddTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _Team_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _Team_DeleteTeam_Handler,
		},
		{
			MethodName: "AddMember",
			Handler:    _Team_AddMember_Handler,
		},
		{
			MethodName: "DeleteMember",
			Handler:    _Team_DeleteMember_Handler,
		},
		{
			MethodName: "AddAdmin",
			Handler:    _Team_AddAdmin_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTeamByIDStream",
			Handler:       _Team_GetTeamByIDStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetTeamsByNameStream",
			Handler:       _Team_GetTeamsByNameStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AddTeamStream",
			Handler:       _Team_AddTeamStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateTeamSteam",
			Handler:       _Team_UpdateTeamSteam_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DeleteTeamStream",
			Handler:       _Team_DeleteTeamStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AddMemberStream",
			Handler:       _Team_AddMemberStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeleteMemberStream",
			Handler:       _Team_DeleteMemberStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AddAdminStream",
			Handler:       _Team_AddAdminStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "team.proto",
}