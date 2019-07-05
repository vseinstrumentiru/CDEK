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

func getStatusReport(clientConf ClientConf, req StatusReportReq) (*StatusReportResp, error) {
	req.setAuth(clientConf.Auth)
	reqByte, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(clientConf.CdekAPIURL)
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
		err = resp.Body.Close()
	}()
	if err != nil {
		return nil, err
	}

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
