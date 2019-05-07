package cdek

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func deleteOrderGetMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_ = req.ParseForm()
		xmlRequest := req.FormValue("xml_request")
		var deleteOrderReq DeleteOrderReq
		_ = xml.Unmarshal([]byte(xmlRequest), &deleteOrderReq)

		if *deleteOrderReq.OrderCount > 1 {
			_, _ = res.Write([]byte(`
				<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
				<response>
					<Order ErrorCode="ERR_NOT_EQUAL_ORDERCOUNT" Msg="Указано неверное количество заказов"/>
					<TraceId>b64b57647abbd3f7</TraceId>
				</response>
			`))

			return
		}

		if *deleteOrderReq.Number != "number-soOEl0" {
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
		client Client
	}
	type args struct {
		req DeleteOrderReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *DeleteOrderResp
		wantErr bool
	}{
		{
			name: "delete",
			fields: fields{
				client: Client{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				req: DeleteOrderReq{
					Number:     strLink("number-soOEl0"),
					OrderCount: intLink(1),
					Order: &DeleteOrder{
						Number: strLink("number-soOEl0"),
					},
				},
			},
			want: &DeleteOrderResp{
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
				client: Client{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				req: DeleteOrderReq{
					Number:     strLink("test"),
					OrderCount: intLink(1),
					Order: &DeleteOrder{
						Number: strLink("test_order_number"),
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong quantity error",
			fields: fields{
				client: Client{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				req: DeleteOrderReq{
					Number:     strLink("test"),
					OrderCount: intLink(2),
					Order: &DeleteOrder{
						Number: strLink("test_order_number"),
					},
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
			args: args{
				req: DeleteOrderReq{},
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
			args: args{
				req: DeleteOrderReq{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong parse url",
			fields: fields{
				client: Client{
					apiURL: " wrong://url ",
				},
			},
			args: args{
				req: DeleteOrderReq{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.fields.client
			got, err := cl.DeleteOrder(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				g, _ := json.Marshal(got)
				w, _ := json.Marshal(tt.want)
				t.Errorf("DeleteOrder() got = %v, want %v", string(g), string(w))
			}
		})
	}
}

func ExampleClient_DeleteOrder() {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")

	result, err := client.DeleteOrder(*NewDeleteOrderReq(
		"number-soOEl0",
		1,
		*NewDeleteOrder().SetNumber("number-soOEl0"),
	))

	_, _ = result, err
}
