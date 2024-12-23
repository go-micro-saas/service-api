// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v3.21.6
// source: api/uuid-service/v1/services/uuid.service.v1.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/go-micro-saas/service-api/api/uuid-service/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvUuidV1NextID = "/saas.api.uuid.servicev1.SrvUuidV1/NextID"

type SrvUuidV1HTTPServer interface {
	// NextID 获取ID
	NextID(context.Context, *resources.NextIDReq) (*resources.NextIDResp, error)
}

func RegisterSrvUuidV1HTTPServer(s *http.Server, srv SrvUuidV1HTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/uuid/next-id", _SrvUuidV1_NextID0_HTTP_Handler(srv))
}

func _SrvUuidV1_NextID0_HTTP_Handler(srv SrvUuidV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.NextIDReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvUuidV1NextID)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NextID(ctx, req.(*resources.NextIDReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.NextIDResp)
		return ctx.Result(200, reply)
	}
}

type SrvUuidV1HTTPClient interface {
	NextID(ctx context.Context, req *resources.NextIDReq, opts ...http.CallOption) (rsp *resources.NextIDResp, err error)
}

type SrvUuidV1HTTPClientImpl struct {
	cc *http.Client
}

func NewSrvUuidV1HTTPClient(client *http.Client) SrvUuidV1HTTPClient {
	return &SrvUuidV1HTTPClientImpl{client}
}

func (c *SrvUuidV1HTTPClientImpl) NextID(ctx context.Context, in *resources.NextIDReq, opts ...http.CallOption) (*resources.NextIDResp, error) {
	var out resources.NextIDResp
	pattern := "/api/v1/uuid/next-id"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvUuidV1NextID))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
