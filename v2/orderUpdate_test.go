package v2

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientImpl_OrderUpdate(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c := createTestClient()

	resp, err := c.OrderRegister(timedCtx, nil)
	require.Error(t, err)
	require.Nil(t, resp)

	registerReq := &OrderRegisterRequest{
		Type:         0,
		Number:       uuid.NewString(),
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
	}
	resp, err = c.OrderRegister(timedCtx, registerReq)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Greater(t, len(resp.Requests), 0)

	updateResp, err := c.OrderUpdate(ctx, &OrderUpdateRequest{
		UUID:       resp.Entity.Uuid,
		Comment:    "updated",
		ToLocation: registerReq.ToLocation,
		Recipient:  registerReq.Recipient,
		TariffCode: registerReq.TariffCode,
		Packages:   registerReq.Packages,
	})
	require.NoError(t, err)

	statusResp, err := c.OrderStatus(ctx, updateResp.Entity.Uuid)
	require.NoError(t, err)
	require.Equal(t, statusResp.Entity.Comment, "updated")
}
