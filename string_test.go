package basic

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
)

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

	aa := `{"gateway":"\u003c{manage_gateway}\u003e","ip":"\u003c{manage_ip}\u003e","ip_src":"static|dhcp","netmask":"\u003c{manage_netmask}\u003e"}`

	fmt.Println(strings.Replace(aa, "\\u003c{manage_gateway}\\u003e", "255.255.255.0", -1))

}

//Test Trim
func TestTrim(t *testing.T) {
	str := "\"\"aa233aa\""
	fmt.Println(strings.Trim(str, "aa"))
	fmt.Println(strings.Trim("aa,bb,cc,dd", ","))
	fmt.Println(strings.TrimLeft(str, "aa"))
	fmt.Println(strings.TrimRight(str, "aa"))
	fmt.Println(strings.Trim(str, "\""))

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
