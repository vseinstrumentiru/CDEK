package cdek

import "encoding/xml"

type RegisterOrderReq struct {
	XMLName     xml.Name     `xml:"DeliveryRequest"`
	Number      *string      `xml:"Number,attr"`
	Date        *string      `xml:"Date,attr"`
	Account     *string      `xml:"Account,attr"`
	Secure      *string      `xml:"Secure,attr"`
	OrderCount  *string      `xml:"OrderCount,attr"`
	Currency    *string      `xml:"Currency,attr,omitempty"`
	Order       *OrderReq    `xml:"Order"`
	CallCourier *CallCourier `xml:"CallCourier"`
}

type OrderReq struct {
	Number                   *string                   `xml:"Number,attr"`
	SendCityCode             *int                      `xml:"SendCityCode,attr"`
	RecCityCode              *int                      `xml:"RecCityCode,attr"`
	SendCityPostCode         *int                      `xml:"SendCityPostCode,attr"`
	RecCityPostCode          *int                      `xml:"RecCityPostCode,attr"`
	SendCountryCode          *int                      `xml:"SendCountryCode,attr,omitempty"`
	RecCountryCode           *int                      `xml:"RecCountryCode,attr,omitempty"`
	SendCityName             *string                   `xml:"SendCityName,attr,omitempty"`
	RecCityName              *string                   `xml:"RecCityName,attr,omitempty"`
	RecipientINN             *string                   `xml:"RecipientINN,attr,omitempty"`
	DateInvoice              *string                   `xml:"DateInvoice,attr,omitempty"`
	ShipperName              *string                   `xml:"ShipperName,attr,omitempty"`
	ShipperAddress           *string                   `xml:"ShipperAddress,attr,omitempty"`
	Passport                 *Passport                 `xml:"Passport,omitempty"`
	Sender                   *Sender                   `xml:"Sender,omitempty"`
	RecipientName            *string                   `xml:"RecipientName,attr"`
	RecipientEmail           *string                   `xml:"RecipientEmail,attr,omitempty"`
	Phone                    *string                   `xml:"Phone,attr"`
	TariffTypeCode           *int                      `xml:"TariffTypeCode,attr"`
	DeliveryRecipientCost    *float64                  `xml:"DeliveryRecipientCost,attr,omitempty"`
	DeliveryRecipientVATRate *string                   `xml:"DeliveryRecipientVATRate,attr,omitempty"`
	DeliveryRecipientVATSum  *float64                  `xml:"DeliveryRecipientVATSum,attr,omitempty"`
	RecipientCurrency        *string                   `xml:"RecipientCurrency,attr,omitempty"`
	ItemsCurrency            *string                   `xml:"ItemsCurrency,attr,omitempty"`
	Seller                   *Seller                   `xml:"Seller,omitempty"`
	Comment                  *string                   `xml:"Comment,attr,omitempty"`
	Address                  *Address                  `xml:"Address,omitempty"`
	Package                  []*OrderPackage           `xml:"Package,omitempty"`
	DeliveryRecipientCostAdv *DeliveryRecipientCostAdv `xml:"DeliveryRecipientCostAdv,omitempty"`
	AddService               *AddService               `xml:"AddService,omitempty"`
	Schedule                 *Schedule                 `xml:"Schedule,omitempty"`
}

type Passport struct {
	Series    *string `xml:"Series,attr"`
	Number    *string `xml:"Number,attr"`
	IssueDate *string `xml:"IssueDate,attr"`
	IssuedBy  *string `xml:"IssuedBy,attr"`
	DateBirth *string `xml:"DateBirth,attr"`
}

type Sender struct {
	Company *string  `xml:"Company,attr"`
	Name    *string  `xml:"Name,attr"`
	Address *Address `xml:"Address,omitempty"`
}

type Address struct {
	Street  *string `xml:"Street,attr"`
	House   *string `xml:"House,attr"`
	Flat    *string `xml:"Flat,attr,omitempty"`
	Phone   *string `xml:"Phone,attr,omitempty"`
	PvzCode *string `xml:"PvzCode,attr,omitempty"`
}

type Seller struct {
	Address       *string `xml:"Address,attr,omitempty"`
	Name          *string `xml:"Name,attr,omitempty"`
	INN           *string `xml:"INN,attr,omitempty"`
	Phone         *string `xml:"Phone,attr,omitempty"`
	OwnershipForm *int    `xml:"OwnershipForm,attr,omitempty"`
}

