// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// merkle avl tree
type LeafNode struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Size                 int32    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeafNode) Reset()         { *m = LeafNode{} }
func (m *LeafNode) String() string { return proto.CompactTextString(m) }
func (*LeafNode) ProtoMessage()    {}
func (*LeafNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{0}
}

func (m *LeafNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeafNode.Unmarshal(m, b)
}
func (m *LeafNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeafNode.Marshal(b, m, deterministic)
}
func (m *LeafNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeafNode.Merge(m, src)
}
func (m *LeafNode) XXX_Size() int {
	return xxx_messageInfo_LeafNode.Size(m)
}
func (m *LeafNode) XXX_DiscardUnknown() {
	xxx_messageInfo_LeafNode.DiscardUnknown(m)
}

var xxx_messageInfo_LeafNode proto.InternalMessageInfo

func (m *LeafNode) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LeafNode) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *LeafNode) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *LeafNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type InnerNode struct {
	LeftHash             []byte   `protobuf:"bytes,1,opt,name=leftHash,proto3" json:"leftHash,omitempty"`
	RightHash            []byte   `protobuf:"bytes,2,opt,name=rightHash,proto3" json:"rightHash,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Size                 int32    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerNode) Reset()         { *m = InnerNode{} }
func (m *InnerNode) String() string { return proto.CompactTextString(m) }
func (*InnerNode) ProtoMessage()    {}
func (*InnerNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{1}
}

func (m *InnerNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerNode.Unmarshal(m, b)
}
func (m *InnerNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerNode.Marshal(b, m, deterministic)
}
func (m *InnerNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerNode.Merge(m, src)
}
func (m *InnerNode) XXX_Size() int {
	return xxx_messageInfo_InnerNode.Size(m)
}
func (m *InnerNode) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerNode.DiscardUnknown(m)
}

var xxx_messageInfo_InnerNode proto.InternalMessageInfo

func (m *InnerNode) GetLeftHash() []byte {
	if m != nil {
		return m.LeftHash
	}
	return nil
}

func (m *InnerNode) GetRightHash() []byte {
	if m != nil {
		return m.RightHash
	}
	return nil
}

func (m *InnerNode) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *InnerNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type MAVLProof struct {
	LeafHash             []byte       `protobuf:"bytes,1,opt,name=leafHash,proto3" json:"leafHash,omitempty"`
	InnerNodes           []*InnerNode `protobuf:"bytes,2,rep,name=innerNodes,proto3" json:"innerNodes,omitempty"`
	RootHash             []byte       `protobuf:"bytes,3,opt,name=rootHash,proto3" json:"rootHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MAVLProof) Reset()         { *m = MAVLProof{} }
func (m *MAVLProof) String() string { return proto.CompactTextString(m) }
func (*MAVLProof) ProtoMessage()    {}
func (*MAVLProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{2}
}

func (m *MAVLProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MAVLProof.Unmarshal(m, b)
}
func (m *MAVLProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MAVLProof.Marshal(b, m, deterministic)
}
func (m *MAVLProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MAVLProof.Merge(m, src)
}
func (m *MAVLProof) XXX_Size() int {
	return xxx_messageInfo_MAVLProof.Size(m)
}
func (m *MAVLProof) XXX_DiscardUnknown() {
	xxx_messageInfo_MAVLProof.DiscardUnknown(m)
}

var xxx_messageInfo_MAVLProof proto.InternalMessageInfo

func (m *MAVLProof) GetLeafHash() []byte {
	if m != nil {
		return m.LeafHash
	}
	return nil
}

func (m *MAVLProof) GetInnerNodes() []*InnerNode {
	if m != nil {
		return m.InnerNodes
	}
	return nil
}

func (m *MAVLProof) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

type StoreNode struct {
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> update chain33
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	LeftHash             []byte   `protobuf:"bytes,3,opt,name=leftHash,proto3" json:"leftHash,omitempty"`
	RightHash            []byte   `protobuf:"bytes,4,opt,name=rightHash,proto3" json:"rightHash,omitempty"`
	Height               int32    `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Size                 int32    `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
<<<<<<< HEAD
=======
	Key        []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value      []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	LeftHash   []byte `protobuf:"bytes,3,opt,name=leftHash,proto3" json:"leftHash,omitempty"`
	RightHash  []byte `protobuf:"bytes,4,opt,name=rightHash,proto3" json:"rightHash,omitempty"`
	Height     int32  `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	Size       int32  `protobuf:"varint,6,opt,name=size" json:"size,omitempty"`
	ParentHash []byte `protobuf:"bytes,7,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
>>>>>>> change protobuf file (use protoc-gen-1.0)
=======
}

func (m *StoreNode) Reset()         { *m = StoreNode{} }
func (m *StoreNode) String() string { return proto.CompactTextString(m) }
func (*StoreNode) ProtoMessage()    {}
func (*StoreNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{3}
>>>>>>> update chain33
}

func (m *StoreNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreNode.Unmarshal(m, b)
}
func (m *StoreNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreNode.Marshal(b, m, deterministic)
}
func (m *StoreNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreNode.Merge(m, src)
}
func (m *StoreNode) XXX_Size() int {
	return xxx_messageInfo_StoreNode.Size(m)
}
func (m *StoreNode) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreNode.DiscardUnknown(m)
}

