module github.com/bendrucker/terraform-credentials-keychain

go 1.14

require (
	github.com/99designs/keyring v1.1.5
	github.com/bendrucker/terraform-credential-helper-sdk v0.2.1
	github.com/godbus/dbus v4.1.0+incompatible // indirect
)

replace github.com/99designs/keyring => github.com/n11c/keyring v1.1.4-0.20200205101003-eb070b66369c
