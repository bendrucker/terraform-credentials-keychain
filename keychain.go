package main

import (
	"flag"

	"github.com/99designs/keyring"
	credentialhelper "github.com/bendrucker/terraform-credential-helper-sdk"
)

// KeychainHelper implements credentialhelper.Helper and uses the system keychain for storage.
// It supports macOS, Linux, and Windows via github.com/99designs/keyring.
type KeychainHelper struct{}

var _ credentialhelper.Helper = (*KeychainHelper)(nil)

// Open opens a keyring for Terraform
func (h *KeychainHelper) Open() (keyring.Keyring, error) {
	return keyring.Open(keyring.Config{
		ServiceName: "terraform",
	})
}

// Get gets the stored credentials from the keyring
func (h *KeychainHelper) Get(hostname string, flags *flag.FlagSet) ([]byte, error) {
	ring, err := h.Open()
	if err != nil {
		return nil, err
	}

	item, err := ring.Get(hostname)
	if err != nil {
		return nil, err
	}

	return item.Data, nil
}

// Store stores new credentials in the keyring
func (h *KeychainHelper) Store(hostname string, credentials []byte, flags *flag.FlagSet) error {
	ring, err := h.Open()
	if err != nil {
		return err
	}

	return ring.Set(keyring.Item{
		Key:         hostname,
		Label:       hostname,
		Data:        credentials,
		Description: "application password",
	})
}

// Forget deletes existing credentials from the keyring
func (h *KeychainHelper) Forget(hostname string, flags *flag.FlagSet) error {
	ring, err := h.Open()
	if err != nil {
		return err
	}

	return ring.Remove(hostname)
}
