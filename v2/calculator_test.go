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
		FromLocation: CalculatorLocation{Code: "53562"},
		ToLocation:   CalculatorLocation{Code: "53562"},
		Packages: []Package{
			{Weight: 1},
		},
	})
	// @todo unsupported media type
	require.NoError(t, err)
	require.NotNil(t, resp)
}
