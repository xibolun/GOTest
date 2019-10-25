package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	var ts = make([]T, 10)
	for i := 0; i < 10; i++ {
		ts[i] = T{i}
	}
	for _, t := range ts {
		go t.Incr(&wg)
	}
	wg.Wait()
	for _, t := range ts {
		fmt.Println(t)
		go t.Print()
	}
	time.Sleep(5 * time.Second)
}

func TestGoRoutine2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	var ts = make([]*T, 10)
	for i := 0; i < 10; i++ {
		ts[i] = &T{i}
	}
	for _, t := range ts {
		go t.Incr(&wg)
	}
	wg.Wait()
	for _, t := range ts {
		fmt.Println(t)
		go t.Print()
	}
	time.Sleep(5 * time.Second)
}
