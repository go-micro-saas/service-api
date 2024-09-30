package nodeidapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

type IDManager interface {
	SetNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) error
	Release(ctx context.Context) error
}

type idManager struct {
	log    *log.Helper
	helper NodeIDHelper

	setOnce        sync.Once
	nodeID         *nodeidresourcev1.GetNodeIdRespData
	renewalManager RenewalManager
}

func NewIDManager(logger log.Logger, helper NodeIDHelper) IDManager {
	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-api/id-manager"))
	return &idManager{
		log:    logHelper,
		helper: helper,
	}
}

func (s *idManager) SetNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) error {
	var err error
	s.setOnce.Do(func() {
		err = s.setNode(ctx, req)
	})
	if err != nil {
		s.setOnce = sync.Once{}
		return err
	}
	return nil
}

func (s *idManager) setNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) error {
	nodeID, renewal, err := s.helper.GetAndAutoRenewalNodeID(ctx, req)
	if err != nil {
		return err
	}
	node, err := idpkg.NewBwmarrinSnowflake(nodeID.NodeId)
	if err != nil {
		_ = renewal.Stop(ctx)
		e := errorpkg.ErrorInternalError(err.Error())
		return errorpkg.WithStack(e)
	}
	s.log.WithContext(ctx).Infow("msg", "set node id", "node_id", nodeID)
	s.nodeID = nodeID
	s.renewalManager = renewal
	idpkg.SetNode(node)
	return nil
}

func (s *idManager) Release(ctx context.Context) error {
	if s.renewalManager != nil {
		err := s.renewalManager.Stop(ctx)
		if err != nil {
			s.log.WithContext(ctx).Warnw("msg", "release node id", "err", err)
			err = nil
		}
	}
	_, err := s.helper.ReleaseNodeId(ctx, s.nodeID)
	if err != nil {
		return err
	}
	return nil
}
