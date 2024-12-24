package snowflakeapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	nodeidapi "github.com/go-micro-saas/service-api/app/nodeid-service"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	"sync"
	"time"
)

var (
	_httpIDManager     IDManager
	_httpIDManagerOnce sync.Once
	_grpcIDManager     IDManager
	_grpcIDManagerOnce sync.Once

	DefaultTries = 3
)

func DefaultOptions(logger log.Logger) []Option {
	return []Option{
		WithLogger(logger),
	}
}

// GetSingletonSnowflakeNode 获取节点
// http GetSingletonIdGeneratorByHTTPAPI
// grpc GetSingletonIdGeneratorByGRPCAPI
func GetSingletonSnowflakeNode(idManager IDManager, req *nodeidresourcev1.GetNodeIdReq) (idpkg.Snowflake, func(), error) {
	return idManager.GetSingletonSnowflakeNode(context.Background(), req)
}

func SetSnowflake(node idpkg.Snowflake) error {
	idpkg.SetNode(node)
	return nil
}

type options struct {
	logger              log.Logger
	serverName          clientutil.ServiceName
	isGetNodeIDFromIPV4 bool
	mustGetNodeIdForAPI bool
	nodeEpoch           time.Time
}

type Option func(*options)

func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func WithServerName(serverName clientutil.ServiceName) Option {
	return func(o *options) {
		o.serverName = serverName
	}
}

func WithGetNodeIdFromIPV4(isGetNodeIDFromIPV4 bool) Option {
	return func(o *options) {
		o.isGetNodeIDFromIPV4 = isGetNodeIDFromIPV4
	}
}

func WithMustGetNodeIdFromAPI(mustGetNodeIdForAPI bool) Option {
	return func(o *options) {
		o.mustGetNodeIdForAPI = mustGetNodeIdForAPI
	}
}

func WithNodeEpoch(nodeEpoch time.Time) Option {
	return func(o *options) {
		o.nodeEpoch = nodeEpoch
	}
}

func GetSingletonIdGeneratorByHTTPAPI(serviceAPIManager clientutil.ServiceAPIManager, req *nodeidresourcev1.GetNodeIdReq, opts ...Option) (idpkg.Snowflake, func(), error) {
	opt := options{}
	opt.logger, _ = logpkg.NewDummyLogger()
	for _, o := range opts {
		o(&opt)
	}
	if opt.isGetNodeIDFromIPV4 {
		return GetIdGeneratorFromIPV4()
	}

	serverName := nodeidapi.NodeidServiceHTTP
	if opt.serverName != "" {
		serverName = opt.serverName
	}
	client, err := nodeidapi.NewHTTPClient(serviceAPIManager, serverName)
	if err != nil {
		return nil, nil, err
	}
	_httpIDManagerOnce.Do(func() {
		_httpIDManager, err = getIDManager(nodeidapi.NewHTTPApi(client), &opt)
	})
	if err != nil {
		_httpIDManager = nil
		_httpIDManagerOnce = sync.Once{}
		return nil, nil, err
	}
	return getIdGeneratorFromAPI(_httpIDManager, req, &opt)
}

func GetSingletonIdGeneratorByGRPCAPI(serviceAPIManager clientutil.ServiceAPIManager, req *nodeidresourcev1.GetNodeIdReq, opts ...Option) (idpkg.Snowflake, func(), error) {
	opt := options{}
	opt.logger, _ = logpkg.NewDummyLogger()
	for _, o := range opts {
		o(&opt)
	}
	if opt.isGetNodeIDFromIPV4 {
		return GetIdGeneratorFromIPV4()
	}

	serverName := nodeidapi.NodeidServiceGRPC
	if opt.serverName != "" {
		serverName = opt.serverName
	}
	client, err := nodeidapi.NewGRPCClient(serviceAPIManager, serverName)
	if err != nil {
		return nil, nil, err
	}
	_grpcIDManagerOnce.Do(func() {
		_grpcIDManager, err = getIDManager(nodeidapi.NewGRPCApi(client), &opt)
	})
	if err != nil {
		_grpcIDManager = nil
		_grpcIDManagerOnce = sync.Once{}
		return nil, nil, err
	}
	return getIdGeneratorFromAPI(_grpcIDManager, req, &opt)
}

func getIDManager(nodeidAPI nodeidapi.NodeIDAPI, opt *options) (IDManager, error) {
	var (
		helperOpts = []nodeidapi.Option{
			nodeidapi.WithTries(DefaultTries),
			nodeidapi.WithRetryDelay(nodeidapi.DefaultRetryDelay),
			nodeidapi.WithHeartbeatInterval(nodeidapi.DefaultHeartbeatInterval),
			nodeidapi.WithLogger(opt.logger),
		}
	)
	helper, err := nodeidapi.NewNodeIDHelper(nodeidAPI, helperOpts...)
	if err != nil {
		return nil, err
	}
	return NewIDManager(opt.logger, helper), nil
}

func getIdGeneratorFromAPI(mgr IDManager, req *nodeidresourcev1.GetNodeIdReq, opt *options) (idpkg.Snowflake, func(), error) {
	var (
		logHelper = log.NewHelper(log.With(opt.logger, "module", "nodeid-api/id-generator"))
	)

	ctx := context.Background()
	node, cleanup, err := mgr.GetSingletonSnowflakeNode(ctx, req)
	if err != nil {
		if opt.mustGetNodeIdForAPI {
			return nil, nil, err
		}
		logHelper.WithContext(ctx).Warnw("msg", "GetSingletonSnowflakeNode failed", "error", err)
		err = nil
		node, cleanup, err = GetIdGeneratorFromIPV4()
	}
	if node != nil {
		idpkg.SetNode(node)
	}
	return node, cleanup, nil
}

func GetIdGeneratorFromIPV4(opts ...Option) (idpkg.Snowflake, func(), error) {
	opt := options{}
	opt.logger, _ = logpkg.NewDummyLogger()
	for _, o := range opts {
		o(&opt)
	}
	if !opt.nodeEpoch.IsZero() {
		idpkg.DefaultEpoch = opt.nodeEpoch
	}
	nodeID, _ := idpkg.GenNodeID()
	if nodeID < 1 {
		nodeID = 1
	}
	node, err := idpkg.NewBwmarrinSnowflake(int64(nodeID))
	if err != nil {
		return nil, nil, err
	}
	idpkg.SetNode(node)
	cleanup := func() {}
	return node, cleanup, nil
}
