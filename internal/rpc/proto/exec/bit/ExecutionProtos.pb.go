// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: ExecutionProtos.proto

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

package bit

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/zeroshade/go-drill/internal/rpc/proto/exec"
	shared "github.com/zeroshade/go-drill/internal/rpc/proto/exec/shared"
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

type FragmentHandle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId         *shared.QueryId `protobuf:"bytes,1,opt,name=query_id,json=queryId" json:"query_id,omitempty"`
	MajorFragmentId *int32          `protobuf:"varint,2,opt,name=major_fragment_id,json=majorFragmentId" json:"major_fragment_id,omitempty"`
	MinorFragmentId *int32          `protobuf:"varint,3,opt,name=minor_fragment_id,json=minorFragmentId" json:"minor_fragment_id,omitempty"`
	ParentQueryId   *shared.QueryId `protobuf:"bytes,4,opt,name=parent_query_id,json=parentQueryId" json:"parent_query_id,omitempty"`
}

func (x *FragmentHandle) Reset() {
	*x = FragmentHandle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExecutionProtos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FragmentHandle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FragmentHandle) ProtoMessage() {}

func (x *FragmentHandle) ProtoReflect() protoreflect.Message {
	mi := &file_ExecutionProtos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FragmentHandle.ProtoReflect.Descriptor instead.
func (*FragmentHandle) Descriptor() ([]byte, []int) {
	return file_ExecutionProtos_proto_rawDescGZIP(), []int{0}
}

func (x *FragmentHandle) GetQueryId() *shared.QueryId {
	if x != nil {
		return x.QueryId
	}
	return nil
}

func (x *FragmentHandle) GetMajorFragmentId() int32 {
	if x != nil && x.MajorFragmentId != nil {
		return *x.MajorFragmentId
	}
	return 0
}

func (x *FragmentHandle) GetMinorFragmentId() int32 {
	if x != nil && x.MinorFragmentId != nil {
		return *x.MinorFragmentId
	}
	return 0
}

func (x *FragmentHandle) GetParentQueryId() *shared.QueryId {
	if x != nil {
		return x.ParentQueryId
	}
	return nil
}

//
// Prepared statement state on server side. Clients do not
// need to know the contents. They just need to submit it back to
// server when executing the prepared statement.
type ServerPreparedStatementState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SqlQuery *string `protobuf:"bytes,1,opt,name=sql_query,json=sqlQuery" json:"sql_query,omitempty"`
}

func (x *ServerPreparedStatementState) Reset() {
	*x = ServerPreparedStatementState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExecutionProtos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerPreparedStatementState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerPreparedStatementState) ProtoMessage() {}

func (x *ServerPreparedStatementState) ProtoReflect() protoreflect.Message {
	mi := &file_ExecutionProtos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerPreparedStatementState.ProtoReflect.Descriptor instead.
func (*ServerPreparedStatementState) Descriptor() ([]byte, []int) {
	return file_ExecutionProtos_proto_rawDescGZIP(), []int{1}
}

func (x *ServerPreparedStatementState) GetSqlQuery() string {
	if x != nil && x.SqlQuery != nil {
		return *x.SqlQuery
	}
	return ""
}

var File_ExecutionProtos_proto protoreflect.FileDescriptor

var file_ExecutionProtos_proto_rawDesc = []byte{
	0x0a, 0x15, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x62, 0x69,
	0x74, 0x1a, 0x12, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x74, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x0e, 0x46,
	0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2f, 0x0a,
	0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x64, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x11, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x5f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x61, 0x6a, 0x6f, 0x72,
	0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x69,
	0x6e, 0x6f, 0x72, 0x5f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x46, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x64, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x1c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x71, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x42, 0x66, 0x0a, 0x1b, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x64, 0x72, 0x69, 0x6c, 0x6c, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x42, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x65, 0x72, 0x6f, 0x73,
	0x68, 0x61, 0x64, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x72, 0x69, 0x6c, 0x6c, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x62, 0x69, 0x74,
}

var (
	file_ExecutionProtos_proto_rawDescOnce sync.Once
	file_ExecutionProtos_proto_rawDescData = file_ExecutionProtos_proto_rawDesc
)

func file_ExecutionProtos_proto_rawDescGZIP() []byte {
	file_ExecutionProtos_proto_rawDescOnce.Do(func() {
		file_ExecutionProtos_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExecutionProtos_proto_rawDescData)
	})
	return file_ExecutionProtos_proto_rawDescData
}

var file_ExecutionProtos_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ExecutionProtos_proto_goTypes = []interface{}{
	(*FragmentHandle)(nil),               // 0: exec.bit.FragmentHandle
	(*ServerPreparedStatementState)(nil), // 1: exec.bit.ServerPreparedStatementState
	(*shared.QueryId)(nil),               // 2: exec.shared.QueryId
}
var file_ExecutionProtos_proto_depIdxs = []int32{
	2, // 0: exec.bit.FragmentHandle.query_id:type_name -> exec.shared.QueryId
	2, // 1: exec.bit.FragmentHandle.parent_query_id:type_name -> exec.shared.QueryId
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ExecutionProtos_proto_init() }
func file_ExecutionProtos_proto_init() {
	if File_ExecutionProtos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ExecutionProtos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FragmentHandle); i {
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
		file_ExecutionProtos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerPreparedStatementState); i {
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
			RawDescriptor: file_ExecutionProtos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExecutionProtos_proto_goTypes,
		DependencyIndexes: file_ExecutionProtos_proto_depIdxs,
		MessageInfos:      file_ExecutionProtos_proto_msgTypes,
	}.Build()
	File_ExecutionProtos_proto = out.File
	file_ExecutionProtos_proto_rawDesc = nil
	file_ExecutionProtos_proto_goTypes = nil
	file_ExecutionProtos_proto_depIdxs = nil
}
