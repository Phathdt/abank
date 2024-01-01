package repo

import (
	"abank/module/models"
	"context"
)

type SqlStorage interface {
	ListAccount(ctx context.Context, userId int) ([]models.Account, error)
}

type CacheStorage interface {
}

type repo struct {
	store      SqlStorage
	cacheStore CacheStorage
}

func (r *repo) ListAccount(ctx context.Context, userId int) ([]models.Account, error) {
	return r.store.ListAccount(ctx, userId)
}

func NewRepo(store SqlStorage, cacheStore CacheStorage) *repo {
	return &repo{store: store, cacheStore: cacheStore}
}
