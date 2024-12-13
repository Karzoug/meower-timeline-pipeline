package codec

import (
	"errors"

	userApi "github.com/Karzoug/meower-timeline-pipeline/pkg/proto/kafka/user/v1"

	"google.golang.org/protobuf/proto"
)

type UserCodec struct{}

func (c UserCodec) Encode(value any) ([]byte, error) {
	v, ok := value.(*userApi.ChangedEvent)
	if !ok {
		return nil, errors.New("relationCodec: value to encode is not of type *userApi.ChangedEvent")
	}

	return proto.Marshal(v)
}

func (c UserCodec) Decode(data []byte) (any, error) {
	event := &userApi.ChangedEvent{}

	return event, proto.Unmarshal(data, event)
}
