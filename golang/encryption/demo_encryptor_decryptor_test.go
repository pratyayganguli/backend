package encryption

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"testing"
)

func AESKeyGenerator(keysize *uint) ([]byte, error) {
	if keysize != nil {
		// generate 128 bit key
		randomBytes := make([]byte, 16)
		if _, err := io.ReadFull(rand.Reader, randomBytes); err == nil {
			return randomBytes, nil
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid key size")
	}
}

func TestAESKeyGenerator(t *testing.T) {
	keySize := uint(16)
	if key, err := AESKeyGenerator(&keySize); err == nil {
		b64Key := base64.StdEncoding.EncodeToString(key)
		fmt.Println(b64Key)
	} else {
		fmt.Errorf("Test case failed %s", err.Error())
	}
}

// will contain generic functions to encrypt and then decrypt data (symetric encryption)

type Key struct {
	KeyBytes []byte
	Iv       []byte
}

type SecureEncryptor interface {
	// the signature may be changed...
	Encrypt(plainText []byte) []byte
}

// supported bits are 128, 192, 256
type AESEncryptor struct {
	KeySize uint
}

func (ae *AESEncryptor) Encrypt(plainText []byte) []byte {
	return nil
}

func TestEncrypt(t *testing.T) {
	message := "Hi, how are you?"
	ae := &AESEncryptor{
		KeySize: 128,
	}
	cipher := ae.Encrypt([]byte(message))
	b64Cipher := base64.StdEncoding.EncodeToString(cipher)
	fmt.Println(b64Cipher)
}

type SecureDecryptor interface {
	Decrypt(encryptText []byte) []byte
}
