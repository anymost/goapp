// Code generated by protoc-gen-go. DO NOT EDIT.
// source: route_guide.proto

package routeguide

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type RouteRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteRequest) Reset()         { *m = RouteRequest{} }
func (m *RouteRequest) String() string { return proto.CompactTextString(m) }
func (*RouteRequest) ProtoMessage()    {}
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d679f20da65b7b, []int{0}
}

func (m *RouteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteRequest.Unmarshal(m, b)
}
func (m *RouteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteRequest.Marshal(b, m, deterministic)
}
func (m *RouteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteRequest.Merge(m, src)
}
func (m *RouteRequest) XXX_Size() int {
	return xxx_messageInfo_RouteRequest.Size(m)
}
func (m *RouteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteRequest proto.InternalMessageInfo

func (m *RouteRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type RouteResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteResponse) Reset()         { *m = RouteResponse{} }
func (m *RouteResponse) String() string { return proto.CompactTextString(m) }
func (*RouteResponse) ProtoMessage()    {}
func (*RouteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d679f20da65b7b, []int{1}
}

func (m *RouteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteResponse.Unmarshal(m, b)
}
func (m *RouteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteResponse.Marshal(b, m, deterministic)
}
func (m *RouteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteResponse.Merge(m, src)
}
func (m *RouteResponse) XXX_Size() int {
	return xxx_messageInfo_RouteResponse.Size(m)
}
func (m *RouteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteResponse proto.InternalMessageInfo

func (m *RouteResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*RouteRequest)(nil), "routeguide.RouteRequest")
	proto.RegisterType((*RouteResponse)(nil), "routeguide.RouteResponse")
}

func init() { proto.RegisterFile("route_guide.proto", fileDescriptor_b7d679f20da65b7b) }

var fileDescriptor_b7d679f20da65b7b = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xca, 0x2f, 0x2d,
	0x49, 0x8d, 0x4f, 0x2f, 0xcd, 0x4c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02,
	0x0b, 0x81, 0x45, 0x94, 0x94, 0xb8, 0x78, 0x82, 0x40, 0xbc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x30, 0x5b, 0x49, 0x99, 0x8b, 0x17, 0xaa, 0xa6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0x9b, 0x22,
	0xa3, 0x55, 0x4c, 0x5c, 0x5c, 0x60, 0x55, 0xee, 0x20, 0x73, 0x85, 0x9c, 0xb9, 0xb8, 0xdc, 0x53,
	0x4b, 0xdc, 0x52, 0x13, 0x4b, 0x4a, 0x8b, 0x52, 0x85, 0x24, 0xf4, 0x10, 0x56, 0xea, 0x21, 0xdb,
	0x27, 0x25, 0x89, 0x45, 0x06, 0x62, 0x8b, 0x12, 0x83, 0x90, 0x3b, 0x17, 0x8f, 0x4f, 0x66, 0x31,
	0xcc, 0x94, 0x62, 0x32, 0x8d, 0x31, 0x60, 0x14, 0x72, 0xe3, 0xe2, 0x0e, 0x4a, 0x4d, 0xce, 0x2f,
	0x4a, 0x01, 0x4b, 0x91, 0x69, 0x8e, 0x06, 0xc8, 0x1c, 0x4e, 0xb0, 0xa0, 0x73, 0x46, 0x62, 0x09,
	0xd9, 0xa6, 0x18, 0x30, 0x3a, 0x19, 0x70, 0x49, 0x67, 0xe6, 0xeb, 0xa5, 0x17, 0x15, 0x24, 0xeb,
	0xa5, 0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4, 0x16, 0x23, 0x69, 0x71, 0xe2, 0x47, 0x04, 0x64, 0x00,
	0x28, 0xc6, 0x02, 0x18, 0x93, 0xd8, 0xc0, 0x51, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x13,
	0xaa, 0x71, 0xd1, 0xcf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteGuideClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) ListFeatures(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[0], "/routeguide.RouteGuide/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListFeaturesClient interface {
	Recv() (*RouteResponse, error)
	grpc.ClientStream
}

type routeGuideListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideListFeaturesClient) Recv() (*RouteResponse, error) {
	m := new(RouteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[1], "/routeguide.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*RouteRequest) error
	CloseAndRecv() (*RouteResponse, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *RouteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[2], "/routeguide.RouteGuide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRouteChatClient{stream}
	return x, nil
}

type RouteGuide_RouteChatClient interface {
	Send(*RouteRequest) error
	Recv() (*RouteResponse, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRouteChatClient) Recv() (*RouteResponse, error) {
	m := new(RouteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServer is the server API for RouteGuide service.
type RouteGuideServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *RouteRequest) (*RouteResponse, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(*RouteRequest, RouteGuide_ListFeaturesServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(RouteGuide_RecordRouteServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(RouteGuide_RouteChatServer) error
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RouteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).ListFeatures(m, &routeGuideListFeaturesServer{stream})
}

type RouteGuide_ListFeaturesServer interface {
	Send(*RouteResponse) error
	grpc.ServerStream
}

type routeGuideListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideListFeaturesServer) Send(m *RouteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordRoute(&routeGuideRecordRouteServer{stream})
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteResponse) error
	Recv() (*RouteRequest, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*RouteRequest, error) {
	m := new(RouteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RouteChat(&routeGuideRouteChatServer{stream})
}

type RouteGuide_RouteChatServer interface {
	Send(*RouteResponse) error
	Recv() (*RouteRequest, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRouteChatServer) Recv() (*RouteRequest, error) {
	m := new(RouteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuide_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "route_guide.proto",
}
