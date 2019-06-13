package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const pvzListURL = "pvzlist/v1/xml"

func getPvzList(clientConfig ClientConf, filter map[PvzListFilter]string) (*PvzList, error) {
	serverURL, err := url.Parse(clientConfig.XmlApiUrl)
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
		err = resp.Body.Close()
	}()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pvzList PvzList
	err = xml.Unmarshal(body, &pvzList)
	if err != nil {
		return nil, err
	}

	return &pvzList, nil
}
