package basic

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestOS(t *testing.T) {
	fmt.Printf("home path: %s\n", os.Getenv("$HOME"))
	fmt.Printf("os env: %s\n", os.Environ())
	fmt.Printf("go path: %s\n", os.Getenv("GOPATH"))
	fmt.Printf("$HOME path: %s\n", os.Getenv("HOME"))
	//fmt.Printf("os user path: %s\n",os.UserCacheDir())

}

func TestSystemIn(t *testing.T) {

}

func TestOSStdin_test(t *testing.T) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
}
