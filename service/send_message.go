package service

import (
	"context"
	"fmt"
	"github.com/guapo-organizations/account-service/lib/send-message"
	"github.com/guapo-organizations/account-service/proto/account"
)

//发送信息

//有效发送验证码
func (this *AccountService) SendEmailCode(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.SendMesaageResponse, error) {

	if in.Email == "" {
		return nil, fmt.Errorf("你这个什么emil？？")
	}

	msg_service := send_message.SendMessageService{}

	_, err := msg_service.SendEmailCode(in.Email)
	if err != nil {
		return nil, err
	}

	//能走到这里就是发送成功了
	return &account.SendMesaageResponse{
		Status: 0,
		Errmsg: "",
	}, nil

}

//校验邮箱验证码
func (this *AccountService) ValidateEmailCode(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.ValidataCodeResponse, error) {

	if in.Email == "" {
		return nil, fmt.Errorf("邮箱在哪里？")
	}

	msg_service := send_message.SendMessageService{}

	validate_resule, err := msg_service.CheckEmailCode(in.Email)

	if err != nil {
		return nil, err
	}
	response := new(account.ValidataCodeResponse)

	if validate_resule {
		response.Status = 0
		response.Errmsg = ""
	} else {
		response.Status = 1
		response.Errmsg = "亲，没有验证码不对哦"
	}
	
	return response, nil
}
