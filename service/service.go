package service

import (
	"context"
	"fmt"
	"zldz/account-service/proto/account"
	"zldz/go-micro-secret/help"
)

type AccountService struct {
}

//实现grpc服务接口
func (this *AccountService) RegisterAccountByEmailOrPhone(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.UserBaseInfo, error) {
	//响应
	response := new(account.UserBaseInfo)

	if help.VerifyEmailFormat(in.Email) {
		//邮箱注册
	} else if help.VerifyMobileFormat(in.Phone) {
		//手机号注册
	} else {
		return nil, fmt.Errorf("手机号和邮箱必填一个")
	}

	response.Account = "测试账号"

	return response, nil
}
