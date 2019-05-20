package cdek

import "time"

func NewDeliveryReqContent() *DeliveryReqContent {
	return new(DeliveryReqContent)
}

func (deliveryReqContent *DeliveryReqContent) SetNumber(number string) *DeliveryReqContent {
	deliveryReqContent.Number = &number

	return deliveryReqContent
}

func (deliveryReqContent *DeliveryReqContent) SetAuth(auth Auth) *DeliveryReqContent {
	deliveryReqContent.Account = &auth.Account
	*deliveryReqContent.Date, *deliveryReqContent.Secure = auth.EncodedSecure()

	return deliveryReqContent
}

func (deliveryReqContent *DeliveryReqContent) SetOrderCount(orderCount string) *DeliveryReqContent {
	deliveryReqContent.OrderCount = &orderCount

	return deliveryReqContent
}

func (deliveryReqContent *DeliveryReqContent) SetCurrency(currency string) *DeliveryReqContent {
	deliveryReqContent.Currency = &currency

	return deliveryReqContent
}

func (deliveryReqContent *DeliveryReqContent) SetOrder(order OrderReq) *DeliveryReqContent {
	deliveryReqContent.Order = &order

	return deliveryReqContent
}

func (deliveryReqContent *DeliveryReqContent) SetCallCourier(callCourier CallCourier) *DeliveryReqContent {
	deliveryReqContent.CallCourier = &callCourier

	return deliveryReqContent
}

func NewOrderReq() *OrderReq {
	return new(OrderReq)
}

func (orderReq *OrderReq) SetNumber(number string) *OrderReq {
	orderReq.Number = &number

	return orderReq
}

func (orderReq *OrderReq) SetSendCityCode(sendCityCode int) *OrderReq {
	orderReq.SendCityCode = &sendCityCode

	return orderReq
}

func (orderReq *OrderReq) SetRecCityCode(recCityCode int) *OrderReq {
	orderReq.RecCityCode = &recCityCode

	return orderReq
}

func (orderReq *OrderReq) SetSendCityPostCode(sendCityPostCode int) *OrderReq {
	orderReq.SendCityPostCode = &sendCityPostCode

	return orderReq
}

func (orderReq *OrderReq) SetRecCityPostCode(recCityPostCode int) *OrderReq {
	orderReq.RecCityPostCode = &recCityPostCode

	return orderReq
}

func (orderReq *OrderReq) SetSendCountryCode(sendCountryCode int) *OrderReq {
	orderReq.SendCountryCode = &sendCountryCode

	return orderReq
}

func (orderReq *OrderReq) SetRecCountryCode(recCountryCode int) *OrderReq {
	orderReq.RecCountryCode = &recCountryCode

	return orderReq
}

func (orderReq *OrderReq) SetSendCityName(sendCityName string) *OrderReq {
	orderReq.SendCityName = &sendCityName

	return orderReq
}

func (orderReq *OrderReq) SetRecCityName(recCityName string) *OrderReq {
	orderReq.RecCityName = &recCityName

	return orderReq
}

func (orderReq *OrderReq) SetRecipientINN(recipientINN string) *OrderReq {
	orderReq.RecipientINN = &recipientINN

	return orderReq
}

func (orderReq *OrderReq) SetDateInvoice(dateInvoice time.Time) *OrderReq {
	dateInvoiceFmt := dateInvoice.Format("2006-01-02")
	orderReq.DateInvoice = &dateInvoiceFmt

	return orderReq
}

func (orderReq *OrderReq) SetShipperName(shipperName string) *OrderReq {
	orderReq.ShipperName = &shipperName

	return orderReq
}

func (orderReq *OrderReq) SetShipperAddress(shipperAddress string) *OrderReq {
	orderReq.ShipperAddress = &shipperAddress

	return orderReq
}

func (orderReq *OrderReq) SetPassport(passport Passport) *OrderReq {
	orderReq.Passport = &passport

	return orderReq
}

func (orderReq *OrderReq) SetSender(sender Sender) *OrderReq {
	orderReq.Sender = &sender

	return orderReq
}

func (orderReq *OrderReq) SetRecipientName(recipientName string) *OrderReq {
	orderReq.RecipientName = &recipientName

	return orderReq
}

