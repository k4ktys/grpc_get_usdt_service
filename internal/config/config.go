package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env string `env:"ENV" env-default:"dev"`

	GrpcPort int `env:"GRPC_PORT"`

	DbDriver   string `env:"DB_DRIVER"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     string `env:"DB_PORT"`
	DbName     string `env:"DB_NAME"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
