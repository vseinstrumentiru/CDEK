package cdek

//StatusReport Order Status Report request
type StatusReport struct {
	securableXML
	ShowHistory            *bool                   `xml:"ShowHistory,attr"`
	ShowReturnOrder        *bool                   `xml:"ShowReturnOrder,attr"`
	ShowReturnOrderHistory *bool                   `xml:"ShowReturnOrderHistory,attr"`
	ChangePeriod           *ChangePeriod           `xml:"ChangePeriod"`
	Order                  []*StatusReportOrderReq `xml:"Order"`
}

//StatusReportResp Order Status Report response
type StatusReportResp struct {
	Error
	StatusReport *StatusReportContentForResp `xml:"StatusReport"`
	Order        []*StatusReportOrderResp    `xml:"Order"`
	ReturnOrder  []*ReturnOrder              `xml:"ReturnOrder,omitempty"`
}

//ChangePeriod The period during which the order status has changed.
type ChangePeriod struct {
	DateFirst *string `xml:"DateFirst,attr"`
	DateLast  *string `xml:"DateLast,attr"`
}

//StatusReportOrderReq Shipment (order)
type StatusReportOrderReq struct {
	DispatchNumber *int    `xml:"DispatchNumber,attr"`
	Number         *string `xml:"Number,attr"`
	Date           *string `xml:"Date,attr"`
}

//StatusReportContentForResp The period during which the order status has changed.
type StatusReportContentForResp struct {
	DateFirst *string `xml:"DateFirst,attr"`
	DateLast  *string `xml:"DateLast,attr"`
}

//StatusReportOrderResp Shipment (order)
type StatusReportOrderResp struct {
	ActNumber            *string      `xml:"ActNumber,attr"`
	Number               *string      `xml:"Number,attr"`
	DispatchNumber       *string      `xml:"DispatchNumber,attr"`
	DeliveryDate         *string      `xml:"DeliveryDate,attr"`
	RecipientName        *string      `xml:"RecipientName,attr"`
	ReturnDispatchNumber *int         `xml:"ReturnDispatchNumber,attr"`
	Status               *Status      `xml:"Status"`
	Reason               *Reason      `xml:"Reason"`
	DelayReason          *DelayReason `xml:"DelayReason"`
	Package              *Package     `xml:"Package"`
	Attempt              *Attempt     `xml:"Attempt,omitempty"`
	Call                 *Call        `xml:"Call"`
}

//Status Current order status
type Status struct {
	Date        *string  `xml:"Date,attr"`
	Code        *string  `xml:"Code,attr"`
	Description *string  `xml:"Description,attr"`
	CityCode    *string  `xml:"CityCode,attr"`
	CityName    *string  `xml:"CityName,attr"`
	State       []*State `xml:"State"`
}

//State Status change history
type State struct {
	Date        *string `xml:"Date,attr"`
	Code        *string `xml:"Code,attr"`
	Description *string `xml:"Description,attr"`
	CityCode    *string `xml:"CityCode,attr,omitempty"`
	CityName    *string `xml:"CityName,attr,omitempty"`
}

//Reason Current additional status
type Reason struct {
	Date        *string `xml:"Date,attr"`
	Code        *string `xml:"Code,attr"`
	Description *string `xml:"Description,attr"`
}

//DelayReason Current delay reason
type DelayReason struct {
	Date        *string `xml:"Date,attr"`
	Code        *string `xml:"Code,attr"`
	Description *string `xml:"Description,attr"`
	State       []State `xml:"State"`
}

//Package Package
type Package struct {
	Number  *string `xml:"Number,attr"`
	BarCode *string `xml:"BarCode,attr"`
	Item    []*Item `xml:"Item"`
}

//Item Items
type Item struct {
	WareKey     *string `xml:"WareKey,attr"`
	Amount      *string `xml:"Amount,attr"`
	DelivAmount *string `xml:"DelivAmount,attr"`
}

//Attempt Delivery time taken from the delivery schedule
type Attempt struct {
	ID                  *int    `xml:"ID,attr"`
	ScheduleCode        *int    `xml:"ScheduleCode,attr"`
	ScheduleDescription *string `xml:"ScheduleDescription,attr"`
}

//Call History of notification calls to the receiver
type Call struct {
	CallGood  *CallGood  `xml:"CallGood"`
	CallFail  *CallFail  `xml:"CallFail"`
	CallDelay *CallDelay `xml:"CallDelay"`
}

//CallGood History of successful calls
type CallGood struct {
	Good []*CallGoodItem `xml:"Good"`
}

//CallGoodItem Successful call
type CallGoodItem struct {
	Date      *string `xml:"Date,attr"`
	DateDeliv *string `xml:"DateDeliv,attr"`
}

//CallFail History of failed calls
type CallFail struct {
	Fail []*CallFailItem `xml:"Fail"`
}

//CallFailItem Failed call
type CallFailItem struct {
	Date              *string `xml:"Date,attr"`
	ReasonCode        *string `xml:"ReasonCode,attr"`
	ReasonDescription *string `xml:"ReasonDescription,attr"`
}

//CallDelay History of call reschedules
type CallDelay struct {
	Delay []*CallDelayItem `xml:"Delay"`
}

//CallDelayItem Call reschedule
type CallDelayItem struct {
	Date     *string `xml:"Date,attr"`
	DateNext *string `xml:"DateNext,attr"`
}

//ReturnOrder Return shipment
type ReturnOrder struct {
	ActNumber      *string      `xml:"ActNumber,attr"`
	Number         *string      `xml:"Number,attr"`
	DispatchNumber *int         `xml:"DispatchNumber,attr"`
	DeliveryDate   *string      `xml:"DeliveryDate,attr"`
	RecipientName  *string      `xml:"RecipientName,attr"`
	Status         *Status      `xml:"Status"`
	Reason         *Reason      `xml:"Reason"`
	DelayReason    *DelayReason `xml:"DelayReason"`
}
