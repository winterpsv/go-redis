package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockAuth struct {
	mock.Mock
}

func (m *MockAuth) GenerateHash(password string) string {
	args := m.Called(password)
	return args.String(0)
}

func (m *MockAuth) comparePasswordHash(nickname, password string) bool {
	args := m.Called(nickname, password)
	return args.Bool(0)
}

func (m *MockAuth) GenerateToken(nickname, password string) (string, error) {
	args := m.Called(nickname, password)
	return args.String(0), args.Error(1)
}

func (m *MockAuth) parseHeader(req *http.Request) (string, error) {
	args := m.Called(req)
	return args.String(0), args.Error(1)
}

func (m *MockAuth) Authenticate(req *http.Request) (*jwt.Token, error) {
	args := m.Called(req)
	return args.Get(0).(*jwt.Token), args.Error(1)
}

func (m *MockAuth) IsAdmin(token *jwt.Token) error {
	args := m.Called(token)
	return args.Error(0)
}
