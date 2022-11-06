package crypto_test

import (
	"testing"

	"github.com/timickb/password-manager/internal/crypto"
)

func TestStubEncrypterEncrypt(t *testing.T) {
	e := crypto.StubEncrypter{}
	if e.Encrypt("value") != "value" {
		t.Fatalf("wrong result")
	}
}

func TestStubEncrypterDecrypt(t *testing.T) {
	e := crypto.StubEncrypter{}
	if e.Decrypt("value") != "value" {
		t.Fatalf("wrong result")
	}
}
