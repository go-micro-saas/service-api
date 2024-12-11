package snowflakeapi

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	nodeidapi "github.com/go-micro-saas/service-api/app/nodeid-service"
	launcher "github.com/go-micro-saas/service-api/testdata/launcher"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"os"
	"testing"
)

var (
	launcherManager  setuputil.LauncherManager
	serviceAPIManger clientutil.ServiceAPIManager
	idManagerHandler *idManager
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
	//httpClient, err := nodeidapi.NewHTTPClient(serviceAPIManger)
	//if err != nil {
	//	fmt.Printf("%+v\n", err)
	//	panic(err)
	//}
	grpcClient, err := nodeidapi.NewGRPCClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	nodeIDHandler := nodeidapi.NewNodeIDHelper(nodeidapi.NewGRPCApi(grpcClient), nodeidapi.WithLogger(log.DefaultLogger))
	idManagerHandler = NewIDManager(log.DefaultLogger, nodeIDHandler).(*idManager)

	os.Exit(m.Run())
}
