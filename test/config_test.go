package test

import (
	"testing"

	"github.com/aarcex3/mygpo-clone/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	cfg := config.LoadConfig("test")

	assert.Equal(t, "8080", cfg.ServerPort, "Expected ServerPort to be '8080'")
	assert.Equal(t, "127.0.0.1", cfg.ServerHost, "Expected ServerHost to be '127.0.0.1'")
	assert.Equal(t, "sqlite3", cfg.DatabaseEngine, "Expected DatabaseEngine to be 'sqlite3'")
	assert.Equal(t, ":memory:", cfg.DatabaseURL, "Expected DatabaseURL to be ':memory:'")
	assert.Equal(t, "supersecretkey", string(cfg.SecretKey), "Expected SecretKey to be 'supersecretkey'")
}
