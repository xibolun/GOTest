package basic

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkSliceCopy(b *testing.B) {
	b.ResetTimer()
	var a []int
	for i := 0; i < 10000000; i++ {
		a = append(a, i)
	}

	i := rand.Intn(1000000)

	copy(a[i:], a[i+1:])

	a = a[:len(a)-1]

	fmt.Println(len(a))

}

func BenchmarkSliceSplit(b *testing.B) {
	b.ResetTimer()
	var a []int
	for i := 0; i < 10000000; i++ {
		a = append(a, i)
	}
	i := rand.Intn(1000000)

	m := a[0:i]

	c := a[i+1:]

	a = append(m, c...)

	fmt.Println(len(a))

}

func TestSL(t *testing.T) {
	sl := make([]int, 3, 5)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))

	var nilS []int
	fmt.Printf("nil slice :%t\n", nilS == nil)

	origin := []int{10, 20, 30, 40, 50}
	current := origin[1:3]
	// 容量: cap(origin)-1, 长度: 3-1
	fmt.Printf("current len %d, cap %d, value %v\n", len(current), cap(current), current)

	// current 与 origin 共享了一套数组，所以值会被连动修改
	current[1] = 35
	fmt.Printf("current len %d, cap %d, value %v\n", len(current), cap(current), current)
	fmt.Printf("origin len %d, cap %d, value %v\n", len(origin), cap(origin), origin)
}

// TestSL2 容量以2倍增长
func TestSL2(t *testing.T) {
	origin := []int{10, 20, 30, 40, 50}
	current := append(origin, 60)

	fmt.Printf("origin len %d, cap %d, value %v\n", len(origin), cap(origin), origin)
	fmt.Printf("current len %d, cap %d, value %v\n", len(current), cap(current), current)

	//origin len 5, cap 5, value [10 20 30 40 50]
	//current len 6, cap 10, value [10 20 30 40 50 60]
}

func TestSL21(t *testing.T) {
	var origin []int
	fmt.Printf("origin len %d, cap %d, value %v\n", len(origin), cap(origin), origin)

	origin = append(origin, 2)
	fmt.Printf("origin len %d, cap %d, value %v\n", len(origin), cap(origin), origin)

	origin = append(origin, 7)
	fmt.Printf("origin len %d, cap %d, value %v\n", len(origin), cap(origin), origin)

	origin = append(origin, 1)
	fmt.Printf("origin len %d, cap %d, value %v\n", len(origin), cap(origin), origin)

	origin = append(origin, 3)
	fmt.Printf("origin len %d, cap %d, value %v\n", len(origin), cap(origin), origin)

	origin = append(origin, 8)
	fmt.Printf("origin len %d, cap %d, value %v\n", len(origin), cap(origin), origin)

	origin = append(origin, 4)
	fmt.Printf("origin len %d, cap %d, value %v\n", len(origin), cap(origin), origin)

	//origin len 0, cap 0, value []
	//origin len 1, cap 1, value [2]
	//origin len 2, cap 2, value [2 7]
	//origin len 3, cap 4, value [2 7 1]
	//origin len 4, cap 4, value [2 7 1 3]
	//origin len 5, cap 8, value [2 7 1 3 8]
	//origin len 6, cap 8, value [2 7 1 3 8 4]
}

func TestSL3(t *testing.T) {
	var origin []int

	for i := 0; i < 1003; i++ {
		origin = append(origin, i)
	}

	fmt.Printf("origin len %d, cap %d\n", len(origin), cap(origin))
	current := append(origin, 10)
	fmt.Printf("current len %d, cap %d\n", len(current), cap(current))

	//=== RUN   TestSL3
	//origin len 1003, cap 1024
	//current len 1004, cap 1024
}

func TestSL4(t *testing.T) {
	origin := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// 引入第三个索引进行创建切片
	// 长度为3-2，容量为4-2
	current := origin[2:3:4]

	fmt.Printf("current len %d, cap %d, value %v\n", len(current), cap(current), current)

	//current len 1, cap 2, value [Plum]
}

func TestSL5(t *testing.T) {
	origin := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	for i, v := range origin {
		fmt.Printf("index: %d, value: %s, value_point %v\n", i, v, &v)
	}
	//	value的指针地址都是同一个
	//index: 0, value: Apple, value_point 0xc00004d110
	//index: 1, value: Orange, value_point 0xc00004d110
	//index: 2, value: Plum, value_point 0xc00004d110
	//index: 3, value: Banana, value_point 0xc00004d110
	//index: 4, value: Grape, value_point 0xc00004d110
}

func TestSL6(t *testing.T) {
	a := []byte("ba")

	a1 := append(a, 'd')
	a2 := append(a, 'g')

	fmt.Println(string(a1)) // bad
	fmt.Println(string(a2)) // bag
}

func TestAppend1(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	var dst2 []*int
	for _, i := range src {
		dst2 = append(dst2, &i)
	}

	for _, p := range dst2 {
		fmt.Print(*p)
	}
}

func TestAppend2(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}

	// 输出55555
	for _, p := range copySlicePoint(src) {
		fmt.Print(*p)
	}

	fmt.Println()

	// 输出55555
	for _, p := range copySlicePoint2(src) {
		fmt.Print(*p)
	}

	fmt.Println()

	// 输出12345
	for _, p := range copySlice(src) {
		fmt.Print(p)
	}
}

func TestAppend3(t *testing.T) {
	sl := make([]int, 0, 2)
	fmt.Printf("current len %d, cap %d, value %v\n", len(sl), cap(sl), sl)
	sl = append(sl, []int{1, 2}...)
	fmt.Printf("current len %d, cap %d, value %v\n", len(sl), cap(sl), sl)

	// len值与cap值一样
	ss := make([]int, 2)
	fmt.Printf("current len %d, cap %d, value %v\n", len(ss), cap(ss), ss)
	// append的时候会在len后面进行累加
	ss = append(ss, []int{1, 2}...)
	fmt.Printf("current len %d, cap %d, value %v\n", len(ss), cap(ss), ss)

	//current len 0, cap 2, value []
	//current len 2, cap 2, value [1 2]
	//current len 2, cap 2, value [0 0]
	//current len 4, cap 4, value [0 0 1 2]

}

func TestRemoveEl(t *testing.T) {
	t.Log(splitKeyword(`aa bb  cc '' " " dd`))
}

func copySlicePoint(src []int) []*int {
	var dst2 []*int
	for _, i := range src {
		dst2 = append(dst2, &i)
	}
	return dst2
}
func copySlicePoint2(src []int) []*int {
	var dst2 []*int
	var j *int
	for _, i := range src {
		j = &i
		dst2 = append(dst2, j)
	}
	return dst2
}

func copySlice(src []int) []int {
	var dst2 []int
	for _, i := range src {
		dst2 = append(dst2, i)
	}
	return dst2
}

func splitKeyword(keyword string) []string {
	str := make([]string, 0)
	if strings.Contains(keyword, ",") {
		str = strings.Split(keyword, ",")
	}
	if strings.Contains(keyword, " ") {
		str = strings.Split(keyword, " ")
	}

	if strings.Contains(keyword, "\n") {
		str = strings.Split(keyword, "\n")
	}
	if len(str) <= 0 {
		return []string{keyword}
	}

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == "" || str[i] == `""` || str[i] == "''" || str[i] == `"` || str[i] == "'" {
			str = append(str[:i], str[i+1:]...)
		}
	}
	fmt.Println(len(str))

	return str
}
