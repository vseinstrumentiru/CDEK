package cdek

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

const (
	regionsURL = "/v1/location/regions/json"
)

//GetRegions This method is used to load detailed information on regions.
func (c Client) GetRegions(ctx context.Context, filter map[RegionFilter]string) (*GetRegionsResp, error) {
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return jsonReq[GetRegionsResp](req)
}
