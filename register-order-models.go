package cdek

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

//RegisterOrderReq Order registration request
type RegisterOrderReq struct {
	securable
	XMLName     xml.Name     `xml:"DeliveryRequest"`
	Number      *string      `xml:"Number,attr"`
	OrderCount  *int         `xml:"OrderCount,attr"`
	Currency    *string      `xml:"Currency,attr,omitempty"`
	Order       *OrderReq    `xml:"Order"`
	CallCourier *CallCourier `xml:"CallCourier"`
}

//OrderReq Shipment (order)
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

//Passport Details of the receiver’s passport. Used to print waybills. Only for international orders.
type Passport struct {
	Series    *string `xml:"Series,attr"`
	Number    *string `xml:"Number,attr"`
	IssueDate *string `xml:"IssueDate,attr"`
	IssuedBy  *string `xml:"IssuedBy,attr"`
	DateBirth *string `xml:"DateBirth,attr"`
}

//Sender Sender. Must be defined if it is different from the online store Client.
// If the online store is a sender, the Sender tag is not available.
type Sender struct {
	Company *string   `xml:"Company,attr"`
	Name    *string   `xml:"Name,attr"`
	Address *Address  `xml:"Address,omitempty"`
	Phone   []*string `xml:"Phone,omitempty"`
}

//Address Address of pickup
type Address struct {
	Street  *string `xml:"Street,attr"`
	House   *string `xml:"House,attr"`
	Flat    *string `xml:"Flat,attr,omitempty"`
	Phone   *string `xml:"Phone,attr,omitempty"`
	PvzCode *string `xml:"PvzCode,attr,omitempty"`
}

//Seller Requisites of the real seller
type Seller struct {
	Address       *string `xml:"Address,attr,omitempty"`
	Name          *string `xml:"Name,attr,omitempty"`
	INN           *string `xml:"INN,attr,omitempty"`
	Phone         *string `xml:"Phone,attr,omitempty"`
	OwnershipForm *int    `xml:"OwnershipForm,attr,omitempty"`
}

//OrderPackage Package (all packages are sent with different Package tags)
type OrderPackage struct {
	Number  *string             `xml:"Number,attr"`
	BarCode *string             `xml:"BarCode,attr"`
	Weight  *int                `xml:"Weight,attr"`
	SizeA   *int                `xml:"SizeA,attr,omitempty"`
	SizeB   *int                `xml:"SizeB,attr,omitempty"`
	SizeC   *int                `xml:"SizeC,attr,omitempty"`
	Item    []*OrderPackageItem `xml:"Item"`
}

//OrderPackageItem Items (goods)
type OrderPackageItem struct {
	Amount         *int     `xml:"Amount,attr"`
	WareKey        *string  `xml:"WareKey,attr"`
	Cost           *float64 `xml:"Cost,attr"`
	Payment        *float64 `xml:"Payment,attr"`
	PaymentVATRate *string  `xml:"PaymentVATRate,attr,omitempty"`
	PaymentVATSum  *float64 `xml:"PaymentVATSum,attr,omitempty"`
	Weight         *int     `xml:"Weight,attr"`
	Comment        *string  `xml:"Comment,attr"`
	WeightBrutto   *int     `xml:"WeightBrutto,attr,omitempty"`
	CommentEx      *string  `xml:"CommentEx,attr,omitempty"`
	Link           *string  `xml:"Link,attr,omitempty"`
}

//DeliveryRecipientCostAdv Additional charge for delivery (E-shop charges the receiver), depending on the order’s sum
type DeliveryRecipientCostAdv struct {
	Threshold *int     `xml:"Threshold,attr"`
	Sum       *float64 `xml:"Sum,attr"`
	VATRate   *string  `xml:"VATRate,attr,omitempty"`
	VATSum    *float64 `xml:"VATSum,attr,omitempty"`
}

//AddService Additional services
type AddService struct {
	ServiceCode *int `xml:"ServiceCode,attr"`
	Count       *int `xml:"Count,attr,omitempty"`
}

//Schedule Schedule for delivery/pickup
type Schedule struct {
	Attempt []*ScheduleAttempt `xml:"Attempt"`
}

//ScheduleAttempt Time of delivery (one time interval not less than 3 hours is permitted for one day)
type ScheduleAttempt struct {
	ID      *string  `xml:"ID,attr"`
	Date    *string  `xml:"Date,attr"`
	Comment *string  `xml:"Comment,attr"`
	TimeBeg *string  `xml:"TimeBeg,attr"`
	TimeEnd *string  `xml:"TimeEnd,attr"`
	Address *Address `xml:"Address"`
}

//CallCourier Call courier
type CallCourier struct {
	Call *CourierCallReq `xml:"Call"`
}

//CourierCallReq Waiting for a courier
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

//RegisterOrderResp Order registration response structure
type RegisterOrderResp struct {
	Order []*OrderResp       `xml:"Order,omitempty"`
	Call  []*CourierCallResp `xml:"Call,omitempty"`
}

//OrderResp Order
type OrderResp struct {
	Error
	DispatchNumber *int    `xml:"DispatchNumber,attr"`
	Number         *string `xml:"Number,attr"`
}

//CourierCallResp Call courier
type CourierCallResp struct {
	Error
	Number *string `xml:"Number,attr"`
}

//GetError returns error supplemented with order data
func (o *OrderResp) GetError() error {
	errorMsgParts := []string{
		o.Error.Error(),
	}
	if o.DispatchNumber != nil {
		errorMsgParts = append(errorMsgParts, fmt.Sprintf("DispatchNumber: %d", *o.DispatchNumber))
	}
	if o.Number != nil {
		errorMsgParts = append(errorMsgParts, fmt.Sprintf("Number: %s", *o.Number))
	}

	return errors.New(strings.Join(errorMsgParts, "; "))
}
