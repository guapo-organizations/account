package service

import (
	"context"
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

	conn, err := my_grpc_connet.GetGrpcConnet("zldz.sms", "sms", consul_config)

	if err != nil {
		return nil, err
	}

	sms_service.NewSmsClient(conn)

	//插入数据库
	_, err = account_service.RegisterAccountByEmail(in.Email, in.Passwd)

	if err != nil {
		return nil, err
	}

	//响应
	response := new(account.RegisterAccountEmailResponse)
	return response, nil
}
