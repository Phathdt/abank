package models

import "github.com/phathdt/service-context/core"

type Account struct {
	core.SQLModel
	Name    string `json:"name" gorm:"column:name"`
	UserId  int    `json:"user_id" gorm:"column:user_id"`
	Balance int64  `json:"balance" gorm:"column:balance"`
}

func (Account) TableName() string {
	return "accounts"
}
