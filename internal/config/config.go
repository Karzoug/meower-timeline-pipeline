package config

import (
	"github.com/Karzoug/meower-common-go/metric/prom"
	"github.com/Karzoug/meower-common-go/trace/otlp"

	"github.com/Karzoug/meower-timeline-pipeline/pkg/grpc"
	"github.com/Karzoug/meower-timeline-pipeline/pkg/kafka"
	"github.com/Karzoug/meower-timeline-pipeline/pkg/redis"

	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel       zerolog.Level     `env:"LOG_LEVEL" envDefault:"info"`
	PromHTTP       prom.ServerConfig `envPrefix:"PROM_"`
	OTLP           otlp.Config       `envPrefix:"OTLP_"`
	Kafka          kafka.Config      `envPrefix:"KAFKA_"`
	RelationClient grpc.Config       `envPrefix:"RELATION_CLIENT_"`
	Redis          redis.Config      `envPrefix:"REDIS_"`
}
