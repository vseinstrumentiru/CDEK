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
	deleteOrderURL = "delete_orders.php"
)

//DeleteOrder The method is designed to cancel/delete an order at the client's initiative.
func (c clientImpl) DeleteOrder(req DeleteOrderReq) (*DeleteOrderResp, error) {
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

	serverURL.Path = path.Join(serverURL.Path, deleteOrderURL)
	reqURL := serverURL.String()

	resp, err := http.Post(reqURL, urlFormEncoded, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	var deleteOrderResp DeleteOrderResp
	err = xml.Unmarshal(body, &deleteOrderResp)
	if err != nil {
		return nil, err
	}

	multiError := &multierror.Error{}
	for _, o := range deleteOrderResp.Order {
		if o.IsErroneous() {
			multiError = multierror.Append(o.GetError())
		}
	}
	if multiError.Len() > 0 {
		return nil, multiError.ErrorOrNil()
	}

	return &deleteOrderResp, nil
}
