package v2

import (
	"context"
	"fmt"
	"net/http"
)

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
	// TariffCode Код тарифа
	TariffCode int `json:"tariff_code"`
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
	DeliveryRecipientCost struct {
		// Value Сумма дополнительного сбора
		Value int `json:"value"`
		// VatSum Сумма НДС
		VatSum int `json:"vat_sum,omitempty"`
		// VatRate Ставка НДС (значение - 0, 10, 20, null - нет НДС)
		VatRate int `json:"vat_rate,omitempty"`
	} `json:"delivery_recipient_cost,omitempty"`
	// DeliveryRecipientCostAdv Доп. сбор за доставку (которую ИМ берет с получателя), в зависимости от суммы заказа
	DeliveryRecipientCostAdv []struct {
		// Sum Доп. сбор за доставку товаров, общая стоимость которых попадает в интервал
		Sum int `json:"sum"`
		// Threshold Порог стоимости товара (действует по условию меньше или равно) в целых единицах валюты
		Threshold int `json:"threshold"`
		// VatSum Сумма НДС
		VatSum int `json:"vat_sum,omitempty"`
		// VatRate Ставка НДС (значение - 0, 10, 20, null - нет НДС)
		VatRate int `json:"vat_rate,omitempty"`
	} `json:"delivery_recipient_cost_adv,omitempty"`
	// Sender Отправитель
	Sender struct {
		// Company Название компании
		Company string `json:"company,omitempty"`
		// Name ФИО контактного лица
		Name string `json:"name"`
		// PassportRequirementsSatisfied Требования по паспортным данным удовлетворены (актуально для
		// международных заказов):
		// true - паспортные данные собраны или не требуются
		// false - паспортные данные требуются и не собраны
		PassportRequirementsSatisfied bool `json:"passport_requirements_satisfied,omitempty"`
		// PassportSeries Серия паспорта
		PassportSeries string `json:"passport_series,omitempty"`
		// PassportSeries Номер паспорта
		PassportNumber string `json:"passport_number,omitempty"`
		// PassportDateOfIssue Номер паспорта date (yyyy-MM-dd)
		PassportDateOfIssue string `json:"passport_date_of_issue,omitempty"`
		// PassportOrganization Орган выдачи паспорта
		PassportOrganization string `json:"passport_organization,omitempty"`
		// Tin ИНН Может содержать 10, либо 12 символов
		Tin string `json:"tin,omitempty"`
		// PassportDateOfBirth Дата рождения (yyyy-MM-dd)
		PassportDateOfBirth string `json:"passport_date_of_birth,omitempty"`
		// Phones Список телефонов, Не более 10 номеров
		Phones []OrderPhone `json:"phones,omitempty"`
	} `json:"sender"`
	// Seller Реквизиты истинного продавца
	Seller OrderSeller `json:"seller,omitempty"`
	// Recipient Получатель
	Recipient OrderSenderRecipient `json:"recipient,omitempty"`
	// FromLocation Адрес отправления. Не может использоваться одновременно с shipment_point
	FromLocation OrderLocation `json:"from_location"`
	// ToLocation Адрес получения. Не может использоваться одновременно с delivery_point
	ToLocation OrderLocation `json:"to_location"`
	// ItemsCostCurrency TODO
	ItemsCostCurrency string `json:"items_cost_currency"`
	// RecipientCurrency TODO
	RecipientCurrency string `json:"recipient_currency"`
	// Services Дополнительные услуги
	Services []OrderService `json:"services,omitempty"`
	// Packages Список информации по местам (упаковкам)
	Packages OrderPackage `json:"packages"`
	// DeliveryProblem Проблемы доставки, с которыми столкнулся курьер при доставке заказа "до двери"
	DeliveryProblem []struct {
		// Code Код проблемы (подробнее см. приложение 4) https://api-docs.cdek.ru/29923975.html
		Code string `json:"code,omitempty"`
		// CreateDate Дата создания проблемы
		CreateDate string `json:"create_date,omitempty"`
	} `json:"delivery_problem,omitempty"`
	// DeliveryDetail Информация о вручении
	DeliveryDetail struct {
		// Date Дата доставки
		Date string `json:"date"`
		// RecipientName получатель при доставке
		RecipientName string `json:"recipient_name"`
		// PaymentSum Сумма наложенного платежа, которую взяли с получателя, в валюте страны получателя с учетом частичной доставки
		PaymentSum float64 `json:"payment_sum,omitempty"`
		// PaymentInfo Тип оплаты наложенного платежа получателем
		PaymentInfo []struct {
			// Type Тип оплаты: CARD - картой, CASH - наличными
			Type string `json:"type"`
			// Sum Сумма в валюте страны получателя
			Sum float64 `json:"sum"`
			// DeliverySum Стоимость услуги доставки (по тарифу)
			DeliverySum float64 `json:"delivery_sum"`
			// TotalSum Итоговая стоимость заказа
			TotalSum float64 `json:"total_sum"`
		} `json:"payment_info,omitempty"`
		// DeliverySum Стоимость услуги доставки (по тарифу)
		DeliverySum float64 `json:"delivery_sum"`
		TotalSum    float64 `json:"total_sum"`
	} `json:"delivery_detail,omitempty"`
	// TransactedPayment Признак того, что по заказу была получена информация о переводе наложенного платежа интернет-магазину
	TransactedPayment bool `json:"transacted_payment,omitempty"`
	// Statuses Список статусов по заказу, отсортированных по дате и времени
	Statuses []struct {
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
	} `json:"statuses"`
	// Calls Информация о прозвонах получателя
	Calls []struct {
		// FailedCalls Информация о неуспешных прозвонах (недозвонах)
		FailedCalls struct {
			// DateTime Дата и время создания недозвона	datetime
			DateTime string `json:"date_time"`
			// ReasonCode Причина недозвона (подробнее см. приложение 5)
			ReasonCode int `json:"reason_code"`
		} `json:"failed_calls,omitempty"`
		// RescheduledCalls Информация о переносах прозвонов
		RescheduledCalls struct {
			// DateTime Дата и время создания переноса прозвона
			DateTime string `json:"date_time"`
			// DateNext Дата, на которую согласован повторный прозвон
			DateNext string `json:"date_next"`
			// TimeNext Время, на которое согласован повторный прозвон
			TimeNext string `json:"time_next"`
			// Comment Комментарий к переносу прозвона
			Comment string `json:"comment,omitempty"`
		} `json:"rescheduled_calls,omitempty"`
	} `json:"calls,omitempty"`

	// @todo ticket SD-735298 - this is not documented but exists in example response https://api-docs.cdek.ru/29923975.html
	DeliveryDate   string `json:"delivery_date,omitempty"`
	ShopSellerName string `json:"shop_seller_name,omitempty"`
}

