package cdek

//Good Location's dimension
type Good struct {
	Weight float64 `json:"weight"`
	Length int     `json:"length"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Volume float64 `json:"volume"`
}

//ServiceReq List of additional service
type ServiceReq struct {
	ID    int `json:"id"`
	Param int `json:"param,omitempty"`
}

//ServiceResp List of transmitted additional services
type ServiceResp struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Rate  float64 `json:"rate,omitempty"`
}

//GetCostReq Cost calculation on tariffs with priority request
type GetCostReq struct {
	Version        *string       `json:"version"`
	AuthLogin      *string       `json:"authLogin,omitempty"`
	Secure         *string       `json:"secure,omitempty"`
	DateExecute    *string       `json:"dateExecute,omitempty"`
	SenderCityID   *int          `json:"senderCityId"`
	ReceiverCityID *int          `json:"receiverCityId"`
	TariffID       *int          `json:"tariffId"`
	Goods          []*Good       `json:"goods"`
	Services       []*ServiceReq `json:"services,omitempty"`
}

//GetCostResp Cost calculation on tariffs with priority response
type GetCostResp struct {
	Error  []ErrorResp `json:"error,omitempty"`
	Result struct {
		Price             float64       `json:"price,string"`
		DeliveryPeriodMin int           `json:"deliveryPeriodMin"`
		DeliveryPeriodMax int           `json:"deliveryPeriodMax"`
		DeliveryDateMin   string        `json:"deliveryDateMin"`
		DeliveryDateMax   string        `json:"deliveryDateMax"`
		TariffID          int           `json:"tariffId"`
		CashOnDelivery    float64       `json:"cashOnDelivery"`
		PriceByCurrency   float64       `json:"priceByCurrency"`
		Currency          string        `json:"currency"`
		PercentVAT        int           `json:"percentVAT"`
		Services          []ServiceResp `json:"services"`
	} `json:"result"`
}

//ErrorResp error response
type ErrorResp struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
