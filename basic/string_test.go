package basic

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func RemoveFunc(s string, f func(rune) bool) string {
	rs := []rune(s)
	var tmp []rune
	for i := 0; i < len(rs); i++ {
		if f(rs[i]) {
			continue
		}
		tmp = append(tmp, rs[i])
	}
	return string(tmp)
}

func TestIntegerConversion(t *testing.T) {
	intI, _ := strconv.ParseInt("10", 0, 0)
	int8I, _ := strconv.ParseInt("10", 0, 8)
	int16I, _ := strconv.ParseInt("10", 0, 16)
	int32I, _ := strconv.ParseInt("10", 0, 32)
	int64I, _ := strconv.ParseInt("10", 0, 64)
	fmt.Printf("parseInt 0 result: value: %d, size: %d\n", intI, unsafe.Sizeof(int(intI)))
	fmt.Printf("parseInt 8 result: value: %d, size: %d\n", int8I, unsafe.Sizeof(int8(int8I)))
	fmt.Printf("parseInt 16 result: value: %d, size: %d\n", int16I, unsafe.Sizeof(int16(int16I)))
	fmt.Printf("parseInt 32 result: value: %d, size: %d\n", int32I, unsafe.Sizeof(int32(int32I)))
	fmt.Printf("parseInt 64 result: value: %d, size: %d\n", int64I, unsafe.Sizeof(int64(int64I)))
}

func TestSubstr(t *testing.T) {
	fmt.Println(Substr("hello", 0, 10))
}

func TestLikeStr(t *testing.T) {
	fmt.Println(fmt.Sprintf("%%%s%%", "hello"))
}

func TestLnTest(t *testing.T) {
	LnTest("pengganyu")
}

func TestUnicodeStr(t *testing.T) {
	EncodeStr()
}

func TestBase64Encode(t *testing.T) {
	//Base64Encode("hello")

	unicode, _ := base64.StdEncoding.DecodeString("aGVsbG93b3JsZAo=")
	fmt.Printf("decode str %s is %s\n ", "34134", string(unicode))
}

func TestStrState(t *testing.T) {
	type aa string
	fmt.Println(reflect.ValueOf(aa("hello")).String())
	fmt.Println(aa("hello"))
}

// 线程不安全
func TestBufferWrite(t *testing.T) {
	strs := []string{"1", "2", "3", "4", "5", "6"}

	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}

	fmt.Println(buffer.String())
}

// replace
func TestReplace(t *testing.T) {
	str := "%s,%s,%s"

	fmt.Println(strings.Replace(str, "%s", "1", 0))
	fmt.Println(strings.Replace(str, "%s", "1", 1))
	fmt.Println(strings.Replace(str, "%s", "1", -1))
	fmt.Println(strings.Replace("hello//world//g", "//", "/", -1))

	aa := `{"gateway":"\u003c{manage_gateway}\u003e","ip":"\u003c{manage_ip}\u003e","ip_src":"static|dhcp","netmask":"\u003c{manage_netmask}\u003e"}`

	fmt.Println(strings.Replace(aa, "\\u003c{manage_gateway}\\u003e", "255.255.255.0", -1))

}

func TestIndex(t *testing.T) {
	var s = "hello.tar.gz"
	fmt.Println(s[strings.Index(s, "."):])
	s = "10m20.434343s"
	fmt.Println(s[:strings.Index(s, ".")] + "s")

}

