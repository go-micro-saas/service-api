package nodeidapi

import (
	"context"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	nodeidservicev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/services"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type httpAPI struct {
	client nodeidservicev1.SrvNodeIDV1HTTPClient
}

func NewHTTPApi(client nodeidservicev1.SrvNodeIDV1HTTPClient) NodeIDAPI {
	return &httpAPI{client: client}
}

func (s *httpAPI) GetNodeId(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error) {
	resp, err := s.client.GetNodeId(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.Data, nil
}

func (s *httpAPI) ReleaseNodeId(ctx context.Context, req *nodeidresourcev1.ReleaseNodeIdReq) (*nodeidresourcev1.ReleaseNodeIdRespData, error) {
	resp, err := s.client.ReleaseNodeId(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.Data, nil
}

func (s *httpAPI) RenewalNodeId(ctx context.Context, req *nodeidresourcev1.RenewalNodeIdReq) (*nodeidresourcev1.RenewalNodeIdRespData, error) {
	resp, err := s.client.RenewalNodeId(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.Data, nil
}
