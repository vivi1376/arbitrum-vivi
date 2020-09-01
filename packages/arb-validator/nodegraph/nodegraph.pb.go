//
// Copyright 2019, Offchain Labs, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: nodegraph.proto

package nodegraph

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	valprotocol "github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	structures "github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
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

type NodeGraphBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes               []*structures.NodeBuf       `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	OldestNodeHash      *common.HashBuf             `protobuf:"bytes,2,opt,name=oldestNodeHash,proto3" json:"oldestNodeHash,omitempty"`
	LatestConfirmedHash *common.HashBuf             `protobuf:"bytes,3,opt,name=latestConfirmedHash,proto3" json:"latestConfirmedHash,omitempty"`
	LeafHashes          []*common.HashBuf           `protobuf:"bytes,4,rep,name=leafHashes,proto3" json:"leafHashes,omitempty"`
	Params              *valprotocol.ChainParamsBuf `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *NodeGraphBuf) Reset() {
	*x = NodeGraphBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodegraph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGraphBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGraphBuf) ProtoMessage() {}

func (x *NodeGraphBuf) ProtoReflect() protoreflect.Message {
	mi := &file_nodegraph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGraphBuf.ProtoReflect.Descriptor instead.
func (*NodeGraphBuf) Descriptor() ([]byte, []int) {
	return file_nodegraph_proto_rawDescGZIP(), []int{0}
}

func (x *NodeGraphBuf) GetNodes() []*structures.NodeBuf {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *NodeGraphBuf) GetOldestNodeHash() *common.HashBuf {
	if x != nil {
		return x.OldestNodeHash
	}
	return nil
}

func (x *NodeGraphBuf) GetLatestConfirmedHash() *common.HashBuf {
	if x != nil {
		return x.LatestConfirmedHash
	}
	return nil
}

func (x *NodeGraphBuf) GetLeafHashes() []*common.HashBuf {
	if x != nil {
		return x.LeafHashes
	}
	return nil
}

func (x *NodeGraphBuf) GetParams() *valprotocol.ChainParamsBuf {
	if x != nil {
		return x.Params
	}
	return nil
}

type StakedNodeGraphBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeGraph  *NodeGraphBuf   `protobuf:"bytes,1,opt,name=nodeGraph,proto3" json:"nodeGraph,omitempty"`
	Stakers    []*StakerBuf    `protobuf:"bytes,2,rep,name=stakers,proto3" json:"stakers,omitempty"`
	Challenges []*ChallengeBuf `protobuf:"bytes,3,rep,name=challenges,proto3" json:"challenges,omitempty"`
}

func (x *StakedNodeGraphBuf) Reset() {
	*x = StakedNodeGraphBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodegraph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakedNodeGraphBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakedNodeGraphBuf) ProtoMessage() {}

func (x *StakedNodeGraphBuf) ProtoReflect() protoreflect.Message {
	mi := &file_nodegraph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakedNodeGraphBuf.ProtoReflect.Descriptor instead.
func (*StakedNodeGraphBuf) Descriptor() ([]byte, []int) {
	return file_nodegraph_proto_rawDescGZIP(), []int{1}
}

func (x *StakedNodeGraphBuf) GetNodeGraph() *NodeGraphBuf {
	if x != nil {
		return x.NodeGraph
	}
	return nil
}

func (x *StakedNodeGraphBuf) GetStakers() []*StakerBuf {
	if x != nil {
		return x.Stakers
	}
	return nil
}

func (x *StakedNodeGraphBuf) GetChallenges() []*ChallengeBuf {
	if x != nil {
		return x.Challenges
	}
	return nil
}

type StakerBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       *common.AddressBuf   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Location      *common.HashBuf      `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	CreationTime  *common.TimeTicksBuf `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	ChallengeAddr *common.AddressBuf   `protobuf:"bytes,4,opt,name=challengeAddr,proto3" json:"challengeAddr,omitempty"`
}

func (x *StakerBuf) Reset() {
	*x = StakerBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodegraph_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakerBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakerBuf) ProtoMessage() {}

