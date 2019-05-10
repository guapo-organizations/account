package lib

import (
	"fmt"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/go-micro-secret/help"
	"zldz/account-service/model"
)

type RegisterService struct {
}

//邮箱注册
func (this RegisterService) RegisterAccountByEmail(email string) (account_model *model.Account, err error) {

	if !help.VerifyEmailFormat(email) {
		return nil, fmt.Errorf("email格式不对")
	}

	db := database.GetMysqlDB()

	db.Where(&model.Account{
		Email: email,
	}).Find(account_model)

	//找到直接返回
	if account_model.Id != 0 {
		return account_model, nil
	}

	//email注册
	account_model.Email = email
	db.Create(account_model)
	return account_model, nil
}

//手机号注册
func (this RegisterService) RegisterAccountByPhone(phone string) (account_model *model.Account, err error) {

	if !help.VerifyMobileFormat(phone) {
		return nil, fmt.Errorf("手机号格式不对")
	}
	db := database.GetMysqlDB()

	db.Where(&model.Account{
		Phone: phone,
	}).Find(account_model)

	//找到直接返回
	if account_model.Id != 0 {
		return account_model, nil
	}

	//手机注册
	account_model.Phone = phone
	db.Create(account_model)
	return account_model, nil
}
