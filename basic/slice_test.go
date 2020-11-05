package basic

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}

	//// array to slice
	//slice := arr[:3]
	//slice = append(slice, 123)
	//
	//fmt.Println(arr)
	//fmt.Println(slice)

	fmt.Println(arr[0:])
	fmt.Println(arr[0:3])

}
