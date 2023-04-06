package interactor

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	repository "task3_4/user-management/internal/adapter/db/mongodb"
	"task3_4/user-management/internal/application/service"
	"task3_4/user-management/internal/controller/http/dto"
	"task3_4/user-management/internal/controller/http/presenter"
	model "task3_4/user-management/internal/entity"
	"task3_4/user-management/pkg/apperror"
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
		return nil, apperror.ErrExistsUser
	}

	u, err := au.UserRepository.Create(uModel)
	if err != nil {
		return nil, apperror.ErrCreateUser
	}

	return au.UserPresenter.ResponseUser(u), nil
}

func (au *AuthInteractor) UpdatePassword(userForm *dto.UpdateUserPasswordDTO, ID string) (*dto.UserDTO, error) {
	u, err := au.UserRepository.FindByID(ID)
	if err != nil {
		return nil, apperror.ErrFindUser
	}

	if !u.Active {
		return nil, apperror.ErrEditDeletedUser
	}

	u.PasswordHash = au.Auth.GenerateHash(userForm.Password)
	u.UpdatedAt = time.Now().Unix()

	u, err = au.UserRepository.Update(u)
	if err != nil {
		return nil, apperror.ErrUpdateUser
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
