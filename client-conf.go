package cdek

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

const calculatorURLDefault = "http://api.cdek.ru/calculator/calculate_price_by_json.php"

//ClientConf SDK client configuration
type ClientConf struct {
	Auth          *Auth
	CdekAPIURL    string
	CalculatorURL string
}

//NewClientConf ClientConf constructor with defaults
func NewClientConf(cdekAPIURL string) *ClientConf {
	clientConf := ClientConf{
		CdekAPIURL:    cdekAPIURL,
		CalculatorURL: calculatorURLDefault,
	}

	return &clientConf
}

//SetAuth set auth data
func (clientConf *ClientConf) SetAuth(auth *Auth) *ClientConf {
	clientConf.Auth = auth

	return clientConf
}

//SetCalculatorURL url for delivery calculation
func (clientConf *ClientConf) SetCalculatorURL(calculatorURL string) *ClientConf {
	clientConf.CalculatorURL = calculatorURL

	return clientConf
}

//Auth account credentials
type Auth struct {
	Account string
	Secure  string
}

//EncodedSecure encode secure according to CDEK api
func (a Auth) EncodedSecure() (date string, encodedSecure string) {
	date = time.Now().Format("2006-01-02")
	encoder := md5.New()
	_, _ = encoder.Write([]byte(date + "&" + a.Secure))

	return date, hex.EncodeToString(encoder.Sum(nil))
}
