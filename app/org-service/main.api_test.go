package orgapi

import (
	"fmt"
	"os"
	"testing"

	orgservicev1 "github.com/go-micro-saas/service-api/api/org-service/v1/services"
	launcher "github.com/go-micro-saas/service-api/testdata/launcher"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

var (
	launcherManager  setuputil.LauncherManager
	serviceAPIManger clientutil.ServiceAPIManager
	orgV1HTTP        orgservicev1.SrvOrgV1HTTPClient
	orgV1GRPC        orgservicev1.SrvOrgV1Client
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
	orgV1HTTP, err = NewOrgV1HTTPClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	orgV1GRPC, err = NewOrgV1GRPCClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}

	os.Exit(m.Run())
}
