package basic

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// A\B交替执行，由于时间短，执行不了多久就结束了
//		=== RUN   TestWaitGroup1
//		test start, procs: 0, wgAdd: 0, second: 0
//		Waiting To Finish
//		B:2
//		B:3
//		B:5
//		B:7
//		done
//		--- PASS: TestWaitGroup1 (0.00s)
//		A:2
//		PASS
//		A:3
//		A:5
func TestWaitGroup1(t *testing.T) {
	WaitGroup(0, 0, 0)
}

// A\B交替执行，设置了sleep时间，能够执行一段时间，直到异常
//		A:4243
//		A:4253
//		B:4999
//		Completed B
//		A:4259
//		panic: sync: negative WaitGroup counter
func TestWaitGroup(t *testing.T) {
	WaitGroup(0, 0, 10)
}

// nothing
//		=== RUN   TestWaitGroup2
//		test start, procs: 1, wgAdd: 0, second: 0
//		Waiting To Finish
//		done
//		--- PASS: TestWaitGroup2 (0.00s)
//		PASS
func TestWaitGroup2(t *testing.T) {
	WaitGroup(1, 0, 0)
}

// 只会执行A，异常
//		A:4993
//		A:4999
//		Completed A
//		panic: sync: negative WaitGroup counter
func TestWaitGroup3(t *testing.T) {
	WaitGroup(1, 0, 5)
}

// 只会执行B
//		B:4999
//		Completed B
//		done
//		--- PASS: TestWaitGroup4 (0.06s)
//		PASS
func TestWaitGroup4(t *testing.T) {
	WaitGroup(1, 1, 0)
}

// A\B依次执行，谁先执行，谁先结束
//		=== RUN   TestWaitGroup5
//		test start, procs: 1, wgAdd: 2, second: 0
//		Waiting To Finish
//		B:2
//		B:3
//
//		B:4999
//		Completed B
//		A:2
//		....
//		A:4999
//		Completed A
//		done
//		--- PASS: TestWaitGroup5 (0.11s)
//		PASS
func TestWaitGroup5(t *testing.T) {
	WaitGroup(1, 2, 0)
}

// B先执行，A\B交换执行完
//		B:4999
//		Completed B
//		A:4931
//		A:4933
// 		....
//		A:4999
//		Completed A
//		done
//		--- PASS: TestWaitGroup6 (0.07s)
//		PASS
func TestWaitGroup6(t *testing.T) {
	WaitGroup(2, 2, 0)
}

// A\B交互执行，但B可以执行完，A不能
//		A:4733
//		B:4999
//		Completed B
//		done
//		A:4751
//		--- PASS: TestWaitGroup7 (0.07s)
//		A:4759
//		PASS
func TestWaitGroup7(t *testing.T) {
	WaitGroup(2, 1, 0)
}

// 先执行A, A\B交互执行，A\B都可以执行完，会发生异常
//		B:4799
//		A:4999
//		Completed A
//		B:4801
//		...
//		B:4999
//		Completed B
//		panic: sync: negative WaitGroup counter
func TestWaitGroup8(t *testing.T) {
	WaitGroup(2, 1, 10)
}

// A\B交替执行，由于CPU核心数增加，所以处理时间比TestWaitGroup6短
func TestWaitGroup9(t *testing.T) {
	WaitGroup(4, 2, 0)
}

// A\B依次执行，谁先抢到谁执行，谁就先执行完，执行完之后waiting
func TestWaitGroup10(t *testing.T) {
	WaitGroup(1, 3, 0)
}

// A\B交替执行，执行完之后waiting
func TestWaitGroup11(t *testing.T) {
	WaitGroup(2, 3, 0)
}

func TestWaitGroup12(t *testing.T) {
	var wg sync.WaitGroup

	f := func(wg *sync.WaitGroup) {
		wg.Done()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f(&wg)
	}

	wg.Wait()

	fmt.Println("program is done")
}

func WaitGroup(procs, wgAdd, second int) {
	fmt.Printf("test start, procs: %d, wgAdd: %d, second: %d\n", procs, wgAdd, second)
	runtime.GOMAXPROCS(procs)
	var wg sync.WaitGroup
	wg.Add(wgAdd)

	go printPrime("A", &wg)
	go printPrime("B", &wg)

	if second > 0 {
		time.Sleep(time.Duration(second) * time.Second)
	}

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("done")
}

// printPrime print prime
func printPrime(prefix string, wg *sync.WaitGroup) {
	defer (*wg).Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

func TestInnerFuncWait(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("program is done")
}
