package v2

import (
	"context"
	"fmt"
	"net/http"
)

type DeliveryPointsResponse []DeliveryPoint

type DeliveryPoint struct {
	Code           string `json:"code"`
	Name           string `json:"name,omitempty"`
	AddressComment string `json:"address_comment,omitempty"`
	WorkTime       string `json:"work_time,omitempty"`
	Phones         []struct {
		Number string `json:"number"`
	} `json:"phones,omitempty"`
	Email          string `json:"email,omitempty"`
	Note           string `json:"note,omitempty"`
	Type           string `json:"type"`
	OwnerCode      string `json:"owner_code"`
	TakeOnly       bool   `json:"take_only"`
	IsHandout      bool   `json:"is_handout,omitempty"`
	IsReception    bool   `json:"is_reception,omitempty"`
	IsDressingRoom bool   `json:"is_dressing_room,omitempty"`
	HaveCashless   bool   `json:"have_cashless"`
	HaveCash       bool   `json:"have_cash"`
	AllowedCod     bool   `json:"allowed_cod"`
	Site           string `json:"site,omitempty"`
	WorkTimeList   []struct {
		Day  int    `json:"day"`
		Time string `json:"time"`
	} `json:"work_time_list,omitempty"`
	WeightMin float64 `json:"weight_min,omitempty"`
	WeightMax float64 `json:"weight_max,omitempty"`
	Location  struct {
		CountryCode string  `json:"country_code"`
		RegionCode  int     `json:"region_code"`
		Region      string  `json:"region,omitempty"`
		CityCode    int     `json:"city_code"`
		City        string  `json:"city,omitempty"`
		FiasGuid    string  `json:"fias_guid,omitempty"`
		PostalCode  string  `json:"postal_code,omitempty"`
		Longitude   float64 `json:"longitude"`
		Latitude    float64 `json:"latitude"`
		Address     string  `json:"address"`
		AddressFull string  `json:"address_full,omitempty"`
	} `json:"location"`
	Fulfillment         bool   `json:"fulfillment"`
	NearestStation      string `json:"nearest_station,omitempty"`
	NearestMetroStation string `json:"nearest_metro_station,omitempty"`
	OfficeImageList     []struct {
		Url string `json:"url"`
	} `json:"office_image_list,omitempty"`
	WorkTimeExceptions []struct {
		Date      string `json:"date"`
		IsWorking bool   `json:"is_working"`
	} `json:"work_time_exceptions,omitempty"`
}

func (c *clientImpl) DeliveryPoints(ctx context.Context) (*DeliveryPointsResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.buildUri("/v2/deliverypoints"),
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

	return jsonReq[DeliveryPointsResponse](req)
}
