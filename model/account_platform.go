package model

import "github.com/jinzhu/gorm"

type AccountPlatform struct {
	gorm.Model
	AccountId  uint  `gorm:"column:account_id"`
	TypeId     uint  `gorm:"column:type_id"`
	PlatformId string `gorm:"column:platform_id"`
	UnionId    string `gorm:"column:union_id"`
	//站在accountplatform的角度上，ACCOUNT 的Id是外键，accountPlatform的accountid是主键
	Account Account `gorm:"foreignkey:ID;AssociationForeignKey:AccountId"`
}

//表的名字
func (a AccountPlatform) TableName() string {
	return "account_platform"
}
