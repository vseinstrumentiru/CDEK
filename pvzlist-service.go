package cdek

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

const (
	pvzListURL = "pvzlist/v1/xml"
)

//GetPvzList The method is used to load the list of active pickup points, from which the client can pick up its order.
func (c Client) GetPvzList(ctx context.Context, filter map[PvzListFilter]string) ([]*Pvz, error) {
	serverURL, err := url.Parse(c.apiURL)
	if err != nil {
		return nil, err
	}
	serverURL.Path = path.Join(serverURL.Path, pvzListURL)

	queryString := serverURL.Query()
	for filterKey, value := range filter {
		queryString.Set(string(filterKey), value)
	}
	serverURL.RawQuery = queryString.Encode()

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL.String(), nil)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", urlFormEncoded)

	resp, err := xmlReq[pvzList](r)
	if err != nil {
		return nil, err
	}

	return resp.Pvz, nil
}
