package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	account_service "github.com/guapo-organizations/account-service/lib/account"
	"github.com/guapo-organizations/account-service/model"
	"github.com/guapo-organizations/account-service/proto/account"
	ly_jwt "github.com/guapo-organizations/go-micro-secret/jwt"
	"log"
)

//邮箱加密码登录，也就是账号加密码登录
func (this *AccountService) LoginByEmailPassword(ctx context.Context, in *account.LoginByPasswordRequest) (*account.LoginByPasswordResponse, error) {
	//执行登录
	account_model, err := account_service.LoginByEmailPassword(in.Email, in.Passwd)
	if err != nil {
		return nil, err
	}

	log.Println(account_model)
	//生成token,有效期是一个小时
	token_map := this.changeToTokenMapByAccountModel(account_model)
	token, err := ly_jwt.JwtTokenEncode(token_map, 60)

	if err != nil {
		return nil, err
	}
	//返回响应

	return &account.LoginByPasswordResponse{
		Token: token,
	}, nil

}

//手机号加密码登录
func (this *AccountService) LoginByPhonePassword(ctx context.Context, in *account.LoginByPasswordRequest) (*account.LoginByPasswordResponse, error) {
	//执行登录
	account_model, err := account_service.LoginByPhonePassword(in.Phone, in.Passwd)
	if err != nil {
		return nil, err
	}

	log.Println(account_model)
	//生成token,有效期是一个小时
	token_map := this.changeToTokenMapByAccountModel(account_model)
	token, err := ly_jwt.JwtTokenEncode(token_map, 60)

	if err != nil {
		return nil, err
	}
	//返回响应

	return &account.LoginByPasswordResponse{
		Token: token,
	}, nil

}

//模型转化为token所需要的map
func (this *AccountService) changeToTokenMapByAccountModel(account_model *model.Account) jwt.MapClaims {
	token_map := jwt.MapClaims{}
	token_map["Phone"] = account_model.Phone
	token_map["Email"] = account_model.Email
	token_map["Name"] = account_model.Name
	token_map["Account_Platform"] = account_model.Account_Platform
	return token_map
}
