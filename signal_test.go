package basic

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func TestSignal(t *testing.T) {

	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1)
	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		for range sigs {
			sig := <-sigs
			fmt.Printf("go signal value %v", sig)
		}
	}()

	for {
		fmt.Print()
	}

}
