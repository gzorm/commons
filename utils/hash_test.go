package utils

import (
	"crypto/sha256"
	"testing"
)

func TestHash(t *testing.T) {
	hash := sha256.New()
	hash.Write([]byte("33333"))
	t.Logf("%x", string(hash.Sum(nil)))
	hash.Write([]byte("333332"))
	t.Logf("%x", string(hash.Sum(nil)))
}

func TestParse(t *testing.T) {
}
