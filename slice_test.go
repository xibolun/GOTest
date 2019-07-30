package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSliceCopy(b *testing.B) {
	b.ResetTimer()
	var a []int
	for i := 0; i < 10000000; i++ {
		a = append(a, i)
	}

	i := rand.Intn(1000000)

	copy(a[i:], a[i+1:])

	a = a[:len(a)-1]

	fmt.Println(len(a))

}

func BenchmarkSliceSplit(b *testing.B) {
	b.ResetTimer()
	var a []int
	for i := 0; i < 10000000; i++ {
		a = append(a, i)
	}
	i := rand.Intn(1000000)

	m := a[0:i]

	c := a[i+1:]

	a = append(m, c...)

	fmt.Println(len(a))

}
