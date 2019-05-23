package service

import (
	"context"
	account_service "github.com/guapo-organizations/account-service/lib/account"
	"github.com/guapo-organizations/account-service/proto/account"
)

//实现grpc服务接口

//用户注册登录

//邮箱或者手机号注册
func (this *AccountService) RegisterAccountByEmailToken(ctx context.Context, in *account.RegisterAccountInfo) (*account.UserBaseInfo, error) {

	email, err := account_service.GetEmailFormCache(in.Token.EmailToken)
	if err != nil {
		return nil, err
	}
	//邮箱注册
	account_model, err := account_service.RegisterAccountByEmail(email, in.Passwd)

	if err != nil {
		return nil, err
	}

	return this.AccountModelChangeUserBaseInfo(account_model), nil

}
