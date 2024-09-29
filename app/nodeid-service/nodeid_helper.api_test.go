package nodeidapi

import (
	"context"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	apiutil "github.com/go-micro-saas/service-api/util"
	"testing"
	"time"
)

// go test -v -count=1 ./app/nodeid-service/ -test.run=Test_nodeIDHelper_GetAndAutoRenewalNodeID
func Test_nodeIDHelper_GetAndAutoRenewalNodeID(t *testing.T) {
	type args struct {
		ctx context.Context
		req *nodeidresourcev1.GetNodeIdReq
	}
	tests := []struct {
		name    string
		args    args
		want    *nodeidresourcev1.GetNodeIdRespData
		want1   RenewalManager
		wantErr bool
	}{
		{
			name: "#GetAndAutoRenewalNodeID",
			args: args{
				ctx: context.Background(),
				req: &nodeidresourcev1.GetNodeIdReq{
					InstanceId:   "my-testdata-id",
					InstanceName: "my-testdata-name",
					Metadata:     map[string]string{"k": "v"},
				},
			},
			want:    nil,
			want1:   nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := nodeIDHandler.GetAndAutoRenewalNodeID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAndAutoRenewalNodeID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetAndAutoRenewalNodeID() got = %v, want %v", got, tt.want)
			//}
			//if !reflect.DeepEqual(got1, tt.want1) {
			//	t.Errorf("GetAndAutoRenewalNodeID() got1 = %v, want %v", got1, tt.want1)
			//}
			t.Log("==> got: ", got)
			t.Log("==> waiting renewal: ", got.HeartbeatInterval.AsDuration())
			apiutil.Sleep(got.HeartbeatInterval.AsDuration() + time.Second)
			t.Log("==> wait: renewal result")
			result := got1.RenewalResult(tt.args.ctx)
			if result.Err != nil {
				t.Error(result.Err)
				t.FailNow()
			}
			if result.Data == nil {
				t.Error("result data is nil")
				t.FailNow()
			}
			t.Log("==> renewal result: ", result.Data)
			t.Log("==> renewal LastTime: ", result.LastTime)
			err = got1.Stop(tt.args.ctx)
			if err != nil {
				t.Log("==> Stop failed: ", err)
				t.FailNow()
			}
			t.Log("==> waiting stop renewal: ", got.HeartbeatInterval.AsDuration())
			apiutil.Sleep(time.Second)
			_, err = nodeIDHandler.ReleaseNodeId(tt.args.ctx, got)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
		})
	}
}
