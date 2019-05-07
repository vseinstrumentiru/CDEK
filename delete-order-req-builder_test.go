package cdek

import (
	"reflect"
	"testing"
)

func TestNewDeleteOrder(t *testing.T) {
	tests := []struct {
		name string
		want *DeleteOrder
	}{
		{
			name: "constructor",
			want: &DeleteOrder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeleteOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteOrder_SetNumber(t *testing.T) {
	type fields struct {
		Number         *string
		DispatchNumber *int
	}
	type args struct {
		number string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteOrder
	}{
		{
			name: "set",
			fields: fields{
				Number: nil,
			},
			args: args{
				number: "number-soOEl0",
			},
			want: &DeleteOrder{
				Number: strLink("number-soOEl0"),
			},
		},
		{
			name: "modify",
			fields: fields{
				Number: strLink("number-previous"),
			},
			args: args{
				number: "number-soOEl0",
			},
			want: &DeleteOrder{
				Number: strLink("number-soOEl0"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &DeleteOrder{
				Number:         tt.fields.Number,
				DispatchNumber: tt.fields.DispatchNumber,
			}
			if got := o.SetNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteOrder.SetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteOrder_SetDispatchNumber(t *testing.T) {
	type fields struct {
		Number         *string
		DispatchNumber *int
	}
	type args struct {
		dispatchNumber int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteOrder
	}{
		{
			name: "set",
			fields: fields{
				DispatchNumber: nil,
			},
			args: args{
				dispatchNumber: 1105048590,
			},
			want: &DeleteOrder{
				DispatchNumber: intLink(1105048590),
			},
		},
		{
			name: "modify",
			fields: fields{
				DispatchNumber: intLink(73472387282),
			},
			args: args{
				dispatchNumber: 1105048590,
			},
			want: &DeleteOrder{
				DispatchNumber: intLink(1105048590),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &DeleteOrder{
				Number:         tt.fields.Number,
				DispatchNumber: tt.fields.DispatchNumber,
			}
			if got := o.SetDispatchNumber(tt.args.dispatchNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteOrder.SetDispatchNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeleteOrderReq(t *testing.T) {
	type args struct {
		number     string
		orderCount int
		order      DeleteOrder
	}
	tests := []struct {
		name string
		args args
		want *DeleteOrderReq
	}{
		{
			name: "constructor is ok",
			args: args{
				number:     "test number",
				orderCount: 1,
				order: DeleteOrder{
					Number: strLink("test order number"),
				},
			},
			want: &DeleteOrderReq{
				Number:     strLink("test number"),
				OrderCount: intLink(1),
				Order: &DeleteOrder{
					Number: strLink("test order number"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeleteOrderReq(tt.args.number, tt.args.orderCount, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteOrderReq() = %v, want %v", got, tt.want)
			}
		})
	}
}
