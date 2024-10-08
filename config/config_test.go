package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	os.Setenv("ServerPort", "8080")
	os.Setenv("ServerHost", "localhost")
	os.Setenv("DatabaseEngine", "postgres")
	os.Setenv("DatabaseURL", "postgres://user:password@localhost:5432/testdb")
	os.Setenv("SecretKey", "supersecretkey")

	cfg := LoadConfig()

	if cfg.ServerPort != "8080" {
		t.Errorf("Expected ServerPort to be '8080', got '%s'", cfg.ServerPort)
	}
	if cfg.ServerHost != "localhost" {
		t.Errorf("Expected ServerHost to be 'localhost', got '%s'", cfg.ServerHost)
	}
	if cfg.DatabaseEngine != "postgres" {
		t.Errorf("Expected DatabaseEngine to be 'postgres', got '%s'", cfg.DatabaseEngine)
	}
	if cfg.DatabaseURL != "postgres://user:password@localhost:5432/testdb" {
		t.Errorf("Expected DatabaseURL to be 'postgres://user:password@localhost:5432/testdb', got '%s'", cfg.DatabaseURL)
	}
	if string(cfg.SecretKey) != "supersecretkey" {
		t.Errorf("Expected SecretKey to be 'supersecretkey', got '%s'", string(cfg.SecretKey))
	}

	// Clean up the environment variables after the test
	os.Clearenv()
}
