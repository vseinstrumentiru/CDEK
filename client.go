package cdek

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

const jsonContentType = "application/json"
const xmlContentType = "application/xml"
const urlFormEncoded = "application/x-www-form-urlencoded"

const calculatorURLDefault = "http://api.cdek.ru/calculator/calculate_price_by_json.php"

//Client SDK Client configuration
type Client struct {
	auth          *auth
	apiURL        string
	calculatorURL string
}

//NewClient Client constructor with defaults
func NewClient(apiURL string) *Client {
	return &Client{
		apiURL:        apiURL,
		calculatorURL: calculatorURLDefault,
	}
}

//SetAuth set auth data
func (c *Client) SetAuth(account, secure string) *Client {
	c.auth = &auth{
		account: account,
		secure:  secure,
	}

	return c
}

//SetCalculatorURL url for delivery calculation
func (c *Client) SetCalculatorURL(calculatorURL string) *Client {
	c.calculatorURL = calculatorURL

	return c
}

type auth struct {
	account string
	secure  string
}

func (a auth) encodedSecure() (date string, encodedSecure string) {
	date = time.Now().Format("2006-01-02")
	encoder := md5.New()
	_, _ = encoder.Write([]byte(date + "&" + a.secure))

	return date, hex.EncodeToString(encoder.Sum(nil))
}