func (orderReq *OrderReq) SetRecipientEmail(recipientEmail string) *OrderReq {
	orderReq.RecipientEmail = &recipientEmail

	return orderReq
}

func (orderReq *OrderReq) SetPhone(phone string) *OrderReq {
	orderReq.Phone = &phone

	return orderReq
}

func (orderReq *OrderReq) SetTariffTypeCode(tariffTypeCode int) *OrderReq {
	orderReq.TariffTypeCode = &tariffTypeCode

	return orderReq
}

func (orderReq *OrderReq) SetDeliveryRecipientCost(deliveryRecipientCost float64) *OrderReq {
	orderReq.DeliveryRecipientCost = &deliveryRecipientCost

	return orderReq
}

func (orderReq *OrderReq) SetDeliveryRecipientVATRate(deliveryRecipientVATRate string) *OrderReq {
	orderReq.DeliveryRecipientVATRate = &deliveryRecipientVATRate

	return orderReq
}

func (orderReq *OrderReq) SetDeliveryRecipientVATSum(deliveryRecipientVATSum float64) *OrderReq {
	orderReq.DeliveryRecipientVATSum = &deliveryRecipientVATSum

	return orderReq
}

func (orderReq *OrderReq) SetRecipientCurrency(recipientCurrency string) *OrderReq {
	orderReq.RecipientCurrency = &recipientCurrency

	return orderReq
}

func (orderReq *OrderReq) SetItemsCurrency(itemsCurrency string) *OrderReq {
	orderReq.ItemsCurrency = &itemsCurrency

	return orderReq
}

func (orderReq *OrderReq) SetSeller(seller Seller) *OrderReq {
	orderReq.Seller = &seller

	return orderReq
}

func (orderReq *OrderReq) SetComment(comment string) *OrderReq {
	orderReq.Comment = &comment

	return orderReq
}

func (orderReq *OrderReq) SetAddress(address Address) *OrderReq {
	orderReq.Address = &address

	return orderReq
}

func (orderReq *OrderReq) AddPackage(pack OrderPackage) *OrderReq {
	orderReq.Package = append(orderReq.Package, &pack)

	return orderReq
}

func (orderReq *OrderReq) SetDeliveryRecipientCostAdv(deliveryRecipientCostAdv DeliveryRecipientCostAdv) *OrderReq {
	orderReq.DeliveryRecipientCostAdv = &deliveryRecipientCostAdv

	return orderReq
}

func (orderReq *OrderReq) SetAddService(addService AddService) *OrderReq {
	orderReq.AddService = &addService

	return orderReq
}

func (orderReq *OrderReq) SetSchedule(schedule Schedule) *OrderReq {
	orderReq.Schedule = &schedule

	return orderReq
}

func NewPassport() *Passport {
	return new(Passport)
}

func (passport *Passport) SetSeries(series string) *Passport {
	passport.Series = &series

	return passport
}

func (passport *Passport) SetNumber(number string) *Passport {
	passport.Number = &number

	return passport
}

func (passport *Passport) SetIssueDate(issueDate time.Time) *Passport {
	issueDateFmt := issueDate.Format("2006-01-02")

	passport.IssueDate = &issueDateFmt

	return passport
}

func (passport *Passport) SetIssuedBy(issuedBy string) *Passport {
	passport.IssuedBy = &issuedBy

	return passport
}

func (passport *Passport) SetDateBirth(dateBirth time.Time) *Passport {
	dateBirthFmt := dateBirth.Format("2006-01-02")
	passport.DateBirth = &dateBirthFmt

	return passport
}

func NewSender() *Sender {
	return new(Sender)
}

func (sender *Sender) SetCompany(company string) *Sender {
	sender.Company = &company

	return sender
}

func (sender *Sender) SetName(name string) *Sender {
	sender.Name = &name

	return sender
}

func (sender *Sender) SetAddress(address Address) *Sender {
	sender.Address = &address

	return sender
}

func NewAddress() *Address {
	return new(Address)
}

func (address *Address) SetStreet(street string) *Address {
	address.Street = &street

	return address
}

func (address *Address) SetHouse(house string) *Address {
	address.House = &house

	return address
}

func (address *Address) SetFlat(flat string) *Address {
	address.Flat = &flat

	return address
}

func (address *Address) SetPhone(phone string) *Address {
	address.Phone = &phone

	return address
}

