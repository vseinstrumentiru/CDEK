package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

type OrderDeliveryRecipientCost struct {
	// Value Сумма дополнительного сбора
	Value int `json:"value"`
	// VatSum Сумма НДС
	VatSum int `json:"vat_sum,omitempty"`
	// VatRate Ставка НДС (значение - 0, 10, 20, null - нет НДС)
	VatRate int `json:"vat_rate,omitempty"`
}

type DeliveryRecipientCostAdv struct {
	// Sum Доп. сбор за доставку товаров, общая стоимость которых попадает в интервал
	Sum int `json:"sum"`
	// Threshold Порог стоимости товара (действует по условию меньше или равно) в целых единицах валюты
	Threshold int `json:"threshold"`
	// VatSum Сумма НДС
	VatSum int `json:"vat_sum,omitempty"`
	// VatRate Ставка НДС (значение - 0, 10, 20, null - нет НДС)
	VatRate int `json:"vat_rate,omitempty"`
}

type OrderLocation struct {
	// Code Код населенного пункта СДЭК (метод "Список населенных пунктов")
	Code int `json:"code,omitempty"`
	// FiasGuid Уникальный идентификатор ФИАС UUID
	FiasGuid string `json:"fias_guid,omitempty"`
	// PostalCode Почтовый индекс
	PostalCode string `json:"postal_code,omitempty"`
	// Longitude Долгота
	Longitude float64 `json:"longitude,omitempty"`
	// Latitude Широта
	Latitude float64 `json:"latitude,omitempty"`
	// CountryCode
	CountryCode string `json:"country_code,omitempty"`
	// Region Название региона
	Region string `json:"region,omitempty"`
	// RegionCode Код региона СДЭК
	RegionCode int `json:"region_code,omitempty"`
	// SubRegion Название района региона
	SubRegion string `json:"sub_region,omitempty"`
	// City Название города
	City string `json:"city,omitempty"`
	// Address Строка адреса
	Address string `json:"address"`
}

type OrderRegisterRequest struct {
	// Type Тип заказа: 1 - "интернет-магазин" (только для договора типа "Договор с ИМ"), 2 - "доставка" (для любого договора)
	Type int `json:"type,omitempty"`
	// Number Только для заказов "интернет-магазин". Номер заказа в ИС Клиента (если не передан, будет присвоен номер заказа в ИС СДЭК - uuid)
	// Может содержать только цифры, буквы латинского алфавита или спецсимволы (формат ASCII)
	Number string `json:"number"`
	// Comment Комментарий к заказу
	Comment string `json:"comment"`
	// TariffCode Код тарифа (подробнее см. приложение 1)
	TariffCode int `json:"tariff_code"`
	// OrderDeliveryRecipientCost Доп. сбор за доставку, которую ИМ берет с получателя. Только для заказов "интернет-магазин".
	DeliveryRecipientCost OrderDeliveryRecipientCost `json:"delivery_recipient_cost"`
	// DeliveryRecipientCostAdv Доп. сбор за доставку (которую ИМ берет с получателя) в зависимости от суммы заказа
	// Только для заказов "интернет-магазин".  Возможно указать несколько порогов.
	DeliveryRecipientCostAdv DeliveryRecipientCostAdv `json:"delivery_recipient_cost_adv"`
	// FromLocation Адрес отправления. Не может использоваться одновременно с shipment_point
	FromLocation OrderLocation `json:"from_location"`
	// ToLocation Адрес получения. Не может использоваться одновременно с delivery_point
	ToLocation OrderLocation `json:"to_location"`
	// Packages Список информации по местам (упаковкам). Количество мест в заказе может быть от 1 до 255
	Packages []OrderPackage `json:"packages,omitempty"`
	// Recipient Получатель
	Recipient OrderSenderRecipient `json:"recipient"`
	// Sender Отправитель. Обязателен если:
	// нет, если заказ типа "интернет-магазин"
	// да, если заказ типа "доставка"
	Sender OrderSenderRecipient `json:"sender,omitempty"`
	// Services Дополнительные услуги
	Services []OrderService `json:"services,omitempty"`
	// Seller Реквизиты истинного продавца. Только для заказов "интернет-магазин"
	Seller OrderSeller `json:"seller,omitempty"`
}

type OrderSeller struct {
	// Name Наименование истинного продавца. Обязателен если заполнен inn
	Name string `json:"name,omitempty"`
	// INN ИНН истинного продавца. Может содержать 10, либо 12 символов
	INN string `json:"inn,omitempty"`
	// Phone Телефон истинного продавца. Обязателен если заполнен inn
	Phone string `json:"phone,omitempty"`
	// OwnershipForm Код формы собственности (подробнее см. приложение 2). Обязателен если заполнен inn
	OwnershipForm int `json:"ownership_form,omitempty"`
	// Address Адрес истинного продавца. Используется при печати инвойсов для отображения адреса настоящего
	// продавца товара, либо торгового названия. Только для международных заказов "интернет-магазин".
	// Обязателен если заказ - международный
	Address string `json:"address,omitempty"`
}

