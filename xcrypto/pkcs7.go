package xcrypto

import (
	"bytes"
	"errors"
)

// ErrPKCS7BlockSizeExceeded is an error that indicates that the block size exceeds the maximum length.
var ErrPKCS7BlockSizeExceeded = errors.New("pkcs7: block size exceeded")

// PKCS7Padding pads a byte array using the PKCS#7 padding method.
//
// The src parameter is the source byte array. The blockSize parameter is the block size.
// The dst parameter is the destination byte array. The err return value is an error.
func PKCS7Padding(src []byte, blockSize int) (dst []byte, err error) {
	if blockSize > 255 {
		return nil, ErrPKCS7BlockSizeExceeded
	}
	pl := blockSize - len(src)%blockSize
	if pl == blockSize {
		return src, nil
	}
	padding := bytes.Repeat([]byte{byte(pl)}, pl)
	dst = append(src, padding...)
	return dst, nil
}

// ErrPKCS7PaddingLengthExceeded is an error that indicates that the length of the padding exceeds the total length.
var ErrPKCS7PaddingLengthExceeded = errors.New("pkcs7: padding length exceeded")

// PKCS7Unpadding removes the padding bytes from a byte array that was padded using the PKCS#7 padding method.
//
// The src parameter is the source byte array.
// The dst parameter is the destination byte array. The err return value is an error.
func PKCS7Unpadding(src []byte) (dst []byte, err error) {
	l := len(src)
	lastByte := src[l-1]
	if int(lastByte) > l {
		return nil, ErrPKCS7PaddingLengthExceeded
	}
	return src[:l-int(lastByte)], nil
}
