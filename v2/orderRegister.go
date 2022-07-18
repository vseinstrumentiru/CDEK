package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

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
	DeliveryRecipientCost Payment `json:"delivery_recipient_cost"`
	// DeliveryRecipientCostAdv Доп. сбор за доставку (которую ИМ берет с получателя) в зависимости от суммы заказа
	// Только для заказов "интернет-магазин".  Возможно указать несколько порогов.
	DeliveryRecipientCostAdv Cost `json:"delivery_recipient_cost_adv"`
	// FromLocation Адрес отправления. Не может использоваться одновременно с shipment_point
	FromLocation Location `json:"from_location"`
	// ToLocation Адрес получения. Не может использоваться одновременно с delivery_point
	ToLocation Location `json:"to_location"`
	// Packages Список информации по местам (упаковкам). Количество мест в заказе может быть от 1 до 255
	Packages []Package `json:"packages,omitempty"`
	// Recipient Получатель
	Recipient RecipientSender `json:"recipient"`
	// Sender Отправитель. Обязателен если:
	// нет, если заказ типа "интернет-магазин"
	// да, если заказ типа "доставка"
	Sender RecipientSender `json:"sender,omitempty"`
	// Services Дополнительные услуги
	Services []Service `json:"services,omitempty"`
	// Seller Реквизиты истинного продавца. Только для заказов "интернет-магазин"
	Seller Seller `json:"seller,omitempty"`
}

func (c *clientImpl) OrderRegister(ctx context.Context, input *OrderRegisterRequest) (*Response, error) {
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

	resp, err := jsonReq[Response](req)
	if err != nil {
		return nil, err
	}

	if err := validateResponse(resp.Requests); err != nil {
		return nil, err
	}

	return resp, nil
}
