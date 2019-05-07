package cdek

import "encoding/xml"

//DeleteOrderReq request structure for deleting order from CDEK
type DeleteOrderReq struct {
	securableXML
	XMLName    xml.Name     `xml:"DeleteRequest"`
	Number     *string      `xml:"Number,attr"`
	OrderCount *int         `xml:"OrderCount,attr"`
	Order      *DeleteOrder `xml:"Order"`
}

//DeleteOrder order model for deleting request
type DeleteOrder struct {
	Number         *string `xml:"Number,attr"`
	DispatchNumber *int    `xml:"DispatchNumber,attr"`
}

//DeleteOrderResp response structure of deleting order from CDEK
type DeleteOrderResp struct {
	Order []*OrderResp `xml:"Order"`
}
