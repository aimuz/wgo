package xcrypto

import "bytes"

// PKCS7Padding ...
func PKCS7Padding(data []byte, blockSize int) ([]byte, error) {
	pl := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(pl)}, pl)...), nil
}

// PKCS7Unpadding ...
func PKCS7Unpadding(data []byte) ([]byte, error) {
	return data[:len(data)-int(data[len(data)-1])], nil
}
