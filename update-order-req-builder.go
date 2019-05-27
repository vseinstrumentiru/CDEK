package cdek

func NewUpdateOrderReq(number string, orderCount string, order UpdateOrder) *UpdateOrderReq {
	return &UpdateOrderReq{
		Number:     &number,
		OrderCount: &orderCount,
		Order:      &order,
	}
}

func (updateOrderReq *UpdateOrderReq) setAuth(auth Auth) *UpdateOrderReq {
	updateOrderReq.Account = &auth.Account

	date, sec := auth.EncodedSecure()
	updateOrderReq.Date = &date
	updateOrderReq.Secure = &sec

	return updateOrderReq
}

func NewUpdateOrder() *UpdateOrder {
	return new(UpdateOrder)
}

func (updateOrder *UpdateOrder) SetNumber(number string) *UpdateOrder {
	updateOrder.Number = &number

	return updateOrder
}

func (updateOrder *UpdateOrder) SetDispatchNumber(dispatchNumber int) *UpdateOrder {
	updateOrder.DispatchNumber = &dispatchNumber

	return updateOrder
}

func (updateOrder *UpdateOrder) SetDeliveryRecipientCost(deliveryRecipientCost float64) *UpdateOrder {
	updateOrder.DeliveryRecipientCost = &deliveryRecipientCost

	return updateOrder
}

func (updateOrder *UpdateOrder) SetDeliveryRecipientVATRate(deliveryRecipientVATRate string) *UpdateOrder {
	updateOrder.DeliveryRecipientVATRate = &deliveryRecipientVATRate

	return updateOrder
}

func (updateOrder *UpdateOrder) SetDeliveryRecipientVATSum(deliveryRecipientVATSum float64) *UpdateOrder {
	updateOrder.DeliveryRecipientVATSum = &deliveryRecipientVATSum

	return updateOrder
}

func (updateOrder *UpdateOrder) SetRecipientName(recipientName string) *UpdateOrder {
	updateOrder.RecipientName = &recipientName

	return updateOrder
}

func (updateOrder *UpdateOrder) SetPhone(phone string) *UpdateOrder {
	updateOrder.Phone = &phone

	return updateOrder
}

func (updateOrder *UpdateOrder) SetRecipientINN(recipientINN string) *UpdateOrder {
	updateOrder.RecipientINN = &recipientINN

	return updateOrder
}

func (updateOrder *UpdateOrder) SetDateInvoice(dateInvoice string) *UpdateOrder {
	updateOrder.DateInvoice = &dateInvoice

	return updateOrder
}

func (updateOrder *UpdateOrder) SetPassport(passport Passport) *UpdateOrder {
	updateOrder.Passport = &passport

	return updateOrder
}

func (updateOrder *UpdateOrder) SetAddress(address Address) *UpdateOrder {
	updateOrder.Address = &address

	return updateOrder
}

func (updateOrder *UpdateOrder) SetDeliveryRecipientCostAdv(deliveryRecipientCostAdv DeliveryRecipientCostAdv) *UpdateOrder {
	updateOrder.DeliveryRecipientCostAdv = &deliveryRecipientCostAdv

	return updateOrder
}

func (updateOrder *UpdateOrder) SetPackage(pack OrderPackage) *UpdateOrder {
	updateOrder.Package = &pack

	return updateOrder
}
