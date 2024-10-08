package nodeidapi

import (
	"context"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	"github.com/stretchr/testify/require"
	"testing"
)

// go test -v -count=1 ./app/nodeid-service/ -test.run=Test_idManager_SetNode
func Test_idManager_SetNode(t *testing.T) {
	type args struct {
		ctx context.Context
		req *nodeidresourcev1.GetNodeIdReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#Test_idManager_SetNode",
			args: args{
				ctx: context.Background(),
				req: &nodeidresourcev1.GetNodeIdReq{
					InstanceId:   "Test_idManager_SetNode_ID",
					InstanceName: "Test_idManager_SetNode_Name",
					Metadata:     nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := idManagerHandler.SetNode(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("SetNode() error = %v, wantErr %v", err, tt.wantErr)
			}
			err := idManagerHandler.Release(tt.args.ctx)
			require.Nil(t, err)
		})
	}
}
