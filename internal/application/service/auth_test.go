package service

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	repository "task3_4/user-management/internal/adapter/db/mongodb"
	model "task3_4/user-management/internal/entity"
	"task3_4/user-management/internal/infrastructure/config"
	"testing"
)

var (
	// Создание mock объекта UserRepository
	userRepository = new(repository.MockUserRepository)
	conf           = new(config.Config)
	// Создание объекта Auth с mock UserRepository
	mockauth = NewAuth(userRepository, conf)
)

func TestAuth_GenerateHash(t *testing.T) {
	actual := mockauth.GenerateHash("password123")

	assert.NotEmpty(t, actual)
}

func TestAuth_СomparePasswordHash(t *testing.T) {
	actual := mockauth.GenerateHash("password")

	// Ожидаемый результат для метода FindByNickname
	ID, _ := primitive.ObjectIDFromHex("123")
	expectedUser := &model.User{ID: ID, Nickname: "test", PasswordHash: actual}
	userRepository.On("FindByNickname", "test").Return(expectedUser, nil)

	// Вызов метода СomparePasswordHash
	resultTrue := mockauth.comparePasswordHash("test", "password")

	// Проверка результата на true
	assert.True(t, resultTrue)

	// Вызов метода СomparePasswordHash
	resultFalse := mockauth.comparePasswordHash("test", "password123")

	// Проверка результата на false
	assert.False(t, resultFalse)
}
