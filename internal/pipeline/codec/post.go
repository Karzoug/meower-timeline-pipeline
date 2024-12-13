package codec

import (
	"errors"

	postApi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/kafka/post/v1"

	"google.golang.org/protobuf/proto"
)

type PostCodec struct{}

func (c PostCodec) Encode(value any) ([]byte, error) {
	v, ok := value.(*postApi.ChangedEvent)
	if !ok {
		return nil, errors.New("postCodec: value to encode is not of type *ChangedEvent")
	}

	return proto.Marshal(v)
}

func (c PostCodec) Decode(data []byte) (any, error) {
	event := &postApi.ChangedEvent{}

	return event, proto.Unmarshal(data, event)
}
