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

func getStatusReport(clientConfig ClientConfig, req StatusReportReq) (*StatusReportRes, error) {
	reqByte, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	serverUrl, err := url.Parse(clientConfig.XmlApiUrl)
	if err != nil {
		return nil, err
	}

	serverUrl.Path = path.Join(serverUrl.Path, statusReportUrl)
	reqUrl := serverUrl.String()

	res, err := http.Post(reqUrl, xmlContentType, bytes.NewReader(reqByte))
	if err != nil {
		return nil, err
	}

	defer func() {
		err = res.Body.Close()
	}()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var statusReportRes StatusReportRes
	err = xml.Unmarshal(body, &statusReportRes)
	if err != nil {
		return nil, err
	}

	return &statusReportRes, nil
}
