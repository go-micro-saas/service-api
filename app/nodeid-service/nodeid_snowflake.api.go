package nodeidapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func SetSnowflake(node idpkg.Snowflake) error {
	idpkg.SetNode(node)
	return nil
}

type IDManager interface {
	GetSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error)
	SetSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error)
}

type idManager struct {
	log    *log.Helper
	helper NodeIDHelper
}

func NewIDManager(logger log.Logger, helper NodeIDHelper) IDManager {
	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-api/id-manager"))
	return &idManager{
		log:    logHelper,
		helper: helper,
	}
}

func (s *idManager) GetSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error) {
	nodeID, renewal, err := s.helper.GetAndAutoRenewalNodeID(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	node, err := idpkg.NewBwmarrinSnowflake(nodeID.NodeId)
	if err != nil {
		_ = renewal.Stop(ctx)
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, nil, errorpkg.WithStack(e)
	}
	cleanup := func() {
		cleanContext := context.Background()
		cleanupErr := renewal.Stop(cleanContext)
		if cleanupErr != nil {
			s.log.WithContext(ctx).Warnw("msg", "[release] stop renewal failed", "err", cleanupErr)
		}
		_, cleanupErr = s.helper.ReleaseNodeId(cleanContext, nodeID)
		if cleanupErr != nil {
			s.log.WithContext(ctx).Warnw("msg", "[release] release node failed", "err", cleanupErr)
		}
	}
	return node, cleanup, nil
}

func (s *idManager) SetSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error) {
	node, cleanup, err := s.GetSnowflakeNode(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	s.log.WithContext(ctx).Infow("msg", "set node id", "node_id", node)
	idpkg.SetNode(node)
	return node, cleanup, err
}
