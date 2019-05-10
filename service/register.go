package service

import (
	"context"
	"fmt"
	"github.com/guapo-organizations/account-service/lib"
	"github.com/guapo-organizations/account-service/proto/account"
)

//实现grpc服务接口
func (this *AccountService) RegisterAccountByEmailOrPhone(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.UserBaseInfo, error) {
	register_service := lib.RegisterService{}

	//邮箱注册
	account_model, err := register_service.RegisterAccountByEmail(in.Email)
	if err == nil {
		return this.AccountModelChangeUserBaseInfo(account_model), nil
	}

	//手机号注册
	account_model, err = register_service.RegisterAccountByPhone(in.Phone)
	if err == nil {
		return this.AccountModelChangeUserBaseInfo(account_model), nil
	}

	return nil, fmt.Errorf("邮箱或者手机号注册失败")
}
