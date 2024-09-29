package nodeidapi

import (
	"context"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
)

type NodeIDAPI interface {
	GetAndRenewalNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, RenewalInterface, error)
	GetNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error)
	RenewalNodeID(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (RenewalInterface, error)
}

type RenewalInterface interface {
	Stop() error
}

type stopRenewal struct {
	cancel func()
}

func (r *stopRenewal) Stop() error {
	if r.cancel != nil {
		r.cancel()
	}
	return nil
}
