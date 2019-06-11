package cdek

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const statusReportUrl = "/status_report_h.php"

func getStatusReport(clientConf ClientConf, req StatusReportReq) (*StatusReportResp, error) {
	req.setAuth(clientConf.Auth)
	reqByte, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	serverUrl, err := url.Parse(clientConf.XmlApiUrl)
	if err != nil {
		return nil, err
	}

	serverUrl.Path = path.Join(serverUrl.Path, statusReportUrl)
	reqUrl := serverUrl.String()

	resp, err := http.Post(reqUrl, xmlContentType, bytes.NewReader(reqByte))
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
