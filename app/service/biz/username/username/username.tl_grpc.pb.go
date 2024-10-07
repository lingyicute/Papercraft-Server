//
// WARNING! All changes made in this file will be lost!
// Created from 'scheme.tl' by 'mtprotoc'
//
// Copyright (c) 2024-present,  Yomi.
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: username.tl.proto

package username

import (
	context "context"
	mtproto "github.com/teamgram/proto/mtproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RPCUsername_UsernameGetAccountUsername_FullMethodName    = "/username.RPCUsername/username_getAccountUsername"
	RPCUsername_UsernameCheckAccountUsername_FullMethodName  = "/username.RPCUsername/username_checkAccountUsername"
	RPCUsername_UsernameGetChannelUsername_FullMethodName    = "/username.RPCUsername/username_getChannelUsername"
	RPCUsername_UsernameCheckChannelUsername_FullMethodName  = "/username.RPCUsername/username_checkChannelUsername"
	RPCUsername_UsernameUpdateUsernameByPeer_FullMethodName  = "/username.RPCUsername/username_updateUsernameByPeer"
	RPCUsername_UsernameCheckUsername_FullMethodName         = "/username.RPCUsername/username_checkUsername"
	RPCUsername_UsernameUpdateUsername_FullMethodName        = "/username.RPCUsername/username_updateUsername"
	RPCUsername_UsernameDeleteUsername_FullMethodName        = "/username.RPCUsername/username_deleteUsername"
	RPCUsername_UsernameResolveUsername_FullMethodName       = "/username.RPCUsername/username_resolveUsername"
	RPCUsername_UsernameGetListByUsernameList_FullMethodName = "/username.RPCUsername/username_getListByUsernameList"
	RPCUsername_UsernameDeleteUsernameByPeer_FullMethodName  = "/username.RPCUsername/username_deleteUsernameByPeer"
	RPCUsername_UsernameSearch_FullMethodName                = "/username.RPCUsername/username_search"
)

// RPCUsernameClient is the client API for RPCUsername service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCUsernameClient interface {
	UsernameGetAccountUsername(ctx context.Context, in *TLUsernameGetAccountUsername, opts ...grpc.CallOption) (*UsernameData, error)
	UsernameCheckAccountUsername(ctx context.Context, in *TLUsernameCheckAccountUsername, opts ...grpc.CallOption) (*UsernameExist, error)
	UsernameGetChannelUsername(ctx context.Context, in *TLUsernameGetChannelUsername, opts ...grpc.CallOption) (*UsernameData, error)
	UsernameCheckChannelUsername(ctx context.Context, in *TLUsernameCheckChannelUsername, opts ...grpc.CallOption) (*UsernameExist, error)
	UsernameUpdateUsernameByPeer(ctx context.Context, in *TLUsernameUpdateUsernameByPeer, opts ...grpc.CallOption) (*mtproto.Bool, error)
	UsernameCheckUsername(ctx context.Context, in *TLUsernameCheckUsername, opts ...grpc.CallOption) (*UsernameExist, error)
	UsernameUpdateUsername(ctx context.Context, in *TLUsernameUpdateUsername, opts ...grpc.CallOption) (*mtproto.Bool, error)
	UsernameDeleteUsername(ctx context.Context, in *TLUsernameDeleteUsername, opts ...grpc.CallOption) (*mtproto.Bool, error)
	UsernameResolveUsername(ctx context.Context, in *TLUsernameResolveUsername, opts ...grpc.CallOption) (*mtproto.Peer, error)
	UsernameGetListByUsernameList(ctx context.Context, in *TLUsernameGetListByUsernameList, opts ...grpc.CallOption) (*Vector_UsernameData, error)
	UsernameDeleteUsernameByPeer(ctx context.Context, in *TLUsernameDeleteUsernameByPeer, opts ...grpc.CallOption) (*mtproto.Bool, error)
	UsernameSearch(ctx context.Context, in *TLUsernameSearch, opts ...grpc.CallOption) (*Vector_UsernameData, error)
}

type rPCUsernameClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCUsernameClient(cc grpc.ClientConnInterface) RPCUsernameClient {
	return &rPCUsernameClient{cc}
}

