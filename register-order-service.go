package cdek

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/hashicorp/go-multierror"
)

const (
	registerOrderURL = "new_orders.php"
)

//RegisterOrder This method is used to register orders to be delivered to clients.
func (c Client) RegisterOrder(req RegisterOrderReq) (*RegisterOrderResp, error) {
	req.setAuth(c.auth)
	reqByte, err := xml.Marshal(req)

	if err != nil {
		return nil, err
	}

	data := make(url.Values)
	data.Add("xml_request", string(reqByte))

	serverURL, err := url.Parse(c.apiURL)
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

	body, _ := ioutil.ReadAll(resp.Body)

	var registerOrderResp RegisterOrderResp
	err = xml.Unmarshal(body, &registerOrderResp)
	if err != nil {
		return nil, err
	}

	multiError := &multierror.Error{}
	for _, o := range registerOrderResp.Order {
		if o.IsErroneous() {
			multiError = multierror.Append(o.GetError())
		}
	}
	for _, c := range registerOrderResp.Call {
		if c.IsErroneous() {
			multiError = multierror.Append(c.Error)
		}
	}
	if multiError.Len() > 0 {
		return nil, multiError.ErrorOrNil()
	}

	return &registerOrderResp, nil
}
