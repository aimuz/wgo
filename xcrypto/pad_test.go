// Copyright 2013-2014 Vasiliy Gorin. All rights reserved.
// Use of this source code is governed by a GNU-style
// license that can be found in the LICENSE file.

// Copy from https://raw.githubusercontent.com/vgorin/cryptogo/master/pad/pad_test.go

package xcrypto

import "testing"

func TestPKCS7Padding(t *testing.T) {
	msg := "this is my test message of length 36"
	msgBytes := []byte(msg)
	t.Logf("message (len=%d): %s", len(msgBytes), msg)
	paddedBytes, err := PKCS7Pad(msgBytes, 17)
	if err != nil {
		t.Error(err)
	}
	padded := string(paddedBytes)
	t.Logf("padded (len=%d): %s", len(paddedBytes), padded)
	t.Logf("padded bytes (len=%d): %v", len(paddedBytes), paddedBytes)
	originalBytes, err := PKCS7Unpad(paddedBytes)
	if err != nil {
		t.Error(err)
	}
	original := string(originalBytes)
	t.Logf("unpadded: %s", original)

	paddedBytes[len(paddedBytes)-5] = 77
	originalBytes, err = PKCS7Unpad(paddedBytes)
	t.Logf("expected error: %s", err)
	if err == nil {
		t.Error("expected error but got nil")
	}
}
