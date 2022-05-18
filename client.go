package cdek

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"time"
)

const (
	jsonContentType = "application/json"
	urlFormEncoded  = "application/x-www-form-urlencoded"

	calculatorURLDefault = "http://api.cdek.ru/calculator/calculate_price_by_json.php"
)

//ServiceAccessСonfigurator allows to configure client for the service
type ServiceAccessСonfigurator interface {
	SetAuth(account, secure string) ServiceProvider
	SetCalculatorURL(calculatorURL string) ServiceProvider
}

//ServiceProvider provides CDEK API functionality
type ServiceProvider interface {
	ServiceAccessСonfigurator

	CalculateDelivery(ctx context.Context, req GetCostReq) (*GetCostRespResult, error)
	GetCities(ctx context.Context, filter map[CityFilter]string) (*GetCitiesResp, error)
	GetPvzList(filter map[PvzListFilter]string) ([]*Pvz, error)
	GetRegions(filter map[RegionFilter]string) (*GetRegionsResp, error)
	RegisterOrder(req RegisterOrderReq) (*RegisterOrderResp, error)
	UpdateOrder(req UpdateOrderReq) (*UpdateOrderResp, error)
	DeleteOrder(req DeleteOrderReq) (*DeleteOrderResp, error)
	GetStatusReport(statusReportReq StatusReport) (*StatusReportResp, error)
}

//Client SDK Client configuration
type Client struct {
	auth          *auth
	apiURL        string
	calculatorURL string
}

//NewClient Client constructor with defaults
func NewClient(apiURL string) ServiceProvider {
	return &Client{
		apiURL:        apiURL,
		calculatorURL: calculatorURLDefault,
	}
}

//SetAuth set auth data
func (c *Client) SetAuth(account, secure string) ServiceProvider {
	c.auth = &auth{
		account: account,
		secure:  secure,
	}

	return c
}

//SetCalculatorURL url for delivery calculation
func (c *Client) SetCalculatorURL(calculatorURL string) ServiceProvider {
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
