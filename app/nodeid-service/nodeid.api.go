package nodeidapi

import (
	"context"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	"time"
)

type NodeIDAPI interface {
	GetNodeId(context.Context, *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error)
	ReleaseNodeId(context.Context, *nodeidresourcev1.ReleaseNodeIdReq) (*nodeidresourcev1.ReleaseNodeIdRespData, error)
	RenewalNodeId(context.Context, *nodeidresourcev1.RenewalNodeIdReq) (*nodeidresourcev1.RenewalNodeIdRespData, error)
}

type NodeIDHelper interface {
	GetAndAutoRenewalNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, RenewalManager, error)
	GetNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error)
	RenewalNodeID(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (RenewalManager, error)
	ReleaseNodeId(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (*nodeidresourcev1.ReleaseNodeIdRespData, error)
}

type NodeID interface {
	Release(ctx context.Context) error
}

type RenewalManager interface {
	Stop(ctx context.Context) error
	RenewalResult(ctx context.Context) *RenewalResult
}

type RenewalResult struct {
	Data     *nodeidresourcev1.RenewalNodeIdRespData
	Err      error
	LastTime time.Time
}

type renewalManager struct {
	data *RenewalResult
	stop func()
}

func (s *renewalManager) Stop(_ context.Context) error {
	if s.stop != nil {
		s.stop()
	}
	return nil
}

func (s *renewalManager) RenewalResult(_ context.Context) *RenewalResult {
	return s.data
}
