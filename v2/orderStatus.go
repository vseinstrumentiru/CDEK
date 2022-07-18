package v2

import (
	"context"
	"fmt"
	"net/http"
)

type OrderStatusFailedCall struct {
	// DateTime Дата и время создания недозвона	datetime
	DateTime string `json:"date_time"`
	// ReasonCode Причина недозвона (подробнее см. приложение 5)
	ReasonCode int `json:"reason_code"`
}

type OrderStatusRescheduledCall struct {
	// DateTime Дата и время создания переноса прозвона
	DateTime string `json:"date_time"`
	// DateNext Дата, на которую согласован повторный прозвон
	DateNext string `json:"date_next"`
	// TimeNext Время, на которое согласован повторный прозвон
	TimeNext string `json:"time_next"`
	// Comment Комментарий к переносу прозвона
	Comment string `json:"comment,omitempty"`
}

type OrderStatusCall struct {
	// FailedCalls Информация о неуспешных прозвонах (недозвонах)
	FailedCalls OrderStatusFailedCall `json:"failed_calls,omitempty"`
	// RescheduledCalls Информация о переносах прозвонов
	RescheduledCalls OrderStatusRescheduledCall `json:"rescheduled_calls,omitempty"`
}

type OrderStatusDeliveryProblem struct {
	// Code Код проблемы (подробнее см. приложение 4) https://api-docs.cdek.ru/29923975.html
	Code string `json:"code,omitempty"`
	// CreateDate Дата создания проблемы
	CreateDate string `json:"create_date,omitempty"`
}

type OrderStatusPaymentInfo struct {
	// Type Тип оплаты: CARD - картой, CASH - наличными
	Type string `json:"type"`
	// Sum Сумма в валюте страны получателя
	Sum float64 `json:"sum"`
	// DeliverySum Стоимость услуги доставки (по тарифу)
	DeliverySum float64 `json:"delivery_sum"`
	// TotalSum Итоговая стоимость заказа
	TotalSum float64 `json:"total_sum"`
}

type OrderStatusDeliveryDetail struct {
	// Date Дата доставки
	Date string `json:"date"`
	// RecipientName получатель при доставке
	RecipientName string `json:"recipient_name"`
	// PaymentSum Сумма наложенного платежа, которую взяли с получателя, в валюте страны получателя с учетом частичной доставки
	PaymentSum float64 `json:"payment_sum,omitempty"`
	// PaymentInfo Тип оплаты наложенного платежа получателем
	PaymentInfo []OrderStatusPaymentInfo `json:"payment_info,omitempty"`
	// DeliverySum Стоимость услуги доставки (по тарифу)
	DeliverySum float64 `json:"delivery_sum"`
	TotalSum    float64 `json:"total_sum"`
}

type OrderStatusInfo struct {
	// Code Код статуса (подробнее см. приложение 1)
	Code string `json:"code"`
	// Name Название статуса
	Name string `json:"name"`
	// DateTime Дата и время установки статуса (формат yyyy-MM-dd'T'HH:mm:ssZ)
	DateTime string `json:"date_time"`
	// ReasonCode Дополнительный код статуса (подробнее см. приложение 2)
	ReasonCode string `json:"reason_code,omitempty"`
	// City Наименование места возникновения статуса
	City string `json:"city"`
}