var xxx_messageInfo_StoreNode proto.InternalMessageInfo

func (m *StoreNode) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *StoreNode) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *StoreNode) GetLeftHash() []byte {
	if m != nil {
		return m.LeftHash
	}
	return nil
}

func (m *StoreNode) GetRightHash() []byte {
	if m != nil {
		return m.RightHash
	}
	return nil
}

func (m *StoreNode) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *StoreNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type LocalDBSet struct {
<<<<<<< HEAD
	KV []*KeyValue `protobuf:"bytes,2,rep,name=KV" json:"KV,omitempty"`
=======
	KV                   []*KeyValue `protobuf:"bytes,2,rep,name=KV,proto3" json:"KV,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LocalDBSet) Reset()         { *m = LocalDBSet{} }
func (m *LocalDBSet) String() string { return proto.CompactTextString(m) }
func (*LocalDBSet) ProtoMessage()    {}
func (*LocalDBSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{4}
}

func (m *LocalDBSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalDBSet.Unmarshal(m, b)
}
func (m *LocalDBSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalDBSet.Marshal(b, m, deterministic)
}
func (m *LocalDBSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalDBSet.Merge(m, src)
}
func (m *LocalDBSet) XXX_Size() int {
	return xxx_messageInfo_LocalDBSet.Size(m)
}
func (m *LocalDBSet) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalDBSet.DiscardUnknown(m)
>>>>>>> update chain33
}

var xxx_messageInfo_LocalDBSet proto.InternalMessageInfo

func (m *LocalDBSet) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

type LocalDBList struct {
	Prefix               []byte   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Direction            int32    `protobuf:"varint,3,opt,name=direction,proto3" json:"direction,omitempty"`
	Count                int32    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalDBList) Reset()         { *m = LocalDBList{} }
func (m *LocalDBList) String() string { return proto.CompactTextString(m) }
func (*LocalDBList) ProtoMessage()    {}
func (*LocalDBList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{5}
}

func (m *LocalDBList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalDBList.Unmarshal(m, b)
}
func (m *LocalDBList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalDBList.Marshal(b, m, deterministic)
}
func (m *LocalDBList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalDBList.Merge(m, src)
}
func (m *LocalDBList) XXX_Size() int {
	return xxx_messageInfo_LocalDBList.Size(m)
}
func (m *LocalDBList) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalDBList.DiscardUnknown(m)
}

var xxx_messageInfo_LocalDBList proto.InternalMessageInfo

func (m *LocalDBList) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *LocalDBList) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LocalDBList) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *LocalDBList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type LocalDBGet struct {
	Keys                 [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalDBGet) Reset()         { *m = LocalDBGet{} }
func (m *LocalDBGet) String() string { return proto.CompactTextString(m) }
func (*LocalDBGet) ProtoMessage()    {}
func (*LocalDBGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{6}
}

func (m *LocalDBGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalDBGet.Unmarshal(m, b)
}
func (m *LocalDBGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalDBGet.Marshal(b, m, deterministic)
}
func (m *LocalDBGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalDBGet.Merge(m, src)
}
func (m *LocalDBGet) XXX_Size() int {
	return xxx_messageInfo_LocalDBGet.Size(m)
}
func (m *LocalDBGet) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalDBGet.DiscardUnknown(m)
}

var xxx_messageInfo_LocalDBGet proto.InternalMessageInfo

func (m *LocalDBGet) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

type LocalReplyValue struct {
	Values               [][]byte `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalReplyValue) Reset()         { *m = LocalReplyValue{} }
func (m *LocalReplyValue) String() string { return proto.CompactTextString(m) }
func (*LocalReplyValue) ProtoMessage()    {}
func (*LocalReplyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{7}
}

func (m *LocalReplyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalReplyValue.Unmarshal(m, b)
}
func (m *LocalReplyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalReplyValue.Marshal(b, m, deterministic)
}
func (m *LocalReplyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalReplyValue.Merge(m, src)
}
func (m *LocalReplyValue) XXX_Size() int {
	return xxx_messageInfo_LocalReplyValue.Size(m)
}
func (m *LocalReplyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalReplyValue.DiscardUnknown(m)
}

var xxx_messageInfo_LocalReplyValue proto.InternalMessageInfo

func (m *LocalReplyValue) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type StoreSet struct {
	StateHash            []byte      `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	KV                   []*KeyValue `protobuf:"bytes,2,rep,name=KV,proto3" json:"KV,omitempty"`
	Height               int64       `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StoreSet) Reset()         { *m = StoreSet{} }
