package launcher

import (
	clientutil "github.com/go-micro-saas/service-kit/cluster_service_api"
	setuputil "github.com/go-micro-saas/service-kit/setup"
)

type TestdataInstance struct {
	LauncherManager   setuputil.LauncherManager
	ServiceAPIManager clientutil.ServiceAPIManager
}

func InitTestdataInstance(configPath string) (*TestdataInstance, error) {
	launcher, err := setuputil.NewLauncherManager(configPath)
	if err != nil {
		return nil, err
	}

	logger, err := launcher.GetLogger()
	if err != nil {
		return nil, err
	}

	apiConfigs, _, err := clientutil.ToConfig(launcher.GetConfig().GetClusterServiceApi())
	if err != nil {
		return nil, err
	}
	var opts = []clientutil.Option{
		clientutil.WithLogger(logger),
	}
	apiManager, err := clientutil.NewServiceAPIManager(apiConfigs, opts...)
	if err != nil {
		return nil, err
	}
	return &TestdataInstance{
		LauncherManager:   launcher,
		ServiceAPIManager: apiManager,
	}, nil
}
