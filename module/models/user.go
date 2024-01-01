package models

import "github.com/phathdt/service-context/core"

type User struct {
	core.SQLModel
	Name       string `json:"name" gorm:"column:name"`
	AccountIds []int  `json:"account_ids" gorm:"-"`
}

func (User) TableName() string {
	return "users"
}
