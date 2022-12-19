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
			name: "8 padding",
			args: args{
				src:       []byte{1, 2, 3, 4, 5, 6, 7, 8},
				blockSize: 8,
			},
			want: []byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		},
		{
			name: "5 padding",
			args: args{
				src:       []byte{1, 2, 3},
				blockSize: 8,
			},
			want: []byte{1, 2, 3, 5, 5, 5, 5, 5},
		},
		{
			name: "full padding",
			args: args{
				src:       []byte{},
				blockSize: 8,
			},
			want: []byte{8, 8, 8, 8, 8, 8, 8, 8},
		},
		{
			name: "1 padding",
			args: args{
				src:       []byte{1, 2, 3, 4, 5, 6, 7},
				blockSize: 8,
			},
			want: []byte{1, 2, 3, 4, 5, 6, 7, 1},
		},
		{
			name: "256 blockSize",
			args: args{
				src:       []byte{1, 2, 3, 4, 5, 6, 7},
				blockSize: 256,
			},
			wantErr: true,
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
			name: "8 padding",
			args: args{
				src:       []byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8},
				blockSize: 8,
			},
			want: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "3 padding",
			args: args{
				src:       []byte{1, 2, 3, 5, 5, 5, 5, 5},
				blockSize: 8,
			},
			want: []byte{1, 2, 3},
		},
		{
			name: "Invalid padding length",
			args: args{
				src:       []byte{1, 2, 3, 4, 5, 6, 7, 9},
				blockSize: 8,
			},
			wantErr: true,
		},
		{
			name: "incorrect padding",
			args: args{
				src:       []byte{1, 2, 3, 4, 5, 6, 7, 8},
				blockSize: 8,
			},
			wantErr: true,
		},
		{
			name: "0 body",
			args: args{
				src:       []byte{8, 8, 8, 8, 8, 8, 8, 8},
				blockSize: 8,
			},
			want:    []byte{},
			wantErr: false,
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
