package cdek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const regionsURL = "/v1/location/regions/json"

func getRegions(clientConfig ClientConf, filter map[RegionFilter]string) (*GetRegionsResp, error) {
	serverURL, err := url.Parse(clientConfig.CdekAPIURL)
	if err != nil {
		return nil, err
	}

	serverURL.Path = path.Join(serverURL.Path, regionsURL)

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

	var regions GetRegionsResp
	err = json.Unmarshal(body, &regions)
	if err != nil {
		return nil, err
	}

	return &regions, nil
}
