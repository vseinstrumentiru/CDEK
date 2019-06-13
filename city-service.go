package cdek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const citiesURL = "v1/location/cities/json"

func getCities(clientConfig ClientConf, filter map[CityFilter]string) (*GetCitiesResp, error) {
	serverURL, err := url.Parse(clientConfig.XmlApiUrl)
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
		err = resp.Body.Close()
	}()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cities GetCitiesResp
	err = json.Unmarshal(body, &cities)
	if err != nil {
		return nil, err
	}

	return &cities, nil
}
