package controller

import "github.com/labstack/echo"

type UserControllerInterface interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUserVote(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeactivateUser(c echo.Context) error
	GetUserByToken(c echo.Context) error
}
