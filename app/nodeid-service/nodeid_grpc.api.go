package nodeidapi

import (
	"context"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	nodeidservicev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/services"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type grpcAPI struct {
	client nodeidservicev1.SrvNodeIDV1Client
}

func NewGRPCApi(client nodeidservicev1.SrvNodeIDV1Client) NodeIDAPI {
	return &grpcAPI{client: client}
}

func (s *grpcAPI) GetNodeId(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error) {
	resp, err := s.client.GetNodeId(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.Data, nil
}

func (s *grpcAPI) ReleaseNodeId(ctx context.Context, req *nodeidresourcev1.ReleaseNodeIdReq) (*nodeidresourcev1.ReleaseNodeIdRespData, error) {
	resp, err := s.client.ReleaseNodeId(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.Data, nil
}

func (s *grpcAPI) RenewalNodeId(ctx context.Context, req *nodeidresourcev1.RenewalNodeIdReq) (*nodeidresourcev1.RenewalNodeIdRespData, error) {
	resp, err := s.client.RenewalNodeId(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.Data, nil
}
