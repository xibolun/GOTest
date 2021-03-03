package basic

import (
	"sync"
	"fmt"
)

func TestLock() {
	var mutex sync.Mutex

	mutex.Unlock()

	for i := 0; i < 10; i++ {
		mutex.Lock()
		fmt.Printf("mutex locked %d \n", i)

		mutex.Unlock()
		fmt.Printf("mutex unlocked %d \n", i)
	}

}
