package main

import (
	"os"
	"multiapi/pkg/env"
)

// AppConfig contains API/business configurations
type AppConfig struct {
	TestMode bool
}

// ServerConfig contains server configurations (HTTP, etc)
type ServerConfig struct {
	Address string
	Port    int
}

// DatabaseConfig contains database configurations
type DatabaseConfig struct {
	URI                  string
	MigrationsDir        string
	MigrationsLogVerbose bool
}

type Config struct {
	Server    ServerConfig
	AppConfig AppConfig
	Database  DatabaseConfig
}

// NewConfig returns a Config object populated with values from environment variables or defaults
func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Address: env.GetEnv("ADDRESS", "localhost"),
			Port:    env.GetEnvAsInt("PORT", 4000),
		},
		AppConfig: AppConfig{
			TestMode: env.GetEnvAsBool("TEST_MODE", false),
		},
		Database: DatabaseConfig{
			URI:                  os.Getenv("DB_URI"),
			MigrationsDir:        env.GetEnv("DB_MIGRATIONS_DIR", "file://migrations"),
			MigrationsLogVerbose: env.GetEnvAsBool("DB_MIGRATIONS_VERBOSE", false),
		},
	}
}
