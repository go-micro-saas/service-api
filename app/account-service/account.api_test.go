package accountapi

import (
	"context"
	accountresourcev1 "github.com/go-micro-saas/service-api/api/account-service/v1/resources"
	"github.com/stretchr/testify/require"
	"testing"
)

// go test -v -count 1 ./app/account-service -run TestHTTPClient_Xxx
func TestHTTPClient_Xxx(t *testing.T) {
	var (
		ctx = context.Background()
		req = &accountresourcev1.PingReq{
			Message: "hello world",
		}
	)
	resp, err := accountV1HTTP.Ping(ctx, req)
	require.Nil(t, err)
	t.Logf("==> TestHTTPClient_Xxx resp: %#v\n", resp)
	t.Logf("==> TestHTTPClient_Xxx resp.Data: %#v\n", resp.GetData())
}

// go test -v -count 1 ./app/account-service -run TestGRPCClient_Xxx
func TestGRPCClient_Xxx(t *testing.T) {
	var (
		ctx = context.Background()
		req = &accountresourcev1.PingReq{
			Message: "hello world",
		}
	)
	resp, err := accountV1GRPC.Ping(ctx, req)
	require.Nil(t, err)
	t.Logf("==> TestGRPCClient_Xxx resp: %#v\n", resp)
	t.Logf("==> TestGRPCClient_Xxx resp.Data: %#v\n", resp.GetData())

}
