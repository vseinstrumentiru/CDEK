package cdek

import (
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
		_, _ = res.Write([]byte(`{"alerts":[{"type":"danger","msg":"API location not available","errorCode":"connector.location.error.send"}]}`))
	}))
}

func Test_client_GetCities(t *testing.T) {
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
		clientConf ClientConf
		filter     map[CityFilter]string
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
				clientConf: ClientConf{
					CdekAPIURL: testServer.URL,
				},
				filter: citiesFilterBuilder.Filter(),
			},
			want: &GetCitiesResp{
				buildTestCity("dece8676-077e-4793-9d67-96112fe96b03",
					"Москва",
					61627,
					"Кировская",
					43,
					44,
					"Верхошижемский",
					"Russia",
					"RU",
					57.9664,
					49.1074,
					"4300700005400",
					"f1c72b9d-a2d7-45b7-b9f5-2222c12d5164",
					0),
				buildTestCity("18bd1ad1-0ed5-4908-9069-db07b805aa53",
					"Москва",
					44,
					"Москва",
					77,
					81,
					"Москва",
					"Russia",
					"RU",
					55.754,
					37.6204,
					"7700000000000",
					"0c5b2444-70a0-4932-980c-b4dc0d3f02b5",
					-1),
			},
			wantErr: false,
		},
		{
			name: "server error",
			args: args{
				clientConf: ClientConf{
					CdekAPIURL: testServerWithError.URL,
				},
				filter: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uncompilable url",
			args: args{
				clientConf: ClientConf{
					CdekAPIURL: " wrong://url ",
				},
				filter: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong url 2",
			args: args{
				clientConf: ClientConf{
					CdekAPIURL: "wrong://url",
				},
				filter: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid error on service",
			args: args{
				clientConf: ClientConf{
					CdekAPIURL: testServerWithValidError.URL,
				},
				filter: filterCausesServerError.Filter(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := client{
				clientConf: tt.args.clientConf,
			}
			got, err := cl.GetCities(tt.args.filter)
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

func buildTestCity(CityUUID string,
	CityName string,
	CityCode int,
	Region string,
	RegionCodeExt int,
	RegionCode int,
	SubRegion string,
	Country string,
	CountryCode string,
	Latitude float64,
	Longitude float64,
	Kladr string,
	FiasGUID string,
	PaymentLimit float64,
) *City {
	return &City{
		&CityUUID,
		&CityName,
		&CityCode,
		&Region,
		&RegionCodeExt,
		&RegionCode,
		&SubRegion,
		&Country,
		&CountryCode,
		&Latitude,
		&Longitude,
		&Kladr,
		&FiasGUID,
		nil,
		&PaymentLimit,
	}
}
