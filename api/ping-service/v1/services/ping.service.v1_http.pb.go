// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v3.21.6
// source: api/ping-service/v1/services/ping.service.v1.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/go-micro-saas/service-api/api/ping-service/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvPingPing = "/saas.api.ping.servicev1.SrvPing/Ping"

type SrvPingHTTPServer interface {
	// Ping Ping ping
	//
	// 测试PingV1
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
}

func RegisterSrvPingHTTPServer(s *http.Server, srv SrvPingHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/ping/{message}", _SrvPing_Ping0_HTTP_Handler(srv))
}

func _SrvPing_Ping0_HTTP_Handler(srv SrvPingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.PingReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvPingPing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*resources.PingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.PingResp)
		return ctx.Result(200, reply)
	}
}

type SrvPingHTTPClient interface {
	Ping(ctx context.Context, req *resources.PingReq, opts ...http.CallOption) (rsp *resources.PingResp, err error)
}

type SrvPingHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvPingHTTPClient(client *http.Client) SrvPingHTTPClient {
	return &SrvPingHTTPClientImpl{client}
}

func (c *SrvPingHTTPClientImpl) Ping(ctx context.Context, in *resources.PingReq, opts ...http.CallOption) (*resources.PingResp, error) {
	var out resources.PingResp
	pattern := "/api/v1/ping/{message}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvPingPing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
