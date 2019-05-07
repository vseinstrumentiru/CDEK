package cdek

import (
	"time"
)

//NewDeliveryRequest Order registration request builder
// number: ID number of the acceptance certificate/waybill,
// accompanying document attached upon the transfer of the cargo to CDEK, generated in the online store's system.
// Identifier of the list of cargoes in the IS of the CDEK Client. By default, you can use 1.
// orderCount: The total number of orders in a document, default value: 1.
// order: Shipment (order)
func NewDeliveryRequest(number string, orderCount int, order *OrderReq) *RegisterOrderReq {
	return &RegisterOrderReq{
		Number:     &number,
		OrderCount: &orderCount,
		Order:      order,
	}
}

//SetCurrency Currency identifier for prices, RUB is a default parameter. Only for international orders
func (registerOrderReq *RegisterOrderReq) SetCurrency(currency string) *RegisterOrderReq {
	registerOrderReq.Currency = &currency

	return registerOrderReq
}

//SetCallCourier Call courier
func (registerOrderReq *RegisterOrderReq) SetCallCourier(callCourier CallCourier) *RegisterOrderReq {
	registerOrderReq.CallCourier = &callCourier

	return registerOrderReq
}

//NewOrderReq Shipment (order) builder
// number: Client shipment number (unique for orders of a particular Client). Order identifier in the Client's IS
// recipientName: Receiver (full name). At least 3 characters.
// phone: Receiver's phone
// tariffTypeCode: Code of tariff type
func NewOrderReq(number string, recipientName string, phone string, tariffTypeCode int) *OrderReq {
	return &OrderReq{
		Number:         &number,
		RecipientName:  &recipientName,
		Phone:          &phone,
		TariffTypeCode: &tariffTypeCode,
	}
}

//SetSendCityCode Sender's city code from the CDEK base
func (orderReq *OrderReq) SetSendCityCode(sendCityCode int) *OrderReq {
	orderReq.SendCityCode = &sendCityCode

	return orderReq
}

//SetRecCityCode Receiver's city code from the CDEK base
func (orderReq *OrderReq) SetRecCityCode(recCityCode int) *OrderReq {
	orderReq.RecCityCode = &recCityCode

	return orderReq
}

//SetSendCityPostCode Postal code of the sender's city
func (orderReq *OrderReq) SetSendCityPostCode(sendCityPostCode int) *OrderReq {
	orderReq.SendCityPostCode = &sendCityPostCode

	return orderReq
}

//SetRecCityPostCode Postal code of the receiver's city
func (orderReq *OrderReq) SetRecCityPostCode(recCityPostCode int) *OrderReq {
	orderReq.RecCityPostCode = &recCityPostCode

	return orderReq
}

//SetSendCountryCode Sender's country code to identify a country by postal code. Default value: RU
func (orderReq *OrderReq) SetSendCountryCode(sendCountryCode int) *OrderReq {
	orderReq.SendCountryCode = &sendCountryCode

	return orderReq
}

//SetRecCountryCode Receiver's country code to identify a country by postal code. Default value: RU
func (orderReq *OrderReq) SetRecCountryCode(recCountryCode int) *OrderReq {
	orderReq.RecCountryCode = &recCountryCode

	return orderReq
}

//SetSendCityName Name of sender's city
func (orderReq *OrderReq) SetSendCityName(sendCityName string) *OrderReq {
	orderReq.SendCityName = &sendCityName

	return orderReq
}

//SetRecCityName Name of receiver's city
func (orderReq *OrderReq) SetRecCityName(recCityName string) *OrderReq {
	orderReq.RecCityName = &recCityName

	return orderReq
}

//SetRecipientINN TIN of the receiver. Only for international orders.
func (orderReq *OrderReq) SetRecipientINN(recipientINN string) *OrderReq {
	orderReq.RecipientINN = &recipientINN

	return orderReq
}

//SetDateInvoice Invoice date. Only for international orders.
func (orderReq *OrderReq) SetDateInvoice(dateInvoice time.Time) *OrderReq {
	dateInvoiceFmt := dateInvoice.Format("2006-01-02")
	orderReq.DateInvoice = &dateInvoiceFmt

	return orderReq
}

//SetShipperName Shipper. Used to print waybills. Only for international orders.
func (orderReq *OrderReq) SetShipperName(shipperName string) *OrderReq {
	orderReq.ShipperName = &shipperName

	return orderReq
}

