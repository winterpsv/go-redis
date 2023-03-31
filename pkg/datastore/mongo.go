package datastore

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"task3_3_new/user-management/internal/infrastructure/config"
	"time"
)

func NewDB(cfg *config.Config) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoUri))

	if err != nil {
		log.Fatalln(err)
	}

	return client.Database(cfg.MongoBase)
}
