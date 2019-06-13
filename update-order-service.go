package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const updateOrderURL = "update"

func updateOrder(clientConf ClientConf, req UpdateOrderReq) (*UpdateOrderResp, error) {
	req.setAuth(clientConf.Auth)
	reqByte, err := xml.Marshal(req)

	if err != nil {
		return nil, err
	}

	data := make(url.Values)
	data.Add("xml_request", string(reqByte))

	serverURL, err := url.Parse(clientConf.XmlApiUrl)
	if err != nil {
		return nil, err
	}

	serverURL.Path = path.Join(serverURL.Path, updateOrderURL)
	reqURL := serverURL.String()

	resp, err := http.Post(reqURL, urlFormEncoded, strings.NewReader(data.Encode()))
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

	var updateOrderResp UpdateOrderResp
	err = xml.Unmarshal(body, &updateOrderResp)
	if err != nil {
		return nil, err
	}

	return &updateOrderResp, nil
}
