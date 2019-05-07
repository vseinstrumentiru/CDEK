package cdek

import (
	"reflect"
	"testing"
)

func TestNewGetCostReq(t *testing.T) {
	apiVersion := apiVersion
	senderCityID := 1
	receiverCityID := 2
	tariffID := 3

	type args struct {
		senderCityID   int
		receiverCityID int
		tariffID       int
	}
	tests := []struct {
		name string
		args args
		want *GetCostReq
	}{
		{
			"constructor is ok",
			args{
				senderCityID:   1,
				receiverCityID: 2,
				tariffID:       3,
			},
			&GetCostReq{
				Version:        &apiVersion,
				SenderCityID:   &senderCityID,
				ReceiverCityID: &receiverCityID,
				TariffID:       &tariffID,
				Goods:          nil,
				Services:       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetCostReq(tt.args.senderCityID, tt.args.receiverCityID, tt.args.tariffID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGetCostReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCostReq_AddService(t *testing.T) {
	type fields struct {
		Version        *string
		AuthLogin      *string
		Secure         *string
		DateExecute    *string
		SenderCityID   *int
		ReceiverCityID *int
		TariffID       *int
		Goods          []*Good
		Services       []*ServiceReq
	}
	type args struct {
		service ServiceReq
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetCostReq
	}{
		{
			"constructor is ok",
			fields{
				Version:        nil,
				AuthLogin:      nil,
				Secure:         nil,
				DateExecute:    nil,
				SenderCityID:   nil,
				ReceiverCityID: nil,
				TariffID:       nil,
				Goods:          nil,
				Services:       nil,
			},
			args{
				service: ServiceReq{
					ID:    1,
					Param: 2,
				},
			},
			&GetCostReq{
				Version:        nil,
				SenderCityID:   nil,
				ReceiverCityID: nil,
				TariffID:       nil,
				Goods:          nil,
				Services: []*ServiceReq{
					{
						ID:    1,
						Param: 2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getCostReq := &GetCostReq{
				Version:        tt.fields.Version,
				SenderCityID:   tt.fields.SenderCityID,
				ReceiverCityID: tt.fields.ReceiverCityID,
				TariffID:       tt.fields.TariffID,
				Goods:          tt.fields.Goods,
				Services:       tt.fields.Services,
			}
			if got := getCostReq.AddService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCostReq_AddGood(t *testing.T) {
	type fields struct {
		Version        *string
		AuthLogin      *string
		Secure         *string
		DateExecute    *string
		SenderCityID   *int
		ReceiverCityID *int
		TariffID       *int
		Goods          []*Good
		Services       []*ServiceReq
	}
	type args struct {
		good Good
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetCostReq
	}{
		{
			"constructor is ok",
			fields{
				Version:        nil,
				AuthLogin:      nil,
				Secure:         nil,
				DateExecute:    nil,
				SenderCityID:   nil,
				ReceiverCityID: nil,
				TariffID:       nil,
				Goods:          nil,
				Services:       nil,
			},
			args{
				good: Good{
					Weight: 1.11,
					Length: 2,
					Width:  3,
					Height: 4,
					Volume: 5.55,
				},
			},
			&GetCostReq{
				Version:        nil,
				SenderCityID:   nil,
				ReceiverCityID: nil,
				TariffID:       nil,
				Goods: []*Good{
					{
						Weight: 1.11,
						Length: 2,
						Width:  3,
						Height: 4,
						Volume: 5.55,
					},
				},
				Services: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getCostReq := &GetCostReq{
				Version:        tt.fields.Version,
				SenderCityID:   tt.fields.SenderCityID,
				ReceiverCityID: tt.fields.ReceiverCityID,
				TariffID:       tt.fields.TariffID,
				Goods:          tt.fields.Goods,
				Services:       tt.fields.Services,
			}
			if got := getCostReq.AddGood(tt.args.good); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
