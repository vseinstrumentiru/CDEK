package cdek

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const statusReportURL = "/status_report_h.php"

//GetStatusReport This method is used to generate an order status report, including order change history.
func (c Client) GetStatusReport(statusReportReq StatusReport) (*StatusReportResp, error) {
	statusReportReq.setAuth(c.auth)
	reqByte, err := xml.Marshal(statusReportReq)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(c.apiURL)
	if err != nil {
		return nil, err
	}

	serverURL.Path = path.Join(serverURL.Path, statusReportURL)
	reqURL := serverURL.String()

	resp, err := http.Post(reqURL, xmlContentType, bytes.NewReader(reqByte))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	var statusReportResp StatusReportResp
	err = xml.Unmarshal(body, &statusReportResp)
	if err != nil {
		return nil, err
	}
	if statusReportResp.IsErroneous() {
		return nil, statusReportResp.Error
	}

	return &statusReportResp, nil
}
