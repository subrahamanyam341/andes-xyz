package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-xyz/hashing"
	"github.com/subrahamanyam341/andes-xyz/hashing/blake2b"
	"github.com/subrahamanyam341/andes-xyz/hashing/sha256"
)

func TestNewHasher(t *testing.T) {
	t.Parallel()

	type res struct {
		hasher hashing.Hasher
		err    error
	}
	testData := make(map[string]res)
	testData["sha256"] = res{
		hasher: sha256.NewSha256(),
		err:    nil,
	}
	testData["blake2b"] = res{
		hasher: blake2b.NewBlake2b(),
		err:    nil,
	}
	testData[""] = res{
		hasher: nil,
		err:    ErrNoHasherInConfig,
	}
	testData["invalid hasher name"] = res{
		hasher: nil,
		err:    ErrNoHasherInConfig,
	}

	for key, value := range testData {
		hasher, err := NewHasher(key)
		assert.Equal(t, value.err, err)
		assert.Equal(t, value.hasher, hasher)
	}
}
