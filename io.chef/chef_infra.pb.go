// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: chef_infra.proto

package io_chef

import (
	context "context"
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

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major int32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch int32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chef_infra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_chef_infra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_chef_infra_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetMajor() int32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() int32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetPatch() int32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

type Cookbook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version *Version `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Cookbook) Reset() {
	*x = Cookbook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chef_infra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookbook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookbook) ProtoMessage() {}

func (x *Cookbook) ProtoReflect() protoreflect.Message {
	mi := &file_chef_infra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookbook.ProtoReflect.Descriptor instead.
func (*Cookbook) Descriptor() ([]byte, []int) {
	return file_chef_infra_proto_rawDescGZIP(), []int{1}
}

func (x *Cookbook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cookbook) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

var File_chef_infra_proto protoreflect.FileDescriptor

var file_chef_infra_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x68, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x22, 0x4b,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x22, 0x4d, 0x0a, 0x08, 0x43,
	0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x4e, 0x0a, 0x09, 0x43, 0x68,
	0x65, 0x66, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x12, 0x41, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x62,
	0x6f, 0x6f, 0x6b, 0x1a, 0x13, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x69, 0x6f,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chef_infra_proto_rawDescOnce sync.Once
	file_chef_infra_proto_rawDescData = file_chef_infra_proto_rawDesc
)

func file_chef_infra_proto_rawDescGZIP() []byte {
	file_chef_infra_proto_rawDescOnce.Do(func() {
		file_chef_infra_proto_rawDescData = protoimpl.X.CompressGZIP(file_chef_infra_proto_rawDescData)
	})
	return file_chef_infra_proto_rawDescData
}

var file_chef_infra_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chef_infra_proto_goTypes = []interface{}{
	(*Version)(nil),  // 0: chef_infra.Version
	(*Cookbook)(nil), // 1: chef_infra.Cookbook
}
var file_chef_infra_proto_depIdxs = []int32{
	0, // 0: chef_infra.Cookbook.version:type_name -> chef_infra.Version
	1, // 1: chef_infra.ChefInfra.GetCookbookVersion:input_type -> chef_infra.Cookbook
	0, // 2: chef_infra.ChefInfra.GetCookbookVersion:output_type -> chef_infra.Version
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chef_infra_proto_init() }
func file_chef_infra_proto_init() {
	if File_chef_infra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chef_infra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_chef_infra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookbook); i {
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
			RawDescriptor: file_chef_infra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chef_infra_proto_goTypes,
		DependencyIndexes: file_chef_infra_proto_depIdxs,
		MessageInfos:      file_chef_infra_proto_msgTypes,
	}.Build()
	File_chef_infra_proto = out.File
	file_chef_infra_proto_rawDesc = nil
	file_chef_infra_proto_goTypes = nil
	file_chef_infra_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChefInfraClient is the client API for ChefInfra service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChefInfraClient interface {
	GetCookbookVersion(ctx context.Context, in *Cookbook, opts ...grpc.CallOption) (*Version, error)
}

type chefInfraClient struct {
	cc grpc.ClientConnInterface
}

func NewChefInfraClient(cc grpc.ClientConnInterface) ChefInfraClient {
	return &chefInfraClient{cc}
}

func (c *chefInfraClient) GetCookbookVersion(ctx context.Context, in *Cookbook, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/chef_infra.ChefInfra/GetCookbookVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChefInfraServer is the server API for ChefInfra service.
type ChefInfraServer interface {
	GetCookbookVersion(context.Context, *Cookbook) (*Version, error)
}

// UnimplementedChefInfraServer can be embedded to have forward compatible implementations.
type UnimplementedChefInfraServer struct {
}

func (*UnimplementedChefInfraServer) GetCookbookVersion(context.Context, *Cookbook) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCookbookVersion not implemented")
}

func RegisterChefInfraServer(s *grpc.Server, srv ChefInfraServer) {
	s.RegisterService(&_ChefInfra_serviceDesc, srv)
}

func _ChefInfra_GetCookbookVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookbook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefInfraServer).GetCookbookVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef_infra.ChefInfra/GetCookbookVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefInfraServer).GetCookbookVersion(ctx, req.(*Cookbook))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChefInfra_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef_infra.ChefInfra",
	HandlerType: (*ChefInfraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCookbookVersion",
			Handler:    _ChefInfra_GetCookbookVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chef_infra.proto",
}
