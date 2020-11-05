package basic

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	var a, b = 1, 2

	swap(&a, &b)

	fmt.Printf("a: %d, b:%d\n", a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
