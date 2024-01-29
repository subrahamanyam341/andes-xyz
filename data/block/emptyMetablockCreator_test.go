package block

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subrahamanyam341/andes-xyz/core/check"
)

func TestNewEmptyMetaBlockCreator(t *testing.T) {
	t.Parallel()

	creator := NewEmptyMetaBlockCreator()
	require.False(t, check.IfNil(creator))
}

func TestEmptyMetaBlockCreator_CreateNewHeader(t *testing.T) {
	t.Parallel()

	creator := NewEmptyMetaBlockCreator()
	header := creator.CreateNewHeader()
	require.False(t, check.IfNil(header))
	assert.Equal(t, "*block.MetaBlock", fmt.Sprintf("%T", header))
}
