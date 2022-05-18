package cdek

import (
	"context"
	"encoding/xml"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	statusReportURL = "/status_report_h.php"
)

//GetStatusReport This method is used to generate an order status report, including order change history.
func (c Client) GetStatusReport(ctx context.Context, statusReportReq StatusReport) (*StatusReportResp, error) {
	statusReportReq.setAuth(c.auth)
	reqByte, err := xml.Marshal(statusReportReq)
	if err != nil {
		return nil, err
	}

	data := make(url.Values)
	data.Add("xml_request", string(reqByte))

	serverURL, err := url.Parse(c.apiURL)
	if err != nil {
		return nil, err
	}

	serverURL.Path = path.Join(serverURL.Path, statusReportURL)

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, serverURL.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", urlFormEncoded)

	resp, err := xmlReq[StatusReportResp](r)
	if err != nil {
		return nil, err
	}

	if resp.IsErroneous() {
		return nil, resp.Error
	}

	return resp, nil
}
