package xcrypto

import (
	"errors"
)

// PKCS#7 Reference implementation
// see: https://www.rfc-editor.org/rfc/rfc2315

// ErrPKCS7BlockSizeExceeded is an error that indicates that the block size exceeds the maximum length.
var ErrPKCS7BlockSizeExceeded = errors.New("pkcs7: block size exceeded")

// PKCS7Padding pads a byte array using the PKCS#7 padding method.
//
// The src parameter is the source byte array. The blockSize parameter is the block size.
// The dst parameter is the destination byte array. The err return value is an error.
func PKCS7Padding(src []byte, blockSize int) (dst []byte, err error) {
	if blockSize < 2 || blockSize > 255 {
		return nil, ErrPKCS7BlockSizeExceeded
	}
	padLen := blockSize - len(src)%blockSize
	dst = make([]byte, len(src)+padLen)
	copy(dst, src)
	for i := len(src); i < len(dst); i++ {
		dst[i] = byte(padLen)
	}
	return dst, nil
}

// ErrPKCS7PaddingLengthExceeded is an error that indicates that the length of the padding exceeds the total length.
var ErrPKCS7PaddingLengthExceeded = errors.New("pkcs7: padding length exceeded")

// ErrPKCS7IncorrectPadding is an error that indicates the padding is incorrect.
var ErrPKCS7IncorrectPadding = errors.New("pkcs7: incorrect padding")

// PKCS7Unpadding removes the padding bytes from a byte array that was padded using the PKCS#7 padding method.
//
// The src parameter is the source byte array.
// The dst parameter is the destination byte array. The err return value is an error.
func PKCS7Unpadding(src []byte) (dst []byte, err error) {
	length := len(src)

	padLen := int(src[length-1])
	if padLen > length {
		return nil, ErrPKCS7PaddingLengthExceeded
	}

	for i := length - padLen; i < length; i++ {
		if src[i] != src[length-1] {
			return nil, ErrPKCS7IncorrectPadding
		}
	}

	dst = src[:length-padLen]
	return dst, nil
}
