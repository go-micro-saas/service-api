package pingapi

import (
	"context"
	pingroucesv1 "github.com/go-micro-saas/service-api/api/ping-service/v1/resources"
	pingservicev1 "github.com/go-micro-saas/service-api/api/ping-service/v1/services"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
)

const (
	PingServiceHTTP clientutil.ServiceName = "ping-service-http"
	PingServiceGRPC clientutil.ServiceName = "ping-service-grpc"
)

// PingAPI 仅供参考，建议在程序中直接使用 client.Ping()，然后
// 使用 apiutil.CheckAPIResponse(resp, err) 进行错误处理
type PingAPI interface {
	Ping(ctx context.Context, req *pingroucesv1.PingReq) (*pingroucesv1.PingRespData, error)
}

// NewGRPCClient ...
func NewGRPCClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (pingservicev1.SrvPingClient, error) {
	serviceName := PingServiceGRPC
	for i := range rewriteServiceName {
		serviceName = rewriteServiceName[i]
	}
	conn, err := clientutil.NewSingletonServiceAPIConnection(serviceAPIManager, serviceName)
	//conn, err := NewServiceAPIConnection(serviceAPIManager, serviceName)
	if err != nil {
		return nil, err
	}
	grpcConn, err := conn.GetGRPCConnection()
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingClient(grpcConn), nil
}

// NewHTTPClient ...
func NewHTTPClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
	serviceName := PingServiceHTTP
	for i := range rewriteServiceName {
		serviceName = rewriteServiceName[i]
	}
	conn, err := clientutil.NewSingletonServiceAPIConnection(serviceAPIManager, serviceName)
	//conn, err := NewServiceAPIConnection(serviceAPIManager, serviceName)
	if err != nil {
		return nil, err
	}
	httpClient, err := conn.GetHTTPClient()
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingHTTPClient(httpClient), nil
}