func (address *Address) SetPvzCode(pvzCode string) *Address {
	address.PvzCode = &pvzCode

	return address
}

func NewSeller() *Seller {
	return new(Seller)
}

func (seller *Seller) SetAddress(address string) *Seller {
	seller.Address = &address

	return seller
}

func (seller *Seller) SetName(name string) *Seller {
	seller.Name = &name

	return seller
}

func (seller *Seller) SetINN(inn string) *Seller {
	seller.INN = &inn

	return seller
}

func (seller *Seller) SetPhone(phone string) *Seller {
	seller.Phone = &phone

	return seller
}

func (seller *Seller) SetOwnershipForm(ownershipForm int) *Seller {
	seller.OwnershipForm = &ownershipForm

	return seller
}

func NewOrderPackage() *OrderPackage {
	return new(OrderPackage)
}

func (orderPackage *OrderPackage) SetNumber(number string) *OrderPackage {
	orderPackage.Number = &number

	return orderPackage
}

func (orderPackage *OrderPackage) SetBarCode(barCode string) *OrderPackage {
	orderPackage.BarCode = &barCode

	return orderPackage
}

func (orderPackage *OrderPackage) SetWeight(weight int) *OrderPackage {
	orderPackage.Weight = &weight

	return orderPackage
}

func (orderPackage *OrderPackage) SetSizeA(sizeA int) *OrderPackage {
	orderPackage.SizeA = &sizeA

	return orderPackage
}

func (orderPackage *OrderPackage) SetSizeB(sizeB int) *OrderPackage {
	orderPackage.SizeB = &sizeB

	return orderPackage
}

func (orderPackage *OrderPackage) SetSizeC(sizeC int) *OrderPackage {
	orderPackage.SizeC = &sizeC

	return orderPackage
}

func (orderPackage *OrderPackage) AddItem(item OrderPackageItem) *OrderPackage {
	orderPackage.Item = append(orderPackage.Item, &item)

	return orderPackage
}

func NewOrderPackageItem() *OrderPackageItem {
	return new(OrderPackageItem)
}

func (item *OrderPackageItem) SetAmount(amount int) *OrderPackageItem {
	item.Amount = &amount

	return item
}

func (item *OrderPackageItem) SetWarekey(warekey string) *OrderPackageItem {
	item.Warekey = &warekey

	return item
}

func (item *OrderPackageItem) SetCost(cost float64) *OrderPackageItem {
	item.Cost = &cost

	return item
}

func (item *OrderPackageItem) SetPayment(payment float64) *OrderPackageItem {
	item.Payment = &payment

	return item
}

func (item *OrderPackageItem) SetPaymentVATRate(paymentVATRate string) *OrderPackageItem {
	item.PaymentVATRate = &paymentVATRate

	return item
}

func (item *OrderPackageItem) SetPaymentVATSum(paymentVATSum float64) *OrderPackageItem {
	item.PaymentVATSum = &paymentVATSum

	return item
}

func (item *OrderPackageItem) SetWeight(weight int) *OrderPackageItem {
	item.Weight = &weight

	return item
}

func (item *OrderPackageItem) SetComment(comment string) *OrderPackageItem {
	item.Comment = &comment

	return item
}

func (item *OrderPackageItem) SetWeightBrutto(weightBrutto string) *OrderPackageItem {
	item.WeightBrutto = &weightBrutto

	return item
}

func (item *OrderPackageItem) SetCommentEx(commentEx string) *OrderPackageItem {
	item.CommentEx = &commentEx

	return item
}

func (item *OrderPackageItem) SetLink(link string) *OrderPackageItem {
	item.Link = &link

	return item
}

func NewDeliveryRecipientCostAdv() *DeliveryRecipientCostAdv {
	return new(DeliveryRecipientCostAdv)
}

func (d *DeliveryRecipientCostAdv) SetThreshold(threshold int) *DeliveryRecipientCostAdv {
	d.Threshold = &threshold

	return d
}

func (d *DeliveryRecipientCostAdv) SetSum(sum float64) *DeliveryRecipientCostAdv {
	d.Sum = &sum

	return d
}

func (d *DeliveryRecipientCostAdv) SetVATRate(vatRate string) *DeliveryRecipientCostAdv {
	d.VATRate = &vatRate

	return d
}

