package main

import (
	"fmt"
	"testing"
)

const (
	v128 = 0x80
	v64  = 0x40
	v32  = 0x20
	v16  = 0x10
)

func Test_OrOper(t *testing.T) {

	a := 10
	// 10 = 1010
	// 16 =  10000   10 | 16 = 11010  = 26
	fmt.Printf(" or operation with 16: %d\n", a|v16)
	// 10 & 16 = 00000 = 0
	fmt.Printf(" and operation with 16: %d\n", a&v16)
	//  10 &^16 = 1010 & 01111 = 10
	fmt.Printf(" and not  operation with 16: %d\n", a&^v16)
	// 左移 加1位 10 = 1010  << 1 = 10100 = 2^4+2^2 = 20  相当于*2
	fmt.Printf(" shift operation with 16: %d\n", a<<1)
	// 右移 减1位 10 = 1010  >> 1 = 101 = 2^2+2^0 = 5  相当于/2
	fmt.Printf(" shift operation with 16: %d\n", a>>1)

	fmt.Println(64 << 10)
}

// 异或运算
func TestXOR(t *testing.T) {
	a := 100
	fmt.Println(a ^ 3)
	fmt.Println(a ^ a)
	fmt.Println(a ^ 3 ^ 3)
	fmt.Println(a ^ 3 ^ a)
}
