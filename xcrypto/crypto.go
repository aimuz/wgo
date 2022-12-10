package xcrypto

// Crypto is an encryption and decryption interface for simple encapsulation
type Crypto interface {
	Encrypt(plain []byte) ([]byte, error)
	Decrypt(encrypt []byte) ([]byte, error)
}
