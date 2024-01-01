package storage

import "github.com/redis/go-redis/v9"

type cacheStore struct {
	client *redis.Client
}

func NewCacheStore(client *redis.Client) *cacheStore {
	return &cacheStore{client: client}
}
