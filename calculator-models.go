package cdek

type Good struct {
	Weight float64 `json:"weight"`
	Length int     `json:"length"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Volume float64 `json:"volume"`
}

type ServiceReq struct {
	Id    int `json:"id"`
	Param int `json:"param,omitempty"`
}

type ServiceResp struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Rate  float64 `json:"rate,omitempty"`
}

type GetCostReq struct {
	Version        *string       `json:"version"`
	AuthLogin      *string       `json:"authLogin,omitempty"`
	Secure         *string       `json:"secure,omitempty"`
	DateExecute    *string       `json:"dateExecute,omitempty"`
	SenderCityId   *int          `json:"senderCityId"`
	ReceiverCityId *int          `json:"receiverCityId"`
	TariffId       *int          `json:"tariffId"`
	Goods          []*Good       `json:"goods"`
	Services       []*ServiceReq `json:"services,omitempty"`
}

type GetCostResp struct {
	Error  []ErrorResp `json:"error,omitempty"`
	Result struct {
		Price             string        `json:"price"`
		DeliveryPeriodMin int           `json:"deliveryPeriodMin"`
		DeliveryPeriodMax int           `json:"deliveryPeriodMax"`
		DeliveryDateMin   string        `json:"deliveryDateMin"`
		DeliveryDateMax   string        `json:"deliveryDateMax"`
		TariffId          int           `json:"tariffId"`
		CashOnDelivery    float64       `json:"cashOnDelivery"`
		PriceByCurrency   float64       `json:"priceByCurrency"`
		Currency          string        `json:"currency"`
		PercentVAT        int           `json:"percentVAT"`
		Services          []ServiceResp `json:"services"`
	} `json:"result"`
}

type ErrorResp struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
