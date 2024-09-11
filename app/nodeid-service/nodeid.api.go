package nodeidapi

import (
	"context"
	nodeiderrorv1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/errors"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	nodeidservicev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/services"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type NodeIDAPI interface {
}

type grpcAPI struct {
	client nodeidservicev1.SrvNodeIDV1Client
}

func NewNodeIDAPI(client nodeidservicev1.SrvNodeIDV1Client) NodeIDAPI {
	return &grpcAPI{client: client}
}

func (s *grpcAPI) GetAndRenewalNodeID(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (*nodeidresourcev1.GetNodeIdRespData, error) {
	resp, err := s.client.GetNodeId(ctx, req)
	if err = apiutil.CheckAPIResponse(resp, err); err != nil {
		return nil, errorpkg.FormatError(err)
	}

	// token错误：ID被其他程序占用
	var tokenIncorrect = nodeiderrorv1.ERROR_S102_ACCESS_TOKEN_INCORRECT

	renewalReq := &nodeidresourcev1.RenewalNodeIdReq{NodeId: resp.Data.NodeId}
	renewalResp, err := s.client.RenewalNodeId(ctx, renewalReq)
	if err != nil {
		if apiutil.IsReason(err, tokenIncorrect) {

		}
	}
	if e := apiutil.CheckResponseCode(renewalResp); e != nil {
		if e.Reason == tokenIncorrect.String() {

		}
	}

	return resp.Data, nil
}
