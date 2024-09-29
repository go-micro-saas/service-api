package nodeidapi

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	launcher "github.com/go-micro-saas/service-api/testdata/launcher"
	clientutil "github.com/go-micro-saas/service-kit/cluster_service_api"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	"os"
	"testing"
)

var (
	launcherManager  setuputil.LauncherManager
	serviceAPIManger clientutil.ServiceAPIManager
	httpAPIHandler   *httpAPI
	grpcAPIHandler   *grpcAPI
	nodeIDHandler    *nodeIDHelper
)

func TestMain(m *testing.M) {
	configPath := "./../../testdata/configs"

	// testdata instance
	testInstance, err := launcher.InitTestdataInstance(configPath)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	launcherManager = testInstance.LauncherManager
	serviceAPIManger = testInstance.ServiceAPIManager
	defer func() { _ = launcherManager.Close() }()

	// client
	httpClient, err := NewHTTPClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	grpcClient, err := NewGRPCClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	httpAPIHandler = NewHTTPApi(httpClient).(*httpAPI)
	grpcAPIHandler = NewGRPCApi(grpcClient).(*grpcAPI)
	nodeIDHandler = NewNodeIDHelper(grpcAPIHandler, WithLogger(log.DefaultLogger)).(*nodeIDHelper)

	os.Exit(m.Run())
}
