package controller

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	interactor "task3_4/user-management/internal/application/usecase"
	"task3_4/user-management/internal/controller/http/dto"
	"task3_4/user-management/pkg/apperror"
)

type AuthController struct {
	authInteractor interactor.AuthInteractorInterface
}

func NewAuthController(us interactor.AuthInteractorInterface) *AuthController {
	return &AuthController{us}
}

func (ac *AuthController) CreateUser(c echo.Context) error {
	d := new(dto.CreateUserDTO)
	if err := c.Bind(d); err != nil {
		return apperror.ErrDecodeData
	}

	if err := c.Validate(d); err != nil {
		return apperror.ErrValidation
	}

	u, err := ac.authInteractor.Create(d)
	if err != nil {
		return err
	}

	err = c.JSON(http.StatusOK, u)
	if err != nil {
		return apperror.SuccessCreatedUser
	}

	return nil
}

func (ac *AuthController) CreateToken(c echo.Context) error {
	d := new(dto.CreateTokenDTO)
	if err := c.Bind(d); err != nil {
		return apperror.ErrDecodeData
	}

	if err := c.Validate(d); err != nil {
		return apperror.ErrValidation
	}

	token, err := ac.authInteractor.CreateToken(d)
	if err != nil {
		return err
	}

	err = c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
	if err != nil {
		return apperror.SuccessCreatedToken
	}

	return nil
}

func (ac *AuthController) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := ac.authInteractor.Authenticate(c.Request())
		if err != nil {
			return err
		}

		c.Set("user", token)
		return next(c)
	}
}

func (ac *AuthController) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		err := ac.authInteractor.IsAdmin(token)
		if err != nil {
			return err
		}

		return next(c)
	}
}

func (ac *AuthController) UpdateUserPassword(c echo.Context) error {
	d := new(dto.UpdateUserPasswordDTO)
	ID := c.Param("id")
	if err := c.Bind(d); err != nil {
		return apperror.ErrDecodeData
	}

	if err := c.Validate(d); err != nil {
		return apperror.ErrValidation
	}

	u, err := ac.authInteractor.UpdatePassword(d, ID)
	if err != nil {
		return err
	}

	err = c.JSON(http.StatusOK, u)
	if err != nil {
		return apperror.SuccessUpdatedUser
	}

	return nil
}
