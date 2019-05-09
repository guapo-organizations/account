package service

import (
	"context"
	"zldz/account-service/proto/account"
)

type AccountService struct {
}

//实现grpc服务接口
func (this *AccountService) RegisterAccountByEmailOrPhone(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.UserBaseInfo, error) {
	return nil, nil
}
