package snowflakeapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	nodeidapi "github.com/go-micro-saas/service-api/app/nodeid-service"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	"time"
)

var (
	DefaultTries = 3
)

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

func Options(logger log.Logger) []Option {
	return []Option{
		WithLogger(logger),
	}
}

func GetIdGeneratorByHTTPAPI(serviceAPIManager clientutil.ServiceAPIManager, req *nodeidresourcev1.GetNodeIdReq, opts ...Option) (idpkg.Snowflake, func(), error) {
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
	return GetIdGeneratorFromAPI(nodeidapi.NewHTTPApi(client), req, &opt)
}

func GetIdGeneratorByGRPCAPI(serviceAPIManager clientutil.ServiceAPIManager, req *nodeidresourcev1.GetNodeIdReq, opts ...Option) (idpkg.Snowflake, func(), error) {
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
	return GetIdGeneratorFromAPI(nodeidapi.NewGRPCApi(client), req, &opt)
}

func GetIdGeneratorFromAPI(nodeidAPI nodeidapi.NodeIDAPI, req *nodeidresourcev1.GetNodeIdReq, opt *options) (idpkg.Snowflake, func(), error) {
	var (
		logHelper  = log.NewHelper(log.With(opt.logger, "module", "nodeid-api/id-generator"))
		helperOpts = []nodeidapi.Option{
			nodeidapi.WithTries(DefaultTries),
			nodeidapi.WithRetryDelay(nodeidapi.DefaultRetryDelay),
			nodeidapi.WithHeartbeatInterval(nodeidapi.DefaultHeartbeatInterval),
			nodeidapi.WithLogger(opt.logger),
		}
	)
	helper := nodeidapi.NewNodeIDHelper(nodeidAPI, helperOpts...)
	mgr := NewIDManager(opt.logger, helper)

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
