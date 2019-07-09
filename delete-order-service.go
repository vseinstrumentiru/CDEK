package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const deleteOrderURL = "delete_orders.php"

func (cl client) DeleteOrder(req DeleteOrderReq) (*DeleteOrderResp, error) {
	req.setAuth(cl.clientConf.Auth)
	reqByte, err := xml.Marshal(req)

	if err != nil {
		return nil, err
	}

	data := make(url.Values)
	data.Add("xml_request", string(reqByte))

	serverURL, err := url.Parse(cl.clientConf.CdekAPIURL)
	if err != nil {
		return nil, err
	}

	serverURL.Path = path.Join(serverURL.Path, deleteOrderURL)
	reqURL := serverURL.String()

	resp, err := http.Post(reqURL, urlFormEncoded, strings.NewReader(data.Encode()))
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

	var deleteOrderResp DeleteOrderResp
	err = xml.Unmarshal(body, &deleteOrderResp)
	if err != nil {
		return nil, err
	}

	return &deleteOrderResp, nil
}
