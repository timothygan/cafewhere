package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/timothygan/cafewhere/backend/internal/models"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, json, expiration).Err()
}

func (r *RedisRepository) Get(ctx context.Context, key string) ([]*models.Cafe, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var shop []*models.Cafe
	err = json.Unmarshal([]byte(val), &shop)
	if err != nil {
		return nil, err
	}
	return shop, nil
}
