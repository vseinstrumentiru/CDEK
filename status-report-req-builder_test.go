package cdek

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestNewStatusReportReq(t *testing.T) {
	tests := []struct {
		name string
		want *StatusReport
	}{
		{
			name: "constructor",
			want: &StatusReport{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatusReportReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatusReportReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusReportReq_SetShowHistory(t *testing.T) {
	type fields struct {
		ShowHistory            *int
		ShowReturnOrder        *bool
		ShowReturnOrderHistory *bool
		ChangePeriod           *ChangePeriod
		Order                  []*StatusReportOrderReq
	}
	type args struct {
		showHistory int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StatusReport
	}{
		{
			name: "set",
			fields: fields{
				ShowHistory:     nil,
				ShowReturnOrder: boolLink(true),
			},
			args: args{
				showHistory: 1,
			},
			want: &StatusReport{
				ShowHistory:     intLink(1),
				ShowReturnOrder: boolLink(true),
			},
		},
		{
			name: "modify",
			fields: fields{
				ShowHistory:     intLink(1),
				ShowReturnOrder: boolLink(true),
			},
			args: args{
				showHistory: 1,
			},
			want: &StatusReport{
				ShowHistory:     intLink(1),
				ShowReturnOrder: boolLink(true),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &StatusReport{
				ShowHistory:            tt.fields.ShowHistory,
				ShowReturnOrder:        tt.fields.ShowReturnOrder,
				ShowReturnOrderHistory: tt.fields.ShowReturnOrderHistory,
				ChangePeriod:           tt.fields.ChangePeriod,
				Order:                  tt.fields.Order,
			}
			if got := req.SetShowHistory(tt.args.showHistory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusReport.SetShowHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusReportReq_SetShowReturnOrder(t *testing.T) {
	type fields struct {
		ShowHistory            *int
		ShowReturnOrder        *bool
		ShowReturnOrderHistory *bool
		ChangePeriod           *ChangePeriod
		Order                  []*StatusReportOrderReq
	}
	type args struct {
		showReturnOrder bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StatusReport
	}{
		{
			name: "set",
			fields: fields{
				ShowHistory:     intLink(1),
				ShowReturnOrder: nil,
			},
			args: args{
				showReturnOrder: true,
			},
			want: &StatusReport{
				ShowHistory:     intLink(1),
				ShowReturnOrder: boolLink(true),
			},
		},
		{
			name: "modify",
			fields: fields{
				ShowHistory:     intLink(1),
				ShowReturnOrder: boolLink(true),
			},
			args: args{
				showReturnOrder: false,
			},
			want: &StatusReport{
				ShowHistory:     intLink(1),
				ShowReturnOrder: boolLink(false),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &StatusReport{
				ShowHistory:            tt.fields.ShowHistory,
				ShowReturnOrder:        tt.fields.ShowReturnOrder,
				ShowReturnOrderHistory: tt.fields.ShowReturnOrderHistory,
				ChangePeriod:           tt.fields.ChangePeriod,
				Order:                  tt.fields.Order,
			}
			if got := req.SetShowReturnOrder(tt.args.showReturnOrder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusReport.SetShowReturnOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusReportReq_SetShowReturnOrderHistory(t *testing.T) {
	type fields struct {
		ShowHistory            *int
		ShowReturnOrder        *bool
		ShowReturnOrderHistory *bool
		ChangePeriod           *ChangePeriod
		Order                  []*StatusReportOrderReq
	}
	type args struct {
		showReturnOrderHistory bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StatusReport
	}{
		{
			name: "set",
			fields: fields{
				ShowHistory:            intLink(1),
				ShowReturnOrderHistory: nil,
			},
			args: args{
				showReturnOrderHistory: true,
			},
			want: &StatusReport{
				ShowHistory:            intLink(1),
				ShowReturnOrderHistory: boolLink(true),
			},
		},
		{
			name: "modify",
			fields: fields{
				ShowHistory:            intLink(1),
				ShowReturnOrderHistory: boolLink(true),
			},
			args: args{
				showReturnOrderHistory: false,
			},
			want: &StatusReport{
				ShowHistory:            intLink(1),
				ShowReturnOrderHistory: boolLink(false),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &StatusReport{
				ShowHistory:            tt.fields.ShowHistory,
				ShowReturnOrder:        tt.fields.ShowReturnOrder,
				ShowReturnOrderHistory: tt.fields.ShowReturnOrderHistory,
				ChangePeriod:           tt.fields.ChangePeriod,
				Order:                  tt.fields.Order,
			}
			if got := req.SetShowReturnOrderHistory(tt.args.showReturnOrderHistory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusReport.SetShowReturnOrderHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusReportReq_SetChangePeriod(t *testing.T) {
	type fields struct {
		ShowHistory            *int
		ShowReturnOrder        *bool
		ShowReturnOrderHistory *bool
		ChangePeriod           *ChangePeriod
		Order                  []*StatusReportOrderReq
	}
	type args struct {
		changePeriod ChangePeriod
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StatusReport
	}{
		{
			name: "set",
			fields: fields{
				ShowHistory:  intLink(1),
				ChangePeriod: nil,
			},
			args: args{
				changePeriod: ChangePeriod{
					DateFirst: strLink("2019-07-15"),
					DateLast:  strLink("2019-07-16"),
				},
			},
			want: &StatusReport{
				ShowHistory: intLink(1),
				ChangePeriod: &ChangePeriod{
					DateFirst: strLink("2019-07-15"),
					DateLast:  strLink("2019-07-16"),
				},
			},
		},
		{
			name: "modify",
			fields: fields{
				ShowHistory: intLink(1),
				ChangePeriod: &ChangePeriod{
					DateFirst: strLink("2019-07-15"),
					DateLast:  strLink("2019-07-16"),
				},
			},
			args: args{
				changePeriod: ChangePeriod{
					DateFirst: strLink("2019-07-16"),
					DateLast:  strLink("2019-07-17"),
				},
			},
			want: &StatusReport{
				ShowHistory: intLink(1),
				ChangePeriod: &ChangePeriod{
					DateFirst: strLink("2019-07-16"),
					DateLast:  strLink("2019-07-17"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &StatusReport{
				ShowHistory:            tt.fields.ShowHistory,
				ShowReturnOrder:        tt.fields.ShowReturnOrder,
				ShowReturnOrderHistory: tt.fields.ShowReturnOrderHistory,
				ChangePeriod:           tt.fields.ChangePeriod,
				Order:                  tt.fields.Order,
			}
			if got := req.SetChangePeriod(tt.args.changePeriod); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusReport.SetChangePeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusReportReq_AddOrder(t *testing.T) {
	type fields struct {
		ShowHistory            *int
		ShowReturnOrder        *bool
		ShowReturnOrderHistory *bool
		ChangePeriod           *ChangePeriod
		Order                  []*StatusReportOrderReq
	}
	type args struct {
		order StatusReportOrderReq
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StatusReport
	}{
		{
			name: "add first",
			fields: fields{
				ShowHistory: intLink(1),
				Order:       nil,
			},
			args: args{
				order: StatusReportOrderReq{
					DispatchNumber: intLink(1),
					Number:         strLink("test_number"),
					Date:           strLink("2019-07-15"),
				},
			},
			want: &StatusReport{
				ShowHistory: intLink(1),
				Order: []*StatusReportOrderReq{
					{
						DispatchNumber: intLink(1),
						Number:         strLink("test_number"),
						Date:           strLink("2019-07-15"),
					},
				},
			},
		},
		{
			name: "add second",
			fields: fields{
				ShowHistory: intLink(1),
				Order: []*StatusReportOrderReq{
					{
						DispatchNumber: intLink(1),
						Number:         strLink("test_number"),
						Date:           strLink("2019-07-15"),
					},
				},
			},
			args: args{
				order: StatusReportOrderReq{
					DispatchNumber: intLink(2),
					Number:         strLink("test_number_2"),
					Date:           strLink("2019-07-16"),
				},
			},
			want: &StatusReport{
				ShowHistory: intLink(1),
				Order: []*StatusReportOrderReq{
					{
						DispatchNumber: intLink(1),
						Number:         strLink("test_number"),
						Date:           strLink("2019-07-15"),
					},
					{
						DispatchNumber: intLink(2),
						Number:         strLink("test_number_2"),
						Date:           strLink("2019-07-16"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &StatusReport{
				ShowHistory:            tt.fields.ShowHistory,
				ShowReturnOrder:        tt.fields.ShowReturnOrder,
				ShowReturnOrderHistory: tt.fields.ShowReturnOrderHistory,
				ChangePeriod:           tt.fields.ChangePeriod,
				Order:                  tt.fields.Order,
			}
			if got := req.AddOrder(tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusReport.AddOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewChangePeriod(t *testing.T) {
	type args struct {
		dateFirst time.Time
	}

	tests := []struct {
		name string
		args args
		want *ChangePeriod
	}{
		{
			name: "constructor",
			args: args{
				dateFirst: time.Date(2019, 7, 15, 0, 0, 0, 0, time.UTC),
			},
			want: &ChangePeriod{
				DateFirst: strLink("2019-07-15"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChangePeriod(tt.args.dateFirst); !reflect.DeepEqual(got, tt.want) {
				g, _ := json.Marshal(got)
				w, _ := json.Marshal(tt.want)
				t.Errorf("NewChangePeriod() = %v, want %v", string(g), string(w))
			}
		})
	}
}

func TestChangePeriod_SetDateLast(t *testing.T) {
	type fields struct {
		DateFirst *string
		DateLast  *string
	}
	type args struct {
		date time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ChangePeriod
	}{
		{
			name: "set",
			fields: fields{
				DateFirst: strLink("2019-07-15"),
				DateLast:  nil,
			},
			args: args{
				date: time.Date(2019, 7, 16, 0, 0, 0, 0, time.UTC),
			},
			want: &ChangePeriod{
				DateFirst: strLink("2019-07-15"),
				DateLast:  strLink("2019-07-16"),
			},
		},
		{
			name: "modify",
			fields: fields{
				DateFirst: strLink("2019-07-15"),
				DateLast:  strLink("2019-07-16"),
			},
			args: args{
				date: time.Date(2019, 7, 17, 0, 0, 0, 0, time.UTC),
			},
			want: &ChangePeriod{
				DateFirst: strLink("2019-07-15"),
				DateLast:  strLink("2019-07-17"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changePeriod := &ChangePeriod{
				DateFirst: tt.fields.DateFirst,
				DateLast:  tt.fields.DateLast,
			}
			if got := changePeriod.SetDateLast(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangePeriod.SetDateLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStatusReportOrderReq(t *testing.T) {
	type args struct {
		dispatchNumber int
		number         string
		date           time.Time
	}
	tests := []struct {
		name string
		args args
		want *StatusReportOrderReq
	}{
		{
			name: "constructor",
			args: args{
				dispatchNumber: 1,
				number:         "test_number",
				date:           time.Date(2019, 7, 15, 0, 0, 0, 0, time.UTC),
			},
			want: &StatusReportOrderReq{
				DispatchNumber: intLink(1),
				Number:         strLink("test_number"),
				Date:           strLink("2019-07-15"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStatusReportOrderReq(tt.args.dispatchNumber, tt.args.number, tt.args.date)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatusReportOrderReq() = %v, want %v", got, tt.want)
			}
		})
	}
}
