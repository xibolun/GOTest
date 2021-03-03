package basic

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	c := make(chan int)
	stop := make(chan struct{})
	go AddData(c, stop)
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-stop:
			fmt.Println("TestSelect is end")
		}
	}
}

func TestTimerSelect(t *testing.T) {
	c := make(chan int)
	stop := make(chan struct{})
	go AddData(c, stop)

	waitDuration := 10 * time.Millisecond
	for {
		select {
		case <-time.After(waitDuration):
			fmt.Println("time duration 10 millisecond")
		case v := <-c:
			fmt.Println(v)
		case <-stop:
			fmt.Println("TestSelect is end")
		}

	}
}

func TestBreakSelect(t *testing.T) {
	c := make(chan int)
	stop := make(chan struct{})
	go AddData(c, stop)

	waitDuration := 10 * time.Millisecond
	//for {
	//	select {
	//	case <-time.After(waitDuration):
	//		fmt.Println("time duration 10 millisecond")
	//	case v := <-c:
	//		fmt.Println(v)
	//	case <-stop:
	//		fmt.Println("TestSelect is end")
	//		return
	//	}
	//
	//}
LOOP:
	for {
		select {
		case <-time.After(waitDuration):
			fmt.Println("time duration 10 millisecond")
		case v := <-c:
			fmt.Println(v)
		case <-stop:
			fmt.Println("TestSelect is end")
			break LOOP
		}

	}
	fmt.Println("program is done")
}

func TestRegisterFunc(t *testing.T) {
	f := func(i string) {
		fmt.Println(i)
	}
	c := make(chan int)
	stop := make(chan struct{})
	go AddData(c, stop)

	waitDuration := 10 * time.Millisecond
	for {
		select {
		case <-time.After(waitDuration):
			fmt.Println("time duration 10 millisecond")
		case v := <-c:
			f(strconv.Itoa(v))
		case <-stop:
			fmt.Println("TestSelect is end")
			return
		}
	}
}

func TestLoopSelect(t *testing.T) {
	f := func(i string) {
		fmt.Println(i)
	}
	c := make(chan int)
	stop := make(chan struct{})
	begin := make(chan struct{})
	chs := []chan struct{}{stop, begin}
	go AddData(c, stop)

	waitDuration := 10 * time.Millisecond
	for {
		for i := range chs {
			select {
			case <-chs[i]:
				fmt.Println("TestSelect is End")
				return
			case <-time.After(waitDuration):
				fmt.Println("time duration 10 millisecond")
			case v := <-c:
				f(strconv.Itoa(v))
			default:
				fmt.Println("there is noting")
			}
		}

	}
}

func AddData(ch chan int, stop chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	stop <- struct{}{}
}
