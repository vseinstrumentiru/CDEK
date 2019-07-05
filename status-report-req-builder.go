package cdek

import "time"

//NewStatusReportReq Order Status Report builder
func NewStatusReportReq() *StatusReportReq {
	statusReportReq := new(StatusReportReq)
	statusReportReq.StatusReport = new(StatusReportContentReq)

	return statusReportReq
}

func (req *StatusReportReq) setAuth(auth Auth) *StatusReportReq {
	req.StatusReport.Account = &auth.Account

	date, sec := auth.EncodedSecure()
	req.StatusReport.Date = &date
	req.StatusReport.Secure = &sec

	return req
}

//SetShowHistory The attribute indicating that the order history must be loaded (1 – yes, 0 – no)
func (req *StatusReportReq) SetShowHistory(showHistory bool) *StatusReportReq {
	req.StatusReport.ShowHistory = &showHistory

	return req
}

//SetShowReturnOrder The attribute indicating that the list of return orders must be loaded (1 – yes, 0 – no)
func (req *StatusReportReq) SetShowReturnOrder(showReturnOrder bool) *StatusReportReq {
	req.StatusReport.ShowReturnOrder = &showReturnOrder

	return req
}

//SetShowReturnOrderHistory The attribute indicating that the history of return orders must be loaded (1 – yes, 0 – no)
func (req *StatusReportReq) SetShowReturnOrderHistory(showReturnOrderHistory bool) *StatusReportReq {
	req.StatusReport.ShowReturnOrderHistory = &showReturnOrderHistory

	return req
}

//SetChangePeriod The period during which the order status has changed.
func (req *StatusReportReq) SetChangePeriod(changePeriod ChangePeriod) *StatusReportReq {
	req.StatusReport.ChangePeriod = &changePeriod

	return req
}

//AddOrder Add Shipment (order)
func (req *StatusReportReq) AddOrder(order StatusReportOrderReq) *StatusReportReq {
	req.StatusReport.Order = append(req.StatusReport.Order, &order)

	return req
}

//NewChangePeriod ChangePeriod builder
func NewChangePeriod() *ChangePeriod {
	return new(ChangePeriod)
}

//SetDateFirst Start date of requested period
func (changePeriod *ChangePeriod) SetDateFirst(date time.Time) *ChangePeriod {
	dateFormatted := date.Format("2006-01-02")
	changePeriod.DateFirst = &dateFormatted

	return changePeriod
}

//SetDateLast End date of requested period
func (changePeriod *ChangePeriod) SetDateLast(date time.Time) *ChangePeriod {
	dateFormatted := date.Format("2006-01-02")
	changePeriod.DateLast = &dateFormatted

	return changePeriod
}

//NewStatusReportOrderReq StatusReportOrderReq builder
func NewStatusReportOrderReq() *StatusReportOrderReq {
	return new(StatusReportOrderReq)
}

//SetDispatchNumber CDEK shipment number (assigned when orders are imported). Order identifier in the CDEK IS.
func (req *StatusReportOrderReq) SetDispatchNumber(dispatchNumber int) *StatusReportOrderReq {
	req.DispatchNumber = &dispatchNumber

	return req
}

//SetNumber Client's shipment number. Order identifier in the IS of the CDEK client.
func (req *StatusReportOrderReq) SetNumber(number string) *StatusReportOrderReq {
	req.Number = &number

	return req
}

//SetDate Date of an acceptance certificate, based on which the order has been transferred
func (req *StatusReportOrderReq) SetDate(date time.Time) *StatusReportOrderReq {
	dateFormatted := date.Format("2006-01-02")
	req.Date = &dateFormatted

	return req
}
