package accountapi

import (
	accountservicev1 "github.com/go-micro-saas/service-api/api/account-service/v1/services"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
)

const (
	AccountServiceHTTP clientutil.ServiceName = "account-service-http"
	AccountServiceGRPC clientutil.ServiceName = "account-service-grpc"
)

// NewAccountV1GRPCClient ...
func NewAccountV1GRPCClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (accountservicev1.SrvAccountV1Client, error) {
	serviceName := AccountServiceGRPC
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
	return accountservicev1.NewSrvAccountV1Client(grpcConn), nil
}

// NewAccountV1HTTPClient ...
func NewAccountV1HTTPClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (accountservicev1.SrvAccountV1HTTPClient, error) {
	serviceName := AccountServiceHTTP
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
	return accountservicev1.NewSrvAccountV1HTTPClient(httpClient), nil
}

// NewUserAuthV1GRPCClient ...
func NewUserAuthV1GRPCClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (accountservicev1.SrvUserAuthV1Client, error) {
	serviceName := AccountServiceGRPC
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
	return accountservicev1.NewSrvUserAuthV1Client(grpcConn), nil
}

// NewUserAuthV1HTTPClient ...
func NewUserAuthV1HTTPClient(serviceAPIManager clientutil.ServiceAPIManager, rewriteServiceName ...clientutil.ServiceName) (accountservicev1.SrvUserAuthV1HTTPClient, error) {
	serviceName := AccountServiceHTTP
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
	return accountservicev1.NewSrvUserAuthV1HTTPClient(httpClient), nil
}
