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
	token_map["Id"] = account_model.ID
	token_map["Account_Platform"] = account_model.Account_Platform
	token, err := ly_jwt.JwtTokenEncode(token_map, 60)
	if err != nil {
		return "", err
	}
	return token, nil
}

//对用户登录的token进行解码
func (this *AccountService) TokenDecode(ctx context.Context, in *account.TokenDecodeRequest) (*account.TokenDecodeResponse, error) {
	decode_map, err := ly_jwt.JwtTokenDecode(in.Token)

	if err != nil {
		return nil, err
	}
	//构建user_base_info
	user_base_info := new(account.UserBaseInfo)
	user_base_info.Name = decode_map["Name"].(string)
	user_base_info.Phone = decode_map["Phone"].(string)
	user_base_info.AccountId = int64(decode_map["Id"].(float64))
	user_base_info.Email = decode_map["Email"].(string)

	var user_open_platform_info_list []*account.UserOpenPlatformInfo
	//构建user_open_platform_info
	account_platform := decode_map["Account_Platform"].([]interface{})

	for _, value := range account_platform {

		platform_map := value.(map[string]interface{})
		user_open_platform_info_list = append(user_open_platform_info_list, &account.UserOpenPlatformInfo{
			Id:         int64(platform_map["Id"].(float64)),
			AccountId:  int64(platform_map["AccountId"].(float64)),
			Type:       int64(platform_map["Type"].(float64)),
			PlatformId: platform_map["PlatformId"].(string),
		})

	}

	response := new(account.TokenDecodeResponse)
	response.UserBaseInfo = user_base_info
	response.UserOpenPlatformInfo = user_open_platform_info_list
	return response, nil
}
