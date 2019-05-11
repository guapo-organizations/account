// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account/account.proto

package account

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//数据库,基本的账户信息
type UserBaseInfo struct {
	//id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//邮箱
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	//手机号
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	//用户的真实姓名
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBaseInfo) Reset()         { *m = UserBaseInfo{} }
func (m *UserBaseInfo) String() string { return proto.CompactTextString(m) }
func (*UserBaseInfo) ProtoMessage()    {}
func (*UserBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{0}
}

func (m *UserBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBaseInfo.Unmarshal(m, b)
}
func (m *UserBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBaseInfo.Marshal(b, m, deterministic)
}
func (m *UserBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBaseInfo.Merge(m, src)
}
func (m *UserBaseInfo) XXX_Size() int {
	return xxx_messageInfo_UserBaseInfo.Size(m)
}
func (m *UserBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserBaseInfo proto.InternalMessageInfo

func (m *UserBaseInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserBaseInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserBaseInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserBaseInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//数据库,用户绑定的三方账号
type UserPlatformInfo struct {
	//账户类型
	Type int64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	//账户id
	AccountId int64 `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	//三方平台标识的用户id
	PlatformId           string   `protobuf:"bytes,3,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPlatformInfo) Reset()         { *m = UserPlatformInfo{} }
func (m *UserPlatformInfo) String() string { return proto.CompactTextString(m) }
func (*UserPlatformInfo) ProtoMessage()    {}
func (*UserPlatformInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{1}
}

func (m *UserPlatformInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPlatformInfo.Unmarshal(m, b)
}
func (m *UserPlatformInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPlatformInfo.Marshal(b, m, deterministic)
}
func (m *UserPlatformInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPlatformInfo.Merge(m, src)
}
func (m *UserPlatformInfo) XXX_Size() int {
	return xxx_messageInfo_UserPlatformInfo.Size(m)
}
func (m *UserPlatformInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPlatformInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserPlatformInfo proto.InternalMessageInfo

func (m *UserPlatformInfo) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *UserPlatformInfo) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *UserPlatformInfo) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

//获取用户的信息
type UserInfo struct {
	UserBaseInfo         *UserBaseInfo       `protobuf:"bytes,1,opt,name=user_base_info,json=userBaseInfo,proto3" json:"user_base_info,omitempty"`
	UserPlatformInfo     []*UserPlatformInfo `protobuf:"bytes,2,rep,name=user_platform_info,json=userPlatformInfo,proto3" json:"user_platform_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{2}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserBaseInfo() *UserBaseInfo {
	if m != nil {
		return m.UserBaseInfo
	}
	return nil
}

func (m *UserInfo) GetUserPlatformInfo() []*UserPlatformInfo {
	if m != nil {
		return m.UserPlatformInfo
	}
	return nil
}

//发送消息的结果
type SendMesaageResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMesaageResponse) Reset()         { *m = SendMesaageResponse{} }
func (m *SendMesaageResponse) String() string { return proto.CompactTextString(m) }
func (*SendMesaageResponse) ProtoMessage()    {}
func (*SendMesaageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{3}
}

func (m *SendMesaageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMesaageResponse.Unmarshal(m, b)
}
func (m *SendMesaageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMesaageResponse.Marshal(b, m, deterministic)
}
func (m *SendMesaageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMesaageResponse.Merge(m, src)
}
func (m *SendMesaageResponse) XXX_Size() int {
	return xxx_messageInfo_SendMesaageResponse.Size(m)
}
func (m *SendMesaageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMesaageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendMesaageResponse proto.InternalMessageInfo

func (m *SendMesaageResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SendMesaageResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

//code校验结果
type ValidataCodeResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidataCodeResponse) Reset()         { *m = ValidataCodeResponse{} }
func (m *ValidataCodeResponse) String() string { return proto.CompactTextString(m) }
func (*ValidataCodeResponse) ProtoMessage()    {}
func (*ValidataCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{4}
}

func (m *ValidataCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidataCodeResponse.Unmarshal(m, b)
}
func (m *ValidataCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidataCodeResponse.Marshal(b, m, deterministic)
}
func (m *ValidataCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidataCodeResponse.Merge(m, src)
}
func (m *ValidataCodeResponse) XXX_Size() int {
	return xxx_messageInfo_ValidataCodeResponse.Size(m)
}
func (m *ValidataCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidataCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidataCodeResponse proto.InternalMessageInfo

func (m *ValidataCodeResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ValidataCodeResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

//邮件或手机号注册请求
type EmailOrPhoneRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Passwd               string   `protobuf:"bytes,3,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailOrPhoneRequest) Reset()         { *m = EmailOrPhoneRequest{} }
func (m *EmailOrPhoneRequest) String() string { return proto.CompactTextString(m) }
func (*EmailOrPhoneRequest) ProtoMessage()    {}
func (*EmailOrPhoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{5}
}

func (m *EmailOrPhoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailOrPhoneRequest.Unmarshal(m, b)
}
func (m *EmailOrPhoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailOrPhoneRequest.Marshal(b, m, deterministic)
}
func (m *EmailOrPhoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailOrPhoneRequest.Merge(m, src)
}
func (m *EmailOrPhoneRequest) XXX_Size() int {
	return xxx_messageInfo_EmailOrPhoneRequest.Size(m)
}
func (m *EmailOrPhoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailOrPhoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailOrPhoneRequest proto.InternalMessageInfo

func (m *EmailOrPhoneRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailOrPhoneRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *EmailOrPhoneRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

//三方类型定义
type Type struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{6}
}

func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (m *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(m, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

func (m *Type) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Type) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//三方类型
type PlatformTypeList struct {
	Type                 []*Type  `protobuf:"bytes,1,rep,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformTypeList) Reset()         { *m = PlatformTypeList{} }
func (m *PlatformTypeList) String() string { return proto.CompactTextString(m) }
func (*PlatformTypeList) ProtoMessage()    {}
func (*PlatformTypeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{7}
}

func (m *PlatformTypeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformTypeList.Unmarshal(m, b)
}
func (m *PlatformTypeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformTypeList.Marshal(b, m, deterministic)
}
func (m *PlatformTypeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformTypeList.Merge(m, src)
}
func (m *PlatformTypeList) XXX_Size() int {
	return xxx_messageInfo_PlatformTypeList.Size(m)
}
func (m *PlatformTypeList) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformTypeList.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformTypeList proto.InternalMessageInfo

func (m *PlatformTypeList) GetType() []*Type {
	if m != nil {
		return m.Type
	}
	return nil
}

//三方注册请求
type PlatformRequest struct {
	Type                 *Type    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	PlatformId           string   `protobuf:"bytes,2,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformRequest) Reset()         { *m = PlatformRequest{} }
func (m *PlatformRequest) String() string { return proto.CompactTextString(m) }
func (*PlatformRequest) ProtoMessage()    {}
func (*PlatformRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{8}
}

func (m *PlatformRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformRequest.Unmarshal(m, b)
}
func (m *PlatformRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformRequest.Marshal(b, m, deterministic)
}
func (m *PlatformRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformRequest.Merge(m, src)
}
func (m *PlatformRequest) XXX_Size() int {
	return xxx_messageInfo_PlatformRequest.Size(m)
}
func (m *PlatformRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformRequest proto.InternalMessageInfo

func (m *PlatformRequest) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *PlatformRequest) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

func init() {
	proto.RegisterType((*UserBaseInfo)(nil), "account.UserBaseInfo")
	proto.RegisterType((*UserPlatformInfo)(nil), "account.UserPlatformInfo")
	proto.RegisterType((*UserInfo)(nil), "account.UserInfo")
	proto.RegisterType((*SendMesaageResponse)(nil), "account.SendMesaageResponse")
	proto.RegisterType((*ValidataCodeResponse)(nil), "account.ValidataCodeResponse")
	proto.RegisterType((*EmailOrPhoneRequest)(nil), "account.EmailOrPhoneRequest")
	proto.RegisterType((*Type)(nil), "account.Type")
	proto.RegisterType((*PlatformTypeList)(nil), "account.PlatformTypeList")
	proto.RegisterType((*PlatformRequest)(nil), "account.PlatformRequest")
}

func init() { proto.RegisterFile("proto/account/account.proto", fileDescriptor_5d492a0187472a3b) }

var fileDescriptor_5d492a0187472a3b = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x25, 0x09, 0xa3, 0xed, 0x85, 0x76, 0xd4, 0xed, 0x50, 0xf6, 0x81, 0x46, 0xfd, 0xc4, 0x13,
	0x93, 0x98, 0xf6, 0xb4, 0xa7, 0x32, 0x75, 0x15, 0x52, 0xab, 0x55, 0xd9, 0x5a, 0x69, 0x2f, 0x43,
	0x6e, 0x73, 0x61, 0x96, 0x20, 0xce, 0x62, 0x47, 0x13, 0xff, 0xa2, 0x3f, 0x79, 0xb2, 0x63, 0x87,
	0x14, 0x65, 0x63, 0x4f, 0xe4, 0x1e, 0xe7, 0x9e, 0x13, 0x9f, 0x73, 0x04, 0xbc, 0x4e, 0x33, 0xa1,
	0xc4, 0x3b, 0xf6, 0xf0, 0x20, 0xf2, 0x44, 0xb9, 0xdf, 0x91, 0x41, 0xc9, 0x9e, 0x1d, 0xe9, 0x0f,
	0xe8, 0xdc, 0x4a, 0xcc, 0x26, 0x4c, 0xe2, 0x34, 0x99, 0x0b, 0x72, 0x04, 0x3e, 0x8f, 0x43, 0x6f,
	0xe0, 0x0d, 0x83, 0xc8, 0xe7, 0x31, 0x39, 0x85, 0x67, 0xb8, 0x62, 0x7c, 0x19, 0xfa, 0x03, 0x6f,
	0x78, 0x10, 0x15, 0x83, 0x46, 0xd3, 0x9f, 0x22, 0xc1, 0x30, 0x28, 0x50, 0x33, 0x10, 0x02, 0xcd,
	0x84, 0xad, 0x30, 0x6c, 0x1a, 0xd0, 0x3c, 0xd3, 0x39, 0x74, 0x35, 0xff, 0xcd, 0x92, 0xa9, 0xb9,
	0xc8, 0x56, 0x46, 0x83, 0x40, 0x53, 0xad, 0x53, 0xb4, 0x2a, 0xe6, 0x99, 0xf4, 0x01, 0xec, 0x27,
	0xcd, 0x78, 0x6c, 0xc4, 0x82, 0xe8, 0xc0, 0x22, 0xd3, 0x98, 0xbc, 0x85, 0x76, 0x6a, 0x29, 0xf4,
	0x79, 0x21, 0x0b, 0x0e, 0x9a, 0xc6, 0xf4, 0xd1, 0x83, 0x7d, 0x2d, 0x64, 0x04, 0x3e, 0xc2, 0x51,
	0x2e, 0x31, 0x9b, 0xdd, 0x33, 0x89, 0x33, 0x9e, 0xcc, 0x85, 0x91, 0x6a, 0x8f, 0x5f, 0x8c, 0x9c,
	0x0b, 0xd5, 0x3b, 0x47, 0x9d, 0xbc, 0xea, 0xc0, 0x25, 0x10, 0xb3, 0xbc, 0xd1, 0xd3, 0x04, 0xfe,
	0x20, 0x18, 0xb6, 0xc7, 0x2f, 0x9f, 0x10, 0x54, 0x2f, 0x15, 0x75, 0xf3, 0x2d, 0x84, 0x5e, 0xc0,
	0xc9, 0x57, 0x4c, 0xe2, 0x6b, 0x94, 0x8c, 0x2d, 0x30, 0x42, 0x99, 0x8a, 0x44, 0x22, 0xe9, 0x41,
	0x4b, 0x2a, 0xa6, 0x72, 0x69, 0xef, 0x6f, 0x27, 0x8d, 0x63, 0x96, 0xad, 0xe4, 0xc2, 0x5a, 0x6d,
	0x27, 0xfa, 0x19, 0x4e, 0xef, 0xd8, 0x92, 0xc7, 0x4c, 0xb1, 0x4f, 0x22, 0xfe, 0x1b, 0xcf, 0xfe,
	0x4e, 0x9e, 0xef, 0x70, 0x72, 0xa1, 0xc3, 0xfb, 0x92, 0xdd, 0xe8, 0xb4, 0x22, 0xfc, 0x95, 0xa3,
	0x54, 0x9b, 0x80, 0xbd, 0xda, 0x80, 0xfd, 0x6a, 0xc0, 0x3d, 0x68, 0xa5, 0x4c, 0xca, 0xdf, 0x2e,
	0x00, 0x3b, 0xd1, 0x11, 0x34, 0xbf, 0xe9, 0x10, 0xeb, 0x82, 0x75, 0xa5, 0xf0, 0x2b, 0xa5, 0xf8,
	0x00, 0x5d, 0xe7, 0x94, 0xde, 0xbb, 0xe2, 0x52, 0x91, 0xb3, 0x72, 0x57, 0x1b, 0x7d, 0x58, 0x1a,
	0xad, 0x5f, 0x28, 0xa8, 0xe8, 0x2d, 0x3c, 0x77, 0x6b, 0xee, 0xeb, 0xcf, 0x2a, 0x8a, 0xf5, 0x5b,
	0xdb, 0xd5, 0xf1, 0xb7, 0xab, 0x33, 0x7e, 0x0c, 0x60, 0xef, 0xbc, 0xd8, 0x23, 0x77, 0xd0, 0x8f,
	0x70, 0xc1, 0xa5, 0xc2, 0xcc, 0x42, 0x93, 0x75, 0xd5, 0x35, 0xf2, 0xa6, 0x94, 0xa8, 0x31, 0xf3,
	0x55, 0x7d, 0xc1, 0x68, 0x83, 0x5c, 0xc3, 0xa1, 0xee, 0x82, 0xd9, 0xd1, 0x29, 0xee, 0xe0, 0xd9,
	0x9c, 0xd6, 0x34, 0x88, 0x36, 0x48, 0x04, 0xc7, 0xb6, 0x13, 0xf8, 0xbf, 0x94, 0xfd, 0xf2, 0xb4,
	0xae, 0x4d, 0xb4, 0x41, 0xce, 0xa1, 0x73, 0x25, 0x16, 0x3c, 0x71, 0x56, 0xfc, 0x9b, 0xee, 0xf8,
	0xc9, 0x4d, 0xed, 0x2d, 0x2f, 0xa1, 0x57, 0xa5, 0x98, 0xac, 0x5d, 0x5c, 0x24, 0x2c, 0x5f, 0xdf,
	0x4a, 0xb0, 0x96, 0xe8, 0xbe, 0x65, 0xfe, 0xa5, 0xde, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe0,
	0x93, 0x9f, 0xf9, 0xc4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	//通过email或者手机号注册一个用户
	RegisterAccountByEmailOrPhone(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*UserBaseInfo, error)
	//邮箱发送验证码
	SendEmailCode(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*SendMesaageResponse, error)
	//邮箱验证校验
	ValidateEmailCode(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*ValidataCodeResponse, error)
	//通过验证码，或者手机登录
	LoginAccount(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*UserInfo, error)
	//通过三方登录
	LoginAccountByPlatform(ctx context.Context, in *PlatformRequest, opts ...grpc.CallOption) (*UserInfo, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) RegisterAccountByEmailOrPhone(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*UserBaseInfo, error) {
	out := new(UserBaseInfo)
	err := c.cc.Invoke(ctx, "/account.Account/RegisterAccountByEmailOrPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) SendEmailCode(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*SendMesaageResponse, error) {
	out := new(SendMesaageResponse)
	err := c.cc.Invoke(ctx, "/account.Account/SendEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) ValidateEmailCode(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*ValidataCodeResponse, error) {
	out := new(ValidataCodeResponse)
	err := c.cc.Invoke(ctx, "/account.Account/ValidateEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) LoginAccount(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/account.Account/LoginAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) LoginAccountByPlatform(ctx context.Context, in *PlatformRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/account.Account/LoginAccountByPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	//通过email或者手机号注册一个用户
	RegisterAccountByEmailOrPhone(context.Context, *EmailOrPhoneRequest) (*UserBaseInfo, error)
	//邮箱发送验证码
	SendEmailCode(context.Context, *EmailOrPhoneRequest) (*SendMesaageResponse, error)
	//邮箱验证校验
	ValidateEmailCode(context.Context, *EmailOrPhoneRequest) (*ValidataCodeResponse, error)
	//通过验证码，或者手机登录
	LoginAccount(context.Context, *EmailOrPhoneRequest) (*UserInfo, error)
	//通过三方登录
	LoginAccountByPlatform(context.Context, *PlatformRequest) (*UserInfo, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) RegisterAccountByEmailOrPhone(ctx context.Context, req *EmailOrPhoneRequest) (*UserBaseInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccountByEmailOrPhone not implemented")
}
func (*UnimplementedAccountServer) SendEmailCode(ctx context.Context, req *EmailOrPhoneRequest) (*SendMesaageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailCode not implemented")
}
func (*UnimplementedAccountServer) ValidateEmailCode(ctx context.Context, req *EmailOrPhoneRequest) (*ValidataCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateEmailCode not implemented")
}
func (*UnimplementedAccountServer) LoginAccount(ctx context.Context, req *EmailOrPhoneRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAccount not implemented")
}
func (*UnimplementedAccountServer) LoginAccountByPlatform(ctx context.Context, req *PlatformRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAccountByPlatform not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_RegisterAccountByEmailOrPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailOrPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).RegisterAccountByEmailOrPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/RegisterAccountByEmailOrPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).RegisterAccountByEmailOrPhone(ctx, req.(*EmailOrPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_SendEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailOrPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SendEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/SendEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SendEmailCode(ctx, req.(*EmailOrPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_ValidateEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailOrPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).ValidateEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/ValidateEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).ValidateEmailCode(ctx, req.(*EmailOrPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_LoginAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailOrPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).LoginAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/LoginAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).LoginAccount(ctx, req.(*EmailOrPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_LoginAccountByPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).LoginAccountByPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/LoginAccountByPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).LoginAccountByPlatform(ctx, req.(*PlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAccountByEmailOrPhone",
			Handler:    _Account_RegisterAccountByEmailOrPhone_Handler,
		},
		{
			MethodName: "SendEmailCode",
			Handler:    _Account_SendEmailCode_Handler,
		},
		{
			MethodName: "ValidateEmailCode",
			Handler:    _Account_ValidateEmailCode_Handler,
		},
		{
			MethodName: "LoginAccount",
			Handler:    _Account_LoginAccount_Handler,
		},
		{
			MethodName: "LoginAccountByPlatform",
			Handler:    _Account_LoginAccountByPlatform_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account/account.proto",
}
