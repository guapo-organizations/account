package service

import (
	"context"
	"fmt"
	account_service "github.com/guapo-organizations/account-service/lib/account"
	"github.com/guapo-organizations/account-service/proto/account"
)

//实现grpc服务接口

//用户注册登录

//邮箱或者手机号注册
func (this *AccountService) RegisterAccountByEmailToken(ctx context.Context, in *account.RegisterAccountInfo) (*account.UserBaseInfo, error) {

	//邮箱注册
	if in.Email != "" {
		account_model, err := account_service.RegisterAccountByEmail(in.Email, in.Passwd)
		//可能是格式出错
		if err != nil {
			return nil, err
		}
		return this.AccountModelChangeUserBaseInfo(account_model), nil
	}

	//手机号注册
	if in.Phone != "" {
		account_model, err := account_service.RegisterAccountByPhone(in.Phone, in.Passwd)
		//可能是格式出错
		if err != nil {
			return nil, err
		}

		return this.AccountModelChangeUserBaseInfo(account_model), nil
	}

	return nil, fmt.Errorf("你想干嘛？你怎么不传参数？")
}

func (this *AccountService) RegisterAccountByPhoneToken(ctx context.Context, in *account.RegisterAccountInfo) (*account.UserBaseInfo, error) {
	return nil, nil
}

