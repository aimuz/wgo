package xcrypto

import (
	"encoding/base64"
	"reflect"
	"testing"
)

func TestAESUseCBCWithPKCS7(t *testing.T) {

}

var testAESKey, _ = base64.StdEncoding.DecodeString("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG=")

var testEncrypt, _ = base64.StdEncoding.DecodeString("hyzAe4OzmOMbd6TvGdIOO6uBmdJoD0Fk53REIHvxYtJlE2B655HuD0m8KUePWB3+LrPXo87wzQ1QLvbeUgmBM4x6F8PGHQHFVAFmOD2LdJF9FrXpbUAh0B5GIItb52sn896wVsMSHGuPE328HnRGBcrS7C41IzDWyWNlZkyyXwon8T332jisa+h6tEDYsVticbSnyU8dKOIbgU6ux5VTjg3yt+WGzjlpKn6NPhRjpA912xMezR4kw6KWwMrCVKSVCZciVGCgavjIQ6X8tCOp3yZbGpy0VxpAe+77TszTfRd5RJSVO/HTnifJpXgCSUdUue1v6h0EIBYYI1BD1DlD+C0CR8e6OewpusjZ4uBl9FyJvnhvQl+q5rv1ixrcpCumEPo5MJSgM9ehVsNPfUM669WuMyVWQLCzpu9GhglF2PE=")

var testPlain, _ = base64.StdEncoding.DecodeString("ODk0NjVjODQwYzVmMTE2ZgAAARg8eG1sPjxUb1VzZXJOYW1lPjwhW0NEQVRBW2doXzEwZjZjM2MzYWM1YV1dPjwvVG9Vc2VyTmFtZT4KPEZyb21Vc2VyTmFtZT48IVtDREFUQVtveU9SbnVQOHE3b3UyZ2ZZanFMelNJV1pmMHJzXV0+PC9Gcm9tVXNlck5hbWU+CjxDcmVhdGVUaW1lPjE0MDk3MzU2Njg8L0NyZWF0ZVRpbWU+CjxNc2dUeXBlPjwhW0NEQVRBW3RleHRdXT48L01zZ1R5cGU+CjxDb250ZW50PjwhW0NEQVRBW2FiY2R0ZVRdXT48L0NvbnRlbnQ+CjxNc2dJZD42MDU0NzY4NTkwMDY0NzEzNzI4PC9Nc2dJZD4KPC94bWw+d3gyYzI3NjlmOGVmZDlhYmMy")

func TestAESUseCBCWithPKCS7_Encrypt(t *testing.T) {
	type fields struct {
		aesKey []byte
		iv     []byte
	}
	type args struct {
		plain []byte
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
				aesKey: testAESKey,
				iv:     testAESKey[:16],
			},
			args: args{plain: testPlain},
			want: testEncrypt,
		},
		{
			name: "fail aes key",
			fields: fields{
				aesKey: testAESKey[:1],
				iv:     testAESKey[:16],
			},
			args:    args{plain: testPlain},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewAESUseCBCWithPKCS7(tt.fields.aesKey, tt.fields.iv)
			got, err := w.Encrypt(tt.args.plain)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAESUseCBCWithPKCS7_Decrypt(t *testing.T) {
	type fields struct {
		aesKey []byte
		iv     []byte
	}
	type args struct {
		encrypt []byte
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
				aesKey: testAESKey,
				iv:     testAESKey[:16],
			},
			args: args{encrypt: testEncrypt},
			want: testPlain,
		},
		{
			name: "fail aes key",
			fields: fields{
				aesKey: testAESKey[:1],
				iv:     testAESKey[:16],
			},
			args:    args{encrypt: testEncrypt},
			wantErr: true,
		},
		{
			name: "fail pad",
			fields: fields{
				aesKey: testAESKey,
				iv:     testAESKey[:16],
			},
			args:    args{encrypt: testEncrypt[:32]},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewAESUseCBCWithPKCS7(tt.fields.aesKey, tt.fields.iv)
			got, err := w.Decrypt(tt.args.encrypt)
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
