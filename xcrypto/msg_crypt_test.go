package xcrypto

import (
	"reflect"
	"testing"
)

var xmlPlain = []byte(`<xml><ToUserName><![CDATA[gh_10f6c3c3ac5a]]></ToUserName>
<FromUserName><![CDATA[oyORnuP8q7ou2gfYjqLzSIWZf0rs]]></FromUserName>
<CreateTime>1409735668</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[abcdteT]]></Content>
<MsgId>6054768590064713728</MsgId>
</xml>`)

func TestWXBizMsgCrypt_Encrypt(t *testing.T) {
	wxc, err := NewWXBizMsgCrypt("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG", "wx2c2769f8efd9abc2")
	if err != nil {
		t.Error(err)
	}

	type fields struct {
		aesCrypto Crypto
		appID     string
	}
	type args struct {
		replyMsg []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				aesCrypto: wxc.aesCrypto,
				appID:     wxc.appID,
			},
			args: args{replyMsg: testPlain},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WXBizMsgCrypt{
				aesCrypto: tt.fields.aesCrypto,
				appID:     tt.fields.appID,
			}
			got, err := w.Encrypt(tt.args.replyMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err = w.Decrypt(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.args.replyMsg) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.args.replyMsg)
			}
		})
	}
}

func TestWXBizMsgCrypt_Decrypt(t *testing.T) {
	wxc, err := NewWXBizMsgCrypt("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG", "wx2c2769f8efd9abc2")
	if err != nil {
		t.Error(err)
	}

	type fields struct {
		aesCrypto Crypto
		appID     string
	}
	type args struct {
		postData []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				aesCrypto: wxc.aesCrypto,
				appID:     wxc.appID,
			},
			args: args{postData: testEncrypt},
			want: xmlPlain,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WXBizMsgCrypt{
				aesCrypto: tt.fields.aesCrypto,
				appID:     tt.fields.appID,
			}

			got, err := w.Decrypt(tt.args.postData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
