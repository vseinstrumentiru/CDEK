package cdek

import "time"

type StatusReportReqBuilder interface {
	SetAuth(auth Auth) *StatusReportReq
	SetShowHistory(showHistory bool) *StatusReportReq
	SetShowReturnOrder(showReturnOrder bool) *StatusReportReq
	SetShowReturnOrderHistory(showReturnOrderHistory bool) *StatusReportReq
	SetChangePeriod(changePeriod ChangePeriod) *StatusReportReq
	AddOrder(order StatusReportOrderReq) *StatusReportReq
}

func NewStatusReportReq() *StatusReportReq {
	statusReportReq := new(StatusReportReq)
	statusReportReq.StatusReport = new(StatusReportContentReq)

	return statusReportReq
}

func (req *StatusReportReq) SetAuth(auth Auth) *StatusReportReq {
	date, encodedSecure := auth.EncodedSecure()
	account := auth.Account
	req.StatusReport.Account = &account
	req.StatusReport.Secure = &encodedSecure
	req.StatusReport.Date = &date

	return req
}

func (req *StatusReportReq) SetShowHistory(showHistory bool) *StatusReportReq {
	req.StatusReport.ShowHistory = &showHistory

	return req
}

func (req *StatusReportReq) SetShowReturnOrder(showReturnOrder bool) *StatusReportReq {
	req.StatusReport.ShowReturnOrder = &showReturnOrder

	return req
}

func (req *StatusReportReq) SetShowReturnOrderHistory(showReturnOrderHistory bool) *StatusReportReq {
	req.StatusReport.ShowReturnOrderHistory = &showReturnOrderHistory

	return req
}

func (req *StatusReportReq) SetChangePeriod(changePeriod ChangePeriod) *StatusReportReq {
	req.StatusReport.ChangePeriod = &changePeriod

	return req
}

func (req *StatusReportReq) AddOrder(order StatusReportOrderReq) *StatusReportReq {
	req.StatusReport.Order = append(req.StatusReport.Order, &order)

	return req
}

func NewChangePeriod() *ChangePeriod {
	return new(ChangePeriod)
}

type ChangePeriodBuilder interface {
	SetDateFirst(date time.Time) *ChangePeriod
	SetDateLast(date time.Time) *ChangePeriod
}

func (changePeriod *ChangePeriod) SetDateFirst(date time.Time) *ChangePeriod {
	dateFormatted := date.Format("2006-01-02")
	changePeriod.DateFirst = &dateFormatted

	return changePeriod
}
func (changePeriod *ChangePeriod) SetDateLast(date time.Time) *ChangePeriod {
	dateFormatted := date.Format("2006-01-02")
	changePeriod.DateLast = &dateFormatted

	return changePeriod
}

func NewStatusReportOrderReq() *StatusReportOrderReq {
	return new(StatusReportOrderReq)
}

type StatusReportOrderReqBuilder interface {
	SetDispatchNumber(dispatchNumber int) *StatusReportOrderReq
	SetNumber(number string) *StatusReportOrderReq
	SetDate(date time.Time) *StatusReportOrderReq
}

func (req *StatusReportOrderReq) SetDispatchNumber(dispatchNumber int) *StatusReportOrderReq {
	req.DispatchNumber = &dispatchNumber

	return req
}

func (req *StatusReportOrderReq) SetNumber(number string) *StatusReportOrderReq {
	req.Number = &number

	return req
}

func (req *StatusReportOrderReq) SetDate(date time.Time) *StatusReportOrderReq {
	dateFormatted := date.Format("2006-01-02")
	req.Date = &dateFormatted

	return req
}
