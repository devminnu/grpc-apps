// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: location.proto

package location

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AvgTemprature     string  `protobuf:"bytes,3,opt,name=avgTemprature,proto3" json:"avgTemprature,omitempty"`
	AvgSoilTemprature float32 `protobuf:"fixed32,4,opt,name=avgSoilTemprature,proto3" json:"avgSoilTemprature,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Location) GetAvgTemprature() string {
	if x != nil {
		return x.AvgTemprature
	}
	return ""
}

func (x *Location) GetAvgSoilTemprature() float32 {
	if x != nil {
		return x.AvgSoilTemprature
	}
	return 0
}

type LocationID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LocationID) Reset() {
	*x = LocationID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationID) ProtoMessage() {}

func (x *LocationID) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationID.ProtoReflect.Descriptor instead.
func (*LocationID) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{1}
}

func (x *LocationID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LocationUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *LocationUpdateResponse) Reset() {
	*x = LocationUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationUpdateResponse) ProtoMessage() {}

func (x *LocationUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationUpdateResponse.ProtoReflect.Descriptor instead.
func (*LocationUpdateResponse) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{2}
}

func (x *LocationUpdateResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

var File_location_proto protoreflect.FileDescriptor

var file_location_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x76, 0x67, 0x54, 0x65, 0x6d, 0x70, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x76, 0x67, 0x54, 0x65, 0x6d,
	0x70, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x76, 0x67, 0x53, 0x6f,
	0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x11, 0x61, 0x76, 0x67, 0x53, 0x6f, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x16, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x38, 0x0a, 0x0b, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x67, 0x72, 0x12, 0x29, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0b, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x62, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_location_proto_rawDescOnce sync.Once
	file_location_proto_rawDescData = file_location_proto_rawDesc
)

func file_location_proto_rawDescGZIP() []byte {
	file_location_proto_rawDescOnce.Do(func() {
		file_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_location_proto_rawDescData)
	})
	return file_location_proto_rawDescData
}

var file_location_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_location_proto_goTypes = []interface{}{
	(*Location)(nil),               // 0: Location
	(*LocationID)(nil),             // 1: LocationID
	(*LocationUpdateResponse)(nil), // 2: LocationUpdateResponse
}
var file_location_proto_depIdxs = []int32{
	1, // 0: LocationMgr.getLocationByID:input_type -> LocationID
	0, // 1: LocationMgr.getLocationByID:output_type -> Location
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_location_proto_init() }
func file_location_proto_init() {
	if File_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationID); i {
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
		file_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationUpdateResponse); i {
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
			RawDescriptor: file_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_location_proto_goTypes,
		DependencyIndexes: file_location_proto_depIdxs,
		MessageInfos:      file_location_proto_msgTypes,
	}.Build()
	File_location_proto = out.File
	file_location_proto_rawDesc = nil
	file_location_proto_goTypes = nil
	file_location_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LocationMgrClient is the client API for LocationMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocationMgrClient interface {
	GetLocationByID(ctx context.Context, in *LocationID, opts ...grpc.CallOption) (*Location, error)
}

type locationMgrClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationMgrClient(cc grpc.ClientConnInterface) LocationMgrClient {
	return &locationMgrClient{cc}
}

func (c *locationMgrClient) GetLocationByID(ctx context.Context, in *LocationID, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/LocationMgr/getLocationByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationMgrServer is the server API for LocationMgr service.
type LocationMgrServer interface {
	GetLocationByID(context.Context, *LocationID) (*Location, error)
}

// UnimplementedLocationMgrServer can be embedded to have forward compatible implementations.
type UnimplementedLocationMgrServer struct {
}

func (*UnimplementedLocationMgrServer) GetLocationByID(context.Context, *LocationID) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationByID not implemented")
}

func RegisterLocationMgrServer(s *grpc.Server, srv LocationMgrServer) {
	s.RegisterService(&_LocationMgr_serviceDesc, srv)
}

func _LocationMgr_GetLocationByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationMgrServer).GetLocationByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocationMgr/GetLocationByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationMgrServer).GetLocationByID(ctx, req.(*LocationID))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocationMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LocationMgr",
	HandlerType: (*LocationMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getLocationByID",
			Handler:    _LocationMgr_GetLocationByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "location.proto",
}
