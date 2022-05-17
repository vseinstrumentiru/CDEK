package cdek

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//CalculateDelivery Cost calculation on tariffs with priority.
func (c clientImpl) CalculateDelivery(req GetCostReq) (*GetCostRespResult, error) {
	req.setAuth(c.auth)
	reqByte, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.calculatorURL, jsonContentType, bytes.NewReader(reqByte))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	var getCostResp getCostResp
	err = json.Unmarshal(body, &getCostResp)
	if err != nil {
		return nil, err
	}

	if getCostResp.ErrorResp != nil {
		var errorMsg string
		for _, err := range getCostResp.ErrorResp {
			errorMsg += fmt.Sprintf("Error code: %s, error text: %s \n", *err.ErrorCode, *err.Msg)
		}

		return nil, errors.New(errorMsg)
	}

	return &getCostResp.Result, nil
}
