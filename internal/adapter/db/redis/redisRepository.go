package redis

import (
	"context"
	"github.com/go-redis/redis"
)

type RedisRepository struct {
	client *redis.Client
	con    context.Context
}

func NewRedisRepository(db *redis.Client) *RedisRepository {
	return &RedisRepository{client: db, con: context.Background()}
}
