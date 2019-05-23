package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const pvzListUrl = "pvzlist/v1/xml"

func getPvzList(clientConfig ClientConf, filter map[PvzListFilter]string) (*PvzList, error) {
	serverUrl, err := url.Parse(clientConfig.XmlApiUrl)
	if err != nil {
		return nil, err
	}

	serverUrl.Path = path.Join(serverUrl.Path, pvzListUrl)

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

	var pvzList PvzList
	err = xml.Unmarshal(body, &pvzList)
	if err != nil {
		return nil, err
	}

	return &pvzList, nil
}
