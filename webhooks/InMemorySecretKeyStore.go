package webhooks

import (
	"errors"
	"strings"
	"sync"

	"github.com/Worldline-Global-Collect/connect-sdk-go/webhooks/validation"
)

type inMemorySecretKeyStore struct{}

var keyStoreMap = map[string]string{}
var keyStoreLock = sync.RWMutex{}
var keyStore = &inMemorySecretKeyStore{}

// GetSecretKey returns the secretKey associated with the given keyID
func (ks *inMemorySecretKeyStore) GetSecretKey(keyID string) (string, error) {
	keyStoreLock.RLock()
	defer keyStoreLock.RUnlock()

	value, exists := keyStoreMap[keyID]

	if !exists {
		return "", validation.NewSecretKeyNotAvailableError(keyID)
	}

	return value, nil
}

// InMemorySecretKeyStore returns an in-memory secret key store.
// This can be used in applications where secret keys can be specified at application startup.
func InMemorySecretKeyStore() validation.SecretKeyStore {
	return keyStore
}

// StoreInMemorySecretKey stores the given keyID, secretKey pair in the in-memory secret key store
func StoreInMemorySecretKey(keyID, secretKey string) error {
	if strings.TrimSpace(keyID) == "" {
		return errors.New("invalid keyID")
	}
	if strings.TrimSpace(secretKey) == "" {
		return errors.New("invalid secret")
	}

	keyStoreLock.Lock()
	defer keyStoreLock.Unlock()

	keyStoreMap[keyID] = secretKey

	return nil
}

// RemoveInMemorySecretKey removes the given keyID and it's associated secretKey from the in-memory secret key store
func RemoveInMemorySecretKey(keyID string) {
	keyStoreLock.Lock()
	defer keyStoreLock.Unlock()

	delete(keyStoreMap, keyID)
}

// ClearInMemorySecretKeys empties the in-memory secret key store
func ClearInMemorySecretKeys() {
	keyStoreLock.Lock()
	defer keyStoreLock.Unlock()

	keyStoreMap = map[string]string{}
}
