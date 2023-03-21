//
// (C) Copyright 2019-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.5.0
// source: mgmt/acl.proto

package mgmt

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

// Access control list and ownership information
type AccessControlList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries    []string `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`                         // List of ACEs in short string format
	OwnerUser  string   `protobuf:"bytes,2,opt,name=owner_user,json=ownerUser,proto3" json:"owner_user,omitempty"`    // Name of user that owns the resource
	OwnerGroup string   `protobuf:"bytes,3,opt,name=owner_group,json=ownerGroup,proto3" json:"owner_group,omitempty"` // Name of group that owns the resource
}

func (x *AccessControlList) Reset() {
	*x = AccessControlList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_acl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessControlList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessControlList) ProtoMessage() {}

func (x *AccessControlList) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_acl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessControlList.ProtoReflect.Descriptor instead.
func (*AccessControlList) Descriptor() ([]byte, []int) {
	return file_mgmt_acl_proto_rawDescGZIP(), []int{0}
}

func (x *AccessControlList) GetEntries() []string {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AccessControlList) GetOwnerUser() string {
	if x != nil {
		return x.OwnerUser
	}
	return ""
}

func (x *AccessControlList) GetOwnerGroup() string {
	if x != nil {
		return x.OwnerGroup
	}
	return ""
}

// Response to ACL-related requests includes the command status and current ACL
type ACLResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32              `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"` // DAOS error code
	Acl    *AccessControlList `protobuf:"bytes,2,opt,name=acl,proto3" json:"acl,omitempty"`
}

func (x *ACLResp) Reset() {
	*x = ACLResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_acl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLResp) ProtoMessage() {}

func (x *ACLResp) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_acl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLResp.ProtoReflect.Descriptor instead.
func (*ACLResp) Descriptor() ([]byte, []int) {
	return file_mgmt_acl_proto_rawDescGZIP(), []int{1}
}

func (x *ACLResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ACLResp) GetAcl() *AccessControlList {
	if x != nil {
		return x.Acl
	}
	return nil
}

// Request to fetch an ACL
type GetACLReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys      string   `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty"`                                   // DAOS system identifier
	Id       string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                     // Target UUID or label
	SvcRanks []uint32 `protobuf:"varint,3,rep,packed,name=svc_ranks,json=svcRanks,proto3" json:"svc_ranks,omitempty"` // List of pool service ranks
}

func (x *GetACLReq) Reset() {
	*x = GetACLReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_acl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetACLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetACLReq) ProtoMessage() {}

func (x *GetACLReq) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_acl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetACLReq.ProtoReflect.Descriptor instead.
func (*GetACLReq) Descriptor() ([]byte, []int) {
	return file_mgmt_acl_proto_rawDescGZIP(), []int{2}
}

func (x *GetACLReq) GetSys() string {
	if x != nil {
		return x.Sys
	}
	return ""
}

func (x *GetACLReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetACLReq) GetSvcRanks() []uint32 {
	if x != nil {
		return x.SvcRanks
	}
	return nil
}

// Request to modify an ACL
// Results depend on the specific modification command.
type ModifyACLReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys      string   `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty"`                                   // DAOS system identifier
	Id       string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                     // Target UUID or label
	Entries  []string `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`                           // List of ACEs to overwrite ACL with
	SvcRanks []uint32 `protobuf:"varint,4,rep,packed,name=svc_ranks,json=svcRanks,proto3" json:"svc_ranks,omitempty"` // List of pool service ranks
}

func (x *ModifyACLReq) Reset() {
	*x = ModifyACLReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_acl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyACLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyACLReq) ProtoMessage() {}

func (x *ModifyACLReq) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_acl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyACLReq.ProtoReflect.Descriptor instead.
func (*ModifyACLReq) Descriptor() ([]byte, []int) {
	return file_mgmt_acl_proto_rawDescGZIP(), []int{3}
}

func (x *ModifyACLReq) GetSys() string {
	if x != nil {
		return x.Sys
	}
	return ""
}

