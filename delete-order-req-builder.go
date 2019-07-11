package cdek

//NewDeleteOrderReq DeleteOrderReq constructor
func NewDeleteOrderReq(number string, orderCount int, order DeleteOrder) *DeleteOrderReq {
	return &DeleteOrderReq{
		Number:     &number,
		OrderCount: &orderCount,
		Order:      &order,
	}
}

//NewDeleteOrder DeleteOrder constructor
func NewDeleteOrder(number string, dispatchNumber int) *DeleteOrder {
	return &DeleteOrder{
		Number:         &number,
		DispatchNumber: &dispatchNumber,
	}
}
