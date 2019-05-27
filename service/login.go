package service

import (
	"context"
	account_service "github.com/guapo-organizations/account-service/lib/account"
	"github.com/guapo-organizations/account-service/proto/account"
	"log"
)

//邮箱加密码登录，也就是账号加密码登录
func (this *AccountService) LoginByEmailPassword(ctx context.Context, in *account.LoginByPasswordRequest) (*account.LoginByPasswordResponse, error) {
	//执行登录
	account_model, err := account_service.LoginByEmailPassword(in.Email, in.Passwd)
	if err != nil {
		return nil, err
	}

	log.Println(account_model)
	//生成token,有效期是一个小时
	token, err := this.TokenEncode(account_model, 60)

	if err != nil {
		return nil, err
	}
	//返回响应

	return &account.LoginByPasswordResponse{
		Token: token,
	}, nil

}

//手机号加密码登录
func (this *AccountService) LoginByPhonePassword(ctx context.Context, in *account.LoginByPasswordRequest) (*account.LoginByPasswordResponse, error) {
	//执行登录
	account_model, err := account_service.LoginByPhonePassword(in.Phone, in.Passwd)
	if err != nil {
		return nil, err
	}

	log.Println(account_model)
	//生成token,有效期是一个小时
	token, err := this.TokenEncode(account_model, 60)

	if err != nil {
		return nil, err
	}
	//返回响应

	return &account.LoginByPasswordResponse{
		Token: token,
	}, nil

}
