package account

import (
	"fmt"
	"github.com/guapo-organizations/account-service/model"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/go-micro-secret/help"
	"strings"
)

//emial加密码登录
func LoginByEmailPassword(email, password string) (account *model.Account, err error) {

	if !help.VerifyEmailFormat(email) {
		return nil, fmt.Errorf("email格式不对")
	}

	//查询用户
	db := database.GetMysqlDB()
	//执行一下new，保证指针不是nil
	account_model := new(model.Account)
	//属性,preload是加载关联的模型，where为什么传的是指针，如果传的是结构体，则会有默认值，也就是golang的临界值，orm官网说的，临界值时golang语言的特点
	db.Preload("Account_Platform").Where(&model.Account{
		Email: email,
	}).First(account_model)

	//找不到
	if account_model.ID == 0 {
		return nil, fmt.Errorf("找不到邮箱为:%s的账号", email)
	}

	if !strings.EqualFold(account_model.Passwd, password) {
		return nil, fmt.Errorf("密码不正确")
	}

	return account_model, nil
}

//手机号加密码登录
func LoginByPhonePassword(phone, password string) (account *model.Account, err error) {
	if !help.VerifyMobileFormat(phone) {
		return nil, fmt.Errorf("手机号格式不对")
	}

	//查询用户
	db := database.GetMysqlDB()
	account_model := new(model.Account)
	db.Preload("Account_Platform").Where(&model.Account{
		Phone: phone,
	}).First(account_model)

	//找不到
	if account_model.ID == 0 {
		return nil, fmt.Errorf("找不到手机号为:%s的账号", phone)
	}

	if !strings.EqualFold(account_model.Passwd, password) {
		return nil, fmt.Errorf("密码不正确")
	}

	return account_model, nil
}
