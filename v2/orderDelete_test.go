package v2

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientImpl_OrderDelete(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c := createTestClient()

	resp, err := c.OrderRegister(timedCtx, nil)
	require.Error(t, err)
	require.Nil(t, resp)

	uid := uuid.NewString()
	resp, err = c.OrderRegister(timedCtx, &OrderRegisterRequest{
		Type:         0,
		Number:       uid,
		Comment:      "test",
		TariffCode:   62,
		FromLocation: Location{Code: 44, Address: "qwe"},
		ToLocation:   Location{Code: 287, Address: "qwe"},
		Sender: RecipientSender{
			Name:    "test",
			Company: "test",
			Email:   "test@test.com",
		},
		Recipient: RecipientSender{
			Name: "test",
			Phones: []Phone{
				{Number: "123"},
			},
		},
		Packages: []Package{
			{
				Number:  "test",
				Weight:  1,
				Comment: "test",
				Items: []PackageItem{
					{
						Name:    "test",
						WareKey: "test",
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Greater(t, len(resp.Requests), 0)

	statusResp, err := c.OrderDelete(ctx, resp.Entity.Uuid)
	require.NoError(t, err)
	require.Equal(t, statusResp.Entity.Uuid, resp.Entity.Uuid)
}
