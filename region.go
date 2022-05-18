package cdek

import (
	"context"
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

//RegionFilter filter key for "List of Regions" request
type RegionFilter string

const (
	//RegionFilterRegionCodeExt  Region code
	RegionFilterRegionCodeExt RegionFilter = "regionCodeExt"

	//RegionFilterRegionCode Region code in the CDEK IS
	RegionFilterRegionCode RegionFilter = "regionCode"

	//RegionFilterRegionFiasGUID Region code according to the Federal Information Address System
	RegionFilterRegionFiasGUID RegionFilter = "regionFiasGuid"

	//RegionFilterCountryCode Country code in the CDEK IS
	RegionFilterCountryCode RegionFilter = "countryCode"

	//RegionFilterCountryCodeExt Code according to the Russian Classifier of Countries of the World
	RegionFilterCountryCodeExt RegionFilter = "countryCodeExt"

	//RegionFilterPage Number of the results page. Default value: 0
	RegionFilterPage RegionFilter = "page"

	//RegionFilterSize Limitation on the number of results displayed. Default value: 1,000
	RegionFilterSize RegionFilter = "size"
)

//GetRegions This method is used to load detailed information on regions.
func (c clientImpl) GetRegions(ctx context.Context, filter map[RegionFilter]string) (*GetRegionsResp, error) {
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

//GetRegionsResp List of Regions response
type GetRegionsResp []Region

//Region Region response
type Region struct {
	RegionUUID     *string `json:"regionUuid"`
	RegionName     *string `json:"regionName"`
	Prefix         *string `json:"prefix,omitempty"`
	RegionCodeExt  *int    `json:"regionCodeExt,string,omitempty"`
	RegionCode     *int    `json:"regionCode,string,omitempty"`
	RegionFiasGUID *string `json:"regionFiasGuid,omitempty"`
	CountryName    *string `json:"countryName"`
	CountryCode    *string `json:"countryCode,omitempty"`
	CountryCodeExt *int    `json:"countryCodeExt,omitempty,string"`
}
