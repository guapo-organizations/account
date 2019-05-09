package service

import (
	"context"
	"fmt"
	"regexp"
	"zldz/account-service/proto/account"
)

type AccountService struct {
}

//实现grpc服务接口
func (this *AccountService) RegisterAccountByEmailOrPhone(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.UserBaseInfo, error) {
	//响应
	response := new(account.UserBaseInfo)

	if this.VerifyEmailFormat(in.Email) {
		//邮箱注册
	} else if this.VerifyMobileFormat(in.Phone) {
		//手机号注册
	} else {
		return nil, fmt.Errorf("手机号和邮箱必填一个")
	}

	response.Account = "测试账号"
	
	return response, nil
}

//验证是否为邮箱
func (this *AccountService) VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//验证是否为手机号
func (this *AccountService) VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
