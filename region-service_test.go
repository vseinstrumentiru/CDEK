package cdek

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_GetRegions(t *testing.T) {
	mockServer := getRegionsMockServer()
	defer mockServer.Close()

	mockServerWithValidError := getRegionsMockServerWithValidError()
	defer mockServerWithValidError.Close()

	mockServerWithError := getRegionsMockServerWithError()
	defer mockServerWithError.Close()

	type fields struct {
		client Client
	}
	type args struct {
		filter map[RegionFilter]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetRegionsResp
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
				filter: map[RegionFilter]string{
					RegionFilterRegionFiasGUID: "61723327-1c20-42fe-8dfa-402638d9b396",
				},
			},
			want: &GetRegionsResp{
				Region{
					RegionUUID:     strLink("18aff43f-58b8-4608-ade7-92fdab7fc39f"),
					RegionName:     strLink("Тверская"),
					Prefix:         strLink("обл"),
					RegionCodeExt:  intLink(69),
					RegionCode:     intLink(50),
					RegionFiasGUID: strLink("61723327-1c20-42fe-8dfa-402638d9b396"),
					CountryName:    strLink("Россия"),
					CountryCode:    strLink("RU"),
					CountryCodeExt: intLink(643),
				},
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
			args: args{
				filter: map[RegionFilter]string{
					RegionFilterRegionFiasGUID: "61723327-1c20-42fe-8dfa-402638d9b396",
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
			got, err := cl.GetRegions(ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetRegions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetRegions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getRegionsMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`[
			{
				"regionName":"Тверская",
				"regionCode":"50",
				"regionUuid":"18aff43f-58b8-4608-ade7-92fdab7fc39f",
				"prefix":"обл",
				"regionCodeExt":"69",
				"regionFiasGuid":"61723327-1c20-42fe-8dfa-402638d9b396",
				"countryName":"Россия",
				"countryCode":"RU",
				"countryCodeExt":"643"
			}
		]`))
	}))
}

func getRegionsMockServerWithValidError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`
			{
				"alerts": [
					{"type":"danger","msg":"API location not available","errorCode":"connector.location.error.send"}
				]
			}
		`))
	}))
}

func getRegionsMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("err"))
	}))
}

func ExampleClient_GetRegions() {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")

	ctx := context.TODO()
	result, err := client.GetRegions(ctx, map[RegionFilter]string{
		RegionFilterPage: "3",
	})

	_, _ = result, err
}
