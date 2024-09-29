package nodeidapi

import (
	"context"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
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
	Data(ctx context.Context) <-chan *nodeidresourcev1.RenewalNodeIdRespData
}

type renewalManager struct {
	data <-chan *nodeidresourcev1.RenewalNodeIdRespData
	stop func()
}

func (s *renewalManager) Stop(_ context.Context) error {
	if s.stop != nil {
		s.stop()
	}
	return nil
}

func (s *renewalManager) Data(_ context.Context) <-chan *nodeidresourcev1.RenewalNodeIdRespData {
	return s.data
}
