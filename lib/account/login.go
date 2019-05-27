package account

import (
	"fmt"
	"github.com/guapo-organizations/account-service/model"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/go-micro-secret/help"
	"strings"
)

//emial加密码登录
func LoginByEmailPasswod(email, passwod string) (account *model.Account, err error) {

	if !help.VerifyEmailFormat(email) {
		return nil, fmt.Errorf("email格式不对")
	}

	//查询用户
	db := database.GetMysqlDB()
	//执行一下new，保证指针不是nil
	account_model := new(model.Account)
	db.Preload("Account_Platform").Where(&model.Account{
		Email: email,
	}).First(account_model)

	//找不到
	if account_model.ID == 0 {
		return nil, fmt.Errorf("找不到邮箱为:%s的账号", email)
	}

	if !strings.EqualFold(account_model.Passwd, passwod) {
		return nil, fmt.Errorf("密码不正确")
	}

	return account_model, nil
}
