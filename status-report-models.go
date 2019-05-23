package cdek

type StatusReportReq struct {
	StatusReport *StatusReportContentReq `xml:"StatusReport"`
}

type StatusReportRes struct {
	StatusReport *StatusReportContentForRes `xml:"StatusReport"`
	Order        []*StatusReportOrderRes    `xml:"Order"`
	ReturnOrder  []*ReturnOrder             `xml:"ReturnOrder,omitempty"`
}

type StatusReportContentReq struct {
	Date                   *string                 `xml:"Date,attr"`
	Account                *string                 `xml:"Account,attr"`
	Secure                 *string                 `xml:"Secure,attr"`
	ShowHistory            *bool                   `xml:"ShowHistory,attr"`
	ShowReturnOrder        *bool                   `xml:"ShowReturnOrder,attr"`
	ShowReturnOrderHistory *bool                   `xml:"ShowReturnOrderHistory,attr"`
	ChangePeriod           *ChangePeriod           `xml:"ChangePeriod"`
	Order                  []*StatusReportOrderReq `xml:"Order"`
}

type ChangePeriod struct {
	DateFirst *string `xml:"DateFirst,attr"`
	DateLast  *string `xml:"DateLast,attr"`
}

type StatusReportOrderReq struct {
	DispatchNumber *int    `xml:"DispatchNumber,attr"`
	Number         *string `xml:"Number,attr"`
	Date           *string `xml:"Date,attr"`
}

type StatusReportContentForRes struct {
	DateFirst *string `xml:"DateFirst,attr"`
	DateLast  *string `xml:"DateLast,attr"`
}

type StatusReportOrderRes struct {
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

type Status struct {
	Date        *string  `xml:"Date,attr"`
	Code        *string  `xml:"Code,attr"`
	Description *string  `xml:"Description,attr"`
	CityCode    *string  `xml:"CityCode,attr"`
	CityName    *string  `xml:"CityName,attr"`
	State       []*State `xml:"State"`
}

type State struct {
	Date        *string `xml:"Date,attr"`
	Code        *string `xml:"Code,attr"`
	Description *string `xml:"Description,attr"`
	CityCode    *string `xml:"CityCode,attr,omitempty"`
	CityName    *string `xml:"CityName,attr,omitempty"`
}

type Reason struct {
	Date        *string `xml:"Date,attr"`
	Code        *string `xml:"Code,attr"`
	Description *string `xml:"Description,attr"`
}

type DelayReason struct {
	Date        *string `xml:"Date,attr"`
	Code        *string `xml:"Code,attr"`
	Description *string `xml:"Description,attr"`
	State       []State `xml:"State"`
}

type Package struct {
	Number  *string `xml:"Number,attr"`
	BarCode *string `xml:"BarCode,attr"`
	Item    []*Item `xml:"Item"`
}

type Item struct {
	WareKey     *string `xml:"WareKey,attr"`
	Amount      *string `xml:"Amount,attr"`
	DelivAmount *string `xml:"DelivAmount,attr"`
}

type Attempt struct {
	ID                  *int    `xml:"ID,attr"`
	ScheduleCode        *int    `xml:"ScheduleCode,attr"`
	ScheduleDescription *string `xml:"ScheduleDescription,attr"`
}

type Call struct {
	CallGood  *CallGood  `xml:"CallGood"`
	CallFail  *CallFail  `xml:"CallFail"`
	CallDelay *CallDelay `xml:"CallDelay"`
}

type CallGood struct {
	Good []*CallGoodItem `xml:"Good"`
}

type CallGoodItem struct {
	Date      *string `xml:"Date,attr"`
	DateDeliv *string `xml:"DateDeliv,attr"`
}

type CallFail struct {
	Fail []*CallFailItem `xml:"Fail"`
}

type CallFailItem struct {
	Date              *string `xml:"Date,attr"`
	ReasonCode        *string `xml:"ReasonCode,attr"`
	ReasonDescription *string `xml:"ReasonDescription,attr"`
}

type CallDelay struct {
	Delay []*CallDelayItem `xml:"Delay"`
}

type CallDelayItem struct {
	Date     *string `xml:"Date,attr"`
	DateNext *string `xml:"DateNext,attr"`
}

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
