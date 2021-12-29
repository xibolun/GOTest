package basic

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestSingleCommand(t *testing.T) {
	stdout, err := exec.Command("uname", "-a").CombinedOutput()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(stdout)
}

func TestSingleTimeoutCommand(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stdout, err := exec.CommandContext(ctx, "ping", "-c 2", "-i 1", "www.baidu.com").CombinedOutput()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(stdout))
}

// can not get stdout
//
// === RUN   TestSingleTimeoutCommand
//    exec_test.go:27: signal: killed
func TestLongTimeoutCommand(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stdout, err := exec.CommandContext(ctx, "ping", "www.baidu.com").CombinedOutput()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(stdout)
}

// can not cancel when timeout
// cat /tmp/a.sh
// 		ping www.baidu.com
//
// === RUN   TestSingleTimeoutCommand
//    exec_test.go:27: signal: killed
func TestTimeoutCancelFailureCommand(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stdout, err := exec.CommandContext(ctx, "/bin/bash", "/tmp/a.sh").CombinedOutput()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(stdout)
}

// cancel command && get stdout\ stderr
// cat /tmp/a.sh
// 		ping www.baidu.com
//
// === RUN   TestSingleTimeoutCommand
//    exec_test.go:27: signal: killed
func TestTimeoutCancelCommand(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//cmd := exec.CommandContext(ctx, "/bin/bash", "/tmp/a.sh")
	cmd := exec.CommandContext(ctx, "uname", "-a")

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	outReader := bufio.NewReader(stdoutPipe)
	errReader := bufio.NewReader(stderrPipe)

	stdoutChan := make(chan string, 0)
	stderrChan := make(chan string, 0)

	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go func() {
		for {
			line, err := outReader.ReadString('\n')

			if line != "" {
				stdoutChan <- line
			}

			if err != nil {
				stderrChan <- err.Error()
				return
			}

			if line == "" {
				break
			}
		}
	}()
	go func() {
		for {
			line, err := errReader.ReadString('\n')

			if line != "" {
				stderrChan <- line
			}

			if err != nil {
				stderrChan <- err.Error()
				return
			}

			if line == "" {
				break
			}
		}
	}()

	var stdoutStr string
	var stderrStr string
LoopBreak:
	for {
		select {
		case <-ctx.Done():
			break LoopBreak
		case str := <-stdoutChan:
			stdoutStr += str
		case str := <-stderrChan:
			stderrStr += str
		case <-time.After(10 * time.Millisecond):
			if err = cmd.Wait(); err != nil {
				if err.Error() == "exec: Wait was already called" {
					break LoopBreak
				}
				if exiterr, ok := err.(*exec.ExitError); ok {
					status := exiterr.Sys().(syscall.WaitStatus)
					if status.ExitStatus() == 0 {
						stderrStr = fmt.Sprintf("wrong exit status: %d", status.ExitStatus())
						break LoopBreak
					}
				}
			}
		}
	}

	fmt.Println(stdoutStr)
	fmt.Println(stderrStr)
	fmt.Println("exec done")
}

// get stdout
//
// === RUN   TestSingleTimeoutCommand
//    exec_test.go:27: signal: killed
func TestLongTimeoutStdoutCommand(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/bash", "/tmp/a.sh")

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	var stdoutBuf, stderrBuf bytes.Buffer

	// 将stdout写入os.stdout当中
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if errStdout != nil || errStderr != nil {
		fmt.Printf("failed to capture stdout or stderr\n")
		return
	}

	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}
