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
	GetAccount(ctx context.Context, accountId int) (*models.Account, error)
}

type CacheStorage interface {
	GetAccount(ctx context.Context, accountId int) (*models.Account, error)
	SetAccount(ctx context.Context, data *models.Account) error
}

type repo struct {
	store      SqlStorage
	cacheStore CacheStorage
}

func (r *repo) GetAccount(ctx context.Context, accountId int) (*models.Account, error) {
	account, err := r.cacheStore.GetAccount(ctx, accountId)
	if err == nil {
		return account, nil
	}

	account, err = r.store.GetAccount(ctx, accountId)
	if err != nil {
		return nil, err
	}

	_ = r.cacheStore.SetAccount(ctx, account)

	return account, nil
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
