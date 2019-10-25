package basic

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func Test(t *testing.T) {

	a := int32(0)

	atomic.AddInt32(&a, 1)
	fmt.Println(a)

	var zero = int32(0)
	fmt.Println(atomic.LoadInt32(&zero))

	b := int32(0)
	fmt.Println(atomic.LoadInt32(&b))
}

func TestUnit(t *testing.T) {
	var i uint
	j := -1
	fmt.Println(i - 1)
	fmt.Println(uint(j))
}

func TestSubtraction(t *testing.T) {
	a1 := uint(2)
	b1 := uint(3)
	fmt.Println(a1 - b1)

	a2 := uint8(2)
	b2 := uint8(3)
	fmt.Println(a2 - b2)

}
