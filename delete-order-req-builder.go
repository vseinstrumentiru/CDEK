package cdek

//NewDeleteOrderReq DeleteOrderReq constructor
func NewDeleteOrderReq(number string, orderCount string, order DeleteOrder) *DeleteOrderReq {
	return &DeleteOrderReq{
		Number:     &number,
		OrderCount: &orderCount,
		Order:      &order,
	}
}

func (deleteOrderReq *DeleteOrderReq) setAuth(auth Auth) *DeleteOrderReq {
	deleteOrderReq.Account = &auth.Account

	date, sec := auth.EncodedSecure()
	deleteOrderReq.Date = &date
	deleteOrderReq.Secure = &sec

	return deleteOrderReq
}

//NewDeleteOrder DeleteOrder constructor
func NewDeleteOrder(number string, dispatchNumber int) *DeleteOrder {
	return &DeleteOrder{
		Number:         &number,
		DispatchNumber: &dispatchNumber,
	}
}
