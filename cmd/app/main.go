package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"task3_4/user-management/internal/controller/http/handler"
	"task3_4/user-management/internal/infrastructure/config"
	registry "task3_4/user-management/internal/infrastructure/registry/app"
	"task3_4/user-management/pkg/datastore"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	redis := datastore.NewRedisDB(cfg)
	db := datastore.NewMongoDB(cfg)

	r := registry.NewRegistry(db, redis, cfg)

	route := handler.NewRoute(echo.New(), r.NewAppControllers())
	route.InitRoutes()

	fmt.Println("Server listen at http://localhost" + ":" + cfg.ServerAddress)
	if err := route.Server.Start(":" + cfg.ServerAddress); err != nil {
		log.Fatalln(err)
	}
}
