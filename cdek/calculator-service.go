package cdek

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const calculatorUrl = "http://api.cdek.ru/calculator/calculate_price_by_json.php"

func calculateDelivery(request GetCostRequest) (*GetCostResponse, error) {
	requestByte, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(calculatorUrl, jsonContentType, bytes.NewReader(requestByte))
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

	var getCostResponse GetCostResponse
	err = json.Unmarshal(body, &getCostResponse)
	if err != nil {
		return nil, err
	}

	return &getCostResponse, nil
}
