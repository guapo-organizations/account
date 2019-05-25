package service

import (
	"context"
	"fmt"
	account_service "github.com/guapo-organizations/account-service/lib/account"
	"github.com/guapo-organizations/account-service/proto/account"
	"github.com/guapo-organizations/go-micro-secret/consul"
	my_grpc_connet "github.com/guapo-organizations/go-micro-secret/grpc"
	sms_service "github.com/guapo-organizations/sms-service/proto/sms"
)

type AccountService struct {
}

//邮箱注册
func (this *AccountService) RegisterAccountByEmail(ctx context.Context, in *account.RegisterAccountEmailRequester) (*account.RegisterAccountEmailResponse, error) {

	consul_config, err := consul.GetConfig()

	if err != nil {
		return nil, err
	}
	//调用sms服务，查看验证吗是否正确
	conn, err := my_grpc_connet.GetGrpcConnet("zldz.sms", "sms", consul_config)

	if err != nil {
		return nil, err
	}
	//关闭资源
	defer conn.Close()

	sms_client := sms_service.NewSmsClient(conn)

	sms_response, err := sms_client.ValidateCode(ctx, &sms_service.ValidateCodeRequest{
		Key:         in.Email,
		Code:        in.Code,
		PublishMode: "email",
		PublishType: "code",
	})

	if err != nil {
		return nil, err
	}

	//验证码不正确
	if !sms_response.Result {
		return nil, fmt.Errorf("验证码不正确")
	}

	//插入数据库
	_, err = account_service.RegisterAccountByEmail(in.Email, in.Passwd)

	if err != nil {
		return nil, err
	}

	//响应
	response := new(account.RegisterAccountEmailResponse)
	return response, nil
}
