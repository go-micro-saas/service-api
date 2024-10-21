package pingapi

import (
	"fmt"
	launcher "github.com/go-micro-saas/service-api/testdata/launcher"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"os"
	"testing"
)

var (
	launcherManager  setuputil.LauncherManager
	serviceAPIManger clientutil.ServiceAPIManager
	httpAPIHandler   *httpAPI
	grpcAPIHandler   *grpcAPI
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

	os.Exit(m.Run())
}
