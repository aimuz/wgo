package hook

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer("", "sucE2FiRTy4nkB5J6FPu", "")
	http.ListenAndServe(":8081", s)
	select {}
	//tests := []struct {
	//	name string
	//	want *Server
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := NewServer(); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("NewServer() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

const xmlVideo = `<xml>
	<ToUserName><![CDATA[oia2TjjewbmiOUlr6X-1crbLOvLw]]></ToUserName>
	<FromUserName><![CDATA[gh_7f083739789a]]></FromUserName>
	<CreateTime>1407743423</CreateTime>
	<MsgType><![CDATA[video]]></MsgType>
	<Video>
		<MediaId><![CDATA[eYJ1MbwPRJtOvIEabaxHs7TX2D-HV71s79GUxqdUkjm6Gs2Ed1KF3ulAOA9H1xG0]]></MediaId>
		<Title><![CDATA[testCallBackReplyVideo]]></Title>
		<Description><![CDATA[testCallBackReplyVideo]]></Description>
	</Video>
</xml>`

const xmlEncrypted = `<xml>
	<Encrypt><![CDATA[dh0NcwUkVnUzto3InGyUGmGputX0KoTIsx1oN+GtqaBPd1DAT6ULZ3usGQHBwkiCAZ0S+ZqQ0G2OflIwfTgwgyU7bS/buG6r7vIoMpzAhviQMP5uVwzTqmDcppsYROr2k9uK+4WrPcN4PNZCJSNOZw+LIvTUTZNlho2aaLfTMzajcVnLzphBWn4Qe/T6G+GhS8HvUpiIUs7tzOEd8FLUtcYiQ1hgBFAiTgz0smOv4Ve11K4HMYXSn1/pUkVnErvyNIEmTQG5gLplDZCC1OdKOplLEckcGMIa+jagogbEhv9186/JY02sPk9tXisEq9RnuITQrbwyicEwdTMaupcXg8E8Rklr7vADk5lUE35rxW8dS5MA4gkWaEarzvdN/6akFcoAkaAzRuQUoC3ZEJ+WD+TOIKddJK9mWIoLDmO4s8F8sTLfSimQZnAJfoRfcNO9zVrAipiBvqymYhHwGO4q/cQEFdG9Srj/6iJuQ8yVl3ZS7Y7a5vBzPyQDfnX+a4svHUPs5s8CqlUMZduF1yymKsDG3LrOnsnMefG3ILH21iHfxVicjzxH3mmBjC5sog6d5GOTAPG6vbicYtTUyFG+OWIIlpR+hq6W8jNmrpKJIhbNDee8DP7RPNajC95udrLV]]></Encrypt>
	<MsgSignature><![CDATA[f7ece64907beca235c89418e3c0fd17719ee2150]]></MsgSignature>
	<TimeStamp>1670604719</TimeStamp>
	<Nonce><![CDATA[1320562132]]></Nonce>
</xml>`

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
			wantMsg: &Message{},
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
			wantMsg: nil,
			wantErr: true,
		},
		{
			name: "empty body",
			args: args{
				token:  "sucE2FiRTy4nkB5J6FPu",
				method: http.MethodPost,
				query: url.Values{
					"echostr":   []string{"138066257067505647"},
					"nonce":     []string{"1117362755"},
					"signature": []string{"249e0702a0ae4d9e12854f0b0ddcd274dc3074cf"},
					"timestamp": []string{"1670509035"},
				},
				body: "",
			},
			wantMsg: nil,
			wantErr: true,
		},
		{
			name: "video",
			args: args{
				token:  "sucE2FiRTy4nkB5J6FPu",
				method: http.MethodPost,
				query: url.Values{
					"echostr":   []string{"138066257067505647"},
					"nonce":     []string{"1117362755"},
					"signature": []string{"249e0702a0ae4d9e12854f0b0ddcd274dc3074cf"},
					"timestamp": []string{"1670509035"},
				},
				body: xmlVideo,
			},
			wantMsg: &Message{
				MsgType:      "video",
				ToUserName:   "oia2TjjewbmiOUlr6X-1crbLOvLw",
				FromUserName: "gh_7f083739789a",
				CreateTime:   1407743423,
			},
			wantErr: false,
		},
		{
			name: "video signature fail",
			args: args{
				token:  "fail",
				method: http.MethodPost,
				query: url.Values{
					"echostr":   []string{"138066257067505647"},
					"nonce":     []string{"1117362755"},
					"signature": []string{"249e0702a0ae4d9e12854f0b0ddcd274dc3074cf"},
					"timestamp": []string{"1670509035"},
				},
				body: xmlVideo,
			},
			wantMsg: nil,
			wantErr: true,
		},
		{
			name: "fail xml",
			args: args{
				token:  "fail",
				method: http.MethodPost,
				query: url.Values{
					"echostr":   []string{"138066257067505647"},
					"nonce":     []string{"1117362755"},
					"signature": []string{"249e0702a0ae4d9e12854f0b0ddcd274dc3074cf"},
					"timestamp": []string{"1670509035"},
				},
				body: xmlVideo[:len(xmlVideo)-10],
			},
			wantMsg: nil,
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
