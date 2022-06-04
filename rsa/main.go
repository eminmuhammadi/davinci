package rsa

import (
	b64 "encoding/base64"
	"fmt"
)

type KeyPair struct {
	PrivateKey string
	PublicKey  string
}

// Generates a new RSA key pair.
func GenerateKeyPair(keySize int, passphrase string) KeyPair {
	privKey, publicKey := generateRsaKeys(keySize)

	privKeyByte := PrivateKeyToBytes(privKey, passphrase)
	publicKeyByte := PublicKeyToBytes(publicKey, passphrase)

	// Keys in PEM format
	return KeyPair{
		PrivateKey: string(privKeyByte),
		PublicKey:  string(publicKeyByte),
	}
}

// RSA Encrypt
func Encrypt(publicKey string, plainText string, passphrase string) string {
	pubKey := PemToPublicKey([]byte(publicKey), passphrase)

	// Encrypt plaintext
	ciphertext := encryptUsingPublicKey(pubKey, []byte(plainText))

	// Convert ciphertext to base64 string.
	return b64.StdEncoding.EncodeToString(ciphertext)
}

// RSA Decrypt
func Decrypt(privateKey string, cipherText string, passphrase string) string {
	privKey := PemToPrivateKey([]byte(privateKey), passphrase)

	// Convert ciphertext from base64 string.
	ciphertext, err := b64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		panic(fmt.Sprintf("failed to decode ciphertext: %s", err))
	}

	// Decrypt ciphertext
	plaintext := decryptUsingPrivateKey(privKey, ciphertext)

	return string(plaintext)
}
