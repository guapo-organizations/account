package account

import (
	"fmt"
	"github.com/guapo-organizations/account-service/model"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/go-micro-secret/help"
	"github.com/jinzhu/gorm"
	"strings"
)

//emial加密码登录
func LoginByEmailPassword(email, password string) (*model.Account, error) {

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
func LoginByPhonePassword(phone, password string) (*model.Account, error) {
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

//三方登录
func LoginByOpenPlatform(name, platform_id, union_id string, type_id uint) (*model.Account, error) {

	//查询是否存在三方登录的信息
	db := database.GetMysqlDB()
	account_type_model := new(model.PlatformType)
	db.Where(&model.PlatformType{
		Model: gorm.Model{
			ID: type_id,
		},
	}).First(account_type_model)

	if account_type_model.ID == 0 {
		return nil, fmt.Errorf("找不到三方平台类型为%d的信息", type_id)
	}

	//查询该用户是否已经注册
	account_platform := new(model.AccountPlatform)
	db.Where(&model.AccountPlatform{
		PlatformId: platform_id,
	}).First(account_platform)

	//已经注册过了,直接返回
	if account_platform.ID != 0 {
		account_model := new(model.Account)
		db.Model(account_platform).Related(account_model)
		return account_model, nil
	}

	//第一次登录，先注册
	db.Begin()
	//注册账户
	account_model := new(model.Account)
	account_model.Name = name
	db.Create(account_model)
	//注册三方
	account_platform.AccountId = account_model.ID
	account_platform.TypeId = type_id
	account_platform.PlatformId = platform_id
	account_platform.UnionId = union_id
	db.Create(account_platform)
	db.Commit()

	return account_model, nil
}
