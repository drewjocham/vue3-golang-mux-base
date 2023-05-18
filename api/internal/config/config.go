package config

import (
	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	HTTPPort      int    `envconfig:"PORT" default:"8081"`
	ServerAddress string `envconfig:"SERVER_ADDRESS" default:"localhost:8082"`
}

type Config struct {
	Server ServerConfig

	Port    int
	Env     string
	Version string

	Limiter struct {
		Enabled bool
		Rps     float64
		Burst   int
	}

	Smtp struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
	}

	Cors struct {
		TrustedOrigins []string
	}
}

func NewConfig() (*Config, error) {
	var c Config

	err := envconfig.Process("", &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
