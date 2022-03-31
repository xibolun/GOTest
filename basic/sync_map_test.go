package basic

import (
	"sync"
	"testing"
	"time"
)

func TestMap1(t *testing.T) {

}

func Test_SyncOnce(t *testing.T) {
	once := sync.Once{}
	for {
		once.Do(func() {
			t.Log("hello world")
		})
		time.Sleep(2 * time.Second)
	}
}
