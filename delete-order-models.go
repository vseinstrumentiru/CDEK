package cdek

import "encoding/xml"

type DeleteOrderReq struct {
	XMLName    xml.Name     `xml:"DeleteRequest"`
	Number     *string      `xml:"Number,attr"`
	Date       *string      `xml:"Date,attr"`
	Account    *string      `xml:"Account,attr"`
	Secure     *string      `xml:"Secure,attr"`
	OrderCount *string      `xml:"OrderCount,attr"`
	Order      *DeleteOrder `xml:"Order"`
}

type DeleteOrder struct {
	Number         *string `xml:"Number,attr"`
	DispatchNumber *int    `xml:"DispatchNumber,attr"`
}

type DeleteOrderRes struct {
	XMLName xml.Name    `xml:"Response"`
	Order   []*OrderRes `xml:"Order"`
}
