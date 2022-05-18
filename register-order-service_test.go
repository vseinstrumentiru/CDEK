package cdek

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestClient_RegisterOrder(t *testing.T) {
	mockServer := registerOrderMockServer()
	defer mockServer.Close()

	mockServerWithValidError := registerOrderMockServerWithValidError()
	defer mockServerWithValidError.Close()

	mockServerWithError := registerOrderMockServerWithError()
	defer mockServerWithError.Close()

	type fields struct {
		client Client
	}
	type args struct {
		req RegisterOrderReq
	}
	testOrder := NewOrderReq("number-soOEl0", "name-soOEl0", "+79138739944", 139).
		SetSendCityCode(44).
		SetRecCityCode(44).
		SetComment("comment-soOEl0").
		SetDeliveryRecipientCost(0).
		SetDeliveryRecipientVATRate("VATX").
		SetDeliveryRecipientVATSum(0).
		SetRecipientEmail("no-reply@cdek.ru").
		SetAddress(*NewAddress("street-soOEl0", "house-soOEl0").
			SetFlat("flat-soOEl0"),
		).
		SetSender(*NewSender().
			SetCompany("company-soOEl0").
			SetName("Отправителев").
			SetAddress(*NewAddress("street-soOEl0", "house-soOEl0").
				SetFlat("flat-soOEl0"),
			).
			AddPhone("+79138739946").
			AddPhone("+79138739945"),
		).
		SetSeller(*NewSeller().
			SetAddress("street_soOEl0 1").
			SetName("seller-soOEl0").
			SetINN("111111111111").
			SetPhone("+79138739947").
			SetOwnershipForm(249),
		).
		AddPackage(*NewOrderPackage("soOEl00", "barcode-soOEl00", 100).
			SetSize(2, 3, 4).
			AddItem(*NewOrderPackageItem(2, "warekey-soOEl000", 8, 10, 1, "comment-soOEl000").
				SetPaymentVATRate("VATX").
				SetPaymentVATSum(0),
			),
		).
		SetDeliveryRecipientCostAdv(*NewDeliveryRecipientCostAdv(2000, 150).
			SetVATRate("vat10").
			SetVATSum(13.64),
		).
		SetAddService(*NewAddService(30)).
		SetSchedule(*NewSchedule().
			AddAttempt(*NewScheduleAttempt("soOEl00", time.Date(2019, 7, 19, 12, 0, 0, 0, time.UTC)).
				SetAddress(*NewAddress("street-prozvon_adr", "house-prozvon_adr").
					SetFlat("flat-prozvon_adr"),
				).
				SetTimeBeg("11:00:00").
				SetTimeEnd("13:00:00"),
			),
		)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RegisterOrderResp
		wantErr bool
	}{
		{
			name: "creation",
			fields: fields{
				client: Client{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				req: *NewDeliveryRequest("soOEl", 1, testOrder),
			},
			want: &RegisterOrderResp{
				Order: []*OrderResp{
					{
						DispatchNumber: intLink(1105068300),
						Number:         strLink("number-soOEl0"),
					},
					{
						Error: Error{
							Msg: strLink("Добавлено заказов 1"),
						},
					},
				},
				Call: nil,
			},
			wantErr: false,
		},
		{
			name: "handle valid error",
			fields: fields{
				client: Client{
					apiURL: mockServerWithValidError.URL,
				},
			},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uncompilable url",
			fields: fields{
				client: Client{
					apiURL: " wrong://url ",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong url",
			fields: fields{
				client: Client{
					apiURL: "wrong://url",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "server error",
			fields: fields{
				client: Client{
					apiURL: mockServerWithError.URL,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.fields.client
			got, err := cl.RegisterOrder(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RegisterOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.RegisterOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func registerOrderMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`
			<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
			<response>
				<Order DispatchNumber="1105068300" Number="number-soOEl0"/>
				<Order Msg="Добавлено заказов 1"/>
				<TraceId>4770a74ae95c3d19</TraceId>
			</response>
		`))
	}))
}

func registerOrderMockServerWithValidError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`
			<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
			<response>
				<Order ErrorCode="ERR_SENDCITYPOSTCODE" Msg="Город отправителя не определен по индексу"/>
				<Call ErrorCode="ERR_EXAMPLEERRCODE" Msg="Какая-то ошибка"/>
				<TraceId>d5f71d55d760b9b7</TraceId>
			</response>
		`))
	}))
}

func registerOrderMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("err"))
	}))
}

func ExampleClient_RegisterOrder() {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")

	exampleOrder := NewOrderReq("number-soOEl0", "name-soOEl0", "+79138739944", 139).
		SetSendCityCode(44).
		SetRecCityCode(44).
		SetComment("comment-soOEl0").
		SetDeliveryRecipientCost(0).
		SetDeliveryRecipientVATRate("VATX").
		SetDeliveryRecipientVATSum(0).
		SetRecipientEmail("no-reply@cdek.ru").
		SetAddress(*NewAddress("street-soOEl0", "house-soOEl0").
			SetFlat("flat-soOEl0"),
		).
		SetSender(*NewSender().
			SetCompany("company-soOEl0").
			SetName("Отправителев").
			SetAddress(*NewAddress("street-soOEl0", "house-soOEl0").
				SetFlat("flat-soOEl0"),
			).
			AddPhone("+79138739946").
			AddPhone("+79138739945"),
		).
		SetSeller(*NewSeller().
			SetAddress("street_soOEl0 1").
			SetName("seller-soOEl0").
			SetINN("111111111111").
			SetPhone("+79138739947").
			SetOwnershipForm(249),
		).
		AddPackage(*NewOrderPackage("soOEl00", "barcode-soOEl00", 100).
			SetSize(2, 3, 4).
			AddItem(*NewOrderPackageItem(2, "warekey-soOEl000", 8, 10, 1, "comment-soOEl000").
				SetPaymentVATRate("VATX").
				SetPaymentVATSum(0),
			),
		).
		SetDeliveryRecipientCostAdv(*NewDeliveryRecipientCostAdv(2000, 150).
			SetVATRate("vat10").
			SetVATSum(13.64),
		).
		SetAddService(*NewAddService(30)).
		SetSchedule(*NewSchedule().
			AddAttempt(*NewScheduleAttempt("soOEl00", time.Date(2019, 7, 19, 12, 0, 0, 0, time.UTC)).
				SetAddress(*NewAddress("street-prozvon_adr", "house-prozvon_adr").
					SetFlat("flat-prozvon_adr"),
				).
				SetTimeBeg("11:00:00").
				SetTimeEnd("13:00:00"),
			),
		)

	ctx := context.TODO()
	result, err := client.RegisterOrder(ctx, *NewDeliveryRequest("soOEl", 1, exampleOrder))

	_, _ = result, err
}