type OrderPhone struct {
	// Number Номер телефона. Должен передаваться в международном формате: код страны (для России +7) и сам номер (10 и более цифр)
	// Обязателен если: нет, если заказ типа "интернет-магазин". да, если заказ типа "доставка"
	Number string `json:"number,omitempty"`
	// Additional Дополнительная информация (добавочный номер)
	Additional string `json:"additional,omitempty"`
}

type OrderSenderRecipient struct {
	// Name нет, если заказ типа "интернет-магазин"; да, если заказ типа "доставка"
	Name string `json:"name,omitempty"`
	// Company Название компании. нет, если заказ типа "интернет-магазин"; да, если заказ типа "доставка"
	Company string `json:"company,omitempty"`
	// Email Эл. адрес. нет, если заказ типа "интернет-магазин"; да, если заказ типа "доставка"
	Email string `json:"email,omitempty"`
	// PassportSeries Серия паспорта
	PassportSeries string `json:"passport_series,omitempty"`
	// PassportNumber Номер паспорта
	PassportNumber string `json:"passport_number,omitempty"`
	// PassportDateOfIssue Дата выдачи паспорта
	PassportDateOfIssue string `json:"passport_date_of_issue,omitempty"`
	// PassportOrganization Орган выдачи паспорта
	PassportOrganization string `json:"passport_organization,omitempty"`
	// Tin ИНН Может содержать 10, либо 12 символов
	Tin string `json:"tin,omitempty"`
	// PassportDateOfBirth Дата рождения (yyyy-MM-dd)
	PassportDateOfBirth string `json:"passport_date_of_birth,omitempty"`
	// Phones Список телефонов, Не более 10 номеров
	Phones []OrderPhone `json:"phones,omitempty"`
}

type OrderPayment struct {
	// Value Сумма наложенного платежа (в случае предоплаты = 0)
	Value int `json:"value"`
	// VatSum Сумма НДС
	VatSum int `json:"vat_sum,omitempty"`
	// VatRate Ставка НДС (значение - 0, 10, 20, null - нет НДС)
	VatRate int `json:"vat_rate,omitempty"`
}

type OrderPackageItem struct {
	// Name Наименование товара (может также содержать описание товара: размер, цвет)
	Name string `json:"name"`
	// WareKey Идентификатор/артикул товара. Артикул товара может содержать только символы: [A-z А-я 0-9 ! @ " # № $ ; % ^ : & ? * () _ - + = ? < > , .{ } [ ] \ / , пробел]
	WareKey string `json:"ware_key"`
	// Marking Маркировка товара. Если для товара/вложения указана маркировка, Amount не может быть больше 1.
	// Для корректного отображения маркировки товара в чеке требуется передавать НЕ РАЗОБРАННЫЙ тип маркировки, который может выглядеть следующим образом:
	// 1) Код товара в формате GS1. Пример: 010468008549838921AAA0005255832GS91EE06GS92VTwGVc7wKCc2tqRncUZ1RU5LeUKSXjWbfNQOpQjKK+A
	// 2) Последовательность допустимых символов общей длиной в 29 символов. Пример: 00000046198488X?io+qCABm8wAYa
	// 3) Меховые изделия. Имеют собственный формат. Пример: RU-430302-AAA7582720
	Marking string `json:"marking,omitempty"`
	// Payment Оплата за товар при получении (за единицу товара в валюте страны получателя, значение >=0) — наложенный платеж, в случае предоплаты значение = 0
	Payment OrderPayment `json:"payment"`
	// Cost Объявленная стоимость товара (за единицу товара в валюте взаиморасчетов, значение >=0). С данного значения рассчитывается страховка
	Cost float64 `json:"cost"`
	// Amount Количество единиц товара (в штуках). Количество одного товара в заказе может быть от 1 до 999
	Amount int `json:"amount"`
	// NameI18N Наименование на иностранном языке. Только для международных заказов
	NameI18N string `json:"name_i18n,omitempty"`
	// Brand Бренд на иностранном языке. Только для международных заказов
	Brand string `json:"brand,omitempty"`
	// CountryCode Бренд на иностранном языке. Только для международных заказов
	CountryCode string `json:"country_code,omitempty"`
	// Weight Вес (за единицу товара, в граммах)
	Weight int `json:"weight"`
	// WeightGross Вес брутто. Только для международных заказов
	WeightGross int `json:"weight_gross,omitempty"`
	// Material Код материала (подробнее см. приложение 4). Только для международных заказов
	Material string `json:"material,omitempty"`
	// WifiGsm Содержит wifi/gsm. Только для международных заказов
	WifiGsm bool `json:"wifi_gsm,omitempty"`
	// Url Ссылка на сайт интернет-магазина с описанием товара. Только для международных заказов
	Url string `json:"url,omitempty"`
}

