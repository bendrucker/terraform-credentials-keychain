package main

import (
	"os/user"

	"github.com/zalando/go-keyring"
)

func Get(hostname string) (string, error) {
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

func Store(hostname string, credentials string) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	return keyring.Set(hostname, user.Username, credentials)
}

func Forget(hostname string) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	return keyring.Delete(hostname, user.Username)
}
