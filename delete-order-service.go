package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const deleteOrderUrl = "delete_orders.php"

func deleteOrder(clientConf ClientConf, req DeleteOrderReq) (*DeleteOrderResp, error) {
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

	serverUrl.Path = path.Join(serverUrl.Path, deleteOrderUrl)
	reqUrl := serverUrl.String()

	resp, err := http.Post(reqUrl, urlFormEncoded, strings.NewReader(data.Encode()))
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

	var deleteOrderResp DeleteOrderResp
	err = xml.Unmarshal(body, &deleteOrderResp)
	if err != nil {
		return nil, err
	}

	return &deleteOrderResp, nil
}
