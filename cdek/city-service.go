package cdek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const citiesUrl = "v1/location/cities/json"

func getCities(clientConfig ClientConfig, filter map[CityFilter]string) (*GetCitiesRes, error) {
	serverUrl, err := url.Parse(clientConfig.XmlApiUrl)
	if err != nil {
		return nil, err
	}

	serverUrl.Path = path.Join(serverUrl.Path, citiesUrl)

	queryString := serverUrl.Query()
	for filterKey, value := range filter {
		queryString.Set(string(filterKey), value)
	}
	serverUrl.RawQuery = queryString.Encode()

	reqUrl := serverUrl.String()

	res, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = res.Body.Close()
	}()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cities GetCitiesRes
	err = json.Unmarshal(body, &cities)
	if err != nil {
		return nil, err
	}

	return &cities, nil
}
