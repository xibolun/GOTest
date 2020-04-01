package benchmarks

import (
	"math"
	"testing"
)

func clear(n uint64, i, j uint8) uint64 {
	return (math.MaxUint64<<j | ((1 << i) - 1)) & n
}
func BenchmarkWrong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clear(1221892080809121, 10, 63)
	}
}

var result uint64

func BenchmarkCorrect(b *testing.B) {
	var r uint64
	for i := 0; i < b.N; i++ {
		r = clear(1221892080809121, 10, 63)
	}
	result = r
}