func (m *StoreSet) String() string { return proto.CompactTextString(m) }
func (*StoreSet) ProtoMessage()    {}
func (*StoreSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{8}
}

func (m *StoreSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreSet.Unmarshal(m, b)
}
func (m *StoreSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreSet.Marshal(b, m, deterministic)
}
func (m *StoreSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreSet.Merge(m, src)
}
func (m *StoreSet) XXX_Size() int {
	return xxx_messageInfo_StoreSet.Size(m)
}
func (m *StoreSet) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreSet.DiscardUnknown(m)
}

var xxx_messageInfo_StoreSet proto.InternalMessageInfo

func (m *StoreSet) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *StoreSet) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

func (m *StoreSet) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type StoreDel struct {
	StateHash            []byte   `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Height               int64    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreDel) Reset()         { *m = StoreDel{} }
func (m *StoreDel) String() string { return proto.CompactTextString(m) }
func (*StoreDel) ProtoMessage()    {}
func (*StoreDel) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{9}
}

func (m *StoreDel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreDel.Unmarshal(m, b)
}
func (m *StoreDel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreDel.Marshal(b, m, deterministic)
}
func (m *StoreDel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreDel.Merge(m, src)
}
func (m *StoreDel) XXX_Size() int {
	return xxx_messageInfo_StoreDel.Size(m)
}
func (m *StoreDel) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreDel.DiscardUnknown(m)
}

var xxx_messageInfo_StoreDel proto.InternalMessageInfo

func (m *StoreDel) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *StoreDel) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type StoreSetWithSync struct {
	Storeset             *StoreSet `protobuf:"bytes,1,opt,name=storeset,proto3" json:"storeset,omitempty"`
	Sync                 bool      `protobuf:"varint,2,opt,name=sync,proto3" json:"sync,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StoreSetWithSync) Reset()         { *m = StoreSetWithSync{} }
func (m *StoreSetWithSync) String() string { return proto.CompactTextString(m) }
func (*StoreSetWithSync) ProtoMessage()    {}
func (*StoreSetWithSync) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{10}
}

func (m *StoreSetWithSync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreSetWithSync.Unmarshal(m, b)
}
func (m *StoreSetWithSync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreSetWithSync.Marshal(b, m, deterministic)
}
func (m *StoreSetWithSync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreSetWithSync.Merge(m, src)
}
func (m *StoreSetWithSync) XXX_Size() int {
	return xxx_messageInfo_StoreSetWithSync.Size(m)
}
func (m *StoreSetWithSync) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreSetWithSync.DiscardUnknown(m)
}

var xxx_messageInfo_StoreSetWithSync proto.InternalMessageInfo

func (m *StoreSetWithSync) GetStoreset() *StoreSet {
	if m != nil {
		return m.Storeset
	}
	return nil
}

func (m *StoreSetWithSync) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

type StoreGet struct {
	StateHash            []byte   `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Keys                 [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreGet) Reset()         { *m = StoreGet{} }
func (m *StoreGet) String() string { return proto.CompactTextString(m) }
func (*StoreGet) ProtoMessage()    {}
func (*StoreGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{11}
}

func (m *StoreGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreGet.Unmarshal(m, b)
}
func (m *StoreGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreGet.Marshal(b, m, deterministic)
}
func (m *StoreGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreGet.Merge(m, src)
}
func (m *StoreGet) XXX_Size() int {
	return xxx_messageInfo_StoreGet.Size(m)
}
func (m *StoreGet) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreGet.DiscardUnknown(m)
}

var xxx_messageInfo_StoreGet proto.InternalMessageInfo

func (m *StoreGet) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *StoreGet) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

type StoreReplyValue struct {
	Values               [][]byte `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreReplyValue) Reset()         { *m = StoreReplyValue{} }
func (m *StoreReplyValue) String() string { return proto.CompactTextString(m) }
func (*StoreReplyValue) ProtoMessage()    {}
func (*StoreReplyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{12}
}

func (m *StoreReplyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreReplyValue.Unmarshal(m, b)
}
func (m *StoreReplyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreReplyValue.Marshal(b, m, deterministic)
}
func (m *StoreReplyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreReplyValue.Merge(m, src)
}
func (m *StoreReplyValue) XXX_Size() int {
	return xxx_messageInfo_StoreReplyValue.Size(m)
}
func (m *StoreReplyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreReplyValue.DiscardUnknown(m)
}

var xxx_messageInfo_StoreReplyValue proto.InternalMessageInfo

func (m *StoreReplyValue) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type StoreList struct {
	StateHash            []byte   `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Start                []byte   `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  []byte   `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Suffix               []byte   `protobuf:"bytes,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Count                int64    `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Mode                 int64    `protobuf:"varint,6,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreList) Reset()         { *m = StoreList{} }
func (m *StoreList) String() string { return proto.CompactTextString(m) }
func (*StoreList) ProtoMessage()    {}
func (*StoreList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{13}
}

