package goblog

import (
	"sync"
	"testing"
	"time"
)

// golang memory_model https://golang.org/ref/mem

func TestSynchronization1(t *testing.T) {
	Synchronization1()
}

func TestSynchronization2(t *testing.T) {
	Synchronization2()
}

func TestIncorrectSynchronization1(t *testing.T) {
	IncorrectSynchronization1()
}

func TestIncorrectSynchronization2(t *testing.T) {
	IncorrectSynchronization2()
}

func TestIncorrectSynchronization3(t *testing.T) {
	IncorrectSynchronization3()
}

func TestIncorrectSynchronization4(t *testing.T) {
	IncorrectSynchronization4()
}

// IncorrectSynchronization4
//
func IncorrectSynchronization4() {
	type T struct {
		msg string
	}

	var g *T

	funcSetup := func() {
		t := new(T)
		t.msg = "hello, world"
		g = t
	}

	funcMain := func() {
		go funcSetup()
		for g == nil {
		}
		print(g.msg)
	}

	funcMain()
}

// IncorrectSynchronization3
//
func IncorrectSynchronization3() {
	var a string
	var done bool

	funcSetup := func() {
		a = "hello, world"
		done = true
	}

	funcMain := func() {
		go funcSetup()
		for !done {
		}
		print(a)
	}

	funcMain()
}

// IncorrectSynchronization2
// 			有可能输出2次，也有可能一次都不输出，无法保证
func IncorrectSynchronization2() {
	var a string
	var done bool
	var once sync.Once

	funcSetup := func() {
		a = "hello, world"
		done = true
	}

	funcDoPrint := func() {
		if !done {
			once.Do(funcSetup)
		}
		print(a)
	}

	funcTwoPrint := func() {
		go funcDoPrint()
		go funcDoPrint()
	}

	funcTwoPrint()
}

// IncorrectSynchronization1
// 			有可能输出00  21  20
func IncorrectSynchronization1() {
	var a, b int

	funcF := func() {
		a = 1
		b = 2
	}

	funcG := func() {
		print(b)
		print(a)
	}

	funcMain := func() {
		go funcF()
		// 由于线程太快了，添加一个休眠时间，更好地观察输出
		time.Sleep(10 * time.Microsecond)
		funcG()
	}

	funcMain()
}

// Synchronization1 有可能会输出空string
func Synchronization1() {
	var a string

	funcF := func() {
		println(a)
	}

	funcHello := func() {
		a = "hello, world"
		go funcF()
	}

	funcHello()
}

// Synchronization2 有可能会输出空string
func Synchronization2() {
	var a string

	funcHello := func() {
		go func() {
			a = "hello world"
		}()
		print(a)
	}

	funcHello()
}
