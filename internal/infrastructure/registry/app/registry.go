package appRegistry

import (
	"go.mongodb.org/mongo-driver/mongo"
	"task3_3_new/user-management/internal/controller/http/v1"
	"task3_3_new/user-management/internal/infrastructure/config"
)

type registry struct {
	db  *mongo.Database
	cfg *config.Config
}

type Controllers struct {
	UserController controller.UserControllerInterface
	AuthController controller.AuthControllerInterface
}

type Registry interface {
	NewAppControllers() *Controllers
}

func NewRegistry(db *mongo.Database, cfg *config.Config) Registry {
	return &registry{db, cfg}
}

func (r *registry) NewAppControllers() *Controllers {
	return &Controllers{r.NewUserController(), r.NewAuthController()}
}
