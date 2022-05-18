package cdek

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

//CalculateDelivery Cost calculation on tariffs with priority.
func (c Client) CalculateDelivery(ctx context.Context, req GetCostReq) (*GetCostRespResult, error) {
	req.setAuth(c.auth)

	payload, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.calculatorURL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", jsonContentType)

	resp, err := jsonReq[getCostResp](r)
	if err != nil {
		return nil, err
	}

	if resp.ErrorResp != nil {
		var errs error
		for _, err := range resp.ErrorResp {
			errs = multierror.Append(
				errs,
				fmt.Errorf("error code: %s, error text: %s", *err.ErrorCode, *err.Msg),
			)
		}

		return nil, errs
	}

	return &resp.Result, nil
}
