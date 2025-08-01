// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: github.com/openconfig/gnoi/debug/debug.proto

package debug

import (
	_ "github.com/openconfig/gnoi/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DebugRequest_Mode int32

const (
	DebugRequest_MODE_UNSPECIFIED DebugRequest_Mode = 0
	DebugRequest_MODE_SHELL       DebugRequest_Mode = 1
	DebugRequest_MODE_CLI         DebugRequest_Mode = 2
)

// Enum value maps for DebugRequest_Mode.
var (
	DebugRequest_Mode_name = map[int32]string{
		0: "MODE_UNSPECIFIED",
		1: "MODE_SHELL",
		2: "MODE_CLI",
	}
	DebugRequest_Mode_value = map[string]int32{
		"MODE_UNSPECIFIED": 0,
		"MODE_SHELL":       1,
		"MODE_CLI":         2,
	}
)

func (x DebugRequest_Mode) Enum() *DebugRequest_Mode {
	p := new(DebugRequest_Mode)
	*p = x
	return p
}

func (x DebugRequest_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DebugRequest_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_openconfig_gnoi_debug_debug_proto_enumTypes[0].Descriptor()
}

func (DebugRequest_Mode) Type() protoreflect.EnumType {
	return &file_github_com_openconfig_gnoi_debug_debug_proto_enumTypes[0]
}

func (x DebugRequest_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DebugRequest_Mode.Descriptor instead.
func (DebugRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP(), []int{0, 0}
}

type DebugRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mode          DebugRequest_Mode      `protobuf:"varint,1,opt,name=mode,proto3,enum=gnoi.debug.DebugRequest_Mode" json:"mode,omitempty"`
	Command       []byte                 `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	ByteLimit     int64                  `protobuf:"varint,3,opt,name=byte_limit,json=byteLimit,proto3" json:"byte_limit,omitempty"`
	Timeout       int64                  `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	RoleAccount   string                 `protobuf:"bytes,5,opt,name=role_account,json=roleAccount,proto3" json:"role_account,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DebugRequest) Reset() {
	*x = DebugRequest{}
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugRequest) ProtoMessage() {}

func (x *DebugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugRequest.ProtoReflect.Descriptor instead.
func (*DebugRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP(), []int{0}
}

func (x *DebugRequest) GetMode() DebugRequest_Mode {
	if x != nil {
		return x.Mode
	}
	return DebugRequest_MODE_UNSPECIFIED
}

func (x *DebugRequest) GetCommand() []byte {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *DebugRequest) GetByteLimit() int64 {
	if x != nil {
		return x.ByteLimit
	}
	return 0
}

func (x *DebugRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *DebugRequest) GetRoleAccount() string {
	if x != nil {
		return x.RoleAccount
	}
	return ""
}

type DebugResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Response:
	//
	//	*DebugResponse_Request
	//	*DebugResponse_Data
	//	*DebugResponse_Status
	Response      isDebugResponse_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DebugResponse) Reset() {
	*x = DebugResponse{}
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugResponse) ProtoMessage() {}

func (x *DebugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugResponse.ProtoReflect.Descriptor instead.
func (*DebugResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP(), []int{1}
}

func (x *DebugResponse) GetResponse() isDebugResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *DebugResponse) GetRequest() *DebugRequest {
	if x != nil {
		if x, ok := x.Response.(*DebugResponse_Request); ok {
			return x.Request
		}
	}
	return nil
}

func (x *DebugResponse) GetData() []byte {
	if x != nil {
		if x, ok := x.Response.(*DebugResponse_Data); ok {
			return x.Data
		}
	}
	return nil
}

func (x *DebugResponse) GetStatus() *DebugStatus {
	if x != nil {
		if x, ok := x.Response.(*DebugResponse_Status); ok {
			return x.Status
		}
	}
	return nil
}

type isDebugResponse_Response interface {
	isDebugResponse_Response()
}

type DebugResponse_Request struct {
	Request *DebugRequest `protobuf:"bytes,100,opt,name=request,proto3,oneof"`
}

type DebugResponse_Data struct {
	Data []byte `protobuf:"bytes,101,opt,name=data,proto3,oneof"`
}

type DebugResponse_Status struct {
	Status *DebugStatus `protobuf:"bytes,102,opt,name=status,proto3,oneof"`
}

func (*DebugResponse_Request) isDebugResponse_Response() {}

func (*DebugResponse_Data) isDebugResponse_Response() {}

func (*DebugResponse_Status) isDebugResponse_Response() {}

type DebugStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details       []*anypb.Any           `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DebugStatus) Reset() {
	*x = DebugStatus{}
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebugStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugStatus) ProtoMessage() {}

func (x *DebugStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugStatus.ProtoReflect.Descriptor instead.
func (*DebugStatus) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP(), []int{2}
}

func (x *DebugStatus) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DebugStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DebugStatus) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_github_com_openconfig_gnoi_debug_debug_proto protoreflect.FileDescriptor

