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

// NewAESUseCBCWithPKCS7 creates a new AESUseCBCWithPKCS7 structure.
//
// The key parameter is the key. The iv parameter is the initialization vector.
// The function returns a pointer to a new AESUseCBCWithPKCS7 structure.
func NewAESUseCBCWithPKCS7(key, iv []byte) *AESUseCBCWithPKCS7 {
	return &AESUseCBCWithPKCS7{
		key: key,
		iv:  iv,
	}
}

// Encrypt will be padded using PKCS7 and finally encrypted using AES CBC mode
func (w *AESUseCBCWithPKCS7) Encrypt(src []byte) (dst []byte, err error) {
	c, err := w.newAESBlack()
	if err != nil {
		return nil, err
	}

	var blockSize = c.BlockSize()
	src, err = PKCS7Padding(src, blockSize)
	if err != nil {
		return nil, err
	}
	encrypter := cipher.NewCBCEncrypter(c, w.iv)
	encrypter.CryptBlocks(dst, src)
	return dst, nil
}

// Decrypt will be padded using PKCS7 and finally decrypted using AES CBC mode
func (w *AESUseCBCWithPKCS7) Decrypt(src []byte) (dst []byte, err error) {
	c, err := w.newAESBlack()
	if err != nil {
		return nil, err
	}

	dst = make([]byte, len(src))
	decrypter := cipher.NewCBCDecrypter(c, w.iv)
	decrypter.CryptBlocks(dst, src)

	dst, err = PKCS7Unpadding(dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (w *AESUseCBCWithPKCS7) newAESBlack() (cipher.Block, error) {
	return aes.NewCipher(w.key)
}
