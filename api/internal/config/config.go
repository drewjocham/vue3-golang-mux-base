package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port    int    `envconfig:"PORT" default:"8081"` // PORT is default to 8080 in App Engine
	Env     string `envconfig:"ENV" default:"development"`
	Version string `envconfig:"VERSION" default:"development"`

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
		TrustedOrigins []string `envconfig:"CORS" default:"http://localhost:3000"`
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
