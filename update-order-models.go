package cdek

import "encoding/xml"

//UpdateOrderReq Order Change request structure
type UpdateOrderReq struct {
	credentialsXML
	XMLName    xml.Name     `xml:"UpdateRequest"`
	Number     *string      `xml:"Number,attr"`
	OrderCount *int         `xml:"OrderCount,attr"`
	Order      *UpdateOrder `xml:"Order"`
}

//UpdateOrder Order Change request
type UpdateOrder struct {
	Number                   *string                   `xml:"Number,attr"`
	DispatchNumber           *int                      `xml:"DispatchNumber,attr"`
	DeliveryRecipientCost    *float64                  `xml:"DeliveryRecipientCost,attr"`
	DeliveryRecipientVATRate *string                   `xml:"DeliveryRecipientVATRate,attr"`
	DeliveryRecipientVATSum  *float64                  `xml:"DeliveryRecipientVATSum,attr"`
	RecipientName            *string                   `xml:"RecipientName,attr"`
	RecipientEmail           *string                   `xml:"RecipientEmail,attr"`
	Phone                    *string                   `xml:"Phone,attr"`
	RecipientINN             *string                   `xml:"RecipientINN,attr"`
	DateInvoice              *string                   `xml:"DateInvoice,attr"`
	Passport                 *Passport                 `xml:"Passport"`
	Address                  *Address                  `xml:"Address"`
	DeliveryRecipientCostAdv *DeliveryRecipientCostAdv `xml:"DeliveryRecipientCostAdv"`
	Package                  *OrderPackage             `xml:"Package"`
}

//UpdateOrderResp Order Change response
type UpdateOrderResp struct {
	Order []*OrderResp `xml:"Order"`
}