//SetShipperAddress Shipper’s address. Used to print waybills. Only for international orders
func (orderReq *OrderReq) SetShipperAddress(shipperAddress string) *OrderReq {
	orderReq.ShipperAddress = &shipperAddress

	return orderReq
}

//SetPassport Details of the receiver’s passport. Used to print waybills. Only for international orders.
func (orderReq *OrderReq) SetPassport(passport Passport) *OrderReq {
	orderReq.Passport = &passport

	return orderReq
}

//SetSender Sender. Must be defined if it is different from the online store Client.
// If the online store is a sender, the Sender tag is not available.
func (orderReq *OrderReq) SetSender(sender Sender) *OrderReq {
	orderReq.Sender = &sender

	return orderReq
}

//SetRecipientEmail Receiver's email for sending order status notifications and contacting in case of failed calls
func (orderReq *OrderReq) SetRecipientEmail(recipientEmail string) *OrderReq {
	orderReq.RecipientEmail = &recipientEmail

	return orderReq
}

//SetDeliveryRecipientCost Extra delivery charge collected by the online store from the receiver
// (in the specified currency)
func (orderReq *OrderReq) SetDeliveryRecipientCost(deliveryRecipientCost float64) *OrderReq {
	orderReq.DeliveryRecipientCost = &deliveryRecipientCost

	return orderReq
}

//SetDeliveryRecipientVATRate VAT rate included in the extra delivery charge
func (orderReq *OrderReq) SetDeliveryRecipientVATRate(deliveryRecipientVATRate string) *OrderReq {
	orderReq.DeliveryRecipientVATRate = &deliveryRecipientVATRate

	return orderReq
}

//SetDeliveryRecipientVATSum VAT amount included in the extra delivery charge
func (orderReq *OrderReq) SetDeliveryRecipientVATSum(deliveryRecipientVATSum float64) *OrderReq {
	orderReq.DeliveryRecipientVATSum = &deliveryRecipientVATSum

	return orderReq
}

//SetRecipientCurrency Code of cash-on-delivery currency:
// extra delivery charge and payment for the goods to be collected from the receiver.
// The currency is considered to be a currency of the receiver's country
func (orderReq *OrderReq) SetRecipientCurrency(recipientCurrency string) *OrderReq {
	orderReq.RecipientCurrency = &recipientCurrency

	return orderReq
}

//SetItemsCurrency Code of declared value currency (all items in the order).
// Currency of settlements with the CDEK Client under contract.
func (orderReq *OrderReq) SetItemsCurrency(itemsCurrency string) *OrderReq {
	orderReq.ItemsCurrency = &itemsCurrency

	return orderReq
}

//SetSeller Requisites of the real seller
func (orderReq *OrderReq) SetSeller(seller Seller) *OrderReq {
	orderReq.Seller = &seller

	return orderReq
}

//SetComment Comments (special notes on the order)
func (orderReq *OrderReq) SetComment(comment string) *OrderReq {
	orderReq.Comment = &comment

	return orderReq
}

//SetAddress Depending on a delivery mode, it is necessary to specify either Street, House, Flat attributes
// (delivery to the receiver's address) or PvzCode (pickup)
func (orderReq *OrderReq) SetAddress(address Address) *OrderReq {
	orderReq.Address = &address

	return orderReq
}

//AddPackage Package (all packages are sent with different Package tags)
func (orderReq *OrderReq) AddPackage(pack OrderPackage) *OrderReq {
	orderReq.Package = append(orderReq.Package, &pack)

	return orderReq
}

//SetDeliveryRecipientCostAdv Additional charge for delivery (E-shop charges the receiver), depending on the order’s sum
func (orderReq *OrderReq) SetDeliveryRecipientCostAdv(deliveryRecipientCostAdv DeliveryRecipientCostAdv) *OrderReq {
	orderReq.DeliveryRecipientCostAdv = &deliveryRecipientCostAdv

	return orderReq
}

//SetAddService Additional services
func (orderReq *OrderReq) SetAddService(addService AddService) *OrderReq {
	orderReq.AddService = &addService

	return orderReq
}

//SetSchedule Schedule for delivery/pickup
func (orderReq *OrderReq) SetSchedule(schedule Schedule) *OrderReq {
	orderReq.Schedule = &schedule

	return orderReq
}

