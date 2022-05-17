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
	regionsURL = "/v1/location/regions/json"
)

//GetRegions This method is used to load detailed information on regions.
func (c clientImpl) GetRegions(filter map[RegionFilter]string) (*GetRegionsResp, error) {
	serverURL, err := url.Parse(c.apiURL)
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
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	var regions GetRegionsResp
	err = json.Unmarshal(body, &regions)
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

	return &regions, nil
}
