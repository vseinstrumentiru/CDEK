package cdek

import (
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
		client Client
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
				client: Client{
					apiURL: mockServer.URL,
				},
			},
			args: args{
				statusReportReq: StatusReport{
					ShowHistory: boolLink(true),
					Order: []*StatusReportOrderReq{
						{
							DispatchNumber: intLink(1105256461),
						},
					},
				},
			},
			// TODO: wait for the CDEK technical support answer and finish the test
			want:    nil,
			wantErr: true,
		},
		{
			name: "handle valid error",
			fields: fields{
				client: Client{
					apiURL: mockServerWithValidError.URL,
				},
			},
			args: args{
				statusReportReq: StatusReport{
					ShowHistory: boolLink(true),
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
				client: Client{
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
				client: Client{
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
				client: Client{
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.fields.client
			got, err := cl.GetStatusReport(tt.args.statusReportReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetStatusReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetStatusReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getStatusReportMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// TODO: wait for the CDEK technical support answer and finish the test
		_, _ = res.Write([]byte(``))
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
