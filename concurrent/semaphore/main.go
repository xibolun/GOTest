package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	c chan int
	w *sync.WaitGroup
}

func main() {
	//WaitGroupModel()
	SemaphoreModel()

}

// 无线程池模式
func WaitGroupModel() {
	start := time.Now()

	w := &sync.WaitGroup{}

	for i := 0; i < 1000000; i++ {
		w.Add(1)

		go func(i int) {
			fmt.Printf("%d start output \n", i)
			w.Done()
		}(i)
	}

	w.Wait()
	fmt.Printf("finished! const %s\n", time.Now().Sub(start).String())
}

func SemaphoreModel() {
	start := time.Now()
	sema := &Semaphore{
		c: make(chan int, 10),
		w: &sync.WaitGroup{},
	}
	for i := 0; i < 1000000; i++ {
		sema.w.Add(1)

		go func(i int) {
			sema.c <- i
			fmt.Printf("%d start insert into channel\n", i)
			fmt.Printf("%d start out  channel\n", <-sema.c)
			sema.w.Done()
		}(i)
	}

	sema.w.Wait()
	fmt.Printf("finished! const %s\n", time.Now().Sub(start).String())
}

func ContextModel() {

}
