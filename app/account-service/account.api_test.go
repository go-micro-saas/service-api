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
		req = &accountresourcev1.PingReq{}
	)
	resp, err := accountV1HTTP.Ping(ctx, req)
	require.Nil(t, err)
	t.Logf("==> resp: %#v\n", resp)
}

func TestGRPCClient_Xxx(t *testing.T) {

}
