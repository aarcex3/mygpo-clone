package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort     string
	ServerHost     string
	DatabaseEngine string
	DatabaseURL    string
	SecretKey      []byte
}

func LoadConfig(env string) *Config {

	var envFile string
	switch env {
	case "prod":
		envFile = ".env.prod"
	case "test":
		envFile = ".env.test"
	default:
		envFile = ".env.dev"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	return &Config{
		ServerPort:     os.Getenv("ServerPort"),
		ServerHost:     os.Getenv("ServerHost"),
		DatabaseEngine: os.Getenv("DatabaseEngine"),
		DatabaseURL:    os.Getenv("DatabaseURL"),
		SecretKey:      []byte(os.Getenv("SecretKey")),
	}
}
