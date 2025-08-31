package security

import (
	"crypto/rand"
	"fmt"
	"io"
)

const AES_128 uint = 16
const AES_192 uint = 24
const AES_256 uint = 32

type AESKey struct {
	Size     uint
	KeyBytes []byte
}

func (k *AESKey) Generate() error {
	if e := validateKeySize(k.Size); e == nil {
		secureRandomBytes := make([]byte, k.Size)
		// we are not using the number of bytes anywhere while, generating the key
		if _, err := io.ReadFull(rand.Reader, secureRandomBytes); err == nil {
			k.KeyBytes = secureRandomBytes
			return nil
		} else {
			return err
		}
	} else {
		return e
	}
}

func validateKeySize(size uint) error {
	if size != AES_128 && size != AES_192 && size != AES_256 {
		return fmt.Errorf("invalid key size: %d", size)
	}
	return nil
}
