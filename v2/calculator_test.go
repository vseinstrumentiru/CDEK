package v2

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientImpl_CalculatorTrafiffList(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c := createTestClient()

	resp, err := c.CalculatorTrafiffList(timedCtx, &CalculatorTrafiffListRequest{
		Lang:         "rus",
		Currency:     1,
		FromLocation: Location{Code: 44},
		ToLocation:   Location{Code: 287},
		Packages: []Package{
			{Weight: 1},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Greater(t, len(resp.TariffCodes), 0)
}
