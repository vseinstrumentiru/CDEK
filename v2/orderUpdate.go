package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderUpdateRequest struct {
	// UUID Идентификатор заказа в ИС СДЭК, который нужно изменить (да, если не заполнен cdek_number)
	UUID string `json:"uuid,omitempty"`
	// CdekNumber Номер заказа СДЭК, который нужно изменить (да, если не заполнен uuid)
	CdekNumber string `json:"cdek_number,omitempty"`
	// Код тарифа (режимы старого и нового тарифа должны совпадать)
	TariffCode int `json:"tariff_code,omitempty"`
	// Comment Комментарий к заказу
	Comment string `json:"comment"`
	// ShipmentPoint Код ПВЗ СДЭК, на который будет производится забор отправления либо самостоятельный привоз клиентом. Не может использоваться одновременно с from_location
	ShipmentPoint string `json:"shipment_point,omitempty"`
	// DeliveryPoint Код ПВЗ СДЭК, на который будет доставлена посылка. Не может использоваться одновременно с to_location
	DeliveryPoint string `json:"delivery_point,omitempty"`
	// OrderDeliveryRecipientCost Доп. сбор за доставку, которую ИМ берет с получателя. Валюта сбора должна совпадать с валютой наложенного платежа
	DeliveryRecipientCost Payment `json:"delivery_recipient_cost"`
	// DeliveryRecipientCostAdv Доп. сбор за доставку (которую ИМ берет с получателя) в зависимости от суммы заказа. Только для заказов "интернет-магазин". Возможно указать несколько порогов.
	DeliveryRecipientCostAdv Cost `json:"delivery_recipient_cost_adv"`
	// Sender Отправитель. Обязателен если:
	// нет, если заказ типа "интернет-магазин"
	// да, если заказ типа "доставка"
	Sender RecipientSender `json:"sender,omitempty"`
	// Seller Реквизиты истинного продавца
	Seller Seller `json:"seller,omitempty"`
	// Recipient Получатель
	Recipient RecipientSender `json:"recipient,omitempty"`
	// ToLocation Адрес получения. Не может использоваться одновременно с delivery_point
	ToLocation Location `json:"to_location"`
	// FromLocation Адрес отправления. Не может использоваться одновременно с shipment_point
	FromLocation Location `json:"from_location"`
	// Services Дополнительные услуги
	Services []Service `json:"services,omitempty"`
	// Packages Список информации по местам (упаковкам)
	Packages []Package `json:"packages,omitempty"`
}

type OrderUpdateResponse struct {
	Entity   ResponseEntity     `json:"entity,omitempty"`
	Requests []ResponseRequests `json:"requests"`
}

func (c *clientImpl) OrderUpdate(ctx context.Context, input *OrderUpdateRequest) (*OrderUpdateResponse, error) {
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

	resp, err := jsonReq[OrderUpdateResponse](req)
	if err != nil {
		return nil, err
	}

	if err := validateResponse(resp.Requests); err != nil {
		return nil, err
	}

	return resp, nil
}
