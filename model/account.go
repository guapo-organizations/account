package model

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Passwd string `gorm:"column:passwd"`
	Email  string `gorm:"column:email"`
	Phone  string `gorm:"column:phone"`
	Name   string `gorm:"column:name"`
	//外键定义，模型的属性,不是表的字段,会自动进行关联
	Account_Platform []AccountPlatform `gorm:"foreignkey:AccountId"`
}

//表的名字
func (a Account) TableName() string {
	return "account"
}
