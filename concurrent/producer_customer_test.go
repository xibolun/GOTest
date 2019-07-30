package concurrent

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func Producer(factor int, ch chan<- int) {
	i := 0
	for {
		ch <- i * factor
		i++
	}
}

func Customer(ch <-chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func TestProducerAndCustomer(t *testing.T) {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Customer(ch)

	////靠体眠是无法保证稳定的输出结果
	//time.Sleep(2 * time.Second)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
