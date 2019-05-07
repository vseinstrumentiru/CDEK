package cdek

import (
	"reflect"
	"testing"
)

func TestPvzListFilterBuilder_AddFilter(t *testing.T) {
	type fields struct {
		filter map[PvzListFilter]string
	}
	type args struct {
		filter PvzListFilter
		value  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PvzListFilterBuilder
	}{
		{
			name: "add single filter",
			fields: fields{
				filter: nil,
			},
			args: args{
				filter: PvzListFilterCityID,
				value:  "33",
			},
			want: &PvzListFilterBuilder{
				filter: map[PvzListFilter]string{
					PvzListFilterCityID: "33",
				},
			},
		},
		{
			name: "rewrite filter",
			fields: fields{
				filter: map[PvzListFilter]string{
					PvzListFilterCityID: "55",
				},
			},
			args: args{
				filter: PvzListFilterCityID,
				value:  "66",
			},
			want: &PvzListFilterBuilder{
				filter: map[PvzListFilter]string{
					PvzListFilterCityID: "66",
				},
			},
		},
		{
			name: "add filter",
			fields: fields{
				filter: map[PvzListFilter]string{
					PvzListFilterCityID: "77",
				},
			},
			args: args{
				filter: PvzListFilterRegionID,
				value:  "b8837188-39ee-4ff9-bc91-fcc9ed451bb3",
			},
			want: &PvzListFilterBuilder{
				filter: map[PvzListFilter]string{
					PvzListFilterCityID:   "77",
					PvzListFilterRegionID: "b8837188-39ee-4ff9-bc91-fcc9ed451bb3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterBuilder := &PvzListFilterBuilder{
				filter: tt.fields.filter,
			}
			if got := filterBuilder.AddFilter(tt.args.filter, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPvzListFilterBuilder_Filter(t *testing.T) {
	type fields struct {
		filter map[PvzListFilter]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[PvzListFilter]string
	}{
		{
			name: "multiple filters",
			fields: fields{
				filter: map[PvzListFilter]string{
					PvzListFilterCityID:   "77",
					PvzListFilterRegionID: "b8837188-39ee-4ff9-bc91-fcc9ed451bb3",
				},
			},
			want: map[PvzListFilter]string{
				PvzListFilterCityID:   "77",
				PvzListFilterRegionID: "b8837188-39ee-4ff9-bc91-fcc9ed451bb3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterBuilder := PvzListFilterBuilder{
				filter: tt.fields.filter,
			}
			if got := filterBuilder.Filter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
