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
