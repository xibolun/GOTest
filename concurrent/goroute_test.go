package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 10个关于并发的测试 https://colobu.com/2019/04/28/go-concurrency-quizzes/
func Test1(t *testing.T) {
	var mu sync.RWMutex
	var count int
	go func() {
		mu.RLock()
		defer mu.RUnlock()
		func() {
			time.Sleep(5 * time.Second)
			func() {
				mu.RLock()
				defer mu.RUnlock()
			}()
		}()
	}()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

func TestLoop10(t *testing.T) {
	//Loop10()
	//Loop10Order()
	LoopGroup()
}
func Loop10Order() {
	for i := 0; i < 10; i++ {
		c := make(chan int)

		go func(c chan int) {
			fmt.Println(<-c)
		}(c)
		c <- i
	}
}

func LoopGroup() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go func(i int){
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Loop10() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	<-c
}
