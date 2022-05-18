package cdek

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func calculateDeliveryGetMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		_, _ = req.Body.Read(body)
		var getCostReq GetCostReq
		_ = json.Unmarshal(body, &getCostReq)

		var errorsResp []Error
		if *getCostReq.Version != apiVersion {
			errorsResp = append(errorsResp, Error{
				ErrorCode: strLink("1"),
				Msg:       strLink("Указанная вами версия API не поддерживается"),
			})
		}
		if *getCostReq.Secure == "" || *getCostReq.Account == "" || *getCostReq.Date == "" {
			errorsResp = append(errorsResp, Error{
				ErrorCode: strLink("2"),
				Msg:       strLink("Ошибка авторизации"),
			})
		}
		if *getCostReq.SenderCityID == 0 {
			errorsResp = append(errorsResp, Error{
				ErrorCode: strLink("7"),
				Msg:       strLink("Не задан город-отправитель"),
			})
		}
		if *getCostReq.ReceiverCityID == 0 {
			errorsResp = append(errorsResp, Error{
				ErrorCode: strLink("8"),
				Msg:       strLink("Не задан город-получатель"),
			})
		}
		if getCostReq.Goods == nil {
			errorsResp = append(errorsResp, Error{
				ErrorCode: strLink("8"),
				Msg:       strLink("Не задано ни одного места для отправления"),
			})
		}

		if *getCostReq.TariffID == 0 {
			errorsResp = append(errorsResp, Error{
				ErrorCode: strLink("6"),
				Msg:       strLink("Не задан тариф или список тарифов"),
			})
		}

		result, _ := json.Marshal(&getCostResp{
			ErrorResp: errorsResp,
			Result: GetCostRespResult{
				Price:             100,
				DeliveryPeriodMin: 1,
				DeliveryPeriodMax: 2,
				TariffID:          3,
			},
		})

		_, _ = res.Write(result)
	}))
}

func calculateDeliveryGetMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("err"))
	}))
}

func TestClient_CalculateDelivery(t *testing.T) {
	testServer := calculateDeliveryGetMockServer()
	defer testServer.Close()

	testServerWithError := calculateDeliveryGetMockServerWithError()
	defer testServerWithError.Close()

	type args struct {
		client Client
		req    GetCostReq
	}
	tests := []struct {
		name    string
		args    args
		want    *GetCostRespResult
		wantErr bool
	}{
		{
			"success deliveries calculated",
			args{
				client: Client{
					auth: &auth{
						account: "123",
						secure:  "123",
					},
					apiURL:        "",
					calculatorURL: testServer.URL,
				},
				req: GetCostReq{
					Version:        strLink(apiVersion),
					SenderCityID:   intLink(1),
					ReceiverCityID: intLink(2),
					TariffID:       intLink(3),
					Goods: []*Good{
						{
							Weight: 1.1,
							Length: 2,
							Width:  3,
							Height: 4,
							Volume: 5.5,
						},
					},
					Services: nil,
				},
			},
			&GetCostRespResult{
				Price:             100,
				DeliveryPeriodMin: 1,
				DeliveryPeriodMax: 2,
				TariffID:          3,
			},
			false,
		},
		{
			"wrong goods",
			args{
				client: Client{
					auth: &auth{
						account: "123",
						secure:  "123",
					},
					apiURL:        "",
					calculatorURL: testServer.URL,
				},
				req: GetCostReq{
					Version:        strLink(apiVersion),
					SenderCityID:   intLink(1),
					ReceiverCityID: intLink(2),
					TariffID:       intLink(3),
					Goods:          nil,
					Services:       nil,
				},
			},
			nil,
			true,
		},
		{
			"server error",
			args{
				client: Client{
					auth: &auth{
						account: "123",
						secure:  "123",
					},
					apiURL:        "",
					calculatorURL: testServerWithError.URL,
				},
				req: GetCostReq{
					Version:        strLink(apiVersion),
					SenderCityID:   intLink(1),
					ReceiverCityID: intLink(2),
					TariffID:       intLink(3),
					Goods:          nil,
					Services:       nil,
				},
			},
			nil,
			true,
		},
		{
			"marshal error",
			args{
				client: Client{
					auth: &auth{
						account: "123",
						secure:  "123",
					},
					apiURL:        "",
					calculatorURL: testServerWithError.URL,
				},
				req: GetCostReq{
					Version:        strLink(apiVersion),
					SenderCityID:   nil,
					ReceiverCityID: intLink(2),
					TariffID:       intLink(3),
					Goods: []*Good{
						{
							Weight: math.Inf(1),
						},
					},
					Services: nil,
				},
			},
			nil,
			true,
		},
		{
			"wrong url",
			args{
				client: Client{
					calculatorURL: "wrong url",
				},
				req: GetCostReq{},
			},
			nil,
			true,
		},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.args.client
			got, err := cl.CalculateDelivery(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateDelivery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateDelivery() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleClient_CalculateDelivery() {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")

	ctx := context.TODO()
	result, err := client.CalculateDelivery(ctx, *NewGetCostReq(61208, 2108, 10))

	_, _ = result, err
}
