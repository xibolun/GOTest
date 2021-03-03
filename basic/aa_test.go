package basic

import (
	"fmt"
	"sync"
	"testing"
)

func TestSwap(t *testing.T) {

	//ch := make(chan int)
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println(<-ch)
	//	}(i)
	//	ch <- i
	//}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
