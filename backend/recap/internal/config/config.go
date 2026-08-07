// Package config loads and validates recap service configuration.
package config

import (
	"fmt"
	"net"
	"time"

	"github.com/caarlos0/env/v11"

	applogger "github.com/avito-hackaton-team/avito-recap/backend/recap/internal/logger"
)

type Config struct {
	HTTP       HTTPConfig       `envPrefix:"HTTP_"`
	PostgreSQL PostgreSQLConfig `envPrefix:"POSTGRES_"`
	App        AppConfig        `envPrefix:"APP_"`
	Repository RepositoryConfig `envPrefix:"REPOSITORY_"`
	Logger     applogger.Config `envPrefix:"LOGGER_"`
}

type HTTPConfig struct {
	Address string `env:"ADDR" envDefault:"0.0.0.0:8080"`
}

type PostgreSQLConfig struct {
	Host           string        `env:"HOST" envDefault:"127.0.0.1"`
	Port           uint16        `env:"PORT" envDefault:"5432"`
	User           string        `env:"USER,notEmpty"`
	Password       string        `env:"PASSWORD,notEmpty"`
	Database       string        `env:"DB,notEmpty"`
	SSLMode        string        `env:"SSL_MODE" envDefault:"disable"`
	MaxConns       int32         `env:"MAX_CONNS" envDefault:"10"`
	MinConns       int32         `env:"MIN_CONNS" envDefault:"2"`
	ConnectTimeout time.Duration `env:"CONNECT_TIMEOUT" envDefault:"5s"`
}

type AppConfig struct {
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
}

type RepositoryConfig struct {
	OperationTimeout time.Duration `env:"OPERATION_TIMEOUT" envDefault:"3s"`
}

func Load() (Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return Config{}, fmt.Errorf("parse environment: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, fmt.Errorf("validate config: %w", err)
	}

	return cfg, nil
}

func (c Config) Validate() error {
	if _, _, err := net.SplitHostPort(c.HTTP.Address); err != nil {
		return fmt.Errorf("HTTP_ADDR: %w", err)
	}

	if c.App.ShutdownTimeout <= 0 {
		return fmt.Errorf("APP_SHUTDOWN_TIMEOUT must be positive")
	}

	if c.Repository.OperationTimeout <= 0 {
		return fmt.Errorf("REPOSITORY_OPERATION_TIMEOUT must be positive")
	}

	if c.PostgreSQL.Port == 0 {
		return fmt.Errorf("POSTGRES_PORT must be positive")
	}

	if c.PostgreSQL.ConnectTimeout <= 0 {
		return fmt.Errorf("POSTGRES_CONNECT_TIMEOUT must be positive")
	}

	if c.PostgreSQL.MaxConns <= 0 {
		return fmt.Errorf("POSTGRES_MAX_CONNS must be positive")
	}

	if c.PostgreSQL.MinConns < 0 {
		return fmt.Errorf("POSTGRES_MIN_CONNS must not be negative")
	}

	if c.PostgreSQL.MinConns > c.PostgreSQL.MaxConns {
		return fmt.Errorf("POSTGRES_MIN_CONNS must not exceed POSTGRES_MAX_CONNS")
	}

	return nil
}
