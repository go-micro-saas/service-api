package nodeidapi

import (
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Option func(*options)

func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func WithTries(tries int) Option {
	return func(o *options) {
		o.tries = tries
	}
}

func WithRetryDelay(delay time.Duration) Option {
	return func(o *options) {
		o.retryDelay = delay
	}
}

func WithHeartbeatInterval(duration time.Duration) Option {
	return func(o *options) {
		o.heartbeatInterval = duration
	}
}
