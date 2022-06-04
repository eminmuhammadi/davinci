package aes

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"

	generator "github.com/eminmuhammadi/davinci/generator"
)

const __NONCE_SIZE__ = 12

// Encrypts plaintext with AES-256-GCM.
func EncryptGCM256(plaintext string, keyString string) string {
	var err error

	// Convert key from base64 string.
	key, err := b64.StdEncoding.DecodeString(keyString)
	if err != nil {
		panic(err)
	}

	// Convert plaintext to byte array.
	plaintextBytes := []byte(plaintext)

	// AES-256-GCM.
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Nonce
	nonce := generator.RandomBytes(__NONCE_SIZE__)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	// Encrypt.
	ciphertext := aesgcm.Seal(nil, nonce, plaintextBytes, nil)

	// Convert ciphertext to base64 string.
	return b64.StdEncoding.EncodeToString(append(nonce, ciphertext...))
}

// Decrypts ciphertext with AES-256-GCM.
func DecryptGCM256(cipherText string, keyString string) string {
	var err error

	// Convert key from base64 string.
	key, err := b64.StdEncoding.DecodeString(keyString)
	if err != nil {
		panic(err)
	}

	// Convert ciphertext from base64 string.
	ciphertext, err := b64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	// AES-256-GCM.
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	// Nonce and ciphertext.
	nonce, ciphertext := ciphertext[:__NONCE_SIZE__], ciphertext[__NONCE_SIZE__:]

	// Decrypt.
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	// Convert plaintext to string.
	return string(plaintext)
}
