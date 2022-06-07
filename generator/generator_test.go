package generator

import (
	b64 "encoding/base64"
	"testing"
)

func TestRandomBytes(t *testing.T) {
	bytes := RandomBytes(32)
	if len(bytes) != 32 {
		t.Error("Expected 32 bytes, got", len(bytes))
	}
}

func TestRandomBytesBase64(t *testing.T) {
	base64 := RandomBytesBase64(32)

	bytes, err := b64.StdEncoding.DecodeString(base64)
	if err != nil {
		t.Error("Expected no error, got", err)
	}

	if len(bytes) != 32 {
		t.Error("Expected 32 bytes, got", len(bytes))
	}
}
