package cdek

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

const (
	jsonContentType = "application/json"
	urlFormEncoded  = "application/x-www-form-urlencoded"

	calculatorURLDefault = "http://api.cdek.ru/calculator/calculate_price_by_json.php"
)

//ClientСonfigurator allows to configure client for the service
type ClientСonfigurator interface {
	SetAuth(account, secure string) Client
	SetCalculatorURL(calculatorURL string) Client
}

//Client provides CDEK API functionality
type Client interface {
	ClientСonfigurator

	CalculateDelivery(req GetCostReq) (*GetCostRespResult, error)
	GetCities(filter map[CityFilter]string) (*GetCitiesResp, error)
	GetPvzList(filter map[PvzListFilter]string) ([]*Pvz, error)
	GetRegions(filter map[RegionFilter]string) (*GetRegionsResp, error)
	RegisterOrder(req RegisterOrderReq) (*RegisterOrderResp, error)
	UpdateOrder(req UpdateOrderReq) (*UpdateOrderResp, error)
	DeleteOrder(req DeleteOrderReq) (*DeleteOrderResp, error)
	GetStatusReport(statusReportReq StatusReport) (*StatusReportResp, error)
}

//clientImpl SDK clientImpl configuration
type clientImpl struct {
	auth          *auth
	apiURL        string
	calculatorURL string
}

//NewClient clientImpl constructor with defaults
func NewClient(apiURL string) Client {
	return &clientImpl{
		apiURL:        apiURL,
		calculatorURL: calculatorURLDefault,
	}
}

//SetAuth set auth data
func (c *clientImpl) SetAuth(account, secure string) Client {
	c.auth = &auth{
		account: account,
		secure:  secure,
	}

	return c
}

//SetCalculatorURL url for delivery calculation
func (c *clientImpl) SetCalculatorURL(calculatorURL string) Client {
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
