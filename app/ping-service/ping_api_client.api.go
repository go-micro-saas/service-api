package pingapi

import (
	clientutil "github.com/go-micro-saas/service-kit/cluster_service_api"
	pingservicev1 "github.com/go-micro-saas/service-kit/testdata/ping-service/api/ping-service/v1/services"
)

const (
	PingServiceHTTP clientutil.ServiceName = "ping-service-http"
	PingServiceGRPC clientutil.ServiceName = "ping-service-grpc"
)

// NewPingGRPCClient ...
func NewPingGRPCClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (pingservicev1.SrvPingClient, error) {
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

// NewPingHTTPClient ...
func NewPingHTTPClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
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
