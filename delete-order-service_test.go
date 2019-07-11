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

		_, _ = res.Write([]byte(`
			<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
			<response>
				<Order ErrorCode="ERR_ORDER_NOTFIND" Msg="Заказ не найден в базе СДЭК" Number="test_order_number"/>
				<Order Msg="Удалено заказов 0"/>
				<TraceId>20f55f887b90d054</TraceId>
			</response>
		`))
	}))
}

func deleteOrderGetMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("err"))
	}))
}

func Test_client_DeleteOrder(t *testing.T) {
	mockServer := deleteOrderGetMockServer()
	defer mockServer.Close()

	mockServerWithError := deleteOrderGetMockServerWithError()
	defer mockServerWithError.Close()

	type fields struct {
		clientConf ClientConf
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
			name: "order not found",
			fields: fields{
				clientConf: ClientConf{
					CdekAPIURL: mockServer.URL,
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
			want: &DeleteOrderResp{
				XMLName: xml.Name{
					Local: "response",
				},
				Order: []*OrderResp{
					{
						Number: strLink("test_order_number"),
						ErrorXML: ErrorXML{
							ErrorCode: strLink("ERR_ORDER_NOTFIND"),
							Msg:       strLink("Заказ не найден в базе СДЭК"),
						},
					},
					{
						ErrorXML: ErrorXML{
							Msg: strLink("Удалено заказов 0"),
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "wrong quantity error",
			fields: fields{
				clientConf: ClientConf{
					CdekAPIURL: mockServer.URL,
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
			want: &DeleteOrderResp{
				XMLName: xml.Name{
					Local: "response",
				},
				Order: []*OrderResp{
					{
						ErrorXML: ErrorXML{
							ErrorCode: strLink("ERR_NOT_EQUAL_ORDERCOUNT"),
							Msg:       strLink("Указано неверное количество заказов"),
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "server error",
			fields: fields{
				clientConf: ClientConf{
					CdekAPIURL: mockServerWithError.URL,
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
				clientConf: ClientConf{
					CdekAPIURL: "wrong://url",
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
				clientConf: ClientConf{
					CdekAPIURL: " wrong://url ",
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
			cl := client{
				clientConf: tt.fields.clientConf,
			}
			got, err := cl.DeleteOrder(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				g, _ := json.Marshal(got)
				w, _ := json.Marshal(tt.want)
				t.Errorf("DeleteOrder() got = %v, want %v", string(g), string(w))
			}
		})
	}
}