const file_github_com_openconfig_gnoi_debug_debug_proto_rawDesc = "" +
	"\n" +
	",github.com/openconfig/gnoi/debug/debug.proto\x12\n" +
	"gnoi.debug\x1a,github.com/openconfig/gnoi/types/types.proto\x1a\x19google/protobuf/any.proto\"\xf3\x01\n" +
	"\fDebugRequest\x121\n" +
	"\x04mode\x18\x01 \x01(\x0e2\x1d.gnoi.debug.DebugRequest.ModeR\x04mode\x12\x18\n" +
	"\acommand\x18\x02 \x01(\fR\acommand\x12\x1d\n" +
	"\n" +
	"byte_limit\x18\x03 \x01(\x03R\tbyteLimit\x12\x18\n" +
	"\atimeout\x18\x04 \x01(\x03R\atimeout\x12!\n" +
	"\frole_account\x18\x05 \x01(\tR\vroleAccount\":\n" +
	"\x04Mode\x12\x14\n" +
	"\x10MODE_UNSPECIFIED\x10\x00\x12\x0e\n" +
	"\n" +
	"MODE_SHELL\x10\x01\x12\f\n" +
	"\bMODE_CLI\x10\x02\"\x9a\x01\n" +
	"\rDebugResponse\x124\n" +
	"\arequest\x18d \x01(\v2\x18.gnoi.debug.DebugRequestH\x00R\arequest\x12\x14\n" +
	"\x04data\x18e \x01(\fH\x00R\x04data\x121\n" +
	"\x06status\x18f \x01(\v2\x17.gnoi.debug.DebugStatusH\x00R\x06statusB\n" +
	"\n" +
	"\bresponse\"k\n" +
	"\vDebugStatus\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12.\n" +
	"\adetails\x18\x03 \x03(\v2\x14.google.protobuf.AnyR\adetails2G\n" +
	"\x05Debug\x12>\n" +
	"\x05Debug\x12\x18.gnoi.debug.DebugRequest\x1a\x19.gnoi.debug.DebugResponse0\x01B*\xd2>\x050.1.0Z github.com/openconfig/gnoi/debugb\x06proto3"

var (
	file_github_com_openconfig_gnoi_debug_debug_proto_rawDescOnce sync.Once
	file_github_com_openconfig_gnoi_debug_debug_proto_rawDescData []byte
)

func file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP() []byte {
	file_github_com_openconfig_gnoi_debug_debug_proto_rawDescOnce.Do(func() {
		file_github_com_openconfig_gnoi_debug_debug_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_github_com_openconfig_gnoi_debug_debug_proto_rawDesc), len(file_github_com_openconfig_gnoi_debug_debug_proto_rawDesc)))
	})
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescData
}

var file_github_com_openconfig_gnoi_debug_debug_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_openconfig_gnoi_debug_debug_proto_goTypes = []any{
	(DebugRequest_Mode)(0), // 0: gnoi.debug.DebugRequest.Mode
	(*DebugRequest)(nil),   // 1: gnoi.debug.DebugRequest
	(*DebugResponse)(nil),  // 2: gnoi.debug.DebugResponse
	(*DebugStatus)(nil),    // 3: gnoi.debug.DebugStatus
	(*anypb.Any)(nil),      // 4: google.protobuf.Any
}
var file_github_com_openconfig_gnoi_debug_debug_proto_depIdxs = []int32{
	0, // 0: gnoi.debug.DebugRequest.mode:type_name -> gnoi.debug.DebugRequest.Mode
	1, // 1: gnoi.debug.DebugResponse.request:type_name -> gnoi.debug.DebugRequest
	3, // 2: gnoi.debug.DebugResponse.status:type_name -> gnoi.debug.DebugStatus
	4, // 3: gnoi.debug.DebugStatus.details:type_name -> google.protobuf.Any
	1, // 4: gnoi.debug.Debug.Debug:input_type -> gnoi.debug.DebugRequest
	2, // 5: gnoi.debug.Debug.Debug:output_type -> gnoi.debug.DebugResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_openconfig_gnoi_debug_debug_proto_init() }
func file_github_com_openconfig_gnoi_debug_debug_proto_init() {
	if File_github_com_openconfig_gnoi_debug_debug_proto != nil {
		return
	}
	file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[1].OneofWrappers = []any{
		(*DebugResponse_Request)(nil),
		(*DebugResponse_Data)(nil),
		(*DebugResponse_Status)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_github_com_openconfig_gnoi_debug_debug_proto_rawDesc), len(file_github_com_openconfig_gnoi_debug_debug_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_openconfig_gnoi_debug_debug_proto_goTypes,
		DependencyIndexes: file_github_com_openconfig_gnoi_debug_debug_proto_depIdxs,
		EnumInfos:         file_github_com_openconfig_gnoi_debug_debug_proto_enumTypes,
		MessageInfos:      file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes,
	}.Build()
	File_github_com_openconfig_gnoi_debug_debug_proto = out.File
	file_github_com_openconfig_gnoi_debug_debug_proto_goTypes = nil
	file_github_com_openconfig_gnoi_debug_debug_proto_depIdxs = nil
}
