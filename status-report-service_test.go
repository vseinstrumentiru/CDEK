package cdek

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_GetStatusReport(t *testing.T) {
	mockServer := getStatusReportMockServer()
	defer mockServer.Close()

	mockServerWithValidError := getStatusReportMockServerWithValidError()
	defer mockServerWithValidError.Close()

	mockServerWithError := getStatusReportMockServerWithError()
	defer mockServerWithError.Close()

	type fields struct {
		client clientImpl
	}
	type args struct {
		statusReportReq StatusReport
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *StatusReportResp
		wantErr bool
	}{
		{
			name: "handle response",
			fields: fields{
				client: clientImpl{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				statusReportReq: StatusReport{
					ShowHistory: intLink(1),
					Order: []*StatusReportOrderReq{
						{
							Number: strLink("number-soOEl0"),
						},
					},
				},
			},
			want: &StatusReportResp{
				Order: []*StatusReportOrderResp{
					{
						ActNumber:      strLink("soOEl"),
						Number:         strLink("number-soOEl0"),
						DispatchNumber: intLink(1105068433),
						Status: &Status{
							Date:        strLink("2019-07-21T17:34:34+00:00"),
							Code:        intLink(1),
							Description: strLink("Создан"),
							CityCode:    intLink(44),
							CityName:    strLink("Москва"),
							State: []*State{
								{
									Date:        strLink("2019-07-21T17:34:34+00:00"),
									Code:        intLink(1),
									Description: strLink("Создан"),
									CityCode:    intLink(44),
									CityName:    strLink("Москва"),
								},
							},
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
				statusReportReq: StatusReport{
					ShowHistory: intLink(1),
					ChangePeriod: &ChangePeriod{
						DateFirst: strLink("14-07-2019"),
					},
				},
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
			args: args{
				statusReportReq: StatusReport{},
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
				statusReportReq: StatusReport{},
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
				statusReportReq: StatusReport{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.fields.client
			got, err := cl.GetStatusReport(ctx, tt.args.statusReportReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientImpl.GetStatusReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				g, _ := json.Marshal(got)
				w, _ := json.Marshal(tt.want)
				t.Errorf("clientImpl.GetStatusReport() = \n %v \n, want \n %v", string(g), string(w))
			}
		})
	}
}

func getStatusReportMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
			<StatusReport DateFirst="2000-12-31T17:00:00+00:00" DateLast="2019-07-21T18:05:39+00:00" >
				<Order
					ActNumber="soOEl" 
					Number="number-soOEl0" 
					DispatchNumber="1105068433"  
					DeliveryDate="" 
					RecipientName="" >
					<Status Date="2019-07-21T17:34:34+00:00" 
						Code="1" 
						Description="Создан" 
						CityCode="44" 
						CityName="Москва">
						<State 
							Date="2019-07-21T17:34:34+00:00" 
							Code="1" 
							Description="Создан" 
							CityCode="44" 
							CityName="Москва" />
					</Status>
					<Reason Code="" Description="" Date=""></Reason>
					<DelayReason Code="" Description="" Date="" ></DelayReason>
				</Order>
			</StatusReport>
		`))
	}))
}

func getStatusReportMockServerWithValidError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`
			<?xml version="1.0" encoding="UTF-8"?>
			<StatusReport ErrorCode="ERR_DATEFORMAT" Msg="Неверный формат даты в параметре Date = "/>
		`))
	}))
}

func getStatusReportMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("err"))
	}))
}
