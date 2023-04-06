package interactor

import (
	"github.com/golang-jwt/jwt/v5"
	"task3_4/user-management/internal/controller/http/dto"
)

type UserInteractorInterface interface {
	GetAll(page, pageSize int64, key string) ([]*dto.UserDTO, error)
	Get(id, key string) (*dto.UserDTO, error)
	UpdateVote(userForm *dto.VoteUserDTO, ID string, token *jwt.Token) (*dto.UserDTO, error)
	Update(userForm *dto.UpdateUserDTO, ID string) (*dto.UserDTO, error)
	Deactivate(ID string) (*dto.UserDTO, error)
	GetUserByToken(token *jwt.Token, key string) (*dto.UserDTO, error)
	CacheGet(key string) ([]byte, error)
}
