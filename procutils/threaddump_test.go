package procutils

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestThreadDumpSignaler(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test.threaddump")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Clean up the file after the test

	// Call the ThreadDumpSignaler function
	ThreadDumpSignaler(tmpFile.Name())

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		panic(errors.Wrap(err, "error finding process"))
	}

	err = p.Signal(syscall.SIGUSR1)
	if err != nil {
		panic(errors.Wrap(err, "error sending signal"))
	}

	// Wait for the signal
	select {
	case <-time.After(100 * time.Millisecond):
	}

	fileInfo, err := os.Stat(tmpFile.Name())
	require.NoError(t, err)

	contents, err := os.ReadFile(tmpFile.Name())
	require.NoError(t, err)

	t.Logf("Thread dump file name=%s size=%d bytes\n", tmpFile.Name(), fileInfo.Size())
	t.Logf("Thread dump contents: %s\n", string(contents))

	require.NotZero(t, fileInfo.Size(), "Thread dump file should not be empty")
}
