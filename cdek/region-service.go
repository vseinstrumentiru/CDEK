package cdek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const regionsUrl = "/v1/location/regions/json"

func getRegions(clientConfig ClientConfig, filter map[RegionFilter]string) (*GetRegionsResponse, error) {
	serverUrl, err := url.Parse(clientConfig.XmlApiUrl)
	if err != nil {
		return nil, err
	}

	serverUrl.Path = path.Join(serverUrl.Path, regionsUrl)

	queryString := serverUrl.Query()
	for filterKey, value := range filter {
		queryString.Set(string(filterKey), value)
	}
	serverUrl.RawQuery = queryString.Encode()

	requestUrl := serverUrl.String()

	resp, err := http.Get(requestUrl)
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

	var regions GetRegionsResponse
	err = json.Unmarshal(body, &regions)
	if err != nil {
		return nil, err
	}

	return &regions, nil
}
