// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/end_device_services.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import context "context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NsDeviceRegistry service

type NsDeviceRegistryClient interface {
	// ListDevices returns the devices that match the given identifiers.
	List(ctx context.Context, in *ListEndDevicesRequest, opts ...grpc.CallOption) (*EndDevices, error)
	// GetDevice returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// SetDevice creates or updates the device.
	Set(ctx context.Context, in *SetDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// DeleteDevice deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type nsDeviceRegistryClient struct {
	cc *grpc.ClientConn
}

func NewNsDeviceRegistryClient(cc *grpc.ClientConn) NsDeviceRegistryClient {
	return &nsDeviceRegistryClient{cc}
}

func (c *nsDeviceRegistryClient) List(ctx context.Context, in *ListEndDevicesRequest, opts ...grpc.CallOption) (*EndDevices, error) {
	out := new(EndDevices)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsDeviceRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsDeviceRegistryClient) Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsDeviceRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsDeviceRegistryClient) Set(ctx context.Context, in *SetDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsDeviceRegistry/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsDeviceRegistryClient) Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsDeviceRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NsDeviceRegistry service

type NsDeviceRegistryServer interface {
	// ListDevices returns the devices that match the given identifiers.
	List(context.Context, *ListEndDevicesRequest) (*EndDevices, error)
	// GetDevice returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(context.Context, *GetEndDeviceRequest) (*EndDevice, error)
	// SetDevice creates or updates the device.
	Set(context.Context, *SetDeviceRequest) (*EndDevice, error)
	// DeleteDevice deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(context.Context, *EndDeviceIdentifiers) (*types.Empty, error)
}

func RegisterNsDeviceRegistryServer(s *grpc.Server, srv NsDeviceRegistryServer) {
	s.RegisterService(&_NsDeviceRegistry_serviceDesc, srv)
}

