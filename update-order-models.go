package cdek

import "encoding/xml"

type UpdateOrderReq struct {
	XMLName    xml.Name     `xml:"UpdateRequest"`
	Number     *string      `xml:"Number,attr"`
	Date       *string      `xml:"Date,attr"`
	Account    *string      `xml:"Account,attr"`
	Secure     *string      `xml:"Secure,attr"`
	OrderCount *string      `xml:"OrderCount,attr"`
	Order      *UpdateOrder `xml:"Order"`
}

type UpdateOrder struct {
	Number                   *string                   `xml:"Number,attr"`
	DispatchNumber           *int                      `xml:"DispatchNumber,attr"`
	DeliveryRecipientCost    *float64                  `xml:"DeliveryRecipientCost,attr"`
	DeliveryRecipientVATRate *string                   `xml:"DeliveryRecipientVATRate,attr"`
	DeliveryRecipientVATSum  *float64                  `xml:"DeliveryRecipientVATSum,attr"`
	RecipientName            *string                   `xml:"RecipientName,attr"`
	Phone                    *string                   `xml:"Phone,attr"`
	RecipientINN             *string                   `xml:"RecipientINN,attr"`
	DateInvoice              *string                   `xml:"DateInvoice,attr"`
	Passport                 *Passport                 `xml:"Passport"`
	Address                  *Address                  `xml:"Address"`
	DeliveryRecipientCostAdv *DeliveryRecipientCostAdv `xml:"DeliveryRecipientCostAdv"`
	Package                  *OrderPackage             `xml:"Package"`
}

type UpdateOrderResp struct {
	XMLName xml.Name     `xml:"Response"`
	Order   []*OrderResp `xml:"Order"`
}
