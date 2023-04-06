package config

import (
	"github.com/caarlos0/env/v7"
)

type Config struct {
	MongoUri        string `env:"MONGO_URI"`
	MongoBase       string `env:"MONGO_BD"`
	MongoCollection string `env:"MONGO_COLLECTION"`
	ServerAddress   string `env:"SERVER_ADDRESS"`
	JwtSecret       string `env:"JWT_SECRET"`
	RedisAddr       string `env:"REDIS_ADDR"`
	RedisPass       string `env:"REDIS_PASS"`
	RedisDB         int    `env:"REDIS_DB"`
}

func ReadConfig() (config *Config, err error) {
	config = &Config{}

	opts := env.Options{RequiredIfNoDef: true}

	if err := env.Parse(config, opts); err != nil {
		return nil, err
	}

	return
}
