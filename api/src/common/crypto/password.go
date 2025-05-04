package crypto

import (
	"encoding/hex"
	"unsafe"

	"golang.org/x/crypto/scrypt"
)

func HashPassword(password string, salt string) (string, error) {
	const keyLen = 64
	derivedKey, err := scrypt.Key(zcString2Bytes(password), zcString2Bytes(salt), 16384, 8, 1, keyLen)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(derivedKey), nil
}

// Copy string to bytes conversion
func zcString2Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
