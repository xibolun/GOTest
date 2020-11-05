package basic

import (
	"fmt"
	"math"
	"testing"
	"time"
	"unsafe"
)

func MathMethods() {
	fmt.Println(math.Round(0.5))       //四舍满五入
	fmt.Println(math.RoundToEven(0.5)) //四舍大于五入
	fmt.Println(math.RoundToEven(0.6)) //四舍大于五入
	fmt.Println(math.Round(0.6 - 0.5)) //只舍不入
	fmt.Println(math.Round(0.6 + 0.5)) //只舍不入
	fmt.Println(math.Ceil(0.6 + 0.5))  //只入不舍
	fmt.Println(math.Ceil(0.6 - 0.5))  //只入不舍

	fmt.Printf("1 mod 3:  %f\n", math.Mod(1, 3))
	fmt.Printf("1/3:   %f\n", 1.0/3)
	fmt.Printf("1/3:   %d\n", int(math.Ceil(1.0/3)))
	fmt.Println(7 % 13) //取余

}

func TestMath(t *testing.T) {
	MathMethods()
}

func TestPercent(t *testing.T) {
	var diskUsed uint32 = 3
	var diskTotal uint32 = 7
	fmt.Println(float64(diskUsed) * 100.0 / float64(diskTotal))
	fmt.Println(fmt.Sprintf("%.2f", float64(diskUsed)*100.0/float64(diskTotal)))
}

func TestMathMax(t *testing.T) {
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(time.Now().Unix())

	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	fmt.Println(unsafe.Sizeof(i1))
	fmt.Println(unsafe.Sizeof(i2))
	fmt.Println(unsafe.Sizeof(i3))
	fmt.Println(unsafe.Sizeof(i4))
	fmt.Println(unsafe.Sizeof(i5))

}
