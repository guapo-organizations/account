package service

import (
	"context"
	"fmt"
	account_service "github.com/guapo-organizations/account-service/lib/account"
	"github.com/guapo-organizations/account-service/lib/message"
	"github.com/guapo-organizations/account-service/proto/account"
)

//有效发送验证码
func (this *AccountService) SendEmailCode(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.SendMesaageResponse, error) {

	if in.Email == "" {
		return nil, fmt.Errorf("你这个什么emil？？")
	}

	//发送emil
	_, err := message.SendEmailCode(in.Email)
	if err != nil {
		return nil, err
	}

	//能走到这里就是发送成功了
	return &account.SendMesaageResponse{
		Status: 0,
		Errmsg: "",
		Token: &account.Token{
			EmailToken: account_service.GetEmailToken(in.Email),
		},
	}, nil

}

//校验邮箱验证码
func (this *AccountService) ValidateEmailCode(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.ValidataCodeResponse, error) {

	if in.Email == "" {
		return nil, fmt.Errorf("邮箱在哪里？")
	}

	//得到验证结果
	validate_resule, err := message.CheckEmailCode(in.Email)

	if err != nil {
		return nil, err
	}

	response := new(account.ValidataCodeResponse)

	if validate_resule {
		response.Status = 0
		response.Errmsg = ""
	} else {
		response.Status = 1
		response.Errmsg = "亲，验证码不对哦"
	}

	return response, nil
}
