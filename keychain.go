package main

import (
	"errors"
	"flag"

	"github.com/99designs/keyring"
	credentialhelper "github.com/bendrucker/terraform-credential-helper-sdk"
)

// KeychainHelper implements credentialhelper.Helper and uses the system keychain for storage.
// It supports macOS, Linux, and Windows via github.com/99designs/keyring.
type KeychainHelper struct {
	// The name of the macOS keychain where credentials will be stored
	Keychain string

	keyring keyring.Config
}

var _ credentialhelper.Helper = (*KeychainHelper)(nil)

// Flags constructs the flags supported by the helper
func (h *KeychainHelper) Flags() *flag.FlagSet {
	flags := flag.NewFlagSet("terraform-credentials-keychain", flag.ContinueOnError)
	flags.StringVar(&h.Keychain, "keychain", "login", "Name of the macOS keychain where credentials will be stored")
	return flags
}

// Open opens a keyring for Terraform
func (h *KeychainHelper) Open() (keyring.Keyring, error) {
	h.keyring.ServiceName = "terraform"
	h.keyring.KeychainName = h.Keychain

	return keyring.Open(h.keyring)
}

// Get gets the stored credentials from the keyring
func (h *KeychainHelper) Get(hostname string) ([]byte, error) {
	ring, err := h.Open()
	if err != nil {
		return nil, err
	}

	item, err := ring.Get(hostname)
	if err != nil {
		if errors.Is(err, keyring.ErrKeyNotFound) {
			return []byte("{}"), nil
		}

		return nil, err
	}

	return item.Data, nil
}

// Store stores new credentials in the keyring
func (h *KeychainHelper) Store(hostname string, credentials []byte) error {
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
func (h *KeychainHelper) Forget(hostname string) error {
	ring, err := h.Open()
	if err != nil {
		return err
	}

	return ring.Remove(hostname)
}
