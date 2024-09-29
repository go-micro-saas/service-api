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
	GetAndAutoRenewalNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error)
	GetNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error)
	RenewalNodeID(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) error
	ReleaseNodeId(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (*nodeidresourcev1.ReleaseNodeIdRespData, error)
}

type NodeID interface {
	Release(ctx context.Context) error
}

type nodeIDInstance struct {
	nodeID      *nodeidresourcev1.GetNodeIdRespData
	stopRenewal func()
	release     func(ctx context.Context) error
}

func (s *nodeIDInstance) StopRenewal(_ context.Context) error {
	if s.stopRenewal != nil {
		s.stopRenewal()
	}
	return nil
}

func (s *nodeIDInstance) Release(ctx context.Context) error {
	_ = s.StopRenewal(ctx)
	if s.release != nil {
		return s.release(ctx)
	}
	return nil
}
