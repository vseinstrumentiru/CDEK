package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ernesto-jimenez/httplogger"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func validateResponse(requests []ResponseRequests) error {
	var result error
	for _, item := range requests {
		if item.State == "INVALID" {
			result = multierror.Append(result, fmt.Errorf("%+v", item))
		}
	}

	return result
}

var client = http.Client{
	Transport: httplogger.NewLoggedTransport(http.DefaultTransport, newLogger()),
}

type RespErrors struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}

func jsonReq[T any](req *http.Request) (*T, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http.DefaultClient.Do")
	}
	defer response.Body.Close()

	var s T
	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}

	var respErr RespErrors
	if err := json.Unmarshal(payload, &respErr); err == nil && len(respErr.Errors) > 0 {
		return nil, fmt.Errorf("json error: %v", respErr)
	}

	if err := json.Unmarshal(payload, &s); err != nil {
		return nil, fmt.Errorf("json error: %s", payload)
	}

	return &s, nil
}

type httpLogger struct {
	log *log.Logger
}

func newLogger() *httpLogger {
	return &httpLogger{
		log: log.New(os.Stderr, "log - ", log.LstdFlags),
	}
}

func (l *httpLogger) LogRequest(req *http.Request) {
	if req.Body != nil {
		body, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewReader(body))

		l.log.Printf(
			"Request %s %s %s",
			req.Method,
			req.URL.String(),
			body,
		)
	} else {
		l.log.Printf(
			"Request %s %s",
			req.Method,
			req.URL.String(),
		)
	}
}

func (l *httpLogger) LogResponse(req *http.Request, res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	if err != nil {
		l.log.Println(err)
	} else {
		l.log.Printf(
			"Response method=%s status=%d durationMs=%d %s",
			req.Method,
			res.StatusCode,
			duration,
			req.URL.String(),
		)
	}
}
