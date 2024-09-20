package nodeidapi

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Option func(*options)

type options struct {
	logger log.Logger
	tries  int
}

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

//func WithRetryDelay(delay time.Duration) Option {
//	return func(o *options) {
//		m.delayFunc = func(tries int) time.Duration {
//			return delay
//		}
//	}
//}
