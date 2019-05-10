package model

import "github.com/jinzhu/gorm"

type AccountPlatform struct {
	gorm.Model
	Id         int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT;column:id"`
	AccountId  int64  `gorm:"column:account_id"`
	Type       int64  `gorm:"column:type"`
	PlatformId string `gorm:"column:platform_id"`
}
