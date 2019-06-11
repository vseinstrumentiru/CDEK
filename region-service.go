package cdek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const regionsUrl = "/v1/location/regions/json"

func getRegions(clientConfig ClientConf, filter map[RegionFilter]string) (*GetRegionsResp, error) {
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

	reqUrl := serverUrl.String()

	resp, err := http.Get(reqUrl)
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
