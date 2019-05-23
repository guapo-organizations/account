package model

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Passwd string `gorm:"column:passwd"`
	Email  string `gorm:"column:email"`
	Phone  string `gorm:"column:phone"`
	Name   string `gorm:"column:name"`
}

//表的名字
func (a Account) TableName() string {
	return "account"
}