type OrderStatusEntity struct {
	// Uuid Идентификатор заказа в ИС СДЭК
	Uuid string `json:"uuid"`
	// IsReturn Признак возвратного заказа: true - возвратный, false - прямой
	IsReturn bool `json:"is_return"`
	// IsReverse Признак реверсного заказа: true - реверсный, false - не реверсный
	IsReverse bool `json:"is_reverse"`
	// Type Тип заказа: 1 - "интернет-магазин" (только для договора типа "Договор с ИМ"), 2 - "доставка" (для любого договора)
	Type int `json:"type"`
	// CdekNumber Номер заказа СДЭК
	CdekNumber string `json:"cdek_number,omitempty"`
	// Number Номер заказа в ИС Клиента. При запросе информации по данному полю возможны варианты:
	// - если не передан, будет присвоен номер заказа в ИС СДЭК - uuid;
	// - если найдено больше 1, то выбирается созданный с самой последней датой.
	// Может содержать только цифры, буквы латинского алфавита или спецсимволы (формат ASCII)
	Number string `json:"number,omitempty"`
	// DeliveryMode Истинный режим заказа:
	// 1 - дверь-дверь
	// 2 - дверь-склад
	// 3 - склад-дверь
	// 4 - склад-склад
	// 6 - дверь-постамат
	// 7 - склад-постамат
	DeliveryMode string `json:"delivery_mode"`
	//// TariffCode Код тарифа
	//TariffCode int `json:"tariff_code"`
	// Comment Комментарий к заказу
	Comment string `json:"comment,omitempty"`
	// DeveloperKey Ключ разработчика
	DeveloperKey string `json:"developer_key,omitempty"`
	// ShipmentPoint Код ПВЗ СДЭК, на который будет производиться самостоятельный привоз клиентом
	ShipmentPoint string `json:"shipment_point,omitempty"`
	// DeliveryPoint Код офиса СДЭК (ПВЗ/постамат), на который будет доставлена посылка
	DeliveryPoint string `json:"delivery_point,omitempty"`
	// DateInvoice Дата инвойса. Только для международных заказов. date (yyyy-MM-dd)
	DateInvoice string `json:"date_invoice,omitempty"`
	// ShipperName Грузоотправитель. Только для международных заказов
	ShipperName string `json:"shipper_name,omitempty"`
	// ShipperAddress Адрес грузоотправителя. Только для международных заказов
	ShipperAddress string `json:"shipper_address,omitempty"`
	// DeliveryRecipientCost Доп. сбор за доставку, которую ИМ берет с получателя.
	DeliveryRecipientCost Payment `json:"delivery_recipient_cost,omitempty"`
	// DeliveryRecipientCostAdv Доп. сбор за доставку (которую ИМ берет с получателя), в зависимости от суммы заказа
	DeliveryRecipientCostAdv []Cost `json:"delivery_recipient_cost_adv,omitempty"`
	// Sender Отправитель
	Sender RecipientSender `json:"sender"`
	// Seller Реквизиты истинного продавца
	Seller Seller `json:"seller,omitempty"`
	// Recipient Получатель
	Recipient RecipientSender `json:"recipient,omitempty"`
	// FromLocation Адрес отправления. Не может использоваться одновременно с shipment_point
	FromLocation Location `json:"from_location"`
	// ToLocation Адрес получения. Не может использоваться одновременно с delivery_point
	ToLocation Location `json:"to_location"`
	// ItemsCostCurrency TODO
	ItemsCostCurrency string `json:"items_cost_currency"`
	// RecipientCurrency TODO
	RecipientCurrency string `json:"recipient_currency"`
	// Services Дополнительные услуги
	Services []Service `json:"services,omitempty"`
	// Packages Список информации по местам (упаковкам)
	Packages []Package `json:"packages"`
	// DeliveryProblem Проблемы доставки, с которыми столкнулся курьер при доставке заказа "до двери"
	DeliveryProblem []OrderStatusDeliveryProblem `json:"delivery_problem,omitempty"`
	// DeliveryDetail Информация о вручении
	DeliveryDetail OrderStatusDeliveryDetail `json:"delivery_detail,omitempty"`
	// TransactedPayment Признак того, что по заказу была получена информация о переводе наложенного платежа интернет-магазину
	TransactedPayment bool `json:"transacted_payment,omitempty"`
	// Statuses Список статусов по заказу, отсортированных по дате и времени
	Statuses []OrderStatusInfo `json:"statuses"`
	// Calls Информация о прозвонах получателя
	Calls []OrderStatusCall `json:"calls,omitempty"`
	// @todo ticket SD-735298 - this is not documented but exists in example response https://api-docs.cdek.ru/29923975.html
	DeliveryDate   string `json:"delivery_date,omitempty"`
	ShopSellerName string `json:"shop_seller_name,omitempty"`
}

func (c *clientImpl) OrderStatus(ctx context.Context, uuid string) (*Response, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.buildUri(fmt.Sprintf("/v2/orders/%s", uuid), nil),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	return jsonReq[Response](req)
}
