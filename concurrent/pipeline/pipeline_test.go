package pipeline

// golang pipeline model from https://blog.golang.org/pipelines
// make some tests

import (
	"fmt"
	"sync"
	"testing"
)

func Test_gen(t *testing.T) {
	out := gen(1, 2, 3, 4)

	for i := range out {
		fmt.Printf("output: %d\n", i)
	}
}

func Test_sq(t *testing.T) {
	for i := range sq(gen(1, 2, 3, 4)) {
		fmt.Printf("output: %d\n", i)
	}
}

// A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels
// onto a single channel that's closed when all the inputs are closed. This is called fan-in.
func gen(nums ...int) <-chan int {
	// 这里需要做缓冲么？why?做缓冲可以提高效率
	//out := make(chan int, len(nums))
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("input: %d\n", n)
			out <- n
		}
		// 若注释掉这行代码，则整个程序会发生deadlock
		close(out)
	}()
	// 这里需要休眠么？ why？
	//time.Sleep(2 * time.Second)
	return out
}

// Multiple functions can read from the same channel until that channel is closed; this is called fan-out.
// This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// merge multiple channel value to one
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
