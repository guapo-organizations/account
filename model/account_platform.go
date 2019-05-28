package model

import "github.com/jinzhu/gorm"

type AccountPlatform struct {
	gorm.Model
	Id         int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT;column:id"`
	AccountId  int64  `gorm:"column:account_id"`
	TypeId       int64  `gorm:"column:type_id"`
	PlatformId string `gorm:"column:platform_id"`
}

//表的名字
func (a AccountPlatform) TableName() string {
	return "account_platform"
}
