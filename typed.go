package cdek

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func jsonReq[T any](req *http.Request) (*T, error) {
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var s T
	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(payload, &s); err != nil {
		return nil, fmt.Errorf("%s", payload)
	}

	return &s, nil
}

func xmlReq[T any](req *http.Request) (*T, error) {
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var s T
	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(payload, &s); err != nil {
		return nil, fmt.Errorf("%s", payload)
	}

	return &s, nil
}
