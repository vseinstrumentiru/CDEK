package cdek

import "testing"

func TestError_IsErroneous(t *testing.T) {
	type fields struct {
		ErrorCode *string
		Msg       *string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "is erroneous",
			fields: fields{
				ErrorCode: strLink("some_code"),
				Msg:       strLink("some message"),
			},
			want: true,
		},
		{
			name: "is not erroneous",
			fields: fields{
				ErrorCode: nil,
				Msg:       nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				ErrorCode: tt.fields.ErrorCode,
				Msg:       tt.fields.Msg,
			}
			if got := e.IsErroneous(); got != tt.want {
				t.Errorf("Error.IsErroneous() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		ErrorCode *string
		Msg       *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "make error string",
			fields: fields{
				ErrorCode: strLink("some_code"),
				Msg:       strLink("some message"),
			},
			want: "some message; ErrorCode: some_code",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				ErrorCode: tt.fields.ErrorCode,
				Msg:       tt.fields.Msg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlert_Error(t *testing.T) {
	type fields struct {
		Type      string
		Msg       string
		ErrorCode string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "make error string",
			fields: fields{
				Type:      "some_type",
				Msg:       "some message",
				ErrorCode: "some_code",
			},
			want: "Type: some_type; Msg: some message; ErrorCode: some_code",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Alert{
				Type:      tt.fields.Type,
				Msg:       tt.fields.Msg,
				ErrorCode: tt.fields.ErrorCode,
			}
			if got := a.Error(); got != tt.want {
				t.Errorf("Alert.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
