package keyValStorage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-xyz/core/keyValStorage"
)

func TestNewKeyValStorage_GetKeyAndVal(t *testing.T) {
	t.Parallel()

	key := []byte("key")
	value := []byte("value")

	keyVal := keyValStorage.NewKeyValStorage(key, value)
	assert.NotNil(t, keyVal)
	assert.Equal(t, key, keyVal.Key())
	assert.Equal(t, value, keyVal.Value())
}
