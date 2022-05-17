package cdek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/hashicorp/go-multierror"
)

const (
	citiesURL = "v1/location/cities/json"
)

//GetCities This method is used to load detailed information on cities.
func (c clientImpl) GetCities(filter map[CityFilter]string) (*GetCitiesResp, error) {
	serverURL, err := url.Parse(c.apiURL)
	if err != nil {
		return nil, err
	}

	serverURL.Path = path.Join(serverURL.Path, citiesURL)

	queryString := serverURL.Query()
	for filterKey, value := range filter {
		queryString.Set(string(filterKey), value)
	}
	serverURL.RawQuery = queryString.Encode()

	reqURL := serverURL.String()

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	var cities GetCitiesResp
	err = json.Unmarshal(body, &cities)
	if err != nil {
		var alertResponse AlertResponse
		err = json.Unmarshal(body, &alertResponse)
		if err != nil {
			return nil, err
		}

		multiError := &multierror.Error{}
		for _, alert := range alertResponse.Alerts {
			multiError = multierror.Append(alert)
		}

		return nil, multiError.ErrorOrNil()
	}

	return &cities, nil
}
