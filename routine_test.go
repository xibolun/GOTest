package basic

import (
	"testing"
	"fmt"
	"time"
)

func TestTestRoutine(t *testing.T) {

	go TestRoutine("hello", 2)
	go TestRoutine("world", 1)

	fmt.Println("start to say word")
	time.Sleep(time.Duration(2) * time.Second)

	fmt.Println("say word done")
}
