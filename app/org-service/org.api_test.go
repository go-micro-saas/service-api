package orgapi

import (
	"context"
	"testing"

	orgresourcev1 "github.com/go-micro-saas/service-api/api/org-service/v1/resources"
	"github.com/stretchr/testify/require"
)

// go test -v -count 1 ./app/org-service -run TestHTTPClient_Xxx
func TestHTTPClient_Xxx(t *testing.T) {
	var (
		ctx = context.Background()
		req = &orgresourcev1.PingReq{
			Message: "hello http",
		}
	)
	resp, err := orgV1HTTP.Ping(ctx, req)
	require.Nil(t, err)
	t.Logf("==> TestHTTPClient_Xxx resp: %#v\n", resp)
	t.Logf("==> TestHTTPClient_Xxx resp.Data: %#v\n", resp.GetData())
}

// go test -v -count 1 ./app/org-service -run TestGRPCClient_Xxx
func TestGRPCClient_Xxx(t *testing.T) {
	var (
		ctx = context.Background()
		req = &orgresourcev1.PingReq{
			Message: "hello grpc",
		}
	)
	resp, err := orgV1GRPC.Ping(ctx, req)
	require.Nil(t, err)
	t.Logf("==> TestGRPCClient_Xxx resp: %#v\n", resp)
	t.Logf("==> TestGRPCClient_Xxx resp.Data: %#v\n", resp.GetData())

}
