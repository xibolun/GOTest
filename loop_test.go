package basic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLoop(t *testing.T) {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(100)
	ok := true
DONE:
	for {

		if r%2 == 0 {
			fmt.Printf("%d mod 2 ==0\n", r)
		}

		if r%2 == 1 {
			fmt.Printf("%d mod 2 ==1\n", r)

			ok = false
			break DONE
		}
		break DONE
	}

	fmt.Println(ok)
}
