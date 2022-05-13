package main

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	App struct {
		ServerAddr string `env:"SERVER_ADDR" envDefault:"localhost:50051"`
		Login      struct {
			Limit    int           `env:"LOGIN_LIMIT" envDefault:"10"`
			Interval time.Duration `env:"LOGIN_INTERVAL" envDefault:"1m0s"`
			Ttl      time.Duration `env:"LOGIN_TTL" envDefault:"2m0s"`
		}
		Password struct {
			Limit    int           `env:"PASSWORD_LIMIT" envDefault:"100"`
			Interval time.Duration `env:"PASSWORD_INTERVAL" envDefault:"1m0s"`
			Ttl      time.Duration `env:"PASSWORD_TTL" envDefault:"2m0s"`
		}
		IP struct {
			Limit    int           `env:"IP_LIMIT" envDefault:"1000"`
			Interval time.Duration `env:"IP_INTERVAL" envDefault:"1m0s"`
			Ttl      time.Duration `env:"IP_TTL" envDefault:"2m0s"`
		}
	}
	DB struct {
		Dsn string `env:"DB_DSN" envDefault:"postgres://test:test@0.0.0.0:5432/test?sslmode=disable"`
	}
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
