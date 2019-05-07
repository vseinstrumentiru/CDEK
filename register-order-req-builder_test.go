package cdek

import (
	"reflect"
	"testing"
	"time"
)

func TestNewDeliveryRequest(t *testing.T) {
	type args struct {
		number     string
		orderCount int
		order      *OrderReq
	}
	tests := []struct {
		name string
		args args
		want *RegisterOrderReq
	}{
		{
			name: "constructor",
			args: args{
				number:     "test_number",
				orderCount: 1,
				order:      &OrderReq{},
			},
			want: &RegisterOrderReq{
				Number:     strLink("test_number"),
				OrderCount: intLink(1),
				Order:      &OrderReq{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeliveryRequest(tt.args.number, tt.args.orderCount, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeliveryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterOrderReq_SetCurrency(t *testing.T) {
	type fields struct {
		Number      *string
		OrderCount  *int
		Currency    *string
		Order       *OrderReq
		CallCourier *CallCourier
	}
	type args struct {
		currency string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RegisterOrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:     strLink("test_number"),
				OrderCount: intLink(1),
				Order:      &OrderReq{},
				Currency:   nil,
			},
			args: args{
				currency: "RUB",
			},
			want: &RegisterOrderReq{
				Number:     strLink("test_number"),
				OrderCount: intLink(1),
				Order:      &OrderReq{},
				Currency:   strLink("RUB"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:     strLink("test_number"),
				OrderCount: intLink(1),
				Order:      &OrderReq{},
				Currency:   strLink("EUR"),
			},
			args: args{
				currency: "RUB",
			},
			want: &RegisterOrderReq{
				Number:     strLink("test_number"),
				OrderCount: intLink(1),
				Order:      &OrderReq{},
				Currency:   strLink("RUB"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registerOrderReq := &RegisterOrderReq{
				Number:      tt.fields.Number,
				OrderCount:  tt.fields.OrderCount,
				Currency:    tt.fields.Currency,
				Order:       tt.fields.Order,
				CallCourier: tt.fields.CallCourier,
			}
			if got := registerOrderReq.SetCurrency(tt.args.currency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterOrderReq.SetCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterOrderReq_SetCallCourier(t *testing.T) {
	type fields struct {
		Number      *string
		OrderCount  *int
		Currency    *string
		Order       *OrderReq
		CallCourier *CallCourier
	}
	type args struct {
		callCourier CallCourier
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RegisterOrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:      strLink("test_number"),
				OrderCount:  intLink(1),
				Order:       &OrderReq{},
				CallCourier: nil,
			},
			args: args{
				callCourier: CallCourier{
					Call: &CourierCallReq{
						Date:       strLink("2019-07-15"),
						TimeBeg:    strLink("10:00"),
						TimeEnd:    strLink("17:00"),
						SendPhone:  strLink("+79138739944"),
						SenderName: strLink("full name"),
						SendAddress: &Address{
							Street: strLink("test street"),
							House:  strLink("10/3"),
						},
					},
				},
			},
			want: &RegisterOrderReq{
				Number:     strLink("test_number"),
				OrderCount: intLink(1),
				Order:      &OrderReq{},
				CallCourier: &CallCourier{
					Call: &CourierCallReq{
						Date:       strLink("2019-07-15"),
						TimeBeg:    strLink("10:00"),
						TimeEnd:    strLink("17:00"),
						SendPhone:  strLink("+79138739944"),
						SenderName: strLink("full name"),
						SendAddress: &Address{
							Street: strLink("test street"),
							House:  strLink("10/3"),
						},
					},
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:     strLink("test_number"),
				OrderCount: intLink(1),
				Order:      &OrderReq{},
				CallCourier: &CallCourier{
					Call: &CourierCallReq{
						Date:       strLink("2019-07-16"),
						TimeBeg:    strLink("11:00"),
						TimeEnd:    strLink("15:00"),
						SendPhone:  strLink("+79138739944"),
						SenderName: strLink("full name"),
						SendAddress: &Address{
							Street: strLink("test street"),
							House:  strLink("10/3"),
						},
					},
				},
			},
			args: args{
				callCourier: CallCourier{
					Call: &CourierCallReq{
						Date:       strLink("2019-07-15"),
						TimeBeg:    strLink("10:00"),
						TimeEnd:    strLink("17:00"),
						SendPhone:  strLink("+79138739944"),
						SenderName: strLink("full name"),
						SendAddress: &Address{
							Street: strLink("test street"),
							House:  strLink("10/3"),
						},
					},
				},
			},
			want: &RegisterOrderReq{
				Number:     strLink("test_number"),
				OrderCount: intLink(1),
				Order:      &OrderReq{},
				CallCourier: &CallCourier{
					Call: &CourierCallReq{
						Date:       strLink("2019-07-15"),
						TimeBeg:    strLink("10:00"),
						TimeEnd:    strLink("17:00"),
						SendPhone:  strLink("+79138739944"),
						SenderName: strLink("full name"),
						SendAddress: &Address{
							Street: strLink("test street"),
							House:  strLink("10/3"),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registerOrderReq := &RegisterOrderReq{
				Number:      tt.fields.Number,
				OrderCount:  tt.fields.OrderCount,
				Currency:    tt.fields.Currency,
				Order:       tt.fields.Order,
				CallCourier: tt.fields.CallCourier,
			}
			if got := registerOrderReq.SetCallCourier(tt.args.callCourier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterOrderReq.SetCallCourier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOrderReq(t *testing.T) {
	type args struct {
		number         string
		recipientName  string
		phone          string
		tariffTypeCode int
	}
	tests := []struct {
		name string
		args args
		want *OrderReq
	}{
		{
			name: "constructor",
			args: args{
				number:         "test_number",
				recipientName:  "recipient name",
				phone:          "+79138739944",
				tariffTypeCode: 10,
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOrderReq(tt.args.number, tt.args.recipientName, tt.args.phone, tt.args.tariffTypeCode)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetSendCityCode(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		sendCityCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				SendCityCode:   nil,
			},
			args: args{
				sendCityCode: 1234321,
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				SendCityCode:   intLink(1234321),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				SendCityCode:   intLink(4321234),
			},
			args: args{
				sendCityCode: 1234321,
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				SendCityCode:   intLink(1234321),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetSendCityCode(tt.args.sendCityCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetSendCityCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetRecCityCode(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		recCityCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCityCode:    nil,
			},
			args: args{
				recCityCode: 1234321,
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCityCode:    intLink(1234321),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCityCode:    intLink(4321234),
			},
			args: args{
				recCityCode: 1234321,
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCityCode:    intLink(1234321),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetRecCityCode(tt.args.recCityCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetRecCityCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetSendCityPostCode(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		sendCityPostCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:           strLink("test_number"),
				RecipientName:    strLink("recipient name"),
				Phone:            strLink("+79138739944"),
				TariffTypeCode:   intLink(10),
				SendCityPostCode: nil,
			},
			args: args{
				sendCityPostCode: 1234321,
			},
			want: &OrderReq{
				Number:           strLink("test_number"),
				RecipientName:    strLink("recipient name"),
				Phone:            strLink("+79138739944"),
				TariffTypeCode:   intLink(10),
				SendCityPostCode: intLink(1234321),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:           strLink("test_number"),
				RecipientName:    strLink("recipient name"),
				Phone:            strLink("+79138739944"),
				TariffTypeCode:   intLink(10),
				SendCityPostCode: intLink(4321234),
			},
			args: args{
				sendCityPostCode: 1234321,
			},
			want: &OrderReq{
				Number:           strLink("test_number"),
				RecipientName:    strLink("recipient name"),
				Phone:            strLink("+79138739944"),
				TariffTypeCode:   intLink(10),
				SendCityPostCode: intLink(1234321),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetSendCityPostCode(tt.args.sendCityPostCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetSendCityPostCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetRecCityPostCode(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		recCityPostCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:          strLink("test_number"),
				RecipientName:   strLink("recipient name"),
				Phone:           strLink("+79138739944"),
				TariffTypeCode:  intLink(10),
				RecCityPostCode: nil,
			},
			args: args{
				recCityPostCode: 1234321,
			},
			want: &OrderReq{
				Number:          strLink("test_number"),
				RecipientName:   strLink("recipient name"),
				Phone:           strLink("+79138739944"),
				TariffTypeCode:  intLink(10),
				RecCityPostCode: intLink(1234321),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:          strLink("test_number"),
				RecipientName:   strLink("recipient name"),
				Phone:           strLink("+79138739944"),
				TariffTypeCode:  intLink(10),
				RecCityPostCode: intLink(4321234),
			},
			args: args{
				recCityPostCode: 1234321,
			},
			want: &OrderReq{
				Number:          strLink("test_number"),
				RecipientName:   strLink("recipient name"),
				Phone:           strLink("+79138739944"),
				TariffTypeCode:  intLink(10),
				RecCityPostCode: intLink(1234321),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetRecCityPostCode(tt.args.recCityPostCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetRecCityPostCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetSendCountryCode(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		sendCountryCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:          strLink("test_number"),
				RecipientName:   strLink("recipient name"),
				Phone:           strLink("+79138739944"),
				TariffTypeCode:  intLink(10),
				SendCountryCode: nil,
			},
			args: args{
				sendCountryCode: 1234321,
			},
			want: &OrderReq{
				Number:          strLink("test_number"),
				RecipientName:   strLink("recipient name"),
				Phone:           strLink("+79138739944"),
				TariffTypeCode:  intLink(10),
				SendCountryCode: intLink(1234321),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:          strLink("test_number"),
				RecipientName:   strLink("recipient name"),
				Phone:           strLink("+79138739944"),
				TariffTypeCode:  intLink(10),
				SendCountryCode: intLink(4321234),
			},
			args: args{
				sendCountryCode: 1234321,
			},
			want: &OrderReq{
				Number:          strLink("test_number"),
				RecipientName:   strLink("recipient name"),
				Phone:           strLink("+79138739944"),
				TariffTypeCode:  intLink(10),
				SendCountryCode: intLink(1234321),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetSendCountryCode(tt.args.sendCountryCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetSendCountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetRecCountryCode(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		recCountryCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCountryCode: nil,
			},
			args: args{
				recCountryCode: 1234321,
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCountryCode: intLink(1234321),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCountryCode: intLink(4321234),
			},
			args: args{
				recCountryCode: 1234321,
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCountryCode: intLink(1234321),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetRecCountryCode(tt.args.recCountryCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetRecCountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetSendCityName(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		sendCityName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				SendCityName:   nil,
			},
			args: args{
				sendCityName: "Багинск",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				SendCityName:   strLink("Багинск"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				SendCityName:   strLink("Суздаль"),
			},
			args: args{
				sendCityName: "Багинск",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				SendCityName:   strLink("Багинск"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetSendCityName(tt.args.sendCityName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetSendCityName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetRecCityName(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		recCityName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCityName:    nil,
			},
			args: args{
				recCityName: "Багинск",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCityName:    strLink("Багинск"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCityName:    strLink("Суздаль"),
			},
			args: args{
				recCityName: "Багинск",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecCityName:    strLink("Багинск"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetRecCityName(tt.args.recCityName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetRecCityName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetRecipientINN(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		recipientINN string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecipientINN:   nil,
			},
			args: args{
				recipientINN: "12345678987",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecipientINN:   strLink("12345678987"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecipientINN:   strLink("98765432123"),
			},
			args: args{
				recipientINN: "12345678987",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecipientINN:   strLink("12345678987"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetRecipientINN(tt.args.recipientINN); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetRecipientINN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetDateInvoice(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		dateInvoice time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				DateInvoice:    nil,
			},
			args: args{
				dateInvoice: time.Date(2019, 7, 20, 0, 0, 0, 0, time.UTC),
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				DateInvoice:    strLink("2019-07-20"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				DateInvoice:    strLink("2019-07-17"),
			},
			args: args{
				dateInvoice: time.Date(2019, 7, 20, 0, 0, 0, 0, time.UTC),
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				DateInvoice:    strLink("2019-07-20"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetDateInvoice(tt.args.dateInvoice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetDateInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetShipperName(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		shipperName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ShipperName:    nil,
			},
			args: args{
				shipperName: "shipper name",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ShipperName:    strLink("shipper name"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ShipperName:    strLink("previous shipper name"),
			},
			args: args{
				shipperName: "shipper name",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ShipperName:    strLink("shipper name"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetShipperName(tt.args.shipperName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetShipperName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetShipperAddress(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		shipperAddress string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ShipperAddress: nil,
			},
			args: args{
				shipperAddress: "shipper address",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ShipperAddress: strLink("shipper address"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ShipperAddress: strLink("previous shipper address"),
			},
			args: args{
				shipperAddress: "shipper address",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ShipperAddress: strLink("shipper address"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetShipperAddress(tt.args.shipperAddress); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetShipperAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetPassport(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		passport Passport
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Passport:       nil,
			},
			args: args{
				passport: Passport{
					Series: strLink("some series"),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Passport: &Passport{
					Series: strLink("some series"),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Passport: &Passport{
					Series: strLink("previous series"),
				},
			},
			args: args{
				passport: Passport{
					Series: strLink("some series"),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Passport: &Passport{
					Series: strLink("some series"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetPassport(tt.args.passport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetPassport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetSender(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		sender Sender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Sender:         nil,
			},
			args: args{
				sender: Sender{
					Name: strLink("sender name"),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Sender: &Sender{
					Name: strLink("sender name"),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Sender: &Sender{
					Name: strLink("previous sender name"),
				},
			},
			args: args{
				sender: Sender{
					Name: strLink("sender name"),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Sender: &Sender{
					Name: strLink("sender name"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetSender(tt.args.sender); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetSender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetRecipientEmail(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		recipientEmail string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecipientEmail: nil,
			},
			args: args{
				recipientEmail: "recipient@actual.email",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecipientEmail: strLink("recipient@actual.email"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecipientEmail: strLink("recipient@previous.email"),
			},
			args: args{
				recipientEmail: "recipient@actual.email",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				RecipientEmail: strLink("recipient@actual.email"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetRecipientEmail(tt.args.recipientEmail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetRecipientEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetDeliveryRecipientCost(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		deliveryRecipientCost float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:                strLink("test_number"),
				RecipientName:         strLink("recipient name"),
				Phone:                 strLink("+79138739944"),
				TariffTypeCode:        intLink(10),
				DeliveryRecipientCost: nil,
			},
			args: args{
				deliveryRecipientCost: 399.99,
			},
			want: &OrderReq{
				Number:                strLink("test_number"),
				RecipientName:         strLink("recipient name"),
				Phone:                 strLink("+79138739944"),
				TariffTypeCode:        intLink(10),
				DeliveryRecipientCost: float64Link(399.99),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:                strLink("test_number"),
				RecipientName:         strLink("recipient name"),
				Phone:                 strLink("+79138739944"),
				TariffTypeCode:        intLink(10),
				DeliveryRecipientCost: float64Link(400.),
			},
			args: args{
				deliveryRecipientCost: 399.99,
			},
			want: &OrderReq{
				Number:                strLink("test_number"),
				RecipientName:         strLink("recipient name"),
				Phone:                 strLink("+79138739944"),
				TariffTypeCode:        intLink(10),
				DeliveryRecipientCost: float64Link(399.99),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetDeliveryRecipientCost(tt.args.deliveryRecipientCost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetDeliveryRecipientCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetDeliveryRecipientVATRate(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		deliveryRecipientVATRate string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:                   strLink("test_number"),
				RecipientName:            strLink("recipient name"),
				Phone:                    strLink("+79138739944"),
				TariffTypeCode:           intLink(10),
				DeliveryRecipientVATRate: nil,
			},
			args: args{
				deliveryRecipientVATRate: "VATX",
			},
			want: &OrderReq{
				Number:                   strLink("test_number"),
				RecipientName:            strLink("recipient name"),
				Phone:                    strLink("+79138739944"),
				TariffTypeCode:           intLink(10),
				DeliveryRecipientVATRate: strLink("VATX"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:                   strLink("test_number"),
				RecipientName:            strLink("recipient name"),
				Phone:                    strLink("+79138739944"),
				TariffTypeCode:           intLink(10),
				DeliveryRecipientVATRate: strLink("VATY"),
			},
			args: args{
				deliveryRecipientVATRate: "VATX",
			},
			want: &OrderReq{
				Number:                   strLink("test_number"),
				RecipientName:            strLink("recipient name"),
				Phone:                    strLink("+79138739944"),
				TariffTypeCode:           intLink(10),
				DeliveryRecipientVATRate: strLink("VATX"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetDeliveryRecipientVATRate(tt.args.deliveryRecipientVATRate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetDeliveryRecipientVATRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetDeliveryRecipientVATSum(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		deliveryRecipientVATSum float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:                  strLink("test_number"),
				RecipientName:           strLink("recipient name"),
				Phone:                   strLink("+79138739944"),
				TariffTypeCode:          intLink(10),
				DeliveryRecipientVATSum: nil,
			},
			args: args{
				deliveryRecipientVATSum: 399.99,
			},
			want: &OrderReq{
				Number:                  strLink("test_number"),
				RecipientName:           strLink("recipient name"),
				Phone:                   strLink("+79138739944"),
				TariffTypeCode:          intLink(10),
				DeliveryRecipientVATSum: float64Link(399.99),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:                  strLink("test_number"),
				RecipientName:           strLink("recipient name"),
				Phone:                   strLink("+79138739944"),
				TariffTypeCode:          intLink(10),
				DeliveryRecipientVATSum: float64Link(400.),
			},
			args: args{
				deliveryRecipientVATSum: 399.99,
			},
			want: &OrderReq{
				Number:                  strLink("test_number"),
				RecipientName:           strLink("recipient name"),
				Phone:                   strLink("+79138739944"),
				TariffTypeCode:          intLink(10),
				DeliveryRecipientVATSum: float64Link(399.99),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetDeliveryRecipientVATSum(tt.args.deliveryRecipientVATSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetDeliveryRecipientVATSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetRecipientCurrency(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		recipientCurrency string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:            strLink("test_number"),
				RecipientName:     strLink("recipient name"),
				Phone:             strLink("+79138739944"),
				TariffTypeCode:    intLink(10),
				RecipientCurrency: nil,
			},
			args: args{
				recipientCurrency: "RUB",
			},
			want: &OrderReq{
				Number:            strLink("test_number"),
				RecipientName:     strLink("recipient name"),
				Phone:             strLink("+79138739944"),
				TariffTypeCode:    intLink(10),
				RecipientCurrency: strLink("RUB"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:            strLink("test_number"),
				RecipientName:     strLink("recipient name"),
				Phone:             strLink("+79138739944"),
				TariffTypeCode:    intLink(10),
				RecipientCurrency: strLink("EUR"),
			},
			args: args{
				recipientCurrency: "RUB",
			},
			want: &OrderReq{
				Number:            strLink("test_number"),
				RecipientName:     strLink("recipient name"),
				Phone:             strLink("+79138739944"),
				TariffTypeCode:    intLink(10),
				RecipientCurrency: strLink("RUB"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetRecipientCurrency(tt.args.recipientCurrency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetRecipientCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetItemsCurrency(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		itemsCurrency string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ItemsCurrency:  nil,
			},
			args: args{
				itemsCurrency: "RUB",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ItemsCurrency:  strLink("RUB"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ItemsCurrency:  strLink("EUR"),
			},
			args: args{
				itemsCurrency: "RUB",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				ItemsCurrency:  strLink("RUB"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetItemsCurrency(tt.args.itemsCurrency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetItemsCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetSeller(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		seller Seller
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Seller:         nil,
			},
			args: args{
				seller: Seller{
					Name: strLink("seller name"),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Seller: &Seller{
					Name: strLink("seller name"),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Seller: &Seller{
					Name: strLink("previous seller name"),
				},
			},
			args: args{
				seller: Seller{
					Name: strLink("seller name"),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Seller: &Seller{
					Name: strLink("seller name"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetSeller(tt.args.seller); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetSeller() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetComment(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		comment string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Comment:        nil,
			},
			args: args{
				comment: "comment",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Comment:        strLink("comment"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Comment:        strLink("previous comment"),
			},
			args: args{
				comment: "comment",
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Comment:        strLink("comment"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetComment(tt.args.comment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetAddress(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		address Address
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Address:        nil,
			},
			args: args{
				address: Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Address: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Address: &Address{
					Street: strLink("previous street"),
					House:  strLink("11/4"),
				},
			},
			args: args{
				address: Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Address: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetAddress(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_AddPackage(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		pack OrderPackage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "add first",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Package:        nil,
			},
			args: args{
				pack: OrderPackage{
					Number:  strLink("package number"),
					BarCode: strLink("package barcode"),
					Weight:  intLink(1500),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Package: []*OrderPackage{
					{
						Number:  strLink("package number"),
						BarCode: strLink("package barcode"),
						Weight:  intLink(1500),
					},
				},
			},
		},
		{
			name: "add second",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Package: []*OrderPackage{
					{
						Number:  strLink("package number"),
						BarCode: strLink("package barcode"),
						Weight:  intLink(1500),
					},
				},
			},
			args: args{
				pack: OrderPackage{
					Number:  strLink("package number 2"),
					BarCode: strLink("package barcode 2"),
					Weight:  intLink(3000),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Package: []*OrderPackage{
					{
						Number:  strLink("package number"),
						BarCode: strLink("package barcode"),
						Weight:  intLink(1500),
					},
					{
						Number:  strLink("package number 2"),
						BarCode: strLink("package barcode 2"),
						Weight:  intLink(3000),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.AddPackage(tt.args.pack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.AddPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetDeliveryRecipientCostAdv(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		deliveryRecipientCostAdv DeliveryRecipientCostAdv
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:                   strLink("test_number"),
				RecipientName:            strLink("recipient name"),
				Phone:                    strLink("+79138739944"),
				TariffTypeCode:           intLink(10),
				DeliveryRecipientCostAdv: nil,
			},
			args: args{
				deliveryRecipientCostAdv: DeliveryRecipientCostAdv{
					Threshold: intLink(999),
					Sum:       float64Link(333.99),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				DeliveryRecipientCostAdv: &DeliveryRecipientCostAdv{
					Threshold: intLink(999),
					Sum:       float64Link(333.99),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				DeliveryRecipientCostAdv: &DeliveryRecipientCostAdv{
					Threshold: intLink(777),
					Sum:       float64Link(666.55),
				},
			},
			args: args{
				deliveryRecipientCostAdv: DeliveryRecipientCostAdv{
					Threshold: intLink(999),
					Sum:       float64Link(333.99),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				DeliveryRecipientCostAdv: &DeliveryRecipientCostAdv{
					Threshold: intLink(999),
					Sum:       float64Link(333.99),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetDeliveryRecipientCostAdv(tt.args.deliveryRecipientCostAdv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetDeliveryRecipientCostAdv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetAddService(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		addService AddService
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				AddService:     nil,
			},
			args: args{
				addService: AddService{
					ServiceCode: intLink(783498),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				AddService: &AddService{
					ServiceCode: intLink(783498),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				AddService: &AddService{
					ServiceCode: intLink(857432),
				},
			},
			args: args{
				addService: AddService{
					ServiceCode: intLink(783498),
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				AddService: &AddService{
					ServiceCode: intLink(783498),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetAddService(tt.args.addService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetAddService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderReq_SetSchedule(t *testing.T) {
	type fields struct {
		Number                   *string
		SendCityCode             *int
		RecCityCode              *int
		SendCityPostCode         *int
		RecCityPostCode          *int
		SendCountryCode          *int
		RecCountryCode           *int
		SendCityName             *string
		RecCityName              *string
		RecipientINN             *string
		DateInvoice              *string
		ShipperName              *string
		ShipperAddress           *string
		Passport                 *Passport
		Sender                   *Sender
		RecipientName            *string
		RecipientEmail           *string
		Phone                    *string
		TariffTypeCode           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientCurrency        *string
		ItemsCurrency            *string
		Seller                   *Seller
		Comment                  *string
		Address                  *Address
		Package                  []*OrderPackage
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		AddService               *AddService
		Schedule                 *Schedule
	}
	type args struct {
		schedule Schedule
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderReq
	}{
		{
			name: "set",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Schedule:       nil,
			},
			args: args{
				schedule: Schedule{
					Attempt: []*ScheduleAttempt{
						{
							ID:   strLink("attempt id"),
							Date: strLink("2019-07-17"),
						},
					},
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Schedule: &Schedule{
					Attempt: []*ScheduleAttempt{
						{
							ID:   strLink("attempt id"),
							Date: strLink("2019-07-17"),
						},
					},
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Schedule: &Schedule{
					Attempt: []*ScheduleAttempt{
						{
							ID:   strLink("previous attempt id"),
							Date: strLink("2019-07-16"),
						},
					},
				},
			},
			args: args{
				schedule: Schedule{
					Attempt: []*ScheduleAttempt{
						{
							ID:   strLink("attempt id"),
							Date: strLink("2019-07-17"),
						},
					},
				},
			},
			want: &OrderReq{
				Number:         strLink("test_number"),
				RecipientName:  strLink("recipient name"),
				Phone:          strLink("+79138739944"),
				TariffTypeCode: intLink(10),
				Schedule: &Schedule{
					Attempt: []*ScheduleAttempt{
						{
							ID:   strLink("attempt id"),
							Date: strLink("2019-07-17"),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderReq := &OrderReq{
				Number:                   tt.fields.Number,
				SendCityCode:             tt.fields.SendCityCode,
				RecCityCode:              tt.fields.RecCityCode,
				SendCityPostCode:         tt.fields.SendCityPostCode,
				RecCityPostCode:          tt.fields.RecCityPostCode,
				SendCountryCode:          tt.fields.SendCountryCode,
				RecCountryCode:           tt.fields.RecCountryCode,
				SendCityName:             tt.fields.SendCityName,
				RecCityName:              tt.fields.RecCityName,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				ShipperName:              tt.fields.ShipperName,
				ShipperAddress:           tt.fields.ShipperAddress,
				Passport:                 tt.fields.Passport,
				Sender:                   tt.fields.Sender,
				RecipientName:            tt.fields.RecipientName,
				RecipientEmail:           tt.fields.RecipientEmail,
				Phone:                    tt.fields.Phone,
				TariffTypeCode:           tt.fields.TariffTypeCode,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientCurrency:        tt.fields.RecipientCurrency,
				ItemsCurrency:            tt.fields.ItemsCurrency,
				Seller:                   tt.fields.Seller,
				Comment:                  tt.fields.Comment,
				Address:                  tt.fields.Address,
				Package:                  tt.fields.Package,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				AddService:               tt.fields.AddService,
				Schedule:                 tt.fields.Schedule,
			}
			if got := orderReq.SetSchedule(tt.args.schedule); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderReq.SetSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPassport(t *testing.T) {
	tests := []struct {
		name string
		want *Passport
	}{
		{
			name: "constructor",
			want: &Passport{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPassport(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPassport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_SetSeries(t *testing.T) {
	type fields struct {
		Series    *string
		Number    *string
		IssueDate *string
		IssuedBy  *string
		DateBirth *string
	}
	type args struct {
		series string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Passport
	}{
		{
			name: "set",
			fields: fields{
				Series: nil,
			},
			args: args{
				series: "passport series",
			},
			want: &Passport{
				Series: strLink("passport series"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Series: strLink("previous passport series"),
			},
			args: args{
				series: "passport series",
			},
			want: &Passport{
				Series: strLink("passport series"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passport := &Passport{
				Series:    tt.fields.Series,
				Number:    tt.fields.Number,
				IssueDate: tt.fields.IssueDate,
				IssuedBy:  tt.fields.IssuedBy,
				DateBirth: tt.fields.DateBirth,
			}
			if got := passport.SetSeries(tt.args.series); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Passport.SetSeries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_SetNumber(t *testing.T) {
	type fields struct {
		Series    *string
		Number    *string
		IssueDate *string
		IssuedBy  *string
		DateBirth *string
	}
	type args struct {
		number string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Passport
	}{
		{
			name: "set",
			fields: fields{
				Number: nil,
			},
			args: args{
				number: "passport number",
			},
			want: &Passport{
				Number: strLink("passport number"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number: strLink("previous passport number"),
			},
			args: args{
				number: "passport number",
			},
			want: &Passport{
				Number: strLink("passport number"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passport := &Passport{
				Series:    tt.fields.Series,
				Number:    tt.fields.Number,
				IssueDate: tt.fields.IssueDate,
				IssuedBy:  tt.fields.IssuedBy,
				DateBirth: tt.fields.DateBirth,
			}
			if got := passport.SetNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Passport.SetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_SetIssueDate(t *testing.T) {
	type fields struct {
		Series    *string
		Number    *string
		IssueDate *string
		IssuedBy  *string
		DateBirth *string
	}
	type args struct {
		issueDate time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Passport
	}{
		{
			name: "set",
			fields: fields{
				IssueDate: nil,
			},
			args: args{
				issueDate: time.Date(2009, 7, 17, 0, 0, 0, 0, time.UTC),
			},
			want: &Passport{
				IssueDate: strLink("2009-07-17"),
			},
		},
		{
			name: "modify",
			fields: fields{
				IssueDate: strLink("2019-07-17"),
			},
			args: args{
				issueDate: time.Date(2009, 7, 17, 0, 0, 0, 0, time.UTC),
			},
			want: &Passport{
				IssueDate: strLink("2009-07-17"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passport := &Passport{
				Series:    tt.fields.Series,
				Number:    tt.fields.Number,
				IssueDate: tt.fields.IssueDate,
				IssuedBy:  tt.fields.IssuedBy,
				DateBirth: tt.fields.DateBirth,
			}
			if got := passport.SetIssueDate(tt.args.issueDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Passport.SetIssueDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_SetIssuedBy(t *testing.T) {
	type fields struct {
		Series    *string
		Number    *string
		IssueDate *string
		IssuedBy  *string
		DateBirth *string
	}
	type args struct {
		issuedBy string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Passport
	}{
		{
			name: "set",
			fields: fields{
				IssuedBy: nil,
			},
			args: args{
				issuedBy: "me",
			},
			want: &Passport{
				IssuedBy: strLink("me"),
			},
		},
		{
			name: "modify",
			fields: fields{
				IssuedBy: strLink("you"),
			},
			args: args{
				issuedBy: "me",
			},
			want: &Passport{
				IssuedBy: strLink("me"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passport := &Passport{
				Series:    tt.fields.Series,
				Number:    tt.fields.Number,
				IssueDate: tt.fields.IssueDate,
				IssuedBy:  tt.fields.IssuedBy,
				DateBirth: tt.fields.DateBirth,
			}
			if got := passport.SetIssuedBy(tt.args.issuedBy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Passport.SetIssuedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_SetDateBirth(t *testing.T) {
	type fields struct {
		Series    *string
		Number    *string
		IssueDate *string
		IssuedBy  *string
		DateBirth *string
	}
	type args struct {
		dateBirth time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Passport
	}{
		{
			name: "set",
			fields: fields{
				DateBirth: nil,
			},
			args: args{
				dateBirth: time.Date(1999, 7, 17, 0, 0, 0, 0, time.UTC),
			},
			want: &Passport{
				DateBirth: strLink("1999-07-17"),
			},
		},
		{
			name: "modify",
			fields: fields{
				DateBirth: strLink("2019-07-17"),
			},
			args: args{
				dateBirth: time.Date(1999, 7, 17, 0, 0, 0, 0, time.UTC),
			},
			want: &Passport{
				DateBirth: strLink("1999-07-17"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passport := &Passport{
				Series:    tt.fields.Series,
				Number:    tt.fields.Number,
				IssueDate: tt.fields.IssueDate,
				IssuedBy:  tt.fields.IssuedBy,
				DateBirth: tt.fields.DateBirth,
			}
			if got := passport.SetDateBirth(tt.args.dateBirth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Passport.SetDateBirth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSender(t *testing.T) {
	tests := []struct {
		name string
		want *Sender
	}{
		{
			name: "constructor",
			want: &Sender{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSender(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSender_SetCompany(t *testing.T) {
	type fields struct {
		Company *string
		Name    *string
		Address *Address
	}
	type args struct {
		company string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Sender
	}{
		{
			name: "set",
			fields: fields{
				Company: nil,
			},
			args: args{
				company: "sender company",
			},
			want: &Sender{
				Company: strLink("sender company"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Company: strLink("previous sender company"),
			},
			args: args{
				company: "sender company",
			},
			want: &Sender{
				Company: strLink("sender company"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sender := &Sender{
				Company: tt.fields.Company,
				Name:    tt.fields.Name,
				Address: tt.fields.Address,
			}
			if got := sender.SetCompany(tt.args.company); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sender.SetCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSender_SetName(t *testing.T) {
	type fields struct {
		Company *string
		Name    *string
		Address *Address
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Sender
	}{
		{
			name: "set",
			fields: fields{
				Name: nil,
			},
			args: args{
				name: "Bobby",
			},
			want: &Sender{
				Name: strLink("Bobby"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Name: strLink("Alex"),
			},
			args: args{
				name: "Bobby",
			},
			want: &Sender{
				Name: strLink("Bobby"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sender := &Sender{
				Company: tt.fields.Company,
				Name:    tt.fields.Name,
				Address: tt.fields.Address,
			}
			if got := sender.SetName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sender.SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSender_SetAddress(t *testing.T) {
	type fields struct {
		Company *string
		Name    *string
		Address *Address
	}
	type args struct {
		address Address
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Sender
	}{
		{
			name: "set",
			fields: fields{
				Address: nil,
			},
			args: args{
				address: Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
			want: &Sender{
				Address: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Address: &Address{
					Street: strLink("previous street"),
					House:  strLink("11/4"),
				},
			},
			args: args{
				address: Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
			want: &Sender{
				Address: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sender := &Sender{
				Company: tt.fields.Company,
				Name:    tt.fields.Name,
				Address: tt.fields.Address,
			}
			if got := sender.SetAddress(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sender.SetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAddress(t *testing.T) {
	type args struct {
		street string
		house  string
	}
	tests := []struct {
		name string
		args args
		want *Address
	}{
		{
			name: "constructor",
			args: args{
				street: "test street",
				house:  "10/3",
			},
			want: &Address{
				Street: strLink("test street"),
				House:  strLink("10/3"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddress(tt.args.street, tt.args.house); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetFlat(t *testing.T) {
	type fields struct {
		Street  *string
		House   *string
		Flat    *string
		Phone   *string
		PvzCode *string
	}
	type args struct {
		flat string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Address
	}{
		{
			name: "set",
			fields: fields{
				Street: strLink("test street"),
				House:  strLink("10/3"),
				Flat:   nil,
			},
			args: args{
				flat: "12b",
			},
			want: &Address{
				Street: strLink("test street"),
				House:  strLink("10/3"),
				Flat:   strLink("12b"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Street: strLink("test street"),
				House:  strLink("10/3"),
				Flat:   strLink("12a"),
			},
			args: args{
				flat: "12b",
			},
			want: &Address{
				Street: strLink("test street"),
				House:  strLink("10/3"),
				Flat:   strLink("12b"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address := &Address{
				Street:  tt.fields.Street,
				House:   tt.fields.House,
				Flat:    tt.fields.Flat,
				Phone:   tt.fields.Phone,
				PvzCode: tt.fields.PvzCode,
			}
			if got := address.SetFlat(tt.args.flat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetFlat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetPhone(t *testing.T) {
	type fields struct {
		Street  *string
		House   *string
		Flat    *string
		Phone   *string
		PvzCode *string
	}
	type args struct {
		phone string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Address
	}{
		{
			name: "set",
			fields: fields{
				Street: strLink("test street"),
				House:  strLink("10/3"),
				Phone:  nil,
			},
			args: args{
				phone: "+79138739944",
			},
			want: &Address{
				Street: strLink("test street"),
				House:  strLink("10/3"),
				Phone:  strLink("+79138739944"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Street: strLink("test street"),
				House:  strLink("10/3"),
				Phone:  strLink("+89138739999"),
			},
			args: args{
				phone: "+79138739944",
			},
			want: &Address{
				Street: strLink("test street"),
				House:  strLink("10/3"),
				Phone:  strLink("+79138739944"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address := &Address{
				Street:  tt.fields.Street,
				House:   tt.fields.House,
				Flat:    tt.fields.Flat,
				Phone:   tt.fields.Phone,
				PvzCode: tt.fields.PvzCode,
			}
			if got := address.SetPhone(tt.args.phone); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetPvzCode(t *testing.T) {
	type fields struct {
		Street  *string
		House   *string
		Flat    *string
		Phone   *string
		PvzCode *string
	}
	type args struct {
		pvzCode string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Address
	}{
		{
			name: "set",
			fields: fields{
				Street:  strLink("test street"),
				House:   strLink("10/3"),
				PvzCode: nil,
			},
			args: args{
				pvzCode: "NSK333",
			},
			want: &Address{
				Street:  strLink("test street"),
				House:   strLink("10/3"),
				PvzCode: strLink("NSK333"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Street:  strLink("test street"),
				House:   strLink("10/3"),
				PvzCode: strLink("NSK777"),
			},
			args: args{
				pvzCode: "NSK333",
			},
			want: &Address{
				Street:  strLink("test street"),
				House:   strLink("10/3"),
				PvzCode: strLink("NSK333"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address := &Address{
				Street:  tt.fields.Street,
				House:   tt.fields.House,
				Flat:    tt.fields.Flat,
				Phone:   tt.fields.Phone,
				PvzCode: tt.fields.PvzCode,
			}
			if got := address.SetPvzCode(tt.args.pvzCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetPvzCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSeller(t *testing.T) {
	tests := []struct {
		name string
		want *Seller
	}{
		{
			name: "constructor",
			want: &Seller{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSeller(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeller() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeller_SetAddress(t *testing.T) {
	type fields struct {
		Address       *string
		Name          *string
		INN           *string
		Phone         *string
		OwnershipForm *int
	}
	type args struct {
		address string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Seller
	}{
		{
			name: "set",
			fields: fields{
				Address: nil,
			},
			args: args{
				address: "test address",
			},
			want: &Seller{
				Address: strLink("test address"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Address: strLink("previous address"),
			},
			args: args{
				address: "test address",
			},
			want: &Seller{
				Address: strLink("test address"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seller := &Seller{
				Address:       tt.fields.Address,
				Name:          tt.fields.Name,
				INN:           tt.fields.INN,
				Phone:         tt.fields.Phone,
				OwnershipForm: tt.fields.OwnershipForm,
			}
			if got := seller.SetAddress(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seller.SetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeller_SetName(t *testing.T) {
	type fields struct {
		Address       *string
		Name          *string
		INN           *string
		Phone         *string
		OwnershipForm *int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Seller
	}{
		{
			name: "set",
			fields: fields{
				Name: nil,
			},
			args: args{
				name: "test name",
			},
			want: &Seller{
				Name: strLink("test name"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Name: strLink("previous name"),
			},
			args: args{
				name: "test name",
			},
			want: &Seller{
				Name: strLink("test name"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seller := &Seller{
				Address:       tt.fields.Address,
				Name:          tt.fields.Name,
				INN:           tt.fields.INN,
				Phone:         tt.fields.Phone,
				OwnershipForm: tt.fields.OwnershipForm,
			}
			if got := seller.SetName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seller.SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeller_SetINN(t *testing.T) {
	type fields struct {
		Address       *string
		Name          *string
		INN           *string
		Phone         *string
		OwnershipForm *int
	}
	type args struct {
		inn string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Seller
	}{
		{
			name: "set",
			fields: fields{
				INN: nil,
			},
			args: args{
				inn: "12345678987",
			},
			want: &Seller{
				INN: strLink("12345678987"),
			},
		},
		{
			name: "modify",
			fields: fields{
				INN: strLink("98765432123"),
			},
			args: args{
				inn: "12345678987",
			},
			want: &Seller{
				INN: strLink("12345678987"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seller := &Seller{
				Address:       tt.fields.Address,
				Name:          tt.fields.Name,
				INN:           tt.fields.INN,
				Phone:         tt.fields.Phone,
				OwnershipForm: tt.fields.OwnershipForm,
			}
			if got := seller.SetINN(tt.args.inn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seller.SetINN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeller_SetPhone(t *testing.T) {
	type fields struct {
		Address       *string
		Name          *string
		INN           *string
		Phone         *string
		OwnershipForm *int
	}
	type args struct {
		phone string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Seller
	}{
		{
			name: "set",
			fields: fields{
				Phone: nil,
			},
			args: args{
				phone: "+79138739944",
			},
			want: &Seller{
				Phone: strLink("+79138739944"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Phone: strLink("+89138739999"),
			},
			args: args{
				phone: "+79138739944",
			},
			want: &Seller{
				Phone: strLink("+79138739944"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seller := &Seller{
				Address:       tt.fields.Address,
				Name:          tt.fields.Name,
				INN:           tt.fields.INN,
				Phone:         tt.fields.Phone,
				OwnershipForm: tt.fields.OwnershipForm,
			}
			if got := seller.SetPhone(tt.args.phone); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seller.SetPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeller_SetOwnershipForm(t *testing.T) {
	type fields struct {
		Address       *string
		Name          *string
		INN           *string
		Phone         *string
		OwnershipForm *int
	}
	type args struct {
		ownershipForm int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Seller
	}{
		{
			name: "set",
			fields: fields{
				OwnershipForm: nil,
			},
			args: args{
				ownershipForm: 1,
			},
			want: &Seller{
				OwnershipForm: intLink(1),
			},
		},
		{
			name: "modify",
			fields: fields{
				OwnershipForm: intLink(2),
			},
			args: args{
				ownershipForm: 1,
			},
			want: &Seller{
				OwnershipForm: intLink(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seller := &Seller{
				Address:       tt.fields.Address,
				Name:          tt.fields.Name,
				INN:           tt.fields.INN,
				Phone:         tt.fields.Phone,
				OwnershipForm: tt.fields.OwnershipForm,
			}
			if got := seller.SetOwnershipForm(tt.args.ownershipForm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seller.SetOwnershipForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOrderPackage(t *testing.T) {
	type args struct {
		number  string
		barCode string
		weight  int
	}
	tests := []struct {
		name string
		args args
		want *OrderPackage
	}{
		{
			name: "constructor",
			args: args{
				number:  "test number",
				barCode: "test barcode",
				weight:  4500,
			},
			want: &OrderPackage{
				Number:  strLink("test number"),
				BarCode: strLink("test barcode"),
				Weight:  intLink(4500),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderPackage(tt.args.number, tt.args.barCode, tt.args.weight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderPackage_SetSize(t *testing.T) {
	type fields struct {
		Number  *string
		BarCode *string
		Weight  *int
		SizeA   *int
		SizeB   *int
		SizeC   *int
		Item    []*OrderPackageItem
	}
	type args struct {
		length int
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderPackage
	}{
		{
			name: "set",
			fields: fields{
				Number:  strLink("test number"),
				BarCode: strLink("test barcode"),
				Weight:  intLink(4500),
				SizeA:   nil,
				SizeB:   nil,
				SizeC:   nil,
			},
			args: args{
				length: 50,
				width:  30,
				height: 20,
			},
			want: &OrderPackage{
				Number:  strLink("test number"),
				BarCode: strLink("test barcode"),
				Weight:  intLink(4500),
				SizeA:   intLink(50),
				SizeB:   intLink(30),
				SizeC:   intLink(20),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number:  strLink("test number"),
				BarCode: strLink("test barcode"),
				Weight:  intLink(4500),
				SizeA:   intLink(100),
				SizeB:   intLink(60),
				SizeC:   intLink(40),
			},
			args: args{
				length: 50,
				width:  30,
				height: 20,
			},
			want: &OrderPackage{
				Number:  strLink("test number"),
				BarCode: strLink("test barcode"),
				Weight:  intLink(4500),
				SizeA:   intLink(50),
				SizeB:   intLink(30),
				SizeC:   intLink(20),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderPackage := &OrderPackage{
				Number:  tt.fields.Number,
				BarCode: tt.fields.BarCode,
				Weight:  tt.fields.Weight,
				SizeA:   tt.fields.SizeA,
				SizeB:   tt.fields.SizeB,
				SizeC:   tt.fields.SizeC,
				Item:    tt.fields.Item,
			}
			if got := orderPackage.SetSize(tt.args.length, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderPackage.SetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderPackage_AddItem(t *testing.T) {
	type fields struct {
		Number  *string
		BarCode *string
		Weight  *int
		SizeA   *int
		SizeB   *int
		SizeC   *int
		Item    []*OrderPackageItem
	}
	type args struct {
		item OrderPackageItem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderPackage
	}{
		{
			name: "add first",
			fields: fields{
				Number:  strLink("test number"),
				BarCode: strLink("test barcode"),
				Weight:  intLink(4500),
				Item:    nil,
			},
			args: args{
				item: OrderPackageItem{
					Amount:  intLink(1),
					WareKey: strLink("00012345"),
					Cost:    float64Link(500),
					Payment: float64Link(500),
					Weight:  intLink(2),
					Comment: strLink("wrap"),
				},
			},
			want: &OrderPackage{
				Number:  strLink("test number"),
				BarCode: strLink("test barcode"),
				Weight:  intLink(4500),
				Item: []*OrderPackageItem{
					{
						Amount:  intLink(1),
						WareKey: strLink("00012345"),
						Cost:    float64Link(500),
						Payment: float64Link(500),
						Weight:  intLink(2),
						Comment: strLink("wrap"),
					},
				},
			},
		},
		{
			name: "add second",
			fields: fields{
				Number:  strLink("test number"),
				BarCode: strLink("test barcode"),
				Weight:  intLink(4500),
				Item: []*OrderPackageItem{
					{
						Amount:  intLink(1),
						WareKey: strLink("00012345"),
						Cost:    float64Link(500),
						Payment: float64Link(500),
						Weight:  intLink(2),
						Comment: strLink("wrap"),
					},
				},
			},
			args: args{
				item: OrderPackageItem{
					Amount:  intLink(1),
					WareKey: strLink("00078945"),
					Cost:    float64Link(5500),
					Payment: float64Link(5500),
					Weight:  intLink(3),
					Comment: strLink("female jacket"),
				},
			},
			want: &OrderPackage{
				Number:  strLink("test number"),
				BarCode: strLink("test barcode"),
				Weight:  intLink(4500),
				Item: []*OrderPackageItem{
					{
						Amount:  intLink(1),
						WareKey: strLink("00012345"),
						Cost:    float64Link(500),
						Payment: float64Link(500),
						Weight:  intLink(2),
						Comment: strLink("wrap"),
					},
					{
						Amount:  intLink(1),
						WareKey: strLink("00078945"),
						Cost:    float64Link(5500),
						Payment: float64Link(5500),
						Weight:  intLink(3),
						Comment: strLink("female jacket"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderPackage := &OrderPackage{
				Number:  tt.fields.Number,
				BarCode: tt.fields.BarCode,
				Weight:  tt.fields.Weight,
				SizeA:   tt.fields.SizeA,
				SizeB:   tt.fields.SizeB,
				SizeC:   tt.fields.SizeC,
				Item:    tt.fields.Item,
			}
			if got := orderPackage.AddItem(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderPackage.AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOrderPackageItem(t *testing.T) {
	type args struct {
		amount  int
		wareKey string
		cost    float64
		payment float64
		weight  int
		comment string
	}
	tests := []struct {
		name string
		args args
		want *OrderPackageItem
	}{
		{
			name: "constructor",
			args: args{
				amount:  1,
				wareKey: "00012345",
				cost:    500,
				payment: 500,
				weight:  2,
				comment: "wrap",
			},
			want: &OrderPackageItem{
				Amount:  intLink(1),
				WareKey: strLink("00012345"),
				Cost:    float64Link(500),
				Payment: float64Link(500),
				Weight:  intLink(2),
				Comment: strLink("wrap"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOrderPackageItem(
				tt.args.amount,
				tt.args.wareKey,
				tt.args.cost,
				tt.args.payment,
				tt.args.weight,
				tt.args.comment,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderPackageItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderPackageItem_SetPaymentVATRate(t *testing.T) {
	type fields struct {
		Amount         *int
		WareKey        *string
		Cost           *float64
		Payment        *float64
		PaymentVATRate *string
		PaymentVATSum  *float64
		Weight         *int
		Comment        *string
		WeightBrutto   *int
		CommentEx      *string
		Link           *string
	}
	type args struct {
		paymentVATRate string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderPackageItem
	}{
		{
			name: "set",
			fields: fields{
				Amount:         intLink(1),
				WareKey:        strLink("00012345"),
				Cost:           float64Link(500),
				Payment:        float64Link(500),
				Weight:         intLink(2),
				Comment:        strLink("wrap"),
				PaymentVATRate: nil,
			},
			args: args{
				paymentVATRate: "VATX",
			},
			want: &OrderPackageItem{
				Amount:         intLink(1),
				WareKey:        strLink("00012345"),
				Cost:           float64Link(500),
				Payment:        float64Link(500),
				Weight:         intLink(2),
				Comment:        strLink("wrap"),
				PaymentVATRate: strLink("VATX"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Amount:         intLink(1),
				WareKey:        strLink("00012345"),
				Cost:           float64Link(500),
				Payment:        float64Link(500),
				Weight:         intLink(2),
				Comment:        strLink("wrap"),
				PaymentVATRate: strLink("VATY"),
			},
			args: args{
				paymentVATRate: "VATX",
			},
			want: &OrderPackageItem{
				Amount:         intLink(1),
				WareKey:        strLink("00012345"),
				Cost:           float64Link(500),
				Payment:        float64Link(500),
				Weight:         intLink(2),
				Comment:        strLink("wrap"),
				PaymentVATRate: strLink("VATX"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &OrderPackageItem{
				Amount:         tt.fields.Amount,
				WareKey:        tt.fields.WareKey,
				Cost:           tt.fields.Cost,
				Payment:        tt.fields.Payment,
				PaymentVATRate: tt.fields.PaymentVATRate,
				PaymentVATSum:  tt.fields.PaymentVATSum,
				Weight:         tt.fields.Weight,
				Comment:        tt.fields.Comment,
				WeightBrutto:   tt.fields.WeightBrutto,
				CommentEx:      tt.fields.CommentEx,
				Link:           tt.fields.Link,
			}
			if got := item.SetPaymentVATRate(tt.args.paymentVATRate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderPackageItem.SetPaymentVATRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderPackageItem_SetPaymentVATSum(t *testing.T) {
	type fields struct {
		Amount         *int
		WareKey        *string
		Cost           *float64
		Payment        *float64
		PaymentVATRate *string
		PaymentVATSum  *float64
		Weight         *int
		Comment        *string
		WeightBrutto   *int
		CommentEx      *string
		Link           *string
	}
	type args struct {
		paymentVATSum float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderPackageItem
	}{
		{
			name: "set",
			fields: fields{
				Amount:        intLink(1),
				WareKey:       strLink("00012345"),
				Cost:          float64Link(500),
				Payment:       float64Link(500),
				Weight:        intLink(2),
				Comment:       strLink("wrap"),
				PaymentVATSum: nil,
			},
			args: args{
				paymentVATSum: 65.8,
			},
			want: &OrderPackageItem{
				Amount:        intLink(1),
				WareKey:       strLink("00012345"),
				Cost:          float64Link(500),
				Payment:       float64Link(500),
				Weight:        intLink(2),
				Comment:       strLink("wrap"),
				PaymentVATSum: float64Link(65.8),
			},
		},
		{
			name: "modify",
			fields: fields{
				Amount:        intLink(1),
				WareKey:       strLink("00012345"),
				Cost:          float64Link(500),
				Payment:       float64Link(500),
				Weight:        intLink(2),
				Comment:       strLink("wrap"),
				PaymentVATSum: float64Link(35.1),
			},
			args: args{
				paymentVATSum: 65.8,
			},
			want: &OrderPackageItem{
				Amount:        intLink(1),
				WareKey:       strLink("00012345"),
				Cost:          float64Link(500),
				Payment:       float64Link(500),
				Weight:        intLink(2),
				Comment:       strLink("wrap"),
				PaymentVATSum: float64Link(65.8),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &OrderPackageItem{
				Amount:         tt.fields.Amount,
				WareKey:        tt.fields.WareKey,
				Cost:           tt.fields.Cost,
				Payment:        tt.fields.Payment,
				PaymentVATRate: tt.fields.PaymentVATRate,
				PaymentVATSum:  tt.fields.PaymentVATSum,
				Weight:         tt.fields.Weight,
				Comment:        tt.fields.Comment,
				WeightBrutto:   tt.fields.WeightBrutto,
				CommentEx:      tt.fields.CommentEx,
				Link:           tt.fields.Link,
			}
			if got := item.SetPaymentVATSum(tt.args.paymentVATSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderPackageItem.SetPaymentVATSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderPackageItem_SetWeightBrutto(t *testing.T) {
	type fields struct {
		Amount         *int
		WareKey        *string
		Cost           *float64
		Payment        *float64
		PaymentVATRate *string
		PaymentVATSum  *float64
		Weight         *int
		Comment        *string
		WeightBrutto   *int
		CommentEx      *string
		Link           *string
	}
	type args struct {
		weightBrutto int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderPackageItem
	}{
		{
			name: "set",
			fields: fields{
				Amount:       intLink(1),
				WareKey:      strLink("00012345"),
				Cost:         float64Link(500),
				Payment:      float64Link(500),
				Weight:       intLink(2),
				Comment:      strLink("wrap"),
				WeightBrutto: nil,
			},
			args: args{
				weightBrutto: 1500,
			},
			want: &OrderPackageItem{
				Amount:       intLink(1),
				WareKey:      strLink("00012345"),
				Cost:         float64Link(500),
				Payment:      float64Link(500),
				Weight:       intLink(2),
				Comment:      strLink("wrap"),
				WeightBrutto: intLink(1500),
			},
		},
		{
			name: "modify",
			fields: fields{
				Amount:       intLink(1),
				WareKey:      strLink("00012345"),
				Cost:         float64Link(500),
				Payment:      float64Link(500),
				Weight:       intLink(2),
				Comment:      strLink("wrap"),
				WeightBrutto: intLink(3000),
			},
			args: args{
				weightBrutto: 1500,
			},
			want: &OrderPackageItem{
				Amount:       intLink(1),
				WareKey:      strLink("00012345"),
				Cost:         float64Link(500),
				Payment:      float64Link(500),
				Weight:       intLink(2),
				Comment:      strLink("wrap"),
				WeightBrutto: intLink(1500),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &OrderPackageItem{
				Amount:         tt.fields.Amount,
				WareKey:        tt.fields.WareKey,
				Cost:           tt.fields.Cost,
				Payment:        tt.fields.Payment,
				PaymentVATRate: tt.fields.PaymentVATRate,
				PaymentVATSum:  tt.fields.PaymentVATSum,
				Weight:         tt.fields.Weight,
				Comment:        tt.fields.Comment,
				WeightBrutto:   tt.fields.WeightBrutto,
				CommentEx:      tt.fields.CommentEx,
				Link:           tt.fields.Link,
			}
			if got := item.SetWeightBrutto(tt.args.weightBrutto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderPackageItem.SetWeightBrutto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderPackageItem_SetCommentEx(t *testing.T) {
	type fields struct {
		Amount         *int
		WareKey        *string
		Cost           *float64
		Payment        *float64
		PaymentVATRate *string
		PaymentVATSum  *float64
		Weight         *int
		Comment        *string
		WeightBrutto   *int
		CommentEx      *string
		Link           *string
	}
	type args struct {
		commentEx string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderPackageItem
	}{
		{
			name: "set",
			fields: fields{
				Amount:    intLink(1),
				WareKey:   strLink("00012345"),
				Cost:      float64Link(500),
				Payment:   float64Link(500),
				Weight:    intLink(2),
				Comment:   strLink("wrap"),
				CommentEx: nil,
			},
			args: args{
				commentEx: "red hat",
			},
			want: &OrderPackageItem{
				Amount:    intLink(1),
				WareKey:   strLink("00012345"),
				Cost:      float64Link(500),
				Payment:   float64Link(500),
				Weight:    intLink(2),
				Comment:   strLink("wrap"),
				CommentEx: strLink("red hat"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Amount:    intLink(1),
				WareKey:   strLink("00012345"),
				Cost:      float64Link(500),
				Payment:   float64Link(500),
				Weight:    intLink(2),
				Comment:   strLink("wrap"),
				CommentEx: strLink("blue hat"),
			},
			args: args{
				commentEx: "red hat",
			},
			want: &OrderPackageItem{
				Amount:    intLink(1),
				WareKey:   strLink("00012345"),
				Cost:      float64Link(500),
				Payment:   float64Link(500),
				Weight:    intLink(2),
				Comment:   strLink("wrap"),
				CommentEx: strLink("red hat"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &OrderPackageItem{
				Amount:         tt.fields.Amount,
				WareKey:        tt.fields.WareKey,
				Cost:           tt.fields.Cost,
				Payment:        tt.fields.Payment,
				PaymentVATRate: tt.fields.PaymentVATRate,
				PaymentVATSum:  tt.fields.PaymentVATSum,
				Weight:         tt.fields.Weight,
				Comment:        tt.fields.Comment,
				WeightBrutto:   tt.fields.WeightBrutto,
				CommentEx:      tt.fields.CommentEx,
				Link:           tt.fields.Link,
			}
			if got := item.SetCommentEx(tt.args.commentEx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderPackageItem.SetCommentEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderPackageItem_SetLink(t *testing.T) {
	type fields struct {
		Amount         *int
		WareKey        *string
		Cost           *float64
		Payment        *float64
		PaymentVATRate *string
		PaymentVATSum  *float64
		Weight         *int
		Comment        *string
		WeightBrutto   *int
		CommentEx      *string
		Link           *string
	}
	type args struct {
		link string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OrderPackageItem
	}{
		{
			name: "set",
			fields: fields{
				Amount:  intLink(1),
				WareKey: strLink("00012345"),
				Cost:    float64Link(500),
				Payment: float64Link(500),
				Weight:  intLink(2),
				Comment: strLink("wrap"),
				Link:    nil,
			},
			args: args{
				link: "some://link",
			},
			want: &OrderPackageItem{
				Amount:  intLink(1),
				WareKey: strLink("00012345"),
				Cost:    float64Link(500),
				Payment: float64Link(500),
				Weight:  intLink(2),
				Comment: strLink("wrap"),
				Link:    strLink("some://link"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Amount:  intLink(1),
				WareKey: strLink("00012345"),
				Cost:    float64Link(500),
				Payment: float64Link(500),
				Weight:  intLink(2),
				Comment: strLink("wrap"),
				Link:    strLink("other://link"),
			},
			args: args{
				link: "some://link",
			},
			want: &OrderPackageItem{
				Amount:  intLink(1),
				WareKey: strLink("00012345"),
				Cost:    float64Link(500),
				Payment: float64Link(500),
				Weight:  intLink(2),
				Comment: strLink("wrap"),
				Link:    strLink("some://link"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &OrderPackageItem{
				Amount:         tt.fields.Amount,
				WareKey:        tt.fields.WareKey,
				Cost:           tt.fields.Cost,
				Payment:        tt.fields.Payment,
				PaymentVATRate: tt.fields.PaymentVATRate,
				PaymentVATSum:  tt.fields.PaymentVATSum,
				Weight:         tt.fields.Weight,
				Comment:        tt.fields.Comment,
				WeightBrutto:   tt.fields.WeightBrutto,
				CommentEx:      tt.fields.CommentEx,
				Link:           tt.fields.Link,
			}
			if got := item.SetLink(tt.args.link); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderPackageItem.SetLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeliveryRecipientCostAdv(t *testing.T) {
	type args struct {
		threshold int
		sum       float64
	}
	tests := []struct {
		name string
		args args
		want *DeliveryRecipientCostAdv
	}{
		{
			name: "constructor",
			args: args{
				threshold: 999,
				sum:       333.99,
			},
			want: &DeliveryRecipientCostAdv{
				Threshold: intLink(999),
				Sum:       float64Link(333.99),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeliveryRecipientCostAdv(tt.args.threshold, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeliveryRecipientCostAdv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeliveryRecipientCostAdv_SetVATRate(t *testing.T) {
	type fields struct {
		Threshold *int
		Sum       *float64
		VATRate   *string
		VATSum    *float64
	}
	type args struct {
		vatRate string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeliveryRecipientCostAdv
	}{
		{
			name: "set",
			fields: fields{
				Threshold: intLink(999),
				Sum:       float64Link(333.99),
				VATRate:   nil,
			},
			args: args{
				vatRate: "VATX",
			},
			want: &DeliveryRecipientCostAdv{
				Threshold: intLink(999),
				Sum:       float64Link(333.99),
				VATRate:   strLink("VATX"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Threshold: intLink(999),
				Sum:       float64Link(333.99),
				VATRate:   strLink("VATY"),
			},
			args: args{
				vatRate: "VATX",
			},
			want: &DeliveryRecipientCostAdv{
				Threshold: intLink(999),
				Sum:       float64Link(333.99),
				VATRate:   strLink("VATX"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DeliveryRecipientCostAdv{
				Threshold: tt.fields.Threshold,
				Sum:       tt.fields.Sum,
				VATRate:   tt.fields.VATRate,
				VATSum:    tt.fields.VATSum,
			}
			if got := d.SetVATRate(tt.args.vatRate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeliveryRecipientCostAdv.SetVATRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeliveryRecipientCostAdv_SetVATSum(t *testing.T) {
	type fields struct {
		Threshold *int
		Sum       *float64
		VATRate   *string
		VATSum    *float64
	}
	type args struct {
		vatSum float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeliveryRecipientCostAdv
	}{
		{
			name: "set",
			fields: fields{
				Threshold: intLink(999),
				Sum:       float64Link(333.99),
				VATSum:    nil,
			},
			args: args{
				vatSum: 54.3,
			},
			want: &DeliveryRecipientCostAdv{
				Threshold: intLink(999),
				Sum:       float64Link(333.99),
				VATSum:    float64Link(54.3),
			},
		},
		{
			name: "modify",
			fields: fields{
				Threshold: intLink(999),
				Sum:       float64Link(333.99),
				VATSum:    float64Link(32.1),
			},
			args: args{
				vatSum: 54.3,
			},
			want: &DeliveryRecipientCostAdv{
				Threshold: intLink(999),
				Sum:       float64Link(333.99),
				VATSum:    float64Link(54.3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DeliveryRecipientCostAdv{
				Threshold: tt.fields.Threshold,
				Sum:       tt.fields.Sum,
				VATRate:   tt.fields.VATRate,
				VATSum:    tt.fields.VATSum,
			}
			if got := d.SetVATSum(tt.args.vatSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeliveryRecipientCostAdv.SetVATSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAddService(t *testing.T) {
	type args struct {
		serviceCode int
	}
	tests := []struct {
		name string
		args args
		want *AddService
	}{
		{
			name: "constructor",
			args: args{
				serviceCode: 123,
			},
			want: &AddService{
				ServiceCode: intLink(123),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddService(tt.args.serviceCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddService_SetCount(t *testing.T) {
	type fields struct {
		ServiceCode *int
		Count       *int
	}
	type args struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddService
	}{
		{
			name: "add",
			fields: fields{
				ServiceCode: intLink(123),
				Count:       nil,
			},
			args: args{
				count: 1,
			},
			want: &AddService{
				ServiceCode: intLink(123),
				Count:       intLink(1),
			},
		},
		{
			name: "modify",
			fields: fields{
				ServiceCode: intLink(123),
				Count:       intLink(3),
			},
			args: args{
				count: 1,
			},
			want: &AddService{
				ServiceCode: intLink(123),
				Count:       intLink(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addService := &AddService{
				ServiceCode: tt.fields.ServiceCode,
				Count:       tt.fields.Count,
			}
			if got := addService.SetCount(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddService.SetCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSchedule(t *testing.T) {
	tests := []struct {
		name string
		want *Schedule
	}{
		{
			name: "constructor",
			want: &Schedule{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSchedule(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchedule_AddAttempt(t *testing.T) {
	type fields struct {
		Attempt []*ScheduleAttempt
	}
	type args struct {
		attempt ScheduleAttempt
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schedule
	}{
		{
			name: "add first",
			fields: fields{
				Attempt: nil,
			},
			args: args{
				attempt: ScheduleAttempt{
					ID:   strLink("attempt id"),
					Date: strLink("2019-07-17"),
				},
			},
			want: &Schedule{
				Attempt: []*ScheduleAttempt{
					{
						ID:   strLink("attempt id"),
						Date: strLink("2019-07-17"),
					},
				},
			},
		},
		{
			name: "add second",
			fields: fields{
				Attempt: []*ScheduleAttempt{
					{
						ID:   strLink("attempt id"),
						Date: strLink("2019-07-17"),
					},
				},
			},
			args: args{
				attempt: ScheduleAttempt{
					ID:   strLink("attempt id 2"),
					Date: strLink("2019-07-20"),
				},
			},
			want: &Schedule{
				Attempt: []*ScheduleAttempt{
					{
						ID:   strLink("attempt id"),
						Date: strLink("2019-07-17"),
					},
					{
						ID:   strLink("attempt id 2"),
						Date: strLink("2019-07-20"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schedule := &Schedule{
				Attempt: tt.fields.Attempt,
			}
			if got := schedule.AddAttempt(tt.args.attempt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schedule.AddAttempt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewScheduleAttempt(t *testing.T) {
	type args struct {
		id   string
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want *ScheduleAttempt
	}{
		{
			name: "constructor",
			args: args{
				id:   "attempt_id",
				date: time.Date(2019, 7, 17, 0, 0, 0, 0, time.UTC),
			},
			want: &ScheduleAttempt{
				ID:   strLink("attempt_id"),
				Date: strLink("2019-07-17"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScheduleAttempt(tt.args.id, tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScheduleAttempt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduleAttempt_SetComment(t *testing.T) {
	type fields struct {
		ID      *string
		Date    *string
		Comment *string
		TimeBeg *string
		TimeEnd *string
		Address *Address
	}
	type args struct {
		comment string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ScheduleAttempt
	}{
		{
			name: "set",
			fields: fields{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				Comment: nil,
			},
			args: args{
				comment: "comment text",
			},
			want: &ScheduleAttempt{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				Comment: strLink("comment text"),
			},
		},
		{
			name: "modify",
			fields: fields{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				Comment: strLink("previous comment text"),
			},
			args: args{
				comment: "comment text",
			},
			want: &ScheduleAttempt{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				Comment: strLink("comment text"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduleAttempt := &ScheduleAttempt{
				ID:      tt.fields.ID,
				Date:    tt.fields.Date,
				Comment: tt.fields.Comment,
				TimeBeg: tt.fields.TimeBeg,
				TimeEnd: tt.fields.TimeEnd,
				Address: tt.fields.Address,
			}
			if got := scheduleAttempt.SetComment(tt.args.comment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScheduleAttempt.SetComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduleAttempt_SetTimeBeg(t *testing.T) {
	type fields struct {
		ID      *string
		Date    *string
		Comment *string
		TimeBeg *string
		TimeEnd *string
		Address *Address
	}
	type args struct {
		timeBeg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ScheduleAttempt
	}{
		{
			name: "set",
			fields: fields{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				TimeBeg: nil,
			},
			args: args{
				timeBeg: "20:14",
			},
			want: &ScheduleAttempt{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				TimeBeg: strLink("20:14"),
			},
		},
		{
			name: "modify",
			fields: fields{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				TimeBeg: strLink("19:10"),
			},
			args: args{
				timeBeg: "20:14",
			},
			want: &ScheduleAttempt{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				TimeBeg: strLink("20:14"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduleAttempt := &ScheduleAttempt{
				ID:      tt.fields.ID,
				Date:    tt.fields.Date,
				Comment: tt.fields.Comment,
				TimeBeg: tt.fields.TimeBeg,
				TimeEnd: tt.fields.TimeEnd,
				Address: tt.fields.Address,
			}
			if got := scheduleAttempt.SetTimeBeg(tt.args.timeBeg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScheduleAttempt.SetTimeBeg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduleAttempt_SetTimeEnd(t *testing.T) {
	type fields struct {
		ID      *string
		Date    *string
		Comment *string
		TimeBeg *string
		TimeEnd *string
		Address *Address
	}
	type args struct {
		timeEnd string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ScheduleAttempt
	}{
		{
			name: "set",
			fields: fields{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				TimeEnd: nil,
			},
			args: args{
				timeEnd: "20:14",
			},
			want: &ScheduleAttempt{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				TimeEnd: strLink("20:14"),
			},
		},
		{
			name: "modify",
			fields: fields{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				TimeEnd: strLink("19:10"),
			},
			args: args{
				timeEnd: "20:14",
			},
			want: &ScheduleAttempt{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				TimeEnd: strLink("20:14"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduleAttempt := &ScheduleAttempt{
				ID:      tt.fields.ID,
				Date:    tt.fields.Date,
				Comment: tt.fields.Comment,
				TimeBeg: tt.fields.TimeBeg,
				TimeEnd: tt.fields.TimeEnd,
				Address: tt.fields.Address,
			}
			if got := scheduleAttempt.SetTimeEnd(tt.args.timeEnd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScheduleAttempt.SetTimeEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduleAttempt_SetAddress(t *testing.T) {
	type fields struct {
		ID      *string
		Date    *string
		Comment *string
		TimeBeg *string
		TimeEnd *string
		Address *Address
	}
	type args struct {
		address Address
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ScheduleAttempt
	}{
		{
			name: "set",
			fields: fields{
				ID:      strLink("attempt_id"),
				Date:    strLink("2019-07-17"),
				Address: nil,
			},
			args: args{
				address: Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
			want: &ScheduleAttempt{
				ID:   strLink("attempt_id"),
				Date: strLink("2019-07-17"),
				Address: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				ID:   strLink("attempt_id"),
				Date: strLink("2019-07-17"),
				Address: &Address{
					Street: strLink("previous street"),
					House:  strLink("11/4"),
				},
			},
			args: args{
				address: Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
			want: &ScheduleAttempt{
				ID:   strLink("attempt_id"),
				Date: strLink("2019-07-17"),
				Address: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduleAttempt := &ScheduleAttempt{
				ID:      tt.fields.ID,
				Date:    tt.fields.Date,
				Comment: tt.fields.Comment,
				TimeBeg: tt.fields.TimeBeg,
				TimeEnd: tt.fields.TimeEnd,
				Address: tt.fields.Address,
			}
			if got := scheduleAttempt.SetAddress(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScheduleAttempt.SetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCallCourier(t *testing.T) {
	type args struct {
		call CourierCallReq
	}
	tests := []struct {
		name string
		args args
		want *CallCourier
	}{
		{
			name: "constructor",
			args: args{
				call: CourierCallReq{
					Date:       strLink("2019-07-15"),
					TimeBeg:    strLink("10:00"),
					TimeEnd:    strLink("17:00"),
					SendPhone:  strLink("+79138739944"),
					SenderName: strLink("full name"),
					SendAddress: &Address{
						Street: strLink("test street"),
						House:  strLink("10/3"),
					},
				},
			},
			want: &CallCourier{
				Call: &CourierCallReq{
					Date:       strLink("2019-07-15"),
					TimeBeg:    strLink("10:00"),
					TimeEnd:    strLink("17:00"),
					SendPhone:  strLink("+79138739944"),
					SenderName: strLink("full name"),
					SendAddress: &Address{
						Street: strLink("test street"),
						House:  strLink("10/3"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCallCourier(tt.args.call); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCallCourier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCourierCall(t *testing.T) {
	type args struct {
		date        time.Time
		timeBeg     string
		timeEnd     string
		sendPhone   string
		senderName  string
		sendAddress Address
	}
	tests := []struct {
		name string
		args args
		want *CourierCallReq
	}{
		{
			name: "constructor",
			args: args{
				date:       time.Date(2019, 7, 15, 0, 0, 0, 0, time.UTC),
				timeBeg:    "10:00",
				timeEnd:    "17:00",
				sendPhone:  "+79138739944",
				senderName: "full name",
				sendAddress: Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCourierCall(
				tt.args.date,
				tt.args.timeBeg,
				tt.args.timeEnd,
				tt.args.sendPhone,
				tt.args.senderName,
				tt.args.sendAddress,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCourierCall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourierCallReq_SetLunchBeg(t *testing.T) {
	type fields struct {
		Date             *string
		TimeBeg          *string
		TimeEnd          *string
		LunchBeg         *string
		LunchEnd         *string
		SendCityCode     *int
		SendCityPostCode *string
		SendCountryCode  *string
		SendCityName     *string
		SendPhone        *string
		SenderName       *string
		Comment          *string
		SendAddress      *Address
	}
	type args struct {
		lunchBeg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CourierCallReq
	}{
		{
			name: "set",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				LunchBeg: nil,
			},
			args: args{
				lunchBeg: "12:00",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				LunchBeg: strLink("12:00"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				LunchBeg: strLink("13:00"),
			},
			args: args{
				lunchBeg: "12:00",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				LunchBeg: strLink("12:00"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := &CourierCallReq{
				Date:             tt.fields.Date,
				TimeBeg:          tt.fields.TimeBeg,
				TimeEnd:          tt.fields.TimeEnd,
				LunchBeg:         tt.fields.LunchBeg,
				LunchEnd:         tt.fields.LunchEnd,
				SendCityCode:     tt.fields.SendCityCode,
				SendCityPostCode: tt.fields.SendCityPostCode,
				SendCountryCode:  tt.fields.SendCountryCode,
				SendCityName:     tt.fields.SendCityName,
				SendPhone:        tt.fields.SendPhone,
				SenderName:       tt.fields.SenderName,
				Comment:          tt.fields.Comment,
				SendAddress:      tt.fields.SendAddress,
			}
			if got := call.SetLunchBeg(tt.args.lunchBeg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourierCallReq.SetLunchBeg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourierCallReq_SetLunchEnd(t *testing.T) {
	type fields struct {
		Date             *string
		TimeBeg          *string
		TimeEnd          *string
		LunchBeg         *string
		LunchEnd         *string
		SendCityCode     *int
		SendCityPostCode *string
		SendCountryCode  *string
		SendCityName     *string
		SendPhone        *string
		SenderName       *string
		Comment          *string
		SendAddress      *Address
	}
	type args struct {
		lunchEnd string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CourierCallReq
	}{
		{
			name: "set",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				LunchEnd: nil,
			},
			args: args{
				lunchEnd: "12:00",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				LunchEnd: strLink("12:00"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				LunchEnd: strLink("14:00"),
			},
			args: args{
				lunchEnd: "13:00",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				LunchEnd: strLink("13:00"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := &CourierCallReq{
				Date:             tt.fields.Date,
				TimeBeg:          tt.fields.TimeBeg,
				TimeEnd:          tt.fields.TimeEnd,
				LunchBeg:         tt.fields.LunchBeg,
				LunchEnd:         tt.fields.LunchEnd,
				SendCityCode:     tt.fields.SendCityCode,
				SendCityPostCode: tt.fields.SendCityPostCode,
				SendCountryCode:  tt.fields.SendCountryCode,
				SendCityName:     tt.fields.SendCityName,
				SendPhone:        tt.fields.SendPhone,
				SenderName:       tt.fields.SenderName,
				Comment:          tt.fields.Comment,
				SendAddress:      tt.fields.SendAddress,
			}
			if got := call.SetLunchEnd(tt.args.lunchEnd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourierCallReq.SetLunchEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourierCallReq_SetSendCityCode(t *testing.T) {
	type fields struct {
		Date             *string
		TimeBeg          *string
		TimeEnd          *string
		LunchBeg         *string
		LunchEnd         *string
		SendCityCode     *int
		SendCityPostCode *string
		SendCountryCode  *string
		SendCityName     *string
		SendPhone        *string
		SenderName       *string
		Comment          *string
		SendAddress      *Address
	}
	type args struct {
		sendCityCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CourierCallReq
	}{
		{
			name: "set",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityCode: nil,
			},
			args: args{
				sendCityCode: 44,
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityCode: intLink(44),
			},
		},
		{
			name: "modify",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityCode: intLink(55),
			},
			args: args{
				sendCityCode: 44,
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityCode: intLink(44),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := &CourierCallReq{
				Date:             tt.fields.Date,
				TimeBeg:          tt.fields.TimeBeg,
				TimeEnd:          tt.fields.TimeEnd,
				LunchBeg:         tt.fields.LunchBeg,
				LunchEnd:         tt.fields.LunchEnd,
				SendCityCode:     tt.fields.SendCityCode,
				SendCityPostCode: tt.fields.SendCityPostCode,
				SendCountryCode:  tt.fields.SendCountryCode,
				SendCityName:     tt.fields.SendCityName,
				SendPhone:        tt.fields.SendPhone,
				SenderName:       tt.fields.SenderName,
				Comment:          tt.fields.Comment,
				SendAddress:      tt.fields.SendAddress,
			}
			if got := call.SetSendCityCode(tt.args.sendCityCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourierCallReq.SetSendCityCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourierCallReq_SetSendCityPostCode(t *testing.T) {
	type fields struct {
		Date             *string
		TimeBeg          *string
		TimeEnd          *string
		LunchBeg         *string
		LunchEnd         *string
		SendCityCode     *int
		SendCityPostCode *string
		SendCountryCode  *string
		SendCityName     *string
		SendPhone        *string
		SenderName       *string
		Comment          *string
		SendAddress      *Address
	}
	type args struct {
		sendCityPostCode string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CourierCallReq
	}{
		{
			name: "set",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityPostCode: nil,
			},
			args: args{
				sendCityPostCode: "98690",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityPostCode: strLink("98690"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityPostCode: strLink("98734"),
			},
			args: args{
				sendCityPostCode: "98690",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityPostCode: strLink("98690"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := &CourierCallReq{
				Date:             tt.fields.Date,
				TimeBeg:          tt.fields.TimeBeg,
				TimeEnd:          tt.fields.TimeEnd,
				LunchBeg:         tt.fields.LunchBeg,
				LunchEnd:         tt.fields.LunchEnd,
				SendCityCode:     tt.fields.SendCityCode,
				SendCityPostCode: tt.fields.SendCityPostCode,
				SendCountryCode:  tt.fields.SendCountryCode,
				SendCityName:     tt.fields.SendCityName,
				SendPhone:        tt.fields.SendPhone,
				SenderName:       tt.fields.SenderName,
				Comment:          tt.fields.Comment,
				SendAddress:      tt.fields.SendAddress,
			}
			if got := call.SetSendCityPostCode(tt.args.sendCityPostCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourierCallReq.SetSendCityPostCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourierCallReq_SetSendCountryCode(t *testing.T) {
	type fields struct {
		Date             *string
		TimeBeg          *string
		TimeEnd          *string
		LunchBeg         *string
		LunchEnd         *string
		SendCityCode     *int
		SendCityPostCode *string
		SendCountryCode  *string
		SendCityName     *string
		SendPhone        *string
		SenderName       *string
		Comment          *string
		SendAddress      *Address
	}
	type args struct {
		sendCountryCode string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CourierCallReq
	}{
		{
			name: "set",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCountryCode: nil,
			},
			args: args{
				sendCountryCode: "RU",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCountryCode: strLink("RU"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCountryCode: strLink("UA"),
			},
			args: args{
				sendCountryCode: "RU",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCountryCode: strLink("RU"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := &CourierCallReq{
				Date:             tt.fields.Date,
				TimeBeg:          tt.fields.TimeBeg,
				TimeEnd:          tt.fields.TimeEnd,
				LunchBeg:         tt.fields.LunchBeg,
				LunchEnd:         tt.fields.LunchEnd,
				SendCityCode:     tt.fields.SendCityCode,
				SendCityPostCode: tt.fields.SendCityPostCode,
				SendCountryCode:  tt.fields.SendCountryCode,
				SendCityName:     tt.fields.SendCityName,
				SendPhone:        tt.fields.SendPhone,
				SenderName:       tt.fields.SenderName,
				Comment:          tt.fields.Comment,
				SendAddress:      tt.fields.SendAddress,
			}
			if got := call.SetSendCountryCode(tt.args.sendCountryCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourierCallReq.SetSendCountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourierCallReq_SetSendCityName(t *testing.T) {
	type fields struct {
		Date             *string
		TimeBeg          *string
		TimeEnd          *string
		LunchBeg         *string
		LunchEnd         *string
		SendCityCode     *int
		SendCityPostCode *string
		SendCountryCode  *string
		SendCityName     *string
		SendPhone        *string
		SenderName       *string
		Comment          *string
		SendAddress      *Address
	}
	type args struct {
		sendCityName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CourierCallReq
	}{
		{
			name: "set",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityName: nil,
			},
			args: args{
				sendCityName: "Орел",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityName: strLink("Орел"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityName: strLink("Саранск"),
			},
			args: args{
				sendCityName: "Орел",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				SendCityName: strLink("Орел"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := &CourierCallReq{
				Date:             tt.fields.Date,
				TimeBeg:          tt.fields.TimeBeg,
				TimeEnd:          tt.fields.TimeEnd,
				LunchBeg:         tt.fields.LunchBeg,
				LunchEnd:         tt.fields.LunchEnd,
				SendCityCode:     tt.fields.SendCityCode,
				SendCityPostCode: tt.fields.SendCityPostCode,
				SendCountryCode:  tt.fields.SendCountryCode,
				SendCityName:     tt.fields.SendCityName,
				SendPhone:        tt.fields.SendPhone,
				SenderName:       tt.fields.SenderName,
				Comment:          tt.fields.Comment,
				SendAddress:      tt.fields.SendAddress,
			}
			if got := call.SetSendCityName(tt.args.sendCityName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourierCallReq.SetSendCityName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourierCallReq_SetComment(t *testing.T) {
	type fields struct {
		Date             *string
		TimeBeg          *string
		TimeEnd          *string
		LunchBeg         *string
		LunchEnd         *string
		SendCityCode     *int
		SendCityPostCode *string
		SendCountryCode  *string
		SendCityName     *string
		SendPhone        *string
		SenderName       *string
		Comment          *string
		SendAddress      *Address
	}
	type args struct {
		comment string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CourierCallReq
	}{
		{
			name: "set",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				Comment: nil,
			},
			args: args{
				comment: "comment",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				Comment: strLink("comment"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				Comment: strLink("previous comment"),
			},
			args: args{
				comment: "comment",
			},
			want: &CourierCallReq{
				Date:       strLink("2019-07-15"),
				TimeBeg:    strLink("10:00"),
				TimeEnd:    strLink("17:00"),
				SendPhone:  strLink("+79138739944"),
				SenderName: strLink("full name"),
				SendAddress: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
				Comment: strLink("comment"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := &CourierCallReq{
				Date:             tt.fields.Date,
				TimeBeg:          tt.fields.TimeBeg,
				TimeEnd:          tt.fields.TimeEnd,
				LunchBeg:         tt.fields.LunchBeg,
				LunchEnd:         tt.fields.LunchEnd,
				SendCityCode:     tt.fields.SendCityCode,
				SendCityPostCode: tt.fields.SendCityPostCode,
				SendCountryCode:  tt.fields.SendCountryCode,
				SendCityName:     tt.fields.SendCityName,
				SendPhone:        tt.fields.SendPhone,
				SenderName:       tt.fields.SenderName,
				Comment:          tt.fields.Comment,
				SendAddress:      tt.fields.SendAddress,
			}
			if got := call.SetComment(tt.args.comment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourierCallReq.SetComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
