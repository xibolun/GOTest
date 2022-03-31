package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestReturnLoop(t *testing.T) {
	for {
		time.Sleep(2 * time.Second)
		return
	}

	fmt.Println("break success")

}
