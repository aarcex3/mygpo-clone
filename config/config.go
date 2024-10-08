package config

import (
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

func LoadConfig() *Config {
	godotenv.Load()
	return &Config{
		ServerPort:     os.Getenv("ServerPort"),
		ServerHost:     os.Getenv("ServerHost"),
		DatabaseEngine: os.Getenv("DatabaseEngine"),
		DatabaseURL:    os.Getenv("DatabaseURL"),
		SecretKey:      []byte(os.Getenv("SecretKey")),
	}
}
