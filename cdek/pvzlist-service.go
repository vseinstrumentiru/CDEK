package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const pvzListUrl = "pvzlist/v1/xml"

func getPvzList(clientConfig ClientConfig, filter map[PvzListFilter]string) (*PvzList, error) {
	serverUrl, err := url.Parse(clientConfig.XmlApiUrl)
	if err != nil {
		return nil, err
	}

	serverUrl.Path = path.Join(serverUrl.Path, pvzListUrl)

	queryString := serverUrl.Query()
	for key, v := range filter {
		queryString.Set(string(key), v)
	}
	serverUrl.RawQuery = queryString.Encode()

	requestUrl := serverUrl.String()

	resp, err := http.Get(requestUrl)
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