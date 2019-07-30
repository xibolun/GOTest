package basic

import (
	"bufio"
	"fmt"
	"os/exec"
)

func Output() {
	cmd := exec.Command("ll")

	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}
