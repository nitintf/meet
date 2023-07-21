package config

import (
	"github.com/friendsofgo/errors"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"meet"`
	Port        int    `envconfig:"SERVICE_PORT" default:"8080"`
}

func New() (*Config, error) {
	cfg := new(Config)

	err := envconfig.Process("", cfg)

	if err != nil {
		return nil, errors.Wrap(err, "error loading env variables")
	}

	return cfg, nil
}
