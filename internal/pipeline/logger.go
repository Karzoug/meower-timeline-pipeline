package pipeline

import (
	"github.com/lovoo/goka"
	"github.com/rs/zerolog"
)

var _ goka.Logger = &logger{}

type logger struct {
	zerolog.Logger
}

func newLogger(l zerolog.Logger) *logger {
	return &logger{l}
}

func (l *logger) Debugf(msg string, args ...any) {
	l.Debug().Msgf(msg, args...)
}
