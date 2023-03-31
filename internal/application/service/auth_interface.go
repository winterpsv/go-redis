package service

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type AuthInterface interface {
	GenerateHash(string) string
	comparePasswordHash(nickname, password string) bool
	GenerateToken(nickname, password string) (string, error)
	parseHeader(*http.Request) (string, error)
	Authenticate(*http.Request) (*jwt.Token, error)
	IsAdmin(*jwt.Token) error
}
