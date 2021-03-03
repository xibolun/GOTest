package basic

import (
	"fmt"
	"time"
)

func TestRoutine(word string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(word)
}
