package basic

import (
	"fmt"
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

func TestDeadLock(t *testing.T) {
	c := make(chan string)
	c <- "hello"
	fmt.Printf("after channel\n")
}
