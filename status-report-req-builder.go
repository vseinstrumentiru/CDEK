package cdek

import "time"

//NewStatusReportReq Order Status Report builder
func NewStatusReportReq() *StatusReport {
	return new(StatusReport)
}

//SetShowHistory The attribute indicating that the order history must be loaded (1 – yes, 0 – no)
func (req *StatusReport) SetShowHistory(showHistory int) *StatusReport {
	req.ShowHistory = &showHistory

	return req
}

//SetShowReturnOrder The attribute indicating that the list of return orders must be loaded (1 – yes, 0 – no)
func (req *StatusReport) SetShowReturnOrder(showReturnOrder bool) *StatusReport {
	req.ShowReturnOrder = &showReturnOrder

	return req
}

//SetShowReturnOrderHistory The attribute indicating that the history of return orders must be loaded (1 – yes, 0 – no)
func (req *StatusReport) SetShowReturnOrderHistory(showReturnOrderHistory bool) *StatusReport {
	req.ShowReturnOrderHistory = &showReturnOrderHistory

	return req
}

//SetChangePeriod The period during which the order status has changed.
func (req *StatusReport) SetChangePeriod(changePeriod ChangePeriod) *StatusReport {
	req.ChangePeriod = &changePeriod

	return req
}

//AddOrder Add Shipment (order)
func (req *StatusReport) AddOrder(order StatusReportOrderReq) *StatusReport {
	req.Order = append(req.Order, &order)

	return req
}

//NewChangePeriod ChangePeriod builder
// dateFirst: start date of requested period
func NewChangePeriod(dateFirst time.Time) *ChangePeriod {
	dateFirstFormatted := dateFirst.Format("2006-01-02")

	return &ChangePeriod{
		DateFirst: &dateFirstFormatted,
	}
}

//SetDateLast End date of requested period
func (changePeriod *ChangePeriod) SetDateLast(date time.Time) *ChangePeriod {
	dateFormatted := date.Format("2006-01-02")
	changePeriod.DateLast = &dateFormatted

	return changePeriod
}

//NewStatusReportByCDEKNumberReq StatusReportOrderReq builder by CDEK order number
func NewStatusReportByCDEKIdentifierReq(dispatchNumber int) *StatusReportOrderReq {
	return &StatusReportOrderReq{
		DispatchNumber: &dispatchNumber,
	}
}

//NewStatusReportByClientNumberReq StatusReportOrderReq builder by client order number
func NewStatusReportByClientIdentifierReq(number string, date time.Time) *StatusReportOrderReq {
	dateFormatted := date.Format("2006-01-02")

	return &StatusReportOrderReq{
		Number: &number,
		Date:   &dateFormatted,
	}
}
