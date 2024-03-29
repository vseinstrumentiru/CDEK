package v2

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientImpl_DeliveryPoints(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	c := createTestClient()

	resp, err := c.DeliveryPoints(timedCtx, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = c.DeliveryPoints(timedCtx, &DeliveryPointsRequest{
		CountryCode: "ru",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}
