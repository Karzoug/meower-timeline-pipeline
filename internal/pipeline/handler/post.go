package handler

import (
	"context"
	"fmt"

	ck "github.com/Karzoug/meower-common-go/kafka"
	relapi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/grpc/relation/v1"
	postapi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/kafka/post/v1"
	tmlineapi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/kafka/timeline/v1"

	"github.com/cenkalti/backoff/v4"
	"github.com/lovoo/goka"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

func Post(relClient relapi.RelationServiceClient, tracer trace.Tracer, logger zerolog.Logger) func(goka.Context, any) {
	return func(gctx goka.Context, msg any) {
		event := msg.(*postapi.ChangedEvent)

		var outEventChangeTaskType tmlineapi.ChangeTaskType
		switch event.ChangeType {
		case postapi.ChangeType_CHANGE_TYPE_CREATED:
			outEventChangeTaskType = tmlineapi.ChangeTaskType_CHANGE_TASK_TYPE_POST_INSERT
		case postapi.ChangeType_CHANGE_TYPE_DELETED:
			outEventChangeTaskType = tmlineapi.ChangeTaskType_CHANGE_TASK_TYPE_POST_DELETE
		default:
			return
		}

		ctx, span := tracer.Start(gctx.Context(), tracePreffix+"PostEventHandler")
		defer span.End()

		logger := logger.With().
			Str("topic", string(gctx.Topic())).
			Str("key", gctx.Key()).
			Str("tarce_id", span.SpanContext().TraceID().String()).
			Logger()

		logger.Info().Msg("event received")

		var followers []*relapi.User
		operation := func() error {
			ctx, cancel := context.WithTimeout(ctx, defaultOperationTimeout)
			defer cancel()

			var err error
			resp, err := relClient.ListFollowers(ctx,
				&relapi.ListFollowersRequest{
					Parent:   event.AuthorId,
					PageSize: -1,
				})
			if err != nil {
				logger.Warn().
					Err(err).
					Msg("failed to list followers")
				return err
			}

			followers = resp.Followers

			return nil
		}
		if err := backoff.Retry(operation,
			backoff.NewExponentialBackOff(
				backoff.WithMaxElapsedTime(maxRetryTimeoutBeforeExit),
			),
		); err != nil {
			logger.Error().
				Err(err).
				Msg("failed to list followers, all retries failed")

			gctx.Fail(fmt.Errorf("all retries for list followers failed: %w", err))
			return
		}

		hdr := goka.Headers{}
		hdr[ck.MessageTypeHeaderKey] = []byte(ck.MessageTypeHeaderValue(&tmlineapi.ChangeTaskEvent{}))

		for _, follower := range followers {
			gctx.Emit(timelineTopic,
				follower.Id,
				&tmlineapi.ChangeTaskEvent{
					UserId:       event.AuthorId,
					TargetUserId: follower.Id,
					PostId:       event.Id,
					ChangeType:   outEventChangeTaskType,
				},
				goka.WithCtxEmitHeaders(hdr),
			)
		}

		logger.Info().Msg("event processed")
	}
}
