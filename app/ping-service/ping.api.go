package pingapi

import (
	"context"
	pingroucesv1 "github.com/go-micro-saas/service-api/api/ping-service/v1/resources"
	pingservicev1 "github.com/go-micro-saas/service-api/api/ping-service/v1/services"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type PingAPI interface {
	Ping(ctx context.Context, req *pingroucesv1.PingReq) (*pingroucesv1.PingRespData, error)
}

type httpAPI struct {
	client pingservicev1.SrvPingHTTPClient
}

func NewPingHTTPAPI(client pingservicev1.SrvPingHTTPClient) PingAPI {
	return &httpAPI{client: client}
}

func (s *httpAPI) Ping(ctx context.Context, req *pingroucesv1.PingReq) (*pingroucesv1.PingRespData, error) {
	resp, err := s.client.Ping(ctx, req)
	if err = apiutil.CheckAPIResponse(resp, err); err != nil {
		return nil, errorpkg.FormatError(err)
	}
	return resp.Data, nil
}

type grpcAPI struct {
	client pingservicev1.SrvPingClient
}

func NewPingGRPCAPI(client pingservicev1.SrvPingClient) PingAPI {
	return &grpcAPI{client: client}
}

func (s *grpcAPI) Ping(ctx context.Context, req *pingroucesv1.PingReq) (*pingroucesv1.PingRespData, error) {
	resp, err := s.client.Ping(ctx, req)
	if err = apiutil.CheckAPIResponse(resp, err); err != nil {
		return nil, errorpkg.FormatError(err)
	}
	return resp.Data, nil
}
