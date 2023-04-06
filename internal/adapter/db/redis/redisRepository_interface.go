package redisRepository

import "time"

type RedisRepositoryInterface interface {
	Get(key string) ([]byte, error)
	Save(key string, data []byte, expiration time.Duration) error
	Delete(key string) error
}
