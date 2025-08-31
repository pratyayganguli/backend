package security

import (
	"fmt"
	"testing"
)

func TestNewAESKey(t *testing.T) {
	key := &AESKey{
		Size: AES_128,
	}
	if e := key.Generate(); e == nil {
		fmt.Println("Key Generation Passed!")
	}
}

func TestAESEncryptDecryptFlow(t *testing.T) {
	key := &AESKey{
		Size: AES_128,
	}
	if e := key.Generate(); e == nil {
		var se SecureEncryptor = &AESEncryptor{
			AESKey: *key,
		}
		if ct, err := se.Encrypt([]byte("Hi, is the connection secure?")); err == nil {
			var sd SecureDecryptor[string] = &AESDecryptor[string]{
				AESKey: *key,
			}
			if _, err := sd.Decrypt(ct); err == nil {
				fmt.Println("success")
			} else {
				fmt.Printf("fail: %s", err.Error())
			}
		} else {
			fmt.Printf("fail: %s", err.Error())
		}
	}
}