//Test Trim
func TestTrim(t *testing.T) {
	str := "\"\"aa233aa\""
	fmt.Printf("Trim : %s\n", strings.Trim(str, "aa"))
	fmt.Printf("TrimSuffix : %s\n", strings.TrimSuffix(str, "aa"))
	fmt.Printf("trim : %s\n", strings.Trim("aa,bb,cc,dd", ","))
	fmt.Printf("trim : %s\n", strings.TrimLeft("token=eyJhbGciOiJIUzI1NiJ9.", "token="))
	fmt.Printf("trim : %s\n", string("token=eyJhbGciOiJIUzI1NiJ9.")[6:])
	fmt.Printf("TrimLeft : %s\n", strings.TrimLeft(str, "aa"))
	fmt.Printf("TrimRight : %s\n", strings.TrimRight(str, "aa"))
	fmt.Printf("trim : %s\n", strings.Trim(str, "\""))
	fmt.Printf("trim : %s\n", strings.TrimSpace("          10.39.240.229"))

	// https://groups.google.com/forum/#!topic/golang-nuts/WAItFEvrhmU
	// TrimRight会截断掉给出的cutset所有的组合，以下TrimRigth都只输出A
	// TrimRight returns a slice of the string s, with all trailing Unicode code points     contained in cutset removed.

	// To remove a suffix, use TrimSuffix instead.
	fmt.Printf("TrimRigth: %s\n", strings.TrimRight("A06-09~06", "-09~06"))
	fmt.Printf("TrimRigth: %s\n", strings.TrimRight("A-9-09~06", "-09~06"))
	fmt.Printf("TrimRigth: %s\n", strings.TrimRight("A9-09~06", "-09~06"))
	fmt.Printf("TrimRigth: %s\n", strings.TrimRight("A~9-09~06", "-09~06"))
	fmt.Printf("TrimRigth: %s\n", strings.TrimRight("A69-09~06", "-09~06"))
	fmt.Printf("TrimSuffix: %s\n", strings.TrimSuffix("A06-09~06", "-09~06"))

	ss := "\"No trouble was found with this adapter.  However Error Log Analysis indicates that there recently may have been a network problem.  If this adapter is connected to a network, and if you you are experiencing problems with network communications, check for a loose or defective cable or connection.  If a switch or another system is directly attached to this adapter, verify it is powered up, configured, and functioning correctly.\""
	fmt.Println(strings.Trim(ss, "\\\""))

	bb := "localhost:8083"
	fmt.Println(strings.Trim(bb, "8083"))
}

// EventData 事件数据详情
type EventData struct {
	No    int    `json:"no"`
	Order string `json:"order"`
	Time  string `json:"time"`
	Info  string `json:"info"`
}

//TestIpmitool Ipmitool采集到的系统事件列表
func TestIpmitool(t *testing.T) {
	str := `   1 | 05/28/2014 | 16:35:20 | System ACPI Power State #0xff | D3 Power State | Asserted
   2 | 05/28/2014 | 17:31:33 | Button #0xff | Power Button pressed | Asserted
   3 | 05/28/2014 | 17:31:37 | System ACPI Power State #0xff | D0 Power State | Asserted
   4 | 05/29/2014 | 01:03:11 | Button #0xff | Power Button pressed | Asserted
   5 | 05/29/2014 | 01:03:13 | System ACPI Power State #0xff | D3 Power State | Asserted
   6 | 05/29/2014 | 01:55:13 | System ACPI Power State #0xff | D3 Power State | Asserted
   7 | 05/29/2014 | 01:55:24 | Button #0xff | Power Button pressed | Asserted
   8 | 05/29/2014 | 01:55:28 | System ACPI Power State #0xff | D0 Power State | Asserted
   9 | 05/29/2014 | 02:02:33 | Button #0xff | Power Button pressed | Asserted
   a | 05/29/2014 | 02:02:35 | System ACPI Power State #0xff | D3 Power State | Asserted
   b | 05/29/2014 | 03:04:34 | Button #0xff | Power Button pressed | Asserted
   c | 05/29/2014 | 03:04:36 | System ACPI Power State #0xff | D3 Power State | Asserted
   d | 08/28/2016 | 14:16:33 | Button #0xff | Power Button pressed | Asserted
   e | 08/28/2016 | 14:16:38 | System ACPI Power State #0xff | D3 Power State | Asserted
   f | 08/28/2016 | 14:32:48 | Button #0xff | Power Button pressed | Asserted
 
`

	var tools []*EventData
	for i, str := range strings.Split(str, "\n") {

		if strings.Count(str, "|") <= 0 {
			continue
		}

		strs := strings.Split(str, "|")

		order := strings.TrimSpace(strs[0])
		time := strings.TrimSpace(strs[1] + strs[2])
		info := strings.TrimSpace(strs[3] + strs[4] + strs[5])

		tools = append(tools, &EventData{
			No:    i,
			Order: order,
			Time:  time,
			Info:  info,
		})
	}

	sort.Slice(tools, func(i, j int) bool {
		return tools[i].Time < tools[j].Time
	})

	for _, item := range tools {
		fmt.Printf("order: %s, time: %s, info: %s\n", item.Order, item.Time, item.Info)
	}

}

