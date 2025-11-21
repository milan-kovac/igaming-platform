package config

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	DB     MySQLEnv
	Server ServerEnv
}

type MySQLEnv struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type ServerEnv struct {
	Port string
}

func getEnvOrFail(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Panicf("Environment variable %s is required but not set", key)
	}
	return val
}

func LoadEnv() *Env {
	cfg := &Env{
		Server: ServerEnv{
			Port: getEnvOrFail("PORT"),
		},
		DB: MySQLEnv{
			Username: getEnvOrFail("DB_USER"),
			Password: getEnvOrFail("DB_PASSWORD"),
			Host:     getEnvOrFail("DB_HOST"),
			Port:     getEnvOrFail("DB_PORT"),
			Database: getEnvOrFail("DB_NAME"),
		},
	}

	return cfg
}
