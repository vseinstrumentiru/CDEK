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
	securableJSON
	Version        *string       `json:"version"`
	SenderCityID   *int          `json:"senderCityId"`
	ReceiverCityID *int          `json:"receiverCityId"`
	TariffID       *int          `json:"tariffId"`
	Goods          []*Good       `json:"goods"`
	Services       []*ServiceReq `json:"services,omitempty"`
}

type getCostResp struct {
	ErrorResp []Error           `json:"error,omitempty"`
	Result    GetCostRespResult `json:"result"`
}

//GetCostRespResult Cost calculation on tariffs with priority result response
type GetCostRespResult struct {
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
}
