package redisRepository

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Get(key string) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockUserRepository) Save(key string, data []byte, expiration time.Duration) error {
	args := m.Called(key, data, expiration)
	return args.Error(1)
}

func (m *MockUserRepository) Delete(key string) error {
	args := m.Called(key)
	return args.Error(1)
}
