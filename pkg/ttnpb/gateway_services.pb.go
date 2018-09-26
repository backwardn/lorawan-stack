// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/gateway_services.proto

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

import strings "strings"
import reflect "reflect"

import io "io"

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

type PullGatewayConfigurationRequest struct {
	GatewayIdentifiers   `protobuf:"bytes,1,opt,name=gateway_ids,json=gatewayIds,embedded=gateway_ids" json:"gateway_ids"`
	FieldMask            types.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask" json:"field_mask"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PullGatewayConfigurationRequest) Reset()      { *m = PullGatewayConfigurationRequest{} }
func (*PullGatewayConfigurationRequest) ProtoMessage() {}
func (*PullGatewayConfigurationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_services_ec773ed0747ed4c5, []int{0}
}
func (m *PullGatewayConfigurationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PullGatewayConfigurationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PullGatewayConfigurationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PullGatewayConfigurationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullGatewayConfigurationRequest.Merge(dst, src)
}
func (m *PullGatewayConfigurationRequest) XXX_Size() int {
	return m.Size()
}
func (m *PullGatewayConfigurationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullGatewayConfigurationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullGatewayConfigurationRequest proto.InternalMessageInfo

func (m *PullGatewayConfigurationRequest) GetFieldMask() types.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return types.FieldMask{}
}

func init() {
	proto.RegisterType((*PullGatewayConfigurationRequest)(nil), "ttn.lorawan.v3.PullGatewayConfigurationRequest")
	golang_proto.RegisterType((*PullGatewayConfigurationRequest)(nil), "ttn.lorawan.v3.PullGatewayConfigurationRequest")
}
func (this *PullGatewayConfigurationRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PullGatewayConfigurationRequest)
	if !ok {
		that2, ok := that.(PullGatewayConfigurationRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.GatewayIdentifiers.Equal(&that1.GatewayIdentifiers) {
		return false
	}
	if !this.FieldMask.Equal(&that1.FieldMask) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GatewayRegistry service

type GatewayRegistryClient interface {
	// Create a new gateway. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*Gateway, error)
	// Get the gateway with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	Get(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*Gateway, error)
	// List gateways. See request message for details.
	List(ctx context.Context, in *ListGatewaysRequest, opts ...grpc.CallOption) (*Gateways, error)
	Update(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*Gateway, error)
	Delete(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type gatewayRegistryClient struct {
	cc *grpc.ClientConn
}

func NewGatewayRegistryClient(cc *grpc.ClientConn) GatewayRegistryClient {
	return &gatewayRegistryClient{cc}
}

func (c *gatewayRegistryClient) Create(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	out := new(Gateway)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) Get(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	out := new(Gateway)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) List(ctx context.Context, in *ListGatewaysRequest, opts ...grpc.CallOption) (*Gateways, error) {
	out := new(Gateways)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) Update(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	out := new(Gateway)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayRegistryClient) Delete(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GatewayRegistry service

type GatewayRegistryServer interface {
	// Create a new gateway. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateGatewayRequest) (*Gateway, error)
	// Get the gateway with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	Get(context.Context, *GetGatewayRequest) (*Gateway, error)
	// List gateways. See request message for details.
	List(context.Context, *ListGatewaysRequest) (*Gateways, error)
	Update(context.Context, *UpdateGatewayRequest) (*Gateway, error)
	Delete(context.Context, *GatewayIdentifiers) (*types.Empty, error)
}

func RegisterGatewayRegistryServer(s *grpc.Server, srv GatewayRegistryServer) {
	s.RegisterService(&_GatewayRegistry_serviceDesc, srv)
}

func _GatewayRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Create(ctx, req.(*CreateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Get(ctx, req.(*GetGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).List(ctx, req.(*ListGatewaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Update(ctx, req.(*UpdateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayRegistryServer).Delete(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GatewayRegistry",
	HandlerType: (*GatewayRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GatewayRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GatewayRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GatewayRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GatewayRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GatewayRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/gateway_services.proto",
}

// Client API for GatewayAccess service

type GatewayAccessClient interface {
	ListRights(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	CreateAPIKey(ctx context.Context, in *CreateGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	ListAPIKeys(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*APIKeys, error)
	// Update the rights of an existing gateway API key. To generate an API key,
	// the CreateGatewayAPIKey should be used. To delete an API key, update it
	// with zero rights.
	UpdateAPIKey(ctx context.Context, in *UpdateGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Set the rights of a collaborator on the gateway. Users or organizations
	// are considered to be a collaborator if they have at least one right on the
	// gateway.
	SetCollaborator(ctx context.Context, in *SetGatewayCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error)
	ListCollaborators(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*Collaborators, error)
}

type gatewayAccessClient struct {
	cc *grpc.ClientConn
}

func NewGatewayAccessClient(cc *grpc.ClientConn) GatewayAccessClient {
	return &gatewayAccessClient{cc}
}

func (c *gatewayAccessClient) ListRights(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) CreateAPIKey(ctx context.Context, in *CreateGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/CreateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) ListAPIKeys(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*APIKeys, error) {
	out := new(APIKeys)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/ListAPIKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) UpdateAPIKey(ctx context.Context, in *UpdateGatewayAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/UpdateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) SetCollaborator(ctx context.Context, in *SetGatewayCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/SetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAccessClient) ListCollaborators(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GatewayAccess/ListCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GatewayAccess service

type GatewayAccessServer interface {
	ListRights(context.Context, *GatewayIdentifiers) (*Rights, error)
	CreateAPIKey(context.Context, *CreateGatewayAPIKeyRequest) (*APIKey, error)
	ListAPIKeys(context.Context, *GatewayIdentifiers) (*APIKeys, error)
	// Update the rights of an existing gateway API key. To generate an API key,
	// the CreateGatewayAPIKey should be used. To delete an API key, update it
	// with zero rights.
	UpdateAPIKey(context.Context, *UpdateGatewayAPIKeyRequest) (*APIKey, error)
	// Set the rights of a collaborator on the gateway. Users or organizations
	// are considered to be a collaborator if they have at least one right on the
	// gateway.
	SetCollaborator(context.Context, *SetGatewayCollaboratorRequest) (*types.Empty, error)
	ListCollaborators(context.Context, *GatewayIdentifiers) (*Collaborators, error)
}

func RegisterGatewayAccessServer(s *grpc.Server, srv GatewayAccessServer) {
	s.RegisterService(&_GatewayAccess_serviceDesc, srv)
}

func _GatewayAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).ListRights(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGatewayAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/CreateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).CreateAPIKey(ctx, req.(*CreateGatewayAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/ListAPIKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).ListAPIKeys(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_UpdateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).UpdateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/UpdateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).UpdateAPIKey(ctx, req.(*UpdateGatewayAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGatewayCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/SetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).SetCollaborator(ctx, req.(*SetGatewayCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GatewayAccess/ListCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAccessServer).ListCollaborators(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GatewayAccess",
	HandlerType: (*GatewayAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _GatewayAccess_ListRights_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _GatewayAccess_CreateAPIKey_Handler,
		},
		{
			MethodName: "ListAPIKeys",
			Handler:    _GatewayAccess_ListAPIKeys_Handler,
		},
		{
			MethodName: "UpdateAPIKey",
			Handler:    _GatewayAccess_UpdateAPIKey_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _GatewayAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _GatewayAccess_ListCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/gateway_services.proto",
}

// Client API for GatewayConfigurator service

type GatewayConfiguratorClient interface {
	PullConfiguration(ctx context.Context, in *PullGatewayConfigurationRequest, opts ...grpc.CallOption) (GatewayConfigurator_PullConfigurationClient, error)
}

type gatewayConfiguratorClient struct {
	cc *grpc.ClientConn
}

func NewGatewayConfiguratorClient(cc *grpc.ClientConn) GatewayConfiguratorClient {
	return &gatewayConfiguratorClient{cc}
}

func (c *gatewayConfiguratorClient) PullConfiguration(ctx context.Context, in *PullGatewayConfigurationRequest, opts ...grpc.CallOption) (GatewayConfigurator_PullConfigurationClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GatewayConfigurator_serviceDesc.Streams[0], "/ttn.lorawan.v3.GatewayConfigurator/PullConfiguration", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayConfiguratorPullConfigurationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GatewayConfigurator_PullConfigurationClient interface {
	Recv() (*Gateway, error)
	grpc.ClientStream
}

type gatewayConfiguratorPullConfigurationClient struct {
	grpc.ClientStream
}

func (x *gatewayConfiguratorPullConfigurationClient) Recv() (*Gateway, error) {
	m := new(Gateway)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GatewayConfigurator service

type GatewayConfiguratorServer interface {
	PullConfiguration(*PullGatewayConfigurationRequest, GatewayConfigurator_PullConfigurationServer) error
}

func RegisterGatewayConfiguratorServer(s *grpc.Server, srv GatewayConfiguratorServer) {
	s.RegisterService(&_GatewayConfigurator_serviceDesc, srv)
}

func _GatewayConfigurator_PullConfiguration_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullGatewayConfigurationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayConfiguratorServer).PullConfiguration(m, &gatewayConfiguratorPullConfigurationServer{stream})
}

type GatewayConfigurator_PullConfigurationServer interface {
	Send(*Gateway) error
	grpc.ServerStream
}

type gatewayConfiguratorPullConfigurationServer struct {
	grpc.ServerStream
}

func (x *gatewayConfiguratorPullConfigurationServer) Send(m *Gateway) error {
	return x.ServerStream.SendMsg(m)
}

var _GatewayConfigurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GatewayConfigurator",
	HandlerType: (*GatewayConfiguratorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullConfiguration",
			Handler:       _GatewayConfigurator_PullConfiguration_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/gateway_services.proto",
}

func (m *PullGatewayConfigurationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PullGatewayConfigurationRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGatewayServices(dAtA, i, uint64(m.GatewayIdentifiers.Size()))
	n1, err := m.GatewayIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGatewayServices(dAtA, i, uint64(m.FieldMask.Size()))
	n2, err := m.FieldMask.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func encodeVarintGatewayServices(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedPullGatewayConfigurationRequest(r randyGatewayServices, easy bool) *PullGatewayConfigurationRequest {
	this := &PullGatewayConfigurationRequest{}
	v1 := NewPopulatedGatewayIdentifiers(r, easy)
	this.GatewayIdentifiers = *v1
	v2 := types.NewPopulatedFieldMask(r, easy)
	this.FieldMask = *v2
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyGatewayServices interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneGatewayServices(r randyGatewayServices) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringGatewayServices(r randyGatewayServices) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneGatewayServices(r)
	}
	return string(tmps)
}
func randUnrecognizedGatewayServices(r randyGatewayServices, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldGatewayServices(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldGatewayServices(dAtA []byte, r randyGatewayServices, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateGatewayServices(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateGatewayServices(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateGatewayServices(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateGatewayServices(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateGatewayServices(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateGatewayServices(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateGatewayServices(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *PullGatewayConfigurationRequest) Size() (n int) {
	var l int
	_ = l
	l = m.GatewayIdentifiers.Size()
	n += 1 + l + sovGatewayServices(uint64(l))
	l = m.FieldMask.Size()
	n += 1 + l + sovGatewayServices(uint64(l))
	return n
}

func sovGatewayServices(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGatewayServices(x uint64) (n int) {
	return sovGatewayServices((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *PullGatewayConfigurationRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PullGatewayConfigurationRequest{`,
		`GatewayIdentifiers:` + strings.Replace(strings.Replace(this.GatewayIdentifiers.String(), "GatewayIdentifiers", "GatewayIdentifiers", 1), `&`, ``, 1) + `,`,
		`FieldMask:` + strings.Replace(strings.Replace(this.FieldMask.String(), "FieldMask", "types.FieldMask", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGatewayServices(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PullGatewayConfigurationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayServices
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PullGatewayConfigurationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PullGatewayConfigurationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGatewayServices
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GatewayIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldMask", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGatewayServices
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FieldMask.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayServices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatewayServices
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGatewayServices(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGatewayServices
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGatewayServices
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGatewayServices
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGatewayServices
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGatewayServices
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGatewayServices(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGatewayServices = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGatewayServices   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("lorawan-stack/api/gateway_services.proto", fileDescriptor_gateway_services_ec773ed0747ed4c5)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/gateway_services.proto", fileDescriptor_gateway_services_ec773ed0747ed4c5)
}

var fileDescriptor_gateway_services_ec773ed0747ed4c5 = []byte{
	// 911 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x3d, 0x8c, 0x1b, 0x45,
	0x14, 0x9e, 0x09, 0x91, 0x45, 0xe6, 0x02, 0x51, 0x06, 0x29, 0xb1, 0x16, 0x32, 0x4e, 0x36, 0x07,
	0x89, 0x4c, 0xbc, 0x7b, 0x38, 0x02, 0x21, 0x10, 0x42, 0x39, 0x03, 0xa7, 0x13, 0x89, 0x14, 0x39,
	0xa2, 0x71, 0x63, 0xad, 0xed, 0xf1, 0x7a, 0xe4, 0xf5, 0x8e, 0xd9, 0x19, 0xfb, 0x64, 0x4e, 0x27,
	0x45, 0x11, 0x45, 0x42, 0x43, 0x24, 0x84, 0xa0, 0x44, 0x54, 0x29, 0x43, 0x17, 0x89, 0x26, 0xe5,
	0x95, 0x27, 0xa5, 0xb9, 0x2a, 0x8a, 0x77, 0x29, 0x52, 0xa6, 0x4c, 0x89, 0x76, 0x76, 0xd7, 0x3f,
	0x6b, 0x9b, 0xdb, 0x93, 0xe8, 0xd6, 0xef, 0xbd, 0xf9, 0xbe, 0xf7, 0xbe, 0xf7, 0xde, 0x8c, 0xd1,
	0x55, 0x87, 0x7b, 0xd6, 0x8e, 0xe5, 0x96, 0x84, 0xb4, 0x9a, 0x5d, 0xd3, 0xea, 0x33, 0xd3, 0xb6,
	0x24, 0xdd, 0xb1, 0x46, 0x75, 0x41, 0xbd, 0x21, 0x6b, 0x52, 0x61, 0xf4, 0x3d, 0x2e, 0x39, 0x7e,
	0x5b, 0x4a, 0xd7, 0x88, 0xa3, 0x8d, 0xe1, 0x75, 0xad, 0x64, 0x33, 0xd9, 0x19, 0x34, 0x8c, 0x26,
	0xef, 0x99, 0x36, 0xb7, 0xb9, 0xa9, 0xc2, 0x1a, 0x83, 0xb6, 0xfa, 0xa5, 0x7e, 0xa8, 0xaf, 0xe8,
	0xb8, 0xf6, 0x9e, 0xcd, 0xb9, 0xed, 0x50, 0xc5, 0x60, 0xb9, 0x2e, 0x97, 0x96, 0x64, 0xdc, 0x8d,
	0xc1, 0xb5, 0x77, 0x63, 0xef, 0x04, 0x83, 0xf6, 0xfa, 0x72, 0x14, 0x3b, 0x2f, 0xa6, 0x9d, 0x6d,
	0x46, 0x9d, 0x56, 0xbd, 0x67, 0x89, 0x6e, 0x1c, 0x51, 0x58, 0x59, 0x45, 0x1c, 0x70, 0x79, 0x31,
	0x80, 0xb5, 0xa8, 0x2b, 0x59, 0x9b, 0x51, 0x2f, 0x49, 0x82, 0x2c, 0x06, 0x79, 0xcc, 0xee, 0xc8,
	0xd8, 0xaf, 0xff, 0x05, 0x51, 0xe1, 0xf6, 0xc0, 0x71, 0xb6, 0x22, 0xe8, 0x0a, 0x77, 0xdb, 0xcc,
	0x1e, 0x78, 0xaa, 0x90, 0x2a, 0xfd, 0x7e, 0x40, 0x85, 0xc4, 0xb7, 0xd0, 0x5a, 0xa2, 0x1f, 0x6b,
	0x89, 0x3c, 0xbc, 0x08, 0xaf, 0xae, 0x95, 0x75, 0x63, 0x5e, 0x3b, 0x23, 0x46, 0xd8, 0x9e, 0xa6,
	0xb0, 0xf9, 0xe6, 0xfe, 0xf3, 0x02, 0x38, 0x78, 0x5e, 0x80, 0x55, 0x64, 0x27, 0x5e, 0x81, 0xbf,
	0x44, 0x68, 0x5a, 0x6c, 0xfe, 0x84, 0x42, 0xd3, 0x8c, 0x48, 0x0f, 0x23, 0xd1, 0xc3, 0xf8, 0x26,
	0x0c, 0xb9, 0x65, 0x89, 0xee, 0xe6, 0xc9, 0x10, 0xa5, 0x7a, 0xaa, 0x9d, 0x18, 0xca, 0x3f, 0xe7,
	0xd0, 0x99, 0x98, 0xad, 0x4a, 0x6d, 0x26, 0xa4, 0x37, 0xc2, 0xcf, 0x20, 0xca, 0x55, 0x3c, 0x6a,
	0x49, 0x8a, 0xd7, 0xd3, 0x99, 0x45, 0xf6, 0xc9, 0x09, 0x55, 0x94, 0x76, 0x7e, 0x45, 0xfe, 0xfa,
	0x03, 0x78, 0xef, 0xd9, 0x3f, 0xbf, 0x9c, 0xb8, 0x07, 0x75, 0xd3, 0x1c, 0x08, 0xea, 0x09, 0x73,
	0xb7, 0xc9, 0x1d, 0xc7, 0x6a, 0x70, 0xcf, 0x92, 0xdc, 0x33, 0x42, 0x5b, 0xa8, 0x43, 0xf2, 0xb1,
	0x97, 0xb4, 0x45, 0x7c, 0x06, 0x8b, 0xb5, 0x9b, 0xfa, 0x96, 0xc9, 0x3d, 0xdb, 0x72, 0xd9, 0x0f,
	0xd1, 0x30, 0xa4, 0x4e, 0xcf, 0xfa, 0x14, 0x4a, 0xca, 0x30, 0x87, 0x86, 0x7b, 0xe8, 0x8d, 0x2d,
	0x2a, 0xf1, 0xa5, 0x85, 0x5c, 0xa9, 0xcc, 0x5a, 0x4e, 0x51, 0x55, 0xb3, 0x8e, 0xf5, 0x09, 0xb0,
	0xb9, 0x3b, 0xd3, 0x4d, 0x63, 0xfa, 0xbd, 0x87, 0x0f, 0x21, 0x3a, 0x79, 0x93, 0x09, 0x89, 0x2f,
	0xa7, 0xd1, 0x42, 0x6b, 0x8c, 0x28, 0x12, 0xca, 0xfc, 0x0a, 0x4a, 0xa1, 0x3f, 0x8c, 0x24, 0x7c,
	0x00, 0xf1, 0xa9, 0x09, 0x6b, 0xed, 0x23, 0x7c, 0x5c, 0x3d, 0x6b, 0xdb, 0xf8, 0xff, 0x12, 0x13,
	0x0f, 0x51, 0xee, 0xbb, 0x7e, 0x6b, 0xe9, 0x78, 0x44, 0xf6, 0xac, 0x7a, 0x96, 0x54, 0x69, 0x57,
	0xb4, 0x25, 0x7a, 0x1a, 0x29, 0x3d, 0xc3, 0x0e, 0xb6, 0x50, 0xee, 0x2b, 0xea, 0x50, 0x49, 0x71,
	0x86, 0x85, 0xd1, 0xce, 0x2d, 0xac, 0xc1, 0xd7, 0xe1, 0x9d, 0xa1, 0x13, 0x45, 0x9a, 0x2f, 0x9e,
	0x5b, 0xda, 0xc4, 0xbd, 0xf2, 0xdf, 0x39, 0xf4, 0x56, 0x0c, 0x77, 0xa3, 0xd9, 0xa4, 0x42, 0x60,
	0x8e, 0x50, 0xd8, 0xb3, 0xaa, 0xda, 0xf5, 0x8c, 0xdc, 0xa9, 0x98, 0xe8, 0xac, 0xfe, 0xbe, 0xe2,
	0x2e, 0xe0, 0x0b, 0xcb, 0xb9, 0xe3, 0xeb, 0x04, 0xff, 0x04, 0xd1, 0xe9, 0x68, 0xd1, 0x6e, 0xdc,
	0xde, 0xfe, 0x96, 0x8e, 0x70, 0xf1, 0x3f, 0xd7, 0x30, 0x0a, 0x4a, 0xd4, 0x5e, 0xe0, 0x8e, 0xdc,
	0xfa, 0x27, 0x8a, 0x7b, 0x43, 0xff, 0xf0, 0xe8, 0xe1, 0x0d, 0xaf, 0xb6, 0x52, 0x97, 0x46, 0x7b,
	0x23, 0xd0, 0x5a, 0x58, 0x7d, 0x84, 0x92, 0xad, 0xfc, 0xf3, 0xcb, 0x53, 0x10, 0xfa, 0x15, 0x95,
	0xc3, 0x25, 0x5c, 0x58, 0x51, 0x7f, 0xc2, 0x8b, 0x7f, 0x83, 0xe8, 0x74, 0x34, 0x4b, 0xab, 0x14,
	0x98, 0x9b, 0xb4, 0x6c, 0x0a, 0x54, 0x14, 0xfb, 0x17, 0xda, 0xa7, 0xc7, 0x50, 0xc0, 0xdc, 0xb5,
	0xfa, 0xac, 0xde, 0xa5, 0xe1, 0x40, 0xaa, 0x21, 0xfc, 0x15, 0xa2, 0x33, 0x77, 0xa8, 0xac, 0xcc,
	0xec, 0x0e, 0x2e, 0xa5, 0x09, 0xef, 0x4c, 0xee, 0x94, 0xd9, 0xb8, 0x69, 0x7e, 0xcb, 0x27, 0xf3,
	0x73, 0x95, 0xdf, 0xc7, 0xda, 0x46, 0x86, 0xfc, 0x66, 0x77, 0x57, 0xb5, 0xe9, 0x47, 0x88, 0xce,
	0x86, 0x7d, 0x9a, 0x25, 0xcc, 0xd6, 0xad, 0x0b, 0x0b, 0xc3, 0x35, 0x0b, 0xa1, 0x5f, 0x53, 0x59,
	0x7d, 0x80, 0xd7, 0x57, 0xf4, 0x6c, 0x2e, 0x93, 0xf2, 0x10, 0xbd, 0xb3, 0xf0, 0xfc, 0x71, 0x0f,
	0xd7, 0xd1, 0xd9, 0xf0, 0x65, 0x9c, 0x7b, 0x12, 0xb1, 0x99, 0x26, 0x3e, 0xe2, 0xf1, 0x5c, 0x79,
	0x91, 0x6c, 0xc0, 0xcd, 0x3f, 0xe1, 0xfe, 0x98, 0xc0, 0x83, 0x31, 0x81, 0x87, 0x63, 0x02, 0x5e,
	0x8c, 0x09, 0x78, 0x39, 0x26, 0xe0, 0xd5, 0x98, 0x80, 0xd7, 0x63, 0x02, 0xef, 0xfa, 0x04, 0xde,
	0xf7, 0x09, 0x78, 0xe4, 0x13, 0xf8, 0xd8, 0x27, 0xe0, 0x89, 0x4f, 0xc0, 0x53, 0x9f, 0x80, 0x7d,
	0x9f, 0xc0, 0x03, 0x9f, 0xc0, 0x43, 0x9f, 0x80, 0x17, 0x3e, 0x81, 0x2f, 0x7d, 0x02, 0x5e, 0xf9,
	0x04, 0xbe, 0xf6, 0x09, 0xb8, 0x1b, 0x10, 0x70, 0x3f, 0x20, 0xf0, 0x61, 0x40, 0xc0, 0xef, 0x01,
	0x81, 0x7f, 0x04, 0x04, 0x3c, 0x0a, 0x08, 0x78, 0x1c, 0x10, 0xf8, 0x24, 0x20, 0xf0, 0x69, 0x40,
	0x60, 0xed, 0x9a, 0xcd, 0x0d, 0xd9, 0xa1, 0xb2, 0xc3, 0x5c, 0x5b, 0x18, 0x2e, 0x95, 0x3b, 0xdc,
	0xeb, 0x9a, 0xf3, 0xff, 0x13, 0xfa, 0x5d, 0xdb, 0x94, 0xd2, 0xed, 0x37, 0x1a, 0x39, 0xd5, 0xf0,
	0xeb, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x56, 0x7d, 0xc2, 0x61, 0x55, 0x09, 0x00, 0x00,
}
