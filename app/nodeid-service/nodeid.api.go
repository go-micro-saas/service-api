package nodeidapi

import (
	"context"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	nodeidservicev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/services"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	"time"
)

const (
	NodeidServiceHTTP clientutil.ServiceName = "nodeid-service-http"
	NodeidServiceGRPC clientutil.ServiceName = "nodeid-service-grpc"
)

type NodeIDAPI interface {
	GetNodeId(context.Context, *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error)
	ReleaseNodeId(context.Context, *nodeidresourcev1.ReleaseNodeIdReq) (*nodeidresourcev1.ReleaseNodeIdRespData, error)
	RenewalNodeId(context.Context, *nodeidresourcev1.RenewalNodeIdReq) (*nodeidresourcev1.RenewalNodeIdRespData, error)
}

type NodeIDHelper interface {
	GetAndAutoRenewalNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, RenewalManager, error)
	GetNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error)
	RenewalNodeID(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (RenewalManager, error)
	ReleaseNodeId(ctx context.Context, dataModel *nodeidresourcev1.GetNodeIdRespData) (*nodeidresourcev1.ReleaseNodeIdRespData, error)
}

type RenewalManager interface {
	Stop(ctx context.Context) error
	RenewalResult(ctx context.Context) *RenewalResult
}

type RenewalResult struct {
	Data     *nodeidresourcev1.RenewalNodeIdRespData
	Err      error
	LastTime time.Time
}

type renewalManager struct {
	data *RenewalResult
	stop func()
}

func (s *renewalManager) Stop(_ context.Context) error {
	if s.stop != nil {
		s.stop()
	}
	return nil
}

func (s *renewalManager) RenewalResult(_ context.Context) *RenewalResult {
	return s.data
}

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

// ===== GRPC =====

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

// ===== GRPC =====

// ===== HTTP =====

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

// ===== HTTP =====
