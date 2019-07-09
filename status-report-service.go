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

func (cl client) GetStatusReport(statusReportReq StatusReportReq) (*StatusReportResp, error) {
	statusReportReq.setAuth(cl.clientConf.Auth)
	reqByte, err := xml.Marshal(statusReportReq)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(cl.clientConf.CdekAPIURL)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var statusReportResp StatusReportResp
	err = xml.Unmarshal(body, &statusReportResp)
	if err != nil {
		return nil, err
	}

	return &statusReportResp, nil
}
