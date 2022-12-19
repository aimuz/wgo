package xcrypto

import (
	"reflect"
	"testing"
)

func TestPKCS7Padding(t *testing.T) {
	type args struct {
		src       []byte
		blockSize int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "no padding",
			args: args{
				src:       []byte{1, 2, 3, 4, 5, 6, 7, 8},
				blockSize: 8,
			},
			want: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "blockSize is 8",
			args: args{
				src:       []byte{1, 2, 3},
				blockSize: 8,
			},
			want: []byte{1, 2, 3, 5, 5, 5, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PKCS7Padding(tt.args.src, tt.args.blockSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("PKCS7Padding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS7Padding() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPKCS7Unpadding(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "blockSize is 8",
			args: args{
				src: []byte{1, 2, 3, 5, 5, 5, 5, 5},
			},
			want: []byte{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := PKCS7Unpadding(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("PKCS7Unpadding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDst, tt.want) {
				t.Errorf("PKCS7Unpadding() got = %v, want %v", gotDst, tt.want)
			}
		})
	}
}
