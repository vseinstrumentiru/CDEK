package v2

import (
	"context"
	"fmt"
	"net/http"
)

func (c *clientImpl) OrderDelete(ctx context.Context, uuid string) (*Response, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		c.buildUri(fmt.Sprintf("/v2/orders/%s", uuid), nil),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := jsonReq[Response](req)
	if err != nil {
		return nil, err
	}

	if err := validateResponse(resp.Requests); err != nil {
		return nil, err
	}

	return resp, nil
}
