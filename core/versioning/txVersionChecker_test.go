package versioning

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subrahamanyam341/andes-xyz/core"
	"github.com/subrahamanyam341/andes-xyz/data/transaction"
)

func TestTxVersionChecker_IsSignedWithHashOptionsZeroShouldReturnFalse(t *testing.T) {
	t.Parallel()

	minTxVersion := uint32(1)
	tx := &transaction.Transaction{
		Options: 0,
		Version: minTxVersion,
	}
	tvc := NewTxVersionChecker(minTxVersion)

	res := tvc.IsSignedWithHash(tx)
	require.False(t, res)
}

func TestTxVersionChecker_IsSignedWithHash(t *testing.T) {
	t.Parallel()

	minTxVersion := uint32(1)
	tx := &transaction.Transaction{
		Options: transaction.MaskSignedWithHash,
		Version: minTxVersion + 1,
	}
	tvc := NewTxVersionChecker(minTxVersion)

	res := tvc.IsSignedWithHash(tx)
	require.True(t, res)
}

func TestTxVersionChecker_IsGuardedTransaction(t *testing.T) {
	t.Parallel()

	minTxVersion := uint32(1)
	tx := &transaction.Transaction{
		Options: transaction.MaskGuardedTransaction,
		Version: minTxVersion + 1,
	}

	tvc := NewTxVersionChecker(minTxVersion)
	res := tvc.IsGuardedTransaction(tx)
	require.True(t, res)

	tx.Options = 0
	res = tvc.IsGuardedTransaction(tx)
	require.False(t, res)
}

func TestTxVersionChecker_CheckTxVersionShouldReturnErrorOptionsNotZero(t *testing.T) {
	minTxVersion := uint32(1)
	tx := &transaction.Transaction{
		Options: transaction.MaskSignedWithHash,
		Version: minTxVersion,
	}

	tvc := NewTxVersionChecker(minTxVersion)
	err := tvc.CheckTxVersion(tx)
	require.Equal(t, core.ErrInvalidTransactionVersion, err)
}

func TestTxVersionChecker_CheckTxVersionShould(t *testing.T) {
	minTxVersion := uint32(1)
	tx := &transaction.Transaction{
		Options: 0,
		Version: minTxVersion,
	}

	tvc := NewTxVersionChecker(minTxVersion)
	err := tvc.CheckTxVersion(tx)
	require.Nil(t, err)
}
