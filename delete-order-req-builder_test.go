package cdek

import (
	"reflect"
	"testing"
)

func TestNewDeleteOrder(t *testing.T) {
	type args struct {
		number         string
		dispatchNumber int
	}
	tests := []struct {
		name string
		args args
		want *DeleteOrder
	}{
		{
			name: "constructor is ok",
			args: args{
				number:         "test number",
				dispatchNumber: 1,
			},
			want: &DeleteOrder{
				Number:         strLink("test number"),
				DispatchNumber: intLink(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeleteOrder(tt.args.number, tt.args.dispatchNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteOrder() = %v, want %v", got, tt.want)
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
