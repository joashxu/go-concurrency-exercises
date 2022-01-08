//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"os"
	"os/signal"
)

func main() {
	// Create a process
	proc := MockProcess{}

	go func() {
		// Create a channel to receive os.Signal values.
		// This channel should be buffered.
		sig := make(chan os.Signal, 1)

		// Register the given channel to receive notifications.
		// We use os.Interrupt instead of syscall.SIGINT so 
		// golang can handle the correct interrupt for the platform
		// it is running on (OS Agnostic). 
		signal.Notify(sig, os.Interrupt)

		// Block until we receive a signal
		<-sig

		// Remove all signal handler.
		// sig will no longer receive any signal
		signal.Reset()

		// Gracefully ask the process to kill itself.
		proc.Stop()
	}()

	// Run the process (blocking)
	proc.Run()
}
