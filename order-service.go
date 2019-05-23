package cdek

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const registerOrderUrl = "new_orders.php"

func registerOrder(clientConf ClientConf, req RegisterOrderReq) (*RegisterOrderRes, error) {
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

	serverUrl.Path = path.Join(serverUrl.Path, registerOrderUrl)
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

	var registerOrderRes RegisterOrderRes
	err = xml.Unmarshal(body, &registerOrderRes)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", string(body))

	if registerOrderRes.ErrorCode != nil {
		return nil, errors.New(*registerOrderRes.Msg)
	}

	return &registerOrderRes, nil
}
