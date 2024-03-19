package validation

// SecretKeyStore represents a store of secret keys.
// Implementations could store secret keys in a database, on disk, etc.
// Thread-safe.
type SecretKeyStore interface {
	// GetSecretKey returns the secretKey for the given keyID
	//
	// Can return any of the following errors:
	// SecretKeyNotAvailableError if there is no secretKey with the given keyID
	GetSecretKey(string) (string, error)
}
