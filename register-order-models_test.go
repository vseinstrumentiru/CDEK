package cdek

import "testing"

func TestOrderResp_GetError(t *testing.T) {
	type fields struct {
		Error          Error
		DispatchNumber *int
		Number         *string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "err",
			fields: fields{
				Error: Error{
					ErrorCode: strLink("err_code"),
					Msg:       strLink("error text"),
				},
				DispatchNumber: intLink(192957),
				Number:         strLink("test_number"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderResp{
				Error:          tt.fields.Error,
				DispatchNumber: tt.fields.DispatchNumber,
				Number:         tt.fields.Number,
			}
			if err := o.GetError(); (err != nil) != tt.wantErr {
				t.Errorf("OrderResp.GetError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
