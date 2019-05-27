package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/guapo-organizations/account-service/model"
	"github.com/guapo-organizations/account-service/proto/account"
	ly_jwt "github.com/guapo-organizations/go-micro-secret/jwt"
	"time"
)

//模型转化为token所需要的map,token编码
func (this *AccountService) TokenEncode(account_model *model.Account, minute time.Duration) (string, error) {
	token_map := jwt.MapClaims{}
	token_map["Phone"] = account_model.Phone
	token_map["Email"] = account_model.Email
	token_map["Name"] = account_model.Name
	token_map["Account_Platform"] = account_model.Account_Platform
	token, err := ly_jwt.JwtTokenEncode(token_map, 60)
	if err != nil {
		return "", err
	}
	return token, nil
}

//对用户登录的token进行解码
func (this *AccountService) TokenDecode(ctx context.Context, in *account.TokenDecodeRequest) (*account.TokenDecodeResponse, error) {
	return nil, nil
}
