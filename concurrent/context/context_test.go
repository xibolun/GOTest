package context

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_contextWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for i := 0; i < 10; i++ {
			OutPut(ctx, "hello,  "+strconv.Itoa(i))
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
}
func Test_contextWithDeadline(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))

	go func() {
		for i := 0; i < 10; i++ {
			OutPut(ctx, "hello,  "+strconv.Itoa(i))
		}
	}()

	time.Sleep(10 * time.Second)

	cancel()
}

func Test_contextWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	go func() {
		for i := 0; i < 10; i++ {
			OutPut(ctx, "hello,  "+strconv.Itoa(i))
		}
	}()

	time.Sleep(10 * time.Second)

	cancel()
}

func Test_contextWithTimeout2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("go routine is timeout")
				//time.Sleep(1 * time.Second)
				return
			default:
				fmt.Println("go routine is running")
				//time.Sleep(1 * time.Second)
			}
		}
		fmt.Println("out for loop!")

	}(ctx)

	cancel()
}

// cancel 会将context下的child同步结束掉
func Test_contextChild(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	subCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

	go func() {
		for i := 0; i < 10; i++ {
			OutPut(ctx, "ctx: hello,  "+strconv.Itoa(i))

			OutPut(subCtx, "subCtx: hello,  "+strconv.Itoa(i))
		}
	}()

	time.Sleep(10 * time.Second)

	cancel()
}

func Test_contextWithValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "golang_version", "1.14.0")

	go func() {
		for i := 0; i < 10; i++ {
			OutPut(ctx, "hello,  "+strconv.Itoa(i))
		}
	}()

	time.Sleep(2 * time.Second)
}

func OutPut(ctx context.Context, msg string) {
	if ctx.Value("golang_version") != nil {
		fmt.Printf("golang_version :%s \n", ctx.Value("golang_version"))
	}
	select {
	case <-ctx.Done():
		fmt.Println("done")
		return
	default:
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
	}
}
