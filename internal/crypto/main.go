package crypto

import (
	"crypto/rand"
)

func RandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
