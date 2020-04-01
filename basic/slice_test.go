package basic

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := [3]int{1, 2, 3}

	// array to slice
	slice := arr[:3]
	slice = append(slice, 123)

	fmt.Println(arr)
	fmt.Println(slice)

}
