package service

import (
	"zldz/account-service/model"
	"zldz/account-service/proto/account"
)

type AccountService struct {
}

//模型转化为响应
func (this *AccountService) AccountModelChangeUserBaseInfo(account_model *model.Account) *account.UserBaseInfo {
	response := new(account.UserBaseInfo)
	response.Phone = account_model.Phone
	response.Email = account_model.Email
	response.Id = int64(account_model.ID)
	response.Name = account_model.Name
	return response
}
