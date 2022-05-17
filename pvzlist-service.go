package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	pvzListURL = "pvzlist/v1/xml"
)

//GetPvzList The method is used to load the list of active pickup points, from which the client can pick up its order.
func (c clientImpl) GetPvzList(filter map[PvzListFilter]string) ([]*Pvz, error) {
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

	reqURL := serverURL.String()

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	var pvzList pvzList
	err = xml.Unmarshal(body, &pvzList)
	if err != nil {
		return nil, err
	}

	return pvzList.Pvz, nil
}
