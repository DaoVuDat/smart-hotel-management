package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Uri               string
	MaxConn           int
	MinConn           int
	MaxConnLifeTime   time.Duration
	MaxIdleTime       time.Duration
	HealthCheckPeriod time.Duration
}

func Load() *Config {
	return &Config{
		DB: DBConfig{
			Uri:               GetEnv("DB_URI", "postgres://postgres:postgres@localhost:5432/postgres"),
			MaxConn:           GetEnvInt("DB_MAX_CONN", 10),
			MinConn:           GetEnvInt("DB_MIN_CONN", 5),
			MaxConnLifeTime:   GetEnvDuration("DB_MAX_CONN_LIFE_TIME", 10*time.Minute),
			MaxIdleTime:       GetEnvDuration("DB_MAX_IDLE_TIME", 10*time.Minute),
			HealthCheckPeriod: GetEnvDuration("DB_HEALTH_CHECK_PERIOD", 1*time.Minute),
		},
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if val, err := strconv.Atoi(value); err == nil {
			return val
		}
	}
	return fallback
}

func GetEnvDuration(key string, fallback time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); ok {
		if val, err := time.ParseDuration(value); err == nil {
			return val
		}
	}
	return fallback
}
