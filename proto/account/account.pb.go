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

//email注册用户
type RegisterAccountEmailRequester struct {
	//邮箱
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	//验证码
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	//密码
	Passwd               string   `protobuf:"bytes,3,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAccountEmailRequester) Reset()         { *m = RegisterAccountEmailRequester{} }
func (m *RegisterAccountEmailRequester) String() string { return proto.CompactTextString(m) }
func (*RegisterAccountEmailRequester) ProtoMessage()    {}
func (*RegisterAccountEmailRequester) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{0}
}

func (m *RegisterAccountEmailRequester) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAccountEmailRequester.Unmarshal(m, b)
}
func (m *RegisterAccountEmailRequester) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAccountEmailRequester.Marshal(b, m, deterministic)
}
func (m *RegisterAccountEmailRequester) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAccountEmailRequester.Merge(m, src)
}
func (m *RegisterAccountEmailRequester) XXX_Size() int {
	return xxx_messageInfo_RegisterAccountEmailRequester.Size(m)
}
func (m *RegisterAccountEmailRequester) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAccountEmailRequester.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAccountEmailRequester proto.InternalMessageInfo

func (m *RegisterAccountEmailRequester) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterAccountEmailRequester) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RegisterAccountEmailRequester) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

//email注册用户响应
type RegisterAccountEmailResponse struct {
	Errcode              int64    `protobuf:"varint,1,opt,name=errcode,proto3" json:"errcode,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAccountEmailResponse) Reset()         { *m = RegisterAccountEmailResponse{} }
func (m *RegisterAccountEmailResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterAccountEmailResponse) ProtoMessage()    {}
func (*RegisterAccountEmailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{1}
}

func (m *RegisterAccountEmailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAccountEmailResponse.Unmarshal(m, b)
}
func (m *RegisterAccountEmailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAccountEmailResponse.Marshal(b, m, deterministic)
}
func (m *RegisterAccountEmailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAccountEmailResponse.Merge(m, src)
}
func (m *RegisterAccountEmailResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterAccountEmailResponse.Size(m)
}
func (m *RegisterAccountEmailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAccountEmailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAccountEmailResponse proto.InternalMessageInfo

func (m *RegisterAccountEmailResponse) GetErrcode() int64 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *RegisterAccountEmailResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

//账号密码登录
type LoginByPasswordRequest struct {
	//邮箱
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	//手机号
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	//密码
	Passwd               string   `protobuf:"bytes,3,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginByPasswordRequest) Reset()         { *m = LoginByPasswordRequest{} }
func (m *LoginByPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*LoginByPasswordRequest) ProtoMessage()    {}
func (*LoginByPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{2}
}

func (m *LoginByPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginByPasswordRequest.Unmarshal(m, b)
}
func (m *LoginByPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginByPasswordRequest.Marshal(b, m, deterministic)
}
func (m *LoginByPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginByPasswordRequest.Merge(m, src)
}
func (m *LoginByPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_LoginByPasswordRequest.Size(m)
}
func (m *LoginByPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginByPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginByPasswordRequest proto.InternalMessageInfo

func (m *LoginByPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginByPasswordRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginByPasswordRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

//账号密码登录响应
type LoginByPasswordResponse struct {
	//错误码
	Error int64 `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	//错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//登录的token
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginByPasswordResponse) Reset()         { *m = LoginByPasswordResponse{} }
func (m *LoginByPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*LoginByPasswordResponse) ProtoMessage()    {}
func (*LoginByPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{3}
}

func (m *LoginByPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginByPasswordResponse.Unmarshal(m, b)
}
func (m *LoginByPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginByPasswordResponse.Marshal(b, m, deterministic)
}
func (m *LoginByPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginByPasswordResponse.Merge(m, src)
}
func (m *LoginByPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_LoginByPasswordResponse.Size(m)
}
func (m *LoginByPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginByPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginByPasswordResponse proto.InternalMessageInfo

func (m *LoginByPasswordResponse) GetError() int64 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *LoginByPasswordResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *LoginByPasswordResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//三方登录请求，没有的时候，自动注册
type OpenPlaformLoginRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//我们系统自己的三方平台标识
	TypeId uint64 `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	//比如微信的openid
	PlatformId string `protobuf:"bytes,3,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	//比如微信的unioncid
	UnionId              string   `protobuf:"bytes,4,opt,name=union_id,json=unionId,proto3" json:"union_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenPlaformLoginRequest) Reset()         { *m = OpenPlaformLoginRequest{} }
func (m *OpenPlaformLoginRequest) String() string { return proto.CompactTextString(m) }
func (*OpenPlaformLoginRequest) ProtoMessage()    {}
func (*OpenPlaformLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{4}
}

func (m *OpenPlaformLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenPlaformLoginRequest.Unmarshal(m, b)
}
func (m *OpenPlaformLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenPlaformLoginRequest.Marshal(b, m, deterministic)
}
func (m *OpenPlaformLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenPlaformLoginRequest.Merge(m, src)
}
func (m *OpenPlaformLoginRequest) XXX_Size() int {
	return xxx_messageInfo_OpenPlaformLoginRequest.Size(m)
}
func (m *OpenPlaformLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenPlaformLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenPlaformLoginRequest proto.InternalMessageInfo

func (m *OpenPlaformLoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OpenPlaformLoginRequest) GetTypeId() uint64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *OpenPlaformLoginRequest) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

func (m *OpenPlaformLoginRequest) GetUnionId() string {
	if m != nil {
		return m.UnionId
	}
	return ""
}

//三方登录响应
type OpenPlaformLoginResponse struct {
	//错误码
	Error int64 `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	//错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//登录的token
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenPlaformLoginResponse) Reset()         { *m = OpenPlaformLoginResponse{} }
func (m *OpenPlaformLoginResponse) String() string { return proto.CompactTextString(m) }
func (*OpenPlaformLoginResponse) ProtoMessage()    {}
func (*OpenPlaformLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{5}
}

func (m *OpenPlaformLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenPlaformLoginResponse.Unmarshal(m, b)
}
func (m *OpenPlaformLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenPlaformLoginResponse.Marshal(b, m, deterministic)
}
func (m *OpenPlaformLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenPlaformLoginResponse.Merge(m, src)
}
func (m *OpenPlaformLoginResponse) XXX_Size() int {
	return xxx_messageInfo_OpenPlaformLoginResponse.Size(m)
}
func (m *OpenPlaformLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenPlaformLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenPlaformLoginResponse proto.InternalMessageInfo

func (m *OpenPlaformLoginResponse) GetError() int64 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *OpenPlaformLoginResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *OpenPlaformLoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//token解码请求
type TokenDecodeRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenDecodeRequest) Reset()         { *m = TokenDecodeRequest{} }
func (m *TokenDecodeRequest) String() string { return proto.CompactTextString(m) }
func (*TokenDecodeRequest) ProtoMessage()    {}
func (*TokenDecodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{6}
}

func (m *TokenDecodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenDecodeRequest.Unmarshal(m, b)
}
func (m *TokenDecodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenDecodeRequest.Marshal(b, m, deterministic)
}
func (m *TokenDecodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenDecodeRequest.Merge(m, src)
}
func (m *TokenDecodeRequest) XXX_Size() int {
	return xxx_messageInfo_TokenDecodeRequest.Size(m)
}
func (m *TokenDecodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenDecodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenDecodeRequest proto.InternalMessageInfo

func (m *TokenDecodeRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//token解码响应
type TokenDecodeResponse struct {
	//用户基本信息
	UserBaseInfo *UserBaseInfo `protobuf:"bytes,1,opt,name=user_base_info,json=userBaseInfo,proto3" json:"user_base_info,omitempty"`
	//用户在三方的信息
	UserOpenPlatformInfo []*UserOpenPlatformInfo `protobuf:"bytes,2,rep,name=user_open_platform_info,json=userOpenPlatformInfo,proto3" json:"user_open_platform_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TokenDecodeResponse) Reset()         { *m = TokenDecodeResponse{} }
func (m *TokenDecodeResponse) String() string { return proto.CompactTextString(m) }
func (*TokenDecodeResponse) ProtoMessage()    {}
func (*TokenDecodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{7}
}

func (m *TokenDecodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenDecodeResponse.Unmarshal(m, b)
}
func (m *TokenDecodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenDecodeResponse.Marshal(b, m, deterministic)
}
func (m *TokenDecodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenDecodeResponse.Merge(m, src)
}
func (m *TokenDecodeResponse) XXX_Size() int {
	return xxx_messageInfo_TokenDecodeResponse.Size(m)
}
func (m *TokenDecodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenDecodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenDecodeResponse proto.InternalMessageInfo

func (m *TokenDecodeResponse) GetUserBaseInfo() *UserBaseInfo {
	if m != nil {
		return m.UserBaseInfo
	}
	return nil
}

func (m *TokenDecodeResponse) GetUserOpenPlatformInfo() []*UserOpenPlatformInfo {
	if m != nil {
		return m.UserOpenPlatformInfo
	}
	return nil
}

//用户基本信息
type UserBaseInfo struct {
	AccountId            int64    `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBaseInfo) Reset()         { *m = UserBaseInfo{} }
func (m *UserBaseInfo) String() string { return proto.CompactTextString(m) }
func (*UserBaseInfo) ProtoMessage()    {}
func (*UserBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{8}
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

func (m *UserBaseInfo) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
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

//用户开放平台的信息
type UserOpenPlatformInfo struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//账户id
	AccountId int64 `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	//三方类型，比如微信
	TypeId int64 `protobuf:"varint,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	//三方的唯一标识，比如微信的openid
	PlatformId string `protobuf:"bytes,4,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	//同一个三方平台，不同应用，用户的唯一标识
	UnionId              string   `protobuf:"bytes,5,opt,name=union_id,json=unionId,proto3" json:"union_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserOpenPlatformInfo) Reset()         { *m = UserOpenPlatformInfo{} }
func (m *UserOpenPlatformInfo) String() string { return proto.CompactTextString(m) }
func (*UserOpenPlatformInfo) ProtoMessage()    {}
func (*UserOpenPlatformInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{9}
}

func (m *UserOpenPlatformInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOpenPlatformInfo.Unmarshal(m, b)
}
func (m *UserOpenPlatformInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOpenPlatformInfo.Marshal(b, m, deterministic)
}
func (m *UserOpenPlatformInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOpenPlatformInfo.Merge(m, src)
}
func (m *UserOpenPlatformInfo) XXX_Size() int {
	return xxx_messageInfo_UserOpenPlatformInfo.Size(m)
}
func (m *UserOpenPlatformInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOpenPlatformInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserOpenPlatformInfo proto.InternalMessageInfo

func (m *UserOpenPlatformInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserOpenPlatformInfo) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *UserOpenPlatformInfo) GetTypeId() int64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *UserOpenPlatformInfo) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

func (m *UserOpenPlatformInfo) GetUnionId() string {
	if m != nil {
		return m.UnionId
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterAccountEmailRequester)(nil), "account.RegisterAccountEmailRequester")
	proto.RegisterType((*RegisterAccountEmailResponse)(nil), "account.RegisterAccountEmailResponse")
	proto.RegisterType((*LoginByPasswordRequest)(nil), "account.LoginByPasswordRequest")
	proto.RegisterType((*LoginByPasswordResponse)(nil), "account.LoginByPasswordResponse")
	proto.RegisterType((*OpenPlaformLoginRequest)(nil), "account.OpenPlaformLoginRequest")
	proto.RegisterType((*OpenPlaformLoginResponse)(nil), "account.OpenPlaformLoginResponse")
	proto.RegisterType((*TokenDecodeRequest)(nil), "account.TokenDecodeRequest")
	proto.RegisterType((*TokenDecodeResponse)(nil), "account.TokenDecodeResponse")
	proto.RegisterType((*UserBaseInfo)(nil), "account.UserBaseInfo")
	proto.RegisterType((*UserOpenPlatformInfo)(nil), "account.UserOpenPlatformInfo")
}

func init() { proto.RegisterFile("proto/account/account.proto", fileDescriptor_5d492a0187472a3b) }

var fileDescriptor_5d492a0187472a3b = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x8f, 0xd2, 0x4c,
	0x14, 0x7e, 0xfb, 0x01, 0xbc, 0x7b, 0xd8, 0x6c, 0xcc, 0x11, 0x01, 0xd9, 0x25, 0x8b, 0x4d, 0x34,
	0x1b, 0x2f, 0xd6, 0x04, 0x2f, 0xbd, 0x92, 0xe8, 0x05, 0xc6, 0x44, 0xd2, 0xac, 0x17, 0xc6, 0x0f,
	0xd2, 0xa5, 0x03, 0x36, 0xd2, 0x99, 0x3a, 0xd3, 0xc6, 0x70, 0xeb, 0xcf, 0xf0, 0x17, 0xf8, 0x87,
	0xfc, 0x3f, 0x66, 0x3e, 0x5a, 0x0a, 0xb4, 0x78, 0xb3, 0x57, 0xcc, 0x73, 0xce, 0x99, 0xe7, 0x3c,
	0x9c, 0x8f, 0x29, 0x9c, 0x27, 0x9c, 0xa5, 0xec, 0x59, 0xb0, 0x58, 0xb0, 0x8c, 0xa6, 0xf9, 0xef,
	0xb5, 0xb2, 0x62, 0xcb, 0x40, 0x2f, 0x80, 0xa1, 0x4f, 0x56, 0x91, 0x48, 0x09, 0x7f, 0xa9, 0x4d,
	0xaf, 0xe3, 0x20, 0x5a, 0xfb, 0xe4, 0x7b, 0x46, 0xa4, 0x11, 0x3b, 0xd0, 0x20, 0xd2, 0xd2, 0xb7,
	0x46, 0xd6, 0xd5, 0x89, 0xaf, 0x01, 0x22, 0xb8, 0x0b, 0x16, 0x92, 0xbe, 0xad, 0x8c, 0xea, 0x8c,
	0x5d, 0x68, 0x26, 0x81, 0x10, 0x3f, 0xc2, 0xbe, 0xa3, 0xac, 0x06, 0x79, 0x33, 0xb8, 0xa8, 0x4e,
	0x21, 0x12, 0x46, 0x05, 0xc1, 0x3e, 0xb4, 0x08, 0xe7, 0x8a, 0x4e, 0xe6, 0x70, 0xfc, 0x1c, 0x4a,
	0x46, 0xc2, 0x79, 0x2c, 0x56, 0x26, 0x8f, 0x41, 0xde, 0x27, 0xe8, 0xbe, 0x65, 0xab, 0x88, 0x4e,
	0x36, 0x33, 0x99, 0x82, 0xf1, 0xd0, 0xe8, 0xad, 0x51, 0xdb, 0x81, 0x46, 0xf2, 0x95, 0xd1, 0x5c,
	0xae, 0x06, 0xb5, 0x7a, 0x3f, 0x43, 0xef, 0x80, 0xdd, 0x48, 0x95, 0xf4, 0x9c, 0x33, 0x6e, 0x84,
	0x6a, 0x50, 0x27, 0x53, 0x46, 0xa7, 0xec, 0x1b, 0xa1, 0x86, 0x5f, 0x03, 0xef, 0xa7, 0x05, 0xbd,
	0x77, 0x09, 0xa1, 0xb3, 0x75, 0xb0, 0x64, 0x3c, 0x56, 0xa9, 0x72, 0xf9, 0x08, 0x2e, 0x0d, 0x62,
	0x62, 0xd4, 0xab, 0x33, 0xf6, 0xa0, 0x95, 0x6e, 0x12, 0x32, 0x8f, 0x42, 0x45, 0xef, 0xfa, 0x4d,
	0x09, 0xa7, 0x21, 0x5e, 0x42, 0x3b, 0x59, 0x07, 0xa9, 0x24, 0x91, 0x4e, 0x9d, 0x04, 0x72, 0xd3,
	0x34, 0xc4, 0x87, 0xf0, 0x7f, 0x46, 0x23, 0x46, 0xa5, 0xd7, 0x55, 0xde, 0x96, 0xc2, 0xd3, 0xd0,
	0xfb, 0x02, 0xfd, 0x43, 0x0d, 0x77, 0xf8, 0x27, 0x9f, 0x02, 0xde, 0xc8, 0xc3, 0x2b, 0x22, 0x1b,
	0x59, 0xea, 0x8e, 0x8e, 0xb5, 0xca, 0xb1, 0xbf, 0x2d, 0xb8, 0xbf, 0x13, 0x6c, 0x74, 0xbc, 0x80,
	0xb3, 0x4c, 0x10, 0x3e, 0xbf, 0x0d, 0x04, 0x99, 0x47, 0x74, 0xc9, 0xd4, 0xb5, 0xf6, 0xf8, 0xc1,
	0x75, 0x3e, 0xcb, 0xef, 0x05, 0xe1, 0x93, 0x40, 0x90, 0x29, 0x5d, 0x32, 0xff, 0x34, 0x2b, 0x21,
	0xbc, 0x81, 0x9e, 0xba, 0xcc, 0x12, 0x42, 0xe7, 0xdb, 0x32, 0x49, 0x16, 0x7b, 0xe4, 0x5c, 0xb5,
	0xc7, 0xc3, 0x1d, 0x16, 0x53, 0x0c, 0x5d, 0x39, 0xc9, 0xd6, 0xc9, 0x2a, 0xac, 0x5e, 0x0c, 0xa7,
	0xe5, 0x9c, 0x38, 0x04, 0x30, 0x2c, 0xb2, 0xc6, 0xba, 0x5e, 0x27, 0xc6, 0x32, 0x0d, 0xb7, 0xd3,
	0x68, 0x57, 0x4e, 0xa3, 0x53, 0x9e, 0xc6, 0xbc, 0xf5, 0xee, 0xb6, 0xf5, 0xde, 0x2f, 0x0b, 0x3a,
	0x55, 0xea, 0xf0, 0x0c, 0xec, 0x22, 0x9f, 0x1d, 0x85, 0x7b, 0x3a, 0xec, 0x7d, 0x1d, 0xa5, 0x11,
	0x72, 0x94, 0xaf, 0x66, 0x84, 0xdc, 0xa3, 0x23, 0xd4, 0xd8, 0x19, 0xa1, 0xf1, 0x1f, 0x07, 0x5a,
	0x66, 0x9f, 0x31, 0x82, 0xee, 0xde, 0x8a, 0x4f, 0x36, 0x6a, 0xc9, 0xf1, 0x49, 0x51, 0xe6, 0xa3,
	0xcf, 0xcc, 0xe0, 0xf1, 0x3f, 0xe2, 0xf4, 0x4c, 0x78, 0xff, 0xe1, 0x47, 0xe8, 0x98, 0xed, 0x54,
	0x9e, 0x7c, 0x45, 0xf1, 0xb2, 0x20, 0xa8, 0x7e, 0x1a, 0x06, 0xa3, 0xfa, 0x80, 0x0a, 0xf2, 0x99,
	0x6c, 0xca, 0xdd, 0x92, 0x7f, 0x80, 0x7b, 0xfb, 0x3b, 0x87, 0xdb, 0x7b, 0x35, 0x4f, 0xc2, 0xe0,
	0xd1, 0x91, 0x88, 0x82, 0xfa, 0x0d, 0xb4, 0x4b, 0x1b, 0x84, 0xe7, 0xc5, 0x9d, 0xc3, 0x25, 0x1c,
	0x5c, 0x54, 0x3b, 0x73, 0xae, 0xdb, 0xa6, 0xfa, 0x42, 0x3c, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xb2, 0x78, 0x4f, 0x00, 0x40, 0x06, 0x00, 0x00,
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
	//通过email注册一个用户
	RegisterAccountByEmail(ctx context.Context, in *RegisterAccountEmailRequester, opts ...grpc.CallOption) (*RegisterAccountEmailResponse, error)
	//邮箱加密码登录
	LoginByEmailPassword(ctx context.Context, in *LoginByPasswordRequest, opts ...grpc.CallOption) (*LoginByPasswordResponse, error)
	//手机号加密码登录
	LoginByPhonePassword(ctx context.Context, in *LoginByPasswordRequest, opts ...grpc.CallOption) (*LoginByPasswordResponse, error)
	//三方登录
	OpenPlaformLogin(ctx context.Context, in *OpenPlaformLoginRequest, opts ...grpc.CallOption) (*OpenPlaformLoginResponse, error)
	//token解码
	TokenDecode(ctx context.Context, in *TokenDecodeRequest, opts ...grpc.CallOption) (*TokenDecodeResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) RegisterAccountByEmail(ctx context.Context, in *RegisterAccountEmailRequester, opts ...grpc.CallOption) (*RegisterAccountEmailResponse, error) {
	out := new(RegisterAccountEmailResponse)
	err := c.cc.Invoke(ctx, "/account.Account/RegisterAccountByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) LoginByEmailPassword(ctx context.Context, in *LoginByPasswordRequest, opts ...grpc.CallOption) (*LoginByPasswordResponse, error) {
	out := new(LoginByPasswordResponse)
	err := c.cc.Invoke(ctx, "/account.Account/LoginByEmailPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) LoginByPhonePassword(ctx context.Context, in *LoginByPasswordRequest, opts ...grpc.CallOption) (*LoginByPasswordResponse, error) {
	out := new(LoginByPasswordResponse)
	err := c.cc.Invoke(ctx, "/account.Account/LoginByPhonePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) OpenPlaformLogin(ctx context.Context, in *OpenPlaformLoginRequest, opts ...grpc.CallOption) (*OpenPlaformLoginResponse, error) {
	out := new(OpenPlaformLoginResponse)
	err := c.cc.Invoke(ctx, "/account.Account/OpenPlaformLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) TokenDecode(ctx context.Context, in *TokenDecodeRequest, opts ...grpc.CallOption) (*TokenDecodeResponse, error) {
	out := new(TokenDecodeResponse)
	err := c.cc.Invoke(ctx, "/account.Account/TokenDecode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	//通过email注册一个用户
	RegisterAccountByEmail(context.Context, *RegisterAccountEmailRequester) (*RegisterAccountEmailResponse, error)
	//邮箱加密码登录
	LoginByEmailPassword(context.Context, *LoginByPasswordRequest) (*LoginByPasswordResponse, error)
	//手机号加密码登录
	LoginByPhonePassword(context.Context, *LoginByPasswordRequest) (*LoginByPasswordResponse, error)
	//三方登录
	OpenPlaformLogin(context.Context, *OpenPlaformLoginRequest) (*OpenPlaformLoginResponse, error)
	//token解码
	TokenDecode(context.Context, *TokenDecodeRequest) (*TokenDecodeResponse, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) RegisterAccountByEmail(ctx context.Context, req *RegisterAccountEmailRequester) (*RegisterAccountEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccountByEmail not implemented")
}
func (*UnimplementedAccountServer) LoginByEmailPassword(ctx context.Context, req *LoginByPasswordRequest) (*LoginByPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByEmailPassword not implemented")
}
func (*UnimplementedAccountServer) LoginByPhonePassword(ctx context.Context, req *LoginByPasswordRequest) (*LoginByPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByPhonePassword not implemented")
}
func (*UnimplementedAccountServer) OpenPlaformLogin(ctx context.Context, req *OpenPlaformLoginRequest) (*OpenPlaformLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenPlaformLogin not implemented")
}
func (*UnimplementedAccountServer) TokenDecode(ctx context.Context, req *TokenDecodeRequest) (*TokenDecodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenDecode not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_RegisterAccountByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAccountEmailRequester)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).RegisterAccountByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/RegisterAccountByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).RegisterAccountByEmail(ctx, req.(*RegisterAccountEmailRequester))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_LoginByEmailPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).LoginByEmailPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/LoginByEmailPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).LoginByEmailPassword(ctx, req.(*LoginByPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_LoginByPhonePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).LoginByPhonePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/LoginByPhonePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).LoginByPhonePassword(ctx, req.(*LoginByPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_OpenPlaformLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenPlaformLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).OpenPlaformLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/OpenPlaformLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).OpenPlaformLogin(ctx, req.(*OpenPlaformLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_TokenDecode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenDecodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).TokenDecode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/TokenDecode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).TokenDecode(ctx, req.(*TokenDecodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAccountByEmail",
			Handler:    _Account_RegisterAccountByEmail_Handler,
		},
		{
			MethodName: "LoginByEmailPassword",
			Handler:    _Account_LoginByEmailPassword_Handler,
		},
		{
			MethodName: "LoginByPhonePassword",
			Handler:    _Account_LoginByPhonePassword_Handler,
		},
		{
			MethodName: "OpenPlaformLogin",
			Handler:    _Account_OpenPlaformLogin_Handler,
		},
		{
			MethodName: "TokenDecode",
			Handler:    _Account_TokenDecode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account/account.proto",
}
