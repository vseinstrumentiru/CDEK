package cdek

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

const (
	citiesURL = "v1/location/cities/json"
)

//CityFilter filter key for "List of Cities" request
type CityFilter string

const (
	//CityFilterRegionCodeExt Код региона
	CityFilterRegionCodeExt CityFilter = "regionCodeExt"

	//CityFilterRegionCode Код региона в ИС СДЭК
	CityFilterRegionCode CityFilter = "regionCode"

	//CityFilterRegionFiasGUID Код региона из ФИАС
	CityFilterRegionFiasGUID CityFilter = "regionFiasGuid"

	//CityFilterPage Номер страницы выборки результата.По умолчанию 0
	CityFilterPage CityFilter = "page"

	//CityFilterSize Ограничение выборки результата.По умолчанию 1000
	CityFilterSize CityFilter = "size"

	//CityFilterCountryCode Код страны в формате ISO 3166-1 alpha-2
	CityFilterCountryCode CityFilter = "countryCode"

	//CityFilterCityName Название города
	CityFilterCityName CityFilter = "cityName"

	//CityFilterPostcode Почтовый индекс
	CityFilterPostcode CityFilter = "postcode"
)

//GetCities This method is used to load detailed information on cities.
func (c *clientImpl) GetCities(ctx context.Context, filter map[CityFilter]string) (*GetCitiesResp, error) {
	serverURL, err := url.Parse(c.apiURL)
	if err != nil {
		return nil, err
	}

	qs := serverURL.Query()
	for k, v := range filter {
		qs.Set(string(k), v)
	}
	serverURL.Path = path.Join(serverURL.Path, citiesURL)
	serverURL.RawQuery = qs.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", serverURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return jsonReq[GetCitiesResp](req)
}
