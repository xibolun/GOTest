package concurrent

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
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
	Loop10Order()
	//LoopGroup()
}

func TestChannelTimeout(t *testing.T) {
	Convey("test channel timeout", t, func() {
		wg := sync.WaitGroup{}
		loopTimes := 10

		sum := 0
		for i := 0; i < loopTimes; i++ {
			lock := make(chan struct{})
			wg.Add(1)

			go func(i int) {
				random := rand.Intn(10)
				fmt.Printf("random is %d\n", random)
				time.Sleep(time.Duration(random) * time.Second)
				sum += i
				fmt.Printf("output is %d\n", i)
				lock <- struct{}{}
			}(i)

			select {
			case <-time.After(2 * time.Second):
				fmt.Printf("after 2s \n")
			case <-lock:
				fmt.Printf("output done\n")
			}
			wg.Done()
		}

		wg.Wait()

		fmt.Printf("all is done, sum is %d\n", sum)

	})
}
func TestGoroutineTimeout(t *testing.T) {
	Convey("test goroutine timeout", t, func() {
		wg := sync.WaitGroup{}
		loopTimes := 10

		sum := 0
		for i := 0; i < loopTimes; i++ {
			wg.Add(1)
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			go func(i int, ctx context.Context) {
				random := rand.Intn(10)
				fmt.Printf("random is %d\n", random)
				time.Sleep(time.Duration(random) * time.Second)
				sum += i
				fmt.Printf("output is %d\n", i)
			}(i, ctx)

			select {
			case <-ctx.Done():
				fmt.Printf("output done\n")
			case <-time.After(2 * time.Second):
				fmt.Printf("after 2s \n")
			}
		}

		wg.Wait()

		fmt.Printf("all is done, sum is %d\n", sum)

	})
}

func Loop10Order() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(c chan int) {
			fmt.Println(<-c)
		}(c)
		c <- i
	}
}

func LoopGroup() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Loop10() {
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	go func() {
		time.Sleep(2 * time.Second)
		close(c)
	}()
	<-c

	fmt.Println("all done")
}
