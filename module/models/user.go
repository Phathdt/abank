package models

import (
	"time"
)

type User struct {
	Id         int        `json:"id" gorm:"column:id;" db:"id"`
	CreatedAt  *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"  db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"  db:"updated_at"`
	Name       string     `json:"name" gorm:"column:name"`
	AccountIds []int      `json:"account_ids" gorm:"-"`
}

func (User) TableName() string {
	return "users"
}
