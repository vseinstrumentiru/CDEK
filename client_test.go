package cdek

import (
	"reflect"
	"testing"
)

func ClientConfForTests() *ClientConf {
	return &ClientConf{
		Auth: Auth{
			Account: "z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd",
			Secure:  "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq",
		},
		CdekAPIURL: "https://integration.edu.cdek.ru",
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		clientConfig ClientConf
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		{
			"client created",
			args{
				clientConfig: *ClientConfForTests(),
			},
			&client{
				clientConfig: *ClientConfForTests(),
			},
		},
		{
			"empty client created",
			args{
				clientConfig: *new(ClientConf),
			},
			&client{
				clientConfig: *new(ClientConf),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.clientConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_GetPvzList(t *testing.T) {
	type fields struct {
		clientConfig ClientConf
	}
	type args struct {
		filter map[PvzListFilter]string
	}
	cityCode := "44"
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"pvz city id filter working",
			fields{
				*ClientConfForTests(),
			},
			args{
				map[PvzListFilter]string{
					PvzListFilterCityID: cityCode,
				},
			},
			cityCode,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := client{
				clientConfig: tt.fields.clientConfig,
			}
			pvzlist, err := cl.GetPvzList(tt.args.filter)
			if pvzlist == nil {
				t.Errorf("client.GetPvzList() error = %v", "nothing received")
				return
			}
			if len(pvzlist) < 1 {
				t.Errorf("client.GetPvzList() error = %v", "received pvz quantity less than 1")
				return
			}
			got := *pvzlist[0].CityCode
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetPvzList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.GetPvzList() = %v, want %v", got, tt.want)
			}
		})
	}
}
