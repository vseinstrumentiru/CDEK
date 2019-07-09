package cdek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func getMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		_, _ = req.Body.Read(body)
		var getCostReq GetCostReq
		_ = json.Unmarshal(body, &getCostReq)

		errorsResp := make([]ErrorResp, 0)
		if *getCostReq.Version != apiVersion {
			errorsResp = append(errorsResp, ErrorResp{
				Code: 1,
				Text: "Указанная вами версия API не поддерживается",
			})
		}
		if *getCostReq.Secure == "" || *getCostReq.AuthLogin == "" || *getCostReq.DateExecute == "" {
			errorsResp = append(errorsResp, ErrorResp{
				Code: 2,
				Text: "Ошибка авторизации",
			})
		}
		if *getCostReq.SenderCityID == 0 {
			errorsResp = append(errorsResp, ErrorResp{
				Code: 7,
				Text: "Не задан город-отправитель",
			})
		}
		if *getCostReq.ReceiverCityID == 0 {
			errorsResp = append(errorsResp, ErrorResp{
				Code: 8,
				Text: "Не задан город-получатель",
			})
		}
		if getCostReq.Goods == nil {
			errorsResp = append(errorsResp, ErrorResp{
				Code: 8,
				Text: "Не задано ни одного места для отправления",
			})
		}

		if *getCostReq.TariffID == 0 {
			errorsResp = append(errorsResp, ErrorResp{
				Code: 6,
				Text: "Не задан тариф или список тарифов",
			})
		}

		result, _ := json.Marshal(&GetCostResp{
			Error: errorsResp,
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

func getMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte{})
	}))
}

func Test_calculateDelivery(t *testing.T) {
	testServer := getMockServer()
	testServerWithError := getMockServerWithError()

	defer testServer.Close()

	apiVersion := apiVersion
	senderCityID := 1
	receiverCityID := 2
	tariffID := 3

	type args struct {
		clientConf ClientConf
		req        GetCostReq
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
				clientConf: ClientConf{
					Auth: &Auth{
						Account: "123",
						Secure:  "123",
					},
					CdekAPIURL:    "",
					CalculatorURL: testServer.URL,
				},
				req: GetCostReq{
					Version:        &apiVersion,
					SenderCityID:   &senderCityID,
					ReceiverCityID: &receiverCityID,
					TariffID:       &tariffID,
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
				clientConf: ClientConf{
					Auth: &Auth{
						Account: "123",
						Secure:  "123",
					},
					CdekAPIURL:    "",
					CalculatorURL: testServer.URL,
				},
				req: GetCostReq{
					Version:        &apiVersion,
					SenderCityID:   &senderCityID,
					ReceiverCityID: &receiverCityID,
					TariffID:       &tariffID,
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
				clientConf: ClientConf{
					Auth: &Auth{
						Account: "123",
						Secure:  "123",
					},
					CdekAPIURL:    "",
					CalculatorURL: testServerWithError.URL,
				},
				req: GetCostReq{
					Version:        &apiVersion,
					SenderCityID:   &senderCityID,
					ReceiverCityID: &receiverCityID,
					TariffID:       &tariffID,
					Goods:          nil,
					Services:       nil,
				},
			},
			nil,
			true,
		},
		{
			"wrong url",
			args{
				clientConf: ClientConf{
					CalculatorURL: "wrong url",
				},
				req: GetCostReq{},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateDelivery(tt.args.clientConf, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateDelivery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				g, _ := json.Marshal(got)
				w, _ := json.Marshal(tt.want)
				t.Errorf("calculateDelivery() got \n %v \n want \n %v", string(g), string(w))
			}
		})
	}
}
