package cdek

import (
	"context"
	"encoding/xml"
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
func (c Client) DeleteOrder(ctx context.Context, req DeleteOrderReq) (*DeleteOrderResp, error) {
	serverURL, err := url.Parse(c.apiURL)
	if err != nil {
		return nil, err
	}
	serverURL.Path = path.Join(serverURL.Path, deleteOrderURL)

	req.setAuth(c.auth)

	xmlBytes, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Add("xml_request", string(xmlBytes))

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, serverURL.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", urlFormEncoded)

	resp, err := xmlReq[DeleteOrderResp](r)
	if err != nil {
		return nil, err
	}

	var errs error
	for _, o := range resp.Order {
		if o.IsErroneous() {
			errs = multierror.Append(errs, o.GetError())
		}
	}

	if errs != nil {
		return nil, errs
	}

	return resp, nil
}
