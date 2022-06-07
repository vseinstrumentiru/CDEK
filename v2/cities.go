package v2

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type CitiesRequest struct {
	// CountryCodes Массив кодов стран в формате  ISO_3166-1_alpha-2
	CountryCodes []string `url:"country_codes,omitempty"`
	// RegionCode Код региона СДЭК
	RegionCode int `url:"region_code,omitempty"`
	// FiasGuid Уникальный идентификатор ФИАС населенного пункта UUID
	FiasGuid string `url:"fias_guid,omitempty"`
	// PostalCode Почтовый индекс
	PostalCode string `url:"postal_code,omitempty"`
	// Code Код населенного пункта СДЭК
	Code string `url:"code,omitempty"`
	// City Название населенного пункта. Должно соответствовать полностью
	City string `url:"city,omitempty"`
	// Size Ограничение выборки результата. По умолчанию 1000
	Size int `url:"size,omitempty"`
	// Page Номер страницы выборки результата. По умолчанию 0
	Page int `url:"page,omitempty"`
	// Lang Локализация офиса. По умолчанию "rus"
	Lang string `url:"lang,omitempty"`
}

type CitiesResponse []*City

type City struct {
	Code         int      `json:"code"`
	City         string   `json:"city"`
	CountryCode  string   `json:"country_code"`
	Country      string   `json:"country"`
	Region       string   `json:"region,omitempty"`
	RegionCode   int      `json:"region_code"`
	SubRegion    string   `json:"sub_region,omitempty"`
	PostalCodes  []string `json:"postal_codes,omitempty"`
	Longitude    float64  `json:"longitude"`
	Latitude     float64  `json:"latitude"`
	TimeZone     string   `json:"time_zone"`
	KladrCode    string   `json:"kladr_code,omitempty"`
	PaymentLimit float64  `json:"payment_limit,omitempty"`
	FiasGuid     string   `json:"fias_guid,omitempty"`
}

func (c *clientImpl) Cities(ctx context.Context, input *CitiesRequest) (*CitiesResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.buildUri("/v2/location/cities", input),
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequestWithContext")
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getAccessToken")
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	return jsonReq[CitiesResponse](req)
}
