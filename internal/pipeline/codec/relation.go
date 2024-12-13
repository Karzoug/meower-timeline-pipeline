package codec

import (
	"errors"

	relationApi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/kafka/relation/v1"

	"google.golang.org/protobuf/proto"
)

type RelationCodec struct{}

func (c RelationCodec) Encode(value any) ([]byte, error) {
	v, ok := value.(*relationApi.ChangedEvent)
	if !ok {
		return nil, errors.New("relationCodec: value to encode is not of type *ChangedEvent")
	}

	return proto.Marshal(v)
}

func (c RelationCodec) Decode(data []byte) (any, error) {
	event := &relationApi.ChangedEvent{}

	return event, proto.Unmarshal(data, event)
}
