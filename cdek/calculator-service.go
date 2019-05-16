package cdek

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const calculatorUrl = "http://api.cdek.ru/calculator/calculate_price_by_json.php"

func calculateDelivery(req GetCostReq) (*GetCostRes, error) {
	reqByte, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(calculatorUrl, jsonContentType, bytes.NewReader(reqByte))
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

	var getCostRes GetCostRes
	err = json.Unmarshal(body, &getCostRes)
	if err != nil {
		return nil, err
	}

	return &getCostRes, nil
}
