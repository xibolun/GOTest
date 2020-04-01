package algorithm

import (
	"fmt"
	"strings"
	"testing"
)

func TestAnd(t *testing.T) {
	a := []byte("hello world")
	binaryStr := fmt.Sprintf("%0b", a)

	fmt.Println(binaryStr)
	fmt.Println(strings.Replace(binaryStr, " ", "", -1))

	//space := "hello world"
	str := strings.FieldsFunc(binaryStr, func(r rune) bool {
		return r == 91 || r == 93 || r==32
	})

	fmt.Println(str)
}
