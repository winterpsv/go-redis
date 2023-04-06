package controller

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	interactor "task3_4/user-management/internal/application/usecase"
	"task3_4/user-management/internal/controller/http/dto"
)

type UserController struct {
	userInteractor interactor.UserInteractorInterface
}

func NewUserController(us interactor.UserInteractorInterface) *UserController {
	return &UserController{us}
}

func (uc *UserController) GetUsers(c echo.Context) error {
	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("per_page")
	key := strings.ToLower(fmt.Sprintf("%s_%s", c.Request().URL.Path, c.Request().URL.RawQuery))

	page, err := strconv.ParseInt(pageStr, 10, 0)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.ParseInt(perPageStr, 10, 0)
	if err != nil {
		pageSize = 10
	}

	u, err := uc.userInteractor.GetAll(page, pageSize, key)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.JSON(http.StatusOK, u)

	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return nil
}

func (uc *UserController) GetUser(c echo.Context) error {
	ID := c.Param("id")
	key := strings.ToLower(fmt.Sprintf("%s_%s", c.Request().URL.Path, c.Request().URL.RawQuery))

	u, err := uc.userInteractor.Get(ID, key)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.JSON(http.StatusOK, u)
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return nil
}

func (uc *UserController) UpdateUserVote(c echo.Context) error {
	d := new(dto.VoteUserDTO)
	targetID := c.Param("id")
	userID := c.Get("user").(*jwt.Token)

	if err := c.Bind(d); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "decode data failed")
	}

	if err := c.Validate(d); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	u, err := uc.userInteractor.UpdateVote(d, targetID, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.JSON(http.StatusOK, u)
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return nil
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	d := new(dto.UpdateUserDTO)
	ID := c.Param("id")

	if err := c.Bind(d); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "decode data failed")
	}

	if err := c.Validate(d); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	u, err := uc.userInteractor.Update(d, ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.JSON(http.StatusOK, u)
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return nil
}

func (uc *UserController) DeactivateUser(c echo.Context) error {
	ID := c.Param("id")

	u, err := uc.userInteractor.Deactivate(ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.JSON(http.StatusOK, u)
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return nil
}

func (uc *UserController) GetUserByToken(c echo.Context) error {
	u := c.Get("user").(*jwt.Token)
	key := strings.ToLower(fmt.Sprintf("%s_%s", c.Request().URL.Path, c.Request().URL.RawQuery))

	user, err := uc.userInteractor.GetUserByToken(u, key)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.JSON(http.StatusOK, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return nil
}

func (uc *UserController) ResponseCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := strings.ToLower(fmt.Sprintf("%s_%s", c.Request().URL.Path, c.Request().URL.RawQuery))

		data, err := uc.userInteractor.CacheGet(key)

		if err == nil {
			return c.JSONBlob(http.StatusOK, data)
		}

		return next(c)
	}
}