func (m *StoreList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreList.Unmarshal(m, b)
}
func (m *StoreList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreList.Marshal(b, m, deterministic)
}
func (m *StoreList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreList.Merge(m, src)
}
func (m *StoreList) XXX_Size() int {
	return xxx_messageInfo_StoreList.Size(m)
}
func (m *StoreList) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreList.DiscardUnknown(m)
}

var xxx_messageInfo_StoreList proto.InternalMessageInfo

func (m *StoreList) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *StoreList) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *StoreList) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *StoreList) GetSuffix() []byte {
	if m != nil {
		return m.Suffix
	}
	return nil
}

func (m *StoreList) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *StoreList) GetMode() int64 {
	if m != nil {
		return m.Mode
	}
	return 0
}

type StoreListReply struct {
	Start                []byte   `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  []byte   `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	Suffix               []byte   `protobuf:"bytes,3,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Count                int64    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Num                  int64    `protobuf:"varint,5,opt,name=num,proto3" json:"num,omitempty"`
	Mode                 int64    `protobuf:"varint,6,opt,name=mode,proto3" json:"mode,omitempty"`
	NextKey              []byte   `protobuf:"bytes,7,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
	Keys                 [][]byte `protobuf:"bytes,8,rep,name=keys,proto3" json:"keys,omitempty"`
	Values               [][]byte `protobuf:"bytes,9,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreListReply) Reset()         { *m = StoreListReply{} }
func (m *StoreListReply) String() string { return proto.CompactTextString(m) }
func (*StoreListReply) ProtoMessage()    {}
func (*StoreListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{14}
}

func (m *StoreListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreListReply.Unmarshal(m, b)
}
func (m *StoreListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreListReply.Marshal(b, m, deterministic)
}
func (m *StoreListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreListReply.Merge(m, src)
}
func (m *StoreListReply) XXX_Size() int {
	return xxx_messageInfo_StoreListReply.Size(m)
}
func (m *StoreListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreListReply.DiscardUnknown(m)
}

var xxx_messageInfo_StoreListReply proto.InternalMessageInfo

func (m *StoreListReply) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *StoreListReply) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *StoreListReply) GetSuffix() []byte {
	if m != nil {
		return m.Suffix
	}
	return nil
}

func (m *StoreListReply) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *StoreListReply) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *StoreListReply) GetMode() int64 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *StoreListReply) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

func (m *StoreListReply) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *StoreListReply) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type PruneData struct {
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> update chain33
	// 该叶子节点的所有父hash
	Hashs                [][]byte `protobuf:"bytes,1,rep,name=hashs,proto3" json:"hashs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
<<<<<<< HEAD
}

func (m *PruneData) Reset()         { *m = PruneData{} }
func (m *PruneData) String() string { return proto.CompactTextString(m) }
func (*PruneData) ProtoMessage()    {}
func (*PruneData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{15}
}

func (m *PruneData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PruneData.Unmarshal(m, b)
}
func (m *PruneData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PruneData.Marshal(b, m, deterministic)
}
func (m *PruneData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PruneData.Merge(m, src)
}
func (m *PruneData) XXX_Size() int {
	return xxx_messageInfo_PruneData.Size(m)
}
func (m *PruneData) XXX_DiscardUnknown() {
	xxx_messageInfo_PruneData.DiscardUnknown(m)
=======
	// 对应keyHash下的区块高度
	Height int64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	// hash+prefix的长度
	Lenth int32 `protobuf:"varint,2,opt,name=lenth" json:"lenth,omitempty"`
>>>>>>> change protobuf file (use protoc-gen-1.0)
=======
>>>>>>> update chain33
}

func (m *PruneData) Reset()         { *m = PruneData{} }
func (m *PruneData) String() string { return proto.CompactTextString(m) }
func (*PruneData) ProtoMessage()    {}
func (*PruneData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{15}
}

<<<<<<< HEAD
=======
func (m *PruneData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PruneData.Unmarshal(m, b)
}
func (m *PruneData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PruneData.Marshal(b, m, deterministic)
}
func (m *PruneData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PruneData.Merge(m, src)
}
func (m *PruneData) XXX_Size() int {
	return xxx_messageInfo_PruneData.Size(m)
}
func (m *PruneData) XXX_DiscardUnknown() {
	xxx_messageInfo_PruneData.DiscardUnknown(m)
}

var xxx_messageInfo_PruneData proto.InternalMessageInfo

>>>>>>> update chain33
func (m *PruneData) GetHashs() [][]byte {
	if m != nil {
		return m.Hashs
	}
	return nil
}

