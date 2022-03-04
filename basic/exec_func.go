package basic

import (
	"os/exec"
	"strings"
	"time"
)

func CmdRunWithTimeout(cmdStr string, timeout time.Duration) (error, bool) {

	cmd := exec.Command("su", strings.Split(cmdStr, " ")...)
	var err error
	err = cmd.Start()
	if err != nil {
		return err, false
	}
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(timeout):
		// timeout
		if err = cmd.Process.Kill(); err != nil {
			return err, false
		}
		go func() {
			<-done // allow goroutine to exit
		}()
		return err, true
	case err = <-done:
		return err, false
	}
}
