package cdek

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	type args struct {
		apiURL string
	}
	tests := []struct {
		name string
		args args
		want *clientImpl
	}{
		{
			"clientImpl created",
			args{
				apiURL: "apiURL",
			},
			&clientImpl{
				apiURL:        "apiURL",
				calculatorURL: calculatorURLDefault,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.apiURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_EncodedSecure(t *testing.T) {
	now := time.Now().Format("2006-01-02")
	encoder := md5.New()
	_, _ = encoder.Write([]byte(now + "&" + "testSecure"))
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
				Account: "testAccount",
				Secure:  "testSecure",
			},
			wantDate:          now,
			wantEncodedSecure: testSecureEncoded,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := auth{
				account: tt.fields.Account,
				secure:  tt.fields.Secure,
			}
			gotDate, gotEncodedSecure := a.encodedSecure()
			if gotDate != tt.wantDate {
				t.Errorf("encodedSecure() gotDate = %v, want %v", gotDate, tt.wantDate)
			}
			if gotEncodedSecure != tt.wantEncodedSecure {
				t.Errorf("encodedSecure() gotEncodedSecure = %v, want %v", gotEncodedSecure, tt.wantEncodedSecure)
			}
		})
	}
}

func TestClient_SetAuth(t *testing.T) {
	type fields struct {
		Auth          *auth
		CdekAPIURL    string
		CalculatorURL string
	}
	type args struct {
		account string
		secure  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *clientImpl
	}{
		{
			name: "auth set",
			fields: fields{
				Auth: nil,
			},
			args: args{
				account: "testAccount",
				secure:  "testSecure",
			},
			want: &clientImpl{
				auth: &auth{
					account: "testAccount",
					secure:  "testSecure",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientConf := &clientImpl{
				auth:          tt.fields.Auth,
				apiURL:        tt.fields.CdekAPIURL,
				calculatorURL: tt.fields.CalculatorURL,
			}
			if got := clientConf.SetAuth(tt.args.account, tt.args.secure); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetCalculatorURL(t *testing.T) {
	type fields struct {
		Auth          *auth
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
		want   *clientImpl
	}{
		{
			name: "set url",
			fields: fields{
				CalculatorURL: "",
			},
			args: args{
				calculatorURL: "testCalcUrl",
			},
			want: &clientImpl{
				calculatorURL: "testCalcUrl",
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
			want: &clientImpl{
				calculatorURL: "testCalcUrl_rewritten",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientConf := &clientImpl{
				auth:          tt.fields.Auth,
				apiURL:        tt.fields.CdekAPIURL,
				calculatorURL: tt.fields.CalculatorURL,
			}
			if got := clientConf.SetCalculatorURL(tt.args.calculatorURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCalculatorURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleNewClient() {
	client := NewClient("https://integration.edu.cdek.ru/")
	client.SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")
}
