package v2

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientImpl_Cities(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c := createTestClient()

	resp, err := c.Cities(timedCtx, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = c.Cities(timedCtx, &CitiesRequest{Page: 1000, Size: 1000})
	require.NoError(t, err)
	require.NotNil(t, resp)
}
