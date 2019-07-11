package cdek

import (
	"reflect"
	"testing"
)

func TestRegionFilterBuilder_AddFilter(t *testing.T) {
	type fields struct {
		filter map[RegionFilter]string
	}
	type args struct {
		filter RegionFilter
		value  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RegionFilterBuilder
	}{
		{
			name: "add single filter",
			fields: fields{
				filter: nil,
			},
			args: args{
				filter: RegionFilterPage,
				value:  "3",
			},
			want: &RegionFilterBuilder{
				filter: map[RegionFilter]string{
					RegionFilterPage: "3",
				},
			},
		},
		{
			name: "rewrite filter",
			fields: fields{
				filter: map[RegionFilter]string{
					RegionFilterPage: "3",
				},
			},
			args: args{
				filter: RegionFilterPage,
				value:  "2",
			},
			want: &RegionFilterBuilder{
				filter: map[RegionFilter]string{
					RegionFilterPage: "2",
				},
			},
		},
		{
			name: "add filter",
			fields: fields{
				filter: map[RegionFilter]string{
					RegionFilterPage: "2",
				},
			},
			args: args{
				filter: RegionFilterRegionCode,
				value:  "b8837188-39ee-4ff9-bc91-fcc9ed451bb3",
			},
			want: &RegionFilterBuilder{
				filter: map[RegionFilter]string{
					RegionFilterPage:       "2",
					RegionFilterRegionCode: "b8837188-39ee-4ff9-bc91-fcc9ed451bb3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterBuilder := &RegionFilterBuilder{
				filter: tt.fields.filter,
			}
			if got := filterBuilder.AddFilter(tt.args.filter, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegionFilterBuilder_Filter(t *testing.T) {
	type fields struct {
		filter map[RegionFilter]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[RegionFilter]string
	}{
		{
			name: "multiple filters",
			fields: fields{
				filter: map[RegionFilter]string{
					RegionFilterPage:       "2",
					RegionFilterRegionCode: "b8837188-39ee-4ff9-bc91-fcc9ed451bb3",
				},
			},
			want: map[RegionFilter]string{
				RegionFilterPage:       "2",
				RegionFilterRegionCode: "b8837188-39ee-4ff9-bc91-fcc9ed451bb3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterBuilder := &RegionFilterBuilder{
				filter: tt.fields.filter,
			}
			if got := filterBuilder.Filter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