//NewPassport Passport builder
func NewPassport() *Passport {
	return new(Passport)
}

//SetSeries Series of the receiver’s passport.
func (passport *Passport) SetSeries(series string) *Passport {
	passport.Series = &series

	return passport
}

//SetNumber Number of the receiver’s passport.
func (passport *Passport) SetNumber(number string) *Passport {
	passport.Number = &number

	return passport
}

//SetIssueDate Date of issue of the receiver’s passport.
func (passport *Passport) SetIssueDate(issueDate time.Time) *Passport {
	issueDateFmt := issueDate.Format("2006-01-02")

	passport.IssueDate = &issueDateFmt

	return passport
}

//SetIssuedBy Issuing authority of the receiver’s passport.
func (passport *Passport) SetIssuedBy(issuedBy string) *Passport {
	passport.IssuedBy = &issuedBy

	return passport
}

//SetDateBirth The receiver’s birthdate
func (passport *Passport) SetDateBirth(dateBirth time.Time) *Passport {
	dateBirthFmt := dateBirth.Format("2006-01-02")
	passport.DateBirth = &dateBirthFmt

	return passport
}

//NewSender Sender builder
func NewSender() *Sender {
	return new(Sender)
}

//SetCompany Name of sender's company
func (sender *Sender) SetCompany(company string) *Sender {
	sender.Company = &company

	return sender
}

//SetName Sender's contact person
func (sender *Sender) SetName(name string) *Sender {
	sender.Name = &name

	return sender
}

//SetAddress Address of pickup builder
func (sender *Sender) SetAddress(address Address) *Sender {
	sender.Address = &address

	return sender
}

//AddPhone Sender's phone
func (sender *Sender) AddPhone(phone string) *Sender {
	sender.Phone = append(sender.Phone, &phone)

	return sender
}

//NewAddress Address of pickup builder
func NewAddress(street string, house string) *Address {
	return &Address{
		Street: &street,
		House:  &house,
	}
}

//SetFlat Flat/office
func (address *Address) SetFlat(flat string) *Address {
	address.Flat = &flat

	return address
}

//SetPhone Sender's phone
func (address *Address) SetPhone(phone string) *Address {
	address.Phone = &phone

	return address
}

//SetPvzCode Pickup code. The attribute is required only for orders with the delivery mode “to warehouse”,
// provided that no additional service “Delivery in the receiver's city” is ordered.
// If the specified pickup point is closed at the time of order creation, the order will be accepted for the nearest
// functioning pickup point. The receiver will be notified about change of the pickup point via SMS.
// If all pickup points in the city that can provide the selected service are closed,
// order registration will be impossible. The relevant error message will be sent.
func (address *Address) SetPvzCode(pvzCode string) *Address {
	address.PvzCode = &pvzCode

	return address
}

//NewSeller Requisites of the real seller builder
func NewSeller() *Seller {
	return new(Seller)
}

//SetAddress Real seller’s address. Used to print invoices to display the address of the true seller or trade name.
// Only for international orders.
func (seller *Seller) SetAddress(address string) *Seller {
	seller.Address = &address

	return seller
}

//SetName Real seller
func (seller *Seller) SetName(name string) *Seller {
	seller.Name = &name

	return seller
}

//SetINN ITN (Individual Taxpayer Number) of the real seller
func (seller *Seller) SetINN(inn string) *Seller {
	seller.INN = &inn

	return seller
}

//SetPhone Telephone of the real seller
func (seller *Seller) SetPhone(phone string) *Seller {
	seller.Phone = &phone

	return seller
}

//SetOwnershipForm Code of type ownership
func (seller *Seller) SetOwnershipForm(ownershipForm int) *Seller {
	seller.OwnershipForm = &ownershipForm

	return seller
}

//NewOrderPackage OrderPackage builder
// number: Package number (ordinal package number or order number can be used), unique for this order.
// Order identifier in the Client's IS.
// barCode: Package barcode, package identifier (if any);
// otherwise, transmit a value of the package number: Packege.Number).
// The parameter is used to handle the cargo at CDEK warehouses), unique for this order.
// Package identifier in the Client's IS.
// weight: Total weight (in grams)
func NewOrderPackage(number string, barCode string, weight int) *OrderPackage {
	return &OrderPackage{
		Number:  &number,
		BarCode: &barCode,
		Weight:  &weight,
	}
}