//TestCompare 字符串匹配的时候可以直接对比
func TestCompare(t *testing.T) {
	fmt.Println("12/20/2018 15:35:50" < "01/15/2019 07:45:27")
}

func TestNullSlice(t *testing.T) {
	//sl := make([]string, 0)
	//sl := []string{""}
	sl := []string{}

	for i := range sl {
		fmt.Printf("hello,%d\n", i)
	}
}

//getField 获取field value
func TestgetField(t *testing.T) string {
	//1 | 07/22/2016 | 15:41:43 | Event Logging Disabled SEL | Log area reset/cleared | Asserted
	var str, char string
	var index int
	if index == 1 {
		return str[0:strings.Index(str, char)]
	}

	if index == 2 {

	}
	return ""
}

func TestSlice(t *testing.T) {
	str := "hello?223"

	fmt.Println(str[0:strings.Index(str, "?")])
}

func TestTimeStrCompare(t *testing.T) {
	fmt.Println(time.Now().String())
	fmt.Println(time.Now().String() > "2020-10-11")
	fmt.Println(time.Now().String() > "0000-00-00")
}

func TestGetPasswd(t *testing.T) {
	str := "ipmitool -I lanplus -H 10.0.10.100 -U huxj -P oW3g^ZZo user set password 3 C2HT~rmb"

	strs := strings.Split(str, " ")
	for i, v := range strs {
		if strings.EqualFold(strings.ToUpper(v), "-P") {
			str = strings.Replace(str, strs[i+1], "******", -1)
		}
		if strings.EqualFold(v, "password") && strings.Contains(str, "set password") {
			str = strings.Replace(str, strs[i+2], "******", -1)
		}
	}
	fmt.Println(str)
}

func TestStr(t *testing.T) {
	str := "123"
	//str[0] = 'a' 虽然string是数组，但是不能给它的index进行赋值
	//([]byte)(str)[1] = "b"  同样的道理
	fmt.Printf("%v", str)
}

