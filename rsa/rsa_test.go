package rsa

import (
	"testing"

	generator "github.com/eminmuhammadi/davinci/generator"
)

func TestGenerateKeyPairWithoutPassphrase(t *testing.T) {
	keyPair := GenerateKeyPair(2048, "")

	if len(keyPair.PrivateKey) == 0 {
		t.Error("Expected private key, got empty string")
	}

	if len(keyPair.PublicKey) == 0 {
		t.Error("Expected public key, got empty string")
	}
}

func TestGenerateKeyPair(t *testing.T) {
	passphrase := generator.RandomBytesBase64(32)
	keyPair := GenerateKeyPair(2048, passphrase)

	if len(keyPair.PrivateKey) == 0 {
		t.Error("Expected private key, got empty string")
	}

	if len(keyPair.PublicKey) == 0 {
		t.Error("Expected public key, got empty string")
	}
}

func TestEncryptAndDecrypt(t *testing.T) {
	passphrase := generator.RandomBytesBase64(32)
	keyPair := GenerateKeyPair(2048, passphrase)
	data := "Hello World!"

	ciphertext := Encrypt(keyPair.PublicKey, data, passphrase)
	plaintext := Decrypt(keyPair.PrivateKey, ciphertext, passphrase)

	if plaintext != data {
		t.Error("Expected plaintext to be", data, ", got", plaintext)
	}
}
