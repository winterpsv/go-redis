package interactor

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	repository "task3_3_new/user-management/internal/adapter/db/mongodb"
	"task3_3_new/user-management/internal/application/service"
	"task3_3_new/user-management/internal/controller/http/dto"
	"task3_3_new/user-management/internal/controller/http/presenter"
	model "task3_3_new/user-management/internal/entity"
	"task3_3_new/user-management/pkg/apperror"
	"time"
)

type AuthInteractor struct {
	UserRepository repository.UserRepositoryInterface
	UserPresenter  presenter.UserPresenterInterface
	Auth           service.AuthInterface
}

func NewAuthInteractor(r repository.UserRepositoryInterface, p presenter.UserPresenterInterface, a service.AuthInterface) *AuthInteractor {
	return &AuthInteractor{r, p, a}
}

func (au *AuthInteractor) Create(userForm *dto.CreateUserDTO) (*dto.UserDTO, error) {
	password := au.Auth.GenerateHash(userForm.Password)
	timestamp := time.Now().Unix()

	uModel := model.User{
		Nickname:     userForm.Nickname,
		FirstName:    userForm.FirstName,
		LastName:     userForm.LastName,
		PasswordHash: password,
		CreatedAt:    timestamp,
		Role:         userForm.Role,
		Active:       true,
		Votes:        []model.UserVote{},
	}

	existingUser, err := au.UserRepository.FindByNickname(userForm.Nickname)
	if existingUser != nil && existingUser.Nickname != "" {
		return nil, apperror.NewAppError(http.StatusConflict, fmt.Sprintf("User with nickname %s already exists and his ID %s", uModel.Nickname, existingUser.ID.Hex()), err)
	}

	u, err := au.UserRepository.Create(uModel)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusInternalServerError, "Failed to create user", err)
	}

	return au.UserPresenter.ResponseUser(u), nil
}

func (au *AuthInteractor) UpdatePassword(userForm *dto.UpdateUserPasswordDTO, ID string) (*dto.UserDTO, error) {
	u, err := au.UserRepository.FindByID(ID)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusInternalServerError, "Failed to find user", err)
	}

	if u.Active == false {
		return nil, apperror.NewAppError(http.StatusNoContent, fmt.Sprintf("user %s is deleted", u.Nickname), err)
	}

	u.PasswordHash = au.Auth.GenerateHash(userForm.Password)
	u.UpdatedAt = time.Now().Unix()

	u, err = au.UserRepository.Update(u)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusInternalServerError, "Failed to update user", err)
	}

	return au.UserPresenter.ResponseUser(u), nil
}

func (au *AuthInteractor) Authenticate(req *http.Request) (*jwt.Token, error) {
	token, err := au.Auth.Authenticate(req)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (au *AuthInteractor) IsAdmin(token *jwt.Token) error {
	err := au.Auth.IsAdmin(token)
	if err != nil {
		return err
	}

	return nil
}

func (au *AuthInteractor) CreateToken(userForm *dto.CreateTokenDTO) (string, error) {
	token, err := au.Auth.GenerateToken(userForm.Nickname, userForm.Password)
	if err != nil {
		return "", err
	}

	return au.UserPresenter.ResponseToken(token), nil
}
