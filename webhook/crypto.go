package webhook

import (
	"encoding/xml"
	"errors"
	"strconv"

	"github.com/aimuz/wgo/xcrypto"
)

// WXBizMsgCrypto ...
type WXBizMsgCrypto interface {
	DecryptMessage(payload []byte, v interface{}) error
	EncryptMessage(payload []byte, timestamp int64, nonce string) ([]byte, error)
}

// WXBizMsgCrypt ...
type WXBizMsgCrypt struct {
	appID string
	token string

	crypto xcrypto.Crypto
}

// NewWXBizMsgCrypt ...
func NewWXBizMsgCrypt(encodingAESKey, appID, token string) (*WXBizMsgCrypt, error) {
	c, err := xcrypto.NewWXCrypt(encodingAESKey, appID)
	if err != nil {
		return nil, err
	}
	return &WXBizMsgCrypt{
		appID:  appID,
		token:  token,
		crypto: c,
	}, nil
}

// DecryptMessage ...
func (wc *WXBizMsgCrypt) DecryptMessage(payload []byte, v interface{}) error {
	if wc.crypto == nil {
		return errors.New("crypto: wc.crypto not init")
	}
	// decrypt payload
	payload, err := wc.crypto.Decrypt(payload)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(payload, v)
	if err != nil {
		return err
	}
	return nil
}

// EncryptMessage ...
func (wc *WXBizMsgCrypt) EncryptMessage(payload []byte, timestamp int64, nonce string) ([]byte, error) {
	sign := SHA1Sign(strconv.FormatInt(timestamp, 10), wc.token, nonce, string(payload))
	msg := &EncryptMessage{
		Encrypt:      string(payload),
		MsgSignature: sign,
		TimeStamp:    timestamp,
		Nonce:        nonce,
	}
	return xml.Marshal(msg)
}
