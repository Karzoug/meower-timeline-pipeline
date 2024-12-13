package handler

import (
	"github.com/lovoo/goka"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"

	ck "github.com/Karzoug/meower-common-go/kafka"
	relapi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/kafka/relation/v1"
	tmlineapi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/kafka/timeline/v1"
)

func Relation(tracer trace.Tracer, logger zerolog.Logger) func(goka.Context, any) {
	return func(gctx goka.Context, msg any) {
		event := msg.(*relapi.ChangedEvent)

		var outEventChangeTaskType tmlineapi.ChangeTaskType
		switch event.ChangeType {
		case relapi.ChangeType_CHANGE_TYPE_FOLLOW, relapi.ChangeType_CHANGE_TYPE_UNMUTE:
			outEventChangeTaskType = tmlineapi.ChangeTaskType_CHANGE_TASK_TYPE_USER_SUBSCRIBE
		case relapi.ChangeType_CHANGE_TYPE_UNFOLLOW, relapi.ChangeType_CHANGE_TYPE_MUTE:
			outEventChangeTaskType = tmlineapi.ChangeTaskType_CHANGE_TASK_TYPE_USER_UNSUBSCRIBE
		default:
			return
		}

		_, span := tracer.Start(gctx.Context(), tracePreffix+"RelationEventHandler")
		defer span.End()

		logger := logger.With().
			Str("topic", string(gctx.Topic())).
			Str("key", gctx.Key()).
			Str("tarce_id", span.SpanContext().TraceID().String()).
			Logger()

		logger.Info().Msg("event received")

		hdr := goka.Headers{}
		hdr[ck.MessageTypeHeaderKey] = []byte(ck.MessageTypeHeaderValue(&tmlineapi.ChangeTaskEvent{}))

		gctx.Emit(timelineTopic,
			event.SourceUserId,
			&tmlineapi.ChangeTaskEvent{
				UserId:       event.SourceUserId,
				TargetUserId: event.TargetUserId,
				ChangeType:   outEventChangeTaskType,
			},
			goka.WithCtxEmitHeaders(hdr),
		)

		logger.Info().Msg("event processed")
	}
}
