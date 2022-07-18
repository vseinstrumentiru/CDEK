package v2

import (
	"context"
	"fmt"
	"net/http"
)

type DeliveryPointsResponse []DeliveryPoint

type DeliveryPointWorkTime struct {
	Day  int    `json:"day"`
	Time string `json:"time"`
}

type DeliveryPointWorkTimeExceptions struct {
	Date      string `json:"date"`
	IsWorking bool   `json:"is_working"`
}

type DeliveryPointOfficeImage struct {
	Url string `json:"url"`
}

type DeliveryPointLocation struct {
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
}

type DeliveryPoint struct {
	Code                string                            `json:"code"`
	Name                string                            `json:"name,omitempty"`
	AddressComment      string                            `json:"address_comment,omitempty"`
	WorkTime            string                            `json:"work_time,omitempty"`
	Phones              []Phone                           `json:"phones,omitempty"`
	Email               string                            `json:"email,omitempty"`
	Note                string                            `json:"note,omitempty"`
	Type                string                            `json:"type"`
	OwnerCode           string                            `json:"owner_code"`
	TakeOnly            bool                              `json:"take_only"`
	IsHandout           bool                              `json:"is_handout,omitempty"`
	IsReception         bool                              `json:"is_reception,omitempty"`
	IsDressingRoom      bool                              `json:"is_dressing_room,omitempty"`
	HaveCashless        bool                              `json:"have_cashless"`
	HaveCash            bool                              `json:"have_cash"`
	AllowedCod          bool                              `json:"allowed_cod"`
	Site                string                            `json:"site,omitempty"`
	WorkTimeList        []DeliveryPointWorkTime           `json:"work_time_list,omitempty"`
	WeightMin           float64                           `json:"weight_min,omitempty"`
	WeightMax           float64                           `json:"weight_max,omitempty"`
	Location            DeliveryPointLocation             `json:"location"`
	Fulfillment         bool                              `json:"fulfillment"`
	NearestStation      string                            `json:"nearest_station,omitempty"`
	NearestMetroStation string                            `json:"nearest_metro_station,omitempty"`
	OfficeImageList     []DeliveryPointOfficeImage        `json:"office_image_list,omitempty"`
	WorkTimeExceptions  []DeliveryPointWorkTimeExceptions `json:"work_time_exceptions,omitempty"`
}

type DeliveryPointsRequest struct {
	// PostalCode Почтовый индекс города, для которого необходим список офисов
	PostalCode int `url:"postal_code,omitempty"`
	// CityCode Код населенного пункта СДЭК (метод "Список населенных пунктов")
	CityCode int `url:"city_code,omitempty"`
	// Type Тип офиса, может принимать значения: «PVZ» - склады, «POSTAMAT» - постаматы, «ALL» - все.
	Type string `url:"type,omitempty"`
	// CountryCode Код страны в формате ISO_3166-1_alpha-2 (см. “Общероссийский классификатор стран мира”)
	CountryCode string `url:"country_code,omitempty"`
	// RegionCode Код региона по базе СДЭК
	RegionCode int `url:"region_code,omitempty"`
	// HaveCashless Наличие терминала оплаты
	HaveCashless bool `url:"have_cashless,omitempty"`
	// HaveCash Есть прием наличных
	HaveCash bool `url:"have_cash,omitempty"`
	// AllowedCod Разрешен наложенный платеж
	AllowedCod bool `url:"allowed_cod,omitempty"`
	// IsDressingRoom Наличие примерочной
	IsDressingRoom bool `url:"is_dressing_room,omitempty"`
	// WeightMax Максимальный вес в кг, который может принять офис (значения больше 0 - передаются офисы, которые принимают этот вес; 0 - офисы с нулевым весом не передаются; значение не указано - все офисы)
	WeightMax bool `url:"weight_max,omitempty"`
	// WeightMin Минимальный вес в кг, который принимает офис (при переданном значении будут выводиться офисы с минимальным весом до указанного значения)
	WeightMin bool `url:"weight_min,omitempty"`
	// Lang Локализация офиса. По умолчанию "rus"
	Lang string `url:"lang,omitempty"`
	// TakeOnly Является ли офис только пунктом выдачи
	TakeOnly bool `url:"take_only,omitempty"`
	// IsHandout Является пунктом выдачи, может принимать значения
	IsHandout bool `url:"is_handout,omitempty"`
	// IsReception Есть ли в офисе приём заказов
	IsReception bool `url:"is_reception,omitempty"`
	// FiasGuid Код города ФИАС	UUID
	FiasGuid string `url:"fias_guid,omitempty"`
}

func (c *clientImpl) DeliveryPoints(ctx context.Context, input *DeliveryPointsRequest) (*DeliveryPointsResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.buildUri("/v2/deliverypoints", input),
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
