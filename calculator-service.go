package cdek

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const calculatorURL = "http://api.cdek.ru/calculator/calculate_price_by_json.php"

func calculateDelivery(clientConf ClientConf, req GetCostReq) (*GetCostResp, error) {
	req.AuthLogin = clientConf.Auth.Account
	req.DateExecute, req.Secure = clientConf.Auth.EncodedSecure()
	reqByte, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(calculatorURL, jsonContentType, bytes.NewReader(reqByte))
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

	var getCostResp GetCostResp
	err = json.Unmarshal(body, &getCostResp)
	if err != nil {
		return nil, err
	}

	return &getCostResp, nil
}
