// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: region.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{0}
}

func (x *IP) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{1}
}

func (x *Region) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type LatLong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
}

func (x *LatLong) Reset() {
	*x = LatLong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatLong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatLong) ProtoMessage() {}

func (x *LatLong) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatLong.ProtoReflect.Descriptor instead.
func (*LatLong) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{2}
}

func (x *LatLong) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LatLong) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_region_proto protoreflect.FileDescriptor

var file_region_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x20, 0x0a, 0x06, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x43, 0x0a,
	0x07, 0x4c, 0x61, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x4c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x32, 0x62, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x09, 0x49, 0x50, 0x32, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x50, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0a, 0x49,
	0x50, 0x32, 0x4c, 0x61, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x50, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x74,
	0x4c, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_region_proto_rawDescOnce sync.Once
	file_region_proto_rawDescData = file_region_proto_rawDesc
)

func file_region_proto_rawDescGZIP() []byte {
	file_region_proto_rawDescOnce.Do(func() {
		file_region_proto_rawDescData = protoimpl.X.CompressGZIP(file_region_proto_rawDescData)
	})
	return file_region_proto_rawDescData
}

var file_region_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_region_proto_goTypes = []interface{}{
	(*IP)(nil),      // 0: proto.IP
	(*Region)(nil),  // 1: proto.Region
	(*LatLong)(nil), // 2: proto.LatLong
}
var file_region_proto_depIdxs = []int32{
	0, // 0: proto.RegionServer.IP2Region:input_type -> proto.IP
	0, // 1: proto.RegionServer.IP2LatLong:input_type -> proto.IP
	1, // 2: proto.RegionServer.IP2Region:output_type -> proto.Region
	2, // 3: proto.RegionServer.IP2LatLong:output_type -> proto.LatLong
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_region_proto_init() }
func file_region_proto_init() {
	if File_region_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_region_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
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
		file_region_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
		file_region_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatLong); i {
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
			RawDescriptor: file_region_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_region_proto_goTypes,
		DependencyIndexes: file_region_proto_depIdxs,
		MessageInfos:      file_region_proto_msgTypes,
	}.Build()
	File_region_proto = out.File
	file_region_proto_rawDesc = nil
	file_region_proto_goTypes = nil
	file_region_proto_depIdxs = nil
}
