package cdek

import (
	"reflect"
	"testing"
)

func TestCityFilterBuilder_AddFilter(t *testing.T) {
	type fields struct {
		filter map[CityFilter]string
	}
	type args struct {
		filter CityFilter
		value  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CityFilterBuilder
	}{
		{
			name: "add single CityFilterPage filter",
			fields: fields{
				filter: nil,
			},
			args: args{
				filter: CityFilterPage,
				value:  "33",
			},
			want: &CityFilterBuilder{
				filter: map[CityFilter]string{
					CityFilterPage: "33",
				},
			},
		},
		{
			name: "add string filter",
			fields: fields{
				filter: nil,
			},
			args: args{
				filter: "page",
				value:  "22",
			},
			want: &CityFilterBuilder{
				filter: map[CityFilter]string{
					CityFilterPage: "22",
				},
			},
		},
		{
			name: "rewrite filter",
			fields: fields{
				filter: map[CityFilter]string{
					CityFilterPage: "55",
				},
			},
			args: args{
				filter: CityFilterPage,
				value:  "66",
			},
			want: &CityFilterBuilder{
				filter: map[CityFilter]string{
					CityFilterPage: "66",
				},
			},
		},
		{
			name: "add filter",
			fields: fields{
				filter: map[CityFilter]string{
					CityFilterPage: "77",
				},
			},
			args: args{
				filter: CityFilterCityName,
				value:  "Moscow",
			},
			want: &CityFilterBuilder{
				filter: map[CityFilter]string{
					CityFilterPage:     "77",
					CityFilterCityName: "Moscow",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterBuilder := &CityFilterBuilder{
				filter: tt.fields.filter,
			}
			if got := filterBuilder.AddFilter(tt.args.filter, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCityFilterBuilder_Filter(t *testing.T) {
	type fields struct {
		filter map[CityFilter]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[CityFilter]string
	}{
		{
			name: "multiple filters",
			fields: fields{
				filter: map[CityFilter]string{
					CityFilterPage:     "22",
					CityFilterCityName: "Moscow",
				},
			},
			want: map[CityFilter]string{
				CityFilterPage:     "22",
				CityFilterCityName: "Moscow",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterBuilder := &CityFilterBuilder{
				filter: tt.fields.filter,
			}
			if got := filterBuilder.Filter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
