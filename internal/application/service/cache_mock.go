package service

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type MockCache struct {
	mock.Mock
}

func (m *MockCache) AddToCache(key string, data []byte, expiration time.Duration) error {
	args := m.Called(key, data, expiration)
	return args.Error(0)
}

func (m *MockCache) GetFromCache(key string) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}
