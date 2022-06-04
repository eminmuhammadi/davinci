package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	aes "github.com/eminmuhammadi/davinci/aes"
)

const __ENCRYPTED_PRIVATE_KEY__ = "ENCRYPTED PRIVATE KEY"
const __ENCRYPTED_PUBLIC_KEY__ = "ENCRYPTED PUBLIC KEY"
const __PRIVATE_KEY__ = "RSA PRIVATE KEY"
const __PUBLIC_KEY__ = "RSA PUBLIC KEY"

// Generates a new RSA key pair.
func generateRsaKeys(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(fmt.Sprintf("failed to generate private key: %s", err))
	}

	return privKey, &privKey.PublicKey
}

// Rsa encrypts plaintext with RSA public key.
func encryptUsingPublicKey(pub *rsa.PublicKey, plaintext []byte) []byte {
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, plaintext, nil)
	if err != nil {
		panic(fmt.Sprintf("failed to encrypt plaintext: %s", err))
	}

	return ciphertext
}

// Rsa decrypts ciphertext with RSA private key.
func decryptUsingPrivateKey(priv *rsa.PrivateKey, ciphertext []byte) []byte {
	hash := sha256.New()

	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		panic(fmt.Sprintf("failed to decrypt ciphertext: %s", err))
	}

	return plaintext
}

// PemToPrivateKey converts pem string to a private key.
func PemToPrivateKey(privKey []byte, passphrase string) *rsa.PrivateKey {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(privKey); block == nil {
		panic("failed to decode PEM block containing the key")
	}

	// Decrypt the private key if a passphrase is provided
	if block.Type == __ENCRYPTED_PRIVATE_KEY__ {
		if passphrase == "" {
			panic("passphrase is required to decrypt the public key")
		}

		block.Bytes = []byte(aes.DecryptGCM256(string(block.Bytes), passphrase))
	}

	// Parse the private key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		panic(fmt.Sprintf("failed to parse private key: %s", err))
	}

	var privKeyParsed *rsa.PrivateKey
	var ok bool
	if privKeyParsed, ok = parsedKey.(*rsa.PrivateKey); !ok {
		panic("failed to parse private key")
	}

	return privKeyParsed
}

// PemToPublicKey converts pem string to a public key.
func PemToPublicKey(pubKey []byte, passphrase string) *rsa.PublicKey {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(pubKey); block == nil {
		panic("failed to parse PEM block containing the key")
	}

	// Decrypt the private key if a passphrase is provided
	if block.Type == __ENCRYPTED_PUBLIC_KEY__ {
		if passphrase == "" {
			panic("passphrase is required to decrypt the public key")
		}

		block.Bytes = []byte(aes.DecryptGCM256(string(block.Bytes), passphrase))
	}

	// Parse the public key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		panic(fmt.Sprintf("failed to parse public key: %s", err))
	}

	var publicKeyParsed *rsa.PublicKey
	var ok bool
	if publicKeyParsed, ok = parsedKey.(*rsa.PublicKey); !ok {
		panic("key is not an RSA public key")
	}

	return publicKeyParsed
}

// PrivateKeyToBytes converts a private key to bytes.
func PrivateKeyToBytes(privKey *rsa.PrivateKey, passphrase string) []byte {
	privASN1, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		panic(fmt.Sprint("marshal private key:", err))
	}

	block := &pem.Block{
		Type:  __PRIVATE_KEY__,
		Bytes: privASN1,
	}

	// Encrypt the private key if a passphrase is provided
	if passphrase != "" {
		block = &pem.Block{
			Type: __ENCRYPTED_PRIVATE_KEY__,
			Headers: map[string]string{
				"Proc-Type": "4,ENCRYPTED",
				"DEK-Info":  fmt.Sprintf("AES-256-GCM,%x", sha256.Sum256([]byte(passphrase))),
			},
			Bytes: []byte(aes.EncryptGCM256(string(privASN1), passphrase)),
		}
	}

	return pem.EncodeToMemory(block)
}

// PublicKeyToBytes converts a public key to bytes.
func PublicKeyToBytes(pub *rsa.PublicKey, passphrase string) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		panic(fmt.Sprint("marshal public key:", err))
	}

	block := &pem.Block{
		Type:  __PUBLIC_KEY__,
		Bytes: pubASN1,
	}

	// Encrypt the private key if a passphrase is provided
	if passphrase != "" {
		block = &pem.Block{
			Type: __ENCRYPTED_PUBLIC_KEY__,
			Headers: map[string]string{
				"Proc-Type": "4,ENCRYPTED",
				"DEK-Info":  fmt.Sprintf("AES-256-GCM,%x", sha256.Sum256([]byte(passphrase))),
			},
			Bytes: []byte(aes.EncryptGCM256(string(pubASN1), passphrase)),
		}
	}

	return pem.EncodeToMemory(block)
}
