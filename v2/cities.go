package v2

import (
	"context"
	"fmt"
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

type CitiesResponse []Region

type City struct {
	Code        string  `json:"code"`
	CountryCode string  `json:"country_code"`
	FiasGuid    string  `json:"fias_guid"`
	Country     string  `json:"country"`
	Region      string  `json:"region"`
	RegionCode  string  `json:"region_code"`
	SubRegion   string  `json:"sub_region"`
	City        string  `json:"city"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	TimeZone    string  `json:"time_zone"`
}

func (c *clientImpl) Cities(ctx context.Context, input *CitiesRequest) (*CitiesResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.buildUri("/v2/location/cities", input),
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

	return jsonReq[CitiesResponse](req)
}
