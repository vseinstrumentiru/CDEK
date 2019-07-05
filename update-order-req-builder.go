package cdek

//NewUpdateOrderReq UpdateOrderReq builder
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

//NewUpdateOrder UpdateOrder builder
func NewUpdateOrder() *UpdateOrder {
	return new(UpdateOrder)
}

//SetNumber Client shipment number (unique for orders of a particular client). Order identifier in the Client's IS.
func (updateOrder *UpdateOrder) SetNumber(number string) *UpdateOrder {
	updateOrder.Number = &number

	return updateOrder
}

//SetDispatchNumber CDEK order number
func (updateOrder *UpdateOrder) SetDispatchNumber(dispatchNumber int) *UpdateOrder {
	updateOrder.DispatchNumber = &dispatchNumber

	return updateOrder
}

//SetDeliveryRecipientCost Additional delivery charge collected by the online store from the receiver
// (in the specified currency)
func (updateOrder *UpdateOrder) SetDeliveryRecipientCost(deliveryRecipientCost float64) *UpdateOrder {
	updateOrder.DeliveryRecipientCost = &deliveryRecipientCost

	return updateOrder
}

//SetDeliveryRecipientVATRate VAT rate included in the extra delivery charge
func (updateOrder *UpdateOrder) SetDeliveryRecipientVATRate(deliveryRecipientVATRate string) *UpdateOrder {
	updateOrder.DeliveryRecipientVATRate = &deliveryRecipientVATRate

	return updateOrder
}

//SetDeliveryRecipientVATSum VAT amount included in the extra. delivery charge
func (updateOrder *UpdateOrder) SetDeliveryRecipientVATSum(deliveryRecipientVATSum float64) *UpdateOrder {
	updateOrder.DeliveryRecipientVATSum = &deliveryRecipientVATSum

	return updateOrder
}

//SetRecipientName Receiver (full name). At least 3 characters.
func (updateOrder *UpdateOrder) SetRecipientName(recipientName string) *UpdateOrder {
	updateOrder.RecipientName = &recipientName

	return updateOrder
}

//SetPhone Receiver's phone
func (updateOrder *UpdateOrder) SetPhone(phone string) *UpdateOrder {
	updateOrder.Phone = &phone

	return updateOrder
}

//SetRecipientINN TIN of the receiver. Only for international orders.
func (updateOrder *UpdateOrder) SetRecipientINN(recipientINN string) *UpdateOrder {
	updateOrder.RecipientINN = &recipientINN

	return updateOrder
}

//SetDateInvoice Invoice date. Only for international orders.
func (updateOrder *UpdateOrder) SetDateInvoice(dateInvoice string) *UpdateOrder {
	updateOrder.DateInvoice = &dateInvoice

	return updateOrder
}

//SetPassport Details of the receiver’s passport.
func (updateOrder *UpdateOrder) SetPassport(passport Passport) *UpdateOrder {
	updateOrder.Passport = &passport

	return updateOrder
}

//SetAddress Delivery address.
func (updateOrder *UpdateOrder) SetAddress(address Address) *UpdateOrder {
	updateOrder.Address = &address

	return updateOrder
}

//SetDeliveryRecipientCostAdv Additional charge for delivery
func (updateOrder *UpdateOrder) SetDeliveryRecipientCostAdv(v DeliveryRecipientCostAdv) *UpdateOrder {
	updateOrder.DeliveryRecipientCostAdv = &v

	return updateOrder
}

//SetPackage Package
func (updateOrder *UpdateOrder) SetPackage(pack OrderPackage) *UpdateOrder {
	updateOrder.Package = &pack

	return updateOrder
}
