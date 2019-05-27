package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const updateOrderUrl = "update"

func updateOrder(clientConf ClientConf, req UpdateOrderReq) (*UpdateOrderRes, error) {
	req.setAuth(clientConf.Auth)
	reqByte, err := xml.Marshal(req)

	if err != nil {
		return nil, err
	}

	data := make(url.Values)
	data.Add("xml_request", string(reqByte))

	serverUrl, err := url.Parse(clientConf.XmlApiUrl)
	if err != nil {
		return nil, err
	}

	serverUrl.Path = path.Join(serverUrl.Path, updateOrderUrl)
	reqUrl := serverUrl.String()

	res, err := http.Post(reqUrl, urlFormEncoded, strings.NewReader(data.Encode()))
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

	var updateOrderRes UpdateOrderRes
	err = xml.Unmarshal(body, &updateOrderRes)
	if err != nil {
		return nil, err
	}

	return &updateOrderRes, nil
}
