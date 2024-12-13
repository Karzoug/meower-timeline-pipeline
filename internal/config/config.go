package config

import (
	"github.com/Karzoug/meower-common-go/metric/prom"

	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel zerolog.Level     `env:"LOG_LEVEL" envDefault:"info"`
	PromHTTP prom.ServerConfig `envPrefix:"PROM_"`
}
