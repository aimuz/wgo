package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
)

// AESUseCBCWithPKCS7 This will be filled using PKCS7 and finally encrypted using AES CBC mode
type AESUseCBCWithPKCS7 struct {
	key []byte
	iv  []byte
}

// NewAESUseCBCWithPKCS7 ...
func NewAESUseCBCWithPKCS7(key, iv []byte) *AESUseCBCWithPKCS7 {
	return &AESUseCBCWithPKCS7{
		key: key,
		iv:  iv,
	}
}

// Encrypt will be padded using PKCS7 and finally encrypted using AES CBC mode
func (w *AESUseCBCWithPKCS7) Encrypt(plain []byte) ([]byte, error) {
	c, err := w.newAESBlack()
	if err != nil {
		return nil, err
	}

	const blockSize = 32
	msg, err := PKCS7Padding(plain, blockSize)
	if err != nil {
		return nil, err
	}
	encrypter := cipher.NewCBCEncrypter(c, w.iv)
	encrypter.CryptBlocks(msg, msg)
	return msg, nil
}

// Decrypt will be padded using PKCS7 and finally decrypted using AES CBC mode
func (w *AESUseCBCWithPKCS7) Decrypt(encrypt []byte) ([]byte, error) {
	c, err := w.newAESBlack()
	if err != nil {
		return nil, err
	}

	encrypter := cipher.NewCBCDecrypter(c, w.iv)
	encrypter.CryptBlocks(encrypt, encrypt)

	plain, err := PKCS7Unpadding(encrypt)
	if err != nil {
		return nil, err
	}
	return plain, nil
}

func (w *AESUseCBCWithPKCS7) newAESBlack() (cipher.Block, error) {
	return aes.NewCipher(w.key)
}
