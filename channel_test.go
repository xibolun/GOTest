package basic

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	time.Sleep(1 * time.Second)
	fmt.Printf("after sleep")
}

func greet(c chan string) {
	fmt.Printf("hello +%s\n", <-c)
}

func TestChannel1(t *testing.T) {
	c := make(chan string)

	// channel stop then main routine util write data
	greet(c)

	fmt.Printf("after greet")
}

func TestChannel2(t *testing.T) {
	c := make(chan string)

	// channel stop then main routine util write data
	go greet(c)

	c <- "world"
	fmt.Printf("after greet")
}

// TestDeadLock 在同一goroutine里面不能又读又取无缓冲通道的channel
// 程序hang住
func TestDeadLock(t *testing.T) {
	c := make(chan string)
	c <- "hello"
	fmt.Printf("after channel: %s\n", <-c)
}

func TestChannel3(t *testing.T) {

	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

func TestChannel4(t *testing.T) {
	var ch chan int
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}
