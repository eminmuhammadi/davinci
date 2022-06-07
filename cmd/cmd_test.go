package cmd

import (
	"testing"

	generator "github.com/eminmuhammadi/davinci/generator"
	rsa "github.com/eminmuhammadi/davinci/rsa"
)

func TestDecryptAndEncryptAction(t *testing.T) {
	passphrase := generator.RandomBytesBase64(32)
	key := generator.RandomBytesBase64(32)
	keyPair := rsa.GenerateKeyPair(2048, passphrase)
	input := []byte("Hello World!")

	ciphertext := EncryptAction([]byte(keyPair.PublicKey), []byte(passphrase), []byte(key), input)
	plaintext := DecryptAction([]byte(keyPair.PrivateKey), []byte(passphrase), []byte(ciphertext))

	if string(plaintext) != string(input) {
		t.Errorf("Expected %s, got %s", input, plaintext)
	}
}

func TestKeyPair(t *testing.T) {
	passphrase := generator.RandomBytesBase64(32)
	keys := KeyPairAction([]byte(passphrase), 2048)

	if len(keys.PrivateKey) == 0 {
		t.Errorf("Expected private key, got %s", keys.PrivateKey)
	}

	if len(keys.PublicKey) == 0 {
		t.Errorf("Expected public key, got %s", keys.PublicKey)
	}
}

func TestPassphraseAction(t *testing.T) {
	passphrase := PassphraseAction()
	if len(passphrase) == 0 {
		t.Errorf("Expected passphrase, got %s", passphrase)
	}
}

func TestKeyAction(t *testing.T) {
	key := KeyAction()
	if len(key) == 0 {
		t.Errorf("Expected key, got %s", key)
	}
}
