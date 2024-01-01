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

func (s *sqlStore) GetAccount(ctx context.Context, accountId int) (*models.Account, error) {
	var data models.Account

	db := s.db.WithContext(ctx)

	result := db.Where("id = ?", accountId).Limit(1).Find(&data)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, errors.WithStack(errors.New("not found"))
	}

	return &data, nil
}

func (s *sqlStore) GetListAccount(ctx context.Context, cond map[string]interface{}) ([]models.Account, error) {
	var data []models.Account

	db := s.db.WithContext(ctx).Table(models.Account{}.TableName()).Order("id ASC")

	if err := db.Where(cond).
		Select("*").
		Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}

func (s *sqlStore) GetUser(ctx context.Context, userId int) (*models.User, error) {
	var data models.User

	db := s.db.WithContext(ctx)

	result := db.Where("id = ?", userId).Limit(1).Find(&data)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, errors.WithStack(errors.New("not found"))
	}

	return &data, nil
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
