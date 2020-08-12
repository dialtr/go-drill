// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: GeneralRPC.proto

//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package rpc

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/zeroshade/go-drill/internal/rpc/proto/exec"
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

type RpcMode int32

const (
	RpcMode_REQUEST          RpcMode = 0
	RpcMode_RESPONSE         RpcMode = 1
	RpcMode_RESPONSE_FAILURE RpcMode = 2
	RpcMode_PING             RpcMode = 3
	RpcMode_PONG             RpcMode = 4
)

// Enum value maps for RpcMode.
var (
	RpcMode_name = map[int32]string{
		0: "REQUEST",
		1: "RESPONSE",
		2: "RESPONSE_FAILURE",
		3: "PING",
		4: "PONG",
	}
	RpcMode_value = map[string]int32{
		"REQUEST":          0,
		"RESPONSE":         1,
		"RESPONSE_FAILURE": 2,
		"PING":             3,
		"PONG":             4,
	}
)

func (x RpcMode) Enum() *RpcMode {
	p := new(RpcMode)
	*p = x
	return p
}

func (x RpcMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RpcMode) Descriptor() protoreflect.EnumDescriptor {
	return file_GeneralRPC_proto_enumTypes[0].Descriptor()
}

func (RpcMode) Type() protoreflect.EnumType {
	return &file_GeneralRPC_proto_enumTypes[0]
}

func (x RpcMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RpcMode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RpcMode(num)
	return nil
}

// Deprecated: Use RpcMode.Descriptor instead.
func (RpcMode) EnumDescriptor() ([]byte, []int) {
	return file_GeneralRPC_proto_rawDescGZIP(), []int{0}
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok *bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GeneralRPC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_GeneralRPC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_GeneralRPC_proto_rawDescGZIP(), []int{0}
}

func (x *Ack) GetOk() bool {
	if x != nil && x.Ok != nil {
		return *x.Ok
	}
	return false
}

type RpcHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode           *RpcMode `protobuf:"varint,1,opt,name=mode,enum=exec.rpc.RpcMode" json:"mode,omitempty"`
	CoordinationId *int32   `protobuf:"varint,2,opt,name=coordination_id,json=coordinationId" json:"coordination_id,omitempty"` // reusable coordination identifier.  Sender defines.  Server returns on return.  Irrelevant for purely single direction rpc.
	RpcType        *int32   `protobuf:"varint,3,opt,name=rpc_type,json=rpcType" json:"rpc_type,omitempty"`                      // a rpc mode specific rpc type.
}

func (x *RpcHeader) Reset() {
	*x = RpcHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GeneralRPC_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcHeader) ProtoMessage() {}

func (x *RpcHeader) ProtoReflect() protoreflect.Message {
	mi := &file_GeneralRPC_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcHeader.ProtoReflect.Descriptor instead.
func (*RpcHeader) Descriptor() ([]byte, []int) {
	return file_GeneralRPC_proto_rawDescGZIP(), []int{1}
}

func (x *RpcHeader) GetMode() RpcMode {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return RpcMode_REQUEST
}

func (x *RpcHeader) GetCoordinationId() int32 {
	if x != nil && x.CoordinationId != nil {
		return *x.CoordinationId
	}
	return 0
}

func (x *RpcHeader) GetRpcType() int32 {
	if x != nil && x.RpcType != nil {
		return *x.RpcType
	}
	return 0
}

type CompleteRpcMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *RpcHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`                                 // required
	ProtobufBody []byte     `protobuf:"bytes,2,opt,name=protobuf_body,json=protobufBody" json:"protobuf_body,omitempty"` // required
	RawBody      []byte     `protobuf:"bytes,3,opt,name=raw_body,json=rawBody" json:"raw_body,omitempty"`                // optional
}

func (x *CompleteRpcMessage) Reset() {
	*x = CompleteRpcMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GeneralRPC_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteRpcMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteRpcMessage) ProtoMessage() {}

func (x *CompleteRpcMessage) ProtoReflect() protoreflect.Message {
	mi := &file_GeneralRPC_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteRpcMessage.ProtoReflect.Descriptor instead.
func (*CompleteRpcMessage) Descriptor() ([]byte, []int) {
	return file_GeneralRPC_proto_rawDescGZIP(), []int{2}
}

func (x *CompleteRpcMessage) GetHeader() *RpcHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CompleteRpcMessage) GetProtobufBody() []byte {
	if x != nil {
		return x.ProtobufBody
	}
	return nil
}

func (x *CompleteRpcMessage) GetRawBody() []byte {
	if x != nil {
		return x.RawBody
	}
	return nil
}

var File_GeneralRPC_proto protoreflect.FileDescriptor

var file_GeneralRPC_proto_rawDesc = []byte{
	0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x50, 0x43, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x12, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x15, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x76, 0x0a, 0x09, 0x52, 0x70, 0x63, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x70,
	0x63, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x70, 0x63, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x81, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x70, 0x63, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x70, 0x63, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x61, 0x77, 0x42,
	0x6f, 0x64, 0x79, 0x2a, 0x4e, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53,
	0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4e,
	0x47, 0x10, 0x04, 0x42, 0x6c, 0x0a, 0x1b, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x64, 0x72, 0x69, 0x6c, 0x6c, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x50, 0x43, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x65, 0x72, 0x6f, 0x73, 0x68, 0x61, 0x64, 0x65, 0x2f, 0x67, 0x6f, 0x2d,
	0x64, 0x72, 0x69, 0x6c, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x72, 0x70,
	0x63,
}

var (
	file_GeneralRPC_proto_rawDescOnce sync.Once
	file_GeneralRPC_proto_rawDescData = file_GeneralRPC_proto_rawDesc
)

func file_GeneralRPC_proto_rawDescGZIP() []byte {
	file_GeneralRPC_proto_rawDescOnce.Do(func() {
		file_GeneralRPC_proto_rawDescData = protoimpl.X.CompressGZIP(file_GeneralRPC_proto_rawDescData)
	})
	return file_GeneralRPC_proto_rawDescData
}

var file_GeneralRPC_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GeneralRPC_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_GeneralRPC_proto_goTypes = []interface{}{
	(RpcMode)(0),               // 0: exec.rpc.RpcMode
	(*Ack)(nil),                // 1: exec.rpc.Ack
	(*RpcHeader)(nil),          // 2: exec.rpc.RpcHeader
	(*CompleteRpcMessage)(nil), // 3: exec.rpc.CompleteRpcMessage
}
var file_GeneralRPC_proto_depIdxs = []int32{
	0, // 0: exec.rpc.RpcHeader.mode:type_name -> exec.rpc.RpcMode
	2, // 1: exec.rpc.CompleteRpcMessage.header:type_name -> exec.rpc.RpcHeader
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GeneralRPC_proto_init() }
func file_GeneralRPC_proto_init() {
	if File_GeneralRPC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GeneralRPC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_GeneralRPC_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcHeader); i {
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
		file_GeneralRPC_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteRpcMessage); i {
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
			RawDescriptor: file_GeneralRPC_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GeneralRPC_proto_goTypes,
		DependencyIndexes: file_GeneralRPC_proto_depIdxs,
		EnumInfos:         file_GeneralRPC_proto_enumTypes,
		MessageInfos:      file_GeneralRPC_proto_msgTypes,
	}.Build()
	File_GeneralRPC_proto = out.File
	file_GeneralRPC_proto_rawDesc = nil
	file_GeneralRPC_proto_goTypes = nil
	file_GeneralRPC_proto_depIdxs = nil
}
