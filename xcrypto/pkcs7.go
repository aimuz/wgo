package xcrypto

import (
	"bytes"
	"errors"
)

// PKCS7Padding ...
func PKCS7Padding(data []byte, blockSize int) ([]byte, error) {
	pl := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(pl)}, pl)...), nil
}

// ErrPKCS7IllegalFormat ...
var ErrPKCS7IllegalFormat = errors.New("pkcs7: Illegal Format")

// PKCS7Unpadding ...
func PKCS7Unpadding(data []byte) ([]byte, error) {
	l := len(data)
	lastByte := data[l-1]
	if int(lastByte) > l {
		return nil, ErrPKCS7IllegalFormat
	}
	return data[:l-int(lastByte)], nil
}
