package models

import (
	"time"
)

type Account struct {
	Id        int        `json:"id" gorm:"column:id;" db:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"  db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"  db:"updated_at"`
	Name      string     `json:"name" gorm:"column:name"`
	UserId    int        `json:"user_id" gorm:"column:user_id"`
	Balance   int64      `json:"balance" gorm:"column:balance"`
}

func (Account) TableName() string {
	return "accounts"
}
