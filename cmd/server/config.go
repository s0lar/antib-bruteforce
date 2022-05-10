package main

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	App struct {
		ServerAddr string `env:"SERVER_ADDR" envDefault:"localhost:50051"`
		Login      struct {
			limit    int           `env:"LOGIN_LIMIT" envDefault:"10"`
			interval time.Duration `env:"LOGIN_INTERVAL" envDefault:"1m0s"`
			ttl      time.Duration `env:"LOGIN_TTL" envDefault:"2m0s"`
		}
		Password struct {
			limit    int           `env:"PASSWORD_LIMIT" envDefault:"1000"`
			interval time.Duration `env:"PASSWORD_INTERVAL" envDefault:"1m0s"`
			ttl      time.Duration `env:"PASSWORD_TTL" envDefault:"2m0s"`
		}
		IP struct {
			limit    int           `env:"IP_LIMIT" envDefault:"1000"`
			interval time.Duration `env:"IP_INTERVAL" envDefault:"1m0s"`
			ttl      time.Duration `env:"IP_TTL" envDefault:"2m0s"`
		}
	}
	DB struct {
		Username  string `env:"DB_USERNAME,required" envDefault:"postgres"`
		Password  string `env:"DB_PASSWORD,required" envDefault:"postgres"`
		Host      string `env:"DB_HOST" envDefault:"localhost"`
		Port      int    `env:"DB_PORT" envDefault:"5432"`
		SSLEnable bool   `env:"DB_SSL_ENABLE" envDefault:"false"`
	}
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
