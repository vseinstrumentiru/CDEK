package cdek

import (
	"reflect"
	"testing"
)

func TestNewUpdateOrderReq(t *testing.T) {
	type args struct {
		number     string
		orderCount int
		order      UpdateOrder
	}
	tests := []struct {
		name string
		args args
		want *UpdateOrderReq
	}{
		{
			name: "constructor",
			args: args{
				number:     "test_number",
				orderCount: 1,
				order:      UpdateOrder{},
			},
			want: &UpdateOrderReq{
				Number:     strLink("test_number"),
				OrderCount: intLink(1),
				Order:      &UpdateOrder{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUpdateOrderReq(tt.args.number, tt.args.orderCount, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateOrderReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUpdateOrder(t *testing.T) {
	tests := []struct {
		name string
		want *UpdateOrder
	}{
		{
			name: "constructor",
			want: &UpdateOrder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUpdateOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetNumber(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		number string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				Number: nil,
			},
			args: args{
				number: "test_number",
			},
			want: &UpdateOrder{
				Number: strLink("test_number"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number: strLink("previous_number"),
			},
			args: args{
				number: "test_number",
			},
			want: &UpdateOrder{
				Number: strLink("test_number"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetDispatchNumber(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		dispatchNumber int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				DispatchNumber: nil,
			},
			args: args{
				dispatchNumber: 183727812,
			},
			want: &UpdateOrder{
				DispatchNumber: intLink(183727812),
			},
		},
		{
			name: "modify",
			fields: fields{
				DispatchNumber: intLink(473623423),
			},
			args: args{
				dispatchNumber: 183727812,
			},
			want: &UpdateOrder{
				DispatchNumber: intLink(183727812),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetDispatchNumber(tt.args.dispatchNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetDispatchNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetDeliveryRecipientCost(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		deliveryRecipientCost float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				DeliveryRecipientCost: nil,
			},
			args: args{
				deliveryRecipientCost: 188.99,
			},
			want: &UpdateOrder{
				DeliveryRecipientCost: float64Link(188.99),
			},
		},
		{
			name: "modify",
			fields: fields{
				DeliveryRecipientCost: float64Link(1773.99),
			},
			args: args{
				deliveryRecipientCost: 188.99,
			},
			want: &UpdateOrder{
				DeliveryRecipientCost: float64Link(188.99),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetDeliveryRecipientCost(tt.args.deliveryRecipientCost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetDeliveryRecipientCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetDeliveryRecipientVATRate(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		deliveryRecipientVATRate string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				DeliveryRecipientVATRate: nil,
			},
			args: args{
				deliveryRecipientVATRate: "VATX",
			},
			want: &UpdateOrder{
				DeliveryRecipientVATRate: strLink("VATX"),
			},
		},
		{
			name: "modify",
			fields: fields{
				DeliveryRecipientVATRate: strLink("VATY"),
			},
			args: args{
				deliveryRecipientVATRate: "VATX",
			},
			want: &UpdateOrder{
				DeliveryRecipientVATRate: strLink("VATX"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			got := updateOrder.SetDeliveryRecipientVATRate(tt.args.deliveryRecipientVATRate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetDeliveryRecipientVATRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetDeliveryRecipientVATSum(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		deliveryRecipientVATSum float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				DeliveryRecipientVATSum: nil,
			},
			args: args{
				deliveryRecipientVATSum: 88.99,
			},
			want: &UpdateOrder{
				DeliveryRecipientVATSum: float64Link(88.99),
			},
		},
		{
			name: "modify",
			fields: fields{
				DeliveryRecipientVATSum: float64Link(73.99),
			},
			args: args{
				deliveryRecipientVATSum: 88.99,
			},
			want: &UpdateOrder{
				DeliveryRecipientVATSum: float64Link(88.99),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetDeliveryRecipientVATSum(tt.args.deliveryRecipientVATSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetDeliveryRecipientVATSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetRecipientName(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		recipientName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				RecipientName: nil,
			},
			args: args{
				recipientName: "Mr Bobby",
			},
			want: &UpdateOrder{
				RecipientName: strLink("Mr Bobby"),
			},
		},
		{
			name: "modify",
			fields: fields{
				RecipientName: strLink("Mrs Sarah"),
			},
			args: args{
				recipientName: "Mr Bobby",
			},
			want: &UpdateOrder{
				RecipientName: strLink("Mr Bobby"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetRecipientName(tt.args.recipientName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetRecipientName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetPhone(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		phone string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				Phone: nil,
			},
			args: args{
				phone: "+79138739944",
			},
			want: &UpdateOrder{
				Phone: strLink("+79138739944"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Phone: strLink("+89138875429"),
			},
			args: args{
				phone: "+79138739944",
			},
			want: &UpdateOrder{
				Phone: strLink("+79138739944"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetPhone(tt.args.phone); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetRecipientINN(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		recipientINN string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				RecipientINN: nil,
			},
			args: args{
				recipientINN: "123432123432",
			},
			want: &UpdateOrder{
				RecipientINN: strLink("123432123432"),
			},
		},
		{
			name: "modify",
			fields: fields{
				RecipientINN: strLink("986538673967"),
			},
			args: args{
				recipientINN: "123432123432",
			},
			want: &UpdateOrder{
				RecipientINN: strLink("123432123432"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetRecipientINN(tt.args.recipientINN); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetRecipientINN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetDateInvoice(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		dateInvoice string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				DateInvoice: nil,
			},
			args: args{
				dateInvoice: "2019-07-18",
			},
			want: &UpdateOrder{
				DateInvoice: strLink("2019-07-18"),
			},
		},
		{
			name: "modify",
			fields: fields{
				DateInvoice: strLink("2017-06-13"),
			},
			args: args{
				dateInvoice: "2019-07-18",
			},
			want: &UpdateOrder{
				DateInvoice: strLink("2019-07-18"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetDateInvoice(tt.args.dateInvoice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetDateInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetPassport(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		passport Passport
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				Passport: nil,
			},
			args: args{
				passport: Passport{
					Series: strLink("passportSeries"),
				},
			},
			want: &UpdateOrder{
				Passport: &Passport{
					Series: strLink("passportSeries"),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Passport: &Passport{
					Series: strLink("previousPassportSeries"),
				},
			},
			args: args{
				passport: Passport{
					Series: strLink("passportSeries"),
				},
			},
			want: &UpdateOrder{
				Passport: &Passport{
					Series: strLink("passportSeries"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetPassport(tt.args.passport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetPassport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetAddress(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		address Address
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
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
			want: &UpdateOrder{
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
					Street: strLink("previous test street"),
					House:  strLink("11/4"),
				},
			},
			args: args{
				address: Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
			want: &UpdateOrder{
				Address: &Address{
					Street: strLink("test street"),
					House:  strLink("10/3"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetAddress(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetDeliveryRecipientCostAdv(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		v DeliveryRecipientCostAdv
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				Address: nil,
			},
			args: args{
				v: DeliveryRecipientCostAdv{
					Threshold: intLink(1000),
					Sum:       float64Link(33.8),
				},
			},
			want: &UpdateOrder{
				DeliveryRecipientCostAdv: &DeliveryRecipientCostAdv{
					Threshold: intLink(1000),
					Sum:       float64Link(33.8),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				DeliveryRecipientCostAdv: &DeliveryRecipientCostAdv{
					Threshold: intLink(1059),
					Sum:       float64Link(12.6),
				},
			},
			args: args{
				v: DeliveryRecipientCostAdv{
					Threshold: intLink(1000),
					Sum:       float64Link(33.8),
				},
			},
			want: &UpdateOrder{
				DeliveryRecipientCostAdv: &DeliveryRecipientCostAdv{
					Threshold: intLink(1000),
					Sum:       float64Link(33.8),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetDeliveryRecipientCostAdv(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetDeliveryRecipientCostAdv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder_SetPackage(t *testing.T) {
	type fields struct {
		Number                   *string
		DispatchNumber           *int
		DeliveryRecipientCost    *float64
		DeliveryRecipientVATRate *string
		DeliveryRecipientVATSum  *float64
		RecipientName            *string
		Phone                    *string
		RecipientINN             *string
		DateInvoice              *string
		Passport                 *Passport
		Address                  *Address
		DeliveryRecipientCostAdv *DeliveryRecipientCostAdv
		Package                  *OrderPackage
	}
	type args struct {
		pack OrderPackage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOrder
	}{
		{
			name: "set",
			fields: fields{
				Address: nil,
			},
			args: args{
				pack: OrderPackage{
					Number:  strLink("test number"),
					BarCode: strLink("test barcode"),
					Weight:  intLink(1000),
				},
			},
			want: &UpdateOrder{
				Package: &OrderPackage{
					Number:  strLink("test number"),
					BarCode: strLink("test barcode"),
					Weight:  intLink(1000),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				Package: &OrderPackage{
					Number:  strLink("previous number"),
					BarCode: strLink("previous barcode"),
					Weight:  intLink(1755),
				},
			},
			args: args{
				pack: OrderPackage{
					Number:  strLink("test number"),
					BarCode: strLink("test barcode"),
					Weight:  intLink(1000),
				},
			},
			want: &UpdateOrder{
				Package: &OrderPackage{
					Number:  strLink("test number"),
					BarCode: strLink("test barcode"),
					Weight:  intLink(1000),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOrder := &UpdateOrder{
				Number:                   tt.fields.Number,
				DispatchNumber:           tt.fields.DispatchNumber,
				DeliveryRecipientCost:    tt.fields.DeliveryRecipientCost,
				DeliveryRecipientVATRate: tt.fields.DeliveryRecipientVATRate,
				DeliveryRecipientVATSum:  tt.fields.DeliveryRecipientVATSum,
				RecipientName:            tt.fields.RecipientName,
				Phone:                    tt.fields.Phone,
				RecipientINN:             tt.fields.RecipientINN,
				DateInvoice:              tt.fields.DateInvoice,
				Passport:                 tt.fields.Passport,
				Address:                  tt.fields.Address,
				DeliveryRecipientCostAdv: tt.fields.DeliveryRecipientCostAdv,
				Package:                  tt.fields.Package,
			}
			if got := updateOrder.SetPackage(tt.args.pack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder.SetPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}
