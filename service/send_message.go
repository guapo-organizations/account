package service

import (
	"context"
	"fmt"
	"github.com/guapo-organizations/account-service/proto/account"
	"github.com/guapo-organizations/go-micro-secret/help"
)

//发送信息

//有效发送验证码
func (this *AccountService) SendEmailCode(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.SendMesaageResponse, error) {

	if !help.VerifyEmailFormat(in.Email){
		return nil,fmt.Errorf("邮箱格式不对")
	}

	return nil, nil
}
