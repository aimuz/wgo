// Copyright 2013-2014 Vasiliy Gorin. All rights reserved.
// Use of this source code is governed by a GNU-style
// license that can be found in the LICENSE file.

// Copy from https://raw.githubusercontent.com/vgorin/cryptogo/master/pad/pad.go

package xcrypto

import (
	"bytes"
	"errors"
	"fmt"
)

// ErrBlockSizeSmall ...
var ErrBlockSizeSmall = errors.New("block size is too small (minimum is 2 bytes")

// ErrSupportedBlockSize ...
var ErrSupportedBlockSize = errors.New("unsupported block size")

// PKCS7Pad adds PKCS7 padding to the data block, http://en.wikipedia.org/wiki/Padding_(cryptography)#PKCS7
func PKCS7Pad(message []byte, blockSize int) ([]byte, error) {
	// block size must be bigger or equal 2
	if blockSize < 1<<1 {
		return nil, ErrBlockSizeSmall
	}
	// block size up to 255 requires 1 byte padding
	if blockSize < 1<<8 {
		// calculate padding length
		padlen := PadLength(len(message), blockSize)

		// define PKCS7 padding block
		padding := bytes.Repeat([]byte{byte(padlen)}, padlen)

		// apply padding
		padded := append(message, padding...)
		return padded, nil
	}
	// block size bigger or equal 256 is not currently supported
	return nil, ErrSupportedBlockSize
}

// ErrInvalidPadding ...
var ErrInvalidPadding = errors.New("invalid padding")

// PKCS7Unpad removes PKCS7 padding from the data block, http://en.wikipedia.org/wiki/Padding_(cryptography)#PKCS7
// this function may return an error id padding is incorrect,
// however it will return unpadded data in any case
func PKCS7Unpad(padded []byte) ([]byte, error) {
	// read padding length
	pLen := len(padded)
	lastByte := padded[pLen-1]
	padLen := int(lastByte)

	if pLen < padLen {
		return nil, ErrInvalidPadding
	}

	// check validity of PKCS7 padding
	for i := padLen; i > 1; i-- {
		if padded[pLen-i] != lastByte {
			err := fmt.Errorf("invalid padding (byte -%d: %d). Is the message supplied PKCS7 padded", i, padded[pLen-i])
			return nil, err
		}
	}

	// remove padding
	return padded[:pLen-padLen], nil
}

// PadLength calculates padding length
func PadLength(sliceLength, blockSize int) (padLen int) {
	padLen = blockSize - sliceLength%blockSize
	if padLen == 0 {
		padLen = blockSize
	}
	return padLen
}
