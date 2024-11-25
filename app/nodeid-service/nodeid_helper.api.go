package nodeidapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeiderrorv1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/errors"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	threadpkg "github.com/ikaiguang/go-srv-kit/kratos/thread"
	"time"
)

const (
	DefaultRetryDelay        = 10 * time.Millisecond
	DefaultHeartbeatInterval = 30 * time.Second
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

type options struct {
	logger            log.Logger
	tries             int
	retryDelay        time.Duration
	heartbeatInterval time.Duration
}

func (s *options) doRetryDelay() {
	if s.retryDelay <= 0 {
		return
	}
	time.Sleep(s.retryDelay)
}

func WithHeartbeatInterval(duration time.Duration) Option {
	return func(o *options) {
		o.heartbeatInterval = duration
	}
}

type nodeIDHelper struct {
	opts   *options
	log    *log.Helper
	client NodeIDAPI
}

func NewNodeIDHelper(client NodeIDAPI, opts ...Option) NodeIDHelper {
	o := &options{
		retryDelay:        DefaultRetryDelay,
		heartbeatInterval: DefaultHeartbeatInterval,
	}
	for i := range opts {
		opts[i](o)
	}
	var logHelper *log.Helper
	if o.logger == nil {
		o.logger, _ = logpkg.NewDummyLogger()
	}
	logHelper = log.NewHelper(log.With(o.logger, "module", "nodeid-api"))
	return &nodeIDHelper{opts: o, log: logHelper, client: client}
}

func (s *nodeIDHelper) GetAndAutoRenewalNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, RenewalManager, error) {
	resp, err := s.GetNodeID(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	// 续订节点ID
	renewalInterface, err := s.RenewalNodeID(ctx, resp)
	if err != nil {
		return nil, nil, err
	}
	return resp, renewalInterface, nil
}

func (s *nodeIDHelper) GetNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error) {
	resp, err := s.client.GetNodeId(ctx, req)
	if err != nil {
		if s.opts.tries <= 1 {
			return nil, err
			//return nil, errorpkg.FormatError(err)
		}
		if s.log != nil {
			//s.log.WithContext(ctx).Warnw("msg", "GetNodeId failed", "try_number", 1, "error", e)
			s.log.WithContext(ctx).Warnw("msg", "GetNodeId failed", "try_number", 1, "error", err)
		}
	} else {
		return resp, nil
	}
	for i := 2; i <= s.opts.tries; i++ {
		s.opts.doRetryDelay()
		resp, err = s.client.GetNodeId(ctx, req)
		if err != nil {
			if s.opts.tries <= i {
				return nil, err
				//return nil, errorpkg.FormatError(err)
			}
			if s.log != nil {
				//s.log.WithContext(ctx).Warnw("msg", "GetNodeId failed", "try_number", 1, "error", e)
				s.log.WithContext(ctx).Warnw("msg", "GetNodeId failed", "try_number", 1, "error", err)
			}
		} else {
			return resp, nil
		}
	}
	if err == nil {
		e := nodeiderrorv1.DefaultErrorS102NoAvailableId()
		return nil, errorpkg.WithStack(e)
	}
	return nil, errorpkg.FormatError(err)
}

