package basic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// golang的rand为什么要设计seed？

// 每次执行结果：
//			81
//			87
//			47
//			59
//			81
func TestRand1(t *testing.T) {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}

//=== RUN   TestRand2
// 每次执行也是同样的结果
//			5
//			87
//			68
//			50
//			23
func TestRand2(t *testing.T) {
	rand.Seed(42)
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
// seed设置为动态的才会保证随机生成
func TestRand3(t *testing.T) {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
