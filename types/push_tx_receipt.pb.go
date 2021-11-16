// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: push_tx_receipt.proto

package types

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

type TxReceipts4SubscribePerBlk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx          []*Transaction `protobuf:"bytes,1,rep,name=tx,proto3" json:"tx,omitempty"`
	ReceiptData []*ReceiptData `protobuf:"bytes,2,rep,name=receiptData,proto3" json:"receiptData,omitempty"`
	// repeated KeyValue    KV          = 3;
	Height       int64  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	BlockHash    []byte `protobuf:"bytes,5,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	ParentHash   []byte `protobuf:"bytes,6,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	PreviousHash []byte `protobuf:"bytes,7,opt,name=previousHash,proto3" json:"previousHash,omitempty"`
	AddDelType   int32  `protobuf:"varint,8,opt,name=addDelType,proto3" json:"addDelType,omitempty"`
	SeqNum       int64  `protobuf:"varint,9,opt,name=seqNum,proto3" json:"seqNum,omitempty"`
}

func (x *TxReceipts4SubscribePerBlk) Reset() {
	*x = TxReceipts4SubscribePerBlk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_tx_receipt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxReceipts4SubscribePerBlk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxReceipts4SubscribePerBlk) ProtoMessage() {}

func (x *TxReceipts4SubscribePerBlk) ProtoReflect() protoreflect.Message {
	mi := &file_push_tx_receipt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxReceipts4SubscribePerBlk.ProtoReflect.Descriptor instead.
func (*TxReceipts4SubscribePerBlk) Descriptor() ([]byte, []int) {
	return file_push_tx_receipt_proto_rawDescGZIP(), []int{0}
}

func (x *TxReceipts4SubscribePerBlk) GetTx() []*Transaction {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *TxReceipts4SubscribePerBlk) GetReceiptData() []*ReceiptData {
	if x != nil {
		return x.ReceiptData
	}
	return nil
}

func (x *TxReceipts4SubscribePerBlk) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TxReceipts4SubscribePerBlk) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *TxReceipts4SubscribePerBlk) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *TxReceipts4SubscribePerBlk) GetPreviousHash() []byte {
	if x != nil {
		return x.PreviousHash
	}
	return nil
}

func (x *TxReceipts4SubscribePerBlk) GetAddDelType() int32 {
	if x != nil {
		return x.AddDelType
	}
	return 0
}

func (x *TxReceipts4SubscribePerBlk) GetSeqNum() int64 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

type TxReceipts4Subscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxReceipts []*TxReceipts4SubscribePerBlk `protobuf:"bytes,1,rep,name=txReceipts,proto3" json:"txReceipts,omitempty"`
}

func (x *TxReceipts4Subscribe) Reset() {
	*x = TxReceipts4Subscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_tx_receipt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxReceipts4Subscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxReceipts4Subscribe) ProtoMessage() {}

func (x *TxReceipts4Subscribe) ProtoReflect() protoreflect.Message {
	mi := &file_push_tx_receipt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxReceipts4Subscribe.ProtoReflect.Descriptor instead.
func (*TxReceipts4Subscribe) Descriptor() ([]byte, []int) {
	return file_push_tx_receipt_proto_rawDescGZIP(), []int{1}
}

func (x *TxReceipts4Subscribe) GetTxReceipts() []*TxReceipts4SubscribePerBlk {
	if x != nil {
		return x.TxReceipts
	}
	return nil
}

type TxHashWithReceiptType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Ty   int32  `protobuf:"varint,2,opt,name=ty,proto3" json:"ty,omitempty"`
}

func (x *TxHashWithReceiptType) Reset() {
	*x = TxHashWithReceiptType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_tx_receipt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxHashWithReceiptType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxHashWithReceiptType) ProtoMessage() {}

func (x *TxHashWithReceiptType) ProtoReflect() protoreflect.Message {
	mi := &file_push_tx_receipt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxHashWithReceiptType.ProtoReflect.Descriptor instead.
func (*TxHashWithReceiptType) Descriptor() ([]byte, []int) {
	return file_push_tx_receipt_proto_rawDescGZIP(), []int{2}
}

func (x *TxHashWithReceiptType) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *TxHashWithReceiptType) GetTy() int32 {
	if x != nil {
		return x.Ty
	}
	return 0
}

type TxResultPerBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*TxHashWithReceiptType `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Height     int64                    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	BlockHash  []byte                   `protobuf:"bytes,3,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	ParentHash []byte                   `protobuf:"bytes,4,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	AddDelType int32                    `protobuf:"varint,5,opt,name=addDelType,proto3" json:"addDelType,omitempty"`
	SeqNum     int64                    `protobuf:"varint,6,opt,name=seqNum,proto3" json:"seqNum,omitempty"`
}

func (x *TxResultPerBlock) Reset() {
	*x = TxResultPerBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_tx_receipt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxResultPerBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxResultPerBlock) ProtoMessage() {}

