package pingapi

import (
	"context"
	pingroucesv1 "github.com/go-micro-saas/service-api/api/ping-service/v1/resources"
	"testing"
)

// go test -v -count 1 ./app/ping-service -run Test_Ping_Xxx
func Test_Ping_Xxx(t *testing.T) {
	type args struct {
		handler PingAPI
		ctx     context.Context
		req     *pingroucesv1.PingReq
	}
	tests := []struct {
		name    string
		args    args
		want    *pingroucesv1.PingRespData
		wantErr bool
	}{
		{
			name: "#TestHTTPPing",
			args: args{
				handler: httpAPIHandler,
				ctx:     context.Background(),
				req:     &pingroucesv1.PingReq{Message: "TestHTTPPing"},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "#TestGRPCPing",
			args: args{
				handler: grpcAPIHandler,
				ctx:     context.Background(),
				req:     &pingroucesv1.PingReq{Message: "TestGRPCPing"},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.handler.Ping(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Ping() got = %v, want %v", got, tt.want)
			//}
			t.Log("==> got: ", got.String())
		})
	}
}
