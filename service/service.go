package service

import (
	"context"
	"fmt"
	"github.com/guapo-organizations/go-micro-secret/database"
	"zldz/account-service/model"
	"zldz/account-service/proto/account"
	"zldz/go-micro-secret/help"
)

type AccountService struct {
}

//实现grpc服务接口
func (this *AccountService) RegisterAccountByEmailOrPhone(ctx context.Context, in *account.EmailOrPhoneRequest) (*account.UserBaseInfo, error) {
	response := new(account.UserBaseInfo)
	account_model := new(model.Account)
	
	if help.VerifyEmailFormat(in.Email) {
		//邮箱注册
		account_model.Email = in.Email
	} else if help.VerifyMobileFormat(in.Phone) {
		//手机号注册
		account_model.Phone = in.Phone
	} else {
		return nil, fmt.Errorf("手机号和邮箱必填一个")
	}

	//从连接池中获取一个db连接,不需要释放它，sync.pool会帮你做处理
	//创建一个用户
	db := database.GetMysqlDB()
	db.Create(account_model)

	//响应
	response.Phone = account_model.Phone
	response.Email = account_model.Email
	response.Id = account_model.Id
	response.Name = account_model.Name

	return response, nil
}
