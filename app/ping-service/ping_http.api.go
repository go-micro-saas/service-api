package pingapi

import (
	"context"
	pingroucesv1 "github.com/go-micro-saas/service-api/api/ping-service/v1/resources"
	pingservicev1 "github.com/go-micro-saas/service-api/api/ping-service/v1/services"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type httpAPI struct {
	client pingservicev1.SrvPingHTTPClient
}

func NewPingHTTPAPI(client pingservicev1.SrvPingHTTPClient) PingAPI {
	return &httpAPI{client: client}
}

func (s *httpAPI) Ping(ctx context.Context, req *pingroucesv1.PingReq) (*pingroucesv1.PingRespData, error) {
	resp, err := s.client.Ping(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.Data, nil
}
