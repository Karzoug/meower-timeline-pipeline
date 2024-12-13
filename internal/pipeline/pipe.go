package pipeline

import (
	"context"

	"github.com/lovoo/goka"
	storage "github.com/lovoo/goka/storage/redis"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/redis.v5"

	"github.com/Karzoug/meower-timeline-pipeline/internal/pipeline/codec"
	"github.com/Karzoug/meower-timeline-pipeline/internal/pipeline/handler"
	"github.com/Karzoug/meower-timeline-pipeline/pkg/kafka"
	relapi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/grpc/relation/v1"
)

const (
	postTopic      goka.Stream = "posts"
	userTopic      goka.Stream = "users"
	relationTopic  goka.Stream = "relations"
	timelineTopic  goka.Stream = "timelines"
	group          goka.Group  = "timeline-stream-group"
	redisNamespace             = "timeline-pipeline"
)

type pipeline struct {
	procesor *goka.Processor
}

func New(cfg kafka.Config,
	relClient relapi.RelationServiceClient,
	redisClient *redis.Client,
	tracer trace.Tracer,
	logger zerolog.Logger,
) (pipeline, error) {
	logger = logger.With().
		Str("component", "pipeline").
		Logger()

	group := goka.DefineGroup(group,
		goka.Input(postTopic, new(codec.PostCodec), handler.Post(relClient, tracer, logger)),
		goka.Input(userTopic, new(codec.UserCodec), handler.User(relClient, tracer, logger)),
		goka.Input(relationTopic, new(codec.RelationCodec), handler.Relation(tracer, logger)),
		goka.Output(timelineTopic, new(codec.TimelineCodec)),
	)

	gp, err := goka.NewProcessor(cfg.Brokers,
		group,
		goka.WithStorageBuilder(
			storage.RedisBuilder(redisClient, redisNamespace),
		),
		goka.WithLogger(newLogger(logger)),
	)
	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to create processor")
		return pipeline{}, err
	}

	return pipeline{procesor: gp}, nil
}

func (p pipeline) Run(ctx context.Context) error {
	return p.procesor.Run(ctx)
}
