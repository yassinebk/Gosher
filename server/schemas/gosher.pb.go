// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/gosher.proto

package schemas

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

type Custom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Custom) Reset() {
	*x = Custom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Custom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Custom) ProtoMessage() {}

func (x *Custom) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Custom.ProtoReflect.Descriptor instead.
func (*Custom) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{0}
}

func (x *Custom) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CustomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CustomResponse) Reset() {
	*x = CustomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomResponse) ProtoMessage() {}

func (x *CustomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomResponse.ProtoReflect.Descriptor instead.
func (*CustomResponse) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{1}
}

func (x *CustomResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mime  string `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"`
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{2}
}

func (x *UploadRequest) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *UploadRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type FileTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FileTransfer) Reset() {
	*x = FileTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTransfer) ProtoMessage() {}

func (x *FileTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileTransfer.ProtoReflect.Descriptor instead.
func (*FileTransfer) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{3}
}

func (x *FileTransfer) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_gosher_proto protoreflect.FileDescriptor

var file_proto_gosher_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x22, 0x1a, 0x0a,
	0x06, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22, 0x0a, 0x0e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x39, 0x0a,
	0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x20, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xb2, 0x01, 0x0a, 0x0a, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x53, 0x61, 0x79,
	0x48, 0x69, 0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x1a, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x34, 0x0a, 0x08, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x1a, 0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x30, 0x01, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gosher_proto_rawDescOnce sync.Once
	file_proto_gosher_proto_rawDescData = file_proto_gosher_proto_rawDesc
)

func file_proto_gosher_proto_rawDescGZIP() []byte {
	file_proto_gosher_proto_rawDescOnce.Do(func() {
		file_proto_gosher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gosher_proto_rawDescData)
	})
	return file_proto_gosher_proto_rawDescData
}

var file_proto_gosher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_gosher_proto_goTypes = []interface{}{
	(*Custom)(nil),         // 0: schemas.Custom
	(*CustomResponse)(nil), // 1: schemas.CustomResponse
	(*UploadRequest)(nil),  // 2: schemas.UploadRequest
	(*FileTransfer)(nil),   // 3: schemas.FileTransfer
}
var file_proto_gosher_proto_depIdxs = []int32{
	0, // 0: schemas.HelloWorld.SayHi:input_type -> schemas.Custom
	2, // 1: schemas.HelloWorld.Upload:input_type -> schemas.UploadRequest
	0, // 2: schemas.HelloWorld.Download:input_type -> schemas.Custom
	1, // 3: schemas.HelloWorld.SayHi:output_type -> schemas.CustomResponse
	1, // 4: schemas.HelloWorld.Upload:output_type -> schemas.CustomResponse
	3, // 5: schemas.HelloWorld.Download:output_type -> schemas.FileTransfer
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_gosher_proto_init() }
func file_proto_gosher_proto_init() {
	if File_proto_gosher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gosher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Custom); i {
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
		file_proto_gosher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomResponse); i {
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
		file_proto_gosher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_proto_gosher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileTransfer); i {
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
			RawDescriptor: file_proto_gosher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gosher_proto_goTypes,
		DependencyIndexes: file_proto_gosher_proto_depIdxs,
		MessageInfos:      file_proto_gosher_proto_msgTypes,
	}.Build()
	File_proto_gosher_proto = out.File
	file_proto_gosher_proto_rawDesc = nil
	file_proto_gosher_proto_goTypes = nil
	file_proto_gosher_proto_depIdxs = nil
}
