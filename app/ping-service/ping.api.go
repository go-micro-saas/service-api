package pingapi

import (
	"context"
	pingroucesv1 "github.com/go-micro-saas/service-api/api/ping-service/v1/resources"
	pingservicev1 "github.com/go-micro-saas/service-api/api/ping-service/v1/services"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// PingAPI 仅供参考，建议在程序中直接使用 client.Ping()，然后
// 使用 apiutil.CheckAPIResponse(resp, err) 进行错误处理
type PingAPI interface {
	GetClient() interface{}
	Ping(ctx context.Context, req *pingroucesv1.PingReq) (*pingroucesv1.PingRespData, error)
}

type httpAPI struct {
	client pingservicev1.SrvPingHTTPClient
}

func NewPingHTTPAPI(client pingservicev1.SrvPingHTTPClient) PingAPI {
	return &httpAPI{client: client}
}

func (s *httpAPI) GetClient() interface{} {
	return s.client
}

func (s *httpAPI) Ping(ctx context.Context, req *pingroucesv1.PingReq) (*pingroucesv1.PingRespData, error) {
	resp, err := s.client.Ping(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.Data, nil
}

type grpcAPI struct {
	client pingservicev1.SrvPingClient
}

func NewPingGRPCAPI(client pingservicev1.SrvPingClient) PingAPI {
	return &grpcAPI{client: client}
}

func (s *grpcAPI) GetClient() interface{} {
	return s.client
}

func (s *grpcAPI) Ping(ctx context.Context, req *pingroucesv1.PingReq) (*pingroucesv1.PingRespData, error) {
	resp, err := s.client.Ping(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.Data, nil
}
