package webhook

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestNewServer(t *testing.T) {
	//s := NewServer("", "sucE2FiRTy4nkB5J6FPu", "")
	//http.ListenAndServe(":8081", s)
	//select {}
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
	<Encrypt><![CDATA[hyzAe4OzmOMbd6TvGdIOO6uBmdJoD0Fk53REIHvxYtJlE2B655HuD0m8KUePWB3+LrPXo87wzQ1QLvbeUgmBM4x6F8PGHQHFVAFmOD2LdJF9FrXpbUAh0B5GIItb52sn896wVsMSHGuPE328HnRGBcrS7C41IzDWyWNlZkyyXwon8T332jisa+h6tEDYsVticbSnyU8dKOIbgU6ux5VTjg3yt+WGzjlpKn6NPhRjpA912xMezR4kw6KWwMrCVKSVCZciVGCgavjIQ6X8tCOp3yZbGpy0VxpAe+77TszTfRd5RJSVO/HTnifJpXgCSUdUue1v6h0EIBYYI1BD1DlD+C0CR8e6OewpusjZ4uBl9FyJvnhvQl+q5rv1ixrcpCumEPo5MJSgM9ehVsNPfUM669WuMyVWQLCzpu9GhglF2PE=]]></Encrypt>
</xml>`

const xmlEncryptedText = "hyzAe4OzmOMbd6TvGdIOO6uBmdJoD0Fk53REIHvxYtJlE2B655HuD0m8KUePWB3+LrPXo87wzQ1QLvbeUgmBM4x6F8PGHQHFVAFmOD2LdJF9FrXpbUAh0B5GIItb52sn896wVsMSHGuPE328HnRGBcrS7C41IzDWyWNlZkyyXwon8T332jisa+h6tEDYsVticbSnyU8dKOIbgU6ux5VTjg3yt+WGzjlpKn6NPhRjpA912xMezR4kw6KWwMrCVKSVCZciVGCgavjIQ6X8tCOp3yZbGpy0VxpAe+77TszTfRd5RJSVO/HTnifJpXgCSUdUue1v6h0EIBYYI1BD1DlD+C0CR8e6OewpusjZ4uBl9FyJvnhvQl+q5rv1ixrcpCumEPo5MJSgM9ehVsNPfUM669WuMyVWQLCzpu9GhglF2PE="

const xmlMixedEncrypted = `<xml>
	<ToUserName><![CDATA[gh_10f6c3c3ac5a1111]]></ToUserName>
	<FromUserName><![CDATA[oyORnuP8q7ou2gfYjqLzSIWZf0rs]]></FromUserName>
	<CreateTime>1409735668</CreateTime>
	<MsgType><![CDATA[text]]></MsgType>
	<Content><![CDATA[abcdteT]]></Content>
	<MsgId>6054768590064713728</MsgId>
	<Encrypt><![CDATA[hyzAe4OzmOMbd6TvGdIOO6uBmdJoD0Fk53REIHvxYtJlE2B655HuD0m8KUePWB3+LrPXo87wzQ1QLvbeUgmBM4x6F8PGHQHFVAFmOD2LdJF9FrXpbUAh0B5GIItb52sn896wVsMSHGuPE328HnRGBcrS7C41IzDWyWNlZkyyXwon8T332jisa+h6tEDYsVticbSnyU8dKOIbgU6ux5VTjg3yt+WGzjlpKn6NPhRjpA912xMezR4kw6KWwMrCVKSVCZciVGCgavjIQ6X8tCOp3yZbGpy0VxpAe+77TszTfRd5RJSVO/HTnifJpXgCSUdUue1v6h0EIBYYI1BD1DlD+C0CR8e6OewpusjZ4uBl9FyJvnhvQl+q5rv1ixrcpCumEPo5MJSgM9ehVsNPfUM669WuMyVWQLCzpu9GhglF2PE=]]></Encrypt>
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
		wantMsg       *EncryptMessage
		wantPayload   []byte
		wantEncrypted bool
		wantErr       bool
	}{
		{
			name: "MethodGet, Not Body",
			args: args{
				token:  "sucE2FiRTy4nkB5J6FPu",
				method: http.MethodGet,
				query: url.Values{
					"nonce":     []string{"1117362755"},
					"signature": []string{"249e0702a0ae4d9e12854f0b0ddcd274dc3074cf"},
					"timestamp": []string{"1670509035"},
				},
			},
			wantMsg: nil,
			wantErr: true,
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
			name: "plain video xml",
			args: args{
				token:  "sucE2FiRTy4nkB5J6FPu",
				method: http.MethodPost,
				query: url.Values{
					"nonce":     []string{"1117362755"},
					"signature": []string{"249e0702a0ae4d9e12854f0b0ddcd274dc3074cf"},
					"timestamp": []string{"1670509035"},
				},
				body: xmlVideo,
			},
			wantMsg: &EncryptMessage{
				PlainMessage: PlainMessage{
					MsgType:      "video",
					ToUserName:   "oia2TjjewbmiOUlr6X-1crbLOvLw",
					FromUserName: "gh_7f083739789a",
					CreateTime:   1407743423,
				},
			},
			wantErr:     false,
			wantPayload: []byte(xmlVideo),
		},
		{
			name: "video signature fail",
			args: args{
				token:  "fail",
				method: http.MethodPost,
				query: url.Values{
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
					"nonce":     []string{"1117362755"},
					"signature": []string{"249e0702a0ae4d9e12854f0b0ddcd274dc3074cf"},
					"timestamp": []string{"1670509035"},
				},
				body: xmlVideo[:len(xmlVideo)-10],
			},
			wantMsg: nil,
			wantErr: true,
		},
		{
			name: "Encrypted xml",
			args: args{
				token:  "spamtest",
				method: http.MethodPost,
				query: url.Values{
					"nonce":     []string{"1320562132"},
					"signature": []string{"5d197aaffba7e9b25a30732f161a50dee96bd5fa"},
					"timestamp": []string{"1409735669"},
				},
				body: xmlEncrypted,
			},
			wantMsg: &EncryptMessage{
				Encrypt:      xmlEncryptedText,
				PlainMessage: PlainMessage{},
			},
			wantErr:       false,
			wantPayload:   []byte(xmlEncryptedText),
			wantEncrypted: true,
		},
		{
			name: "Encrypted and plain xml",
			args: args{
				token:  "spamtest",
				method: http.MethodPost,
				query: url.Values{
					"nonce":     []string{"1320562132"},
					"signature": []string{"5d197aaffba7e9b25a30732f161a50dee96bd5fa"},
					"timestamp": []string{"1409735669"},
				},
				body: xmlMixedEncrypted,
			},
			wantMsg: &EncryptMessage{
				Encrypt: xmlEncryptedText,
				PlainMessage: PlainMessage{
					ToUserName:   "gh_10f6c3c3ac5a1111",
					FromUserName: "oyORnuP8q7ou2gfYjqLzSIWZf0rs",
					MsgType:      "text",
					Content:      "abcdteT",
					MsgID:        6054768590064713728,
					CreateTime:   1409735668,
				},
			},
			wantErr:       false,
			wantPayload:   []byte(xmlEncryptedText),
			wantEncrypted: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_url := "/?" + tt.args.query.Encode()
			r := httptest.NewRequest(tt.args.method, _url, strings.NewReader(tt.args.body))
			gotMsg, gotPayload, gotEncrypted, err := ValidateWebhook(r, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateWebhook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("ValidateWebhook() gotMsg = %v, want %v", gotMsg, tt.wantMsg)
			}
			if !reflect.DeepEqual(gotPayload, tt.wantPayload) {
				t.Errorf("ValidateWebhook() gotPayload = %v, want %v", gotPayload, tt.wantPayload)
			}
			if gotEncrypted != tt.wantEncrypted {
				t.Errorf("ValidateWebhook() gotEncrypted = %v, want %v", gotEncrypted, tt.wantEncrypted)
			}
		})
	}
}
