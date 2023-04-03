package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	controller "task3_4/user-management/internal/controller/http/v1"
	registry "task3_4/user-management/internal/infrastructure/registry/app"
	"task3_4/user-management/pkg/echovalidator"
)

type Router struct {
	Server         *echo.Echo
	userController controller.UserControllerInterface
	authController controller.AuthControllerInterface
}

func NewRoute(Server *echo.Echo, registry *registry.Controllers) *Router {
	return &Router{
		Server:         Server,
		userController: registry.UserController,
		authController: registry.AuthController,
	}
}

func (r *Router) InitRoutes() {
	v := validator.New()
	r.Server.Validator = &echovalidator.CustomValidator{Validator: v}
	r.Server.Use(r.authMiddleware)
	r.Server.Use(middleware.Logger())
	r.Server.Use(middleware.Recover())

	//API v1
	v1 := r.Server.Group("/v1")

	//User rotes
	userRoutes := v1.Group("/user")
	userRoutes.GET("", r.userController.GetUsers)
	userRoutes.GET("/:id", r.userController.GetUser)
	userRoutes.PUT("/:id/vote", r.userController.UpdateUserVote)
	userRoutes.PUT("/:id", r.userController.UpdateUser, r.authController.IsAdmin)
	userRoutes.DELETE("/:id", r.userController.DeactivateUser, r.authController.IsAdmin)
	userRoutes.GET("/profile", r.userController.GetUserByToken)

	//Auth rotes
	authRoutes := v1.Group("/auth")
	authRoutes.POST("/register", r.authController.CreateUser)
	authRoutes.POST("/login", r.authController.CreateToken)
	authRoutes.PUT("/:id", r.authController.UpdateUserPassword, r.authController.IsAdmin)
}

func (r *Router) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path == "/v1/auth/register" || c.Request().URL.Path == "/v1/auth/login" {
			return next(c)
		}
		return r.authController.Authenticate(next)(c)
	}
}
