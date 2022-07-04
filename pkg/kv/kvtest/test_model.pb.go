// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: test_model.proto

package kvtest

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type TestModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          []byte                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AnotherString string                 `protobuf:"bytes,2,opt,name=another_string,json=anotherString,proto3" json:"another_string,omitempty"`
	ADouble       float64                `protobuf:"fixed64,3,opt,name=a_double,json=aDouble,proto3" json:"a_double,omitempty"`
	TestTime      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=test_time,json=testTime,proto3" json:"test_time,omitempty"`
	TestMap       map[string]int32       `protobuf:"bytes,5,rep,name=test_map,json=testMap,proto3" json:"test_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	TestList      []bool                 `protobuf:"varint,6,rep,packed,name=test_list,json=testList,proto3" json:"test_list,omitempty"`
}

func (x *TestModel) Reset() {
	*x = TestModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestModel) ProtoMessage() {}

func (x *TestModel) ProtoReflect() protoreflect.Message {
	mi := &file_test_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestModel.ProtoReflect.Descriptor instead.
func (*TestModel) Descriptor() ([]byte, []int) {
	return file_test_model_proto_rawDescGZIP(), []int{0}
}

func (x *TestModel) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *TestModel) GetAnotherString() string {
	if x != nil {
		return x.AnotherString
	}
	return ""
}

func (x *TestModel) GetADouble() float64 {
	if x != nil {
		return x.ADouble
	}
	return 0
}

func (x *TestModel) GetTestTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TestTime
	}
	return nil
}

func (x *TestModel) GetTestMap() map[string]int32 {
	if x != nil {
		return x.TestMap
	}
	return nil
}

func (x *TestModel) GetTestList() []bool {
	if x != nil {
		return x.TestList
	}
	return nil
}

var File_test_model_proto protoreflect.FileDescriptor

var file_test_model_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x69, 0x6f, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x2e, 0x6c, 0x61, 0x6b, 0x65, 0x66, 0x73, 0x2e, 0x6b, 0x76, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc2, 0x02, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6e, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x5f, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x08,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x69, 0x6f, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x6c, 0x61,
	0x6b, 0x65, 0x66, 0x73, 0x2e, 0x6b, 0x76, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x08, 0x52, 0x08,
	0x74, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x6c, 0x61, 0x6b,
	0x65, 0x66, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6b, 0x76, 0x2f, 0x6b, 0x76, 0x74, 0x65, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_model_proto_rawDescOnce sync.Once
	file_test_model_proto_rawDescData = file_test_model_proto_rawDesc
)

func file_test_model_proto_rawDescGZIP() []byte {
	file_test_model_proto_rawDescOnce.Do(func() {
		file_test_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_model_proto_rawDescData)
	})
	return file_test_model_proto_rawDescData
}

var file_test_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_model_proto_goTypes = []interface{}{
	(*TestModel)(nil),             // 0: io.treeverse.lakefs.kvtest.TestModel
	nil,                           // 1: io.treeverse.lakefs.kvtest.TestModel.TestMapEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_test_model_proto_depIdxs = []int32{
	2, // 0: io.treeverse.lakefs.kvtest.TestModel.test_time:type_name -> google.protobuf.Timestamp
	1, // 1: io.treeverse.lakefs.kvtest.TestModel.test_map:type_name -> io.treeverse.lakefs.kvtest.TestModel.TestMapEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_test_model_proto_init() }
func file_test_model_proto_init() {
	if File_test_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestModel); i {
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
			RawDescriptor: file_test_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_model_proto_goTypes,
		DependencyIndexes: file_test_model_proto_depIdxs,
		MessageInfos:      file_test_model_proto_msgTypes,
	}.Build()
	File_test_model_proto = out.File
	file_test_model_proto_rawDesc = nil
	file_test_model_proto_goTypes = nil
	file_test_model_proto_depIdxs = nil
}