func (x *StakerBuf) ProtoReflect() protoreflect.Message {
	mi := &file_nodegraph_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakerBuf.ProtoReflect.Descriptor instead.
func (*StakerBuf) Descriptor() ([]byte, []int) {
	return file_nodegraph_proto_rawDescGZIP(), []int{2}
}

func (x *StakerBuf) GetAddress() *common.AddressBuf {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *StakerBuf) GetLocation() *common.HashBuf {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *StakerBuf) GetCreationTime() *common.TimeTicksBuf {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *StakerBuf) GetChallengeAddr() *common.AddressBuf {
	if x != nil {
		return x.ChallengeAddr
	}
	return nil
}

type ChallengeBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId          *common.BlockIdBuf `protobuf:"bytes,1,opt,name=blockId,proto3" json:"blockId,omitempty"`
	LogIndex         uint64             `protobuf:"varint,2,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Asserter         *common.AddressBuf `protobuf:"bytes,3,opt,name=asserter,proto3" json:"asserter,omitempty"`
	Challenger       *common.AddressBuf `protobuf:"bytes,4,opt,name=challenger,proto3" json:"challenger,omitempty"`
	Contract         *common.AddressBuf `protobuf:"bytes,5,opt,name=contract,proto3" json:"contract,omitempty"`
	ConflictNodeHash *common.HashBuf    `protobuf:"bytes,6,opt,name=conflictNodeHash,proto3" json:"conflictNodeHash,omitempty"`
}

func (x *ChallengeBuf) Reset() {
	*x = ChallengeBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodegraph_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeBuf) ProtoMessage() {}

func (x *ChallengeBuf) ProtoReflect() protoreflect.Message {
	mi := &file_nodegraph_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeBuf.ProtoReflect.Descriptor instead.
func (*ChallengeBuf) Descriptor() ([]byte, []int) {
	return file_nodegraph_proto_rawDescGZIP(), []int{3}
}

func (x *ChallengeBuf) GetBlockId() *common.BlockIdBuf {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *ChallengeBuf) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *ChallengeBuf) GetAsserter() *common.AddressBuf {
	if x != nil {
		return x.Asserter
	}
	return nil
}

func (x *ChallengeBuf) GetChallenger() *common.AddressBuf {
	if x != nil {
		return x.Challenger
	}
	return nil
}

func (x *ChallengeBuf) GetContract() *common.AddressBuf {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *ChallengeBuf) GetConflictNodeHash() *common.HashBuf {
	if x != nil {
		return x.ConflictNodeHash
	}
	return nil
}

var File_nodegraph_proto protoreflect.FileDescriptor

var file_nodegraph_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x1a, 0x1c, 0x61, 0x72,
	0x62, 0x2d, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x61, 0x72, 0x62, 0x2d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x61, 0x72, 0x62, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x42, 0x75, 0x66, 0x12, 0x29, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x75, 0x66, 0x52, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0e, 0x6f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0e, 0x6f, 0x6c,
	0x64, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x41, 0x0a, 0x13,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x13, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x2f, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x66, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73,
	0x68, 0x42, 0x75, 0x66, 0x52, 0x0a, 0x6c, 0x65, 0x61, 0x66, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73,
	0x12, 0x33, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x75, 0x66, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x64,
	0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x42, 0x75, 0x66, 0x12, 0x35, 0x0a, 0x09,
	0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x42, 0x75, 0x66, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x75, 0x66, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6b,
	0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x75, 0x66,
	0x52, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x73, 0x22, 0xda, 0x01, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x75, 0x66, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x75, 0x66, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x42, 0x75,
	0x66, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x75, 0x66, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0xa9, 0x02, 0x0a, 0x0c, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x75, 0x66, 0x12, 0x2c, 0x0a, 0x07, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x42, 0x75, 0x66, 0x52,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x2e, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x75, 0x66, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65,
	0x72, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x75, 0x66, 0x52, 0x0a, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x75, 0x66, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66,
	0x6c, 0x69, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68,
	0x42, 0x75, 0x66, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x75, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x61, 0x72, 0x62, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_nodegraph_proto_rawDescOnce sync.Once
	file_nodegraph_proto_rawDescData = file_nodegraph_proto_rawDesc
)

func file_nodegraph_proto_rawDescGZIP() []byte {
	file_nodegraph_proto_rawDescOnce.Do(func() {
		file_nodegraph_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodegraph_proto_rawDescData)
	})
	return file_nodegraph_proto_rawDescData
}

var file_nodegraph_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nodegraph_proto_goTypes = []interface{}{
	(*NodeGraphBuf)(nil),               // 0: nodegraph.NodeGraphBuf
	(*StakedNodeGraphBuf)(nil),         // 1: nodegraph.StakedNodeGraphBuf
	(*StakerBuf)(nil),                  // 2: nodegraph.StakerBuf
	(*ChallengeBuf)(nil),               // 3: nodegraph.ChallengeBuf
	(*structures.NodeBuf)(nil),         // 4: structures.NodeBuf
	(*common.HashBuf)(nil),             // 5: common.HashBuf
	(*valprotocol.ChainParamsBuf)(nil), // 6: valprotocol.ChainParamsBuf
	(*common.AddressBuf)(nil),          // 7: common.AddressBuf
	(*common.TimeTicksBuf)(nil),        // 8: common.TimeTicksBuf
	(*common.BlockIdBuf)(nil),          // 9: common.BlockIdBuf
}
var file_nodegraph_proto_depIdxs = []int32{
	4,  // 0: nodegraph.NodeGraphBuf.nodes:type_name -> structures.NodeBuf
	5,  // 1: nodegraph.NodeGraphBuf.oldestNodeHash:type_name -> common.HashBuf
	5,  // 2: nodegraph.NodeGraphBuf.latestConfirmedHash:type_name -> common.HashBuf
	5,  // 3: nodegraph.NodeGraphBuf.leafHashes:type_name -> common.HashBuf
	6,  // 4: nodegraph.NodeGraphBuf.params:type_name -> valprotocol.ChainParamsBuf
	0,  // 5: nodegraph.StakedNodeGraphBuf.nodeGraph:type_name -> nodegraph.NodeGraphBuf
	2,  // 6: nodegraph.StakedNodeGraphBuf.stakers:type_name -> nodegraph.StakerBuf
	3,  // 7: nodegraph.StakedNodeGraphBuf.challenges:type_name -> nodegraph.ChallengeBuf
	7,  // 8: nodegraph.StakerBuf.address:type_name -> common.AddressBuf
	5,  // 9: nodegraph.StakerBuf.location:type_name -> common.HashBuf
	8,  // 10: nodegraph.StakerBuf.creationTime:type_name -> common.TimeTicksBuf
	7,  // 11: nodegraph.StakerBuf.challengeAddr:type_name -> common.AddressBuf
	9,  // 12: nodegraph.ChallengeBuf.blockId:type_name -> common.BlockIdBuf
	7,  // 13: nodegraph.ChallengeBuf.asserter:type_name -> common.AddressBuf
	7,  // 14: nodegraph.ChallengeBuf.challenger:type_name -> common.AddressBuf
	7,  // 15: nodegraph.ChallengeBuf.contract:type_name -> common.AddressBuf
	5,  // 16: nodegraph.ChallengeBuf.conflictNodeHash:type_name -> common.HashBuf
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_nodegraph_proto_init() }
func file_nodegraph_proto_init() {
	if File_nodegraph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodegraph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGraphBuf); i {
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
		file_nodegraph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakedNodeGraphBuf); i {
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
		file_nodegraph_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakerBuf); i {
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
		file_nodegraph_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeBuf); i {
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
			RawDescriptor: file_nodegraph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nodegraph_proto_goTypes,
		DependencyIndexes: file_nodegraph_proto_depIdxs,
		MessageInfos:      file_nodegraph_proto_msgTypes,
	}.Build()
	File_nodegraph_proto = out.File
	file_nodegraph_proto_rawDesc = nil
	file_nodegraph_proto_goTypes = nil
	file_nodegraph_proto_depIdxs = nil
}