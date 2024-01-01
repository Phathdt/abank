package repo

import (
	"abank/module/models"
	"context"
	"github.com/samber/lo"
)

type SqlStorage interface {
	ListAccount(ctx context.Context, userId int) ([]models.Account, error)
	GetUser(ctx context.Context, userId int) (*models.User, error)
	GetListAccount(ctx context.Context, cond map[string]interface{}) ([]models.Account, error)
}

type CacheStorage interface {
}

type repo struct {
	store      SqlStorage
	cacheStore CacheStorage
}

func (r *repo) GetUser(ctx context.Context, userId int) (*models.User, error) {
	user, err := r.store.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	accounts, err := r.store.GetListAccount(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}

	user.AccountIds = lo.Map(accounts, func(item models.Account, index int) int {
		return item.Id
	})

	return user, nil
}

func (r *repo) ListAccount(ctx context.Context, userId int) ([]models.Account, error) {
	return r.store.ListAccount(ctx, userId)
}

func NewRepo(store SqlStorage, cacheStore CacheStorage) *repo {
	return &repo{store: store, cacheStore: cacheStore}
}
