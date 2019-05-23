package cdek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const regionsUrl = "/v1/location/regions/json"

func getRegions(clientConfig ClientConf, filter map[RegionFilter]string) (*GetRegionsRes, error) {
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

	var regions GetRegionsRes
	err = json.Unmarshal(body, &regions)
	if err != nil {
		return nil, err
	}

	return &regions, nil
}
