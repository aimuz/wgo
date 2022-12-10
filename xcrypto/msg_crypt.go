package xcrypto

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
)

// WXBizMsgCrypt ...
type WXBizMsgCrypt struct {
	aesCrypto Crypto
	appID     string
}

// NewWXBizMsgCrypt ...
func NewWXBizMsgCrypt(encodingAESKey, appID string) (*WXBizMsgCrypt, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return nil, err
	}

	aesCrypto := NewAESUseCBCWithPKCS7(aesKey, aesKey[:16])
	wx := &WXBizMsgCrypt{
		aesCrypto: aesCrypto,
		appID:     appID,
	}
	return wx, err
}

var binaryEndian = binary.BigEndian

// Encrypt ...
func (w *WXBizMsgCrypt) Encrypt(plain []byte) ([]byte, error) {
	fullLen := 20 + len(plain) + len(w.appID)
	full := make([]byte, fullLen)
	_, err := rand.Read(full[:16])
	if err != nil {
		return nil, err
	}
	binaryEndian.PutUint32(full[16:20], uint32(len(plain)))
	copy(full[20:], plain)
	copy(full[20+len(plain):], w.appID)
	return w.aesCrypto.Encrypt(full)
}

// Decrypt ...
func (w *WXBizMsgCrypt) Decrypt(encrypt []byte) ([]byte, error) {
	plain, err := w.aesCrypto.Decrypt(encrypt)
	if err != nil {
		return nil, err
	}
	Len := plain[16:20]
	msgLen := int(binaryEndian.Uint32(Len))
	return plain[20 : msgLen+20], nil
}