//SetSize Package dimensions
// length: Length (in centimeters)
// width: Width (in centimeters)
// height: Height (in centimeters)
func (orderPackage *OrderPackage) SetSize(length int, width int, height int) *OrderPackage {
	orderPackage.SizeA = &length
	orderPackage.SizeB = &width
	orderPackage.SizeC = &height

	return orderPackage
}

//AddItem Add OrderPackageItem to OrderPackageItems list
func (orderPackage *OrderPackage) AddItem(item OrderPackageItem) *OrderPackage {
	orderPackage.Item = append(orderPackage.Item, &item)

	return orderPackage
}

//NewOrderPackageItem OrderPackageItem builder
// amount: Quantity of goods of the same type (pcs). The maximum number is 999.
// wareKey: Identifier/item number of the goods (Unique for this Package).
// cost: Declared value of the goods (per item in the specified currency, value >=0).
// This value is used to calculate the amount of insurance.
// payment: Cash on delivery (per item in the specified currency, value >=0) — cash on delivery;
// in case of prepayment, value = 0.
// weight: Weight (per item, in grams)
// comment: Name of the goods (may contain description of the goods: size, color)
func NewOrderPackageItem(
	amount int,
	wareKey string,
	cost float64,
	payment float64,
	weight int,
	comment string,
) *OrderPackageItem {
	return &OrderPackageItem{
		Amount:  &amount,
		WareKey: &wareKey,
		Cost:    &cost,
		Payment: &payment,
		Weight:  &weight,
		Comment: &comment,
	}
}

//SetPaymentVATRate VAT rate included in the value of the goods
func (item *OrderPackageItem) SetPaymentVATRate(paymentVATRate string) *OrderPackageItem {
	item.PaymentVATRate = &paymentVATRate

	return item
}

//SetPaymentVATSum VAT amount included in the value (payment) of the goods.
func (item *OrderPackageItem) SetPaymentVATSum(paymentVATSum float64) *OrderPackageItem {
	item.PaymentVATSum = &paymentVATSum

	return item
}

//SetWeightBrutto Gross weight (for one unit of goods, in grams). Only for international orders.
func (item *OrderPackageItem) SetWeightBrutto(weightBrutto int) *OrderPackageItem {
	item.WeightBrutto = &weightBrutto

	return item
}

//SetCommentEx Name of the goods, in English (also can contain description of the goods, such as size and color).
// Only for international orders.
func (item *OrderPackageItem) SetCommentEx(commentEx string) *OrderPackageItem {
	item.CommentEx = &commentEx

	return item
}

//SetLink Link to the e-shop’s website with the goods’ description. Only for international orders.
func (item *OrderPackageItem) SetLink(link string) *OrderPackageItem {
	item.Link = &link

	return item
}

//NewDeliveryRecipientCostAdv DeliveryRecipientCostAdv builder
// threshold: Goods price threshold (valid by condition less than or equal to), in even monetary units
// sum: Additional charge for delivery with total amount that falls within the interval
func NewDeliveryRecipientCostAdv(threshold int, sum float64) *DeliveryRecipientCostAdv {
	return &DeliveryRecipientCostAdv{
		Threshold: &threshold,
		Sum:       &sum,
	}
}

//SetVATRate VAT rate included in the additional charge for delivery.
// If the value is unknown, then default value “No VAT” is applied
func (d *DeliveryRecipientCostAdv) SetVATRate(vatRate string) *DeliveryRecipientCostAdv {
	d.VATRate = &vatRate

	return d
}

//SetVATSum VAT sum included in the additional charge for delivery
func (d *DeliveryRecipientCostAdv) SetVATSum(vatSum float64) *DeliveryRecipientCostAdv {
	d.VATSum = &vatSum

	return d
}

//NewAddService Additional services	 builder
// serviceCode: Type of additional service
func NewAddService(serviceCode int) *AddService {
	return &AddService{
		ServiceCode: &serviceCode,
	}
}

//SetCount Number of packages. It is used and is mandatory only for the additional service "Package 1".
func (addService *AddService) SetCount(count int) *AddService {
	addService.Count = &count

	return addService
}

//NewSchedule Schedule for delivery/pickup builder
func NewSchedule() *Schedule {
	return new(Schedule)
}

