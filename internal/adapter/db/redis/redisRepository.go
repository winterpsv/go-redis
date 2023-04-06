package redisRepository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository struct {
	client *redis.Client
	con    context.Context
}

func NewRedisRepository(db *redis.Client) *RedisRepository {
	return &RedisRepository{client: db, con: context.Background()}
}

func (r *RedisRepository) Get(key string) ([]byte, error) {
	data, err := r.client.Get(r.con, key).Bytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *RedisRepository) Save(key string, data []byte, expiration time.Duration) error {
	return r.client.Set(r.con, key, data, expiration).Err()
}

func (r *RedisRepository) Delete(key string) error {
	return r.client.Del(r.con, key).Err()
}