func TestDecodeRuby(t *testing.T) {
	rubyBytes := []byte{4, 8, 123, 15, 58, 9, 98, 111, 100, 121, 34, 2, 10, 1, 4, 8, 123, 9, 58, 10, 97, 103, 101, 110, 116, 73, 34, 10, 115, 104, 101, 108, 108, 6, 58, 6, 69, 84, 58, 11, 97, 99, 116, 105, 111, 110, 73, 34, 8, 114, 117, 110, 6, 59, 6, 84, 58, 11, 99, 97, 108, 108, 101, 114, 73, 34, 15, 99, 108, 111, 117, 100, 45, 97, 99, 116, 50, 6, 59, 6, 84, 58, 9, 100, 97, 116, 97, 123, 15, 58, 9, 116, 121, 112, 101, 73, 34, 11, 115, 99, 114, 105, 112, 116, 6, 59, 6, 84, 58, 9, 117, 115, 101, 114, 73, 34, 9, 114, 111, 111, 116, 6, 59, 6, 84, 58, 12, 99, 111, 109, 109, 97, 110, 100, 73, 34, 16, 115, 99, 114, 105, 112, 116, 58, 47, 116, 109, 112, 6, 59, 6, 84, 58, 13, 102, 105, 108, 101, 110, 97, 109, 101, 73, 34, 8, 116, 109, 112, 6, 59, 6, 84, 58, 12, 99, 111, 110, 116, 101, 110, 116, 73, 34, 13, 90, 87, 78, 111, 98, 119, 61, 61, 6, 59, 6, 84, 58, 11, 98, 97, 115, 101, 54, 52, 84, 58, 11, 112, 97, 114, 97, 109, 115, 73, 34, 0, 6, 59, 6, 84, 58, 15, 115, 99, 114, 105, 112, 116, 84, 121, 112, 101, 73, 34, 9, 66, 97, 115, 104, 6, 59, 6, 84, 58, 19, 112, 114, 111, 99, 101, 115, 115, 95, 114, 101, 115, 117, 108, 116, 84, 58, 16, 101, 110, 118, 105, 114, 111, 110, 109, 101, 110, 116, 73, 34, 0, 6, 59, 6, 84, 58, 13, 115, 101, 110, 100, 101, 114, 105, 100, 73, 34, 15, 99, 108, 111, 117, 100, 45, 97, 99, 116, 50, 6, 58, 6, 69, 84, 58, 14, 114, 101, 113, 117, 101, 115, 116, 105, 100, 73, 34, 46, 97, 99, 116, 50, 45, 102, 99, 53, 100, 56, 57, 54, 54, 45, 102, 54, 53, 53, 45, 56, 51, 101, 102, 45, 97, 53, 56, 51, 45, 56, 51, 97, 53, 50, 48, 99, 48, 97, 101, 48, 102, 6, 59, 7, 84, 58, 11, 102, 105, 108, 116, 101, 114, 123, 7, 73, 34, 10, 97, 103, 101, 110, 116, 6, 59, 7, 84, 91, 6, 73, 34, 10, 115, 104, 101, 108, 108, 6, 59, 7, 84, 58, 15, 99, 111, 108, 108, 101, 99, 116, 105, 118, 101, 73, 34, 16, 109, 99, 111, 108, 108, 101, 99, 116, 105, 118, 101, 6, 59, 7, 84, 59, 10, 73, 34, 16, 109, 99, 111, 108, 108, 101, 99, 116, 105, 118, 101, 6, 59, 7, 84, 58, 10, 97, 103, 101, 110, 116, 73, 34, 10, 115, 104, 101, 108, 108, 6, 59, 7, 84, 58, 13, 99, 97, 108, 108, 101, 114, 105, 100, 73, 34, 20, 99, 101, 114, 116, 61, 99, 108, 111, 117, 100, 45, 97, 99, 116, 50, 6, 59, 7, 84, 58, 8, 116, 116, 108, 105, 2, 16, 14, 58, 12, 109, 115, 103, 116, 105, 109, 101, 108, 43, 7, 101, 103, 111, 93, 58, 9, 104, 97, 115, 104, 73, 34, 37, 49, 57, 57, 51, 101, 98, 98, 52, 56, 97, 53, 101, 49, 49, 98, 97, 57, 102, 100, 48, 57, 56, 49, 53, 99, 102, 102, 101, 100, 55, 49, 102, 6, 59, 7, 70}
	fmt.Printf("%s\n", string(rubyBytes))
}

// TODO 写一个提取字符串的工具
func TestFieldsFunc(t *testing.T) {
	dsn := "root:Yunjikeji#123@tcp(10.0.2.8:3306)/cloudboot_3.0.0?charset=utf8&parseTime=True&loc=Local"

	schema := "$:$@tcp($)/$?charset=utf8&parseTime=True&loc=Local"

	var result []string
	var start int
	for _, it := range strings.Split(schema, "$") {
		if it == "" {
			continue
		}
		itIndex := strings.Index(dsn, it)
		result = append(result, dsn[start:itIndex])
		start = itIndex + len(it)
	}

	fmt.Println(result)

	userIndex := strings.Index(dsn, ":")
	passwordIndex := strings.Index(dsn, "@tcp(")
	hostportIndex := strings.Index(dsn, ")/")
	dbIndex := strings.Index(dsn, "?charset")

	fmt.Println(userIndex)
	fmt.Println(passwordIndex)
	fmt.Println(hostportIndex)
	fmt.Println(dbIndex)

	fmt.Println(dsn[:userIndex])
	fmt.Println(dsn[userIndex+1 : passwordIndex])
	fmt.Println(dsn[passwordIndex+5 : hostportIndex])
	fmt.Println(dsn[hostportIndex+2 : dbIndex])

	//for _, s := range strs {
	//	fmt.Println(s)
	//}
}

