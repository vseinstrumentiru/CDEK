package cdek

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

// CalculateDelivery Cost calculation on tariffs with priority.
func (c clientImpl) CalculateDelivery(ctx context.Context, req *CalculateDeliveryRequest) (*CalculateDeliveryResult, error) {
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

	resp, err := jsonReq[CalculateDeliveryResponse](r)
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

//CalculateDeliveryGood Location's dimension
type CalculateDeliveryGood struct {
	Weight float64 `json:"weight"`
	Length int     `json:"length"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Volume float64 `json:"volume"`
}

//CalculateDeliveryService List of additional service
type CalculateDeliveryService struct {
	ID    int `json:"id"`
	Param int `json:"param,omitempty"`
}

//ServiceResp List of transmitted additional services
type ServiceResp struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Rate  float64 `json:"rate,omitempty"`
}

//CalculateDeliveryRequest Cost calculation on tariffs with priority request
type CalculateDeliveryRequest struct {
	credentialsJSON
	Version        string                      `json:"version"`
	SenderCityID   int                         `json:"senderCityId"`
	ReceiverCityID int                         `json:"receiverCityId"`
	TariffID       int                         `json:"tariffId"`
	Goods          []*CalculateDeliveryGood    `json:"goods"`
	Services       []*CalculateDeliveryService `json:"services,omitempty"`
}

type CalculateDeliveryResponse struct {
	ErrorResp []Error                 `json:"error,omitempty"`
	Result    CalculateDeliveryResult `json:"result"`
}

//CalculateDeliveryResult Cost calculation on tariffs with priority result response
type CalculateDeliveryResult struct {
	Price             float64       `json:"price,string"`
	DeliveryPeriodMin int           `json:"deliveryPeriodMin"`
	DeliveryPeriodMax int           `json:"deliveryPeriodMax"`
	DeliveryDateMin   string        `json:"deliveryDateMin"`
	DeliveryDateMax   string        `json:"deliveryDateMax"`
	TariffID          int           `json:"tariffId"`
	CashOnDelivery    float64       `json:"cashOnDelivery"`
	PriceByCurrency   float64       `json:"priceByCurrency"`
	Currency          string        `json:"currency"`
	PercentVAT        int           `json:"percentVAT"`
	Services          []ServiceResp `json:"services"`
}
