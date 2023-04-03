package controller

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
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
		return apperror.NewAppError(http.StatusBadRequest, "Failed to decode data", err)
	}

	if err := c.Validate(d); err != nil {
		return apperror.NewAppError(http.StatusUnprocessableEntity, "Validation failed", err)
	}

	u, err := ac.authInteractor.Create(d)
	if err != nil {
		return err
	}

	err = c.JSON(http.StatusOK, u)
	if err != nil {
		return apperror.NewAppError(http.StatusCreated, "User created successfully", err)
	}

	return nil
}

func (ac *AuthController) CreateToken(c echo.Context) error {
	d := new(dto.CreateTokenDTO)
	if err := c.Bind(d); err != nil {
		return apperror.NewAppError(http.StatusBadRequest, "Decode data failed", err)
	}

	if err := c.Validate(d); err != nil {
		return apperror.NewAppError(http.StatusUnprocessableEntity, "Validation failed", err)
	}

	token, err := ac.authInteractor.CreateToken(d)
	if err != nil {
		return err
	}

	err = c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
	if err != nil {
		return apperror.NewAppError(http.StatusCreated, "JWT token created successfully", err)
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
		return apperror.NewAppError(http.StatusBadRequest, "Failed to decode data", err)
	}

	if err := c.Validate(d); err != nil {
		return apperror.NewAppError(http.StatusUnprocessableEntity, "Validation failed", err)
	}

	u, err := ac.authInteractor.UpdatePassword(d, ID)
	if err != nil {
		return err
	}

	err = c.JSON(http.StatusOK, u)
	if err != nil {
		return apperror.NewAppError(http.StatusOK, "User updated successfully", err)
	}

	return nil
}
