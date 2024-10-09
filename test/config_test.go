package test

import (
	"testing"

	"github.com/aarcex3/mygpo-clone/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	cfg := config.LoadConfig("test")

	assert.Equal(t, "8080", cfg.ServerPort, "Expected ServerPort to be '8080'")
	assert.Equal(t, "localhost", cfg.ServerHost, "Expected ServerHost to be 'localhost'")
	assert.Equal(t, "postgres", cfg.DatabaseEngine, "Expected DatabaseEngine to be 'postgres'")
	assert.Equal(t, "postgres://user:password@localhost:5432/testdb", cfg.DatabaseURL, "Expected DatabaseURL to be 'postgres://user:password@localhost:5432/testdb'")
	assert.Equal(t, "supersecretkey", string(cfg.SecretKey), "Expected SecretKey to be 'supersecretkey'")
}
