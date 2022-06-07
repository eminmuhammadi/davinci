package generator

import (
	"crypto/rand"
	b64 "encoding/base64"
)

// Generates Random bytes in base64
func RandomBytesBase64(size int) string {
	bytes := RandomBytes(32)

	return b64.StdEncoding.EncodeToString(bytes)
}

// Generates Random bytes
func RandomBytes(size int) []byte {
	bytes := make([]byte, size)
	rand.Read(bytes)

	return bytes
}
