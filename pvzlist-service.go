package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const pvzListURL = "pvzlist/v1/xml"

func (cl client) GetPvzList(filter map[PvzListFilter]string) ([]*Pvz, error) {
	serverURL, err := url.Parse(cl.clientConf.CdekAPIURL)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pvzList pvzList
	err = xml.Unmarshal(body, &pvzList)
	if err != nil {
		return nil, err
	}

	return pvzList.Pvz, nil
}