func (d *DeliveryRecipientCostAdv) SetVATSum(vatSum float64) *DeliveryRecipientCostAdv {
	d.VATSum = &vatSum

	return d
}

func NewAddService() *AddService {
	return new(AddService)
}

func (addService *AddService) SetServiceCode(serviceCode int) *AddService {
	addService.ServiceCode = &serviceCode

	return addService
}

func (addService *AddService) SetCount(count int) *AddService {
	addService.Count = &count

	return addService
}

func NewSchedule() *Schedule {
	return new(Schedule)
}

func (schedule *Schedule) AddAttempt(attempt ScheduleAttempt) *Schedule {
	schedule.Attempt = append(schedule.Attempt, &attempt)

	return schedule
}

func NewScheduleAttempt() *ScheduleAttempt {
	return new(ScheduleAttempt)
}

func (scheduleAttempt *ScheduleAttempt) SetID(id string) *ScheduleAttempt {
	scheduleAttempt.ID = &id

	return scheduleAttempt
}

func (scheduleAttempt *ScheduleAttempt) SetDate(date time.Time) *ScheduleAttempt {
	dateFmt := date.Format("2006-01-02")
	scheduleAttempt.Date = &dateFmt

	return scheduleAttempt
}

func (scheduleAttempt *ScheduleAttempt) SetComment(comment string) *ScheduleAttempt {
	scheduleAttempt.Comment = &comment

	return scheduleAttempt
}

func (scheduleAttempt *ScheduleAttempt) SetTimeBeg(timeBeg string) *ScheduleAttempt {
	scheduleAttempt.TimeBeg = &timeBeg

	return scheduleAttempt
}

func (scheduleAttempt *ScheduleAttempt) SetTimeEnd(timeEnd string) *ScheduleAttempt {
	scheduleAttempt.TimeEnd = &timeEnd

	return scheduleAttempt
}

func (scheduleAttempt *ScheduleAttempt) SetAddress(address Address) *ScheduleAttempt {
	scheduleAttempt.Address = &address

	return scheduleAttempt
}

func NewCallCourier() *CallCourier {
	return new(CallCourier)
}

func (callCourier *CallCourier) SetCall(call CourierCall) *CallCourier {
	callCourier.Call = &call

	return callCourier
}

func NewCourierCall() *CourierCall {
	return new(CourierCall)
}

func (call *CourierCall) SetDate(date time.Time) *CourierCall {
	dateFmt := date.Format("2006-01-02")
	call.Date = &dateFmt

	return call
}

func (call *CourierCall) SetTimeBeg(timeBeg string) *CourierCall {
	call.TimeBeg = &timeBeg

	return call
}

func (call *CourierCall) SetTimeEnd(end string) *CourierCall {
	call.TimeEnd = &end

	return call
}

func (call *CourierCall) SetLunchBeg(lunchBeg string) *CourierCall {
	call.LunchBeg = &lunchBeg

	return call
}

func (call *CourierCall) SetLunchEnd(lunchEnd string) *CourierCall {
	call.LunchEnd = &lunchEnd

	return call
}

func (call *CourierCall) SetSendCityCode(sendCityCode int) *CourierCall {
	call.SendCityCode = &sendCityCode

	return call
}

func (call *CourierCall) SetSendCityPostCode(sendCityPostCode string) *CourierCall {
	call.SendCityPostCode = &sendCityPostCode

	return call
}

func (call *CourierCall) SetSendCountryCode(sendCountryCode string) *CourierCall {
	call.SendCountryCode = &sendCountryCode

	return call
}

func (call *CourierCall) SetSendCityName(sendCityName string) *CourierCall {
	call.SendCityName = &sendCityName

	return call
}

func (call *CourierCall) SetSendPhone(sendPhone string) *CourierCall {
	call.SendPhone = &sendPhone

	return call
}

func (call *CourierCall) SetSenderName(senderName string) *CourierCall {
	call.SenderName = &senderName

	return call
}

func (call *CourierCall) SetComment(comment string) *CourierCall {
	call.Comment = &comment

	return call
}

func (call *CourierCall) SetSendAddress(sendAddress Address) *CourierCall {
	call.SendAddress = &sendAddress

	return call
}
