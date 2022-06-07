package aes

import (
	"testing"

	generator "github.com/eminmuhammadi/davinci/generator"
)

func TestEncryptAndDecrypt(t *testing.T) {
	key := generator.RandomBytesBase64(32)
	data := "Hello World!"

	ciphertext := EncryptGCM256(data, key)
	plaintext := DecryptGCM256(ciphertext, key)

	if plaintext != data {
		t.Error("Expected plaintext to be", data, ", got", plaintext)
	}
}
