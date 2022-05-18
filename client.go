package cdek

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"
)

const (
	ApiVersionV1 = "1.0"

	jsonContentType = "application/json"
	urlFormEncoded  = "application/x-www-form-urlencoded"

	calculatorURLDefault = "http://api.cdek.ru/calculator/calculate_price_by_json.php"
)

//ClientСonfigurator allows to configure client for the service
type ClientСonfigurator interface {
	SetAuth(account, secure string)
	SetCalculatorURL(calculatorURL string)
}

//Client provides CDEK API functionality
type Client interface {
	ClientСonfigurator

	CalculateDelivery(ctx context.Context, req *CalculateDeliveryRequest) (*CalculateDeliveryResult, error)
	GetCities(ctx context.Context, filter map[CityFilter]string) (*GetCitiesResp, error)
	GetPvzList(ctx context.Context, filter map[PvzListFilter]string) ([]*Pvz, error)
	GetRegions(ctx context.Context, filter map[RegionFilter]string) (*GetRegionsResp, error)
	RegisterOrder(ctx context.Context, req RegisterOrderReq) (*RegisterOrderResp, error)
	UpdateOrder(ctx context.Context, req UpdateOrderReq) (*UpdateOrderResp, error)
	DeleteOrder(ctx context.Context, req DeleteOrderReq) (*DeleteOrderResp, error)
	GetStatusReport(ctx context.Context, statusReportReq StatusReport) (*StatusReportResp, error)
}

//clientImpl SDK clientImpl configuration
type clientImpl struct {
	auth          *auth
	apiURL        string
	calculatorURL string
	httpClient    *http.Client
}

//NewClient clientImpl constructor with defaults
func NewClient(apiURL string) Client {
	return &clientImpl{
		apiURL:        apiURL,
		calculatorURL: calculatorURLDefault,
	}
}

//SetAuth set auth data
func (c *clientImpl) SetAuth(account, secure string) {
	c.auth = &auth{
		account: account,
		secure:  secure,
	}
}

//SetCalculatorURL url for delivery calculation
func (c *clientImpl) SetCalculatorURL(calculatorURL string) {
	c.calculatorURL = calculatorURL
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
