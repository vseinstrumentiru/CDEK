package v2

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func jsonReq[T any](req *http.Request) (*T, error) {
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http.DefaultClient.Do")
	}
	defer response.Body.Close()

	var s T
	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}

	//fmt.Printf("%s\n", payload)

	if err := json.Unmarshal(payload, &s); err != nil {
		return nil, fmt.Errorf("%s", payload)
	}

	return &s, nil
}
