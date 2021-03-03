package basic

import (
	"fmt"
)

//func main() {
//	// defineArray()
//	// eachArray()
//	// testMake()
//	// appendFuc()
//	rangeFunc()
//}

func defineArray() {
	var a [2]string ///// 定义数组
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a)
}

func eachArray() {
	p := []int{1, 3, 5, 7, 9, 11} //// 定义数组
	fmt.Println(p)

	for i := 0; i < len(p); i++ {
		fmt.Println("the index and value is ", i, p[i])
	}

	//// go语言的区间是左闭右开的区间
	fmt.Println("the index 0-4 value are ", p[:4])   /// the index 0-4 value are  [1 3 5 7]
	fmt.Println("the index 1-4 value are ", p[1:4])  /// the index 1-4 value are  [3 5 7]
	fmt.Println("the index 4-end value are ", p[4:]) /// the index 4-end value are  [9 11]
}

func testMake() {
	a := make([]int, 0, 10)
	b := make([]int, 5)
	fmt.Println(a)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
}

func appendFuc() {
	a := make([]int, 0, 2)
	fmt.Println(a)
	fmt.Println(append(a, 1))
	fmt.Println(append(a, 1, 3, 4)) //append会自动增加数组的容量，可以增加多个
}

func rangeFunc() {
	var a = []int{1, 2, 3, 4, 5, 6, 7}
	// for循环的range格式
	for i, v := range a {
		fmt.Println(i, v)
	}
}
