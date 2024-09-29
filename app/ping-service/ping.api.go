package pingapi

import (
	"context"
	pingroucesv1 "github.com/go-micro-saas/service-api/api/ping-service/v1/resources"
)

// PingAPI 仅供参考，建议在程序中直接使用 client.Ping()，然后
// 使用 apiutil.CheckAPIResponse(resp, err) 进行错误处理
type PingAPI interface {
	Ping(ctx context.Context, req *pingroucesv1.PingReq) (*pingroucesv1.PingRespData, error)
}