type OrderStatusRequests struct {
	RequestUuid string `json:"request_uuid"`
	Type        string `json:"type"`
	DateTime    string `json:"date_time"`
	State       string `json:"state"`
}

type OrderStatusResponse struct {
	Entity OrderStatusEntity `json:"entity,omitempty"`
	// Requests Информация о запросе/запросах над заказом
	Requests []struct {
		// RequestUuid Идентификатор запроса в ИС СДЭК
		RequestUuid string `json:"request_uuid,omitempty"`
		// Type Тип запроса. Может принимать значения: CREATE, UPDATE, DELETE, AUTH, GET
		Type string `json:"type"`
		// State Текущее состояние запроса. Может принимать значения:
		// ACCEPTED - пройдена предварительная валидация и запрос принят
		// WAITING - запрос ожидает обработки (зависит от выполнения другого запроса)
		// SUCCESSFUL - запрос обработан успешно
		// INVALID - запрос обработался с ошибкой
		State string `json:"state"`
		// DateTime Дата и время установки текущего состояния запроса (формат yyyy-MM-dd'T'HH:mm:ssZ)
		DateTime string `json:"date_time"`
		// Errors Ошибки, возникшие в ходе выполнения запроса
		Errors []OrderRegisterError `json:"errors,omitempty"`
		// Warnings Предупреждения, возникшие в ходе выполнения запроса
		Warnings []OrderRegisterError `json:"warnings,omitempty"`
	} `json:"requests"`
	// RelatedEntities Связанные сущности (если в запросе был передан корректный print)
	RelatedEntities []struct {
		// Type Тип связанной сущности. Может принимать значения: waybill - квитанция к заказу, barcode - ШК места к заказу
		Type string `json:"type"`
		// Uuid Идентификатор сущности, связанной с заказом
		Uuid string `json:"uuid"`
		// Url Ссылка на скачивание печатной формы в статусе "Сформирован", только для type = waybill, barcode
		Url string `json:"url,omitempty"`
		// CdekNumber Номер заказа СДЭК. Может возвращаться для return_order, direct_order, reverse_order
		CdekNumber string `json:"cdek_number,omitempty"`
		// Date Дата доставки, согласованная с получателем. Только для типа delivery
		Date string `json:"date,omitempty"`
		// TimeFrom Время начала ожидания курьера (согласованное с получателем). Только для типа delivery
		TimeFrom string `json:"time_from,omitempty"`
		// Date Время окончания ожидания курьера (согласованное с получателем). Только для типа delivery
		TimeTo string `json:"time_to,omitempty"`
	} `json:"related_entities,omitempty"`
}

func (c *clientImpl) OrderStatus(ctx context.Context, uuid string) (*OrderStatusResponse, error) {
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

	return jsonReq[OrderStatusResponse](req)
}
