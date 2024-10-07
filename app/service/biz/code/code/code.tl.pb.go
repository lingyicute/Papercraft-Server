//
// WARNING! All changes made in this file will be lost!
// Created from 'scheme.tl' by 'mtprotoc'
//
// Copyright (c) 2024-present,  Yomi.
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: code.tl.proto

package code

import (
	mtproto "github.com/teamgram/proto/mtproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TLConstructor int32

const (
	TLConstructor_CRC32_UNKNOWN                  TLConstructor = 0
	TLConstructor_CRC32_phoneCodeTransaction     TLConstructor = -2089576808
	TLConstructor_CRC32_code_createPhoneCode     TLConstructor = 1612963998
	TLConstructor_CRC32_code_getPhoneCode        TLConstructor = 1638179065
	TLConstructor_CRC32_code_deletePhoneCode     TLConstructor = -1498387888
	TLConstructor_CRC32_code_updatePhoneCodeData TLConstructor = -1231746411
)

// Enum value maps for TLConstructor.
var (
	TLConstructor_name = map[int32]string{
		0:           "CRC32_UNKNOWN",
		-2089576808: "CRC32_phoneCodeTransaction",
		1612963998:  "CRC32_code_createPhoneCode",
		1638179065:  "CRC32_code_getPhoneCode",
		-1498387888: "CRC32_code_deletePhoneCode",
		-1231746411: "CRC32_code_updatePhoneCodeData",
	}
	TLConstructor_value = map[string]int32{
		"CRC32_UNKNOWN":                  0,
		"CRC32_phoneCodeTransaction":     -2089576808,
		"CRC32_code_createPhoneCode":     1612963998,
		"CRC32_code_getPhoneCode":        1638179065,
		"CRC32_code_deletePhoneCode":     -1498387888,
		"CRC32_code_updatePhoneCodeData": -1231746411,
	}
)

func (x TLConstructor) Enum() *TLConstructor {
	p := new(TLConstructor)
	*p = x
	return p
}

func (x TLConstructor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TLConstructor) Descriptor() protoreflect.EnumDescriptor {
	return file_code_tl_proto_enumTypes[0].Descriptor()
}

func (TLConstructor) Type() protoreflect.EnumType {
	return &file_code_tl_proto_enumTypes[0]
}

func (x TLConstructor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TLConstructor.Descriptor instead.
func (TLConstructor) EnumDescriptor() ([]byte, []int) {
	return file_code_tl_proto_rawDescGZIP(), []int{0}
}

// PhoneCodeTransaction <--
//   - TL_phoneCodeTransaction
type PhoneCodeTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PredicateName         string        `protobuf:"bytes,1,opt,name=predicate_name,json=predicateName,proto3" json:"predicate_name,omitempty"`
	Constructor           TLConstructor `protobuf:"varint,2,opt,name=constructor,proto3,enum=code.TLConstructor" json:"constructor,omitempty"`
	AuthKeyId             int64         `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	SessionId             int64         `protobuf:"varint,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Phone                 string        `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	PhoneNumberRegistered bool          `protobuf:"varint,6,opt,name=phone_number_registered,json=phoneNumberRegistered,proto3" json:"phone_number_registered,omitempty"`
	PhoneCode             string        `protobuf:"bytes,7,opt,name=phone_code,json=phoneCode,proto3" json:"phone_code,omitempty"`
	PhoneCodeHash         string        `protobuf:"bytes,8,opt,name=phone_code_hash,json=phoneCodeHash,proto3" json:"phone_code_hash,omitempty"`
	PhoneCodeExpired      int32         `protobuf:"varint,9,opt,name=phone_code_expired,json=phoneCodeExpired,proto3" json:"phone_code_expired,omitempty"`
	PhoneCodeExtraData    string        `protobuf:"bytes,10,opt,name=phone_code_extra_data,json=phoneCodeExtraData,proto3" json:"phone_code_extra_data,omitempty"`
	SentCodeType          int32         `protobuf:"varint,11,opt,name=sent_code_type,json=sentCodeType,proto3" json:"sent_code_type,omitempty"`
	FlashCallPattern      string        `protobuf:"bytes,12,opt,name=flash_call_pattern,json=flashCallPattern,proto3" json:"flash_call_pattern,omitempty"`
	NextCodeType          int32         `protobuf:"varint,13,opt,name=next_code_type,json=nextCodeType,proto3" json:"next_code_type,omitempty"`
	State                 int32         `protobuf:"varint,14,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *PhoneCodeTransaction) Reset() {
	*x = PhoneCodeTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_tl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneCodeTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneCodeTransaction) ProtoMessage() {}

func (x *PhoneCodeTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_code_tl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneCodeTransaction.ProtoReflect.Descriptor instead.
func (*PhoneCodeTransaction) Descriptor() ([]byte, []int) {
	return file_code_tl_proto_rawDescGZIP(), []int{0}
}

func (x *PhoneCodeTransaction) GetPredicateName() string {
	if x != nil {
		return x.PredicateName
	}
	return ""
}

func (x *PhoneCodeTransaction) GetConstructor() TLConstructor {
	if x != nil {
		return x.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (x *PhoneCodeTransaction) GetAuthKeyId() int64 {
	if x != nil {
		return x.AuthKeyId
	}
	return 0
}

func (x *PhoneCodeTransaction) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *PhoneCodeTransaction) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *PhoneCodeTransaction) GetPhoneNumberRegistered() bool {
	if x != nil {
		return x.PhoneNumberRegistered
	}
	return false
}

func (x *PhoneCodeTransaction) GetPhoneCode() string {
	if x != nil {
		return x.PhoneCode
	}
	return ""
}

func (x *PhoneCodeTransaction) GetPhoneCodeHash() string {
	if x != nil {
		return x.PhoneCodeHash
	}
	return ""
}

func (x *PhoneCodeTransaction) GetPhoneCodeExpired() int32 {
	if x != nil {
		return x.PhoneCodeExpired
	}
	return 0
}

func (x *PhoneCodeTransaction) GetPhoneCodeExtraData() string {
	if x != nil {
		return x.PhoneCodeExtraData
	}
	return ""
}

func (x *PhoneCodeTransaction) GetSentCodeType() int32 {
	if x != nil {
		return x.SentCodeType
	}
	return 0
}

func (x *PhoneCodeTransaction) GetFlashCallPattern() string {
	if x != nil {
		return x.FlashCallPattern
	}
	return ""
}

func (x *PhoneCodeTransaction) GetNextCodeType() int32 {
	if x != nil {
		return x.NextCodeType
	}
	return 0
}

func (x *PhoneCodeTransaction) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type TLPhoneCodeTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data2 *PhoneCodeTransaction `protobuf:"bytes,1,opt,name=data2,proto3" json:"data2,omitempty"`
}

func (x *TLPhoneCodeTransaction) Reset() {
	*x = TLPhoneCodeTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_tl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLPhoneCodeTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLPhoneCodeTransaction) ProtoMessage() {}

func (x *TLPhoneCodeTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_code_tl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLPhoneCodeTransaction.ProtoReflect.Descriptor instead.
func (*TLPhoneCodeTransaction) Descriptor() ([]byte, []int) {
	return file_code_tl_proto_rawDescGZIP(), []int{1}
}

func (x *TLPhoneCodeTransaction) GetData2() *PhoneCodeTransaction {
	if x != nil {
		return x.Data2
	}
	return nil
}

// --------------------------------------------------------------------------------------------
type TLCodeCreatePhoneCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constructor           TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=code.TLConstructor" json:"constructor,omitempty"`
	AuthKeyId             int64         `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	SessionId             int64         `protobuf:"varint,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Phone                 string        `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	PhoneNumberRegistered bool          `protobuf:"varint,6,opt,name=phone_number_registered,json=phoneNumberRegistered,proto3" json:"phone_number_registered,omitempty"`
	SentCodeType          int32         `protobuf:"varint,7,opt,name=sent_code_type,json=sentCodeType,proto3" json:"sent_code_type,omitempty"`
	NextCodeType          int32         `protobuf:"varint,8,opt,name=next_code_type,json=nextCodeType,proto3" json:"next_code_type,omitempty"`
	State                 int32         `protobuf:"varint,9,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *TLCodeCreatePhoneCode) Reset() {
	*x = TLCodeCreatePhoneCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_tl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLCodeCreatePhoneCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLCodeCreatePhoneCode) ProtoMessage() {}

func (x *TLCodeCreatePhoneCode) ProtoReflect() protoreflect.Message {
	mi := &file_code_tl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLCodeCreatePhoneCode.ProtoReflect.Descriptor instead.
func (*TLCodeCreatePhoneCode) Descriptor() ([]byte, []int) {
	return file_code_tl_proto_rawDescGZIP(), []int{2}
}

func (x *TLCodeCreatePhoneCode) GetConstructor() TLConstructor {
	if x != nil {
		return x.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (x *TLCodeCreatePhoneCode) GetAuthKeyId() int64 {
	if x != nil {
		return x.AuthKeyId
	}
	return 0
}

func (x *TLCodeCreatePhoneCode) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *TLCodeCreatePhoneCode) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *TLCodeCreatePhoneCode) GetPhoneNumberRegistered() bool {
	if x != nil {
		return x.PhoneNumberRegistered
	}
	return false
}

func (x *TLCodeCreatePhoneCode) GetSentCodeType() int32 {
	if x != nil {
		return x.SentCodeType
	}
	return 0
}

func (x *TLCodeCreatePhoneCode) GetNextCodeType() int32 {
	if x != nil {
		return x.NextCodeType
	}
	return 0
}

func (x *TLCodeCreatePhoneCode) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

// --------------------------------------------------------------------------------------------
type TLCodeGetPhoneCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constructor   TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=code.TLConstructor" json:"constructor,omitempty"`
	AuthKeyId     int64         `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	Phone         string        `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	PhoneCodeHash string        `protobuf:"bytes,5,opt,name=phone_code_hash,json=phoneCodeHash,proto3" json:"phone_code_hash,omitempty"`
}

func (x *TLCodeGetPhoneCode) Reset() {
	*x = TLCodeGetPhoneCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_tl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLCodeGetPhoneCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLCodeGetPhoneCode) ProtoMessage() {}

func (x *TLCodeGetPhoneCode) ProtoReflect() protoreflect.Message {
	mi := &file_code_tl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLCodeGetPhoneCode.ProtoReflect.Descriptor instead.
func (*TLCodeGetPhoneCode) Descriptor() ([]byte, []int) {
	return file_code_tl_proto_rawDescGZIP(), []int{3}
}

func (x *TLCodeGetPhoneCode) GetConstructor() TLConstructor {
	if x != nil {
		return x.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (x *TLCodeGetPhoneCode) GetAuthKeyId() int64 {
	if x != nil {
		return x.AuthKeyId
	}
	return 0
}

func (x *TLCodeGetPhoneCode) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *TLCodeGetPhoneCode) GetPhoneCodeHash() string {
	if x != nil {
		return x.PhoneCodeHash
	}
	return ""
}

// --------------------------------------------------------------------------------------------
type TLCodeDeletePhoneCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constructor   TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=code.TLConstructor" json:"constructor,omitempty"`
	AuthKeyId     int64         `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	Phone         string        `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	PhoneCodeHash string        `protobuf:"bytes,5,opt,name=phone_code_hash,json=phoneCodeHash,proto3" json:"phone_code_hash,omitempty"`
}

func (x *TLCodeDeletePhoneCode) Reset() {
	*x = TLCodeDeletePhoneCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_tl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLCodeDeletePhoneCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLCodeDeletePhoneCode) ProtoMessage() {}

func (x *TLCodeDeletePhoneCode) ProtoReflect() protoreflect.Message {
	mi := &file_code_tl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLCodeDeletePhoneCode.ProtoReflect.Descriptor instead.
func (*TLCodeDeletePhoneCode) Descriptor() ([]byte, []int) {
	return file_code_tl_proto_rawDescGZIP(), []int{4}
}

func (x *TLCodeDeletePhoneCode) GetConstructor() TLConstructor {
	if x != nil {
		return x.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (x *TLCodeDeletePhoneCode) GetAuthKeyId() int64 {
	if x != nil {
		return x.AuthKeyId
	}
	return 0
}

func (x *TLCodeDeletePhoneCode) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *TLCodeDeletePhoneCode) GetPhoneCodeHash() string {
	if x != nil {
		return x.PhoneCodeHash
	}
	return ""
}

// --------------------------------------------------------------------------------------------
type TLCodeUpdatePhoneCodeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constructor   TLConstructor         `protobuf:"varint,1,opt,name=constructor,proto3,enum=code.TLConstructor" json:"constructor,omitempty"`
	AuthKeyId     int64                 `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	Phone         string                `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	PhoneCodeHash string                `protobuf:"bytes,5,opt,name=phone_code_hash,json=phoneCodeHash,proto3" json:"phone_code_hash,omitempty"`
	CodeData      *PhoneCodeTransaction `protobuf:"bytes,6,opt,name=code_data,json=codeData,proto3" json:"code_data,omitempty"`
}

func (x *TLCodeUpdatePhoneCodeData) Reset() {
	*x = TLCodeUpdatePhoneCodeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_tl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLCodeUpdatePhoneCodeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLCodeUpdatePhoneCodeData) ProtoMessage() {}

func (x *TLCodeUpdatePhoneCodeData) ProtoReflect() protoreflect.Message {
	mi := &file_code_tl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLCodeUpdatePhoneCodeData.ProtoReflect.Descriptor instead.
func (*TLCodeUpdatePhoneCodeData) Descriptor() ([]byte, []int) {
	return file_code_tl_proto_rawDescGZIP(), []int{5}
}

func (x *TLCodeUpdatePhoneCodeData) GetConstructor() TLConstructor {
	if x != nil {
		return x.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (x *TLCodeUpdatePhoneCodeData) GetAuthKeyId() int64 {
	if x != nil {
		return x.AuthKeyId
	}
	return 0
}

func (x *TLCodeUpdatePhoneCodeData) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *TLCodeUpdatePhoneCodeData) GetPhoneCodeHash() string {
	if x != nil {
		return x.PhoneCodeHash
	}
	return ""
}

func (x *TLCodeUpdatePhoneCodeData) GetCodeData() *PhoneCodeTransaction {
	if x != nil {
		return x.CodeData
	}
	return nil
}

var File_code_tl_proto protoreflect.FileDescriptor

var file_code_tl_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x74, 0x6c,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x74, 0x6c, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x74, 0x6c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb9, 0x04, 0x0a, 0x14, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x54,
	0x4c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x36,
	0x0a, 0x17, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x15, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2c, 0x0a,
	0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24,
	0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x4b,
	0x0a, 0x17, 0x54, 0x4c, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x32, 0x22, 0xbf, 0x02, 0x0a, 0x17,
	0x54, 0x4c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x54, 0x4c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1e,
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xab, 0x01,
	0x0a, 0x14, 0x54, 0x4c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x54, 0x4c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a,
	0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x22, 0xae, 0x01, 0x0a, 0x17,
	0x54, 0x4c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x54, 0x4c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1e,
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x22, 0xeb, 0x01, 0x0a,
	0x1b, 0x54, 0x4c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x54, 0x4c, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2a, 0xe6, 0x01, 0x0a, 0x0d, 0x54,
	0x4c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x52, 0x43, 0x33, 0x32, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x27, 0x0a, 0x1a, 0x43, 0x52, 0x43, 0x33, 0x32, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x98, 0xad,
	0xce, 0x9b, 0xf8, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x22, 0x0a, 0x1a, 0x43, 0x52, 0x43, 0x33,
	0x32, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x10, 0x9e, 0xc1, 0x8f, 0x81, 0x06, 0x12, 0x1f, 0x0a, 0x17,
	0x43, 0x52, 0x43, 0x33, 0x32, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x10, 0xf9, 0xc1, 0x92, 0x8d, 0x06, 0x12, 0x27, 0x0a,
	0x1a, 0x43, 0x52, 0x43, 0x33, 0x32, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x10, 0xd0, 0xd4, 0xc1, 0xb5,
	0xfa, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x2b, 0x0a, 0x1e, 0x43, 0x52, 0x43, 0x33, 0x32, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x10, 0x95, 0x95, 0xd4, 0xb4, 0xfb, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x32, 0xc5, 0x02, 0x0a, 0x07, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x53, 0x0a, 0x14, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x54,
	0x4c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x11, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x65, 0x74,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x54, 0x4c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x14, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x54, 0x4c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x0d, 0x2e, 0x6d, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x18, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x54,
	0x4c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6d, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x67, 0x72,
	0x61, 0x6d, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x67, 0x72, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62,
	0x69, 0x7a, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_code_tl_proto_rawDescOnce sync.Once
	file_code_tl_proto_rawDescData = file_code_tl_proto_rawDesc
)

func file_code_tl_proto_rawDescGZIP() []byte {
	file_code_tl_proto_rawDescOnce.Do(func() {
		file_code_tl_proto_rawDescData = protoimpl.X.CompressGZIP(file_code_tl_proto_rawDescData)
	})
	return file_code_tl_proto_rawDescData
}

var file_code_tl_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_code_tl_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_code_tl_proto_goTypes = []interface{}{
	(TLConstructor)(0),                // 0: code.TLConstructor
	(*PhoneCodeTransaction)(nil),      // 1: code.PhoneCodeTransaction
	(*TLPhoneCodeTransaction)(nil),    // 2: code.TL_phoneCodeTransaction
	(*TLCodeCreatePhoneCode)(nil),     // 3: code.TL_code_createPhoneCode
	(*TLCodeGetPhoneCode)(nil),        // 4: code.TL_code_getPhoneCode
	(*TLCodeDeletePhoneCode)(nil),     // 5: code.TL_code_deletePhoneCode
	(*TLCodeUpdatePhoneCodeData)(nil), // 6: code.TL_code_updatePhoneCodeData
	(*mtproto.Bool)(nil),              // 7: mtproto.Bool
}
var file_code_tl_proto_depIdxs = []int32{
	0,  // 0: code.PhoneCodeTransaction.constructor:type_name -> code.TLConstructor
	1,  // 1: code.TL_phoneCodeTransaction.data2:type_name -> code.PhoneCodeTransaction
	0,  // 2: code.TL_code_createPhoneCode.constructor:type_name -> code.TLConstructor
	0,  // 3: code.TL_code_getPhoneCode.constructor:type_name -> code.TLConstructor
	0,  // 4: code.TL_code_deletePhoneCode.constructor:type_name -> code.TLConstructor
	0,  // 5: code.TL_code_updatePhoneCodeData.constructor:type_name -> code.TLConstructor
	1,  // 6: code.TL_code_updatePhoneCodeData.code_data:type_name -> code.PhoneCodeTransaction
	3,  // 7: code.RPCCode.code_createPhoneCode:input_type -> code.TL_code_createPhoneCode
	4,  // 8: code.RPCCode.code_getPhoneCode:input_type -> code.TL_code_getPhoneCode
	5,  // 9: code.RPCCode.code_deletePhoneCode:input_type -> code.TL_code_deletePhoneCode
	6,  // 10: code.RPCCode.code_updatePhoneCodeData:input_type -> code.TL_code_updatePhoneCodeData
	1,  // 11: code.RPCCode.code_createPhoneCode:output_type -> code.PhoneCodeTransaction
	1,  // 12: code.RPCCode.code_getPhoneCode:output_type -> code.PhoneCodeTransaction
	7,  // 13: code.RPCCode.code_deletePhoneCode:output_type -> mtproto.Bool
	7,  // 14: code.RPCCode.code_updatePhoneCodeData:output_type -> mtproto.Bool
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_code_tl_proto_init() }
func file_code_tl_proto_init() {
	if File_code_tl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_code_tl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneCodeTransaction); i {
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
		file_code_tl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLPhoneCodeTransaction); i {
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
		file_code_tl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLCodeCreatePhoneCode); i {
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
		file_code_tl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLCodeGetPhoneCode); i {
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
		file_code_tl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLCodeDeletePhoneCode); i {
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
		file_code_tl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLCodeUpdatePhoneCodeData); i {
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
			RawDescriptor: file_code_tl_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_code_tl_proto_goTypes,
		DependencyIndexes: file_code_tl_proto_depIdxs,
		EnumInfos:         file_code_tl_proto_enumTypes,
		MessageInfos:      file_code_tl_proto_msgTypes,
	}.Build()
	File_code_tl_proto = out.File
	file_code_tl_proto_rawDesc = nil
	file_code_tl_proto_goTypes = nil
	file_code_tl_proto_depIdxs = nil
}
