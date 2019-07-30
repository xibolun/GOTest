package basic

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

type Stringer interface {
	String() string
}

func TestInterface(t *testing.T) {
	fmt.Println(UpperString("hello, world"))
	fmt.Fprintln(os.Stdout, UpperString("hello, world").String())
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}
