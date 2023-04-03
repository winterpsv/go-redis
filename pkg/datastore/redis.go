package datastore

import (
	"github.com/go-redis/redis"
	"task3_4/user-management/internal/infrastructure/config"
)

func NewRedisDB(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr, // host:port of the redis server
		Password: cfg.RedisPass, // no password set
		DB:       cfg.RedisDB,   // use default DB
	})
}
