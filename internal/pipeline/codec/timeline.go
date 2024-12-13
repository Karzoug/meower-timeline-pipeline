package codec

import (
	"errors"

	timelineApi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/kafka/timeline/v1"

	"google.golang.org/protobuf/proto"
)

type TimelineCodec struct{}

func (c TimelineCodec) Encode(value any) ([]byte, error) {
	v, ok := value.(*timelineApi.ChangeTaskEvent)
	if !ok {
		return nil, errors.New("timelineCodec: value to encode is not of type *ChangedEvent")
	}

	return proto.Marshal(v)
}

func (c TimelineCodec) Decode(data []byte) (any, error) {
	event := &timelineApi.ChangeTaskEvent{}

	return event, proto.Unmarshal(data, event)
}
