package basic

import (
	"fmt"
	"testing"
)

// TestLoop_1: 正常并发测试输出
// TestLoop_2: 函数内部并发模式，输出异常
// TestLoop_3: 函数内部并发模式，添加 v:=v，重新给予指针信息
// TestLoop_4: range的时候函数内部是unsafe的

var keyMap = map[string]string{
	"A_key": "A_value",
	"B_key": "B_value",
	"C_key": "C_value",
	"D_key": "D_value",
	"E_key": "E_value",
	"F_key": "F_value",
}

//--- PASS: TestLoop_1 (0.00s)
//    --- PASS: TestLoop_1/B_key (0.00s)
//        test_test.go:21: key:B_key, value:B_value
//    --- PASS: TestLoop_1/C_key (0.00s)
//        test_test.go:21: key:C_key, value:C_value
//    --- PASS: TestLoop_1/D_key (0.00s)
//        test_test.go:21: key:D_key, value:D_value
//    --- PASS: TestLoop_1/E_key (0.00s)
//        test_test.go:21: key:E_key, value:E_value
//    --- PASS: TestLoop_1/F_key (0.00s)
//        test_test.go:21: key:F_key, value:F_value
//    --- PASS: TestLoop_1/A_key (0.00s)
//        test_test.go:21: key:A_key, value:A_value
func TestLoop_1(t *testing.T) {
	t.Parallel()
	for k, v := range keyMap {
		t.Run(k, func(t *testing.T) {
			t.Log(fmt.Sprintf("key:%s, value:%s", k, v))
		})
	}
}

//--- PASS: TestLoop_2 (0.00s)
//    --- PASS: TestLoop_2/A_key (0.00s)
//        test_test.go:44: key:F_key, value:F_value
//    --- PASS: TestLoop_2/D_key (0.00s)
//        test_test.go:44: key:F_key, value:F_value
//    --- PASS: TestLoop_2/C_key (0.00s)
//        test_test.go:44: key:F_key, value:F_value
//    --- PASS: TestLoop_2/B_key (0.00s)
//        test_test.go:44: key:F_key, value:F_value
//    --- PASS: TestLoop_2/E_key (0.00s)
//        test_test.go:44: key:F_key, value:F_value
//    --- PASS: TestLoop_2/F_key (0.00s)
//        test_test.go:44: key:F_key, value:F_value
func TestLoop_2(t *testing.T) {

	for k, v := range keyMap {
		t.Run(k, func(t *testing.T) {
			t.Parallel()
			t.Log(fmt.Sprintf("key:%s, value:%s", k, v))
		})
	}
}

//--- PASS: TestLoop_2 (0.00s)
//    --- PASS: TestLoop_2/D_key (0.00s)
//        test_test.go:58: key:C_key, value:D_value
//    --- PASS: TestLoop_2/C_key (0.00s)
//        test_test.go:58: key:C_key, value:C_value
//    --- PASS: TestLoop_2/A_key (0.00s)
//        test_test.go:58: key:C_key, value:A_value
//    --- PASS: TestLoop_2/B_key (0.00s)
//        test_test.go:58: key:C_key, value:B_value
//    --- PASS: TestLoop_2/F_key (0.00s)
//        test_test.go:58: key:C_key, value:F_value
//    --- PASS: TestLoop_2/E_key (0.00s)
//        test_test.go:58: key:C_key, value:E_value
func TestLoop_3(t *testing.T) {

	for k, v := range keyMap {
		t.Run(k, func(t *testing.T) {
			v := v
			t.Parallel()
			t.Log(fmt.Sprintf("key:%s, value:%s", k, v))
		})
	}
}

//key:F_key, value:F_value
//key:F_key, value:F_value
//key:F_key, value:F_value
//key:F_key, value:F_value
//key:F_key, value:F_value
//key:F_key, value:F_value
func TestLoop_4(t *testing.T) {
	var funcs []func()
	for k, v := range keyMap {
		funcs = append(funcs, func() {
			fmt.Println(fmt.Sprintf("key:%s, value:%s", k, v))
			//t.Log(fmt.Sprintf("key:%s, value:%s", k, v))
		})
	}

	for _, f := range funcs {
		f()
	}
}

// https://github.com/kyoh86/scopelint/blob/master/README.md
func TestLoop_5(t *testing.T) {
	var copies []*string
	for _, val := range keyMap {
		copies = append(copies, &val)
	}

	for _, item := range copies {
		fmt.Println(*item)
	}
}

// https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
func TestTLog(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1", value: 1},
		{name: "test 2", value: 2},
		{name: "test 3", value: 3},
		{name: "test 4", value: 4},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Here you test tc.value against a test function.
			// Let's use t.Log as our test function :-)
			t.Log(tc.value)
		})
	}
}
