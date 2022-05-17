package cdek

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getCitiesGetMockServer() *httptest.Server {
	fn := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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
	})
	return httptest.NewServer(fn)
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
				client: &clientImpl{apiURL: testServer.URL},
				filter: map[CityFilter]string{CityFilterCityName: "Москва"},
			},
			want: &GetCitiesResp{
				City{
					CityUUID:      strLink("dece8676-077e-4793-9d67-96112fe96b03"),
					CityName:      strLink("Москва"),
					CityCode:      strLink("61627"),
					Region:        strLink("Кировская"),
					RegionCodeExt: intLink(43),
					RegionCode:    intLink(44),
					SubRegion:     strLink("Верхошижемский"),
					Country:       strLink("Russia"),
					CountryCode:   strLink("RU"),
					Latitude:      float64Link(57.9664),
					Longitude:     float64Link(49.1074),
					Kladr:         strLink("4300700005400"),
					FiasGUID:      strLink("f1c72b9d-a2d7-45b7-b9f5-2222c12d5164"),
					PaymentLimit:  float64Link(0),
				},
				City{
					CityUUID:      strLink("18bd1ad1-0ed5-4908-9069-db07b805aa53"),
					CityName:      strLink("Москва"),
					CityCode:      strLink("44"),
					Region:        strLink("Москва"),
					RegionCodeExt: intLink(77),
					RegionCode:    intLink(81),
					SubRegion:     strLink("Москва"),
					Country:       strLink("Russia"),
					CountryCode:   strLink("RU"),
					Latitude:      float64Link(55.754),
					Longitude:     float64Link(37.6204),
					Kladr:         strLink("7700000000000"),
					FiasGUID:      strLink("0c5b2444-70a0-4932-980c-b4dc0d3f02b5"),
					PaymentLimit:  float64Link(-1),
				},
			},
			wantErr: false,
		},
		{
			name: "server error",
			args: args{
				client: &clientImpl{apiURL: testServerWithError.URL},
				filter: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uncompilable url",
			args: args{
				client: &clientImpl{apiURL: " wrong://url "},
				filter: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong url",
			args: args{
				client: &clientImpl{apiURL: "wrong://url"},
				filter: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid error on service",
			args: args{
				client: &clientImpl{apiURL: testServerWithValidError.URL},
				filter: map[CityFilter]string{CityFilterPage: "2000"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.client.GetCities(ctx, tt.args.filter)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_ExampleClient_GetCities(t *testing.T) {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")

	ctx := context.TODO()
	result, err := client.GetCities(ctx, map[CityFilter]string{
		CityFilterCityName: "Москва",
	})
	require.NoError(t, err)
	require.NotNil(t, result)
}