func (x *TxResultPerBlock) ProtoReflect() protoreflect.Message {
	mi := &file_push_tx_receipt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxResultPerBlock.ProtoReflect.Descriptor instead.
func (*TxResultPerBlock) Descriptor() ([]byte, []int) {
	return file_push_tx_receipt_proto_rawDescGZIP(), []int{3}
}

func (x *TxResultPerBlock) GetItems() []*TxHashWithReceiptType {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TxResultPerBlock) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TxResultPerBlock) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *TxResultPerBlock) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *TxResultPerBlock) GetAddDelType() int32 {
	if x != nil {
		return x.AddDelType
	}
	return 0
}

func (x *TxResultPerBlock) GetSeqNum() int64 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

type TxResultSeqs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*TxResultPerBlock `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TxResultSeqs) Reset() {
	*x = TxResultSeqs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_tx_receipt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxResultSeqs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxResultSeqs) ProtoMessage() {}

func (x *TxResultSeqs) ProtoReflect() protoreflect.Message {
	mi := &file_push_tx_receipt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxResultSeqs.ProtoReflect.Descriptor instead.
func (*TxResultSeqs) Descriptor() ([]byte, []int) {
	return file_push_tx_receipt_proto_rawDescGZIP(), []int{4}
}

func (x *TxResultSeqs) GetItems() []*TxResultPerBlock {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_push_tx_receipt_proto protoreflect.FileDescriptor

var file_push_tx_receipt_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x78, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x11,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x1a, 0x54, 0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73,
	0x34, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x42, 0x6c, 0x6b,
	0x12, 0x22, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x02, 0x74, 0x78, 0x12, 0x34, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x44, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x44, 0x65, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x22, 0x59, 0x0a, 0x14,
	0x54, 0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x34, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x74, 0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x54, 0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x34, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x42, 0x6c, 0x6b, 0x52, 0x0a, 0x74, 0x78, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x22, 0x3b, 0x0a, 0x15, 0x54, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x74, 0x79, 0x22, 0xd4, 0x01, 0x0a, 0x10, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x50, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x44, 0x65, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x44, 0x65, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x22, 0x3d, 0x0a, 0x0c, 0x54,
	0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x71, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x65, 0x72, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x33, 0x33, 0x63, 0x6e, 0x2f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x33, 0x33, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_push_tx_receipt_proto_rawDescOnce sync.Once
	file_push_tx_receipt_proto_rawDescData = file_push_tx_receipt_proto_rawDesc
)

func file_push_tx_receipt_proto_rawDescGZIP() []byte {
	file_push_tx_receipt_proto_rawDescOnce.Do(func() {
		file_push_tx_receipt_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_tx_receipt_proto_rawDescData)
	})
	return file_push_tx_receipt_proto_rawDescData
}

var file_push_tx_receipt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_push_tx_receipt_proto_goTypes = []interface{}{
	(*TxReceipts4SubscribePerBlk)(nil), // 0: types.TxReceipts4SubscribePerBlk
	(*TxReceipts4Subscribe)(nil),       // 1: types.TxReceipts4Subscribe
	(*TxHashWithReceiptType)(nil),      // 2: types.TxHashWithReceiptType
	(*TxResultPerBlock)(nil),           // 3: types.TxResultPerBlock
	(*TxResultSeqs)(nil),               // 4: types.TxResultSeqs
	(*Transaction)(nil),                // 5: types.Transaction
	(*ReceiptData)(nil),                // 6: types.ReceiptData
}
var file_push_tx_receipt_proto_depIdxs = []int32{
	5, // 0: types.TxReceipts4SubscribePerBlk.tx:type_name -> types.Transaction
	6, // 1: types.TxReceipts4SubscribePerBlk.receiptData:type_name -> types.ReceiptData
	0, // 2: types.TxReceipts4Subscribe.txReceipts:type_name -> types.TxReceipts4SubscribePerBlk
	2, // 3: types.TxResultPerBlock.items:type_name -> types.TxHashWithReceiptType
	3, // 4: types.TxResultSeqs.items:type_name -> types.TxResultPerBlock
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_push_tx_receipt_proto_init() }
func file_push_tx_receipt_proto_init() {
	if File_push_tx_receipt_proto != nil {
		return
	}
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_push_tx_receipt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxReceipts4SubscribePerBlk); i {
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
		file_push_tx_receipt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxReceipts4Subscribe); i {
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
		file_push_tx_receipt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxHashWithReceiptType); i {
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
		file_push_tx_receipt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxResultPerBlock); i {
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
		file_push_tx_receipt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxResultSeqs); i {
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
			RawDescriptor: file_push_tx_receipt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_push_tx_receipt_proto_goTypes,
		DependencyIndexes: file_push_tx_receipt_proto_depIdxs,
		MessageInfos:      file_push_tx_receipt_proto_msgTypes,
	}.Build()
	File_push_tx_receipt_proto = out.File
	file_push_tx_receipt_proto_rawDesc = nil
	file_push_tx_receipt_proto_goTypes = nil
	file_push_tx_receipt_proto_depIdxs = nil
}
