package accountapi

import (
	"fmt"
	accountservicev1 "github.com/go-micro-saas/service-api/api/account-service/v1/services"
	launcher "github.com/go-micro-saas/service-api/testdata/launcher"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"os"
	"testing"
)

var (
	launcherManager  setuputil.LauncherManager
	serviceAPIManger clientutil.ServiceAPIManager
	accountV1HTTP    accountservicev1.SrvAccountV1HTTPClient
	accountV1GRPC    accountservicev1.SrvAccountV1Client
	userAuthV1HTTP   accountservicev1.SrvUserAuthV1HTTPClient
	userAuthV1GRPC   accountservicev1.SrvUserAuthV1Client
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
	accountV1HTTP, err = NewAccountV1HTTPClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	accountV1GRPC, err = NewAccountV1GRPCClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	// client
	userAuthV1HTTP, err = NewUserAuthV1HTTPClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	userAuthV1GRPC, err = NewUserAuthV1GRPCClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}

	os.Exit(m.Run())
}
