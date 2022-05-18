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

//GetCities This method is used to load detailed information on cities.
func (c Client) GetCities(ctx context.Context, filter map[CityFilter]string) (*GetCitiesResp, error) {
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return jsonReq[GetCitiesResp](req)
}
