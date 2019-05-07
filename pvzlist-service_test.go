package cdek

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

func getPvzListGetMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte(`
			<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
			<PvzList>
				<Pvz 
					Code="MIA2" PostalCode="456300" Name="На Романенко" CountryCode="1" countryCodeIso="RU" 
					CountryName="Россия" RegionCode="3" RegionName="Челябинская обл." CityCode="7" City="Миасс" 
					WorkTime="Пн-Пт 09:00-19:00, Сб 10:00-16:00" Address="ул. Романенко, 93"
					FullAddress="Россия, Челябинская обл., Миасс, ул. Романенко, 93" 
					AddressComment="Бывший магазин Престиж, выше ТРК «Слон»" Phone="+79511247307, +73513284466" 
					Email="n.andruschuk@cdek.ru" qqId="" Note="" coordX="60.112205" coordY="55.047874" Type="PVZ" 
					ownerCode="cdek" IsDressingRoom="true" HaveCashless="true" AllowedCod="true" 
					NearestStation="«Центр», «Лихачева»" MetroStation="" Site=""
				>
					<PhoneDetail number="+79511247307"/>
					<PhoneDetail number="+73513284466"/>
					<OfficeImage number="1" 
						url="edu.api-pvz.imageRepository.service.cdek.tech:8008/images/2638/3222_1_MIA2"/>
					<OfficeImage number="2" 
						url="edu.api-pvz.imageRepository.service.cdek.tech:8008/images/2638/3230_2_MIA2"/>
					<OfficeImage number="3" 
						url="edu.api-pvz.imageRepository.service.cdek.tech:8008/images/2638/3231_3_MIA2"/>
					<OfficeImage number="4" 
						url="edu.api-pvz.imageRepository.service.cdek.tech:8008/images/2638/3232_4_MIA2"/>
					<WorkTimeY day="1" periods="09:00/19:00"/>
					<WorkTimeY day="2" periods="09:00/19:00"/>
					<WorkTimeY day="3" periods="09:00/19:00"/>
					<WorkTimeY day="4" periods="09:00/19:00"/>
					<WorkTimeY day="5" periods="09:00/19:00"/>
					<WorkTimeY day="6" periods="10:00/16:00"/>
				</Pvz>
			</PvzList>
		`))
	}))
}

func getPvzListGetMockServerWithError() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("err"))
	}))
}

func TestClient_GetPvzList(t *testing.T) {
	mockServer := getPvzListGetMockServer()
	defer mockServer.Close()

	mockServerWithError := getPvzListGetMockServerWithError()
	defer mockServerWithError.Close()

	type fields struct {
		client Client
	}
	type args struct {
		filter map[PvzListFilter]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Pvz
		wantErr bool
	}{
		{
			"server ok",
			fields{
				Client{
					apiURL: mockServer.URL,
				},
			},
			args{
				map[PvzListFilter]string{
					PvzListFilterCityID: strconv.Itoa(7),
				},
			},
			[]*Pvz{
				{
					Code:           strLink("MIA2"),
					PostalCode:     strLink("456300"),
					Name:           strLink("На Романенко"),
					CountryCode:    strLink("1"),
					CountryCodeIso: strLink("RU"),
					CountryName:    strLink("Россия"),
					RegionCode:     strLink("3"),
					RegionName:     strLink("Челябинская обл."),
					CityCode:       intLink(7),
					City:           strLink("Миасс"),
					WorkTime:       strLink("Пн-Пт 09:00-19:00, Сб 10:00-16:00"),
					Address:        strLink("ул. Романенко, 93"),
					FullAddress:    strLink("Россия, Челябинская обл., Миасс, ул. Романенко, 93"),
					AddressComment: strLink("Бывший магазин Престиж, выше ТРК «Слон»"),
					Phone:          strLink("+79511247307, +73513284466"),
					Email:          strLink("n.andruschuk@cdek.ru"),
					QqID:           strLink(""),
					Note:           strLink(""),
					CoordX:         float64Link(60.112205),
					CoordY:         float64Link(55.047874),
					Type:           strLink("PVZ"),
					OwnerCode:      strLink("cdek"),
					IsDressingRoom: boolLink(true),
					HaveCashless:   boolLink(true),
					AllowedCod:     boolLink(true),
					NearestStation: strLink("«Центр», «Лихачева»"),
					MetroStation:   strLink(""),
					Site:           strLink(""),
					OfficeImage: []*OfficeImage{
						{
							strLink("edu.api-pvz.imageRepository.service.cdek.tech:8008/images/2638/3222_1_MIA2"),
						},
						{
							strLink("edu.api-pvz.imageRepository.service.cdek.tech:8008/images/2638/3230_2_MIA2"),
						},
						{
							strLink("edu.api-pvz.imageRepository.service.cdek.tech:8008/images/2638/3231_3_MIA2"),
						},
						{
							strLink("edu.api-pvz.imageRepository.service.cdek.tech:8008/images/2638/3232_4_MIA2"),
						},
					},
					WorkTimeY: []*WorkTimeY{
						{
							Day:     intLink(1),
							Periods: strLink("09:00/19:00"),
						},
						{
							Day:     intLink(2),
							Periods: strLink("09:00/19:00"),
						},
						{
							Day:     intLink(3),
							Periods: strLink("09:00/19:00"),
						},
						{
							Day:     intLink(4),
							Periods: strLink("09:00/19:00"),
						},
						{
							Day:     intLink(5),
							Periods: strLink("09:00/19:00"),
						},
						{
							Day:     intLink(6),
							Periods: strLink("10:00/16:00"),
						},
					},
				},
			},
			false,
		},
		{
			"server error",
			fields{
				Client{
					apiURL: mockServerWithError.URL,
				},
			},
			args{},
			nil,
			true,
		},
		{
			"wrong url parse error",
			fields{
				Client{
					apiURL: " wrong://url ",
				},
			},
			args{},
			nil,
			true,
		},
		{
			"wrong url",
			fields{
				Client{
					apiURL: "wrong://url",
				},
			},
			args{},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.fields.client
			got, err := cl.GetPvzList(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetPvzList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g, _ := json.Marshal(got)
			w, _ := json.Marshal(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetPvzList() = \n %v \n, want \n %v", string(g), string(w))
			}
		})
	}
}

func ExampleClient_GetPvzList() {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")

	result, err := client.GetPvzList(map[PvzListFilter]string{
		PvzListFilterCityID: "44",
	})

	_, _ = result, err
}
