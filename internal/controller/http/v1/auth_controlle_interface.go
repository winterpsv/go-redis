package controller

import "github.com/labstack/echo/v4"

type AuthControllerInterface interface {
	CreateUser(c echo.Context) error
	CreateToken(c echo.Context) error
	UpdateUserPassword(c echo.Context) error
	Authenticate(next echo.HandlerFunc) echo.HandlerFunc
	IsAdmin(next echo.HandlerFunc) echo.HandlerFunc
}
