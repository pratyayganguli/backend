package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type SecureEncryptor interface {
	Encrypt(plainText []byte) (*CipherText, error)
}

type AESEncryptor struct {
	AESKey AESKey
}

type CipherText struct {
	Cipher []byte
	Nonce  []byte
}

// type RSAEncryptor struct {
// 	RSAKey RSAKey
// }

func (ae *AESEncryptor) Encrypt(plainText []byte) (*CipherText, error) {
	if ae.AESKey.KeyBytes == nil {
		return nil, fmt.Errorf("key cannot be null or empty")
	}
	if block, err := aes.NewCipher(ae.AESKey.KeyBytes); err == nil {
		if gcm, err := cipher.NewGCM(block); err == nil {
			nonceBytes := make([]byte, gcm.NonceSize())
			if _, err := io.ReadFull(rand.Reader, nonceBytes); err == nil {
				cipher := gcm.Seal(nil, nonceBytes, plainText, nil)
				ct := &CipherText{
					Cipher: cipher,
					Nonce:  nonceBytes,
				}
				return ct, nil
			} else {
				return nil, fmt.Errorf("could not read nonce: %s", err.Error())
			}
		} else {
			return nil, fmt.Errorf("encryption error, failed to create GCM: %s", err.Error())
		}

	} else {
		return nil, fmt.Errorf("encryption error: %s", err.Error())
	}
}