func TestFieldsSpace(t *testing.T) {
	str := "root:  Yunjikeji#123@tcp(10.0.2.8:33  06)/cloudboot_3.0.0?charset=utf8&parseTime=True&loc=Local"
	//
	//strs := strings.FieldsFunc(str, func(r rune) bool {
	//	return unicode.IsSpace(r)
	//})

	strs := strings.Fields(str)

	for _, s := range strs {
		fmt.Println(s)
	}
}

func TestReplaceAll(t *testing.T) {
	//origin := []byte("hello_world_cloudboot")
	//numberSequence := regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
	//dest:=numberSequence.ReplaceAll(origin, []byte(`$3 $2 $1`))
	//fmt.Println(string(dest))

	reg := regexp.MustCompile(`(\w+)\s(\w+)\s(\w+)`)
	str := "hello world cloudboot"
	dest := reg.ReplaceAll([]byte(str), []byte(`$3 $2 $1`))
	// output: cloudboot world hello
	fmt.Println(string(dest))

	dest = reg.ReplaceAll([]byte(str), []byte(`$1 $2 $3`))
	fmt.Println(reg.NumSubexp())
	fmt.Println(reg.SubexpNames())
	fmt.Println(string(dest))

}

func TestEq(t *testing.T) {
	var ss []string
	var s string
	ss = nil
	println(ss)
	println(s)

	println(strings.Join(ss, ","))
	println(strings.Join(ss, ",") == s)
}

func TestField(t *testing.T) {
	for _, item := range strings.Fields("hello world") {
		fmt.Println(item)
	}
}

func TestTrimFunc(t *testing.T) {
	a := []byte("hello world")
	binaryStr := fmt.Sprintf("%8b", a)

	fmt.Println(strings.Replace(binaryStr, " ", "", -1))

	// rune与byte的区别是，rune是四个字节，而byte是一个字节
	// rune是unicode编码
	str := strings.TrimFunc(binaryStr, func(r rune) bool {
		return r == 91 || r == 93
	})

	fmt.Println(str)
}

func TestRemoveFunc(t *testing.T) {
	a := []byte("hello world")
	binaryStr := fmt.Sprintf("%8b", a)

	str := RemoveFunc(binaryStr, func(r rune) bool {
		return r == 91 || r == 93 || r == 32
	})

	fmt.Println(str)
}

// BenchmarkEmptyString-4   	2000000000	         0.42 ns/op  空string测试结果
// BenchmarkEmptyString-4   	2000000000	         0.36 ns/op  uuid测试结果
func BenchmarkEmptyString(t *testing.B) {
	ss := "2BB42903-8127-4A6B-A034-B0252FAE97E5"
	for i := 0; i < t.N; i++ {
		_ = len(ss) <= 0
	}
}

// BenchmarkEmptyString2-4   	2000000000	         0.84 ns/op  空string测试结果
// BenchmarkEmptyString2-4   	2000000000	         0.35 ns/op  uuid测试结果
func BenchmarkEmptyString2(t *testing.B) {
	ss := "2BB42903-8127-4A6B-A034-B0252FAE97E5"
	for i := 0; i < t.N; i++ {
		_ = ss == ""
	}
}

// BenchmarkIntToStr-4   	10000000	       132 ns/op
func BenchmarkIntToStr(t *testing.B) {
	var a = 10
	for i := 0; i < t.N; i++ {
		fmt.Sprintf("%d", a)
	}
}

// BenchmarkIntToStr2-4   	300000000	         4.51 ns/op
func BenchmarkIntToStr2(t *testing.B) {
	var a = 10
	for i := 0; i < t.N; i++ {
		strconv.Itoa(a)
	}
}