//用于存储db Pool数据的Value
type StoreValuePool struct {
	Values               [][]byte `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreValuePool) Reset()         { *m = StoreValuePool{} }
func (m *StoreValuePool) String() string { return proto.CompactTextString(m) }
func (*StoreValuePool) ProtoMessage()    {}
func (*StoreValuePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{16}
}

func (m *StoreValuePool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreValuePool.Unmarshal(m, b)
}
func (m *StoreValuePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreValuePool.Marshal(b, m, deterministic)
}
func (m *StoreValuePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreValuePool.Merge(m, src)
}
func (m *StoreValuePool) XXX_Size() int {
	return xxx_messageInfo_StoreValuePool.Size(m)
}
func (m *StoreValuePool) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreValuePool.DiscardUnknown(m)
}

var xxx_messageInfo_StoreValuePool proto.InternalMessageInfo

func (m *StoreValuePool) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*LeafNode)(nil), "types.LeafNode")
	proto.RegisterType((*InnerNode)(nil), "types.InnerNode")
	proto.RegisterType((*MAVLProof)(nil), "types.MAVLProof")
	proto.RegisterType((*StoreNode)(nil), "types.StoreNode")
	proto.RegisterType((*LocalDBSet)(nil), "types.LocalDBSet")
	proto.RegisterType((*LocalDBList)(nil), "types.LocalDBList")
	proto.RegisterType((*LocalDBGet)(nil), "types.LocalDBGet")
	proto.RegisterType((*LocalReplyValue)(nil), "types.LocalReplyValue")
	proto.RegisterType((*StoreSet)(nil), "types.StoreSet")
	proto.RegisterType((*StoreDel)(nil), "types.StoreDel")
	proto.RegisterType((*StoreSetWithSync)(nil), "types.StoreSetWithSync")
	proto.RegisterType((*StoreGet)(nil), "types.StoreGet")
	proto.RegisterType((*StoreReplyValue)(nil), "types.StoreReplyValue")
	proto.RegisterType((*StoreList)(nil), "types.StoreList")
	proto.RegisterType((*StoreListReply)(nil), "types.StoreListReply")
	proto.RegisterType((*PruneData)(nil), "types.PruneData")
	proto.RegisterType((*StoreValuePool)(nil), "types.StoreValuePool")
}

<<<<<<< HEAD
func init() { proto.RegisterFile("db.proto", fileDescriptor3) }

<<<<<<< HEAD
=======
func init() { proto.RegisterFile("db.proto", fileDescriptor_8817812184a13374) }

>>>>>>> update chain33
var fileDescriptor_8817812184a13374 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x45, 0x92, 0x9d, 0x48, 0x93, 0xd0, 0x18, 0x11, 0x8a, 0x08, 0x29, 0x71, 0x75, 0x72, 0x29,
	0x75, 0x4a, 0x7d, 0xed, 0xa1, 0x0d, 0x81, 0xb4, 0xd8, 0x2d, 0x41, 0x06, 0x17, 0x7a, 0x28, 0x28,
	0xf2, 0x38, 0x12, 0xb1, 0x77, 0x5d, 0xed, 0xaa, 0x44, 0xfd, 0x1b, 0x3d, 0xf5, 0x6f, 0xf5, 0x17,
	0x95, 0x9d, 0x5d, 0x7d, 0xa4, 0x38, 0x71, 0x73, 0xdb, 0x37, 0x1e, 0xcf, 0x7b, 0xf3, 0xe6, 0x81,
	0xc0, 0x9d, 0x5f, 0x0d, 0xd7, 0x39, 0x97, 0xdc, 0xef, 0xca, 0x72, 0x8d, 0xe2, 0x68, 0x3f, 0xe1,
	0xab, 0x15, 0x67, 0xba, 0x18, 0x7e, 0x03, 0x77, 0x82, 0xf1, 0xe2, 0x33, 0x9f, 0xa3, 0xdf, 0x03,
	0xe7, 0x06, 0xcb, 0xc0, 0xea, 0x5b, 0x83, 0xfd, 0x48, 0x3d, 0xfd, 0x43, 0xe8, 0xfe, 0x88, 0x97,
	0x05, 0x06, 0x36, 0xd5, 0x34, 0xf0, 0x9f, 0xc2, 0x4e, 0x8a, 0xd9, 0x75, 0x2a, 0x03, 0xa7, 0x6f,
	0x0d, 0xba, 0x91, 0x41, 0xbe, 0x0f, 0x1d, 0x91, 0xfd, 0xc4, 0xa0, 0x43, 0x55, 0x7a, 0x87, 0xdf,
	0xc1, 0xfb, 0xc8, 0x18, 0xe6, 0x44, 0x70, 0x04, 0xee, 0x12, 0x17, 0xf2, 0x43, 0x2c, 0x52, 0xc3,
	0x52, 0x63, 0xff, 0x18, 0xbc, 0x5c, 0x4d, 0xa1, 0x1f, 0x35, 0x5d, 0x53, 0x78, 0x14, 0x65, 0x01,
	0xde, 0xa7, 0xf7, 0xb3, 0xc9, 0x65, 0xce, 0xf9, 0x42, 0x53, 0xc6, 0x8b, 0xbb, 0x94, 0x1a, 0xfb,
	0xaf, 0x01, 0xb2, 0x4a, 0x9b, 0x08, 0xec, 0xbe, 0x33, 0xd8, 0x7b, 0xd3, 0x1b, 0x92, 0x4b, 0xc3,
	0x5a, 0x74, 0xd4, 0xea, 0x51, 0xd3, 0x72, 0xce, 0xb5, 0x46, 0x47, 0x4f, 0xab, 0x70, 0xf8, 0xdb,
	0x02, 0x6f, 0x2a, 0x79, 0x8e, 0x8f, 0xf2, 0xb2, 0x6d, 0x89, 0xf3, 0x90, 0x25, 0x9d, 0xfb, 0x2d,
	0xe9, 0x6e, 0xb4, 0x64, 0xa7, 0x65, 0xc9, 0x2b, 0x80, 0x09, 0x4f, 0xe2, 0xe5, 0xf9, 0xd9, 0x14,
	0xa5, 0x7f, 0x02, 0xf6, 0x78, 0x66, 0xf6, 0x3d, 0x30, 0xfb, 0x8e, 0xb1, 0x9c, 0x29, 0x41, 0x91,
	0x3d, 0x9e, 0x85, 0x37, 0xb0, 0x67, 0xda, 0x27, 0x99, 0x90, 0x8a, 0x69, 0x9d, 0xe3, 0x22, 0xbb,
	0x35, 0xeb, 0x18, 0x54, 0xed, 0x68, 0x37, 0x3b, 0x1e, 0x83, 0x37, 0xcf, 0x72, 0x4c, 0x64, 0xc6,
	0x99, 0xb9, 0x54, 0x53, 0x50, 0x0e, 0x24, 0xbc, 0x60, 0xd2, 0x5c, 0x4b, 0x83, 0xb0, 0x5f, 0x6b,
	0xbb, 0x40, 0x52, 0x7f, 0x83, 0xa5, 0xbe, 0xc6, 0x7e, 0x44, 0xef, 0xf0, 0x05, 0x1c, 0x50, 0x47,
	0x84, 0xeb, 0xa5, 0x56, 0xa9, 0x24, 0x91, 0x7f, 0x55, 0xa3, 0x41, 0x61, 0x0c, 0x2e, 0xdd, 0x40,
	0xad, 0x79, 0x0c, 0x9e, 0x90, 0xb1, 0xc4, 0xd6, 0xed, 0x9b, 0xc2, 0x56, 0x13, 0xfe, 0x89, 0x9c,
	0x53, 0xf9, 0x1b, 0xbe, 0x33, 0x14, 0xe7, 0xb8, 0xdc, 0x42, 0xd1, 0x4c, 0xb0, 0xef, 0x4c, 0x98,
	0x42, 0xaf, 0x12, 0xf9, 0x25, 0x93, 0xe9, 0xb4, 0x64, 0x89, 0xff, 0x12, 0x5c, 0xa1, 0x6a, 0x02,
	0x25, 0x0d, 0x6a, 0x44, 0x55, 0xad, 0x51, 0xdd, 0x40, 0x27, 0x2e, 0x59, 0x42, 0x63, 0xdd, 0x88,
	0xde, 0xe1, 0x5b, 0x23, 0xeb, 0x62, 0xeb, 0xe6, 0xf7, 0x58, 0x4c, 0xff, 0xfe, 0x0f, 0x8b, 0x7f,
	0x55, 0x39, 0xa7, 0x6c, 0x3c, 0x4c, 0x75, 0x08, 0x5d, 0x21, 0xe3, 0x5c, 0x56, 0x99, 0x27, 0xa0,
	0x72, 0x83, 0x6c, 0x6e, 0xe2, 0xae, 0x9e, 0x8a, 0x4b, 0x14, 0x0b, 0x95, 0x30, 0x1d, 0x73, 0x83,
	0x9a, 0xc4, 0x74, 0xc9, 0x40, 0x0d, 0xd4, 0x02, 0x2b, 0x3e, 0xd7, 0x09, 0x77, 0x22, 0x7a, 0x87,
	0x7f, 0x2c, 0x78, 0x52, 0xab, 0xa2, 0x2d, 0x1a, 0x72, 0x6b, 0x03, 0xb9, 0xbd, 0x89, 0xdc, 0xd9,
	0x4c, 0xde, 0x69, 0x93, 0xf7, 0xc0, 0x61, 0xc5, 0xca, 0x08, 0x52, 0xcf, 0x4d, 0x72, 0xfc, 0x00,
	0x76, 0x19, 0xde, 0xca, 0x31, 0x96, 0xc1, 0x2e, 0x0d, 0xad, 0x60, 0xed, 0xbe, 0xdb, 0xb8, 0xdf,
	0xb2, 0xda, 0xbb, 0x63, 0xf5, 0x73, 0xf0, 0x2e, 0xf3, 0x82, 0xe1, 0x79, 0x2c, 0x63, 0x25, 0x27,
	0x8d, 0x45, 0x2a, 0x02, 0x8b, 0x7a, 0x34, 0x08, 0x07, 0x66, 0x6d, 0xba, 0xd9, 0x25, 0xe7, 0xcb,
	0xd6, 0x30, 0xab, 0x3d, 0xec, 0xec, 0xe4, 0xeb, 0xb3, 0xeb, 0x4c, 0xa6, 0xc5, 0xd5, 0x30, 0xe1,
	0xab, 0xd3, 0xd1, 0x28, 0x61, 0xa7, 0x49, 0x1a, 0x67, 0x6c, 0x34, 0x3a, 0xa5, 0xa0, 0x5d, 0xed,
	0xd0, 0x17, 0x61, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x81, 0x98, 0x2a, 0x32, 0x06, 0x00,
	0x00,
<<<<<<< HEAD
=======
var fileDescriptor3 = []byte{
	// 647 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0x45, 0x92, 0x95, 0x48, 0x93, 0xf0, 0xc5, 0x88, 0xf0, 0x61, 0x42, 0xda, 0x18, 0x5d, 0xb9,
	0x94, 0x3a, 0xa5, 0xbe, 0x2a, 0xf4, 0xa2, 0x0d, 0x81, 0xb4, 0xd8, 0x2d, 0x41, 0x06, 0x17, 0x7a,
	0x51, 0xd8, 0xc8, 0xe3, 0x58, 0x44, 0xde, 0x75, 0xa5, 0x55, 0x89, 0xfa, 0x1a, 0x7d, 0x9a, 0xbe,
	0x46, 0x9f, 0xa8, 0xec, 0xec, 0xea, 0x27, 0xc5, 0x89, 0x9b, 0xbb, 0x39, 0xe3, 0xf1, 0x9c, 0x33,
	0x67, 0x67, 0x6c, 0xf0, 0xe6, 0x57, 0xc3, 0x75, 0x26, 0xa4, 0x08, 0x5c, 0x59, 0xae, 0x31, 0x3f,
	0xda, 0x8f, 0xc5, 0x6a, 0x25, 0xb8, 0x4e, 0x86, 0x5f, 0xc1, 0x9b, 0x20, 0x5b, 0x7c, 0x12, 0x73,
	0x0c, 0xba, 0xe0, 0xdc, 0x60, 0xd9, 0xb3, 0xfa, 0xd6, 0x60, 0x3f, 0x52, 0x61, 0x70, 0x08, 0xee,
	0x77, 0x96, 0x16, 0xd8, 0xb3, 0x29, 0xa7, 0x41, 0xf0, 0x3f, 0xec, 0x2c, 0x31, 0xb9, 0x5e, 0xca,
	0x9e, 0xd3, 0xb7, 0x06, 0x6e, 0x64, 0x50, 0x10, 0x40, 0x27, 0x4f, 0x7e, 0x60, 0xaf, 0x43, 0x59,
	0x8a, 0xc3, 0x6f, 0xe0, 0x7f, 0xe0, 0x1c, 0x33, 0x22, 0x38, 0x02, 0x2f, 0xc5, 0x85, 0x7c, 0xcf,
	0xf2, 0xa5, 0x61, 0xa9, 0x71, 0x70, 0x0c, 0x7e, 0xa6, 0xba, 0xd0, 0x87, 0x9a, 0xae, 0x49, 0x3c,
	0x8a, 0xb2, 0x00, 0xff, 0xe3, 0xbb, 0xd9, 0xe4, 0x32, 0x13, 0x62, 0xa1, 0x29, 0xd9, 0xe2, 0x2e,
	0xa5, 0xc6, 0xc1, 0x4b, 0x80, 0xa4, 0xd2, 0x96, 0xf7, 0xec, 0xbe, 0x33, 0xd8, 0x7b, 0xd5, 0x1d,
	0x92, 0x4b, 0xc3, 0x5a, 0x74, 0xd4, 0xaa, 0x51, 0xdd, 0x32, 0x21, 0xb4, 0x46, 0x47, 0x77, 0xab,
	0x70, 0xf8, 0xcb, 0x02, 0x7f, 0x2a, 0x45, 0x86, 0x8f, 0xf2, 0xb2, 0x6d, 0x89, 0xf3, 0x90, 0x25,
	0x9d, 0xfb, 0x2d, 0x71, 0x37, 0x5a, 0xb2, 0xd3, 0x58, 0x12, 0x3c, 0x05, 0x58, 0xb3, 0x0c, 0xb9,
	0x6e, 0xb5, 0x4b, 0xad, 0x5a, 0x99, 0xf0, 0x05, 0xc0, 0x44, 0xc4, 0x2c, 0x3d, 0x3f, 0x9b, 0xa2,
	0x0c, 0x4e, 0xc0, 0x1e, 0xcf, 0x8c, 0x1f, 0x07, 0xc6, 0x8f, 0x31, 0x96, 0x33, 0x25, 0x38, 0xb2,
	0xc7, 0xb3, 0xf0, 0x06, 0xf6, 0x4c, 0xf9, 0x24, 0xc9, 0xa5, 0x52, 0xb2, 0xce, 0x70, 0x91, 0xdc,
	0x9a, 0x71, 0x0d, 0xaa, 0x3c, 0xb0, 0x1b, 0x0f, 0x8e, 0xc1, 0x9f, 0x27, 0x19, 0xc6, 0x32, 0x11,
	0xdc, 0xbc, 0x64, 0x93, 0x50, 0x0e, 0xc5, 0xa2, 0xe0, 0xd2, 0xbc, 0xa6, 0x06, 0x61, 0xbf, 0xd6,
	0x76, 0x81, 0x34, 0xdd, 0x0d, 0x96, 0xfa, 0xb5, 0xf6, 0x23, 0x8a, 0xc3, 0x67, 0x70, 0x40, 0x15,
	0x11, 0xae, 0x53, 0xad, 0x52, 0x49, 0x22, 0x7f, 0xab, 0x42, 0x83, 0x42, 0x06, 0x1e, 0xbd, 0x91,
	0x1a, 0xf3, 0x18, 0xfc, 0x5c, 0x32, 0x89, 0xad, 0xdd, 0x68, 0x12, 0x5b, 0x4d, 0xf8, 0x6b, 0x25,
	0x9d, 0xca, 0xff, 0xf0, 0xad, 0xa1, 0x38, 0xc7, 0x74, 0x0b, 0x45, 0xd3, 0xc1, 0xbe, 0xd3, 0x61,
	0x0a, 0xdd, 0x4a, 0xe4, 0xe7, 0x44, 0x2e, 0xa7, 0x25, 0x8f, 0x83, 0xe7, 0xe0, 0xe5, 0x2a, 0x97,
	0xa3, 0xa4, 0x46, 0x8d, 0xa8, 0xaa, 0x34, 0xaa, 0x0b, 0x68, 0x05, 0x4a, 0x1e, 0x53, 0x5b, 0x2f,
	0xa2, 0x38, 0x7c, 0x63, 0x64, 0x5d, 0x6c, 0x9d, 0xfc, 0x1e, 0x8b, 0xe9, 0xdb, 0xff, 0x60, 0xf1,
	0xcf, 0xea, 0x0e, 0x68, 0x37, 0x1e, 0xa6, 0x3a, 0x04, 0x37, 0x97, 0x2c, 0x93, 0xd5, 0x4d, 0x10,
	0x50, 0x7b, 0x83, 0x7c, 0x6e, 0xce, 0x41, 0x85, 0x8a, 0x2b, 0x2f, 0x16, 0x6a, 0xc3, 0xf4, 0x19,
	0x18, 0xd4, 0x6c, 0x8c, 0x4b, 0x06, 0x6a, 0xa0, 0x06, 0x58, 0x89, 0xb9, 0xbe, 0x00, 0x27, 0xa2,
	0x38, 0xfc, 0x6d, 0xc1, 0x7f, 0xb5, 0x2a, 0x9a, 0xa2, 0x21, 0xb7, 0x36, 0x90, 0xdb, 0x9b, 0xc8,
	0x9d, 0xcd, 0xe4, 0x9d, 0x36, 0x79, 0x17, 0x1c, 0x5e, 0xac, 0x8c, 0x20, 0x15, 0x6e, 0x92, 0x13,
	0xf4, 0x60, 0x97, 0xe3, 0xad, 0x1c, 0x63, 0x69, 0xae, 0xb1, 0x82, 0xb5, 0xfb, 0x5e, 0xe3, 0x7e,
	0xcb, 0x6a, 0xff, 0x8e, 0xd5, 0xaf, 0xc1, 0xbf, 0xcc, 0x0a, 0x8e, 0xe7, 0x4c, 0xb2, 0xd6, 0x36,
	0x59, 0xed, 0x6d, 0x52, 0x32, 0x53, 0xe4, 0x52, 0xff, 0xa8, 0xba, 0x91, 0x06, 0xe1, 0xc0, 0xd8,
	0x41, 0x6f, 0x79, 0x29, 0x44, 0xda, 0x22, 0xb1, 0xda, 0x24, 0x67, 0x27, 0x5f, 0x9e, 0x5c, 0x27,
	0x72, 0x59, 0x5c, 0x0d, 0x63, 0xb1, 0x3a, 0x1d, 0x8d, 0x62, 0x7e, 0x1a, 0x2f, 0x59, 0xc2, 0x47,
	0xa3, 0x53, 0x5a, 0xc0, 0xab, 0x1d, 0xfa, 0x27, 0x19, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x30,
	0x37, 0x20, 0x41, 0x6a, 0x06, 0x00, 0x00,
>>>>>>> change protobuf file (use protoc-gen-1.0)
=======
>>>>>>> update chain33
}
