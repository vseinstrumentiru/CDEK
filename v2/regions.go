package v2

import (
	"context"
	"fmt"
	"net/http"
)

type RegionsRequest struct {
	// CountryCodes Массив кодов стран в формате  ISO_3166-1_alpha-2
	CountryCodes []string `url:"country_codes,omitempty"`
	// Size Ограничение выборки результата. По умолчанию 1000
	Size int `url:"size,omitempty"`
	// Page Номер страницы выборки результата. По умолчанию 0
	Page int `url:"page,omitempty"`
	// Lang Локализация офиса. По умолчанию "rus"
	Lang string `url:"lang,omitempty"`
}

type RegionsResponse []Region

type Region struct {
	CountryCode string `json:"country_code"`
	Region      string `json:"region"`
	Country     string `json:"country"`
	RegionCode  int    `json:"region_code,omitempty"`
}

func (c *clientImpl) Regions(ctx context.Context, input *RegionsRequest) (*RegionsResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.buildUri("/v2/location/regions", input),
		nil,
	)
	if err != nil {
		return nil, err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	return jsonReq[RegionsResponse](req)
}
