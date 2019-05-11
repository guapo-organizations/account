package register

import (
	"fmt"
	"github.com/guapo-organizations/account-service/model"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/go-micro-secret/help"
)

type RegisterService struct {
}

//邮箱注册
func (this RegisterService) RegisterAccountByEmail(email, passwd string) (*model.Account, error) {

	if !help.VerifyEmailFormat(email) {
		return nil, fmt.Errorf("email格式不对")
	}

	db := database.GetMysqlDB()
	//执行一下new，保证指针不是nil
	account_model := new(model.Account)
	db.Where(&model.Account{
		Email: email,
	}).First(account_model)

	//找到直接返回
	if account_model.ID != 0 {
		return account_model, nil
	}

	//email注册
	account_model.Email = email
	account_model.Passwd = passwd
	db.Create(account_model)

	return account_model, nil
}

//手机号注册
func (this RegisterService) RegisterAccountByPhone(phone, passwd string) (*model.Account, error) {

	if !help.VerifyMobileFormat(phone) {
		return nil, fmt.Errorf("手机号格式不对")
	}

	db := database.GetMysqlDB()
	account_model := new(model.Account)

	db.Where(&model.Account{
		Phone: phone,
	}).First(account_model)

	//找到直接返回
	if account_model.ID != 0 {
		return account_model, nil
	}

	//手机注册
	account_model.Phone = phone
	account_model.Passwd = passwd
	db.Create(account_model)

	return account_model, nil
}
