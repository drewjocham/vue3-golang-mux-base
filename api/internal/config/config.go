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
		TrustedOrigins []string `envconfig:"CORS" default:"http://127.0.0.1:5173"`
	}

	DatabaseConfig struct {
		Uri      string `envconfig:"DB_URI" default:"jdbc:postgresql://localhost:5434/fullstackguru"`
		Password string `envconfig:"DB_PASSWORD" default:"admin"`
		User     string `envconfig:"DB_USER" default:"admin"`
		DataName string `envconfig:"DB_NAME" default:"fullstackguru"`
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Port     int    `envconfig:"DB_PORT" default:"5434"`
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
