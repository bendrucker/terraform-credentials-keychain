package credentialhelper

// Helper represents a Terraform credential helper
// https://www.terraform.io/docs/internals/credentials-helpers.html
type Helper interface {
	// Retrieve the credentials for the given hostname
	Get(hostname string) (credentials []byte, err error)
	// Store new credentials for the given hostname
	Store(hostname string, credentials []byte) error
	// Delete any stored credentials for the given hostname
	Forget(hostname string) error
}
