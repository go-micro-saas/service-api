package nodeidapi

import (
	nodeidservicev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/services"
	clientutil "github.com/go-micro-saas/service-kit/cluster_service_api"
)

const (
	NodeidServiceHTTP clientutil.ServiceName = "nodeid-service-http"
	NodeidServiceGRPC clientutil.ServiceName = "nodeid-service-grpc"
)

// NewGRPCClient ...
func NewGRPCClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (nodeidservicev1.SrvNodeIDV1Client, error) {
	serviceName := NodeidServiceGRPC
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
	return nodeidservicev1.NewSrvNodeIDV1Client(grpcConn), nil
}

// NewHTTPClient ...
func NewHTTPClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (nodeidservicev1.SrvNodeIDV1HTTPClient, error) {
	serviceName := NodeidServiceHTTP
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
	return nodeidservicev1.NewSrvNodeIDV1HTTPClient(httpClient), nil
}
