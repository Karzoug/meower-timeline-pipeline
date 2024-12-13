package handler

import (
	"time"

	"github.com/lovoo/goka"
)

const (
	tracePreffix                          = "TimelinePipeline."
	timelineTopic             goka.Stream = "timelines"
	defaultOperationTimeout               = 5 * time.Second
	maxRetryTimeoutBeforeExit             = 60 * time.Second
)
