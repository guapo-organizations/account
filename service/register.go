package service

import (
	"context"
	"fmt"
	"github.com/guapo-organizations/account-service/lib"
	"github.com/guapo-organizations/account-service/proto/account"
)

//实现grpc服务接口

//邮箱或者手机号注册
func (this *AccountService) RegisterAccountByEmailOrPhone(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.UserBaseInfo, error) {
	register_service := lib.RegisterService{}

	//邮箱注册
	if in.Email != "" {
		account_model, err := register_service.RegisterAccountByEmail(in.Email)
		//可能是格式出错
		if err != nil {
			return nil, err
		}
		return this.AccountModelChangeUserBaseInfo(account_model), nil
	}

	//手机号注册
	if in.Phone != "" {
		account_model, err := register_service.RegisterAccountByPhone(in.Phone)
		//可能是格式出错
		if err != nil {
			return nil, err
		}

		return this.AccountModelChangeUserBaseInfo(account_model), nil
	}

	return nil, fmt.Errorf("你想干嘛？你怎么不传参数？")
}

//第三方注册
func (this *AccountService) RegisterAccountByPlatform(ctx context.Context, in *account.PlatformRequest) (*account.UserPlatformInfo, error) {
	return nil, nil
}
