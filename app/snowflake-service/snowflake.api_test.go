package snowflakeapi

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	nodeidapi "github.com/go-micro-saas/service-api/app/nodeid-service"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	"github.com/stretchr/testify/require"
	"testing"
)

func getTestingIDManager() IDManager {
	grpcClient, err := nodeidapi.NewGRPCClient(serviceAPIManger)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	logger := log.DefaultLogger
	grpcAPIHandler := nodeidapi.NewGRPCApi(grpcClient)
	nodeIDHandler, err := nodeidapi.NewNodeIDHelper(grpcAPIHandler, nodeidapi.WithLogger(logger))
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	return NewIDManager(logger, nodeIDHandler)
}

// go test -v -count=1 ./app/snowflake-service/ -test.run=TestGetSingletonSnowflakeNode
func TestGetSingletonSnowflakeNode(t *testing.T) {
	type args struct {
		idManager IDManager
		req       *nodeidresourcev1.GetNodeIdReq
	}
	tests := []struct {
		name    string
		args    args
		want    idpkg.Snowflake
		wantErr bool
	}{
		{
			name: "#TestGetSingletonSnowflakeNode",
			args: args{
				idManager: getTestingIDManager(),
				req: &nodeidresourcev1.GetNodeIdReq{
					InstanceId:   "testdata",
					InstanceName: "testdata",
					Metadata:     nil,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetSingletonSnowflakeNode(tt.args.idManager, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSingletonSnowflakeNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer got1()
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetSingletonSnowflakeNode() got = %v, want %v", got, tt.want)
			//}
			id, err := got.NextID()
			require.Nil(t, err)
			t.Logf("==> got NextID: %d\n", id)
		})
	}
}
