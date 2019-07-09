package cdek

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const registerOrderURL = "new_orders.php"

func (cl client) RegisterOrder(req RegisterOrderReq) (*RegisterOrderResp, error) {
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

	serverURL.Path = path.Join(serverURL.Path, registerOrderURL)
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

	var registerOrderResp RegisterOrderResp
	err = xml.Unmarshal(body, &registerOrderResp)
	if err != nil {
		return nil, err
	}

	if registerOrderResp.ErrorCode != nil {
		return nil, errors.New(*registerOrderResp.Msg)
	}

	return &registerOrderResp, nil
}
