package config

import (
	"fmt"
	"os"
)

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type AppConfig struct {
	Name     string
	Env      string
	Port     string
	LogLevel string
}

type PostgresConfig struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
	SSLMode  string
}

type RedisConfig struct {
	Host string
	Port string
}

func Load() (Config, error) {
	cfg := Config{
		App: AppConfig{
			Name:     getEnv("APP_NAME", "water-delivery"),
			Env:      getEnv("APP_ENV", "local"),
			Port:     getEnv("APP_PORT", "8080"),
			LogLevel: getEnv("APP_LOG_LEVEL", "debug"),
		},
		Postgres: PostgresConfig{
			Host:     getEnv("POSTGRES_HOST", ""),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			DB:       getEnv("POSTGRES_DB", ""),
			User:     getEnv("POSTGRES_USER", ""),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
		Redis: RedisConfig{
			Host: getEnv("REDIS_HOST", ""),
			Port: getEnv("REDIS_PORT", "6379"),
		},
	}

	if cfg.Postgres.Host == "" || cfg.Postgres.DB == "" || cfg.Postgres.User == "" {
		return Config{}, fmt.Errorf("postgres config is incomplete")
	}

	if cfg.Redis.Host == "" {
		return Config{}, fmt.Errorf("redis config is incomplete")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
