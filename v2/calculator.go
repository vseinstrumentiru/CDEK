package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Package struct {
	Weight int `json:"weight"`
	Length int `json:"length,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Location struct {
	// Code Код населенного пункта СДЭК (метод "Список населенных пунктов")
	Code string `json:"code,omitempty"`
	// PostalCode Почтовый индекс города, для которого необходим список офисов
	PostalCode int `json:"postal_code,omitempty"`
	// City Название города
	City string `json:"city,omitempty"`
	// Address Название города
	Address string `json:"address,omitempty"`
}

type CalculatorTrafiffListRequest struct {
	// Date Дата и время планируемой передачи заказа. По умолчанию - текущая
	Date string `json:"date,omitempty"`
	// Type Тип заказа: 1 - "интернет-магазин", 2 - "доставка". По умолчанию - 1
	Type string `json:"type,omitempty"`
	// Валюта, в которой необходимо произвести расчет. По умолчанию - валюта договора
	Currency string `json:"currency,omitempty"`
	// Lang Локализация офиса. По умолчанию "rus"
	Lang string `url:"lang,omitempty"`
	// FromLocation Адрес отправления
	FromLocation Location `json:"from_location,omitempty"`
	// ToLocation Адрес получения
	ToLocation Location `json:"to_location"`
	// Packages Список информации по местам (упаковкам)
	Packages []Package `json:"packages"`
}

type Tariff struct {
	TariffCode        int     `json:"tariff_code"`
	TariffName        string  `json:"tariff_name"`
	TariffDescription string  `json:"tariff_description"`
	DeliveryMode      int     `json:"delivery_mode"`
	DeliverySum       float64 `json:"delivery_sum"`
	PeriodMin         int     `json:"period_min"`
	PeriodMax         int     `json:"period_max"`
}

type CalculatorTrafiffListResponse struct {
	TariffCodes []Tariff `json:"tariff_codes"`
}

func (c *clientImpl) CalculatorTrafiffList(ctx context.Context, input *CalculatorTrafiffListRequest) (*CalculatorTrafiffListResponse, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.buildUri("/v2/calculator/tarifflist", nil),
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	return jsonReq[CalculatorTrafiffListResponse](req)
}
