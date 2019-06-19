package cdek

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const calculatorURL = "http://api.cdek.ru/calculator/calculate_price_by_json.php"

func calculateDelivery(clientConf ClientConf, req GetCostReq) (*GetCostResp, error) {
	req.setAuth(clientConf.Auth)
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

	if getCostResp.Error != nil {
		var errorMsg string
		for _, errorResp := range getCostResp.Error {
			errorMsg += fmt.Sprintf("Error code: %d, error text: %s \n", errorResp.Code, errorResp.Text, )
		}

		return nil, errors.New(errorMsg)
	}

	return &getCostResp, nil
}