type OrderPackage struct {
	Number  *string             `xml:"Number,attr"`
	BarCode *string             `xml:"BarCode,attr"`
	Weight  *int                `xml:"Weight,attr"`
	SizeA   *int                `xml:"SizeA,attr,omitempty"`
	SizeB   *int                `xml:"SizeB,attr,omitempty"`
	SizeC   *int                `xml:"SizeC,attr,omitempty"`
	Item    []*OrderPackageItem `xml:"Item"`
}

type OrderPackageItem struct {
	Amount         *int     `xml:"Amount,attr"`
	Warekey        *string  `xml:"Warekey,attr"`
	Cost           *float64 `xml:"Cost,attr"`
	Payment        *float64 `xml:"Payment,attr"`
	PaymentVATRate *string  `xml:"PaymentVATRate,attr,omitempty"`
	PaymentVATSum  *float64 `xml:"PaymentVATSum,attr,omitempty"`
	Weight         *int     `xml:"Weight,attr"`
	Comment        *string  `xml:"Comment,attr"`
	WeightBrutto   *string  `xml:"WeightBrutto,attr,omitempty"`
	CommentEx      *string  `xml:"CommentEx,attr,omitempty"`
	Link           *string  `xml:"Link,attr,omitempty"`
}

type DeliveryRecipientCostAdv struct {
	Threshold *int     `xml:"Threshold,attr"`
	Sum       *float64 `xml:"Sum,attr"`
	VATRate   *string  `xml:"VATRate,attr,omitempty"`
	VATSum    *float64 `xml:"VATSum,attr,omitempty"`
}

type AddService struct {
	ServiceCode *int `xml:"ServiceCode,attr"`
	Count       *int `xml:"Count,attr,omitempty"`
}

type Schedule struct {
	Attempt []*ScheduleAttempt `xml:"Attempt"`
}

type ScheduleAttempt struct {
	ID      *string  `xml:"ID,attr"`
	Date    *string  `xml:"Date,attr"`
	Comment *string  `xml:"Comment,attr"`
	TimeBeg *string  `xml:"TimeBeg,attr"`
	TimeEnd *string  `xml:"TimeEnd,attr"`
	Address *Address `xml:"Address"`
}

type CallCourier struct {
	Call *CourierCallReq `xml:"Call"`
}

type CourierCallReq struct {
	Date             *string  `xml:"Date,attr"`
	TimeBeg          *string  `xml:"TimeBeg,attr"`
	TimeEnd          *string  `xml:"TimeEnd,attr"`
	LunchBeg         *string  `xml:"LunchBeg,attr"`
	LunchEnd         *string  `xml:"LunchEnd,attr"`
	SendCityCode     *int     `xml:"SendCityCode,attr"`
	SendCityPostCode *string  `xml:"SendCityPostCode,attr"`
	SendCountryCode  *string  `xml:"SendCountryCode,attr"`
	SendCityName     *string  `xml:"SendCityName,attr"`
	SendPhone        *string  `xml:"SendPhone,attr"`
	SenderName       *string  `xml:"SenderName,attr"`
	Comment          *string  `xml:"Comment,attr"`
	SendAddress      *Address `xml:"SendAddress"`
}

type RegisterOrderRes struct {
	ErrorCode       *string               `xml:"ErrorCode,attr,omitempty"`
	Msg             *string               `xml:"Msg,attr,omitempty"`
	DeliveryRequest []*DeliveryRequestRes `xml:"DeliveryRequest,omitempty"`
	Order           []*OrderRes           `xml:"Order,omitempty"`
	Call            *CourierCallRes       `xml:"Call,omitempty"`
}

type DeliveryRequestRes struct {
	Number    string `xml:"Number,attr"`
	ErrorCode string `xml:"ErrorCode,attr"`
	Msg       string `xml:"Msg,attr"`
}

type OrderRes struct {
	DispatchNumber *int    `xml:"DispatchNumber,attr"`
	Number         *string `xml:"Number,attr"`
	ErrorCode      *string `xml:"ErrorCode,attr,omitempty"`
	Msg            *string `xml:"Msg,attr"`
}

type CourierCallRes struct {
	Number    *string `xml:"Number,attr"`
	ErrorCode *string `xml:"ErrorCode,attr,omitempty"`
	Msg       *string `xml:"Msg,attr"`
}
