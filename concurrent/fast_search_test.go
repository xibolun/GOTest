package concurrent

import (
	"fmt"
	"testing"
	"time"
)

/**
	多个协程进行处理，取最快返回的那一个
 */
func TestFastSearch(t *testing.T) {

	ch := make(chan int, 64)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 2
	}()

	go func() {
		time.Sleep(1 * time.Second)

		ch <- 3
	}()

	go func() {
		time.Sleep(1 * time.Second)

		ch <- 4
	}()
	go func() {
		time.Sleep(1 * time.Second)

		ch <- 5
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 6
	}()

	fmt.Println(<-ch)
}
