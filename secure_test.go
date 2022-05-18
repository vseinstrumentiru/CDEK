package cdek

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"testing"
	"time"
)

func Test_securableJSON_setAuth(t *testing.T) {
	now := time.Now().Format("2006-01-02")
	encoder := md5.New()
	_, _ = encoder.Write([]byte(now + "&" + "testSecure"))
	testSecureEncoded := hex.EncodeToString(encoder.Sum(nil))

	type fields struct {
		AuthLogin   *string
		Secure      *string
		DateExecute *string
	}
	type args struct {
		auth *auth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *securable
	}{
		{
			name:   "secure fields pass correct",
			fields: fields{},
			args: args{
				auth: &auth{
					account: "testAccount",
					secure:  "testSecure",
				},
			},
			want: &securable{
				Account: strLink("testAccount"),
				Secure:  &testSecureEncoded,
				Date:    &now,
			},
		},
		{
			name:   "empty auth",
			fields: fields{},
			args:   args{},
			want:   &securable{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &securable{
				Account: tt.fields.AuthLogin,
				Secure:  tt.fields.Secure,
				Date:    tt.fields.DateExecute,
			}
			if got := s.setAuth(tt.args.auth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_securableXML_setAuth(t *testing.T) {
	testAccount := "testAccount"
	testSecure := "testSecure"
	now := time.Now().Format("2006-01-02")
	encoder := md5.New()
	_, _ = encoder.Write([]byte(now + "&" + testSecure))
	testSecureEncoded := hex.EncodeToString(encoder.Sum(nil))

	type fields struct {
		Account *string
		Date    *string
		Secure  *string
	}
	type args struct {
		auth *auth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *securable
	}{
		{
			name:   "secure fields pass correct",
			fields: fields{},
			args: args{
				auth: &auth{
					account: testAccount,
					secure:  testSecure,
				},
			},
			want: &securable{
				Account: &testAccount,
				Date:    &now,
				Secure:  &testSecureEncoded,
			},
		},
		{
			name:   "empty auth",
			fields: fields{},
			args:   args{},
			want:   &securable{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &securable{
				Account: tt.fields.Account,
				Date:    tt.fields.Date,
				Secure:  tt.fields.Secure,
			}
			if got := s.setAuth(tt.args.auth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
