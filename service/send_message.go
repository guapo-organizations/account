package service

import (
	"context"
	"github.com/guapo-organizations/account-service/proto/account"
)

//发送信息

//有效发送验证码
func (this *AccountService) SendEmailCode(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.SendMesaageResponse, error) {
	return nil, nil
}
