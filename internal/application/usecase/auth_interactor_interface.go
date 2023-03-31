package interactor

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"task3_3_new/user-management/internal/controller/http/dto"
)

type AuthInteractorInterface interface {
	Create(userForm *dto.CreateUserDTO) (*dto.UserDTO, error)
	UpdatePassword(userForm *dto.UpdateUserPasswordDTO, ID string) (*dto.UserDTO, error)
	Authenticate(req *http.Request) (*jwt.Token, error)
	IsAdmin(token *jwt.Token) error
	CreateToken(userForm *dto.CreateTokenDTO) (string, error)
}
