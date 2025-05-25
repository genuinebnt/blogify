package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port int64
	Env  string
	DB   DBConfig
}

type DBConfig struct {
	Host        string
	User        string
	Password    string
	Port        int64
	Database    string
	MaxOpenConn int64
	MaxIdleConn int64
	MaxIdleTime time.Duration
}

func LoadConfig() (*Config, error) {
	portStr := os.Getenv("PORT")
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		return nil, errors.New("invalid PORT: " + err.Error())
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	} else if env != "development" && env != "production" {
		return nil, errors.New("ENVIRONMENT must be 'development' or 'production'")
	}

	if err != nil {
		return nil, err
	}

	host := parseEnvString("POSTGRES_HOST", "localhost")
	user := parseEnvString("POSTGRES_USER", "postgres")
	password := parseEnvString("POSTGRES_PASSWORD", "password")
	dbport := parseEnvInt("POSTGRES_PORT", 5432)
	database := parseEnvString("POSTGRES_DATABASE", "blogify")
	maxOpenConn := parseEnvInt("DATABASE_MAX_OPEN_CONN", 25)
	maxIdleConn := parseEnvInt("DATABASE_MAX_IDLE_CONN", 25)
	maxIdleTime := parseEnvDuration("DATABASE_MAX_IDLE_TIME", 15*time.Minute)

	return &Config{
		Port: port,
		Env:  env,
		DB: DBConfig{
			Host:        host,
			User:        user,
			Password:    password,
			Port:        dbport,
			Database:    database,
			MaxOpenConn: maxOpenConn,
			MaxIdleConn: maxIdleConn,
			MaxIdleTime: maxIdleTime,
		},
	}, nil
}

func (cfg *Config) ConnectionStringFromEnv() string {
	return fmt.Sprintf("postgres//%s:%s@%s:%d/%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)
}

func parseEnvInt(key string, defaultVal int64) int64 {
	if valStr := os.Getenv(key); valStr != "" {
		if val, err := strconv.ParseInt(valStr, 10, 64); err == nil {
			return val
		}
	}
	return defaultVal
}

func parseEnvString(key string, defaultVal string) string {
	if valStr := os.Getenv(key); key != "" {
		return valStr
	}

	return defaultVal
}

func parseEnvDuration(key string, defaultVal time.Duration) time.Duration {
	if valStr := os.Getenv(key); valStr != "" {
		if val, err := time.ParseDuration(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}
