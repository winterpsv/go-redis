package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	MongoUri        string `env:"MONGO_URI"`
	MongoBase       string `env:"MONGO_BD"`
	MongoCollection string `env:"MONGO_COLLECTION"`
	ServerAddress   string `env:"SERVER_ADDRESS"`
	JwtSecret       string `env:"JWT_SECRET"`
}

func ReadConfig() (config *Config, err error) {
	config = &Config{}

	opts := env.Options{RequiredIfNoDef: true}
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := env.Parse(config, opts); err != nil {
		return nil, err
	}

	return
}
