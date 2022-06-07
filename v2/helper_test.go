package v2

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClientImpl_CitiesAll(t *testing.T) {
	ctx := context.Background()

	c := createTestClient()

	resp, err := HelperCitiesAll(ctx, c, nil, 100)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