func (c *rPCUsernameClient) UsernameGetAccountUsername(ctx context.Context, in *TLUsernameGetAccountUsername, opts ...grpc.CallOption) (*UsernameData, error) {
	out := new(UsernameData)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameGetAccountUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameCheckAccountUsername(ctx context.Context, in *TLUsernameCheckAccountUsername, opts ...grpc.CallOption) (*UsernameExist, error) {
	out := new(UsernameExist)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameCheckAccountUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameGetChannelUsername(ctx context.Context, in *TLUsernameGetChannelUsername, opts ...grpc.CallOption) (*UsernameData, error) {
	out := new(UsernameData)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameGetChannelUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameCheckChannelUsername(ctx context.Context, in *TLUsernameCheckChannelUsername, opts ...grpc.CallOption) (*UsernameExist, error) {
	out := new(UsernameExist)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameCheckChannelUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameUpdateUsernameByPeer(ctx context.Context, in *TLUsernameUpdateUsernameByPeer, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameUpdateUsernameByPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameCheckUsername(ctx context.Context, in *TLUsernameCheckUsername, opts ...grpc.CallOption) (*UsernameExist, error) {
	out := new(UsernameExist)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameCheckUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameUpdateUsername(ctx context.Context, in *TLUsernameUpdateUsername, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameUpdateUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameDeleteUsername(ctx context.Context, in *TLUsernameDeleteUsername, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameDeleteUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameResolveUsername(ctx context.Context, in *TLUsernameResolveUsername, opts ...grpc.CallOption) (*mtproto.Peer, error) {
	out := new(mtproto.Peer)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameResolveUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameGetListByUsernameList(ctx context.Context, in *TLUsernameGetListByUsernameList, opts ...grpc.CallOption) (*Vector_UsernameData, error) {
	out := new(Vector_UsernameData)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameGetListByUsernameList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameDeleteUsernameByPeer(ctx context.Context, in *TLUsernameDeleteUsernameByPeer, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameDeleteUsernameByPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUsernameClient) UsernameSearch(ctx context.Context, in *TLUsernameSearch, opts ...grpc.CallOption) (*Vector_UsernameData, error) {
	out := new(Vector_UsernameData)
	err := c.cc.Invoke(ctx, RPCUsername_UsernameSearch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCUsernameServer is the server API for RPCUsername service.
// All implementations should embed UnimplementedRPCUsernameServer
// for forward compatibility
type RPCUsernameServer interface {
	UsernameGetAccountUsername(context.Context, *TLUsernameGetAccountUsername) (*UsernameData, error)
	UsernameCheckAccountUsername(context.Context, *TLUsernameCheckAccountUsername) (*UsernameExist, error)
	UsernameGetChannelUsername(context.Context, *TLUsernameGetChannelUsername) (*UsernameData, error)
	UsernameCheckChannelUsername(context.Context, *TLUsernameCheckChannelUsername) (*UsernameExist, error)
	UsernameUpdateUsernameByPeer(context.Context, *TLUsernameUpdateUsernameByPeer) (*mtproto.Bool, error)
	UsernameCheckUsername(context.Context, *TLUsernameCheckUsername) (*UsernameExist, error)
	UsernameUpdateUsername(context.Context, *TLUsernameUpdateUsername) (*mtproto.Bool, error)
	UsernameDeleteUsername(context.Context, *TLUsernameDeleteUsername) (*mtproto.Bool, error)
	UsernameResolveUsername(context.Context, *TLUsernameResolveUsername) (*mtproto.Peer, error)
	UsernameGetListByUsernameList(context.Context, *TLUsernameGetListByUsernameList) (*Vector_UsernameData, error)
	UsernameDeleteUsernameByPeer(context.Context, *TLUsernameDeleteUsernameByPeer) (*mtproto.Bool, error)
	UsernameSearch(context.Context, *TLUsernameSearch) (*Vector_UsernameData, error)
}

// UnimplementedRPCUsernameServer should be embedded to have forward compatible implementations.
type UnimplementedRPCUsernameServer struct {
}

func (UnimplementedRPCUsernameServer) UsernameGetAccountUsername(context.Context, *TLUsernameGetAccountUsername) (*UsernameData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameGetAccountUsername not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameCheckAccountUsername(context.Context, *TLUsernameCheckAccountUsername) (*UsernameExist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameCheckAccountUsername not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameGetChannelUsername(context.Context, *TLUsernameGetChannelUsername) (*UsernameData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameGetChannelUsername not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameCheckChannelUsername(context.Context, *TLUsernameCheckChannelUsername) (*UsernameExist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameCheckChannelUsername not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameUpdateUsernameByPeer(context.Context, *TLUsernameUpdateUsernameByPeer) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameUpdateUsernameByPeer not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameCheckUsername(context.Context, *TLUsernameCheckUsername) (*UsernameExist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameCheckUsername not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameUpdateUsername(context.Context, *TLUsernameUpdateUsername) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameUpdateUsername not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameDeleteUsername(context.Context, *TLUsernameDeleteUsername) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameDeleteUsername not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameResolveUsername(context.Context, *TLUsernameResolveUsername) (*mtproto.Peer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameResolveUsername not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameGetListByUsernameList(context.Context, *TLUsernameGetListByUsernameList) (*Vector_UsernameData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameGetListByUsernameList not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameDeleteUsernameByPeer(context.Context, *TLUsernameDeleteUsernameByPeer) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameDeleteUsernameByPeer not implemented")
}
func (UnimplementedRPCUsernameServer) UsernameSearch(context.Context, *TLUsernameSearch) (*Vector_UsernameData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameSearch not implemented")
}

// UnsafeRPCUsernameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCUsernameServer will
// result in compilation errors.
type UnsafeRPCUsernameServer interface {
	mustEmbedUnimplementedRPCUsernameServer()
}

func RegisterRPCUsernameServer(s grpc.ServiceRegistrar, srv RPCUsernameServer) {
	s.RegisterService(&RPCUsername_ServiceDesc, srv)
}

func _RPCUsername_UsernameGetAccountUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameGetAccountUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameGetAccountUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameGetAccountUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameGetAccountUsername(ctx, req.(*TLUsernameGetAccountUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameCheckAccountUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameCheckAccountUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameCheckAccountUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameCheckAccountUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameCheckAccountUsername(ctx, req.(*TLUsernameCheckAccountUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameGetChannelUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameGetChannelUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameGetChannelUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameGetChannelUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameGetChannelUsername(ctx, req.(*TLUsernameGetChannelUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameCheckChannelUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameCheckChannelUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameCheckChannelUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameCheckChannelUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameCheckChannelUsername(ctx, req.(*TLUsernameCheckChannelUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameUpdateUsernameByPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameUpdateUsernameByPeer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameUpdateUsernameByPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameUpdateUsernameByPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameUpdateUsernameByPeer(ctx, req.(*TLUsernameUpdateUsernameByPeer))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameCheckUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameCheckUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameCheckUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameCheckUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameCheckUsername(ctx, req.(*TLUsernameCheckUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameUpdateUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameUpdateUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameUpdateUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameUpdateUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameUpdateUsername(ctx, req.(*TLUsernameUpdateUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameDeleteUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameDeleteUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameDeleteUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameDeleteUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameDeleteUsername(ctx, req.(*TLUsernameDeleteUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameResolveUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameResolveUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameResolveUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameResolveUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameResolveUsername(ctx, req.(*TLUsernameResolveUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameGetListByUsernameList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameGetListByUsernameList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameGetListByUsernameList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameGetListByUsernameList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameGetListByUsernameList(ctx, req.(*TLUsernameGetListByUsernameList))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameDeleteUsernameByPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameDeleteUsernameByPeer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameDeleteUsernameByPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameDeleteUsernameByPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameDeleteUsernameByPeer(ctx, req.(*TLUsernameDeleteUsernameByPeer))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUsername_UsernameSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUsernameSearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUsernameServer).UsernameSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUsername_UsernameSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUsernameServer).UsernameSearch(ctx, req.(*TLUsernameSearch))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCUsername_ServiceDesc is the grpc.ServiceDesc for RPCUsername service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCUsername_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "username.RPCUsername",
	HandlerType: (*RPCUsernameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "username_getAccountUsername",
			Handler:    _RPCUsername_UsernameGetAccountUsername_Handler,
		},
		{
			MethodName: "username_checkAccountUsername",
			Handler:    _RPCUsername_UsernameCheckAccountUsername_Handler,
		},
		{
			MethodName: "username_getChannelUsername",
			Handler:    _RPCUsername_UsernameGetChannelUsername_Handler,
		},
		{
			MethodName: "username_checkChannelUsername",
			Handler:    _RPCUsername_UsernameCheckChannelUsername_Handler,
		},
		{
			MethodName: "username_updateUsernameByPeer",
			Handler:    _RPCUsername_UsernameUpdateUsernameByPeer_Handler,
		},
		{
			MethodName: "username_checkUsername",
			Handler:    _RPCUsername_UsernameCheckUsername_Handler,
		},
		{
			MethodName: "username_updateUsername",
			Handler:    _RPCUsername_UsernameUpdateUsername_Handler,
		},
		{
			MethodName: "username_deleteUsername",
			Handler:    _RPCUsername_UsernameDeleteUsername_Handler,
		},
		{
			MethodName: "username_resolveUsername",
			Handler:    _RPCUsername_UsernameResolveUsername_Handler,
		},
		{
			MethodName: "username_getListByUsernameList",
			Handler:    _RPCUsername_UsernameGetListByUsernameList_Handler,
		},
		{
			MethodName: "username_deleteUsernameByPeer",
			Handler:    _RPCUsername_UsernameDeleteUsernameByPeer_Handler,
		},
		{
			MethodName: "username_search",
			Handler:    _RPCUsername_UsernameSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "username.tl.proto",
}
