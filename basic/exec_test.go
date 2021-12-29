package basic

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
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
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/bash", "/tmp/a.sh")

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
		}
	}

	fmt.Println(stdoutStr)
	fmt.Println(stderrStr)
	fmt.Println("exec done")
}
