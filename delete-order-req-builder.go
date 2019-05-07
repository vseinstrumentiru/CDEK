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
func NewDeleteOrder() *DeleteOrder {
	return new(DeleteOrder)
}

//SetNumber Client's shipment number. Order identifier in the IS of the CDEK client.
func (o *DeleteOrder) SetNumber(number string) *DeleteOrder {
	o.Number = &number

	return o
}

//SetDispatchNumber CDEK order number
func (o *DeleteOrder) SetDispatchNumber(dispatchNumber int) *DeleteOrder {
	o.DispatchNumber = &dispatchNumber

	return o
}
