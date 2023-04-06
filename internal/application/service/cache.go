package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	redis "task3_4/user-management/internal/adapter/db/redis"
	"time"
)

type Cache struct {
	RedisRepository redis.RedisRepositoryInterface
}

func NewCache(r redis.RedisRepositoryInterface) *Cache {
	return &Cache{r}
}

func (au *Cache) AddToCache(key string, data []byte, expiration time.Duration) error {
	return au.RedisRepository.Save(key, data, expiration)
}

func (au *Cache) GetFromCache(key string) ([]byte, error) {
	data, err := au.RedisRepository.Get(key)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return data, nil
}
