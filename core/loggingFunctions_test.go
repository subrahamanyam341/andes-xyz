package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subrahamanyam341/andes-xyz/core/mock"
)

func TestDumpGoRoutinesToLogShouldNotPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()
		if r != nil {
			require.Fail(t, fmt.Sprintf("should have not paniced %v", r))
		}
	}()

	DumpGoRoutinesToLog(0, &mock.LoggerMock{})
}

func TestGetRunningGoRoutines(t *testing.T) {
	t.Parallel()

	res := GetRunningGoRoutines(&mock.LoggerMock{})
	require.NotNil(t, res)
}
