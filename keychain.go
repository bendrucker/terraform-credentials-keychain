package main

import (
	"fmt"

	"github.com/99designs/keyring"
)

type KeychainHelper struct{}

func (h *KeychainHelper) Open() (keyring.Keyring, error) {
	return keyring.Open(keyring.Config{ServiceName: "terraform"})
}

func (h *KeychainHelper) Get(hostname string) ([]byte, error) {
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

func (h *KeychainHelper) Store(hostname string, credentials []byte) error {
	ring, err := h.Open()
	if err != nil {
		return err
	}

	return ring.Set(keyring.Item{
		Key:         hostname,
		Data:        credentials,
		Description: fmt.Sprintf("Terraform Cloud API credentials for %s, created via 'terraform login'", hostname),
	})
}

func (h *KeychainHelper) Forget(hostname string) error {
	ring, err := h.Open()
	if err != nil {
		return err
	}

	return ring.Remove(hostname)
}
