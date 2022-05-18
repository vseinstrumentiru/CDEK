package cdek

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func getCitiesGetMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`[{
			"cityUuid":"dece8676-077e-4793-9d67-96112fe96b03",
			"cityName":"Москва",
			"cityCode":"61627",
			"region":"Кировская",
			"regionCodeExt":"43",
			"regionCode":"44",
			"subRegion":"Верхошижемский",
			"country":"Russia",
			"countryCode":"RU",
			"latitude":57.9664,
			"longitude":49.1074,
			"kladr":"4300700005400",
			"fiasGuid":"f1c72b9d-a2d7-45b7-b9f5-2222c12d5164",
			"regionFiasGuid":null,"paymentLimit":0
		},{
			"cityUuid":"18bd1ad1-0ed5-4908-9069-db07b805aa53",
			"cityName":"Москва",
			"cityCode":"44",
			"region":"Москва",
			"regionCodeExt":"77",
			"regionCode":"81",
			"subRegion":"Москва",
			"country":"Russia",
			"countryCode":"RU",
			"latitude":55.754,
			"longitude":37.6204,
			"kladr":"7700000000000",
			"fiasGuid":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5",
			"regionFiasGuid":null,
			"paymentLimit":-1
		}]`))
	}))
}

func getCitiesGetMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("err"))
	}))
}
func getCitiesGetMockServerWithValidError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`
			{
				"alerts":[
					{"type":"danger","msg":"API location not available","errorCode":"connector.location.error.send"}
				]
			}
		`))
	}))
}

func TestClient_GetCities(t *testing.T) {
	testServer := getCitiesGetMockServer()
	defer testServer.Close()

	testServerWithError := getCitiesGetMockServerWithError()
	defer testServerWithError.Close()

	testServerWithValidError := getCitiesGetMockServerWithValidError()
	defer testServerWithError.Close()

	var citiesFilterBuilder CityFilterBuilder
	citiesFilterBuilder.AddFilter(CityFilterCityName, "Москва")

	var filterCausesServerError CityFilterBuilder
	filterCausesServerError.AddFilter(CityFilterPage, "2000")

	type args struct {
		client Client
		filter map[CityFilter]string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetCitiesResp
		wantErr bool
	}{
		{
			name: "got cities",
			args: args{
				client: Client{
					apiURL: testServer.URL,
				},
				filter: citiesFilterBuilder.Filter(),
			},
			want: &GetCitiesResp{
				&City{
					strLink("dece8676-077e-4793-9d67-96112fe96b03"),
					strLink("Москва"),
					strLink("61627"),
					strLink("Кировская"),
					intLink(43),
					intLink(44),
					strLink("Верхошижемский"),
					strLink("Russia"),
					strLink("RU"),
					float64Link(57.9664),
					float64Link(49.1074),
					strLink("4300700005400"),
					strLink("f1c72b9d-a2d7-45b7-b9f5-2222c12d5164"),
					nil,
					float64Link(0),
				},
				&City{
					strLink("18bd1ad1-0ed5-4908-9069-db07b805aa53"),
					strLink("Москва"),
					strLink("44"),
					strLink("Москва"),
					intLink(77),
					intLink(81),
					strLink("Москва"),
					strLink("Russia"),
					strLink("RU"),
					float64Link(55.754),
					float64Link(37.6204),
					strLink("7700000000000"),
					strLink("0c5b2444-70a0-4932-980c-b4dc0d3f02b5"),
					nil,
					float64Link(-1),
				},
			},
			wantErr: false,
		},
		{
			name: "server error",
			args: args{
				client: Client{
					apiURL: testServerWithError.URL,
				},
				filter: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uncompilable url",
			args: args{
				client: Client{
					apiURL: " wrong://url ",
				},
				filter: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong url",
			args: args{
				client: Client{
					apiURL: "wrong://url",
				},
				filter: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid error on service",
			args: args{
				client: Client{
					apiURL: testServerWithValidError.URL,
				},
				filter: filterCausesServerError.Filter(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.args.client
			got, err := cl.GetCities(ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCities() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleClient_GetCities() {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")

	ctx := context.TODO()
	result, err := client.GetCities(ctx, map[CityFilter]string{
		CityFilterPage: "3",
	})

	_, _ = result, err
}
