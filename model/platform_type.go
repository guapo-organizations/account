package model

import "github.com/jinzhu/gorm"

type PlatformType struct {
	gorm.Model
	Type      int64  `gorm:"column:type"`
	AppId     string  `gorm:"column:app_id"`
	AppSecret string `gorm:"column:app_secret"`
}

//表的名字
func (a PlatformType) TableName() string {
	return "platform_type"
}
