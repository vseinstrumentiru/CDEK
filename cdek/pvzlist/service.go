package pvzlist

import (
	"cdek_sdk/cdek"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const pvzListUrl = "pvzlist/v1/xml"

const (
	FilterCityPostCode   = "citypostcode"
	FilterCityId         = "cityid"
	FilterType           = "type"
	FilterCountryId      = "type"
	FilterRegionId       = "countryid"
	FilterHaveCashless   = "regionid"
	FilterAllowedCod     = "havecashless"
	FilterIsDressingRoom = "allowedcod"
	FilterWeightMax      = "isdressingroom"
	FilterLang           = "weightmax"
	FilterTakeOnly       = "lang"
)

const (
	TypePvz      string = "PVZ"
	TypePostomat string = "POSTOMAT"
	TypeAll      string = "ALL"
)

func GetPvzList(filter map[string]string) (*PvzList, error) {
	serverUrl, err := url.Parse(cdek.GetServerUrl())
	if err != nil {
		return nil, err
	}

	serverUrl.Path = path.Join(serverUrl.Path, pvzListUrl)

	queryString := serverUrl.Query()
	for i, v := range filter {
		queryString.Set(i, v)
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
