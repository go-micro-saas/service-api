package snowflakeapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	nodeidapi "github.com/go-micro-saas/service-api/app/nodeid-service"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var (
	// NewSnowflakeNode func(int64) (idpkg.Snowflake, error)
	// idpkg.NewSonySonyflake
	// idpkg.NewBwmarrinSnowflake
	NewSnowflakeNode = idpkg.NewSonySonyflake
)

type IDManager interface {
	GetSingletonSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error)
	GetSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error)
	SetSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error)
}

type idManager struct {
	log    *log.Helper
	helper nodeidapi.NodeIDHelper

	rwMutex           sync.RWMutex
	singletonMutexMap map[string]*sync.Once
	singletonNodeMap  map[string]*SnowflakeNodeInfo
}

type SnowflakeNodeInfo struct {
	Node    idpkg.Snowflake
	Cleanup func()
}

func NewIDManager(logger log.Logger, helper nodeidapi.NodeIDHelper) IDManager {
	logHelper := log.NewHelper(log.With(logger, "module", "nodeid-api/id-manager"))
	return &idManager{
		log:    logHelper,
		helper: helper,

		singletonMutexMap: make(map[string]*sync.Once),
		singletonNodeMap:  make(map[string]*SnowflakeNodeInfo),
	}
}

func (s *idManager) GetSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error) {
	nodeID, renewal, err := s.helper.GetAndAutoRenewalNodeID(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	if nodeID.GetNodeEpoch().GetSeconds() > 0 {
		idpkg.DefaultEpoch = nodeID.GetNodeEpoch().AsTime()
	}
	node, err := NewSnowflakeNode(nodeID.NodeId)
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

func (s *idManager) GetSingletonSnowflakeNode(ctx context.Context, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error) {
	s.rwMutex.RLock()
	_, ok := s.singletonMutexMap[req.InstanceId]
	s.rwMutex.RUnlock()
	if !ok {
		s.rwMutex.Lock()
		s.singletonMutexMap[req.InstanceId] = &sync.Once{}
		s.singletonNodeMap[req.InstanceId] = &SnowflakeNodeInfo{}
		s.rwMutex.Unlock()
	}
	var (
		err error
	)
	s.singletonMutexMap[req.InstanceId].Do(func() {
		s.singletonNodeMap[req.InstanceId].Node, s.singletonNodeMap[req.InstanceId].Cleanup, err = s.GetSnowflakeNode(ctx, req)
	})
	if err != nil {
		s.rwMutex.Lock()
		s.singletonMutexMap[req.InstanceId] = &sync.Once{}
		s.rwMutex.Unlock()
	}
	return s.singletonNodeMap[req.InstanceId].Node, s.singletonNodeMap[req.InstanceId].Cleanup, err
}
