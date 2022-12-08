package hook

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestValidateWebhook(t *testing.T) {
	type args struct {
		token string

		method string
		query  url.Values
		body   string
	}
	tests := []struct {
		name          string
		args          args
		wantMsg       *Message
		wantEncrypted bool
		wantErr       bool
	}{
		{
			name: "validate success",
			args: args{
				token:  "sucE2FiRTy4nkB5J6FPu",
				method: http.MethodGet,
				query: url.Values{
					"echostr":   []string{"138066257067505647"},
					"nonce":     []string{"1117362755"},
					"signature": []string{"249e0702a0ae4d9e12854f0b0ddcd274dc3074cf"},
					"timestamp": []string{"1670509035"},
				},
			},
			wantMsg: &Message{
				MsgType: "Validate",
			},
		},
		{
			name: "signature mismatch",
			args: args{
				token:  "fail",
				method: http.MethodGet,
				query: url.Values{
					"echostr":   []string{"138066257067505647"},
					"nonce":     []string{"1117362755"},
					"signature": []string{"249e0702a0ae4d9e12854f0b0ddcd274dc3074cf"},
					"timestamp": []string{"1670509035"},
				},
			},
			wantMsg: &Message{
				MsgType: "Validate",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_url := "/?" + tt.args.query.Encode()
			r := httptest.NewRequest(tt.args.method, _url, strings.NewReader(tt.args.body))
			msg, encrypted, err := ValidateWebhook(r, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateWebhook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(msg, tt.wantMsg) {
				t.Errorf("ValidateWebhook() msg = %v, wantMsg %v", msg, tt.wantMsg)
			}
			if encrypted != tt.wantEncrypted {
				t.Errorf("ValidateWebhook() encrypted = %v, wantEncrypted %v", encrypted, tt.wantEncrypted)
			}
		})
	}
}
