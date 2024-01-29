package block

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subrahamanyam341/andes-xyz/core/check"
)

func TestNewEmptyHeaderV2Creator(t *testing.T) {
	t.Parallel()

	creator := NewEmptyHeaderV2Creator()
	require.False(t, check.IfNil(creator))
}

func TestEmptyHeaderV2Creator_CreateNewHeader(t *testing.T) {
	t.Parallel()

	creator := NewEmptyHeaderV2Creator()
	header := creator.CreateNewHeader()
	require.False(t, check.IfNil(header))
	require.False(t, check.IfNil(header.(*HeaderV2)))
	require.False(t, check.IfNil(header.(*HeaderV2).Header))
	assert.Equal(t, "*block.HeaderV2", fmt.Sprintf("%T", header))
}
