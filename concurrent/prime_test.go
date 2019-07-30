package concurrent

import (
	"fmt"
	"testing"
)

func GenerateCh() chan int {

	ch := make(chan int)

	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
	}()

	return ch
}

func FilterPrime(seq chan int, prime int) chan int {
	ch := make(chan int)

	go func() {
		for {

			if num := <-seq; num%prime != 0 {
				ch <- num
			}
		}
	}()
	return ch
}

func Test_prime(t *testing.T) {
	ch := GenerateCh()
	for i := 0; i < 100; i++ {
		prime := <-ch
		if prime == 0 || prime == 1 {
			continue
		}
		fmt.Println(prime)
		ch = FilterPrime(ch, prime)
	}

}