type OrderPackage struct {
	// Number Номер упаковки (можно использовать порядковый номер упаковки заказа или номер заказа), уникален в пределах заказа. Идентификатор заказа в ИС Клиента
	Number string `json:"number"`
	// Weight Общий вес (в граммах)
	Weight int `json:"weight"`
	// Comment Комментарий к упаковке. Обязательно и только для заказа типа "доставка"
	Comment string `json:"comment,omitempty"`
	// Height Габариты упаковки. Высота (в сантиметрах). Поле обязательно если:
	// если указаны остальные габариты
	// если заказ до постамата
	// если общий вес >=100 гр
	Height int `json:"height,omitempty"`
	// Length Габариты упаковки. Длина (в сантиметрах). Поле обязательно если:
	// если указаны остальные габариты
	// если заказ до постамата
	// если общий вес >=100 гр
	Length int `json:"length,omitempty"`
	// Width Габариты упаковки. Ширина (в сантиметрах). Поле обязательно если:
	// если указаны остальные габариты
	// если заказ до постамата
	// если общий вес >=100 гр
	Width int `json:"width,omitempty"`
	// Items Позиции товаров в упаковке. Только для заказов "интернет-магазин". Максимум 126 уникальных позиций в заказе. Общее количество товаров в заказе может быть от 1 до 10000
	Items []OrderPackageItem `json:"items,omitempty"`
}

type OrderService struct {
	// Code Тип дополнительной услуги (подробнее см. приложение 3)
	Code string `json:"code"`
	// Parameter Параметр дополнительной услуги:
	// количество для услуг
	// PACKAGE_1, COURIER_PACKAGE_A2, SECURE_PACKAGE_A2, SECURE_PACKAGE_A3, SECURE_PACKAGE_A4,
	// SECURE_PACKAGE_A5, CARTON_BOX_XS, CARTON_BOX_S, CARTON_BOX_M, CARTON_BOX_L, CARTON_BOX_500GR,
	// CARTON_BOX_1KG, CARTON_BOX_2KG, CARTON_BOX_3KG, CARTON_BOX_5KG, CARTON_BOX_10KG, CARTON_BOX_15KG,
	// CARTON_BOX_20KG, CARTON_BOX_30KG, CARTON_FILLER (для всех типов заказа)
	// объявленная стоимость заказа для услуги INSURANCE (только для заказов с типом "доставка")
	// длина для услуг BUBBLE_WRAP, WASTE_PAPER (для всех типов заказа)
	// номер телефона для услуги SMS
	// код фотопроекта для услуги PHOTO_DOCUMENT
	Parameter string `json:"parameter,omitempty"`
}

type OrderRegisterEntity struct {
	// Uuid Идентификатор заказа в ИС СДЭК
	Uuid string `json:"uuid,omitempty"`
}

type OrderRegisterRequests struct {
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
}

type OrderRegisterRelatedEntities struct {
	// Type Тип связанной сущности. Может принимать значения: waybill - квитанция к заказу, barcode - ШК места к заказу
	Type string `json:"type"`
	// Uuid Идентификатор сущности, связанной с заказом
	Uuid string `json:"uuid"`
}

type OrderRegisterError struct {
	// Message Описание ошибки
	Message string `json:"message"`
	// Code string Код ошибки
	Code string `json:"code"`
}

type OrderRegisterResponse struct {
	// Entity Информация о заказе
	Entity OrderRegisterEntity `json:"entity,omitempty"`
	// Requests Информация о запросе над заказом
	Requests []OrderRegisterRequests `json:"requests"`
	// RelatedEntities Связанные сущности (если в запросе был передан корректный print)
	RelatedEntities OrderRegisterRelatedEntities `json:"related_entities,omitempty"`
}

func (c *clientImpl) OrderRegister(ctx context.Context, input *OrderRegisterRequest) (*OrderRegisterResponse, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.buildUri("/v2/orders", nil),
		bytes.NewReader(payload),
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

	resp, err := jsonReq[OrderRegisterResponse](req)
	if err != nil {
		return nil, err
	}

	var result error
	for _, item := range resp.Requests {
		if item.State == "INVALID" {
			result = multierror.Append(result, fmt.Errorf("%+v", item))
		}
	}

	if result != nil {
		return nil, result
	}

	return resp, nil
}
