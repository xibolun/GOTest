package basic

import (
	"fmt"
	"math"
	"testing"
)

func MathMethods() {
	fmt.Println(math.Round(0.5))       //四舍满五入
	fmt.Println(math.RoundToEven(0.5)) //四舍大于五入
	fmt.Println(math.RoundToEven(0.6)) //四舍大于五入
	fmt.Println(math.Round(0.6 - 0.5)) //只舍不入
	fmt.Println(math.Round(0.6 + 0.5)) //只入不舍

	fmt.Println(7 % 13) //取余

}

func TestMath(t *testing.T) {
	MathMethods()
}
