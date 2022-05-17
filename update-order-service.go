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
	updateOrderURL = "update"
)

//UpdateOrder This method is used to change a created order.
func (c clientImpl) UpdateOrder(req UpdateOrderReq) (*UpdateOrderResp, error) {
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

	serverURL.Path = path.Join(serverURL.Path, updateOrderURL)
	reqURL := serverURL.String()

	resp, err := http.Post(reqURL, urlFormEncoded, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	var updateOrderResp UpdateOrderResp
	err = xml.Unmarshal(body, &updateOrderResp)
	if err != nil {
		return nil, err
	}

	multiError := &multierror.Error{}
	for _, o := range updateOrderResp.Order {
		if o.IsErroneous() {
			multiError = multierror.Append(o.GetError())
		}
	}
	if multiError.Len() > 0 {
		return nil, multiError.ErrorOrNil()
	}

	return &updateOrderResp, nil
}
