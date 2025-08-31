package security

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type SecureDecryptor interface {
	Decrypt(*CipherText) ([]byte, error)
}

type AESDecryptor struct {
	AESKey AESKey
}

func (ad *AESDecryptor) Decrypt(cipherText *CipherText) ([]byte, error) {
	if ad.AESKey.KeyBytes != nil {
		if block, err := aes.NewCipher(ad.AESKey.KeyBytes); err == nil {
			if gcm, err := cipher.NewGCM(block); err == nil {
				if plainBytes, err := gcm.Open(nil, cipherText.Nonce, cipherText.Cipher, nil); err == nil {
					return plainBytes, nil
				} else {
					return nil, fmt.Errorf("could not decrypt data: %s", err.Error())
				}
			} else {
				return nil, fmt.Errorf("could not create new gcm: %s", err.Error())
			}
		} else {
			return nil, fmt.Errorf("cipher block could not be created: %s", err.Error())
		}
	} else {
		return nil, fmt.Errorf("key cannot be empty or nil")
	}
}
