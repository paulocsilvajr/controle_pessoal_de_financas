package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSenhaSha256(senha string) string {
	senhaSha256 := sha256.Sum256([]byte(senha))

	dst := make([]byte, hex.EncodedLen(len(senhaSha256)))
	hex.Encode(dst, senhaSha256[:])

	return string(dst[:len(dst)])
}
