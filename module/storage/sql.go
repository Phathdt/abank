package storage

import (
	"abank/module/models"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func (s *sqlStore) ListAccount(ctx context.Context, userId int) ([]models.Account, error) {
	var accounts []models.Account

	db := s.db.WithContext(ctx).Table(models.Account{}.TableName())

	if err := db.Where("user_id = ?", userId).
		Scan(&accounts).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return accounts, nil
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
