package service

import (
	"time"
)

type CacheInterface interface {
	AddToCache(key string, data []byte, expiration time.Duration) error
	GetFromCache(key string) ([]byte, error)
}
