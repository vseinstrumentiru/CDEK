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
	registerOrderURL = "new_orders.php"
)

//RegisterOrder This method is used to register orders to be delivered to clients.
func (c Client) RegisterOrder(ctx context.Context, req RegisterOrderReq) (*RegisterOrderResp, error) {
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

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, serverURL.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", urlFormEncoded)

	resp, err := xmlReq[RegisterOrderResp](r)
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
