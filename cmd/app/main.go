package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"task3_3_new/user-management/internal/controller/http/handler"
	"task3_3_new/user-management/internal/infrastructure/config"
	registry "task3_3_new/user-management/internal/infrastructure/registry/app"
	"task3_3_new/user-management/pkg/datastore"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db := datastore.NewDB(cfg)

	r := registry.NewRegistry(db, cfg)

	route := handler.NewRoute(echo.New(), r.NewAppControllers())
	route.InitRoutes()

	fmt.Println("Server listen at http://localhost" + ":" + cfg.ServerAddress)
	if err := route.Server.Start(":" + cfg.ServerAddress); err != nil {
		log.Fatalln(err)
	}
}
