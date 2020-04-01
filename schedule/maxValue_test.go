package schedule

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestP(t *testing.T) {
	fmt.Printf("GOMAXPROCS output: %s \n", os.Getenv("GOMAXPROCS"))
	// set max p
	fmt.Printf("runtime.GOMAXPROCS output: %d \n", runtime.GOMAXPROCS(10))
	// get cpu num
	//fmt.Println(runtime.NumCPU())
	fmt.Printf("runtime.NumCPU() output: %d \n", runtime.NumCPU)
}

func TestNot(t *testing.T) {
	// 32位操作系统
	// 0001 取反  1110 二进制负数转10进制，最高位为符号位，其他位取反+1
	// 110取反 001
	// 001 +1 = 010 = 2^1 = 2

	// 十进制负数(-1)转二进制；十进制转二进制，取反后再+1
	// 1转二进制 0001 取反 1110 +1 = 1111
	fmt.Printf("^1 output: %d\n", ^1)

	// 64个0取非
	fmt.Printf("^uint64(0) output: %d\n", ^uint64(0))

	// uintptr 32位= uint32  64位 = uint64
	fmt.Printf("^uintptr(0) output: %d\n", ^uintptr(0))

	fmt.Printf("^uintptr(0) >> 63 output: %d\n", ^uintptr(0)>>63)

	fmt.Printf("4 << (^uintptr(0) >> 63)   output: %d\n", 4<<(^uintptr(0)>>63))
	fmt.Printf("4 >> 1  output: %d\n", 4>>1)
}
