package etcd

import (
	"fmt"
	"testing"
	"time"
)

func TestPub(t *testing.T) {
	for i := 0; i < 100; i++ {
		time.Sleep(3 * time.Second)

		_ = Pub(Queue, fmt.Sprintf("hello: %d", i))
	}
}

func TestSub(t *testing.T) {
	Sub(Queue)
}