func _NsDeviceRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsDeviceRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsDeviceRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsDeviceRegistryServer).List(ctx, req.(*ListEndDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsDeviceRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsDeviceRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsDeviceRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsDeviceRegistryServer).Get(ctx, req.(*GetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsDeviceRegistry_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsDeviceRegistryServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsDeviceRegistry/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsDeviceRegistryServer).Set(ctx, req.(*SetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsDeviceRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsDeviceRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsDeviceRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsDeviceRegistryServer).Delete(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _NsDeviceRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsDeviceRegistry",
	HandlerType: (*NsDeviceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _NsDeviceRegistry_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NsDeviceRegistry_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _NsDeviceRegistry_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NsDeviceRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/end_device_services.proto",
}

// Client API for AsDeviceRegistry service

type AsDeviceRegistryClient interface {
	// ListDevices returns the devices that match the given identifiers.
	List(ctx context.Context, in *ListEndDevicesRequest, opts ...grpc.CallOption) (*EndDevices, error)
	// GetDevice returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// SetDevice creates or updates the device.
	Set(ctx context.Context, in *SetDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// DeleteDevice deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type asDeviceRegistryClient struct {
	cc *grpc.ClientConn
}

func NewAsDeviceRegistryClient(cc *grpc.ClientConn) AsDeviceRegistryClient {
	return &asDeviceRegistryClient{cc}
}

func (c *asDeviceRegistryClient) List(ctx context.Context, in *ListEndDevicesRequest, opts ...grpc.CallOption) (*EndDevices, error) {
	out := new(EndDevices)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsDeviceRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asDeviceRegistryClient) Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsDeviceRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asDeviceRegistryClient) Set(ctx context.Context, in *SetDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsDeviceRegistry/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asDeviceRegistryClient) Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsDeviceRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AsDeviceRegistry service

type AsDeviceRegistryServer interface {
	// ListDevices returns the devices that match the given identifiers.
	List(context.Context, *ListEndDevicesRequest) (*EndDevices, error)
	// GetDevice returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(context.Context, *GetEndDeviceRequest) (*EndDevice, error)
	// SetDevice creates or updates the device.
	Set(context.Context, *SetDeviceRequest) (*EndDevice, error)
	// DeleteDevice deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(context.Context, *EndDeviceIdentifiers) (*types.Empty, error)
}

func RegisterAsDeviceRegistryServer(s *grpc.Server, srv AsDeviceRegistryServer) {
	s.RegisterService(&_AsDeviceRegistry_serviceDesc, srv)
}

func _AsDeviceRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsDeviceRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsDeviceRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsDeviceRegistryServer).List(ctx, req.(*ListEndDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsDeviceRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsDeviceRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsDeviceRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsDeviceRegistryServer).Get(ctx, req.(*GetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsDeviceRegistry_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsDeviceRegistryServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsDeviceRegistry/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsDeviceRegistryServer).Set(ctx, req.(*SetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsDeviceRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsDeviceRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsDeviceRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsDeviceRegistryServer).Delete(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _AsDeviceRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.AsDeviceRegistry",
	HandlerType: (*AsDeviceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AsDeviceRegistry_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AsDeviceRegistry_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _AsDeviceRegistry_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AsDeviceRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/end_device_services.proto",
}

// Client API for JsDeviceRegistry service

type JsDeviceRegistryClient interface {
	// ListDevices returns the devices that match the given identifiers.
	List(ctx context.Context, in *ListEndDevicesRequest, opts ...grpc.CallOption) (*EndDevices, error)
	// GetDevice returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// SetDevice creates or updates the device.
	Set(ctx context.Context, in *SetDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// DeleteDevice deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type jsDeviceRegistryClient struct {
	cc *grpc.ClientConn
}

func NewJsDeviceRegistryClient(cc *grpc.ClientConn) JsDeviceRegistryClient {
	return &jsDeviceRegistryClient{cc}
}

func (c *jsDeviceRegistryClient) List(ctx context.Context, in *ListEndDevicesRequest, opts ...grpc.CallOption) (*EndDevices, error) {
	out := new(EndDevices)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.JsDeviceRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jsDeviceRegistryClient) Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.JsDeviceRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jsDeviceRegistryClient) Set(ctx context.Context, in *SetDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.JsDeviceRegistry/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jsDeviceRegistryClient) Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.JsDeviceRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JsDeviceRegistry service

type JsDeviceRegistryServer interface {
	// ListDevices returns the devices that match the given identifiers.
	List(context.Context, *ListEndDevicesRequest) (*EndDevices, error)
	// GetDevice returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(context.Context, *GetEndDeviceRequest) (*EndDevice, error)
	// SetDevice creates or updates the device.
	Set(context.Context, *SetDeviceRequest) (*EndDevice, error)
	// DeleteDevice deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(context.Context, *EndDeviceIdentifiers) (*types.Empty, error)
}

func RegisterJsDeviceRegistryServer(s *grpc.Server, srv JsDeviceRegistryServer) {
	s.RegisterService(&_JsDeviceRegistry_serviceDesc, srv)
}

func _JsDeviceRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JsDeviceRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.JsDeviceRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JsDeviceRegistryServer).List(ctx, req.(*ListEndDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JsDeviceRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JsDeviceRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.JsDeviceRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JsDeviceRegistryServer).Get(ctx, req.(*GetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JsDeviceRegistry_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JsDeviceRegistryServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.JsDeviceRegistry/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JsDeviceRegistryServer).Set(ctx, req.(*SetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JsDeviceRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JsDeviceRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.JsDeviceRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JsDeviceRegistryServer).Delete(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _JsDeviceRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.JsDeviceRegistry",
	HandlerType: (*JsDeviceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _JsDeviceRegistry_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _JsDeviceRegistry_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _JsDeviceRegistry_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _JsDeviceRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/end_device_services.proto",
}

func init() {
	proto.RegisterFile("lorawan-stack/api/end_device_services.proto", fileDescriptor_end_device_services_3ede3ea37530bef6)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/end_device_services.proto", fileDescriptor_end_device_services_3ede3ea37530bef6)
}

var fileDescriptor_end_device_services_3ede3ea37530bef6 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x96, 0x3d, 0x4c, 0x14, 0x41,
	0x14, 0x80, 0xe7, 0x89, 0x52, 0x6c, 0x61, 0xcc, 0x16, 0x26, 0xae, 0xe6, 0xc5, 0x1c, 0xda, 0xa0,
	0xcc, 0x26, 0x52, 0x89, 0x15, 0x06, 0x42, 0x30, 0x6a, 0x01, 0x1d, 0x16, 0x64, 0xee, 0x6e, 0x58,
	0xe6, 0xee, 0x98, 0x59, 0x77, 0x06, 0x08, 0x31, 0x24, 0xc4, 0x8a, 0xc2, 0xc2, 0xc4, 0x98, 0x58,
	0x1a, 0x2b, 0x12, 0x1b, 0x42, 0x45, 0x49, 0x49, 0x49, 0x62, 0x43, 0xc9, 0xce, 0x5a, 0x5c, 0x49,
	0x27, 0xa5, 0xb9, 0xbd, 0x3d, 0xee, 0x2f, 0x67, 0xd4, 0x1c, 0xe4, 0xa8, 0x6e, 0xde, 0xcc, 0x7b,
	0xf3, 0xbe, 0x77, 0xf3, 0x15, 0xeb, 0x3c, 0xaa, 0xa8, 0x88, 0xad, 0x33, 0x39, 0xa6, 0x0d, 0x2b,
	0x94, 0x7d, 0x16, 0x0a, 0x9f, 0xcb, 0xe2, 0x62, 0x91, 0xaf, 0x89, 0x02, 0x5f, 0xd4, 0x3c, 0xaa,
	0xfd, 0x6a, 0x1a, 0x46, 0xca, 0x28, 0xf7, 0xa6, 0x31, 0x92, 0x66, 0x05, 0x74, 0x6d, 0xdc, 0x1b,
	0x0b, 0x84, 0x59, 0x5e, 0xcd, 0xd3, 0x82, 0x5a, 0xf1, 0x03, 0x15, 0x28, 0x3f, 0x4d, 0xcb, 0xaf,
	0x2e, 0xa5, 0x51, 0x1a, 0xa4, 0xab, 0x7a, 0xb9, 0x77, 0x2f, 0x50, 0x2a, 0xa8, 0xf0, 0xb4, 0x09,
	0x93, 0x52, 0x19, 0x66, 0x84, 0x92, 0xd9, 0xe5, 0xde, 0xdd, 0xec, 0xf4, 0xfc, 0x0e, 0xbe, 0x12,
	0x9a, 0x8d, 0xec, 0x30, 0xf7, 0x27, 0xcc, 0x2c, 0x67, 0xa4, 0x3b, 0x47, 0x14, 0xb9, 0x34, 0x62,
	0x49, 0xf0, 0x28, 0xeb, 0xf2, 0xa4, 0x7a, 0xc3, 0xb9, 0xf5, 0x5a, 0x4f, 0xa5, 0x75, 0x73, 0x3c,
	0x10, 0xda, 0x44, 0x1b, 0xee, 0x07, 0x70, 0xae, 0xbf, 0x14, 0xda, 0xb8, 0x0f, 0x69, 0xfb, 0x84,
	0xb4, 0xb6, 0x3b, 0x2d, 0x8b, 0xf5, 0x7c, 0x3d, 0xc7, 0xdf, 0xae, 0x72, 0x6d, 0x3c, 0xaf, 0x33,
	0xad, 0x99, 0x92, 0x9b, 0x7c, 0xff, 0xe3, 0xe7, 0xa7, 0x6b, 0xcf, 0xdc, 0xa7, 0xbe, 0xd4, 0x3e,
	0x0b, 0xc3, 0x8a, 0x28, 0xd4, 0xc7, 0xf4, 0xdf, 0xb5, 0x44, 0x8b, 0xa2, 0xa8, 0x69, 0x7b, 0xbc,
	0xe9, 0xd7, 0xa7, 0xd1, 0xee, 0x1e, 0x38, 0x43, 0x33, 0xdc, 0xb8, 0x23, 0x9d, 0x6d, 0x66, 0x78,
	0x13, 0xa6, 0xc1, 0x72, 0xa7, 0x27, 0x4b, 0xae, 0x9c, 0xa2, 0x70, 0xb7, 0xd0, 0x8d, 0xd2, 0xf2,
	0xc6, 0xdd, 0x24, 0x3d, 0xc9, 0xba, 0xea, 0xce, 0x97, 0x9b, 0xee, 0x2f, 0x70, 0x86, 0xe6, 0xb9,
	0x71, 0xef, 0x77, 0xf2, 0xcc, 0x73, 0xf3, 0xd7, 0xc4, 0x7b, 0x90, 0x22, 0x7f, 0x07, 0xef, 0x4d,
	0x37, 0x73, 0xf6, 0xd8, 0xff, 0xc4, 0xdb, 0x52, 0xd3, 0x64, 0x9d, 0x80, 0xd1, 0x85, 0xd9, 0xdc,
	0x54, 0x3f, 0x3a, 0x4c, 0xc0, 0xa8, 0xfb, 0x19, 0x9c, 0xe1, 0x29, 0x5e, 0xe1, 0x86, 0xbb, 0x0f,
	0x7a, 0x8e, 0x36, 0xdb, 0x34, 0xd1, 0xbb, 0x4d, 0xeb, 0xaa, 0xd3, 0x86, 0xea, 0x74, 0xba, 0xa6,
	0x7a, 0xee, 0x55, 0x3a, 0xfc, 0xcc, 0xe8, 0xf4, 0x7f, 0xab, 0xd3, 0x18, 0xa1, 0xb6, 0x97, 0xaa,
	0x3e, 0x79, 0x39, 0xaa, 0xb3, 0xc1, 0x51, 0x9d, 0x5d, 0x41, 0xd5, 0xd9, 0x85, 0xab, 0xce, 0x06,
	0x56, 0x75, 0xd6, 0x2f, 0xd5, 0x5f, 0x5c, 0x8e, 0xea, 0xa5, 0xc1, 0x51, 0xbd, 0x74, 0x05, 0x55,
	0x2f, 0x5d, 0xb8, 0xea, 0xa5, 0x81, 0x55, 0xbd, 0xd4, 0x07, 0xd5, 0x9f, 0x7f, 0x83, 0xc3, 0x18,
	0xe1, 0x28, 0x46, 0x38, 0x8e, 0x91, 0x9c, 0xc4, 0x48, 0xaa, 0x31, 0x92, 0xd3, 0x18, 0xc9, 0x59,
	0x8c, 0xb0, 0x65, 0x11, 0xb6, 0x2d, 0x92, 0x1d, 0x8b, 0xb0, 0x6b, 0x91, 0xec, 0x5b, 0x24, 0x07,
	0x16, 0xc9, 0xa1, 0x45, 0x38, 0xb2, 0x08, 0xc7, 0x16, 0xc9, 0x89, 0x45, 0xa8, 0x5a, 0x24, 0xa7,
	0x16, 0xe1, 0xcc, 0x22, 0xd9, 0x4a, 0x90, 0x6c, 0x27, 0x08, 0x1f, 0x13, 0x24, 0x5f, 0x12, 0x84,
	0xaf, 0x09, 0x92, 0x9d, 0x04, 0xc9, 0x6e, 0x82, 0xb0, 0x9f, 0x20, 0x1c, 0x24, 0x08, 0x0b, 0x8f,
	0x03, 0x45, 0xcd, 0x32, 0x37, 0xcb, 0x42, 0x06, 0x9a, 0x4a, 0x6e, 0xd6, 0x55, 0x54, 0xf6, 0xdb,
	0xbf, 0xb8, 0xc2, 0x72, 0xe0, 0x1b, 0x23, 0xc3, 0x7c, 0x7e, 0x38, 0xfd, 0x0f, 0xc6, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xdb, 0x61, 0x96, 0x69, 0x5e, 0x0a, 0x00, 0x00,
}
