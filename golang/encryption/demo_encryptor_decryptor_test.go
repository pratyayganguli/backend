package encryption

import "testing"

// will contain generic functions to encrypt and then decrypt data (symetric encryption)

type Key struct {
	KeyBytes []byte
	Iv [] byte
}

type SecureEncryptor interface {
	// the signature may be changed...
	Encrypt(plainText []byte) []byte
}

type AESEncryotor struct {
	KeySize uint
}

func (ae *AESEncryotor) Encrypt(plainText []byte) []byte {
	return nil
}

func TestEncrypt(t *testing.T) {

}

type SecureDecryptor interface {
	Decrypt(encryptText []byte) []byte
}