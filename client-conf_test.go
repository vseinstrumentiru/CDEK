package cdek

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"testing"
	"time"
)

func TestAuth_EncodedSecure(t *testing.T) {
	testAccount := "testAccount"
	testSecure := "testSecure"
	now := time.Now().Format("2006-01-02")
	encoder := md5.New()
	_, _ = encoder.Write([]byte(now + "&" + testSecure))
	testSecureEncoded := hex.EncodeToString(encoder.Sum(nil))

	type fields struct {
		Account string
		Secure  string
	}
	tests := []struct {
		name              string
		fields            fields
		wantDate          string
		wantEncodedSecure string
	}{
		{
			name: "successful encoding",
			fields: fields{
				Account: testAccount,
				Secure:  testSecure,
			},
			wantDate:          now,
			wantEncodedSecure: testSecureEncoded,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Auth{
				Account: tt.fields.Account,
				Secure:  tt.fields.Secure,
			}
			gotDate, gotEncodedSecure := a.EncodedSecure()
			if gotDate != tt.wantDate {
				t.Errorf("EncodedSecure() gotDate = %v, want %v", gotDate, tt.wantDate)
			}
			if gotEncodedSecure != tt.wantEncodedSecure {
				t.Errorf("EncodedSecure() gotEncodedSecure = %v, want %v", gotEncodedSecure, tt.wantEncodedSecure)
			}
		})
	}
}

func TestClientConf_SetAuth(t *testing.T) {
	type fields struct {
		Auth          *Auth
		CdekAPIURL    string
		CalculatorURL string
	}
	type args struct {
		auth *Auth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ClientConf
	}{
		{
			name: "auth set",
			fields: fields{
				Auth: nil,
			},
			args: args{
				auth: &Auth{
					Account: "testAccount",
					Secure:  "testSecure",
				},
			},
			want: &ClientConf{
				Auth: &Auth{
					Account: "testAccount",
					Secure:  "testSecure",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientConf := &ClientConf{
				Auth:          tt.fields.Auth,
				CdekAPIURL:    tt.fields.CdekAPIURL,
				CalculatorURL: tt.fields.CalculatorURL,
			}
			if got := clientConf.SetAuth(tt.args.auth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientConf_SetCalculatorURL(t *testing.T) {
	type fields struct {
		Auth          *Auth
		CdekAPIURL    string
		CalculatorURL string
	}
	type args struct {
		calculatorURL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ClientConf
	}{
		{
			name: "set url",
			fields: fields{
				CalculatorURL: "",
			},
			args: args{
				calculatorURL: "testCalcUrl",
			},
			want: &ClientConf{
				CalculatorURL: "testCalcUrl",
			},
		},
		{
			name: "rewrite url",
			fields: fields{
				CalculatorURL: "",
			},
			args: args{
				calculatorURL: "testCalcUrl_rewritten",
			},
			want: &ClientConf{
				CalculatorURL: "testCalcUrl_rewritten",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientConf := &ClientConf{
				Auth:          tt.fields.Auth,
				CdekAPIURL:    tt.fields.CdekAPIURL,
				CalculatorURL: tt.fields.CalculatorURL,
			}
			if got := clientConf.SetCalculatorURL(tt.args.calculatorURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCalculatorURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClientConf(t *testing.T) {
	type args struct {
		cdekAPIURL string
	}
	tests := []struct {
		name string
		args args
		want *ClientConf
	}{
		{
			name: "client created",
			args: args{
				cdekAPIURL: "testUrl",
			},
			want: &ClientConf{
				CdekAPIURL:    "testUrl",
				CalculatorURL: calculatorURLDefault,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientConf(tt.args.cdekAPIURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
