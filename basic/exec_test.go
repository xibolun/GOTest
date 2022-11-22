package basic

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
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

	cmd := exec.CommandContext(ctx, "/bin/bash", "/tmp/a.sh")
	//cmd := exec.CommandContext(ctx, "lspci")

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	outReader := bufio.NewReaderSize(stdoutPipe, 2*1024*1024)
	errReader := bufio.NewReaderSize(stderrPipe, 2*1024*1024)

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

			if err != nil && err != io.EOF && err.Error() != "read |0: file already closed" {
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

			if err != nil && err != io.EOF && err.Error() != "read |0: file already closed" {
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
		case <-time.After(100 * time.Millisecond):
			if err = cmd.Wait(); err != nil {
				if err.Error() == "exec: Wait was already called" || err == os.ErrClosed {
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

	fmt.Println("stdout", stdoutStr)
	fmt.Println("stderr", stderrStr)
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

func TestOutputTimeout(t *testing.T) {
	argv := []string{"-c", "/tmp/a.sh"}
	attr := new(os.ProcAttr)
	newProcess, err := os.StartProcess("/bin/bash", argv, attr) //运行脚本
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Process PID", newProcess.Pid)
	processState, err := newProcess.Wait() //等待命令执行完
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("processState PID:", processState.Pid()) //获取PID
	fmt.Println("ProcessExit:", processState.Exited())   //获取进程是否退出
}

func TestCmdRunTimeout2(t *testing.T) {
	err, _ := CmdRunWithTimeout("ping www.baidu.com", 10*time.Second)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCmdOutput(t *testing.T) {
	ast := assert.New(t)

	output, err := exec.Command("cat", "/tmp/aa.sh").CombinedOutput()
	ast.NotNil(err)      // exit status 1
	ast.NotEmpty(output) // cat: /tmp/aa.sh: No such file or directory
}

func TestCmdOutputPipeline(t *testing.T) {
	ast := assert.New(t)

	output, err := exec.Command("cat", "/tmp/aa.sh", "|", "grep", "aa").CombinedOutput()
	ast.NotNil(err) // exit status 1
	//cat: /tmp/aa.sh: No such file or directory
	//cat: |: No such file or directory
	//cat: grep: No such file
	ast.NotEmpty(output)
}

func TestCmdOutputPipeline1(t *testing.T) {
	ast := assert.New(t)

	output, err := exec.Command(`cat /tmp/aa.sh | grep aa`).CombinedOutput()
	ast.NotNil(err)   // fork/exec cat /tmp/aa.sh | grep aa: no such file or directory
	ast.Empty(output) // nil
}

func TestCmdOutputPipeline2(t *testing.T) {
	ast := assert.New(t)
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(`cat /tmp/aa.sh | grep aa`)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	ast.NotNil(err) // fork/exec cat /tmp/aa.sh | grep aa: no such file or directory
	ast.Empty(stderr.Bytes())
	ast.Empty(stdout.Bytes())
}

func TestCmdOutputPipeline3(t *testing.T) {
	ast := assert.New(t)
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", `cat /tmp/aa.sh | grep aa`)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	ast.NotNil(err)              // exit status 1
	ast.NotEmpty(stderr.Bytes()) // cat: /tmp/aa.sh: No such file or directory
	ast.Empty(stdout.Bytes())
}