//AddAttempt add Time of delivery
func (schedule *Schedule) AddAttempt(attempt ScheduleAttempt) *Schedule {
	schedule.Attempt = append(schedule.Attempt, &attempt)

	return schedule
}

//NewScheduleAttempt Time of delivery builder
// id: ID number of a schedule from the online store's database. You may use 1 as a default value
// date: Date of delivery as agreed by the receiver
func NewScheduleAttempt(id string, date time.Time) *ScheduleAttempt {
	dateFmt := date.Format("2006-01-02")

	return &ScheduleAttempt{
		ID:   &id,
		Date: &dateFmt,
	}
}

//SetComment Comment
func (scheduleAttempt *ScheduleAttempt) SetComment(comment string) *ScheduleAttempt {
	scheduleAttempt.Comment = &comment

	return scheduleAttempt
}

//SetTimeBeg Start of a delivery time range (receiver's time)
func (scheduleAttempt *ScheduleAttempt) SetTimeBeg(timeBeg string) *ScheduleAttempt {
	scheduleAttempt.TimeBeg = &timeBeg

	return scheduleAttempt
}

//SetTimeEnd End of a delivery time range (receiver's time)
func (scheduleAttempt *ScheduleAttempt) SetTimeEnd(timeEnd string) *ScheduleAttempt {
	scheduleAttempt.TimeEnd = &timeEnd

	return scheduleAttempt
}

//SetAddress New delivery address (if change is required).
// Depending on a delivery mode, Street or House attributes should be specified.
// Flat – delivery to the receiver's address, or PvzCode – pickup
func (scheduleAttempt *ScheduleAttempt) SetAddress(address Address) *ScheduleAttempt {
	scheduleAttempt.Address = &address

	return scheduleAttempt
}

//NewCallCourier Call courier builder
// call: Waiting for a courier
func NewCallCourier(call CourierCallReq) *CallCourier {
	return &CallCourier{
		Call: &call,
	}
}

//NewCourierCall NewCourierCall builder
// date: Date of courier waiting
// timeBeg: Start time of courier waiting
// timeEnd: End time of courier waiting
// sendPhone: Sender's contact phone
// senderName: Sender (full name)
// sendAddress: Sender's address
func NewCourierCall(
	date time.Time,
	timeBeg string,
	timeEnd string,
	sendPhone string,
	senderName string,
	sendAddress Address,
) *CourierCallReq {
	dateFmt := date.Format("2006-01-02")

	return &CourierCallReq{
		Date:        &dateFmt,
		TimeBeg:     &timeBeg,
		TimeEnd:     &timeEnd,
		SendPhone:   &sendPhone,
		SenderName:  &senderName,
		SendAddress: &sendAddress,
	}
}

//SetLunchBeg Start time of a lunch break, if it is within a time range [TimeBeg; TimeEnd]
func (call *CourierCallReq) SetLunchBeg(lunchBeg string) *CourierCallReq {
	call.LunchBeg = &lunchBeg

	return call
}

//SetLunchEnd End time of a lunch break, if it is within a time range [TimeBeg; TimeEnd]
func (call *CourierCallReq) SetLunchEnd(lunchEnd string) *CourierCallReq {
	call.LunchEnd = &lunchEnd

	return call
}

//SetSendCityCode Sender's city code from the CDEK base
func (call *CourierCallReq) SetSendCityCode(sendCityCode int) *CourierCallReq {
	call.SendCityCode = &sendCityCode

	return call
}

//SetSendCityPostCode Postal code of the sender's city
func (call *CourierCallReq) SetSendCityPostCode(sendCityPostCode string) *CourierCallReq {
	call.SendCityPostCode = &sendCityPostCode

	return call
}

//SetSendCountryCode Sender's country code to identify a country by postal code
func (call *CourierCallReq) SetSendCountryCode(sendCountryCode string) *CourierCallReq {
	call.SendCountryCode = &sendCountryCode

	return call
}

//SetSendCityName Name of sender's city
func (call *CourierCallReq) SetSendCityName(sendCityName string) *CourierCallReq {
	call.SendCityName = &sendCityName

	return call
}

//SetComment Comment for courier
func (call *CourierCallReq) SetComment(comment string) *CourierCallReq {
	call.Comment = &comment

	return call
}
