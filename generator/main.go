package generator

import (
	"crypto/rand"
	b64 "encoding/base64"
)

// Generates a random 32-bit string.
func RandomBytesBase64(size int) string {
	bytes := RandomBytes(32)

	return b64.StdEncoding.EncodeToString(bytes)
}

// Generates random 12 bit
func RandomBytes(size int) []byte {
	bytes := make([]byte, size)
	rand.Read(bytes)

	return bytes
}
