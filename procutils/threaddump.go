package procutils

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/cwdot/go-bark/logging"
	"github.com/cwdot/go-stdlib/colors"
)

// ThreadDumpSignaler sets up a signal handler to dump goroutines to a file
func ThreadDumpSignaler(outputPath string) {
	// Set up signal handler
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGUSR1)

	go func() {
		for {
			<-sigCh

			f, err := os.Create(outputPath)
			if err != nil {
				logging.Error("Error creating thread dump file", "output_path", outputPath, "error", err)
				break
			}
			defer f.Close()

			if err := pprof.Lookup("goroutine").WriteTo(f, 1); err != nil {
				logging.Error("Error writing thread dump", "output_path", outputPath, "error", err)
				break
			}
			logging.Info("Dumped goroutines to %s", "output_path", outputPath)
		}
	}()

	command := fmt.Sprintf("kill -SIGUSR1 %d", os.Getpid())
	logging.Infof("Dump signal handler: %s", colors.Green.Wrap(command))
}