func (x *ModifyACLReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ModifyACLReq) GetEntries() []string {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *ModifyACLReq) GetSvcRanks() []uint32 {
	if x != nil {
		return x.SvcRanks
	}
	return nil
}

// Delete a principal's entry from the ACL
type DeleteACLReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys       string   `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty"`                                   // DAOS system identifier
	Id        string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                     // Target UUID or label
	Principal string   `protobuf:"bytes,3,opt,name=principal,proto3" json:"principal,omitempty"`                       // Principal whose entry is to be deleted
	SvcRanks  []uint32 `protobuf:"varint,4,rep,packed,name=svc_ranks,json=svcRanks,proto3" json:"svc_ranks,omitempty"` // List of pool service ranks
}

func (x *DeleteACLReq) Reset() {
	*x = DeleteACLReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_acl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteACLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteACLReq) ProtoMessage() {}

func (x *DeleteACLReq) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_acl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteACLReq.ProtoReflect.Descriptor instead.
func (*DeleteACLReq) Descriptor() ([]byte, []int) {
	return file_mgmt_acl_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteACLReq) GetSys() string {
	if x != nil {
		return x.Sys
	}
	return ""
}

func (x *DeleteACLReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteACLReq) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *DeleteACLReq) GetSvcRanks() []uint32 {
	if x != nil {
		return x.SvcRanks
	}
	return nil
}

var File_mgmt_acl_proto protoreflect.FileDescriptor

var file_mgmt_acl_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x6d, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x4c, 0x0a, 0x07, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x03, 0x61, 0x63, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x03,
	0x61, 0x63, 0x6c, 0x22, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x79, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x76, 0x63, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x76, 0x63, 0x52, 0x61, 0x6e, 0x6b, 0x73, 0x22,
	0x67, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x79,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x76, 0x63, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08,
	0x73, 0x76, 0x63, 0x52, 0x61, 0x6e, 0x6b, 0x73, 0x22, 0x6b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x79, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x76, 0x63, 0x5f,
	0x72, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x76, 0x63,
	0x52, 0x61, 0x6e, 0x6b, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6f, 0x73, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x64,
	0x61, 0x6f, 0x73, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x67, 0x6d,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mgmt_acl_proto_rawDescOnce sync.Once
	file_mgmt_acl_proto_rawDescData = file_mgmt_acl_proto_rawDesc
)

func file_mgmt_acl_proto_rawDescGZIP() []byte {
	file_mgmt_acl_proto_rawDescOnce.Do(func() {
		file_mgmt_acl_proto_rawDescData = protoimpl.X.CompressGZIP(file_mgmt_acl_proto_rawDescData)
	})
	return file_mgmt_acl_proto_rawDescData
}

var file_mgmt_acl_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mgmt_acl_proto_goTypes = []interface{}{
	(*AccessControlList)(nil), // 0: mgmt.AccessControlList
	(*ACLResp)(nil),           // 1: mgmt.ACLResp
	(*GetACLReq)(nil),         // 2: mgmt.GetACLReq
	(*ModifyACLReq)(nil),      // 3: mgmt.ModifyACLReq
	(*DeleteACLReq)(nil),      // 4: mgmt.DeleteACLReq
}
var file_mgmt_acl_proto_depIdxs = []int32{
	0, // 0: mgmt.ACLResp.acl:type_name -> mgmt.AccessControlList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mgmt_acl_proto_init() }
func file_mgmt_acl_proto_init() {
	if File_mgmt_acl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mgmt_acl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessControlList); i {
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
		file_mgmt_acl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACLResp); i {
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
		file_mgmt_acl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetACLReq); i {
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
		file_mgmt_acl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyACLReq); i {
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
		file_mgmt_acl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteACLReq); i {
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
			RawDescriptor: file_mgmt_acl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mgmt_acl_proto_goTypes,
		DependencyIndexes: file_mgmt_acl_proto_depIdxs,
		MessageInfos:      file_mgmt_acl_proto_msgTypes,
	}.Build()
	File_mgmt_acl_proto = out.File
	file_mgmt_acl_proto_rawDesc = nil
	file_mgmt_acl_proto_goTypes = nil
	file_mgmt_acl_proto_depIdxs = nil
}
