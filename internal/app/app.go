package app

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Karzoug/meower-common-go/metric/prom"
	"github.com/Karzoug/meower-common-go/trace/otlp"
	"github.com/Karzoug/meower-timeline-pipeline/internal/config"
	"github.com/Karzoug/meower-timeline-pipeline/internal/pipeline"
	"github.com/Karzoug/meower-timeline-pipeline/pkg/buildinfo"
	grpccfg "github.com/Karzoug/meower-timeline-pipeline/pkg/grpc"
	relapi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/grpc/relation/v1"
	"github.com/Karzoug/meower-timeline-pipeline/pkg/redis"
)

const (
	serviceName     = "TimelinePipeline"
	metricNamespace = "timeline_pipeline"
	pkgName         = "github.com/Karzoug/meower-timeline-pipeline"
	initTimeout     = 10 * time.Second
	shutdownTimeout = 10 * time.Second
)

var serviceVersion = buildinfo.Get().ServiceVersion

func Run(ctx context.Context, logger zerolog.Logger) error {
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(cfg.LogLevel)

	logger.Info().
		Int("GOMAXPROCS", runtime.GOMAXPROCS(0)).
		Str("log level", cfg.LogLevel.String()).
		Msg("starting up")

	// set timeout for initialization
	ctxInit, closeCtx := context.WithTimeout(ctx, initTimeout)
	defer closeCtx()

	// set up tracer
	cfg.OTLP.ServiceName = serviceName
	cfg.OTLP.ServiceVersion = serviceVersion
	shutdownTracer, err := otlp.RegisterGlobal(ctxInit, cfg.OTLP)
	if err != nil {
		return err
	}
	defer doClose(shutdownTracer, logger)

	tracer := otel.GetTracerProvider().Tracer(pkgName)

	// set up meter
	shutdownMeter, err := prom.RegisterGlobal(ctxInit, serviceName, serviceVersion, metricNamespace)
	if err != nil {
		return err
	}
	defer doClose(shutdownMeter, logger)

	relClient, err := buildRelationClient(cfg.RelationClient)
	if err != nil {
		return err
	}

	redisClient, shutdownRedis, err := redis.New(cfg.Redis)
	if err != nil {
		return err
	}
	defer doClose(shutdownRedis, logger)

	pipe, err := pipeline.New(cfg.Kafka, relClient, redisClient, tracer, logger)
	if err != nil {
		return err
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return pipe.Run(ctx)
	})
	// run prometheus metrics http server
	eg.Go(func() error {
		return prom.Serve(ctx, cfg.PromHTTP, logger)
	})

	return eg.Wait()
}

func doClose(fn func(context.Context) error, logger zerolog.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := fn(ctx); err != nil {
		logger.Error().
			Err(err).
			Msg("error closing")
	}
}

func buildRelationClient(cfg grpccfg.Config) (relapi.RelationServiceClient, error) {
	relConn, err := grpc.NewClient(
		cfg.URI,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("could not connect to relation microservice: %w", err)
	}

	return relapi.NewRelationServiceClient(relConn), nil
}
