package v2

import (
	"context"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClientImpl_Auth(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c := NewClient(&Options{
		Endpoint: EndpointTest,
		Credentials: &Credentials{
			ClientID:     os.Getenv("CDEK_CLIENT_ID"),
			ClientSecret: os.Getenv("CDEK_SECRET_ID"),
		},
	})

	resp, err := c.Auth(timedCtx)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestClientImpl_DeliveryPoints(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c := NewClient(&Options{
		Endpoint: EndpointTest,
		Credentials: &Credentials{
			ClientID:     os.Getenv("CDEK_CLIENT_ID"),
			ClientSecret: os.Getenv("CDEK_SECRET_ID"),
		},
	})

	resp, err := c.DeliveryPoints(timedCtx)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
