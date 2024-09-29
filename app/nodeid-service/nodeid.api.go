package nodeidapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeiderrorv1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/errors"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	nodeidservicev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/services"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	threadpkg "github.com/ikaiguang/go-srv-kit/kratos/thread"
	"time"
)

const (
	DefaultRetryDelay        = 10 * time.Millisecond
	DefaultHeartbeatInterval = 30 * time.Second
)

type options struct {
	logger            log.Logger
	tries             int
	retryDelay        time.Duration
	heartbeatInterval time.Duration
}

type grpcAPI struct {
	opts   *options
	log    *log.Helper
	client nodeidservicev1.SrvNodeIDV1Client
}

func NewNodeIDAPI(client nodeidservicev1.SrvNodeIDV1Client, opts ...Option) NodeIDAPI {
	o := &options{
		retryDelay:        DefaultRetryDelay,
		heartbeatInterval: DefaultHeartbeatInterval,
	}
	for i := range opts {
		opts[i](o)
	}
	var logHelper *log.Helper
	if o.logger != nil {
		logHelper = log.NewHelper(log.With(o.logger, "module", "nodeid-api"))
	}
	return &grpcAPI{opts: o, log: logHelper, client: client}
}

func (s *grpcAPI) GetClient() interface{} {
	return s.client
}

func (s *grpcAPI) GetAndAutoRenewalNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, NodeIDInterface, error) {
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

func (s *grpcAPI) GetNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error) {
	resp, err := s.client.GetNodeId(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		if s.opts.tries <= 1 {
			return nil, errorpkg.WithStack(e)
		}
		if s.log != nil {
			s.log.WithContext(ctx).Warnw("msg", "GetNodeId failed", "try_number", 1, "error", e)
		}
	} else {
		return resp.Data, nil
	}
	for i := 2; i <= s.opts.tries; i++ {
		time.Sleep(s.opts.retryDelay)
		resp, err = s.client.GetNodeId(ctx, req)
		if e := apiutil.CheckAPIResponse(resp, err); e != nil {
			if s.opts.tries <= i {
				return nil, errorpkg.WithStack(e)
			}
			if s.log != nil {
				s.log.WithContext(ctx).Warnw("msg", "GetNodeId failed", "try_number", 1, "error", e)
			}
		} else {
			return resp.Data, nil
		}
	}
	if err == nil {
		e := nodeiderrorv1.DefaultErrorS102NoAvailableId()
		return nil, errorpkg.WithStack(e)
	}
	return nil, errorpkg.FormatError(err)
}

func (s *grpcAPI) RenewalNodeID(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (NodeIDInterface, error) {
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
	)
	if interval <= time.Second {
		interval = DefaultHeartbeatInterval
	}
	if s.opts.heartbeatInterval > 0 {
		interval = s.opts.heartbeatInterval
	}
	if s.log != nil {
		s.log.WithContext(ctx).Infow("msg", "RenewalNodeID started", "heartbeat_interval", interval.String())
	}
	threadpkg.GoSafe(func() {
		var (
			ticker = time.NewTicker(interval)
		)
		defer ticker.Stop()
		defer cancel()
		for {
			select {
			case <-newCtx.Done():
				s.log.WithContext(ctx).Infow("msg", "RenewalNodeID stopped")
				return
			case <-ticker.C:
				resp, err := s.renewalNodeID(newCtx, req)
				if err != nil {
					if s.log != nil {
						s.log.WithContext(ctx).Warnw("msg", "RenewalNodeID failed", "error", err)
					}
					break
				}
				//req.AccessToken = resp.AccessToken
				_ = resp
			}
		}
	})

	// 新的 AccessToken
	dataModel.AccessToken = req.AccessToken
	releaseFunc := func(releaseContext context.Context) error {
		_, err := s.ReleaseNodeId(releaseContext, dataModel)
		if err != nil {
			return err
		}
		return nil
	}
	nodeID := &nodeIDInstance{
		stopRenewal: cancel,
		release:     releaseFunc,
	}
	return nodeID, nil
}

func (s *grpcAPI) renewalNodeID(ctx context.Context, req *nodeidresourcev1.RenewalNodeIdReq) (*nodeidresourcev1.RenewalNodeIdRespData, error) {
	resp, err := s.client.RenewalNodeId(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		if s.opts.tries <= 1 {
			return nil, errorpkg.WithStack(e)
		}
		if s.log != nil {
			s.log.WithContext(ctx).Warnw("msg", "RenewalNodeId failed", "try_number", 1, "error", e)
		}
	} else {
		return resp.Data, nil
	}
	for i := 2; i <= s.opts.tries; i++ {
		time.Sleep(s.opts.retryDelay)
		resp, err = s.client.RenewalNodeId(ctx, req)
		if e := apiutil.CheckAPIResponse(resp, err); e != nil {
			if s.opts.tries <= i {
				return nil, errorpkg.WithStack(e)
			}
			if s.log != nil {
				s.log.WithContext(ctx).Warnw("msg", "RenewalNodeId failed", "try_number", 1, "error", e)
			}
		} else {
			return resp.Data, nil
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

func (s *grpcAPI) ReleaseNodeId(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (*nodeidresourcev1.ReleaseNodeIdRespData, error) {
	req := &nodeidresourcev1.ReleaseNodeIdReq{
		InstanceId:  dataModel.InstanceId,
		NodeId:      dataModel.NodeId,
		AccessToken: dataModel.AccessToken,
	}
	resp, err := s.client.ReleaseNodeId(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		if s.opts.tries <= 1 {
			return nil, errorpkg.WithStack(e)
		}
		if s.log != nil {
			s.log.WithContext(ctx).Warnw("msg", "ReleaseNodeId failed", "try_number", 1, "error", e)
		}
	} else {
		return resp.Data, nil
	}
	for i := 2; i <= s.opts.tries; i++ {
		time.Sleep(s.opts.retryDelay)
		resp, err = s.client.ReleaseNodeId(ctx, req)
		if e := apiutil.CheckAPIResponse(resp, err); e != nil {
			if s.opts.tries <= i {
				return nil, errorpkg.WithStack(e)
			}
			if s.log != nil {
				s.log.WithContext(ctx).Warnw("msg", "ReleaseNodeId failed", "try_number", 1, "error", e)
			}
		} else {
			return resp.Data, nil
		}
	}
	if err == nil {
		e := errorpkg.ErrorInternalServer("ReleaseNodeId failed")
		return nil, errorpkg.WithStack(e)
	}
	return nil, errorpkg.FormatError(err)
}
