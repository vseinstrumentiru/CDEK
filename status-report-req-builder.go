package cdek

import "time"

//NewStatusReportReq Order Status Report builder
func NewStatusReportReq() *StatusReport {
	return new(StatusReport)
}

//SetShowHistory The attribute indicating that the order history must be loaded (1 – yes, 0 – no)
func (req *StatusReport) SetShowHistory(showHistory bool) *StatusReport {
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

//NewStatusReportOrderReq StatusReportOrderReq builder
// dispatchNumber: CDEK shipment number (assigned when orders are imported). Order identifier in the CDEK IS.
// number: Client's shipment number. Order identifier in the IS of the CDEK Client.
// date: Date of an acceptance certificate, based on which the order has been transferred.
func NewStatusReportOrderReq(dispatchNumber int, number string, date time.Time) *StatusReportOrderReq {
	dateFormatted := date.Format("2006-01-02")

	return &StatusReportOrderReq{
		DispatchNumber: &dispatchNumber,
		Number:         &number,
		Date:           &dateFormatted,
	}
}
