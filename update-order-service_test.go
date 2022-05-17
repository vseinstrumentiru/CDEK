package cdek

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_UpdateOrder(t *testing.T) {
	mockServer := updateOrderMockServer()
	defer mockServer.Close()

	mockServerWithValidError := updateOrderMockServerWithValidError()
	defer mockServerWithValidError.Close()

	mockServerWithError := updateOrderMockServerWithError()
	defer mockServerWithError.Close()

	type fields struct {
		client clientImpl
	}
	type args struct {
		req UpdateOrderReq
	}

	testOrderToUpdate := NewUpdateOrder().
		SetDeliveryRecipientCost(10.02).
		SetDeliveryRecipientVATRate("VATX").
		SetDeliveryRecipientVATSum(0.0).
		SetNumber("number-s785558445").
		SetPackage(*NewOrderPackage("soOEl00", "barcode-soOEl00", 100).
			SetSize(2, 3, 4).
			AddItem(*NewOrderPackageItem(2, "warekey-soOEl000", 8, 10, 1, "comment-soOEl000").
				SetPaymentVATRate("VATX").
				SetPaymentVATSum(0),
			),
		)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UpdateOrderResp
		wantErr bool
	}{
		{
			name: "creation",
			fields: fields{
				client: clientImpl{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				req: *NewUpdateOrderReq("soOEl", 1, *testOrderToUpdate),
			},
			want: &UpdateOrderResp{
				Order: []*OrderResp{
					{
						DispatchNumber: intLink(1105062403),
						Number:         strLink("number-s785558445"),
					},
					{
						Error: Error{
							Msg: strLink("Изменено заказов 1"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "handle valid error",
			fields: fields{
				client: clientImpl{
					apiURL: mockServerWithValidError.URL,
				},
			},
			args: args{
				req: *NewUpdateOrderReq("soOEl", 1, *testOrderToUpdate.SetNumber("notFoundNumber")),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uncompilable url",
			fields: fields{
				client: clientImpl{
					apiURL: " wrong://url ",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong url",
			fields: fields{
				client: clientImpl{
					apiURL: "wrong://url",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "server error",
			fields: fields{
				client: clientImpl{
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
			got, err := cl.UpdateOrder(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientImpl.UpdateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clientImpl.UpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func updateOrderMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`
			<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
			<response>
				<Order DispatchNumber="1105062403" Number="number-s785558445"/>
				<Order Msg="Изменено заказов 1"/>
				<TraceId>6a02b7f9bfa85283</TraceId>
			</response>
		`))
	}))
}

func updateOrderMockServerWithValidError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`
			<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
			<response>
				<Order ErrorCode="ERR_ORDER_NOTFIND" Msg="Заказ не найден в базе СДЭК" Number="number"/>
				<Order Msg="Изменено заказов 0"/>
				<TraceId>1e183b7496f3f3d2</TraceId>
			</response>
		`))
	}))
}

func updateOrderMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("err"))
	}))
}

func ExampleClient_UpdateOrder() {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")

	exampleOrderToUpdate := NewUpdateOrder().
		SetDeliveryRecipientCost(10.02).
		SetDeliveryRecipientVATRate("VATX").
		SetDeliveryRecipientVATSum(0.0).
		SetNumber("number-s785558445").
		SetPackage(*NewOrderPackage("soOEl00", "barcode-soOEl00", 100).
			SetSize(2, 3, 4).
			AddItem(*NewOrderPackageItem(2, "warekey-soOEl000", 8, 10, 1, "comment-soOEl000").
				SetPaymentVATRate("VATX").
				SetPaymentVATSum(0),
			),
		)

	ctx := context.TODO()
	result, err := client.UpdateOrder(ctx, *NewUpdateOrderReq("soOEl", 1, *exampleOrderToUpdate))

	_, _ = result, err
}
