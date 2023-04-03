package appRegistry

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	controller "task3_4/user-management/internal/controller/http/v1"
	"task3_4/user-management/internal/infrastructure/config"
)

type registry struct {
	db    *mongo.Database
	redis *redis.Client
	cfg   *config.Config
}

type Controllers struct {
	UserController controller.UserControllerInterface
	AuthController controller.AuthControllerInterface
}

type Registry interface {
	NewAppControllers() *Controllers
}

func NewRegistry(db *mongo.Database, redis *redis.Client, cfg *config.Config) Registry {
	return &registry{db, redis, cfg}
}

func (r *registry) NewAppControllers() *Controllers {
	return &Controllers{r.NewUserController(), r.NewAuthController()}
}
