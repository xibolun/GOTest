package basic

import (
	"fmt"
	"testing"
)

func Test_defer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	panic("error")
}
