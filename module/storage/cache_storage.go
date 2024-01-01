package storage

import (
	"abank/module/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type cacheStore struct {
	client *redis.Client
}

func (s *cacheStore) GetUser(ctx context.Context, userId int) (*models.User, error) {
	key := fmt.Sprintf("/users/%d", userId)

	result, err := s.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, errors.New("not found")
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	var user models.User
	if err = json.Unmarshal([]byte(result), &user); err != nil {
		return nil, errors.WithStack(err)
	}

	return &user, nil
}

func (s *cacheStore) SetUser(ctx context.Context, data *models.User) error {
	key := fmt.Sprintf("/users/%d", data.Id)

	bytes, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = s.client.Set(ctx, key, bytes, time.Minute*60).Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *cacheStore) GetAccount(ctx context.Context, accountId int) (*models.Account, error) {
	key := fmt.Sprintf("/accounts/%d", accountId)

	result, err := s.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, errors.New("not found")
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	var account models.Account
	if err = json.Unmarshal([]byte(result), &account); err != nil {
		return nil, errors.WithStack(err)
	}

	return &account, nil
}

func (s *cacheStore) SetAccount(ctx context.Context, data *models.Account) error {
	key := fmt.Sprintf("/accounts/%d", data.Id)

	bytes, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = s.client.Set(ctx, key, bytes, time.Minute*60).Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func NewCacheStore(client *redis.Client) *cacheStore {
	return &cacheStore{client: client}
}
