package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

type Row[T any] struct {
	Key   string
	Value T
}

func DummyRow() *Row[string] {
	return &Row[string]{
		Key:   "secret",
		Value: "password123",
	}
}

func EncryptRow(row *Row[string]) (string, []byte, []byte) {
	keyBytes := make([]byte, 32)
	if _, err := rand.Read(keyBytes); err != nil {
		panic(err)
	}
	if block, err := aes.NewCipher(keyBytes); err == nil {
		if gcm, err := cipher.NewGCM(block); err != nil {
			panic(err)
		} else {
			nonceBytes := make([]byte, gcm.NonceSize())
			if _, err := io.ReadFull(rand.Reader, nonceBytes); err == nil {
				cipherText := gcm.Seal(nil, nonceBytes, []byte(row.Value), nil)
				b64cipherText := base64.StdEncoding.EncodeToString(cipherText)
				log.Printf("Encrypted Data: %s", b64cipherText)
				return b64cipherText, keyBytes, nonceBytes
			} else {
				panic(err)
			}
		}
	} else {
		panic(err)
	}

}

func DecryptRow(b64CipherText string, nonce, key []byte) error {
	return nil
}
