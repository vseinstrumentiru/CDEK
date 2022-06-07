package v2

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientImpl_Regions(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c := createTestClient()

	resp, err := c.Regions(timedCtx, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = c.Regions(timedCtx, &RegionsRequest{Page: 1000, Size: 1000})
	require.NoError(t, err)
	require.NotNil(t, resp)
}
