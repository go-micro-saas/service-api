package orgapi

import (
	orgservicev1 "github.com/go-micro-saas/service-api/api/org-service/v1/services"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
)

const (
	OrgServiceHTTP clientutil.ServiceName = "org-service-http"
	OrgServiceGRPC clientutil.ServiceName = "org-service-grpc"
)

// NewOrgV1GRPCClient ...
func NewOrgV1GRPCClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (orgservicev1.SrvOrgV1Client, error) {
	serviceName := OrgServiceGRPC
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
	return orgservicev1.NewSrvOrgV1Client(grpcConn), nil
}

// NewOrgV1HTTPClient ...
func NewOrgV1HTTPClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (orgservicev1.SrvOrgV1HTTPClient, error) {
	serviceName := OrgServiceHTTP
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
	return orgservicev1.NewSrvOrgV1HTTPClient(httpClient), nil
}
