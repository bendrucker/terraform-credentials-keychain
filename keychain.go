package main

import (
	"os/user"

	"github.com/zalando/go-keyring"
)

type KeychainHelper struct{}

func (h *KeychainHelper) Get(hostname string) (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	secret, err := keyring.Get(hostname, user.Username)
	if err != nil {
		return "", err
	}

	return secret, nil
}

func (h *KeychainHelper) Store(hostname string, credentials string) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	return keyring.Set(hostname, user.Username, credentials)
}

func (h *KeychainHelper) Forget(hostname string) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	return keyring.Delete(hostname, user.Username)
}
