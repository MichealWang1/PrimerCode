// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.9.0
// source: protocol.proto

package protocol

import (
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

type CmdId int32

const (
	CmdId_Cmd_Head   CmdId = 0
	CmdId_Cmd_Login  CmdId = 1001
	CmdId_Cmd_Logout CmdId = 1002
	CmdId_Cmd_Kick   CmdId = 1003
)

// Enum value maps for CmdId.
var (
	CmdId_name = map[int32]string{
		0:    "Cmd_Head",
		1001: "Cmd_Login",
		1002: "Cmd_Logout",
		1003: "Cmd_Kick",
	}
	CmdId_value = map[string]int32{
		"Cmd_Head":   0,
		"Cmd_Login":  1001,
		"Cmd_Logout": 1002,
		"Cmd_Kick":   1003,
	}
)

func (x CmdId) Enum() *CmdId {
	p := new(CmdId)
	*p = x
	return p
}

func (x CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_proto_enumTypes[0].Descriptor()
}

func (CmdId) Type() protoreflect.EnumType {
	return &file_protocol_proto_enumTypes[0]
}

func (x CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdId.Descriptor instead.
func (CmdId) EnumDescriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MyString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MyString) Reset() {
	*x = MyString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyString) ProtoMessage() {}

func (x *MyString) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyString.ProtoReflect.Descriptor instead.
func (*MyString) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *MyString) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protocol_proto protoreflect.FileDescriptor

var file_protocol_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x24, 0x0a, 0x08, 0x4d, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x45, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6d, 0x64, 0x5f, 0x48, 0x65, 0x61, 0x64, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x09, 0x43, 0x6d, 0x64, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0xe9, 0x07, 0x12, 0x0f,
	0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x5f, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x10, 0xea, 0x07, 0x12,
	0x0d, 0x0a, 0x08, 0x43, 0x6d, 0x64, 0x5f, 0x4b, 0x69, 0x63, 0x6b, 0x10, 0xeb, 0x07, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_proto_rawDescOnce sync.Once
	file_protocol_proto_rawDescData = file_protocol_proto_rawDesc
)

func file_protocol_proto_rawDescGZIP() []byte {
	file_protocol_proto_rawDescOnce.Do(func() {
		file_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_proto_rawDescData)
	})
	return file_protocol_proto_rawDescData
}

var file_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_proto_goTypes = []interface{}{
	(CmdId)(0),       // 0: protocol.CmdId
	(*Error)(nil),    // 1: protocol.Error
	(*MyString)(nil), // 2: protocol.MyString
}
var file_protocol_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_proto_init() }
func file_protocol_proto_init() {
	if File_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyString); i {
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
			RawDescriptor: file_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_proto_goTypes,
		DependencyIndexes: file_protocol_proto_depIdxs,
		EnumInfos:         file_protocol_proto_enumTypes,
		MessageInfos:      file_protocol_proto_msgTypes,
	}.Build()
	File_protocol_proto = out.File
	file_protocol_proto_rawDesc = nil
	file_protocol_proto_goTypes = nil
	file_protocol_proto_depIdxs = nil
}
