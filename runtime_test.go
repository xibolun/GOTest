package basic

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {
	fmt.Printf("runtime.NumCPU(): %d\n", runtime.NumCPU())
	fmt.Printf("runtime.NumCgoCall(): %d\n", runtime.NumCgoCall())
	fmt.Printf("runtime.Version(): %s\n", runtime.Version())
	fmt.Printf("runtime.GOROOT(): %s\n", runtime.GOROOT())
	fmt.Printf("runtime.GOOS: %s\n", runtime.GOOS)
	fmt.Printf("runtime.GOARCH: %s\n", runtime.GOARCH)

}
