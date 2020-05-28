package main

import (
	"testing"

	"github.com/99designs/keyring"
)

func TestKeychain(t *testing.T) {
	helper := &KeychainHelper{
		Keychain: "test-terraform-credentials-keychain",
		keyring: keyring.Config{
			KeychainPasswordFunc: func(prompt string) (string, error) {
				return "", nil
			},
		},
	}

	if err := helper.Store("test.terraform.io", []byte(`{"token":"secret"}`)); err != nil {
		t.Fatalf("failed to store: %v", err)
	}

	b, err := helper.Get("test.terraform.io")
	if err != nil {
		t.Fatalf("failed to get: %v", err)
	}

	expected := `{"token":"secret"}`
	if got := string(b); got != expected {
		t.Errorf("wrong credentials, expected %s, got %s", expected, got)
	}

	if err := helper.Forget("test.terraform.io"); err != nil {
		t.Fatalf("failed to forget: %v", err)
	}
}
