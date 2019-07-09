package cdek

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		clientConf ClientConf
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		{
			"client created",
			args{
				clientConf: ClientConf{
					Auth: &Auth{
						Account: "Account",
						Secure:  "Secure",
					},
					CdekAPIURL:    "CdekAPIURL",
					CalculatorURL: "CalculatorURL",
				},
			},
			&client{
				clientConf: ClientConf{
					Auth: &Auth{
						Account: "Account",
						Secure:  "Secure",
					},
					CdekAPIURL:    "CdekAPIURL",
					CalculatorURL: "CalculatorURL",
				},
			},
		},
		{
			"empty client created",
			args{
				clientConf: ClientConf{},
			},
			&client{
				clientConf: ClientConf{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.clientConf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
