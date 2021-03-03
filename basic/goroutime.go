package basic

import (
	"fmt"
	"time"
)
//
//func main() {
//	//	go say("world")
//	// 	say("hello")
//
//	channelFunc()
//	cacheChannel()
//}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func channelFunc() {
	a := []int{1, 2, 3, 4}
	c := make(chan int)

	// 	fmt.Println(sum(a, c))

	go sum(a, c)

	fmt.Println("channel the results C = : ", <-c) // <-c 将c的值从channel当中导出
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将sum的值放入channel当中
}

func cacheChannel() {
	c := make(chan int, 2)

	c <- 1
	c <- 2

	fmt.Println("cacheChannel the value of C is :", <-c)
	fmt.Println("cacheChannel the value of C is :", <-c)

}
