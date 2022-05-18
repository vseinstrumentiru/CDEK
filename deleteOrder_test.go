package cdek

import (
	"context"
	"encoding/xml"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func deleteOrderGetMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_ = req.ParseForm()
		xmlRequest := req.FormValue("xml_request")
		var deleteOrderReq DeleteOrderRequest
		_ = xml.Unmarshal([]byte(xmlRequest), &deleteOrderReq)

		if deleteOrderReq.OrderCount > 1 {
			_, _ = res.Write([]byte(`
				<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
				<response>
					<Order ErrorCode="ERR_NOT_EQUAL_ORDERCOUNT" Msg="Указано неверное количество заказов"/>
					<TraceId>b64b57647abbd3f7</TraceId>
				</response>
			`))

			return
		}

		if deleteOrderReq.Number != "number-soOEl0" {
			_, _ = res.Write([]byte(`
				<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
				<response>
					<Order ErrorCode="ERR_ORDER_NOTFIND" Msg="Заказ не найден в базе СДЭК" Number="test_order_number"/>
					<Order Msg="Удалено заказов 0"/>
					<TraceId>20f55f887b90d054</TraceId>
				</response>
			`))

			return
		}

		_, _ = res.Write([]byte(`
			<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
			<response>
				<Order DispatchNumber="1105048590" Number="number-soOEl0"/>
				<Order Msg="Удалено заказов 1"/>
				<TraceId>2a706c5c926fc49d</TraceId>
			</response>
		`))
	}))
}

func deleteOrderGetMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("err"))
	}))
}

func TestClient_DeleteOrder(t *testing.T) {
	mockServer := deleteOrderGetMockServer()
	defer mockServer.Close()

	mockServerWithError := deleteOrderGetMockServerWithError()
	defer mockServerWithError.Close()

	type fields struct {
		client clientImpl
	}
	type args struct {
		req DeleteOrderRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *DeleteOrderResponse
		wantErr bool
	}{
		{
			name: "delete",
			fields: fields{
				client: clientImpl{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				req: DeleteOrderRequest{
					Number:     "number-soOEl0",
					OrderCount: 1,
					Order: DeleteOrder{
						Number: "number-soOEl0",
					},
				},
			},
			want: &DeleteOrderResponse{
				Order: []*OrderResp{
					{
						DispatchNumber: intLink(1105048590),
						Number:         strLink("number-soOEl0"),
					},
					{
						Error: Error{
							Msg: strLink("Удалено заказов 1"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "order not found",
			fields: fields{
				client: clientImpl{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				req: DeleteOrderRequest{
					Number:     "test",
					OrderCount: 1,
					Order: DeleteOrder{
						Number: "test_order_number",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong quantity error",
			fields: fields{
				client: clientImpl{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				req: DeleteOrderRequest{
					Number:     "test",
					OrderCount: 2,
					Order: DeleteOrder{
						Number: "test_order_number",
					},
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
			args: args{
				req: DeleteOrderRequest{},
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
			args: args{
				req: DeleteOrderRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong parse url",
			fields: fields{
				client: clientImpl{
					apiURL: " wrong://url ",
				},
			},
			args: args{
				req: DeleteOrderRequest{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.fields.client
			got, err := cl.DeleteOrder(ctx, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func ExampleClient_DeleteOrder() {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")

	ctx := context.TODO()
	result, err := client.DeleteOrder(ctx, DeleteOrderRequest{
		Number:     "number-soOEl0",
		OrderCount: 1,
		Order: DeleteOrder{
			Number:         "number-soOEl0",
			DispatchNumber: 1,
		},
	})

	_, _ = result, err
}