func (s *nodeIDHelper) RenewalNodeID(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (RenewalManager, error) {
	// token错误：ID被其他程序占用
	var (
		req = &nodeidresourcev1.RenewalNodeIdReq{
			InstanceId:  dataModel.InstanceId,
			NodeId:      dataModel.NodeId,
			AccessToken: dataModel.AccessToken,
		}
	)
	var (
		newCtx, cancel = context.WithCancel(context.Background())
		interval       = dataModel.HeartbeatInterval.AsDuration()
		renewalResult  = &RenewalResult{}
	)
	if interval <= time.Millisecond {
		interval = DefaultHeartbeatInterval
		if s.opts.heartbeatInterval > 0 {
			interval = s.opts.heartbeatInterval
		}
	}
	if s.log != nil {
		s.log.WithContext(ctx).Infow("msg", "RenewalNodeID started", "heartbeat_interval", interval.String())
	}
	threadpkg.GoSafe(func() {
		var (
			ticker = time.NewTicker(interval)
		)
		defer func() {
			ticker.Stop()
			cancel()
		}()
		for {
			select {
			case <-newCtx.Done():
				if s.log != nil {
					s.log.WithContext(ctx).Infow("msg", "RenewalNodeID stopped")
				}
				return
			case <-ticker.C:
				renewalResult.LastTime = time.Now()
				resp, err := s.renewalNodeID(newCtx, req)
				if err != nil {
					if s.log != nil {
						s.log.WithContext(ctx).Warnw("msg", "RenewalNodeID failed", "error", err)
					}
					renewalResult.Err = err
					break
				}
				renewalResult.Data = resp
			}
		}
	})

	renewalHandler := &renewalManager{
		data: renewalResult,
		stop: cancel,
	}
	return renewalHandler, nil
}

func (s *nodeIDHelper) renewalNodeID(ctx context.Context, req *nodeidresourcev1.RenewalNodeIdReq) (*nodeidresourcev1.RenewalNodeIdRespData, error) {
	resp, err := s.client.RenewalNodeId(ctx, req)
	if err != nil {
		if s.opts.tries <= 1 {
			return nil, err
			//return nil, errorpkg.FormatError(err)
		}
		if s.log != nil {
			//s.log.WithContext(ctx).Warnw("msg", "RenewalNodeId failed", "try_number", 1, "error", e)
			s.log.WithContext(ctx).Warnw("msg", "RenewalNodeId failed", "try_number", 1, "error", err)
		}
	} else {
		return resp, nil
	}
	for i := 2; i <= s.opts.tries; i++ {
		s.opts.doRetryDelay()
		resp, err = s.client.RenewalNodeId(ctx, req)
		if err != nil {
			if s.opts.tries <= i {
				return nil, err
				//return nil, errorpkg.FormatError(err)
			}
			if s.log != nil {
				//s.log.WithContext(ctx).Warnw("msg", "RenewalNodeId failed", "try_number", 1, "error", e)
				s.log.WithContext(ctx).Warnw("msg", "RenewalNodeId failed", "try_number", 1, "error", err)
			}
		} else {
			return resp, nil
		}
	}
	if err == nil {
		e := nodeiderrorv1.DefaultErrorS102NodeIdRenewalFailed()
		return nil, errorpkg.WithStack(e)
	}
	return nil, errorpkg.FormatError(err)
	//var tokenIncorrect = nodeiderrorv1.ERROR_S102_ACCESS_TOKEN_INCORRECT
	//if err != nil {
	//	if apiutil.IsReason(err, tokenIncorrect) {
	//
	//	}
	//}
	//if e := apiutil.CheckResponseCode(resp); e != nil {
	//	if e.Reason == tokenIncorrect.String() {
	//
	//	}
	//}
}

// ReleaseNodeId 释放节点，记得调用 RenewalManager.Stop 停止刷新节点ID
func (s *nodeIDHelper) ReleaseNodeId(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (*nodeidresourcev1.ReleaseNodeIdRespData, error) {
	req := &nodeidresourcev1.ReleaseNodeIdReq{
		InstanceId:  dataModel.InstanceId,
		NodeId:      dataModel.NodeId,
		AccessToken: dataModel.AccessToken,
	}
	resp, err := s.client.ReleaseNodeId(ctx, req)
	if err != nil {
		if s.opts.tries <= 1 {
			return nil, err
			//return nil, errorpkg.FormatError(err)
		}
		if s.log != nil {
			//s.log.WithContext(ctx).Warnw("msg", "ReleaseNodeId failed", "try_number", 1, "error", e)
			s.log.WithContext(ctx).Warnw("msg", "ReleaseNodeId failed", "try_number", 1, "error", err)
		}
	} else {
		return resp, nil
	}
	for i := 2; i <= s.opts.tries; i++ {
		s.opts.doRetryDelay()
		resp, err = s.client.ReleaseNodeId(ctx, req)
		if err != nil {
			if s.opts.tries <= i {
				return nil, err
				//return nil, errorpkg.FormatError(err)
			}
			if s.log != nil {
				//s.log.WithContext(ctx).Warnw("msg", "ReleaseNodeId failed", "try_number", 1, "error", e)
				s.log.WithContext(ctx).Warnw("msg", "ReleaseNodeId failed", "try_number", 1, "error", err)
			}
		} else {
			return resp, nil
		}
	}
	if err == nil {
		e := errorpkg.ErrorInternalServer("ReleaseNodeId failed")
		return nil, errorpkg.WithStack(e)
	}
	return nil, errorpkg.FormatError(err)
}
