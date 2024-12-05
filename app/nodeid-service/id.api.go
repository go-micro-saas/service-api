package nodeidapi

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
)

var (
	DefaultTries = 3
)

type idOptions struct {
	logger              log.Logger
	serverName          clientutil.ServiceName
	mustGetNodeIdForAPI bool
}
type IdOption func(*idOptions)

func WithLoggerForIdOption(logger log.Logger) IdOption {
	return func(o *idOptions) {
		o.logger = logger
	}
}

func WithServerName(serverName clientutil.ServiceName) IdOption {
	return func(o *idOptions) {
		o.serverName = serverName
	}
}

func WithMustGetNodeIdForAPI(mustGetNodeIdForAPI bool) IdOption {
	return func(o *idOptions) {
		o.mustGetNodeIdForAPI = mustGetNodeIdForAPI
	}
}

func GetIdGeneratorByHTTPAPI(serviceAPIManager clientutil.ServiceAPIManager, req *nodeidresourcev1.GetNodeIdReq, opts ...IdOption) (idpkg.Snowflake, func(), error) {
	opt := idOptions{}
	opt.logger, _ = logpkg.NewDummyLogger()
	for _, o := range opts {
		o(&opt)
	}

	serverName := NodeidServiceHTTP
	if opt.serverName != "" {
		serverName = opt.serverName
	}
	client, err := NewHTTPClient(serviceAPIManager, serverName)
	if err != nil {
		return nil, nil, err
	}
	return getIdGenerator(NewHTTPApi(client), req, &opt)
}

func GetIdGeneratorByGRPCAPI(serviceAPIManager clientutil.ServiceAPIManager, req *nodeidresourcev1.GetNodeIdReq, opts ...IdOption) (idpkg.Snowflake, func(), error) {
	opt := idOptions{}
	opt.logger, _ = logpkg.NewDummyLogger()
	for _, o := range opts {
		o(&opt)
	}
	serverName := NodeidServiceGRPC
	if opt.serverName != "" {
		serverName = opt.serverName
	}
	client, err := NewGRPCClient(serviceAPIManager, serverName)
	if err != nil {
		return nil, nil, err
	}
	return getIdGenerator(NewGRPCApi(client), req, &opt)
}

func getIdGenerator(nodeidAPI NodeIDAPI, req *nodeidresourcev1.GetNodeIdReq, opt *idOptions) (idpkg.Snowflake, func(), error) {
	var (
		logHelper  = log.NewHelper(log.With(opt.logger, "module", "nodeid-api/id-generator"))
		helperOpts = []Option{
			WithTries(DefaultTries),
			WithRetryDelay(DefaultRetryDelay),
			WithHeartbeatInterval(DefaultHeartbeatInterval),
			WithLogger(opt.logger),
		}
	)
	helper := NewNodeIDHelper(nodeidAPI, helperOpts...)
	mgr := NewIDManager(opt.logger, helper)

	ctx := context.Background()
	node, cleanup, err := mgr.GetSingletonSnowflakeNode(ctx, req)
	if err != nil {
		if opt.mustGetNodeIdForAPI {
			return nil, nil, err
		}
		logHelper.WithContext(ctx).Warnw("msg", "GetSingletonSnowflakeNode failed", "error", err)
		err = nil
		nodeID, _ := idpkg.GenNodeID()
		if nodeID < 1 {
			nodeID = 1
		}
		node, err = idpkg.NewBwmarrinSnowflake(int64(nodeID))
		if err != nil {
			return nil, nil, err
		}
		cleanup = func() {}
	}
	_ = SetSnowflake(node)
	return node, cleanup, nil
}
