// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: proto/grpc/examples/wallet/wallet.proto

package wallet

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

type BalancePerAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance int64  `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *BalancePerAddress) Reset() {
	*x = BalancePerAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_examples_wallet_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancePerAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancePerAddress) ProtoMessage() {}

func (x *BalancePerAddress) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_examples_wallet_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancePerAddress.ProtoReflect.Descriptor instead.
func (*BalancePerAddress) Descriptor() ([]byte, []int) {
	return file_proto_grpc_examples_wallet_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *BalancePerAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BalancePerAddress) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type BalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeBalancePerAddress bool `protobuf:"varint,1,opt,name=include_balance_per_address,json=includeBalancePerAddress,proto3" json:"include_balance_per_address,omitempty"`
}

func (x *BalanceRequest) Reset() {
	*x = BalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_examples_wallet_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceRequest) ProtoMessage() {}

func (x *BalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_examples_wallet_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceRequest.ProtoReflect.Descriptor instead.
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_examples_wallet_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *BalanceRequest) GetIncludeBalancePerAddress() bool {
	if x != nil {
		return x.IncludeBalancePerAddress
	}
	return false
}

type BalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance   int64                `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Addresses []*BalancePerAddress `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *BalanceResponse) Reset() {
	*x = BalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_examples_wallet_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceResponse) ProtoMessage() {}

func (x *BalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_examples_wallet_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceResponse.ProtoReflect.Descriptor instead.
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_examples_wallet_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *BalanceResponse) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *BalanceResponse) GetAddresses() []*BalancePerAddress {
	if x != nil {
		return x.Addresses
	}
	return nil
}

var File_proto_grpc_examples_wallet_wallet_proto protoreflect.FileDescriptor

var file_proto_grpc_examples_wallet_wallet_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22,
	0x47, 0x0a, 0x11, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x4f, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x18, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x72, 0x0a, 0x0f, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x32, 0xc8, 0x01,
	0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x5d, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x61, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x42, 0x0b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_examples_wallet_wallet_proto_rawDescOnce sync.Once
	file_proto_grpc_examples_wallet_wallet_proto_rawDescData = file_proto_grpc_examples_wallet_wallet_proto_rawDesc
)

func file_proto_grpc_examples_wallet_wallet_proto_rawDescGZIP() []byte {
	file_proto_grpc_examples_wallet_wallet_proto_rawDescOnce.Do(func() {
		file_proto_grpc_examples_wallet_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_examples_wallet_wallet_proto_rawDescData)
	})
	return file_proto_grpc_examples_wallet_wallet_proto_rawDescData
}

var file_proto_grpc_examples_wallet_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_grpc_examples_wallet_wallet_proto_goTypes = []interface{}{
	(*BalancePerAddress)(nil), // 0: grpc.examples.wallet.BalancePerAddress
	(*BalanceRequest)(nil),    // 1: grpc.examples.wallet.BalanceRequest
	(*BalanceResponse)(nil),   // 2: grpc.examples.wallet.BalanceResponse
}
var file_proto_grpc_examples_wallet_wallet_proto_depIdxs = []int32{
	0, // 0: grpc.examples.wallet.BalanceResponse.addresses:type_name -> grpc.examples.wallet.BalancePerAddress
	1, // 1: grpc.examples.wallet.Wallet.FetchBalance:input_type -> grpc.examples.wallet.BalanceRequest
	1, // 2: grpc.examples.wallet.Wallet.WatchBalance:input_type -> grpc.examples.wallet.BalanceRequest
	2, // 3: grpc.examples.wallet.Wallet.FetchBalance:output_type -> grpc.examples.wallet.BalanceResponse
	2, // 4: grpc.examples.wallet.Wallet.WatchBalance:output_type -> grpc.examples.wallet.BalanceResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_grpc_examples_wallet_wallet_proto_init() }
func file_proto_grpc_examples_wallet_wallet_proto_init() {
	if File_proto_grpc_examples_wallet_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_examples_wallet_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancePerAddress); i {
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
		file_proto_grpc_examples_wallet_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceRequest); i {
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
		file_proto_grpc_examples_wallet_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceResponse); i {
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
			RawDescriptor: file_proto_grpc_examples_wallet_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_examples_wallet_wallet_proto_goTypes,
		DependencyIndexes: file_proto_grpc_examples_wallet_wallet_proto_depIdxs,
		MessageInfos:      file_proto_grpc_examples_wallet_wallet_proto_msgTypes,
	}.Build()
	File_proto_grpc_examples_wallet_wallet_proto = out.File
	file_proto_grpc_examples_wallet_wallet_proto_rawDesc = nil
	file_proto_grpc_examples_wallet_wallet_proto_goTypes = nil
	file_proto_grpc_examples_wallet_wallet_proto_depIdxs = nil
}
